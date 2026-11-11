package fhires

// The free/busy status of the slot.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/slotstatus
type SlotStatus string

const (
	// Indicates that the time interval is busy because one  or more events have been scheduled for that interval.
	SlotStatusBusy SlotStatus = "busy"
	// Indicates that the time interval is free for scheduling.
	SlotStatusFree SlotStatus = "free"
	// Indicates that the time interval is busy and that the interval cannot be scheduled.
	SlotStatusBusyUnavailable SlotStatus = "busy-unavailable"
	// Indicates that the time interval is busy because one or more events have been tentatively scheduled for that interval.
	SlotStatusBusyTentative SlotStatus = "busy-tentative"
	// This instance should not have been part of this patient's medical record.
	SlotStatusEnteredInError SlotStatus = "entered-in-error"
)

func (code SlotStatus) String() string {
	return code.Code()
}

func (code SlotStatus) Code() string {
	return string(code)
}

func (code SlotStatus) Display() string {
	switch code {
	case SlotStatusBusy:
		return "Busy"
	case SlotStatusFree:
		return "Free"
	case SlotStatusBusyUnavailable:
		return "Busy (Unavailable)"
	case SlotStatusBusyTentative:
		return "Busy (Tentative)"
	case SlotStatusEnteredInError:
		return "Entered in error"
	}
	return "<unknown>"
}

func (code SlotStatus) Definition() string {
	switch code {
	case SlotStatusBusy:
		return "Indicates that the time interval is busy because one  or more events have been scheduled for that interval."
	case SlotStatusFree:
		return "Indicates that the time interval is free for scheduling."
	case SlotStatusBusyUnavailable:
		return "Indicates that the time interval is busy and that the interval cannot be scheduled."
	case SlotStatusBusyTentative:
		return "Indicates that the time interval is busy because one or more events have been tentatively scheduled for that interval."
	case SlotStatusEnteredInError:
		return "This instance should not have been part of this patient's medical record."
	}
	return "<unknown>"
}
