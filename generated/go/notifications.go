// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/sayari-analytics/sayari-go/generated/go/core"
)

type ProjectNotifications struct {
	// Limit total notifications in the response. Defaults to 100.
	Limit *int `json:"-" url:"limit,omitempty"`
	// Offset which notifications are returned. Defaults to 0.
	Offset *int                    `json:"-" url:"offset,omitempty"`
	Sort   *NotificationsSortField `json:"-" url:"sort,omitempty"`
}

type ResourceNotifications struct {
	// Limit total notifications in the response. Defaults to 100.
	Limit *int `json:"-" url:"limit,omitempty"`
	// Offset which notifications are returned. Defaults to 0.
	Offset *int `json:"-" url:"offset,omitempty"`
}

// Defines a sort order on a field. The value should begin with a '-' to indicate a descending sort, followed by a field name to sort on.
type NotificationsSortField string

const (
	// The date the notification was generated, descending.
	NotificationsSortFieldDateDesc NotificationsSortField = "-date"
)

func NewNotificationsSortFieldFromString(s string) (NotificationsSortField, error) {
	switch s {
	case "-date":
		return NotificationsSortFieldDateDesc, nil
	}
	var t NotificationsSortField
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (n NotificationsSortField) Ptr() *NotificationsSortField {
	return &n
}

// OK
type ProjectNotificationsResponse struct {
	Offset int                        `json:"offset" url:"offset"`
	Limit  int                        `json:"limit" url:"limit"`
	Next   bool                       `json:"next" url:"next"`
	Data   []*ProjectNotificationData `json:"data,omitempty" url:"data,omitempty"`
	Size   *QualifiedCount            `json:"size,omitempty" url:"size,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *ProjectNotificationsResponse) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *ProjectNotificationsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ProjectNotificationsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = ProjectNotificationsResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *ProjectNotificationsResponse) String() string {
	if len(p._rawJSON) > 0 {
		if value, err := core.StringifyJSON(p._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

// OK
type ResourceNotificationsResponse struct {
	Offset int                         `json:"offset" url:"offset"`
	Limit  int                         `json:"limit" url:"limit"`
	Next   bool                        `json:"next" url:"next"`
	Data   []*ResourceNotificationData `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (r *ResourceNotificationsResponse) GetExtraProperties() map[string]interface{} {
	return r.extraProperties
}

func (r *ResourceNotificationsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ResourceNotificationsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = ResourceNotificationsResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties

	r._rawJSON = json.RawMessage(data)
	return nil
}

func (r *ResourceNotificationsResponse) String() string {
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
