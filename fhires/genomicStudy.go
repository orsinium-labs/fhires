package fhires

import "encoding/xml"

// A set of analyses performed to analyze and generate genomic data.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/GenomicStudy
type GenomicStudy struct {
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
	// Identifiers for this genomic study.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The status of the genomic study.
	Status GenomicStudyStatus `bson:"status" json:"status" xml:"status"`
	// The type of the study, e.g., Familial variant segregation, Functional variation detection, or Gene expression profiling.
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The primary subject of the genomic study.
	Subject Reference `bson:"subject" json:"subject" xml:"subject"`
	// The healthcare event with which this genomics study is associated.
	Encounter *Reference `bson:"encounter,omitempty" json:"encounter,omitempty" xml:"encounter,omitempty"`
	// When the genomic study was started.
	StartDate *string `bson:"startDate,omitempty" json:"startDate,omitempty" xml:"startDate,omitempty"`
	// Event resources that the genomic study is based on.
	BasedOn []Reference `bson:"basedOn,omitempty" json:"basedOn,omitempty" xml:"basedOn,omitempty"`
	// Healthcare professional who requested or referred the genomic study.
	Referrer *Reference `bson:"referrer,omitempty" json:"referrer,omitempty" xml:"referrer,omitempty"`
	// Healthcare professionals who interpreted the genomic study.
	Interpreter []Reference `bson:"interpreter,omitempty" json:"interpreter,omitempty" xml:"interpreter,omitempty"`
	// Why the genomic study was performed.
	Reason []CodeableReference `bson:"reason,omitempty" json:"reason,omitempty" xml:"reason,omitempty"`
	// The defined protocol that describes the study.
	InstantiatesCanonical *string `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty" xml:"instantiatesCanonical,omitempty"`
	// The URL pointing to an externally maintained protocol that describes the study.
	InstantiatesUri *string `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty" xml:"instantiatesUri,omitempty"`
	// Comments related to the genomic study.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// Description of the genomic study.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// The details about a specific analysis that was performed in this GenomicStudy.
	Analysis []GenomicStudyAnalysis `bson:"analysis,omitempty" json:"analysis,omitempty" xml:"analysis,omitempty"`
}

type GenomicStudyAnalysis struct {
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
	// Identifiers for the analysis event.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Type of the methods used in the analysis, e.g., Fluorescence in situ hybridization (FISH), Karyotyping, or Microsatellite instability testing (MSI).
	MethodType []CodeableConcept `bson:"methodType,omitempty" json:"methodType,omitempty" xml:"methodType,omitempty"`
	// Type of the genomic changes studied in the analysis, e.g., DNA, RNA, or amino acid change.
	ChangeType []CodeableConcept `bson:"changeType,omitempty" json:"changeType,omitempty" xml:"changeType,omitempty"`
	// The reference genome build that is used in this analysis.
	GenomeBuild *CodeableConcept `bson:"genomeBuild,omitempty" json:"genomeBuild,omitempty" xml:"genomeBuild,omitempty"`
	// The defined protocol that describes the analysis.
	InstantiatesCanonical *string `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty" xml:"instantiatesCanonical,omitempty"`
	// The URL pointing to an externally maintained protocol that describes the analysis.
	InstantiatesUri *string `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty" xml:"instantiatesUri,omitempty"`
	// Name of the analysis event (human friendly).
	Title *string `bson:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// The focus of a genomic analysis when it is not the patient of record representing something or someone associated with the patient such as a spouse, parent, child, or sibling.
	//
	// For example, in trio testing, the GenomicStudy.subject would be the child (proband) and the GenomicStudy.analysis.focus of a specific analysis would be the parent.
	//
	// Comment: If the focus of the analysis is not known, the value of this field SHALL use the data absent extension.
	Focus []Reference `bson:"focus,omitempty" json:"focus,omitempty" xml:"focus,omitempty"`
	// The specimen used in the analysis event.
	Specimen []Reference `bson:"specimen,omitempty" json:"specimen,omitempty" xml:"specimen,omitempty"`
	// The date of the analysis event.
	Date *string `bson:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// Any notes capture with the analysis event.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// The protocol that was performed for the analysis event.
	ProtocolPerformed *Reference `bson:"protocolPerformed,omitempty" json:"protocolPerformed,omitempty" xml:"protocolPerformed,omitempty"`
	// The genomic regions to be studied in the analysis (BED file).
	RegionsStudied []Reference `bson:"regionsStudied,omitempty" json:"regionsStudied,omitempty" xml:"regionsStudied,omitempty"`
	// Genomic regions actually called in the analysis event (BED file).
	RegionsCalled []Reference `bson:"regionsCalled,omitempty" json:"regionsCalled,omitempty" xml:"regionsCalled,omitempty"`
	// Inputs for the analysis event.
	Input []GenomicStudyAnalysisInput `bson:"input,omitempty" json:"input,omitempty" xml:"input,omitempty"`
	// Outputs for the analysis event.
	Output []GenomicStudyAnalysisOutput `bson:"output,omitempty" json:"output,omitempty" xml:"output,omitempty"`
	// Performer for the analysis event.
	Performer []GenomicStudyAnalysisPerformer `bson:"performer,omitempty" json:"performer,omitempty" xml:"performer,omitempty"`
	// Devices used for the analysis (e.g., instruments, software), with settings and parameters.
	Device []GenomicStudyAnalysisDevice `bson:"device,omitempty" json:"device,omitempty" xml:"device,omitempty"`
}

