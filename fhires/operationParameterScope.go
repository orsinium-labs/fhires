package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// OperationParameterScope is documented here http://hl7.org/fhir/ValueSet/operation-parameter-scope
type OperationParameterScope int

const (
	OperationParameterScopeInstance OperationParameterScope = iota
	OperationParameterScopeType
	OperationParameterScopeSystem
)

func (code OperationParameterScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *OperationParameterScope) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "instance":
		*code = OperationParameterScopeInstance
	case "type":
		*code = OperationParameterScopeType
	case "system":
		*code = OperationParameterScopeSystem
	default:
		return fmt.Errorf("unknown OperationParameterScope code `%s`", s)
	}
	return nil
}

func (code OperationParameterScope) String() string {
	return code.Code()
}

func (code OperationParameterScope) Code() string {
	switch code {
	case OperationParameterScopeInstance:
		return "instance"
	case OperationParameterScopeType:
		return "type"
	case OperationParameterScopeSystem:
		return "system"
	}
	return "<unknown>"
}

func (code OperationParameterScope) Display() string {
	switch code {
	case OperationParameterScopeInstance:
		return "Instance"
	case OperationParameterScopeType:
		return "Type"
	case OperationParameterScopeSystem:
		return "System"
	}
	return "<unknown>"
}

func (code OperationParameterScope) Definition() string {
	switch code {
	case OperationParameterScopeInstance:
		return "This is a parameter that can be used at the instance level."
	case OperationParameterScopeType:
		return "This is a parameter that can be used at the type level."
	case OperationParameterScopeSystem:
		return "This is a parameter that can be used at the system level."
	}
	return "<unknown>"
}
