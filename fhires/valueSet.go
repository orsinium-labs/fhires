package fhires

import (
	v2 "encoding/json/v2"

	decimal "github.com/shopspring/decimal"
)

// A ValueSet resource instance specifies a set of codes drawn from one or more code systems, intended for use in a particular context.
//
// Value sets link between [[[CodeSystem]]] definitions and their use in [coded elements](terminologies.html).
//
// Status: active
//
// http://hl7.org/fhir/StructureDefinition/ValueSet
type ValueSet struct {
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
	// An absolute URI that is used to identify this value set when it is referenced in a specification, model, design or an instance; also called its canonical identifier.
	//
	// This SHOULD be globally unique and SHOULD be a literal address at which an authoritative instance of this value set is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the value set is stored on different servers.
	//
	// Comment: Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred.  Multiple instances may share the same URL if they have a distinct version.
	//
	// The determination of when to create a new version of a resource (same url, new version) vs. defining a new artifact is up to the author.  Considerations for making this decision are found in [Technical and Business Versions](resource.html#versions).
	//
	// In some cases, the resource can no longer be found at the stated url, but the url itself cannot change. Implementations can use the [meta.source](resource.html#meta) element to indicate where the current master source of the resource can be found.
	URL *string `bson:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
	// A formal identifier that is used to identify this value set when it is represented in other formats, or referenced in a specification, model, design or an instance.
	//
	// Comment: Typically, this is used for identifiers that can go in an HL7 V3 II (instance identifier) data type, and can then identify this value set outside of FHIR, where it is not possible to use the logical URI.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The identifier that is used to identify this version of the value set when it is referenced in a specification, model, design or instance.
	//
	// This is an arbitrary value managed by the value set author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	//
	// Comment: There may be different value set instances that have the same identifier but different versions.  The version can be appended to the url in a reference to allow a reference to a particular business version of the value set with the format [url]|[version]. The version SHOULD NOT contain a '#' - see [Business Version](resource.html#bv-format).
	Version *string `bson:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
	// Indicates the mechanism used to compare versions to determine which ValueSet is more current.
	//
	// Comment: If set as a string, this is a FHIRPath expression that has two additional context variables passed in - %version1 and %version2 and will return a negative number if version1 is newer, a positive number if version2 and a 0 if the version ordering can't be successfully be determined.
	VersionAlgorithmString *string `bson:"versionAlgorithmString,omitempty" json:"versionAlgorithmString,omitempty" xml:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding `bson:"versionAlgorithmCoding,omitempty" json:"versionAlgorithmCoding,omitempty" xml:"versionAlgorithmCoding,omitempty"`
	// A natural language name identifying the value set.
	//
	// This name should be usable as an identifier for the module by machine processing applications such as code generation.
	//
	// Comment: The name is not expected to be globally unique. The name should be a simple alphanumeric type name to ensure that it is machine-processing friendly.A name should be provided unless the value set is a contained resource (e.g. an anonymous value set in a profile). Most registries will require a name.
	Name *string `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// A short, descriptive, user-friendly title for the value set.
	//
	// Comment: This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc.
	Title *string `bson:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// The status of this value set.
	//
	// Enables tracking the life-cycle of the content. The status of the value set applies to the value set definition (ValueSet.compose) and the associated ValueSet metadata. Expansions do not have a state.
	//
	// Comment: Allows filtering of value sets that are appropriate for use versus not.See also the [http://hl7.org/fhir/StructureDefinition/valueset-workflowStatusDescription](http://hl7.org/fhir/extensions/StructureDefinition-valueset-workflowStatusDescription.html) extension for additional status information related to the editorial process.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Status PublicationStatus `bson:"status" json:"status" xml:"status"`
	// A Boolean value to indicate that this value set is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	//
	// Comment: Allows filtering of value sets that are appropriate for use versus not.
	Experimental *bool `bson:"experimental,omitempty" json:"experimental,omitempty" xml:"experimental,omitempty"`
	// The date (and optionally time) when the value set metadata or content logical definition (.compose) was created or revised.
	//
	// Comment: Note that this is not the same as the meta.lastUpdated which is specific to an instance of a value set resource on a server.  Additional specific dates may be added as extensions or be found by consulting Provenances associated with past versions of the resource.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Date *string `bson:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// The name of the organization or individual responsible for the release and ongoing maintenance of the value set.
	//
	// Comment: Usually an organization but may be an individual. This is not necessarily the same individual or organization that developed and initially authored the content. The publisher is the primary point of contact for questions or issues with the value set. This item SHOULD be populated unless the information is available from context.
	Publisher *string `bson:"publisher,omitempty" json:"publisher,omitempty" xml:"publisher,omitempty"`
	// Contact details to assist a user in finding and communicating with the publisher.
	//
	// Comment: May be a web site, an email address, a telephone number, etc.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty" xml:"contact,omitempty"`
	// A free text natural language description of the value set from a consumer's perspective.
	//
	// The textual description specifies the span of meanings for concepts to be included within the Value Set Expansion, and also may specify the intended use and limitations of the Value Set.
	//
	// Comment: Description SHOULD contain instructions for clinical or administrative use and interpretation and information about misuse. Description SHOULD be populated unless the information is available from context. Description SHOULD be populated unless the value set is a contained resource (e.g. a value set defined solely in the context of a profile).
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// The content was developed with a focus and intent of supporting the contexts that are listed.
	//
	// These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate value set instances.
	//
	// Comment: When multiple useContexts are specified, there is no expectation that all or any of the contexts apply.
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty" xml:"useContext,omitempty"`
	// A legal or geographic region in which the value set is intended to be used.
	//
	// Comment: It may be possible for the value set to be used in jurisdictions other than those for which it was originally designed or intended.
	//
	// DEPRECATION NOTE: For consistency, implementations are encouraged to migrate to using the new 'jurisdiction' code in the useContext element.  (I.e. useContext.code indicating http://terminology.hl7.org/CodeSystem/usage-context-type#jurisdiction and useContext.valueCodeableConcept indicating the jurisdiction.)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty" xml:"jurisdiction,omitempty"`
	// If this is set to 'true', then no new versions of the content logical definition can be created.
	//
	//  Note: Other metadata might still change.
	//
	// Comment: Normally immutability is set to 'false', which is the default assumption if it is not populated.  Note that the implication is that if this is set to 'true', there may be only one ValueSet version for this definition. Immutability tends to be set to 'true' in one of two cases: - Where the value set, by the nature of its usage, cannot change.  For example "All specializations of ACT in ActClassCode" - Where there's no safe way to express the "Purpose" such that someone else could safely make changes to the value set definition. Source workflow control must guarantee that the same URI always yields the same definition.
	Immutable *bool `bson:"immutable,omitempty" json:"immutable,omitempty" xml:"immutable,omitempty"`
	// Explanation of why this value set is needed and why it has been designed as it has.
	//
	// Comment: This element does not describe the usage of the value set. Instead, it provides traceability of ''why'' the resource is either needed or ''why'' it is defined as it is.  This may be used to point to source materials or specifications that drove the structure of this value set.
	Purpose *string `bson:"purpose,omitempty" json:"purpose,omitempty" xml:"purpose,omitempty"`
	// A copyright statement relating to the value set and/or its contents.
	//
	// Copyright statements are generally legal restrictions on the use and publishing of the value set.
	//
	// Comment: Frequently, the copyright differs between the value set and the codes that are included. The copyright statement should clearly differentiate between these when required.
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
	// The period during which the ValueSet content was or is planned to be in active use.
	//
	// Comment: The effective period for a ValueSet determines when the content is applicable for usage and is independent of publication and review dates. For example, a valueset intended to be used for the year 2016 might be published in 2015.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	EffectivePeriod *Period `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty" xml:"effectivePeriod,omitempty"`
	// Descriptions related to the content of the ValueSet.
	//
	// Topics provide a high-level categorization as well as keywords for the ValueSet that can be useful for filtering and searching.
	//
	// Comment: DEPRECATION NOTE: For consistency, implementations are encouraged to migrate to using the new 'topic' code in the useContext element.  (I.e. useContext.code indicating http://terminology.hl7.org/CodeSystem/usage-context-type#topic and useContext.valueCodeableConcept indicating the topic)
	Topic []CodeableConcept `bson:"topic,omitempty" json:"topic,omitempty" xml:"topic,omitempty"`
	// An individiual or organization primarily involved in the creation and maintenance of the ValueSet.
	Author []ContactDetail `bson:"author,omitempty" json:"author,omitempty" xml:"author,omitempty"`
	// An individual or organization primarily responsible for internal coherence of the ValueSet.
	Editor []ContactDetail `bson:"editor,omitempty" json:"editor,omitempty" xml:"editor,omitempty"`
	// An individual or organization asserted by the publisher to be primarily responsible for review of some aspect of the ValueSet.
	//
	// Comment: See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Reviewer []ContactDetail `bson:"reviewer,omitempty" json:"reviewer,omitempty" xml:"reviewer,omitempty"`
	// An individual or organization asserted by the publisher to be responsible for officially endorsing the ValueSet for use in some setting.
	//
	// Comment: See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Endorser []ContactDetail `bson:"endorser,omitempty" json:"endorser,omitempty" xml:"endorser,omitempty"`
	// Related artifacts such as additional documentation, justification, dependencies, bibliographic references, and predecessor and successor artifacts.
	//
	// Comment: Each related artifact is either an attachment, or a reference to another resource, but not both.
	RelatedArtifact []RelatedArtifact `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty" xml:"relatedArtifact,omitempty"`
	// A set of criteria that define the contents of the value set by including or excluding codes selected from the specified code system(s) that the value set draws from.
	//
	// This is also known as the Content Logical Definition (CLD).
	Compose *ValueSetCompose `bson:"compose,omitempty" json:"compose,omitempty" xml:"compose,omitempty"`
	// A value set can also be "expanded", where the value set is turned into a simple collection of enumerated codes.
	//
	// This element holds the expansion, if it has been performed.
	//
	// Comment: Expansion is performed to produce a collection of codes that are ready to use for data entry or validation. Value set expansions are always considered to be stateless - they are a record of the set of codes in the value set at a point in time under a given set of conditions, and are not subject to ongoing maintenance.
	//
	// Expansion.parameter is  a simplified list of parameters - a subset of the features of the [Parameters](parameters.html) resource.
	Expansion *ValueSetExpansion `bson:"expansion,omitempty" json:"expansion,omitempty" xml:"expansion,omitempty"`
	// Description of the semantic space the Value Set Expansion is intended to cover and should further clarify the text in ValueSet.description.
	Scope *ValueSetScope `bson:"scope,omitempty" json:"scope,omitempty" xml:"scope,omitempty"`
}

