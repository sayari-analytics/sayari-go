// This file was auto-generated by Fern from our API Definition.

package resolution

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	errors "errors"
	fmt "fmt"
	generatedgo "github.com/sayari-analytics/sayari-go/generated/go"
	core "github.com/sayari-analytics/sayari-go/generated/go/core"
	io "io"
	http "net/http"
	url "net/url"
)

type Client struct {
	baseURL     string
	httpClient  core.HTTPClient
	header      http.Header
	rateLimiter *core.RateLimiter
}

func NewClient(opts ...core.ClientOption) *Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &Client{
		baseURL:     options.BaseURL,
		httpClient:  options.HTTPClient,
		header:      options.ToHeader(),
		rateLimiter: options.RateLimiter,
	}
}

// The resolution endpoints allow users to search for matching entities against a provided list of attributes. The endpoint is similar to the search endpoint, except it's tuned to only return the best match so the client doesn't need to do as much or any post-processing work to filter down results.
func (c *Client) Resolution(ctx context.Context, request *generatedgo.Resolution) (*generatedgo.ResolutionResponse, error) {
	baseURL := "https://api.sayari.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "v1/resolution"

	queryParams := make(url.Values)
	for _, value := range request.Name {
		queryParams.Add("name", fmt.Sprintf("%v", *value))
	}
	for _, value := range request.Identifier {
		queryParams.Add("identifier", fmt.Sprintf("%v", *value))
	}
	for _, value := range request.Country {
		queryParams.Add("country", fmt.Sprintf("%v", *value))
	}
	for _, value := range request.Address {
		queryParams.Add("address", fmt.Sprintf("%v", *value))
	}
	for _, value := range request.DateOfBirth {
		queryParams.Add("date_of_birth", fmt.Sprintf("%v", *value))
	}
	for _, value := range request.Contact {
		queryParams.Add("contact", fmt.Sprintf("%v", *value))
	}
	for _, value := range request.Type {
		queryParams.Add("type", fmt.Sprintf("%v", *value))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 404:
			value := new(generatedgo.NotFound)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 429:
			value := new(generatedgo.RateLimitExceeded)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 401:
			value := new(generatedgo.Unauthorized)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *generatedgo.ResolutionResponse
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
		nil,
		&response,
		false,
		c.header,
		errorDecoder,
		c.rateLimiter,
	); err != nil {
		return response, err
	}
	return response, nil
}
