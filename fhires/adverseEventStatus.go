package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// Codes identifying the lifecycle stage of an adverse event.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/adverse-event-status
type AdverseEventStatus int

const (
	AdverseEventStatusPreparation AdverseEventStatus = iota
	// The event is currently occurring.
	AdverseEventStatusInProgress
	// The event was terminated prior to any activity beyond preparation.  I.e. The 'main' activity has not yet begun.  The boundary between preparatory and the 'main' activity is context-specific.
	AdverseEventStatusNotDone
	// The event has been temporarily stopped but is expected to resume in the future.
	AdverseEventStatusOnHold
	// The event was terminated prior to the full completion of the intended activity but after at least some of the 'main' activity (beyond preparation) has occurred.
	AdverseEventStatusStopped
	// The event has now concluded.
	AdverseEventStatusCompleted
	// This electronic record should never have existed, though it is possible that real-world decisions were based on it.  (If real-world activity has occurred, the status should be "stopped" rather than "entered-in-error".).
	AdverseEventStatusEnteredInError
	// The authoring/source system does not know which of the status values currently applies for this event.  Note: This concept is not to be used for "other" - one of the listed statuses is presumed to apply,  but the authoring/source system does not know which.
	AdverseEventStatusUnknown
)

func (code AdverseEventStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *AdverseEventStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "preparation":
		*code = AdverseEventStatusPreparation
	case "in-progress":
		*code = AdverseEventStatusInProgress
	case "not-done":
		*code = AdverseEventStatusNotDone
	case "on-hold":
		*code = AdverseEventStatusOnHold
	case "stopped":
		*code = AdverseEventStatusStopped
	case "completed":
		*code = AdverseEventStatusCompleted
	case "entered-in-error":
		*code = AdverseEventStatusEnteredInError
	case "unknown":
		*code = AdverseEventStatusUnknown
	default:
		return fmt.Errorf("unknown AdverseEventStatus code `%s`", s)
	}
	return nil
}

func (code AdverseEventStatus) String() string {
	return code.Code()
}

func (code AdverseEventStatus) Code() string {
	switch code {
	case AdverseEventStatusPreparation:
		return "preparation"
	case AdverseEventStatusInProgress:
		return "in-progress"
	case AdverseEventStatusNotDone:
		return "not-done"
	case AdverseEventStatusOnHold:
		return "on-hold"
	case AdverseEventStatusStopped:
		return "stopped"
	case AdverseEventStatusCompleted:
		return "completed"
	case AdverseEventStatusEnteredInError:
		return "entered-in-error"
	case AdverseEventStatusUnknown:
		return "unknown"
	}
	return "<unknown>"
}

func (code AdverseEventStatus) Display() string {
	switch code {
	case AdverseEventStatusPreparation:
		return "Preparation"
	case AdverseEventStatusInProgress:
		return "In Progress"
	case AdverseEventStatusNotDone:
		return "Not Done"
	case AdverseEventStatusOnHold:
		return "On Hold"
	case AdverseEventStatusStopped:
		return "Stopped"
	case AdverseEventStatusCompleted:
		return "Completed"
	case AdverseEventStatusEnteredInError:
		return "Entered in Error"
	case AdverseEventStatusUnknown:
		return "Unknown"
	}
	return "<unknown>"
}

func (code AdverseEventStatus) Definition() string {
	switch code {
	case AdverseEventStatusPreparation:
		return "The core event has not started yet, but some staging activities have begun (e.g. surgical suite preparation).  Preparation stages may be tracked for billing purposes."
	case AdverseEventStatusInProgress:
		return "The event is currently occurring."
	case AdverseEventStatusNotDone:
		return "The event was terminated prior to any activity beyond preparation.  I.e. The 'main' activity has not yet begun.  The boundary between preparatory and the 'main' activity is context-specific."
	case AdverseEventStatusOnHold:
		return "The event has been temporarily stopped but is expected to resume in the future."
	case AdverseEventStatusStopped:
		return "The event was terminated prior to the full completion of the intended activity but after at least some of the 'main' activity (beyond preparation) has occurred."
	case AdverseEventStatusCompleted:
		return "The event has now concluded."
	case AdverseEventStatusEnteredInError:
		return "This electronic record should never have existed, though it is possible that real-world decisions were based on it.  (If real-world activity has occurred, the status should be \"stopped\" rather than \"entered-in-error\".)."
	case AdverseEventStatusUnknown:
		return "The authoring/source system does not know which of the status values currently applies for this event.  Note: This concept is not to be used for \"other\" - one of the listed statuses is presumed to apply,  but the authoring/source system does not know which."
	}
	return "<unknown>"
}
