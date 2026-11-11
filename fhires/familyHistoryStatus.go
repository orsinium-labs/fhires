package fhires

// A code that identifies the status of the family history record.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/history-status
type FamilyHistoryStatus string

const (
	// Some health information is known and captured, but not complete - see notes for details.
	FamilyHistoryStatusPartial FamilyHistoryStatus = "partial"
	// All available related health information is captured as of the date (and possibly time) when the family member history was taken.
	FamilyHistoryStatusCompleted FamilyHistoryStatus = "completed"
	// This instance should not have been part of this patient's medical record.
	FamilyHistoryStatusEnteredInError FamilyHistoryStatus = "entered-in-error"
	// Health information for this family member is unavailable/unknown.
	FamilyHistoryStatusHealthUnknown FamilyHistoryStatus = "health-unknown"
)

func (code FamilyHistoryStatus) String() string {
	return code.Code()
}

func (code FamilyHistoryStatus) Code() string {
	return string(code)
}

func (code FamilyHistoryStatus) Display() string {
	switch code {
	case FamilyHistoryStatusPartial:
		return "Partial"
	case FamilyHistoryStatusCompleted:
		return "Completed"
	case FamilyHistoryStatusEnteredInError:
		return "Entered in Error"
	case FamilyHistoryStatusHealthUnknown:
		return "Health Unknown"
	}
	return "<unknown>"
}

func (code FamilyHistoryStatus) Definition() string {
	switch code {
	case FamilyHistoryStatusPartial:
		return "Some health information is known and captured, but not complete - see notes for details."
	case FamilyHistoryStatusCompleted:
		return "All available related health information is captured as of the date (and possibly time) when the family member history was taken."
	case FamilyHistoryStatusEnteredInError:
		return "This instance should not have been part of this patient's medical record."
	case FamilyHistoryStatusHealthUnknown:
		return "Health information for this family member is unavailable/unknown."
	}
	return "<unknown>"
}
