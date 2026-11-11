package fhires

// Whether an operation is a normal operation or a query.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/operation-kind
type OperationKind string

const (
	// This operation is invoked as an operation.
	OperationKindOperation OperationKind = "operation"
	// This operation is a named query, invoked using the search mechanism.
	OperationKindQuery OperationKind = "query"
)

func (code OperationKind) String() string {
	return code.Code()
}

func (code OperationKind) Code() string {
	return string(code)
}

func (code OperationKind) Display() string {
	switch code {
	case OperationKindOperation:
		return "Operation"
	case OperationKindQuery:
		return "Query"
	}
	return "<unknown>"
}

func (code OperationKind) Definition() string {
	switch code {
	case OperationKindOperation:
		return "This operation is invoked as an operation."
	case OperationKindQuery:
		return "This operation is a named query, invoked using the search mechanism."
	}
	return "<unknown>"
}
