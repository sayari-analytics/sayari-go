package main

import (
	"context"
	"encoding/csv"
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

	headers := []string{
		"field_name", "field_address", "field_country", "field_type",
		"corporate_entity_id", "supplier_entity_id", "search_entity_id",
		"corporate_name", "supplier_name", "search_name",
		"corporate_address", "supplier_address", "search_address",
		"corporate_country", "supplier_country", "search_country",
		"corporate_type", "supplier_type", "search_type",
		"corporate_strength", "supplier_strength",
	}
	w.Write(headers)

	// Process each row
	for i, row := range rows {
		// skip first row
		if i == 0 {
			log.Println("Headers: ", row)
			continue
		}

		fieldValues := getFieldInfo(attributeColMap, row)

		// Resolve corporate profile
		_, resp1, err := resolveEntity(client, sayari.ProfileEnumCorporate, attributeColMap, row)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		r1 := getResolveData(resp1)

		// Resolve supplies profile
		_, resp2, err := resolveEntity(client, sayari.ProfileEnumSupplier, attributeColMap, row)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		r2 := getResolveData(resp2)

		// Search
		_, resp3, err := searchEntity(client, attributeColMap, row)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		r3 := getSearchData(resp3)

		results := []string{
			fieldValues[0], fieldValues[1], fieldValues[2], fieldValues[3], // fieldName, fieldAddress, fieldCountry, fieldType
			r1[0], r2[0], r3[0], // ID
			r1[1], r2[1], r3[1], // Name
			r1[2], r2[2], r3[2], // Address
			r1[3], r2[3], r3[3], // Country
			r1[4], r2[4], r3[4], // Type
			r1[5], r2[5], // match strength
		}

		err = w.Write(results)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}
}

func getFieldInfo(attributeFieldsMap map[string][]int, row []string) []string {
	fieldName := row[attributeFieldsMap[name][0]]
	fieldAddress := row[attributeFieldsMap[address][0]]
	fieldCountry := row[attributeFieldsMap[country][0]]

	var fieldType string
	if len(attributeFieldsMap[entityType]) > 0 {
		fieldType = row[attributeFieldsMap[entityType][0]]
	}
	return []string{fieldName, fieldAddress, fieldCountry, fieldType}
}

func getResolveData(resp *sayari.ResolutionResponse) []string {
	if len(resp.Data) == 0 {
		emptyResp := make([]string, 6)
		return emptyResp
	}
	e1 := resp.Data[0]
	var e1addr string
	if len(e1.Addresses) > 0 {
		e1addr = e1.Addresses[0]
	}
	var e1country string
	if len(e1.Countries) > 0 {
		e1country = fmt.Sprint(e1.Countries[0])
	}
	return []string{e1.EntityId, e1.Label, e1addr, e1country, fmt.Sprint(e1.Type), e1.MatchStrength.Value}
}

func getSearchData(resp *sayari.EntitySearchResponse) []string {
	if len(resp.Data) == 0 {
		emptyResp := make([]string, 5)
		return emptyResp
	}
	e1 := resp.Data[0]
	var e1addr string
	if len(e1.Addresses) > 0 {
		e1addr = e1.Addresses[0]
	}
	var e1country string
	if len(e1.Countries) > 0 {
		e1country = fmt.Sprint(e1.Countries[0])
	}
	return []string{e1.Id, e1.Label, e1addr, e1country, fmt.Sprint(e1.Type)}

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

func resolveEntity(client *sdk.Connection, profile sayari.ProfileEnum, attributeColMap map[string][]int, row []string) (time.Duration, *sayari.ResolutionResponse, error) {
	var entityInfo sayari.Resolution
	entityInfo.Profile = &profile

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

	return duration, resp, nil
}
