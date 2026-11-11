package fhires

// A code to indicate if the substance is actively used.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/substance-status
type FHIRSubstanceStatus string

const (
	// The substance is considered for use or reference.
	FHIRSubstanceStatusActive FHIRSubstanceStatus = "active"
	// The substance is considered for reference, but not for use.
	FHIRSubstanceStatusInactive FHIRSubstanceStatus = "inactive"
	// The substance was entered in error.
	FHIRSubstanceStatusEnteredInError FHIRSubstanceStatus = "entered-in-error"
)

func (code FHIRSubstanceStatus) String() string {
	return code.Code()
}

func (code FHIRSubstanceStatus) Code() string {
	return string(code)
}

func (code FHIRSubstanceStatus) Display() string {
	switch code {
	case FHIRSubstanceStatusActive:
		return "Active"
	case FHIRSubstanceStatusInactive:
		return "Inactive"
	case FHIRSubstanceStatusEnteredInError:
		return "Entered in Error"
	}
	return "<unknown>"
}

func (code FHIRSubstanceStatus) Definition() string {
	switch code {
	case FHIRSubstanceStatusActive:
		return "The substance is considered for use or reference."
	case FHIRSubstanceStatusInactive:
		return "The substance is considered for reference, but not for use."
	case FHIRSubstanceStatusEnteredInError:
		return "The substance was entered in error."
	}
	return "<unknown>"
}
