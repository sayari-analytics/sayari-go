package main

import (
	"context"
	"log"
	"os"
	"time"

	sayari "github.com/sayari-analytics/sayari-go/generated/go"
	"github.com/sayari-analytics/sayari-go/sdk"

	"github.com/joho/godotenv"
)

func main() {
	// load ENV file if ENV vars are not set
	if os.Getenv("CLIENT_ID") == "" || os.Getenv("CLIENT_SECRET") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Failed to load .env file. Err: %v", err)
		}
	}

	// NOTE: To connect you most provide your client ID and client secret. To avoid accidentally checking these into git,
	// it is recommended to use ENV variables

	// Create a client to auth against the API
	client, err := sdk.Connect(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// List sources
	sources, err := client.Source.ListSources(context.Background(), &sayari.ListSources{})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Printf("Found %v sources", len(sources.Data))

	// Get the first source
	firstSource, err := client.Source.GetSource(context.Background(), sources.Data[0].Id)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	// uncomment to view data
	//spew.Dump(firstSource)
	log.Printf("First source is: %v", firstSource.Label)

	// Search for an entity
	searchTerm := "Slickdeals"
	entitySearchResults, err := client.Search.SearchEntity(context.Background(), &sayari.SearchEntity{Q: searchTerm})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	// uncomment to view data
	//spew.Dump(searchResults)
	log.Printf("Found %v entity results for %v", len(entitySearchResults.Data), searchTerm)

	firstEntityResult := entitySearchResults.Data[0].Id

	// Get the entity summary
	entitySummary, err := client.Entity.EntitySummary(context.Background(), firstEntityResult)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	// uncomment to view data
	//spew.Dump(entitySummary)
	log.Printf("Has address '%v'", entitySummary.Addresses[0])

	// Get the full entity
	entityDetails, err := client.Entity.GetEntity(context.Background(), firstEntityResult, &sayari.GetEntity{})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	// uncomment to view data
	//spew.Dump(entitySummary)
	log.Printf("Is referenced by %v sources.", len(entityDetails.ReferencedBy.Data))

	// Resolve
	resolution, err := client.Resolution.Resolution(context.Background(), &sayari.Resolution{Name: []*string{sayari.String(searchTerm)}})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	// uncomment to view data
	//spew.Dump(resolution)
	log.Printf("Resolved to %v entities.", len(resolution.Data))

	// Search for record
	recordSearch, err := client.Search.SearchRecord(context.Background(), &sayari.SearchRecord{Q: searchTerm})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	// uncomment to view data
	//spew.Dump(recordSearch)
	log.Printf("Found %v records.", len(recordSearch.Data))

	// Get record
	record, err := client.Record.GetRecord(context.Background(), sdk.EncodeRecordID(recordSearch.Data[0].Id), &sayari.GetRecord{})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	// uncomment to view data
	//spew.Dump(record)
	log.Printf("Found record: %v.", record.Label)

	// Do traversal
	traversal, err := client.Traversal.Traversal(context.Background(), firstEntityResult, &sayari.Traversal{})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	// uncomment to view data
	//spew.Dump(traversal)
	log.Printf("Did traversal of entity %v. Found %v related things.", firstEntityResult, len(traversal.Data))

	// Do UBO traversal
	ubo, err := client.Traversal.Ubo(context.Background(), firstEntityResult, &sayari.Ubo{})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	// uncomment to view data
	//spew.Dump(ubo)
	log.Printf("Found %v beneficial owners.", len(ubo.Data))

	// Ownership
	downstream, err := client.Traversal.Ownership(context.Background(), ubo.Data[0].Target.Id, &sayari.Ownership{})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	// uncomment to view data
	//spew.Dump(downstream)
	log.Printf("Found %v downstream things owned by the first UBO of %v.", len(downstream.Data), searchTerm)

	// Fetch an entity likely to be associated with watch lists
	putinResult, err := client.Search.SearchEntity(context.Background(), &sayari.SearchEntity{Q: "putin"})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	// Check watchlist
	watchlist, err := client.Traversal.Watchlist(context.Background(), putinResult.Data[0].Id, &sayari.Watchlist{})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	// uncomment to view data
	//spew.Dump(watchlist)
	log.Printf("Found %v watchlist resulsts for entity %v.", len(watchlist.Data), putinResult.Data[0].Id)

	// Shortest Path
	shortestPath, err := client.Traversal.ShortestPath(context.Background(), &sayari.ShortestPath{Entities: []string{firstEntityResult, ubo.Data[0].Target.Id}})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	// uncomment to view data
	//spew.Dump(shortestPath)
	log.Printf("Found path with %v hops", len(shortestPath.Data[0].Path))

	// Check usage
	usage, err := client.Info.GetUsage(context.Background(), &sayari.GetUsage{})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("Entity summary usage: %v", *usage.Usage.EntitySummary)

	// Check history
	historyParams := sayari.GetHistory{
		Size: sayari.Int(10000),
		From: sayari.Time(time.Now().AddDate(0, 0, -2)),
		To:   sayari.Time(time.Now().AddDate(0, 0, -1)),
	}

	history, err := client.Info.GetHistory(context.Background(), &historyParams)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("Found %v events from %v to %v", len(history.Events), historyParams.From.Format("2006-01-02"), historyParams.To.Format("2006-01-02"))
}
