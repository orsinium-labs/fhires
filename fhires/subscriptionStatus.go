package fhires

import "encoding/json/v2"

// The SubscriptionStatus resource describes the state of a Subscription during notifications.
//
// It is not persisted.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/SubscriptionStatus
type SubscriptionStatus struct {
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
	// The status of the subscription, which marks the server state for managing the subscription.
	Status *SubscriptionStatusCodes `bson:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// The type of event being conveyed with this notification.
	Type SubscriptionNotificationType `bson:"type" json:"type" xml:"type"`
	// The total number of actual events which have been generated since the Subscription was created (inclusive of this notification) - regardless of how many have been successfully communicated.
	//
	//  This number is NOT incremented for handshake and heartbeat notifications.
	EventsSinceSubscriptionStart *int64 `bson:"eventsSinceSubscriptionStart,omitempty" json:"eventsSinceSubscriptionStart,omitempty,string" xml:"eventsSinceSubscriptionStart,omitempty"`
	// Detailed information about events relevant to this subscription notification.
	NotificationEvent []SubscriptionStatusNotificationEvent `bson:"notificationEvent,omitempty" json:"notificationEvent,omitempty" xml:"notificationEvent,omitempty"`
	// The reference to the Subscription which generated this notification.
	Subscription Reference `bson:"subscription" json:"subscription" xml:"subscription"`
	// The reference to the SubscriptionTopic for the Subscription which generated this notification.
	//
	// Comment: This value SHOULD NOT be present when using `empty` payloads, MAY be present when using `id-only` payloads, and SHOULD be present when using `full-resource` payloads.
	Topic *string `bson:"topic,omitempty" json:"topic,omitempty" xml:"topic,omitempty"`
	// A record of errors that occurred when the server processed a notification.
	//
	// Comment: Recommended practice: clear errors when status is updated.
	Error []CodeableConcept `bson:"error,omitempty" json:"error,omitempty" xml:"error,omitempty"`
}

type SubscriptionStatusNotificationEvent struct {
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
	// Either the sequential number of this event in this subscription context or a relative event number for this notification.
	//
	// Comment: In subscriptions where delivery of notifications IS NOT guaranteed (e.g., REST-Hook), this number is a unique and monotonically-increasing event number for a subscription. In channels where delivery of notifications IS guaranteed, this number is a relative index for the events present in the notification (e.g., 1, 2, etc.).
	EventNumber int64 `bson:"eventNumber" json:"eventNumber,string" xml:"eventNumber"`
	// The actual time this event occurred on the server.
	Timestamp *string `bson:"timestamp,omitempty" json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// The focus of this event.
	//
	// While this will usually be a reference to the focus resource of the event, it MAY contain a reference to a non-FHIR object.
	Focus *Reference `bson:"focus,omitempty" json:"focus,omitempty" xml:"focus,omitempty"`
	// Additional context information for this event.
	//
	// Generally, this will contain references to additional resources included with the event (e.g., the Patient relevant to an Encounter), however it MAY refer to non-FHIR objects.
	AdditionalContext []Reference `bson:"additionalContext,omitempty" json:"additionalContext,omitempty" xml:"additionalContext,omitempty"`
}

func (r SubscriptionStatus) ResourceType() string {
	return "SubscriptionStatus"
}

// MarshalJSON implements json.Marshaler
func (r SubscriptionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		SubscriptionStatus
		ResourceType string `json:"resourceType"`
	}{
		ResourceType:       r.ResourceType(),
		SubscriptionStatus: r,
	})
}
