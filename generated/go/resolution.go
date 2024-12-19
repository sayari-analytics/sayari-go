// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/sayari-analytics/sayari-go/generated/go/core"
)

type Resolution struct {
	// A limit on the number of objects to be returned with a range between 1 and 10 inclusive. Defaults to 10.
	Limit *int `json:"-" url:"limit,omitempty"`
	// Number of results to skip before returning response. Defaults to 0.
	Offset *int `json:"-" url:"offset,omitempty"`
	// Entity name
	Name []*string `json:"-" url:"name,omitempty"`
	// Entity address. For optimal matching results, it's recommended to concatenate the full address string (street, city, state, postal code).
	Address []*string `json:"-" url:"address,omitempty"`
	// Entity city that contains the provided city name.
	City []*string `json:"-" url:"city,omitempty"`
	// Entity state that contains the provided state name.
	State []*string `json:"-" url:"state,omitempty"`
	// Entity country - must be ISO (3166) Trigram i.e., `USA`. See complete list [here](/sayari-library/ontology/enumerated-types#country)
	Country []*Country `json:"-" url:"country,omitempty"`
	// Entity identifier. Can be from either the [Identifier Type](/sayari-library/ontology/enumerated-types#identifier-type) or [Weak Identifier Type](/sayari-library/ontology/enumerated-types#weak-identifier-type) enums.
	Identifier []*BothIdentifierTypes `json:"-" url:"identifier,omitempty"`
	// Entity date of birth
	DateOfBirth []*string `json:"-" url:"date_of_birth,omitempty"`
	// Entity contact
	Contact []*string `json:"-" url:"contact,omitempty"`
	// [Entity type](/sayari-library/ontology/entities). If multiple values are passed for any field, the endpoint will match entities with ANY of the values.
	Type []*Entities `json:"-" url:"type,omitempty"`
	// Specifies the search algorithm to use. `corporate` (default) is optimized for accurate entity attribute matching, ideal for business verification. `suppliers` is tailored for matching entities with trade data, suitable for supply chain use cases. `search` mimics /search/entity behavior, best for name-only matches.
	Profile *ProfileEnum `json:"-" url:"profile,omitempty"`
	// Adding this param enables an alternative matching logic. It will set a minimum percentage of tokens needed to match with user input to be considered a "hit". Accepts integers from 0 to 100 inclusive.
	NameMinPercentage *int `json:"-" url:"name_min_percentage,omitempty"`
	// Adding this param enables an alternative matching logic. It sets the minimum number of matching tokens the resolved hits need to have in common with the user input to be considered a "hit". Accepts non-negative integers.
	NameMinTokens *int `json:"-" url:"name_min_tokens,omitempty"`
	// Specifies the minimum score required to pass, which controls the strictness of the matching threshold. The default value is 77, and tuned for general use-case accuracy. Increase the value for stricter matching, reduce to loosen.
	MinimumScoreThreshold *int `json:"-" url:"minimum_score_threshold,omitempty"`
	// Enables a name search fallback when either the corporate or supplier profiles fails to find a match. When invoked, the fallback will make a call similar to /search/entity on name only. By default set to false.
	SearchFallback *bool `json:"-" url:"search_fallback,omitempty"`
	// Specifies the window of similar results returned in the match group. Increase for fewer multiple matches, decrease to open the aperture and allow for more matches. Default is .8
	CutoffThreshold *int `json:"-" url:"cutoff_threshold,omitempty"`
	// Specifies the maximum number of entity candidates considered during search. Default is 50. Higher values increase match pool size but also increase latency.
	CandidatePoolSize *int `json:"-" url:"candidate_pool_size,omitempty"`
	// Bypasses the post-processing setps and re-ranking. Useful for debugging. By default set to false, set to true to enable.
	SkipPostProcess *bool `json:"-" url:"skip_post_process,omitempty"`
}

type ResolutionPost struct {
	// A limit on the number of objects to be returned with a range between 1 and 10 inclusive. Defaults to 10.
	Limit *int `json:"-" url:"limit,omitempty"`
	// Number of results to skip before returning response. Defaults to 0.
	Offset *int            `json:"-" url:"offset,omitempty"`
	Body   *ResolutionBody `json:"-" url:"-"`
}

func (r *ResolutionPost) UnmarshalJSON(data []byte) error {
	body := new(ResolutionBody)
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	r.Body = body
	return nil
}

func (r *ResolutionPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Body)
}

type ResolutionPersisted struct {
	// A limit on the number of objects to be returned with a range between 1 and 10 inclusive. Defaults to 10.
	Limit *int `json:"-" url:"limit,omitempty"`
	// Number of results to skip before returning response. Defaults to 0.
	Offset *int            `json:"-" url:"offset,omitempty"`
	Body   *ResolutionBody `json:"-" url:"-"`
}

