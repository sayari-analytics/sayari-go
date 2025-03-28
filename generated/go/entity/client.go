// This file was auto-generated by Fern from our API Definition.

package entity

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

// <Note>To retrieve a L1 Due Dilligence PDF Report. Include 'Accept: application/pdf' in request headers.</Note> Retrieve an entity profile from the database based on the entity ID. This endpoint returns the full profile, entity_summary returns the same payload minus relationships.
func (c *Client) GetEntity(
	ctx context.Context,
	// Unique identifier of the entity
	id string,
	request *generatedgo.GetEntity,
	opts ...option.RequestOption,
) (*generatedgo.GetEntityResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.sayari.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v1/entity/%v",
		id,
	)
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
		404: func(apiError *core.APIError) error {
			return &generatedgo.NotFound{
				APIError: apiError,
			}
		},
		405: func(apiError *core.APIError) error {
			return &generatedgo.MethodNotAllowed{
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

	var response *generatedgo.GetEntityResponse
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

// The Entity Summary endpoint returns a similar payload, minus relationships.
func (c *Client) EntitySummary(
	ctx context.Context,
	// Unique identifier of the entity
	id string,
	opts ...option.RequestOption,
) (*generatedgo.EntitySummaryResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.sayari.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v1/entity_summary/%v",
		id,
	)
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
		404: func(apiError *core.APIError) error {
			return &generatedgo.NotFound{
				APIError: apiError,
			}
		},
		405: func(apiError *core.APIError) error {
			return &generatedgo.MethodNotAllowed{
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

	var response *generatedgo.EntitySummaryResponse
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
