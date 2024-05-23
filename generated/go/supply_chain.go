// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/sayari-analytics/sayari-go/generated/go/core"
)

type UpstreamTradeTraversalResponse struct {
	Status   *int                               `json:"status,omitempty" url:"status,omitempty"`
	Success  *bool                              `json:"success,omitempty" url:"success,omitempty"`
	Message  *string                            `json:"message,omitempty" url:"message,omitempty"`
	Entities map[EntityId]*TradeTraversalEntity `json:"entities,omitempty" url:"entities,omitempty"`
	Paths    *TradeTraversalPathOrSegment       `json:"paths,omitempty" url:"paths,omitempty"`

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
	// Country leaf node filter. Only return supply chains that end with a supplier in 1+ of the specified countries.
	Countries []Country `json:"-" url:"countries,omitempty"`
	// Country leaf node filter. Only return supply chains that end with a supplier in none of the specified countries.
	NotCountries []Country `json:"-" url:"-countries,omitempty"`
	// Risk leaf node filter. Only return supply chains that end with a supplier that has 1+ of the specified risk factors.
	Risks []Risk `json:"-" url:"risks,omitempty"`
	// Risk leaf node filter. Only return supply chains that end with a supplier that has none of the specified risk factors.
	NotRisk []Risk `json:"-" url:"-risks,omitempty"`
	// Product root edge filter. Only return supply chains that start with an edge that has 1+ of the specified HS codes.
	HsCode []string `json:"-" url:"hs_code,omitempty"`
	// Product root edge filter. Only return supply chains that start with an edge that has none of the specified HS codes.
	NotHsCode []string `json:"-" url:"-hs_code,omitempty"`
	// Component node filter. Only return supply chains that contain at least one edge with 1+ of the specified HS codes.
	Components []string `json:"-" url:"components,omitempty"`
	// Component node filter. Only return supply chains that contain no edges with any of the specified HS codes.
	NotComponents []string `json:"-" url:"-components,omitempty"`
	// The maximum depth of the traversal, from 1 to 4 inclusive. Default is 4. Reduce if query is timing out.
	MaxDepth *int `json:"-" url:"max_depth,omitempty"`
	// The date range to filter the supply chain by by only considering shipments within the specified date range, inclusive. The date range is formatted as "YYYY-MM-DD|YYYY-MM-DD", where the first date is the start date and the second date is the end date. Both dates are optional, e.g. "|2022-01-01" will return all shipments up to and including 2022-01-01.
	Date *string `json:"-" url:"date,omitempty"`
}
