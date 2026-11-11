package fhires

// Codes providing the combined status of a specimen.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/specimen-combined
type SpecimenCombined string

const (
	// The specimen is in a group.
	SpecimenCombinedGrouped SpecimenCombined = "grouped"
	// The specimen is pooled.
	SpecimenCombinedPooled SpecimenCombined = "pooled"
)

func (code SpecimenCombined) String() string {
	return code.Code()
}

func (code SpecimenCombined) Code() string {
	return string(code)
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
