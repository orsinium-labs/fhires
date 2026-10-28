package fhires

import "encoding/json/v2"

// A booking of a healthcare event among patient(s), practitioner(s), related person(s) and/or device(s) for a specific date/time.
//
// This may result in one or more Encounter(s).
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Appointment
type Appointment struct {
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
	// This records identifiers associated with this appointment concern that are defined by business processes and/or used to refer to it when a direct URL reference to the resource itself is not appropriate (e.g.
	//
	// in CDA documents, or in written / printed documentation).
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The overall status of the Appointment.
	//
	// Each of the participants has their own participation status which indicates their involvement in the process, however this status indicates the shared status.
	//
	// Comment: If the Appointment's status is "cancelled" then all participants are expected to have their calendars released for the appointment period, and as such any Slots that were marked as BUSY can be re-set to FREE.
	//
	// This element is labeled as a modifier because the status contains the code entered-in-error that mark the Appointment as not currently valid.
	Status AppointmentStatus `bson:"status" json:"status" xml:"status"`
	// The coded reason for the appointment being cancelled.
	//
	// This is often used in reporting/billing/futher processing to determine if further actions are required, or specific fees apply.
	CancellationReason *CodeableConcept `bson:"cancellationReason,omitempty" json:"cancellationReason,omitempty" xml:"cancellationReason,omitempty"`
	// Concepts representing classification of patient encounter such as ambulatory (outpatient), inpatient, emergency, home health or others due to local variations.
	Class []CodeableConcept `bson:"class,omitempty" json:"class,omitempty" xml:"class,omitempty"`
	// A broad categorization of the service that is to be performed during this appointment.
	ServiceCategory []CodeableConcept `bson:"serviceCategory,omitempty" json:"serviceCategory,omitempty" xml:"serviceCategory,omitempty"`
	// The specific service that is to be performed during this appointment.
	//
	// Comment: For a provider to provider appointment the code "FOLLOWUP" may be appropriate, as this is expected to be discussing some patient that was seen in the past.
	ServiceType []CodeableReference `bson:"serviceType,omitempty" json:"serviceType,omitempty" xml:"serviceType,omitempty"`
	// The specialty of a practitioner that would be required to perform the service requested in this appointment.
	Specialty []CodeableConcept `bson:"specialty,omitempty" json:"specialty,omitempty" xml:"specialty,omitempty"`
	// The style of appointment or patient that has been booked in the slot (not service type).
	AppointmentType *CodeableConcept `bson:"appointmentType,omitempty" json:"appointmentType,omitempty" xml:"appointmentType,omitempty"`
	// The reason that this appointment is being scheduled.
	//
	// This is more clinical than administrative. This can be coded, or as specified using information from another resource. When the patient arrives and the encounter begins it may be used as the admission diagnosis. The indication will typically be a Condition (with other resources referenced in the evidence.detail), or a Procedure.
	Reason []CodeableReference `bson:"reason,omitempty" json:"reason,omitempty" xml:"reason,omitempty"`
	// The priority of the appointment.
	//
	// Can be used to make informed decisions if needing to re-prioritize appointments. (The iCal Standard specifies 0 as undefined, 1 as highest, 9 as lowest priority).
	//
	// Comment: Seeking implementer feedback on this property and how interoperable it is.Using an extension to record a CodeableConcept for named values may be tested at a future connectathon.
	Priority *CodeableConcept `bson:"priority,omitempty" json:"priority,omitempty" xml:"priority,omitempty"`
	// The brief description of the appointment as would be shown on a subject line in a meeting request, or appointment list.
	//
	// Detailed or expanded information should be put in the note field.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Appointment replaced by this Appointment in cases where there is a cancellation, the details of the cancellation can be found in the cancellationReason property (on the referenced resource).
	Replaces []Reference `bson:"replaces,omitempty" json:"replaces,omitempty" xml:"replaces,omitempty"`
	// Connection details of a virtual service (e.g. conference call).
	//
	// Comment: There are two types of virtual meetings that often exist:* a persistent, virtual meeting room that can only be used for a single purpose at a time, * and a dynamic virtual meeting room that is generated on demand for a specific purpose.Implementers may consider using Location.virtualService for persistent meeting rooms.If each participant would have a different meeting link, an extension using the VirtualServiceContactDetail  can be applied to the Appointment.participant BackboneElement.
	VirtualService []VirtualServiceDetail `bson:"virtualService,omitempty" json:"virtualService,omitempty" xml:"virtualService,omitempty"`
	// Additional information to support the appointment provided when making the appointment.
	SupportingInformation []Reference `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty" xml:"supportingInformation,omitempty"`
	// The previous appointment in a series of related appointments.
	//
	// Comment: This property is intended for use when representing a series of related appointments.  For example, in a nuclear medicine procedure, where there is an appointment for the injection of the isotopes, and then a subsequent appointment for the scan, the scan appointment would refer to the injection appointment via Appointment.previousAppointment.  For representing recurring appointments, see the guidance on recurring vs. series appointments.
	PreviousAppointment *Reference `bson:"previousAppointment,omitempty" json:"previousAppointment,omitempty" xml:"previousAppointment,omitempty"`
	// The originating appointment in a recurring set of related appointments.
	//
	// Comment: This property is intended for use when representing a recurring set of related appointments.For example, a patient undergoing physical therapy may have a recurring appointment every Tuesday and Thursday.  Each occurrence of the set will refer to the originating appointment, which contains the recurring template information.  For representing appointment series, see the guidance on recurring vs. series appointments.
	OriginatingAppointment *Reference `bson:"originatingAppointment,omitempty" json:"originatingAppointment,omitempty" xml:"originatingAppointment,omitempty"`
	// Date/Time that the appointment is to take place.
	//
	// Comment: If actual start/end times are not known then the requested period start/end can be used.
	Start *string `bson:"start,omitempty" json:"start,omitempty" xml:"start,omitempty"`
	// Date/Time that the appointment is to conclude.
	//
	// Comment: If actual start/end times are not known then the requested period start/end can be used.
	End *string `bson:"end,omitempty" json:"end,omitempty" xml:"end,omitempty"`
	// Number of minutes that the appointment is to take.
	//
	// This can be less than the duration between the start and end times.  For example, where the actual time of appointment is only an estimate or if a 30 minute appointment is being requested, but any time would work.  Also, if there is, for example, a planned 15 minute break in the middle of a long appointment, the duration may be 15 minutes less than the difference between the start and end.
	MinutesDuration *int `bson:"minutesDuration,omitempty" json:"minutesDuration,omitempty" xml:"minutesDuration,omitempty"`
	// A set of date ranges (potentially including times) that the appointment is preferred to be scheduled within.
	//
	// The duration (usually in minutes) could also be provided to indicate the length of the appointment to fill and populate the start/end times for the actual allocated time.
	//
	// However, in other situations the duration may be calculated by the scheduling system.
	//
	// Comment: This does not introduce a capacity for recurring appointments.
	RequestedPeriod []Period `bson:"requestedPeriod,omitempty" json:"requestedPeriod,omitempty" xml:"requestedPeriod,omitempty"`
	// The slots from the participants' schedules that will be filled by the appointment.
	Slot []Reference `bson:"slot,omitempty" json:"slot,omitempty" xml:"slot,omitempty"`
	// The set of accounts that is expected to be used for billing the activities that result from this Appointment.
	//
	// Comment: The specified account(s) could be those identified during pre-registration workflows in preparation for an upcoming Encounter.
	Account []Reference `bson:"account,omitempty" json:"account,omitempty" xml:"account,omitempty"`
	// The date that this appointment was initially created.
	//
	// This could be different to the meta.lastModified value on the initial entry, as this could have been before the resource was created on the FHIR server, and should remain unchanged over the lifespan of the appointment.
	//
	// Comment: This property is required for many use cases where the age of an appointment is considered in processing workflows for scheduling and billing of appointments.
	Created *string `bson:"created,omitempty" json:"created,omitempty" xml:"created,omitempty"`
	// The date/time describing when the appointment was cancelled.
	//
	// Comment: If the appointment was cancelled multiple times, this would reflect the most recent cancellation.
	CancellationDate *string `bson:"cancellationDate,omitempty" json:"cancellationDate,omitempty" xml:"cancellationDate,omitempty"`
	// Additional notes/comments about the appointment.
	//
	// Comment: Additional text to aid in facilitating the appointment. For instance, a note might be, "patient should proceed immediately to infusion room upon arrival"Where this is a planned appointment and the start/end dates are not set then this field can be used to provide additional guidance on the details of the appointment request, including any restrictions on when to book it.Typically only the concept.text will be used, however occasionally a reference to some generic documentation (or specific) and also supports coded instructions if/when they are required.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// While Appointment.note contains information for internal use, Appointment.patientInstructions is used to capture patient facing information about the Appointment (e.g.
	//
	// please bring your referral or fast from 8pm night before).
	PatientInstruction []CodeableReference `bson:"patientInstruction,omitempty" json:"patientInstruction,omitempty" xml:"patientInstruction,omitempty"`
	// The request this appointment is allocated to assess (e.g.
	//
	// incoming referral or procedure request).
	BasedOn []Reference `bson:"basedOn,omitempty" json:"basedOn,omitempty" xml:"basedOn,omitempty"`
	// The patient or group associated with the appointment, if they are to be present (usually) then they should also be included in the participant backbone element.
	Subject *Reference `bson:"subject,omitempty" json:"subject,omitempty" xml:"subject,omitempty"`
	// List of participants involved in the appointment.
	Participant []AppointmentParticipant `bson:"participant" json:"participant" xml:"participant"`
	// The sequence number that identifies a specific appointment in a recurring pattern.
	RecurrenceID *int `bson:"recurrenceId,omitempty" json:"recurrenceId,omitempty" xml:"recurrenceId,omitempty"`
	// This appointment varies from the recurring pattern.
	//
	// Comment: For example, if a patient has physical therapy appointments every Tuesday and Thursday, but due to a national holiday, a single occurrence was rescheduled to a different day, while the rest of the recurring series remained on Tuesday / Thursday.
	OccurrenceChanged *bool `bson:"occurrenceChanged,omitempty" json:"occurrenceChanged,omitempty" xml:"occurrenceChanged,omitempty"`
	// The details of the recurrence pattern or template that is used to generate recurring appointments.
	RecurrenceTemplate []AppointmentRecurrenceTemplate `bson:"recurrenceTemplate,omitempty" json:"recurrenceTemplate,omitempty" xml:"recurrenceTemplate,omitempty"`
}

