package fhires

import "encoding/xml"

// This resource provides the insurance enrollment details to the insurer regarding a specified coverage.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/EnrollmentRequest
type EnrollmentRequest struct {
	DomainResource
	// The Response business identifier.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The status of the resource instance.
	//
	// Comment: This element is labeled as a modifier because the status contains codes that mark the request as not currently valid.
	Status *FinancialResourceStatusCodes `bson:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// The date when this resource was created.
	Created *string `bson:"created,omitempty" json:"created,omitempty" xml:"created,omitempty"`
	// The Insurer who is target  of the request.
	Insurer *Reference `bson:"insurer,omitempty" json:"insurer,omitempty" xml:"insurer,omitempty"`
	// The practitioner who is responsible for the services rendered to the patient.
	Provider *Reference `bson:"provider,omitempty" json:"provider,omitempty" xml:"provider,omitempty"`
	// Patient Resource.
	Candidate *Reference `bson:"candidate,omitempty" json:"candidate,omitempty" xml:"candidate,omitempty"`
	// Reference to the program or plan identification, underwriter or payor.
	Coverage *Reference `bson:"coverage,omitempty" json:"coverage,omitempty" xml:"coverage,omitempty"`
}

func (r EnrollmentRequest) ResourceType() string {
	return "EnrollmentRequest"
}

// MarshalJSON implements [json.Marshaler]
func (r EnrollmentRequest) MarshalJSON() ([]byte, error) {
	type Wrapper EnrollmentRequest
	return marshalJSON(r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Marshaler].
func (r EnrollmentRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Wrapper EnrollmentRequest
	return marshalXML(e, start, r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Unmarshaler].
func (r *EnrollmentRequest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper EnrollmentRequest
	w := Wrapper(*r)
	err := unmarshalXML(d, start, r.ResourceType(), &w)
	*r = EnrollmentRequest(w)
	return err
}