type ValueSetCompose struct {
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
	// The Locked Date is  the effective date that is used to determine the version of all referenced Code Systems and Value Set Definitions included in the compose that are not already tied to a specific version.
	//
	// Comment: With a defined lockedDate the value set is considered "Locked". Otherwise, the value set may have different expansions as underlying code systems and/or value sets evolve.  The interpretation of lockedDate is often dependent on the context - e.g. a SNOMED CT derived value set with a lockedDate will have a different expansion in USA than in UK.  If a value set specifies a version for include and exclude statements, and also specifies a locked date, the specified versions need to be available that date, or the value set will not be usable.
	LockedDate *string `bson:"lockedDate,omitempty" json:"lockedDate,omitempty" xml:"lockedDate,omitempty"`
	// Whether inactive codes - codes that are not approved for current use - are in the value set.
	//
	// If inactive = true, inactive codes are to be included in the expansion, if inactive = false, the inactive codes will not be included in the expansion. If absent, the behavior is determined by the implementation, or by the applicable $expand parameters (but generally, inactive codes would be expected to be included).
	//
	// Comment: Note that in the FHIR terminology framework, "deprecated" does not mean inactive, but in some code systems, e.g. LOINC, "deprecated" does mean inactive. Code systems should define what codes are considered to be inactive. If this is not clearly defined (including in the FHIR code system resource), then all codes are assumed to be active.
	//
	// The Value Set Definition specification defines an ActiveOnly element, which is the reverse of this element e.g. (ValueSet.compose.inactive=FALSE) is the same as (VSD.ActiveOnly=TRUE).
	Inactive *bool `bson:"inactive,omitempty" json:"inactive,omitempty" xml:"inactive,omitempty"`
	// Include one or more codes from a code system or other value set(s).
	//
	// Comment: All the conditions in an include must be true. If a system is listed, all the codes from the system are listed. If one or more filters are listed, all of the filters must apply. If one or more value sets are listed, the codes must be in all the value sets. E.g. each include is 'include all the codes that meet all these conditions'.
	Include []ValueSetComposeInclude `bson:"include" json:"include" xml:"include"`
	// Exclude one or more codes from the value set based on code system filters and/or other value sets.
	//
	// Comment: Usually this is used to selectively exclude codes that were included by subsumption in the inclusions. Any display names specified for the codes are ignored.
	Exclude []ValueSetComposeInclude `bson:"exclude,omitempty" json:"exclude,omitempty"`
	// A property to return in the expansion, if the client doesn't ask for any particular properties.
	//
	// May be either a code from the code system definition (convenient) or a the formal URI that refers to the property. The special value '*' means all properties known to the server.
	//
	// Comment: Note that property names can clash, so using a URI is recommended.
	Property []string `bson:"property,omitempty" json:"property,omitempty" xml:"property,omitempty"`
}