type AppointmentParticipant struct {
	// Unique id for the element within a resource (for internal references).
	//
	// This may be any string value that does not contain spaces.
	ID *string `bson:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element.
	//
	// To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	//
	// Comment: There can be no stigma associated with the use of extensions by any application, project, or standard - regardless of the institution or jurisdiction that uses or defines the extensions.  The use of extensions is what allows the FHIR specification to retain a core level of simplicity for everyone.
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty" xml:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants.
	//
	// Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	//
	// Comment: There can be no stigma associated with the use of extensions by any application, project, or standard - regardless of the institution or jurisdiction that uses or defines the extensions.  The use of extensions is what allows the FHIR specification to retain a core level of simplicity for everyone.
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty" xml:"modifierExtension,omitempty"`
	// Role of participant in the appointment.
	//
	// Comment: The role of the participant can be used to declare what the actor will be doing in the scope of this appointment.If the actor is not specified, then it is expected that the actor will be filled in at a later stage of planning.This value SHALL be the same when creating an AppointmentResponse so that they can be matched, and subsequently update the Appointment.
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Participation period of the actor.
	Period *Period `bson:"period,omitempty" json:"period,omitempty" xml:"period,omitempty"`
	// The individual, device, location, or service participating in the appointment.
	//
	// Comment: Where a CareTeam is provided, this does not imply that the entire team is included, just a single member from the group with the appropriate role. Where multiple members are required, please include the CareTeam the required number of times.
	Actor *Reference `bson:"actor,omitempty" json:"actor,omitempty" xml:"actor,omitempty"`
	// Whether this participant is required to be present at the meeting.
	//
	// If false, the participant is optional.
	//
	// Comment: For the use-case where two doctors need to meet to discuss the results for a specific patient, and the patient is not required to be present include the patient in the subject field, but do not include them as a participant - this was formerly done prior to R5 with required='information-only'.
	Required *bool `bson:"required,omitempty" json:"required,omitempty" xml:"required,omitempty"`
	// Participation status of the actor.
	Status ParticipationStatus `bson:"status" json:"status" xml:"status"`
}

