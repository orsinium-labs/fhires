package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// The type of operator to use for assertion.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/assert-operator-codes
type AssertionOperatorType int

const (
	AssertionOperatorTypeEquals AssertionOperatorType = iota
	// Not equals comparison.
	AssertionOperatorTypeNotEquals
	// Compare value within a known set of values.
	AssertionOperatorTypeIn
	// Compare value not within a known set of values.
	AssertionOperatorTypeNotIn
	// Compare value to be greater than a known value.
	AssertionOperatorTypeGreaterThan
	// Compare value to be less than a known value.
	AssertionOperatorTypeLessThan
	// Compare value is empty.
	AssertionOperatorTypeEmpty
	// Compare value is not empty.
	AssertionOperatorTypeNotEmpty
	// Compare value string contains a known value.
	AssertionOperatorTypeContains
	// Compare value string does not contain a known value.
	AssertionOperatorTypeNotContains
	// Evaluate the FHIRPath expression as a boolean condition.
	AssertionOperatorTypeEval
	// Manually evaluate the condition described by this assert. The test engine SHALL pause and provide an input mechanism to set the outcome of this assert to 'pass', 'fail', 'skip' or 'stop'.
	AssertionOperatorTypeManualEval
)

func (code AssertionOperatorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *AssertionOperatorType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "equals":
		*code = AssertionOperatorTypeEquals
	case "notEquals":
		*code = AssertionOperatorTypeNotEquals
	case "in":
		*code = AssertionOperatorTypeIn
	case "notIn":
		*code = AssertionOperatorTypeNotIn
	case "greaterThan":
		*code = AssertionOperatorTypeGreaterThan
	case "lessThan":
		*code = AssertionOperatorTypeLessThan
	case "empty":
		*code = AssertionOperatorTypeEmpty
	case "notEmpty":
		*code = AssertionOperatorTypeNotEmpty
	case "contains":
		*code = AssertionOperatorTypeContains
	case "notContains":
		*code = AssertionOperatorTypeNotContains
	case "eval":
		*code = AssertionOperatorTypeEval
	case "manualEval":
		*code = AssertionOperatorTypeManualEval
	default:
		return fmt.Errorf("unknown AssertionOperatorType code `%s`", s)
	}
	return nil
}

func (code AssertionOperatorType) String() string {
	return code.Code()
}

func (code AssertionOperatorType) Code() string {
	switch code {
	case AssertionOperatorTypeEquals:
		return "equals"
	case AssertionOperatorTypeNotEquals:
		return "notEquals"
	case AssertionOperatorTypeIn:
		return "in"
	case AssertionOperatorTypeNotIn:
		return "notIn"
	case AssertionOperatorTypeGreaterThan:
		return "greaterThan"
	case AssertionOperatorTypeLessThan:
		return "lessThan"
	case AssertionOperatorTypeEmpty:
		return "empty"
	case AssertionOperatorTypeNotEmpty:
		return "notEmpty"
	case AssertionOperatorTypeContains:
		return "contains"
	case AssertionOperatorTypeNotContains:
		return "notContains"
	case AssertionOperatorTypeEval:
		return "eval"
	case AssertionOperatorTypeManualEval:
		return "manualEval"
	}
	return "<unknown>"
}

func (code AssertionOperatorType) Display() string {
	switch code {
	case AssertionOperatorTypeEquals:
		return "equals"
	case AssertionOperatorTypeNotEquals:
		return "notEquals"
	case AssertionOperatorTypeIn:
		return "in"
	case AssertionOperatorTypeNotIn:
		return "notIn"
	case AssertionOperatorTypeGreaterThan:
		return "greaterThan"
	case AssertionOperatorTypeLessThan:
		return "lessThan"
	case AssertionOperatorTypeEmpty:
		return "empty"
	case AssertionOperatorTypeNotEmpty:
		return "notEmpty"
	case AssertionOperatorTypeContains:
		return "contains"
	case AssertionOperatorTypeNotContains:
		return "notContains"
	case AssertionOperatorTypeEval:
		return "evaluate"
	case AssertionOperatorTypeManualEval:
		return "manualEvaluate"
	}
	return "<unknown>"
}

func (code AssertionOperatorType) Definition() string {
	switch code {
	case AssertionOperatorTypeEquals:
		return "Default value. Equals comparison."
	case AssertionOperatorTypeNotEquals:
		return "Not equals comparison."
	case AssertionOperatorTypeIn:
		return "Compare value within a known set of values."
	case AssertionOperatorTypeNotIn:
		return "Compare value not within a known set of values."
	case AssertionOperatorTypeGreaterThan:
		return "Compare value to be greater than a known value."
	case AssertionOperatorTypeLessThan:
		return "Compare value to be less than a known value."
	case AssertionOperatorTypeEmpty:
		return "Compare value is empty."
	case AssertionOperatorTypeNotEmpty:
		return "Compare value is not empty."
	case AssertionOperatorTypeContains:
		return "Compare value string contains a known value."
	case AssertionOperatorTypeNotContains:
		return "Compare value string does not contain a known value."
	case AssertionOperatorTypeEval:
		return "Evaluate the FHIRPath expression as a boolean condition."
	case AssertionOperatorTypeManualEval:
		return "Manually evaluate the condition described by this assert. The test engine SHALL pause and provide an input mechanism to set the outcome of this assert to 'pass', 'fail', 'skip' or 'stop'."
	}
	return "<unknown>"
}
