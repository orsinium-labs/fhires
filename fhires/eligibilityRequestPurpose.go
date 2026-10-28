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
// http://hl7.org/fhir/ValueSet/eligibilityrequest-purpose
type EligibilityRequestPurpose int

const (
	EligibilityRequestPurposeAuthRequirements EligibilityRequestPurpose = iota
	// The plan benefits and optionally benefits consumed  for the listed, or discovered if specified, converages are requested.
	EligibilityRequestPurposeBenefits
	// The insurer is requested to report on any coverages which they are aware of in addition to any specifed.
	EligibilityRequestPurposeDiscovery
	// A check that the specified coverages are in-force is requested.
	EligibilityRequestPurposeValidation
)

func (code EligibilityRequestPurpose) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *EligibilityRequestPurpose) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "auth-requirements":
		*code = EligibilityRequestPurposeAuthRequirements
	case "benefits":
		*code = EligibilityRequestPurposeBenefits
	case "discovery":
		*code = EligibilityRequestPurposeDiscovery
	case "validation":
		*code = EligibilityRequestPurposeValidation
	default:
		return fmt.Errorf("unknown EligibilityRequestPurpose code `%s`", s)
	}
	return nil
}

func (code EligibilityRequestPurpose) String() string {
	return code.Code()
}

func (code EligibilityRequestPurpose) Code() string {
	switch code {
	case EligibilityRequestPurposeAuthRequirements:
		return "auth-requirements"
	case EligibilityRequestPurposeBenefits:
		return "benefits"
	case EligibilityRequestPurposeDiscovery:
		return "discovery"
	case EligibilityRequestPurposeValidation:
		return "validation"
	}
	return "<unknown>"
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
