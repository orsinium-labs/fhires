package fhires

import (
	"encoding/json"
	"fmt"
	"strings"
)

// SpecimenCombined is documented here http://hl7.org/fhir/ValueSet/specimen-combined
type SpecimenCombined int

const (
	SpecimenCombinedGrouped SpecimenCombined = iota
	SpecimenCombinedPooled
)

func (code SpecimenCombined) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *SpecimenCombined) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "grouped":
		*code = SpecimenCombinedGrouped
	case "pooled":
		*code = SpecimenCombinedPooled
	default:
		return fmt.Errorf("unknown SpecimenCombined code `%s`", s)
	}
	return nil
}

func (code SpecimenCombined) String() string {
	return code.Code()
}

func (code SpecimenCombined) Code() string {
	switch code {
	case SpecimenCombinedGrouped:
		return "grouped"
	case SpecimenCombinedPooled:
		return "pooled"
	}
	return "<unknown>"
}

func (code SpecimenCombined) Display() string {
	switch code {
	case SpecimenCombinedGrouped:
		return "Grouped"
	case SpecimenCombinedPooled:
		return "Pooled"
	}
	return "<unknown>"
}

func (code SpecimenCombined) Definition() string {
	switch code {
	case SpecimenCombinedGrouped:
		return "The specimen is in a group."
	case SpecimenCombinedPooled:
		return "The specimen is pooled."
	}
	return "<unknown>"
}
