package fhires

import v2 "encoding/json/v2"

// The Citation Resource enables reference to any knowledge artifact for purposes of identification and attribution.
//
// The Citation Resource supports existing reference structures and developing publication practices such as versioning, expressing complex contributorship roles, and referencing computable resources.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Citation
type Citation struct {
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
	// An absolute URI that is used to identify this citation record when it is referenced in a specification, model, design or an instance; also called its canonical identifier.
	//
	// This SHOULD be globally unique and SHOULD be a literal address at which an authoritative instance of this summary is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the summary is stored on different servers.
	//
	// Comment: In some cases, the resource can no longer be found at the stated url, but the url itself cannot change. Implementations can use the meta.source element to indicate where the current master source of the resource can be found.
	URL *string `bson:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
	// A formal identifier that is used to identify this citation record when it is represented in other formats, or referenced in a specification, model, design or an instance.
	//
	// Comment: Use this element if you need to identify the citation record independently from identifying the cited artifact.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The identifier that is used to identify this version of the citation record when it is referenced in a specification, model, design or instance.
	//
	// This is an arbitrary value managed by the citation record author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	//
	// Comment: There may be different citation record instances that have the same identifier but different versions. The version can be appended to the url in a reference to allow a reference to a particular business version of the citation record with the format [url]|[version]. The version SHOULD NOT contain a '#' - see [Business Version](resource.html#bv-format).
	Version *string `bson:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
	// Indicates the mechanism used to compare versions to determine which is more current.
	//
	// Comment: If set as a string, this is a FHIRPath expression that has two additional context variables passed in - %version1 and %version2 and will return a negative number if version1 is newer, a positive number if version2 is newer, and a 0 if the version ordering can't successfully be determined.
	VersionAlgorithmString *string `bson:"versionAlgorithmString,omitempty" json:"versionAlgorithmString,omitempty" xml:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding `bson:"versionAlgorithmCoding,omitempty" json:"versionAlgorithmCoding,omitempty" xml:"versionAlgorithmCoding,omitempty"`
	// A natural language name identifying the citation record.
	//
	// This name should be usable as an identifier for the module by machine processing applications such as code generation.
	//
	// Comment: The name is not expected to be globally unique. The name should be a simple alphanumeric type name to ensure that it is machine-processing friendly.
	Name *string `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// A short, descriptive, user-friendly title for the citation record.
	//
	// Comment: This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc. Use this element if you have a need for naming the citation record for human use, such as a display name for the citation record that is used when selecting citations from a list.
	Title *string `bson:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// The status of this summary. Enables tracking the life-cycle of the content.
	//
	// Comment: Allows filtering of summaries that are appropriate for use versus not. Use "active" when the citation record is the version for others to actively use, "draft" while you are developing the content, and "retired" when this record should no longer be used for any purpose.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Status PublicationStatus `bson:"status" json:"status" xml:"status"`
	// A Boolean value to indicate that this citation record is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	//
	// Comment: Allows filtering of citation records that are appropriate for use versus not.
	Experimental *bool `bson:"experimental,omitempty" json:"experimental,omitempty" xml:"experimental,omitempty"`
	// The date (and optionally time) when the citation record was last significantly changed.
	//
	// The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the citation record changes.
	//
	// Comment: The date is often not tracked until the resource is published, but may be present on draft content. Note that this is not the same as the resource last-modified-date, since the resource may be a secondary representation of the citation record. Additional specific dates may be added as extensions or be found by consulting Provenances associated with past versions of the resource.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Date *string `bson:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// The name of the organization or individual that published the citation record.
	//
	// Comment: Usually an organization but may be an individual. The publisher (or steward) of the citation record is the organization or individual primarily responsible for the maintenance and upkeep of the citation record. This is not necessarily the same individual or organization that developed and initially authored the content. The publisher is the primary point of contact for questions or issues with the citation record. This item SHOULD be populated unless the information is available from context.
	Publisher *string `bson:"publisher,omitempty" json:"publisher,omitempty" xml:"publisher,omitempty"`
	// Contact details to assist a user in finding and communicating with the publisher.
	//
	// Comment: May be a web site, an email address, a telephone number, etc.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty" xml:"contact,omitempty"`
	// A free text natural language description of the citation from a consumer's perspective.
	//
	// Comment: This description can be used to capture details such as comments about misuse, instructions for clinical use and interpretation, literature references, examples from the paper world, etc. It is not a rendering of the citation as conveyed in the 'text' field of the resource itself. This item SHOULD be populated unless the information is available from context (e.g. the language of the citation is presumed to be the predominant language in the place the citation was created).
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// The content was developed with a focus and intent of supporting the contexts that are listed.
	//
	// These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate citation record instances.
	//
	// Comment: When multiple useContexts are specified, there is no expectation that all or any of the contexts apply.
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty" xml:"useContext,omitempty"`
	// A legal or geographic region in which the citation record is intended to be used.
	//
	// Comment: It may be possible for the citation record to be used in jurisdictions other than those for which it was originally designed or intended.
	//
	// DEPRECATION NOTE: For consistency, implementations are encouraged to migrate to using the new 'jurisdiction' code in the useContext element.  (I.e. useContext.code indicating http://terminology.hl7.org/CodeSystem/usage-context-type#jurisdiction and useContext.valueCodeableConcept indicating the jurisdiction.)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty" xml:"jurisdiction,omitempty"`
	// Explanation of why this citation is needed and why it has been designed as it has.
	//
	// Comment: This element does not describe the usage of the citation. Instead, it provides traceability of ''why'' the resource is either needed or ''why'' it is defined as it is.  This may be used to point to source materials or specifications that drove the structure of this citation.
	Purpose *string `bson:"purpose,omitempty" json:"purpose,omitempty" xml:"purpose,omitempty"`
	// Use and/or publishing restrictions for the citation record, not for the cited artifact.
	Copyright *string `bson:"copyright,omitempty" json:"copyright,omitempty" xml:"copyright,omitempty"`
	// A short string (<50 characters), suitable for inclusion in a page footer that identifies the copyright holder, effective period, and optionally whether rights are resctricted.
	//
	// (e.g. 'All rights reserved', 'Some rights reserved').
	//
	// Comment: The (c) symbol should NOT be included in this string. It will be added by software when rendering the notation. Full details about licensing, restrictions, warrantees, etc. goes in the more general 'copyright' element.
	CopyrightLabel *string `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty" xml:"copyrightLabel,omitempty"`
	// The date on which the resource content was approved by the publisher.
	//
	// Approval happens once when the content is officially approved for usage.
	//
	// Comment: The 'date' element may be more recent than the approval date because of minor changes or editorial corrections.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	ApprovalDate *string `bson:"approvalDate,omitempty" json:"approvalDate,omitempty" xml:"approvalDate,omitempty"`
	// The date on which the resource content was last reviewed.
	//
	// Review happens periodically after approval but does not change the original approval date.
	//
	// Comment: If specified, this date follows the original approval date.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	LastReviewDate *string `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty" xml:"lastReviewDate,omitempty"`
	// The period during which the citation record content was or is planned to be in active use.
	//
	// Comment: The effective period for a citation record determines when the content is applicable for usage and is independent of publication and review dates. For example, a citation intended to be used for the year 2016 might be published in 2015.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	EffectivePeriod *Period `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty" xml:"effectivePeriod,omitempty"`
	// Who authored or created the citation record.
	//
	// Comment: The Citation.author element is structured to support a name (with a string datatype) and contact point information for the author/creator, and is not structured to support unique identifiers for the author. If detailed contributorship data is needed for the authorship of the citation record, then one can create a Citation Resource for the Citation Resource.
	Author []ContactDetail `bson:"author,omitempty" json:"author,omitempty" xml:"author,omitempty"`
	// Who edited or revised the citation record.
	//
	// Comment: The Citation.editor element is structured to support a name (with a string datatype) and contact point information for the editor, and is not structured to support unique identifiers for the editor. If detailed contributorship data is needed for the editing of the citation record, then one can create a Citation Resource for the Citation Resource.
	Editor []ContactDetail `bson:"editor,omitempty" json:"editor,omitempty" xml:"editor,omitempty"`
	// Who reviewed the citation record.
	//
	// Comment: The Citation.reviewer element is structured to support a name (with a string datatype) and contact point information for the reviewer, and is not structured to support unique identifiers for the reviewer. If detailed contributorship data is needed for the review of the citation record, then one can create a Citation Resource for the Citation Resource.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Reviewer []ContactDetail `bson:"reviewer,omitempty" json:"reviewer,omitempty" xml:"reviewer,omitempty"`
	// Who endorsed the citation record.
	//
	// Comment: The Citation.endorser element is structured to support a name (with a string datatype) and contact point information for the endorser, and is not structured to support unique identifiers for the endorser. If detailed contributorship data is needed for the endorsement of the citation record, then one can create a Citation Resource for the Citation Resource.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Endorser []ContactDetail `bson:"endorser,omitempty" json:"endorser,omitempty" xml:"endorser,omitempty"`
	// A human-readable display of key concepts to represent the citation.
	Summary []CitationSummary `bson:"summary,omitempty" json:"summary,omitempty" xml:"summary,omitempty"`
	// The assignment to an organizing scheme.
	//
	// Comment: Use this element if you need to classify the citation record independently from classifying the cited artifact.
	Classification []CitationClassification `bson:"classification,omitempty" json:"classification,omitempty" xml:"classification,omitempty"`
	// Used for general notes and annotations not coded elsewhere.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// The status of the citation record.
	//
	// Comment: Use this if needed for reporting the state or status of the citation record, NOT FOR reporting the state or status of the cited article.
	CurrentState []CodeableConcept `bson:"currentState,omitempty" json:"currentState,omitempty" xml:"currentState,omitempty"`
	// The state or status of the citation record paired with an effective date or period for that state.
	//
	// Comment: Use this if needed for reporting the state or status of the citation record, NOT FOR reporting the state or status of the cited article.
	StatusDate []CitationStatusDate `bson:"statusDate,omitempty" json:"statusDate,omitempty" xml:"statusDate,omitempty"`
	// Artifact related to the citation record.
	//
	// Comment: Use this if needed for reporting artifacts related to the citation record, NOT FOR reporting artifacts related to the cited article. For example, one may use this element to report the data source for generation of the Citation Resource instance if it was automatically generated, such as conversion from a citation repository.
	RelatedArtifact []RelatedArtifact `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty" xml:"relatedArtifact,omitempty"`
	// The article or artifact being described.
	CitedArtifact *CitationCitedArtifact `bson:"citedArtifact,omitempty" json:"citedArtifact,omitempty" xml:"citedArtifact,omitempty"`
}

