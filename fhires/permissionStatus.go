package fhires

// Codes identifying the lifecycle stage of a product.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/permission-status
type PermissionStatus string

const (
	// Permission is given.
	PermissionStatusActive PermissionStatus = "active"
	// Permission was entered in error and is not active.
	PermissionStatusEnteredInError PermissionStatus = "entered-in-error"
	// Permission is being defined.
	PermissionStatusDraft PermissionStatus = "draft"
	// Permission not granted.
	PermissionStatusRejected PermissionStatus = "rejected"
)

func (code PermissionStatus) String() string {
	return code.Code()
}

func (code PermissionStatus) Code() string {
	return string(code)
}

func (code PermissionStatus) Display() string {
	switch code {
	case PermissionStatusActive:
		return "Active"
	case PermissionStatusEnteredInError:
		return "Entered in Error"
	case PermissionStatusDraft:
		return "Draft"
	case PermissionStatusRejected:
		return "Rejected"
	}
	return "<unknown>"
}

func (code PermissionStatus) Definition() string {
	switch code {
	case PermissionStatusActive:
		return "Permission is given."
	case PermissionStatusEnteredInError:
		return "Permission was entered in error and is not active."
	case PermissionStatusDraft:
		return "Permission is being defined."
	case PermissionStatusRejected:
		return "Permission not granted."
	}
	return "<unknown>"
}
