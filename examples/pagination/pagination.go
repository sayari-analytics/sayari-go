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

	// Traversal
	entity, err := client.Search.SearchEntity(context.Background(), &sayari.SearchEntity{Q: "David Konigsberg", Limit: sayari.Int(1)})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	allTraversals, err := client.GetAllTraversalResults(context.Background(), entity.Data[0].Id, &sayari.Traversal{Limit: sayari.Int(1)})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("Got %v results", len(allTraversals))

	// Entities
	allEntities, err := client.GetAllEntitySearchResults(context.Background(), &sayari.SearchEntity{Q: "David Konigsberg", Limit: sayari.Int(5)})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("Got %v results", len(allEntities))

	allEntities, err = client.GetAllEntitySearchResults(context.Background(), &sayari.SearchEntity{Q: "David John Smith", Limit: sayari.Int(5)})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("Got %v results", len(allEntities))

	// Records
	allRecords, err := client.GetAllRecordSearchResults(context.Background(), &sayari.SearchRecord{Q: "David Konigsberg"})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("Got %v results", len(allRecords))
}
