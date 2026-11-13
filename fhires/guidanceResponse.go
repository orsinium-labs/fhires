package fhires

import "encoding/xml"

// A guidance response is the formal response to a guidance request, including any output parameters returned by the evaluation, as well as the description of any proposed actions to be taken.
//
// Purpose: The GuidanceResponse resource supports recording the results of decision support interactions, reportability determination for public health, as well as the communication of additional data requirements for subsequent interactions.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/GuidanceResponse
type GuidanceResponse struct {
	DomainResource
	// The identifier of the request associated with this response.
	//
	// If an identifier was given as part of the request, it will be reproduced here to enable the requester to more easily identify the response in a multi-request scenario.
	RequestIdentifier *Identifier `bson:"requestIdentifier,omitempty" json:"requestIdentifier,omitempty" xml:"requestIdentifier,omitempty"`
	// Allows a service to provide  unique, business identifiers for the response.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// An identifier, CodeableConcept or canonical reference to the guidance that was requested.
	ModuleUri             string          `bson:"moduleUri" json:"moduleUri" xml:"moduleUri"`
	ModuleCanonical       string          `bson:"moduleCanonical" json:"moduleCanonical" xml:"moduleCanonical"`
	ModuleCodeableConcept CodeableConcept `bson:"moduleCodeableConcept" json:"moduleCodeableConcept" xml:"moduleCodeableConcept"`
	// The status of the response.
	//
	// If the evaluation is completed successfully, the status will indicate success. However, in order to complete the evaluation, the engine may require more information. In this case, the status will be data-required, and the response will contain a description of the additional required information. If the evaluation completed successfully, but the engine determines that a potentially more accurate response could be provided if more data was available, the status will be data-requested, and the response will contain a description of the additional requested information.
	//
	// Comment: This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid.
	Status GuidanceResponseStatus `bson:"status" json:"status" xml:"status"`
	// The patient for which the request was processed.
	Subject *Reference `bson:"subject,omitempty" json:"subject,omitempty" xml:"subject,omitempty"`
	// The encounter during which this response was created or to which the creation of this record is tightly associated.
	//
	// Comment: This will typically be the encounter the event occurred within, but some activities may be initiated prior to or after the official copmletion of an encounter but still be tied to the context of the encounter.
	Encounter *Reference `bson:"encounter,omitempty" json:"encounter,omitempty" xml:"encounter,omitempty"`
	// Indicates when the guidance response was processed.
	OccurrenceDateTime *string `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty" xml:"occurrenceDateTime,omitempty"`
	// Provides a reference to the device that performed the guidance.
	Performer *Reference `bson:"performer,omitempty" json:"performer,omitempty" xml:"performer,omitempty"`
	// Describes the reason for the guidance response in coded or textual form, or Indicates the reason the request was initiated.
	//
	// This is typically provided as a parameter to the evaluation and echoed by the service, although for some use cases, such as subscription- or event-based scenarios, it may provide an indication of the cause for the response.
	//
	// Comment: Although this reference can be used to point to any resource, it is typically expected to refer to subject-specific data appropriate to the subject of the guidance. For example, patient-based decision support would be expected to reference patient-level data. In addition, implementations should provide as much detail as possible by using the [http://hl7.org/fhir/StructureDefinition/targetElement](http://hl7.org/fhir/extensions/StructureDefinition-targetElement.html) and [http://hl7.org/fhir/StructureDefinition/targetPath](http://hl7.org/fhir/extensions/StructureDefinition-targetPath.html) extensions to indicate the specific elements relevant to providing the reason for the guidance.
	Reason []CodeableReference `bson:"reason,omitempty" json:"reason,omitempty" xml:"reason,omitempty"`
	// Provides a mechanism to communicate additional information about the response.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// Messages resulting from the evaluation of the artifact or artifacts.
	//
	// As part of evaluating the request, the engine may produce informational or warning messages. These messages will be provided by this element.
	EvaluationMessage *Reference `bson:"evaluationMessage,omitempty" json:"evaluationMessage,omitempty" xml:"evaluationMessage,omitempty"`
	// The output parameters of the evaluation, if any.
	//
	// Many modules will result in the return of specific resources such as procedure or communication requests that are returned as part of the operation result. However, modules may define specific outputs that would be returned as the result of the evaluation, and these would be returned in this element.
	OutputParameters *Reference `bson:"outputParameters,omitempty" json:"outputParameters,omitempty" xml:"outputParameters,omitempty"`
	// The actions, if any, produced by the evaluation of the artifact.
	Result []Reference `bson:"result,omitempty" json:"result,omitempty" xml:"result,omitempty"`
	// If the evaluation could not be completed due to lack of information, or additional information would potentially result in a more accurate response, this element will a description of the data required in order to proceed with the evaluation.
	//
	// A subsequent request to the service should include this data.
	DataRequirement []DataRequirement `bson:"dataRequirement,omitempty" json:"dataRequirement,omitempty" xml:"dataRequirement,omitempty"`
}

func (r GuidanceResponse) ResourceType() string {
	return "GuidanceResponse"
}

// MarshalJSON implements [json.Marshaler]
func (r GuidanceResponse) MarshalJSON() ([]byte, error) {
	type Wrapper GuidanceResponse
	return marshalJSON(r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Marshaler].
func (r GuidanceResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Wrapper GuidanceResponse
	return marshalXML(e, start, r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Unmarshaler].
func (r *GuidanceResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper GuidanceResponse
	w := Wrapper(*r)
	err := unmarshalXML(d, start, r.ResourceType(), &w)
	*r = GuidanceResponse(w)
	return err
}