type ValueSetComposeInclude struct {
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
	// An absolute URI which is the code system from which the selected codes come from.
	//
	// Comment: If there are no codes or filters, the entire code system is included. Note that the set of codes that are included may contain abstract codes. See ''Coding.system'' for further documentation about the correct value for the system element.
	System *string `bson:"system,omitempty" json:"system,omitempty" xml:"system,omitempty"`
	// The version of the code system that the codes are selected from, or the special version '*' for all versions.
	//
	// Comment: This is used when selecting the descendants of a concept - they may change between versions. If no version is specified, then the exact contents of the value set might not be known until a context of use binds it to a particular version. The special value '*' means all versions; It is at server discretion regarding expansions and which versions must be supported.
	Version *string `bson:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
	// Specifies a concept to be included or excluded.
	//
	// Comment: The list of concepts is considered ordered, though the order might not have any particular significance. Typically, the order of an expansion follows that defined in the compose element.
	Concept []ValueSetComposeIncludeConcept `bson:"concept,omitempty" json:"concept,omitempty" xml:"concept,omitempty"`
	// Select concepts by specifying a matching criterion based on the properties (including relationships) defined by the system, or on filters defined by the system.
	//
	// If multiple filters are specified within the include, they SHALL all be true.
	//
	// Comment: Selecting codes by specifying filters based on properties is only possible where the underlying code system defines appropriate properties. Note that in some cases, the underlying code system defines the logical concepts but not the literal codes for the concepts. In such cases, the literal definitions may be provided by a third party.
	Filter []ValueSetComposeIncludeFilter `bson:"filter,omitempty" json:"filter,omitempty" xml:"filter,omitempty"`
	// Selects the concepts found in this value set (based on its value set definition).
	//
	// This is an absolute URI that is a reference to ValueSet.url.  If multiple value sets are specified this includes the intersection of the contents of all of the referenced value sets.
	//
	// Comment: The value set URI is either a logical reference to a defined value set such as a [SNOMED CT reference set](https://terminology.hl7.org/SNOMEDCT.html), or a direct reference to a value set definition using ValueSet.url. The reference might not refer to an actual FHIR ValueSet resource; in this case, whatever is referred to is an implicit definition of a value set that needs to be clear about how versions are resolved.
	ValueSet []string `bson:"valueSet,omitempty" json:"valueSet,omitempty" xml:"valueSet,omitempty"`
	// A copyright statement for the specific code system asserted by the containing ValueSet.compose.include element's system value (if the associated ValueSet.compose.include.version element is not present); or the code system and version combination (if the associated ValueSet.compose.include.version element is present).
	Copyright *string `bson:"copyright,omitempty" json:"copyright,omitempty" xml:"copyright,omitempty"`
}

type ValueSetComposeIncludeConcept struct {
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
	// Specifies a code for the concept to be included or excluded.
	//
	// Comment: Expressions are allowed if defined by the underlying code system.
	Code string `bson:"code" json:"code" xml:"code"`
	// The text to display to the user for this concept in the context of this valueset.
	//
	// If no display is provided, then applications using the value set use the display specified for the code by the system.
	//
	// Comment: The value set resource allows for an alternative display to be specified for when this concept is used in this particular value set. See notes in the value set narrative about the correct use of this element.
	Display *string `bson:"display,omitempty" json:"display,omitempty" xml:"display,omitempty"`
	// Additional representations for this concept when used in this value set - other languages, aliases, specialized purposes, used for particular purposes, etc.
	//
	// Comment: Concepts have both a ```display``` and an array of ```designation```. The display is equivalent to a special designation with an implied ```designation.use``` of "primary code" and a language equal to the [Resource Language](resource.html#language).
	Designation []ValueSetComposeIncludeConceptDesignation `bson:"designation,omitempty" json:"designation,omitempty" xml:"designation,omitempty"`
}

type ValueSetComposeIncludeConceptDesignation struct {
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
	// A code that represents types of uses of designations.
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

type ValueSetComposeIncludeFilter struct {
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
	// A code that identifies a property or a filter defined in the code system.
	Property string `bson:"property" json:"property" xml:"property"`
	// The kind of operation to perform as a part of the filter criteria.
	//
	// Comment: In case filter.property represents a property of the system, the operation applies to the selected property. In case filter.property represents a filter of the system, the operation SHALL match one of the CodeSystem.filter.operator values.
	Op FilterOperator `bson:"op" json:"op" xml:"op"`
	// The match value may be either a code defined by the system, or a string value, which is a regex match on the literal string of the property value  (if the filter represents a property defined in CodeSystem) or of the system filter value (if the filter represents a filter defined in CodeSystem) when the operation is 'regex', or one of the values (true and false), when the operation is 'exists'.
	//
	// Comment: Use regex matching with care - full regex matching on every SNOMED CT term is prohibitive, for example.
	Value string `bson:"value" json:"value" xml:"value"`
}

type ValueSetExpansion struct {
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
	// An identifier that uniquely identifies this expansion of the valueset, based on a unique combination of the provided parameters, the system default parameters, and the underlying system code system versions etc.
	//
	// Systems may re-use the same identifier as long as those factors remain the same, and the expansion is the same, but are not required to do so. This is a business identifier.
	//
	// Comment: Typically, this uri is a UUID (e.g. urn:uuid:8230ff20-c97a-4167-a59d-dc2cb9df16dd).
	Identifier *string `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// As per paging Search results, the next URLs are opaque to the client, have no dictated structure, and only the server understands them.
	//
	// Comment: Clients SHOULD use the next link, if present, to page through expansion results in preference to using the offset and count parameters. Due to the optional nature of the next link, its absence does not necessarily indicate that it is the last page of results. Instead, as the offset and count parameters SHALL be populated when paging, clients can reliably use the count/offset parameters to determine whether the whole expansion is returned.
	Next *string `bson:"next,omitempty" json:"next,omitempty" xml:"next,omitempty"`
	// The time at which the expansion was produced by the expanding system.
	//
	// Comment: This SHOULD be a fully populated instant, but in some circumstances, value sets are expanded by hand, and the expansion is published without that precision.
	Timestamp string `bson:"timestamp" json:"timestamp" xml:"timestamp"`
	// The total number of concepts in the expansion.
	//
	// If the number of concept nodes in this resource is less than the stated number, then the server can return more using the offset parameter.
	//
	// Comment: Paging only applies to flat expansions.
	Total *int `bson:"total,omitempty" json:"total,omitempty" xml:"total,omitempty"`
	// If paging is being used, the offset at which this resource starts.
	//
	//  I.e. this resource is a partial view into the expansion. If paging is not being used, this element SHALL NOT be present.
	//
	// Comment: Paging only applies to flat expansions. If a filter is applied, the count is the number of concepts that matched the filter, not the number of concepts in an unfiltered view of the expansion.
	Offset *int `bson:"offset,omitempty" json:"offset,omitempty" xml:"offset,omitempty"`
	// A parameter that controlled the expansion process.
	//
	// These parameters may be used by users of expanded value sets to check whether the expansion is suitable for a particular purpose, or to pick the correct expansion.
	//
	// Comment: The server decides which parameters to include here, but at a minimum, the list SHOULD include all of the parameters that affect the $expand operation. If the expansion will be persisted all of these parameters SHALL be included. If the codeSystem on the server has a specified version then this version SHALL be provided as a parameter in the expansion (note that not all code systems have a version).
	Parameter []ValueSetExpansionParameter `bson:"parameter,omitempty" json:"parameter,omitempty" xml:"parameter,omitempty"`
	// A property defines an additional slot through which additional information can be provided about a concept.
	Property []ValueSetExpansionProperty `bson:"property,omitempty" json:"property,omitempty" xml:"property,omitempty"`
	// The codes that are contained in the value set expansion.
	Contains []ValueSetExpansionContains `bson:"contains,omitempty" json:"contains,omitempty" xml:"contains,omitempty"`
}

