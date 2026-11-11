package fhires

// The allowable request method or HTTP operation codes.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/http-operations
type TestScriptRequestMethodCode string

const (
	// HTTP DELETE operation.
	TestScriptRequestMethodCodeDelete TestScriptRequestMethodCode = "delete"
	// HTTP GET operation.
	TestScriptRequestMethodCodeGet TestScriptRequestMethodCode = "get"
	// HTTP OPTIONS operation.
	TestScriptRequestMethodCodeOptions TestScriptRequestMethodCode = "options"
	// HTTP PATCH operation.
	TestScriptRequestMethodCodePatch TestScriptRequestMethodCode = "patch"
	// HTTP POST operation.
	TestScriptRequestMethodCodePost TestScriptRequestMethodCode = "post"
	// HTTP PUT operation.
	TestScriptRequestMethodCodePut TestScriptRequestMethodCode = "put"
	// HTTP HEAD operation.
	TestScriptRequestMethodCodeHead TestScriptRequestMethodCode = "head"
)

func (code TestScriptRequestMethodCode) String() string {
	return code.Code()
}

func (code TestScriptRequestMethodCode) Code() string {
	return string(code)
}

func (code TestScriptRequestMethodCode) Display() string {
	switch code {
	case TestScriptRequestMethodCodeDelete:
		return "DELETE"
	case TestScriptRequestMethodCodeGet:
		return "GET"
	case TestScriptRequestMethodCodeOptions:
		return "OPTIONS"
	case TestScriptRequestMethodCodePatch:
		return "PATCH"
	case TestScriptRequestMethodCodePost:
		return "POST"
	case TestScriptRequestMethodCodePut:
		return "PUT"
	case TestScriptRequestMethodCodeHead:
		return "HEAD"
	}
	return "<unknown>"
}

func (code TestScriptRequestMethodCode) Definition() string {
	switch code {
	case TestScriptRequestMethodCodeDelete:
		return "HTTP DELETE operation."
	case TestScriptRequestMethodCodeGet:
		return "HTTP GET operation."
	case TestScriptRequestMethodCodeOptions:
		return "HTTP OPTIONS operation."
	case TestScriptRequestMethodCodePatch:
		return "HTTP PATCH operation."
	case TestScriptRequestMethodCodePost:
		return "HTTP POST operation."
	case TestScriptRequestMethodCodePut:
		return "HTTP PUT operation."
	case TestScriptRequestMethodCodeHead:
		return "HTTP HEAD operation."
	}
	return "<unknown>"
}
