package main

import (
	"context"
	"encoding/csv"
	"fmt"
	arg "github.com/alexflint/go-arg"
	"github.com/joho/godotenv"
	sayari "github.com/sayari-analytics/sayari-go/generated/go"
	"github.com/sayari-analytics/sayari-go/sdk"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	name        = "name"
	identifier  = "identifier"
	country     = "country"
	address     = "address"
	dateOfBirth = "dateofbirth"
	contact     = "contact"
	entityType  = "type"
	tag         = "tag"

	noHitMsg = "no-hit" // the string that is returned when a search yields no hits.
)

type apiResult struct {
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
	tag:         "An arbitrary tag passed through to the output",
}

type supplyChainInfo struct {
	hasSupplyChain    string
	numSuppliers      string
	avgSupplyChainLen string
}

type Job struct {
	rowNum int
	row    []string
}

type Results struct {
	rowNum int
	output [][]string
}

var numWorkers = 7
var maxResults = 3
var cloudflareRetry bool
var runInDev bool
var args struct {
	MaxResults         int  // the maximum number of results to return for each search (defaults to 3)
	MeasureSupplyChain bool // set to true if you want to include supply chain metrics
	LogTimes           bool
	NumWorkers         int
	CloudflareRetry    bool // retry if we get a cloudflare error
	Dev                bool // run against dev ENV (requires DEV_CLIENT_ID, DEV_CLIENT_SECRET, and DEV_BASE_URL)
}

var scCache supplyChainCache

type supplyChainCache struct {
	sync.Mutex
	data map[string]supplyChainInfo
}

func main() {
	// load ENV file if ENV vars are not set
	if os.Getenv("CLIENT_ID") == "" || os.Getenv("CLIENT_SECRET") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Failed to load .env file. Err: %v", err)
		}
	}

	// parse cmd line args (use --help to display)
	arg.MustParse(&args)
	if args.MaxResults != 0 {
		log.Println("Setting max results to ", args.MaxResults)
		maxResults = args.MaxResults
	}
	if args.NumWorkers != 0 {
		log.Println("Setting num workers to ", args.NumWorkers)
		numWorkers = args.NumWorkers
	}
	if args.CloudflareRetry {
		log.Println("Cloudflare retry enabled")
		cloudflareRetry = true
	}
	if args.Dev {
		log.Println("Running against Dev")
		runInDev = true
	}

	time.Sleep(time.Second)

	// Use the base URL ENV var if provided
	baseURL := sayari.Environments.Production
	if os.Getenv("BASE_URL") != "" {
		baseURL = os.Getenv("BASE_URL")
	}
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	if runInDev {
		baseURL = os.Getenv("DEV_BASE_URL")
		clientID = os.Getenv("DEV_CLIENT_ID")
		clientSecret = os.Getenv("DEV_CLIENT_SECRET")
	}

	// Create a client to auth against the API
	client, err := sdk.ConnectTo(clientID, clientSecret, baseURL)
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
	if err != nil {
		log.Fatalln("failed to map rows to attribute columns", err)
	}

	headers := []string{
		"field_name", "field_address", "field_country", "field_identifier", "field_type", "result index",
		"corporate_entity_id", "supplier_entity_id", "search_entity_id",
		"corporate_name", "supplier_name", "search_name",
		"corporate_address", "supplier_address", "search_address",
		"corporate_country", "supplier_country", "search_country",
		"corporate_type", "supplier_type", "search_type",
		"corporate_strength", "supplier_strength",
	}

	// include supply chain metrics if desired
	if args.MeasureSupplyChain {
		headers = append(headers,
			"corporate_supply_chain", "corporate_suppliers_count", "corporate_avg_supply_chain_len",
			"supplier_supply_chain", "supplier_suppliers_count", "supplier_avg_supply_chain_len",
			"search_supply_chain", "search_suppliers_count", "search_avg_supply_chain_len",
		)
		// initialize cache
		scCache.data = make(map[string]supplyChainInfo)
	}

	// include tag if provided
	var includeTags bool
	if _, ok := attributeColMap[tag]; ok {
		includeTags = true
		headers = append([]string{"tag"}, headers...)
	}

	w.Write(headers)

	// setups workers
	jobChan := make(chan Job, numWorkers*5)
	resultChan := make(chan Results, numWorkers*5)
	doneChan := make(chan bool)
	for i := 0; i < numWorkers; i++ {
		go processRows(i, client, jobChan, resultChan, doneChan, attributeColMap, includeTags)
	}

	// collate and write results
	resultsDone := make(chan bool)
	go handleResults(w, len(rows), resultChan, resultsDone)

	// Process each row
	for i, row := range rows {
		// skip first row
		if i == 0 {
			log.Println("Input Headers: ", row)
			continue
		}

		jobChan <- Job{
			rowNum: i,
			row:    row,
		}
	}

	// close job chan (workers will run until chan is drained)
	close(jobChan)

	// wait until all workers have responded done
	for i := 0; i < numWorkers; i++ {
		<-doneChan
	}
	log.Println("All workers done.")

	<-resultsDone
	log.Println("Done processing results")
}

type buffer struct {
	sync.Mutex
	data map[int]Results
}

func bufferResults(buffer *buffer, resultsChan chan Results) {
	// make map for buffer
	buffer.Lock()
	buffer.data = make(map[int]Results)
	buffer.Unlock()

	// put results into buffer
	for results := range resultsChan {
		buffer.Lock()
		buffer.data[results.rowNum] = results
		buffer.Unlock()
	}
}

