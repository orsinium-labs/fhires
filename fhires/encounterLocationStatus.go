package fhires

// The status of the location.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/encounter-location-status
type EncounterLocationStatus string

const (
	// The patient is planned to be moved to this location at some point in the future.
	EncounterLocationStatusPlanned EncounterLocationStatus = "planned"
	// The patient is currently at this location, or was between the period specified.A system may update these records when the patient leaves the location to either reserved, or completed.
	EncounterLocationStatusActive EncounterLocationStatus = "active"
	// This location is held empty for this patient.
	EncounterLocationStatusReserved EncounterLocationStatus = "reserved"
	// The patient was at this location during the period specified.Not to be used when the patient is currently at the location.
	EncounterLocationStatusCompleted EncounterLocationStatus = "completed"
)

func (code EncounterLocationStatus) String() string {
	return code.Code()
}

func (code EncounterLocationStatus) Code() string {
	return string(code)
}

func (code EncounterLocationStatus) Display() string {
	switch code {
	case EncounterLocationStatusPlanned:
		return "Planned"
	case EncounterLocationStatusActive:
		return "Active"
	case EncounterLocationStatusReserved:
		return "Reserved"
	case EncounterLocationStatusCompleted:
		return "Completed"
	}
	return "<unknown>"
}

func (code EncounterLocationStatus) Definition() string {
	switch code {
	case EncounterLocationStatusPlanned:
		return "The patient is planned to be moved to this location at some point in the future."
	case EncounterLocationStatusActive:
		return "The patient is currently at this location, or was between the period specified.\r\rA system may update these records when the patient leaves the location to either reserved, or completed."
	case EncounterLocationStatusReserved:
		return "This location is held empty for this patient."
	case EncounterLocationStatusCompleted:
		return "The patient was at this location during the period specified.\r\rNot to be used when the patient is currently at the location."
	}
	return "<unknown>"
}
