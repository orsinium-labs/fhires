package fhires

import "encoding/json/v2"

// This resource allows for the definition of some activity to be performed, independent of a particular patient, practitioner, or other performance context.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/ActivityDefinition
type ActivityDefinition struct {
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
	// An absolute URI that is used to identify this activity definition when it is referenced in a specification, model, design or an instance; also called its canonical identifier.
	//
	// This SHOULD be globally unique and SHOULD be a literal address at which an authoritative instance of this activity definition is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the activity definition is stored on different servers.
	//
	// Comment: Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred.  Multiple instances may share the same URL if they have a distinct version.
	//
	// The determination of when to create a new version of a resource (same url, new version) vs. defining a new artifact is up to the author.  Considerations for making this decision are found in [Technical and Business Versions](resource.html#versions).
	//
	// In some cases, the resource can no longer be found at the stated url, but the url itself cannot change. Implementations can use the [meta.source](resource.html#meta) element to indicate where the current master source of the resource can be found.
	URL *string `bson:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
	// A formal identifier that is used to identify this activity definition when it is represented in other formats, or referenced in a specification, model, design or an instance.
	//
	// Comment: Typically, this is used for identifiers that can go in an HL7 V3 II (instance identifier) data type, and can then identify this activity definition outside of FHIR, where it is not possible to use the logical URI.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The identifier that is used to identify this version of the activity definition when it is referenced in a specification, model, design or instance.
	//
	// This is an arbitrary value managed by the activity definition author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence. To provide a version consistent with the Decision Support Service specification, use the format Major.Minor.Revision (e.g. 1.0.0). For more information on versioning knowledge assets, refer to the Decision Support Service specification. Note that a version is required for non-experimental active assets.
	//
	// Comment: There may be different activity definition instances that have the same identifier but different versions.  The version can be appended to the url in a reference to allow a reference to a particular business version of the activity definition with the format [url]|[version]. The version SHOULD NOT contain a '#' - see [Business Version](resource.html#bv-format).
	Version *string `bson:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
	// Indicates the mechanism used to compare versions to determine which is more current.
	//
	// Comment: If set as a string, this is a FHIRPath expression that has two additional context variables passed in - %version1 and %version2 and will return a negative number if version1 is newer, a positive number if version2 and a 0 if the version ordering can't be successfully be determined.
	VersionAlgorithmString *string `bson:"versionAlgorithmString,omitempty" json:"versionAlgorithmString,omitempty" xml:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding `bson:"versionAlgorithmCoding,omitempty" json:"versionAlgorithmCoding,omitempty" xml:"versionAlgorithmCoding,omitempty"`
	// A natural language name identifying the activity definition.
	//
	// This name should be usable as an identifier for the module by machine processing applications such as code generation.
	//
	// Comment: The name is not expected to be globally unique. The name should be a simple alphanumeric type name to ensure that it is machine-processing friendly.
	Name *string `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// A short, descriptive, user-friendly title for the activity definition.
	//
	// Comment: This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc.
	Title *string `bson:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// An explanatory or alternate title for the activity definition giving additional information about its content.
	Subtitle *string `bson:"subtitle,omitempty" json:"subtitle,omitempty" xml:"subtitle,omitempty"`
	// The status of this activity definition. Enables tracking the life-cycle of the content.
	//
	// Comment: Allows filtering of activity definitions that are appropriate for use versus not.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Status PublicationStatus `bson:"status" json:"status" xml:"status"`
	// A Boolean value to indicate that this activity definition is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	//
	// Comment: Allows filtering of activity definitions that are appropriate for use versus not.
	Experimental *bool `bson:"experimental,omitempty" json:"experimental,omitempty" xml:"experimental,omitempty"`
	// A code, group definition, or canonical reference that describes  or identifies the intended subject of the activity being defined.
	//
	//  Canonical references are allowed to support the definition of protocols for drug and substance quality specifications, and is allowed to reference a MedicinalProductDefinition, SubstanceDefinition, AdministrableProductDefinition, ManufacturedItemDefinition, or PackagedProductDefinition resource.
	//
	// Comment: Note that the choice of canonical for the subject element was introduced in R4B to support pharmaceutical quality use cases. To ensure as much backwards-compatibility as possible, it is recommended to only use the new canonical type with these use cases.
	SubjectCodeableConcept *CodeableConcept `bson:"subjectCodeableConcept,omitempty" json:"subjectCodeableConcept,omitempty" xml:"subjectCodeableConcept,omitempty"`
	SubjectReference       *Reference       `bson:"subjectReference,omitempty" json:"subjectReference,omitempty" xml:"subjectReference,omitempty"`
	SubjectCanonical       *string          `bson:"subjectCanonical,omitempty" json:"subjectCanonical,omitempty" xml:"subjectCanonical,omitempty"`
	// The date  (and optionally time) when the activity definition was last significantly changed.
	//
	// The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the activity definition changes.
	//
	// Comment: The date is often not tracked until the resource is published, but may be present on draft content. Note that this is not the same as the resource last-modified-date, since the resource may be a secondary representation of the activity definition. Additional specific dates may be added as extensions or be found by consulting Provenances associated with past versions of the resource.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Date *string `bson:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// The name of the organization or individual responsible for the release and ongoing maintenance of the activity definition.
	//
	// Comment: Usually an organization but may be an individual. The publisher (or steward) of the activity definition is the organization or individual primarily responsible for the maintenance and upkeep of the activity definition. This is not necessarily the same individual or organization that developed and initially authored the content. The publisher is the primary point of contact for questions or issues with the activity definition. This item SHOULD be populated unless the information is available from context.
	Publisher *string `bson:"publisher,omitempty" json:"publisher,omitempty" xml:"publisher,omitempty"`
	// Contact details to assist a user in finding and communicating with the publisher.
	//
	// Comment: May be a web site, an email address, a telephone number, etc.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty" xml:"contact,omitempty"`
	// A free text natural language description of the activity definition from a consumer's perspective.
	//
	// Comment: This description can be used to capture details such as comments about misuse, instructions for clinical use and interpretation, literature references, examples from the paper world, etc. It is not a rendering of the activity definition as conveyed in the 'text' field of the resource itself. This item SHOULD be populated unless the information is available from context (e.g. the language of the activity definition is presumed to be the predominant language in the place the activity definition was created).
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// The content was developed with a focus and intent of supporting the contexts that are listed.
	//
	// These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate activity definition instances.
	//
	// Comment: When multiple useContexts are specified, there is no expectation that all or any of the contexts apply.
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty" xml:"useContext,omitempty"`
	// A legal or geographic region in which the activity definition is intended to be used.
	//
	// Comment: It may be possible for the activity definition to be used in jurisdictions other than those for which it was originally designed or intended.
	//
	// DEPRECATION NOTE: For consistency, implementations are encouraged to migrate to using the new 'jurisdiction' code in the useContext element.  (I.e. useContext.code indicating http://terminology.hl7.org/CodeSystem/usage-context-type#jurisdiction and useContext.valueCodeableConcept indicating the jurisdiction.)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty" xml:"jurisdiction,omitempty"`
	// Explanation of why this activity definition is needed and why it has been designed as it has.
	//
	// Comment: This element does not describe the usage of the activity definition. Instead, it provides traceability of ''why'' the resource is either needed or ''why'' it is defined as it is.  This may be used to point to source materials or specifications that drove the structure of this activity definition.
	Purpose *string `bson:"purpose,omitempty" json:"purpose,omitempty" xml:"purpose,omitempty"`
	// A detailed description of how the activity definition is used from a clinical perspective.
	Usage *string `bson:"usage,omitempty" json:"usage,omitempty" xml:"usage,omitempty"`
	// A copyright statement relating to the activity definition and/or its contents.
	//
	// Copyright statements are generally legal restrictions on the use and publishing of the activity definition.
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
	// The period during which the activity definition content was or is planned to be in active use.
	//
	// Comment: The effective period for an activity definition  determines when the content is applicable for usage and is independent of publication and review dates. For example, a activity intended to be used for the year 2016 might be published in 2015.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	EffectivePeriod *Period `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty" xml:"effectivePeriod,omitempty"`
	// Descriptive topics related to the content of the activity.
	//
	// Topics provide a high-level categorization of the activity that can be useful for filtering and searching.
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
	// A reference to a Library resource containing any formal logic used by the activity definition.
	Library []string `bson:"library,omitempty" json:"library,omitempty" xml:"library,omitempty"`
	// A description of the kind of resource the activity definition is representing.
	//
	// For example, a MedicationRequest, a ServiceRequest, or a CommunicationRequest.
	//
	// Comment: The kind element may only specify Request resource types to facilitate considering user input as part of processing the result of any automated clinical reasoning processes. To support creation of event resources, such as Observations, RiskAssessments, and DetectedIssues, use a Task resource with the focus of the task set to the event resource to be created. Note that the kind of resource to be created may determine what types of extensions are permitted.
	Kind *RequestResourceTypes `bson:"kind,omitempty" json:"kind,omitempty" xml:"kind,omitempty"`
	// A profile to which the target of the activity definition is expected to conform.
	Profile *string `bson:"profile,omitempty" json:"profile,omitempty" xml:"profile,omitempty"`
	// Detailed description of the type of activity; e.g.
	//
	// What lab test, what procedure, what kind of encounter.
	//
	// Comment: Tends to be less relevant for activities involving particular products.
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// Indicates the level of authority/intentionality associated with the activity and where the request should fit into the workflow chain.
	Intent *RequestIntent `bson:"intent,omitempty" json:"intent,omitempty" xml:"intent,omitempty"`
	// Indicates how quickly the activity  should be addressed with respect to other requests.
	Priority *RequestPriority `bson:"priority,omitempty" json:"priority,omitempty" xml:"priority,omitempty"`
	// Set this to true if the definition is to indicate that a particular activity should NOT be performed.
	//
	// If true, this element should be interpreted to reinforce a negative coding. For example NPO as a code with a doNotPerform of true would still indicate to NOT perform the action.
	//
	// Comment: This element is not intended to be used to communicate a decision support response to cancel an order in progress. That should be done with the "remove" type of a PlanDefinition or RequestOrchestration.
	DoNotPerform *bool `bson:"doNotPerform,omitempty" json:"doNotPerform,omitempty" xml:"doNotPerform,omitempty"`
	// The timing or frequency upon which the described activity is to occur.
	//
	// Comment: The intent of the timing element is to provide timing when the action should be performed. As a definitional resource, this timing is interpreted as part of an apply operation so that the timing of the result actions in a CarePlan or RequestOrchestration, for example, would be specified by evaluating the timing definition in the context of the apply and setting the resulting timing on the appropriate elements of the target resource. If the timing is an Age, the activity is expected to be performed when the subject is the given Age. When the timing is a Duration, the activity is expected to be performed within the specified duration from the apply. When the timing is a Range, it may be a range of Ages or Durations, providing a range for the expected timing of the resulting activity. When the timing is a Timing, it is establishing a schedule for the timing of the resulting activity.
	TimingTiming   *Timing   `bson:"timingTiming,omitempty" json:"timingTiming,omitempty" xml:"timingTiming,omitempty"`
	TimingAge      *Age      `bson:"timingAge,omitempty" json:"timingAge,omitempty" xml:"timingAge,omitempty"`
	TimingRange    *Range    `bson:"timingRange,omitempty" json:"timingRange,omitempty" xml:"timingRange,omitempty"`
	TimingDuration *Duration `bson:"timingDuration,omitempty" json:"timingDuration,omitempty" xml:"timingDuration,omitempty"`
	// If a CodeableConcept is present, it indicates the pre-condition for performing the service.
	//
	//  For example "pain", "on flare-up", etc.
	AsNeededBoolean         *bool            `bson:"asNeededBoolean,omitempty" json:"asNeededBoolean,omitempty" xml:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept *CodeableConcept `bson:"asNeededCodeableConcept,omitempty" json:"asNeededCodeableConcept,omitempty" xml:"asNeededCodeableConcept,omitempty"`
	// Identifies the facility where the activity will occur; e.g.
	//
	// home, hospital, specific clinic, etc.
	//
	// Comment: May reference a specific clinical location or may just identify a type of location.
	Location *CodeableReference `bson:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// Indicates who should participate in performing the action described.
	Participant []ActivityDefinitionParticipant `bson:"participant,omitempty" json:"participant,omitempty" xml:"participant,omitempty"`
	// Identifies the food, drug or other product being consumed or supplied in the activity.
	ProductReference       *Reference       `bson:"productReference,omitempty" json:"productReference,omitempty" xml:"productReference,omitempty"`
	ProductCodeableConcept *CodeableConcept `bson:"productCodeableConcept,omitempty" json:"productCodeableConcept,omitempty" xml:"productCodeableConcept,omitempty"`
	// Identifies the quantity expected to be consumed at once (per dose, per meal, etc.).
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty" xml:"quantity,omitempty"`
	// Provides detailed dosage instructions in the same way that they are described for MedicationRequest resources.
	//
	// Comment: If a dosage instruction is used, the definition should not specify timing or quantity.
	Dosage []Dosage `bson:"dosage,omitempty" json:"dosage,omitempty" xml:"dosage,omitempty"`
	// Indicates the sites on the subject's body where the procedure should be performed (I.e.
	//
	// the target sites).
	//
	// Comment: Only used if not implicit in the code found in ServiceRequest.type.
	BodySite []CodeableConcept `bson:"bodySite,omitempty" json:"bodySite,omitempty" xml:"bodySite,omitempty"`
	// Defines specimen requirements for the action to be performed, such as required specimens for a lab test.
	SpecimenRequirement []string `bson:"specimenRequirement,omitempty" json:"specimenRequirement,omitempty" xml:"specimenRequirement,omitempty"`
	// Defines observation requirements for the action to be performed, such as body weight or surface area.
	ObservationRequirement []string `bson:"observationRequirement,omitempty" json:"observationRequirement,omitempty" xml:"observationRequirement,omitempty"`
	// Defines the observations that are expected to be produced by the action.
	ObservationResultRequirement []string `bson:"observationResultRequirement,omitempty" json:"observationResultRequirement,omitempty" xml:"observationResultRequirement,omitempty"`
	// A reference to a StructureMap resource that defines a transform that can be executed to produce the intent resource using the ActivityDefinition instance as the input.
	//
	// Comment: Note that if both a transform and dynamic values are specified, the dynamic values will be applied to the result of the transform.
	Transform *string `bson:"transform,omitempty" json:"transform,omitempty" xml:"transform,omitempty"`
	// Dynamic values that will be evaluated to produce values for elements of the resulting resource.
	//
	// For example, if the dosage of a medication must be computed based on the patient's weight, a dynamic value would be used to specify an expression that calculated the weight, and the path on the request resource that would contain the result.
	//
	// Comment: Dynamic values are applied in the order in which they are defined in the ActivityDefinition. Note that if both a transform and dynamic values are specified, the dynamic values will be applied to the result of the transform.
	DynamicValue []ActivityDefinitionDynamicValue `bson:"dynamicValue,omitempty" json:"dynamicValue,omitempty" xml:"dynamicValue,omitempty"`
}

type ActivityDefinitionParticipant struct {
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
	// Indicates how the actor will be involved in the action - author, reviewer, witness, etc.
	Function *CodeableConcept `bson:"function,omitempty" json:"function,omitempty" xml:"function,omitempty"`
}

type ActivityDefinitionDynamicValue struct {
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
	// Comment: The path attribute contains a [Simple FHIRPath Subset](fhirpath.html#simple) that allows path traversal, but not calculation.
	Path string `bson:"path" json:"path" xml:"path"`
	// An expression specifying the value of the customized element.
	//
	// Comment: The expression may be inlined, or may be a reference to a named expression within a logic library referenced by the library element.
	Expression Expression `bson:"expression" json:"expression" xml:"expression"`
}

func (r ActivityDefinition) ResourceType() string {
	return "ActivityDefinition"
}

// MarshalJSON implements json.Marshaler
func (r ActivityDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ActivityDefinition
		ResourceType string `json:"resourceType"`
	}{
		ActivityDefinition: r,
		ResourceType:       r.ResourceType(),
	})
}
