package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// A code specifying the types of information being requested.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/eligibilityresponse-purpose
type EligibilityResponsePurpose int

const (
	EligibilityResponsePurposeAuthRequirements EligibilityResponsePurpose = iota
	// The plan benefits and optionally benefits consumed  for the listed, or discovered if specified, converages are requested.
	EligibilityResponsePurposeBenefits
	// The insurer is requested to report on any coverages which they are aware of in addition to any specifed.
	EligibilityResponsePurposeDiscovery
	// A check that the specified coverages are in-force is requested.
	EligibilityResponsePurposeValidation
)

func (code EligibilityResponsePurpose) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *EligibilityResponsePurpose) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "auth-requirements":
		*code = EligibilityResponsePurposeAuthRequirements
	case "benefits":
		*code = EligibilityResponsePurposeBenefits
	case "discovery":
		*code = EligibilityResponsePurposeDiscovery
	case "validation":
		*code = EligibilityResponsePurposeValidation
	default:
		return fmt.Errorf("unknown EligibilityResponsePurpose code `%s`", s)
	}
	return nil
}

func (code EligibilityResponsePurpose) String() string {
	return code.Code()
}

func (code EligibilityResponsePurpose) Code() string {
	switch code {
	case EligibilityResponsePurposeAuthRequirements:
		return "auth-requirements"
	case EligibilityResponsePurposeBenefits:
		return "benefits"
	case EligibilityResponsePurposeDiscovery:
		return "discovery"
	case EligibilityResponsePurposeValidation:
		return "validation"
	}
	return "<unknown>"
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
