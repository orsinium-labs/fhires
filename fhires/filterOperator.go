package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// The kind of operation to perform as a part of a property based filter.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/filter-operator
type FilterOperator int

const (
	FilterOperatorEquals FilterOperator = iota
	// Includes all concept ids that have a transitive is-a relationship with the concept Id provided as the value, including the provided concept itself (include descendant codes and self).
	FilterOperatorIsA
	// Includes all concept ids that have a transitive is-a relationship with the concept Id provided as the value, excluding the provided concept itself (i.e. include descendant codes only).
	FilterOperatorDescendentOf
	// The specified property of the code does not have an is-a relationship with the provided value.
	FilterOperatorIsNotA
	// The specified property of the code  matches the regex specified in the provided value.
	FilterOperatorRegex
	// The specified property of the code is in the set of codes or concepts specified in the provided value (comma-separated list).
	FilterOperatorIn
	// The specified property of the code is not in the set of codes or concepts specified in the provided value (comma-separated list).
	FilterOperatorNotIn
	// Includes all concept ids that have a transitive is-a relationship from the concept Id provided as the value, including the provided concept itself (i.e. include ancestor codes and self).
	FilterOperatorGeneralizes
	// Only concepts with a direct hierarchical relationship to the index code and no other concepts. This does not include the index code in the output.
	FilterOperatorChildOf
	// Includes concept ids that have a transitive is-a relationship with the concept Id provided as the value, but which do not have any concept ids with transitive is-a relationships with themselves.
	FilterOperatorDescendentLeaf
	// The specified property of the code has at least one value (if the specified value is true; if the specified value is false, then matches when the specified property of the code has no values).
	FilterOperatorExists
)

func (code FilterOperator) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *FilterOperator) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "=":
		*code = FilterOperatorEquals
	case "is-a":
		*code = FilterOperatorIsA
	case "descendent-of":
		*code = FilterOperatorDescendentOf
	case "is-not-a":
		*code = FilterOperatorIsNotA
	case "regex":
		*code = FilterOperatorRegex
	case "in":
		*code = FilterOperatorIn
	case "not-in":
		*code = FilterOperatorNotIn
	case "generalizes":
		*code = FilterOperatorGeneralizes
	case "child-of":
		*code = FilterOperatorChildOf
	case "descendent-leaf":
		*code = FilterOperatorDescendentLeaf
	case "exists":
		*code = FilterOperatorExists
	default:
		return fmt.Errorf("unknown FilterOperator code `%s`", s)
	}
	return nil
}

func (code FilterOperator) String() string {
	return code.Code()
}

func (code FilterOperator) Code() string {
	switch code {
	case FilterOperatorEquals:
		return "="
	case FilterOperatorIsA:
		return "is-a"
	case FilterOperatorDescendentOf:
		return "descendent-of"
	case FilterOperatorIsNotA:
		return "is-not-a"
	case FilterOperatorRegex:
		return "regex"
	case FilterOperatorIn:
		return "in"
	case FilterOperatorNotIn:
		return "not-in"
	case FilterOperatorGeneralizes:
		return "generalizes"
	case FilterOperatorChildOf:
		return "child-of"
	case FilterOperatorDescendentLeaf:
		return "descendent-leaf"
	case FilterOperatorExists:
		return "exists"
	}
	return "<unknown>"
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
