package fhires

import "encoding/json/v2"

// An action that is or was performed on or for a patient, practitioner, device, organization, or location.
//
// For example, this can be a physical intervention on a patient like an operation, or less invasive like long term services, counseling, or hypnotherapy.  This can be a quality or safety inspection for a location, organization, or device.  This can be an accreditation procedure on a practitioner for licensing.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Procedure
type Procedure struct {
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
	// Business identifiers assigned to this procedure by the performer or other systems which remain constant as the resource is updated and is propagated from server to server.
	//
	// Comment: This is a business identifier, not a resource identifier (see [discussion](resource.html#identifiers)).  It is best practice for the identifier to only appear on a single resource instance, however business practices may occasionally dictate that multiple resource instances with the same identifier can exist - possibly even with different resource types.  For example, multiple Patient and Person resource instances might share the same social insurance number.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The URL pointing to a FHIR-defined protocol, guideline, order set or other definition that is adhered to in whole or in part by this Procedure.
	InstantiatesCanonical []string `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty" xml:"instantiatesCanonical,omitempty"`
	// The URL pointing to an externally maintained protocol, guideline, order set or other definition that is adhered to in whole or in part by this Procedure.
	//
	// Comment: This might be an HTML page, PDF, etc. or could just be a non-resolvable URI identifier.
	InstantiatesUri []string `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty" xml:"instantiatesUri,omitempty"`
	// A reference to a resource that contains details of the request for this procedure.
	BasedOn []Reference `bson:"basedOn,omitempty" json:"basedOn,omitempty" xml:"basedOn,omitempty"`
	// A larger event of which this particular procedure is a component or step.
	//
	// Comment: The MedicationAdministration resource has a partOf reference to Procedure, but this is not a circular reference.   For example, the anesthesia MedicationAdministration is part of the surgical Procedure (MedicationAdministration.partOf = Procedure).  For example, the procedure to insert the IV port for an IV medication administration is part of the medication administration (Procedure.partOf = MedicationAdministration).
	PartOf []Reference `bson:"partOf,omitempty" json:"partOf,omitempty" xml:"partOf,omitempty"`
	// A code specifying the state of the procedure.
	//
	// Generally, this will be the in-progress or completed state.
	//
	// Comment: The "unknown" code is not to be used to convey other statuses.  The "unknown" code should be used when one of the statuses applies, but the authoring system doesn't know the current state of the procedure.
	//
	// This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid.
	Status EventStatus `bson:"status" json:"status" xml:"status"`
	// Captures the reason for the current state of the procedure.
	//
	// Comment: This is generally only used for "exception" statuses such as "not-done", "suspended" or "aborted". The reason for performing the event at all is captured in reasonCode, not here.
	StatusReason *CodeableConcept `bson:"statusReason,omitempty" json:"statusReason,omitempty" xml:"statusReason,omitempty"`
	// A code that classifies the procedure for searching, sorting and display purposes (e.g.
	//
	// "Surgical Procedure").
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
	// The specific procedure that is performed.
	//
	// Use text if the exact nature of the procedure cannot be coded (e.g. "Laparoscopic Appendectomy").
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// On whom or on what the procedure was performed.
	//
	// This is usually an individual human, but can also be performed on animals, groups of humans or animals, organizations or practitioners (for licensing), locations or devices (for safety inspections or regulatory authorizations).  If the actual focus of the procedure is different from the subject, the focus element specifies the actual focus of the procedure.
	Subject Reference `bson:"subject" json:"subject" xml:"subject"`
	// Who is the target of the procedure when it is not the subject of record only.
	//
	//  If focus is not present, then subject is the focus.  If focus is present and the subject is one of the targets of the procedure, include subject as a focus as well. If focus is present and the subject is not included in focus, it implies that the procedure was only targeted on the focus. For example, when a caregiver is given education for a patient, the caregiver would be the focus and the procedure record is associated with the subject (e.g. patient).  For example, use focus when recording the target of the education, training, or counseling is the parent or relative of a patient.
	Focus *Reference `bson:"focus,omitempty" json:"focus,omitempty" xml:"focus,omitempty"`
	// The Encounter during which this Procedure was created or performed or to which the creation of this record is tightly associated.
	//
	// Comment: This will typically be the encounter the event occurred within, but some activities may be initiated prior to or after the official completion of an encounter but still be tied to the context of the encounter.
	Encounter *Reference `bson:"encounter,omitempty" json:"encounter,omitempty" xml:"encounter,omitempty"`
	// Estimated or actual date, date-time, period, or age when the procedure did occur or is occurring.
	//
	//  Allows a period to support complex procedures that span more than one date, and also allows for the length of the procedure to be captured.
	//
	// Comment: This indicates when the procedure actually occurred or is occurring, not when it was asked/requested/ordered to occur. For the latter, look at the occurence element of the Request this Procedure is "basedOn". The status code allows differentiation of whether the timing reflects a historic event or an ongoing event. Ongoing events should not include an upper bound in the Period or Timing.bounds.
	//
	// Age is generally used when the patient reports an age at which the procedure was performed. Range is generally used when the patient reports an age range when the procedure was performed, such as sometime between 20-25 years old.  dateTime supports a range of precision due to some procedures being reported as past procedures that might not have millisecond precision while other procedures performed and documented during the encounter might have more precise UTC timestamps with timezone.
	OccurrenceDateTime *string `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty" xml:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *Period `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty" xml:"occurrencePeriod,omitempty"`
	OccurrenceString   *string `bson:"occurrenceString,omitempty" json:"occurrenceString,omitempty" xml:"occurrenceString,omitempty"`
	OccurrenceAge      *Age    `bson:"occurrenceAge,omitempty" json:"occurrenceAge,omitempty" xml:"occurrenceAge,omitempty"`
	OccurrenceRange    *Range  `bson:"occurrenceRange,omitempty" json:"occurrenceRange,omitempty" xml:"occurrenceRange,omitempty"`
	OccurrenceTiming   *Timing `bson:"occurrenceTiming,omitempty" json:"occurrenceTiming,omitempty" xml:"occurrenceTiming,omitempty"`
	// The date the occurrence of the procedure was first captured in the record regardless of Procedure.status (potentially after the occurrence of the event).
	Recorded *string `bson:"recorded,omitempty" json:"recorded,omitempty" xml:"recorded,omitempty"`
	// Individual who recorded the record and takes responsibility for its content.
	Recorder *Reference `bson:"recorder,omitempty" json:"recorder,omitempty" xml:"recorder,omitempty"`
	// Indicates if this record was captured as a secondary 'reported' record rather than as an original primary source-of-truth record.
	//
	//  It may also indicate the source of the report.
	ReportedBoolean   *bool      `bson:"reportedBoolean,omitempty" json:"reportedBoolean,omitempty" xml:"reportedBoolean,omitempty"`
	ReportedReference *Reference `bson:"reportedReference,omitempty" json:"reportedReference,omitempty" xml:"reportedReference,omitempty"`
	// Indicates who or what performed the procedure and how they were involved.
	Performer []ProcedurePerformer `bson:"performer,omitempty" json:"performer,omitempty" xml:"performer,omitempty"`
	// The location where the procedure actually happened.
	//
	//  E.g. a newborn at home, a tracheostomy at a restaurant.
	Location *Reference `bson:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// The coded reason or reference why the procedure was performed.
	//
	// This may be a coded entity of some type, be present as text, or be a reference to one of several resources that justify the procedure.
	//
	// Comment: Use Procedure.reason.concept when a code sufficiently describes the reason. Use Procedure.reason.reference when referencing a resource, which allows more information to be conveyed, such as onset date.  For a single Procedure.reason, if both Procedure.reason.concept and Procedure.reason.reference are present, they are expected to be consistent with each other.
	Reason []CodeableReference `bson:"reason,omitempty" json:"reason,omitempty" xml:"reason,omitempty"`
	// Detailed and structured anatomical location information.
	//
	// Multiple locations are allowed - e.g. multiple punch biopsies of a lesion.
	//
	// Comment: If the use case requires attributes from the BodySite resource (e.g. to identify and track separately) then use the standard extension [http://hl7.org/fhir/StructureDefinition/procedure-targetBodyStructure](http://hl7.org/fhir/extensions/StructureDefinition-procedure-targetBodyStructure.html).
	BodySite []CodeableConcept `bson:"bodySite,omitempty" json:"bodySite,omitempty" xml:"bodySite,omitempty"`
	// The outcome of the procedure - did it resolve the reasons for the procedure being performed?
	//
	// Comment: If outcome contains narrative text only, it can be captured using the CodeableConcept.text.
	Outcome *CodeableConcept `bson:"outcome,omitempty" json:"outcome,omitempty" xml:"outcome,omitempty"`
	// This could be a histology result, pathology report, surgical report, etc.
	//
	// Comment: There could potentially be multiple reports - e.g. if this was a procedure which took multiple biopsies resulting in a number of anatomical pathology reports.
	Report []Reference `bson:"report,omitempty" json:"report,omitempty" xml:"report,omitempty"`
	// Any complications that occurred during the procedure, or in the immediate post-performance period.
	//
	// These are generally tracked separately from the notes, which will typically describe the procedure itself rather than any 'post procedure' issues.
	//
	// Comment: If complications are only expressed by the narrative text, they can be captured using the CodeableReference.concept.text.
	Complication []CodeableReference `bson:"complication,omitempty" json:"complication,omitempty" xml:"complication,omitempty"`
	// If the procedure required specific follow up - e.g.
	//
	// removal of sutures. The follow up may be represented as a simple note or could potentially be more complex, in which case the CarePlan resource can be used.
	FollowUp []CodeableConcept `bson:"followUp,omitempty" json:"followUp,omitempty" xml:"followUp,omitempty"`
	// Any other notes and comments about the procedure.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// A device that is implanted, removed or otherwise manipulated (calibration, battery replacement, fitting a prosthesis, attaching a wound-vac, etc.) as a focal portion of the Procedure.
	FocalDevice []ProcedureFocalDevice `bson:"focalDevice,omitempty" json:"focalDevice,omitempty" xml:"focalDevice,omitempty"`
	// Identifies medications, devices and any other substance used as part of the procedure.
	//
	// Comment: For devices actually implanted or removed, use Procedure.focalDevice.manipulated.
	Used []CodeableReference `bson:"used,omitempty" json:"used,omitempty" xml:"used,omitempty"`
	// Other resources from the patient record that may be relevant to the procedure.
	//
	//  The information from these resources was either used to create the instance or is provided to help with its interpretation. This extension should not be used if more specific inline elements or extensions are available.
	SupportingInfo []Reference `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty" xml:"supportingInfo,omitempty"`
}