type ValueSetExpansionParameter struct {
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
	// Name of the input parameter to the $expand operation; may be a server-assigned name for additional default or other server-supplied parameters used to control the expansion process.
	//
	// Comment: The names are assigned at the discretion of the server.
	Name string `bson:"name" json:"name" xml:"name"`
	// The value of the parameter.
	ValueString   *string          `bson:"valueString,omitempty" json:"valueString,omitempty" xml:"valueString,omitempty"`
	ValueBoolean  *bool            `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty" xml:"valueBoolean,omitempty"`
	ValueInteger  *int             `bson:"valueInteger,omitempty" json:"valueInteger,omitempty" xml:"valueInteger,omitempty"`
	ValueDecimal  *decimal.Decimal `bson:"valueDecimal,omitempty" json:"valueDecimal,omitempty" xml:"valueDecimal,omitempty"`
	ValueUri      *string          `bson:"valueUri,omitempty" json:"valueUri,omitempty" xml:"valueUri,omitempty"`
	ValueCode     *string          `bson:"valueCode,omitempty" json:"valueCode,omitempty" xml:"valueCode,omitempty"`
	ValueDateTime *string          `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty" xml:"valueDateTime,omitempty"`
}

type ValueSetExpansionProperty struct {
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
	// The code is used in ValueSet.expansion.contains.property.code.
	Code string `bson:"code" json:"code" xml:"code"`
	// Reference to the formal meaning of the property.
	//
	// One possible source of meaning is the [Concept Properties](codesystem-concept-properties.html) code system.
	Uri *string `bson:"uri,omitempty" json:"uri,omitempty" xml:"uri,omitempty"`
}

