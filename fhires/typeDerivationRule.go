package fhires

// How a type relates to its baseDefinition.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/type-derivation-rule
type TypeDerivationRule string

const (
	// This definition defines a new type that adds additional elements and optionally additional rules to the base type.
	TypeDerivationRuleSpecialization TypeDerivationRule = "specialization"
	// This definition adds additional rules to an existing concrete type.
	TypeDerivationRuleConstraint TypeDerivationRule = "constraint"
)

func (code TypeDerivationRule) String() string {
	return code.Code()
}

func (code TypeDerivationRule) Code() string {
	return string(code)
}

func (code TypeDerivationRule) Display() string {
	switch code {
	case TypeDerivationRuleSpecialization:
		return "Specialization"
	case TypeDerivationRuleConstraint:
		return "Constraint"
	}
	return "<unknown>"
}

func (code TypeDerivationRule) Definition() string {
	switch code {
	case TypeDerivationRuleSpecialization:
		return "This definition defines a new type that adds additional elements and optionally additional rules to the base type."
	case TypeDerivationRuleConstraint:
		return "This definition adds additional rules to an existing concrete type."
	}
	return "<unknown>"
}
