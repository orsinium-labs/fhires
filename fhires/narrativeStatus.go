package fhires

// The status of a resource narrative.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/narrative-status
type NarrativeStatus string

const (
	// The contents of the narrative are entirely generated from the core elements in the content.
	NarrativeStatusGenerated NarrativeStatus = "generated"
	// The contents of the narrative are entirely generated from the core elements in the content and some of the content is generated from extensions. The narrative SHALL reflect the impact of all modifier extensions.
	NarrativeStatusExtensions NarrativeStatus = "extensions"
	// The contents of the narrative may contain additional information not found in the structured data. Note that there is no computable way to determine what the extra information is, other than by human inspection.
	NarrativeStatusAdditional NarrativeStatus = "additional"
	// The contents of the narrative are some equivalent of "No human-readable text provided in this case".
	NarrativeStatusEmpty NarrativeStatus = "empty"
)

func (code NarrativeStatus) String() string {
	return code.Code()
}

func (code NarrativeStatus) Code() string {
	return string(code)
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
