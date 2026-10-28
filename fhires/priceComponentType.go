package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// Codes indicating the kind of the price component.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/price-component-type
type PriceComponentType int

const (
	PriceComponentTypeBase PriceComponentType = iota
	// the amount is a surcharge applied on the base price.
	PriceComponentTypeSurcharge
	// the amount is a deduction applied on the base price.
	PriceComponentTypeDeduction
	// the amount is a discount applied on the base price.
	PriceComponentTypeDiscount
	// the amount is the tax component of the total price.
	PriceComponentTypeTax
	// the amount is of informational character, it has not been applied in the calculation of the total price.
	PriceComponentTypeInformational
)

func (code PriceComponentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *PriceComponentType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "base":
		*code = PriceComponentTypeBase
	case "surcharge":
		*code = PriceComponentTypeSurcharge
	case "deduction":
		*code = PriceComponentTypeDeduction
	case "discount":
		*code = PriceComponentTypeDiscount
	case "tax":
		*code = PriceComponentTypeTax
	case "informational":
		*code = PriceComponentTypeInformational
	default:
		return fmt.Errorf("unknown PriceComponentType code `%s`", s)
	}
	return nil
}

func (code PriceComponentType) String() string {
	return code.Code()
}

func (code PriceComponentType) Code() string {
	switch code {
	case PriceComponentTypeBase:
		return "base"
	case PriceComponentTypeSurcharge:
		return "surcharge"
	case PriceComponentTypeDeduction:
		return "deduction"
	case PriceComponentTypeDiscount:
		return "discount"
	case PriceComponentTypeTax:
		return "tax"
	case PriceComponentTypeInformational:
		return "informational"
	}
	return "<unknown>"
}

func (code PriceComponentType) Display() string {
	switch code {
	case PriceComponentTypeBase:
		return "base price"
	case PriceComponentTypeSurcharge:
		return "surcharge"
	case PriceComponentTypeDeduction:
		return "deduction"
	case PriceComponentTypeDiscount:
		return "discount"
	case PriceComponentTypeTax:
		return "tax"
	case PriceComponentTypeInformational:
		return "informational"
	}
	return "<unknown>"
}

func (code PriceComponentType) Definition() string {
	switch code {
	case PriceComponentTypeBase:
		return "the amount is the base price used for calculating the total price before applying surcharges, discount or taxes."
	case PriceComponentTypeSurcharge:
		return "the amount is a surcharge applied on the base price."
	case PriceComponentTypeDeduction:
		return "the amount is a deduction applied on the base price."
	case PriceComponentTypeDiscount:
		return "the amount is a discount applied on the base price."
	case PriceComponentTypeTax:
		return "the amount is the tax component of the total price."
	case PriceComponentTypeInformational:
		return "the amount is of informational character, it has not been applied in the calculation of the total price."
	}
	return "<unknown>"
}
