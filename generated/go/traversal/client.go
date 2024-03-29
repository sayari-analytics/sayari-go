// This file was auto-generated by Fern from our API Definition.

package traversal

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	errors "errors"
	fmt "fmt"
	generatedgo "github.com/sayari-analytics/sayari-go/generated/go"
	core "github.com/sayari-analytics/sayari-go/generated/go/core"
	option "github.com/sayari-analytics/sayari-go/generated/go/option"
	io "io"
	http "net/http"
	url "net/url"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	return &Client{
		baseURL: options.BaseURL,
		caller: core.NewCaller(
			&core.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
			options.RateLimiter,
		),
		header: options.ToHeader(),
	}
}

// The Traversal endpoint returns paths from a single target entity to up to 50 directly or indirectly-related entities. Each path includes information on the 0 to 10 intermediary entities, as well as their connecting relationships. The response's explored_count field indicates the size of the graph subset the application searched. Running a traversal on a highly connected entity with a restrictive set of argument filters and a high max depth will require the application to explore a higher number of traversal paths, which may affect performance. In cases where a traversal searches over a very large, highly-connected subgraph, a partial result set may be returned containing only the most relevant results. This will be indicated in the response by the partial_results field.
func (c *Client) Traversal(
	ctx context.Context,
	// Unique identifier of the entity
	id string,
	request *generatedgo.Traversal,
	opts ...option.RequestOption,
) (*generatedgo.TraversalResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.sayari.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"v1/traversal/%v", id)

	queryParams := make(url.Values)
	if request.Limit != nil {
		queryParams.Add("limit", fmt.Sprintf("%v", *request.Limit))
	}
	if request.Offset != nil {
		queryParams.Add("offset", fmt.Sprintf("%v", *request.Offset))
	}
	if request.MinDepth != nil {
		queryParams.Add("min_depth", fmt.Sprintf("%v", *request.MinDepth))
	}
	if request.MaxDepth != nil {
		queryParams.Add("max_depth", fmt.Sprintf("%v", *request.MaxDepth))
	}
	for _, value := range request.Relationships {
		queryParams.Add("relationships", fmt.Sprintf("%v", *value))
	}
	if request.Psa != nil {
		queryParams.Add("psa", fmt.Sprintf("%v", *request.Psa))
	}
	for _, value := range request.Countries {
		queryParams.Add("countries", fmt.Sprintf("%v", *value))
	}
	for _, value := range request.Types {
		queryParams.Add("types", fmt.Sprintf("%v", *value))
	}
	if request.Sanctioned != nil {
		queryParams.Add("sanctioned", fmt.Sprintf("%v", *request.Sanctioned))
	}
	if request.Pep != nil {
		queryParams.Add("pep", fmt.Sprintf("%v", *request.Pep))
	}
	if request.MinShares != nil {
		queryParams.Add("min_shares", fmt.Sprintf("%v", *request.MinShares))
	}
	if request.IncludeUnknownShares != nil {
		queryParams.Add("include_unknown_shares", fmt.Sprintf("%v", *request.IncludeUnknownShares))
	}
	if request.ExcludeFormerRelationships != nil {
		queryParams.Add("exclude_former_relationships", fmt.Sprintf("%v", *request.ExcludeFormerRelationships))
	}
	if request.ExcludeClosedEntities != nil {
		queryParams.Add("exclude_closed_entities", fmt.Sprintf("%v", *request.ExcludeClosedEntities))
	}
	if request.EuHighRiskThird != nil {
		queryParams.Add("eu_high_risk_third", fmt.Sprintf("%v", *request.EuHighRiskThird))
	}
	if request.ReputationalRiskModernSlavery != nil {
		queryParams.Add("reputational_risk_modern_slavery", fmt.Sprintf("%v", *request.ReputationalRiskModernSlavery))
	}
	if request.StateOwned != nil {
		queryParams.Add("state_owned", fmt.Sprintf("%v", *request.StateOwned))
	}
	if request.FormerlySanctioned != nil {
		queryParams.Add("formerly_sanctioned", fmt.Sprintf("%v", *request.FormerlySanctioned))
	}
	if request.ReputationalRiskTerrorism != nil {
		queryParams.Add("reputational_risk_terrorism", fmt.Sprintf("%v", *request.ReputationalRiskTerrorism))
	}
	if request.ReputationalRiskOrganizedCrime != nil {
		queryParams.Add("reputational_risk_organized_crime", fmt.Sprintf("%v", *request.ReputationalRiskOrganizedCrime))
	}
	if request.ReputationalRiskFinancialCrime != nil {
		queryParams.Add("reputational_risk_financial_crime", fmt.Sprintf("%v", *request.ReputationalRiskFinancialCrime))
	}
	if request.ReputationalRiskBriberyAndCorruption != nil {
		queryParams.Add("reputational_risk_bribery_and_corruption", fmt.Sprintf("%v", *request.ReputationalRiskBriberyAndCorruption))
	}
	if request.ReputationalRiskOther != nil {
		queryParams.Add("reputational_risk_other", fmt.Sprintf("%v", *request.ReputationalRiskOther))
	}
	if request.ReputationalRiskCybercrime != nil {
		queryParams.Add("reputational_risk_cybercrime", fmt.Sprintf("%v", *request.ReputationalRiskCybercrime))
	}
	if request.RegulatoryAction != nil {
		queryParams.Add("regulatory_action", fmt.Sprintf("%v", *request.RegulatoryAction))
	}
	if request.LawEnforcementAction != nil {
		queryParams.Add("law_enforcement_action", fmt.Sprintf("%v", *request.LawEnforcementAction))
	}
	if request.XinjiangGeospatial != nil {
		queryParams.Add("xinjiang_geospatial", fmt.Sprintf("%v", *request.XinjiangGeospatial))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(generatedgo.BadRequest)
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
		case 404:
			value := new(generatedgo.NotFound)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 405:
			value := new(generatedgo.MethodNotAllowed)
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
		case 500:
			value := new(generatedgo.InternalServerError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 502:
			value := new(generatedgo.BadGateway)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 520:
			value := new(generatedgo.ConnectionError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *generatedgo.TraversalResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodGet,
			MaxAttempts:  options.MaxAttempts,
			Headers:      headers,
			Client:       options.HTTPClient,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// The UBO endpoint returns paths from a single target entity to up to 50 beneficial owners. The endpoint is a shorthand for the equivalent traversal query.
func (c *Client) Ubo(
	ctx context.Context,
	// Unique identifier of the entity
	id string,
	request *generatedgo.Ubo,
	opts ...option.RequestOption,
) (*generatedgo.TraversalResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.sayari.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"v1/ubo/%v", id)

	queryParams := make(url.Values)
	if request.Limit != nil {
		queryParams.Add("limit", fmt.Sprintf("%v", *request.Limit))
	}
	if request.Offset != nil {
		queryParams.Add("offset", fmt.Sprintf("%v", *request.Offset))
	}
	if request.MinDepth != nil {
		queryParams.Add("min_depth", fmt.Sprintf("%v", *request.MinDepth))
	}
	if request.MaxDepth != nil {
		queryParams.Add("max_depth", fmt.Sprintf("%v", *request.MaxDepth))
	}
	if request.Psa != nil {
		queryParams.Add("psa", fmt.Sprintf("%v", *request.Psa))
	}
	for _, value := range request.Countries {
		queryParams.Add("countries", fmt.Sprintf("%v", *value))
	}
	for _, value := range request.Types {
		queryParams.Add("types", fmt.Sprintf("%v", *value))
	}
	if request.Sanctioned != nil {
		queryParams.Add("sanctioned", fmt.Sprintf("%v", *request.Sanctioned))
	}
	if request.Pep != nil {
		queryParams.Add("pep", fmt.Sprintf("%v", *request.Pep))
	}
	if request.MinShares != nil {
		queryParams.Add("min_shares", fmt.Sprintf("%v", *request.MinShares))
	}
	if request.IncludeUnknownShares != nil {
		queryParams.Add("include_unknown_shares", fmt.Sprintf("%v", *request.IncludeUnknownShares))
	}
	if request.ExcludeFormerRelationships != nil {
		queryParams.Add("exclude_former_relationships", fmt.Sprintf("%v", *request.ExcludeFormerRelationships))
	}
	if request.ExcludeClosedEntities != nil {
		queryParams.Add("exclude_closed_entities", fmt.Sprintf("%v", *request.ExcludeClosedEntities))
	}
	if request.EuHighRiskThird != nil {
		queryParams.Add("eu_high_risk_third", fmt.Sprintf("%v", *request.EuHighRiskThird))
	}
	if request.ReputationalRiskModernSlavery != nil {
		queryParams.Add("reputational_risk_modern_slavery", fmt.Sprintf("%v", *request.ReputationalRiskModernSlavery))
	}
	if request.StateOwned != nil {
		queryParams.Add("state_owned", fmt.Sprintf("%v", *request.StateOwned))
	}
	if request.FormerlySanctioned != nil {
		queryParams.Add("formerly_sanctioned", fmt.Sprintf("%v", *request.FormerlySanctioned))
	}
	if request.ReputationalRiskTerrorism != nil {
		queryParams.Add("reputational_risk_terrorism", fmt.Sprintf("%v", *request.ReputationalRiskTerrorism))
	}
	if request.ReputationalRiskOrganizedCrime != nil {
		queryParams.Add("reputational_risk_organized_crime", fmt.Sprintf("%v", *request.ReputationalRiskOrganizedCrime))
	}
	if request.ReputationalRiskFinancialCrime != nil {
		queryParams.Add("reputational_risk_financial_crime", fmt.Sprintf("%v", *request.ReputationalRiskFinancialCrime))
	}
	if request.ReputationalRiskBriberyAndCorruption != nil {
		queryParams.Add("reputational_risk_bribery_and_corruption", fmt.Sprintf("%v", *request.ReputationalRiskBriberyAndCorruption))
	}
	if request.ReputationalRiskOther != nil {
		queryParams.Add("reputational_risk_other", fmt.Sprintf("%v", *request.ReputationalRiskOther))
	}
	if request.ReputationalRiskCybercrime != nil {
		queryParams.Add("reputational_risk_cybercrime", fmt.Sprintf("%v", *request.ReputationalRiskCybercrime))
	}
	if request.RegulatoryAction != nil {
		queryParams.Add("regulatory_action", fmt.Sprintf("%v", *request.RegulatoryAction))
	}
	if request.LawEnforcementAction != nil {
		queryParams.Add("law_enforcement_action", fmt.Sprintf("%v", *request.LawEnforcementAction))
	}
	if request.XinjiangGeospatial != nil {
		queryParams.Add("xinjiang_geospatial", fmt.Sprintf("%v", *request.XinjiangGeospatial))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(generatedgo.BadRequest)
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
		case 404:
			value := new(generatedgo.NotFound)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 405:
			value := new(generatedgo.MethodNotAllowed)
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
		case 500:
			value := new(generatedgo.InternalServerError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 502:
			value := new(generatedgo.BadGateway)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 520:
			value := new(generatedgo.ConnectionError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *generatedgo.TraversalResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodGet,
			MaxAttempts:  options.MaxAttempts,
			Headers:      headers,
			Client:       options.HTTPClient,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// The Ownership endpoint returns paths from a single target entity to up to 50 entities directly or indirectly owned by that entity. The endpoint is a shorthand for the equivalent traversal query.
func (c *Client) Ownership(
	ctx context.Context,
	// Unique identifier of the entity
	id string,
	request *generatedgo.Ownership,
	opts ...option.RequestOption,
) (*generatedgo.TraversalResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.sayari.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"v1/downstream/%v", id)

	queryParams := make(url.Values)
	if request.Limit != nil {
		queryParams.Add("limit", fmt.Sprintf("%v", *request.Limit))
	}
	if request.Offset != nil {
		queryParams.Add("offset", fmt.Sprintf("%v", *request.Offset))
	}
	if request.MinDepth != nil {
		queryParams.Add("min_depth", fmt.Sprintf("%v", *request.MinDepth))
	}
	if request.MaxDepth != nil {
		queryParams.Add("max_depth", fmt.Sprintf("%v", *request.MaxDepth))
	}
	if request.Psa != nil {
		queryParams.Add("psa", fmt.Sprintf("%v", *request.Psa))
	}
	for _, value := range request.Countries {
		queryParams.Add("countries", fmt.Sprintf("%v", *value))
	}
	for _, value := range request.Types {
		queryParams.Add("types", fmt.Sprintf("%v", *value))
	}
	if request.Sanctioned != nil {
		queryParams.Add("sanctioned", fmt.Sprintf("%v", *request.Sanctioned))
	}
	if request.Pep != nil {
		queryParams.Add("pep", fmt.Sprintf("%v", *request.Pep))
	}
	if request.MinShares != nil {
		queryParams.Add("min_shares", fmt.Sprintf("%v", *request.MinShares))
	}
	if request.IncludeUnknownShares != nil {
		queryParams.Add("include_unknown_shares", fmt.Sprintf("%v", *request.IncludeUnknownShares))
	}
	if request.ExcludeFormerRelationships != nil {
		queryParams.Add("exclude_former_relationships", fmt.Sprintf("%v", *request.ExcludeFormerRelationships))
	}
	if request.ExcludeClosedEntities != nil {
		queryParams.Add("exclude_closed_entities", fmt.Sprintf("%v", *request.ExcludeClosedEntities))
	}
	if request.EuHighRiskThird != nil {
		queryParams.Add("eu_high_risk_third", fmt.Sprintf("%v", *request.EuHighRiskThird))
	}
	if request.ReputationalRiskModernSlavery != nil {
		queryParams.Add("reputational_risk_modern_slavery", fmt.Sprintf("%v", *request.ReputationalRiskModernSlavery))
	}
	if request.StateOwned != nil {
		queryParams.Add("state_owned", fmt.Sprintf("%v", *request.StateOwned))
	}
	if request.FormerlySanctioned != nil {
		queryParams.Add("formerly_sanctioned", fmt.Sprintf("%v", *request.FormerlySanctioned))
	}
	if request.ReputationalRiskTerrorism != nil {
		queryParams.Add("reputational_risk_terrorism", fmt.Sprintf("%v", *request.ReputationalRiskTerrorism))
	}
	if request.ReputationalRiskOrganizedCrime != nil {
		queryParams.Add("reputational_risk_organized_crime", fmt.Sprintf("%v", *request.ReputationalRiskOrganizedCrime))
	}
	if request.ReputationalRiskFinancialCrime != nil {
		queryParams.Add("reputational_risk_financial_crime", fmt.Sprintf("%v", *request.ReputationalRiskFinancialCrime))
	}
	if request.ReputationalRiskBriberyAndCorruption != nil {
		queryParams.Add("reputational_risk_bribery_and_corruption", fmt.Sprintf("%v", *request.ReputationalRiskBriberyAndCorruption))
	}
	if request.ReputationalRiskOther != nil {
		queryParams.Add("reputational_risk_other", fmt.Sprintf("%v", *request.ReputationalRiskOther))
	}
	if request.ReputationalRiskCybercrime != nil {
		queryParams.Add("reputational_risk_cybercrime", fmt.Sprintf("%v", *request.ReputationalRiskCybercrime))
	}
	if request.RegulatoryAction != nil {
		queryParams.Add("regulatory_action", fmt.Sprintf("%v", *request.RegulatoryAction))
	}
	if request.LawEnforcementAction != nil {
		queryParams.Add("law_enforcement_action", fmt.Sprintf("%v", *request.LawEnforcementAction))
	}
	if request.XinjiangGeospatial != nil {
		queryParams.Add("xinjiang_geospatial", fmt.Sprintf("%v", *request.XinjiangGeospatial))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(generatedgo.BadRequest)
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
		case 404:
			value := new(generatedgo.NotFound)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 405:
			value := new(generatedgo.MethodNotAllowed)
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
		case 500:
			value := new(generatedgo.InternalServerError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 502:
			value := new(generatedgo.BadGateway)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 520:
			value := new(generatedgo.ConnectionError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *generatedgo.TraversalResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodGet,
			MaxAttempts:  options.MaxAttempts,
			Headers:      headers,
			Client:       options.HTTPClient,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// The Watchlist endpoint returns paths from a single target entity to up to 50 other entities that appear on a watchlist or are peps. The endpoint is a shorthand for the equivalent traversal query.
func (c *Client) Watchlist(
	ctx context.Context,
	// Unique identifier of the entity
	id string,
	request *generatedgo.Watchlist,
	opts ...option.RequestOption,
) (*generatedgo.TraversalResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.sayari.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"v1/watchlist/%v", id)

	queryParams := make(url.Values)
	if request.Limit != nil {
		queryParams.Add("limit", fmt.Sprintf("%v", *request.Limit))
	}
	if request.Offset != nil {
		queryParams.Add("offset", fmt.Sprintf("%v", *request.Offset))
	}
	if request.MinDepth != nil {
		queryParams.Add("min_depth", fmt.Sprintf("%v", *request.MinDepth))
	}
	if request.MaxDepth != nil {
		queryParams.Add("max_depth", fmt.Sprintf("%v", *request.MaxDepth))
	}
	for _, value := range request.Relationships {
		queryParams.Add("relationships", fmt.Sprintf("%v", *value))
	}
	if request.Psa != nil {
		queryParams.Add("psa", fmt.Sprintf("%v", *request.Psa))
	}
	for _, value := range request.Countries {
		queryParams.Add("countries", fmt.Sprintf("%v", *value))
	}
	for _, value := range request.Types {
		queryParams.Add("types", fmt.Sprintf("%v", *value))
	}
	if request.Sanctioned != nil {
		queryParams.Add("sanctioned", fmt.Sprintf("%v", *request.Sanctioned))
	}
	if request.Pep != nil {
		queryParams.Add("pep", fmt.Sprintf("%v", *request.Pep))
	}
	if request.MinShares != nil {
		queryParams.Add("min_shares", fmt.Sprintf("%v", *request.MinShares))
	}
	if request.IncludeUnknownShares != nil {
		queryParams.Add("include_unknown_shares", fmt.Sprintf("%v", *request.IncludeUnknownShares))
	}
	if request.ExcludeFormerRelationships != nil {
		queryParams.Add("exclude_former_relationships", fmt.Sprintf("%v", *request.ExcludeFormerRelationships))
	}
	if request.ExcludeClosedEntities != nil {
		queryParams.Add("exclude_closed_entities", fmt.Sprintf("%v", *request.ExcludeClosedEntities))
	}
	if request.EuHighRiskThird != nil {
		queryParams.Add("eu_high_risk_third", fmt.Sprintf("%v", *request.EuHighRiskThird))
	}
	if request.ReputationalRiskModernSlavery != nil {
		queryParams.Add("reputational_risk_modern_slavery", fmt.Sprintf("%v", *request.ReputationalRiskModernSlavery))
	}
	if request.StateOwned != nil {
		queryParams.Add("state_owned", fmt.Sprintf("%v", *request.StateOwned))
	}
	if request.FormerlySanctioned != nil {
		queryParams.Add("formerly_sanctioned", fmt.Sprintf("%v", *request.FormerlySanctioned))
	}
	if request.ReputationalRiskTerrorism != nil {
		queryParams.Add("reputational_risk_terrorism", fmt.Sprintf("%v", *request.ReputationalRiskTerrorism))
	}
	if request.ReputationalRiskOrganizedCrime != nil {
		queryParams.Add("reputational_risk_organized_crime", fmt.Sprintf("%v", *request.ReputationalRiskOrganizedCrime))
	}
	if request.ReputationalRiskFinancialCrime != nil {
		queryParams.Add("reputational_risk_financial_crime", fmt.Sprintf("%v", *request.ReputationalRiskFinancialCrime))
	}
	if request.ReputationalRiskBriberyAndCorruption != nil {
		queryParams.Add("reputational_risk_bribery_and_corruption", fmt.Sprintf("%v", *request.ReputationalRiskBriberyAndCorruption))
	}
	if request.ReputationalRiskOther != nil {
		queryParams.Add("reputational_risk_other", fmt.Sprintf("%v", *request.ReputationalRiskOther))
	}
	if request.ReputationalRiskCybercrime != nil {
		queryParams.Add("reputational_risk_cybercrime", fmt.Sprintf("%v", *request.ReputationalRiskCybercrime))
	}
	if request.RegulatoryAction != nil {
		queryParams.Add("regulatory_action", fmt.Sprintf("%v", *request.RegulatoryAction))
	}
	if request.LawEnforcementAction != nil {
		queryParams.Add("law_enforcement_action", fmt.Sprintf("%v", *request.LawEnforcementAction))
	}
	if request.XinjiangGeospatial != nil {
		queryParams.Add("xinjiang_geospatial", fmt.Sprintf("%v", *request.XinjiangGeospatial))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(generatedgo.BadRequest)
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
		case 404:
			value := new(generatedgo.NotFound)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 405:
			value := new(generatedgo.MethodNotAllowed)
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
		case 500:
			value := new(generatedgo.InternalServerError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 502:
			value := new(generatedgo.BadGateway)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 520:
			value := new(generatedgo.ConnectionError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *generatedgo.TraversalResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodGet,
			MaxAttempts:  options.MaxAttempts,
			Headers:      headers,
			Client:       options.HTTPClient,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// The Shortest Path endpoint returns a response identifying the shortest traversal path connecting each pair of entities.
func (c *Client) ShortestPath(
	ctx context.Context,
	request *generatedgo.ShortestPath,
	opts ...option.RequestOption,
) (*generatedgo.ShortestPathResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.sayari.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/" + "v1/shortest_path"

	queryParams := make(url.Values)
	for _, value := range request.Entities {
		queryParams.Add("entities", fmt.Sprintf("%v", value))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(generatedgo.BadRequest)
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
		case 404:
			value := new(generatedgo.NotFound)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 405:
			value := new(generatedgo.MethodNotAllowed)
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
		case 500:
			value := new(generatedgo.InternalServerError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 502:
			value := new(generatedgo.BadGateway)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 520:
			value := new(generatedgo.ConnectionError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *generatedgo.ShortestPathResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodGet,
			MaxAttempts:  options.MaxAttempts,
			Headers:      headers,
			Client:       options.HTTPClient,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
