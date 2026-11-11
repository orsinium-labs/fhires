package fhires

// The Participation status of an appointment.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/participationstatus
type ParticipationStatus string

const (
	// The participant has accepted the appointment.
	ParticipationStatusAccepted ParticipationStatus = "accepted"
	// The participant has declined the appointment and will not participate in the appointment.
	ParticipationStatusDeclined ParticipationStatus = "declined"
	// The participant has  tentatively accepted the appointment. This could be automatically created by a system and requires further processing before it can be accepted. There is no commitment that attendance will occur.
	ParticipationStatusTentative ParticipationStatus = "tentative"
	// The participant needs to indicate if they accept the appointment by changing this status to one of the other statuses.
	ParticipationStatusNeedsAction ParticipationStatus = "needs-action"
)

func (code ParticipationStatus) String() string {
	return code.Code()
}

func (code ParticipationStatus) Code() string {
	return string(code)
}

func (code ParticipationStatus) Display() string {
	switch code {
	case ParticipationStatusAccepted:
		return "Accepted"
	case ParticipationStatusDeclined:
		return "Declined"
	case ParticipationStatusTentative:
		return "Tentative"
	case ParticipationStatusNeedsAction:
		return "Needs Action"
	}
	return "<unknown>"
}

func (code ParticipationStatus) Definition() string {
	switch code {
	case ParticipationStatusAccepted:
		return "The participant has accepted the appointment."
	case ParticipationStatusDeclined:
		return "The participant has declined the appointment and will not participate in the appointment."
	case ParticipationStatusTentative:
		return "The participant has  tentatively accepted the appointment. This could be automatically created by a system and requires further processing before it can be accepted. There is no commitment that attendance will occur."
	case ParticipationStatusNeedsAction:
		return "The participant needs to indicate if they accept the appointment by changing this status to one of the other statuses."
	}
	return "<unknown>"
}
