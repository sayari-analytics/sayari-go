package sdk

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	sayari "github.com/sayari-analytics/sayari-go/generated/go"
)

var ErrNoMatchFound = errors.New("no match found for this entity")

const (
	name        = "name"
	identifier  = "identifier" //TODO: support column headers that match identifier types, TBD
	country     = "country"
	address     = "address"
	dateOfBirth = "dateofbirth"
	contact     = "contact"
	entityType  = "type"
)

// TODO: add better explanations for these columns
var attributeFieldsMap = map[string]string{
	name:        "The name of the entity",
	identifier:  "...",
	country:     "Must be from the enum set",
	address:     "...",
	dateOfBirth: "...",
	contact:     "...",
	entityType:  "Must be from the enum set",
}

func screenCSV(ctx context.Context, c *Connection, attributeColMap map[string][]int, csvDataChan chan []string,
	summaryChan chan sayari.EntitySummaryResponse, unresolvedChan chan []string, wg *sync.WaitGroup, errChan chan error) {
	for row := range csvDataChan {
		// Attempt to resolve each entity
		entityID, err := resolveEntity(ctx, c, attributeColMap, row)
		if err != nil {
			// return if we have an error more serious than being unable to resolve
			if err != ErrNoMatchFound {
				errChan <- err
				return
			}
			// track unresolved rows
			unresolvedChan <- row
			continue
		}
		// Get risks
		entitySummary, err := c.Entity.EntitySummary(ctx, entityID)
		if err != nil {
			log.Printf("Error fetching entity '%v' ID '%v'", row[attributeColMap[name][0]], entityID)
			errChan <- err
			return
		}
		summaryChan <- *entitySummary
	}
	wg.Done()
}

func (c *Connection) ScreenCSVEntities(ctx context.Context, csvPath string) ([]*sayari.EntitySummaryResponse, []*sayari.EntitySummaryResponse, [][]string, error) {
	var riskyEntities []*sayari.EntitySummaryResponse
	var nonRiskyEntities []*sayari.EntitySummaryResponse
	var unresolved [][]string

	// Load in CSV
	rows, err := loadCSV(csvPath)
	if err != nil {
		return nil, nil, nil, err
	}

	// create a map of the columns to be able to look up which contains which attribute
	attributeColMap := make(map[string][]int)

	// create channels to handle this work concurrently
	// concurrency reduced from 3 -> 1 to cope with rate limiting
	numWorkers := 1
	csvDataChan := make(chan []string, numWorkers)
	summaryChan := make(chan sayari.EntitySummaryResponse, numWorkers)
	unresolvedChan := make(chan []string, numWorkers)
	errChan := make(chan error)
	var wg = sync.WaitGroup{}

	// start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go screenCSV(ctx, c, attributeColMap, csvDataChan, summaryChan, unresolvedChan, &wg, errChan)
	}

	// send data to workers
	go func() {
		for i, row := range rows {
			// Use the first row to map the columns
			if i == 0 {
				err = mapCSV(row, attributeColMap)
				if err != nil {
					errChan <- err
				}
				continue
			}

			// Send row to workers
			csvDataChan <- row
		}
		close(csvDataChan)
	}()

	// signal when all the workers are done
	done := make(chan bool)
	go func() {
		wg.Wait()
		done <- true
	}()

	// read results off the channels and return when done
	for {
		select {
		case summary := <-summaryChan:
			if len(summary.Risk) > 0 {
				riskyEntities = append(riskyEntities, &summary)
			} else {
				nonRiskyEntities = append(nonRiskyEntities, &summary)
			}
		case unresolvedRow := <-unresolvedChan:
			unresolved = append(unresolved, unresolvedRow)
		case err = <-errChan:
			return nil, nil, nil, err
		case <-done:
			return riskyEntities, nonRiskyEntities, unresolved, nil
		}
	}
}

func loadCSV(csvPath string) ([][]string, error) {
	data, err := os.Open(csvPath)
	if err != nil {
		return nil, err
	}
	defer data.Close()

	reader := csv.NewReader(data)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return rows, nil
}

// determine which fields are in which columns
func mapCSV(row []string, attributeColMap map[string][]int) error {
	for colNum, colName := range row {
		// Remove all spaces from the column name and convert to lowercase
		colName = strings.ToLower(strings.ReplaceAll(colName, " ", ""))
		// If the column is valid, note its position(s)
		if _, ok := attributeFieldsMap[colName]; !ok {
			return fmt.Errorf("column '%v' does not match to resolution field", colName)
		}
		if existingColNums, ok := attributeColMap[colName]; ok {
			attributeColMap[colName] = append(existingColNums, colNum)
		} else {
			attributeColMap[colName] = []int{colNum}
		}
	}
	return nil
}

func resolveEntity(ctx context.Context, c *Connection, attributeColMap map[string][]int, row []string) (string, error) {
	var entityInfo sayari.Resolution

	if colNums, ok := attributeColMap[name]; ok {
		for _, colNum := range colNums {
			entityInfo.Name = append(entityInfo.Name, &row[colNum])
		}
	}

	if colNums, ok := attributeColMap[identifier]; ok {
		for _, colNum := range colNums {
			var bothIdentifier *sayari.BothIdentifierTypes
			providedIdentifier := strings.ToLower(row[colNum])
			// attempt to get Identifier
			thisIdentifier, err := sayari.NewIdentifierTypeFromString(providedIdentifier)
			if err != nil {
				bothIdentifier = sayari.NewBothIdentifierTypesFromIdentifierType(thisIdentifier)
			}

			// if that didn't work, attempt to get weak identifier
			thisWeakIdentifier, err := sayari.NewWeakIdentifierTypeFromString(providedIdentifier)
			if err != nil {
				bothIdentifier = sayari.NewBothIdentifierTypesFromWeakIdentifierType(thisWeakIdentifier)
			}

			if bothIdentifier == nil {
				log.Printf("Failed to resolve entity. '%v' is not a valid strong or weak identifier type.", row[colNum])
				return "", fmt.Errorf("'%v' is not a valid strong or weak identifier type", row[colNum])
			}

			entityInfo.Identifier = append(entityInfo.Identifier, bothIdentifier)
		}
	}

	if colNums, ok := attributeColMap[country]; ok {
		for _, colNum := range colNums {
			country, err := sayari.NewCountryFromString(row[colNum])
			if err != nil {
				return "", err
			}
			entityInfo.Country = append(entityInfo.Country, &country)
		}
	}

	if colNums, ok := attributeColMap[address]; ok {
		for _, colNum := range colNums {
			entityInfo.Address = append(entityInfo.Address, &row[colNum])
		}
	}

	if colNums, ok := attributeColMap[dateOfBirth]; ok {
		for _, colNum := range colNums {
			entityInfo.DateOfBirth = append(entityInfo.DateOfBirth, &row[colNum])
		}
	}

	if colNums, ok := attributeColMap[contact]; ok {
		for _, colNum := range colNums {
			entityInfo.Contact = append(entityInfo.Contact, &row[colNum])
		}
	}

	if colNums, ok := attributeColMap[entityType]; ok {
		for _, colNum := range colNums {
			entityType, err := sayari.NewEntitiesFromString(row[colNum])
			if err != nil {
				return "", err
			}
			entityInfo.Type = append(entityInfo.Type, &entityType)
		}
	}

	entity, err := c.Resolution.Resolution(ctx, &entityInfo)
	if err != nil {
		return "", err
	}

	if len(entity.Data) == 0 {
		return "", ErrNoMatchFound
	}
	return entity.Data[0].EntityId, nil
}