type GenomicStudyAnalysisInput struct {
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
	// File containing input data.
	File *Reference `bson:"file,omitempty" json:"file,omitempty" xml:"file,omitempty"`
	// Type of input data, e.g., BAM, CRAM, or FASTA.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The analysis event or other GenomicStudy that generated this input file.
	GeneratedByIdentifier *Identifier `bson:"generatedByIdentifier,omitempty" json:"generatedByIdentifier,omitempty" xml:"generatedByIdentifier,omitempty"`
	GeneratedByReference  *Reference  `bson:"generatedByReference,omitempty" json:"generatedByReference,omitempty" xml:"generatedByReference,omitempty"`
}

type GenomicStudyAnalysisOutput struct {
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
	// File containing output data.
	File *Reference `bson:"file,omitempty" json:"file,omitempty" xml:"file,omitempty"`
	// Type of output data, e.g., VCF, MAF, or BAM.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
}

type GenomicStudyAnalysisPerformer struct {
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
	// The organization, healthcare professional, or others who participated in performing this analysis.
	Actor *Reference `bson:"actor,omitempty" json:"actor,omitempty" xml:"actor,omitempty"`
	// Role of the actor for this analysis.
	Role *CodeableConcept `bson:"role,omitempty" json:"role,omitempty" xml:"role,omitempty"`
}

type GenomicStudyAnalysisDevice struct {
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
	// Device used for the analysis.
	Device *Reference `bson:"device,omitempty" json:"device,omitempty" xml:"device,omitempty"`
	// Specific function for the device used for the analysis.
	Function *CodeableConcept `bson:"function,omitempty" json:"function,omitempty" xml:"function,omitempty"`
}

func (r GenomicStudy) ResourceType() string {
	return "GenomicStudy"
}

// MarshalJSON implements [json.Marshaler]
func (r GenomicStudy) MarshalJSON() ([]byte, error) {
	type Wrapper GenomicStudy
	return marshalJSON(r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Marshaler].
func (r GenomicStudy) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Wrapper GenomicStudy
	return marshalXML(e, start, r.ResourceType(), Wrapper(r))
}


// UnmarshalXML implements [xml.Unmarshaler].
func (r *GenomicStudy) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper GenomicStudy
	w := Wrapper(*r)
	err := unmarshalXML(d, start, r.ResourceType(), &w)
	*r = GenomicStudy(w)
	return err
}