type ValueSetExpansionContains struct {
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
	// An absolute URI which is the code system in which the code for this item in the expansion is defined.
	System *string `bson:"system,omitempty" json:"system,omitempty" xml:"system,omitempty"`
	// If true, this entry is included in the expansion for navigational purposes, and the user cannot select the code directly as a proper value.
	//
	// Comment: This should not be understood to exclude its use for searching (e.g. by subsumption testing). The client should know whether it is appropriate for the user to select an abstract code or not.
	Abstract *bool `bson:"abstract,omitempty" json:"abstract,omitempty" xml:"abstract,omitempty"`
	// If the concept is inactive in the code system that defines it.
	//
	// Inactive codes are those that are no longer to be used, but are maintained by the code system for understanding legacy data. It might not be known or specified whether a concept is inactive (and it may depend on the context of use).
	//
	// Comment: This should only have a value if the concept is inactive.
	Inactive *bool `bson:"inactive,omitempty" json:"inactive,omitempty" xml:"inactive,omitempty"`
	// The version of the code system from this code was taken.
	//
	// Note that a well-maintained code system does not need the version reported, because the meaning of codes is consistent across versions. However this cannot consistently be assured, and when the meaning is not guaranteed to be consistent, the version SHOULD be exchanged.
	//
	// Comment: The exact value of the version string is specified by the system from which the code is derived.
	Version *string `bson:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
	// The code for this item in the expansion hierarchy.
	//
	// If this code is missing the entry in the hierarchy is a place holder (abstract) and does not represent a valid code in the value set.
	Code *string `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// The recommended display for this item in the expansion.
	Display *string `bson:"display,omitempty" json:"display,omitempty" xml:"display,omitempty"`
	// Additional representations for this item - other languages, aliases, specialized purposes, used for particular purposes, etc.
	//
	// These are relevant when the conditions of the expansion do not fix to a single correct representation.
	//
	// Comment: The designations provided must be based on the value set and code system definitions.
	Designation []ValueSetComposeIncludeConceptDesignation `bson:"designation,omitempty" json:"designation,omitempty"`
	// A property value for this concept.
	Property []ValueSetExpansionContainsProperty `bson:"property,omitempty" json:"property,omitempty" xml:"property,omitempty"`
	// Other codes and entries contained under this entry in the hierarchy.
	//
	// Comment: If the expansion uses this element, there is  no implication about the logical relationship between them, and the  structure cannot be used for logical inferencing. The structure  exists to provide navigational assistance for helping human users to  locate codes in the expansion.
	Contains []ValueSetExpansionContains `bson:"contains,omitempty" json:"contains,omitempty"`
}