type AppointmentRecurrenceTemplate struct {
	// Unique id for the element within a resource (for internal references).
	//
	// This may be any string value that does not contain spaces.
	ID *string `bson:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element.
	//
	// To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	//
	// Comment: There can be no stigma associated with the use of extensions by any application, project, or standard - regardless of the institution or jurisdiction that uses or defines the extensions.  The use of extensions is what allows the FHIR specification to retain a core level of simplicity for everyone.
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty" xml:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants.
	//
	// Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	//
	// Comment: There can be no stigma associated with the use of extensions by any application, project, or standard - regardless of the institution or jurisdiction that uses or defines the extensions.  The use of extensions is what allows the FHIR specification to retain a core level of simplicity for everyone.
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty" xml:"modifierExtension,omitempty"`
	// The timezone of the recurring appointment occurrences.
	Timezone *CodeableConcept `bson:"timezone,omitempty" json:"timezone,omitempty" xml:"timezone,omitempty"`
	// How often the appointment series should recur.
	RecurrenceType CodeableConcept `bson:"recurrenceType" json:"recurrenceType" xml:"recurrenceType"`
	// Recurring appointments will not occur after this date.
	LastOccurrenceDate *string `bson:"lastOccurrenceDate,omitempty" json:"lastOccurrenceDate,omitempty" xml:"lastOccurrenceDate,omitempty"`
	// How many appointments are planned in the recurrence.
	OccurrenceCount *int `bson:"occurrenceCount,omitempty" json:"occurrenceCount,omitempty" xml:"occurrenceCount,omitempty"`
	// The list of specific dates that will have appointments generated.
	OccurrenceDate []string `bson:"occurrenceDate,omitempty" json:"occurrenceDate,omitempty" xml:"occurrenceDate,omitempty"`
	// Information about weekly recurring appointments.
	WeeklyTemplate *AppointmentRecurrenceTemplateWeeklyTemplate `bson:"weeklyTemplate,omitempty" json:"weeklyTemplate,omitempty" xml:"weeklyTemplate,omitempty"`
	// Information about monthly recurring appointments.
	MonthlyTemplate *AppointmentRecurrenceTemplateMonthlyTemplate `bson:"monthlyTemplate,omitempty" json:"monthlyTemplate,omitempty" xml:"monthlyTemplate,omitempty"`
	// Information about yearly recurring appointments.
	YearlyTemplate *AppointmentRecurrenceTemplateYearlyTemplate `bson:"yearlyTemplate,omitempty" json:"yearlyTemplate,omitempty" xml:"yearlyTemplate,omitempty"`
	// Any dates, such as holidays, that should be excluded from the recurrence.
	//
	// Comment: excludingDate is an alternative to excludingRecurrenceId.  You may say either that 12/25/XXXX should be excluded using excludingDate, or you may indicate the recurrenceId of the appointment that would fall on the excluded date in the excludingRecurrenceId property.
	ExcludingDate []string `bson:"excludingDate,omitempty" json:"excludingDate,omitempty" xml:"excludingDate,omitempty"`
	// Any dates, such as holidays, that should be excluded from the recurrence.
	//
	// Comment: excludingRecurrenceId is an alternative to excludingDate.  You may say either that 12/25/XXXX should be excluded using excludingDate, or you may indicate the recurrenceId of the appointment that would fall on the excluded date in the excludingRecurrenceId property.
	ExcludingRecurrenceID []int `bson:"excludingRecurrenceId,omitempty" json:"excludingRecurrenceId,omitempty" xml:"excludingRecurrenceId,omitempty"`
}

