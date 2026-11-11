package fhires

// How a search parameter relates to the set of elements returned by evaluating its expression query.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/search-processingmode
type SearchProcessingModeType string

const (
	// The search parameter is derived directly from the selected nodes based on the type definitions.
	SearchProcessingModeTypeNormal SearchProcessingModeType = "normal"
	// The search parameter is derived by a phonetic transform from the selected nodes.
	SearchProcessingModeTypePhonetic SearchProcessingModeType = "phonetic"
	// The interpretation of the xpath statement is unknown (and can't be automated).
	SearchProcessingModeTypeOther SearchProcessingModeType = "other"
)

func (code SearchProcessingModeType) String() string {
	return code.Code()
}

func (code SearchProcessingModeType) Code() string {
	return string(code)
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
