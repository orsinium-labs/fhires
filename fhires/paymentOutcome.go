package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// PaymentOutcome is documented here http://hl7.org/fhir/ValueSet/payment-outcome
type PaymentOutcome int

const (
	PaymentOutcomeQueued PaymentOutcome = iota
	PaymentOutcomeComplete
	PaymentOutcomeError
	PaymentOutcomePartial
)

func (code PaymentOutcome) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *PaymentOutcome) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "queued":
		*code = PaymentOutcomeQueued
	case "complete":
		*code = PaymentOutcomeComplete
	case "error":
		*code = PaymentOutcomeError
	case "partial":
		*code = PaymentOutcomePartial
	default:
		return fmt.Errorf("unknown PaymentOutcome code `%s`", s)
	}
	return nil
}

func (code PaymentOutcome) String() string {
	return code.Code()
}

func (code PaymentOutcome) Code() string {
	switch code {
	case PaymentOutcomeQueued:
		return "queued"
	case PaymentOutcomeComplete:
		return "complete"
	case PaymentOutcomeError:
		return "error"
	case PaymentOutcomePartial:
		return "partial"
	}
	return "<unknown>"
}

func (code PaymentOutcome) Display() string {
	switch code {
	case PaymentOutcomeQueued:
		return "Queued"
	case PaymentOutcomeComplete:
		return "Processing Complete"
	case PaymentOutcomeError:
		return "Error"
	case PaymentOutcomePartial:
		return "Partial Processing"
	}
	return "<unknown>"
}

func (code PaymentOutcome) Definition() string {
	switch code {
	case PaymentOutcomeQueued:
		return "The Claim/Pre-authorization/Pre-determination has been received but processing has not begun."
	case PaymentOutcomeComplete:
		return "The processing has completed without errors"
	case PaymentOutcomeError:
		return "One or more errors have been detected in the Claim"
	case PaymentOutcomePartial:
		return "No errors have been detected in the Claim and some of the adjudication has been performed."
	}
	return "<unknown>"
}
