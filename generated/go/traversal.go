// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/sayari-analytics/sayari-go/generated/go/internal"
)

type Ownership struct {
	// Limit total values for traversal. Defaults to 10. Max of 50.
	Limit *int `json:"-" url:"limit,omitempty"`
	// Offset values for traversal. Defaults to 0. Max of 1000.
	Offset *int `json:"-" url:"offset,omitempty"`
	// Set minimum depth for traversal. Defaults to 1.
	MinDepth *int `json:"-" url:"min_depth,omitempty"`
	// Set maximum depth for traversal. Defaults to 4.
	MaxDepth *int `json:"-" url:"max_depth,omitempty"`
	// Set relationship type(s) to follow when traversing related entities. Defaults to shareholder_of, beneficial_owner_of, partner_of, has_subsidiary, and has_branch.
	Relationships []*Relationships `json:"-" url:"relationships,omitempty"`
	// Also traverse relationships from entities that are possibly the same as any entity that appears in the path. Defaults to traversing possibly same as relationships.
	Psa *bool `json:"-" url:"psa,omitempty"`
	// Filter paths to only those that end at an entity associated with the specified country(ies). Defaults to returning paths that end in any country.
	Countries []*Country `json:"-" url:"countries,omitempty"`
	// Filter paths to only those that end at an entity of the specified type(s). Defaults to returning paths that end at any type.
	Types []*Entities `json:"-" url:"types,omitempty"`
	// Filter paths to only those that end at an entity appearing on a watchlist. Defaults to not filtering paths by sanctioned status.
	Sanctioned *bool `json:"-" url:"sanctioned,omitempty"`
	// Filter paths to only those that end at an entity appearing on a pep list. Defaults to not filtering paths by pep status.
	Pep *bool `json:"-" url:"pep,omitempty"`
	// Set minimum percentage of share ownership for traversal. Defaults to 0.
	MinShares *int `json:"-" url:"min_shares,omitempty"`
	// Also traverse relationships when share percentages are unknown. Only useful when min_shares is set greater than 0. Defaults to true.
	IncludeUnknownShares *bool `json:"-" url:"include_unknown_shares,omitempty"`
	// Include relationships that were valid in the past but not at the present time. Defaults to true.
	ExcludeFormerRelationships *bool `json:"-" url:"exclude_former_relationships,omitempty"`
	// Include entities that existed in the past but not at the present time. Defaults to false.
	ExcludeClosedEntities *bool `json:"-" url:"exclude_closed_entities,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with any risk factor of these categories
	RiskCategories *TraversalRiskCategory `json:"-" url:"risk_categories,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	EuHighRiskThird *bool `json:"-" url:"eu_high_risk_third,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	ReputationalRiskModernSlavery *bool `json:"-" url:"reputational_risk_modern_slavery,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	StateOwned *bool `json:"-" url:"state_owned,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	FormerlySanctioned *bool `json:"-" url:"formerly_sanctioned,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	ReputationalRiskTerrorism *bool `json:"-" url:"reputational_risk_terrorism,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	ReputationalRiskOrganizedCrime *bool `json:"-" url:"reputational_risk_organized_crime,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	ReputationalRiskFinancialCrime *bool `json:"-" url:"reputational_risk_financial_crime,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	ReputationalRiskBriberyAndCorruption *bool `json:"-" url:"reputational_risk_bribery_and_corruption,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	ReputationalRiskOther *bool `json:"-" url:"reputational_risk_other,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	ReputationalRiskCybercrime *bool `json:"-" url:"reputational_risk_cybercrime,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	RegulatoryAction *bool `json:"-" url:"regulatory_action,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	LawEnforcementAction *bool `json:"-" url:"law_enforcement_action,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	XinjiangGeospatial *bool `json:"-" url:"xinjiang_geospatial,omitempty"`
}

type ShortestPath struct {
	// A list of Sayari entity IDs specifying the source and target entities for the shortest path calculation. The list must contain exactly two entity IDs The first entity ID represents the source.The second entity ID represents the target.
	Entities []string `json:"-" url:"entities"`
}