func (r *ResolutionPersisted) UnmarshalJSON(data []byte) error {
	body := new(ResolutionBody)
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	r.Body = body
	return nil
}

func (r *ResolutionPersisted) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Body)
}

type BothIdentifierTypes struct {
	IdentifierType     IdentifierType
	WeakIdentifierType WeakIdentifierType

	typ string
}

func NewBothIdentifierTypesFromIdentifierType(value IdentifierType) *BothIdentifierTypes {
	return &BothIdentifierTypes{typ: "IdentifierType", IdentifierType: value}
}

func NewBothIdentifierTypesFromWeakIdentifierType(value WeakIdentifierType) *BothIdentifierTypes {
	return &BothIdentifierTypes{typ: "WeakIdentifierType", WeakIdentifierType: value}
}

func (b *BothIdentifierTypes) UnmarshalJSON(data []byte) error {
	var valueIdentifierType IdentifierType
	if err := json.Unmarshal(data, &valueIdentifierType); err == nil {
		b.typ = "IdentifierType"
		b.IdentifierType = valueIdentifierType
		return nil
	}
	var valueWeakIdentifierType WeakIdentifierType
	if err := json.Unmarshal(data, &valueWeakIdentifierType); err == nil {
		b.typ = "WeakIdentifierType"
		b.WeakIdentifierType = valueWeakIdentifierType
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, b)
}

