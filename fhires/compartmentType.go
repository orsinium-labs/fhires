package fhires

// Which type a compartment definition describes.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/compartment-type
type CompartmentType string

const (
	// The compartment definition is for the patient compartment.
	CompartmentTypePatient CompartmentType = "Patient"
	// The compartment definition is for the encounter compartment.
	CompartmentTypeEncounter CompartmentType = "Encounter"
	// The compartment definition is for the related-person compartment.
	CompartmentTypeRelatedPerson CompartmentType = "RelatedPerson"
	// The compartment definition is for the practitioner compartment.
	CompartmentTypePractitioner CompartmentType = "Practitioner"
	// The compartment definition is for the device compartment.
	CompartmentTypeDevice CompartmentType = "Device"
	// The compartment definition is for the episodeofcare compartment.
	CompartmentTypeEpisodeOfCare CompartmentType = "EpisodeOfCare"
)

func (code CompartmentType) String() string {
	return code.Code()
}

func (code CompartmentType) Code() string {
	return string(code)
}

func (code CompartmentType) Display() string {
	switch code {
	case CompartmentTypePatient:
		return "Patient"
	case CompartmentTypeEncounter:
		return "Encounter"
	case CompartmentTypeRelatedPerson:
		return "RelatedPerson"
	case CompartmentTypePractitioner:
		return "Practitioner"
	case CompartmentTypeDevice:
		return "Device"
	case CompartmentTypeEpisodeOfCare:
		return "EpisodeOfCare"
	}
	return "<unknown>"
}

func (code CompartmentType) Definition() string {
	switch code {
	case CompartmentTypePatient:
		return "The compartment definition is for the patient compartment."
	case CompartmentTypeEncounter:
		return "The compartment definition is for the encounter compartment."
	case CompartmentTypeRelatedPerson:
		return "The compartment definition is for the related-person compartment."
	case CompartmentTypePractitioner:
		return "The compartment definition is for the practitioner compartment."
	case CompartmentTypeDevice:
		return "The compartment definition is for the device compartment."
	case CompartmentTypeEpisodeOfCare:
		return "The compartment definition is for the episodeofcare compartment."
	}
	return "<unknown>"
}
