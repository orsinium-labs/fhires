package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// Whether a reference needs to be version specific or version independent, or whether either can be used.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/reference-version-rules
type ReferenceVersionRules int

const (
	ReferenceVersionRulesEither ReferenceVersionRules = iota
	// The reference must be version independent.
	ReferenceVersionRulesIndependent
	// The reference must be version specific.
	ReferenceVersionRulesSpecific
)

func (code ReferenceVersionRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *ReferenceVersionRules) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "either":
		*code = ReferenceVersionRulesEither
	case "independent":
		*code = ReferenceVersionRulesIndependent
	case "specific":
		*code = ReferenceVersionRulesSpecific
	default:
		return fmt.Errorf("unknown ReferenceVersionRules code `%s`", s)
	}
	return nil
}

func (code ReferenceVersionRules) String() string {
	return code.Code()
}

func (code ReferenceVersionRules) Code() string {
	switch code {
	case ReferenceVersionRulesEither:
		return "either"
	case ReferenceVersionRulesIndependent:
		return "independent"
	case ReferenceVersionRulesSpecific:
		return "specific"
	}
	return "<unknown>"
}

func (code ReferenceVersionRules) Display() string {
	switch code {
	case ReferenceVersionRulesEither:
		return "Either Specific or independent"
	case ReferenceVersionRulesIndependent:
		return "Version independent"
	case ReferenceVersionRulesSpecific:
		return "Version Specific"
	}
	return "<unknown>"
}

func (code ReferenceVersionRules) Definition() string {
	switch code {
	case ReferenceVersionRulesEither:
		return "The reference may be either version independent or version specific."
	case ReferenceVersionRulesIndependent:
		return "The reference must be version independent."
	case ReferenceVersionRulesSpecific:
		return "The reference must be version specific."
	}
	return "<unknown>"
}
