package main

import (
	"context"
	"log"
	"os"

	sayari "github.com/sayari-analytics/sayari-go/generated/go"
	"github.com/sayari-analytics/sayari-go/sdk"

	"github.com/joho/godotenv"
)

func main() {
	// This example is intended to provide a basic use case for the sayari go SDK

	// Load ENV file
	godotenv.Load()

	// To connect you most provide your client ID and client secret. To avoid accidentally checking these into git,
	// it is recommended to use ENV variables
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	if len(clientID) == 0 || len(clientSecret) == 0 {
		log.Fatalln("Both client ID and client secret must be set")
	}

	// Create a client to auth against the API
	client, err := sdk.Connect(clientID, clientSecret)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Println("Connection successful")

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
	searchTerm := "DAKLabb"
	searchResults, err := client.Search.SearchEntity(context.Background(), &sayari.SearchEntity{Q: searchTerm})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	// uncomment to view data
	//spew.Dump(searchResults)
	log.Printf("Found %v entity results for %v", len(searchResults.Data), searchTerm)

	// Get the entity summary
	entitySummary, err := client.Entity.EntitySummary(context.Background(), searchResults.Data[0].Id)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	// uncomment to view data
	//spew.Dump(entitySummary)
	log.Printf("Has address %v: ", entitySummary.Addresses[0])

	// Get the full entity
	entityDetails, err := client.Entity.GetEntity(context.Background(), searchResults.Data[0].Id, &sayari.GetEntity{})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	// uncomment to view data
	//spew.Dump(entitySummary)
	log.Printf("Is referenced by %v sources.", len(entityDetails.ReferencedBy.Data))
}
