package fhires

import (
	v2 "encoding/json/v2"

	decimal "github.com/shopspring/decimal"
)

// The Evidence Resource provides a machine-interpretable expression of an evidence concept including the evidence variables (e.g., population, exposures/interventions, comparators, outcomes, measured variables, confounding variables), the statistics, and the certainty of this evidence.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Evidence
type Evidence struct {
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
	// An absolute URI that is used to identify this evidence when it is referenced in a specification, model, design or an instance; also called its canonical identifier.
	//
	// This SHOULD be globally unique and SHOULD be a literal address at which an authoritative instance of this summary is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the summary is stored on different servers.
	//
	// Comment: In some cases, the resource can no longer be found at the stated url, but the url itself cannot change. Implementations can use the meta.source element to indicate where the current master source of the resource can be found.
	URL *string `bson:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
	// A formal identifier that is used to identify this summary when it is represented in other formats, or referenced in a specification, model, design or an instance.
	//
	// Comment: Typically, this is used for identifiers that can go in an HL7 V3 II (instance identifier) data type, and can then identify this summary outside of FHIR, where it is not possible to use the logical URI.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The identifier that is used to identify this version of the summary when it is referenced in a specification, model, design or instance.
	//
	// This is an arbitrary value managed by the summary author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	//
	// Comment: The determination of when to create a new version of a resource (same url, new version) vs. defining a new artifact is up to the author. Considerations for making this decision are found in Technical and Business Versions.
	Version *string `bson:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
	// Indicates the mechanism used to compare versions to determine which is more current.
	//
	// Comment: If set as a string, this is a FHIRPath expression that has two additional context variables passed in - %version1 and %version2 and will return a negative number if version1 is newer, a positive number if version2 is newer, and a 0 if the version ordering can't successfully be determined.
	VersionAlgorithmString *string `bson:"versionAlgorithmString,omitempty" json:"versionAlgorithmString,omitempty" xml:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding `bson:"versionAlgorithmCoding,omitempty" json:"versionAlgorithmCoding,omitempty" xml:"versionAlgorithmCoding,omitempty"`
	// A natural language name identifying the evidence.
	//
	// This name should be usable as an identifier for the module by machine processing applications such as code generation.
	//
	// Comment: The name is not expected to be globally unique. The name should be a simple alphanumeric type name to ensure that it is machine-processing friendly.
	Name *string `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// A short, descriptive, user-friendly title for the summary.
	Title *string `bson:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Citation Resource or display of suggested citation for this evidence.
	CiteAsReference *Reference `bson:"citeAsReference,omitempty" json:"citeAsReference,omitempty" xml:"citeAsReference,omitempty"`
	CiteAsMarkdown  *string    `bson:"citeAsMarkdown,omitempty" json:"citeAsMarkdown,omitempty" xml:"citeAsMarkdown,omitempty"`
	// The status of this summary. Enables tracking the life-cycle of the content.
	//
	// Comment: Allows filtering of summarys s that are appropriate for use versus not.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Status PublicationStatus `bson:"status" json:"status" xml:"status"`
	// A Boolean value to indicate that this resource is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	Experimental *bool `bson:"experimental,omitempty" json:"experimental,omitempty" xml:"experimental,omitempty"`
	// The date  (and optionally time) when the summary was last significantly changed.
	//
	// The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the summary changes.
	//
	// Comment: The date is often not tracked until the resource is published, but may be present on draft content. Note that this is not the same as the resource last-modified-date, since the resource may be a secondary representation of the summary. Additional specific dates may be added as extensions or be found by consulting Provenances associated with past versions of the resource.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Date *string `bson:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
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
	// The name of the organization or individual responsible for the release and ongoing maintenance of the evidence.
	//
	// Comment: Usually an organization but may be an individual. The publisher (or steward) of the evidence is the organization or individual primarily responsible for the maintenance and upkeep of the evidence. This is not necessarily the same individual or organization that developed and initially authored the content. The publisher is the primary point of contact for questions or issues with the evidence. This item SHOULD be populated unless the information is available from context.
	Publisher *string `bson:"publisher,omitempty" json:"publisher,omitempty" xml:"publisher,omitempty"`
	// Contact details to assist a user in finding and communicating with the publisher.
	//
	// Comment: Extensions to ContactDetail include: contactReference, contactAddress, and contributionTime (see [Clinical Reasoning Module](clinicalreasoning-module.html)).
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty" xml:"contact,omitempty"`
	// An individiual, organization, or device primarily involved in the creation and maintenance of the content.
	//
	// Comment: Extensions to ContactDetail include: contactReference, contactAddress, and contributionTime (see [Clinical Reasoning Module](clinicalreasoning-module.html)).
	Author []ContactDetail `bson:"author,omitempty" json:"author,omitempty" xml:"author,omitempty"`
	// An individiual, organization, or device primarily responsible for internal coherence of the content.
	//
	// Comment: Extensions to ContactDetail include: contactReference, contactAddress, and contributionTime (see [Clinical Reasoning Module](clinicalreasoning-module.html)).
	Editor []ContactDetail `bson:"editor,omitempty" json:"editor,omitempty" xml:"editor,omitempty"`
	// An individiual, organization, or device primarily responsible for review of some aspect of the content.
	//
	// Comment: Extensions to ContactDetail include: contactReference, contactAddress, and contributionTime (see [Clinical Reasoning Module](clinicalreasoning-module.html)).
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Reviewer []ContactDetail `bson:"reviewer,omitempty" json:"reviewer,omitempty" xml:"reviewer,omitempty"`
	// An individiual, organization, or device responsible for officially endorsing the content for use in some setting.
	//
	// Comment: Extensions to ContactDetail include: contactReference, contactAddress, and contributionTime (see [Clinical Reasoning Module](clinicalreasoning-module.html)).
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Endorser []ContactDetail `bson:"endorser,omitempty" json:"endorser,omitempty" xml:"endorser,omitempty"`
	// The content was developed with a focus and intent of supporting the contexts that are listed.
	//
	// These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate evidence instances.
	//
	// Comment: When multiple useContexts are specified, there is no expectation that all or any of the contexts apply.
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty" xml:"useContext,omitempty"`
	// Explanation of why this Evidence is needed and why it has been designed as it has.
	//
	// Comment: This element does not describe the usage of the Evidence. Instead, it provides traceability of ''why'' the resource is either needed or ''why'' it is defined as it is. This may be used to point to source materials or specifications that drove the structure of this Evidence.
	Purpose *string `bson:"purpose,omitempty" json:"purpose,omitempty" xml:"purpose,omitempty"`
	// A copyright statement relating to the Evidence and/or its contents.
	//
	// Copyright statements are generally legal restrictions on the use and publishing of the Evidence.
	//
	// Comment: The short copyright declaration (e.g. (c) '2015+ xyz organization' should be sent in the copyrightLabel element.
	Copyright *string `bson:"copyright,omitempty" json:"copyright,omitempty" xml:"copyright,omitempty"`
	// A short string (<50 characters), suitable for inclusion in a page footer that identifies the copyright holder, effective period, and optionally whether rights are resctricted.
	//
	// (e.g. 'All rights reserved', 'Some rights reserved').
	//
	// Comment: The (c) symbol should NOT be included in this string. It will be added by software when rendering the notation. Full details about licensing, restrictions, warrantees, etc. goes in the more general 'copyright' element.
	CopyrightLabel *string `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty" xml:"copyrightLabel,omitempty"`
	// Link or citation to artifact associated with the summary.
	RelatedArtifact []RelatedArtifact `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty" xml:"relatedArtifact,omitempty"`
	// A free text natural language description of the evidence from a consumer's perspective.
	//
	// Comment: This description can be used to capture details such as comments about misuse, instructions for clinical use and interpretation, literature references, examples from the paper world, etc. It is not a rendering of the evidence as conveyed in the 'text' field of the resource itself. This item SHOULD be populated unless the information is available from context (e.g. the language of the evidence is presumed to be the predominant language in the place the evidence was created).
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Declarative description of the Evidence.
	Assertion *string `bson:"assertion,omitempty" json:"assertion,omitempty" xml:"assertion,omitempty"`
	// Footnotes and/or explanatory notes.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// Evidence variable such as population, exposure, or outcome.
	VariableDefinition []EvidenceVariableDefinition `bson:"variableDefinition" json:"variableDefinition" xml:"variableDefinition"`
	// The method to combine studies.
	SynthesisType *CodeableConcept `bson:"synthesisType,omitempty" json:"synthesisType,omitempty" xml:"synthesisType,omitempty"`
	// The design of the study that produced this evidence.
	//
	// The design is described with any number of study design characteristics.
	StudyDesign []CodeableConcept `bson:"studyDesign,omitempty" json:"studyDesign,omitempty" xml:"studyDesign,omitempty"`
	// Values and parameters for a single statistic.
	Statistic []EvidenceStatistic `bson:"statistic,omitempty" json:"statistic,omitempty" xml:"statistic,omitempty"`
	// Assessment of certainty, confidence in the estimates, or quality of the evidence.
	Certainty []EvidenceCertainty `bson:"certainty,omitempty" json:"certainty,omitempty" xml:"certainty,omitempty"`
}

