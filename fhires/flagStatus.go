package fhires

// Indicates whether this flag is active and needs to be displayed to a user, or whether it is no longer needed or was entered in error.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/flag-status
type FlagStatus string

const (
	// A current flag that should be displayed to a user. A system may use the category to determine which user roles should view the flag.
	FlagStatusActive FlagStatus = "active"
	// The flag no longer needs to be displayed.
	FlagStatusInactive FlagStatus = "inactive"
	// The flag was added in error and should no longer be displayed.
	FlagStatusEnteredInError FlagStatus = "entered-in-error"
)

func (code FlagStatus) String() string {
	return code.Code()
}

func (code FlagStatus) Code() string {
	return string(code)
}

func (code FlagStatus) Display() string {
	switch code {
	case FlagStatusActive:
		return "Active"
	case FlagStatusInactive:
		return "Inactive"
	case FlagStatusEnteredInError:
		return "Entered in Error"
	}
	return "<unknown>"
}

func (code FlagStatus) Definition() string {
	switch code {
	case FlagStatusActive:
		return "A current flag that should be displayed to a user. A system may use the category to determine which user roles should view the flag."
	case FlagStatusInactive:
		return "The flag no longer needs to be displayed."
	case FlagStatusEnteredInError:
		return "The flag was added in error and should no longer be displayed."
	}
	return "<unknown>"
}
