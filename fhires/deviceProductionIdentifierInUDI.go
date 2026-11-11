package fhires

// Device Production Identifier in UDI
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/device-productidentifierinudi
type DeviceProductionIdentifierInUDI string

const (
	// The label includes the lot number.
	DeviceProductionIdentifierInUDILotNumber DeviceProductionIdentifierInUDI = "lot-number"
	// The label includes the manufacture date.
	DeviceProductionIdentifierInUDIManufacturedDate DeviceProductionIdentifierInUDI = "manufactured-date"
	// The label includes the serial number.
	DeviceProductionIdentifierInUDISerialNumber DeviceProductionIdentifierInUDI = "serial-number"
	// The label includes the expiration date.
	DeviceProductionIdentifierInUDIExpirationDate DeviceProductionIdentifierInUDI = "expiration-date"
	// The label includes the biological source identifier.
	DeviceProductionIdentifierInUDIBiologicalSource DeviceProductionIdentifierInUDI = "biological-source"
	// The label includes the software version.
	DeviceProductionIdentifierInUDISoftwareVersion DeviceProductionIdentifierInUDI = "software-version"
)

func (code DeviceProductionIdentifierInUDI) String() string {
	return code.Code()
}

func (code DeviceProductionIdentifierInUDI) Code() string {
	return string(code)
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
