// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/sayari-analytics/sayari-go/generated/go/core"
	time "time"
)

type GetEntity struct {
	// The pagination token for the next page of attribute `name`.
	AttributesNameNext *string `json:"-"`
	// The pagination token for the previous page of attribute `name`.
	AttributesNamePrev *string `json:"-"`
	// Limit total values returned for attribute `name`. Defaults to 100.
	AttributesNameLimit *int `json:"-"`
	// The pagination token for the next page of attribute `address`.
	AttributesAddressNext *string `json:"-"`
	// The pagination token for the previous page of attribute `address`.
	AttributesAddressPrev *string `json:"-"`
	// Limit total values returned for attribute `address`. Defaults to 100.
	AttributesAddressLimit *int `json:"-"`
	// The pagination token for the next page of attribute `country`.
	AttributesCountryNext *string `json:"-"`
	// The pagination token for the previous page of attribute `country`.
	AttributesCountryPrev *string `json:"-"`
	// Limit total values returned for attribute `country`. Defaults to 100.
	AttributesCountryLimit *int `json:"-"`
	// The pagination token for the next page of relationship results
	RelationshipsNext *string `json:"-"`
	// The pagination token for the previous page of relationship results
	RelationshipsPrev *string `json:"-"`
	// Limit total relationship values. Defaults to 50.
	RelationshipsLimit *int `json:"-"`
	// Filter relationships to [relationship type](/sayari-library/ontology/relationships), e.g. director_of or has_shareholder
	RelationshipsType *Relationships `json:"-"`
	// Sorts relationships by As Of date or Shareholder percentage, e.g. date or -shares
	RelationshipsSort *string `json:"-"`
	// Filters relationships to after a date
	RelationshipsStartDate *time.Time `json:"-"`
	// Filters relationships to before a date
	RelationshipsEndDate *time.Time `json:"-"`
	// Filters relationships to greater than or equal to a Shareholder percentage
	RelationshipsMinShares *int `json:"-"`
	// Filters relationships to a list of [countries](/sayari-library/ontology/enumerated-types#country)
	RelationshipsCountry []*Country `json:"-"`
	// Filters shipment relationships to a list of arrival [countries](/sayari-library/ontology/enumerated-types#country)
	RelationshipsArrivalCountry []*Country `json:"-"`
	// Filters shipment relationships to an arrival state
	RelationshipsArrivalState *string `json:"-"`
	// Filters shipment relationships to an arrival city
	RelationshipsArrivalCity *string `json:"-"`
	// Filters shipment relationships to a list of departure [countries](/sayari-library/ontology/enumerated-types#country)
	RelationshipsDepartureCountry []*Country `json:"-"`
	// Filters shipment relationships to a departure state
	RelationshipsDepartureState *string `json:"-"`
	// Filters shipment relationships to a departure city
	RelationshipsDepartureCity *string `json:"-"`
	// Filters shipment relationships to a trade partner name
	RelationshipsPartnerName *string `json:"-"`
	// Filters shipment relationships to a trade partner [risk tag](/sayari-library/ontology/enumerated-types#tag)
	RelationshipsPartnerRisk []*Tag `json:"-"`
	// Filters shipment relationships to an HS code
	RelationshipsHsCode *string `json:"-"`
	// The pagination token for the next page of possibly same entities.
	PossiblySameAsNext *string `json:"-"`
	// The pagination token for the previous page of possibly same entities.
	PossiblySameAsPrev *string `json:"-"`
	// Limit total possibly same as entities. Defaults to 100.
	PossiblySameAsLimit *int `json:"-"`
	// The pagination token for the next page of the entity's referencing records
	ReferencedByNext *string `json:"-"`
	// The pagination token for the previous page of the entity's referencing records
	ReferencedByPrev *string `json:"-"`
	// Limit totals values returned for entity's referencing records. Defaults to 100.
	ReferencedByLimit *int `json:"-"`
}

// OK
type EntitySummaryResponse struct {
	// Unique identifier of the entity
	Id string `json:"id"`
	// Display name of the entity
	Label string `json:"label"`
	// Number of outgoing relationships
	Degree int `json:"degree"`
	// True if the entity existed in the past but not at the present time, otherwise false. Always false for data curation.
	Closed bool `json:"closed"`
	// Convenience URL to the entity in the API.
	EntityUrl string `json:"entity_url"`
	// True if the entity has the ["Politically Exposed Person (PEP)" risk factor](/sayari-library/ontology/risk-factors#politically-exposed-person-pep-), otherwise false.
	Pep   bool    `json:"pep"`
	PsaId *string `json:"psa_id,omitempty"`
	// Number of entities that are Possibly the Same As (PSA) the entity.
	PsaCount int `json:"psa_count"`
	// True if the entity has the ["Sanctioned" risk factor](/sayari-library/ontology/risk-factors#sanctioned), otherwise false.
	Sanctioned bool `json:"sanctioned"`
	// The entity type. See detailed explanations [here](/sayari-library/ontology/entities).
	Type        Entities      `json:"type,omitempty"`
	Identifiers []*Identifier `json:"identifiers,omitempty"`
	// Entity [country](/sayari-library/ontology/enumerated-types#country)
	Countries []Country `json:"countries,omitempty"`
	// Number of records associated with the entity, grouped by source.
	SourceCount map[string]*SourceCountInfo `json:"source_count,omitempty"`
	// List of physical addresses associated with the entity. See more [here](/sayari-library/ontology/attributes#address)
	Addresses []string `json:"addresses,omitempty"`
	// Birth date of a person. See more [here](/sayari-library/ontology/attributes#date-of-birth)
	DateOfBirth *string `json:"date_of_birth,omitempty"`
	// Count of related entities for a given relationship type.
	RelationshipCount map[Relationships]int `json:"relationship_count,omitempty"`
	// Count of related entities for a given relationship type.
	UserRelationshipCount    map[Relationships]int   `json:"user_relationship_count,omitempty"`
	AttributeCounts          interface{}             `json:"attribute_counts,omitempty"`
	UserAttributeCounts      interface{}             `json:"user_attribute_counts,omitempty"`
	RelatedEntitiesCount     int                     `json:"related_entities_count"`
	UserRelatedEntitiesCount int                     `json:"user_related_entities_count"`
	UserRecordCount          int                     `json:"user_record_count"`
	RegistrationDate         *EntityRegistrationDate `json:"registration_date,omitempty"`
	TranslatedLabel          *EntityTranslatedLabel  `json:"translated_label,omitempty"`
	HsCode                   *EntityHsCode           `json:"hs_code,omitempty"`
	ShipmentArrival          *ShipmentArrival        `json:"shipment_arrival,omitempty"`
	ShipmentDeparture        *ShipmentDepartue       `json:"shipment_departure,omitempty"`
	CompanyType              *CompanyType            `json:"company_type,omitempty"`
	LatestStatus             *Status                 `json:"latest_status,omitempty"`
	// [Risk factors](/sayari-library/ontology/risk-factors) associated with the entity.
	Risk           EntityRisk           `json:"risk,omitempty"`
	Attributes     *AttributeDetails    `json:"attributes,omitempty"`
	Relationships  *EntityRelationships `json:"relationships,omitempty"`
	PossiblySameAs *PossiblySameAs      `json:"possibly_same_as,omitempty"`
	ReferencedBy   *ReferencedBy        `json:"referenced_by,omitempty"`

	_rawJSON json.RawMessage
}

