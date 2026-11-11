package fhires

// MedicationStatement Status Codes
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/medication-statement-status
type MedicationStatementStatusCodes string

const (
	// The action of recording the medication statement is finished.
	MedicationStatementStatusCodesRecorded MedicationStatementStatusCodes = "recorded"
	// Some of the actions that are implied by the medication usage may have occurred.  For example, the patient may have taken some of the medication.  Clinical decision support systems should take this status into account.
	MedicationStatementStatusCodesEnteredInError MedicationStatementStatusCodes = "entered-in-error"
	// The medication usage is draft or preliminary.
	MedicationStatementStatusCodesDraft MedicationStatementStatusCodes = "draft"
)

func (code MedicationStatementStatusCodes) String() string {
	return code.Code()
}

func (code MedicationStatementStatusCodes) Code() string {
	return string(code)
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
