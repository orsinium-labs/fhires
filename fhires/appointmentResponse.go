package fhires

import "encoding/xml"

// A reply to an appointment request for a patient and/or practitioner(s), such as a confirmation or rejection.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/AppointmentResponse
type AppointmentResponse struct {
	DomainResource
	// This records identifiers associated with this appointment response concern that are defined by business processes and/ or used to refer to it when a direct URL reference to the resource itself is not appropriate.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Appointment that this response is replying to.
	Appointment Reference `bson:"appointment" json:"appointment" xml:"appointment"`
	// Indicates that the response is proposing a different time that was initially requested.
	//
	//  The new proposed time will be indicated in the start and end properties.
	ProposedNewTime *bool `bson:"proposedNewTime,omitempty" json:"proposedNewTime,omitempty" xml:"proposedNewTime,omitempty"`
	// Date/Time that the appointment is to take place, or requested new start time.
	//
	// Comment: This may be either the same as the appointment request to confirm the details of the appointment, or alternately a new time to request a re-negotiation of the start time.
	Start *string `bson:"start,omitempty" json:"start,omitempty" xml:"start,omitempty"`
	// This may be either the same as the appointment request to confirm the details of the appointment, or alternately a new time to request a re-negotiation of the end time.
	End *string `bson:"end,omitempty" json:"end,omitempty" xml:"end,omitempty"`
	// Role of participant in the appointment.
	//
	// Comment: The role of the participant can be used to declare what the actor will be doing in the scope of the referenced appointment.If the actor is not specified, then it is expected that the actor will be filled in at a later stage of planning.This value SHALL be the same as specified on the referenced Appointment so that they can be matched, and subsequently updated.
	ParticipantType []CodeableConcept `bson:"participantType,omitempty" json:"participantType,omitempty" xml:"participantType,omitempty"`
	// A Person, Location, HealthcareService, or Device that is participating in the appointment.
	Actor *Reference `bson:"actor,omitempty" json:"actor,omitempty" xml:"actor,omitempty"`
	// Participation status of the participant.
	//
	// When the status is declined or tentative if the start/end times are different to the appointment, then these times should be interpreted as a requested time change. When the status is accepted, the times can either be the time of the appointment (as a confirmation of the time) or can be empty.
	//
	// Comment: This element is labeled as a modifier because the status contains the code entered-in-error that marks the participant as not currently valid.
	ParticipantStatus string `bson:"participantStatus" json:"participantStatus" xml:"participantStatus"`
	// Additional comments about the appointment.
	//
	// Comment: This comment is particularly important when the responder is declining, tentatively accepting or requesting another time to indicate the reasons why.
	Comment *string `bson:"comment,omitempty" json:"comment,omitempty" xml:"comment,omitempty"`
	// Indicates that this AppointmentResponse applies to all occurrences in a recurring request.
	//
	// Comment: When a recurring appointment is requested, the participant may choose to respond to each individual occurrence, in which case AppointmentResponse.recurring should be false.  If the participant chooses to respond the same way for all occurrences, they may instead use a single AppointmentResponse with recurring set to true.  These may be combined as the participant chooses.  For example, they may accept all occurrences (recurring = true), but then send a decline for a specific occurrence (recurring=false).
	Recurring *bool `bson:"recurring,omitempty" json:"recurring,omitempty" xml:"recurring,omitempty"`
	// The original date within a recurring request.
	//
	// This could be used in place of the recurrenceId to be more direct (or where the template is provided through the simple list of dates in `Appointment.occurrenceDate`).
	OccurrenceDate *string `bson:"occurrenceDate,omitempty" json:"occurrenceDate,omitempty" xml:"occurrenceDate,omitempty"`
	// The recurrence ID (sequence number) of the specific appointment when responding to a recurring request.
	//
	// Comment: If the recurrence template was defined using the list of occurrenceDates then this property might not be used.
	RecurrenceID *int `bson:"recurrenceId,omitempty" json:"recurrenceId,omitempty" xml:"recurrenceId,omitempty"`
}

func (r AppointmentResponse) ResourceType() string {
	return "AppointmentResponse"
}

// MarshalJSON implements [json.Marshaler]
func (r AppointmentResponse) MarshalJSON() ([]byte, error) {
	type Wrapper AppointmentResponse
	return marshalJSON(r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Marshaler].
func (r AppointmentResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Wrapper AppointmentResponse
	return marshalXML(e, start, r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Unmarshaler].
func (r *AppointmentResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper AppointmentResponse
	w := Wrapper(*r)
	err := unmarshalXML(d, start, r.ResourceType(), &w)
	*r = AppointmentResponse(w)
	return err
}
