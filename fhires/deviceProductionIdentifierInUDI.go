package fhires

import (
	"encoding/json"
	"fmt"
	"strings"
)

// DeviceProductionIdentifierInUDI is documented here http://hl7.org/fhir/ValueSet/device-productidentifierinudi
type DeviceProductionIdentifierInUDI int

const (
	DeviceProductionIdentifierInUDILotNumber DeviceProductionIdentifierInUDI = iota
	DeviceProductionIdentifierInUDIManufacturedDate
	DeviceProductionIdentifierInUDISerialNumber
	DeviceProductionIdentifierInUDIExpirationDate
	DeviceProductionIdentifierInUDIBiologicalSource
	DeviceProductionIdentifierInUDISoftwareVersion
)

func (code DeviceProductionIdentifierInUDI) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *DeviceProductionIdentifierInUDI) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "lot-number":
		*code = DeviceProductionIdentifierInUDILotNumber
	case "manufactured-date":
		*code = DeviceProductionIdentifierInUDIManufacturedDate
	case "serial-number":
		*code = DeviceProductionIdentifierInUDISerialNumber
	case "expiration-date":
		*code = DeviceProductionIdentifierInUDIExpirationDate
	case "biological-source":
		*code = DeviceProductionIdentifierInUDIBiologicalSource
	case "software-version":
		*code = DeviceProductionIdentifierInUDISoftwareVersion
	default:
		return fmt.Errorf("unknown DeviceProductionIdentifierInUDI code `%s`", s)
	}
	return nil
}

func (code DeviceProductionIdentifierInUDI) String() string {
	return code.Code()
}

func (code DeviceProductionIdentifierInUDI) Code() string {
	switch code {
	case DeviceProductionIdentifierInUDILotNumber:
		return "lot-number"
	case DeviceProductionIdentifierInUDIManufacturedDate:
		return "manufactured-date"
	case DeviceProductionIdentifierInUDISerialNumber:
		return "serial-number"
	case DeviceProductionIdentifierInUDIExpirationDate:
		return "expiration-date"
	case DeviceProductionIdentifierInUDIBiologicalSource:
		return "biological-source"
	case DeviceProductionIdentifierInUDISoftwareVersion:
		return "software-version"
	}
	return "<unknown>"
}

func (code DeviceProductionIdentifierInUDI) Display() string {
	switch code {
	case DeviceProductionIdentifierInUDILotNumber:
		return "Lot Number"
	case DeviceProductionIdentifierInUDIManufacturedDate:
		return "Manufactured date"
	case DeviceProductionIdentifierInUDISerialNumber:
		return "Serial Number"
	case DeviceProductionIdentifierInUDIExpirationDate:
		return "Expiration date"
	case DeviceProductionIdentifierInUDIBiologicalSource:
		return "Biological source"
	case DeviceProductionIdentifierInUDISoftwareVersion:
		return "Software Version"
	}
	return "<unknown>"
}

func (code DeviceProductionIdentifierInUDI) Definition() string {
	switch code {
	case DeviceProductionIdentifierInUDILotNumber:
		return "The label includes the lot number."
	case DeviceProductionIdentifierInUDIManufacturedDate:
		return "The label includes the manufacture date."
	case DeviceProductionIdentifierInUDISerialNumber:
		return "The label includes the serial number."
	case DeviceProductionIdentifierInUDIExpirationDate:
		return "The label includes the expiration date."
	case DeviceProductionIdentifierInUDIBiologicalSource:
		return "The label includes the biological source identifier."
	case DeviceProductionIdentifierInUDISoftwareVersion:
		return "The label includes the software version."
	}
	return "<unknown>"
}
