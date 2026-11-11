package fhires

// The gender of a person used for administrative purposes.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/administrative-gender
type AdministrativeGender string

const (
	// Male.
	AdministrativeGenderMale AdministrativeGender = "male"
	// Female.
	AdministrativeGenderFemale AdministrativeGender = "female"
	// Other.
	AdministrativeGenderOther AdministrativeGender = "other"
	// Unknown.
	AdministrativeGenderUnknown AdministrativeGender = "unknown"
)

func (code AdministrativeGender) String() string {
	return code.Code()
}

func (code AdministrativeGender) Code() string {
	return string(code)
}

func (code AdministrativeGender) Display() string {
	switch code {
	case AdministrativeGenderMale:
		return "Male"
	case AdministrativeGenderFemale:
		return "Female"
	case AdministrativeGenderOther:
		return "Other"
	case AdministrativeGenderUnknown:
		return "Unknown"
	}
	return "<unknown>"
}

func (code AdministrativeGender) Definition() string {
	switch code {
	case AdministrativeGenderMale:
		return "Male."
	case AdministrativeGenderFemale:
		return "Female."
	case AdministrativeGenderOther:
		return "Other."
	case AdministrativeGenderUnknown:
		return "Unknown."
	}
	return "<unknown>"
}
