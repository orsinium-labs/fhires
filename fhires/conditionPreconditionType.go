package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// Kind of precondition for the condition.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/condition-precondition-type
type ConditionPreconditionType int

const (
	ConditionPreconditionTypeSensitive ConditionPreconditionType = iota
	// The observation is very specific for this condition, but not particularly sensitive.
	ConditionPreconditionTypeSpecific
)

func (code ConditionPreconditionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *ConditionPreconditionType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "sensitive":
		*code = ConditionPreconditionTypeSensitive
	case "specific":
		*code = ConditionPreconditionTypeSpecific
	default:
		return fmt.Errorf("unknown ConditionPreconditionType code `%s`", s)
	}
	return nil
}

func (code ConditionPreconditionType) String() string {
	return code.Code()
}

func (code ConditionPreconditionType) Code() string {
	switch code {
	case ConditionPreconditionTypeSensitive:
		return "sensitive"
	case ConditionPreconditionTypeSpecific:
		return "specific"
	}
	return "<unknown>"
}

func (code ConditionPreconditionType) Display() string {
	switch code {
	case ConditionPreconditionTypeSensitive:
		return "Sensitive"
	case ConditionPreconditionTypeSpecific:
		return "Specific"
	}
	return "<unknown>"
}

func (code ConditionPreconditionType) Definition() string {
	switch code {
	case ConditionPreconditionTypeSensitive:
		return "The observation is very sensitive for the condition, but may also indicate other conditions."
	case ConditionPreconditionTypeSpecific:
		return "The observation is very specific for this condition, but not particularly sensitive."
	}
	return "<unknown>"
}
