package sdk

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strings"

	sayari "github.com/sayari-analytics/sayari-go/generated/go"
)

var ErrNoMatchFound = errors.New("no match found for this entity")

const (
	Name        = "name"
	Identifier  = "identifier" //TODO: support column headers that match identifier types, TBD
	Country     = "country"
	Address     = "address"
	DateOfBirth = "dateofbirth"
	Contact     = "contact"
	Type        = "type"
)

// TODO: add better explanations for these columns
var attributeFieldsMap = map[string]string{
	Name:        "The name of the entity",
	Identifier:  "...",
	Country:     "Must be from the enum set",
	Address:     "...",
	DateOfBirth: "...",
	Contact:     "...",
	Type:        "Must be from the enum set",
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

	// Read through the rows of data
	for i, row := range rows {
		// Use the first row to map the columns
		if i == 0 {
			err = mapCSV(row, attributeColMap)
			if err != nil {
				return nil, nil, nil, err
			}
			continue
		}

		//TODO: Introduce concurrency

		// Attempt to resolve each entity
		entityID, err := resolveEntity(ctx, c, attributeColMap, row)
		if err != nil {
			// return if we have an error more serious than being unable to resolve
			if err != ErrNoMatchFound {
				return nil, nil, nil, err
			}
			// track unresolved rows
			unresolved = append(unresolved, row)
			continue
		}

		// Get risks
		entitySummary, err := c.Entity.EntitySummary(ctx, entityID)
		if err != nil {
			return nil, nil, nil, err
		}

		if len(entitySummary.Risk) > 0 {
			riskyEntities = append(riskyEntities, entitySummary)
		} else {
			nonRiskyEntities = append(nonRiskyEntities, entitySummary)
		}
	}
	return riskyEntities, nonRiskyEntities, unresolved, err
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
		if _, ok := attributeFieldsMap[colName]; ok {
			if existingColNums, ok := attributeColMap[colName]; ok {
				attributeColMap[colName] = append(existingColNums, colNum)
			} else {
				attributeColMap[colName] = []int{colNum}
			}
		} else {
			return fmt.Errorf("column '%v' does not match to resolution field", colName)
		}
	}
	return nil
}

func resolveEntity(ctx context.Context, c *Connection, attributeColMap map[string][]int, row []string) (string, error) {
	var entityInfo sayari.Resolution

	if colNums, ok := attributeColMap[Name]; ok {
		for _, colNum := range colNums {
			entityInfo.Name = append(entityInfo.Name, &row[colNum])
		}
	}

	if colNums, ok := attributeColMap[Identifier]; ok {
		for _, colNum := range colNums {
			entityInfo.Identifier = append(entityInfo.Identifier, &row[colNum])
		}
	}

	if colNums, ok := attributeColMap[Country]; ok {
		for _, colNum := range colNums {
			country, err := sayari.NewCountryFromString(row[colNum])
			if err != nil {
				return "", err
			}
			entityInfo.Country = append(entityInfo.Country, &country)
		}
	}

	if colNums, ok := attributeColMap[Address]; ok {
		for _, colNum := range colNums {
			entityInfo.Address = append(entityInfo.Address, &row[colNum])
		}
	}

	if colNums, ok := attributeColMap[DateOfBirth]; ok {
		for _, colNum := range colNums {
			entityInfo.DateOfBirth = append(entityInfo.DateOfBirth, &row[colNum])
		}
	}

	if colNums, ok := attributeColMap[Contact]; ok {
		for _, colNum := range colNums {
			entityInfo.Contact = append(entityInfo.Contact, &row[colNum])
		}
	}

	if colNums, ok := attributeColMap[Type]; ok {
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
