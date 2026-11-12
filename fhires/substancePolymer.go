package fhires

import "encoding/xml"

// Properties of a substance specific to it being a polymer.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/SubstancePolymer
type SubstancePolymer struct {
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
	// A business idenfier for this polymer, but typically this is handled by a SubstanceDefinition identifier.
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Overall type of the polymer.
	Class *CodeableConcept `bson:"class,omitempty" json:"class,omitempty" xml:"class,omitempty"`
	// Polymer geometry, e.g. linear, branched, cross-linked, network or dendritic.
	Geometry *CodeableConcept `bson:"geometry,omitempty" json:"geometry,omitempty" xml:"geometry,omitempty"`
	// Descrtibes the copolymer sequence type (polymer connectivity).
	CopolymerConnectivity []CodeableConcept `bson:"copolymerConnectivity,omitempty" json:"copolymerConnectivity,omitempty" xml:"copolymerConnectivity,omitempty"`
	// Todo - this is intended to connect to a repeating full modification structure, also used by Protein and Nucleic Acid .
	//
	// String is just a placeholder.
	Modification *string `bson:"modification,omitempty" json:"modification,omitempty" xml:"modification,omitempty"`
	// Todo.
	MonomerSet []SubstancePolymerMonomerSet `bson:"monomerSet,omitempty" json:"monomerSet,omitempty" xml:"monomerSet,omitempty"`
	// Specifies and quantifies the repeated units and their configuration.
	Repeat []SubstancePolymerRepeat `bson:"repeat,omitempty" json:"repeat,omitempty" xml:"repeat,omitempty"`
}

type SubstancePolymerMonomerSet struct {
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
	// Captures the type of ratio to the entire polymer, e.g.
	//
	// Monomer/Polymer ratio, SRU/Polymer Ratio.
	RatioType *CodeableConcept `bson:"ratioType,omitempty" json:"ratioType,omitempty" xml:"ratioType,omitempty"`
	// The starting materials - monomer(s) used in the synthesis of the polymer.
	StartingMaterial []SubstancePolymerMonomerSetStartingMaterial `bson:"startingMaterial,omitempty" json:"startingMaterial,omitempty" xml:"startingMaterial,omitempty"`
}

type SubstancePolymerMonomerSetStartingMaterial struct {
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
	// The type of substance for this starting material.
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// Substance high level category, e.g. chemical substance.
	Category *CodeableConcept `bson:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
	// Used to specify whether the attribute described is a defining element for the unique identification of the polymer.
	IsDefining *bool `bson:"isDefining,omitempty" json:"isDefining,omitempty" xml:"isDefining,omitempty"`
	// A percentage.
	Amount *Quantity `bson:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
}

type SubstancePolymerRepeat struct {
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
	// A representation of an (average) molecular formula from a polymer.
	AverageMolecularFormula *string `bson:"averageMolecularFormula,omitempty" json:"averageMolecularFormula,omitempty" xml:"averageMolecularFormula,omitempty"`
	// How the quantitative amount of Structural Repeat Units is captured (e.g.
	//
	// Exact, Numeric, Average).
	RepeatUnitAmountType *CodeableConcept `bson:"repeatUnitAmountType,omitempty" json:"repeatUnitAmountType,omitempty" xml:"repeatUnitAmountType,omitempty"`
	// An SRU - Structural Repeat Unit.
	RepeatUnit []SubstancePolymerRepeatRepeatUnit `bson:"repeatUnit,omitempty" json:"repeatUnit,omitempty" xml:"repeatUnit,omitempty"`
}

type SubstancePolymerRepeatRepeatUnit struct {
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
	// Structural repeat units are essential elements for defining polymers.
	Unit *string `bson:"unit,omitempty" json:"unit,omitempty" xml:"unit,omitempty"`
	// The orientation of the polymerisation, e.g. head-tail, head-head, random.
	Orientation *CodeableConcept `bson:"orientation,omitempty" json:"orientation,omitempty" xml:"orientation,omitempty"`
	// Number of repeats of this unit.
	Amount *int `bson:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
	// Applies to homopolymer and block co-polymers where the degree of polymerisation within a block can be described.
	DegreeOfPolymerisation []SubstancePolymerRepeatRepeatUnitDegreeOfPolymerisation `bson:"degreeOfPolymerisation,omitempty" json:"degreeOfPolymerisation,omitempty" xml:"degreeOfPolymerisation,omitempty"`
	// A graphical structure for this SRU.
	StructuralRepresentation []SubstancePolymerRepeatRepeatUnitStructuralRepresentation `bson:"structuralRepresentation,omitempty" json:"structuralRepresentation,omitempty" xml:"structuralRepresentation,omitempty"`
}

type SubstancePolymerRepeatRepeatUnitDegreeOfPolymerisation struct {
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
	// The type of the degree of polymerisation shall be described, e.g. SRU/Polymer Ratio.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// An average amount of polymerisation.
	Average *int `bson:"average,omitempty" json:"average,omitempty" xml:"average,omitempty"`
	// A low expected limit of the amount.
	Low *int `bson:"low,omitempty" json:"low,omitempty" xml:"low,omitempty"`
	// A high expected limit of the amount.
	High *int `bson:"high,omitempty" json:"high,omitempty" xml:"high,omitempty"`
}

type SubstancePolymerRepeatRepeatUnitStructuralRepresentation struct {
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
	// The type of structure (e.g. Full, Partial, Representative).
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The structural representation as text string in a standard format e.g.
	//
	// InChI, SMILES, MOLFILE, CDX, SDF, PDB, mmCIF.
	Representation *string `bson:"representation,omitempty" json:"representation,omitempty" xml:"representation,omitempty"`
	// The format of the representation e.g. InChI, SMILES, MOLFILE, CDX, SDF, PDB, mmCIF.
	Format *CodeableConcept `bson:"format,omitempty" json:"format,omitempty" xml:"format,omitempty"`
	// An attached file with the structural representation.
	Attachment *Attachment `bson:"attachment,omitempty" json:"attachment,omitempty" xml:"attachment,omitempty"`
}

func (r SubstancePolymer) ResourceType() string {
	return "SubstancePolymer"
}

// MarshalJSON implements json.Marshaler
func (r SubstancePolymer) MarshalJSON() ([]byte, error) {
	type Wrapper SubstancePolymer
	return marshalJSON(r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Unmarshaler].
func (r *SubstancePolymer) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper SubstancePolymer
	w := Wrapper(*r)
	err := unmarshalXML(d, start, &w)
	*r = SubstancePolymer(w)
	return err
}
