package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// CriteriaNotExistsBehavior is documented here http://hl7.org/fhir/ValueSet/subscriptiontopic-cr-behavior
type CriteriaNotExistsBehavior int

const (
	CriteriaNotExistsBehaviorTestPasses CriteriaNotExistsBehavior = iota
	CriteriaNotExistsBehaviorTestFails
)

func (code CriteriaNotExistsBehavior) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *CriteriaNotExistsBehavior) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "test-passes":
		*code = CriteriaNotExistsBehaviorTestPasses
	case "test-fails":
		*code = CriteriaNotExistsBehaviorTestFails
	default:
		return fmt.Errorf("unknown CriteriaNotExistsBehavior code `%s`", s)
	}
	return nil
}

func (code CriteriaNotExistsBehavior) String() string {
	return code.Code()
}

func (code CriteriaNotExistsBehavior) Code() string {
	switch code {
	case CriteriaNotExistsBehaviorTestPasses:
		return "test-passes"
	case CriteriaNotExistsBehaviorTestFails:
		return "test-fails"
	}
	return "<unknown>"
}

func (code CriteriaNotExistsBehavior) Display() string {
	switch code {
	case CriteriaNotExistsBehaviorTestPasses:
		return "Test passes"
	case CriteriaNotExistsBehaviorTestFails:
		return "Test fails"
	}
	return "<unknown>"
}

func (code CriteriaNotExistsBehavior) Definition() string {
	switch code {
	case CriteriaNotExistsBehaviorTestPasses:
		return "The requested conditional statement will pass if a matching state does not exist (e.g., previous state during create)."
	case CriteriaNotExistsBehaviorTestFails:
		return "The requested conditional statement will fail if a matching state does not exist (e.g., previous state during create)."
	}
	return "<unknown>"
}
