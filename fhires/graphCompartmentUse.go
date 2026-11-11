package fhires

// Defines how a compartment rule is used.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/graph-compartment-use
type GraphCompartmentUse string

const (
	// This compartment rule is a condition for whether the rule applies.
	GraphCompartmentUseWhere GraphCompartmentUse = "where"
	// This compartment rule is enforced on any relationships that meet the conditions.
	GraphCompartmentUseRequires GraphCompartmentUse = "requires"
)

func (code GraphCompartmentUse) String() string {
	return code.Code()
}

func (code GraphCompartmentUse) Code() string {
	return string(code)
}

func (code GraphCompartmentUse) Display() string {
	switch code {
	case GraphCompartmentUseWhere:
		return "Where"
	case GraphCompartmentUseRequires:
		return "requires"
	}
	return "<unknown>"
}

func (code GraphCompartmentUse) Definition() string {
	switch code {
	case GraphCompartmentUseWhere:
		return "This compartment rule is a condition for whether the rule applies."
	case GraphCompartmentUseRequires:
		return "This compartment rule is enforced on any relationships that meet the conditions."
	}
	return "<unknown>"
}
