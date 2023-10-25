package main

import (
	"context"
	"log"
	"os"

	sayari "github.com/sayari-analytics/sayari-go/generated/go"
	"github.com/sayari-analytics/sayari-go/sdk"
)

func main() {
	// NOTE: To connect you most provide your client ID and client secret. To avoid accidentally checking these into git,
	// it is recommended to use ENV variables

	// Create a client to auth against the API
	client, err := sdk.Connect(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Create the request body
	resolutionRequest := sayari.Resolution{
		Name: []*string{sayari.String("Victoria Beckham")},
	}

	// Make the request and handle the error
	resolution, err := client.Resolution.Resolution(context.Background(), &resolutionRequest)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Get the entity details for the best match
	entityDetails, err := client.Entity.GetEntity(context.Background(), resolution.Data[0].EntityId, &sayari.GetEntity{})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Printf("%+v", entityDetails)
}
