package fhires

// Signature Type: A signature along with supporting context.
//
// The signature may be a digital signature that is cryptographic in nature, or some other signature acceptable to the domain. This other signature may be as simple as a graphical image representing a hand-written signature, or a signature ceremony Different signature approaches have different utilities.
//
// Purpose: There are a number of places where content must be signed in healthcare.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Signature
type Signature struct {
	// Unique id for the element within a resource (for internal references).
	//
	// This may be any string value that does not contain spaces.
	ID *string `bson:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element.
	//
	// To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	//
	// Comment: There can be no stigma associated with the use of extensions by any application, project, or standard - regardless of the institution or jurisdiction that uses or defines the extensions.  The use of extensions is what allows the FHIR specification to retain a core level of simplicity for everyone.
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty" xml:"extension,omitempty"`
	// An indication of the reason that the entity signed this document.
	//
	// This may be explicitly included as part of the signature information and can be used when determining accountability for various actions concerning the document.
	//
	// Comment: Examples include attesting to: authorship, correct transcription, and witness of specific event. Also known as a &quot;Commitment Type Indication&quot;.
	Type []Coding `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// When the digital signature was signed.
	//
	// Comment: This should agree with the information in the signature.
	When *string `bson:"when,omitempty" json:"when,omitempty" xml:"when,omitempty"`
	// A reference to an application-usable description of the identity that signed  (e.g.
	//
	// the signature used their private key).
	//
	// Comment: This should agree with the information in the signature.
	Who *Reference `bson:"who,omitempty" json:"who,omitempty" xml:"who,omitempty"`
	// A reference to an application-usable description of the identity that is represented by the signature.
	//
	// Comment: The party that can't sign. For example a child.
	OnBehalfOf *Reference `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty" xml:"onBehalfOf,omitempty"`
	// A mime type that indicates the technical format of the target resources signed by the signature.
	//
	// Comment: "xml", "json" and "ttl" are allowed, which describe the simple encodings described in the specification (and imply appropriate bundle support). Otherwise, mime types are legal here.
	TargetFormat *string `bson:"targetFormat,omitempty" json:"targetFormat,omitempty" xml:"targetFormat,omitempty"`
	// A mime type that indicates the technical format of the signature.
	//
	// Important mime types are application/signature+xml for X ML DigSig, application/jose for JWS, and image/* for a graphical image of a signature, etc.
	SigFormat *string `bson:"sigFormat,omitempty" json:"sigFormat,omitempty" xml:"sigFormat,omitempty"`
	// The base64 encoding of the Signature content.
	//
	// When signature is not recorded electronically this element would be empty.
	//
	// Comment: Where the signature type is an XML DigSig, the signed content is a FHIR Resource(s), the signature is of the XML form of the Resource(s) using  XML-Signature (XMLDIG) "Detached Signature" form.
	Data *string `bson:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
}
