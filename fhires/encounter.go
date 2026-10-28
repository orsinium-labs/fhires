package fhires

import "encoding/json/v2"

// An interaction between healthcare provider(s), and/or patient(s) for the purpose of providing healthcare service(s) or assessing the health status of patient(s).
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Encounter
type Encounter struct {
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
	// Identifier(s) by which this encounter is known.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The current state of the encounter (not the state of the patient within the encounter - that is subjectState).
	//
	// Comment: Note that internal business rules will determine the appropriate transitions that may occur between statuses (and also classes).
	Status EncounterStatus `bson:"status" json:"status" xml:"status"`
	// Concepts representing classification of patient encounter such as ambulatory (outpatient), inpatient, emergency, home health or others due to local variations.
	Class []CodeableConcept `bson:"class,omitempty" json:"class,omitempty" xml:"class,omitempty"`
	// Indicates the urgency of the encounter.
	Priority *CodeableConcept `bson:"priority,omitempty" json:"priority,omitempty" xml:"priority,omitempty"`
	// Specific type of encounter (e.g.
	//
	// e-mail consultation, surgical day-care, skilled nursing, rehabilitation).
	//
	// Comment: Since there are many ways to further classify encounters, this element is 0..*.
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Broad categorization of the service that is to be provided (e.g. cardiology).
	ServiceType []CodeableReference `bson:"serviceType,omitempty" json:"serviceType,omitempty" xml:"serviceType,omitempty"`
	// The patient or group related to this encounter.
	//
	// In some use-cases the patient MAY not be present, such as a case meeting about a patient between several practitioners or a careteam.
	//
	// Comment: While the encounter is always about the patient, the patient might not actually be known in all contexts of use, and there may be a group of patients that could be anonymous (such as in a group therapy for Alcoholics Anonymous - where the recording of the encounter could be used for billing on the number of people/staff and not important to the context of the specific patients) or alternately in veterinary care a herd of sheep receiving treatment (where the animals are not individually tracked).
	Subject *Reference `bson:"subject,omitempty" json:"subject,omitempty" xml:"subject,omitempty"`
	// The subjectStatus value can be used to track the patient's status within the encounter.
	//
	// It details whether the patient has arrived or departed, has been triaged or is currently in a waiting status.
	//
	// Comment: Different use-cases are likely to have different permitted transitions between states, such as an Emergency department could use `arrived` when the patient first presents, then `triaged` once has been assessed by a nurse, then `receiving-care` once treatment begins, however other sectors may use a different set of these values, or their own custom set in place of this example valueset provided.
	SubjectStatus *CodeableConcept `bson:"subjectStatus,omitempty" json:"subjectStatus,omitempty" xml:"subjectStatus,omitempty"`
	// Where a specific encounter should be classified as a part of a specific episode(s) of care this field should be used.
	//
	// This association can facilitate grouping of related encounters together for a specific purpose, such as government reporting, issue tracking, association via a common problem.  The association is recorded on the encounter as these are typically created after the episode of care and grouped on entry rather than editing the episode of care to append another encounter to it (the episode of care could span years).
	EpisodeOfCare []Reference `bson:"episodeOfCare,omitempty" json:"episodeOfCare,omitempty" xml:"episodeOfCare,omitempty"`
	// The request this encounter satisfies (e.g. incoming referral or procedure request).
	BasedOn []Reference `bson:"basedOn,omitempty" json:"basedOn,omitempty" xml:"basedOn,omitempty"`
	// The group(s) of individuals, organizations that are allocated to participate in this encounter.
	//
	// The participants backbone will record the actuals of when these individuals participated during the encounter.
	CareTeam []Reference `bson:"careTeam,omitempty" json:"careTeam,omitempty" xml:"careTeam,omitempty"`
	// Another Encounter of which this encounter is a part of (administratively or in time).
	//
	// Comment: This is also used for associating a child's encounter back to the mother's encounter.Refer to the Notes section in the Patient resource for further details.
	PartOf *Reference `bson:"partOf,omitempty" json:"partOf,omitempty" xml:"partOf,omitempty"`
	// The organization that is primarily responsible for this Encounter's services.
	//
	// This MAY be the same as the organization on the Patient record, however it could be different, such as if the actor performing the services was from an external organization (which may be billed seperately) for an external consultation.  Refer to the colonoscopy example on the Encounter examples tab.
	ServiceProvider *Reference `bson:"serviceProvider,omitempty" json:"serviceProvider,omitempty" xml:"serviceProvider,omitempty"`
	// The list of people responsible for providing the service.
	//
	// Comment: Any Patient or Group present in the participation.actor must also be the subject, though the subject may be absent from the participation.actor for cases where the patient (or group) is not present, such as during a case review conference.
	Participant []EncounterParticipant `bson:"participant,omitempty" json:"participant,omitempty" xml:"participant,omitempty"`
	// The appointment that scheduled this encounter.
	Appointment []Reference `bson:"appointment,omitempty" json:"appointment,omitempty" xml:"appointment,omitempty"`
	// Connection details of a virtual service (e.g. conference call).
	//
	// Comment: There are two types of virtual meetings that often exist:* a persistent, virtual meeting room that can only be used for a single purpose at a time, * and a dynamic virtual meeting room that is generated on demand for a specific purpose.Implementers may consider using Location.virtualService for persistent meeting rooms.If each participant would have a different meeting link, an extension using the VirtualServiceContactDetail  can be applied to the Encounter.participant BackboneElement.
	VirtualService []VirtualServiceDetail `bson:"virtualService,omitempty" json:"virtualService,omitempty" xml:"virtualService,omitempty"`
	// The actual start and end time of the encounter.
	//
	// Comment: If not (yet) known, the end of the Period may be omitted.
	ActualPeriod *Period `bson:"actualPeriod,omitempty" json:"actualPeriod,omitempty" xml:"actualPeriod,omitempty"`
	// The planned start date/time (or admission date) of the encounter.
	PlannedStartDate *string `bson:"plannedStartDate,omitempty" json:"plannedStartDate,omitempty" xml:"plannedStartDate,omitempty"`
	// The planned end date/time (or discharge date) of the encounter.
	PlannedEndDate *string `bson:"plannedEndDate,omitempty" json:"plannedEndDate,omitempty" xml:"plannedEndDate,omitempty"`
	// Actual quantity of time the encounter lasted.
	//
	// This excludes the time during leaves of absence.When missing it is the time in between the start and end values.
	//
	// Comment: If the precision on these values is low (e.g. to the day only) then this may be considered was an all day (or multi-day) encounter, unless the duration is included, where that amount of time occurred sometime during the interval.May differ from the time in `Encounter.period` due to leave of absence(s).
	Length *Duration `bson:"length,omitempty" json:"length,omitempty" xml:"length,omitempty"`
	// The list of medical reasons that are expected to be addressed during the episode of care.
	//
	// Comment: The reason communicates what medical problem the patient has that should be addressed during the episode of care.  This reason could be patient reported complaint, a clinical indication that was determined in a previous encounter or episode of care, or some planned care such as an immunization recommendation.  In the case where you have a primary reason, but are expecting to also address other problems, you can list the primary reason with a use code of 'Chief Complaint', while the other problems being addressed would have a use code of 'Reason for Visit'.Examples: * pregnancy would use HealthcareService or a coding as the reason * patient home monitoring could use Condition as the reason
	Reason []EncounterReason `bson:"reason,omitempty" json:"reason,omitempty" xml:"reason,omitempty"`
	// The list of diagnosis relevant to this encounter.
	//
	// Comment: Also note that for the purpose of billing, the diagnoses are recorded in the account where they can be ranked appropriately for how the invoicing/claiming documentation needs to be prepared.
	Diagnosis []EncounterDiagnosis `bson:"diagnosis,omitempty" json:"diagnosis,omitempty" xml:"diagnosis,omitempty"`
	// The set of accounts that may be used for billing for this Encounter.
	//
	// Comment: The billing system may choose to allocate billable items associated with the Encounter to different referenced Accounts based on internal business rules.
	Account []Reference `bson:"account,omitempty" json:"account,omitempty" xml:"account,omitempty"`
	// Diet preferences reported by the patient.
	//
	// Comment: For example, a patient may request both a dairy-free and nut-free diet preference (not mutually exclusive).
	DietPreference []CodeableConcept `bson:"dietPreference,omitempty" json:"dietPreference,omitempty" xml:"dietPreference,omitempty"`
	// Any special requests that have been made for this encounter, such as the provision of specific equipment or other things.
	SpecialArrangement []CodeableConcept `bson:"specialArrangement,omitempty" json:"specialArrangement,omitempty" xml:"specialArrangement,omitempty"`
	// Special courtesies that may be provided to the patient during the encounter (VIP, board member, professional courtesy).
	//
	// Comment: Although the specialCourtesy property can contain values like VIP, the purpose of this field is intended to be used for flagging additional `benefits` that might occur for the patient during the encounter.It could include things like the patient is to have a private room, special room features, receive a friendly visit from hospital adminisitration, or should be briefed on treatment by senior staff during the stay.It is not specifically intended to be used for securing the specific record - that is the purpose of the security meta tag, and where appropriate, both fields could be used.
	SpecialCourtesy []CodeableConcept `bson:"specialCourtesy,omitempty" json:"specialCourtesy,omitempty" xml:"specialCourtesy,omitempty"`
	// Details about the stay during which a healthcare service is provided.This does not describe the event of admitting the patient, but rather any information that is relevant from the time of admittance until the time of discharge.
	//
	// Comment: An Encounter may cover more than just the inpatient stay. Contexts such as outpatients, community clinics, and aged care facilities are also included.The duration recorded in the period of this encounter covers the entire scope of this admission record.
	Admission *EncounterAdmission `bson:"admission,omitempty" json:"admission,omitempty" xml:"admission,omitempty"`
	// List of locations where  the patient has been during this encounter.
	//
	// Comment: Virtual encounters can be recorded in the Encounter by specifying a location reference to a location of type "kind" such as "client's home" and an encounter.class = "virtual".
	Location []EncounterLocation `bson:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
}

type EncounterParticipant struct {
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
	// Role of participant in encounter.
	//
	// Comment: The participant type indicates how an individual actor participates in an encounter. It includes non-practitioner participants, and for practitioners this is to describe the action type in the context of this encounter (e.g. Admitting Dr, Attending Dr, Translator, Consulting Dr). This is different to the practitioner roles which are functional roles, derived from terms of employment, education, licensing, etc.
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The period of time that the specified participant participated in the encounter.
	//
	// These can overlap or be sub-sets of the overall encounter's period.
	Period *Period `bson:"period,omitempty" json:"period,omitempty" xml:"period,omitempty"`
	// Person involved in the encounter, the patient/group is also included here to indicate that the patient was actually participating in the encounter.
	//
	// Not including the patient here covers use cases such as a case meeting between practitioners about a patient - non contact times.
	//
	// Comment: For planning purposes, Appointments may include a CareTeam participant to indicate that one specific person from the CareTeam will be assigned, but that assignment might not happen until the Encounter begins. Hence CareTeam is not included in Encounter.participant, as the specific individual should be assigned and represented as a Practitioner or other person resource.Similarly, Location can be included in Appointment.participant to assist with planning.  However, the patient location is tracked on the Encounter in the Encounter.location property to allow for additional metadata and history to be recorded.The role of the participant can be used to declare what the actor will be doing in the scope of this encounter participation.If the individual is not specified during planning, then it is expected that the individual will be filled in at a later stage prior to the encounter commencing.
	Actor *Reference `bson:"actor,omitempty" json:"actor,omitempty" xml:"actor,omitempty"`
}

type EncounterReason struct {
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
	// What the reason value should be used as e.g.
	//
	// Chief Complaint, Health Concern, Health Maintenance (including screening).
	Use []CodeableConcept `bson:"use,omitempty" json:"use,omitempty" xml:"use,omitempty"`
	// Reason the encounter takes place, expressed as a code or a reference to another resource.
	//
	// For admissions, this can be used for a coded admission diagnosis.
	Value []CodeableReference `bson:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

type EncounterDiagnosis struct {
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
	// The coded diagnosis or a reference to a Condition (with other resources referenced in the evidence.detail), the use property will indicate the purpose of this specific diagnosis.
	Condition []CodeableReference `bson:"condition,omitempty" json:"condition,omitempty" xml:"condition,omitempty"`
	// Role that this diagnosis has within the encounter (e.g.
	//
	// admission, billing, discharge …).
	Use []CodeableConcept `bson:"use,omitempty" json:"use,omitempty" xml:"use,omitempty"`
}

type EncounterAdmission struct {
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
	// Pre-admission identifier.
	PreAdmissionIdentifier *Identifier `bson:"preAdmissionIdentifier,omitempty" json:"preAdmissionIdentifier,omitempty" xml:"preAdmissionIdentifier,omitempty"`
	// The location/organization from which the patient came before admission.
	Origin *Reference `bson:"origin,omitempty" json:"origin,omitempty" xml:"origin,omitempty"`
	// From where patient was admitted (physician referral, transfer).
	AdmitSource *CodeableConcept `bson:"admitSource,omitempty" json:"admitSource,omitempty" xml:"admitSource,omitempty"`
	// Indicates that this encounter is directly related to a prior admission, often because the conditions addressed in the prior admission were not fully addressed.
	ReAdmission *CodeableConcept `bson:"reAdmission,omitempty" json:"reAdmission,omitempty" xml:"reAdmission,omitempty"`
	// Location/organization to which the patient is discharged.
	Destination *Reference `bson:"destination,omitempty" json:"destination,omitempty" xml:"destination,omitempty"`
	// Category or kind of location after discharge.
	DischargeDisposition *CodeableConcept `bson:"dischargeDisposition,omitempty" json:"dischargeDisposition,omitempty" xml:"dischargeDisposition,omitempty"`
}

type EncounterLocation struct {
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
	// The location where the encounter takes place.
	Location Reference `bson:"location" json:"location" xml:"location"`
	// The status of the participants' presence at the specified location during the period specified.
	//
	// If the participant is no longer at the location, then the period will have an end date/time.
	//
	// Comment: When the patient is no longer active at a location, then the period end date is entered, and the status may be changed to completed.
	Status *EncounterLocationStatus `bson:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// This will be used to specify the required levels (bed/ward/room/etc.) desired to be recorded to simplify either messaging or query.
	//
	// Comment: This information is de-normalized from the Location resource to support the easier understanding of the encounter resource and processing in messaging or query.
	//
	// There may be many levels in the hierachy, and this may only pic specific levels that are required for a specific usage scenario.
	Form *CodeableConcept `bson:"form,omitempty" json:"form,omitempty" xml:"form,omitempty"`
	// Time period during which the patient was present at the location.
	Period *Period `bson:"period,omitempty" json:"period,omitempty" xml:"period,omitempty"`
}

func (r Encounter) ResourceType() string {
	return "Encounter"
}

// MarshalJSON implements json.Marshaler
func (r Encounter) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Encounter
		ResourceType string `json:"resourceType"`
	}{
		Encounter:    r,
		ResourceType: r.ResourceType(),
	})
}
