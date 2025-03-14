// This file was auto-generated by Fern from our API Definition.

package resource

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

// Save an entity to a project.
func (c *Client) SaveEntity(
	ctx context.Context,
	request *generatedgo.SaveEntityRequest,
	opts ...option.RequestOption,
) (*generatedgo.SaveEntityResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.sayari.com",
	)
	endpointURL := baseURL + "/v1/resource/entity"
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

	var response *generatedgo.SaveEntityResponse
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

// Deletes an existing saved resource from a project.
func (c *Client) DeleteResource(
	ctx context.Context,
	type_ generatedgo.ResourceType,
	resourceId string,
	opts ...option.RequestOption,
) (*generatedgo.DeleteResourceResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.sayari.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v1/resource/%v/%v",
		type_,
		resourceId,
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

	var response *generatedgo.DeleteResourceResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodDelete,
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