func handleResults(w *csv.Writer, numResult int, resultsChan chan Results, doneChan chan bool) {
	// read all results into a buffer
	resultBuffer := new(buffer)
	go bufferResults(resultBuffer, resultsChan)

	// check the buffer periodically for the next result
	ticker := time.NewTicker(200 * time.Millisecond)
	rowNum := 1
	for rowNum < numResult {
		select {
		case <-ticker.C:
			resultBuffer.Lock()
			rowResults, found := resultBuffer.data[rowNum]
			resultBuffer.Unlock()
			if found {
				for _, result := range rowResults.output {
					err := w.Write(result)
					if err != nil {
						log.Fatalf("Error writing output: %v", err)
					}
				}
				// clean out item from map
				resultBuffer.Lock()
				delete(resultBuffer.data, rowNum)
				resultBuffer.Unlock()

				// write results and increment desired row
				w.Flush()
				rowNum++
			}
		}
	}
	doneChan <- true
}

func processRows(workerID int, client *sdk.Connection, jobChan chan Job, resultsChan chan Results, doneChan chan bool,
	attributeColMap map[string][]int, includeTags bool) {

	log.Printf("Starting worker %d", workerID+1)

	// step through job chan until it closes
	for job := range jobChan {
		log.Printf("Processing row %v", job.rowNum)

		rowStart := time.Now()
		jobResults := Results{rowNum: job.rowNum}
		fieldValues := getFieldInfo(attributeColMap, job.row)

		// Resolve corporate profile
		resp1, err := resolveEntity(client, sayari.ProfileEnumCorporate, attributeColMap, job.row)
		if err != nil {
			log.Fatalf("Error resolving entity w/ corporate profile. Error: %v", err)
		}
		r1 := getResolveData(resp1)

		// Resolve supplies profile
		resp2, err := resolveEntity(client, sayari.ProfileEnumSuppliers, attributeColMap, job.row)
		if err != nil {
			log.Fatalf("Error resolving entity w/ suppliers profile. Error: %v", err)
		}
		r2 := getResolveData(resp2)

		// Search
		resp3, err := searchEntity(client, attributeColMap, job.row)
		if err != nil {
			log.Fatalf("Error searching for entity. Error: %v", err)
		}
		r3 := getSearchData(resp3)

		// loop through the results
		for i := 0; i < maxResults; i++ {
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
					scCache.Lock()
					cachedData, ok := scCache.data[entityID]
					scCache.Unlock()
					if ok {
						results = append(results, cachedData.hasSupplyChain, cachedData.numSuppliers, cachedData.avgSupplyChainLen)
						continue
					}

					// calculate if not in cache
					hasSupplyChain, numSuppliers, avgSupplyChainLen := getSupplyChainInfo(client, entityID)
					results = append(results, hasSupplyChain, numSuppliers, avgSupplyChainLen)

					// add results to cache
					scCache.Lock()
					scCache.data[entityID] = supplyChainInfo{
						hasSupplyChain:    hasSupplyChain,
						numSuppliers:      numSuppliers,
						avgSupplyChainLen: avgSupplyChainLen,
					}
					scCache.Unlock()
				}
			}

			// include tag in result if desired
			if includeTags {
				results = append([]string{job.row[attributeColMap[tag][0]]}, results...)
			}
			jobResults.output = append(jobResults.output, results)
		}
		resultsChan <- jobResults
		if args.LogTimes {
			log.Printf("Processed row %v in %v", job.rowNum, time.Since(rowStart))
		}
	}

	log.Printf("Worker %d done", workerID+1)
	doneChan <- true
	return
}

func getSupplyChainInfo(client *sdk.Connection, entityID string) (string, string, string) {
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
	if hasSupplyChain {
		return fmt.Sprint(hasSupplyChain), fmt.Sprint(len(suppliers)), fmt.Sprintf("%.2f", avgSupplyChainLen)
	}
	return fmt.Sprint(hasSupplyChain), "", ""
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

func getResolveData(resp *sayari.ResolutionResponse) []apiResult {
	results := make([]apiResult, maxResults)
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

func getSearchData(resp *sayari.EntitySearchResponse) []apiResult {
	results := make([]apiResult, maxResults)
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

func resolveEntity(client *sdk.Connection, profile sayari.ProfileEnum, attributeColMap map[string][]int, row []string) (*sayari.ResolutionResponse, error) {
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

	resp, err := client.Resolution.Resolution(context.Background(), &entityInfo)
	if err != nil {
		// got 400 cloudflare error, reattempt
		if cloudflareRetry && strings.Contains(err.Error(), "cloudflare") {
			log.Println("Got cloudflare error on resolve query, retrying in 1 min.")
			time.Sleep(time.Minute)
			return resolveEntity(client, profile, attributeColMap, row)
		}
		return nil, err
	}

	return resp, nil
}

func searchEntity(client *sdk.Connection, attributeColMap map[string][]int, row []string) (*sayari.EntitySearchResponse, error) {
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

	resp, err := client.Search.SearchEntity(context.Background(), &entityInfo)
	if err != nil {
		// got 400 cloudflare error, reattempt
		if cloudflareRetry && strings.Contains(err.Error(), "cloudflare") {
			log.Println("Got cloudflare error on search query, retrying in 1 min.")
			time.Sleep(time.Minute)
			return searchEntity(client, attributeColMap, row)
		}
		return nil, err
	}

	return resp, nil
}
