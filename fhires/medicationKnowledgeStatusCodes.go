package fhires

// MedicationKnowledge Status Codes
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/medicationknowledge-status
type MedicationKnowledgeStatusCodes string

const (
	// The medication referred to by this MedicationKnowledge is in active use within the drug database or inventory system.
	MedicationKnowledgeStatusCodesActive MedicationKnowledgeStatusCodes = "active"
	// The medication referred to by this MedicationKnowledge was entered in error within the drug database or inventory system.
	MedicationKnowledgeStatusCodesEnteredInError MedicationKnowledgeStatusCodes = "entered-in-error"
	// The medication referred to by this MedicationKnowledge is not in active use within the drug database or inventory system.
	MedicationKnowledgeStatusCodesInactive MedicationKnowledgeStatusCodes = "inactive"
)

func (code MedicationKnowledgeStatusCodes) String() string {
	return code.Code()
}

func (code MedicationKnowledgeStatusCodes) Code() string {
	return string(code)
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
