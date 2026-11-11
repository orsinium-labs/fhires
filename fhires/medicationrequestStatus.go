package fhires

// MedicationRequest Status Codes
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/medicationrequest-status
type MedicationrequestStatus string

const (
	// The request is 'actionable', but not all actions that are implied by it have occurred yet.
	MedicationrequestStatusActive MedicationrequestStatus = "active"
	// Actions implied by the request are to be temporarily halted. The request might or might not be resumed. May also be called 'suspended'.
	MedicationrequestStatusOnHold MedicationrequestStatus = "on-hold"
	// The request is no longer active and the subject should no longer be taking the medication.
	MedicationrequestStatusEnded MedicationrequestStatus = "ended"
	// Actions implied by the request are to be permanently halted, before all of the administrations occurred. This should not be used if the original order was entered in error
	MedicationrequestStatusStopped MedicationrequestStatus = "stopped"
	// All actions that are implied by the request have occurred.
	MedicationrequestStatusCompleted MedicationrequestStatus = "completed"
	// The request has been withdrawn before any administrations have occurred
	MedicationrequestStatusCancelled MedicationrequestStatus = "cancelled"
	// The request was recorded against the wrong patient or for some reason should not have been recorded (e.g. wrong medication, wrong dose, etc.). Some of the actions that are implied by the medication request may have occurred. For example, the medication may have been dispensed and the patient may have taken some of the medication.
	MedicationrequestStatusEnteredInError MedicationrequestStatus = "entered-in-error"
	// The request is not yet 'actionable', e.g. it is a work in progress, requires sign-off, verification or needs to be run through decision support process.
	MedicationrequestStatusDraft MedicationrequestStatus = "draft"
	// The authoring/source system does not know which of the status values currently applies for this request. Note: This concept is not to be used for 'other' - one of the listed statuses is presumed to apply, but the authoring/source system does not know which.
	MedicationrequestStatusUnknown MedicationrequestStatus = "unknown"
)

func (code MedicationrequestStatus) String() string {
	return code.Code()
}

func (code MedicationrequestStatus) Code() string {
	return string(code)
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
