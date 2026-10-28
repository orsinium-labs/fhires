package fhires

import v2 "encoding/json/v2"

// This resource allows for the definition of various types of plans as a sharable, consumable, and executable artifact.
//
// The resource is general enough to support the description of a broad range of clinical and non-clinical artifacts such as clinical decision support rules, order sets, protocols, and drug quality specifications.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/PlanDefinition
type PlanDefinition struct {
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
	// An absolute URI that is used to identify this plan definition when it is referenced in a specification, model, design or an instance; also called its canonical identifier.
	//
	// This SHOULD be globally unique and SHOULD be a literal address at which an authoritative instance of this plan definition is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the plan definition is stored on different servers.
	//
	// Comment: Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred.  Multiple instances may share the same URL if they have a distinct version.
	//
	// The determination of when to create a new version of a resource (same url, new version) vs. defining a new artifact is up to the author.  Considerations for making this decision are found in [Technical and Business Versions](resource.html#versions).
	//
	// In some cases, the resource can no longer be found at the stated url, but the url itself cannot change. Implementations can use the [meta.source](resource.html#meta) element to indicate where the current master source of the resource can be found.
	URL *string `bson:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
	// A formal identifier that is used to identify this plan definition when it is represented in other formats, or referenced in a specification, model, design or an instance.
	//
	// Comment: Typically, this is used for identifiers that can go in an HL7 V3 II (instance identifier) data type, and can then identify this plan definition outside of FHIR, where it is not possible to use the logical URI.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The identifier that is used to identify this version of the plan definition when it is referenced in a specification, model, design or instance.
	//
	// This is an arbitrary value managed by the plan definition author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence. To provide a version consistent with the Decision Support Service specification, use the format Major.Minor.Revision (e.g. 1.0.0). For more information on versioning knowledge assets, refer to the Decision Support Service specification. Note that a version is required for non-experimental active artifacts.
	//
	// Comment: There may be different plan definition instances that have the same identifier but different versions.  The version can be appended to the url in a reference to allow a reference to a particular business version of the plan definition with the format [url]|[version]. The version SHOULD NOT contain a '#' - see [Business Version](resource.html#bv-format).
	Version *string `bson:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
	// Indicates the mechanism used to compare versions to determine which is more current.
	//
	// Comment: If set as a string, this is a FHIRPath expression that has two additional context variables passed in - %version1 and %version2 and will return a negative number if version1 is newer, a positive number if version2 and a 0 if the version ordering can't be successfully be determined.
	VersionAlgorithmString *string `bson:"versionAlgorithmString,omitempty" json:"versionAlgorithmString,omitempty" xml:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding `bson:"versionAlgorithmCoding,omitempty" json:"versionAlgorithmCoding,omitempty" xml:"versionAlgorithmCoding,omitempty"`
	// A natural language name identifying the plan definition.
	//
	// This name should be usable as an identifier for the module by machine processing applications such as code generation.
	//
	// Comment: The name is not expected to be globally unique. The name should be a simple alphanumeric type name to ensure that it is machine-processing friendly.
	Name *string `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// A short, descriptive, user-friendly title for the plan definition.
	//
	// Comment: This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc.
	Title *string `bson:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// An explanatory or alternate title for the plan definition giving additional information about its content.
	Subtitle *string `bson:"subtitle,omitempty" json:"subtitle,omitempty" xml:"subtitle,omitempty"`
	// A high-level category for the plan definition that distinguishes the kinds of systems that would be interested in the plan definition.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The status of this plan definition. Enables tracking the life-cycle of the content.
	//
	// Comment: Allows filtering of plan definitions that are appropriate for use versus not.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Status PublicationStatus `bson:"status" json:"status" xml:"status"`
	// A Boolean value to indicate that this plan definition is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	//
	// Comment: Allows filtering of plan definitions that are appropriate for use versus not.
	Experimental *bool `bson:"experimental,omitempty" json:"experimental,omitempty" xml:"experimental,omitempty"`
	// A code, group definition, or canonical reference that describes  or identifies the intended subject of the plan definition.
	//
	// Canonical references are allowed to support the definition of protocols for drug and substance quality specifications, and is allowed to reference a MedicinalProductDefinition, SubstanceDefinition, AdministrableProductDefinition, ManufacturedItemDefinition, or PackagedProductDefinition resource.
	//
	// Comment: Note that the choice of canonical for the subject element was introduced in R4B to support pharmaceutical quality use cases. To ensure as much backwards-compatibility as possible, it is recommended to only use the new canonical type with these use cases.
	SubjectCodeableConcept *CodeableConcept `bson:"subjectCodeableConcept,omitempty" json:"subjectCodeableConcept,omitempty" xml:"subjectCodeableConcept,omitempty"`
	SubjectReference       *Reference       `bson:"subjectReference,omitempty" json:"subjectReference,omitempty" xml:"subjectReference,omitempty"`
	SubjectCanonical       *string          `bson:"subjectCanonical,omitempty" json:"subjectCanonical,omitempty" xml:"subjectCanonical,omitempty"`
	// The date  (and optionally time) when the plan definition was last significantly changed.
	//
	// The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the plan definition changes.
	//
	// Comment: The date is often not tracked until the resource is published, but may be present on draft content. Note that this is not the same as the resource last-modified-date, since the resource may be a secondary representation of the plan definition. Additional specific dates may be added as extensions or be found by consulting Provenances associated with past versions of the resource.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Date *string `bson:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// The name of the organization or individual responsible for the release and ongoing maintenance of the plan definition.
	//
	// Comment: Usually an organization but may be an individual. The publisher (or steward) of the plan definition is the organization or individual primarily responsible for the maintenance and upkeep of the plan definition. This is not necessarily the same individual or organization that developed and initially authored the content. The publisher is the primary point of contact for questions or issues with the plan definition. This item SHOULD be populated unless the information is available from context.
	Publisher *string `bson:"publisher,omitempty" json:"publisher,omitempty" xml:"publisher,omitempty"`
	// Contact details to assist a user in finding and communicating with the publisher.
	//
	// Comment: May be a web site, an email address, a telephone number, etc.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty" xml:"contact,omitempty"`
	// A free text natural language description of the plan definition from a consumer's perspective.
	//
	// Comment: This description can be used to capture details such as comments about misuse, instructions for clinical use and interpretation, literature references, examples from the paper world, etc. It is not a rendering of the plan definition as conveyed in the 'text' field of the resource itself. This item SHOULD be populated unless the information is available from context (e.g. the language of the plan definition is presumed to be the predominant language in the place the plan definition was created).
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// The content was developed with a focus and intent of supporting the contexts that are listed.
	//
	// These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate plan definition instances.
	//
	// Comment: When multiple useContexts are specified, there is no expectation that all or any of the contexts apply.
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty" xml:"useContext,omitempty"`
	// A legal or geographic region in which the plan definition is intended to be used.
	//
	// Comment: It may be possible for the plan definition to be used in jurisdictions other than those for which it was originally designed or intended.
	//
	// DEPRECATION NOTE: For consistency, implementations are encouraged to migrate to using the new 'jurisdiction' code in the useContext element.  (I.e. useContext.code indicating http://terminology.hl7.org/CodeSystem/usage-context-type#jurisdiction and useContext.valueCodeableConcept indicating the jurisdiction.)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty" xml:"jurisdiction,omitempty"`
	// Explanation of why this plan definition is needed and why it has been designed as it has.
	//
	// Comment: This element does not describe the usage of the plan definition. Instead, it provides traceability of ''why'' the resource is either needed or ''why'' it is defined as it is.  This may be used to point to source materials or specifications that drove the structure of this plan definition.
	Purpose *string `bson:"purpose,omitempty" json:"purpose,omitempty" xml:"purpose,omitempty"`
	// A detailed description of how the plan definition is used from a clinical perspective.
	Usage *string `bson:"usage,omitempty" json:"usage,omitempty" xml:"usage,omitempty"`
	// A copyright statement relating to the plan definition and/or its contents.
	//
	// Copyright statements are generally legal restrictions on the use and publishing of the plan definition.
	//
	// Comment: The short copyright declaration (e.g. (c) '2015+ xyz organization' should be sent in the copyrightLabel element.
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
	// The period during which the plan definition content was or is planned to be in active use.
	//
	// Comment: The effective period for a plan definition  determines when the content is applicable for usage and is independent of publication and review dates. For example, a plan definition intended to be used for the year 2016 might be published in 2015.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	EffectivePeriod *Period `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty" xml:"effectivePeriod,omitempty"`
	// Descriptive topics related to the content of the plan definition.
	//
	// Topics provide a high-level categorization of the definition that can be useful for filtering and searching.
	//
	// Comment: DEPRECATION NOTE: For consistency, implementations are encouraged to migrate to using the new 'topic' code in the useContext element.  (I.e. useContext.code indicating http://terminology.hl7.org/CodeSystem/usage-context-type#topic and useContext.valueCodeableConcept indicating the topic)
	Topic []CodeableConcept `bson:"topic,omitempty" json:"topic,omitempty" xml:"topic,omitempty"`
	// An individiual or organization primarily involved in the creation and maintenance of the content.
	Author []ContactDetail `bson:"author,omitempty" json:"author,omitempty" xml:"author,omitempty"`
	// An individual or organization primarily responsible for internal coherence of the content.
	Editor []ContactDetail `bson:"editor,omitempty" json:"editor,omitempty" xml:"editor,omitempty"`
	// An individual or organization asserted by the publisher to be primarily responsible for review of some aspect of the content.
	//
	// Comment: See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Reviewer []ContactDetail `bson:"reviewer,omitempty" json:"reviewer,omitempty" xml:"reviewer,omitempty"`
	// An individual or organization asserted by the publisher to be responsible for officially endorsing the content for use in some setting.
	//
	// Comment: See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Endorser []ContactDetail `bson:"endorser,omitempty" json:"endorser,omitempty" xml:"endorser,omitempty"`
	// Related artifacts such as additional documentation, justification, or bibliographic references.
	//
	// Comment: Each related artifact is either an attachment, or a reference to another resource, but not both.
	RelatedArtifact []RelatedArtifact `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty" xml:"relatedArtifact,omitempty"`
	// A reference to a Library resource containing any formal logic used by the plan definition.
	Library []string `bson:"library,omitempty" json:"library,omitempty" xml:"library,omitempty"`
	// A goal describes an expected outcome that activities within the plan are intended to achieve.
	//
	// For example, weight loss, restoring an activity of daily living, obtaining herd immunity via immunization, meeting a process improvement objective, meeting the acceptance criteria for a test as specified by a quality specification, etc.
	Goal []PlanDefinitionGoal `bson:"goal,omitempty" json:"goal,omitempty" xml:"goal,omitempty"`
	// Actors represent the individuals or groups involved in the execution of the defined set of activities.
	Actor []PlanDefinitionActor `bson:"actor,omitempty" json:"actor,omitempty" xml:"actor,omitempty"`
	// An action or group of actions to be taken as part of the plan.
	//
	// For example, in clinical care, an action would be to prescribe a particular indicated medication, or perform a particular test as appropriate. In pharmaceutical quality, an action would be the test that needs to be performed on a drug product as defined in the quality specification.
	//
	// Comment: Note that there is overlap between many of the elements defined here and the ActivityDefinition resource. When an ActivityDefinition is referenced (using the definition element), the overlapping elements in the plan override the content of the referenced ActivityDefinition unless otherwise documented in the specific elements. See the PlanDefinition resource for more detailed information.
	Action []PlanDefinitionAction `bson:"action,omitempty" json:"action,omitempty" xml:"action,omitempty"`
	// If a CodeableConcept is present, it indicates the pre-condition for performing the service.
	//
	//  For example "pain", "on flare-up", etc.
	AsNeededBoolean         *bool            `bson:"asNeededBoolean,omitempty" json:"asNeededBoolean,omitempty" xml:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept *CodeableConcept `bson:"asNeededCodeableConcept,omitempty" json:"asNeededCodeableConcept,omitempty" xml:"asNeededCodeableConcept,omitempty"`
}

type PlanDefinitionGoal struct {
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
	// Indicates a category the goal falls within.
	Category *CodeableConcept `bson:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
	// Human-readable and/or coded description of a specific desired objective of care, such as "control blood pressure" or "negotiate an obstacle course" or "dance with child at wedding".
	//
	// Comment: If no code is available, use CodeableConcept.text.
	Description CodeableConcept `bson:"description" json:"description" xml:"description"`
	// Identifies the expected level of importance associated with reaching/sustaining the defined goal.
	Priority *CodeableConcept `bson:"priority,omitempty" json:"priority,omitempty" xml:"priority,omitempty"`
	// The event after which the goal should begin being pursued.
	Start *CodeableConcept `bson:"start,omitempty" json:"start,omitempty" xml:"start,omitempty"`
	// Identifies problems, conditions, issues, or concerns the goal is intended to address.
	Addresses []CodeableConcept `bson:"addresses,omitempty" json:"addresses,omitempty" xml:"addresses,omitempty"`
	// Didactic or other informational resources associated with the goal that provide further supporting information about the goal.
	//
	// Information resources can include inline text commentary and links to web resources.
	Documentation []RelatedArtifact `bson:"documentation,omitempty" json:"documentation,omitempty" xml:"documentation,omitempty"`
	// Indicates what should be done and within what timeframe.
	Target []PlanDefinitionGoalTarget `bson:"target,omitempty" json:"target,omitempty" xml:"target,omitempty"`
}

