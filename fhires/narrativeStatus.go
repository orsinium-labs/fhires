package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// The status of a resource narrative.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/narrative-status
type NarrativeStatus int

const (
	NarrativeStatusGenerated NarrativeStatus = iota
	// The contents of the narrative are entirely generated from the core elements in the content and some of the content is generated from extensions. The narrative SHALL reflect the impact of all modifier extensions.
	NarrativeStatusExtensions
	// The contents of the narrative may contain additional information not found in the structured data. Note that there is no computable way to determine what the extra information is, other than by human inspection.
	NarrativeStatusAdditional
	// The contents of the narrative are some equivalent of "No human-readable text provided in this case".
	NarrativeStatusEmpty
)

func (code NarrativeStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *NarrativeStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "generated":
		*code = NarrativeStatusGenerated
	case "extensions":
		*code = NarrativeStatusExtensions
	case "additional":
		*code = NarrativeStatusAdditional
	case "empty":
		*code = NarrativeStatusEmpty
	default:
		return fmt.Errorf("unknown NarrativeStatus code `%s`", s)
	}
	return nil
}

func (code NarrativeStatus) String() string {
	return code.Code()
}

func (code NarrativeStatus) Code() string {
	switch code {
	case NarrativeStatusGenerated:
		return "generated"
	case NarrativeStatusExtensions:
		return "extensions"
	case NarrativeStatusAdditional:
		return "additional"
	case NarrativeStatusEmpty:
		return "empty"
	}
	return "<unknown>"
}

func (code NarrativeStatus) Display() string {
	switch code {
	case NarrativeStatusGenerated:
		return "Generated"
	case NarrativeStatusExtensions:
		return "Extensions"
	case NarrativeStatusAdditional:
		return "Additional"
	case NarrativeStatusEmpty:
		return "Empty"
	}
	return "<unknown>"
}

func (code NarrativeStatus) Definition() string {
	switch code {
	case NarrativeStatusGenerated:
		return "The contents of the narrative are entirely generated from the core elements in the content."
	case NarrativeStatusExtensions:
		return "The contents of the narrative are entirely generated from the core elements in the content and some of the content is generated from extensions. The narrative SHALL reflect the impact of all modifier extensions."
	case NarrativeStatusAdditional:
		return "The contents of the narrative may contain additional information not found in the structured data. Note that there is no computable way to determine what the extra information is, other than by human inspection."
	case NarrativeStatusEmpty:
		return "The contents of the narrative are some equivalent of \"No human-readable text provided in this case\"."
	}
	return "<unknown>"
}
