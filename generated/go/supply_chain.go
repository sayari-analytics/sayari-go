// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/sayari-analytics/sayari-go/generated/go/core"
)

type UpstreamTradeTraversalResponse struct {
	Status  *int                  `json:"status,omitempty" url:"status,omitempty"`
	Success *bool                 `json:"success,omitempty" url:"success,omitempty"`
	Message *string               `json:"message,omitempty" url:"message,omitempty"`
	Data    []*TradeTraversalPath `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (u *UpstreamTradeTraversalResponse) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UpstreamTradeTraversalResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UpstreamTradeTraversalResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpstreamTradeTraversalResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties

	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UpstreamTradeTraversalResponse) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UpstreamTradeTraversalRequest struct {
	// Risk leaf node filter. Only return supply chains that end with a supplier that has 1+ of the specified [risk factors](/sayari-library/ontology/risk-factors).
	Risk []Risk `json:"-" url:"risk,omitempty"`
	// Risk leaf node filter. Only return supply chains that end with a supplier that has none of the specified [risk factors](/sayari-library/ontology/risk-factors).
	NotRisk []Risk `json:"-" url:"-risk,omitempty"`
	// Country leaf node filter. Only return supply chains that end with a supplier in 1+ of the specified countries.
	Countries []Country `json:"-" url:"countries,omitempty"`
	// Country leaf node filter. Only return supply chains that end with a supplier in none of the specified countries.
	NotCountries []Country `json:"-" url:"-countries,omitempty"`
	// Product root edge filter. Only return supply chains that start with an edge that has 1+ of the specified HS codes.
	Product []string `json:"-" url:"product,omitempty"`
	// Product root edge filter. Only return supply chains that start with an edge that has none of the specified HS codes.
	NotProduct []string `json:"-" url:"-product,omitempty"`
	// Component node filter. Only return supply chains that contain at least one edge with 1+ of the specified HS codes.
	Component []string `json:"-" url:"component,omitempty"`
	// Component node filter. Only return supply chains that contain no edges with any of the specified HS codes.
	NotComponent []string `json:"-" url:"-component,omitempty"`
	// Minimum date edge filter. Only return supply chains with edge dates that are greater than or equal to this date.
	MinDate *string `json:"-" url:"min_date,omitempty"`
	// Maximum date edge filter. Only return supply chains with edge dates that are less than or equal to this date.
	MaxDate *string `json:"-" url:"max_date,omitempty"`
	// The maximum depth of the traversal, from 1 to 4 inclusive. Default is 4. Reduce if query is timing out.
	MaxDepth *int `json:"-" url:"max_depth,omitempty"`
	// The maximum number of results to return. Default and maximum values are 25,000.
	Limit *int `json:"-" url:"limit,omitempty"`
}
