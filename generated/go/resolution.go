// This file was auto-generated by Fern from our API Definition.

package api

type Resolution struct {
	// Entity name
	Name []*string `json:"-"`
	// Entity identifier
	Identifier []*string `json:"-"`
	// Entity country
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

type ResolutionResponse struct {
	Fields *ResolutionResponseFields `json:"fields,omitempty"`
	Data   []*ResolutionResult       `json:"data,omitempty"`
}