type EvidenceVariableDefinition struct {
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
	// A text description or summary of the variable.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Footnotes and/or explanatory notes.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// population | subpopulation | exposure | referenceExposure | measuredVariable | confounder.
	VariableRole CodeableConcept `bson:"variableRole" json:"variableRole" xml:"variableRole"`
	// Definition of the actual variable related to the statistic(s).
	Observed *Reference `bson:"observed,omitempty" json:"observed,omitempty" xml:"observed,omitempty"`
	// Definition of the intended variable related to the Evidence.
	Intended *Reference `bson:"intended,omitempty" json:"intended,omitempty" xml:"intended,omitempty"`
	// Indication of quality of match between intended variable to actual variable.
	DirectnessMatch *CodeableConcept `bson:"directnessMatch,omitempty" json:"directnessMatch,omitempty" xml:"directnessMatch,omitempty"`
}

type EvidenceStatistic struct {
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
	// A description of the content value of the statistic.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Footnotes and/or explanatory notes.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// Type of statistic, e.g., relative risk.
	StatisticType *CodeableConcept `bson:"statisticType,omitempty" json:"statisticType,omitempty" xml:"statisticType,omitempty"`
	// When the measured variable is handled categorically, the category element is used to define which category the statistic is reporting.
	//
	// Comment: Simple strings can be used for descriptive purposes. Exact matching to EvidenceVariable.category.name for the Evidence.variableDefinition[variableRole=measuredVariable].observed=Reference(EvidenceVariable) could facilitate validation within datasets.
	Category *CodeableConcept `bson:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
	// Statistic value.
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty" xml:"quantity,omitempty"`
	// The number of events associated with the statistic, where the unit of analysis is different from numberAffected, sampleSize.knownDataCount and sampleSize.numberOfParticipants.
	//
	// Comment: When the number of events is the statistic, use Evidence.statistic.quantity and set Evidence.statistic.type.coding.code=C25463 and Evidence.statistic.type.coding.display=Count. When the statistic is an Event Rate (where individual participants may have 2 or more events), use Evidence.statistic.numberOfEvents to record the total number of events rather than the number of participants with events.
	NumberOfEvents *int `bson:"numberOfEvents,omitempty" json:"numberOfEvents,omitempty" xml:"numberOfEvents,omitempty"`
	// The number of participants affected where the unit of analysis is the same as sampleSize.knownDataCount and sampleSize.numberOfParticipants.
	//
	// Comment: When the number affected is the statistic, use Evidence.statistic.quantity and set Evidence.statistic.type.coding.code=C25463 and Evidence.statistic.type.coding.display=Count. When the statistic is a Proportion, use Evidence.statistic.numberAffected and enter an integer as the value. When the statistic is an Event Rate (where individual participants may have 2 or more events), use Evidence.statistic.numberAffected to record the number of participants with events rather than the total number of events.
	NumberAffected *int `bson:"numberAffected,omitempty" json:"numberAffected,omitempty" xml:"numberAffected,omitempty"`
	// Number of samples in the statistic.
	SampleSize *EvidenceStatisticSampleSize `bson:"sampleSize,omitempty" json:"sampleSize,omitempty" xml:"sampleSize,omitempty"`
	// A statistical attribute of the statistic such as a measure of heterogeneity.
	AttributeEstimate []EvidenceStatisticAttributeEstimate `bson:"attributeEstimate,omitempty" json:"attributeEstimate,omitempty" xml:"attributeEstimate,omitempty"`
	// A component of the method to generate the statistic.
	ModelCharacteristic []EvidenceStatisticModelCharacteristic `bson:"modelCharacteristic,omitempty" json:"modelCharacteristic,omitempty" xml:"modelCharacteristic,omitempty"`
}

