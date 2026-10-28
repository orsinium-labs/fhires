package fhires

import (
	"encoding/json/v2"

	decimal "github.com/shopspring/decimal"
)

// An assessment of the likely outcome(s) for a patient or other subject as well as the likelihood of each outcome.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/RiskAssessment
type RiskAssessment struct {
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
	// Business identifier assigned to the risk assessment.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// A reference to the request that is fulfilled by this risk assessment.
	BasedOn *Reference `bson:"basedOn,omitempty" json:"basedOn,omitempty" xml:"basedOn,omitempty"`
	// A reference to a resource that this risk assessment is part of, such as a Procedure.
	Parent *Reference `bson:"parent,omitempty" json:"parent,omitempty" xml:"parent,omitempty"`
	// The status of the RiskAssessment, using the same statuses as an Observation.
	Status ObservationStatus `bson:"status" json:"status" xml:"status"`
	// The algorithm, process or mechanism used to evaluate the risk.
	Method *CodeableConcept `bson:"method,omitempty" json:"method,omitempty" xml:"method,omitempty"`
	// The type of the risk assessment performed.
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// The patient or group the risk assessment applies to.
	Subject Reference `bson:"subject" json:"subject" xml:"subject"`
	// The encounter where the assessment was performed.
	Encounter *Reference `bson:"encounter,omitempty" json:"encounter,omitempty" xml:"encounter,omitempty"`
	// The date (and possibly time) the risk assessment was performed.
	OccurrenceDateTime *string `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty" xml:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *Period `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty" xml:"occurrencePeriod,omitempty"`
	// For assessments or prognosis specific to a particular condition, indicates the condition being assessed.
	Condition *Reference `bson:"condition,omitempty" json:"condition,omitempty" xml:"condition,omitempty"`
	// The provider, patient, related person, or software application that performed the assessment.
	Performer *Reference `bson:"performer,omitempty" json:"performer,omitempty" xml:"performer,omitempty"`
	// The reason the risk assessment was performed.
	Reason []CodeableReference `bson:"reason,omitempty" json:"reason,omitempty" xml:"reason,omitempty"`
	// Indicates the source data considered as part of the assessment (for example, FamilyHistory, Observations, Procedures, Conditions, etc.).
	Basis []Reference `bson:"basis,omitempty" json:"basis,omitempty" xml:"basis,omitempty"`
	// Describes the expected outcome for the subject.
	//
	// Comment: Multiple repetitions can be used to identify the same type of outcome in different timeframes as well as different types of outcomes.
	Prediction []RiskAssessmentPrediction `bson:"prediction,omitempty" json:"prediction,omitempty" xml:"prediction,omitempty"`
	// A description of the steps that might be taken to reduce the identified risk(s).
	Mitigation *string `bson:"mitigation,omitempty" json:"mitigation,omitempty" xml:"mitigation,omitempty"`
	// Additional comments about the risk assessment.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
}

type RiskAssessmentPrediction struct {
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
	// One of the potential outcomes for the patient (e.g.
	//
	// remission, death,  a particular condition).
	Outcome *CodeableConcept `bson:"outcome,omitempty" json:"outcome,omitempty" xml:"outcome,omitempty"`
	// Indicates how likely the outcome is (in the specified timeframe).
	//
	// Comment: If range is used, it represents the lower and upper bounds of certainty; e.g. 40-60%  Decimal values are expressed as percentages as well (max = 100).
	ProbabilityDecimal *decimal.Decimal `bson:"probabilityDecimal,omitempty" json:"probabilityDecimal,omitempty" xml:"probabilityDecimal,omitempty"`
	ProbabilityRange   *Range           `bson:"probabilityRange,omitempty" json:"probabilityRange,omitempty" xml:"probabilityRange,omitempty"`
	// Indicates how likely the outcome is (in the specified timeframe), expressed as a qualitative value (e.g.
	//
	// low, medium, or high).
	QualitativeRisk *CodeableConcept `bson:"qualitativeRisk,omitempty" json:"qualitativeRisk,omitempty" xml:"qualitativeRisk,omitempty"`
	// Indicates the risk for this particular subject (with their specific characteristics) divided by the risk of the population in general.
	//
	//  (Numbers greater than 1 = higher risk than the population, numbers less than 1 = lower risk.).
	RelativeRisk *decimal.Decimal `bson:"relativeRisk,omitempty" json:"relativeRisk,omitempty" xml:"relativeRisk,omitempty"`
	// Indicates the period of time or age range of the subject to which the specified probability applies.
	//
	// Comment: If not specified, the risk applies "over the subject's lifespan".
	WhenPeriod *Period `bson:"whenPeriod,omitempty" json:"whenPeriod,omitempty" xml:"whenPeriod,omitempty"`
	WhenRange  *Range  `bson:"whenRange,omitempty" json:"whenRange,omitempty" xml:"whenRange,omitempty"`
	// Additional information explaining the basis for the prediction.
	Rationale *string `bson:"rationale,omitempty" json:"rationale,omitempty" xml:"rationale,omitempty"`
}

func (r RiskAssessment) ResourceType() string {
	return "RiskAssessment"
}

// MarshalJSON implements json.Marshaler
func (r RiskAssessment) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		RiskAssessment
		ResourceType string `json:"resourceType"`
	}{
		ResourceType:   r.ResourceType(),
		RiskAssessment: r,
	})
}
