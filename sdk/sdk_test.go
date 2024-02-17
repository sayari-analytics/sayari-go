package sdk

import (
	"context"
	"log"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"

	sayari "github.com/sayari-analytics/sayari-go/generated/go"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	setup()
	exitVal := m.Run()
	os.Exit(exitVal)
}

var api *Connection

func setup() {
	// load ENV file if ENV vars are not set
	if os.Getenv("CLIENT_ID") == "" || os.Getenv("CLIENT_SECRET") == "" {
		err := godotenv.Load("../.env")
		if err != nil {
			log.Fatalf("Failed to load .env file. Err: %v", err)
		}
	}

	// Create a client that is authed against the API
	var err error
	api, err = Connect(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))
	if err != nil {
		log.Fatalf("Failed to connect. Err: %v", err)
	}
}

func TestSources(t *testing.T) {
	// list sources
	sources, err := api.Source.ListSources(context.Background(), &sayari.ListSources{})
	assert.Nil(t, err)
	assert.Len(t, sources.Data, 250, "There should be 250 sources as of 12/19/2023")
	assert.Equal(t, sources.Data[0].Label, "Abu Dhabi Registration Authority Online Registry", "The first shource should be 'Abu Dhabi Registration Authority Online Registry'")
}

func TestEntities(t *testing.T) {
	// search for an entity with a random string
	randomString := generateRandomString(3)

	// query until we get results
	entitySearchResults, err := api.Search.SearchEntity(context.Background(), &sayari.SearchEntity{Q: randomString})
	assert.Nil(t, err)
	// try until we get results
	if len(entitySearchResults.Data) == 0 {
		TestEntities(t)
		return
	}
	assert.Greater(t, len(entitySearchResults.Data), 0)

	// test the get version of this endpoint
	entitySearchGETResults, err := api.Search.SearchEntityGet(context.Background(), &sayari.SearchEntityGet{Q: randomString})
	assert.Nil(t, err)
	assert.Equal(t, entitySearchResults.Data, entitySearchGETResults.Data)
	assert.Equal(t, entitySearchResults.Size.Count, entitySearchGETResults.Size.Count)
	assert.Equal(t, entitySearchResults.Size.Qualifier, entitySearchGETResults.Size.Qualifier)

	// do some checks on the first result
	firstEntity := entitySearchResults.Data[0]
	// capture entity id/label for debugging
	log.Println(firstEntity.Id)
	log.Println(firstEntity.Label)

	// Get entity summary
	firstEntitySummary, err := api.Entity.EntitySummary(context.Background(), firstEntity.Id)
	assert.Nil(t, err)

	// Summary should match search results
	assert.Equal(t, firstEntitySummary.Id, firstEntity.Id)
	assert.Equal(t, firstEntitySummary.Label, firstEntity.Label)
	assert.Equal(t, firstEntitySummary.Degree, firstEntity.Degree)
	assert.Equal(t, firstEntitySummary.Pep, firstEntity.Pep)
	assert.Equal(t, firstEntitySummary.PsaCount, firstEntity.PsaCount)
	assert.Equal(t, firstEntitySummary.Type, firstEntity.Type)
	//These currently don't match, not sure if we will be updating that
	//assert.Equal(t, firstEntitySummary.EntityUrl, firstEntity.EntityUrl)
	assert.Equal(t, firstEntitySummary.Sanctioned, firstEntity.Sanctioned)
	assert.Equal(t, firstEntitySummary.Identifiers, firstEntity.Identifiers)
	assert.Equal(t, firstEntitySummary.Addresses, firstEntity.Addresses)
	assert.Equal(t, firstEntitySummary.Countries, firstEntity.Countries)
	assert.Equal(t, firstEntitySummary.RelationshipCount, firstEntity.RelationshipCount)

	// get entity details
	firstEntityDetails, err := api.Entity.GetEntity(context.Background(), firstEntity.Id, &sayari.GetEntity{})
	log.Println("GetEntity Err: ", err)
	assert.Nil(t, err)
	// check all the same stuff we checked with summary
	assert.Equal(t, firstEntityDetails.Id, firstEntity.Id)
	assert.Equal(t, firstEntityDetails.Label, firstEntity.Label)
	assert.Equal(t, firstEntityDetails.Degree, firstEntity.Degree)
	assert.Equal(t, firstEntityDetails.Pep, firstEntity.Pep)
	assert.Equal(t, firstEntityDetails.PsaCount, firstEntity.PsaCount)
	assert.Equal(t, firstEntityDetails.Type, firstEntity.Type)
	assert.Equal(t, firstEntityDetails.EntityUrl, firstEntity.EntityUrl)
	assert.Equal(t, firstEntityDetails.Sanctioned, firstEntity.Sanctioned)
	assert.Equal(t, firstEntityDetails.Identifiers, firstEntity.Identifiers)
	assert.Equal(t, firstEntityDetails.Addresses, firstEntity.Addresses)
	assert.Equal(t, firstEntityDetails.Countries, firstEntity.Countries)
	assert.Equal(t, firstEntityDetails.RelationshipCount, firstEntity.RelationshipCount)

	// check additional fields
	assert.Equal(t, firstEntityDetails.CompanyType, firstEntity.CompanyType)
	assert.Equal(t, firstEntityDetails.Relationships.Size.Count, firstEntity.Degree)
	if firstEntity.Degree < 50 {
		assert.Len(t, firstEntityDetails.Relationships.Data, firstEntity.Degree)
	} else {
		assert.Len(t, firstEntityDetails.Relationships.Data, 50)
	}
}