type AppointmentRecurrenceTemplateWeeklyTemplate struct {
	// Unique id for the element within a resource (for internal references).
	//
	// This may be any string value that does not contain spaces.
	ID *string `bson:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element.
	//
	// To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	//
	// Comment: There can be no stigma associated with the use of extensions by any application, project, or standard - regardless of the institution or jurisdiction that uses or defines the extensions.  The use of extensions is what allows the FHIR specification to retain a core level of simplicity for everyone.
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty" xml:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants.
	//
	// Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	//
	// Comment: There can be no stigma associated with the use of extensions by any application, project, or standard - regardless of the institution or jurisdiction that uses or defines the extensions.  The use of extensions is what allows the FHIR specification to retain a core level of simplicity for everyone.
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty" xml:"modifierExtension,omitempty"`
	// Indicates that recurring appointments should occur on Mondays.
	Monday *bool `bson:"monday,omitempty" json:"monday,omitempty" xml:"monday,omitempty"`
	// Indicates that recurring appointments should occur on Tuesdays.
	Tuesday *bool `bson:"tuesday,omitempty" json:"tuesday,omitempty" xml:"tuesday,omitempty"`
	// Indicates that recurring appointments should occur on Wednesdays.
	Wednesday *bool `bson:"wednesday,omitempty" json:"wednesday,omitempty" xml:"wednesday,omitempty"`
	// Indicates that recurring appointments should occur on Thursdays.
	Thursday *bool `bson:"thursday,omitempty" json:"thursday,omitempty" xml:"thursday,omitempty"`
	// Indicates that recurring appointments should occur on Fridays.
	Friday *bool `bson:"friday,omitempty" json:"friday,omitempty" xml:"friday,omitempty"`
	// Indicates that recurring appointments should occur on Saturdays.
	Saturday *bool `bson:"saturday,omitempty" json:"saturday,omitempty" xml:"saturday,omitempty"`
	// Indicates that recurring appointments should occur on Sundays.
	Sunday *bool `bson:"sunday,omitempty" json:"sunday,omitempty" xml:"sunday,omitempty"`
	// The interval defines if the recurrence is every nth week.
	//
	// The default is every week, so it is expected that this value will be 2 or more.e.g. For recurring every second week this interval would be 2, or every third week the interval would be 3.
	WeekInterval *int `bson:"weekInterval,omitempty" json:"weekInterval,omitempty" xml:"weekInterval,omitempty"`
}

