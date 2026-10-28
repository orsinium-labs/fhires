package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// MedicationStatement Status Codes
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/medication-statement-status
type MedicationStatementStatusCodes int

const (
	MedicationStatementStatusCodesRecorded MedicationStatementStatusCodes = iota
	// Some of the actions that are implied by the medication usage may have occurred.  For example, the patient may have taken some of the medication.  Clinical decision support systems should take this status into account.
	MedicationStatementStatusCodesEnteredInError
	// The medication usage is draft or preliminary.
	MedicationStatementStatusCodesDraft
)

func (code MedicationStatementStatusCodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *MedicationStatementStatusCodes) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "recorded":
		*code = MedicationStatementStatusCodesRecorded
	case "entered-in-error":
		*code = MedicationStatementStatusCodesEnteredInError
	case "draft":
		*code = MedicationStatementStatusCodesDraft
	default:
		return fmt.Errorf("unknown MedicationStatementStatusCodes code `%s`", s)
	}
	return nil
}

func (code MedicationStatementStatusCodes) String() string {
	return code.Code()
}

func (code MedicationStatementStatusCodes) Code() string {
	switch code {
	case MedicationStatementStatusCodesRecorded:
		return "recorded"
	case MedicationStatementStatusCodesEnteredInError:
		return "entered-in-error"
	case MedicationStatementStatusCodesDraft:
		return "draft"
	}
	return "<unknown>"
}

func (code MedicationStatementStatusCodes) Display() string {
	switch code {
	case MedicationStatementStatusCodesRecorded:
		return "Recorded"
	case MedicationStatementStatusCodesEnteredInError:
		return "Entered in Error"
	case MedicationStatementStatusCodesDraft:
		return "Draft"
	}
	return "<unknown>"
}

func (code MedicationStatementStatusCodes) Definition() string {
	switch code {
	case MedicationStatementStatusCodesRecorded:
		return "The action of recording the medication statement is finished."
	case MedicationStatementStatusCodesEnteredInError:
		return "Some of the actions that are implied by the medication usage may have occurred.  For example, the patient may have taken some of the medication.  Clinical decision support systems should take this status into account."
	case MedicationStatementStatusCodesDraft:
		return "The medication usage is draft or preliminary."
	}
	return "<unknown>"
}