type EvidenceStatisticSampleSize struct {
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
	// Human-readable summary of population sample size.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Footnote or explanatory note about the sample size.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// Number of participants in the population.
	NumberOfStudies *int `bson:"numberOfStudies,omitempty" json:"numberOfStudies,omitempty" xml:"numberOfStudies,omitempty"`
	// A human-readable string to clarify or explain concepts about the sample size.
	NumberOfParticipants *int `bson:"numberOfParticipants,omitempty" json:"numberOfParticipants,omitempty" xml:"numberOfParticipants,omitempty"`
	// Number of participants with known results for measured variables.
	KnownDataCount *int `bson:"knownDataCount,omitempty" json:"knownDataCount,omitempty" xml:"knownDataCount,omitempty"`
}

type EvidenceStatisticAttributeEstimate struct {
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
	// Human-readable summary of the estimate.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Footnote or explanatory note about the estimate.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// The type of attribute estimate, e.g., confidence interval or p value.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The singular quantity of the attribute estimate, for attribute estimates represented as single values; also used to report unit of measure.
	//
	// Comment: Often the p value
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty" xml:"quantity,omitempty"`
	// Use 95 for a 95% confidence interval.
	Level *decimal.Decimal `bson:"level,omitempty" json:"level,omitempty" xml:"level,omitempty"`
	// Lower bound of confidence interval.
	Range *Range `bson:"range,omitempty" json:"range,omitempty" xml:"range,omitempty"`
	// A nested attribute estimate; which is the attribute estimate of an attribute estimate.
	//
	// Comment: A nested attribute estimate; which is the attribute estimate of an attribute estimate
	AttributeEstimate []EvidenceStatisticAttributeEstimate `bson:"attributeEstimate,omitempty" json:"attributeEstimate,omitempty"`
}

