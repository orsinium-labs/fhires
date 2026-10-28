package fhires

// UsageContext Type: Specifies clinical/business/etc.
//
// metadata that can be used to retrieve, index and/or categorize an artifact. This metadata can either be specific to the applicable population (e.g., age category, DRG) or the specific context of care (e.g., venue, care setting, provider of care).
//
// Purpose: Consumers of the resource must be able to determine the intended applicability for the resource. Ideally, this information would be used programmatically to determine when and how it should be incorporated or exposed.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/UsageContext
type UsageContext struct {
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
	// A code that identifies the type of context being specified by this usage context.
	Code Coding `bson:"code" json:"code" xml:"code"`
	// A value that defines the context specified in this context of use.
	//
	// The interpretation of the value is defined by the code.
	ValueCodeableConcept CodeableConcept `bson:"valueCodeableConcept" json:"valueCodeableConcept" xml:"valueCodeableConcept"`
	ValueQuantity        Quantity        `bson:"valueQuantity" json:"valueQuantity" xml:"valueQuantity"`
	ValueRange           Range           `bson:"valueRange" json:"valueRange" xml:"valueRange"`
	ValueReference       Reference       `bson:"valueReference" json:"valueReference" xml:"valueReference"`
}
