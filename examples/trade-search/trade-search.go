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

	// Do Shipment search
	shipments, err := client.Trade.SearchShipments(context.Background(), &sayari.SearchShipments{Q: "microcenter"})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("Found %v shipments", len(shipments.Data))

	// Do supplier search
	suppliers, err := client.Trade.SearchSuppliers(context.Background(), &sayari.SearchSuppliers{Q: "microcenter"})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("Found %v suppliers", len(suppliers.Data))

	// Do buyer search
	buyers, err := client.Trade.SearchBuyers(context.Background(), &sayari.SearchBuyers{Q: "microcenter"})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("Found %v buyers", len(buyers.Data))
}
