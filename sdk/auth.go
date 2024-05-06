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
	id      string
	secret  string
	baseURL string
}

// Connect generates a connection to the production Sayari API
func Connect(id, secret string) (*Connection, error) {
	return ConnectTo(id, secret, sayari.Environments.Production)
}

// ConnectTo generates a connection to the specified instance of the Sayari API
func ConnectTo(id, secret, baseURL string) (*Connection, error) {
	// Connect to auth endpoint and get a token
	tokenResponse, err := getToken(id, secret, baseURL)
	if err != nil {
		return nil, err
	}

	connection := &Connection{
		client.NewClient(
			option.WithHTTPHeader(map[string][]string{"Authorization": {tokenResponse.AccessToken}}),
			option.WithBaseURL(baseURL),
		),
		id,
		secret,
		baseURL,
	}

	// Maintain the token
	go connection.maintainToken(tokenResponse.ExpiresIn)

	// Create clients
	return connection, nil
}

func getToken(id, secret, baseURL string) (*sayari.AuthResponse, error) {
	authClient := auth.NewClient(option.WithBaseURL(baseURL))
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
	tokenResponse, err := getToken(c.id, c.secret, c.baseURL)
	if err != nil {
		log.Fatalf("Error maintining token. Err: %v", err)
	}

	// update client
	c.Client = client.NewClient(option.WithHTTPHeader(map[string][]string{"Authorization": {tokenResponse.AccessToken}}))

	// recurse
	c.maintainToken(tokenResponse.ExpiresIn)
}
