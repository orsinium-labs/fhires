package fhires

import (
	"encoding/json"
	"fmt"
	"strings"
)

// MedicationStatusCodes is documented here http://hl7.org/fhir/ValueSet/medication-status
type MedicationStatusCodes int

const (
	MedicationStatusCodesActive MedicationStatusCodes = iota
	MedicationStatusCodesInactive
	MedicationStatusCodesEnteredInError
)

func (code MedicationStatusCodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *MedicationStatusCodes) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "active":
		*code = MedicationStatusCodesActive
	case "inactive":
		*code = MedicationStatusCodesInactive
	case "entered-in-error":
		*code = MedicationStatusCodesEnteredInError
	default:
		return fmt.Errorf("unknown MedicationStatusCodes code `%s`", s)
	}
	return nil
}

func (code MedicationStatusCodes) String() string {
	return code.Code()
}

func (code MedicationStatusCodes) Code() string {
	switch code {
	case MedicationStatusCodesActive:
		return "active"
	case MedicationStatusCodesInactive:
		return "inactive"
	case MedicationStatusCodesEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}

func (code MedicationStatusCodes) Display() string {
	switch code {
	case MedicationStatusCodesActive:
		return "Active"
	case MedicationStatusCodesInactive:
		return "Inactive"
	case MedicationStatusCodesEnteredInError:
		return "Entered in Error"
	}
	return "<unknown>"
}

func (code MedicationStatusCodes) Definition() string {
	switch code {
	case MedicationStatusCodesActive:
		return "The medication record is current and is appropriate for reference in new instances."
	case MedicationStatusCodesInactive:
		return "The medication record is not current and is not is appropriate for reference in new instances."
	case MedicationStatusCodesEnteredInError:
		return "The medication record was created erroneously and is not appropriated for reference in new instances."
	}
	return "<unknown>"
}