func (e *EntitySummaryResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler EntitySummaryResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = EntitySummaryResponse(value)
	e._rawJSON = json.RawMessage(data)
	return nil
}

func (e *EntitySummaryResponse) String() string {
	if len(e._rawJSON) > 0 {
		if value, err := core.StringifyJSON(e._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

// OK
type GetEntityResponse struct {
	// Unique identifier of the entity
	Id string `json:"id"`
	// Display name of the entity
	Label string `json:"label"`
	// Number of outgoing relationships
	Degree int `json:"degree"`
	// True if the entity existed in the past but not at the present time, otherwise false. Always false for data curation.
	Closed bool `json:"closed"`
	// Convenience URL to the entity in the API.
	EntityUrl string `json:"entity_url"`
	// True if the entity has the ["Politically Exposed Person (PEP)" risk factor](/sayari-library/ontology/risk-factors#politically-exposed-person-pep-), otherwise false.
	Pep   bool    `json:"pep"`
	PsaId *string `json:"psa_id,omitempty"`
	// Number of entities that are Possibly the Same As (PSA) the entity.
	PsaCount int `json:"psa_count"`
	// True if the entity has the ["Sanctioned" risk factor](/sayari-library/ontology/risk-factors#sanctioned), otherwise false.
	Sanctioned bool `json:"sanctioned"`
	// The entity type. See detailed explanations [here](/sayari-library/ontology/entities).
	Type        Entities      `json:"type,omitempty"`
	Identifiers []*Identifier `json:"identifiers,omitempty"`
	// Entity [country](/sayari-library/ontology/enumerated-types#country)
	Countries []Country `json:"countries,omitempty"`
	// Number of records associated with the entity, grouped by source.
	SourceCount map[string]*SourceCountInfo `json:"source_count,omitempty"`
	// List of physical addresses associated with the entity. See more [here](/sayari-library/ontology/attributes#address)
	Addresses []string `json:"addresses,omitempty"`
	// Birth date of a person. See more [here](/sayari-library/ontology/attributes#date-of-birth)
	DateOfBirth *string `json:"date_of_birth,omitempty"`
	// Count of related entities for a given relationship type.
	RelationshipCount map[Relationships]int `json:"relationship_count,omitempty"`
	// Count of related entities for a given relationship type.
	UserRelationshipCount    map[Relationships]int   `json:"user_relationship_count,omitempty"`
	AttributeCounts          interface{}             `json:"attribute_counts,omitempty"`
	UserAttributeCounts      interface{}             `json:"user_attribute_counts,omitempty"`
	RelatedEntitiesCount     int                     `json:"related_entities_count"`
	UserRelatedEntitiesCount int                     `json:"user_related_entities_count"`
	UserRecordCount          int                     `json:"user_record_count"`
	RegistrationDate         *EntityRegistrationDate `json:"registration_date,omitempty"`
	TranslatedLabel          *EntityTranslatedLabel  `json:"translated_label,omitempty"`
	HsCode                   *EntityHsCode           `json:"hs_code,omitempty"`
	ShipmentArrival          *ShipmentArrival        `json:"shipment_arrival,omitempty"`
	ShipmentDeparture        *ShipmentDepartue       `json:"shipment_departure,omitempty"`
	CompanyType              *CompanyType            `json:"company_type,omitempty"`
	LatestStatus             *Status                 `json:"latest_status,omitempty"`
	// [Risk factors](/sayari-library/ontology/risk-factors) associated with the entity.
	Risk           EntityRisk           `json:"risk,omitempty"`
	Attributes     *AttributeDetails    `json:"attributes,omitempty"`
	Relationships  *EntityRelationships `json:"relationships,omitempty"`
	PossiblySameAs *PossiblySameAs      `json:"possibly_same_as,omitempty"`
	ReferencedBy   *ReferencedBy        `json:"referenced_by,omitempty"`

	_rawJSON json.RawMessage
}

func (g *GetEntityResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GetEntityResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetEntityResponse(value)
	g._rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetEntityResponse) String() string {
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
