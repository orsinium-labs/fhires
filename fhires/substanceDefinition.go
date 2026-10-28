package fhires

import v2 "encoding/json/v2"

// The detailed description of a substance, typically at a level beyond what is used for prescribing.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/SubstanceDefinition
type SubstanceDefinition struct {
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
	// Identifier by which this substance is known.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// A business level version identifier of the substance.
	Version *string `bson:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
	// Status of substance within the catalogue e.g. active, retired.
	Status *CodeableConcept `bson:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// A high level categorization, e.g.
	//
	// polymer or nucleic acid, or food, chemical, biological, or a lower level such as the general types of polymer (linear or branch chain) or type of impurity (process related or contaminant).
	Classification []CodeableConcept `bson:"classification,omitempty" json:"classification,omitempty" xml:"classification,omitempty"`
	// If the substance applies to human or veterinary use.
	Domain *CodeableConcept `bson:"domain,omitempty" json:"domain,omitempty" xml:"domain,omitempty"`
	// The quality standard, established benchmark, to which substance complies (e.g.
	//
	// USP/NF, Ph. Eur, JP, BP, Company Standard).
	Grade []CodeableConcept `bson:"grade,omitempty" json:"grade,omitempty" xml:"grade,omitempty"`
	// Textual description of the substance.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Supporting literature.
	InformationSource []Reference `bson:"informationSource,omitempty" json:"informationSource,omitempty" xml:"informationSource,omitempty"`
	// Textual comment about the substance's catalogue or registry record.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// The entity that creates, makes, produces or fabricates the substance.
	//
	// This is a set of potential manufacturers but is not necessarily comprehensive.
	Manufacturer []Reference `bson:"manufacturer,omitempty" json:"manufacturer,omitempty" xml:"manufacturer,omitempty"`
	// An entity that is the source for the substance.
	//
	// It may be different from the manufacturer. Supplier is synonymous to a distributor.
	Supplier []Reference `bson:"supplier,omitempty" json:"supplier,omitempty" xml:"supplier,omitempty"`
	// Moiety, for structural modifications.
	Moiety []SubstanceDefinitionMoiety `bson:"moiety,omitempty" json:"moiety,omitempty" xml:"moiety,omitempty"`
	// General specifications for this substance.
	Characterization []SubstanceDefinitionCharacterization `bson:"characterization,omitempty" json:"characterization,omitempty" xml:"characterization,omitempty"`
	// General specifications for this substance.
	Property []SubstanceDefinitionProperty `bson:"property,omitempty" json:"property,omitempty" xml:"property,omitempty"`
	// General information detailing this substance.
	ReferenceInformation *Reference `bson:"referenceInformation,omitempty" json:"referenceInformation,omitempty" xml:"referenceInformation,omitempty"`
	// The average mass of a molecule of a compound compared to 1/12 the mass of carbon 12 and calculated as the sum of the atomic weights of the constituent atoms.
	MolecularWeight []SubstanceDefinitionMolecularWeight `bson:"molecularWeight,omitempty" json:"molecularWeight,omitempty" xml:"molecularWeight,omitempty"`
	// Structural information.
	Structure *SubstanceDefinitionStructure `bson:"structure,omitempty" json:"structure,omitempty" xml:"structure,omitempty"`
	// Codes associated with the substance.
	Code []SubstanceDefinitionCode `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// Names applicable to this substance.
	Name []SubstanceDefinitionName `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// A link between this substance and another, with details of the relationship.
	Relationship []SubstanceDefinitionRelationship `bson:"relationship,omitempty" json:"relationship,omitempty" xml:"relationship,omitempty"`
	// Data items specific to nucleic acids.
	NucleicAcid *Reference `bson:"nucleicAcid,omitempty" json:"nucleicAcid,omitempty" xml:"nucleicAcid,omitempty"`
	// Data items specific to polymers.
	Polymer *Reference `bson:"polymer,omitempty" json:"polymer,omitempty" xml:"polymer,omitempty"`
	// Data items specific to proteins.
	Protein *Reference `bson:"protein,omitempty" json:"protein,omitempty" xml:"protein,omitempty"`
	// Material or taxonomic/anatomical source for the substance.
	SourceMaterial *SubstanceDefinitionSourceMaterial `bson:"sourceMaterial,omitempty" json:"sourceMaterial,omitempty" xml:"sourceMaterial,omitempty"`
}

