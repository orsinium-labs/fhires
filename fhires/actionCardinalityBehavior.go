package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// Defines behavior for an action or a group for how many times that item may be repeated.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/action-cardinality-behavior
type ActionCardinalityBehavior int

const (
	ActionCardinalityBehaviorSingle ActionCardinalityBehavior = iota
	// The action may be selected multiple times.
	ActionCardinalityBehaviorMultiple
)

func (code ActionCardinalityBehavior) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *ActionCardinalityBehavior) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "single":
		*code = ActionCardinalityBehaviorSingle
	case "multiple":
		*code = ActionCardinalityBehaviorMultiple
	default:
		return fmt.Errorf("unknown ActionCardinalityBehavior code `%s`", s)
	}
	return nil
}

func (code ActionCardinalityBehavior) String() string {
	return code.Code()
}

func (code ActionCardinalityBehavior) Code() string {
	switch code {
	case ActionCardinalityBehaviorSingle:
		return "single"
	case ActionCardinalityBehaviorMultiple:
		return "multiple"
	}
	return "<unknown>"
}

func (code ActionCardinalityBehavior) Display() string {
	switch code {
	case ActionCardinalityBehaviorSingle:
		return "Single"
	case ActionCardinalityBehaviorMultiple:
		return "Multiple"
	}
	return "<unknown>"
}

func (code ActionCardinalityBehavior) Definition() string {
	switch code {
	case ActionCardinalityBehaviorSingle:
		return "The action may only be selected one time."
	case ActionCardinalityBehaviorMultiple:
		return "The action may be selected multiple times."
	}
	return "<unknown>"
}
