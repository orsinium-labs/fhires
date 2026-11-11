package fhires

// The outcome of the processing.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/payment-outcome
type PaymentOutcome string

const (
	// The Claim/Pre-authorization/Pre-determination has been received but processing has not begun.
	PaymentOutcomeQueued PaymentOutcome = "queued"
	// The processing has completed without errors
	PaymentOutcomeComplete PaymentOutcome = "complete"
	// One or more errors have been detected in the Claim
	PaymentOutcomeError PaymentOutcome = "error"
	// No errors have been detected in the Claim and some of the adjudication has been performed.
	PaymentOutcomePartial PaymentOutcome = "partial"
)

func (code PaymentOutcome) String() string {
	return code.Code()
}

func (code PaymentOutcome) Code() string {
	return string(code)
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
