// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/sayari-analytics/sayari-go/generated/go/core"
	time "time"
)

type GetEntity struct {
	// The pagination token for the next page of attribute `additional_information`.
	AttributesAdditionalInformationNext *string `json:"-" url:"attributes.additional_information.next,omitempty"`
	// The pagination token for the previous page of attribute `additional_information`.
	AttributesAdditionalInformationPrev *string `json:"-" url:"attributes.additional_information.prev,omitempty"`
	// Limit total values returned for attribute `additional_information`. Defaults to 100.
	AttributesAdditionalInformationLimit *int `json:"-" url:"attributes.additional_information.limit,omitempty"`
	// The pagination token for the next page of attribute `address`.
	AttributesAddressNext *string `json:"-" url:"attributes.address.next,omitempty"`
	// The pagination token for the previous page of attribute `address`.
	AttributesAddressPrev *string `json:"-" url:"attributes.address.prev,omitempty"`
	// Limit total values returned for attribute `address`. Defaults to 100.
	AttributesAddressLimit *int `json:"-" url:"attributes.address.limit,omitempty"`
	// The pagination token for the next page of attribute `business_purpose`.
	AttributesBusinessPurposeNext *string `json:"-" url:"attributes.business_purpose.next,omitempty"`
	// The pagination token for the previous page of attribute `business_purpose`.
	AttributesBusinessPurposePrev *string `json:"-" url:"attributes.business_purpose.prev,omitempty"`
	// Limit total values returned for attribute `business_purpose`. Defaults to 100.
	AttributesBusinessPurposeLimit *int `json:"-" url:"attributes.business_purpose.limit,omitempty"`
	// The pagination token for the next page of attribute `company_type`.
	AttributesCompanyTypeNext *string `json:"-" url:"attributes.company_type.next,omitempty"`
	// The pagination token for the previous page of attribute `company_type`.
	AttributesCompanyTypePrev *string `json:"-" url:"attributes.company_type.prev,omitempty"`
	// Limit total values returned for attribute `company_type`. Defaults to 100.
	AttributesCompanyTypeLimit *int `json:"-" url:"attributes.company_type.limit,omitempty"`
	// The pagination token for the next page of attribute `country`.
	AttributesCountryNext *string `json:"-" url:"attributes.country.next,omitempty"`
	// The pagination token for the previous page of attribute `country`.
	AttributesCountryPrev *string `json:"-" url:"attributes.country.prev,omitempty"`
	// Limit total values returned for attribute `country`. Defaults to 100.
	AttributesCountryLimit *int `json:"-" url:"attributes.country.limit,omitempty"`
	// The pagination token for the next page of attribute `identifier`.
	AttributesIdentifierNext *string `json:"-" url:"attributes.identifier.next,omitempty"`
	// The pagination token for the previous page of attribute `identifier`.
	AttributesIdentifierPrev *string `json:"-" url:"attributes.identifier.prev,omitempty"`
	// Limit total values returned for attribute `identifier`. Defaults to 100.
	AttributesIdentifierLimit *int `json:"-" url:"attributes.identifier.limit,omitempty"`
	// The pagination token for the next page of attribute `name`.
	AttributesNameNext *string `json:"-" url:"attributes.name.next,omitempty"`
	// The pagination token for the previous page of attribute `name`.
	AttributesNamePrev *string `json:"-" url:"attributes.name.prev,omitempty"`
	// Limit total values returned for attribute `name`. Defaults to 100.
	AttributesNameLimit *int `json:"-" url:"attributes.name.limit,omitempty"`
	// The pagination token for the next page of attribute `status`.
	AttributesStatusNext *string `json:"-" url:"attributes.status.next,omitempty"`
	// The pagination token for the previous page of attribute `status`.
	AttributesStatusPrev *string `json:"-" url:"attributes.status.prev,omitempty"`
	// Limit total values returned for attribute `status`. Defaults to 100.
	AttributesStatusLimit *int `json:"-" url:"attributes.status.limit,omitempty"`
	// The pagination token for the next page of relationship results
	RelationshipsNext *string `json:"-" url:"relationships.next,omitempty"`
	// The pagination token for the previous page of relationship results
	RelationshipsPrev *string `json:"-" url:"relationships.prev,omitempty"`
	// Limit total relationship values. Defaults to 50.
	RelationshipsLimit *int `json:"-" url:"relationships.limit,omitempty"`
	// Filter relationships to [relationship type](/sayari-library/ontology/relationships), e.g. director_of or has_shareholder
	RelationshipsType *Relationships `json:"-" url:"relationships.type,omitempty"`
	// Sorts relationships by As Of date or Shareholder percentage, e.g. date or -shares
	RelationshipsSort *string `json:"-" url:"relationships.sort,omitempty"`
	// Filters relationships to after a date
	RelationshipsStartDate *time.Time `json:"-" url:"relationships.startDate,omitempty" format:"date"`
	// Filters relationships to before a date
	RelationshipsEndDate *time.Time `json:"-" url:"relationships.endDate,omitempty" format:"date"`
	// Filters relationships to greater than or equal to a Shareholder percentage
	RelationshipsMinShares *int `json:"-" url:"relationships.minShares,omitempty"`
	// Filters relationships to a list of [countries](/sayari-library/ontology/enumerated-types#country)
	RelationshipsCountry []*Country `json:"-" url:"relationships.country,omitempty"`
	// Filters shipment relationships to a list of arrival [countries](/sayari-library/ontology/enumerated-types#country)
	RelationshipsArrivalCountry []*Country `json:"-" url:"relationships.arrivalCountry,omitempty"`
	// Filters shipment relationships to an arrival state
	RelationshipsArrivalState *string `json:"-" url:"relationships.arrivalState,omitempty"`
	// Filters shipment relationships to an arrival city
	RelationshipsArrivalCity *string `json:"-" url:"relationships.arrivalCity,omitempty"`
	// Filters shipment relationships to a list of departure [countries](/sayari-library/ontology/enumerated-types#country)
	RelationshipsDepartureCountry []*Country `json:"-" url:"relationships.departureCountry,omitempty"`
	// Filters shipment relationships to a departure state
	RelationshipsDepartureState *string `json:"-" url:"relationships.departureState,omitempty"`
	// Filters shipment relationships to a departure city
	RelationshipsDepartureCity *string `json:"-" url:"relationships.departureCity,omitempty"`
	// Filters shipment relationships to a trade partner name
	RelationshipsPartnerName *string `json:"-" url:"relationships.partnerName,omitempty"`
	// Filters shipment relationships to a trade partner [risk tag](/sayari-library/ontology/enumerated-types#tag)
	RelationshipsPartnerRisk []*Risk `json:"-" url:"relationships.partnerRisk,omitempty"`
	// Filters shipment relationships to an HS code
	RelationshipsHsCode *string `json:"-" url:"relationships.hsCode,omitempty"`
	// The pagination token for the next page of possibly same entities.
	PossiblySameAsNext *string `json:"-" url:"possibly_same_as.next,omitempty"`
	// The pagination token for the previous page of possibly same entities.
	PossiblySameAsPrev *string `json:"-" url:"possibly_same_as.prev,omitempty"`
	// Limit total possibly same as entities. Defaults to 100.
	PossiblySameAsLimit *int `json:"-" url:"possibly_same_as.limit,omitempty"`
	// The pagination token for the next page of the entity's referencing records
	ReferencedByNext *string `json:"-" url:"referenced_by.next,omitempty"`
	// The pagination token for the previous page of the entity's referencing records
	ReferencedByPrev *string `json:"-" url:"referenced_by.prev,omitempty"`
	// Limit totals values returned for entity's referencing records. Defaults to 100.
	ReferencedByLimit *int `json:"-" url:"referenced_by.limit,omitempty"`
}

