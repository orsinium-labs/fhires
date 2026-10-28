package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// OrientationType is documented here http://hl7.org/fhir/ValueSet/orientation-type
type OrientationType int

const (
	OrientationTypeSense OrientationType = iota
	OrientationTypeAntisense
)

func (code OrientationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *OrientationType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "sense":
		*code = OrientationTypeSense
	case "antisense":
		*code = OrientationTypeAntisense
	default:
		return fmt.Errorf("unknown OrientationType code `%s`", s)
	}
	return nil
}

func (code OrientationType) String() string {
	return code.Code()
}

func (code OrientationType) Code() string {
	switch code {
	case OrientationTypeSense:
		return "sense"
	case OrientationTypeAntisense:
		return "antisense"
	}
	return "<unknown>"
}

func (code OrientationType) Display() string {
	switch code {
	case OrientationTypeSense:
		return "Sense orientation of referenceSeq"
	case OrientationTypeAntisense:
		return "Antisense orientation of referenceSeq"
	}
	return "<unknown>"
}

func (code OrientationType) Definition() string {
	switch code {
	case OrientationTypeSense:
		return "Sense orientation of reference sequence."
	case OrientationTypeAntisense:
		return "Antisense orientation of reference sequence."
	}
	return "<unknown>"
}