type EvidenceStatisticModelCharacteristic struct {
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
	// Description of a component of the method to generate the statistic.
	Code CodeableConcept `bson:"code" json:"code" xml:"code"`
	// Further specification of the quantified value of the component of the method to generate the statistic.
	Value *Quantity `bson:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
	// A variable adjusted for in the adjusted analysis.
	Variable []EvidenceStatisticModelCharacteristicVariable `bson:"variable,omitempty" json:"variable,omitempty" xml:"variable,omitempty"`
	// An attribute of the statistic used as a model characteristic.
	AttributeEstimate []EvidenceStatisticAttributeEstimate `bson:"attributeEstimate,omitempty" json:"attributeEstimate,omitempty"`
}

type EvidenceStatisticModelCharacteristicVariable struct {
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
	// Description of the variable.
	VariableDefinition Reference `bson:"variableDefinition" json:"variableDefinition" xml:"variableDefinition"`
	// How the variable is classified for use in adjusted analysis.
	Handling *EvidenceVariableHandling `bson:"handling,omitempty" json:"handling,omitempty" xml:"handling,omitempty"`
	// Description for grouping of ordinal or polychotomous variables.
	ValueCategory []CodeableConcept `bson:"valueCategory,omitempty" json:"valueCategory,omitempty" xml:"valueCategory,omitempty"`
	// Discrete value for grouping of ordinal or polychotomous variables.
	ValueQuantity []Quantity `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty" xml:"valueQuantity,omitempty"`
	// Range of values for grouping of ordinal or polychotomous variables.
	ValueRange []Range `bson:"valueRange,omitempty" json:"valueRange,omitempty" xml:"valueRange,omitempty"`
}

type EvidenceCertainty struct {
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
	// Textual description of certainty.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Footnotes and/or explanatory notes.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// Aspect of certainty being rated.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Assessment or judgement of the aspect.
	Rating *CodeableConcept `bson:"rating,omitempty" json:"rating,omitempty" xml:"rating,omitempty"`
	// Individual or group who did the rating.
	Rater *string `bson:"rater,omitempty" json:"rater,omitempty" xml:"rater,omitempty"`
	// A domain or subdomain of certainty.
	Subcomponent []EvidenceCertainty `bson:"subcomponent,omitempty" json:"subcomponent,omitempty"`
}

func (r Evidence) ResourceType() string {
	return "Evidence"
}

// MarshalJSON implements json.Marshaler
func (r Evidence) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		Evidence
		ResourceType string `json:"resourceType"`
	}{
		Evidence:     r,
		ResourceType: r.ResourceType(),
	})
}
