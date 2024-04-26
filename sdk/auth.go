package sdk

import (
	"context"
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
		client.NewClient(option.WithToken(tokenResponse.AccessToken),
			option.WithClientName(string(sayari.ClientNameGo)),
		),
		id,
		secret,
	}

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
