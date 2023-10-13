package main

import (
	"context"
	"fmt"
	sayari "github.com/sayari-analytics/sayari-go/generated/go"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"

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

	numWorkers := 100
	numRequests := 1000
	perSecondLimit := 100

	// queue up triggers
	triggers := make(chan bool, 2*numWorkers)
	go func(triggers chan bool) {
		for i := 0; i < numRequests; i++ {
			triggers <- true
		}
		close(triggers)
	}(triggers)

	// monitor and limit rate
	fires := make(chan bool, 2*numWorkers)
	go func(fires chan bool) {
		ticker := time.NewTicker(time.Second)
		var count int
		for {
			select {
			case <-fires:
				count++
			case <-ticker.C:
				fmt.Printf("%v requests per second\n", count)
				count = 0
			}
			if count > perSecondLimit {
				time.Sleep(time.Second)
			}
		}
	}(fires)

	// start workers
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go doWork(client, triggers, fires, &wg)
	}
	wg.Wait()

}

func doWork(c *sdk.Connection, triggers chan bool, fires chan bool, wg *sync.WaitGroup) {
	for range triggers {
		doRandomSearch(c)
		fires <- true
	}
	wg.Done()
}

func doRandomSearch(c *sdk.Connection) {
	randomString := generateRandomString(3)
	_, err := c.Search.SearchEntity(context.Background(), &sayari.SearchEntity{Q: randomString})
	if err != nil {
		log.Fatalf("Failed to search for string '%v'. Err: %v", randomString, err)
	}
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func generateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
