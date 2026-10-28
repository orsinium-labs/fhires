package fhires

import v2 "encoding/json/v2"

// A record of dispensation of a biologically derived product.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/BiologicallyDerivedProductDispense
type BiologicallyDerivedProductDispense struct {
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
	// Unique instance identifiers assigned to a biologically derived product dispense.
	//
	// Note: This is a business identifier, not a resource identifier.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The order or request that the dispense is fulfilling.
	//
	// This is a reference to a ServiceRequest resource.
	BasedOn []Reference `bson:"basedOn,omitempty" json:"basedOn,omitempty" xml:"basedOn,omitempty"`
	// A larger event of which this particular event is a component.
	PartOf []Reference `bson:"partOf,omitempty" json:"partOf,omitempty" xml:"partOf,omitempty"`
	// A code specifying the state of the dispense event.
	Status BiologicallyDerivedProductDispenseCodes `bson:"status" json:"status" xml:"status"`
	// Indicates the relationship between the donor of the biologically derived product and the intended recipient.
	OriginRelationshipType *CodeableConcept `bson:"originRelationshipType,omitempty" json:"originRelationshipType,omitempty" xml:"originRelationshipType,omitempty"`
	// A link to a resource identifying the biologically derived product that is being dispensed.
	Product Reference `bson:"product" json:"product" xml:"product"`
	// A link to a resource representing the patient that the product is dispensed for.
	Patient Reference `bson:"patient" json:"patient" xml:"patient"`
	// Indicates the type of matching associated with the dispense.
	MatchStatus *CodeableConcept `bson:"matchStatus,omitempty" json:"matchStatus,omitempty" xml:"matchStatus,omitempty"`
	// Indicates who or what performed an action.
	Performer []BiologicallyDerivedProductDispensePerformer `bson:"performer,omitempty" json:"performer,omitempty" xml:"performer,omitempty"`
	// The physical location where the dispense was performed.
	Location *Reference `bson:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// The amount of product in the dispense.
	//
	// Quantity will depend on the product being dispensed. Examples are: volume; cell count; concentration.
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty" xml:"quantity,omitempty"`
	// When the product was selected/ matched.
	PreparedDate *string `bson:"preparedDate,omitempty" json:"preparedDate,omitempty" xml:"preparedDate,omitempty"`
	// When the product was dispatched for clinical use.
	WhenHandedOver *string `bson:"whenHandedOver,omitempty" json:"whenHandedOver,omitempty" xml:"whenHandedOver,omitempty"`
	// Link to a resource identifying the physical location that the product was dispatched to.
	Destination *Reference `bson:"destination,omitempty" json:"destination,omitempty" xml:"destination,omitempty"`
	// Additional notes.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// Specific instructions for use.
	UsageInstruction *string `bson:"usageInstruction,omitempty" json:"usageInstruction,omitempty" xml:"usageInstruction,omitempty"`
}

type BiologicallyDerivedProductDispensePerformer struct {
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
	// Identifies the function of the performer during the dispense.
	Function *CodeableConcept `bson:"function,omitempty" json:"function,omitempty" xml:"function,omitempty"`
	// Identifies the person responsible for the action.
	Actor Reference `bson:"actor" json:"actor" xml:"actor"`
}

func (r BiologicallyDerivedProductDispense) ResourceType() string {
	return "BiologicallyDerivedProductDispense"
}

// MarshalJSON implements json.Marshaler
func (r BiologicallyDerivedProductDispense) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		BiologicallyDerivedProductDispense
		ResourceType string `json:"resourceType"`
	}{
		BiologicallyDerivedProductDispense: r,
		ResourceType:                       r.ResourceType(),
	})
}
