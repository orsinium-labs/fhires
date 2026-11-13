package fhires

import "encoding/xml"

// This resource provides the status of the payment for goods and services rendered, and the request and response resource references.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/PaymentNotice
type PaymentNotice struct {
	DomainResource
	// A unique identifier assigned to this payment notice.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The status of the resource instance.
	//
	// Comment: This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid.
	Status FinancialResourceStatusCodes `bson:"status" json:"status" xml:"status"`
	// Reference of resource for which payment is being made.
	Request *Reference `bson:"request,omitempty" json:"request,omitempty" xml:"request,omitempty"`
	// Reference of response to resource for which payment is being made.
	Response *Reference `bson:"response,omitempty" json:"response,omitempty" xml:"response,omitempty"`
	// The date when this resource was created.
	Created string `bson:"created" json:"created" xml:"created"`
	// The party who reports the payment notice.
	Reporter *Reference `bson:"reporter,omitempty" json:"reporter,omitempty" xml:"reporter,omitempty"`
	// A reference to the payment which is the subject of this notice.
	Payment *Reference `bson:"payment,omitempty" json:"payment,omitempty" xml:"payment,omitempty"`
	// The date when the above payment action occurred.
	PaymentDate *string `bson:"paymentDate,omitempty" json:"paymentDate,omitempty" xml:"paymentDate,omitempty"`
	// The party who will receive or has received payment that is the subject of this notification.
	Payee *Reference `bson:"payee,omitempty" json:"payee,omitempty" xml:"payee,omitempty"`
	// The party who is notified of the payment status.
	Recipient Reference `bson:"recipient" json:"recipient" xml:"recipient"`
	// The amount sent to the payee.
	Amount Money `bson:"amount" json:"amount" xml:"amount"`
	// A code indicating whether payment has been sent or cleared.
	//
	// Comment: Typically paid: payment sent, cleared: payment received.
	PaymentStatus *CodeableConcept `bson:"paymentStatus,omitempty" json:"paymentStatus,omitempty" xml:"paymentStatus,omitempty"`
}

func (r PaymentNotice) ResourceType() string {
	return "PaymentNotice"
}

// MarshalJSON implements [json.Marshaler]
func (r PaymentNotice) MarshalJSON() ([]byte, error) {
	type Wrapper PaymentNotice
	return marshalJSON(r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Marshaler].
func (r PaymentNotice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Wrapper PaymentNotice
	return marshalXML(e, start, r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Unmarshaler].
func (r *PaymentNotice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper PaymentNotice
	w := Wrapper(*r)
	err := unmarshalXML(d, start, r.ResourceType(), &w)
	*r = PaymentNotice(w)
	return err
}
