// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/sayari-analytics/sayari-go/generated/go/core"
)

type SearchEntity struct {
	// A limit on the number of objects to be returned with a range between 1 and 100. Defaults to 100.
	Limit *int `json:"-"`
	// Number of results to skip before returning response. Defaults to 0.
	Offset *int `json:"-"`
	// Query term. The syntax for the query parameter follows elasticsearch simple query string syntax. The includes the ability to use search operators and to perform nested queries. Must be url encoded.
	Q string `json:"q"`
	// Record or entity fields to search against.
	Fields []SearchField `json:"fields,omitempty"`
	// Filters to be applied to search query to limit the result-set.
	Filter *FilterList `json:"filter,omitempty"`
	// Whether or not to return search facets in results giving counts by field. Defaults to false.
	Facets *bool `json:"facets,omitempty"`
	// Whether or not to return search geo bound facets in results giving counts by geo tile. Defaults to false.
	GeoFacets *bool `json:"geo_facets,omitempty"`
	// Set to true to enable full elasticsearch query string syntax which allows for fielded search and more complex operators. Note that the syntax is more strict and can result in empty result-sets. Defaults to false.
	Advanced *bool `json:"advanced,omitempty"`
}

type SearchEntityGet struct {
	// A limit on the number of objects to be returned with a range between 1 and 100. Defaults to 100.
	Limit *int `json:"-"`
	// Number of results to skip before returning response. Defaults to 0.
	Offset *int `json:"-"`
	// Query term. The syntax for the query parameter follows elasticsearch simple query string syntax. The includes the ability to use search operators and to perform nested queries. Must be url encoded.
	Q string `json:"-"`
	// Record or entity fields to search against.
	Fields []*SearchField `json:"-"`
	// Whether or not to return search facets in results giving counts by field. Defaults to false.
	Facets *bool `json:"-"`
	// Whether or not to return search geo bound facets in results giving counts by geo tile. Defaults to false.
	GeoFacets *bool `json:"-"`
	// Set to true to enable full elasticsearch query string syntax which allows for fielded search and more complex operators. Note that the syntax is more strict and can result in empty result-sets. Defaults to false.
	Advanced *bool `json:"-"`
}

type SearchRecord struct {
	// A limit on the number of objects to be returned with a range between 1 and 100. Defaults to 100.
	Limit *int `json:"-"`
	// Number of results to skip before returning response. Defaults to 0.
	Offset *int `json:"-"`
	// Query term. The syntax for the query parameter follows elasticsearch simple query string syntax. The includes the ability to use search operators and to perform nested queries. Must be url encoded.
	Q string `json:"q"`
	// Record or entity fields to search against.
	Fields []SearchField `json:"fields,omitempty"`
	// Filters to be applied to search query to limit the result-set.
	Filter *FilterList `json:"filter,omitempty"`
	// Whether or not to return search facets in results giving counts by field. Defaults to false.
	Facets *bool `json:"facets,omitempty"`
	// Set to true to enable full elasticsearch query string syntax which allows for fielded search and more complex operators. Note that the syntax is more strict and can result in empty result-sets. Defaults to false.
	Advanced *bool `json:"advanced,omitempty"`
}

type SearchRecordGet struct {
	// A limit on the number of objects to be returned with a range between 1 and 100. Defaults to 100.
	Limit *int `json:"-"`
	// Number of results to skip before returning response. Defaults to 0.
	Offset *int `json:"-"`
	// Query term. The syntax for the query parameter follows elasticsearch simple query string syntax. The includes the ability to use search operators and to perform nested queries. Must be url encoded.
	Q string `json:"-"`
	// Record or entity fields to search against.
	Fields []*SearchField `json:"-"`
	// Whether or not to return search facets in results giving counts by field. Defaults to false.
	Facets *bool `json:"-"`
	// Set to true to enable full elasticsearch query string syntax which allows for fielded search and more complex operators. Note that the syntax is more strict and can result in empty result-sets. Defaults to false.
	Advanced *bool `json:"-"`
}

// OK
type EntitySearchResponse struct {
	Limit  int              `json:"limit"`
	Size   *QualifiedCount  `json:"size,omitempty"`
	Offset int              `json:"offset"`
	Next   bool             `json:"next"`
	Data   []*SearchResults `json:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (e *EntitySearchResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler EntitySearchResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = EntitySearchResponse(value)
	e._rawJSON = json.RawMessage(data)
	return nil
}

func (e *EntitySearchResponse) String() string {
	if len(e._rawJSON) > 0 {
		if value, err := core.StringifyJSON(e._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

// Filter your search on the following attributes.
type FilterList struct {
	Source     []SourceId `json:"source,omitempty"`
	Country    []Country  `json:"country,omitempty"`
	State      []string   `json:"state,omitempty"`
	City       []string   `json:"city,omitempty"`
	EntityType []Entities `json:"entity_type,omitempty"`
	Bounds     []string   `json:"bounds,omitempty"`
	Risk       []Risk     `json:"risk,omitempty"`

	_rawJSON json.RawMessage
}

func (f *FilterList) UnmarshalJSON(data []byte) error {
	type unmarshaler FilterList
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*f = FilterList(value)
	f._rawJSON = json.RawMessage(data)
	return nil
}

func (f *FilterList) String() string {
	if len(f._rawJSON) > 0 {
		if value, err := core.StringifyJSON(f._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(f); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", f)
}

// OK
type RecordSearchResponse struct {
	Limit  int              `json:"limit"`
	Size   *QualifiedCount  `json:"size,omitempty"`
	Offset int              `json:"offset"`
	Next   bool             `json:"next"`
	Data   []*RecordDetails `json:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (r *RecordSearchResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler RecordSearchResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RecordSearchResponse(value)
	r._rawJSON = json.RawMessage(data)
	return nil
}

func (r *RecordSearchResponse) String() string {
	if len(r._rawJSON) > 0 {
		if value, err := core.StringifyJSON(r._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}
