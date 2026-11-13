package fhires

import (
	"encoding/xml"

	decimal "github.com/shopspring/decimal"
)

// An assessment of the likely outcome(s) for a patient or other subject as well as the likelihood of each outcome.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/RiskAssessment
type RiskAssessment struct {
	DomainResource
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

// MarshalJSON implements [json.Marshaler]
func (r RiskAssessment) MarshalJSON() ([]byte, error) {
	type Wrapper RiskAssessment
	return marshalJSON(r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Marshaler].
func (r RiskAssessment) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Wrapper RiskAssessment
	return marshalXML(e, start, r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Unmarshaler].
func (r *RiskAssessment) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper RiskAssessment
	w := Wrapper(*r)
	err := unmarshalXML(d, start, r.ResourceType(), &w)
	*r = RiskAssessment(w)
	return err
}
