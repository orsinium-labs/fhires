package fhires

// Regulatory Identifier type
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/devicedefinition-regulatory-identifier-type
type DeviceDefinitionRegulatoryIdentifierType string

const (
	// EUDAMED's basic UDI-DI identifier.
	DeviceDefinitionRegulatoryIdentifierTypeBasic DeviceDefinitionRegulatoryIdentifierType = "basic"
	// EUDAMED's master UDI-DI identifier.
	DeviceDefinitionRegulatoryIdentifierTypeMaster DeviceDefinitionRegulatoryIdentifierType = "master"
	// The identifier is a license number.
	DeviceDefinitionRegulatoryIdentifierTypeLicense DeviceDefinitionRegulatoryIdentifierType = "license"
)

func (code DeviceDefinitionRegulatoryIdentifierType) String() string {
	return code.Code()
}

func (code DeviceDefinitionRegulatoryIdentifierType) Code() string {
	return string(code)
}

func (code DeviceDefinitionRegulatoryIdentifierType) Display() string {
	switch code {
	case DeviceDefinitionRegulatoryIdentifierTypeBasic:
		return "Basic"
	case DeviceDefinitionRegulatoryIdentifierTypeMaster:
		return "Master"
	case DeviceDefinitionRegulatoryIdentifierTypeLicense:
		return "License"
	}
	return "<unknown>"
}

func (code DeviceDefinitionRegulatoryIdentifierType) Definition() string {
	switch code {
	case DeviceDefinitionRegulatoryIdentifierTypeBasic:
		return "EUDAMED's basic UDI-DI identifier."
	case DeviceDefinitionRegulatoryIdentifierTypeMaster:
		return "EUDAMED's master UDI-DI identifier."
	case DeviceDefinitionRegulatoryIdentifierTypeLicense:
		return "The identifier is a license number."
	}
	return "<unknown>"
}
