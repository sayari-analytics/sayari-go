package main

import (
	"context"
	sayari "github.com/sayari-analytics/sayari-go/generated/go"
	"github.com/sayari-analytics/sayari-go/sdk"
	"log"
	"os"

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

	// Use the base URL ENV var if provided
	baseURL := sayari.Environments.Production
	if os.Getenv("BASE_URL") != "" {
		baseURL = os.Getenv("BASE_URL")
	}

	// Create a client to auth against the API
	client, err := sdk.ConnectTo(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), baseURL)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	searchQuery := "Victoria Beckham"
	entitySearchResults, err := client.Search.SearchEntity(context.Background(), &sayari.SearchEntity{Q: searchQuery})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for _, result := range entitySearchResults.Data {
		log.Printf("Name: %v", result.Label)
		log.Printf("Entity_ID: %v", result.Id)

	}

}
