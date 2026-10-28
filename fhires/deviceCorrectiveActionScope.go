package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// Device - Corrective action scope
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/device-correctiveactionscope
type DeviceCorrectiveActionScope int

const (
	DeviceCorrectiveActionScopeModel DeviceCorrectiveActionScope = iota
	// The corrective action was intended for a specific batch of units identified by a lot number.
	DeviceCorrectiveActionScopeLotNumbers
	// The corrective action was intended for an individual unit (or a set of units) individually identified by serial number.
	DeviceCorrectiveActionScopeSerialNumbers
)

func (code DeviceCorrectiveActionScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *DeviceCorrectiveActionScope) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "model":
		*code = DeviceCorrectiveActionScopeModel
	case "lot-numbers":
		*code = DeviceCorrectiveActionScopeLotNumbers
	case "serial-numbers":
		*code = DeviceCorrectiveActionScopeSerialNumbers
	default:
		return fmt.Errorf("unknown DeviceCorrectiveActionScope code `%s`", s)
	}
	return nil
}

func (code DeviceCorrectiveActionScope) String() string {
	return code.Code()
}

func (code DeviceCorrectiveActionScope) Code() string {
	switch code {
	case DeviceCorrectiveActionScopeModel:
		return "model"
	case DeviceCorrectiveActionScopeLotNumbers:
		return "lot-numbers"
	case DeviceCorrectiveActionScopeSerialNumbers:
		return "serial-numbers"
	}
	return "<unknown>"
}

func (code DeviceCorrectiveActionScope) Display() string {
	switch code {
	case DeviceCorrectiveActionScopeModel:
		return "Model"
	case DeviceCorrectiveActionScopeLotNumbers:
		return "Lot Numbers"
	case DeviceCorrectiveActionScopeSerialNumbers:
		return "Serial Numbers"
	}
	return "<unknown>"
}

func (code DeviceCorrectiveActionScope) Definition() string {
	switch code {
	case DeviceCorrectiveActionScopeModel:
		return "The corrective action was intended for all units of the same model."
	case DeviceCorrectiveActionScopeLotNumbers:
		return "The corrective action was intended for a specific batch of units identified by a lot number."
	case DeviceCorrectiveActionScopeSerialNumbers:
		return "The corrective action was intended for an individual unit (or a set of units) individually identified by serial number."
	}
	return "<unknown>"
}
