// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"

	core "github.com/sayari-analytics/sayari-go/generated/go/core"
)

type GetToken struct {
	ClientId     string `json:"client_id" url:"-"`
	ClientSecret string `json:"client_secret" url:"-"`
	audience     string
	grantType    string
}

func (g *GetToken) Audience() string {
	return g.audience
}

func (g *GetToken) GrantType() string {
	return g.grantType
}

func (g *GetToken) UnmarshalJSON(data []byte) error {
	type unmarshaler GetToken
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*g = GetToken(body)
	g.audience = "sayari.com"
	g.grantType = "client_credentials"
	return nil
}

func (g *GetToken) MarshalJSON() ([]byte, error) {
	type embed GetToken
	var marshaler = struct {
		embed
		Audience  string `json:"audience"`
		GrantType string `json:"grant_type"`
	}{
		embed:     embed(*g),
		Audience:  "sayari.com",
		GrantType: "client_credentials",
	}
	return json.Marshal(marshaler)
}

type AuthResponse struct {
	// The bearer token you will pass in to subsequent API calls to authenticate.
	AccessToken string `json:"access_token" url:"access_token"`
	// Tells you how long (in seconds) until your bearer token expires.
	ExpiresIn int `json:"expires_in" url:"expires_in"`
	// Will always be "Bearer"
	TokenType string `json:"token_type" url:"token_type"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (a *AuthResponse) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *AuthResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler AuthResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AuthResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties

	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *AuthResponse) String() string {
	if len(a._rawJSON) > 0 {
		if value, err := core.StringifyJSON(a._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}
