package fhires

import "encoding/xml"

// This resource describes a product or service that is available through a program and includes the conditions and constraints of availability.
//
//	All of the information in this resource is specific to the inclusion of the item in the formulary and is not inherent to the item itself.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/FormularyItem
type FormularyItem struct {
	DomainResource
	// Business identifier for this formulary item.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// A code (or set of codes) that specify the product or service that is identified by this formulary item.
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// The validity about the information of the formulary item and not of the underlying product or service itself.
	//
	// Comment: This status is intended to identify if the formulary item in a local system is in active use within the formulary.  For example, a formulary might include a medicinal product and the status relates to its inclusion.  At some point in the future, if the status of the formulary item is changed to "inactive" that means that it is no longer valid in the formulary.  But the product itself might still be in active use and might be included in other formularies.
	Status *FormularyItemStatusCodes `bson:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
}

func (r FormularyItem) ResourceType() string {
	return "FormularyItem"
}

// MarshalJSON implements [json.Marshaler]
func (r FormularyItem) MarshalJSON() ([]byte, error) {
	type Wrapper FormularyItem
	return marshalJSON(r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Marshaler].
func (r FormularyItem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Wrapper FormularyItem
	return marshalXML(e, start, r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Unmarshaler].
func (r *FormularyItem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper FormularyItem
	w := Wrapper(*r)
	err := unmarshalXML(d, start, r.ResourceType(), &w)
	*r = FormularyItem(w)
	return err
}