type Traversal struct {
	// Limit total values for traversal. Defaults to 10. Max of 50.
	Limit *int `json:"-" url:"limit,omitempty"`
	// Offset values for traversal. Defaults to 0. Max of 1000.
	Offset *int `json:"-" url:"offset,omitempty"`
	// Set minimum depth for traversal. Defaults to 1.
	MinDepth *int `json:"-" url:"min_depth,omitempty"`
	// Set maximum depth for traversal. Defaults to 4.
	MaxDepth *int `json:"-" url:"max_depth,omitempty"`
	// Set relationship type(s) to follow when traversing related entities. Defaults to following all relationship types.
	Relationships []*Relationships `json:"-" url:"relationships,omitempty"`
	// Also traverse relationships from entities that are possibly the same as any entity that appears in the path. Defaults to traversing possibly same as relationships.
	Psa *bool `json:"-" url:"psa,omitempty"`
	// Filter paths to only those that end at an entity associated with the specified country(ies). Defaults to returning paths that end in any [country](/sayari-library/ontology/enumerated-types#country).
	Countries []*Country `json:"-" url:"countries,omitempty"`
	// Filter paths to only those that end at an entity of the specified type(s). Defaults to returning paths that end at any type.
	Types []*Entities `json:"-" url:"types,omitempty"`
	// Filter paths to only those that end at an entity appearing on a watchlist. Defaults to not filtering paths by sanctioned status.
	Sanctioned *bool `json:"-" url:"sanctioned,omitempty"`
	// Filter paths to only those that end at an entity appearing on a pep list. Defaults to not filtering paths by pep status.
	Pep *bool `json:"-" url:"pep,omitempty"`
	// Set minimum percentage of share ownership for traversal. Defaults to 0.
	MinShares *int `json:"-" url:"min_shares,omitempty"`
	// Also traverse relationships when share percentages are unknown. Only useful when min_shares is set greater than 0. Defaults to true.
	IncludeUnknownShares *bool `json:"-" url:"include_unknown_shares,omitempty"`
	// Include relationships that were valid in the past but not at the present time. Defaults to true.
	ExcludeFormerRelationships *bool `json:"-" url:"exclude_former_relationships,omitempty"`
	// Include entities that existed in the past but not at the present time. Defaults to false.
	ExcludeClosedEntities *bool `json:"-" url:"exclude_closed_entities,omitempty"`
	// Filter paths to only those that include an entity associated with any risk factor belonging to one of the specified categories.
	RiskCategories []RiskCategory `json:"-" url:"risk_categories,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	EuHighRiskThird *bool `json:"-" url:"eu_high_risk_third,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	ReputationalRiskModernSlavery *bool `json:"-" url:"reputational_risk_modern_slavery,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	StateOwned *bool `json:"-" url:"state_owned,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	FormerlySanctioned *bool `json:"-" url:"formerly_sanctioned,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	ReputationalRiskTerrorism *bool `json:"-" url:"reputational_risk_terrorism,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	ReputationalRiskOrganizedCrime *bool `json:"-" url:"reputational_risk_organized_crime,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	ReputationalRiskFinancialCrime *bool `json:"-" url:"reputational_risk_financial_crime,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	ReputationalRiskBriberyAndCorruption *bool `json:"-" url:"reputational_risk_bribery_and_corruption,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	ReputationalRiskOther *bool `json:"-" url:"reputational_risk_other,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	ReputationalRiskCybercrime *bool `json:"-" url:"reputational_risk_cybercrime,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	RegulatoryAction *bool `json:"-" url:"regulatory_action,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	LawEnforcementAction *bool `json:"-" url:"law_enforcement_action,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	XinjiangGeospatial *bool `json:"-" url:"xinjiang_geospatial,omitempty"`
}

type ShortestPathData struct {
	Source string           `json:"source" url:"source"`
	Target *EntityDetails   `json:"target,omitempty" url:"target,omitempty"`
	Path   []*TraversalPath `json:"path,omitempty" url:"path,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (s *ShortestPathData) GetSource() string {
	if s == nil {
		return ""
	}
	return s.Source
}

func (s *ShortestPathData) GetTarget() *EntityDetails {
	if s == nil {
		return nil
	}
	return s.Target
}

func (s *ShortestPathData) GetPath() []*TraversalPath {
	if s == nil {
		return nil
	}
	return s.Path
}

