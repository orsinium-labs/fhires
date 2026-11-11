package fhires

// The type of operator to use for assertion.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/assert-operator-codes
type AssertionOperatorType string

const (
	// Default value. Equals comparison.
	AssertionOperatorTypeEquals AssertionOperatorType = "equals"
	// Not equals comparison.
	AssertionOperatorTypeNotEquals AssertionOperatorType = "notEquals"
	// Compare value within a known set of values.
	AssertionOperatorTypeIn AssertionOperatorType = "in"
	// Compare value not within a known set of values.
	AssertionOperatorTypeNotIn AssertionOperatorType = "notIn"
	// Compare value to be greater than a known value.
	AssertionOperatorTypeGreaterThan AssertionOperatorType = "greaterThan"
	// Compare value to be less than a known value.
	AssertionOperatorTypeLessThan AssertionOperatorType = "lessThan"
	// Compare value is empty.
	AssertionOperatorTypeEmpty AssertionOperatorType = "empty"
	// Compare value is not empty.
	AssertionOperatorTypeNotEmpty AssertionOperatorType = "notEmpty"
	// Compare value string contains a known value.
	AssertionOperatorTypeContains AssertionOperatorType = "contains"
	// Compare value string does not contain a known value.
	AssertionOperatorTypeNotContains AssertionOperatorType = "notContains"
	// Evaluate the FHIRPath expression as a boolean condition.
	AssertionOperatorTypeEval AssertionOperatorType = "eval"
	// Manually evaluate the condition described by this assert. The test engine SHALL pause and provide an input mechanism to set the outcome of this assert to 'pass', 'fail', 'skip' or 'stop'.
	AssertionOperatorTypeManualEval AssertionOperatorType = "manualEval"
)

func (code AssertionOperatorType) String() string {
	return code.Code()
}

func (code AssertionOperatorType) Code() string {
	return string(code)
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