type ValueSetExpansionContainsProperty struct {
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
	// A code that is a reference to ValueSet.expansion.property.code.
	Code string `bson:"code" json:"code" xml:"code"`
	// The value of this property.
	ValueCode     string          `bson:"valueCode" json:"valueCode" xml:"valueCode"`
	ValueCoding   Coding          `bson:"valueCoding" json:"valueCoding" xml:"valueCoding"`
	ValueString   string          `bson:"valueString" json:"valueString" xml:"valueString"`
	ValueInteger  int             `bson:"valueInteger" json:"valueInteger" xml:"valueInteger"`
	ValueBoolean  bool            `bson:"valueBoolean" json:"valueBoolean" xml:"valueBoolean"`
	ValueDateTime string          `bson:"valueDateTime" json:"valueDateTime" xml:"valueDateTime"`
	ValueDecimal  decimal.Decimal `bson:"valueDecimal" json:"valueDecimal" xml:"valueDecimal"`
	// A subproperty value for this concept.
	SubProperty []ValueSetExpansionContainsPropertySubProperty `bson:"subProperty,omitempty" json:"subProperty,omitempty" xml:"subProperty,omitempty"`
}

type ValueSetExpansionContainsPropertySubProperty struct {
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
	// A code that is a reference to ValueSet.expansion.property.code.
	Code string `bson:"code" json:"code" xml:"code"`
	// The value of this subproperty.
	ValueCode     string          `bson:"valueCode" json:"valueCode" xml:"valueCode"`
	ValueCoding   Coding          `bson:"valueCoding" json:"valueCoding" xml:"valueCoding"`
	ValueString   string          `bson:"valueString" json:"valueString" xml:"valueString"`
	ValueInteger  int             `bson:"valueInteger" json:"valueInteger" xml:"valueInteger"`
	ValueBoolean  bool            `bson:"valueBoolean" json:"valueBoolean" xml:"valueBoolean"`
	ValueDateTime string          `bson:"valueDateTime" json:"valueDateTime" xml:"valueDateTime"`
	ValueDecimal  decimal.Decimal `bson:"valueDecimal" json:"valueDecimal" xml:"valueDecimal"`
}

type ValueSetScope struct {
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
	// Criteria describing which concepts or codes should be included and why.
	InclusionCriteria *string `bson:"inclusionCriteria,omitempty" json:"inclusionCriteria,omitempty" xml:"inclusionCriteria,omitempty"`
	// Criteria describing which concepts or codes should be excluded and why.
	ExclusionCriteria *string `bson:"exclusionCriteria,omitempty" json:"exclusionCriteria,omitempty" xml:"exclusionCriteria,omitempty"`
}

func (r ValueSet) ResourceType() string {
	return "ValueSet"
}

// MarshalJSON implements json.Marshaler
func (r ValueSet) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		ValueSet
		ResourceType string `json:"resourceType"`
	}{
		ResourceType: r.ResourceType(),
		ValueSet:     r,
	})
}
