// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"

	core "github.com/sayari-analytics/sayari-go/generated/go/core"
)

type DeleteResourceResponse struct {
	Data *EntityResponseData `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (d *DeleteResourceResponse) GetExtraProperties() map[string]interface{} {
	return d.extraProperties
}

func (d *DeleteResourceResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler DeleteResourceResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DeleteResourceResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties

	d._rawJSON = json.RawMessage(data)
	return nil
}

func (d *DeleteResourceResponse) String() string {
	if len(d._rawJSON) > 0 {
		if value, err := core.StringifyJSON(d._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

type ResourceType string

const (
	ResourceTypeEntity ResourceType = "entity"
)

func NewResourceTypeFromString(s string) (ResourceType, error) {
	switch s {
	case "entity":
		return ResourceTypeEntity, nil
	}
	var t ResourceType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (r ResourceType) Ptr() *ResourceType {
	return &r
}

type SaveEntityRequest struct {
	Type ResourceType `json:"type" url:"type"`
	// The project identifier.
	Project string `json:"project" url:"project"`
	// The entity identifier.
	EntityId string `json:"entity_id" url:"entity_id"`
	// <Warning>This property is in beta and is subject to change. It is provided for early access and testing purposes only.</Warning> custom user key/value pairs (key must be prefixed with "custom\_" and value must be "string" type)
	CustomFields interface{} `json:"custom_fields,omitempty" url:"custom_fields,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *SaveEntityRequest) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SaveEntityRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler SaveEntityRequest
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SaveEntityRequest(value)

	extraProperties, err := core.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SaveEntityRequest) String() string {
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

type SaveEntityResponse struct {
	Data *EntityResponseData `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *SaveEntityResponse) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SaveEntityResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler SaveEntityResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SaveEntityResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SaveEntityResponse) String() string {
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
