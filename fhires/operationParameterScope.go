package fhires

// Indicates that a parameter applies when the operation is being invoked at the specified level
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/operation-parameter-scope
type OperationParameterScope string

const (
	// This is a parameter that can be used at the instance level.
	OperationParameterScopeInstance OperationParameterScope = "instance"
	// This is a parameter that can be used at the type level.
	OperationParameterScopeType OperationParameterScope = "type"
	// This is a parameter that can be used at the system level.
	OperationParameterScopeSystem OperationParameterScope = "system"
)

func (code OperationParameterScope) String() string {
	return code.Code()
}

func (code OperationParameterScope) Code() string {
	return string(code)
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
