package fhires

import "encoding/xml"

// The technical details of an endpoint that can be used for electronic services, such as for web services providing XDS.b, a REST endpoint for another FHIR server, or a s/Mime email address.
//
// This may include any security context information.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Endpoint
type Endpoint struct {
	DomainResource
	// Identifier for the organization that is used to identify the endpoint across multiple disparate systems.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The endpoint status represents the general expected availability of an endpoint.
	//
	// Comment: This element is labeled as a modifier because the status contains codes that mark the endpoint as not currently valid. Temporary downtimes or other unexpected short-term changes in availability would not be represented in this property.
	Status EndpointStatus `bson:"status" json:"status" xml:"status"`
	// A coded value that represents the technical details of the usage of this endpoint, such as what WSDLs should be used in what way.
	//
	// (e.g. XDS.b/DICOM/cds-hook).
	//
	// Comment: For additional connectivity details for the protocol, extensions will be used at this point, as in the XDS example. If there are multiple payload types or mimetypes they are all applicable for all connection types, and all have the same status.
	ConnectionType []CodeableConcept `bson:"connectionType" json:"connectionType" xml:"connectionType"`
	// A friendly name that this endpoint can be referred to with.
	Name *string `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// The description of the endpoint and what it is for (typically used as supplemental information in an endpoint directory describing its usage/purpose).
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// The type of environment(s) exposed at this endpoint (dev, prod, test, etc.).
	EnvironmentType []CodeableConcept `bson:"environmentType,omitempty" json:"environmentType,omitempty" xml:"environmentType,omitempty"`
	// The organization that manages this endpoint (even if technically another organization is hosting this in the cloud, it is the organization associated with the data).
	//
	// Comment: This property is not typically used when searching for Endpoint resources for usage. The typical usage is via the reference from an applicable Organization/Location/Practitioner resource, which is where the context is provided. Multiple Locations may reference a single endpoint, and don't have to be within the same organization resource, but most likely within the same organizational hierarchy.
	ManagingOrganization *Reference `bson:"managingOrganization,omitempty" json:"managingOrganization,omitempty" xml:"managingOrganization,omitempty"`
	// Contact details for a human to contact about the endpoint.
	//
	// The primary use of this for system administrator troubleshooting.
	Contact []ContactPoint `bson:"contact,omitempty" json:"contact,omitempty" xml:"contact,omitempty"`
	// The interval during which the endpoint is expected to be operational.
	//
	// Comment: Temporary downtimes or other short-term changes in availability would not be represented in this property.
	Period *Period `bson:"period,omitempty" json:"period,omitempty" xml:"period,omitempty"`
	// The set of payloads that are provided/available at this endpoint.
	//
	// Comment: Note that not all mimetypes or types will be listed under the one endpoint resource, there may be multiple instances that information for cases where other header data such as the endpoint address, active status/period etc. is different.
	Payload []EndpointPayload `bson:"payload,omitempty" json:"payload,omitempty" xml:"payload,omitempty"`
	// The uri that describes the actual end-point to connect to.
	//
	// Comment: For rest-hook, and websocket, the end-point must be an http: or https: URL; for email, a mailto: url, for sms, a tel: url, and for message the endpoint can be in any form of url the server understands (usually, http: or mllp:). The URI is allowed to be relative; in which case, it is relative to the server end-point (since there may be more than one, clients should avoid using relative URIs)
	//
	// This address will be to the service base, without any parameters, or sub-services or resources tacked on.
	//
	// E.g. for a WADO-RS endpoint, the url should be "https://pacs.hospital.org/wado-rs"
	//
	// and not "https://pacs.hospital.org/wado-rs/studies/1.2.250.1.59.40211.12345678.678910/series/1.2.250.1.59.40211.789001276.14556172.67789/instances/...".
	Address string `bson:"address" json:"address" xml:"address"`
	// Additional headers / information to send as part of the notification.
	//
	// Comment: Exactly what these mean depends on the channel type. The can convey additional information to the recipient and/or meet security requirements.
	Header []string `bson:"header,omitempty" json:"header,omitempty" xml:"header,omitempty"`
}

type EndpointPayload struct {
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
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants.
	//
	// Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	//
	// Comment: There can be no stigma associated with the use of extensions by any application, project, or standard - regardless of the institution or jurisdiction that uses or defines the extensions.  The use of extensions is what allows the FHIR specification to retain a core level of simplicity for everyone.
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty" xml:"modifierExtension,omitempty"`
	// The payload type describes the acceptable content that can be communicated on the endpoint.
	//
	// Comment: The mimeType describes the serialization format of the data, where the payload.type indicates the specific document/schema that is being transferred; e.g. DischargeSummary or CarePlan.
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The mime type to send the payload in - e.g.
	//
	// application/fhir+xml, application/fhir+json. If the mime type is not specified, then the sender could send any content (including no content depending on the connectionType).
	//
	// Comment: Sending the payload has obvious security consequences. The server is responsible for ensuring that the content is appropriately secured.
	MimeType []string `bson:"mimeType,omitempty" json:"mimeType,omitempty" xml:"mimeType,omitempty"`
}

func (r Endpoint) ResourceType() string {
	return "Endpoint"
}

// MarshalJSON implements [json.Marshaler]
func (r Endpoint) MarshalJSON() ([]byte, error) {
	type Wrapper Endpoint
	return marshalJSON(r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Marshaler].
func (r Endpoint) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Wrapper Endpoint
	return marshalXML(e, start, r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Unmarshaler].
func (r *Endpoint) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper Endpoint
	w := Wrapper(*r)
	err := unmarshalXML(d, start, r.ResourceType(), &w)
	*r = Endpoint(w)
	return err
}