type PlanDefinitionGoalTarget struct {
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
	// The parameter whose value is to be tracked, e.g.
	//
	// body weight, blood pressure, or hemoglobin A1c level.
	Measure *CodeableConcept `bson:"measure,omitempty" json:"measure,omitempty" xml:"measure,omitempty"`
	// The target value of the measure to be achieved to signify fulfillment of the goal, e.g.
	//
	// 150 pounds or 7.0%, or in the case of pharmaceutical quality - NMT 0.6%, Clear solution, etc. Either the high or low or both values of the range can be specified. When a low value is missing, it indicates that the goal is achieved at any value at or below the high value. Similarly, if the high value is missing, it indicates that the goal is achieved at any value at or above the low value.
	//
	// Comment: A CodeableConcept with just a text would be used instead of a string if the field was usually coded, or if the type associated with the goal.target.measure defines a coded value.
	DetailQuantity        *Quantity        `bson:"detailQuantity,omitempty" json:"detailQuantity,omitempty" xml:"detailQuantity,omitempty"`
	DetailRange           *Range           `bson:"detailRange,omitempty" json:"detailRange,omitempty" xml:"detailRange,omitempty"`
	DetailCodeableConcept *CodeableConcept `bson:"detailCodeableConcept,omitempty" json:"detailCodeableConcept,omitempty" xml:"detailCodeableConcept,omitempty"`
	DetailString          *string          `bson:"detailString,omitempty" json:"detailString,omitempty" xml:"detailString,omitempty"`
	DetailBoolean         *bool            `bson:"detailBoolean,omitempty" json:"detailBoolean,omitempty" xml:"detailBoolean,omitempty"`
	DetailInteger         *int             `bson:"detailInteger,omitempty" json:"detailInteger,omitempty" xml:"detailInteger,omitempty"`
	DetailRatio           *Ratio           `bson:"detailRatio,omitempty" json:"detailRatio,omitempty" xml:"detailRatio,omitempty"`
	// Indicates the timeframe after the start of the goal in which the goal should be met.
	Due *Duration `bson:"due,omitempty" json:"due,omitempty" xml:"due,omitempty"`
}

