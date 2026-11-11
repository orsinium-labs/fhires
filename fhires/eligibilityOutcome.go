package fhires

// The outcome of the processing.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/eligibility-outcome
type EligibilityOutcome string

const (
	// The Claim/Pre-authorization/Pre-determination has been received but processing has not begun.
	EligibilityOutcomeQueued EligibilityOutcome = "queued"
	// The processing has completed without errors
	EligibilityOutcomeComplete EligibilityOutcome = "complete"
	// One or more errors have been detected in the Claim
	EligibilityOutcomeError EligibilityOutcome = "error"
	// No errors have been detected in the Claim and some of the adjudication has been performed.
	EligibilityOutcomePartial EligibilityOutcome = "partial"
)

func (code EligibilityOutcome) String() string {
	return code.Code()
}

func (code EligibilityOutcome) Code() string {
	return string(code)
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