type AppointmentRecurrenceTemplateMonthlyTemplate struct {
	// Unique id for the element within a resource (for internal references).
	//
	// This may be any string value that does not contain spaces.
	ID *string `bson:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element.
	//
	// To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	//
	// Comment: There can be no stigma associated with the use of extensions by any application, project, or standard - regardless of the institution or jurisdiction that uses or defines the extensions.  The use of extensions is what allows the FHIR specification to retain a core level of simplicity for everyone.
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty" xml:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants.
	//
	// Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	//
	// Comment: There can be no stigma associated with the use of extensions by any application, project, or standard - regardless of the institution or jurisdiction that uses or defines the extensions.  The use of extensions is what allows the FHIR specification to retain a core level of simplicity for everyone.
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty" xml:"modifierExtension,omitempty"`
	// Indicates that appointments in the series of recurring appointments should occur on a specific day of the month.
	DayOfMonth *int `bson:"dayOfMonth,omitempty" json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
	// Indicates which week within a month the appointments in the series of recurring appointments should occur on.
	NthWeekOfMonth *Coding `bson:"nthWeekOfMonth,omitempty" json:"nthWeekOfMonth,omitempty" xml:"nthWeekOfMonth,omitempty"`
	// Indicates which day of the week the recurring appointments should occur each nth week.
	//
	// Comment: This property is intended to be used with Appointment.recurrenceTemplate.monthly.nthWeek.
	DayOfWeek *Coding `bson:"dayOfWeek,omitempty" json:"dayOfWeek,omitempty" xml:"dayOfWeek,omitempty"`
	// Indicates that recurring appointments should occur every nth month.
	MonthInterval int `bson:"monthInterval" json:"monthInterval" xml:"monthInterval"`
}

type AppointmentRecurrenceTemplateYearlyTemplate struct {
	// Unique id for the element within a resource (for internal references).
	//
	// This may be any string value that does not contain spaces.
	ID *string `bson:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element.
	//
	// To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	//
	// Comment: There can be no stigma associated with the use of extensions by any application, project, or standard - regardless of the institution or jurisdiction that uses or defines the extensions.  The use of extensions is what allows the FHIR specification to retain a core level of simplicity for everyone.
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty" xml:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants.
	//
	// Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	//
	// Comment: There can be no stigma associated with the use of extensions by any application, project, or standard - regardless of the institution or jurisdiction that uses or defines the extensions.  The use of extensions is what allows the FHIR specification to retain a core level of simplicity for everyone.
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty" xml:"modifierExtension,omitempty"`
	// Appointment recurs every nth year.
	YearInterval int `bson:"yearInterval" json:"yearInterval" xml:"yearInterval"`
}

func (r Appointment) ResourceType() string {
	return "Appointment"
}

// MarshalJSON implements json.Marshaler
func (r Appointment) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Appointment
		ResourceType string `json:"resourceType"`
	}{
		Appointment:  r,
		ResourceType: r.ResourceType(),
	})
}