type CitationSummary struct {
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
	// Format for display of the citation summary.
	Style *CodeableConcept `bson:"style,omitempty" json:"style,omitempty" xml:"style,omitempty"`
	// The human-readable display of the citation summary.
	Text string `bson:"text" json:"text" xml:"text"`
}

type CitationClassification struct {
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
	// The kind of classifier (e.g. publication type, keyword).
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The specific classification value.
	Classifier []CodeableConcept `bson:"classifier,omitempty" json:"classifier,omitempty" xml:"classifier,omitempty"`
}

type CitationStatusDate struct {
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
	// The state or status of the citation record (that will be paired with the period).
	Activity CodeableConcept `bson:"activity" json:"activity" xml:"activity"`
	// Whether the status date is actual (has occurred) or expected (estimated or anticipated).
	Actual *bool `bson:"actual,omitempty" json:"actual,omitempty" xml:"actual,omitempty"`
	// When the status started and/or ended.
	//
	// Comment: For an instance, place the same value in both start and end elements.
	Period Period `bson:"period" json:"period" xml:"period"`
}

type CitationCitedArtifact struct {
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
	// A formal identifier that is used to identify the cited artifact when it is represented in other formats, or referenced in a specification, model, design or an instance.
	//
	// Comment: This element will contain unique identifiers that support de-duplication of citations. Commonly used identifiers for study citations include digital object identifier (DOI), PubMed ID (PMID), and PubMed Central ID (PMCID). To maintain interoperability, the suggested identifier.system values to use for these 3 identifiers are: DOI = "https://doi.org" PMID = "https://pubmed.ncbi.nlm.nih.gov" PMCID = "https://www.ncbi.nlm.nih.gov/pmc/"
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// A formal identifier that is used to identify things closely related to the cited artifact.
	//
	// Comment: May include trial registry identifiers, e.g. NCT04372602 from clinicaltrials.gov. This related identifier would not be used as unique identifier for the cited artifact itself. Multiple published articles (each with their own unique identifier) may relate to the same study record in a research registry. Commonly used related identifiers for study registries include ClinicalTrials.gov identifier (NCTId) and PROSPERO identifier (used for registration of systematic reviews). To maintain interoperability, the suggested identifier.system values to use for these 2 identifiers are: ClinicalTrials.gov = "https://clinicaltrials.gov" PROSPERO = "https://www.crd.york.ac.uk/prospero/"
	RelatedIdentifier []Identifier `bson:"relatedIdentifier,omitempty" json:"relatedIdentifier,omitempty" xml:"relatedIdentifier,omitempty"`
	// When the cited artifact was accessed.
	//
	// Comment: Use this element when citing an artifact from a potentially unstable source, such as a webpage, to note the date (and time) the source was accessed.
	DateAccessed *string `bson:"dateAccessed,omitempty" json:"dateAccessed,omitempty" xml:"dateAccessed,omitempty"`
	// The defined version of the cited artifact.
	Version *CitationCitedArtifactVersion `bson:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
	// The status of the cited artifact.
	CurrentState []CodeableConcept `bson:"currentState,omitempty" json:"currentState,omitempty" xml:"currentState,omitempty"`
	// An effective date or period, historical or future, actual or expected, for a status of the cited artifact.
	StatusDate []CitationCitedArtifactStatusDate `bson:"statusDate,omitempty" json:"statusDate,omitempty" xml:"statusDate,omitempty"`
	// The title details of the article or artifact.
	Title []CitationCitedArtifactTitle `bson:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// The abstract may be used to convey article-contained abstracts, externally-created abstracts, or other descriptive summaries.
	Abstract []CitationCitedArtifactAbstract `bson:"abstract,omitempty" json:"abstract,omitempty" xml:"abstract,omitempty"`
	// The component of the article or artifact.
	Part *CitationCitedArtifactPart `bson:"part,omitempty" json:"part,omitempty" xml:"part,omitempty"`
	// The artifact related to the cited artifact.
	//
	// Comment: The citedArtifact.relatesTo element uses a BackboneElement instead of the RelatedArtifact Datatype to enable use of an extended value set for the required code for the type of relationship.
	RelatesTo []CitationCitedArtifactRelatesTo `bson:"relatesTo,omitempty" json:"relatesTo,omitempty" xml:"relatesTo,omitempty"`
	// If multiple, used to represent alternative forms of the article that are not separate citations.
	//
	// Comment: A common use is a journal article with a publication date and pagination for a print version and a different publication date for the online version of the same article.
	PublicationForm []CitationCitedArtifactPublicationForm `bson:"publicationForm,omitempty" json:"publicationForm,omitempty" xml:"publicationForm,omitempty"`
	// Used for any URL for the article or artifact cited.
	WebLocation []CitationCitedArtifactWebLocation `bson:"webLocation,omitempty" json:"webLocation,omitempty" xml:"webLocation,omitempty"`
	// The assignment to an organizing scheme.
	Classification []CitationCitedArtifactClassification `bson:"classification,omitempty" json:"classification,omitempty" xml:"classification,omitempty"`
	// This element is used to list authors and other contributors, their contact information, specific contributions, and summary statements.
	Contributorship *CitationCitedArtifactContributorship `bson:"contributorship,omitempty" json:"contributorship,omitempty" xml:"contributorship,omitempty"`
	// Any additional information or content for the article or artifact.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
}