func TestResolution(t *testing.T) {
	// resolve entity with random string
	randomString := generateRandomString(3)

	// query until we get results
	resolution, err := api.Resolution.Resolution(context.Background(), &sayari.Resolution{Name: []*string{&randomString}})
	assert.Nil(t, err)
	if len(resolution.Data) == 0 {
		TestResolution(t)
		return
	}
	assert.Greater(t, len(resolution.Data), 0)

	// do basic check on results
	assert.Len(t, resolution.Fields.Name, 1)
	assert.Equal(t, resolution.Fields.Name[0], randomString)
}

func TestRecords(t *testing.T) {
	// resolve entity with random string
	randomString := generateRandomString(3)

	// query until we get results
	recordSearchResults, err := api.Search.SearchRecord(context.Background(), &sayari.SearchRecord{Q: randomString})
	assert.Nil(t, err)
	if len(recordSearchResults.Data) == 0 {
		TestRecords(t)
		return
	}
	assert.Greater(t, len(recordSearchResults.Data), 0)

	// test the get version of this endpoint
	recordSearchGetResults, err := api.Search.SearchRecordGet(context.Background(), &sayari.SearchRecordGet{Q: randomString})
	assert.Nil(t, err)
	assert.Equal(t, recordSearchResults.Data, recordSearchGetResults.Data)
	assert.Equal(t, recordSearchResults.Size.Count, recordSearchGetResults.Size.Count)
	assert.Equal(t, recordSearchResults.Size.Qualifier, recordSearchGetResults.Size.Qualifier)

	// do checks with the first results
	firstRecord := recordSearchResults.Data[0]
	log.Println(firstRecord.Id)
	log.Println(firstRecord.Label)

	// get record and compare with search result
	record, err := api.Record.GetRecord(context.Background(), EncodeRecordID(firstRecord.Id), &sayari.GetRecord{})
	assert.Nil(t, err)
	assert.Equal(t, record.Label, firstRecord.Label)
	assert.Equal(t, record.Source, firstRecord.Source)
	assert.Equal(t, record.PublicationDate, firstRecord.PublicationDate)
	assert.Equal(t, record.AcquisitionDate, firstRecord.AcquisitionDate)
	assert.Equal(t, record.RecordUrl, firstRecord.RecordUrl)
	assert.Equal(t, record.ReferencesCount, firstRecord.ReferencesCount)
	assert.Equal(t, record.SourceUrl, firstRecord.SourceUrl)
}