func (b BothIdentifierTypes) MarshalJSON() ([]byte, error) {
	if b.typ == "IdentifierType" || b.IdentifierType != "" {
		return json.Marshal(b.IdentifierType)
	}
	if b.typ == "WeakIdentifierType" || b.WeakIdentifierType != "" {
		return json.Marshal(b.WeakIdentifierType)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", b)
}

type BothIdentifierTypesVisitor interface {
	VisitIdentifierType(IdentifierType) error
	VisitWeakIdentifierType(WeakIdentifierType) error
}

func (b *BothIdentifierTypes) Accept(visitor BothIdentifierTypesVisitor) error {
	if b.typ == "IdentifierType" || b.IdentifierType != "" {
		return visitor.VisitIdentifierType(b.IdentifierType)
	}
	if b.typ == "WeakIdentifierType" || b.WeakIdentifierType != "" {
		return visitor.VisitWeakIdentifierType(b.WeakIdentifierType)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", b)
}

type ProfileEnum string

const (
	ProfileEnumCorporate ProfileEnum = "corporate"
	ProfileEnumSuppliers ProfileEnum = "suppliers"
	ProfileEnumSearch    ProfileEnum = "search"
	ProfileEnumScreen    ProfileEnum = "screen"
)

func NewProfileEnumFromString(s string) (ProfileEnum, error) {
	switch s {
	case "corporate":
		return ProfileEnumCorporate, nil
	case "suppliers":
		return ProfileEnumSuppliers, nil
	case "search":
		return ProfileEnumSearch, nil
	case "screen":
		return ProfileEnumScreen, nil
	}
	var t ProfileEnum
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (p ProfileEnum) Ptr() *ProfileEnum {
	return &p
}

type ResolutionBody struct {
	// Entity name
	Name []string `json:"name,omitempty" url:"name,omitempty"`
	// Entity identifier. Can be from either the [Identifier Type](/sayari-library/ontology/enumerated-types#identifier-type) or [Weak Identifier Type](/sayari-library/ontology/enumerated-types#weak-identifier-type) enums.
	Identifier *BothIdentifierTypes `json:"identifier,omitempty" url:"identifier,omitempty"`
	// Entity address
	Address []string `json:"address,omitempty" url:"address,omitempty"`
	// Entity city that contains the provided city name.
	City *string `json:"city,omitempty" url:"city,omitempty"`
	// Entity state that contains the provided state name.
	State *string `json:"state,omitempty" url:"state,omitempty"`
	// Entity country - must be ISO (3166) Trigram i.e., `USA`. See complete list [here](/sayari-library/ontology/enumerated-types#country)
	Country []Country `json:"country,omitempty" url:"country,omitempty"`
	// Entity date of birth
	DateOfBirth []string `json:"date_of_birth,omitempty" url:"date_of_birth,omitempty"`
	// Entity contact
	Contact []string `json:"contact,omitempty" url:"contact,omitempty"`
	// [Entity type](/sayari-library/ontology/entities). If multiple values are passed for any field, the endpoint will match entities with ANY of the values.
	Type []Entities `json:"type,omitempty" url:"type,omitempty"`
	// Specifies the search algorithm to use. `corporate` (default) is optimized for accurate entity attribute matching, ideal for business verification. `suppliers` is tailored for matching entities with trade data, suitable for supply chain use cases. `search` mimics /search/entity behavior, best for name-only matches.
	Profile *ProfileEnum `json:"profile,omitempty" url:"profile,omitempty"`
	// Adding this param enables an alternative matching logic. It will set a minimum percentage of tokens needed to match with user input to be considered a "hit". Accepts integers from 0 to 100 inclusive.
	NameMinPercentage *int `json:"name_min_percentage,omitempty" url:"name_min_percentage,omitempty"`
	// Adding this param enables an alternative matching logic. It sets the minimum number of matching tokens the resolved hits need to have in common with the user input to be considered a "hit". Accepts non-negative integers.
	NameMinTokens *int `json:"name_min_tokens,omitempty" url:"name_min_tokens,omitempty"`
	// An array of tag labels to associate with each resolved entity
	Tags []string `json:"tags,omitempty" url:"tags,omitempty"`
	// Specifies the minimum score required to pass, which controls the strictness of the matching threshold. The default value is 77, and tuned for general use-case accuracy. Increase the value for stricter matching, reduce to loosen.
	MinimumScoreThreshold *int `json:"minimum_score_threshold,omitempty" url:"minimum_score_threshold,omitempty"`
	// Enables a name search fallback when either the corporate or supplier profiles fails to find a match. When invoked, the fallback will make a call similar to /search/entity on name only. By default set to false.
	SearchFallback *bool `json:"search_fallback,omitempty" url:"search_fallback,omitempty"`
	// Specifies the window of similar results returned in the match group. Increase for fewer multiple matches, decrease to open the aperture and allow for more matches. Default is .8
	CutoffThreshold *int `json:"cutoff_threshold,omitempty" url:"cutoff_threshold,omitempty"`
	// Specifies the maximum number of entity candidates considered during search. Default is 50. Higher values increase match pool size but also increase latency.
	CandidatePoolSize *int `json:"candidate_pool_size,omitempty" url:"candidate_pool_size,omitempty"`
	// Bypasses the post-processing setps and re-ranking. Useful for debugging. By default set to false, set to true to enable.
	SkipPostProcess *bool `json:"skip_post_process,omitempty" url:"skip_post_process,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (r *ResolutionBody) GetExtraProperties() map[string]interface{} {
	return r.extraProperties
}

func (r *ResolutionBody) UnmarshalJSON(data []byte) error {
	type unmarshaler ResolutionBody
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = ResolutionBody(value)

	extraProperties, err := core.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties

	r._rawJSON = json.RawMessage(data)
	return nil
}

func (r *ResolutionBody) String() string {
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

type ResolutionPersistedResponse struct {
	Fields *ResolutionPersistedResponseFields `json:"fields,omitempty" url:"fields,omitempty"`
	Data   []*ResolutionPersistedResult       `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (r *ResolutionPersistedResponse) GetExtraProperties() map[string]interface{} {
	return r.extraProperties
}

func (r *ResolutionPersistedResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ResolutionPersistedResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = ResolutionPersistedResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties

	r._rawJSON = json.RawMessage(data)
	return nil
}

func (r *ResolutionPersistedResponse) String() string {
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

type ResolutionResponse struct {
	Fields *ResolutionResponseFields `json:"fields,omitempty" url:"fields,omitempty"`
	Data   []*ResolutionResult       `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (r *ResolutionResponse) GetExtraProperties() map[string]interface{} {
	return r.extraProperties
}

func (r *ResolutionResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ResolutionResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = ResolutionResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties

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

type ResolutionUploadBody struct {
	Filename string            `json:"filename" url:"filename"`
	Data     []*ResolutionBody `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (r *ResolutionUploadBody) GetExtraProperties() map[string]interface{} {
	return r.extraProperties
}

func (r *ResolutionUploadBody) UnmarshalJSON(data []byte) error {
	type unmarshaler ResolutionUploadBody
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = ResolutionUploadBody(value)

	extraProperties, err := core.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties

	r._rawJSON = json.RawMessage(data)
	return nil
}

func (r *ResolutionUploadBody) String() string {
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

type ResolutionUploadResponse struct {
	File     string `json:"file" url:"file"`
	Uploaded string `json:"uploaded" url:"uploaded"`
	Count    int    `json:"count" url:"count"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (r *ResolutionUploadResponse) GetExtraProperties() map[string]interface{} {
	return r.extraProperties
}

func (r *ResolutionUploadResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ResolutionUploadResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = ResolutionUploadResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties

	r._rawJSON = json.RawMessage(data)
	return nil
}

func (r *ResolutionUploadResponse) String() string {
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
