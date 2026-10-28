package fhires

import v2 "encoding/json/v2"

// The technical details of an endpoint that can be used for electronic services, such as for web services providing XDS.b, a REST endpoint for another FHIR server, or a s/Mime email address.
//
// This may include any security context information.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Endpoint
type Endpoint struct {
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

// MarshalJSON implements json.Marshaler
func (r Endpoint) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		Endpoint
		ResourceType string `json:"resourceType"`
	}{
		Endpoint:     r,
		ResourceType: r.ResourceType(),
	})
}
