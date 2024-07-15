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

	arg "github.com/alexflint/go-arg"
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

type result struct {
	entityID      string
	name          string
	address       string
	entityCountry string
	entityType    string
	matchStrength string
}

var attributeFieldsMap = map[string]string{
	name:        "The name of the entity",
	identifier:  "...",
	country:     "Must be from the enum set",
	address:     "...",
	dateOfBirth: "...",
	contact:     "...",
	entityType:  "Must be from the enum set",
}

var args struct {
	MaxResults         int  // the maximum number of results to return for each search (defaults to 3)
	MeasureSupplyChain bool // set to true if you want to include supply chain metrics
}

type supplyChainInfo struct {
	hasSupplyChain    bool
	numSuppliers      int
	avgSupplyChainLen float64
}

var supplyChainCache map[string]supplyChainInfo

func main() {
	// load ENV file if ENV vars are not set
	if os.Getenv("CLIENT_ID") == "" || os.Getenv("CLIENT_SECRET") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Failed to load .env file. Err: %v", err)
		}
	}

	arg.MustParse(&args)
	if args.MaxResults == 0 {
		log.Println("MaxResults not set, will use default of 3")
		args.MaxResults = 3
	} else {
		log.Printf("MaxResults set to %v", args.MaxResults)
	}

	// Use the base URL ENV var if provided
	baseURL := sayari.Environments.Production
	if os.Getenv("BASE_URL") != "" {
		baseURL = os.Getenv("BASE_URL")
	}

	// Create a client to auth against the API
	client, err := sdk.ConnectTo(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), baseURL)
	if err != nil {
		log.Fatalf("Error creating client. Error: %v", err)
	}

	// Load in CSV
	rows, err := loadCSV("examples/benchmarking/input.csv")
	if err != nil {
		log.Fatalf("Error loading CSV. Error: %v", err)
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
		"field_name", "field_address", "field_country", "field_identifier", "field_type", "result index",
		"corporate_entity_id", "supplier_entity_id", "search_entity_id",
		"corporate_name", "supplier_name", "search_name",
		"corporate_address", "supplier_address", "search_address",
		"corporate_country", "supplier_country", "search_country",
		"corporate_type", "supplier_type", "search_type",
		"corporate_strength", "supplier_strength",
	}
	if args.MeasureSupplyChain {
		headers = append(headers,
			"corporate_supply_chain", "corporate_suppliers_count", "corporate_avg_supply_chain_len",
			"supplier_supply_chain", "supplier_suppliers_count", "supplier_avg_supply_chain_len",
			"search_supply_chain", "search_suppliers_count", "search_avg_supply_chain_len",
		)
		// initialize cache
		supplyChainCache = make(map[string]supplyChainInfo)
	}
	w.Write(headers)

	// Process each row
	for i, row := range rows {
		log.Printf("Processing line %v of %v", i+1, len(rows))
		// skip first row
		if i == 0 {
			log.Println("Input Headers: ", row)
			continue
		}

		fieldValues := getFieldInfo(attributeColMap, row)

		// Resolve corporate profile
		_, resp1, err := resolveEntity(client, sayari.ProfileEnumCorporate, attributeColMap, row)
		if err != nil {
			log.Fatalf("Error resolving entity w/ corporate profile. Error: %v", err)
		}
		r1 := getResolveData(resp1)

		// Resolve supplies profile
		_, resp2, err := resolveEntity(client, sayari.ProfileEnumSuppliers, attributeColMap, row)
		if err != nil {
			log.Fatalf("Error resolving entity w/ suppliers profile. Error: %v", err)
		}
		r2 := getResolveData(resp2)

		// Search
		_, resp3, err := searchEntity(client, attributeColMap, row)
		if err != nil {
			log.Fatalf("Error searching for entity. Error: %v", err)
		}
		r3 := getSearchData(resp3)

		// loop through the results
		for i := 0; i < args.MaxResults; i++ {
			if i > 0 {
				// if there are not second or third match, break
				if len(r1[i].entityID) == 0 && len(r2[i].entityID) == 0 && len(r3[i].entityID) == 0 {
					break
				}
				// remove field values for 2nd and 3rd result to make things easier to read
				fieldValues = []string{"", "", "", "", ""}
			}
			results := []string{
				fieldValues[0], fieldValues[1], fieldValues[2], fieldValues[3], fieldValues[4], // fieldName, fieldAddress, fieldCountry, fieldIdentifier, fieldType
				fmt.Sprintf("Result %v", i+1),                  // Result index
				r1[i].entityID, r2[i].entityID, r3[i].entityID, // ID
				r1[i].name, r2[i].name, r3[i].name, // Name
				r1[i].address, r2[i].address, r3[i].address, // Address
				r1[i].entityCountry, r2[i].entityCountry, r3[i].entityCountry, // Country
				r1[i].entityType, r2[i].entityType, r3[i].entityType, // Type
				r1[i].matchStrength, r2[i].matchStrength, // match strength
			}

			if args.MeasureSupplyChain {
				for _, entityID := range []string{r1[i].entityID, r2[i].entityID, r3[i].entityID} {
					if entityID == "" {
						results = append(results, "", "", "")
						continue
					}

					// use data from cache if exists
					if cachedData, ok := supplyChainCache[entityID]; ok {
						results = append(results, fmt.Sprint(cachedData.hasSupplyChain))
						if cachedData.hasSupplyChain {
							results = append(results, fmt.Sprint(cachedData.numSuppliers), fmt.Sprintf("%.2f", cachedData.avgSupplyChainLen))
						}
						continue
					}

					// calculate if not in cache
					hasSupplyChain, numSuppliers, avgSupplyChainLen := getSupplyChainInfo(client, entityID)
					results = append(results, fmt.Sprint(hasSupplyChain))
					if hasSupplyChain {
						results = append(results, fmt.Sprint(numSuppliers), fmt.Sprintf("%.2f", avgSupplyChainLen))
					}
					// add results to cache
					supplyChainCache[entityID] = supplyChainInfo{
						hasSupplyChain:    hasSupplyChain,
						numSuppliers:      numSuppliers,
						avgSupplyChainLen: avgSupplyChainLen,
					}
				}
			}

			err = w.Write(results)
			if err != nil {
				log.Fatalf("Error writing results. Error: %v", err)
			}
		}
	}
}

