package fhires

// The level of confidence that this link represents the same actual person, based on NIST Authentication Levels.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/identity-assuranceLevel
type IdentityAssuranceLevel string

const (
	// Little or no confidence in the asserted identity's accuracy.
	IdentityAssuranceLevelLevel1 IdentityAssuranceLevel = "level1"
	// Some confidence in the asserted identity's accuracy.
	IdentityAssuranceLevelLevel2 IdentityAssuranceLevel = "level2"
	// High confidence in the asserted identity's accuracy.
	IdentityAssuranceLevelLevel3 IdentityAssuranceLevel = "level3"
	// Very high confidence in the asserted identity's accuracy.
	IdentityAssuranceLevelLevel4 IdentityAssuranceLevel = "level4"
)

func (code IdentityAssuranceLevel) String() string {
	return code.Code()
}

func (code IdentityAssuranceLevel) Code() string {
	return string(code)
}

func (code IdentityAssuranceLevel) Display() string {
	switch code {
	case IdentityAssuranceLevelLevel1:
		return "Level 1"
	case IdentityAssuranceLevelLevel2:
		return "Level 2"
	case IdentityAssuranceLevelLevel3:
		return "Level 3"
	case IdentityAssuranceLevelLevel4:
		return "Level 4"
	}
	return "<unknown>"
}

func (code IdentityAssuranceLevel) Definition() string {
	switch code {
	case IdentityAssuranceLevelLevel1:
		return "Little or no confidence in the asserted identity's accuracy."
	case IdentityAssuranceLevelLevel2:
		return "Some confidence in the asserted identity's accuracy."
	case IdentityAssuranceLevelLevel3:
		return "High confidence in the asserted identity's accuracy."
	case IdentityAssuranceLevelLevel4:
		return "Very high confidence in the asserted identity's accuracy."
	}
	return "<unknown>"
}
