// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/sayari-analytics/sayari-go/generated/go/core"
)

type Resolution struct {
	// Entity name
	Name []*string `json:"-"`
	// Entity identifier. Can be from either the [Identifier Type](/sayari-library/ontology/enumerated-types#identifier-type) or [Weak Identifier Type](/sayari-library/ontology/enumerated-types#weak-identifier-type) enums.
	Identifier []*string `json:"-"`
	// Entity country - must be ISO (3166) Trigram i.e., `USA`. See complete list [here](/sayari-library/ontology/enumerated-types#country)
	Country []*Country `json:"-"`
	// Entity address
	Address []*string `json:"-"`
	// Entity date of birth
	DateOfBirth []*string `json:"-"`
	// Entity contact
	Contact []*string `json:"-"`
	// Entity type. If multiple values are passed for any field, the endpoint will match entities with ANY of the values.
	Type []*Entities `json:"-"`
}

// OK
type ResolutionResponse struct {
	Fields *ResolutionResponseFields `json:"fields,omitempty"`
	Data   []*ResolutionResult       `json:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (r *ResolutionResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ResolutionResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = ResolutionResponse(value)
	r._rawJSON = json.RawMessage(data)
	return nil
}

func (r *ResolutionResponse) String() string {
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
