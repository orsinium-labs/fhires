package fhires

// The value set to instantiate this attribute should be drawn from a terminologically robust code system that consists of or contains concepts to support describing the current status of the administered dose of vaccine.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/immunization-status
type ImmunizationStatusCodes string

const (
	// The core event has not started yet, but some staging activities have begun (e.g. surgical suite preparation).  Preparation stages may be tracked for billing purposes.
	ImmunizationStatusCodesPreparation ImmunizationStatusCodes = "preparation"
	// The event is currently occurring.
	ImmunizationStatusCodesInProgress ImmunizationStatusCodes = "in-progress"
	// The event was terminated prior to any activity beyond preparation.  I.e. The 'main' activity has not yet begun.  The boundary between preparatory and the 'main' activity is context-specific.
	ImmunizationStatusCodesNotDone ImmunizationStatusCodes = "not-done"
	// The event has been temporarily stopped but is expected to resume in the future.
	ImmunizationStatusCodesOnHold ImmunizationStatusCodes = "on-hold"
	// The event was terminated prior to the full completion of the intended activity but after at least some of the 'main' activity (beyond preparation) has occurred.
	ImmunizationStatusCodesStopped ImmunizationStatusCodes = "stopped"
	// The event has now concluded.
	ImmunizationStatusCodesCompleted ImmunizationStatusCodes = "completed"
	// This electronic record should never have existed, though it is possible that real-world decisions were based on it.  (If real-world activity has occurred, the status should be "stopped" rather than "entered-in-error".).
	ImmunizationStatusCodesEnteredInError ImmunizationStatusCodes = "entered-in-error"
	// The authoring/source system does not know which of the status values currently applies for this event.  Note: This concept is not to be used for "other" - one of the listed statuses is presumed to apply,  but the authoring/source system does not know which.
	ImmunizationStatusCodesUnknown ImmunizationStatusCodes = "unknown"
)

func (code ImmunizationStatusCodes) String() string {
	return code.Code()
}

func (code ImmunizationStatusCodes) Code() string {
	return string(code)
}

func (code ImmunizationStatusCodes) Display() string {
	switch code {
	case ImmunizationStatusCodesPreparation:
		return "Preparation"
	case ImmunizationStatusCodesInProgress:
		return "In Progress"
	case ImmunizationStatusCodesNotDone:
		return "Not Done"
	case ImmunizationStatusCodesOnHold:
		return "On Hold"
	case ImmunizationStatusCodesStopped:
		return "Stopped"
	case ImmunizationStatusCodesCompleted:
		return "Completed"
	case ImmunizationStatusCodesEnteredInError:
		return "Entered in Error"
	case ImmunizationStatusCodesUnknown:
		return "Unknown"
	}
	return "<unknown>"
}

func (code ImmunizationStatusCodes) Definition() string {
	switch code {
	case ImmunizationStatusCodesPreparation:
		return "The core event has not started yet, but some staging activities have begun (e.g. surgical suite preparation).  Preparation stages may be tracked for billing purposes."
	case ImmunizationStatusCodesInProgress:
		return "The event is currently occurring."
	case ImmunizationStatusCodesNotDone:
		return "The event was terminated prior to any activity beyond preparation.  I.e. The 'main' activity has not yet begun.  The boundary between preparatory and the 'main' activity is context-specific."
	case ImmunizationStatusCodesOnHold:
		return "The event has been temporarily stopped but is expected to resume in the future."
	case ImmunizationStatusCodesStopped:
		return "The event was terminated prior to the full completion of the intended activity but after at least some of the 'main' activity (beyond preparation) has occurred."
	case ImmunizationStatusCodesCompleted:
		return "The event has now concluded."
	case ImmunizationStatusCodesEnteredInError:
		return "This electronic record should never have existed, though it is possible that real-world decisions were based on it.  (If real-world activity has occurred, the status should be \"stopped\" rather than \"entered-in-error\".)."
	case ImmunizationStatusCodesUnknown:
		return "The authoring/source system does not know which of the status values currently applies for this event.  Note: This concept is not to be used for \"other\" - one of the listed statuses is presumed to apply,  but the authoring/source system does not know which."
	}
	return "<unknown>"
}
