package fhires

import v2 "encoding/json/v2"

// A record of a clinical assessment performed to determine what problem(s) may affect the patient and before planning the treatments or management strategies that are best to manage a patient's condition.
//
// Assessments are often 1:1 with a clinical consultation / encounter,  but this varies greatly depending on the clinical workflow. This resource is called "ClinicalImpression" rather than "ClinicalAssessment" to avoid confusion with the recording of assessment tools such as Apgar score.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/ClinicalImpression
type ClinicalImpression struct {
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
	// Business identifiers assigned to this clinical impression by the performer or other systems which remain constant as the resource is updated and propagates from server to server.
	//
	// Comment: This is a business identifier, not a resource identifier (see [discussion](resource.html#identifiers)).  It is best practice for the identifier to only appear on a single resource instance, however business practices may occasionally dictate that multiple resource instances with the same identifier can exist - possibly even with different resource types.  For example, multiple Patient and a Person resource instance might share the same social insurance number.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Identifies the workflow status of the assessment.
	//
	// Comment: This element is labeled as a modifier because the status contains the code entered-in-error that marks the clinical impression as not currently valid.
	Status EventStatus `bson:"status" json:"status" xml:"status"`
	// Captures the reason for the current state of the ClinicalImpression.
	//
	// Comment: This is generally only used for "exception" statuses such as "not-done", "on-hold" or "stopped".
	StatusReason *CodeableConcept `bson:"statusReason,omitempty" json:"statusReason,omitempty" xml:"statusReason,omitempty"`
	// A summary of the context and/or cause of the assessment - why / where it was performed, and what patient events/status prompted it.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// The patient or group of individuals assessed as part of this record.
	Subject Reference `bson:"subject" json:"subject" xml:"subject"`
	// The Encounter during which this ClinicalImpression was created or to which the creation of this record is tightly associated.
	//
	// Comment: This will typically be the encounter the event occurred within, but some activities may be initiated prior to or after the official completion of an encounter but still be tied to the context of the encounter.
	Encounter *Reference `bson:"encounter,omitempty" json:"encounter,omitempty" xml:"encounter,omitempty"`
	// The point in time or period over which the subject was assessed.
	//
	// Comment: This SHOULD be accurate to at least the minute, though some assessments only have a known date.
	EffectiveDateTime *string `bson:"effectiveDateTime,omitempty" json:"effectiveDateTime,omitempty" xml:"effectiveDateTime,omitempty"`
	EffectivePeriod   *Period `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty" xml:"effectivePeriod,omitempty"`
	// Indicates when the documentation of the assessment was complete.
	Date *string `bson:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// The clinician performing the assessment.
	Performer *Reference `bson:"performer,omitempty" json:"performer,omitempty" xml:"performer,omitempty"`
	// A reference to the last assessment that was conducted on this patient.
	//
	// Assessments are often/usually ongoing in nature; a care provider (practitioner or team) will make new assessments on an ongoing basis as new data arises or the patient's conditions changes.
	//
	// Comment: It is always likely that multiple previous assessments exist for a patient. The point of quoting a previous assessment is that this assessment is relative to it (see resolved).
	Previous *Reference `bson:"previous,omitempty" json:"previous,omitempty" xml:"previous,omitempty"`
	// A list of the relevant problems/conditions for a patient.
	//
	// Comment: e.g. The patient is a pregnant, has congestive heart failure, has an ‎Adenocarcinoma, and is allergic to penicillin.
	Problem []Reference `bson:"problem,omitempty" json:"problem,omitempty" xml:"problem,omitempty"`
	// Change in the status/pattern of a subject's condition since previously assessed, such as worsening, improving, or no change.
	//
	//  It is a subjective assessment of the direction of the change.
	ChangePattern *CodeableConcept `bson:"changePattern,omitempty" json:"changePattern,omitempty" xml:"changePattern,omitempty"`
	// Reference to a specific published clinical protocol that was followed during this assessment, and/or that provides evidence in support of the diagnosis.
	Protocol []string `bson:"protocol,omitempty" json:"protocol,omitempty" xml:"protocol,omitempty"`
	// A text summary of the investigations and the diagnosis.
	Summary *string `bson:"summary,omitempty" json:"summary,omitempty" xml:"summary,omitempty"`
	// Specific findings or diagnoses that were considered likely or relevant to ongoing treatment.
	Finding []ClinicalImpressionFinding `bson:"finding,omitempty" json:"finding,omitempty" xml:"finding,omitempty"`
	// Estimate of likely outcome.
	PrognosisCodeableConcept []CodeableConcept `bson:"prognosisCodeableConcept,omitempty" json:"prognosisCodeableConcept,omitempty" xml:"prognosisCodeableConcept,omitempty"`
	// RiskAssessment expressing likely outcome.
	PrognosisReference []Reference `bson:"prognosisReference,omitempty" json:"prognosisReference,omitempty" xml:"prognosisReference,omitempty"`
	// Information supporting the clinical impression, which can contain investigation results.
	SupportingInfo []Reference `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty" xml:"supportingInfo,omitempty"`
	// Commentary about the impression, typically recorded after the impression itself was made, though supplemental notes by the original author could also appear.
	//
	// Comment: Don't use this element for content that should more properly appear as one of the specific elements of the impression.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
}

type ClinicalImpressionFinding struct {
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
	// Specific text, code or reference for finding or diagnosis, which may include ruled-out or resolved conditions.
	Item *CodeableReference `bson:"item,omitempty" json:"item,omitempty" xml:"item,omitempty"`
	// Which investigations support finding or diagnosis.
	Basis *string `bson:"basis,omitempty" json:"basis,omitempty" xml:"basis,omitempty"`
}

func (r ClinicalImpression) ResourceType() string {
	return "ClinicalImpression"
}

// MarshalJSON implements json.Marshaler
func (r ClinicalImpression) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		ClinicalImpression
		ResourceType string `json:"resourceType"`
	}{
		ClinicalImpression: r,
		ResourceType:       r.ResourceType(),
	})
}
