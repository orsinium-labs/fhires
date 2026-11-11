package fhires

// A code specifying the state of the resource instance.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/explanationofbenefit-status
type ExplanationOfBenefitStatus string

const (
	// The resource instance is currently in-force.
	ExplanationOfBenefitStatusActive ExplanationOfBenefitStatus = "active"
	// The resource instance is withdrawn, rescinded or reversed.
	ExplanationOfBenefitStatusCancelled ExplanationOfBenefitStatus = "cancelled"
	// A new resource instance the contents of which is not complete.
	ExplanationOfBenefitStatusDraft ExplanationOfBenefitStatus = "draft"
	// The resource instance was entered in error.
	ExplanationOfBenefitStatusEnteredInError ExplanationOfBenefitStatus = "entered-in-error"
)

func (code ExplanationOfBenefitStatus) String() string {
	return code.Code()
}

func (code ExplanationOfBenefitStatus) Code() string {
	return string(code)
}

func (code ExplanationOfBenefitStatus) Display() string {
	switch code {
	case ExplanationOfBenefitStatusActive:
		return "Active"
	case ExplanationOfBenefitStatusCancelled:
		return "Cancelled"
	case ExplanationOfBenefitStatusDraft:
		return "Draft"
	case ExplanationOfBenefitStatusEnteredInError:
		return "Entered In Error"
	}
	return "<unknown>"
}

func (code ExplanationOfBenefitStatus) Definition() string {
	switch code {
	case ExplanationOfBenefitStatusActive:
		return "The resource instance is currently in-force."
	case ExplanationOfBenefitStatusCancelled:
		return "The resource instance is withdrawn, rescinded or reversed."
	case ExplanationOfBenefitStatusDraft:
		return "A new resource instance the contents of which is not complete."
	case ExplanationOfBenefitStatusEnteredInError:
		return "The resource instance was entered in error."
	}
	return "<unknown>"
}
