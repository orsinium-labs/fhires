package fhires

// The current state of the list.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/list-status
type ListStatus string

const (
	// The list is considered to be an active part of the patient's record.
	ListStatusCurrent ListStatus = "current"
	// The list is "old" and should no longer be considered accurate or relevant.
	ListStatusRetired ListStatus = "retired"
	// The list was never accurate.  It is retained for medico-legal purposes only.
	ListStatusEnteredInError ListStatus = "entered-in-error"
)

func (code ListStatus) String() string {
	return code.Code()
}

func (code ListStatus) Code() string {
	return string(code)
}

func (code ListStatus) Display() string {
	switch code {
	case ListStatusCurrent:
		return "Current"
	case ListStatusRetired:
		return "Retired"
	case ListStatusEnteredInError:
		return "Entered In Error"
	}
	return "<unknown>"
}

func (code ListStatus) Definition() string {
	switch code {
	case ListStatusCurrent:
		return "The list is considered to be an active part of the patient's record."
	case ListStatusRetired:
		return "The list is \"old\" and should no longer be considered accurate or relevant."
	case ListStatusEnteredInError:
		return "The list was never accurate.  It is retained for medico-legal purposes only."
	}
	return "<unknown>"
}
