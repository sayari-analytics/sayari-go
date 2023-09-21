// This file was auto-generated by Fern from our API Definition.

package api

import (
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
	// Limit total relationship values. Defaults to 100.
	RelationshipsLimit *int `json:"-"`
	// Filter relationships to relationship type, e.g. director_of or has_shareholder
	RelationshipsType *string `json:"-"`
	// Sorts relationships by As Of date or Shareholder percentage, e.g. date or -shares
	RelationshipsSort *string `json:"-"`
	// Filters relationships to after a date
	RelationshipsStartDate *time.Time `json:"-"`
	// Filters relationships to before a date
	RelationshipsEndDate *time.Time `json:"-"`
	// Filters relationships to greater than or equal to a Shareholder percentage
	RelationshipsMinShares *int `json:"-"`
	// Filters relationships to a list of countries
	RelationshipsCountry *string `json:"-"`
	// Filters shipment relationships to a list of arrival countries
	RelationshipsArrivalCountry *string `json:"-"`
	// Filters shipment relationships to a list of departure countries
	RelationshipsDepartureCountry *string `json:"-"`
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

type EntityDetails struct {
	Id                EntityId           `json:"id"`
	Label             string             `json:"label"`
	Degree            int                `json:"degree"`
	Closed            bool               `json:"closed"`
	EntityUrl         string             `json:"entity_url"`
	Pep               bool               `json:"pep"`
	PsaCount          int                `json:"psa_count"`
	Sanctioned        bool               `json:"sanctioned"`
	RegistrationDate  string             `json:"registration_date"`
	LatestStatus      *Status            `json:"latest_status,omitempty"`
	Type              EntityType         `json:"type,omitempty"`
	Identifiers       []*Identifier      `json:"identifiers,omitempty"`
	Addresses         []string           `json:"addresses,omitempty"`
	Countries         []string           `json:"countries,omitempty"`
	RelationshipCount *RelationshipCount `json:"relationship_count,omitempty"`
	SourceCount       SourceCount        `json:"source_count,omitempty"`
	Risk              Risk               `json:"risk,omitempty"`
	Attributes        *Attributes        `json:"attributes,omitempty"`
	Relationships     *Relationships     `json:"relationships,omitempty"`
	PossiblySameAs    *PossiblySameAs    `json:"possibly_same_as,omitempty"`
	ReferencedBy      *ReferencedBy      `json:"referenced_by,omitempty"`
}

// The unique identifier for an entity in the database
type EntityId = string