func (s *ShortestPathData) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *ShortestPathData) UnmarshalJSON(data []byte) error {
	type unmarshaler ShortestPathData
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = ShortestPathData(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s *ShortestPathData) String() string {
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

// OK
type ShortestPathResponse struct {
	Entities []string            `json:"entities,omitempty" url:"entities,omitempty"`
	Data     []*ShortestPathData `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (s *ShortestPathResponse) GetEntities() []string {
	if s == nil {
		return nil
	}
	return s.Entities
}

func (s *ShortestPathResponse) GetData() []*ShortestPathData {
	if s == nil {
		return nil
	}
	return s.Data
}

func (s *ShortestPathResponse) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *ShortestPathResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ShortestPathResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = ShortestPathResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s *ShortestPathResponse) String() string {
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

type TraversalData struct {
	Source string           `json:"source" url:"source"`
	Target *EntityDetails   `json:"target,omitempty" url:"target,omitempty"`
	Path   []*TraversalPath `json:"path,omitempty" url:"path,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (t *TraversalData) GetSource() string {
	if t == nil {
		return ""
	}
	return t.Source
}

func (t *TraversalData) GetTarget() *EntityDetails {
	if t == nil {
		return nil
	}
	return t.Target
}

func (t *TraversalData) GetPath() []*TraversalPath {
	if t == nil {
		return nil
	}
	return t.Path
}

func (t *TraversalData) GetExtraProperties() map[string]interface{} {
	return t.extraProperties
}

func (t *TraversalData) UnmarshalJSON(data []byte) error {
	type unmarshaler TraversalData
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TraversalData(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	t.rawJSON = json.RawMessage(data)
	return nil
}

func (t *TraversalData) String() string {
	if len(t.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(t.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}

type TraversalPath struct {
	Field         string                                       `json:"field" url:"field"`
	Entity        *EntityDetails                               `json:"entity,omitempty" url:"entity,omitempty"`
	Relationships map[Relationships]*TraversalRelationshipData `json:"relationships,omitempty" url:"relationships,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (t *TraversalPath) GetField() string {
	if t == nil {
		return ""
	}
	return t.Field
}

func (t *TraversalPath) GetEntity() *EntityDetails {
	if t == nil {
		return nil
	}
	return t.Entity
}

func (t *TraversalPath) GetRelationships() map[Relationships]*TraversalRelationshipData {
	if t == nil {
		return nil
	}
	return t.Relationships
}

func (t *TraversalPath) GetExtraProperties() map[string]interface{} {
	return t.extraProperties
}

func (t *TraversalPath) UnmarshalJSON(data []byte) error {
	type unmarshaler TraversalPath
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TraversalPath(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	t.rawJSON = json.RawMessage(data)
	return nil
}

func (t *TraversalPath) String() string {
	if len(t.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(t.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}

type TraversalRelationshipData struct {
	Values       []*RelationshipInfo `json:"values,omitempty" url:"values,omitempty"`
	Former       *bool               `json:"former,omitempty" url:"former,omitempty"`
	StartDate    *string             `json:"start_date,omitempty" url:"start_date,omitempty"`
	LastObserved *string             `json:"last_observed,omitempty" url:"last_observed,omitempty"`
	EndDate      *string             `json:"end_date,omitempty" url:"end_date,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (t *TraversalRelationshipData) GetValues() []*RelationshipInfo {
	if t == nil {
		return nil
	}
	return t.Values
}

func (t *TraversalRelationshipData) GetFormer() *bool {
	if t == nil {
		return nil
	}
	return t.Former
}

func (t *TraversalRelationshipData) GetStartDate() *string {
	if t == nil {
		return nil
	}
	return t.StartDate
}

func (t *TraversalRelationshipData) GetLastObserved() *string {
	if t == nil {
		return nil
	}
	return t.LastObserved
}

func (t *TraversalRelationshipData) GetEndDate() *string {
	if t == nil {
		return nil
	}
	return t.EndDate
}

func (t *TraversalRelationshipData) GetExtraProperties() map[string]interface{} {
	return t.extraProperties
}

func (t *TraversalRelationshipData) UnmarshalJSON(data []byte) error {
	type unmarshaler TraversalRelationshipData
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TraversalRelationshipData(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	t.rawJSON = json.RawMessage(data)
	return nil
}

func (t *TraversalRelationshipData) String() string {
	if len(t.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(t.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}

// OK
type TraversalResponse struct {
	MinDepth      int             `json:"min_depth" url:"min_depth"`
	MaxDepth      int             `json:"max_depth" url:"max_depth"`
	Relationships []Relationships `json:"relationships,omitempty" url:"relationships,omitempty"`
	Countries     []Country       `json:"countries,omitempty" url:"countries,omitempty"`
	Types         []string        `json:"types,omitempty" url:"types,omitempty"`
	Name          string          `json:"name" url:"name"`
	// <Warning>This field is deprecated.</Warning>
	Watchlist      bool             `json:"watchlist" url:"watchlist"`
	Psa            bool             `json:"psa" url:"psa"`
	Offset         int              `json:"offset" url:"offset"`
	Limit          int              `json:"limit" url:"limit"`
	Next           bool             `json:"next" url:"next"`
	PartialResults bool             `json:"partial_results" url:"partial_results"`
	Data           []*TraversalData `json:"data,omitempty" url:"data,omitempty"`
	Sanctioned     *bool            `json:"sanctioned,omitempty" url:"sanctioned,omitempty"`
	Pep            *bool            `json:"pep,omitempty" url:"pep,omitempty"`
	ExploredCount  int              `json:"explored_count" url:"explored_count"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (t *TraversalResponse) GetMinDepth() int {
	if t == nil {
		return 0
	}
	return t.MinDepth
}

func (t *TraversalResponse) GetMaxDepth() int {
	if t == nil {
		return 0
	}
	return t.MaxDepth
}

func (t *TraversalResponse) GetRelationships() []Relationships {
	if t == nil {
		return nil
	}
	return t.Relationships
}

func (t *TraversalResponse) GetCountries() []Country {
	if t == nil {
		return nil
	}
	return t.Countries
}

func (t *TraversalResponse) GetTypes() []string {
	if t == nil {
		return nil
	}
	return t.Types
}

func (t *TraversalResponse) GetName() string {
	if t == nil {
		return ""
	}
	return t.Name
}

func (t *TraversalResponse) GetWatchlist() bool {
	if t == nil {
		return false
	}
	return t.Watchlist
}

func (t *TraversalResponse) GetPsa() bool {
	if t == nil {
		return false
	}
	return t.Psa
}

func (t *TraversalResponse) GetOffset() int {
	if t == nil {
		return 0
	}
	return t.Offset
}

func (t *TraversalResponse) GetLimit() int {
	if t == nil {
		return 0
	}
	return t.Limit
}

func (t *TraversalResponse) GetNext() bool {
	if t == nil {
		return false
	}
	return t.Next
}

func (t *TraversalResponse) GetPartialResults() bool {
	if t == nil {
		return false
	}
	return t.PartialResults
}

func (t *TraversalResponse) GetData() []*TraversalData {
	if t == nil {
		return nil
	}
	return t.Data
}

func (t *TraversalResponse) GetSanctioned() *bool {
	if t == nil {
		return nil
	}
	return t.Sanctioned
}

func (t *TraversalResponse) GetPep() *bool {
	if t == nil {
		return nil
	}
	return t.Pep
}

func (t *TraversalResponse) GetExploredCount() int {
	if t == nil {
		return 0
	}
	return t.ExploredCount
}

func (t *TraversalResponse) GetExtraProperties() map[string]interface{} {
	return t.extraProperties
}

func (t *TraversalResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler TraversalResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TraversalResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	t.rawJSON = json.RawMessage(data)
	return nil
}

func (t *TraversalResponse) String() string {
	if len(t.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(t.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}

type TraversalRiskCategory struct {
	// Filter paths to only those that include an entity associated with any risk factor belonging to one of the specified categories.
	RiskCategoryList []RiskCategory
	// Matches a custom risk category
	String string

	typ string
}

func NewTraversalRiskCategoryFromRiskCategoryList(value []RiskCategory) *TraversalRiskCategory {
	return &TraversalRiskCategory{typ: "RiskCategoryList", RiskCategoryList: value}
}

func NewTraversalRiskCategoryFromString(value string) *TraversalRiskCategory {
	return &TraversalRiskCategory{typ: "String", String: value}
}

func (t *TraversalRiskCategory) GetRiskCategoryList() []RiskCategory {
	if t == nil {
		return nil
	}
	return t.RiskCategoryList
}

func (t *TraversalRiskCategory) GetString() string {
	if t == nil {
		return ""
	}
	return t.String
}

func (t *TraversalRiskCategory) UnmarshalJSON(data []byte) error {
	var valueRiskCategoryList []RiskCategory
	if err := json.Unmarshal(data, &valueRiskCategoryList); err == nil {
		t.typ = "RiskCategoryList"
		t.RiskCategoryList = valueRiskCategoryList
		return nil
	}
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		t.typ = "String"
		t.String = valueString
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, t)
}

func (t TraversalRiskCategory) MarshalJSON() ([]byte, error) {
	if t.typ == "RiskCategoryList" || t.RiskCategoryList != nil {
		return json.Marshal(t.RiskCategoryList)
	}
	if t.typ == "String" || t.String != "" {
		return json.Marshal(t.String)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", t)
}

type TraversalRiskCategoryVisitor interface {
	VisitRiskCategoryList([]RiskCategory) error
	VisitString(string) error
}

func (t *TraversalRiskCategory) Accept(visitor TraversalRiskCategoryVisitor) error {
	if t.typ == "RiskCategoryList" || t.RiskCategoryList != nil {
		return visitor.VisitRiskCategoryList(t.RiskCategoryList)
	}
	if t.typ == "String" || t.String != "" {
		return visitor.VisitString(t.String)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", t)
}

type Ubo struct {
	// Limit total values for traversal. Defaults to 10. Max of 50.
	Limit *int `json:"-" url:"limit,omitempty"`
	// Offset values for traversal. Defaults to 0. Max of 1000.
	Offset *int `json:"-" url:"offset,omitempty"`
	// Set minimum depth for traversal. Defaults to 1.
	MinDepth *int `json:"-" url:"min_depth,omitempty"`
	// Set maximum depth for traversal. Defaults to 4.
	MaxDepth *int `json:"-" url:"max_depth,omitempty"`
	// Set relationship type(s) to follow when traversing related entities. Defaults to has_shareholder, has_beneficial_owner, has_partner, subsidiary_of, and branch_of.
	Relationships []*Relationships `json:"-" url:"relationships,omitempty"`
	// Also traverse relationships from entities that are possibly the same as any entity that appears in the path. Defaults to traversing possibly same as relationships.
	Psa *bool `json:"-" url:"psa,omitempty"`
	// Filter paths to only those that end at an entity associated with the specified country(ies). Defaults to returning paths that end in any [country](/sayari-library/ontology/enumerated-types#country).
	Countries []*Country `json:"-" url:"countries,omitempty"`
	// Filter paths to only those that end at an entity of the specified type(s). Defaults to returning paths that end at any type.
	Types []*Entities `json:"-" url:"types,omitempty"`
	// Filter paths to only those that end at an entity appearing on a watchlist. Defaults to not filtering paths by sanctioned status.
	Sanctioned *bool `json:"-" url:"sanctioned,omitempty"`
	// Filter paths to only those that end at an entity appearing on a pep list. Defaults to not filtering paths by pep status.
	Pep *bool `json:"-" url:"pep,omitempty"`
	// Set minimum percentage of share ownership for traversal. Defaults to 0.
	MinShares *int `json:"-" url:"min_shares,omitempty"`
	// Also traverse relationships when share percentages are unknown. Only useful when min_shares is set greater than 0. Defaults to true.
	IncludeUnknownShares *bool `json:"-" url:"include_unknown_shares,omitempty"`
	// Include relationships that were valid in the past but not at the present time. Defaults to true.
	ExcludeFormerRelationships *bool `json:"-" url:"exclude_former_relationships,omitempty"`
	// Include entities that existed in the past but not at the present time. Defaults to false.
	ExcludeClosedEntities *bool `json:"-" url:"exclude_closed_entities,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with any risk factor of these categories
	RiskCategories *TraversalRiskCategory `json:"-" url:"risk_categories,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	EuHighRiskThird *bool `json:"-" url:"eu_high_risk_third,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	ReputationalRiskModernSlavery *bool `json:"-" url:"reputational_risk_modern_slavery,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	StateOwned *bool `json:"-" url:"state_owned,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	FormerlySanctioned *bool `json:"-" url:"formerly_sanctioned,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	ReputationalRiskTerrorism *bool `json:"-" url:"reputational_risk_terrorism,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	ReputationalRiskOrganizedCrime *bool `json:"-" url:"reputational_risk_organized_crime,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	ReputationalRiskFinancialCrime *bool `json:"-" url:"reputational_risk_financial_crime,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	ReputationalRiskBriberyAndCorruption *bool `json:"-" url:"reputational_risk_bribery_and_corruption,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	ReputationalRiskOther *bool `json:"-" url:"reputational_risk_other,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	ReputationalRiskCybercrime *bool `json:"-" url:"reputational_risk_cybercrime,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	RegulatoryAction *bool `json:"-" url:"regulatory_action,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	LawEnforcementAction *bool `json:"-" url:"law_enforcement_action,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	XinjiangGeospatial *bool `json:"-" url:"xinjiang_geospatial,omitempty"`
}

type Watchlist struct {
	// Limit total values for traversal. Defaults to 10. Max of 50.
	Limit *int `json:"-" url:"limit,omitempty"`
	// Offset values for traversal. Defaults to 0. Max of 1000.
	Offset *int `json:"-" url:"offset,omitempty"`
	// Set minimum depth for traversal. Defaults to 1.
	MinDepth *int `json:"-" url:"min_depth,omitempty"`
	// Set maximum depth for traversal. Defaults to 4.
	MaxDepth *int `json:"-" url:"max_depth,omitempty"`
	// Set relationship type(s) to follow when traversing related entities. Defaults to following 31 relevant relationship types covering ownership, control, and trade.
	Relationships []*Relationships `json:"-" url:"relationships,omitempty"`
	// Also traverse relationships from entities that are possibly the same as any entity that appears in the path. Defaults to traversing possibly same as relationships.
	Psa *bool `json:"-" url:"psa,omitempty"`
	// Filter paths to only those that end at an entity associated with the specified country(ies). Defaults to returning paths that end in any country.
	Countries []*Country `json:"-" url:"countries,omitempty"`
	// Filter paths to only those that end at an entity of the specified type(s). Defaults to returning paths that end at any type.
	Types []*Entities `json:"-" url:"types,omitempty"`
	// Filter paths to only those that end at an entity appearing on a watchlist. Defaults to not filtering paths by sanctioned status.
	Sanctioned *bool `json:"-" url:"sanctioned,omitempty"`
	// Filter paths to only those that end at an entity appearing on a pep list. Defaults to not filtering paths by pep status.
	Pep *bool `json:"-" url:"pep,omitempty"`
	// Set minimum percentage of share ownership for traversal. Defaults to 0.
	MinShares *int `json:"-" url:"min_shares,omitempty"`
	// Also traverse relationships when share percentages are unknown. Only useful when min_shares is set greater than 0. Defaults to true.
	IncludeUnknownShares *bool `json:"-" url:"include_unknown_shares,omitempty"`
	// Include relationships that were valid in the past but not at the present time. Defaults to false.
	ExcludeFormerRelationships *bool `json:"-" url:"exclude_former_relationships,omitempty"`
	// Include entities that existed in the past but not at the present time. Defaults to false.
	ExcludeClosedEntities *bool `json:"-" url:"exclude_closed_entities,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with any risk factor of these categories
	RiskCategories *TraversalRiskCategory `json:"-" url:"risk_categories,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	EuHighRiskThird *bool `json:"-" url:"eu_high_risk_third,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	ReputationalRiskModernSlavery *bool `json:"-" url:"reputational_risk_modern_slavery,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	StateOwned *bool `json:"-" url:"state_owned,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	FormerlySanctioned *bool `json:"-" url:"formerly_sanctioned,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	ReputationalRiskTerrorism *bool `json:"-" url:"reputational_risk_terrorism,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	ReputationalRiskOrganizedCrime *bool `json:"-" url:"reputational_risk_organized_crime,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	ReputationalRiskFinancialCrime *bool `json:"-" url:"reputational_risk_financial_crime,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	ReputationalRiskBriberyAndCorruption *bool `json:"-" url:"reputational_risk_bribery_and_corruption,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	ReputationalRiskOther *bool `json:"-" url:"reputational_risk_other,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	ReputationalRiskCybercrime *bool `json:"-" url:"reputational_risk_cybercrime,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	RegulatoryAction *bool `json:"-" url:"regulatory_action,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	LawEnforcementAction *bool `json:"-" url:"law_enforcement_action,omitempty"`
	// Filter paths to only those that relate with an entity that we have flagged with this risk factor
	XinjiangGeospatial *bool `json:"-" url:"xinjiang_geospatial,omitempty"`
}
