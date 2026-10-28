package fhires

import v2 "encoding/json/v2"

// A biological material originating from a biological entity intended to be transplanted or infused into another (possibly the same) biological entity.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/BiologicallyDerivedProduct
type BiologicallyDerivedProduct struct {
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
	// Broad category of this product.
	ProductCategory *Coding `bson:"productCategory,omitempty" json:"productCategory,omitempty" xml:"productCategory,omitempty"`
	// A codified value that systematically supports characterization and classification of medical products of human origin inclusive of processing conditions such as additives, volumes and handling conditions.
	ProductCode *CodeableConcept `bson:"productCode,omitempty" json:"productCode,omitempty" xml:"productCode,omitempty"`
	// Parent product (if any) for this biologically-derived product.
	//
	// Comment: For products that have multiple collections. For example Peripheral Blood Stem Cells may be collected over several days from a single donor and the donation split into in multiple containers which must be linked to the parent donation.
	Parent []Reference `bson:"parent,omitempty" json:"parent,omitempty" xml:"parent,omitempty"`
	// Request to obtain and/or infuse this biologically derived product.
	Request []Reference `bson:"request,omitempty" json:"request,omitempty" xml:"request,omitempty"`
	// Unique instance identifiers assigned to a biologically derived product.
	//
	// Note: This is a business identifier, not a resource identifier.
	//
	// Comment: This identifier should uniquely identify the product instance in the business domain.  Ideally it should be a globally unique identifier under the control of an ISO/IEC 15459 Issuing Agency.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// An identifier that supports traceability to the event during which material in this product from one or more biological entities was obtained or pooled.
	//
	// Comment: Necessary to support mandatory requirements for traceability from donor/source to recipient and vice versa, while also satisfying donor anonymity requirements. The element is defined consistently across BiologicallyDerivedProduct, NutritionProduct, and Device.  The identifier references an event that links to a single biological entity such as a blood donor, or to multiple biological entities (e.g. when the product is an embryo or a pooled platelet product).  A single biologicalSourceEvent identifier may appear on multiple products of many types derived from a single donation event or source extraction.  As an example, a single donation event may provide 2 kidneys and a liver for organ transplantation, 2 corneas for eye surgery, heart valves and arterial tissue for cardiovascular surgery, multiple skin grafts, tendons, multiple shaped bone grafts and a large number of bone putty/paste products; and each of them may be assigned to the same biological source event identifier.
	BiologicalSourceEvent *Identifier `bson:"biologicalSourceEvent,omitempty" json:"biologicalSourceEvent,omitempty" xml:"biologicalSourceEvent,omitempty"`
	// Processing facilities responsible for the labeling and distribution of this biologically derived product.
	ProcessingFacility []Reference `bson:"processingFacility,omitempty" json:"processingFacility,omitempty" xml:"processingFacility,omitempty"`
	// A unique identifier for an aliquot of a product.
	//
	//  Used to distinguish individual aliquots of a product carrying the same biologicalSource and productCode identifiers.
	Division *string `bson:"division,omitempty" json:"division,omitempty" xml:"division,omitempty"`
	// Whether the product is currently available.
	ProductStatus *Coding `bson:"productStatus,omitempty" json:"productStatus,omitempty" xml:"productStatus,omitempty"`
	// Date, and where relevant time, of expiration.
	ExpirationDate *string `bson:"expirationDate,omitempty" json:"expirationDate,omitempty" xml:"expirationDate,omitempty"`
	// How this product was collected.
	Collection *BiologicallyDerivedProductCollection `bson:"collection,omitempty" json:"collection,omitempty" xml:"collection,omitempty"`
	// The temperature requirements for storage of the biologically-derived product.
	//
	// Comment: May be extracted from information held in the Product Description Code.
	StorageTempRequirements *Range `bson:"storageTempRequirements,omitempty" json:"storageTempRequirements,omitempty" xml:"storageTempRequirements,omitempty"`
	// A property that is specific to this BiologicallyDerviedProduct instance.
	//
	// Comment: Property can be used to provide information on a wide range of additional information specific to a particular biologicallyDerivedProduct.
	Property []BiologicallyDerivedProductProperty `bson:"property,omitempty" json:"property,omitempty" xml:"property,omitempty"`
}

type BiologicallyDerivedProductCollection struct {
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
	// Healthcare professional who is performing the collection.
	Collector *Reference `bson:"collector,omitempty" json:"collector,omitempty" xml:"collector,omitempty"`
	// The patient or entity, such as a hospital or vendor in the case of a processed/manipulated/manufactured product, providing the product.
	Source *Reference `bson:"source,omitempty" json:"source,omitempty" xml:"source,omitempty"`
	// Time of product collection.
	CollectedDateTime *string `bson:"collectedDateTime,omitempty" json:"collectedDateTime,omitempty" xml:"collectedDateTime,omitempty"`
	CollectedPeriod   *Period `bson:"collectedPeriod,omitempty" json:"collectedPeriod,omitempty" xml:"collectedPeriod,omitempty"`
}

type BiologicallyDerivedProductProperty struct {
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
	// Code that specifies the property. It should reference an established coding system.
	//
	// Comment: The element is identified by name and system URI in the type. Some types may have multiple occurrences.
	Type CodeableConcept `bson:"type" json:"type" xml:"type"`
	// Property values.
	//
	// Comment: The value should be provided as a boolean, integer, CodeableConcept, period, quantity, range, ratio, or attachment. The description can be a string only when these others are not available. The type of value will depend on the property type and is specified in ST-027.
	ValueBoolean         bool            `bson:"valueBoolean" json:"valueBoolean" xml:"valueBoolean"`
	ValueInteger         int             `bson:"valueInteger" json:"valueInteger" xml:"valueInteger"`
	ValueCodeableConcept CodeableConcept `bson:"valueCodeableConcept" json:"valueCodeableConcept" xml:"valueCodeableConcept"`
	ValuePeriod          Period          `bson:"valuePeriod" json:"valuePeriod" xml:"valuePeriod"`
	ValueQuantity        Quantity        `bson:"valueQuantity" json:"valueQuantity" xml:"valueQuantity"`
	ValueRange           Range           `bson:"valueRange" json:"valueRange" xml:"valueRange"`
	ValueRatio           Ratio           `bson:"valueRatio" json:"valueRatio" xml:"valueRatio"`
	ValueString          string          `bson:"valueString" json:"valueString" xml:"valueString"`
	ValueAttachment      Attachment      `bson:"valueAttachment" json:"valueAttachment" xml:"valueAttachment"`
}

func (r BiologicallyDerivedProduct) ResourceType() string {
	return "BiologicallyDerivedProduct"
}

// MarshalJSON implements json.Marshaler
func (r BiologicallyDerivedProduct) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		BiologicallyDerivedProduct
		ResourceType string `json:"resourceType"`
	}{
		BiologicallyDerivedProduct: r,
		ResourceType:               r.ResourceType(),
	})
}
