package fhires

import "encoding/xml"

// An ingredient of a manufactured item or pharmaceutical product.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Ingredient
type Ingredient struct {
	DomainResource
	// The identifier(s) of this Ingredient that are assigned by business processes and/or used to refer to it when a direct URL reference to the resource itself is not appropriate.
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The status of this ingredient. Enables tracking the life-cycle of the content.
	//
	// Comment: Allows filtering of ingredient that are appropriate for use versus not.
	Status PublicationStatus `bson:"status" json:"status" xml:"status"`
	// The product which this ingredient is a constituent part of.
	For []Reference `bson:"for,omitempty" json:"for,omitempty" xml:"for,omitempty"`
	// A classification of the ingredient identifying its purpose within the product, e.g.
	//
	// active, inactive.
	Role CodeableConcept `bson:"role" json:"role" xml:"role"`
	// A classification of the ingredient identifying its precise purpose(s) in the drug product.
	//
	// This extends the Ingredient.role to add more detail. Example: antioxidant, alkalizing agent.
	Function []CodeableConcept `bson:"function,omitempty" json:"function,omitempty" xml:"function,omitempty"`
	// A classification of the ingredient according to where in the physical item it tends to be used, such the outer shell of a tablet, inner body or ink.
	Group *CodeableConcept `bson:"group,omitempty" json:"group,omitempty" xml:"group,omitempty"`
	// If the ingredient is a known or suspected allergen.
	//
	// Note that this is a property of the substance, so if a reference to a SubstanceDefinition is used to decribe that (rather than just a code), the allergen information should go there, not here.
	AllergenicIndicator *bool `bson:"allergenicIndicator,omitempty" json:"allergenicIndicator,omitempty" xml:"allergenicIndicator,omitempty"`
	// A place for providing any notes that are relevant to the component, e.g.
	//
	// removed during process, adjusted for loss on drying.
	Comment *string `bson:"comment,omitempty" json:"comment,omitempty" xml:"comment,omitempty"`
	// The organization(s) that manufacture this ingredient.
	//
	// Can be used to indicate:         1) Organizations we are aware of that manufacture this ingredient         2) Specific Manufacturer(s) currently being used         3) Set of organisations allowed to manufacture this ingredient for this product         Users must be clear on the application of context relevant to their use case.
	Manufacturer []IngredientManufacturer `bson:"manufacturer,omitempty" json:"manufacturer,omitempty" xml:"manufacturer,omitempty"`
	// The substance that comprises this ingredient.
	Substance IngredientSubstance `bson:"substance" json:"substance" xml:"substance"`
}

type IngredientManufacturer struct {
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
	// The way in which this manufacturer is associated with the ingredient.
	//
	// For example whether it is a possible one (others allowed), or an exclusive authorized one for this ingredient. Note that this is not the manufacturing process role.
	Role *IngredientManufacturerRole `bson:"role,omitempty" json:"role,omitempty" xml:"role,omitempty"`
	// An organization that manufactures this ingredient.
	Manufacturer Reference `bson:"manufacturer" json:"manufacturer" xml:"manufacturer"`
}

type IngredientSubstance struct {
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
	// A code or full resource that represents the ingredient's substance.
	Code CodeableReference `bson:"code" json:"code" xml:"code"`
	// The quantity of substance in the unit of presentation, or in the volume (or mass) of the single pharmaceutical product or manufactured item.
	//
	// The allowed repetitions do not represent different strengths, but are different representations - mathematically equivalent - of a single strength.
	Strength []IngredientSubstanceStrength `bson:"strength,omitempty" json:"strength,omitempty" xml:"strength,omitempty"`
}

