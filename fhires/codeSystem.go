package fhires

import (
	v2 "encoding/json/v2"

	decimal "github.com/shopspring/decimal"
)

// The CodeSystem resource is used to declare the existence of and describe a code system or code system supplement and its key properties, and optionally define a part or all of its content.
//
// Status: active
//
// http://hl7.org/fhir/StructureDefinition/CodeSystem
type CodeSystem struct {
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
	// An absolute URI that is used to identify this code system when it is referenced in a specification, model, design or an instance; also called its canonical identifier.
	//
	// This SHOULD be globally unique and SHOULD be a literal address at which an authoritative instance of this code system is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the code system is stored on different servers. This is used in [Coding](datatypes.html#Coding).system.
	//
	// Comment: Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred.  Multiple instances may share the same URL if they have a distinct version.
	//
	// The determination of when to create a new version of a resource (same url, new version) vs. defining a new artifact is up to the author.  Considerations for making this decision are found in [Technical and Business Versions](resource.html#versions).
	//
	// In some cases, the resource can no longer be found at the stated url, but the url itself cannot change. Implementations can use the [meta.source](resource.html#meta) element to indicate where the current master source of the resource can be found.
	URL *string `bson:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
	// A formal identifier that is used to identify this code system when it is represented in other formats, or referenced in a specification, model, design or an instance.
	//
	// Comment: Typically, this is used for identifiers that can go in an HL7 V3 II (instance identifier) data type, and can then identify this code system outside of FHIR, where it is not possible to use the logical URI.  Note that HL7 defines at least three identifiers for many of its code systems - the FHIR canonical URL, the OID and the V2 Table 0396 mnemonic code.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The identifier that is used to identify this version of the code system when it is referenced in a specification, model, design or instance.
	//
	// This is an arbitrary value managed by the code system author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence. This is used in [Coding](datatypes.html#Coding).version.
	//
	// Comment: There may be different code system instances that have the same identifier but different versions.  The version can be appended to the url in a reference to allow a reference to a particular business version of the code system with the format [url]|[version]. The version SHOULD NOT contain a '#' - see [Business Version](resource.html#bv-format).
	Version *string `bson:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
	// Indicates the mechanism used to compare versions to determine which CodeSystem is more current.
	//
	// Comment: If set as a string, this is a FHIRPath expression that has two additional context variables passed in - %version1 and %version2 and will return a negative number if version1 is newer, a positive number if version2 and a 0 if the version ordering can't be successfully be determined.
	VersionAlgorithmString *string `bson:"versionAlgorithmString,omitempty" json:"versionAlgorithmString,omitempty" xml:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding `bson:"versionAlgorithmCoding,omitempty" json:"versionAlgorithmCoding,omitempty" xml:"versionAlgorithmCoding,omitempty"`
	// A natural language name identifying the code system.
	//
	// This name should be usable as an identifier for the module by machine processing applications such as code generation.
	//
	// Comment: The name is not expected to be globally unique. The name should be a simple alphanumeric type name to ensure that it is machine-processing friendly.
	Name *string `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// A short, descriptive, user-friendly title for the code system.
	//
	// Comment: This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc.
	Title *string `bson:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// The status of this code system. Enables tracking the life-cycle of the content.
	//
	// Comment: Allows filtering of code systems that are appropriate for use versus not.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Status PublicationStatus `bson:"status" json:"status" xml:"status"`
	// A Boolean value to indicate that this code system is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	//
	// Comment: Allows filtering of code systems that are appropriate for use versus not.
	Experimental *bool `bson:"experimental,omitempty" json:"experimental,omitempty" xml:"experimental,omitempty"`
	// The date  (and optionally time) when the code system was last significantly changed.
	//
	// The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the code system changes.
	//
	// Comment: The date is often not tracked until the resource is published, but may be present on draft content. Note that this is not the same as the resource last-modified-date, since the resource may be a secondary representation of the code system. Additional specific dates may be added as extensions or be found by consulting Provenances associated with past versions of the resource.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Date *string `bson:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// The name of the organization or individual responsible for the release and ongoing maintenance of the code system.
	//
	// Comment: Usually an organization but may be an individual. The publisher (or steward) of the code system is the organization or individual primarily responsible for the maintenance and upkeep of the code system. This is not necessarily the same individual or organization that developed and initially authored the content. The publisher is the primary point of contact for questions or issues with the code system. This item SHOULD be populated unless the information is available from context.
	Publisher *string `bson:"publisher,omitempty" json:"publisher,omitempty" xml:"publisher,omitempty"`
	// Contact details to assist a user in finding and communicating with the publisher.
	//
	// Comment: May be a web site, an email address, a telephone number, etc.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty" xml:"contact,omitempty"`
	// A free text natural language description of the code system from a consumer's perspective.
	//
	// Comment: This description can be used to capture details such as comments about misuse, instructions for clinical use and interpretation, literature references, examples from the paper world, etc. It is not a rendering of the code system as conveyed in the 'text' field of the resource itself. This item SHOULD be populated unless the information is available from context (e.g. the language of the code system is presumed to be the predominant language in the place the code system was created).
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// The content was developed with a focus and intent of supporting the contexts that are listed.
	//
	// These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate code system instances.
	//
	// Comment: When multiple useContexts are specified, there is no expectation that all or any of the contexts apply.
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty" xml:"useContext,omitempty"`
	// A legal or geographic region in which the code system is intended to be used.
	//
	// Comment: It may be possible for the code system to be used in jurisdictions other than those for which it was originally designed or intended.
	//
	// DEPRECATION NOTE: For consistency, implementations are encouraged to migrate to using the new 'jurisdiction' code in the useContext element.  (I.e. useContext.code indicating http://terminology.hl7.org/CodeSystem/usage-context-type#jurisdiction and useContext.valueCodeableConcept indicating the jurisdiction.)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty" xml:"jurisdiction,omitempty"`
	// Explanation of why this code system is needed and why it has been designed as it has.
	//
	// Comment: This element does not describe the usage of the code system. Instead, it provides traceability of ''why'' the resource is either needed or ''why'' it is defined as it is.  This may be used to point to source materials or specifications that drove the structure of this code system.
	Purpose *string `bson:"purpose,omitempty" json:"purpose,omitempty" xml:"purpose,omitempty"`
	// A copyright statement relating to the code system and/or its contents.
	//
	// Copyright statements are generally legal restrictions on the use and publishing of the code system.
	//
	// Comment: ... Sometimes, the copyright differs between the code system and the codes that are included. The copyright statement should clearly differentiate between these when required.
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
	// The period during which the CodeSystem content was or is planned to be in active use.
	//
	// Comment: The effective period for a CodeSystem  determines when the content is applicable for usage and is independent of publication and review dates. For example, a system intended to be used for the year 2016 might be published in 2015.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	EffectivePeriod *Period `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty" xml:"effectivePeriod,omitempty"`
	// Descriptions related to the content of the CodeSystem.
	//
	// Topics provide a high-level categorization as well as keywords for the CodeSystem that can be useful for filtering and searching.
	//
	// Comment: DEPRECATION NOTE: For consistency, implementations are encouraged to migrate to using the new 'topic' code in the useContext element.  (I.e. useContext.code indicating http://terminology.hl7.org/CodeSystem/usage-context-type#topic and useContext.valueCodeableConcept indicating the topic)
	Topic []CodeableConcept `bson:"topic,omitempty" json:"topic,omitempty" xml:"topic,omitempty"`
	// An individiual or organization primarily involved in the creation and maintenance of the CodeSystem.
	Author []ContactDetail `bson:"author,omitempty" json:"author,omitempty" xml:"author,omitempty"`
	// An individual or organization primarily responsible for internal coherence of the CodeSystem.
	Editor []ContactDetail `bson:"editor,omitempty" json:"editor,omitempty" xml:"editor,omitempty"`
	// An individual or organization asserted by the publisher to be primarily responsible for review of some aspect of the CodeSystem.
	//
	// Comment: See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Reviewer []ContactDetail `bson:"reviewer,omitempty" json:"reviewer,omitempty" xml:"reviewer,omitempty"`
	// An individual or organization asserted by the publisher to be responsible for officially endorsing the CodeSystem for use in some setting.
	//
	// Comment: See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Endorser []ContactDetail `bson:"endorser,omitempty" json:"endorser,omitempty" xml:"endorser,omitempty"`
	// Related artifacts such as additional documentation, justification, dependencies, bibliographic references, and predecessor and successor artifacts.
	//
	// Comment: Each related artifact is either an attachment, or a reference to another resource, but not both.
	RelatedArtifact []RelatedArtifact `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty" xml:"relatedArtifact,omitempty"`
	// If code comparison is case sensitive when codes within this system are compared to each other.
	//
	// Comment: If this value is missing, then it is not specified whether a code system is case sensitive or not. When the rule is not known, Postel's law should be followed: produce codes with the correct case, and accept codes in any case. This element is primarily provided to support validation software.
	CaseSensitive *bool `bson:"caseSensitive,omitempty" json:"caseSensitive,omitempty" xml:"caseSensitive,omitempty"`
	// Canonical reference to the value set that contains all codes in the code system independent of code status.
	//
	// Comment: The definition of the value set SHALL include all codes from this code system, independent of concept status, and only codes from this code system. It SHALL be immutable, and operations SHOULD support validation of any allowed code compositions.
	ValueSet *string `bson:"valueSet,omitempty" json:"valueSet,omitempty" xml:"valueSet,omitempty"`
	// The meaning of the hierarchy of concepts as represented in this resource.
	//
	// Comment: Note that other representations might have a different hierarchy or none at all, and represent the information using properties.
	HierarchyMeaning *CodeSystemHierarchyMeaning `bson:"hierarchyMeaning,omitempty" json:"hierarchyMeaning,omitempty" xml:"hierarchyMeaning,omitempty"`
	// The code system defines a compositional (post-coordination) grammar.
	//
	// Comment: Note that the code system resource does not define what the compositional grammar is, only whether or not there is one.
	Compositional *bool `bson:"compositional,omitempty" json:"compositional,omitempty" xml:"compositional,omitempty"`
	// This flag is used to signify that the code system does not commit to concept permanence across versions.
	//
	// If true, a version must be specified when referencing this code system.
	//
	// Comment: Best practice is that code systems do not redefine concepts, or that if concepts are redefined, a new code system definition is created. But this is not always possible, so some code systems may be defined as 'versionNeeded'.
	//
	// Most code systems occasionally refine the displays defined for concepts between versions. Contexts in which the concept display values are validated may require that the version be specified for some code systems irrespective of the value of this property.
	VersionNeeded *bool `bson:"versionNeeded,omitempty" json:"versionNeeded,omitempty" xml:"versionNeeded,omitempty"`
	// The extent of the content of the code system (the concepts and codes it defines) are represented in this resource instance.
	Content CodeSystemContentMode `bson:"content" json:"content" xml:"content"`
	// The canonical URL of the code system that this code system supplement is adding designations and properties to.
	//
	// Comment: The most common use of a code system supplement is to add additional language support.
	Supplements *string `bson:"supplements,omitempty" json:"supplements,omitempty" xml:"supplements,omitempty"`
	// The total number of concepts defined by the code system.
	//
	// Where the code system has a compositional grammar, the basis of this count is defined by the system steward.
	//
	// Comment: The count of concepts defined in this resource cannot be more than this value but may be less for several reasons - see the content element.
	Count *int `bson:"count,omitempty" json:"count,omitempty" xml:"count,omitempty"`
	// A filter that can be used in a value set compose statement when selecting concepts using a filter.
	//
	// Comment: Note that filters defined in code systems usually require custom code on the part of any terminology engine that will make them available for use in value set filters. For this reason, they are generally only seen in high value published terminologies.
	Filter []CodeSystemFilter `bson:"filter,omitempty" json:"filter,omitempty" xml:"filter,omitempty"`
	// A property defines an additional slot through which additional information can be provided about a concept.
	//
	// Comment: Multiple occurrences of CodeSystem.concept.property may occur for a CodeSystem.property where     CodeSystem.concept.property.code is the same and CodeSystem.concept.property.value differs. For example: multiple designations for a single concept.
	Property []CodeSystemProperty `bson:"property,omitempty" json:"property,omitempty" xml:"property,omitempty"`
	// Concepts that are in the code system.
	//
	// The concept definitions are inherently hierarchical, but the definitions must be consulted to determine what the meanings of the hierarchical relationships are.
	//
	// Comment: If this is empty, it means that the code system resource does not represent the content of the code system.
	Concept []CodeSystemConcept `bson:"concept,omitempty" json:"concept,omitempty" xml:"concept,omitempty"`
}

