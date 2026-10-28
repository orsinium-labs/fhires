package fhires

import v2 "encoding/json/v2"

// A reply to an appointment request for a patient and/or practitioner(s), such as a confirmation or rejection.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/AppointmentResponse
type AppointmentResponse struct {
	// The logical id of the resource, as used in the URL for the resource.
	//
	// Once assigned, this value never changes.
	//
	// Comment: Within the context of the FHIR RESTful interactions, the resource has an id except for cases like the create and conditional update. Otherwise, the use of the resouce id depends on the given use case.
	ID *string `bson:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// The metadata about the resource.
	//
	// This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *Meta `bson:"meta,omitempty" json:"meta,omitempty" xml:"meta,omitempty"`
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content.
	//
	// Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	//
	// Comment: Asserting this rule set restricts the content to be only understood by a limited set of trading partners. This inherently limits the usefulness of the data in the long term. However, the existing health eco-system is highly fractured, and not yet ready to define, collect, and exchange data in a generally computable sense. Wherever possible, implementers and/or specification writers should avoid using this element. Often, when used, the URL is a reference to an implementation guide that defines these special rules as part of its narrative along with other profiles, value sets, etc.
	ImplicitRules *string `bson:"implicitRules,omitempty" json:"implicitRules,omitempty" xml:"implicitRules,omitempty"`
	// The base language in which the resource is written.
	//
	// Comment: Language is provided to support indexing and accessibility (typically, services such as text to speech use the language tag). The html language tag in the narrative applies  to the narrative. The language tag on the resource may be used to specify the language of other presentations generated from the data in the resource. Not all the content has to be in the base language. The Resource.language should not be assumed to apply to the narrative automatically. If a language is specified, it should it also be specified on the div element in the html (see rules in HTML5 for information about the relationship between xml:lang and the html lang attribute).
	Language *string `bson:"language,omitempty" json:"language,omitempty" xml:"language,omitempty"`
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human.
	//
	// The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	//
	// Comment: Contained resources do not have a narrative. Resources that are not contained SHOULD have a narrative. In some cases, a resource may only have text with little or no additional discrete data (as long as all minOccurs=1 elements are satisfied).  This may be necessary for data from legacy systems where information is captured as a "text blob" or where text is additionally entered raw or narrated and encoded information is added later.
	Text *Narrative `bson:"text,omitempty" json:"text,omitempty" xml:"text,omitempty"`
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, nor can they have their own independent transaction scope.
	//
	// This is allowed to be a Parameters resource if and only if it is referenced by a resource that provides context/meaning.
	//
	// Comment: This should never be done when the content can be identified properly, as once identification is lost, it is extremely difficult (and context dependent) to restore it again. Contained resources may have profiles and tags in their meta elements, but SHALL NOT have security labels.
	// May be used to represent additional information that is not part of the basic definition of the resource.
	//
	// To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	//
	// Comment: There can be no stigma associated with the use of extensions by any application, project, or standard - regardless of the institution or jurisdiction that uses or defines the extensions.  The use of extensions is what allows the FHIR specification to retain a core level of simplicity for everyone.
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty" xml:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants.
	//
	// Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	//
	// Comment: There can be no stigma associated with the use of extensions by any application, project, or standard - regardless of the institution or jurisdiction that uses or defines the extensions.  The use of extensions is what allows the FHIR specification to retain a core level of simplicity for everyone.
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty" xml:"modifierExtension,omitempty"`
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

// MarshalJSON implements json.Marshaler
func (r AppointmentResponse) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		AppointmentResponse
		ResourceType string `json:"resourceType"`
	}{
		AppointmentResponse: r,
		ResourceType:        r.ResourceType(),
	})
}
