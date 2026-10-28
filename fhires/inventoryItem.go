package fhires

import (
	"encoding/json/v2"

	"github.com/shopspring/decimal"
)

// functional description of an inventory item used in inventory and supply-related workflows.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/InventoryItem
type InventoryItem struct {
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
	// Business identifier for the inventory item.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Status of the item entry.
	Status InventoryItemStatusCodes `bson:"status" json:"status" xml:"status"`
	// Category or class of the item.
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
	// Code designating the specific type of item.
	Code []CodeableConcept `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// The item name(s) - the brand name, or common name, functional name, generic name.
	Name []InventoryItemName `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Organization(s) responsible for the product.
	ResponsibleOrganization []InventoryItemResponsibleOrganization `bson:"responsibleOrganization,omitempty" json:"responsibleOrganization,omitempty" xml:"responsibleOrganization,omitempty"`
	// The descriptive characteristics of the inventory item.
	Description *InventoryItemDescription `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// The usage status e.g.
	//
	// recalled, in use, discarded... This can be used to indicate that the items have been taken out of inventory, or are in use, etc.
	InventoryStatus []CodeableConcept `bson:"inventoryStatus,omitempty" json:"inventoryStatus,omitempty" xml:"inventoryStatus,omitempty"`
	// The base unit of measure - the unit in which the product is used or counted.
	BaseUnit *CodeableConcept `bson:"baseUnit,omitempty" json:"baseUnit,omitempty" xml:"baseUnit,omitempty"`
	// Net content or amount present in the item.
	NetContent *Quantity `bson:"netContent,omitempty" json:"netContent,omitempty" xml:"netContent,omitempty"`
	// Association with other items or products.
	Association []InventoryItemAssociation `bson:"association,omitempty" json:"association,omitempty" xml:"association,omitempty"`
	// The descriptive or identifying characteristics of the item.
	Characteristic []InventoryItemCharacteristic `bson:"characteristic,omitempty" json:"characteristic,omitempty" xml:"characteristic,omitempty"`
	// Instances or occurrences of the product.
	Instance *InventoryItemInstance `bson:"instance,omitempty" json:"instance,omitempty" xml:"instance,omitempty"`
	// Link to a product resource used in clinical workflows.
	ProductReference *Reference `bson:"productReference,omitempty" json:"productReference,omitempty" xml:"productReference,omitempty"`
}

type InventoryItemName struct {
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
	// The type of name e.g. 'brand-name', 'functional-name', 'common-name'.
	NameType Coding `bson:"nameType" json:"nameType" xml:"nameType"`
	// The language that the item name is expressed in.
	Language string `bson:"language" json:"language" xml:"language"`
	// The name or designation that the item is given.
	Name string `bson:"name" json:"name" xml:"name"`
}

type InventoryItemResponsibleOrganization struct {
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
	// The role of the organization e.g. manufacturer, distributor, etc.
	Role CodeableConcept `bson:"role" json:"role" xml:"role"`
	// An organization that has an association with the item, e.g.
	//
	// manufacturer, distributor, responsible, etc.
	Organization Reference `bson:"organization" json:"organization" xml:"organization"`
}

type InventoryItemDescription struct {
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
	// The language for the item description, when an item must be described in different languages and those languages may be authoritative and not translations of a 'main' language.
	Language *string `bson:"language,omitempty" json:"language,omitempty" xml:"language,omitempty"`
	// Textual description of the item.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
}

type InventoryItemAssociation struct {
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
	// This attribute defined the type of association when establishing associations or relations between items, e.g.
	//
	// 'packaged within' or 'used with' or 'to be mixed with.
	AssociationType CodeableConcept `bson:"associationType" json:"associationType" xml:"associationType"`
	// The related item or product.
	RelatedItem Reference `bson:"relatedItem" json:"relatedItem" xml:"relatedItem"`
	// The quantity of the related product in this product - Numerator is the quantity of the related product.
	//
	// Denominator is the quantity of the present product. For example a value of 20 means that this product contains 20 units of the related product; a value of 1:20 means the inverse - that the contained product contains 20 units of the present product.
	Quantity Ratio `bson:"quantity" json:"quantity" xml:"quantity"`
}

type InventoryItemCharacteristic struct {
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
	// The type of characteristic that is being defined.
	CharacteristicType CodeableConcept `bson:"characteristicType" json:"characteristicType" xml:"characteristicType"`
	// The value of the attribute.
	//
	// Comment: The string value is used for characteristics that are descriptive and not codeable information. CodeableConcept.text is used when the characteristic is discrete and could otherwise be coded but for which there is no code available.
	ValueString          string          `bson:"valueString" json:"valueString" xml:"valueString"`
	ValueInteger         int             `bson:"valueInteger" json:"valueInteger" xml:"valueInteger"`
	ValueDecimal         decimal.Decimal `bson:"valueDecimal" json:"valueDecimal" xml:"valueDecimal"`
	ValueBoolean         bool            `bson:"valueBoolean" json:"valueBoolean" xml:"valueBoolean"`
	ValueUURL            string          `bson:"valueUrl" json:"valueUrl" xml:"valueUrl"`
	ValueDateTime        string          `bson:"valueDateTime" json:"valueDateTime" xml:"valueDateTime"`
	ValueQuantity        Quantity        `bson:"valueQuantity" json:"valueQuantity" xml:"valueQuantity"`
	ValueRange           Range           `bson:"valueRange" json:"valueRange" xml:"valueRange"`
	ValueRatio           Ratio           `bson:"valueRatio" json:"valueRatio" xml:"valueRatio"`
	ValueAnnotation      Annotation      `bson:"valueAnnotation" json:"valueAnnotation" xml:"valueAnnotation"`
	ValueAddress         Address         `bson:"valueAddress" json:"valueAddress" xml:"valueAddress"`
	ValueDuration        Duration        `bson:"valueDuration" json:"valueDuration" xml:"valueDuration"`
	ValueCodeableConcept CodeableConcept `bson:"valueCodeableConcept" json:"valueCodeableConcept" xml:"valueCodeableConcept"`
}

type InventoryItemInstance struct {
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
	// The identifier for the physical instance, typically a serial number.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The lot or batch number of the item.
	LotNumber *string `bson:"lotNumber,omitempty" json:"lotNumber,omitempty" xml:"lotNumber,omitempty"`
	// The expiry date or date and time for the product.
	Expiry *string `bson:"expiry,omitempty" json:"expiry,omitempty" xml:"expiry,omitempty"`
	// The subject that the item is associated with.
	Subject *Reference `bson:"subject,omitempty" json:"subject,omitempty" xml:"subject,omitempty"`
	// The location that the item is associated with.
	Location *Reference `bson:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
}

func (r InventoryItem) ResourceType() string {
	return "InventoryItem"
}

// MarshalJSON implements json.Marshaler
func (r InventoryItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		InventoryItem
		ResourceType string `json:"resourceType"`
	}{
		InventoryItem: r,
		ResourceType:  r.ResourceType(),
	})
}
