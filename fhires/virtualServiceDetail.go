package fhires

// VirtualServiceDetail Type: Virtual Service Contact Details.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/VirtualServiceDetail
type VirtualServiceDetail struct {
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
	// The type of virtual service to connect to (i.e.
	//
	// Teams, Zoom, Specific VMR technology, WhatsApp).
	//
	// Comment: As this type defines what the virtual service is, it then defines what address and additional information might be relevant.
	//
	// Some profiles could be used for specific types to define what the other values could/should be.
	ChannelType *Coding `bson:"channelType,omitempty" json:"channelType,omitempty" xml:"channelType,omitempty"`
	// What address or number needs to be used for a user to connect to the virtual service to join.
	//
	// The channelType informs as to which datatype is appropriate to use (requires knowledge of the specific type).
	AddressUURL                  *string                `bson:"addressUrl,omitempty" json:"addressUrl,omitempty" xml:"addressUrl,omitempty"`
	AddressString                *string                `bson:"addressString,omitempty" json:"addressString,omitempty" xml:"addressString,omitempty"`
	AddressContactPoint          *ContactPoint          `bson:"addressContactPoint,omitempty" json:"addressContactPoint,omitempty" xml:"addressContactPoint,omitempty"`
	AddressExtendedContactDetail *ExtendedContactDetail `bson:"addressExtendedContactDetail,omitempty" json:"addressExtendedContactDetail,omitempty" xml:"addressExtendedContactDetail,omitempty"`
	// Address to see alternative connection details.
	//
	// Comment: This web address can be used to provide additional details on the call, such as alternative/regional call in numbers, or other associated services.
	AdditionalInfo []string `bson:"additionalInfo,omitempty" json:"additionalInfo,omitempty" xml:"additionalInfo,omitempty"`
	// Maximum number of participants supported by the virtual service.
	//
	// Comment: Some services have only a limitted number of participants permitted to connect at a time, this is typically used in "shared" virtual services that you might find on location resources.
	//
	// Typically on-demand services might not have these restrictions.
	MaxParticipants *int `bson:"maxParticipants,omitempty" json:"maxParticipants,omitempty" xml:"maxParticipants,omitempty"`
	// Session Key required by the virtual service.
	//
	// Comment: Some services require a session key to be able to access the service once connected.
	//
	// This could be in cases where a shared number is used, and a session key is added to put into a private line, or to identify the participant in a call.
	SessionKey *string `bson:"sessionKey,omitempty" json:"sessionKey,omitempty" xml:"sessionKey,omitempty"`
}
