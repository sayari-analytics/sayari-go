// This file was auto-generated by Fern from our API Definition.

package api

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
	// Whether or not to return search geo bound facets in results giving counts by geo tile. Defaults to false.
	GeoFacets *bool `json:"geo_facets,omitempty"`
	// Set to true to enable full elasticsearch query string syntax which allows for fielded search and more complex operators. Note that the syntax is more strict and can result in empty result-sets. Defaults to false.
	Advanced *bool `json:"advanced,omitempty"`
}

type EntitySearchResults struct {
	Limit  int              `json:"limit"`
	Size   *SizeInfo        `json:"size,omitempty"`
	Offset int              `json:"offset"`
	Next   bool             `json:"next"`
	Data   []*EntityDetails `json:"data,omitempty"`
}

// Filter your search on the following attributes.
type FilterList struct {
	Source     []SourceId `json:"source,omitempty"`
	Country    []Country  `json:"country,omitempty"`
	State      []string   `json:"state,omitempty"`
	City       []string   `json:"city,omitempty"`
	EntityType []Entities `json:"entity_type,omitempty"`
	Bounds     []string   `json:"bounds,omitempty"`
	Risk       []Tag      `json:"risk,omitempty"`
}

type RecordSearchResults struct {
	Limit  int              `json:"limit"`
	Size   *SizeInfo        `json:"size,omitempty"`
	Offset int              `json:"offset"`
	Next   bool             `json:"next"`
	Data   []*RecordDetails `json:"data,omitempty"`
}
