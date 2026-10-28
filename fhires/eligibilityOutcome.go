package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// The outcome of the processing.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/eligibility-outcome
type EligibilityOutcome int

const (
	EligibilityOutcomeQueued EligibilityOutcome = iota
	// The processing has completed without errors
	EligibilityOutcomeComplete
	// One or more errors have been detected in the Claim
	EligibilityOutcomeError
	// No errors have been detected in the Claim and some of the adjudication has been performed.
	EligibilityOutcomePartial
)

func (code EligibilityOutcome) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *EligibilityOutcome) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "queued":
		*code = EligibilityOutcomeQueued
	case "complete":
		*code = EligibilityOutcomeComplete
	case "error":
		*code = EligibilityOutcomeError
	case "partial":
		*code = EligibilityOutcomePartial
	default:
		return fmt.Errorf("unknown EligibilityOutcome code `%s`", s)
	}
	return nil
}

func (code EligibilityOutcome) String() string {
	return code.Code()
}

func (code EligibilityOutcome) Code() string {
	switch code {
	case EligibilityOutcomeQueued:
		return "queued"
	case EligibilityOutcomeComplete:
		return "complete"
	case EligibilityOutcomeError:
		return "error"
	case EligibilityOutcomePartial:
		return "partial"
	}
	return "<unknown>"
}

func (code EligibilityOutcome) Display() string {
	switch code {
	case EligibilityOutcomeQueued:
		return "Queued"
	case EligibilityOutcomeComplete:
		return "Processing Complete"
	case EligibilityOutcomeError:
		return "Error"
	case EligibilityOutcomePartial:
		return "Partial Processing"
	}
	return "<unknown>"
}

func (code EligibilityOutcome) Definition() string {
	switch code {
	case EligibilityOutcomeQueued:
		return "The Claim/Pre-authorization/Pre-determination has been received but processing has not begun."
	case EligibilityOutcomeComplete:
		return "The processing has completed without errors"
	case EligibilityOutcomeError:
		return "One or more errors have been detected in the Claim"
	case EligibilityOutcomePartial:
		return "No errors have been detected in the Claim and some of the adjudication has been performed."
	}
	return "<unknown>"
}
