package fhires

// A code specifying the types of information being requested.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/eligibilityrequest-purpose
type EligibilityRequestPurpose string

const (
	// The prior authorization requirements for the listed, or discovered if specified, converages for the categories of service and/or specifed biling codes are requested.
	EligibilityRequestPurposeAuthRequirements EligibilityRequestPurpose = "auth-requirements"
	// The plan benefits and optionally benefits consumed  for the listed, or discovered if specified, converages are requested.
	EligibilityRequestPurposeBenefits EligibilityRequestPurpose = "benefits"
	// The insurer is requested to report on any coverages which they are aware of in addition to any specifed.
	EligibilityRequestPurposeDiscovery EligibilityRequestPurpose = "discovery"
	// A check that the specified coverages are in-force is requested.
	EligibilityRequestPurposeValidation EligibilityRequestPurpose = "validation"
)

func (code EligibilityRequestPurpose) String() string {
	return code.Code()
}

func (code EligibilityRequestPurpose) Code() string {
	return string(code)
}

func (code EligibilityRequestPurpose) Display() string {
	switch code {
	case EligibilityRequestPurposeAuthRequirements:
		return "Coverage auth-requirements"
	case EligibilityRequestPurposeBenefits:
		return "Coverage benefits"
	case EligibilityRequestPurposeDiscovery:
		return "Coverage Discovery"
	case EligibilityRequestPurposeValidation:
		return "Coverage Validation"
	}
	return "<unknown>"
}

func (code EligibilityRequestPurpose) Definition() string {
	switch code {
	case EligibilityRequestPurposeAuthRequirements:
		return "The prior authorization requirements for the listed, or discovered if specified, converages for the categories of service and/or specifed biling codes are requested."
	case EligibilityRequestPurposeBenefits:
		return "The plan benefits and optionally benefits consumed  for the listed, or discovered if specified, converages are requested."
	case EligibilityRequestPurposeDiscovery:
		return "The insurer is requested to report on any coverages which they are aware of in addition to any specifed."
	case EligibilityRequestPurposeValidation:
		return "A check that the specified coverages are in-force is requested."
	}
	return "<unknown>"
}
