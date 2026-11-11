package fhires

// Estimate of the potential clinical harm, or seriousness, of a reaction to an identified substance.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/allergy-intolerance-criticality
type AllergyIntoleranceCriticality string

const (
	// Worst case result of a future exposure is not assessed to be life-threatening or having high potential for organ system failure.
	AllergyIntoleranceCriticalityLow AllergyIntoleranceCriticality = "low"
	// Worst case result of a future exposure is assessed to be life-threatening or having high potential for organ system failure.
	AllergyIntoleranceCriticalityHigh AllergyIntoleranceCriticality = "high"
	// Unable to assess the worst case result of a future exposure.
	AllergyIntoleranceCriticalityUnableToAssess AllergyIntoleranceCriticality = "unable-to-assess"
)

func (code AllergyIntoleranceCriticality) String() string {
	return code.Code()
}

func (code AllergyIntoleranceCriticality) Code() string {
	return string(code)
}

func (code AllergyIntoleranceCriticality) Display() string {
	switch code {
	case AllergyIntoleranceCriticalityLow:
		return "Low Risk"
	case AllergyIntoleranceCriticalityHigh:
		return "High Risk"
	case AllergyIntoleranceCriticalityUnableToAssess:
		return "Unable to Assess Risk"
	}
	return "<unknown>"
}

func (code AllergyIntoleranceCriticality) Definition() string {
	switch code {
	case AllergyIntoleranceCriticalityLow:
		return "Worst case result of a future exposure is not assessed to be life-threatening or having high potential for organ system failure."
	case AllergyIntoleranceCriticalityHigh:
		return "Worst case result of a future exposure is assessed to be life-threatening or having high potential for organ system failure."
	case AllergyIntoleranceCriticalityUnableToAssess:
		return "Unable to assess the worst case result of a future exposure."
	}
	return "<unknown>"
}