type CitationCitedArtifactVersion struct {
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
	// The version number or other version identifier.
	Value string `bson:"value" json:"value" xml:"value"`
	// Citation for the main version of the cited artifact.
	//
	// Comment: When referencing a baseCitation, one may inherit any data from the referenced Citation Resource and consider any data in the current Citation Resource as superceding to be version-specific.
	BaseCitation *Reference `bson:"baseCitation,omitempty" json:"baseCitation,omitempty" xml:"baseCitation,omitempty"`
}

type CitationCitedArtifactStatusDate struct {
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
	// A definition of the status associated with a date or period.
	Activity CodeableConcept `bson:"activity" json:"activity" xml:"activity"`
	// Either occurred or expected.
	Actual *bool `bson:"actual,omitempty" json:"actual,omitempty" xml:"actual,omitempty"`
	// When the status started and/or ended.
	//
	// Comment: For an instance, place the same value in both start and end elements.
	Period Period `bson:"period" json:"period" xml:"period"`
}

type CitationCitedArtifactTitle struct {
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
	// Used to express the reason for or classification of the title.
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Used to express the specific language of the title.
	Language *CodeableConcept `bson:"language,omitempty" json:"language,omitempty" xml:"language,omitempty"`
	// The title of the article or artifact.
	Text string `bson:"text" json:"text" xml:"text"`
}

