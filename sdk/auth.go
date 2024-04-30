package sdk

import (
	"context"
	"log"
	"time"

	"github.com/sayari-analytics/sayari-go/generated/go/option"

	sayari "github.com/sayari-analytics/sayari-go/generated/go"
	"github.com/sayari-analytics/sayari-go/generated/go/auth"
	"github.com/sayari-analytics/sayari-go/generated/go/client"
)

type Connection struct {
	*client.Client
	id     string
	secret string
}

func Connect(id, secret string) (*Connection, error) {
	// Connect to auth endpoint and get a token
	tokenResponse, err := getToken(id, secret)
	if err != nil {
		return nil, err
	}

	connection := &Connection{
		client.NewClient(
			option.WithHTTPHeader(map[string][]string{"Authorization": {tokenResponse.AccessToken}}),
			option.WithClientName(string(sayari.ClientNameGo)),
		),
		id,
		secret,
	}

	// Maintain the token
	go connection.maintainToken(tokenResponse.ExpiresIn)

	// Create clients
	return connection, nil
}

func getToken(id, secret string) (*sayari.AuthResponse, error) {
	authClient := auth.NewClient(option.WithClientName(string(sayari.ClientNameGo)))
	return authClient.GetToken(context.Background(), &sayari.GetToken{
		ClientId:     id,
		ClientSecret: secret,
	})
}

func (c *Connection) maintainToken(expiresIn int) {
	// wait until 1 hr before expiring
	expiresIn -= 3600
	if expiresIn < 0 {
		expiresIn = 0
	}
	time.Sleep(time.Duration(expiresIn) * time.Second)

	// get updated token
	tokenResponse, err := getToken(c.id, c.secret)
	if err != nil {
		log.Fatalf("Error maintining token. Err: %v", err)
	}

	// update client
	c.Client = client.NewClient(
		option.WithHTTPHeader(map[string][]string{"Authorization": {tokenResponse.AccessToken}}),
		option.WithClientName(string(sayari.ClientNameGo)),
	)

	// recurse
	c.maintainToken(tokenResponse.ExpiresIn)
}
