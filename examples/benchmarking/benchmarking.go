package main

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	sayari "github.com/sayari-analytics/sayari-go/generated/go"
	"github.com/sayari-analytics/sayari-go/sdk"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

const (
	name        = "name"
	identifier  = "identifier"
	country     = "country"
	address     = "address"
	dateOfBirth = "dateofbirth"
	contact     = "contact"
	entityType  = "type"
)

var attributeFieldsMap = map[string]string{
	name:        "The name of the entity",
	identifier:  "...",
	country:     "Must be from the enum set",
	address:     "...",
	dateOfBirth: "...",
	contact:     "...",
	entityType:  "Must be from the enum set",
}

func main() {
	// load ENV file if ENV vars are not set
	if os.Getenv("CLIENT_ID") == "" || os.Getenv("CLIENT_SECRET") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Failed to load .env file. Err: %v", err)
		}
	}

	// Create a client to auth against the API
	client, err := sdk.Connect(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Load in CSV
	rows, err := loadCSV("examples/benchmarking/input.csv")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Open output CSV
	file, err := os.Create("examples/benchmarking/output.csv")
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(file)
	defer w.Flush()

	// create a map of the columns to be able to look up which contains which attribute
	attributeColMap := make(map[string][]int)

	// Map CSV
	err = mapCSV(rows[0], attributeColMap)

	headers := []string{"t1", "t2", "t3", "r1", "r2", "r3"}
	w.Write(headers)

	// Process each row
	for i, row := range rows {
		// skip first row
		if i == 0 {
			log.Println("Headers: ", row)
			continue
		}

		// Resolve corporate profile
		t1, resp1, err := resolveEntity(client, "corporate", attributeColMap, row)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		// Resolve supplies profile
		t2, resp2, err := resolveEntity(client, "suppliers", attributeColMap, row)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		// Search
		t3, resp3, err := searchEntity(client, attributeColMap, row)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		resultsRow := []string{
			fmt.Sprint(t1.Seconds()),
			fmt.Sprint(t2.Seconds()),
			fmt.Sprint(t3.Seconds()),
			resp1.String(),
			resp2.String(),
			resp3.String(),
		}
		err = w.Write(resultsRow)
		if err != nil {
			log.Fatalf("Error: %v", err)
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

var ErrNoMatchFound = errors.New("no match found for this entity")

func resolveEntity(client *sdk.Connection, profile string, attributeColMap map[string][]int, row []string) (time.Duration, *sayari.ResolutionResponse, error) {
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
				log.Fatalf("Failed to resolve entity. '%v' is not a valid strong or weak identifier type.", row[colNum])
			}

			entityInfo.Identifier = append(entityInfo.Identifier, bothIdentifier)
		}
	}

	if colNums, ok := attributeColMap[country]; ok {
		for _, colNum := range colNums {
			country, err := sayari.NewCountryFromString(row[colNum])
			if err != nil {
				log.Fatalf("Error: %v", err)
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
				log.Fatalf("Error: %v", err)
			}
			entityInfo.Type = append(entityInfo.Type, &entityType)
		}
	}

	start := time.Now()
	resp, err := client.Resolution.Resolution(context.Background(), &entityInfo)
	duration := time.Since(start)
	if err != nil {
		return duration, nil, err
	}

	if len(resp.Data) == 0 {
		return duration, nil, ErrNoMatchFound
	}
	return duration, resp, nil
}

func searchEntity(client *sdk.Connection, attributeColMap map[string][]int, row []string) (time.Duration, *sayari.EntitySearchResponse, error) {
	var entityInfo sayari.SearchEntity

	entityInfo.Q = row[attributeColMap[name][0]]
	entityInfo.Fields = []sayari.SearchField{sayari.SearchFieldName}
	filterList := sayari.FilterList{}
	filterList.EntityType = []sayari.Entities{sayari.EntitiesCompany}

	if colNums, ok := attributeColMap[country]; ok {
		for _, colNum := range colNums {
			country, err := sayari.NewCountryFromString(row[colNum])
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
			filterList.Country = append(filterList.Country, country)
		}
	}
	entityInfo.Filter = &filterList

	start := time.Now()
	resp, err := client.Search.SearchEntity(context.Background(), &entityInfo)
	duration := time.Since(start)
	if err != nil {
		return duration, nil, err
	}

	if len(resp.Data) == 0 {
		return duration, nil, ErrNoMatchFound
	}
	return duration, resp, nil
}
