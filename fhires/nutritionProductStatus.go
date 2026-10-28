package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// Codes identifying the lifecycle stage of a product.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/nutritionproduct-status
type NutritionProductStatus int

const (
	NutritionProductStatusActive NutritionProductStatus = iota
	// The product is not expected or allowed to be used.
	NutritionProductStatusInactive
	// This electronic record should never have existed, though it is possible that real-world decisions were based on it.  (If real-world activity has occurred, the status should be "cancelled" rather than "entered-in-error".).
	NutritionProductStatusEnteredInError
)

func (code NutritionProductStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *NutritionProductStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "active":
		*code = NutritionProductStatusActive
	case "inactive":
		*code = NutritionProductStatusInactive
	case "entered-in-error":
		*code = NutritionProductStatusEnteredInError
	default:
		return fmt.Errorf("unknown NutritionProductStatus code `%s`", s)
	}
	return nil
}

func (code NutritionProductStatus) String() string {
	return code.Code()
}

func (code NutritionProductStatus) Code() string {
	switch code {
	case NutritionProductStatusActive:
		return "active"
	case NutritionProductStatusInactive:
		return "inactive"
	case NutritionProductStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}

func (code NutritionProductStatus) Display() string {
	switch code {
	case NutritionProductStatusActive:
		return "Active"
	case NutritionProductStatusInactive:
		return "Inactive"
	case NutritionProductStatusEnteredInError:
		return "Entered in Error"
	}
	return "<unknown>"
}

func (code NutritionProductStatus) Definition() string {
	switch code {
	case NutritionProductStatusActive:
		return "The product can be used."
	case NutritionProductStatusInactive:
		return "The product is not expected or allowed to be used."
	case NutritionProductStatusEnteredInError:
		return "This electronic record should never have existed, though it is possible that real-world decisions were based on it.  (If real-world activity has occurred, the status should be \"cancelled\" rather than \"entered-in-error\".)."
	}
	return "<unknown>"
}
