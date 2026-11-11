package fhires

// The kind of operation to perform as a part of a property based filter.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/filter-operator
type FilterOperator string

const (
	// The specified property of the code equals the provided value.
	FilterOperatorEquals FilterOperator = "="
	// Includes all concept ids that have a transitive is-a relationship with the concept Id provided as the value, including the provided concept itself (include descendant codes and self).
	FilterOperatorIsA FilterOperator = "is-a"
	// Includes all concept ids that have a transitive is-a relationship with the concept Id provided as the value, excluding the provided concept itself (i.e. include descendant codes only).
	FilterOperatorDescendentOf FilterOperator = "descendent-of"
	// The specified property of the code does not have an is-a relationship with the provided value.
	FilterOperatorIsNotA FilterOperator = "is-not-a"
	// The specified property of the code  matches the regex specified in the provided value.
	FilterOperatorRegex FilterOperator = "regex"
	// The specified property of the code is in the set of codes or concepts specified in the provided value (comma-separated list).
	FilterOperatorIn FilterOperator = "in"
	// The specified property of the code is not in the set of codes or concepts specified in the provided value (comma-separated list).
	FilterOperatorNotIn FilterOperator = "not-in"
	// Includes all concept ids that have a transitive is-a relationship from the concept Id provided as the value, including the provided concept itself (i.e. include ancestor codes and self).
	FilterOperatorGeneralizes FilterOperator = "generalizes"
	// Only concepts with a direct hierarchical relationship to the index code and no other concepts. This does not include the index code in the output.
	FilterOperatorChildOf FilterOperator = "child-of"
	// Includes concept ids that have a transitive is-a relationship with the concept Id provided as the value, but which do not have any concept ids with transitive is-a relationships with themselves.
	FilterOperatorDescendentLeaf FilterOperator = "descendent-leaf"
	// The specified property of the code has at least one value (if the specified value is true; if the specified value is false, then matches when the specified property of the code has no values).
	FilterOperatorExists FilterOperator = "exists"
)

func (code FilterOperator) String() string {
	return code.Code()
}

func (code FilterOperator) Code() string {
	return string(code)
}

func (code FilterOperator) Display() string {
	switch code {
	case FilterOperatorEquals:
		return "Equals"
	case FilterOperatorIsA:
		return "Is A (by subsumption)"
	case FilterOperatorDescendentOf:
		return "Descendent Of (by subsumption)"
	case FilterOperatorIsNotA:
		return "Not (Is A) (by subsumption)"
	case FilterOperatorRegex:
		return "Regular Expression"
	case FilterOperatorIn:
		return "In Set"
	case FilterOperatorNotIn:
		return "Not in Set"
	case FilterOperatorGeneralizes:
		return "Generalizes (by Subsumption)"
	case FilterOperatorChildOf:
		return "Child Of"
	case FilterOperatorDescendentLeaf:
		return "Descendent Leaf"
	case FilterOperatorExists:
		return "Exists"
	}
	return "<unknown>"
}

func (code FilterOperator) Definition() string {
	switch code {
	case FilterOperatorEquals:
		return "The specified property of the code equals the provided value."
	case FilterOperatorIsA:
		return "Includes all concept ids that have a transitive is-a relationship with the concept Id provided as the value, including the provided concept itself (include descendant codes and self)."
	case FilterOperatorDescendentOf:
		return "Includes all concept ids that have a transitive is-a relationship with the concept Id provided as the value, excluding the provided concept itself (i.e. include descendant codes only)."
	case FilterOperatorIsNotA:
		return "The specified property of the code does not have an is-a relationship with the provided value."
	case FilterOperatorRegex:
		return "The specified property of the code  matches the regex specified in the provided value."
	case FilterOperatorIn:
		return "The specified property of the code is in the set of codes or concepts specified in the provided value (comma-separated list)."
	case FilterOperatorNotIn:
		return "The specified property of the code is not in the set of codes or concepts specified in the provided value (comma-separated list)."
	case FilterOperatorGeneralizes:
		return "Includes all concept ids that have a transitive is-a relationship from the concept Id provided as the value, including the provided concept itself (i.e. include ancestor codes and self)."
	case FilterOperatorChildOf:
		return "Only concepts with a direct hierarchical relationship to the index code and no other concepts. This does not include the index code in the output."
	case FilterOperatorDescendentLeaf:
		return "Includes concept ids that have a transitive is-a relationship with the concept Id provided as the value, but which do not have any concept ids with transitive is-a relationships with themselves."
	case FilterOperatorExists:
		return "The specified property of the code has at least one value (if the specified value is true; if the specified value is false, then matches when the specified property of the code has no values)."
	}
	return "<unknown>"
}
