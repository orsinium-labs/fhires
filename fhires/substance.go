package fhires

import "encoding/xml"

// A homogeneous material with a definite composition.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Substance
type Substance struct {
	DomainResource
	// Unique identifier for the substance.
	//
	// For an instance, an identifier associated with the package/container (usually a label affixed directly).
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// A boolean to indicate if this an instance of a substance or a kind of one (a definition).
	Instance bool `bson:"instance" json:"instance" xml:"instance"`
	// A code to indicate if the substance is actively used.
	Status *FHIRSubstanceStatus `bson:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// A code that classifies the general type of substance.
	//
	//  This is used  for searching, sorting and display purposes.
	//
	// Comment: The level of granularity is defined by the category concepts in the value set.   More fine-grained filtering can be performed using the metadata and/or terminology hierarchy in Substance.code.
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
	// A code (or set of codes) that identify this substance.
	//
	// Comment: This could be a reference to an externally defined code.  It could also be a locally assigned code (e.g. a formulary),  optionally with translations to the standard drug codes.
	Code CodeableReference `bson:"code" json:"code" xml:"code"`
	// A description of the substance - its appearance, handling requirements, and other usage notes.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// When the substance is no longer valid to use.
	//
	// For some substances, a single arbitrary date is used for expiry.
	Expiry *string `bson:"expiry,omitempty" json:"expiry,omitempty" xml:"expiry,omitempty"`
	// The amount of the substance.
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty" xml:"quantity,omitempty"`
	// A substance can be composed of other substances.
	Ingredient []SubstanceIngredient `bson:"ingredient,omitempty" json:"ingredient,omitempty" xml:"ingredient,omitempty"`
}

type SubstanceIngredient struct {
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
	// The amount of the ingredient in the substance - a concentration ratio.
	Quantity *Ratio `bson:"quantity,omitempty" json:"quantity,omitempty" xml:"quantity,omitempty"`
	// Another substance that is a component of this substance.
	SubstanceCodeableConcept CodeableConcept `bson:"substanceCodeableConcept" json:"substanceCodeableConcept" xml:"substanceCodeableConcept"`
	SubstanceReference       Reference       `bson:"substanceReference" json:"substanceReference" xml:"substanceReference"`
}

func (r Substance) ResourceType() string {
	return "Substance"
}

// MarshalJSON implements [json.Marshaler]
func (r Substance) MarshalJSON() ([]byte, error) {
	type Wrapper Substance
	return marshalJSON(r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Marshaler].
func (r Substance) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Wrapper Substance
	return marshalXML(e, start, r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Unmarshaler].
func (r *Substance) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper Substance
	w := Wrapper(*r)
	err := unmarshalXML(d, start, r.ResourceType(), &w)
	*r = Substance(w)
	return err
}
