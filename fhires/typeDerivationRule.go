package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// How a type relates to its baseDefinition.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/type-derivation-rule
type TypeDerivationRule int

const (
	TypeDerivationRuleSpecialization TypeDerivationRule = iota
	// This definition adds additional rules to an existing concrete type.
	TypeDerivationRuleConstraint
)

func (code TypeDerivationRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *TypeDerivationRule) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "specialization":
		*code = TypeDerivationRuleSpecialization
	case "constraint":
		*code = TypeDerivationRuleConstraint
	default:
		return fmt.Errorf("unknown TypeDerivationRule code `%s`", s)
	}
	return nil
}

func (code TypeDerivationRule) String() string {
	return code.Code()
}

func (code TypeDerivationRule) Code() string {
	switch code {
	case TypeDerivationRuleSpecialization:
		return "specialization"
	case TypeDerivationRuleConstraint:
		return "constraint"
	}
	return "<unknown>"
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
