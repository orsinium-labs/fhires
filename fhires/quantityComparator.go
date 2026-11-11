package fhires

// How the Quantity should be understood and represented.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/quantity-comparator
type QuantityComparator string

const (
	// The actual value is less than the given value.
	QuantityComparatorLessThan QuantityComparator = "<"
	// The actual value is less than or equal to the given value.
	QuantityComparatorLessOrEquals QuantityComparator = "<="
	// The actual value is greater than or equal to the given value.
	QuantityComparatorGreaterOrEquals QuantityComparator = ">="
	// The actual value is greater than the given value.
	QuantityComparatorGreaterThan QuantityComparator = ">"
	// The actual value is sufficient for the total quantity to equal the given value.
	QuantityComparatorAd QuantityComparator = "ad"
)

func (code QuantityComparator) String() string {
	return code.Code()
}

func (code QuantityComparator) Code() string {
	return string(code)
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
