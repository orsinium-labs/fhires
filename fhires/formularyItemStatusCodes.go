package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// FormularyItem Status Codes
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/formularyitem-status
type FormularyItemStatusCodes int

const (
	FormularyItemStatusCodesActive FormularyItemStatusCodes = iota
	// The service or product referred to by this FormularyItem was entered in error within the drug database or inventory system.
	FormularyItemStatusCodesEnteredInError
	// The service or product referred to by this FormularyItem is not in active use within the drug database or inventory system.
	FormularyItemStatusCodesInactive
)

func (code FormularyItemStatusCodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *FormularyItemStatusCodes) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "active":
		*code = FormularyItemStatusCodesActive
	case "entered-in-error":
		*code = FormularyItemStatusCodesEnteredInError
	case "inactive":
		*code = FormularyItemStatusCodesInactive
	default:
		return fmt.Errorf("unknown FormularyItemStatusCodes code `%s`", s)
	}
	return nil
}

func (code FormularyItemStatusCodes) String() string {
	return code.Code()
}

func (code FormularyItemStatusCodes) Code() string {
	switch code {
	case FormularyItemStatusCodesActive:
		return "active"
	case FormularyItemStatusCodesEnteredInError:
		return "entered-in-error"
	case FormularyItemStatusCodesInactive:
		return "inactive"
	}
	return "<unknown>"
}

func (code FormularyItemStatusCodes) Display() string {
	switch code {
	case FormularyItemStatusCodesActive:
		return "Active"
	case FormularyItemStatusCodesEnteredInError:
		return "Entered in Error"
	case FormularyItemStatusCodesInactive:
		return "Inactive"
	}
	return "<unknown>"
}

func (code FormularyItemStatusCodes) Definition() string {
	switch code {
	case FormularyItemStatusCodesActive:
		return "The service or product referred to by this FormularyItem is in active use within the drug database or inventory system."
	case FormularyItemStatusCodesEnteredInError:
		return "The service or product referred to by this FormularyItem was entered in error within the drug database or inventory system."
	case FormularyItemStatusCodesInactive:
		return "The service or product referred to by this FormularyItem is not in active use within the drug database or inventory system."
	}
	return "<unknown>"
}