func getSupplyChainInfo(client *sdk.Connection, entityID string) (bool, int, float64) {
	var hasSupplyChain bool
	var avgSupplyChainLen float64
	suppliers := make(map[string]interface{})

	// get supply chain data for supplier profile entity
	supplyChainData, err := client.SupplyChain.UpstreamTradeTraversal(context.Background(), entityID, nil)
	if err != nil {
		log.Fatalf("Error getting supply chain data. Error: %v", err)
	}

	// if there is a supply chain, gather metrics
	if len(supplyChainData.Data) > 0 {
		hasSupplyChain = true
		var totalHops int
		for _, supplyChain := range supplyChainData.Data {
			totalHops += len(supplyChain.Path)
			// get all unique entities
			suppliers[supplyChain.Target.Id] = nil
			for _, paths := range supplyChain.Path {
				suppliers[paths.Entity.Id] = nil
			}
		}
		avgSupplyChainLen = float64(totalHops) / float64(len(supplyChainData.Data))
	}

	// remove the entity from its suppliers so it isn't counted
	delete(suppliers, entityID)

	return hasSupplyChain, len(suppliers), avgSupplyChainLen
}

func getFieldInfo(attributeFieldsMap map[string][]int, row []string) []string {
	var fieldName string
	if val, ok := attributeFieldsMap[name]; ok {
		fieldName = row[val[0]]
	}

	var fieldAddress string
	if val, ok := attributeFieldsMap[address]; ok {
		fieldAddress = row[val[0]]
	}

	var fieldCountry string
	if val, ok := attributeFieldsMap[country]; ok {
		fieldCountry = row[val[0]]
	}

	var fieldIdentifier string
	if val, ok := attributeFieldsMap[identifier]; ok {
		fieldIdentifier = row[val[0]]
	}

	var fieldType string
	if val, ok := attributeFieldsMap[entityType]; ok {
		fieldType = row[val[0]]
	}
	return []string{fieldName, fieldAddress, fieldCountry, fieldIdentifier, fieldType}
}

