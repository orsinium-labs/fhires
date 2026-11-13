package fhires

import "encoding/xml"

// Basic is used for handling concepts not yet defined in FHIR, narrative-only resources that don't map to an existing resource, and custom resources not appropriate for inclusion in the FHIR specification.
//
// Purpose: Need some way to safely (without breaking interoperability) allow implementers to exchange content not supported by the initial set of declared resources.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Basic
type Basic struct {
	DomainResource
	// Identifier assigned to the resource for business purposes, outside the context of FHIR.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Identifies the 'type' of resource - equivalent to the resource name for other resources.
	//
	// Comment: Because resource references will only be able to indicate 'Basic', the type of reference will need to be specified in a Profile identified as part of the resource.  Refer to the resource notes section for information on appropriate terminologies for this code.
	//
	// This element is labeled as a modifier because it defines the meaning of the resource and cannot be ignored.
	Code CodeableConcept `bson:"code" json:"code" xml:"code"`
	// Identifies the patient, practitioner, device or any other resource that is the "focus" of this resource.
	//
	// Comment: Optional as not all potential resources will have subjects.  Resources associated with multiple subjects can handle this via extension.
	Subject *Reference `bson:"subject,omitempty" json:"subject,omitempty" xml:"subject,omitempty"`
	// Identifies when the resource was first created.
	Created *string `bson:"created,omitempty" json:"created,omitempty" xml:"created,omitempty"`
	// Indicates who was responsible for creating the resource instance.
	Author *Reference `bson:"author,omitempty" json:"author,omitempty" xml:"author,omitempty"`
}

func (r Basic) ResourceType() string {
	return "Basic"
}

// MarshalJSON implements [json.Marshaler]
func (r Basic) MarshalJSON() ([]byte, error) {
	type Wrapper Basic
	return marshalJSON(r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Marshaler].
func (r Basic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Wrapper Basic
	return marshalXML(e, start, r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Unmarshaler].
func (r *Basic) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper Basic
	w := Wrapper(*r)
	err := unmarshalXML(d, start, r.ResourceType(), &w)
	*r = Basic(w)
	return err
}