type PlanDefinitionActor struct {
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
	// A descriptive label for the actor.
	Title *string `bson:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// A description of how the actor fits into the overall actions of the plan definition.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// The characteristics of the candidates that could serve as the actor.
	Option []PlanDefinitionActorOption `bson:"option" json:"option" xml:"option"`
}

type PlanDefinitionActorOption struct {
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
	// The type of participant in the action.
	Type *ActionParticipantType `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The type of participant in the action.
	TypeCanonical *string `bson:"typeCanonical,omitempty" json:"typeCanonical,omitempty" xml:"typeCanonical,omitempty"`
	// The type of participant in the action.
	//
	// Comment: When this element is a reference, it SHOULD be a reference to a definitional resource (for example, a location type, rather than a specific location).
	TypeReference *Reference `bson:"typeReference,omitempty" json:"typeReference,omitempty" xml:"typeReference,omitempty"`
	// The role the participant should play in performing the described action.
	Role *CodeableConcept `bson:"role,omitempty" json:"role,omitempty" xml:"role,omitempty"`
}

type PlanDefinitionAction struct {
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
	// An identifier that is unique within the PlanDefinition to allow linkage within the realized CarePlan and/or RequestOrchestration.
	LinkID *string `bson:"linkId,omitempty" json:"linkId,omitempty" xml:"linkId,omitempty"`
	// A user-visible prefix for the action.
	//
	// For example a section or item numbering such as 1. or A.
	Prefix *string `bson:"prefix,omitempty" json:"prefix,omitempty" xml:"prefix,omitempty"`
	// The textual description of the action displayed to a user.
	//
	// For example, when the action is a test to be performed, the title would be the title of the test such as Assay by HPLC.
	Title *string `bson:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// A brief description of the action used to provide a summary to display to the user.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// A text equivalent of the action to be performed.
	//
	// This provides a human-interpretable description of the action when the definition is consumed by a system that might not be capable of interpreting it dynamically.
	TextEquivalent *string `bson:"textEquivalent,omitempty" json:"textEquivalent,omitempty" xml:"textEquivalent,omitempty"`
	// Indicates how quickly the action should be addressed with respect to other actions.
	Priority *RequestPriority `bson:"priority,omitempty" json:"priority,omitempty" xml:"priority,omitempty"`
	// A code that provides a meaning, grouping, or classification for the action or action group.
	//
	// For example, a section may have a LOINC code for the section of a documentation template. In pharmaceutical quality, an action (Test) such as pH could be classified as a physical property.
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// A description of why this action is necessary or appropriate.
	//
	// Comment: This is different than the clinical evidence documentation, it's an actual business description of the reason for performing the action.
	Reason []CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty" xml:"reason,omitempty"`
	// Didactic or other informational resources associated with the action that can be provided to the CDS recipient.
	//
	// Information resources can include inline text commentary and links to web resources.
	Documentation []RelatedArtifact `bson:"documentation,omitempty" json:"documentation,omitempty" xml:"documentation,omitempty"`
	// Identifies goals that this action supports.
	//
	// The reference must be to a goal element defined within this plan definition. In pharmaceutical quality, a goal represents acceptance criteria (Goal) for a given action (Test), so the goalId would be the unique id of a defined goal element establishing the acceptance criteria for the action.
	GoalID []string `bson:"goalId,omitempty" json:"goalId,omitempty" xml:"goalId,omitempty"`
	// A code, group definition, or canonical reference that describes the intended subject of the action and its children, if any.
	//
	// Canonical references are allowed to support the definition of protocols for drug and substance quality specifications, and is allowed to reference a MedicinalProductDefinition, SubstanceDefinition, AdministrableProductDefinition, ManufacturedItemDefinition, or PackagedProductDefinition resource.
	//
	// Comment: The subject of an action overrides the subject at a parent action or on the root of the PlanDefinition if specified.
	//
	// In addition, because the subject needs to be resolved during realization, use of subjects in actions (or in the ActivityDefinition referenced by the action) resolves based on the set of subjects supplied in context and by type (i.e. the patient subject would resolve to a resource of type Patient).
	SubjectCodeableConcept *CodeableConcept `bson:"subjectCodeableConcept,omitempty" json:"subjectCodeableConcept,omitempty" xml:"subjectCodeableConcept,omitempty"`
	SubjectReference       *Reference       `bson:"subjectReference,omitempty" json:"subjectReference,omitempty" xml:"subjectReference,omitempty"`
	SubjectCanonical       *string          `bson:"subjectCanonical,omitempty" json:"subjectCanonical,omitempty" xml:"subjectCanonical,omitempty"`
	// A description of when the action should be triggered.
	//
	// When multiple triggers are specified on an action, any triggering event invokes the action.
	Trigger []TriggerDefinition `bson:"trigger,omitempty" json:"trigger,omitempty" xml:"trigger,omitempty"`
	// An expression that describes applicability criteria or start/stop conditions for the action.
	//
	// Comment: When multiple conditions of the same kind are present, the effects are combined using AND semantics, so the overall condition is true only if all the conditions are true.
	Condition []PlanDefinitionActionCondition `bson:"condition,omitempty" json:"condition,omitempty" xml:"condition,omitempty"`
	// Defines input data requirements for the action.
	Input []PlanDefinitionActionInput `bson:"input,omitempty" json:"input,omitempty" xml:"input,omitempty"`
	// Defines the outputs of the action, if any.
	Output []PlanDefinitionActionOutput `bson:"output,omitempty" json:"output,omitempty" xml:"output,omitempty"`
	// A relationship to another action such as "before" or "30-60 minutes after start of".
	//
	// Comment: When an action depends on multiple actions, the meaning is that all actions are dependencies, rather than that any of the actions are a dependency.
	RelatedAction []PlanDefinitionActionRelatedAction `bson:"relatedAction,omitempty" json:"relatedAction,omitempty" xml:"relatedAction,omitempty"`
	// An optional value describing when the action should be performed.
	//
	// Comment: The intent of the timing element is to provide timing when the action should be performed. As a definitional resource, this timing is interpreted as part of an apply operation so that the timing of the result actions in a CarePlan or RequestOrchestration, for example, would be specified by evaluating the timing definition in the context of the apply and setting the resulting timing on the appropriate elements of the target resource. If the timing is an Age, the activity is expected to be performed when the subject is the given Age. When the timing is a Duration, the activity is expected to be performed within the specified duration from the apply. When the timing is a Range, it may be a range of Ages or Durations, providing a range for the expected timing of the resulting activity. When the timing is a Timing, it is establishing a schedule for the timing of the resulting activity.
	TimingAge      *Age      `bson:"timingAge,omitempty" json:"timingAge,omitempty" xml:"timingAge,omitempty"`
	TimingDuration *Duration `bson:"timingDuration,omitempty" json:"timingDuration,omitempty" xml:"timingDuration,omitempty"`
	TimingRange    *Range    `bson:"timingRange,omitempty" json:"timingRange,omitempty" xml:"timingRange,omitempty"`
	TimingTiming   *Timing   `bson:"timingTiming,omitempty" json:"timingTiming,omitempty" xml:"timingTiming,omitempty"`
	// Identifies the facility where the action will occur; e.g.
	//
	// home, hospital, specific clinic, etc.
	//
	// Comment: May reference a specific clinical location or may just identify a type of location.
	Location *CodeableReference `bson:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// Indicates who should participate in performing the action described.
	Participant []PlanDefinitionActionParticipant `bson:"participant,omitempty" json:"participant,omitempty" xml:"participant,omitempty"`
	// The type of action to perform (create, update, remove).
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Defines the grouping behavior for the action and its children.
	GroupingBehavior *ActionGroupingBehavior `bson:"groupingBehavior,omitempty" json:"groupingBehavior,omitempty" xml:"groupingBehavior,omitempty"`
	// Defines the selection behavior for the action and its children.
	SelectionBehavior *ActionSelectionBehavior `bson:"selectionBehavior,omitempty" json:"selectionBehavior,omitempty" xml:"selectionBehavior,omitempty"`
	// Defines the required behavior for the action.
	RequiredBehavior *ActionRequiredBehavior `bson:"requiredBehavior,omitempty" json:"requiredBehavior,omitempty" xml:"requiredBehavior,omitempty"`
	// Defines whether the action should usually be preselected.
	PrecheckBehavior *ActionPrecheckBehavior `bson:"precheckBehavior,omitempty" json:"precheckBehavior,omitempty" xml:"precheckBehavior,omitempty"`
	// Defines whether the action can be selected multiple times.
	CardinalityBehavior *ActionCardinalityBehavior `bson:"cardinalityBehavior,omitempty" json:"cardinalityBehavior,omitempty" xml:"cardinalityBehavior,omitempty"`
	// A reference to an ActivityDefinition that describes the action to be taken in detail, a MessageDefinition describing a message to be snet, a PlanDefinition that describes a series of actions to be taken, a Questionnaire that should be filled out, a SpecimenDefinition describing a specimen to be collected, or an ObservationDefinition that specifies what observation should be captured.
	//
	// Comment: Note that the definition is optional, and if no definition is specified, a dynamicValue with a root ($this) path can be used to define the entire resource dynamically.
	DefinitionCanonical *string `bson:"definitionCanonical,omitempty" json:"definitionCanonical,omitempty" xml:"definitionCanonical,omitempty"`
	DefinitionUri       *string `bson:"definitionUri,omitempty" json:"definitionUri,omitempty" xml:"definitionUri,omitempty"`
	// A reference to a StructureMap resource that defines a transform that can be executed to produce the intent resource using the ActivityDefinition instance as the input.
	//
	// Comment: Note that when a referenced ActivityDefinition also defines a transform, the transform specified here generally takes precedence. In addition, if both a transform and dynamic values are specific, the dynamic values are applied to the result of the transform.
	Transform *string `bson:"transform,omitempty" json:"transform,omitempty" xml:"transform,omitempty"`
	// Customizations that should be applied to the statically defined resource.
	//
	// For example, if the dosage of a medication must be computed based on the patient's weight, a customization would be used to specify an expression that calculated the weight, and the path on the resource that would contain the result.
	//
	// Comment: Dynamic values are applied in the order in which they are defined in the PlanDefinition resource. Note that when dynamic values are also specified by a referenced ActivityDefinition, the dynamicValues from the ActivityDefinition are applied first, followed by the dynamicValues specified here. In addition, if both a transform and dynamic values are specific, the dynamic values are applied to the result of the transform.
	DynamicValue []PlanDefinitionActionDynamicValue `bson:"dynamicValue,omitempty" json:"dynamicValue,omitempty" xml:"dynamicValue,omitempty"`
	// Sub actions that are contained within the action.
	//
	// The behavior of this action determines the functionality of the sub-actions. For example, a selection behavior of at-most-one indicates that of the sub-actions, at most one may be chosen as part of realizing the action definition.
	Action []PlanDefinitionAction `bson:"action,omitempty" json:"action,omitempty"`
}

type PlanDefinitionActionCondition struct {
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
	// The kind of condition.
	//
	// Comment: Applicability criteria are used to determine immediate applicability when a plan definition is applied to a given context. Start and stop criteria are carried through application and used to describe enter/exit criteria for an action.
	Kind ActionConditionKind `bson:"kind" json:"kind" xml:"kind"`
	// An expression that returns true or false, indicating whether the condition is satisfied.
	//
	// Comment: The expression may be inlined or may be a reference to a named expression within a logic library referenced by the library element.
	Expression *Expression `bson:"expression,omitempty" json:"expression,omitempty" xml:"expression,omitempty"`
}

type PlanDefinitionActionInput struct {
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
	// A human-readable label for the data requirement used to label data flows in BPMN or similar diagrams.
	//
	// Also provides a human readable label when rendering the data requirement that conveys its purpose to human readers.
	Title *string `bson:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Defines the data that is to be provided as input to the action.
	Requirement *DataRequirement `bson:"requirement,omitempty" json:"requirement,omitempty" xml:"requirement,omitempty"`
	// Points to an existing input or output element that provides data to this input.
	//
	// Comment: The relatedData element allows indicating that an input to a parent action is an input to specific child actions. It also allows the output of one action to be identified as the input to a different action
	RelatedData *string `bson:"relatedData,omitempty" json:"relatedData,omitempty" xml:"relatedData,omitempty"`
}

