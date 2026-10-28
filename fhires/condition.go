package fhires

import "encoding/json/v2"

// A clinical condition, problem, diagnosis, or other event, situation, issue, or clinical concept that has risen to a level of concern.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Condition
type Condition struct {
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
	// Business identifiers assigned to this condition by the performer or other systems which remain constant as the resource is updated and propagates from server to server.
	//
	// Comment: This is a business identifier, not a resource identifier (see [discussion](resource.html#identifiers)).  It is best practice for the identifier to only appear on a single resource instance, however business practices may occasionally dictate that multiple resource instances with the same identifier can exist - possibly even with different resource types.  For example, multiple Patient and a Person resource instance might share the same social insurance number.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The clinical status of the condition.
	//
	// Comment: The data type is CodeableConcept because clinicalStatus has some clinical judgment involved, such that there might need to be more specificity than the required FHIR value set allows. For example, a SNOMED coding might allow for additional specificity.  clinicalStatus is required since it is a modifier element.  For conditions that are problems list items, the clinicalStatus should not be unknown.  For conditions that are not problem list items, the clinicalStatus may be unknown.  For example, conditions derived from a claim are point in time, so those conditions may have a clinicalStatus of unknown
	ClinicalStatus CodeableConcept `bson:"clinicalStatus" json:"clinicalStatus" xml:"clinicalStatus"`
	// The verification status to support the clinical status of the condition.
	//
	//  The verification status pertains to the condition, itself, not to any specific condition attribute.
	//
	// Comment: verificationStatus is not required.  For example, when a patient has abdominal pain in the ED, there is not likely going to be a verification status.
	//
	// The data type is CodeableConcept because verificationStatus has some clinical judgment involved, such that there might need to be more specificity than the required FHIR value set allows. For example, a SNOMED coding might allow for additional specificity.
	VerificationStatus *CodeableConcept `bson:"verificationStatus,omitempty" json:"verificationStatus,omitempty" xml:"verificationStatus,omitempty"`
	// A category assigned to the condition.
	//
	// Comment: The categorization is often highly contextual and may appear poorly differentiated or not very useful in other contexts.
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
	// A subjective assessment of the severity of the condition as evaluated by the clinician.
	//
	// Comment: Coding of the severity with a terminology is preferred, where possible.
	Severity *CodeableConcept `bson:"severity,omitempty" json:"severity,omitempty" xml:"severity,omitempty"`
	// Identification of the condition, problem or diagnosis.
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// The anatomical location where this condition manifests itself.
	//
	// Comment: Only used if not implicit in code found in Condition.code. If the use case requires attributes from the BodyStructure resource (e.g. to identify and track separately) then use the standard extension [http://hl7.org/fhir/StructureDefinition/bodySite](http://hl7.org/fhir/extensions/StructureDefinition-bodySite.html).  May be a summary code, or a reference to a very precise definition of the location, or both.
	BodySite []CodeableConcept `bson:"bodySite,omitempty" json:"bodySite,omitempty" xml:"bodySite,omitempty"`
	// Indicates the patient or group who the condition record is associated with.
	Subject Reference `bson:"subject" json:"subject" xml:"subject"`
	// The Encounter during which this Condition was created or to which the creation of this record is tightly associated.
	//
	// Comment: This will typically be the encounter the event occurred within, but some activities may be initiated prior to or after the official completion of an encounter but still be tied to the context of the encounter. This record indicates the encounter this particular record is associated with.  In the case of a "new" diagnosis reflecting ongoing/revised information about the condition, this might be distinct from the first encounter in which the underlying condition was first "known".
	Encounter *Reference `bson:"encounter,omitempty" json:"encounter,omitempty" xml:"encounter,omitempty"`
	// Estimated or actual date or date-time  the condition began, in the opinion of the clinician.
	//
	// Comment: Age is generally used when the patient reports an age at which the Condition began to occur.  Period is generally used to convey an imprecise onset that occurred within the time period.  For example, Period is not intended to convey the transition period before the chronic bronchitis or COPD condition was diagnosed, but Period can be used to convey an imprecise diagnosis date.  Range is generally used to convey an imprecise age range (e.g. 4 to 6 years old).  Because a Condition.code can represent multiple levels of granularity and can be modified over time, the onset and abatement dates can have ambiguity whether those dates apply to the current Condition.code or an earlier representation of that Condition.code.   For example, if the Condition.code was initially documented as severe asthma, then it is ambiguous whether the onset and abatement dates apply to asthma (overall in that subject's lifetime) or when asthma transitioned to become severe.
	OnsetDateTime *string `bson:"onsetDateTime,omitempty" json:"onsetDateTime,omitempty" xml:"onsetDateTime,omitempty"`
	OnsetAge      *Age    `bson:"onsetAge,omitempty" json:"onsetAge,omitempty" xml:"onsetAge,omitempty"`
	OnsetPeriod   *Period `bson:"onsetPeriod,omitempty" json:"onsetPeriod,omitempty" xml:"onsetPeriod,omitempty"`
	OnsetRange    *Range  `bson:"onsetRange,omitempty" json:"onsetRange,omitempty" xml:"onsetRange,omitempty"`
	OnsetString   *string `bson:"onsetString,omitempty" json:"onsetString,omitempty" xml:"onsetString,omitempty"`
	// The date or estimated date that the condition resolved or went into remission.
	//
	// This is called "abatement" because of the many overloaded connotations associated with "remission" or "resolution" - Some conditions, such as chronic conditions, are never really resolved, but they can abate.
	//
	// Comment: There is no explicit distinction between resolution and remission because in many cases the distinction is not clear. Age is generally used when the patient reports an age at which the Condition abated.  If there is no abatement element, it is unknown whether the condition has resolved or entered remission; applications and users should generally assume that the condition is still valid.  When abatementString exists, it implies the condition is abated.
	//
	// Because a Condition.code can represent multiple levels of granularity and can be modified over time, the onset and abatement dates can have ambiguity whether those dates apply to the current Condition.code or an earlier representation of that Condition.code.   For example, if the Condition.code was initially documented as severe asthma, then it is ambiguous whether the onset and abatement dates apply to asthma (overall in that subject's lifetime) or when asthma transitioned to become severe.
	AbatementDateTime *string `bson:"abatementDateTime,omitempty" json:"abatementDateTime,omitempty" xml:"abatementDateTime,omitempty"`
	AbatementAge      *Age    `bson:"abatementAge,omitempty" json:"abatementAge,omitempty" xml:"abatementAge,omitempty"`
	AbatementPeriod   *Period `bson:"abatementPeriod,omitempty" json:"abatementPeriod,omitempty" xml:"abatementPeriod,omitempty"`
	AbatementRange    *Range  `bson:"abatementRange,omitempty" json:"abatementRange,omitempty" xml:"abatementRange,omitempty"`
	AbatementString   *string `bson:"abatementString,omitempty" json:"abatementString,omitempty" xml:"abatementString,omitempty"`
	// The recordedDate represents when this particular Condition record was created in the system, which is often a system-generated date.
	//
	// Comment: When onset date is unknown, recordedDate can be used to establish if the condition was present on or before a given date.  If the recordedDate is known and provided by a sending system, it is preferred that the receiving system preserve that recordedDate value. If the recordedDate is not provided by the sending system, the receipt timestamp is sometimes used as the recordedDate.
	RecordedDate *string `bson:"recordedDate,omitempty" json:"recordedDate,omitempty" xml:"recordedDate,omitempty"`
	// Indicates who or what participated in the activities related to the condition and how they were involved.
	Participant []ConditionParticipant `bson:"participant,omitempty" json:"participant,omitempty" xml:"participant,omitempty"`
	// A simple summary of the stage such as "Stage 3" or "Early Onset".
	//
	// The determination of the stage is disease-specific, such as cancer, retinopathy of prematurity, kidney diseases, Alzheimer's, or Parkinson disease.
	Stage []ConditionStage `bson:"stage,omitempty" json:"stage,omitempty" xml:"stage,omitempty"`
	// Supporting evidence / manifestations that are the basis of the Condition's verification status, such as evidence that confirmed or refuted the condition.
	//
	// Comment: If the condition was confirmed, but subsequently refuted, then the evidence can be cumulative including all evidence over time.  The evidence may be a simple list of coded symptoms/manifestations, or references to observations or formal assessments, or both.  For example, if the Condition.code is pneumonia, then there could be an evidence list where Condition.evidence.concept = fever (CodeableConcept), Condition.evidence.concept = cough (CodeableConcept), and Condition.evidence.reference = bronchitis (reference to Condition).
	Evidence []CodeableReference `bson:"evidence,omitempty" json:"evidence,omitempty" xml:"evidence,omitempty"`
	// Additional information about the Condition.
	//
	// This is a general notes/comments entry  for description of the Condition, its diagnosis and prognosis.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
}

type ConditionParticipant struct {
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
	// Distinguishes the type of involvement of the actor in the activities related to the condition.
	Function *CodeableConcept `bson:"function,omitempty" json:"function,omitempty" xml:"function,omitempty"`
	// Indicates who or what participated in the activities related to the condition.
	Actor Reference `bson:"actor" json:"actor" xml:"actor"`
}

type ConditionStage struct {
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
	// A simple summary of the stage such as "Stage 3" or "Early Onset".
	//
	// The determination of the stage is disease-specific, such as cancer, retinopathy of prematurity, kidney diseases, Alzheimer's, or Parkinson disease.
	Summary *CodeableConcept `bson:"summary,omitempty" json:"summary,omitempty" xml:"summary,omitempty"`
	// Reference to a formal record of the evidence on which the staging assessment is based.
	Assessment []Reference `bson:"assessment,omitempty" json:"assessment,omitempty" xml:"assessment,omitempty"`
	// The kind of staging, such as pathological or clinical staging.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
}

func (r Condition) ResourceType() string {
	return "Condition"
}

// MarshalJSON implements json.Marshaler
func (r Condition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Condition
		ResourceType string `json:"resourceType"`
	}{
		Condition:    r,
		ResourceType: r.ResourceType(),
	})
}
