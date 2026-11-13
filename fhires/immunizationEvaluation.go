package fhires

import "encoding/xml"

// Describes a comparison of an immunization event against published recommendations to determine if the administration is "valid" in relation to those  recommendations.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/ImmunizationEvaluation
type ImmunizationEvaluation struct {
	DomainResource
	// A unique identifier assigned to this immunization evaluation record.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Indicates the current status of the evaluation of the vaccination administration event.
	Status ImmunizationEvaluationStatusCodes `bson:"status" json:"status" xml:"status"`
	// The individual for whom the evaluation is being done.
	Patient Reference `bson:"patient" json:"patient" xml:"patient"`
	// The date the evaluation of the vaccine administration event was performed.
	Date *string `bson:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// Indicates the authority who published the protocol (e.g. ACIP).
	Authority *Reference `bson:"authority,omitempty" json:"authority,omitempty" xml:"authority,omitempty"`
	// The vaccine preventable disease the dose is being evaluated against.
	TargetDisease CodeableConcept `bson:"targetDisease" json:"targetDisease" xml:"targetDisease"`
	// The vaccine administration event being evaluated.
	ImmunizationEvent Reference `bson:"immunizationEvent" json:"immunizationEvent" xml:"immunizationEvent"`
	// Indicates if the dose is valid or not valid with respect to the published recommendations.
	DoseStatus CodeableConcept `bson:"doseStatus" json:"doseStatus" xml:"doseStatus"`
	// Provides an explanation as to why the vaccine administration event is valid or not relative to the published recommendations.
	DoseStatusReason []CodeableConcept `bson:"doseStatusReason,omitempty" json:"doseStatusReason,omitempty" xml:"doseStatusReason,omitempty"`
	// Additional information about the evaluation.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// One possible path to achieve presumed immunity against a disease - within the context of an authority.
	Series *string `bson:"series,omitempty" json:"series,omitempty" xml:"series,omitempty"`
	// Nominal position in a series as determined by the outcome of the evaluation process.
	//
	// Comment: The use of an integer is preferred if known. A string should only be used in cases where an integer is not available (such as when documenting a recurring booster dose).
	DoseNumber *string `bson:"doseNumber,omitempty" json:"doseNumber,omitempty" xml:"doseNumber,omitempty"`
	// The recommended number of doses to achieve immunity as determined by the outcome of the evaluation process.
	//
	// Comment: The use of an integer is preferred if known.
	SeriesDoses *string `bson:"seriesDoses,omitempty" json:"seriesDoses,omitempty" xml:"seriesDoses,omitempty"`
}

func (r ImmunizationEvaluation) ResourceType() string {
	return "ImmunizationEvaluation"
}

// MarshalJSON implements [json.Marshaler]
func (r ImmunizationEvaluation) MarshalJSON() ([]byte, error) {
	type Wrapper ImmunizationEvaluation
	return marshalJSON(r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Marshaler].
func (r ImmunizationEvaluation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Wrapper ImmunizationEvaluation
	return marshalXML(e, start, r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Unmarshaler].
func (r *ImmunizationEvaluation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper ImmunizationEvaluation
	w := Wrapper(*r)
	err := unmarshalXML(d, start, r.ResourceType(), &w)
	*r = ImmunizationEvaluation(w)
	return err
}
