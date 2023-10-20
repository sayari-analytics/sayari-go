package sdk

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
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
	summaryChan chan sayari.EntityDetails, unresolvedChan chan []string, wg *sync.WaitGroup, errChan chan error) {
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
			errChan <- err
			return
		}
		summaryChan <- *entitySummary
	}
	wg.Done()
}

func (c *Connection) ScreenCSVEntities(ctx context.Context, csvPath string) ([]*sayari.EntityDetails, []*sayari.EntityDetails, [][]string, error) {
	var riskyEntities []*sayari.EntityDetails
	var nonRiskyEntities []*sayari.EntityDetails
	var unresolved [][]string

	// Load in CSV
	rows, err := loadCSV(csvPath)
	if err != nil {
		return nil, nil, nil, err
	}

	// create a map of the columns to be able to look up which contains which attribute
	attributeColMap := make(map[string][]int)

	// create channels to handle this work concurrently
	numWorkers := 3
	csvDataChan := make(chan []string, numWorkers)
	summaryChan := make(chan sayari.EntityDetails, numWorkers)
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

	// close all the channels once we are done
	var done bool
	go func() {
		wg.Wait()
		close(summaryChan)
		close(unresolvedChan)
		close(errChan)
		done = true
	}()

	// read results off the channels
	for !done {
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
		}
	}

	return riskyEntities, nonRiskyEntities, unresolved, nil
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
		// If the column is valid, note its position(s)
		colName = strings.ToLower(colName)
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
			entityInfo.Identifier = append(entityInfo.Identifier, &row[colNum])
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
