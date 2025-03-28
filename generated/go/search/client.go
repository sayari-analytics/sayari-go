// This file was auto-generated by Fern from our API Definition.

package search

import (
	context "context"
	generatedgo "github.com/sayari-analytics/sayari-go/generated/go"
	core "github.com/sayari-analytics/sayari-go/generated/go/core"
	internal "github.com/sayari-analytics/sayari-go/generated/go/internal"
	option "github.com/sayari-analytics/sayari-go/generated/go/option"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *internal.Caller
	header  http.Header
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	return &Client{
		baseURL: options.BaseURL,
		caller: internal.NewCaller(
			&internal.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header: options.ToHeader(),
	}
}

// Search for an entity. Please note, searches are limited to a maximum of 10,000 results.
func (c *Client) SearchEntity(
	ctx context.Context,
	request *generatedgo.SearchEntity,
	opts ...option.RequestOption,
) (*generatedgo.EntitySearchResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.sayari.com",
	)
	endpointURL := baseURL + "/v1/search/entity"
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &generatedgo.BadRequest{
				APIError: apiError,
			}
		},
		401: func(apiError *core.APIError) error {
			return &generatedgo.Unauthorized{
				APIError: apiError,
			}
		},
		405: func(apiError *core.APIError) error {
			return &generatedgo.MethodNotAllowed{
				APIError: apiError,
			}
		},
		406: func(apiError *core.APIError) error {
			return &generatedgo.NotAcceptable{
				APIError: apiError,
			}
		},
		429: func(apiError *core.APIError) error {
			return &generatedgo.RateLimitExceeded{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &generatedgo.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response *generatedgo.EntitySearchResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Search for an entity. Please note, searches are limited to a maximum of 10,000 results.
func (c *Client) SearchEntityGet(
	ctx context.Context,
	request *generatedgo.SearchEntityGet,
	opts ...option.RequestOption,
) (*generatedgo.EntitySearchResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.sayari.com",
	)
	endpointURL := baseURL + "/v1/search/entity"
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &generatedgo.BadRequest{
				APIError: apiError,
			}
		},
		401: func(apiError *core.APIError) error {
			return &generatedgo.Unauthorized{
				APIError: apiError,
			}
		},
		405: func(apiError *core.APIError) error {
			return &generatedgo.MethodNotAllowed{
				APIError: apiError,
			}
		},
		406: func(apiError *core.APIError) error {
			return &generatedgo.NotAcceptable{
				APIError: apiError,
			}
		},
		429: func(apiError *core.APIError) error {
			return &generatedgo.RateLimitExceeded{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &generatedgo.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response *generatedgo.EntitySearchResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Search for a record. Please note, searches are limited to a maximum of 10,000 results.
func (c *Client) SearchRecord(
	ctx context.Context,
	request *generatedgo.SearchRecord,
	opts ...option.RequestOption,
) (*generatedgo.RecordSearchResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.sayari.com",
	)
	endpointURL := baseURL + "/v1/search/record"
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &generatedgo.BadRequest{
				APIError: apiError,
			}
		},
		401: func(apiError *core.APIError) error {
			return &generatedgo.Unauthorized{
				APIError: apiError,
			}
		},
		405: func(apiError *core.APIError) error {
			return &generatedgo.MethodNotAllowed{
				APIError: apiError,
			}
		},
		406: func(apiError *core.APIError) error {
			return &generatedgo.NotAcceptable{
				APIError: apiError,
			}
		},
		429: func(apiError *core.APIError) error {
			return &generatedgo.RateLimitExceeded{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &generatedgo.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response *generatedgo.RecordSearchResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Search for a record. Please note, searches are limited to a maximum of 10,000 results.
func (c *Client) SearchRecordGet(
	ctx context.Context,
	request *generatedgo.SearchRecordGet,
	opts ...option.RequestOption,
) (*generatedgo.RecordSearchResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.sayari.com",
	)
	endpointURL := baseURL + "/v1/search/record"
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &generatedgo.BadRequest{
				APIError: apiError,
			}
		},
		401: func(apiError *core.APIError) error {
			return &generatedgo.Unauthorized{
				APIError: apiError,
			}
		},
		405: func(apiError *core.APIError) error {
			return &generatedgo.MethodNotAllowed{
				APIError: apiError,
			}
		},
		406: func(apiError *core.APIError) error {
			return &generatedgo.NotAcceptable{
				APIError: apiError,
			}
		},
		429: func(apiError *core.APIError) error {
			return &generatedgo.RateLimitExceeded{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &generatedgo.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response *generatedgo.RecordSearchResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