type PlanDefinitionActionOutput struct {
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
	// A human-readable label for the data requirement used to label data flows in BPMN or similar diagrams.
	//
	// Also provides a human readable label when rendering the data requirement that conveys its purpose to human readers.
	Title *string `bson:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Defines the data that results as output from the action.
	Requirement *DataRequirement `bson:"requirement,omitempty" json:"requirement,omitempty" xml:"requirement,omitempty"`
	// Points to an existing input or output element that is results as output from the action.
	//
	// Comment: The relatedData element allows indicating that the output of child action is also the output of a parent action. It also allows the input of one action to be identified as the output of a different action
	RelatedData *string `bson:"relatedData,omitempty" json:"relatedData,omitempty" xml:"relatedData,omitempty"`
}

type PlanDefinitionActionRelatedAction struct {
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
	// The element id of the target related action.
	TargetID string `bson:"targetId" json:"targetId" xml:"targetId"`
	// The relationship of the start of this action to the related action.
	Relationship ActionRelationshipType `bson:"relationship" json:"relationship" xml:"relationship"`
	// The relationship of the end of this action to the related action.
	EndRelationship *ActionRelationshipType `bson:"endRelationship,omitempty" json:"endRelationship,omitempty" xml:"endRelationship,omitempty"`
	// A duration or range of durations to apply to the relationship.
	//
	// For example, 30-60 minutes before.
	OffsetDuration *Duration `bson:"offsetDuration,omitempty" json:"offsetDuration,omitempty" xml:"offsetDuration,omitempty"`
	OffsetRange    *Range    `bson:"offsetRange,omitempty" json:"offsetRange,omitempty" xml:"offsetRange,omitempty"`
}

type PlanDefinitionActionParticipant struct {
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
	// A reference to the id element of the actor who will participate in this action.
	ActorID *string `bson:"actorId,omitempty" json:"actorId,omitempty" xml:"actorId,omitempty"`
	// The type of participant in the action.
	Type *ActionParticipantType `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The type of participant in the action.
	TypeCanonical *string `bson:"typeCanonical,omitempty" json:"typeCanonical,omitempty" xml:"typeCanonical,omitempty"`
	// The type of participant in the action.
	//
	// Comment: When this element is a reference, it SHOULD be a reference to a definitional resource (for example, a location type, rather than a specific location).
	TypeReference *Reference `bson:"typeReference,omitempty" json:"typeReference,omitempty" xml:"typeReference,omitempty"`
	// The role the participant should play in performing the described action.
	Role *CodeableConcept `bson:"role,omitempty" json:"role,omitempty" xml:"role,omitempty"`
	// Indicates how the actor will be involved in the action - author, reviewer, witness, etc.
	Function *CodeableConcept `bson:"function,omitempty" json:"function,omitempty" xml:"function,omitempty"`
}

