// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/sayari-analytics/sayari-go/generated/go/core"
)

type MetadataResponse struct {
	// Currently deployed version of the application.
	Version string `json:"version" url:"version"`
	// Currently deployed main data release.
	MasterRelease string `json:"master_release" url:"master_release"`
	// Currently deployed watchlist release.
	WatchlistRelease string    `json:"watchlist_release" url:"watchlist_release"`
	User             *UserInfo `json:"user,omitempty" url:"user,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (m *MetadataResponse) GetExtraProperties() map[string]interface{} {
	return m.extraProperties
}

func (m *MetadataResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler MetadataResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = MetadataResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties

	m._rawJSON = json.RawMessage(data)
	return nil
}

func (m *MetadataResponse) String() string {
	if len(m._rawJSON) > 0 {
		if value, err := core.StringifyJSON(m._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(m); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", m)
}
