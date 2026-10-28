package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// MedicationKnowledge Status Codes
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/medicationknowledge-status
type MedicationKnowledgeStatusCodes int

const (
	MedicationKnowledgeStatusCodesActive MedicationKnowledgeStatusCodes = iota
	// The medication referred to by this MedicationKnowledge was entered in error within the drug database or inventory system.
	MedicationKnowledgeStatusCodesEnteredInError
	// The medication referred to by this MedicationKnowledge is not in active use within the drug database or inventory system.
	MedicationKnowledgeStatusCodesInactive
)

func (code MedicationKnowledgeStatusCodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *MedicationKnowledgeStatusCodes) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "active":
		*code = MedicationKnowledgeStatusCodesActive
	case "entered-in-error":
		*code = MedicationKnowledgeStatusCodesEnteredInError
	case "inactive":
		*code = MedicationKnowledgeStatusCodesInactive
	default:
		return fmt.Errorf("unknown MedicationKnowledgeStatusCodes code `%s`", s)
	}
	return nil
}

func (code MedicationKnowledgeStatusCodes) String() string {
	return code.Code()
}

func (code MedicationKnowledgeStatusCodes) Code() string {
	switch code {
	case MedicationKnowledgeStatusCodesActive:
		return "active"
	case MedicationKnowledgeStatusCodesEnteredInError:
		return "entered-in-error"
	case MedicationKnowledgeStatusCodesInactive:
		return "inactive"
	}
	return "<unknown>"
}

func (code MedicationKnowledgeStatusCodes) Display() string {
	switch code {
	case MedicationKnowledgeStatusCodesActive:
		return "Active"
	case MedicationKnowledgeStatusCodesEnteredInError:
		return "Entered in Error"
	case MedicationKnowledgeStatusCodesInactive:
		return "Inactive"
	}
	return "<unknown>"
}

func (code MedicationKnowledgeStatusCodes) Definition() string {
	switch code {
	case MedicationKnowledgeStatusCodesActive:
		return "The medication referred to by this MedicationKnowledge is in active use within the drug database or inventory system."
	case MedicationKnowledgeStatusCodesEnteredInError:
		return "The medication referred to by this MedicationKnowledge was entered in error within the drug database or inventory system."
	case MedicationKnowledgeStatusCodesInactive:
		return "The medication referred to by this MedicationKnowledge is not in active use within the drug database or inventory system."
	}
	return "<unknown>"
}
