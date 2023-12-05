// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/sayari-analytics/sayari-go/generated/go/core"
	time "time"
)

type GetHistory struct {
	// The type of events to filter on.
	Events []*string `json:"-"`
	// An ISO 8601 encoded date string indicating the starting time period for the events. In the format YYYY-MM-DD
	From *time.Time `json:"-"`
	// An ISO 8601 encoded date string indicating the ending time period for the events. In the format YYYY-MM-DD
	To *time.Time `json:"-"`
	// Size to limit number of events returned
	Size *int `json:"-"`
	// Pagination token to retrieve the next page of results
	Token *string `json:"-"`
}

type GetUsage struct {
	// An ISO 8601 encoded date string indicating the starting time period to obtain usage stats. In the format YYYY-MM-DD
	From *time.Time `json:"-"`
	// An ISO 8601 encoded date string indicating the ending time period to obtain usage stats. In the format YYYY-MM-DD
	To *time.Time `json:"-"`
}

type HistoryResponse struct {
	Size      int            `json:"size"`
	NextToken string         `json:"next_token"`
	Events    []*HistoryInfo `json:"events,omitempty"`

	_rawJSON json.RawMessage
}

func (h *HistoryResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler HistoryResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*h = HistoryResponse(value)
	h._rawJSON = json.RawMessage(data)
	return nil
}

func (h *HistoryResponse) String() string {
	if len(h._rawJSON) > 0 {
		if value, err := core.StringifyJSON(h._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(h); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", h)
}

type UsageResponse struct {
	Usage *UsageInfo `json:"usage,omitempty"`
	From  string     `json:"from"`
	To    string     `json:"to"`

	_rawJSON json.RawMessage
}

func (u *UsageResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UsageResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UsageResponse(value)
	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UsageResponse) String() string {
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
