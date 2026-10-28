package fhires

import "encoding/json/v2"

// Representation of a molecular sequence.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/MolecularSequence
type MolecularSequence struct {
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
	// A unique identifier for this particular sequence instance.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Amino Acid Sequence/ DNA Sequence / RNA Sequence.
	Type *SequenceType `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Indicates the subject this sequence is associated too.
	Subject *Reference `bson:"subject,omitempty" json:"subject,omitempty" xml:"subject,omitempty"`
	// The actual focus of a molecular sequence when it is not the patient of record representing something or someone associated with the patient such as a spouse, parent, child, or sibling.
	//
	// For example, in trio testing, the subject would be the child (proband) and the focus would be the parent.
	Focus []Reference `bson:"focus,omitempty" json:"focus,omitempty" xml:"focus,omitempty"`
	// Specimen used for sequencing.
	Specimen *Reference `bson:"specimen,omitempty" json:"specimen,omitempty" xml:"specimen,omitempty"`
	// The method for sequencing, for example, chip information.
	Device *Reference `bson:"device,omitempty" json:"device,omitempty" xml:"device,omitempty"`
	// The organization or lab that should be responsible for this result.
	Performer *Reference `bson:"performer,omitempty" json:"performer,omitempty" xml:"performer,omitempty"`
	// Sequence that was observed.
	Literal *string `bson:"literal,omitempty" json:"literal,omitempty" xml:"literal,omitempty"`
	// Sequence that was observed as file content.
	//
	// Can be an actual file contents, or referenced by a URL to an external system.
	Formatted []Attachment `bson:"formatted,omitempty" json:"formatted,omitempty" xml:"formatted,omitempty"`
	// A sequence defined relative to another sequence.
	Relative []MolecularSequenceRelative `bson:"relative,omitempty" json:"relative,omitempty" xml:"relative,omitempty"`
}

type MolecularSequenceRelative struct {
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
	// These are different ways of identifying nucleotides or amino acids within a sequence.
	//
	// Different databases and file types may use different systems. For detail definitions, see https://loinc.org/92822-6/ for more detail.
	CoordinateSystem CodeableConcept `bson:"coordinateSystem" json:"coordinateSystem" xml:"coordinateSystem"`
	// Indicates the order in which the sequence should be considered when putting multiple 'relative' elements together.
	OrdinalPosition *int `bson:"ordinalPosition,omitempty" json:"ordinalPosition,omitempty" xml:"ordinalPosition,omitempty"`
	// Indicates the nucleotide range in the composed sequence when multiple 'relative' elements are used together.
	SequenceRange *Range `bson:"sequenceRange,omitempty" json:"sequenceRange,omitempty" xml:"sequenceRange,omitempty"`
	// A sequence that is used as a starting sequence to describe variants that are present in a sequence analyzed.
	StartingSequence *MolecularSequenceRelativeStartingSequence `bson:"startingSequence,omitempty" json:"startingSequence,omitempty" xml:"startingSequence,omitempty"`
	// Changes in sequence from the starting sequence.
	Edit []MolecularSequenceRelativeEdit `bson:"edit,omitempty" json:"edit,omitempty" xml:"edit,omitempty"`
}

type MolecularSequenceRelativeStartingSequence struct {
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
	// The genome assembly used for starting sequence, e.g. GRCh38.
	GenomeAssembly *CodeableConcept `bson:"genomeAssembly,omitempty" json:"genomeAssembly,omitempty" xml:"genomeAssembly,omitempty"`
	// Structural unit composed of a nucleic acid molecule which controls its own replication through the interaction of specific proteins at one or more origins of replication ([SO:0000340](http://www.sequenceontology.org/browser/current_svn/term/SO:0000340)).
	Chromosome *CodeableConcept `bson:"chromosome,omitempty" json:"chromosome,omitempty" xml:"chromosome,omitempty"`
	// The reference sequence that represents the starting sequence.
	//
	// Comment: A starting sequence may be represented in one of three ways:
	//
	//   1.  CodeableConcept, using NCBI, LRG or other appropriate code systems
	//
	//   1.  a simple string of IUPAC codes
	//
	//   1.  a reference to another MolecularSequence resource.
	SequenceCodeableConcept *CodeableConcept `bson:"sequenceCodeableConcept,omitempty" json:"sequenceCodeableConcept,omitempty" xml:"sequenceCodeableConcept,omitempty"`
	SequenceString          *string          `bson:"sequenceString,omitempty" json:"sequenceString,omitempty" xml:"sequenceString,omitempty"`
	SequenceReference       *Reference       `bson:"sequenceReference,omitempty" json:"sequenceReference,omitempty" xml:"sequenceReference,omitempty"`
	// Start position of the window on the starting sequence.
	//
	// This value should honor the rules of the coordinateSystem.
	WindowStart *int `bson:"windowStart,omitempty" json:"windowStart,omitempty" xml:"windowStart,omitempty"`
	// End position of the window on the starting sequence.
	//
	// This value should honor the rules of the  coordinateSystem.
	WindowEnd *int `bson:"windowEnd,omitempty" json:"windowEnd,omitempty" xml:"windowEnd,omitempty"`
	// A relative reference to a DNA strand based on gene orientation.
	//
	// The strand that contains the open reading frame of the gene is the "sense" strand, and the opposite complementary strand is the "antisense" strand.
	Orientation *OrientationType `bson:"orientation,omitempty" json:"orientation,omitempty" xml:"orientation,omitempty"`
	// An absolute reference to a strand.
	//
	// The Watson strand is the strand whose 5'-end is on the short arm of the chromosome, and the Crick strand as the one whose 5'-end is on the long arm.
	Strand *StrandType `bson:"strand,omitempty" json:"strand,omitempty" xml:"strand,omitempty"`
}

type MolecularSequenceRelativeEdit struct {
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
	// Start position of the edit on the starting sequence.
	//
	// If the coordinate system is either 0-based or 1-based, then start position is inclusive.
	Start *int `bson:"start,omitempty" json:"start,omitempty" xml:"start,omitempty"`
	// End position of the edit on the starting sequence.
	//
	// If the coordinate system is 0-based then end is exclusive and does not include the last position. If the coordinate system is 1-base, then end is inclusive and includes the last position.
	End *int `bson:"end,omitempty" json:"end,omitempty" xml:"end,omitempty"`
	// Allele that was observed.
	//
	// Nucleotide(s)/amino acids from start position of sequence to stop position of sequence on the positive (+) strand of the observed sequence. When the sequence type is DNA, it should be the sequence on the positive (+) strand. This will lay in the range between variant.start and variant.end.
	ReplacementSequence *string `bson:"replacementSequence,omitempty" json:"replacementSequence,omitempty" xml:"replacementSequence,omitempty"`
	// Allele in the starting sequence.
	//
	// Nucleotide(s)/amino acids from start position of sequence to stop position of sequence on the positive (+) strand of the starting sequence. When the sequence  type is DNA, it should be the sequence on the positive (+) strand. This will lay in the range between variant.start and variant.end.
	ReplacedSequence *string `bson:"replacedSequence,omitempty" json:"replacedSequence,omitempty" xml:"replacedSequence,omitempty"`
}

func (r MolecularSequence) ResourceType() string {
	return "MolecularSequence"
}

// MarshalJSON implements json.Marshaler
func (r MolecularSequence) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		MolecularSequence
		ResourceType string `json:"resourceType"`
	}{
		MolecularSequence: r,
		ResourceType:      r.ResourceType(),
	})
}
