package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// Degree of preference of a type of conditioned specimen.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/specimen-contained-preference
type SpecimenContainedPreference int

const (
	SpecimenContainedPreferencePreferred SpecimenContainedPreference = iota
	// This type of conditioned specimen is an alternate.
	SpecimenContainedPreferenceAlternate
)

func (code SpecimenContainedPreference) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *SpecimenContainedPreference) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "preferred":
		*code = SpecimenContainedPreferencePreferred
	case "alternate":
		*code = SpecimenContainedPreferenceAlternate
	default:
		return fmt.Errorf("unknown SpecimenContainedPreference code `%s`", s)
	}
	return nil
}

func (code SpecimenContainedPreference) String() string {
	return code.Code()
}

func (code SpecimenContainedPreference) Code() string {
	switch code {
	case SpecimenContainedPreferencePreferred:
		return "preferred"
	case SpecimenContainedPreferenceAlternate:
		return "alternate"
	}
	return "<unknown>"
}

func (code SpecimenContainedPreference) Display() string {
	switch code {
	case SpecimenContainedPreferencePreferred:
		return "Preferred"
	case SpecimenContainedPreferenceAlternate:
		return "Alternate"
	}
	return "<unknown>"
}

func (code SpecimenContainedPreference) Definition() string {
	switch code {
	case SpecimenContainedPreferencePreferred:
		return "This type of contained specimen is preferred to collect this kind of specimen."
	case SpecimenContainedPreferenceAlternate:
		return "This type of conditioned specimen is an alternate."
	}
	return "<unknown>"
}
