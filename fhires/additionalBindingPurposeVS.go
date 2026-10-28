package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// Additional Binding Purpose
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/additional-binding-purpose
type AdditionalBindingPurposeVS int

const (
	AdditionalBindingPurposeVSMaximum AdditionalBindingPurposeVS = iota
	// The minimum allowable value set - any conformant system SHALL support all these codes
	AdditionalBindingPurposeVSMinimum
	// This value set is used as a required binding (in addition to the base binding (not a replacement), usually in a particular usage context)
	AdditionalBindingPurposeVSRequired
	// This value set is used as an extensible binding (in addition to the base binding (not a replacement), usually in a particular usage context)
	AdditionalBindingPurposeVSExtensible
	// This value set is a candidate to substitute for the overall conformance value set in some situations; usually these are defined in the documentation
	AdditionalBindingPurposeVSCandidate
	// New records are required to use this value set, but legacy records may use other codes. The definition of 'new record' is difficult, since systems often create new records based on pre-existing data. Usually 'current' bindings are mandated by an external authority that makes clear rules around this
	AdditionalBindingPurposeVSCurrent
	// This is the value set that is preferred in a given context (documentation should explain why)
	AdditionalBindingPurposeVSPreferred
	// This value set is provided for user look up in a given context. Typically, these valuesets only include a subset of codes relevant for input in a context
	AdditionalBindingPurposeVSUi
	// This value set is a good set of codes to start with when designing your system
	AdditionalBindingPurposeVSStarter
	// This value set is a component of the base value set. Usually this is called out so that documentation can be written about a portion of the value set
	AdditionalBindingPurposeVSComponent
)

func (code AdditionalBindingPurposeVS) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *AdditionalBindingPurposeVS) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "maximum":
		*code = AdditionalBindingPurposeVSMaximum
	case "minimum":
		*code = AdditionalBindingPurposeVSMinimum
	case "required":
		*code = AdditionalBindingPurposeVSRequired
	case "extensible":
		*code = AdditionalBindingPurposeVSExtensible
	case "candidate":
		*code = AdditionalBindingPurposeVSCandidate
	case "current":
		*code = AdditionalBindingPurposeVSCurrent
	case "preferred":
		*code = AdditionalBindingPurposeVSPreferred
	case "ui":
		*code = AdditionalBindingPurposeVSUi
	case "starter":
		*code = AdditionalBindingPurposeVSStarter
	case "component":
		*code = AdditionalBindingPurposeVSComponent
	default:
		return fmt.Errorf("unknown AdditionalBindingPurposeVS code `%s`", s)
	}
	return nil
}

func (code AdditionalBindingPurposeVS) String() string {
	return code.Code()
}

func (code AdditionalBindingPurposeVS) Code() string {
	switch code {
	case AdditionalBindingPurposeVSMaximum:
		return "maximum"
	case AdditionalBindingPurposeVSMinimum:
		return "minimum"
	case AdditionalBindingPurposeVSRequired:
		return "required"
	case AdditionalBindingPurposeVSExtensible:
		return "extensible"
	case AdditionalBindingPurposeVSCandidate:
		return "candidate"
	case AdditionalBindingPurposeVSCurrent:
		return "current"
	case AdditionalBindingPurposeVSPreferred:
		return "preferred"
	case AdditionalBindingPurposeVSUi:
		return "ui"
	case AdditionalBindingPurposeVSStarter:
		return "starter"
	case AdditionalBindingPurposeVSComponent:
		return "component"
	}
	return "<unknown>"
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
