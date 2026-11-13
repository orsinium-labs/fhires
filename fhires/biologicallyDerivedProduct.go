package fhires

import "encoding/xml"

// A biological material originating from a biological entity intended to be transplanted or infused into another (possibly the same) biological entity.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/BiologicallyDerivedProduct
type BiologicallyDerivedProduct struct {
	DomainResource
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

// MarshalJSON implements [json.Marshaler]
func (r BiologicallyDerivedProduct) MarshalJSON() ([]byte, error) {
	type Wrapper BiologicallyDerivedProduct
	return marshalJSON(r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Marshaler].
func (r BiologicallyDerivedProduct) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Wrapper BiologicallyDerivedProduct
	return marshalXML(e, start, r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Unmarshaler].
func (r *BiologicallyDerivedProduct) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper BiologicallyDerivedProduct
	w := Wrapper(*r)
	err := unmarshalXML(d, start, r.ResourceType(), &w)
	*r = BiologicallyDerivedProduct(w)
	return err
}
