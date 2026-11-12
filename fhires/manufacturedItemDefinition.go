package fhires

// The definition and characteristics of a medicinal manufactured item, such as a tablet or capsule, as contained in a packaged medicinal product.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/ManufacturedItemDefinition
type ManufacturedItemDefinition struct {
	BaseResource
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
	// Unique identifier.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The status of this item. Enables tracking the life-cycle of the content.
	//
	// Comment: Allows filtering of manufactured items that are appropriate for use versus not.
	Status PublicationStatus `bson:"status" json:"status" xml:"status"`
	// A descriptive name applied to this item.
	Name *string `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Dose form as manufactured and before any transformation into the pharmaceutical product.
	ManufacturedDoseForm CodeableConcept `bson:"manufacturedDoseForm" json:"manufacturedDoseForm" xml:"manufacturedDoseForm"`
	// The “real-world” units in which the quantity of the manufactured item is described.
	UnitOfPresentation *CodeableConcept `bson:"unitOfPresentation,omitempty" json:"unitOfPresentation,omitempty" xml:"unitOfPresentation,omitempty"`
	// Manufacturer of the item, one of several possible.
	Manufacturer []Reference `bson:"manufacturer,omitempty" json:"manufacturer,omitempty" xml:"manufacturer,omitempty"`
	// Allows specifying that an item is on the market for sale, or that it is not available, and the dates and locations associated.
	MarketingStatus []MarketingStatus `bson:"marketingStatus,omitempty" json:"marketingStatus,omitempty" xml:"marketingStatus,omitempty"`
	// The ingredients of this manufactured item.
	//
	// This is only needed if the ingredients are not specified by incoming references from the Ingredient resource.
	Ingredient []CodeableConcept `bson:"ingredient,omitempty" json:"ingredient,omitempty" xml:"ingredient,omitempty"`
	// General characteristics of this item.
	Property []ManufacturedItemDefinitionProperty `bson:"property,omitempty" json:"property,omitempty" xml:"property,omitempty"`
	// Physical parts of the manufactured item, that it is intrisically made from.
	//
	// This is distinct from the ingredients that are part of its chemical makeup.
	Component []ManufacturedItemDefinitionComponent `bson:"component,omitempty" json:"component,omitempty" xml:"component,omitempty"`
}

type ManufacturedItemDefinitionProperty struct {
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
	// A code expressing the type of characteristic.
	Type CodeableConcept `bson:"type" json:"type" xml:"type"`
	// A value for the characteristic.
	ValueCodeableConcept *CodeableConcept `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty" xml:"valueCodeableConcept,omitempty"`
	ValueQuantity        *Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty" xml:"valueQuantity,omitempty"`
	ValueDate            *string          `bson:"valueDate,omitempty" json:"valueDate,omitempty" xml:"valueDate,omitempty"`
	ValueBoolean         *bool            `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty" xml:"valueBoolean,omitempty"`
	ValueMarkdown        *string          `bson:"valueMarkdown,omitempty" json:"valueMarkdown,omitempty" xml:"valueMarkdown,omitempty"`
	ValueAttachment      *Attachment      `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty" xml:"valueAttachment,omitempty"`
	ValueReference       *Reference       `bson:"valueReference,omitempty" json:"valueReference,omitempty" xml:"valueReference,omitempty"`
}

type ManufacturedItemDefinitionComponent struct {
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
	// Defining type of the component e.g. shell, layer, ink.
	Type CodeableConcept `bson:"type" json:"type" xml:"type"`
	// The function of this component within the item e.g.
	//
	// delivers active ingredient, masks taste.
	Function []CodeableConcept `bson:"function,omitempty" json:"function,omitempty" xml:"function,omitempty"`
	// The measurable amount of total quantity of all substances in the component, expressable in different ways (e.g.
	//
	// by mass or volume).
	Amount []Quantity `bson:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
	// A reference to a constituent of the manufactured item as a whole, linked here so that its component location within the item can be indicated.
	//
	// This not where the item's ingredient are primarily stated (for which see Ingredient.for or ManufacturedItemDefinition.ingredient).
	Constituent []ManufacturedItemDefinitionComponentConstituent `bson:"constituent,omitempty" json:"constituent,omitempty" xml:"constituent,omitempty"`
	// General characteristics of this component.
	Property []ManufacturedItemDefinitionProperty `bson:"property,omitempty" json:"property,omitempty"`
	// A component that this component contains or is made from.
	Component []ManufacturedItemDefinitionComponent `bson:"component,omitempty" json:"component,omitempty"`
}

type ManufacturedItemDefinitionComponentConstituent struct {
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
	// The measurable amount of the substance, expressable in different ways (e.g.
	//
	// by mass or volume).
	Amount []Quantity `bson:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
	// The physical location of the constituent/ingredient within the component.
	//
	// Example – if the component is the bead in the capsule, then the location would be where the ingredient resides within the product part – intragranular, extra-granular, etc.
	Location []CodeableConcept `bson:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// The function of this constituent within the component e.g. binder.
	Function []CodeableConcept `bson:"function,omitempty" json:"function,omitempty" xml:"function,omitempty"`
	// The ingredient that is the constituent of the given component.
	HasIngredient []CodeableReference `bson:"hasIngredient,omitempty" json:"hasIngredient,omitempty" xml:"hasIngredient,omitempty"`
}

func (r ManufacturedItemDefinition) ResourceType() string {
	return "ManufacturedItemDefinition"
}

// MarshalJSON implements json.Marshaler
func (r ManufacturedItemDefinition) MarshalJSON() ([]byte, error) {
	type Wrapper ManufacturedItemDefinition
	return marshalJSON(r.ResourceType(), Wrapper(r))
}
