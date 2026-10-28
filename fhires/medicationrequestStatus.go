package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// MedicationRequest Status Codes
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/medicationrequest-status
type MedicationrequestStatus int

const (
	MedicationrequestStatusActive MedicationrequestStatus = iota
	// Actions implied by the request are to be temporarily halted. The request might or might not be resumed. May also be called 'suspended'.
	MedicationrequestStatusOnHold
	// The request is no longer active and the subject should no longer be taking the medication.
	MedicationrequestStatusEnded
	// Actions implied by the request are to be permanently halted, before all of the administrations occurred. This should not be used if the original order was entered in error
	MedicationrequestStatusStopped
	// All actions that are implied by the request have occurred.
	MedicationrequestStatusCompleted
	// The request has been withdrawn before any administrations have occurred
	MedicationrequestStatusCancelled
	// The request was recorded against the wrong patient or for some reason should not have been recorded (e.g. wrong medication, wrong dose, etc.). Some of the actions that are implied by the medication request may have occurred. For example, the medication may have been dispensed and the patient may have taken some of the medication.
	MedicationrequestStatusEnteredInError
	// The request is not yet 'actionable', e.g. it is a work in progress, requires sign-off, verification or needs to be run through decision support process.
	MedicationrequestStatusDraft
	// The authoring/source system does not know which of the status values currently applies for this request. Note: This concept is not to be used for 'other' - one of the listed statuses is presumed to apply, but the authoring/source system does not know which.
	MedicationrequestStatusUnknown
)

func (code MedicationrequestStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *MedicationrequestStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "active":
		*code = MedicationrequestStatusActive
	case "on-hold":
		*code = MedicationrequestStatusOnHold
	case "ended":
		*code = MedicationrequestStatusEnded
	case "stopped":
		*code = MedicationrequestStatusStopped
	case "completed":
		*code = MedicationrequestStatusCompleted
	case "cancelled":
		*code = MedicationrequestStatusCancelled
	case "entered-in-error":
		*code = MedicationrequestStatusEnteredInError
	case "draft":
		*code = MedicationrequestStatusDraft
	case "unknown":
		*code = MedicationrequestStatusUnknown
	default:
		return fmt.Errorf("unknown MedicationrequestStatus code `%s`", s)
	}
	return nil
}

func (code MedicationrequestStatus) String() string {
	return code.Code()
}

func (code MedicationrequestStatus) Code() string {
	switch code {
	case MedicationrequestStatusActive:
		return "active"
	case MedicationrequestStatusOnHold:
		return "on-hold"
	case MedicationrequestStatusEnded:
		return "ended"
	case MedicationrequestStatusStopped:
		return "stopped"
	case MedicationrequestStatusCompleted:
		return "completed"
	case MedicationrequestStatusCancelled:
		return "cancelled"
	case MedicationrequestStatusEnteredInError:
		return "entered-in-error"
	case MedicationrequestStatusDraft:
		return "draft"
	case MedicationrequestStatusUnknown:
		return "unknown"
	}
	return "<unknown>"
}

func (code MedicationrequestStatus) Display() string {
	switch code {
	case MedicationrequestStatusActive:
		return "Active"
	case MedicationrequestStatusOnHold:
		return "On Hold"
	case MedicationrequestStatusEnded:
		return "Ended"
	case MedicationrequestStatusStopped:
		return "Stopped"
	case MedicationrequestStatusCompleted:
		return "Completed"
	case MedicationrequestStatusCancelled:
		return "Cancelled"
	case MedicationrequestStatusEnteredInError:
		return "Entered in Error"
	case MedicationrequestStatusDraft:
		return "Draft"
	case MedicationrequestStatusUnknown:
		return "Unknown"
	}
	return "<unknown>"
}

func (code MedicationrequestStatus) Definition() string {
	switch code {
	case MedicationrequestStatusActive:
		return "The request is 'actionable', but not all actions that are implied by it have occurred yet."
	case MedicationrequestStatusOnHold:
		return "Actions implied by the request are to be temporarily halted. The request might or might not be resumed. May also be called 'suspended'."
	case MedicationrequestStatusEnded:
		return "The request is no longer active and the subject should no longer be taking the medication."
	case MedicationrequestStatusStopped:
		return "Actions implied by the request are to be permanently halted, before all of the administrations occurred. This should not be used if the original order was entered in error"
	case MedicationrequestStatusCompleted:
		return "All actions that are implied by the request have occurred."
	case MedicationrequestStatusCancelled:
		return "The request has been withdrawn before any administrations have occurred"
	case MedicationrequestStatusEnteredInError:
		return "The request was recorded against the wrong patient or for some reason should not have been recorded (e.g. wrong medication, wrong dose, etc.). Some of the actions that are implied by the medication request may have occurred. For example, the medication may have been dispensed and the patient may have taken some of the medication."
	case MedicationrequestStatusDraft:
		return "The request is not yet 'actionable', e.g. it is a work in progress, requires sign-off, verification\u00a0or\u00a0needs to be run through decision support process."
	case MedicationrequestStatusUnknown:
		return "The authoring/source system does not know which of the status values currently applies for this request. Note: This concept is not to be used for 'other' - one of the listed statuses is presumed to apply, but the authoring/source system does not know which."
	}
	return "<unknown>"
}
