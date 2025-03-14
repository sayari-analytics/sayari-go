// This file was auto-generated by Fern from our API Definition.

package project

import (
	context "context"
	fmt "fmt"
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

// Create a new project
func (c *Client) CreateProject(
	ctx context.Context,
	request *generatedgo.CreateProjectRequest,
	opts ...option.RequestOption,
) (*generatedgo.CreateProjectResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.sayari.com",
	)
	endpointURL := baseURL + "/v1/projects"
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

	var response *generatedgo.CreateProjectResponse
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

// Retrieve a list of projects including upload progress info.
func (c *Client) GetProjects(
	ctx context.Context,
	request *generatedgo.GetProjects,
	opts ...option.RequestOption,
) (*generatedgo.GetProjectsResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.sayari.com",
	)
	endpointURL := baseURL + "/v1/projects"
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

	var response *generatedgo.GetProjectsResponse
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

// Retrieve a list of entities in a project.
func (c *Client) GetProjectEntities(
	ctx context.Context,
	// The project identifier.
	id string,
	request *generatedgo.GetProjectEntities,
	opts ...option.RequestOption,
) (*generatedgo.GetProjectEntitiesResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.sayari.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v1/projects/%v/contents/entity",
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
	headers.Add("Accept", fmt.Sprintf("%v", request.Accept))
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

	var response *generatedgo.GetProjectEntitiesResponse
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

// Deletes an existing project.
func (c *Client) DeleteProject(
	ctx context.Context,
	projectId string,
	opts ...option.RequestOption,
) (*generatedgo.DeleteProjectResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.sayari.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v1/projects/%v",
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

	var response *generatedgo.DeleteProjectResponse
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
