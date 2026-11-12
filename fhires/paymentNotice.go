package fhires

import "encoding/xml"

// This resource provides the status of the payment for goods and services rendered, and the request and response resource references.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/PaymentNotice
type PaymentNotice struct {
	BaseResource
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

// MarshalJSON implements json.Marshaler
func (r PaymentNotice) MarshalJSON() ([]byte, error) {
	type Wrapper PaymentNotice
	return marshalJSON(r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Unmarshaler].
func (r *PaymentNotice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper PaymentNotice
	w := Wrapper(*r)
	err := unmarshalXML(d, start, &w)
	*r = PaymentNotice(w)
	return err
}