type CitationCitedArtifactAbstract struct {
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
	// Used to express the reason for or classification of the abstract.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Used to express the specific language of the abstract.
	Language *CodeableConcept `bson:"language,omitempty" json:"language,omitempty" xml:"language,omitempty"`
	// Abstract content.
	Text string `bson:"text" json:"text" xml:"text"`
	// Copyright notice for the abstract.
	//
	// Comment: There is also a copyright element in the publicationForm element, so this copyright element is intended for copyright of the abstract.
	Copyright *string `bson:"copyright,omitempty" json:"copyright,omitempty" xml:"copyright,omitempty"`
}

type CitationCitedArtifactPart struct {
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
	// The kind of component.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The specification of the component.
	Value *string `bson:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
	// The citation for the full article or artifact.
	//
	// Comment: When referencing a baseCitation, one may inherit any data from the referenced Citation Resource and consider any data in the current Citation Resource as superceding to be specific to the citation of the specified part. For example, the part-specific Citation may have additional contributorship data that supercedes the contributorship data for the full citation when there is a need to uniquely report the contributorship for an image or supplement associated with the fuller citation record.
	BaseCitation *Reference `bson:"baseCitation,omitempty" json:"baseCitation,omitempty" xml:"baseCitation,omitempty"`
}

type CitationCitedArtifactRelatesTo struct {
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
	// The type of relationship to the related artifact.
	Type string `bson:"type" json:"type" xml:"type"`
	// Provides additional classifiers of the related artifact.
	Classifier []CodeableConcept `bson:"classifier,omitempty" json:"classifier,omitempty" xml:"classifier,omitempty"`
	// A short label that can be used to reference the related artifact from elsewhere in the containing artifact, such as a footnote index.
	//
	// Comment: Though not commonly used in citation records, this could be used for referencing other articles in the abstract or explanatory notes.
	Label *string `bson:"label,omitempty" json:"label,omitempty" xml:"label,omitempty"`
	// A brief description of the document or knowledge resource being referenced, suitable for display to a consumer.
	Display *string `bson:"display,omitempty" json:"display,omitempty" xml:"display,omitempty"`
	// A bibliographic citation for the related artifact.
	//
	// This text SHOULD be formatted according to an accepted citation format.
	Citation *string `bson:"citation,omitempty" json:"citation,omitempty" xml:"citation,omitempty"`
	// The document being referenced, represented as an attachment.
	//
	// Do not use this element if using the resource element to provide the canonical to the related artifact.
	//
	// Comment: To refer to related artifacts by URL, the FHIR Attachment Datatype includes a url element, so the path would be Citation.citedArtifact.relatesTo.document.url
	Document *Attachment `bson:"document,omitempty" json:"document,omitempty" xml:"document,omitempty"`
	// The related artifact, such as a library, value set, profile, or other knowledge resource.
	//
	// Comment: If the type is predecessor, this is a reference to the succeeding knowledge resource. If the type is successor, this is a reference to the prior knowledge resource
	Resource *string `bson:"resource,omitempty" json:"resource,omitempty" xml:"resource,omitempty"`
	// The related artifact, if the artifact is not a canonical resource, or a resource reference to a canonical resource.
	//
	// Comment: If both resource and resourceReference are present, they SHOULD be consistent and reference the same resource. Although relatedArtifact is intended to support references to definitional resources, there are cases where non-definitional resources can be definitional (such as Location where the kind is mode). Related artifacts SHOULD be used to reference definitional resources, and profiles SHOULD be used to make that explicit for particular use cases.
	ResourceReference *Reference `bson:"resourceReference,omitempty" json:"resourceReference,omitempty" xml:"resourceReference,omitempty"`
}

type CitationCitedArtifactPublicationForm struct {
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
	// The collection the cited article or artifact is published in.
	PublishedIn *CitationCitedArtifactPublicationFormPublishedIn `bson:"publishedIn,omitempty" json:"publishedIn,omitempty" xml:"publishedIn,omitempty"`
	// Describes the form of the medium cited.
	//
	// Common codes are "Internet" or "Print". The CitedMedium value set has 6 codes. The codes internet, print, and offline-digital-storage are the common codes for a typical publication form, though internet and print are more common for study citations. Three additional codes (each appending one of the primary codes with "-without-issue" are used for situations when a study is published both within an issue (of a periodical release as commonly done for journals) AND is published separately from the issue (as commonly done with early online publication), to represent specific identification of the publication form not associated with the issue.
	CitedMedium *CodeableConcept `bson:"citedMedium,omitempty" json:"citedMedium,omitempty" xml:"citedMedium,omitempty"`
	// Volume number of journal or other collection in which the article is published.
	Volume *string `bson:"volume,omitempty" json:"volume,omitempty" xml:"volume,omitempty"`
	// Issue, part or supplement of journal or other collection in which the article is published.
	Issue *string `bson:"issue,omitempty" json:"issue,omitempty" xml:"issue,omitempty"`
	// The date the article was added to the database, or the date the article was released.
	//
	// Comment: The articleDate is the preferred element for expressing the publication date as structured data.
	ArticleDate *string `bson:"articleDate,omitempty" json:"articleDate,omitempty" xml:"articleDate,omitempty"`
	// Text representation of the date on which the issue of the cited artifact was published.
	//
	// Comment: The publicationDateText element is prefererntially used when a date is not represented in a form that can be handled as structured data in other elements.
	PublicationDateText *string `bson:"publicationDateText,omitempty" json:"publicationDateText,omitempty" xml:"publicationDateText,omitempty"`
	// Spring, Summer, Fall/Autumn, Winter.
	PublicationDateSeason *string `bson:"publicationDateSeason,omitempty" json:"publicationDateSeason,omitempty" xml:"publicationDateSeason,omitempty"`
	// The date the article was last revised or updated in the database.
	//
	// Comment: The lastRevisionDate is used for the cited article and not the date the Citation Resource is last revised.
	LastRevisionDate *string `bson:"lastRevisionDate,omitempty" json:"lastRevisionDate,omitempty" xml:"lastRevisionDate,omitempty"`
	// The language or languages in which this form of the article is published.
	Language []CodeableConcept `bson:"language,omitempty" json:"language,omitempty" xml:"language,omitempty"`
	// Entry number or identifier for inclusion in a database.
	//
	// Comment: Accession numbers may be used instead of or in addition to page numbers.  Accession numbers are functionally identifiers when coupled with the identity of the database for which they are accession numbers. Accession numbers which are used for indexing citations in a dataset of citations should NOT be entered in Citation.citedArtifact.publicationForm.accessionNumber -- this use would be a type of Citation.identifier for the citation record itself.
	AccessionNumber *string `bson:"accessionNumber,omitempty" json:"accessionNumber,omitempty" xml:"accessionNumber,omitempty"`
	// Used for full display of pagination.
	PageString *string `bson:"pageString,omitempty" json:"pageString,omitempty" xml:"pageString,omitempty"`
	// Used for isolated representation of first page.
	FirstPage *string `bson:"firstPage,omitempty" json:"firstPage,omitempty" xml:"firstPage,omitempty"`
	// Used for isolated representation of last page.
	LastPage *string `bson:"lastPage,omitempty" json:"lastPage,omitempty" xml:"lastPage,omitempty"`
	// Actual or approximate number of pages or screens.
	//
	// Distinct from reporting the page numbers.
	PageCount *string `bson:"pageCount,omitempty" json:"pageCount,omitempty" xml:"pageCount,omitempty"`
	// Copyright notice for the full article or artifact.
	//
	// Comment: It is possible that different publication forms have different copyright notices.
	Copyright *string `bson:"copyright,omitempty" json:"copyright,omitempty" xml:"copyright,omitempty"`
}

type CitationCitedArtifactPublicationFormPublishedIn struct {
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
	// Kind of container (e.g. Periodical, database, or book).
	//
	// Comment: Journals and newspapers are coded as Periodical.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Journal identifiers include ISSN, ISO Abbreviation and NLMuniqueID; Book identifiers include ISBN.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Name of the database or title of the book or journal.
	//
	// Comment: ClinicalTrials.gov is a name of a database.
	Title *string `bson:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Name of or resource describing the publisher.
	Publisher *Reference `bson:"publisher,omitempty" json:"publisher,omitempty" xml:"publisher,omitempty"`
	// Geographic location of the publisher.
	PublisherLocation *string `bson:"publisherLocation,omitempty" json:"publisherLocation,omitempty" xml:"publisherLocation,omitempty"`
}

