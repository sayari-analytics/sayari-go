package sdk

import (
	"context"
	generatedgo "github.com/sayari-analytics/sayari-go/generated/go"
	"github.com/sayari-analytics/sayari-go/generated/go/auth"
	"github.com/sayari-analytics/sayari-go/generated/go/client"
)

func Connect(id, secret string) (*client.Client, error) {
	// Connect to auth endpoint and get a token
	authClient := auth.NewClient()
	results, err := authClient.GetToken(context.Background(), &generatedgo.GetToken{
		ClientId:     id,
		ClientSecret: secret,
		Audience:     "sayari.com",
		GrantType:    "client_credentials",
	})
	if err != nil {
		return nil, err
	}

	// Create clients
	return client.NewClient(client.WithAuthToken(results.AccessToken)), nil
}
