package fhires

import decimal "github.com/shopspring/decimal"

// Attachment Type: For referring to data content defined in other formats.
//
// Purpose: Many models need to include data defined in other specifications that is complex and opaque to the healthcare model. This includes documents, media recordings, structured data, etc.
//
// Status: active
//
// http://hl7.org/fhir/StructureDefinition/Attachment
type Attachment struct {
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
	// Identifies the type of the data in the attachment and allows a method to be chosen to interpret or render the data.
	//
	// Includes mime type parameters such as charset where appropriate.
	ContentType *string `bson:"contentType,omitempty" json:"contentType,omitempty" xml:"contentType,omitempty"`
	// The human language of the content. The value can be any valid value according to BCP 47.
	Language *string `bson:"language,omitempty" json:"language,omitempty" xml:"language,omitempty"`
	// The actual data of the attachment - a sequence of bytes, base64 encoded.
	//
	// Comment: The base64-encoded data SHALL be expressed in the same character set as the base resource XML or JSON.
	Data *string `bson:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	// A location where the data can be accessed.
	//
	// Comment: If both data and url are provided, the url SHALL point to the same content as the data contains. Urls may be relative references or may reference transient locations such as a wrapping envelope using cid: though this has ramifications for using signatures. Relative URLs are interpreted relative to the service url, like a resource reference, rather than relative to the resource itself. If a URL is provided, it SHALL resolve to actual data.
	URL *string `bson:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
	// The number of bytes of data that make up this attachment (before base64 encoding, if that is done).
	//
	// Comment: The number of bytes is redundant if the data is provided as a base64binary, but is useful if the data is provided as a url reference.
	Size *int64 `bson:"size,omitempty" json:"size,omitempty,string" xml:"size,omitempty"`
	// The calculated hash of the data using SHA-1. Represented using base64.
	//
	// Comment: The hash is calculated on the data prior to base64 encoding, if the data is based64 encoded. The hash is not intended to support digital signatures. Where protection against malicious threats a digital signature should be considered, see [Provenance.signature](provenance-definitions.html#Provenance.signature) for mechanism to protect a resource with a digital signature.
	Hash *string `bson:"hash,omitempty" json:"hash,omitempty" xml:"hash,omitempty"`
	// A label or set of text to display in place of the data.
	//
	// Comment: May sometimes be derived from the source filename.
	Title *string `bson:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// The date that the attachment was first created.
	Creation *string `bson:"creation,omitempty" json:"creation,omitempty" xml:"creation,omitempty"`
	// Height of the image in pixels (photo/video).
	Height *int `bson:"height,omitempty" json:"height,omitempty" xml:"height,omitempty"`
	// Width of the image in pixels (photo/video).
	Width *int `bson:"width,omitempty" json:"width,omitempty" xml:"width,omitempty"`
	// The number of frames in a photo.
	//
	// This is used with a multi-page fax, or an imaging acquisition context that takes multiple slices in a single image, or an animated gif. If there is more than one frame, this SHALL have a value in order to alert interface software that a multi-frame capable rendering widget is required.
	//
	// Comment: if the number of frames is not supplied, the value may be unknown. Applications should not assume that there is only one frame unless it is explicitly stated.
	Frames *int `bson:"frames,omitempty" json:"frames,omitempty" xml:"frames,omitempty"`
	// The duration of the recording in seconds - for audio and video.
	//
	// Comment: The duration might differ from occurrencePeriod if recording was paused.
	Duration *decimal.Decimal `bson:"duration,omitempty" json:"duration,omitempty" xml:"duration,omitempty"`
	// The number of pages when printed.
	Pages *int `bson:"pages,omitempty" json:"pages,omitempty" xml:"pages,omitempty"`
}
