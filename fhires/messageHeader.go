package fhires

import v2 "encoding/json/v2"

// The header for a message exchange that is either requesting or responding to an action.
//
//	The reference(s) that are the subject of the action as well as other information related to the action are typically transmitted in a bundle in which the MessageHeader resource instance is the first resource in the bundle.
//
// Purpose: Many implementations are not prepared to use REST and need a messaging based infrastructure.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/MessageHeader
type MessageHeader struct {
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
	// Code that identifies the event this message represents and connects it with its definition.
	//
	// Events defined as part of the FHIR specification are defined by the implementation.  Alternatively a canonical uri to the EventDefinition.
	//
	// Comment: The time of the event will be found in the focus resource. The time of the message will be found in [Bundle.timestamp](bundle-definitions.html#Bundle.timestamp).
	EventCoding    Coding `bson:"eventCoding" json:"eventCoding" xml:"eventCoding"`
	EventCanonical string `bson:"eventCanonical" json:"eventCanonical" xml:"eventCanonical"`
	// The destination application which the message is intended for.
	//
	// Comment: There SHOULD be at least one destination, but in some circumstances, the source system is unaware of any particular destination system.
	Destination []MessageHeaderDestination `bson:"destination,omitempty" json:"destination,omitempty" xml:"destination,omitempty"`
	// Identifies the sending system to allow the use of a trust relationship.
	//
	// Comment: Use case is for where a (trusted) sending system is responsible for multiple organizations, and therefore cannot differentiate based on source endpoint / authentication alone. Proposing to remove and rely on Task to convey this information.
	Sender *Reference `bson:"sender,omitempty" json:"sender,omitempty" xml:"sender,omitempty"`
	// The logical author of the message - the personor device that decided the described event should happen.
	//
	// When there is more than one candidate, pick the most proximal to the MessageHeader. Can provide other authors in extensions.
	//
	// Comment: Usually only for the request but can be used in a response.Proposing to remove and rely on Task to convey this information.
	Author *Reference `bson:"author,omitempty" json:"author,omitempty" xml:"author,omitempty"`
	// The source application from which this message originated.
	Source MessageHeaderSource `bson:"source" json:"source" xml:"source"`
	// The person or organization that accepts overall responsibility for the contents of the message.
	//
	// The implication is that the message event happened under the policies of the responsible party.
	//
	// Comment: Usually only for the request but can be used in a response.Proposing to remove and rely on Task to convey this information.
	Responsible *Reference `bson:"responsible,omitempty" json:"responsible,omitempty" xml:"responsible,omitempty"`
	// Coded indication of the cause for the event - indicates  a reason for the occurrence of the event that is a focus of this message.
	Reason *CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty" xml:"reason,omitempty"`
	// Information about the message that this message is a response to.
	//
	//  Only present if this message is a response.
	Response *MessageHeaderResponse `bson:"response,omitempty" json:"response,omitempty" xml:"response,omitempty"`
	// The actual data of the message - a reference to the root/focus class of the event.
	//
	// This is allowed to be a Parameters resource.
	//
	// Comment: The data is defined where the transaction type is defined. The transaction data is always included in the bundle that is the full message.  Only the root resource is specified.  The resources it references should be contained in the bundle but are not also listed here.  Multiple repetitions are allowed to cater for merges and other situations with multiple focal targets.
	Focus []Reference `bson:"focus,omitempty" json:"focus,omitempty" xml:"focus,omitempty"`
	// Permanent link to the MessageDefinition for this message.
	Definition *string `bson:"definition,omitempty" json:"definition,omitempty" xml:"definition,omitempty"`
}

type MessageHeaderDestination struct {
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
	// Indicates where the message should be routed.
	//
	// Comment: The url may be a non-resolvable URI for systems that do not use standard network-based addresses.
	EndpointUURL      *string    `bson:"endpointUrl,omitempty" json:"endpointUrl,omitempty" xml:"endpointUrl,omitempty"`
	EndpointReference *Reference `bson:"endpointReference,omitempty" json:"endpointReference,omitempty" xml:"endpointReference,omitempty"`
	// Human-readable name for the target system.
	Name *string `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Identifies the target end system in situations where the initial message transmission is to an intermediary system.
	Target *Reference `bson:"target,omitempty" json:"target,omitempty" xml:"target,omitempty"`
	// Allows data conveyed by a message to be addressed to a particular person or department when routing to a specific application isn't sufficient.
	Receiver *Reference `bson:"receiver,omitempty" json:"receiver,omitempty" xml:"receiver,omitempty"`
}

type MessageHeaderSource struct {
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
	// Identifies the routing target to send acknowledgements to.
	//
	// Comment: The url may be a non-resolvable URI for systems that do not use standard network-based addresses.
	EndpointUURL      *string    `bson:"endpointUrl,omitempty" json:"endpointUrl,omitempty" xml:"endpointUrl,omitempty"`
	EndpointReference *Reference `bson:"endpointReference,omitempty" json:"endpointReference,omitempty" xml:"endpointReference,omitempty"`
	// Human-readable name for the source system.
	Name *string `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// May include configuration or other information useful in debugging.
	Software *string `bson:"software,omitempty" json:"software,omitempty" xml:"software,omitempty"`
	// Can convey versions of multiple systems in situations where a message passes through multiple hands.
	Version *string `bson:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
	// An e-mail, phone, website or other contact point to use to resolve issues with message communications.
	Contact *ContactPoint `bson:"contact,omitempty" json:"contact,omitempty" xml:"contact,omitempty"`
}

type MessageHeaderResponse struct {
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
	// The Bundle.identifier of the message to which this message is a response.
	Identifier Identifier `bson:"identifier" json:"identifier" xml:"identifier"`
	// Code that identifies the type of response to the message - whether it was successful or not, and whether it should be resent or not.
	//
	// Comment: This is a generic response to the request message. Specific data for the response will be found in MessageHeader.focus.
	Code ResponseType `bson:"code" json:"code" xml:"code"`
	// Full details of any issues found in the message.
	//
	// Comment: This SHALL be contained in the bundle. If any of the issues are errors, the response code SHALL be an error.
	Details *Reference `bson:"details,omitempty" json:"details,omitempty" xml:"details,omitempty"`
}

func (r MessageHeader) ResourceType() string {
	return "MessageHeader"
}

// MarshalJSON implements json.Marshaler
func (r MessageHeader) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		MessageHeader
		ResourceType string `json:"resourceType"`
	}{
		MessageHeader: r,
		ResourceType:  r.ResourceType(),
	})
}
