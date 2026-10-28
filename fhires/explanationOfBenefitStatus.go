package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// A code specifying the state of the resource instance.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/explanationofbenefit-status
type ExplanationOfBenefitStatus int

const (
	ExplanationOfBenefitStatusActive ExplanationOfBenefitStatus = iota
	// The resource instance is withdrawn, rescinded or reversed.
	ExplanationOfBenefitStatusCancelled
	// A new resource instance the contents of which is not complete.
	ExplanationOfBenefitStatusDraft
	// The resource instance was entered in error.
	ExplanationOfBenefitStatusEnteredInError
)

func (code ExplanationOfBenefitStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *ExplanationOfBenefitStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "active":
		*code = ExplanationOfBenefitStatusActive
	case "cancelled":
		*code = ExplanationOfBenefitStatusCancelled
	case "draft":
		*code = ExplanationOfBenefitStatusDraft
	case "entered-in-error":
		*code = ExplanationOfBenefitStatusEnteredInError
	default:
		return fmt.Errorf("unknown ExplanationOfBenefitStatus code `%s`", s)
	}
	return nil
}

func (code ExplanationOfBenefitStatus) String() string {
	return code.Code()
}

func (code ExplanationOfBenefitStatus) Code() string {
	switch code {
	case ExplanationOfBenefitStatusActive:
		return "active"
	case ExplanationOfBenefitStatusCancelled:
		return "cancelled"
	case ExplanationOfBenefitStatusDraft:
		return "draft"
	case ExplanationOfBenefitStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}

func (code ExplanationOfBenefitStatus) Display() string {
	switch code {
	case ExplanationOfBenefitStatusActive:
		return "Active"
	case ExplanationOfBenefitStatusCancelled:
		return "Cancelled"
	case ExplanationOfBenefitStatusDraft:
		return "Draft"
	case ExplanationOfBenefitStatusEnteredInError:
		return "Entered In Error"
	}
	return "<unknown>"
}

func (code ExplanationOfBenefitStatus) Definition() string {
	switch code {
	case ExplanationOfBenefitStatusActive:
		return "The resource instance is currently in-force."
	case ExplanationOfBenefitStatusCancelled:
		return "The resource instance is withdrawn, rescinded or reversed."
	case ExplanationOfBenefitStatusDraft:
		return "A new resource instance the contents of which is not complete."
	case ExplanationOfBenefitStatusEnteredInError:
		return "The resource instance was entered in error."
	}
	return "<unknown>"
}
