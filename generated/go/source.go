// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/sayari-analytics/sayari-go/generated/go/core"
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
	_rawJSON        json.RawMessage
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

	extraProperties, err := core.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties

	g._rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetSourceResponse) String() string {
	if len(g._rawJSON) > 0 {
		if value, err := core.StringifyJSON(g._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(g); err == nil {
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
	_rawJSON        json.RawMessage
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

	extraProperties, err := core.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties

	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListSourcesResponse) String() string {
	if len(l._rawJSON) > 0 {
		if value, err := core.StringifyJSON(l._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(l); err == nil {
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
	_rawJSON        json.RawMessage
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

	extraProperties, err := core.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *Source) String() string {
	if len(s._rawJSON) > 0 {
		if value, err := core.StringifyJSON(s._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}
