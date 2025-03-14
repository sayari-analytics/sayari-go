// This file was auto-generated by Fern from our API Definition.

package notifications

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

// The Project Notifications endpoint returns a list of notifications on all entities saved to a project.
func (c *Client) ProjectNotifications(
	ctx context.Context,
	// Unique identifier of the project
	id string,
	request *generatedgo.ProjectNotifications,
	opts ...option.RequestOption,
) (*generatedgo.ProjectNotificationsResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.sayari.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v1/notifications/projects/%v",
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
		502: func(apiError *core.APIError) error {
			return &generatedgo.BadGateway{
				APIError: apiError,
			}
		},
		520: func(apiError *core.APIError) error {
			return &generatedgo.ConnectionError{
				APIError: apiError,
			}
		},
	}

	var response *generatedgo.ProjectNotificationsResponse
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

// <Warning>This endpoint is in beta and is subject to change. It is provided for early access and testing purposes only.</Warning> The Resource Notifications endpoint returns a list of notifications for a saved entity.
func (c *Client) ResourceNotifications(
	ctx context.Context,
	// Unique identifier of the resource
	id string,
	request *generatedgo.ResourceNotifications,
	opts ...option.RequestOption,
) (*generatedgo.ResourceNotificationsResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.sayari.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v1/notifications/resources/%v",
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
		502: func(apiError *core.APIError) error {
			return &generatedgo.BadGateway{
				APIError: apiError,
			}
		},
		520: func(apiError *core.APIError) error {
			return &generatedgo.ConnectionError{
				APIError: apiError,
			}
		},
	}

	var response *generatedgo.ResourceNotificationsResponse
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

// Deletes all notifications from a project.
func (c *Client) DeleteProjectNotifications(
	ctx context.Context,
	projectId string,
	opts ...option.RequestOption,
) error {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.sayari.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v1/notifications/projects/%v",
		projectId,
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
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return err
	}
	return nil
}

// Deletes notifications for saved resources of an entity.
func (c *Client) DeleteEntityNotifications(
	ctx context.Context,
	entityId string,
	opts ...option.RequestOption,
) error {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.sayari.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v1/notifications/entity/%v",
		entityId,
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
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return err
	}
	return nil
}

// Deletes notifications for a saved resource.
func (c *Client) DeleteResourceNotifications(
	ctx context.Context,
	resourceId string,
	opts ...option.RequestOption,
) error {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.sayari.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v1/notifications/resources/%v",
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
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return err
	}
	return nil
}
