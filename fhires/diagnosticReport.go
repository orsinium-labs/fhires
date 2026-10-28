package fhires

import v2 "encoding/json/v2"

// The findings and interpretation of diagnostic tests performed on patients, groups of patients, products, substances, devices, and locations, and/or specimens derived from these.
//
// The report includes clinical context such as requesting provider information, and some mix of atomic results, images, textual and coded interpretations, and formatted representation of diagnostic reports. The report also includes non-clinical context such as batch analysis and stability reporting of products and substances.
//
// Purpose: To support reporting for any diagnostic report into a clinical data repository.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/DiagnosticReport
type DiagnosticReport struct {
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
	// Identifiers assigned to this report by the performer or other systems.
	//
	// Comment: Usually assigned by the Information System of the diagnostic service provider (filler id).
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Details concerning a service requested.
	//
	// Comment: Note: Usually there is one test request for each result, however in some circumstances multiple test requests may be represented using a single test result resource. Note that there are also cases where one request leads to multiple reports.
	BasedOn []Reference `bson:"basedOn,omitempty" json:"basedOn,omitempty" xml:"basedOn,omitempty"`
	// The status of the diagnostic report.
	Status DiagnosticReportStatus `bson:"status" json:"status" xml:"status"`
	// A code that classifies the clinical discipline, department or diagnostic service that created the report (e.g.
	//
	// cardiology, biochemistry, hematology, MRI). This is used for searching, sorting and display purposes.
	//
	// Comment: Multiple categories are allowed using various categorization schemes.   The level of granularity is defined by the category concepts in the value set. More fine-grained filtering can be performed using the metadata and/or terminology hierarchy in DiagnosticReport.code.
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
	// A code or name that describes this diagnostic report.
	Code CodeableConcept `bson:"code" json:"code" xml:"code"`
	// The subject of the report.
	//
	// Usually, but not always, this is a patient. However, diagnostic services also perform analyses on specimens collected from a variety of other sources.
	Subject *Reference `bson:"subject,omitempty" json:"subject,omitempty" xml:"subject,omitempty"`
	// The healthcare event  (e.g.
	//
	// a patient and healthcare provider interaction) which this DiagnosticReport is about.
	//
	// Comment: This will typically be the encounter the event occurred within, but some events may be initiated prior to or after the official completion of an encounter  but still be tied to the context of the encounter  (e.g. pre-admission laboratory tests).
	Encounter *Reference `bson:"encounter,omitempty" json:"encounter,omitempty" xml:"encounter,omitempty"`
	// The time or time-period the observed values are related to.
	//
	// When the subject of the report is a patient, this is usually either the time of the procedure or of specimen collection(s), but very often the source of the date/time is not known, only the date/time itself.
	//
	// Comment: If the diagnostic procedure was performed on the patient, this is the time it was performed. If there are specimens, the diagnostically relevant time can be derived from the specimen collection times, but the specimen information is not always available, and the exact relationship between the specimens and the diagnostically relevant time is not always automatic.
	EffectiveDateTime *string `bson:"effectiveDateTime,omitempty" json:"effectiveDateTime,omitempty" xml:"effectiveDateTime,omitempty"`
	EffectivePeriod   *Period `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty" xml:"effectivePeriod,omitempty"`
	// The date and time that this version of the report was made available to providers, typically after the report was reviewed and verified.
	//
	// Comment: May be different from the update time of the resource itself, because that is the status of the record (potentially a secondary copy), not the actual release time of the report.
	Issued *string `bson:"issued,omitempty" json:"issued,omitempty" xml:"issued,omitempty"`
	// The diagnostic service that is responsible for issuing the report.
	//
	// Comment: This is not necessarily the source of the atomic data items or the entity that interpreted the results. It is the entity that takes responsibility for the clinical report.
	Performer []Reference `bson:"performer,omitempty" json:"performer,omitempty" xml:"performer,omitempty"`
	// The practitioner or organization that is responsible for the report's conclusions and interpretations.
	//
	// Comment: Might not be the same entity that takes responsibility for the clinical report.
	ResultsInterpreter []Reference `bson:"resultsInterpreter,omitempty" json:"resultsInterpreter,omitempty" xml:"resultsInterpreter,omitempty"`
	// Details about the specimens on which this diagnostic report is based.
	//
	// Comment: If the specimen is sufficiently specified with a code in the test result name, then this additional data may be redundant. If there are multiple specimens, these may be represented per observation or group.
	Specimen []Reference `bson:"specimen,omitempty" json:"specimen,omitempty" xml:"specimen,omitempty"`
	// [Observations](observation.html)  that are part of this diagnostic report.
	//
	// Comment: Observations can contain observations.
	Result []Reference `bson:"result,omitempty" json:"result,omitempty" xml:"result,omitempty"`
	// Comments about the diagnostic report.
	//
	// Comment: May include general statements about the diagnostic report, or statements about significant, unexpected or unreliable results values contained within the diagnostic report, or information about its source when relevant to its interpretation.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// One or more links to full details of any study performed during the diagnostic investigation.
	//
	// An ImagingStudy might comprise a set of radiologic images obtained via a procedure that are analyzed as a group. Typically, this is imaging performed by DICOM enabled modalities, but this is not required. A fully enabled PACS viewer can use this information to provide views of the source images. A GenomicStudy might comprise one or more analyses, each serving a specific purpose. These analyses may vary in method (e.g., karyotyping, CNV, or SNV detection), performer, software, devices used, or regions targeted.
	//
	// Comment: For laboratory-type studies like GenomeStudy, type resources will be used for tracking additional metadata and workflow aspects of complex studies. ImagingStudy and the media element are somewhat overlapping - typically, the list of image references in the media element will also be found in one of the imaging study resources. However, each caters to different types of displays for different types of purposes. Neither, either, or both may be provided.
	Study []Reference `bson:"study,omitempty" json:"study,omitempty" xml:"study,omitempty"`
	// This backbone element contains supporting information that was used in the creation of the report not included in the results already included in the report.
	SupportingInfo []DiagnosticReportSupportingInfo `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty" xml:"supportingInfo,omitempty"`
	// A list of key images or data associated with this report.
	//
	// The images or data are generally created during the diagnostic process, and may be directly of the patient, or of treated specimens (i.e. slides of interest).
	Media []DiagnosticReportMedia `bson:"media,omitempty" json:"media,omitempty" xml:"media,omitempty"`
	// Reference to a Composition resource instance that provides structure for organizing the contents of the DiagnosticReport.
	//
	// Comment: The Composition provides structure to the content of the DiagnosticReport (and only contains contents referenced in the DiagnosticReport) - e.g., to order the sections of an anatomic pathology structured report.
	Composition *Reference `bson:"composition,omitempty" json:"composition,omitempty" xml:"composition,omitempty"`
	// Concise and clinically contextualized summary conclusion (interpretation/impression) of the diagnostic report.
	Conclusion *string `bson:"conclusion,omitempty" json:"conclusion,omitempty" xml:"conclusion,omitempty"`
	// One or more codes that represent the summary conclusion (interpretation/impression) of the diagnostic report.
	ConclusionCode []CodeableConcept `bson:"conclusionCode,omitempty" json:"conclusionCode,omitempty" xml:"conclusionCode,omitempty"`
	// Rich text representation of the entire result as issued by the diagnostic service.
	//
	// Multiple formats are allowed but they SHALL be semantically equivalent.
	//
	// Comment: "application/pdf" is recommended as the most reliable and interoperable in this context.
	PresentedForm []Attachment `bson:"presentedForm,omitempty" json:"presentedForm,omitempty" xml:"presentedForm,omitempty"`
}

type DiagnosticReportSupportingInfo struct {
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
	// The code value for the role of the supporting information in the diagnostic report.
	Type CodeableConcept `bson:"type" json:"type" xml:"type"`
	// The reference for the supporting information in the diagnostic report.
	Reference Reference `bson:"reference" json:"reference" xml:"reference"`
}

type DiagnosticReportMedia struct {
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
	// A comment about the image or data.
	//
	// Typically, this is used to provide an explanation for why the image or data is included, or to draw the viewer's attention to important features.
	//
	// Comment: The comment should be displayed with the image or data. It would be common for the report to include additional discussion of the image or data contents or in other sections such as the conclusion.
	Comment *string `bson:"comment,omitempty" json:"comment,omitempty" xml:"comment,omitempty"`
	// Reference to the image or data source.
	Link Reference `bson:"link" json:"link" xml:"link"`
}

func (r DiagnosticReport) ResourceType() string {
	return "DiagnosticReport"
}

// MarshalJSON implements json.Marshaler
func (r DiagnosticReport) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		DiagnosticReport
		ResourceType string `json:"resourceType"`
	}{
		DiagnosticReport: r,
		ResourceType:     r.ResourceType(),
	})
}
