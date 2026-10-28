package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// The type of manual completion to use for assertion.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/assert-manual-completion-codes
type AssertionManualCompletionType int

const (
	AssertionManualCompletionTypeFail AssertionManualCompletionType = iota
	// Mark the currently waiting test passed (if the test is not failed already) and proceed with the next action in the TestScript.
	AssertionManualCompletionTypePass
	// Mark this assert as skipped and proceed with the next action in the TestScript.
	AssertionManualCompletionTypeSkip
	// Stop execution of this TestScript. The overall status of this TestScript is evaluated based on the status of the completed tests.
	AssertionManualCompletionTypeStop
)

func (code AssertionManualCompletionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *AssertionManualCompletionType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "fail":
		*code = AssertionManualCompletionTypeFail
	case "pass":
		*code = AssertionManualCompletionTypePass
	case "skip":
		*code = AssertionManualCompletionTypeSkip
	case "stop":
		*code = AssertionManualCompletionTypeStop
	default:
		return fmt.Errorf("unknown AssertionManualCompletionType code `%s`", s)
	}
	return nil
}

func (code AssertionManualCompletionType) String() string {
	return code.Code()
}

func (code AssertionManualCompletionType) Code() string {
	switch code {
	case AssertionManualCompletionTypeFail:
		return "fail"
	case AssertionManualCompletionTypePass:
		return "pass"
	case AssertionManualCompletionTypeSkip:
		return "skip"
	case AssertionManualCompletionTypeStop:
		return "stop"
	}
	return "<unknown>"
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
