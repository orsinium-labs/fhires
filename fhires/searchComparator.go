package fhires

// What Search Comparator Codes are supported in search.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/search-comparator
type SearchComparator string

const (
	// the value for the parameter in the resource is equal to the provided value.
	SearchComparatorEq SearchComparator = "eq"
	// the value for the parameter in the resource is not equal to the provided value.
	SearchComparatorNe SearchComparator = "ne"
	// the value for the parameter in the resource is greater than the provided value.
	SearchComparatorGt SearchComparator = "gt"
	// the value for the parameter in the resource is less than the provided value.
	SearchComparatorLt SearchComparator = "lt"
	// the value for the parameter in the resource is greater or equal to the provided value.
	SearchComparatorGe SearchComparator = "ge"
	// the value for the parameter in the resource is less or equal to the provided value.
	SearchComparatorLe SearchComparator = "le"
	// the value for the parameter in the resource starts after the provided value.
	SearchComparatorSa SearchComparator = "sa"
	// the value for the parameter in the resource ends before the provided value.
	SearchComparatorEb SearchComparator = "eb"
	// the value for the parameter in the resource is approximately the same to the provided value.
	SearchComparatorAp SearchComparator = "ap"
)

func (code SearchComparator) String() string {
	return code.Code()
}

func (code SearchComparator) Code() string {
	return string(code)
}

func (code SearchComparator) Display() string {
	switch code {
	case SearchComparatorEq:
		return "Equals"
	case SearchComparatorNe:
		return "Not Equals"
	case SearchComparatorGt:
		return "Greater Than"
	case SearchComparatorLt:
		return "Less Than"
	case SearchComparatorGe:
		return "Greater or Equals"
	case SearchComparatorLe:
		return "Less of Equal"
	case SearchComparatorSa:
		return "Starts After"
	case SearchComparatorEb:
		return "Ends Before"
	case SearchComparatorAp:
		return "Approximately"
	}
	return "<unknown>"
}

func (code SearchComparator) Definition() string {
	switch code {
	case SearchComparatorEq:
		return "the value for the parameter in the resource is equal to the provided value."
	case SearchComparatorNe:
		return "the value for the parameter in the resource is not equal to the provided value."
	case SearchComparatorGt:
		return "the value for the parameter in the resource is greater than the provided value."
	case SearchComparatorLt:
		return "the value for the parameter in the resource is less than the provided value."
	case SearchComparatorGe:
		return "the value for the parameter in the resource is greater or equal to the provided value."
	case SearchComparatorLe:
		return "the value for the parameter in the resource is less or equal to the provided value."
	case SearchComparatorSa:
		return "the value for the parameter in the resource starts after the provided value."
	case SearchComparatorEb:
		return "the value for the parameter in the resource ends before the provided value."
	case SearchComparatorAp:
		return "the value for the parameter in the resource is approximately the same to the provided value."
	}
	return "<unknown>"
}
