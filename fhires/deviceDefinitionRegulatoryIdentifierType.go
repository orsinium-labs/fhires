package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// DeviceDefinitionRegulatoryIdentifierType is documented here http://hl7.org/fhir/ValueSet/devicedefinition-regulatory-identifier-type
type DeviceDefinitionRegulatoryIdentifierType int

const (
	DeviceDefinitionRegulatoryIdentifierTypeBasic DeviceDefinitionRegulatoryIdentifierType = iota
	DeviceDefinitionRegulatoryIdentifierTypeMaster
	DeviceDefinitionRegulatoryIdentifierTypeLicense
)

func (code DeviceDefinitionRegulatoryIdentifierType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *DeviceDefinitionRegulatoryIdentifierType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "basic":
		*code = DeviceDefinitionRegulatoryIdentifierTypeBasic
	case "master":
		*code = DeviceDefinitionRegulatoryIdentifierTypeMaster
	case "license":
		*code = DeviceDefinitionRegulatoryIdentifierTypeLicense
	default:
		return fmt.Errorf("unknown DeviceDefinitionRegulatoryIdentifierType code `%s`", s)
	}
	return nil
}

func (code DeviceDefinitionRegulatoryIdentifierType) String() string {
	return code.Code()
}

func (code DeviceDefinitionRegulatoryIdentifierType) Code() string {
	switch code {
	case DeviceDefinitionRegulatoryIdentifierTypeBasic:
		return "basic"
	case DeviceDefinitionRegulatoryIdentifierTypeMaster:
		return "master"
	case DeviceDefinitionRegulatoryIdentifierTypeLicense:
		return "license"
	}
	return "<unknown>"
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
