package fhires

// Additional Binding Purpose
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/additional-binding-purpose
type AdditionalBindingPurposeVS string

const (
	// A required binding, for use when the binding strength is 'extensible' or 'preferred'
	AdditionalBindingPurposeVSMaximum AdditionalBindingPurposeVS = "maximum"
	// The minimum allowable value set - any conformant system SHALL support all these codes
	AdditionalBindingPurposeVSMinimum AdditionalBindingPurposeVS = "minimum"
	// This value set is used as a required binding (in addition to the base binding (not a replacement), usually in a particular usage context)
	AdditionalBindingPurposeVSRequired AdditionalBindingPurposeVS = "required"
	// This value set is used as an extensible binding (in addition to the base binding (not a replacement), usually in a particular usage context)
	AdditionalBindingPurposeVSExtensible AdditionalBindingPurposeVS = "extensible"
	// This value set is a candidate to substitute for the overall conformance value set in some situations; usually these are defined in the documentation
	AdditionalBindingPurposeVSCandidate AdditionalBindingPurposeVS = "candidate"
	// New records are required to use this value set, but legacy records may use other codes. The definition of 'new record' is difficult, since systems often create new records based on pre-existing data. Usually 'current' bindings are mandated by an external authority that makes clear rules around this
	AdditionalBindingPurposeVSCurrent AdditionalBindingPurposeVS = "current"
	// This is the value set that is preferred in a given context (documentation should explain why)
	AdditionalBindingPurposeVSPreferred AdditionalBindingPurposeVS = "preferred"
	// This value set is provided for user look up in a given context. Typically, these valuesets only include a subset of codes relevant for input in a context
	AdditionalBindingPurposeVSUi AdditionalBindingPurposeVS = "ui"
	// This value set is a good set of codes to start with when designing your system
	AdditionalBindingPurposeVSStarter AdditionalBindingPurposeVS = "starter"
	// This value set is a component of the base value set. Usually this is called out so that documentation can be written about a portion of the value set
	AdditionalBindingPurposeVSComponent AdditionalBindingPurposeVS = "component"
)

func (code AdditionalBindingPurposeVS) String() string {
	return code.Code()
}

func (code AdditionalBindingPurposeVS) Code() string {
	return string(code)
}

func (code AdditionalBindingPurposeVS) Display() string {
	switch code {
	case AdditionalBindingPurposeVSMaximum:
		return "Maximum Binding"
	case AdditionalBindingPurposeVSMinimum:
		return "Minimum Binding"
	case AdditionalBindingPurposeVSRequired:
		return "Required Binding"
	case AdditionalBindingPurposeVSExtensible:
		return "Conformance Binding"
	case AdditionalBindingPurposeVSCandidate:
		return "Candidate Binding"
	case AdditionalBindingPurposeVSCurrent:
		return "Current Binding"
	case AdditionalBindingPurposeVSPreferred:
		return "Preferred Binding"
	case AdditionalBindingPurposeVSUi:
		return "UI Suggested Binding"
	case AdditionalBindingPurposeVSStarter:
		return "Starter Binding"
	case AdditionalBindingPurposeVSComponent:
		return "Component Binding"
	}
	return "<unknown>"
}

func (code AdditionalBindingPurposeVS) Definition() string {
	switch code {
	case AdditionalBindingPurposeVSMaximum:
		return "A required binding, for use when the binding strength is 'extensible' or 'preferred'"
	case AdditionalBindingPurposeVSMinimum:
		return "The minimum allowable value set - any conformant system SHALL support all these codes"
	case AdditionalBindingPurposeVSRequired:
		return "This value set is used as a required binding (in addition to the base binding (not a replacement), usually in a particular usage context)"
	case AdditionalBindingPurposeVSExtensible:
		return "This value set is used as an extensible binding (in addition to the base binding (not a replacement), usually in a particular usage context)"
	case AdditionalBindingPurposeVSCandidate:
		return "This value set is a candidate to substitute for the overall conformance value set in some situations; usually these are defined in the documentation"
	case AdditionalBindingPurposeVSCurrent:
		return "New records are required to use this value set, but legacy records may use other codes. The definition of 'new record' is difficult, since systems often create new records based on pre-existing data. Usually 'current' bindings are mandated by an external authority that makes clear rules around this"
	case AdditionalBindingPurposeVSPreferred:
		return "This is the value set that is preferred in a given context (documentation should explain why)"
	case AdditionalBindingPurposeVSUi:
		return "This value set is provided for user look up in a given context. Typically, these valuesets only include a subset of codes relevant for input in a context"
	case AdditionalBindingPurposeVSStarter:
		return "This value set is a good set of codes to start with when designing your system"
	case AdditionalBindingPurposeVSComponent:
		return "This value set is a component of the base value set. Usually this is called out so that documentation can be written about a portion of the value set"
	}
	return "<unknown>"
}