type ProcedurePerformer struct {
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
	// Distinguishes the type of involvement of the performer in the procedure.
	//
	// For example, surgeon, anaesthetist, endoscopist.
	Function *CodeableConcept `bson:"function,omitempty" json:"function,omitempty" xml:"function,omitempty"`
	// Indicates who or what performed the procedure.
	Actor Reference `bson:"actor" json:"actor" xml:"actor"`
	// The Organization the Patient, RelatedPerson, Device, CareTeam, and HealthcareService was acting on behalf of.
	//
	// Comment: Organization, Patient, RelatedPerson, Device, CareTeam, and HealthcareService can be associated with multiple organizations. This element indicates which organization they were acting on behalf of when performing the action.
	OnBehalfOf *Reference `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty" xml:"onBehalfOf,omitempty"`
	// Time period during which the performer performed the procedure.
	Period *Period `bson:"period,omitempty" json:"period,omitempty" xml:"period,omitempty"`
}

type ProcedureFocalDevice struct {
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
	// The kind of change that happened to the device during the procedure.
	Action *CodeableConcept `bson:"action,omitempty" json:"action,omitempty" xml:"action,omitempty"`
	// The device that was manipulated (changed) during the procedure.
	Manipulated Reference `bson:"manipulated" json:"manipulated" xml:"manipulated"`
}

func (r Procedure) ResourceType() string {
	return "Procedure"
}

// MarshalJSON implements json.Marshaler
func (r Procedure) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Procedure
		ResourceType string `json:"resourceType"`
	}{
		Procedure:    r,
		ResourceType: r.ResourceType(),
	})
}
