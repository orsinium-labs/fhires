package fhires

// MedicationAdministration Status Codes
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/medication-admin-status
type MedicationAdministrationStatusCodes string

const (
	// The administration has started but has not yet completed.
	MedicationAdministrationStatusCodesInProgress MedicationAdministrationStatusCodes = "in-progress"
	// The administration was terminated prior to any impact on the subject (though preparatory actions may have been taken)
	MedicationAdministrationStatusCodesNotDone MedicationAdministrationStatusCodes = "not-done"
	// Actions implied by the administration have been temporarily halted, but are expected to continue later. May also be called 'suspended'.
	MedicationAdministrationStatusCodesOnHold MedicationAdministrationStatusCodes = "on-hold"
	// All actions that are implied by the administration have occurred.
	MedicationAdministrationStatusCodesCompleted MedicationAdministrationStatusCodes = "completed"
	// The administration was entered in error and therefore nullified.
	MedicationAdministrationStatusCodesEnteredInError MedicationAdministrationStatusCodes = "entered-in-error"
	// Actions implied by the administration have been permanently halted, before all of them occurred.
	MedicationAdministrationStatusCodesStopped MedicationAdministrationStatusCodes = "stopped"
	// The authoring system does not know which of the status values currently applies for this request. Note: This concept is not to be used for 'other' - one of the listed statuses is presumed to apply, it's just not known which one.
	MedicationAdministrationStatusCodesUnknown MedicationAdministrationStatusCodes = "unknown"
)

func (code MedicationAdministrationStatusCodes) String() string {
	return code.Code()
}

func (code MedicationAdministrationStatusCodes) Code() string {
	return string(code)
}

func (code MedicationAdministrationStatusCodes) Display() string {
	switch code {
	case MedicationAdministrationStatusCodesInProgress:
		return "In Progress"
	case MedicationAdministrationStatusCodesNotDone:
		return "Not Done"
	case MedicationAdministrationStatusCodesOnHold:
		return "On Hold"
	case MedicationAdministrationStatusCodesCompleted:
		return "Completed"
	case MedicationAdministrationStatusCodesEnteredInError:
		return "Entered in Error"
	case MedicationAdministrationStatusCodesStopped:
		return "Stopped"
	case MedicationAdministrationStatusCodesUnknown:
		return "Unknown"
	}
	return "<unknown>"
}

func (code MedicationAdministrationStatusCodes) Definition() string {
	switch code {
	case MedicationAdministrationStatusCodesInProgress:
		return "The administration has started but has not yet completed."
	case MedicationAdministrationStatusCodesNotDone:
		return "The administration was terminated prior to any impact on the subject (though preparatory actions may have been taken)"
	case MedicationAdministrationStatusCodesOnHold:
		return "Actions implied by the administration have been temporarily halted, but are expected to continue later. May also be called 'suspended'."
	case MedicationAdministrationStatusCodesCompleted:
		return "All actions that are implied by the administration have occurred."
	case MedicationAdministrationStatusCodesEnteredInError:
		return "The administration was entered in error and therefore nullified."
	case MedicationAdministrationStatusCodesStopped:
		return "Actions implied by the administration have been permanently halted, before all of them occurred."
	case MedicationAdministrationStatusCodesUnknown:
		return "The authoring system does not know which of the status values currently applies for this request. Note: This concept is not to be used for 'other' - one of the listed statuses is presumed to apply, it's just not known which one."
	}
	return "<unknown>"
}
