package fhires

// How a compartment must be linked.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/graph-compartment-rule
type GraphCompartmentRule string

const (
	// The compartment must be identical (the same literal reference).
	GraphCompartmentRuleIdentical GraphCompartmentRule = "identical"
	// The compartment must be the same - the record must be about the same patient, but the reference may be different.
	GraphCompartmentRuleMatching GraphCompartmentRule = "matching"
	// The compartment must be different.
	GraphCompartmentRuleDifferent GraphCompartmentRule = "different"
	// The compartment rule is defined in the accompanying FHIRPath expression.
	GraphCompartmentRuleCustom GraphCompartmentRule = "custom"
)

func (code GraphCompartmentRule) String() string {
	return code.Code()
}

func (code GraphCompartmentRule) Code() string {
	return string(code)
}

func (code GraphCompartmentRule) Display() string {
	switch code {
	case GraphCompartmentRuleIdentical:
		return "Identical"
	case GraphCompartmentRuleMatching:
		return "Matching"
	case GraphCompartmentRuleDifferent:
		return "Different"
	case GraphCompartmentRuleCustom:
		return "Custom"
	}
	return "<unknown>"
}

func (code GraphCompartmentRule) Definition() string {
	switch code {
	case GraphCompartmentRuleIdentical:
		return "The compartment must be identical (the same literal reference)."
	case GraphCompartmentRuleMatching:
		return "The compartment must be the same - the record must be about the same patient, but the reference may be different."
	case GraphCompartmentRuleDifferent:
		return "The compartment must be different."
	case GraphCompartmentRuleCustom:
		return "The compartment rule is defined in the accompanying FHIRPath expression."
	}
	return "<unknown>"
}
