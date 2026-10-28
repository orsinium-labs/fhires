package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// How a search parameter relates to the set of elements returned by evaluating its expression query.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/search-processingmode
type SearchProcessingModeType int

const (
	SearchProcessingModeTypeNormal SearchProcessingModeType = iota
	// The search parameter is derived by a phonetic transform from the selected nodes.
	SearchProcessingModeTypePhonetic
	// The interpretation of the xpath statement is unknown (and can't be automated).
	SearchProcessingModeTypeOther
)

func (code SearchProcessingModeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *SearchProcessingModeType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "normal":
		*code = SearchProcessingModeTypeNormal
	case "phonetic":
		*code = SearchProcessingModeTypePhonetic
	case "other":
		*code = SearchProcessingModeTypeOther
	default:
		return fmt.Errorf("unknown SearchProcessingModeType code `%s`", s)
	}
	return nil
}

func (code SearchProcessingModeType) String() string {
	return code.Code()
}

func (code SearchProcessingModeType) Code() string {
	switch code {
	case SearchProcessingModeTypeNormal:
		return "normal"
	case SearchProcessingModeTypePhonetic:
		return "phonetic"
	case SearchProcessingModeTypeOther:
		return "other"
	}
	return "<unknown>"
}

func (code SearchProcessingModeType) Display() string {
	switch code {
	case SearchProcessingModeTypeNormal:
		return "Normal"
	case SearchProcessingModeTypePhonetic:
		return "Phonetic"
	case SearchProcessingModeTypeOther:
		return "Other"
	}
	return "<unknown>"
}

func (code SearchProcessingModeType) Definition() string {
	switch code {
	case SearchProcessingModeTypeNormal:
		return "The search parameter is derived directly from the selected nodes based on the type definitions."
	case SearchProcessingModeTypePhonetic:
		return "The search parameter is derived by a phonetic transform from the selected nodes."
	case SearchProcessingModeTypeOther:
		return "The interpretation of the xpath statement is unknown (and can't be automated)."
	}
	return "<unknown>"
}
