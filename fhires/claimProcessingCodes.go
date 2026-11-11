package fhires

// This value set includes Claim Processing Outcome codes.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/claim-outcome
type ClaimProcessingCodes string

const (
	// The Claim/Pre-authorization/Pre-determination has been received but processing has not begun.
	ClaimProcessingCodesQueued ClaimProcessingCodes = "queued"
	// The processing has completed without errors
	ClaimProcessingCodesComplete ClaimProcessingCodes = "complete"
	// One or more errors have been detected in the Claim
	ClaimProcessingCodesError ClaimProcessingCodes = "error"
	// No errors have been detected in the Claim and some of the adjudication has been performed.
	ClaimProcessingCodesPartial ClaimProcessingCodes = "partial"
)

func (code ClaimProcessingCodes) String() string {
	return code.Code()
}

func (code ClaimProcessingCodes) Code() string {
	return string(code)
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
