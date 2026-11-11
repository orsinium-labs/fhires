package fhires

// The outcome of the processing.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/enrollment-outcome
type EnrollmentOutcome string

const (
	// The Claim/Pre-authorization/Pre-determination has been received but processing has not begun.
	EnrollmentOutcomeQueued EnrollmentOutcome = "queued"
	// The processing has completed without errors
	EnrollmentOutcomeComplete EnrollmentOutcome = "complete"
	// One or more errors have been detected in the Claim
	EnrollmentOutcomeError EnrollmentOutcome = "error"
	// No errors have been detected in the Claim and some of the adjudication has been performed.
	EnrollmentOutcomePartial EnrollmentOutcome = "partial"
)

func (code EnrollmentOutcome) String() string {
	return code.Code()
}

func (code EnrollmentOutcome) Code() string {
	return string(code)
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
