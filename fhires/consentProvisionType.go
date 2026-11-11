package fhires

// How a rule statement is applied, such as adding additional consent or removing consent.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/consent-provision-type
type ConsentProvisionType string

const (
	// Consent is denied for actions meeting these rules.
	ConsentProvisionTypeDeny ConsentProvisionType = "deny"
	// Consent is provided for actions meeting these rules.
	ConsentProvisionTypePermit ConsentProvisionType = "permit"
)

func (code ConsentProvisionType) String() string {
	return code.Code()
}

func (code ConsentProvisionType) Code() string {
	return string(code)
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
