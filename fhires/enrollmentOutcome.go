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
// http://hl7.org/fhir/ValueSet/enrollment-outcome
type EnrollmentOutcome int

const (
	EnrollmentOutcomeQueued EnrollmentOutcome = iota
	// The processing has completed without errors
	EnrollmentOutcomeComplete
	// One or more errors have been detected in the Claim
	EnrollmentOutcomeError
	// No errors have been detected in the Claim and some of the adjudication has been performed.
	EnrollmentOutcomePartial
)

func (code EnrollmentOutcome) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *EnrollmentOutcome) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "queued":
		*code = EnrollmentOutcomeQueued
	case "complete":
		*code = EnrollmentOutcomeComplete
	case "error":
		*code = EnrollmentOutcomeError
	case "partial":
		*code = EnrollmentOutcomePartial
	default:
		return fmt.Errorf("unknown EnrollmentOutcome code `%s`", s)
	}
	return nil
}

func (code EnrollmentOutcome) String() string {
	return code.Code()
}

func (code EnrollmentOutcome) Code() string {
	switch code {
	case EnrollmentOutcomeQueued:
		return "queued"
	case EnrollmentOutcomeComplete:
		return "complete"
	case EnrollmentOutcomeError:
		return "error"
	case EnrollmentOutcomePartial:
		return "partial"
	}
	return "<unknown>"
}

func (code EnrollmentOutcome) Display() string {
	switch code {
	case EnrollmentOutcomeQueued:
		return "Queued"
	case EnrollmentOutcomeComplete:
		return "Processing Complete"
	case EnrollmentOutcomeError:
		return "Error"
	case EnrollmentOutcomePartial:
		return "Partial Processing"
	}
	return "<unknown>"
}

func (code EnrollmentOutcome) Definition() string {
	switch code {
	case EnrollmentOutcomeQueued:
		return "The Claim/Pre-authorization/Pre-determination has been received but processing has not begun."
	case EnrollmentOutcomeComplete:
		return "The processing has completed without errors"
	case EnrollmentOutcomeError:
		return "One or more errors have been detected in the Claim"
	case EnrollmentOutcomePartial:
		return "No errors have been detected in the Claim and some of the adjudication has been performed."
	}
	return "<unknown>"
}
