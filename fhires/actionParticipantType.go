package fhires

// The type of participant for the action.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/action-participant-type
type ActionParticipantType string

const (
	// The participant is a care team caring for the patient under evaluation.
	ActionParticipantTypeCareteam ActionParticipantType = "careteam"
	// The participant is a system or device used in the care of the patient.
	ActionParticipantTypeDevice ActionParticipantType = "device"
	// The participant is a group of participants involved in the care of the patient.
	ActionParticipantTypeGroup ActionParticipantType = "group"
	// The participant is an institution that can provide the given healthcare service used in the care of the patient.
	ActionParticipantTypeHealthcareservice ActionParticipantType = "healthcareservice"
	// The participant is a location involved in the care of the patient.
	ActionParticipantTypeLocation ActionParticipantType = "location"
	// The participant is an organization involved in the care of the patient.
	ActionParticipantTypeOrganization ActionParticipantType = "organization"
	// The participant is the patient under evaluation.
	ActionParticipantTypePatient ActionParticipantType = "patient"
	// The participant is a practitioner involved in the patient's care.
	ActionParticipantTypePractitioner ActionParticipantType = "practitioner"
	// The participant is a particular practitioner role involved in the patient's care.
	ActionParticipantTypePractitionerrole ActionParticipantType = "practitionerrole"
	// The participant is a person related to the patient.
	ActionParticipantTypeRelatedperson ActionParticipantType = "relatedperson"
)

func (code ActionParticipantType) String() string {
	return code.Code()
}

func (code ActionParticipantType) Code() string {
	return string(code)
}

func (code ActionParticipantType) Display() string {
	switch code {
	case ActionParticipantTypeCareteam:
		return "CareTeam"
	case ActionParticipantTypeDevice:
		return "Device"
	case ActionParticipantTypeGroup:
		return "Group"
	case ActionParticipantTypeHealthcareservice:
		return "HealthcareService"
	case ActionParticipantTypeLocation:
		return "Location"
	case ActionParticipantTypeOrganization:
		return "Organization"
	case ActionParticipantTypePatient:
		return "Patient"
	case ActionParticipantTypePractitioner:
		return "Practitioner"
	case ActionParticipantTypePractitionerrole:
		return "PractitionerRole"
	case ActionParticipantTypeRelatedperson:
		return "RelatedPerson"
	}
	return "<unknown>"
}

func (code ActionParticipantType) Definition() string {
	switch code {
	case ActionParticipantTypeCareteam:
		return "The participant is a care team caring for the patient under evaluation."
	case ActionParticipantTypeDevice:
		return "The participant is a system or device used in the care of the patient."
	case ActionParticipantTypeGroup:
		return "The participant is a group of participants involved in the care of the patient."
	case ActionParticipantTypeHealthcareservice:
		return "The participant is an institution that can provide the given healthcare service used in the care of the patient."
	case ActionParticipantTypeLocation:
		return "The participant is a location involved in the care of the patient."
	case ActionParticipantTypeOrganization:
		return "The participant is an organization involved in the care of the patient."
	case ActionParticipantTypePatient:
		return "The participant is the patient under evaluation."
	case ActionParticipantTypePractitioner:
		return "The participant is a practitioner involved in the patient's care."
	case ActionParticipantTypePractitionerrole:
		return "The participant is a particular practitioner role involved in the patient's care."
	case ActionParticipantTypeRelatedperson:
		return "The participant is a person related to the patient."
	}
	return "<unknown>"
}
