package fhires

// The status of the Device record.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/device-status
type FHIRDeviceStatus string

const (
	// The device record is current and is appropriate for reference in new instances.
	FHIRDeviceStatusActive FHIRDeviceStatus = "active"
	// The device record is not current and is not appropriate for reference in new instances.
	FHIRDeviceStatusInactive FHIRDeviceStatus = "inactive"
	// The device record is not current and is not appropriate for reference in new instances.
	FHIRDeviceStatusEnteredInError FHIRDeviceStatus = "entered-in-error"
)

func (code FHIRDeviceStatus) String() string {
	return code.Code()
}

func (code FHIRDeviceStatus) Code() string {
	return string(code)
}

func (code FHIRDeviceStatus) Display() string {
	switch code {
	case FHIRDeviceStatusActive:
		return "Active"
	case FHIRDeviceStatusInactive:
		return "Inactive"
	case FHIRDeviceStatusEnteredInError:
		return "Entered in Error"
	}
	return "<unknown>"
}

func (code FHIRDeviceStatus) Definition() string {
	switch code {
	case FHIRDeviceStatusActive:
		return "The device record is current and is appropriate for reference in new instances."
	case FHIRDeviceStatusInactive:
		return "The device record is not current and is not appropriate for reference in new instances."
	case FHIRDeviceStatusEnteredInError:
		return "The device record is not current and is not appropriate for reference in new instances."
	}
	return "<unknown>"
}
