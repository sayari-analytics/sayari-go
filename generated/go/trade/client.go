// This file was auto-generated by Fern from our API Definition.

package trade

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
		baseURL:    options.BaseURL,
		httpClient: options.HTTPClient,
		header:     options.ToHeader(),
	}
}

// Search for a shipment
func (c *Client) SearchShipments(ctx context.Context, request *generatedgo.SearchShipments) (*generatedgo.ShipmentSearchResults, error) {
	baseURL := "https://api.sayari.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "v1/trade/search/shipments"

	queryParams := make(url.Values)
	if request.Limit != nil {
		queryParams.Add("limit", fmt.Sprintf("%v", *request.Limit))
	}
	if request.Offset != nil {
		queryParams.Add("offset", fmt.Sprintf("%v", *request.Offset))
	}
	queryParams.Add("q", fmt.Sprintf("%v", request.Q))
	for _, value := range request.Fields {
		queryParams.Add("fields", fmt.Sprintf("%v", *value))
	}
	if request.Facets != nil {
		queryParams.Add("facets", fmt.Sprintf("%v", *request.Facets))
	}
	if request.GeoFacets != nil {
		queryParams.Add("geo_facets", fmt.Sprintf("%v", *request.GeoFacets))
	}
	if request.Advanced != nil {
		queryParams.Add("advanced", fmt.Sprintf("%v", *request.Advanced))
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

	var response *generatedgo.ShipmentSearchResults
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
