package fhires

// The type of manual completion to use for assertion.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/assert-manual-completion-codes
type AssertionManualCompletionType string

const (
	// Mark the currently waiting test failed and proceed with the next assert if the stopTestOnFail is false or the next test in the TestScript if the stopTestOnFail is true.
	AssertionManualCompletionTypeFail AssertionManualCompletionType = "fail"
	// Mark the currently waiting test passed (if the test is not failed already) and proceed with the next action in the TestScript.
	AssertionManualCompletionTypePass AssertionManualCompletionType = "pass"
	// Mark this assert as skipped and proceed with the next action in the TestScript.
	AssertionManualCompletionTypeSkip AssertionManualCompletionType = "skip"
	// Stop execution of this TestScript. The overall status of this TestScript is evaluated based on the status of the completed tests.
	AssertionManualCompletionTypeStop AssertionManualCompletionType = "stop"
)

func (code AssertionManualCompletionType) String() string {
	return code.Code()
}

func (code AssertionManualCompletionType) Code() string {
	return string(code)
}

func (code AssertionManualCompletionType) Display() string {
	switch code {
	case AssertionManualCompletionTypeFail:
		return "Fail"
	case AssertionManualCompletionTypePass:
		return "Pass"
	case AssertionManualCompletionTypeSkip:
		return "Skip"
	case AssertionManualCompletionTypeStop:
		return "Stop"
	}
	return "<unknown>"
}

func (code AssertionManualCompletionType) Definition() string {
	switch code {
	case AssertionManualCompletionTypeFail:
		return "Mark the currently waiting test failed and proceed with the next assert if the stopTestOnFail is false or the next test in the TestScript if the stopTestOnFail is true."
	case AssertionManualCompletionTypePass:
		return "Mark the currently waiting test passed (if the test is not failed already) and proceed with the next action in the TestScript."
	case AssertionManualCompletionTypeSkip:
		return "Mark this assert as skipped and proceed with the next action in the TestScript."
	case AssertionManualCompletionTypeStop:
		return "Stop execution of this TestScript. The overall status of this TestScript is evaluated based on the status of the completed tests."
	}
	return "<unknown>"
}
