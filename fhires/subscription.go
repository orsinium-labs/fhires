package fhires

import "encoding/json/v2"

// The subscription resource describes a particular client's request to be notified about a SubscriptionTopic.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Subscription
type Subscription struct {
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
	// A formal identifier that is used to identify this code system when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// A natural language name identifying the subscription.
	Name *string `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// The status of the subscription, which marks the server state for managing the subscription.
	//
	// Comment: A client can only submit subscription resources in the requested or off state. Only the server can  move a subscription from requested to active, and then to error. Either the server or the client can turn a subscription off.
	//
	// This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid.
	Status SubscriptionStatusCodes `bson:"status" json:"status" xml:"status"`
	// The reference to the subscription topic to be notified about.
	Topic string `bson:"topic" json:"topic" xml:"topic"`
	// Contact details for a human to contact about the subscription.
	//
	// The primary use of this for system administrator troubleshooting.
	Contact []ContactPoint `bson:"contact,omitempty" json:"contact,omitempty" xml:"contact,omitempty"`
	// The time for the server to turn the subscription off.
	//
	// Comment: The server is permitted to deviate from this time but should observe it.
	End *string `bson:"end,omitempty" json:"end,omitempty" xml:"end,omitempty"`
	// Entity with authorization to make subsequent revisions to the Subscription and also determines what data the subscription is authorized to disclose.
	ManagingEntity *Reference `bson:"managingEntity,omitempty" json:"managingEntity,omitempty" xml:"managingEntity,omitempty"`
	// A description of why this subscription is defined.
	Reason *string `bson:"reason,omitempty" json:"reason,omitempty" xml:"reason,omitempty"`
	// The filter properties to be applied to narrow the subscription topic stream.
	//
	//  When multiple filters are applied, evaluates to true if all the conditions applicable to that resource are met; otherwise it returns false (i.e., logical AND).
	FilterBy []SubscriptionFilterBy `bson:"filterBy,omitempty" json:"filterBy,omitempty" xml:"filterBy,omitempty"`
	// The type of channel to send notifications on.
	ChannelType Coding `bson:"channelType" json:"channelType" xml:"channelType"`
	// The url that describes the actual end-point to send notifications to.
	//
	// Comment: For rest-hook the end-point must be an `http:` or `https:` URL; for websockets, `ws:` or `wss:`; for email, a `mailto:` url; and for message the endpoint can be in any form of url the server understands (usually, `http/s`: or `mllp:`). The URI is allowed to be relative; in which case, it is relative to the server end-point (since there may be more than one, clients should avoid using relative URIs).
	Endpoint *string `bson:"endpoint,omitempty" json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// Channel-dependent information to send as part of the notification (e.g., HTTP Headers).
	//
	// Comment: Exactly what these mean depend on the channel type. They can convey additional information to the server or recipient and/or meet security requirements; for example, support of multiple headers in the outgoing notifications for rest-hook type subscriptions. Note that names are not required to be unique, but channel definitions can impose restrictions.
	Parameter []SubscriptionParameter `bson:"parameter,omitempty" json:"parameter,omitempty" xml:"parameter,omitempty"`
	// If present, a 'heartbeat' notification (keep-alive) is sent via this channel with an interval period equal to this elements integer value in seconds.
	//
	//  If not present, a heartbeat notification is not sent.
	HeartbeatPeriod *int `bson:"heartbeatPeriod,omitempty" json:"heartbeatPeriod,omitempty" xml:"heartbeatPeriod,omitempty"`
	// If present, the maximum amount of time a server will allow before failing a notification attempt.
	Timeout *int `bson:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// The MIME type to send the payload in - e.g., `application/fhir+xml` or `application/fhir+json`.
	//
	// Note that:
	//
	// * clients may request notifications in a specific FHIR version by using the [FHIR Version Parameter](http.html#version-parameter) - e.g., `application/fhir+json; fhirVersion=4.0`.
	//
	// * additional MIME types can be allowed by channels - e.g., `text/plain` and `text/html` are defined by the Email channel.
	ContentType *string `bson:"contentType,omitempty" json:"contentType,omitempty" xml:"contentType,omitempty"`
	// How much of the resource content to deliver in the notification payload.
	//
	// The choices are an empty payload, only the resource id, or the full resource content.
	//
	// Comment: Sending the payload has obvious security implications. The server is responsible for ensuring that the content is appropriately secured.
	Content *SubscriptionPayloadContent `bson:"content,omitempty" json:"content,omitempty" xml:"content,omitempty"`
	// If present, the maximum number of events that will be included in a notification bundle.
	//
	// Note that this is not a strict limit on the number of entries in a bundle, as dependent resources can be included.
	MaxCount *int `bson:"maxCount,omitempty" json:"maxCount,omitempty" xml:"maxCount,omitempty"`
}

type SubscriptionFilterBy struct {
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
	// A resource listed in the `SubscriptionTopic` this `Subscription` references (`SubscriptionTopic.canFilterBy.resource`).
	//
	// This element can be used to differentiate filters for topics that include more than one resource type.
	ResourceType *string `bson:"resourceType,omitempty" json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The filter as defined in the `SubscriptionTopic.canFilterBy.filterParameter` element.
	FilterParameter string `bson:"filterParameter" json:"filterParameter" xml:"filterParameter"`
	// Comparator applied to this filter parameter.
	//
	// Comment: Must be a comparator allowed by the SubscriptionTopic relevant to this Subscription filter.
	Comparator *SearchComparator `bson:"comparator,omitempty" json:"comparator,omitempty" xml:"comparator,omitempty"`
	// Modifier applied to this filter parameter.
	//
	// Comment: Must be a modifier allowed by the SubscriptionTopic relevant to this Subscription filter.
	Modifier *SearchModifierCode `bson:"modifier,omitempty" json:"modifier,omitempty" xml:"modifier,omitempty"`
	// The literal value or resource path as is legal in search - for example, `Patient/123` or `le1950`.
	Value string `bson:"value" json:"value" xml:"value"`
}

type SubscriptionParameter struct {
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
	// Parameter name for information passed to the channel for notifications, for example in the case of a REST hook wanting to pass through an authorization header, the name would be Authorization.
	Name string `bson:"name" json:"name" xml:"name"`
	// Parameter value for information passed to the channel for notifications, for example in the case of a REST hook wanting to pass through an authorization header, the value would be `Bearer 0193...`.
	Value string `bson:"value" json:"value" xml:"value"`
}

func (r Subscription) ResourceType() string {
	return "Subscription"
}

// MarshalJSON implements json.Marshaler
func (r Subscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Subscription
		ResourceType string `json:"resourceType"`
	}{
		ResourceType: r.ResourceType(),
		Subscription: r,
	})
}
