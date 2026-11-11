package fhires

// Clinical assessment of the severity of a reaction event as a whole, potentially considering multiple different manifestations.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/reaction-event-severity
type AllergyIntoleranceSeverity string

const (
	// Causes mild physiological effects.
	AllergyIntoleranceSeverityMild AllergyIntoleranceSeverity = "mild"
	// Causes moderate physiological effects.
	AllergyIntoleranceSeverityModerate AllergyIntoleranceSeverity = "moderate"
	// Causes severe physiological effects.
	AllergyIntoleranceSeveritySevere AllergyIntoleranceSeverity = "severe"
)

func (code AllergyIntoleranceSeverity) String() string {
	return code.Code()
}

func (code AllergyIntoleranceSeverity) Code() string {
	return string(code)
}

func (code AllergyIntoleranceSeverity) Display() string {
	switch code {
	case AllergyIntoleranceSeverityMild:
		return "Mild"
	case AllergyIntoleranceSeverityModerate:
		return "Moderate"
	case AllergyIntoleranceSeveritySevere:
		return "Severe"
	}
	return "<unknown>"
}

func (code AllergyIntoleranceSeverity) Definition() string {
	switch code {
	case AllergyIntoleranceSeverityMild:
		return "Causes mild physiological effects."
	case AllergyIntoleranceSeverityModerate:
		return "Causes moderate physiological effects."
	case AllergyIntoleranceSeveritySevere:
		return "Causes severe physiological effects."
	}
	return "<unknown>"
}
