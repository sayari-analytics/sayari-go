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
