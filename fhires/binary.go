package fhires

import "encoding/xml"

// A resource that represents the data of a single raw artifact as digital content accessible in its native format.
//
//	A Binary resource can contain any content, whether text, image, pdf, zip archive, etc.
//
// Purpose: There are situations where it is useful or required to handle pure binary content using the same framework as other resources.
//
// Status: active
//
// http://hl7.org/fhir/StructureDefinition/Binary
type Binary struct {
	BaseResource
	// MimeType of the binary content represented as a standard MimeType (BCP 13).
	ContentType string `bson:"contentType" json:"contentType" xml:"contentType"`
	// This element identifies another resource that can be used as a proxy of the security sensitivity to use when deciding and enforcing access control rules for the Binary resource.
	//
	// Given that the Binary resource contains very few elements that can be used to determine the sensitivity of the data and relationships to individuals, the referenced resource stands in as a proxy equivalent for this purpose. This referenced resource may be related to the Binary (e.g. DocumentReference), or may be some non-related Resource purely as a security proxy. E.g. to identify that the binary resource relates to a patient, and access should only be granted to applications that have access to the patient.
	//
	// Comment: Very often, a server will also know of a resource that references the binary, and can automatically apply the appropriate access rules based on that reference. However, there are some circumstances where this is not appropriate, e.g. the binary is uploaded directly to the server without any linking resource, the binary is referred to from multiple different resources, and/or the binary is content such as an application logo that has less protection than any of the resources that reference it.
	SecurityContext *Reference `bson:"securityContext,omitempty" json:"securityContext,omitempty" xml:"securityContext,omitempty"`
	// The actual content, base64 encoded.
	//
	// Comment: If the content type is itself base64 encoding, then this will be base64 encoded twice - what is created by un-base64ing the content must be the specified content type.
	Data *string `bson:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
}

func (r Binary) ResourceType() string {
	return "Binary"
}

// MarshalJSON implements [json.Marshaler]
func (r Binary) MarshalJSON() ([]byte, error) {
	type Wrapper Binary
	return marshalJSON(r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Marshaler].
func (r Binary) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Wrapper Binary
	return marshalXML(e, start, r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Unmarshaler].
func (r *Binary) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper Binary
	w := Wrapper(*r)
	err := unmarshalXML(d, start, r.ResourceType(), &w)
	*r = Binary(w)
	return err
}
