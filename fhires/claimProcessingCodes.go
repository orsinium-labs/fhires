package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// This value set includes Claim Processing Outcome codes.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/claim-outcome
type ClaimProcessingCodes int

const (
	ClaimProcessingCodesQueued ClaimProcessingCodes = iota
	// The processing has completed without errors
	ClaimProcessingCodesComplete
	// One or more errors have been detected in the Claim
	ClaimProcessingCodesError
	// No errors have been detected in the Claim and some of the adjudication has been performed.
	ClaimProcessingCodesPartial
)

func (code ClaimProcessingCodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *ClaimProcessingCodes) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "queued":
		*code = ClaimProcessingCodesQueued
	case "complete":
		*code = ClaimProcessingCodesComplete
	case "error":
		*code = ClaimProcessingCodesError
	case "partial":
		*code = ClaimProcessingCodesPartial
	default:
		return fmt.Errorf("unknown ClaimProcessingCodes code `%s`", s)
	}
	return nil
}

func (code ClaimProcessingCodes) String() string {
	return code.Code()
}

func (code ClaimProcessingCodes) Code() string {
	switch code {
	case ClaimProcessingCodesQueued:
		return "queued"
	case ClaimProcessingCodesComplete:
		return "complete"
	case ClaimProcessingCodesError:
		return "error"
	case ClaimProcessingCodesPartial:
		return "partial"
	}
	return "<unknown>"
}

func (code ClaimProcessingCodes) Display() string {
	switch code {
	case ClaimProcessingCodesQueued:
		return "Queued"
	case ClaimProcessingCodesComplete:
		return "Processing Complete"
	case ClaimProcessingCodesError:
		return "Error"
	case ClaimProcessingCodesPartial:
		return "Partial Processing"
	}
	return "<unknown>"
}

func (code ClaimProcessingCodes) Definition() string {
	switch code {
	case ClaimProcessingCodesQueued:
		return "The Claim/Pre-authorization/Pre-determination has been received but processing has not begun."
	case ClaimProcessingCodesComplete:
		return "The processing has completed without errors"
	case ClaimProcessingCodesError:
		return "One or more errors have been detected in the Claim"
	case ClaimProcessingCodesPartial:
		return "No errors have been detected in the Claim and some of the adjudication has been performed."
	}
	return "<unknown>"
}
