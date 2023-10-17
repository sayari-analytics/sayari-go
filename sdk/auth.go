package sdk

import (
	"context"
	"log"
	"time"

	sayari "github.com/sayari-analytics/sayari-go/generated/go"
	"github.com/sayari-analytics/sayari-go/generated/go/auth"
	"github.com/sayari-analytics/sayari-go/generated/go/client"
	"github.com/sayari-analytics/sayari-go/generated/go/core"
)

type Connection struct {
	*client.Client
	id     string
	secret string
}

const clientHeader = "sayari-go"

func Connect(id, secret string) (*Connection, error) {
	// Connect to auth endpoint and get a token
	results, err := getToken(id, secret)
	if err != nil {
		return nil, err
	}

	rateLimter := core.NewRateLimiter()

	connection := &Connection{
		client.NewClient(client.WithAuthToken(results.AccessToken),
			client.WithHeaderClient(clientHeader),
			client.WithRateLimiter(rateLimter),
		),
		id,
		secret,
	}

	// Maintain the token
	go connection.maintainToken(results.ExpiresIn)

	// Create clients
	return connection, nil
}

func (c *Connection) maintainToken(expiresIn int) {
	// wait until 1 hr before expiring
	expiresIn -= 3600
	if expiresIn < 0 {
		expiresIn = 0
	}
	time.Sleep(time.Duration(expiresIn) * time.Second)

	// get updated token
	results, err := getToken(c.id, c.secret)
	if err != nil {
		log.Fatalf("Error maintining token. Err: %v", err)
	}

	// update client
	c.Client = client.NewClient(client.WithAuthToken(results.AccessToken))

	// recurse
	c.maintainToken(results.ExpiresIn)
}

func getToken(id, secret string) (*sayari.AuthResponse, error) {
	authClient := auth.NewClient(client.WithHeaderClient(clientHeader))
	return authClient.GetToken(context.Background(), &sayari.GetToken{
		ClientId:     id,
		ClientSecret: secret,
		Audience:     "sayari.com",
		GrantType:    "client_credentials",
	})
}
