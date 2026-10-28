package fhires

import v2 "encoding/json/v2"

// A medically related item or items, in a container or package.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/PackagedProductDefinition
type PackagedProductDefinition struct {
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
	// A unique identifier for this package as whole - not the the content of the package.
	//
	// Unique instance identifiers assigned to a package by manufacturers, regulators, drug catalogue custodians or other organizations.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// A name for this package.
	//
	// Typically what it would be listed as in a drug formulary or catalogue, inventory etc.
	Name *string `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// A high level category e.g.
	//
	// medicinal product, raw material, shipping/transport container, etc.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The product this package model relates to, not the contents of the package (for which see package.containedItem).
	PackageFor []Reference `bson:"packageFor,omitempty" json:"packageFor,omitempty" xml:"packageFor,omitempty"`
	// The status within the lifecycle of this item.
	//
	// A high level status, this is not intended to duplicate details carried elsewhere such as legal status, or authorization or marketing status.
	Status *CodeableConcept `bson:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// The date at which the given status became applicable.
	StatusDate *string `bson:"statusDate,omitempty" json:"statusDate,omitempty" xml:"statusDate,omitempty"`
	// A total of the complete count of contained items of a particular type/form, independent of sub-packaging or organization.
	//
	// This can be considered as the pack size. This attribute differs from containedItem.amount in that it can give a single aggregated count of all tablet types in a pack, even when these are different manufactured items. For example a pill pack of 21 tablets plus 7 sugar tablets, can be denoted here as '28 tablets'. This attribute is repeatable so that the different item types in one pack type can be counted (e.g. a count of vials and count of syringes). Each repeat must have different units, so that it is clear what the different sets of counted items are, and it is not intended to allow different counts of similar items (e.g. not '2 tubes and 3 tubes'). Repeats are not to be used to represent different pack sizes (e.g. 20 pack vs. 50 pack) - which would be different instances of this resource.
	ContainedItemQuantity []Quantity `bson:"containedItemQuantity,omitempty" json:"containedItemQuantity,omitempty" xml:"containedItemQuantity,omitempty"`
	// Textual description. Note that this is not the name of the package or product.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// The legal status of supply of the packaged item as classified by the regulator.
	LegalStatusOfSupply []PackagedProductDefinitionLegalStatusOfSupply `bson:"legalStatusOfSupply,omitempty" json:"legalStatusOfSupply,omitempty" xml:"legalStatusOfSupply,omitempty"`
	// Allows specifying that an item is on the market for sale, or that it is not available, and the dates and locations associated.
	MarketingStatus []MarketingStatus `bson:"marketingStatus,omitempty" json:"marketingStatus,omitempty" xml:"marketingStatus,omitempty"`
	// Identifies if the package contains different items, such as when a drug product is supplied with another item e.g.
	//
	// a diluent or adjuvant.
	CopackagedIndicator *bool `bson:"copackagedIndicator,omitempty" json:"copackagedIndicator,omitempty" xml:"copackagedIndicator,omitempty"`
	// Manufacturer of this package type.
	//
	// When there are multiple it means these are all possible manufacturers.
	Manufacturer []Reference `bson:"manufacturer,omitempty" json:"manufacturer,omitempty" xml:"manufacturer,omitempty"`
	// Additional information or supporting documentation about the packaged product.
	AttachedDocument []Reference `bson:"attachedDocument,omitempty" json:"attachedDocument,omitempty" xml:"attachedDocument,omitempty"`
	// A packaging item, as a container for medically related items, possibly with other packaging items within, or a packaging component, such as bottle cap (which is not a device or a medication manufactured item).
	Packaging *PackagedProductDefinitionPackaging `bson:"packaging,omitempty" json:"packaging,omitempty" xml:"packaging,omitempty"`
	// Allows the key features to be recorded, such as "hospital pack", "nurse prescribable", "calendar pack".
	Characteristic []PackagedProductDefinitionPackagingProperty `bson:"characteristic,omitempty" json:"characteristic,omitempty"`
}

