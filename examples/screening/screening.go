package main

import (
	"context"
	"log"
	"os"

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

	// Do screening
	riskyEntities, nonRiskyEntities, unresolved, err := client.ScreenCSVEntities(context.Background(), "examples/screening/entities_to_screen.csv")
	if err != nil {
		log.Fatalf("Failed to screen entities. Err: %v", err)
	}

	log.Printf("Found %v entities with risks.", len(riskyEntities))
	log.Printf("Found %v entities without risks.", len(nonRiskyEntities))
	log.Printf("%v records could not be resolved to entities.", len(unresolved))
}
