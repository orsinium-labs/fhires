package fhires

// The type of direction to use for assertion.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/assert-direction-codes
type AssertionDirectionType string

const (
	// The assertion is evaluated on the response. This is the default value.
	AssertionDirectionTypeResponse AssertionDirectionType = "response"
	// The assertion is evaluated on the request.
	AssertionDirectionTypeRequest AssertionDirectionType = "request"
)

func (code AssertionDirectionType) String() string {
	return code.Code()
}

func (code AssertionDirectionType) Code() string {
	return string(code)
}

func (code AssertionDirectionType) Display() string {
	switch code {
	case AssertionDirectionTypeResponse:
		return "response"
	case AssertionDirectionTypeRequest:
		return "request"
	}
	return "<unknown>"
}

func (code AssertionDirectionType) Definition() string {
	switch code {
	case AssertionDirectionTypeResponse:
		return "The assertion is evaluated on the response. This is the default value."
	case AssertionDirectionTypeRequest:
		return "The assertion is evaluated on the request."
	}
	return "<unknown>"
}
