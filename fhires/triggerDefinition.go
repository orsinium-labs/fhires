package fhires

// TriggerDefinition Type: A description of a triggering event.
//
// Triggering events can be named events, data events, or periodic, as determined by the type element.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/TriggerDefinition
type TriggerDefinition struct {
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
	// The type of triggering event.
	Type TriggerType `bson:"type" json:"type" xml:"type"`
	// A formal name for the event.
	//
	// This may be an absolute URI that identifies the event formally (e.g. from a trigger registry), or a simple relative URI that identifies the event in a local context.
	//
	// Comment: An event name can be provided for all event types, but is required for named events. If a name is provided for a type other than named events, it is considered to be a shorthand for the semantics described by the formal description of the event.
	Name *string `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// A code that identifies the event.
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// A reference to a SubscriptionTopic resource that defines the event.
	//
	// If this element is provided, no other information about the trigger definition may be supplied.
	SubscriptionTopic *string `bson:"subscriptionTopic,omitempty" json:"subscriptionTopic,omitempty" xml:"subscriptionTopic,omitempty"`
	// The timing of the event (if this is a periodic trigger).
	TimingTiming    *Timing    `bson:"timingTiming,omitempty" json:"timingTiming,omitempty" xml:"timingTiming,omitempty"`
	TimingReference *Reference `bson:"timingReference,omitempty" json:"timingReference,omitempty" xml:"timingReference,omitempty"`
	TimingDate      *string    `bson:"timingDate,omitempty" json:"timingDate,omitempty" xml:"timingDate,omitempty"`
	TimingDateTime  *string    `bson:"timingDateTime,omitempty" json:"timingDateTime,omitempty" xml:"timingDateTime,omitempty"`
	// The triggering data of the event (if this is a data trigger).
	//
	// If more than one data is requirement is specified, then all the data requirements must be true.
	//
	// Comment: This element shall be present for any data type trigger.
	Data []DataRequirement `bson:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	// A boolean-valued expression that is evaluated in the context of the container of the trigger definition and returns whether or not the trigger fires.
	//
	// Comment: This element can be only be specified for data type triggers and provides additional semantics for the trigger. The context available within the condition is based on the type of data event. For all events, the current resource will be available as context. In addition, for modification events, the previous resource will also be available. The expression may be inlined, or may be a simple absolute URI, which is a reference to a named expression within a logic library referenced by a library element or extension within the containing resource. If the expression is a FHIR Path expression, it evaluates in the context of a resource of one of the type identified in the data requirement, and may also refer to the variable %previous for delta comparisons on events of type data-changed, data-modified, and data-deleted which will always have the same type.
	Condition *Expression `bson:"condition,omitempty" json:"condition,omitempty" xml:"condition,omitempty"`
}
