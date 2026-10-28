package fhires

import v2 "encoding/json/v2"

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
	// The logical id of the resource, as used in the URL for the resource.
	//
	// Once assigned, this value never changes.
	//
	// Comment: Within the context of the FHIR RESTful interactions, the resource has an id except for cases like the create and conditional update. Otherwise, the use of the resouce id depends on the given use case.
	ID *string `bson:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// The metadata about the resource.
	//
	// This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *Meta `bson:"meta,omitempty" json:"meta,omitempty" xml:"meta,omitempty"`
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content.
	//
	// Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	//
	// Comment: Asserting this rule set restricts the content to be only understood by a limited set of trading partners. This inherently limits the usefulness of the data in the long term. However, the existing health eco-system is highly fractured, and not yet ready to define, collect, and exchange data in a generally computable sense. Wherever possible, implementers and/or specification writers should avoid using this element. Often, when used, the URL is a reference to an implementation guide that defines these special rules as part of its narrative along with other profiles, value sets, etc.
	ImplicitRules *string `bson:"implicitRules,omitempty" json:"implicitRules,omitempty" xml:"implicitRules,omitempty"`
	// The base language in which the resource is written.
	//
	// Comment: Language is provided to support indexing and accessibility (typically, services such as text to speech use the language tag). The html language tag in the narrative applies  to the narrative. The language tag on the resource may be used to specify the language of other presentations generated from the data in the resource. Not all the content has to be in the base language. The Resource.language should not be assumed to apply to the narrative automatically. If a language is specified, it should it also be specified on the div element in the html (see rules in HTML5 for information about the relationship between xml:lang and the html lang attribute).
	Language *string `bson:"language,omitempty" json:"language,omitempty" xml:"language,omitempty"`
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

// MarshalJSON implements json.Marshaler
func (r Binary) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		Binary
		ResourceType string `json:"resourceType"`
	}{
		Binary:       r,
		ResourceType: r.ResourceType(),
	})
}
