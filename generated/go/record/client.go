// This file was auto-generated by Fern from our API Definition.

package record

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
	baseURL    string
	httpClient core.HTTPClient
	header     http.Header
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

// Retrieve a record from the database based on the ID
func (c *Client) GetRecord(ctx context.Context, id generatedgo.RecordId, request *generatedgo.GetRecord) (*generatedgo.RecordDetails, error) {
	baseURL := "https://api.sayari.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"v1/record/%v", id)

	queryParams := make(url.Values)
	if request.ReferencesLimit != nil {
		queryParams.Add("references.limit", fmt.Sprintf("%v", *request.ReferencesLimit))
	}
	if request.ReferencesOffset != nil {
		queryParams.Add("references.offset", fmt.Sprintf("%v", *request.ReferencesOffset))
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
			value := new(generatedgo.RatLimitExceeded)
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

	var response *generatedgo.RecordDetails
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
	); err != nil {
		return response, err
	}
	return response, nil
}
