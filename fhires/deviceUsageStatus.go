package fhires

// A coded concept indicating the current status of the Device Usage.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/deviceusage-status
type DeviceUsageStatus string

const (
	// The device is still being used.
	DeviceUsageStatusActive DeviceUsageStatus = "active"
	// The device is no longer being used.
	DeviceUsageStatusCompleted DeviceUsageStatus = "completed"
	// The device was not used.
	DeviceUsageStatusNotDone DeviceUsageStatus = "not-done"
	// The statement was recorded incorrectly.
	DeviceUsageStatusEnteredInError DeviceUsageStatus = "entered-in-error"
	// The device may be used at some time in the future.
	DeviceUsageStatusIntended DeviceUsageStatus = "intended"
	// Actions implied by the statement have been permanently halted, before all of them occurred.
	DeviceUsageStatusStopped DeviceUsageStatus = "stopped"
	// Actions implied by the statement have been temporarily halted, but are expected to continue later. May also be called "suspended".
	DeviceUsageStatusOnHold DeviceUsageStatus = "on-hold"
)

func (code DeviceUsageStatus) String() string {
	return code.Code()
}

func (code DeviceUsageStatus) Code() string {
	return string(code)
}

func (code DeviceUsageStatus) Display() string {
	switch code {
	case DeviceUsageStatusActive:
		return "Active"
	case DeviceUsageStatusCompleted:
		return "Completed"
	case DeviceUsageStatusNotDone:
		return "Not done"
	case DeviceUsageStatusEnteredInError:
		return "Entered in Error"
	case DeviceUsageStatusIntended:
		return "Intended"
	case DeviceUsageStatusStopped:
		return "Stopped"
	case DeviceUsageStatusOnHold:
		return "On Hold"
	}
	return "<unknown>"
}

func (code DeviceUsageStatus) Definition() string {
	switch code {
	case DeviceUsageStatusActive:
		return "The device is still being used."
	case DeviceUsageStatusCompleted:
		return "The device is no longer being used."
	case DeviceUsageStatusNotDone:
		return "The device was not used."
	case DeviceUsageStatusEnteredInError:
		return "The statement was recorded incorrectly."
	case DeviceUsageStatusIntended:
		return "The device may be used at some time in the future."
	case DeviceUsageStatusStopped:
		return "Actions implied by the statement have been permanently halted, before all of them occurred."
	case DeviceUsageStatusOnHold:
		return "Actions implied by the statement have been temporarily halted, but are expected to continue later. May also be called \"suspended\"."
	}
	return "<unknown>"
}
