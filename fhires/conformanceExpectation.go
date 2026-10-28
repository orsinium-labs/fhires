package fhires

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ConformanceExpectation is documented here http://hl7.org/fhir/ValueSet/conformance-expectation
type ConformanceExpectation int

const (
	ConformanceExpectationSHALL ConformanceExpectation = iota
	ConformanceExpectationSHOULD
	ConformanceExpectationMAY
	ConformanceExpectationSHOULDNOT
)

func (code ConformanceExpectation) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *ConformanceExpectation) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "SHALL":
		*code = ConformanceExpectationSHALL
	case "SHOULD":
		*code = ConformanceExpectationSHOULD
	case "MAY":
		*code = ConformanceExpectationMAY
	case "SHOULD-NOT":
		*code = ConformanceExpectationSHOULDNOT
	default:
		return fmt.Errorf("unknown ConformanceExpectation code `%s`", s)
	}
	return nil
}

func (code ConformanceExpectation) String() string {
	return code.Code()
}

func (code ConformanceExpectation) Code() string {
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