type SubstanceDefinitionMoiety struct {
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
	// Role that the moiety is playing.
	Role *CodeableConcept `bson:"role,omitempty" json:"role,omitempty" xml:"role,omitempty"`
	// Identifier by which this moiety substance is known.
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Textual name for this moiety substance.
	Name *string `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Stereochemistry type.
	Stereochemistry *CodeableConcept `bson:"stereochemistry,omitempty" json:"stereochemistry,omitempty" xml:"stereochemistry,omitempty"`
	// Optical activity type.
	OpticalActivity *CodeableConcept `bson:"opticalActivity,omitempty" json:"opticalActivity,omitempty" xml:"opticalActivity,omitempty"`
	// Molecular formula for this moiety of this substance, typically using the Hill system.
	MolecularFormula *string `bson:"molecularFormula,omitempty" json:"molecularFormula,omitempty" xml:"molecularFormula,omitempty"`
	// Quantitative value for this moiety.
	AmountQuantity *Quantity `bson:"amountQuantity,omitempty" json:"amountQuantity,omitempty" xml:"amountQuantity,omitempty"`
	AmountString   *string   `bson:"amountString,omitempty" json:"amountString,omitempty" xml:"amountString,omitempty"`
	// The measurement type of the quantitative value.
	//
	// In capturing the actual relative amounts of substances or molecular fragments it may be necessary to indicate whether the amount refers to, for example, a mole ratio or weight ratio.
	MeasurementType *CodeableConcept `bson:"measurementType,omitempty" json:"measurementType,omitempty" xml:"measurementType,omitempty"`
}

type SubstanceDefinitionCharacterization struct {
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
	// The method used to elucidate the characterization of the drug substance. Example: HPLC.
	Technique *CodeableConcept `bson:"technique,omitempty" json:"technique,omitempty" xml:"technique,omitempty"`
	// Describes the nature of the chemical entity and explains, for instance, whether this is a base or a salt form.
	Form *CodeableConcept `bson:"form,omitempty" json:"form,omitempty" xml:"form,omitempty"`
	// The description or justification in support of the interpretation of the data file.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// The data produced by the analytical instrument or a pictorial representation of that data.
	//
	// Examples: a JCAMP, JDX, or ADX file, or a chromatogram or spectrum analysis.
	File []Attachment `bson:"file,omitempty" json:"file,omitempty" xml:"file,omitempty"`
}

type SubstanceDefinitionProperty struct {
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
	// A code expressing the type of property.
	Type CodeableConcept `bson:"type" json:"type" xml:"type"`
	// A value for the property.
	ValueCodeableConcept *CodeableConcept `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty" xml:"valueCodeableConcept,omitempty"`
	ValueQuantity        *Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty" xml:"valueQuantity,omitempty"`
	ValueDate            *string          `bson:"valueDate,omitempty" json:"valueDate,omitempty" xml:"valueDate,omitempty"`
	ValueBoolean         *bool            `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty" xml:"valueBoolean,omitempty"`
	ValueAttachment      *Attachment      `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty" xml:"valueAttachment,omitempty"`
}

type SubstanceDefinitionMolecularWeight struct {
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
	// The method by which the molecular weight was determined.
	Method *CodeableConcept `bson:"method,omitempty" json:"method,omitempty" xml:"method,omitempty"`
	// Type of molecular weight such as exact, average (also known as.
	//
	// number average), weight average.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Used to capture quantitative values for a variety of elements.
	//
	// If only limits are given, the arithmetic mean would be the average. If only a single definite value for a given element is given, it would be captured in this field.
	Amount Quantity `bson:"amount" json:"amount" xml:"amount"`
}

