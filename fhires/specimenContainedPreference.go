package fhires

// Degree of preference of a type of conditioned specimen.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/specimen-contained-preference
type SpecimenContainedPreference string

const (
	// This type of contained specimen is preferred to collect this kind of specimen.
	SpecimenContainedPreferencePreferred SpecimenContainedPreference = "preferred"
	// This type of conditioned specimen is an alternate.
	SpecimenContainedPreferenceAlternate SpecimenContainedPreference = "alternate"
)

func (code SpecimenContainedPreference) String() string {
	return code.Code()
}

func (code SpecimenContainedPreference) Code() string {
	return string(code)
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
