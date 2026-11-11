package fhires

// The type of comparator operator to use
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/value-filter-comparator
type ValueFilterComparator string

const (
	// the value for the parameter in the resource is equal to the provided value.
	ValueFilterComparatorEq ValueFilterComparator = "eq"
	// the value for the parameter in the resource is not equal to the provided value.
	ValueFilterComparatorNe ValueFilterComparator = "ne"
	// the value for the parameter in the resource is greater than the provided value.
	ValueFilterComparatorGt ValueFilterComparator = "gt"
	// the value for the parameter in the resource is less than the provided value.
	ValueFilterComparatorLt ValueFilterComparator = "lt"
	// the value for the parameter in the resource is greater or equal to the provided value.
	ValueFilterComparatorGe ValueFilterComparator = "ge"
	// the value for the parameter in the resource is less or equal to the provided value.
	ValueFilterComparatorLe ValueFilterComparator = "le"
	// the value for the parameter in the resource starts after the provided value.
	ValueFilterComparatorSa ValueFilterComparator = "sa"
	// the value for the parameter in the resource ends before the provided value.
	ValueFilterComparatorEb ValueFilterComparator = "eb"
	// the value for the parameter in the resource is approximately the same to the provided value.
	ValueFilterComparatorAp ValueFilterComparator = "ap"
)

func (code ValueFilterComparator) String() string {
	return code.Code()
}

func (code ValueFilterComparator) Code() string {
	return string(code)
}

func (code ValueFilterComparator) Display() string {
	switch code {
	case ValueFilterComparatorEq:
		return "Equals"
	case ValueFilterComparatorNe:
		return "Not Equals"
	case ValueFilterComparatorGt:
		return "Greater Than"
	case ValueFilterComparatorLt:
		return "Less Than"
	case ValueFilterComparatorGe:
		return "Greater or Equals"
	case ValueFilterComparatorLe:
		return "Less of Equal"
	case ValueFilterComparatorSa:
		return "Starts After"
	case ValueFilterComparatorEb:
		return "Ends Before"
	case ValueFilterComparatorAp:
		return "Approximately"
	}
	return "<unknown>"
}

func (code ValueFilterComparator) Definition() string {
	switch code {
	case ValueFilterComparatorEq:
		return "the value for the parameter in the resource is equal to the provided value."
	case ValueFilterComparatorNe:
		return "the value for the parameter in the resource is not equal to the provided value."
	case ValueFilterComparatorGt:
		return "the value for the parameter in the resource is greater than the provided value."
	case ValueFilterComparatorLt:
		return "the value for the parameter in the resource is less than the provided value."
	case ValueFilterComparatorGe:
		return "the value for the parameter in the resource is greater or equal to the provided value."
	case ValueFilterComparatorLe:
		return "the value for the parameter in the resource is less or equal to the provided value."
	case ValueFilterComparatorSa:
		return "the value for the parameter in the resource starts after the provided value."
	case ValueFilterComparatorEb:
		return "the value for the parameter in the resource ends before the provided value."
	case ValueFilterComparatorAp:
		return "the value for the parameter in the resource is approximately the same to the provided value."
	}
	return "<unknown>"
}
