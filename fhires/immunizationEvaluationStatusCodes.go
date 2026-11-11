package fhires

// The value set to instantiate this attribute should be drawn from a terminologically robust code system that consists of or contains concepts to support describing the current status of the evaluation for vaccine administration event.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/immunization-evaluation-status
type ImmunizationEvaluationStatusCodes string

const (
	// The administration has started but has not yet completed.
	ImmunizationEvaluationStatusCodesInProgress ImmunizationEvaluationStatusCodes = "in-progress"
	// The administration was terminated prior to any impact on the subject (though preparatory actions may have been taken)
	ImmunizationEvaluationStatusCodesNotDone ImmunizationEvaluationStatusCodes = "not-done"
	// Actions implied by the administration have been temporarily halted, but are expected to continue later. May also be called 'suspended'.
	ImmunizationEvaluationStatusCodesOnHold ImmunizationEvaluationStatusCodes = "on-hold"
	// All actions that are implied by the administration have occurred.
	ImmunizationEvaluationStatusCodesCompleted ImmunizationEvaluationStatusCodes = "completed"
	// The administration was entered in error and therefore nullified.
	ImmunizationEvaluationStatusCodesEnteredInError ImmunizationEvaluationStatusCodes = "entered-in-error"
	// Actions implied by the administration have been permanently halted, before all of them occurred.
	ImmunizationEvaluationStatusCodesStopped ImmunizationEvaluationStatusCodes = "stopped"
	// The authoring system does not know which of the status values currently applies for this request. Note: This concept is not to be used for 'other' - one of the listed statuses is presumed to apply, it's just not known which one.
	ImmunizationEvaluationStatusCodesUnknown ImmunizationEvaluationStatusCodes = "unknown"
)

func (code ImmunizationEvaluationStatusCodes) String() string {
	return code.Code()
}

func (code ImmunizationEvaluationStatusCodes) Code() string {
	return string(code)
}

func (code ImmunizationEvaluationStatusCodes) Display() string {
	switch code {
	case ImmunizationEvaluationStatusCodesInProgress:
		return "In Progress"
	case ImmunizationEvaluationStatusCodesNotDone:
		return "Not Done"
	case ImmunizationEvaluationStatusCodesOnHold:
		return "On Hold"
	case ImmunizationEvaluationStatusCodesCompleted:
		return "Completed"
	case ImmunizationEvaluationStatusCodesEnteredInError:
		return "Entered in Error"
	case ImmunizationEvaluationStatusCodesStopped:
		return "Stopped"
	case ImmunizationEvaluationStatusCodesUnknown:
		return "Unknown"
	}
	return "<unknown>"
}

func (code ImmunizationEvaluationStatusCodes) Definition() string {
	switch code {
	case ImmunizationEvaluationStatusCodesInProgress:
		return "The administration has started but has not yet completed."
	case ImmunizationEvaluationStatusCodesNotDone:
		return "The administration was terminated prior to any impact on the subject (though preparatory actions may have been taken)"
	case ImmunizationEvaluationStatusCodesOnHold:
		return "Actions implied by the administration have been temporarily halted, but are expected to continue later. May also be called 'suspended'."
	case ImmunizationEvaluationStatusCodesCompleted:
		return "All actions that are implied by the administration have occurred."
	case ImmunizationEvaluationStatusCodesEnteredInError:
		return "The administration was entered in error and therefore nullified."
	case ImmunizationEvaluationStatusCodesStopped:
		return "Actions implied by the administration have been permanently halted, before all of them occurred."
	case ImmunizationEvaluationStatusCodesUnknown:
		return "The authoring system does not know which of the status values currently applies for this request. Note: This concept is not to be used for 'other' - one of the listed statuses is presumed to apply, it's just not known which one."
	}
	return "<unknown>"
}
