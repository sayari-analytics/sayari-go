// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/sayari-analytics/sayari-go/generated/go/internal"
)

type ListSources struct {
	// A limit on the number of objects to be returned with a range between 1 and 100. Defaults to 100.
	Limit *int `json:"-" url:"limit,omitempty"`
	// Number of results to skip before returning response. Defaults to 0.
	Offset *int `json:"-" url:"offset,omitempty"`
}

// OK
type GetSourceResponse struct {
	// The unique identifier of the source
	Id          string `json:"id" url:"id"`
	Label       string `json:"label" url:"label"`
	Description string `json:"description" url:"description"`
	// Source [country](/sayari-library/ontology/enumerated-types#country)
	Country    Country `json:"country" url:"country"`
	Region     string  `json:"region" url:"region"`
	DateAdded  string  `json:"date_added" url:"date_added"`
	SourceType string  `json:"source_type" url:"source_type"`
	RecordType string  `json:"record_type" url:"record_type"`
	Structure  string  `json:"structure" url:"structure"`
	SourceUrl  *string `json:"source_url,omitempty" url:"source_url,omitempty"`
	Pep        bool    `json:"pep" url:"pep"`
	Watchlist  bool    `json:"watchlist" url:"watchlist"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (g *GetSourceResponse) GetId() string {
	if g == nil {
		return ""
	}
	return g.Id
}

func (g *GetSourceResponse) GetLabel() string {
	if g == nil {
		return ""
	}
	return g.Label
}

func (g *GetSourceResponse) GetDescription() string {
	if g == nil {
		return ""
	}
	return g.Description
}

func (g *GetSourceResponse) GetCountry() Country {
	if g == nil {
		return ""
	}
	return g.Country
}

func (g *GetSourceResponse) GetRegion() string {
	if g == nil {
		return ""
	}
	return g.Region
}

func (g *GetSourceResponse) GetDateAdded() string {
	if g == nil {
		return ""
	}
	return g.DateAdded
}

func (g *GetSourceResponse) GetSourceType() string {
	if g == nil {
		return ""
	}
	return g.SourceType
}

func (g *GetSourceResponse) GetRecordType() string {
	if g == nil {
		return ""
	}
	return g.RecordType
}

func (g *GetSourceResponse) GetStructure() string {
	if g == nil {
		return ""
	}
	return g.Structure
}

func (g *GetSourceResponse) GetSourceUrl() *string {
	if g == nil {
		return nil
	}
	return g.SourceUrl
}

func (g *GetSourceResponse) GetPep() bool {
	if g == nil {
		return false
	}
	return g.Pep
}

func (g *GetSourceResponse) GetWatchlist() bool {
	if g == nil {
		return false
	}
	return g.Watchlist
}

func (g *GetSourceResponse) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *GetSourceResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GetSourceResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetSourceResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	g.rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetSourceResponse) String() string {
	if len(g.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(g.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}

// OK
type ListSourcesResponse struct {
	Limit  int             `json:"limit" url:"limit"`
	Size   *QualifiedCount `json:"size,omitempty" url:"size,omitempty"`
	Offset int             `json:"offset" url:"offset"`
	Next   bool            `json:"next" url:"next"`
	Data   []*Source       `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListSourcesResponse) GetLimit() int {
	if l == nil {
		return 0
	}
	return l.Limit
}

func (l *ListSourcesResponse) GetSize() *QualifiedCount {
	if l == nil {
		return nil
	}
	return l.Size
}

func (l *ListSourcesResponse) GetOffset() int {
	if l == nil {
		return 0
	}
	return l.Offset
}

func (l *ListSourcesResponse) GetNext() bool {
	if l == nil {
		return false
	}
	return l.Next
}

func (l *ListSourcesResponse) GetData() []*Source {
	if l == nil {
		return nil
	}
	return l.Data
}

func (l *ListSourcesResponse) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *ListSourcesResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListSourcesResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListSourcesResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListSourcesResponse) String() string {
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

type Source struct {
	// The unique identifier of the source
	Id          string `json:"id" url:"id"`
	Label       string `json:"label" url:"label"`
	Description string `json:"description" url:"description"`
	// Source [country](/sayari-library/ontology/enumerated-types#country)
	Country    Country `json:"country" url:"country"`
	Region     string  `json:"region" url:"region"`
	DateAdded  string  `json:"date_added" url:"date_added"`
	SourceType string  `json:"source_type" url:"source_type"`
	RecordType string  `json:"record_type" url:"record_type"`
	Structure  string  `json:"structure" url:"structure"`
	SourceUrl  *string `json:"source_url,omitempty" url:"source_url,omitempty"`
	Pep        bool    `json:"pep" url:"pep"`
	Watchlist  bool    `json:"watchlist" url:"watchlist"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (s *Source) GetId() string {
	if s == nil {
		return ""
	}
	return s.Id
}

func (s *Source) GetLabel() string {
	if s == nil {
		return ""
	}
	return s.Label
}

func (s *Source) GetDescription() string {
	if s == nil {
		return ""
	}
	return s.Description
}

func (s *Source) GetCountry() Country {
	if s == nil {
		return ""
	}
	return s.Country
}

func (s *Source) GetRegion() string {
	if s == nil {
		return ""
	}
	return s.Region
}

func (s *Source) GetDateAdded() string {
	if s == nil {
		return ""
	}
	return s.DateAdded
}

func (s *Source) GetSourceType() string {
	if s == nil {
		return ""
	}
	return s.SourceType
}

func (s *Source) GetRecordType() string {
	if s == nil {
		return ""
	}
	return s.RecordType
}

func (s *Source) GetStructure() string {
	if s == nil {
		return ""
	}
	return s.Structure
}

func (s *Source) GetSourceUrl() *string {
	if s == nil {
		return nil
	}
	return s.SourceUrl
}

func (s *Source) GetPep() bool {
	if s == nil {
		return false
	}
	return s.Pep
}

func (s *Source) GetWatchlist() bool {
	if s == nil {
		return false
	}
	return s.Watchlist
}

func (s *Source) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *Source) UnmarshalJSON(data []byte) error {
	type unmarshaler Source
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = Source(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s *Source) String() string {
	if len(s.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(s.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}
