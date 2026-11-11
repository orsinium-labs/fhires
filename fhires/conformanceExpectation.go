package fhires

// Description Needed Here
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/conformance-expectation
type ConformanceExpectation string

const (
	// Support for the specified capability is required to be considered conformant.
	ConformanceExpectationSHALL ConformanceExpectation = "SHALL"
	// Support for the specified capability is strongly encouraged, and failure to support it should only occur after careful consideration.
	ConformanceExpectationSHOULD ConformanceExpectation = "SHOULD"
	// Support for the specified capability is not necessary to be considered conformant, and the requirement should be considered strictly optional.
	ConformanceExpectationMAY ConformanceExpectation = "MAY"
	// Support for the specified capability is strongly discouraged and should occur only after careful consideration.
	ConformanceExpectationSHOULDNOT ConformanceExpectation = "SHOULD-NOT"
)

func (code ConformanceExpectation) String() string {
	return code.Code()
}

func (code ConformanceExpectation) Code() string {
	return string(code)
}

func (code ConformanceExpectation) Display() string {
	switch code {
	case ConformanceExpectationSHALL:
		return "SHALL"
	case ConformanceExpectationSHOULD:
		return "SHOULD"
	case ConformanceExpectationMAY:
		return "MAY"
	case ConformanceExpectationSHOULDNOT:
		return "SHOULD-NOT"
	}
	return "<unknown>"
}

func (code ConformanceExpectation) Definition() string {
	switch code {
	case ConformanceExpectationSHALL:
		return "Support for the specified capability is required to be considered conformant."
	case ConformanceExpectationSHOULD:
		return "Support for the specified capability is strongly encouraged, and failure to support it should only occur after careful consideration."
	case ConformanceExpectationMAY:
		return "Support for the specified capability is not necessary to be considered conformant, and the requirement should be considered strictly optional."
	case ConformanceExpectationSHOULDNOT:
		return "Support for the specified capability is strongly discouraged and should occur only after careful consideration."
	}
	return "<unknown>"
}