type CitationCitedArtifactWebLocation struct {
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
	// A characterization of the object expected at the web location.
	//
	// Comment: Categories that may be frequently used for study citations may include abstract, full-text, supplement, webpage, and doi-based.
	Classifier []CodeableConcept `bson:"classifier,omitempty" json:"classifier,omitempty" xml:"classifier,omitempty"`
	// The specific URL.
	//
	// Comment: Persistent URLs, like doi, are preferred.
	URL *string `bson:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
}

type CitationCitedArtifactClassification struct {
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
	// The kind of classifier (e.g. publication type, keyword).
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The specific classification value.
	Classifier []CodeableConcept `bson:"classifier,omitempty" json:"classifier,omitempty" xml:"classifier,omitempty"`
	// Complex or externally created classification.
	//
	// Comment: Examples relevant to study citations include MeSH Headings or Keywords that contain qualifier codings or sub-classifications, and include risk of bias assessments created by persons different from the creator of the citation record.
	ArtifactAssessment []Reference `bson:"artifactAssessment,omitempty" json:"artifactAssessment,omitempty" xml:"artifactAssessment,omitempty"`
}

type CitationCitedArtifactContributorship struct {
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
	// Indicates if the list includes all authors and/or contributors.
	//
	// Comment: If list is incomplete, "et al" can be appended to Contributorship.summary.value for display purposes.
	Complete *bool `bson:"complete,omitempty" json:"complete,omitempty" xml:"complete,omitempty"`
	// An individual entity named as a contributor, for example in the author list or contributor list.
	//
	// Comment: Used to report contributorship in individualized ways.
	Entry []CitationCitedArtifactContributorshipEntry `bson:"entry,omitempty" json:"entry,omitempty" xml:"entry,omitempty"`
	// Used to record a display of the author/contributor list without separate data element for each list member.
	Summary []CitationCitedArtifactContributorshipSummary `bson:"summary,omitempty" json:"summary,omitempty" xml:"summary,omitempty"`
}

type CitationCitedArtifactContributorshipEntry struct {
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
	// The identity of the individual contributor.
	//
	// Comment: If naming a contributor without additional detail or further structured data, the display element within the Reference datatype may be used for a simple string without referencing another resource.
	Contributor Reference `bson:"contributor" json:"contributor" xml:"contributor"`
	// For citation styles that use initials.
	//
	// Comment: This element is used to support a data element used for citations that is not otherwise supported fully by FHIR HumanName datatype. In some citation summary styles (e.g. Medline), authors are presented with full spelling of their last name (family name) and initials for all forenames (first names, middle names, preceding names, given names). As the FHIR HumanName Datatype does not have a specific element for forename initials and this is common in citation usage, the element is added to contributorship.entry directly.
	ForenameInitials *string `bson:"forenameInitials,omitempty" json:"forenameInitials,omitempty" xml:"forenameInitials,omitempty"`
	// Organization affiliated with the contributor.
	//
	// Comment: If naming an organization without additional detail or further structured data, the display element within the Reference datatype may be used for a simple string without referencing another resource.
	Affiliation []Reference `bson:"affiliation,omitempty" json:"affiliation,omitempty" xml:"affiliation,omitempty"`
	// This element identifies the specific nature of an individual’s contribution with respect to the cited work.
	//
	// Comment: The CreDiT taxonomy provides a way to encode contribution information. The purpose of CRediT is to provide transparency in contributions to scholarly published work, to enable improved systems of attribution, credit, and accountability.
	ContributionType []CodeableConcept `bson:"contributionType,omitempty" json:"contributionType,omitempty" xml:"contributionType,omitempty"`
	// The role of the contributor (e.g. author, editor, reviewer, funder).
	Role *CodeableConcept `bson:"role,omitempty" json:"role,omitempty" xml:"role,omitempty"`
	// Contributions with accounting for time or number.
	ContributionInstance []CitationCitedArtifactContributorshipEntryContributionInstance `bson:"contributionInstance,omitempty" json:"contributionInstance,omitempty" xml:"contributionInstance,omitempty"`
	// Whether the contributor is the corresponding contributor for the role.
	//
	// Comment: If there are multiple authors and one serves as the corresponding author, then that one would have correspondingContact = true.
	CorrespondingContact *bool `bson:"correspondingContact,omitempty" json:"correspondingContact,omitempty" xml:"correspondingContact,omitempty"`
	// Provides a numerical ranking to represent the degree of contributorship relative to other contributors, such as 1 for first author and 2 for second author.
	//
	// Comment: Used to retain order of authorship as structured data
	RankingOrder *int `bson:"rankingOrder,omitempty" json:"rankingOrder,omitempty" xml:"rankingOrder,omitempty"`
}

type CitationCitedArtifactContributorshipEntryContributionInstance struct {
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
	// The specific contribution.
	Type CodeableConcept `bson:"type" json:"type" xml:"type"`
	// The time that the contribution was made.
	Time *string `bson:"time,omitempty" json:"time,omitempty" xml:"time,omitempty"`
}

type CitationCitedArtifactContributorshipSummary struct {
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
	// Used most commonly to express an author list or a contributorship statement.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The format for the display string, such as author last name with first letter capitalized followed by forename initials.
	Style *CodeableConcept `bson:"style,omitempty" json:"style,omitempty" xml:"style,omitempty"`
	// Used to code the producer or rule for creating the display string.
	Source *CodeableConcept `bson:"source,omitempty" json:"source,omitempty" xml:"source,omitempty"`
	// The display string for the author list, contributor list, or contributorship statement.
	Value string `bson:"value" json:"value" xml:"value"`
}

func (r Citation) ResourceType() string {
	return "Citation"
}

// MarshalJSON implements json.Marshaler
func (r Citation) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		Citation
		ResourceType string `json:"resourceType"`
	}{
		Citation:     r,
		ResourceType: r.ResourceType(),
	})
}
