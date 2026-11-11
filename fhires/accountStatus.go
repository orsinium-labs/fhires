package fhires

// Indicates whether the account is available to be used.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/account-status
type AccountStatus string

const (
	// This account is active and may be used.
	AccountStatusActive AccountStatus = "active"
	// This account is inactive and should not be used to track financial information.
	AccountStatusInactive AccountStatus = "inactive"
	// This instance should not have been part of this patient's medical record.
	AccountStatusEnteredInError AccountStatus = "entered-in-error"
	// This account is on hold.
	AccountStatusOnHold AccountStatus = "on-hold"
	// The account status is unknown.
	AccountStatusUnknown AccountStatus = "unknown"
)

func (code AccountStatus) String() string {
	return code.Code()
}

func (code AccountStatus) Code() string {
	return string(code)
}

func (code AccountStatus) Display() string {
	switch code {
	case AccountStatusActive:
		return "Active"
	case AccountStatusInactive:
		return "Inactive"
	case AccountStatusEnteredInError:
		return "Entered in error"
	case AccountStatusOnHold:
		return "On Hold"
	case AccountStatusUnknown:
		return "Unknown"
	}
	return "<unknown>"
}

func (code AccountStatus) Definition() string {
	switch code {
	case AccountStatusActive:
		return "This account is active and may be used."
	case AccountStatusInactive:
		return "This account is inactive and should not be used to track financial information."
	case AccountStatusEnteredInError:
		return "This instance should not have been part of this patient's medical record."
	case AccountStatusOnHold:
		return "This account is on hold."
	case AccountStatusUnknown:
		return "The account status is unknown."
	}
	return "<unknown>"
}
