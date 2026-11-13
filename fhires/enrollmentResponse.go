package fhires

import "encoding/xml"

// This resource provides enrollment and plan details from the processing of an EnrollmentRequest resource.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/EnrollmentResponse
type EnrollmentResponse struct {
	DomainResource
	// The Response business identifier.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The status of the resource instance.
	//
	// Comment: This element is labeled as a modifier because the status contains codes that mark the response as not currently valid.
	Status *FinancialResourceStatusCodes `bson:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Original request resource reference.
	Request *Reference `bson:"request,omitempty" json:"request,omitempty" xml:"request,omitempty"`
	// Processing status: error, complete.
	Outcome *EnrollmentOutcome `bson:"outcome,omitempty" json:"outcome,omitempty" xml:"outcome,omitempty"`
	// A description of the status of the adjudication.
	Disposition *string `bson:"disposition,omitempty" json:"disposition,omitempty" xml:"disposition,omitempty"`
	// The date when the enclosed suite of services were performed or completed.
	Created *string `bson:"created,omitempty" json:"created,omitempty" xml:"created,omitempty"`
	// The Insurer who produced this adjudicated response.
	Organization *Reference `bson:"organization,omitempty" json:"organization,omitempty" xml:"organization,omitempty"`
	// The practitioner who is responsible for the services rendered to the patient.
	RequestProvider *Reference `bson:"requestProvider,omitempty" json:"requestProvider,omitempty" xml:"requestProvider,omitempty"`
}

func (r EnrollmentResponse) ResourceType() string {
	return "EnrollmentResponse"
}

// MarshalJSON implements [json.Marshaler]
func (r EnrollmentResponse) MarshalJSON() ([]byte, error) {
	type Wrapper EnrollmentResponse
	return marshalJSON(r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Marshaler].
func (r EnrollmentResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Wrapper EnrollmentResponse
	return marshalXML(e, start, r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Unmarshaler].
func (r *EnrollmentResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper EnrollmentResponse
	w := Wrapper(*r)
	err := unmarshalXML(d, start, r.ResourceType(), &w)
	*r = EnrollmentResponse(w)
	return err
}