type CodeSystemFilter struct {
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
	// The code that identifies this filter when it is used as a filter in [ValueSet](valueset.html#).compose.include.filter.
	Code string `bson:"code" json:"code" xml:"code"`
	// A description of how or why the filter is used.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// A list of operators that can be used with the filter.
	Operator []FilterOperator `bson:"operator" json:"operator" xml:"operator"`
	// A description of what the value for the filter should be.
	Value string `bson:"value" json:"value" xml:"value"`
}

type CodeSystemProperty struct {
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
	// A code that is used to identify the property.
	//
	// The code is used internally (in CodeSystem.concept.property.code) and also externally, such as in property filters.
	Code string `bson:"code" json:"code" xml:"code"`
	// Reference to the formal meaning of the property.
	//
	// One possible source of meaning is the [Concept Properties](codesystem-concept-properties.html) code system.
	Uri *string `bson:"uri,omitempty" json:"uri,omitempty" xml:"uri,omitempty"`
	// A description of the property- why it is defined, and how its value might be used.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// The type of the property value.
	//
	// Properties of type "code" contain a code defined by the code system (e.g. a reference to another defined concept).
	Type PropertyType `bson:"type" json:"type" xml:"type"`
}

type CodeSystemConcept struct {
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
	// A code - a text symbol - that uniquely identifies the concept within the code system.
	Code string `bson:"code" json:"code" xml:"code"`
	// A human readable string that is the recommended default way to present this concept to a user.
	Display *string `bson:"display,omitempty" json:"display,omitempty" xml:"display,omitempty"`
	// The formal definition of the concept.
	//
	// The code system resource does not make formal definitions required, because of the prevalence of legacy systems. However, they are highly recommended, as without them there is no formal meaning associated with the concept.
	Definition *string `bson:"definition,omitempty" json:"definition,omitempty" xml:"definition,omitempty"`
	// Additional representations for the concept - other languages, aliases, specialized purposes, used for particular purposes, etc.
	//
	// Comment: Concepts have both a ```display``` and an array of ```designation```. The display is equivalent to a special designation with an implied ```designation.use``` of "primary code" and a language equal to the [Resource Language](resource.html#language).
	Designation []CodeSystemConceptDesignation `bson:"designation,omitempty" json:"designation,omitempty" xml:"designation,omitempty"`
	// A property value for this concept.
	Property []CodeSystemConceptProperty `bson:"property,omitempty" json:"property,omitempty" xml:"property,omitempty"`
	// Defines children of a concept to produce a hierarchy of concepts.
	//
	// The nature of the relationships is variable (is-a/contains/categorizes) - see hierarchyMeaning.
	Concept []CodeSystemConcept `bson:"concept,omitempty" json:"concept,omitempty"`
}

