package fhires

// Whether an operation parameter is an input or an output parameter.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/operation-parameter-use
type OperationParameterUse string

const (
	// This is an input parameter.
	OperationParameterUseIn OperationParameterUse = "in"
	// This is an output parameter.
	OperationParameterUseOut OperationParameterUse = "out"
)

func (code OperationParameterUse) String() string {
	return code.Code()
}

func (code OperationParameterUse) Code() string {
	return string(code)
}

func (code OperationParameterUse) Display() string {
	switch code {
	case OperationParameterUseIn:
		return "In"
	case OperationParameterUseOut:
		return "Out"
	}
	return "<unknown>"
}

func (code OperationParameterUse) Definition() string {
	switch code {
	case OperationParameterUseIn:
		return "This is an input parameter."
	case OperationParameterUseOut:
		return "This is an output parameter."
	}
	return "<unknown>"
}