// OK
type EntitySummaryResponse struct {
	// Unique identifier of the entity
	Id string `json:"id" url:"id"`
	// Display name of the entity
	Label string `json:"label" url:"label"`
	// Number of outgoing relationships
	Degree int `json:"degree" url:"degree"`
	// True if the entity existed in the past but not at the present time, otherwise false. Always false for data curation.
	Closed bool `json:"closed" url:"closed"`
	// Convenience URL to the entity in the API.
	EntityUrl string `json:"entity_url" url:"entity_url"`
	// True if the entity has the ["Politically Exposed Person (PEP)" risk factor](/sayari-library/ontology/risk-factors#politically-exposed-person-pep-), otherwise false.
	Pep   bool    `json:"pep" url:"pep"`
	PsaId *string `json:"psa_id,omitempty" url:"psa_id,omitempty"`
	// Number of entities that are Possibly the Same As (PSA) the entity.
	PsaCount int `json:"psa_count" url:"psa_count"`
	// True if the entity has the ["Sanctioned" risk factor](/sayari-library/ontology/risk-factors#sanctioned), otherwise false.
	Sanctioned bool `json:"sanctioned" url:"sanctioned"`
	// The [entity type](/sayari-library/ontology/entities).
	Type        Entities      `json:"type" url:"type"`
	Identifiers []*Identifier `json:"identifiers,omitempty" url:"identifiers,omitempty"`
	// Entity [country](/sayari-library/ontology/enumerated-types#country)
	Countries []Country `json:"countries,omitempty" url:"countries,omitempty"`
	// Number of records associated with the entity, grouped by source.
	SourceCount map[string]*SourceCountInfo `json:"source_count,omitempty" url:"source_count,omitempty"`
	// List of physical addresses associated with the entity. See more [here](/sayari-library/ontology/attributes#address)
	Addresses  []string       `json:"addresses,omitempty" url:"addresses,omitempty"`
	TradeCount map[string]int `json:"trade_count,omitempty" url:"trade_count,omitempty"`
	// Birth date of a person. See more [here](/sayari-library/ontology/attributes#date-of-birth)
	DateOfBirth           *string           `json:"date_of_birth,omitempty" url:"date_of_birth,omitempty"`
	RelationshipCount     RelationshipCount `json:"relationship_count,omitempty" url:"relationship_count,omitempty"`
	UserRelationshipCount RelationshipCount `json:"user_relationship_count,omitempty" url:"user_relationship_count,omitempty"`
	// Count of attributes for a given [attribute type](/sayari-library/ontology/attributes)
	AttributeCount map[Attributes]int `json:"attribute_count,omitempty" url:"attribute_count,omitempty"`
	// Count of user-created attributes for a given [attribute type](/sayari-library/ontology/attributes)
	UserAttributeCount map[Attributes]int `json:"user_attribute_count,omitempty" url:"user_attribute_count,omitempty"`
	// Count of attributes for a given [attribute type](/sayari-library/ontology/attributes)
	AttributeCounts map[Attributes]int `json:"attribute_counts,omitempty" url:"attribute_counts,omitempty"`
	// Count of user-created attributes for a given [attribute type](/sayari-library/ontology/attributes)
	UserAttributeCounts      map[Attributes]int      `json:"user_attribute_counts,omitempty" url:"user_attribute_counts,omitempty"`
	RelatedEntitiesCount     int                     `json:"related_entities_count" url:"related_entities_count"`
	UserRelatedEntitiesCount int                     `json:"user_related_entities_count" url:"user_related_entities_count"`
	UserRecordCount          int                     `json:"user_record_count" url:"user_record_count"`
	ReferenceId              *string                 `json:"reference_id,omitempty" url:"reference_id,omitempty"`
	RegistrationDate         *EntityRegistrationDate `json:"registration_date,omitempty" url:"registration_date,omitempty"`
	TranslatedLabel          *EntityTranslatedLabel  `json:"translated_label,omitempty" url:"translated_label,omitempty"`
	HsCode                   *EntityHsCode           `json:"hs_code,omitempty" url:"hs_code,omitempty"`
	ShipmentArrival          *ShipmentArrival        `json:"shipment_arrival,omitempty" url:"shipment_arrival,omitempty"`
	ShipmentDeparture        *ShipmentDeparture      `json:"shipment_departure,omitempty" url:"shipment_departure,omitempty"`
	CompanyType              *CompanyType            `json:"company_type,omitempty" url:"company_type,omitempty"`
	LatestStatus             *Status                 `json:"latest_status,omitempty" url:"latest_status,omitempty"`
	// [Risk factors](/sayari-library/ontology/risk-factors) associated with the entity.
	Risk EntityRisk `json:"risk,omitempty" url:"risk,omitempty"`
	// Detailed information about the entity's [attributes](/sayari-library/ontology/attributes).
	Attributes *AttributeDetails `json:"attributes,omitempty" url:"attributes,omitempty"`
	// Detailed information about the entity's [relationships](/sayari-library/ontology/relationships).
	Relationships  *EntityRelationships `json:"relationships,omitempty" url:"relationships,omitempty"`
	PossiblySameAs *PossiblySameAs      `json:"possibly_same_as,omitempty" url:"possibly_same_as,omitempty"`
	ReferencedBy   *ReferencedBy        `json:"referenced_by,omitempty" url:"referenced_by,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (e *EntitySummaryResponse) GetExtraProperties() map[string]interface{} {
	return e.extraProperties
}

func (e *EntitySummaryResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler EntitySummaryResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = EntitySummaryResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties

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
	Id string `json:"id" url:"id"`
	// Display name of the entity
	Label string `json:"label" url:"label"`
	// Number of outgoing relationships
	Degree int `json:"degree" url:"degree"`
	// True if the entity existed in the past but not at the present time, otherwise false. Always false for data curation.
	Closed bool `json:"closed" url:"closed"`
	// Convenience URL to the entity in the API.
	EntityUrl string `json:"entity_url" url:"entity_url"`
	// True if the entity has the ["Politically Exposed Person (PEP)" risk factor](/sayari-library/ontology/risk-factors#politically-exposed-person-pep-), otherwise false.
	Pep   bool    `json:"pep" url:"pep"`
	PsaId *string `json:"psa_id,omitempty" url:"psa_id,omitempty"`
	// Number of entities that are Possibly the Same As (PSA) the entity.
	PsaCount int `json:"psa_count" url:"psa_count"`
	// True if the entity has the ["Sanctioned" risk factor](/sayari-library/ontology/risk-factors#sanctioned), otherwise false.
	Sanctioned bool `json:"sanctioned" url:"sanctioned"`
	// The [entity type](/sayari-library/ontology/entities).
	Type        Entities      `json:"type" url:"type"`
	Identifiers []*Identifier `json:"identifiers,omitempty" url:"identifiers,omitempty"`
	// Entity [country](/sayari-library/ontology/enumerated-types#country)
	Countries []Country `json:"countries,omitempty" url:"countries,omitempty"`
	// Number of records associated with the entity, grouped by source.
	SourceCount map[string]*SourceCountInfo `json:"source_count,omitempty" url:"source_count,omitempty"`
	// List of physical addresses associated with the entity. See more [here](/sayari-library/ontology/attributes#address)
	Addresses  []string       `json:"addresses,omitempty" url:"addresses,omitempty"`
	TradeCount map[string]int `json:"trade_count,omitempty" url:"trade_count,omitempty"`
	// Birth date of a person. See more [here](/sayari-library/ontology/attributes#date-of-birth)
	DateOfBirth           *string           `json:"date_of_birth,omitempty" url:"date_of_birth,omitempty"`
	RelationshipCount     RelationshipCount `json:"relationship_count,omitempty" url:"relationship_count,omitempty"`
	UserRelationshipCount RelationshipCount `json:"user_relationship_count,omitempty" url:"user_relationship_count,omitempty"`
	// Count of attributes for a given [attribute type](/sayari-library/ontology/attributes)
	AttributeCount map[Attributes]int `json:"attribute_count,omitempty" url:"attribute_count,omitempty"`
	// Count of user-created attributes for a given [attribute type](/sayari-library/ontology/attributes)
	UserAttributeCount map[Attributes]int `json:"user_attribute_count,omitempty" url:"user_attribute_count,omitempty"`
	// Count of attributes for a given [attribute type](/sayari-library/ontology/attributes)
	AttributeCounts map[Attributes]int `json:"attribute_counts,omitempty" url:"attribute_counts,omitempty"`
	// Count of user-created attributes for a given [attribute type](/sayari-library/ontology/attributes)
	UserAttributeCounts      map[Attributes]int      `json:"user_attribute_counts,omitempty" url:"user_attribute_counts,omitempty"`
	RelatedEntitiesCount     int                     `json:"related_entities_count" url:"related_entities_count"`
	UserRelatedEntitiesCount int                     `json:"user_related_entities_count" url:"user_related_entities_count"`
	UserRecordCount          int                     `json:"user_record_count" url:"user_record_count"`
	ReferenceId              *string                 `json:"reference_id,omitempty" url:"reference_id,omitempty"`
	RegistrationDate         *EntityRegistrationDate `json:"registration_date,omitempty" url:"registration_date,omitempty"`
	TranslatedLabel          *EntityTranslatedLabel  `json:"translated_label,omitempty" url:"translated_label,omitempty"`
	HsCode                   *EntityHsCode           `json:"hs_code,omitempty" url:"hs_code,omitempty"`
	ShipmentArrival          *ShipmentArrival        `json:"shipment_arrival,omitempty" url:"shipment_arrival,omitempty"`
	ShipmentDeparture        *ShipmentDeparture      `json:"shipment_departure,omitempty" url:"shipment_departure,omitempty"`
	CompanyType              *CompanyType            `json:"company_type,omitempty" url:"company_type,omitempty"`
	LatestStatus             *Status                 `json:"latest_status,omitempty" url:"latest_status,omitempty"`
	// [Risk factors](/sayari-library/ontology/risk-factors) associated with the entity.
	Risk EntityRisk `json:"risk,omitempty" url:"risk,omitempty"`
	// Detailed information about the entity's [attributes](/sayari-library/ontology/attributes).
	Attributes *AttributeDetails `json:"attributes,omitempty" url:"attributes,omitempty"`
	// Detailed information about the entity's [relationships](/sayari-library/ontology/relationships).
	Relationships  *EntityRelationships `json:"relationships,omitempty" url:"relationships,omitempty"`
	PossiblySameAs *PossiblySameAs      `json:"possibly_same_as,omitempty" url:"possibly_same_as,omitempty"`
	ReferencedBy   *ReferencedBy        `json:"referenced_by,omitempty" url:"referenced_by,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (g *GetEntityResponse) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *GetEntityResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GetEntityResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetEntityResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties

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
