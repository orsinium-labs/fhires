package fhires

// A code specifying the types of information being requested.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/eligibilityresponse-purpose
type EligibilityResponsePurpose string

const (
	// The prior authorization requirements for the listed, or discovered if specified, converages for the categories of service and/or specifed biling codes are requested.
	EligibilityResponsePurposeAuthRequirements EligibilityResponsePurpose = "auth-requirements"
	// The plan benefits and optionally benefits consumed  for the listed, or discovered if specified, converages are requested.
	EligibilityResponsePurposeBenefits EligibilityResponsePurpose = "benefits"
	// The insurer is requested to report on any coverages which they are aware of in addition to any specifed.
	EligibilityResponsePurposeDiscovery EligibilityResponsePurpose = "discovery"
	// A check that the specified coverages are in-force is requested.
	EligibilityResponsePurposeValidation EligibilityResponsePurpose = "validation"
)

func (code EligibilityResponsePurpose) String() string {
	return code.Code()
}

func (code EligibilityResponsePurpose) Code() string {
	return string(code)
}

func (code EligibilityResponsePurpose) Display() string {
	switch code {
	case EligibilityResponsePurposeAuthRequirements:
		return "Coverage auth-requirements"
	case EligibilityResponsePurposeBenefits:
		return "Coverage benefits"
	case EligibilityResponsePurposeDiscovery:
		return "Coverage Discovery"
	case EligibilityResponsePurposeValidation:
		return "Coverage Validation"
	}
	return "<unknown>"
}

func (code EligibilityResponsePurpose) Definition() string {
	switch code {
	case EligibilityResponsePurposeAuthRequirements:
		return "The prior authorization requirements for the listed, or discovered if specified, converages for the categories of service and/or specifed biling codes are requested."
	case EligibilityResponsePurposeBenefits:
		return "The plan benefits and optionally benefits consumed  for the listed, or discovered if specified, converages are requested."
	case EligibilityResponsePurposeDiscovery:
		return "The insurer is requested to report on any coverages which they are aware of in addition to any specifed."
	case EligibilityResponsePurposeValidation:
		return "A check that the specified coverages are in-force is requested."
	}
	return "<unknown>"
}
