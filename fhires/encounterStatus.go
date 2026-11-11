package fhires

// Current state of the encounter.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/encounter-status
type EncounterStatus string

const (
	// The Encounter has not yet started.
	EncounterStatusPlanned EncounterStatus = "planned"
	// The Encounter has begun and the patient is present / the practitioner and the patient are meeting.
	EncounterStatusInProgress EncounterStatus = "in-progress"
	// The Encounter has begun, but is currently on hold, e.g. because the patient is temporarily on leave.
	EncounterStatusOnHold EncounterStatus = "on-hold"
	// The Encounter has been clinically completed, the patient has been discharged from the facility or the visit has ended, and the patient may have departed (refer to subjectStatus). While the encounter is in this status, administrative activities are usually performed, collating all required documentation and charge information before being released for billing, at which point the status will move to completed.
	EncounterStatusDischarged EncounterStatus = "discharged"
	// The Encounter has ended.
	EncounterStatusCompleted EncounterStatus = "completed"
	// The Encounter has ended before it has begun.
	EncounterStatusCancelled EncounterStatus = "cancelled"
	// The Encounter has started, but was not able to be completed. Further action may need to be performed, such as rescheduling appointments related to this encounter.
	EncounterStatusDiscontinued EncounterStatus = "discontinued"
	// This instance should not have been part of this patient's medical record.
	EncounterStatusEnteredInError EncounterStatus = "entered-in-error"
	// The encounter status is unknown. Note that "unknown" is a value of last resort and every attempt should be made to provide a meaningful value other than "unknown".
	EncounterStatusUnknown EncounterStatus = "unknown"
)

func (code EncounterStatus) String() string {
	return code.Code()
}

func (code EncounterStatus) Code() string {
	return string(code)
}

func (code EncounterStatus) Display() string {
	switch code {
	case EncounterStatusPlanned:
		return "Planned"
	case EncounterStatusInProgress:
		return "In Progress"
	case EncounterStatusOnHold:
		return "On Hold"
	case EncounterStatusDischarged:
		return "Discharged"
	case EncounterStatusCompleted:
		return "Completed"
	case EncounterStatusCancelled:
		return "Cancelled"
	case EncounterStatusDiscontinued:
		return "Discontinued"
	case EncounterStatusEnteredInError:
		return "Entered in Error"
	case EncounterStatusUnknown:
		return "Unknown"
	}
	return "<unknown>"
}

func (code EncounterStatus) Definition() string {
	switch code {
	case EncounterStatusPlanned:
		return "The Encounter has not yet started."
	case EncounterStatusInProgress:
		return "The Encounter has begun and the patient is present / the practitioner and the patient are meeting."
	case EncounterStatusOnHold:
		return "The Encounter has begun, but is currently on hold, e.g. because the patient is temporarily on leave."
	case EncounterStatusDischarged:
		return "The Encounter has been clinically completed, the patient has been discharged from the facility or the visit has ended, and the patient may have departed (refer to subjectStatus). While the encounter is in this status, administrative activities are usually performed, collating all required documentation and charge information before being released for billing, at which point the status will move to completed."
	case EncounterStatusCompleted:
		return "The Encounter has ended."
	case EncounterStatusCancelled:
		return "The Encounter has ended before it has begun."
	case EncounterStatusDiscontinued:
		return "The Encounter has started, but was not able to be completed. Further action may need to be performed, such as rescheduling appointments related to this encounter."
	case EncounterStatusEnteredInError:
		return "This instance should not have been part of this patient's medical record."
	case EncounterStatusUnknown:
		return "The encounter status is unknown. Note that \"unknown\" is a value of last resort and every attempt should be made to provide a meaningful value other than \"unknown\"."
	}
	return "<unknown>"
}