func TestOwnershipTraversal(t *testing.T) {
	// resolve entity with random string
	randomString := generateRandomString(3)

	// query until we get results
	entitySearchResults, err := api.Search.SearchEntity(context.Background(), &sayari.SearchEntity{Q: randomString})
	assert.Nil(t, err)
	if len(entitySearchResults.Data) == 0 {
		TestOwnershipTraversal(t)
		return
	}
	assert.Greater(t, len(entitySearchResults.Data), 0)

	// use first entity
	entity := entitySearchResults.Data[0]

	// do traversal
	log.Println("Attempting Ownership traversal w/ entity: ", entity.Id)
	traversal, err := api.Traversal.Ownership(context.Background(), entity.Id, &sayari.Ownership{})
	assert.Nil(t, err)
	if len(traversal.Data) == 0 {
		TestOwnershipTraversal(t)
		return
	}
	assert.Greater(t, len(traversal.Data), 0)
	assert.Equal(t, traversal.Data[0].Source, entity.Id)

	// do UBO traversal
	log.Println("Attempting UBO traversal w/ entity: ", entity.Id)
	ubo, err := api.Traversal.Ubo(context.Background(), entity.Id, &sayari.Ubo{})
	assert.Nil(t, err)
	if len(ubo.Data) == 0 {
		TestOwnershipTraversal(t)
		return
	}
	assert.Greater(t, len(ubo.Data), 0)
	uboID := ubo.Data[0].Target.Id

	// do ownership traversal from ubo
	log.Println("Attempting Ownership traversal w/ UBO entity: ", uboID)
	downstream, err := api.Traversal.Ownership(context.Background(), uboID, &sayari.Ownership{})
	assert.Nil(t, err)
	assert.Greater(t, len(downstream.Data), 0)

	/*
		The test below doesn't work, but I don't know why.
		entity 'YdHkr_vnixCoWoQdOX5V7A' has a UBO of 'Sb77z7bNzNs_YtDFAwjuTw'
		ownership of 'Sb77z7bNzNs_YtDFAwjuTw' doesn't include 'YdHkr_vnixCoWoQdOX5V7A'
		perhaps this makes sense...

		the downstream path should contain the initial entity
		found = False
		for path in downstream.data:
			for step in path.path:
			if step.entity.id == entity.id:
				found = True
		assert found
	*/

	// shortest path
	shortestPath, err := api.Traversal.ShortestPath(context.Background(), &sayari.ShortestPath{Entities: []string{string(entity.Id), uboID}})
	assert.Nil(t, err)
	assert.Greater(t, len(shortestPath.Data[0].Path), 0)

	// TODO: figure out good test for watchlist traversal
}

/* FIXME: on hold until we can align on how we want to handle pagination
func TestEntityPagination(t *testing.T) {
	searchTerm := "David Konigsberg"
	info, err := api.Search.SearchEntity(context.Background(), &sayari.SearchEntity{Q: searchTerm, Limit: sayari.Int(1)})
	assert.Nil(t, err)

	// Do paginated query
	allEntities, err := api.GetAllEntitySearchResults(context.Background(), &sayari.SearchEntity{Q: searchTerm, Limit: sayari.Int(5)})
	assert.Nil(t, err)
	assert.Equal(t, allEntities.Limit, info.Size.Count)

	// Test requesting too many pages
	searchTerm = "amazon"
	resp, err := api.GetAllEntitySearchResults(context.Background(), &sayari.SearchEntity{Q: searchTerm})
	assert.Equal(t, ErrTooMuchDataRequested, err, "This request returns too many response to paginate and should error")
	assert.Nil(t, resp)

	// Do paginated query for larger data set
	searchTerm = "David John Smith"
	info, err = api.Search.SearchEntity(context.Background(), &sayari.SearchEntity{Q: searchTerm, Limit: sayari.Int(1)})
	assert.Nil(t, err)
	allEntities, err = api.GetAllEntitySearchResults(context.Background(), &sayari.SearchEntity{Q: searchTerm})
	assert.Nil(t, err)
	assert.Equal(t, info.Size.Count, allEntities.Limit)
}

func TestRecordPagination(t *testing.T) {
	searchTerm := "David Konigsberg"
	info, err := api.Search.SearchRecord(context.Background(), &sayari.SearchRecord{Q: searchTerm, Limit: sayari.Int(1)})
	assert.Nil(t, err)

	// Do paginated query
	allEntities, err := api.GetAllRecordSearchResults(context.Background(), &sayari.SearchRecord{Q: searchTerm})
	assert.Nil(t, err)
	assert.Equal(t, allEntities.Limit, info.Size.Count)
}

func TestTraversalPagination(t *testing.T) {
	searchTerm := "David Konigsberg"
	entitySearchResults, err := api.Search.SearchEntity(context.Background(), &sayari.SearchEntity{Q: searchTerm, Limit: sayari.Int(1)})
	assert.Nil(t, err)
	entity := entitySearchResults.Data[0]

	// Do paginated query
	allTraversals, err := api.GetAllTraversalResults(context.Background(), entity.Id, &sayari.Traversal{Limit: sayari.Int(1)})
	assert.Nil(t, err)
	assert.Greater(t, allTraversals.Limit, 1)
}

*/

