package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// How a rule statement is applied, such as adding additional consent or removing consent.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/consent-provision-type
type ConsentProvisionType int

const (
	ConsentProvisionTypeDeny ConsentProvisionType = iota
	// Consent is provided for actions meeting these rules.
	ConsentProvisionTypePermit
)

func (code ConsentProvisionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *ConsentProvisionType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "deny":
		*code = ConsentProvisionTypeDeny
	case "permit":
		*code = ConsentProvisionTypePermit
	default:
		return fmt.Errorf("unknown ConsentProvisionType code `%s`", s)
	}
	return nil
}

func (code ConsentProvisionType) String() string {
	return code.Code()
}

func (code ConsentProvisionType) Code() string {
	switch code {
	case ConsentProvisionTypeDeny:
		return "deny"
	case ConsentProvisionTypePermit:
		return "permit"
	}
	return "<unknown>"
}

func (code ConsentProvisionType) Display() string {
	switch code {
	case ConsentProvisionTypeDeny:
		return "Deny"
	case ConsentProvisionTypePermit:
		return "Permit"
	}
	return "<unknown>"
}

func (code ConsentProvisionType) Definition() string {
	switch code {
	case ConsentProvisionTypeDeny:
		return "Consent is denied for actions meeting these rules."
	case ConsentProvisionTypePermit:
		return "Consent is provided for actions meeting these rules."
	}
	return "<unknown>"
}
