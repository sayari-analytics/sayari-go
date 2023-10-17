// This file was auto-generated by Fern from our API Definition.

package entity

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

// Retrieve an entity from the database based on the ID
func (c *Client) GetEntity(ctx context.Context, id generatedgo.EntityId, request *generatedgo.GetEntity) (*generatedgo.EntityDetails, error) {
	baseURL := "https://api.sayari.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"v1/entity/%v", id)

	queryParams := make(url.Values)
	if request.AttributesNameNext != nil {
		queryParams.Add("attributes.name.next", fmt.Sprintf("%v", *request.AttributesNameNext))
	}
	if request.AttributesNamePrev != nil {
		queryParams.Add("attributes.name.prev", fmt.Sprintf("%v", *request.AttributesNamePrev))
	}
	if request.AttributesNameLimit != nil {
		queryParams.Add("attributes.name.limit", fmt.Sprintf("%v", *request.AttributesNameLimit))
	}
	if request.AttributesAddressNext != nil {
		queryParams.Add("attributes.address.next", fmt.Sprintf("%v", *request.AttributesAddressNext))
	}
	if request.AttributesAddressPrev != nil {
		queryParams.Add("attributes.address.prev", fmt.Sprintf("%v", *request.AttributesAddressPrev))
	}
	if request.AttributesAddressLimit != nil {
		queryParams.Add("attributes.address.limit", fmt.Sprintf("%v", *request.AttributesAddressLimit))
	}
	if request.AttributesCountryNext != nil {
		queryParams.Add("attributes.country.next", fmt.Sprintf("%v", *request.AttributesCountryNext))
	}
	if request.AttributesCountryPrev != nil {
		queryParams.Add("attributes.country.prev", fmt.Sprintf("%v", *request.AttributesCountryPrev))
	}
	if request.AttributesCountryLimit != nil {
		queryParams.Add("attributes.country.limit", fmt.Sprintf("%v", *request.AttributesCountryLimit))
	}
	if request.RelationshipsNext != nil {
		queryParams.Add("relationships.next", fmt.Sprintf("%v", *request.RelationshipsNext))
	}
	if request.RelationshipsPrev != nil {
		queryParams.Add("relationships.prev", fmt.Sprintf("%v", *request.RelationshipsPrev))
	}
	if request.RelationshipsLimit != nil {
		queryParams.Add("relationships.limit", fmt.Sprintf("%v", *request.RelationshipsLimit))
	}
	if request.RelationshipsType != nil {
		queryParams.Add("relationships.type", fmt.Sprintf("%v", *request.RelationshipsType))
	}
	if request.RelationshipsSort != nil {
		queryParams.Add("relationships.sort", fmt.Sprintf("%v", *request.RelationshipsSort))
	}
	if request.RelationshipsStartDate != nil {
		queryParams.Add("relationships.startDate", fmt.Sprintf("%v", request.RelationshipsStartDate.Format("2006-01-02")))
	}
	if request.RelationshipsEndDate != nil {
		queryParams.Add("relationships.endDate", fmt.Sprintf("%v", request.RelationshipsEndDate.Format("2006-01-02")))
	}
	if request.RelationshipsMinShares != nil {
		queryParams.Add("relationships.minShares", fmt.Sprintf("%v", *request.RelationshipsMinShares))
	}
	for _, value := range request.RelationshipsCountry {
		queryParams.Add("relationships.country", fmt.Sprintf("%v", *value))
	}
	for _, value := range request.RelationshipsArrivalCountry {
		queryParams.Add("relationships.arrivalCountry", fmt.Sprintf("%v", *value))
	}
	for _, value := range request.RelationshipsDepartureCountry {
		queryParams.Add("relationships.departureCountry", fmt.Sprintf("%v", *value))
	}
	if request.RelationshipsHsCode != nil {
		queryParams.Add("relationships.hsCode", fmt.Sprintf("%v", *request.RelationshipsHsCode))
	}
	if request.PossiblySameAsNext != nil {
		queryParams.Add("possibly_same_as.next", fmt.Sprintf("%v", *request.PossiblySameAsNext))
	}
	if request.PossiblySameAsPrev != nil {
		queryParams.Add("possibly_same_as.prev", fmt.Sprintf("%v", *request.PossiblySameAsPrev))
	}
	if request.PossiblySameAsLimit != nil {
		queryParams.Add("possibly_same_as.limit", fmt.Sprintf("%v", *request.PossiblySameAsLimit))
	}
	if request.ReferencedByNext != nil {
		queryParams.Add("referenced_by.next", fmt.Sprintf("%v", *request.ReferencedByNext))
	}
	if request.ReferencedByPrev != nil {
		queryParams.Add("referenced_by.prev", fmt.Sprintf("%v", *request.ReferencedByPrev))
	}
	if request.ReferencedByLimit != nil {
		queryParams.Add("referenced_by.limit", fmt.Sprintf("%v", *request.ReferencedByLimit))
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

	var response *generatedgo.EntityDetails
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

// The Entity Summary endpoint returns a smaller entity payload
func (c *Client) EntitySummary(ctx context.Context, id generatedgo.EntityId) (*generatedgo.EntityDetails, error) {
	baseURL := "https://api.sayari.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"v1/entity_summary/%v", id)

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

	var response *generatedgo.EntityDetails
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
