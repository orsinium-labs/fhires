package fhires

import v2 "encoding/json/v2"

// Record of delivery of what is supplied.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/SupplyDelivery
type SupplyDelivery struct {
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
	// Identifier for the supply delivery event that is used to identify it across multiple disparate systems.
	//
	// Comment: This identifier is typically assigned by the supplier, and may be used to reference the delivery when exchanging information about it with other systems.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// A plan, proposal or order that is fulfilled in whole or in part by this event.
	BasedOn []Reference `bson:"basedOn,omitempty" json:"basedOn,omitempty" xml:"basedOn,omitempty"`
	// A larger event of which this particular event is a component or step.
	//
	// Comment: Not to be used to link an event to an Encounter - use Event.context for that.[The allowed reference resources may be adjusted as appropriate for the event resource].
	PartOf []Reference `bson:"partOf,omitempty" json:"partOf,omitempty" xml:"partOf,omitempty"`
	// A code specifying the state of the dispense event.
	//
	// Comment: This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid.
	Status *SupplyDeliveryStatus `bson:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// A link to a resource representing the person whom the delivered item is for.
	Patient *Reference `bson:"patient,omitempty" json:"patient,omitempty" xml:"patient,omitempty"`
	// Indicates the type of supply being provided.
	//
	//  Examples include: Medication, Device, Biologically Derived Product.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The item that is being delivered or has been supplied.
	SuppliedItem []SupplyDeliverySuppliedItem `bson:"suppliedItem,omitempty" json:"suppliedItem,omitempty" xml:"suppliedItem,omitempty"`
	// The date or time(s) the activity occurred.
	//
	// Comment: [The list of types may be constrained as appropriate for the type of event].
	OccurrenceDateTime *string `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty" xml:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *Period `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty" xml:"occurrencePeriod,omitempty"`
	OccurrenceTiming   *Timing `bson:"occurrenceTiming,omitempty" json:"occurrenceTiming,omitempty" xml:"occurrenceTiming,omitempty"`
	// The individual or organization responsible for supplying the delivery.
	Supplier *Reference `bson:"supplier,omitempty" json:"supplier,omitempty" xml:"supplier,omitempty"`
	// Identification of the facility/location where the delivery was shipped to.
	Destination *Reference `bson:"destination,omitempty" json:"destination,omitempty" xml:"destination,omitempty"`
	// Identifies the individual or organization that received the delivery.
	Receiver []Reference `bson:"receiver,omitempty" json:"receiver,omitempty" xml:"receiver,omitempty"`
}

type SupplyDeliverySuppliedItem struct {
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
	// The amount of the item that has been supplied.  Unit of measure may be included.
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty" xml:"quantity,omitempty"`
	// Identifies the medication, substance, device or biologically derived product being supplied.
	//
	// This is either a link to a resource representing the details of the item or a code that identifies the item from a known list.
	ItemCodeableConcept *CodeableConcept `bson:"itemCodeableConcept,omitempty" json:"itemCodeableConcept,omitempty" xml:"itemCodeableConcept,omitempty"`
	ItemReference       *Reference       `bson:"itemReference,omitempty" json:"itemReference,omitempty" xml:"itemReference,omitempty"`
}

func (r SupplyDelivery) ResourceType() string {
	return "SupplyDelivery"
}

// MarshalJSON implements json.Marshaler
func (r SupplyDelivery) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		SupplyDelivery
		ResourceType string `json:"resourceType"`
	}{
		ResourceType:   r.ResourceType(),
		SupplyDelivery: r,
	})
}