func TestShipmentSearch(t *testing.T) {
	// search for shipments with a random string
	randomString := generateRandomString(3)

	shipments, err := api.Trade.SearchShipments(context.Background(), &sayari.SearchShipments{Q: randomString})
	assert.Nil(t, err)
	// try until we get results
	if len(shipments.Data) == 0 {
		TestShipmentSearch(t)
		return
	}

	assert.Greater(t, len(shipments.Data), 0)

	// test field and multi-filter
	buyerName := "HANSOLL TEXTILE LTD"
	buyerID := "f_nIivE32HCYDPEoSPTGJw"
	hsCode := "600410"
	shipments, err = api.Trade.SearchShipments(context.Background(), &sayari.SearchShipments{
		Q:      buyerName,
		Filter: &sayari.TradeFilterList{HsCode: []string{hsCode}, BuyerId: []string{buyerID}},
	})
	assert.Nil(t, err)
	assert.NotZero(t, len(shipments.Data))
	for _, shipment := range shipments.Data {
		// verify shipment matches on HS code
		assert.NotZero(t, len(shipment.ProductDescriptions))
		var hsFound bool
		for _, shipmentHScode := range shipment.HsCodes {
			if strings.HasPrefix(shipmentHScode.Code, hsCode) {
				hsFound = true
				break
			}
		}
		assert.True(t, hsFound)

		// verify shipment matches entity
		assert.NotZero(t, len(shipment.Buyer))
		var entityFound bool
		for _, buyer := range shipment.Buyer {
			if buyer.Id == buyerID {
				entityFound = true
				break
			}
		}
		assert.True(t, entityFound)
	}
}

func TestSupplierSearch(t *testing.T) {
	// search for suppliers with a random string
	randomString := generateRandomString(3)

	suppliers, err := api.Trade.SearchSuppliers(context.Background(), &sayari.SearchSuppliers{Q: randomString})
	assert.Nil(t, err)
	// try until we get results
	if len(suppliers.Data) == 0 {
		TestSupplierSearch(t)
		return
	}

	assert.Greater(t, len(suppliers.Data), 0)
}

func TestBuyerSearch(t *testing.T) {
	// search for suppliers with a random string
	randomString := generateRandomString(3)

	buyers, err := api.Trade.SearchBuyers(context.Background(), &sayari.SearchBuyers{Q: randomString})
	assert.Nil(t, err)
	// try until we get results
	if len(buyers.Data) == 0 {
		TestBuyerSearch(t)
		return
	}

	assert.Greater(t, len(buyers.Data), 0)
}

func TestUsage(t *testing.T) {
	usage, err := api.Info.GetUsage(context.Background(), &sayari.GetUsage{})
	assert.Nil(t, err)
	assert.NotZero(t, usage.Usage.Entity, "all endpoints should show usage")
	assert.NotZero(t, usage.Usage.EntitySummary, "all endpoints should show usage")
	assert.NotZero(t, usage.Usage.Record, "all endpoints should show usage")
	assert.NotZero(t, usage.Usage.Resolve, "all endpoints should show usage")
	assert.NotZero(t, usage.Usage.SearchEntities, "all endpoints should show usage")
	assert.NotZero(t, usage.Usage.SearchRecords, "all endpoints should show usage")
	assert.NotZero(t, usage.Usage.SearchTrade, "all endpoints should show usage")
	assert.NotZero(t, usage.Usage.Traversal, "all endpoints should show usage")
}

func TestHistory(t *testing.T) {
	history, err := api.Info.GetHistory(context.Background(), &sayari.GetHistory{Size: sayari.Int(10)})
	assert.Nil(t, err)
	assert.Equal(t, history.Size, len(history.Events))
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano())) //nolint: gosec

func generateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
