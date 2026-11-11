package fhires

// Indicates the status of the care team.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/care-team-status
type CareTeamStatus string

const (
	// The care team has been drafted and proposed, but not yet participating in the coordination and delivery of patient care.
	CareTeamStatusProposed CareTeamStatus = "proposed"
	// The care team is currently participating in the coordination and delivery of care.
	CareTeamStatusActive CareTeamStatus = "active"
	// The care team is temporarily on hold or suspended and not participating in the coordination and delivery of care.
	CareTeamStatusSuspended CareTeamStatus = "suspended"
	// The care team was, but is no longer, participating in the coordination and delivery of care.
	CareTeamStatusInactive CareTeamStatus = "inactive"
	// The care team should have never existed.
	CareTeamStatusEnteredInError CareTeamStatus = "entered-in-error"
)

func (code CareTeamStatus) String() string {
	return code.Code()
}

func (code CareTeamStatus) Code() string {
	return string(code)
}

func (code CareTeamStatus) Display() string {
	switch code {
	case CareTeamStatusProposed:
		return "Proposed"
	case CareTeamStatusActive:
		return "Active"
	case CareTeamStatusSuspended:
		return "Suspended"
	case CareTeamStatusInactive:
		return "Inactive"
	case CareTeamStatusEnteredInError:
		return "Entered in Error"
	}
	return "<unknown>"
}

func (code CareTeamStatus) Definition() string {
	switch code {
	case CareTeamStatusProposed:
		return "The care team has been drafted and proposed, but not yet participating in the coordination and delivery of patient care."
	case CareTeamStatusActive:
		return "The care team is currently participating in the coordination and delivery of care."
	case CareTeamStatusSuspended:
		return "The care team is temporarily on hold or suspended and not participating in the coordination and delivery of care."
	case CareTeamStatusInactive:
		return "The care team was, but is no longer, participating in the coordination and delivery of care."
	case CareTeamStatusEnteredInError:
		return "The care team should have never existed."
	}
	return "<unknown>"
}
