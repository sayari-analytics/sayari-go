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
	results, err := client.Source.ListSources(context.Background(), &sayari.ListSources{})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Printf("Found %v sources", len(results.Data))

	// Get the first source
}
