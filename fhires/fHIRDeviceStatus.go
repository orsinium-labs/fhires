package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// The status of the Device record.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/device-status
type FHIRDeviceStatus int

const (
	FHIRDeviceStatusActive FHIRDeviceStatus = iota
	// The device record is not current and is not appropriate for reference in new instances.
	FHIRDeviceStatusInactive
	// The device record is not current and is not appropriate for reference in new instances.
	FHIRDeviceStatusEnteredInError
)

func (code FHIRDeviceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *FHIRDeviceStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "active":
		*code = FHIRDeviceStatusActive
	case "inactive":
		*code = FHIRDeviceStatusInactive
	case "entered-in-error":
		*code = FHIRDeviceStatusEnteredInError
	default:
		return fmt.Errorf("unknown FHIRDeviceStatus code `%s`", s)
	}
	return nil
}

func (code FHIRDeviceStatus) String() string {
	return code.Code()
}

func (code FHIRDeviceStatus) Code() string {
	switch code {
	case FHIRDeviceStatusActive:
		return "active"
	case FHIRDeviceStatusInactive:
		return "inactive"
	case FHIRDeviceStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
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
