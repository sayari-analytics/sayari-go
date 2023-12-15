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
	*core.RateLimiter
}

const clientName = "sayari-go"

func Connect(id, secret string) (*Connection, error) {
	// Connect to auth endpoint and get a token
	tokenResponse, err := getToken(id, secret)
	if err != nil {
		return nil, err
	}

	rateLimter := core.NewRateLimiter()

	connection := &Connection{
		client.NewClient(client.WithToken(tokenResponse.AccessToken),
			client.WithClientName(clientName),
			client.WithRateLimiter(rateLimter),
		),
		id,
		secret,
		rateLimter,
	}

	// Maintain the token
	go connection.maintainToken(tokenResponse.ExpiresIn)

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
	tokenResponse, err := getToken(c.id, c.secret)
	if err != nil {
		log.Fatalf("Error maintining token. Err: %v", err)
	}

	// update client
	c.RateLimiter.Block()
	c.Client = client.NewClient(
		client.WithToken(tokenResponse.AccessToken),
		client.WithClientName(clientName),
		client.WithRateLimiter(c.RateLimiter),
	)
	c.RateLimiter.UnBlock()

	// recurse
	c.maintainToken(tokenResponse.ExpiresIn)
}

func getToken(id, secret string) (*sayari.AuthResponse, error) {
	authClient := auth.NewClient(client.WithClientName(clientName))
	return authClient.GetToken(context.Background(), &sayari.GetToken{
		ClientId:     id,
		ClientSecret: secret,
		Audience:     sayari.AudienceSayari,
		GrantType:    sayari.GrantTypeClientCredentials,
	})
}