type SubstanceDefinitionStructure struct {
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
	// Stereochemistry type.
	Stereochemistry *CodeableConcept `bson:"stereochemistry,omitempty" json:"stereochemistry,omitempty" xml:"stereochemistry,omitempty"`
	// Optical activity type.
	OpticalActivity *CodeableConcept `bson:"opticalActivity,omitempty" json:"opticalActivity,omitempty" xml:"opticalActivity,omitempty"`
	// An expression which states the number and type of atoms present in a molecule of a substance.
	MolecularFormula *string `bson:"molecularFormula,omitempty" json:"molecularFormula,omitempty" xml:"molecularFormula,omitempty"`
	// Specified per moiety according to the Hill system, i.e.
	//
	// first C, then H, then alphabetical, each moiety separated by a dot.
	MolecularFormulaByMoiety *string `bson:"molecularFormulaByMoiety,omitempty" json:"molecularFormulaByMoiety,omitempty" xml:"molecularFormulaByMoiety,omitempty"`
	// The molecular weight or weight range (for proteins, polymers or nucleic acids).
	MolecularWeight *SubstanceDefinitionMolecularWeight `bson:"molecularWeight,omitempty" json:"molecularWeight,omitempty"`
	// The method used to elucidate the structure of the drug substance.
	//
	// Examples: X-ray, NMR, Peptide mapping, Ligand binding assay.
	Technique []CodeableConcept `bson:"technique,omitempty" json:"technique,omitempty" xml:"technique,omitempty"`
	// The source of information about the structure.
	SourceDocument []Reference `bson:"sourceDocument,omitempty" json:"sourceDocument,omitempty" xml:"sourceDocument,omitempty"`
	// A depiction of the structure of the substance.
	Representation []SubstanceDefinitionStructureRepresentation `bson:"representation,omitempty" json:"representation,omitempty" xml:"representation,omitempty"`
}

type SubstanceDefinitionStructureRepresentation struct {
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
	// The kind of structural representation (e.g. full, partial).
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The structural representation as a text string in a standard format.
	Representation *string `bson:"representation,omitempty" json:"representation,omitempty" xml:"representation,omitempty"`
	// The format of the representation e.g.
	//
	// InChI, SMILES, MOLFILE, CDX, SDF, PDB, mmCIF. The logical content type rather than the physical file format of a document.
	Format *CodeableConcept `bson:"format,omitempty" json:"format,omitempty" xml:"format,omitempty"`
	// An attached file with the structural representation e.g.
	//
	// a molecular structure graphic of the substance, a JCAMP or AnIML file.
	Document *Reference `bson:"document,omitempty" json:"document,omitempty" xml:"document,omitempty"`
}

type SubstanceDefinitionCode struct {
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
	// The specific code.
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// Status of the code assignment, for example 'provisional', 'approved'.
	Status *CodeableConcept `bson:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// The date at which the code status was changed as part of the terminology maintenance.
	StatusDate *string `bson:"statusDate,omitempty" json:"statusDate,omitempty" xml:"statusDate,omitempty"`
	// Any comment can be provided in this field, if necessary.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// Supporting literature.
	Source []Reference `bson:"source,omitempty" json:"source,omitempty" xml:"source,omitempty"`
}

type SubstanceDefinitionName struct {
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
	// The actual name.
	Name string `bson:"name" json:"name" xml:"name"`
	// Name type, for example 'systematic',  'scientific, 'brand'.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The status of the name, for example 'current', 'proposed'.
	Status *CodeableConcept `bson:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// If this is the preferred name for this substance.
	Preferred *bool `bson:"preferred,omitempty" json:"preferred,omitempty" xml:"preferred,omitempty"`
	// Human language that the name is written in.
	Language []CodeableConcept `bson:"language,omitempty" json:"language,omitempty" xml:"language,omitempty"`
	// The use context of this name for example if there is a different name a drug active ingredient as opposed to a food colour additive.
	Domain []CodeableConcept `bson:"domain,omitempty" json:"domain,omitempty" xml:"domain,omitempty"`
	// The jurisdiction where this name applies.
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty" xml:"jurisdiction,omitempty"`
	// A synonym of this particular name, by which the substance is also known.
	Synonym []SubstanceDefinitionName `bson:"synonym,omitempty" json:"synonym,omitempty"`
	// A translation for this name into another human language.
	Translation []SubstanceDefinitionName `bson:"translation,omitempty" json:"translation,omitempty"`
	// Details of the official nature of this name.
	Official []SubstanceDefinitionNameOfficial `bson:"official,omitempty" json:"official,omitempty" xml:"official,omitempty"`
	// Supporting literature.
	Source []Reference `bson:"source,omitempty" json:"source,omitempty" xml:"source,omitempty"`
}

