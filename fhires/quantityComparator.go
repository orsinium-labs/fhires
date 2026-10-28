package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// How the Quantity should be understood and represented.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/quantity-comparator
type QuantityComparator int

const (
	QuantityComparatorLessThan QuantityComparator = iota
	// The actual value is less than or equal to the given value.
	QuantityComparatorLessOrEquals
	// The actual value is greater than or equal to the given value.
	QuantityComparatorGreaterOrEquals
	// The actual value is greater than the given value.
	QuantityComparatorGreaterThan
	// The actual value is sufficient for the total quantity to equal the given value.
	QuantityComparatorAd
)

func (code QuantityComparator) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *QuantityComparator) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "<":
		*code = QuantityComparatorLessThan
	case "<=":
		*code = QuantityComparatorLessOrEquals
	case ">=":
		*code = QuantityComparatorGreaterOrEquals
	case ">":
		*code = QuantityComparatorGreaterThan
	case "ad":
		*code = QuantityComparatorAd
	default:
		return fmt.Errorf("unknown QuantityComparator code `%s`", s)
	}
	return nil
}

func (code QuantityComparator) String() string {
	return code.Code()
}

func (code QuantityComparator) Code() string {
	switch code {
	case QuantityComparatorLessThan:
		return "<"
	case QuantityComparatorLessOrEquals:
		return "<="
	case QuantityComparatorGreaterOrEquals:
		return ">="
	case QuantityComparatorGreaterThan:
		return ">"
	case QuantityComparatorAd:
		return "ad"
	}
	return "<unknown>"
}

func (code QuantityComparator) Display() string {
	switch code {
	case QuantityComparatorLessThan:
		return "Less than"
	case QuantityComparatorLessOrEquals:
		return "Less or Equal to"
	case QuantityComparatorGreaterOrEquals:
		return "Greater or Equal to"
	case QuantityComparatorGreaterThan:
		return "Greater than"
	case QuantityComparatorAd:
		return "Sufficient to achieve this total quantity"
	}
	return "<unknown>"
}

func (code QuantityComparator) Definition() string {
	switch code {
	case QuantityComparatorLessThan:
		return "The actual value is less than the given value."
	case QuantityComparatorLessOrEquals:
		return "The actual value is less than or equal to the given value."
	case QuantityComparatorGreaterOrEquals:
		return "The actual value is greater than or equal to the given value."
	case QuantityComparatorGreaterThan:
		return "The actual value is greater than the given value."
	case QuantityComparatorAd:
		return "The actual value is sufficient for the total quantity to equal the given value."
	}
	return "<unknown>"
}
