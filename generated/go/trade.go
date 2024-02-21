// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/sayari-analytics/sayari-go/generated/go/core"
)

type SearchBuyers struct {
	// A limit on the number of objects to be returned with a range between 1 and 100. Defaults to 100.
	Limit *int `json:"-"`
	// Number of results to skip before returning response. Defaults to 0.
	Offset *int `json:"-"`
	// Query term. The syntax for the query parameter follows elasticsearch simple query string syntax. The includes the ability to use search operators and to perform nested queries. Must be url encoded.
	Q string `json:"q"`
	// Filters to be applied to search query to limit the result-set.
	Filter *TradeFilterList `json:"filter,omitempty"`
	// Whether or not to return search facets in results giving counts by field. Defaults to false.
	Facets *bool `json:"facets,omitempty"`
	// Set to true to enable full elasticsearch query string syntax which allows for fielded search and more complex operators. Note that the syntax is more strict and can result in empty result-sets. Defaults to false.
	Advanced *bool `json:"advanced,omitempty"`
}

type SearchShipments struct {
	// A limit on the number of objects to be returned with a range between 1 and 100. Defaults to 100.
	Limit *int `json:"-"`
	// Number of results to skip before returning response. Defaults to 0.
	Offset *int `json:"-"`
	// Query term. The syntax for the query parameter follows elasticsearch simple query string syntax. The includes the ability to use search operators and to perform nested queries. Must be url encoded.
	Q string `json:"q"`
	// Filters to be applied to search query to limit the result-set.
	Filter *TradeFilterList `json:"filter,omitempty"`
	// Whether or not to return search facets in results giving counts by field. Defaults to false.
	Facets *bool `json:"facets,omitempty"`
	// Set to true to enable full elasticsearch query string syntax which allows for fielded search and more complex operators. Note that the syntax is more strict and can result in empty result-sets. Defaults to false.
	Advanced *bool `json:"advanced,omitempty"`
}

type SearchSuppliers struct {
	// A limit on the number of objects to be returned with a range between 1 and 100. Defaults to 100.
	Limit *int `json:"-"`
	// Number of results to skip before returning response. Defaults to 0.
	Offset *int `json:"-"`
	// Query term. The syntax for the query parameter follows elasticsearch simple query string syntax. The includes the ability to use search operators and to perform nested queries. Must be url encoded.
	Q string `json:"q"`
	// Filters to be applied to search query to limit the result-set.
	Filter *TradeFilterList `json:"filter,omitempty"`
	// Whether or not to return search facets in results giving counts by field. Defaults to false.
	Facets *bool `json:"facets,omitempty"`
	// Set to true to enable full elasticsearch query string syntax which allows for fielded search and more complex operators. Note that the syntax is more strict and can result in empty result-sets. Defaults to false.
	Advanced *bool `json:"advanced,omitempty"`
}

// OK
type BuyerSearchResponse struct {
	Limit  int                `json:"limit"`
	Size   *SizeInfo          `json:"size,omitempty"`
	Offset int                `json:"offset"`
	Next   bool               `json:"next"`
	Data   []*SupplierOrBuyer `json:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (b *BuyerSearchResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler BuyerSearchResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BuyerSearchResponse(value)
	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BuyerSearchResponse) String() string {
	if len(b._rawJSON) > 0 {
		if value, err := core.StringifyJSON(b._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(b); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", b)
}

// OK
type ShipmentSearchResponse struct {
	Limit  int         `json:"limit"`
	Size   *SizeInfo   `json:"size,omitempty"`
	Offset int         `json:"offset"`
	Next   bool        `json:"next"`
	Data   []*Shipment `json:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (s *ShipmentSearchResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ShipmentSearchResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = ShipmentSearchResponse(value)
	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *ShipmentSearchResponse) String() string {
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

// OK
type SupplierSearchResponse struct {
	Limit  int                `json:"limit"`
	Size   *SizeInfo          `json:"size,omitempty"`
	Offset int                `json:"offset"`
	Next   bool               `json:"next"`
	Data   []*SupplierOrBuyer `json:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (s *SupplierSearchResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler SupplierSearchResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SupplierSearchResponse(value)
	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SupplierSearchResponse) String() string {
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

// Filter your search on the following attributes.
type TradeFilterList struct {
	// Exact match against the entity_id of the buyer. The buyer is the receiver_of shipments.
	BuyerId []string `json:"buyer_id,omitempty"`
	// Exact match against the entity_id of the supplier. The supplier is the shipper_of shipments.
	SupplierId []string `json:"supplier_id,omitempty"`
	// Buyers whose name contains the provided string.
	BuyerName []string `json:"buyer_name,omitempty"`
	// Shipper whose name contains the provided string.
	SupplierName []string `json:"supplier_name,omitempty"`
	// Buyer with an exact match for the provided [risk factor](/sayari-library/ontology/risk-factors).
	BuyerRisk []Risk `json:"buyer_risk,omitempty"`
	// Shipper with an exact match for the provided [risk factor](/sayari-library/ontology/risk-factors).
	SupplierRisk []Risk `json:"supplier_risk,omitempty"`
	// Buyer with an exact match for the provided [country code](/sayari-library/ontology/enumerated-types#country).
	BuyerCountry []Country `json:"buyer_country,omitempty"`
	// Supplier with an exact match for the provided [country code](/sayari-library/ontology/enumerated-types#country).
	SupplierCountry []Country `json:"supplier_country,omitempty"`
	// Shipment departs from a country with an exact match for the provided [country code](/sayari-library/ontology/enumerated-types#country).
	DepartureCountry []Country `json:"departure_country,omitempty"`
	// Shipment departs from a state that contains the provided state name.
	DepartureState []string `json:"departure_state,omitempty"`
	// Shipment departs from a city that contains the provided city name.
	DepartureCity []string `json:"departure_city,omitempty"`
	// Shipment arrives at a country with an exact match for the provided [country code](/sayari-library/ontology/enumerated-types#country).
	ArrivalCountry []Country `json:"arrival_country,omitempty"`
	// Shipment arrives at a state that contains the provided state name.
	ArrivalState []string `json:"arrival_state,omitempty"`
	// Shipment arrives at a city that contains the provided city name.
	ArrivalCity []string `json:"arrival_city,omitempty"`
	// The shipment HS code starts with the provided HS code.
	HsCode []string `json:"hs_code,omitempty"`
	// The HS description contains the provided string.
	HsDescription []string `json:"hs_description,omitempty"`
	// The supplier purpose contains the provided string.
	SupplierPurpose []string `json:"supplier_purpose,omitempty"`
	// The buyer purpose contains the provided string.
	BuyerPurpose []string `json:"buyer_purpose,omitempty"`
	// The arrival date is within the provided range.
	ArrivalDate []string `json:"arrival_date,omitempty"`
	// The departure date is within the provided range.
	DepartureDate []string `json:"departure_date,omitempty"`
	// The shipment identifier starts with the provided string.
	ShipmentIdentifier []string `json:"shipment_identifier,omitempty"`
	// The shipment weight is within the provided range.
	Weight []string `json:"weight,omitempty"`
	// An exact match for the provided sources.
	Sources []string `json:"sources,omitempty"`

	_rawJSON json.RawMessage
}

func (t *TradeFilterList) UnmarshalJSON(data []byte) error {
	type unmarshaler TradeFilterList
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TradeFilterList(value)
	t._rawJSON = json.RawMessage(data)
	return nil
}

func (t *TradeFilterList) String() string {
	if len(t._rawJSON) > 0 {
		if value, err := core.StringifyJSON(t._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}