func getResolveData(resp *sayari.ResolutionResponse) []result {
	results := make([]result, args.MaxResults)
	if len(resp.Data) == 0 {
		return results
	}
	for i := range results {
		if len(resp.Data) <= i {
			return results
		}
		thisEntity := resp.Data[i]
		results[i].entityID = thisEntity.EntityId
		results[i].name = thisEntity.Label
		results[i].entityType = fmt.Sprint(thisEntity.Type)
		results[i].matchStrength = thisEntity.MatchStrength.Value

		if len(thisEntity.Addresses) > 0 {
			results[i].address = thisEntity.Addresses[0]
		}
		if len(thisEntity.Countries) > 0 {
			results[i].entityCountry = fmt.Sprint(thisEntity.Countries[0])
		}
	}

	return results
}

func getSearchData(resp *sayari.EntitySearchResponse) []result {
	results := make([]result, args.MaxResults)
	if len(resp.Data) == 0 {
		return results
	}
	for i := range results {
		if len(resp.Data) <= i {
			return results
		}
		thisEntity := resp.Data[i]
		results[i].entityID = thisEntity.Id
		results[i].name = thisEntity.Label
		results[i].entityType = fmt.Sprint(thisEntity.Type)
		if len(thisEntity.Addresses) > 0 {
			results[i].address = thisEntity.Addresses[0]
		}
		if len(thisEntity.Countries) > 0 {
			results[i].entityCountry = fmt.Sprint(thisEntity.Countries[0])
		}
	}

	return results
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
			colVal := row[colNum]
			if colVal == "" {
				continue
			}
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
			countryStr := row[colNum]
			if countryStr == "" {
				continue
			}
			country, err := sayari.NewCountryFromString(countryStr)
			if err != nil {
				log.Fatalf("Error getting country. Error: %v", err)
			}
			entityInfo.Country = append(entityInfo.Country, &country)
		}
	}

	if colNums, ok := attributeColMap[address]; ok {
		for _, colNum := range colNums {
			colVal := row[colNum]
			if colVal == "" {
				continue
			}
			entityInfo.Address = append(entityInfo.Address, &row[colNum])
		}
	}

	if colNums, ok := attributeColMap[dateOfBirth]; ok {
		for _, colNum := range colNums {
			colVal := row[colNum]
			if colVal == "" {
				continue
			}
			entityInfo.DateOfBirth = append(entityInfo.DateOfBirth, &row[colNum])
		}
	}

	if colNums, ok := attributeColMap[contact]; ok {
		for _, colNum := range colNums {
			colVal := row[colNum]
			if colVal == "" {
				continue
			}
			entityInfo.Contact = append(entityInfo.Contact, &row[colNum])
		}
	}

	if colNums, ok := attributeColMap[entityType]; ok {
		for _, colNum := range colNums {
			entityTypeStr := row[colNum]
			if entityTypeStr == "" {
				continue
			}
			entityType, err := sayari.NewEntitiesFromString(entityTypeStr)
			if err != nil {
				log.Fatalf("Error getting entity type. Error: %v", err)
			}
			entityInfo.Type = append(entityInfo.Type, &entityType)
		}
	}

	start := time.Now()
	resp, err := client.Resolution.Resolution(context.Background(), &entityInfo)
	duration := time.Since(start)
	if err != nil {
		log.Println("Error calling resolve function.")
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
			countryStr := row[colNum]
			if countryStr == "" {
				continue
			}
			country, err := sayari.NewCountryFromString(countryStr)
			if err != nil {
				log.Fatalf("Error getting country. Error: %v", err)
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
