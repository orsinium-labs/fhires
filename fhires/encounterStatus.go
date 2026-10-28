package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// EncounterStatus is documented here http://hl7.org/fhir/ValueSet/encounter-status
type EncounterStatus int

const (
	EncounterStatusPlanned EncounterStatus = iota
	EncounterStatusInProgress
	EncounterStatusOnHold
	EncounterStatusDischarged
	EncounterStatusCompleted
	EncounterStatusCancelled
	EncounterStatusDiscontinued
	EncounterStatusEnteredInError
	EncounterStatusUnknown
)

func (code EncounterStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *EncounterStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "planned":
		*code = EncounterStatusPlanned
	case "in-progress":
		*code = EncounterStatusInProgress
	case "on-hold":
		*code = EncounterStatusOnHold
	case "discharged":
		*code = EncounterStatusDischarged
	case "completed":
		*code = EncounterStatusCompleted
	case "cancelled":
		*code = EncounterStatusCancelled
	case "discontinued":
		*code = EncounterStatusDiscontinued
	case "entered-in-error":
		*code = EncounterStatusEnteredInError
	case "unknown":
		*code = EncounterStatusUnknown
	default:
		return fmt.Errorf("unknown EncounterStatus code `%s`", s)
	}
	return nil
}

func (code EncounterStatus) String() string {
	return code.Code()
}

func (code EncounterStatus) Code() string {
	switch code {
	case EncounterStatusPlanned:
		return "planned"
	case EncounterStatusInProgress:
		return "in-progress"
	case EncounterStatusOnHold:
		return "on-hold"
	case EncounterStatusDischarged:
		return "discharged"
	case EncounterStatusCompleted:
		return "completed"
	case EncounterStatusCancelled:
		return "cancelled"
	case EncounterStatusDiscontinued:
		return "discontinued"
	case EncounterStatusEnteredInError:
		return "entered-in-error"
	case EncounterStatusUnknown:
		return "unknown"
	}
	return "<unknown>"
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