type SubstanceDefinitionNameOfficial struct {
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
	// Which authority uses this official name.
	Authority *CodeableConcept `bson:"authority,omitempty" json:"authority,omitempty" xml:"authority,omitempty"`
	// The status of the official name, for example 'draft', 'active', 'retired'.
	Status *CodeableConcept `bson:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Date of the official name change.
	Date *string `bson:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
}

type SubstanceDefinitionRelationship struct {
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
	// A pointer to another substance, as a resource or just a representational code.
	SubstanceDefinitionReference       *Reference       `bson:"substanceDefinitionReference,omitempty" json:"substanceDefinitionReference,omitempty" xml:"substanceDefinitionReference,omitempty"`
	SubstanceDefinitionCodeableConcept *CodeableConcept `bson:"substanceDefinitionCodeableConcept,omitempty" json:"substanceDefinitionCodeableConcept,omitempty" xml:"substanceDefinitionCodeableConcept,omitempty"`
	// For example "salt to parent", "active moiety", "starting material", "polymorph", "impurity of".
	Type CodeableConcept `bson:"type" json:"type" xml:"type"`
	// For example where an enzyme strongly bonds with a particular substance, this is a defining relationship for that enzyme, out of several possible substance relationships.
	IsDefining *bool `bson:"isDefining,omitempty" json:"isDefining,omitempty" xml:"isDefining,omitempty"`
	// A numeric factor for the relationship, for instance to express that the salt of a substance has some percentage of the active substance in relation to some other.
	AmountQuantity *Quantity `bson:"amountQuantity,omitempty" json:"amountQuantity,omitempty" xml:"amountQuantity,omitempty"`
	AmountRatio    *Ratio    `bson:"amountRatio,omitempty" json:"amountRatio,omitempty" xml:"amountRatio,omitempty"`
	AmountString   *string   `bson:"amountString,omitempty" json:"amountString,omitempty" xml:"amountString,omitempty"`
	// For use when the numeric has an uncertain range.
	RatioHighLimitAmount *Ratio `bson:"ratioHighLimitAmount,omitempty" json:"ratioHighLimitAmount,omitempty" xml:"ratioHighLimitAmount,omitempty"`
	// An operator for the amount, for example "average", "approximately", "less than".
	Comparator *CodeableConcept `bson:"comparator,omitempty" json:"comparator,omitempty" xml:"comparator,omitempty"`
	// Supporting literature.
	Source []Reference `bson:"source,omitempty" json:"source,omitempty" xml:"source,omitempty"`
}

type SubstanceDefinitionSourceMaterial struct {
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
	// A classification that provides the origin of the raw material.
	//
	// Example: cat hair would be an Animal source type.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The genus of an organism, typically referring to the Latin epithet of the genus element of the plant/animal scientific name.
	Genus *CodeableConcept `bson:"genus,omitempty" json:"genus,omitempty" xml:"genus,omitempty"`
	// The species of an organism, typically referring to the Latin epithet of the species of the plant/animal.
	Species *CodeableConcept `bson:"species,omitempty" json:"species,omitempty" xml:"species,omitempty"`
	// An anatomical origin of the source material within an organism.
	Part *CodeableConcept `bson:"part,omitempty" json:"part,omitempty" xml:"part,omitempty"`
	// The country or countries where the material is harvested.
	CountryOfOrigin []CodeableConcept `bson:"countryOfOrigin,omitempty" json:"countryOfOrigin,omitempty" xml:"countryOfOrigin,omitempty"`
}

func (r SubstanceDefinition) ResourceType() string {
	return "SubstanceDefinition"
}

// MarshalJSON implements json.Marshaler
func (r SubstanceDefinition) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		SubstanceDefinition
		ResourceType string `json:"resourceType"`
	}{
		ResourceType:        r.ResourceType(),
		SubstanceDefinition: r,
	})
}
