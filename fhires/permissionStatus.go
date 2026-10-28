package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// Codes identifying the lifecycle stage of a product.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/permission-status
type PermissionStatus int

const (
	PermissionStatusActive PermissionStatus = iota
	// Permission was entered in error and is not active.
	PermissionStatusEnteredInError
	// Permission is being defined.
	PermissionStatusDraft
	// Permission not granted.
	PermissionStatusRejected
)

func (code PermissionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *PermissionStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "active":
		*code = PermissionStatusActive
	case "entered-in-error":
		*code = PermissionStatusEnteredInError
	case "draft":
		*code = PermissionStatusDraft
	case "rejected":
		*code = PermissionStatusRejected
	default:
		return fmt.Errorf("unknown PermissionStatus code `%s`", s)
	}
	return nil
}

func (code PermissionStatus) String() string {
	return code.Code()
}

func (code PermissionStatus) Code() string {
	switch code {
	case PermissionStatusActive:
		return "active"
	case PermissionStatusEnteredInError:
		return "entered-in-error"
	case PermissionStatusDraft:
		return "draft"
	case PermissionStatusRejected:
		return "rejected"
	}
	return "<unknown>"
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
