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
	// Shipment fields to search against.
	Fields []ShipmentField `json:"fields,omitempty"`
	// Filters to be applied to search query to limit the result-set.
	Filter *TradeFilterList `json:"filter,omitempty"`
	// Whether or not to return search facets in results giving counts by field. Defaults to false.
	Facets *bool `json:"facets,omitempty"`
	// Whether or not to return search geo bound facets in results giving counts by geo tile. Defaults to false.
	GeoFacets *bool `json:"geo_facets,omitempty"`
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
	// Shipment fields to search against.
	Fields []ShipmentField `json:"fields,omitempty"`
	// Filters to be applied to search query to limit the result-set.
	Filter *TradeFilterList `json:"filter,omitempty"`
	// Whether or not to return search facets in results giving counts by field. Defaults to false.
	Facets *bool `json:"facets,omitempty"`
	// Whether or not to return search geo bound facets in results giving counts by geo tile. Defaults to false.
	GeoFacets *bool `json:"geo_facets,omitempty"`
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
	// Shipment fields to search against.
	Fields []ShipmentField `json:"fields,omitempty"`
	// Filters to be applied to search query to limit the result-set.
	Filter *TradeFilterList `json:"filter,omitempty"`
	// Whether or not to return search facets in results giving counts by field. Defaults to false.
	Facets *bool `json:"facets,omitempty"`
	// Whether or not to return search geo bound facets in results giving counts by geo tile. Defaults to false.
	GeoFacets *bool `json:"geo_facets,omitempty"`
	// Set to true to enable full elasticsearch query string syntax which allows for fielded search and more complex operators. Note that the syntax is more strict and can result in empty result-sets. Defaults to false.
	Advanced *bool `json:"advanced,omitempty"`
}

type ShipmentField string

const (
	ShipmentFieldBuyerName        ShipmentField = "buyer_name"
	ShipmentFieldSupplierName     ShipmentField = "supplier_name"
	ShipmentFieldBuyerRisk        ShipmentField = "buyer_risk"
	ShipmentFieldSupplierRisk     ShipmentField = "supplier_risk"
	ShipmentFieldBuyerCountry     ShipmentField = "buyer_country"
	ShipmentFieldSupplierCountry  ShipmentField = "supplier_country"
	ShipmentFieldDepartureCountry ShipmentField = "departure_country"
	ShipmentFieldDepartureState   ShipmentField = "departure_state"
	ShipmentFieldDepartureCity    ShipmentField = "departure_city"
	ShipmentFieldArrivalCountry   ShipmentField = "arrival_country"
	ShipmentFieldArrivalState     ShipmentField = "arrival_state"
	ShipmentFieldArrivalCity      ShipmentField = "arrival_city"
	ShipmentFieldHsCode           ShipmentField = "hs_code"
	ShipmentFieldHsDescription    ShipmentField = "hs_description"
	ShipmentFieldSupplierPurpose  ShipmentField = "supplier_purpose"
	ShipmentFieldBuyerPurpose     ShipmentField = "buyer_purpose"
	ShipmentFieldArrivalDate      ShipmentField = "arrival_date"
	ShipmentFieldWeight           ShipmentField = "weight"
)

func NewShipmentFieldFromString(s string) (ShipmentField, error) {
	switch s {
	case "buyer_name":
		return ShipmentFieldBuyerName, nil
	case "supplier_name":
		return ShipmentFieldSupplierName, nil
	case "buyer_risk":
		return ShipmentFieldBuyerRisk, nil
	case "supplier_risk":
		return ShipmentFieldSupplierRisk, nil
	case "buyer_country":
		return ShipmentFieldBuyerCountry, nil
	case "supplier_country":
		return ShipmentFieldSupplierCountry, nil
	case "departure_country":
		return ShipmentFieldDepartureCountry, nil
	case "departure_state":
		return ShipmentFieldDepartureState, nil
	case "departure_city":
		return ShipmentFieldDepartureCity, nil
	case "arrival_country":
		return ShipmentFieldArrivalCountry, nil
	case "arrival_state":
		return ShipmentFieldArrivalState, nil
	case "arrival_city":
		return ShipmentFieldArrivalCity, nil
	case "hs_code":
		return ShipmentFieldHsCode, nil
	case "hs_description":
		return ShipmentFieldHsDescription, nil
	case "supplier_purpose":
		return ShipmentFieldSupplierPurpose, nil
	case "buyer_purpose":
		return ShipmentFieldBuyerPurpose, nil
	case "arrival_date":
		return ShipmentFieldArrivalDate, nil
	case "weight":
		return ShipmentFieldWeight, nil
	}
	var t ShipmentField
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (s ShipmentField) Ptr() *ShipmentField {
	return &s
}

type BuyerSearchResults struct {
	Limit  int                  `json:"limit"`
	Size   *SizeInfo            `json:"size,omitempty"`
	Offset int                  `json:"offset"`
	Next   bool                 `json:"next"`
	Data   *SupplierOrBuyerHits `json:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (b *BuyerSearchResults) UnmarshalJSON(data []byte) error {
	type unmarshaler BuyerSearchResults
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BuyerSearchResults(value)
	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BuyerSearchResults) String() string {
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

type ShipmentSearchResults struct {
	Limit  int           `json:"limit"`
	Size   *SizeInfo     `json:"size,omitempty"`
	Offset int           `json:"offset"`
	Next   bool          `json:"next"`
	Data   *ShipmentHits `json:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (s *ShipmentSearchResults) UnmarshalJSON(data []byte) error {
	type unmarshaler ShipmentSearchResults
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = ShipmentSearchResults(value)
	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *ShipmentSearchResults) String() string {
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

type SupplierSearchResults struct {
	Limit  int                  `json:"limit"`
	Size   *SizeInfo            `json:"size,omitempty"`
	Offset int                  `json:"offset"`
	Next   bool                 `json:"next"`
	Data   *SupplierOrBuyerHits `json:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (s *SupplierSearchResults) UnmarshalJSON(data []byte) error {
	type unmarshaler SupplierSearchResults
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SupplierSearchResults(value)
	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SupplierSearchResults) String() string {
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
	BuyerId          []string  `json:"buyer_id,omitempty"`
	SupplierId       []string  `json:"supplier_id,omitempty"`
	BuyerName        []string  `json:"buyer_name,omitempty"`
	SupplierName     []string  `json:"supplier_name,omitempty"`
	BuyerRisk        []string  `json:"buyer_risk,omitempty"`
	SupplierRisk     []string  `json:"supplier_risk,omitempty"`
	BuyerCountry     []Country `json:"buyer_country,omitempty"`
	SupplierCountry  []Country `json:"supplier_country,omitempty"`
	DepartureCountry []Country `json:"departure_country,omitempty"`
	DepartureState   []string  `json:"departure_state,omitempty"`
	DepartureCity    []string  `json:"departure_city,omitempty"`
	ArrivalCountry   []Country `json:"arrival_country,omitempty"`
	ArrivalState     []string  `json:"arrival_state,omitempty"`
	ArrivalCity      []string  `json:"arrival_city,omitempty"`
	HsCode           []string  `json:"hs_code,omitempty"`
	HsDescription    []string  `json:"hs_description,omitempty"`
	SupplierPurpose  []string  `json:"supplier_purpose,omitempty"`
	BuyerPurpose     []string  `json:"buyer_purpose,omitempty"`
	ArrivalDate      []string  `json:"arrival_date,omitempty"`
	Weight           []string  `json:"weight,omitempty"`
	Sources          []string  `json:"sources,omitempty"`

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