type CodeSystemConceptDesignation struct {
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
	// The language this designation is defined for.
	//
	// Comment: In the absence of a language, the resource language applies.
	Language *string `bson:"language,omitempty" json:"language,omitempty" xml:"language,omitempty"`
	// A code that details how this designation would be used.
	//
	// Comment: If no use is provided, the designation can be assumed to be suitable for general display to a human user.
	Use *Coding `bson:"use,omitempty" json:"use,omitempty" xml:"use,omitempty"`
	// Additional codes that detail how this designation would be used, if there is more than one use.
	//
	// Comment: This was added rather than increasing the cardinality of .use to 0..* in order to maintain backward compatibility.
	AdditionalUse []Coding `bson:"additionalUse,omitempty" json:"additionalUse,omitempty" xml:"additionalUse,omitempty"`
	// The text value for this designation.
	Value string `bson:"value" json:"value" xml:"value"`
}

type CodeSystemConceptProperty struct {
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
	// A code that is a reference to CodeSystem.property.code.
	Code string `bson:"code" json:"code" xml:"code"`
	// The value of this property.
	ValueCode     string          `bson:"valueCode" json:"valueCode" xml:"valueCode"`
	ValueCoding   Coding          `bson:"valueCoding" json:"valueCoding" xml:"valueCoding"`
	ValueString   string          `bson:"valueString" json:"valueString" xml:"valueString"`
	ValueInteger  int             `bson:"valueInteger" json:"valueInteger" xml:"valueInteger"`
	ValueBoolean  bool            `bson:"valueBoolean" json:"valueBoolean" xml:"valueBoolean"`
	ValueDateTime string          `bson:"valueDateTime" json:"valueDateTime" xml:"valueDateTime"`
	ValueDecimal  decimal.Decimal `bson:"valueDecimal" json:"valueDecimal" xml:"valueDecimal"`
}

func (r CodeSystem) ResourceType() string {
	return "CodeSystem"
}

// MarshalJSON implements json.Marshaler
func (r CodeSystem) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		CodeSystem
		ResourceType string `json:"resourceType"`
	}{
		CodeSystem:   r,
		ResourceType: r.ResourceType(),
	})
}