type IngredientSubstanceStrength struct {
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
	// The quantity of substance in the unit of presentation, or in the volume (or mass) of the single pharmaceutical product or manufactured item.
	//
	// Unit of presentation refers to the quantity that the item occurs in e.g. a strength per tablet size, perhaps 'per 20mg' (the size of the tablet). It is not generally normalized as a unitary unit, which would be 'per mg').
	PresentationRatio           *Ratio           `bson:"presentationRatio,omitempty" json:"presentationRatio,omitempty" xml:"presentationRatio,omitempty"`
	PresentationRatioRange      *RatioRange      `bson:"presentationRatioRange,omitempty" json:"presentationRatioRange,omitempty" xml:"presentationRatioRange,omitempty"`
	PresentationCodeableConcept *CodeableConcept `bson:"presentationCodeableConcept,omitempty" json:"presentationCodeableConcept,omitempty" xml:"presentationCodeableConcept,omitempty"`
	PresentationQuantity        *Quantity        `bson:"presentationQuantity,omitempty" json:"presentationQuantity,omitempty" xml:"presentationQuantity,omitempty"`
	// A textual represention of either the whole of the presentation strength or a part of it - with the rest being in Strength.presentation as a ratio.
	TextPresentation *string `bson:"textPresentation,omitempty" json:"textPresentation,omitempty" xml:"textPresentation,omitempty"`
	// The strength per unitary volume (or mass).
	ConcentrationRatio           *Ratio           `bson:"concentrationRatio,omitempty" json:"concentrationRatio,omitempty" xml:"concentrationRatio,omitempty"`
	ConcentrationRatioRange      *RatioRange      `bson:"concentrationRatioRange,omitempty" json:"concentrationRatioRange,omitempty" xml:"concentrationRatioRange,omitempty"`
	ConcentrationCodeableConcept *CodeableConcept `bson:"concentrationCodeableConcept,omitempty" json:"concentrationCodeableConcept,omitempty" xml:"concentrationCodeableConcept,omitempty"`
	ConcentrationQuantity        *Quantity        `bson:"concentrationQuantity,omitempty" json:"concentrationQuantity,omitempty" xml:"concentrationQuantity,omitempty"`
	// A textual represention of either the whole of the concentration strength or a part of it - with the rest being in Strength.concentration as a ratio.
	TextConcentration *string `bson:"textConcentration,omitempty" json:"textConcentration,omitempty" xml:"textConcentration,omitempty"`
	// A code that indicates if the strength is, for example, based on the ingredient substance as stated or on the substance base (when the ingredient is a salt).
	Basis *CodeableConcept `bson:"basis,omitempty" json:"basis,omitempty" xml:"basis,omitempty"`
	// For when strength is measured at a particular point or distance.
	//
	// There are products where strength is measured at a particular point. For example, the strength of the ingredient in some inhalers is measured at a particular position relative to the point of aerosolization.
	MeasurementPoint *string `bson:"measurementPoint,omitempty" json:"measurementPoint,omitempty" xml:"measurementPoint,omitempty"`
	// The country or countries for which the strength range applies.
	Country []CodeableConcept `bson:"country,omitempty" json:"country,omitempty" xml:"country,omitempty"`
	// Strength expressed in terms of a reference substance.
	//
	// For when the ingredient strength is additionally expressed as equivalent to the strength of some other closely related substance (e.g. salt vs. base). Reference strength represents the strength (quantitative composition) of the active moiety of the active substance. There are situations when the active substance and active moiety are different, therefore both a strength and a reference strength are needed.
	ReferenceStrength []IngredientSubstanceStrengthReferenceStrength `bson:"referenceStrength,omitempty" json:"referenceStrength,omitempty" xml:"referenceStrength,omitempty"`
}

type IngredientSubstanceStrengthReferenceStrength struct {
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
	// Relevant reference substance.
	Substance CodeableReference `bson:"substance" json:"substance" xml:"substance"`
	// Strength expressed in terms of a reference substance.
	StrengthRatio      Ratio      `bson:"strengthRatio" json:"strengthRatio" xml:"strengthRatio"`
	StrengthRatioRange RatioRange `bson:"strengthRatioRange" json:"strengthRatioRange" xml:"strengthRatioRange"`
	StrengthQuantity   Quantity   `bson:"strengthQuantity" json:"strengthQuantity" xml:"strengthQuantity"`
	// For when strength is measured at a particular point or distance.
	MeasurementPoint *string `bson:"measurementPoint,omitempty" json:"measurementPoint,omitempty" xml:"measurementPoint,omitempty"`
	// The country or countries for which the strength range applies.
	Country []CodeableConcept `bson:"country,omitempty" json:"country,omitempty" xml:"country,omitempty"`
}

func (r Ingredient) ResourceType() string {
	return "Ingredient"
}

// MarshalJSON implements [json.Marshaler]
func (r Ingredient) MarshalJSON() ([]byte, error) {
	type Wrapper Ingredient
	return marshalJSON(r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Marshaler].
func (r Ingredient) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Wrapper Ingredient
	return marshalXML(e, start, r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Unmarshaler].
func (r *Ingredient) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper Ingredient
	w := Wrapper(*r)
	err := unmarshalXML(d, start, r.ResourceType(), &w)
	*r = Ingredient(w)
	return err
}