type PackagedProductDefinitionLegalStatusOfSupply struct {
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
	// The actual status of supply.
	//
	// Conveys in what situation this package type may be supplied for use.
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// The place where the legal status of supply applies.
	//
	// When not specified, this indicates it is unknown in this context.
	Jurisdiction *CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty" xml:"jurisdiction,omitempty"`
}

type PackagedProductDefinitionPackaging struct {
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
	// A business identifier that is specific to this particular part of the packaging, often assigned by the manufacturer.
	//
	// Including possibly Data Carrier Identifier (a GS1 barcode).
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The physical type of the container of the items.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Is this a part of the packaging (e.g.
	//
	// a cap or bottle stopper), rather than the packaging itself (e.g. a bottle or vial). The latter type are designed be a container, but the former are not.
	ComponentPart *bool `bson:"componentPart,omitempty" json:"componentPart,omitempty" xml:"componentPart,omitempty"`
	// The quantity of packaging items contained at this layer of the package.
	//
	// This does not relate to the number of contained items but relates solely to the number of packaging items. When looking at the outermost layer it is always 1. If there are two boxes within, at the next layer it would be 2.
	Quantity *int `bson:"quantity,omitempty" json:"quantity,omitempty" xml:"quantity,omitempty"`
	// Material type of the package item.
	Material []CodeableConcept `bson:"material,omitempty" json:"material,omitempty" xml:"material,omitempty"`
	// A possible alternate material for this part of the packaging, that is allowed to be used instead of the usual material (e.g.
	//
	// different types of plastic for a blister sleeve).
	AlternateMaterial []CodeableConcept `bson:"alternateMaterial,omitempty" json:"alternateMaterial,omitempty" xml:"alternateMaterial,omitempty"`
	// Shelf Life and storage information.
	ShelfLifeStorage []ProductShelfLife `bson:"shelfLifeStorage,omitempty" json:"shelfLifeStorage,omitempty" xml:"shelfLifeStorage,omitempty"`
	// Manufacturer of this packaging item.
	//
	// When there are multiple values each one is a potential manufacturer of this packaging item.
	Manufacturer []Reference `bson:"manufacturer,omitempty" json:"manufacturer,omitempty" xml:"manufacturer,omitempty"`
	// General characteristics of this item.
	Property []PackagedProductDefinitionPackagingProperty `bson:"property,omitempty" json:"property,omitempty" xml:"property,omitempty"`
	// The item(s) within the packaging.
	ContainedItem []PackagedProductDefinitionPackagingContainedItem `bson:"containedItem,omitempty" json:"containedItem,omitempty" xml:"containedItem,omitempty"`
	// Allows containers (and parts of containers) within containers, still as a part of a single packaged product.
	//
	// See also PackagedProductDefinition.packaging.containedItem.item(PackagedProductDefinition).
	Packaging []PackagedProductDefinitionPackaging `bson:"packaging,omitempty" json:"packaging,omitempty"`
}

type PackagedProductDefinitionPackagingProperty struct {
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
	ValueAttachment      *Attachment      `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty" xml:"valueAttachment,omitempty"`
}

type PackagedProductDefinitionPackagingContainedItem struct {
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
	// The actual item(s) of medication, as manufactured, or a device (typically, but not necessarily, a co-packaged one), or other medically related item (such as food, biologicals, raw materials, medical fluids, gases etc.), as contained in the package.
	//
	// This also allows another whole packaged product to be included, which is solely for the case where a package of other entire packages is wanted - such as a wholesale or distribution pack (for layers within one package, use PackagedProductDefinition.packaging.packaging).
	Item CodeableReference `bson:"item" json:"item" xml:"item"`
	// The number of this type of item within this packaging or for continuous items such as liquids it is the quantity (for example 25ml).
	//
	// See also PackagedProductDefinition.containedItemQuantity (especially the long definition).
	Amount *Quantity `bson:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
}

func (r PackagedProductDefinition) ResourceType() string {
	return "PackagedProductDefinition"
}

// MarshalJSON implements json.Marshaler
func (r PackagedProductDefinition) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		PackagedProductDefinition
		ResourceType string `json:"resourceType"`
	}{
		PackagedProductDefinition: r,
		ResourceType:              r.ResourceType(),
	})
}
