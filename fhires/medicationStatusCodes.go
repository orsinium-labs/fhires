package fhires

// Medication Status Codes
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/medication-status
type MedicationStatusCodes string

const (
	// The medication record is current and is appropriate for reference in new instances.
	MedicationStatusCodesActive MedicationStatusCodes = "active"
	// The medication record is not current and is not is appropriate for reference in new instances.
	MedicationStatusCodesInactive MedicationStatusCodes = "inactive"
	// The medication record was created erroneously and is not appropriated for reference in new instances.
	MedicationStatusCodesEnteredInError MedicationStatusCodes = "entered-in-error"
)

func (code MedicationStatusCodes) String() string {
	return code.Code()
}

func (code MedicationStatusCodes) Code() string {
	return string(code)
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
