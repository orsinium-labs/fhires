package fhires

// Device - Corrective action scope
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/device-correctiveactionscope
type DeviceCorrectiveActionScope string

const (
	// The corrective action was intended for all units of the same model.
	DeviceCorrectiveActionScopeModel DeviceCorrectiveActionScope = "model"
	// The corrective action was intended for a specific batch of units identified by a lot number.
	DeviceCorrectiveActionScopeLotNumbers DeviceCorrectiveActionScope = "lot-numbers"
	// The corrective action was intended for an individual unit (or a set of units) individually identified by serial number.
	DeviceCorrectiveActionScopeSerialNumbers DeviceCorrectiveActionScope = "serial-numbers"
)

func (code DeviceCorrectiveActionScope) String() string {
	return code.Code()
}

func (code DeviceCorrectiveActionScope) Code() string {
	return string(code)
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
