package fhires

// Indicates the state of the consent.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/consent-state-codes
type ConsentState string

const (
	// The consent is in development or awaiting use but is not yet intended to be acted upon.
	ConsentStateDraft ConsentState = "draft"
	// The consent is to be followed and enforced.
	ConsentStateActive ConsentState = "active"
	// The consent is terminated or replaced.
	ConsentStateInactive ConsentState = "inactive"
	// The consent development has been terminated prior to completion.
	ConsentStateNotDone ConsentState = "not-done"
	// The consent was created wrongly (e.g. wrong patient) and should be ignored.
	ConsentStateEnteredInError ConsentState = "entered-in-error"
	// The resource is in an indeterminate state.
	ConsentStateUnknown ConsentState = "unknown"
)

func (code ConsentState) String() string {
	return code.Code()
}

func (code ConsentState) Code() string {
	return string(code)
}

func (code ConsentState) Display() string {
	switch code {
	case ConsentStateDraft:
		return "Pending"
	case ConsentStateActive:
		return "Active"
	case ConsentStateInactive:
		return "Inactive"
	case ConsentStateNotDone:
		return "Abandoned"
	case ConsentStateEnteredInError:
		return "Entered in Error"
	case ConsentStateUnknown:
		return "Unknown"
	}
	return "<unknown>"
}

func (code ConsentState) Definition() string {
	switch code {
	case ConsentStateDraft:
		return "The consent is in development or awaiting use but is not yet intended to be acted upon."
	case ConsentStateActive:
		return "The consent is to be followed and enforced."
	case ConsentStateInactive:
		return "The consent is terminated or replaced."
	case ConsentStateNotDone:
		return "The consent development has been terminated prior to completion."
	case ConsentStateEnteredInError:
		return "The consent was created wrongly (e.g. wrong patient) and should be ignored."
	case ConsentStateUnknown:
		return "The resource is in an indeterminate state."
	}
	return "<unknown>"
}
