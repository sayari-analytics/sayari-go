// This file was auto-generated by Fern from our API Definition.

package api

type ListSources struct {
	// A limit on the number of objects to be returned with a range between 1 and 100. Defaults to 100.
	Limit *int `json:"-"`
	// Number of results to skip before returning response. Defaults to 0.
	Offset *int `json:"-"`
}

// The unique identifier for a source in the database
type SourceId = string

type Source struct {
	Id          string  `json:"id"`
	Label       string  `json:"label"`
	Description string  `json:"description"`
	Country     string  `json:"country"`
	Region      string  `json:"region"`
	DateAdded   string  `json:"date_added"`
	SourceType  string  `json:"source_type"`
	RecordType  string  `json:"record_type"`
	Structure   string  `json:"structure"`
	SourceUrl   *string `json:"source_url,omitempty"`
	Pep         bool    `json:"pep"`
	Watchlist   bool    `json:"watchlist"`
}

type SourceList struct {
	Limit  int       `json:"limit"`
	Size   *SizeInfo `json:"size,omitempty"`
	Offset int       `json:"offset"`
	Next   bool      `json:"next"`
	Data   []*Source `json:"data,omitempty"`
}