type PlanDefinitionActionDynamicValue struct {
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
	// The path to the element to be customized.
	//
	// This is the path on the resource that will hold the result of the calculation defined by the expression. The specified path SHALL be a FHIRPath resolvable on the specified target type of the ActivityDefinition, and SHALL consist only of identifiers, constant indexers, and a restricted subset of functions. The path is allowed to contain qualifiers (.) to traverse sub-elements, as well as indexers ([x]) to traverse multiple-cardinality sub-elements (see the [Simple FHIRPath Profile](fhirpath.html#simple) for full details).
	//
	// Comment: To specify the path to the current action being realized, the %action environment variable is available in this path. For example, to specify the description element of the target action, the path would be %action.description. The path attribute contains a [Simple FHIRPath Subset](fhirpath.html#simple) that allows path traversal, but not calculation.
	Path *string `bson:"path,omitempty" json:"path,omitempty" xml:"path,omitempty"`
	// An expression specifying the value of the customized element.
	//
	// Comment: The expression may be inlined or may be a reference to a named expression within a logic library referenced by the library element.
	Expression *Expression `bson:"expression,omitempty" json:"expression,omitempty" xml:"expression,omitempty"`
}

func (r PlanDefinition) ResourceType() string {
	return "PlanDefinition"
}

// MarshalJSON implements json.Marshaler
func (r PlanDefinition) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		PlanDefinition
		ResourceType string `json:"resourceType"`
	}{
		PlanDefinition: r,
		ResourceType:   r.ResourceType(),
	})
}
