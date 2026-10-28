package fhires

import "encoding/json/v2"

// A Capability Statement documents a set of capabilities (behaviors) of a FHIR Server or Client for a particular version of FHIR that may be used as a statement of actual server functionality or a statement of required or desired server implementation.
//
// Status: active
//
// http://hl7.org/fhir/StructureDefinition/CapabilityStatement
type CapabilityStatement struct {
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
	// An absolute URI that is used to identify this capability statement when it is referenced in a specification, model, design or an instance; also called its canonical identifier.
	//
	// This SHOULD be globally unique and SHOULD be a literal address at which an authoritative instance of this capability statement is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the capability statement is stored on different servers.
	//
	// Comment: Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred.  Multiple instances may share the same URL if they have a distinct version.
	//
	// The determination of when to create a new version of a resource (same url, new version) vs. defining a new artifact is up to the author.  Considerations for making this decision are found in [Technical and Business Versions](resource.html#versions).
	//
	// In some cases, the resource can no longer be found at the stated url, but the url itself cannot change. Implementations can use the [meta.source](resource.html#meta) element to indicate where the current master source of the resource can be found.
	URL *string `bson:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
	// A formal identifier that is used to identify this CapabilityStatement when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The identifier that is used to identify this version of the capability statement when it is referenced in a specification, model, design or instance.
	//
	// This is an arbitrary value managed by the capability statement author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	//
	// Comment: There may be different capability statement instances that have the same identifier but different versions.  The version can be appended to the url in a reference to allow a reference to a particular business version of the capability statement with the format [url]|[version]. The version SHOULD NOT contain a '#' - see [Business Version](resource.html#bv-format).
	Version *string `bson:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
	// Indicates the mechanism used to compare versions to determine which is more current.
	//
	// Comment: If set as a string, this is a FHIRPath expression that has two additional context variables passed in - %version1 and %version2 and will return a negative number if version1 is newer, a positive number if version2 and a 0 if the version ordering can't be successfully be determined.
	VersionAlgorithmString *string `bson:"versionAlgorithmString,omitempty" json:"versionAlgorithmString,omitempty" xml:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding `bson:"versionAlgorithmCoding,omitempty" json:"versionAlgorithmCoding,omitempty" xml:"versionAlgorithmCoding,omitempty"`
	// A natural language name identifying the capability statement.
	//
	// This name should be usable as an identifier for the module by machine processing applications such as code generation.
	//
	// Comment: The name is not expected to be globally unique. The name should be a simple alphanumeric type name to ensure that it is machine-processing friendly.
	Name *string `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// A short, descriptive, user-friendly title for the capability statement.
	//
	// Comment: This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc.
	Title *string `bson:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// The status of this capability statement. Enables tracking the life-cycle of the content.
	//
	// Comment: Allows filtering of capability statements that are appropriate for use versus not.This is not intended for use with actual capability statements, but where capability statements are used to describe possible or desired systems.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Status PublicationStatus `bson:"status" json:"status" xml:"status"`
	// A Boolean value to indicate that this capability statement is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	//
	// Comment: Allows filtering of capability statements that are appropriate for use versus not.
	Experimental *bool `bson:"experimental,omitempty" json:"experimental,omitempty" xml:"experimental,omitempty"`
	// The date  (and optionally time) when the capability statement was last significantly changed.
	//
	// The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the capability statement changes.
	//
	// Comment: The date is often not tracked until the resource is published, but may be present on draft content. Note that this is not the same as the resource last-modified-date, since the resource may be a secondary representation of the capability statement. Additional specific dates may be added as extensions or be found by consulting Provenances associated with past versions of the resource.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Date string `bson:"date" json:"date" xml:"date"`
	// The name of the organization or individual responsible for the release and ongoing maintenance of the capability statement.
	//
	// Comment: Usually an organization but may be an individual. The publisher (or steward) of the capability statement is the organization or individual primarily responsible for the maintenance and upkeep of the capability statement. This is not necessarily the same individual or organization that developed and initially authored the content. The publisher is the primary point of contact for questions or issues with the capability statement. This item SHOULD be populated unless the information is available from context.
	Publisher *string `bson:"publisher,omitempty" json:"publisher,omitempty" xml:"publisher,omitempty"`
	// Contact details to assist a user in finding and communicating with the publisher.
	//
	// Comment: May be a web site, an email address, a telephone number, etc.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty" xml:"contact,omitempty"`
	// A free text natural language description of the capability statement from a consumer's perspective.
	//
	// Typically, this is used when the capability statement describes a desired rather than an actual solution, for example as a formal expression of requirements as part of an RFP.
	//
	// Comment: This description can be used to capture details such as comments about misuse, instructions for clinical use and interpretation, literature references, examples from the paper world, etc. It is not a rendering of the capability statement as conveyed in the 'text' field of the resource itself. This item SHOULD be populated unless the information is available from context (e.g. the language of the capability statement is presumed to be the predominant language in the place the capability statement was created).This does not need to be populated if the description is adequately implied by the software or implementation details.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// The content was developed with a focus and intent of supporting the contexts that are listed.
	//
	// These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate capability statement instances.
	//
	// Comment: When multiple useContexts are specified, there is no expectation that all or any of the contexts apply.
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty" xml:"useContext,omitempty"`
	// A legal or geographic region in which the capability statement is intended to be used.
	//
	// Comment: It may be possible for the capability statement to be used in jurisdictions other than those for which it was originally designed or intended.
	//
	// DEPRECATION NOTE: For consistency, implementations are encouraged to migrate to using the new 'jurisdiction' code in the useContext element.  (I.e. useContext.code indicating http://terminology.hl7.org/CodeSystem/usage-context-type#jurisdiction and useContext.valueCodeableConcept indicating the jurisdiction.)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty" xml:"jurisdiction,omitempty"`
	// Explanation of why this capability statement is needed and why it has been designed as it has.
	//
	// Comment: This element does not describe the usage of the capability statement. Instead, it provides traceability of ''why'' the resource is either needed or ''why'' it is defined as it is.  This may be used to point to source materials or specifications that drove the structure of this capability statement.
	Purpose *string `bson:"purpose,omitempty" json:"purpose,omitempty" xml:"purpose,omitempty"`
	// A copyright statement relating to the capability statement and/or its contents.
	//
	// Copyright statements are generally legal restrictions on the use and publishing of the capability statement.
	//
	// Comment: ...
	Copyright *string `bson:"copyright,omitempty" json:"copyright,omitempty" xml:"copyright,omitempty"`
	// A short string (<50 characters), suitable for inclusion in a page footer that identifies the copyright holder, effective period, and optionally whether rights are resctricted.
	//
	// (e.g. 'All rights reserved', 'Some rights reserved').
	//
	// Comment: The (c) symbol should NOT be included in this string. It will be added by software when rendering the notation. Full details about licensing, restrictions, warrantees, etc. goes in the more general 'copyright' element.
	CopyrightLabel *string `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty" xml:"copyrightLabel,omitempty"`
	// The way that this statement is intended to be used, to describe an actual running instance of software, a particular product (kind, not instance of software) or a class of implementation (e.g.
	//
	// a desired purchase).
	Kind CapabilityStatementKind `bson:"kind" json:"kind" xml:"kind"`
	// Reference to a canonical URL of another CapabilityStatement that this software implements.
	//
	// This capability statement is a published API description that corresponds to a business service. The server may actually implement a subset of the capability statement it claims to implement, so the capability statement must specify the full capability details.
	//
	// Comment: HL7 defines the following Services: [Terminology Service](terminology-service.html).
	//
	// Many [Implementation Guides](http://fhir.org/guides/registry) define additional services.
	Instantiates []string `bson:"instantiates,omitempty" json:"instantiates,omitempty" xml:"instantiates,omitempty"`
	// Reference to a canonical URL of another CapabilityStatement that this software adds to.
	//
	// The capability statement automatically includes everything in the other statement, and it is not duplicated, though the server may repeat the same resources, interactions and operations to add additional details to them.
	//
	// Comment: the contents of any directly or indirectly imported CapabilityStatements SHALL NOT overlap, i.e. they cannot refer to the same rest/resource, operations/name, searchparam/name, interaction/code, messaging/endpoint, document/mode pair.
	//
	// A capability statement that imports another CapabilityStatement automatically instantiates it too (though this is often not a very useful statement for the kinds of CapabilityStatements that are suitable for importing).
	Imports []string `bson:"imports,omitempty" json:"imports,omitempty" xml:"imports,omitempty"`
	// Software that is covered by this capability statement.
	//
	//  It is used when the capability statement describes the capabilities of a particular software version, independent of an installation.
	Software *CapabilityStatementSoftware `bson:"software,omitempty" json:"software,omitempty" xml:"software,omitempty"`
	// Identifies a specific implementation instance that is described by the capability statement - i.e.
	//
	// a particular installation, rather than the capabilities of a software program.
	Implementation *CapabilityStatementImplementation `bson:"implementation,omitempty" json:"implementation,omitempty" xml:"implementation,omitempty"`
	// The version of the FHIR specification that this CapabilityStatement describes (which SHALL be the same as the FHIR version of the CapabilityStatement itself).
	//
	// There is no default value.
	//
	// Comment: Servers may implement multiple versions (see [Managing Multiple Versions](versioning.html), and the [$versions](capabilitystatement-operation-versions.html) operation). If they do, and the CapabilityStatement is requested from the server, then this fhirVersion will be either the version requested, or the server's default version.
	FhirVersion FHIRVersion `bson:"fhirVersion" json:"fhirVersion" xml:"fhirVersion"`
	// A list of the formats supported by this implementation using their content types.
	//
	// Comment: "xml", "json" and "ttl" are allowed, which describe the simple encodings described in the specification (and imply appropriate bundle support). Otherwise, mime types are legal here.
	Format []string `bson:"format" json:"format" xml:"format"`
	// A list of the patch formats supported by this implementation using their content types.
	//
	// Comment: At present, the patch mime types application/json-patch+json and application/xml-patch+xml are legal. Generally, if a server supports PATCH, it would be expected to support the patch formats and match the formats it supports, but this is not always possible or necessary.
	PatchFormat []string `bson:"patchFormat,omitempty" json:"patchFormat,omitempty" xml:"patchFormat,omitempty"`
	// A list of the languages supported by this implementation that are usefully supported in the ```Accept-Language``` header.
	//
	// Comment: In general, if a server gets a request with an Accept-Language that it doesn't support, it should still reutrn the resource, just in its default language for the resource.
	AcceptLanguage []string `bson:"acceptLanguage,omitempty" json:"acceptLanguage,omitempty" xml:"acceptLanguage,omitempty"`
	// A list of implementation guides that the server does (or should) support in their entirety.
	//
	// Comment: Note: this is primarily only relevant in terms of ImplementationGuides that don't define specific CapabilityStatements declaring the expectation of distinct roles.  (E.g. generic IGs that establish privacy policies.)  In situations where an ImplementationGuide does define CapabilityStatements, asserting CapabilityStatement.implementationGuide means that the implementation adheres to any Implementation.global definitions present in that IG as well as any textual requirements around security or other general interoperability behaviors. However, it does not make any assertions as to conformance with any of the CapabilityStatements defined in the IG. To assert conformance with CapabilityStatements in a referenced IG, it is necessary to use the CapabilityStatement.instantiates element.
	ImplementationGuide []string `bson:"implementationGuide,omitempty" json:"implementationGuide,omitempty" xml:"implementationGuide,omitempty"`
	// A definition of the restful capabilities of the solution, if any.
	//
	// Comment: Multiple repetitions allow definition of both client and/or server behaviors or possibly behaviors under different configuration settings (for software or requirements statements).
	Rest []CapabilityStatementRest `bson:"rest,omitempty" json:"rest,omitempty" xml:"rest,omitempty"`
	// A description of the messaging capabilities of the solution.
	//
	// Comment: Multiple repetitions allow the documentation of multiple endpoints per solution.
	Messaging []CapabilityStatementMessaging `bson:"messaging,omitempty" json:"messaging,omitempty" xml:"messaging,omitempty"`
	// A document definition.
	Document []CapabilityStatementDocument `bson:"document,omitempty" json:"document,omitempty" xml:"document,omitempty"`
}

type CapabilityStatementSoftware struct {
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
	// Name the software is known by.
	Name string `bson:"name" json:"name" xml:"name"`
	// The version identifier for the software covered by this statement.
	//
	// Comment: If possible, a version should be specified, as statements are likely to be different for different versions of software.
	Version *string `bson:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
	// Date this version of the software was released.
	ReleaseDate *string `bson:"releaseDate,omitempty" json:"releaseDate,omitempty" xml:"releaseDate,omitempty"`
}

type CapabilityStatementImplementation struct {
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
	// Information about the specific installation that this capability statement relates to.
	Description string `bson:"description" json:"description" xml:"description"`
	// An absolute base URL for the implementation.
	//
	//  This forms the base for REST interfaces as well as the mailbox and document interfaces.
	URL *string `bson:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
	// The organization responsible for the management of the instance and oversight of the data on the server at the specified URL.
	Custodian *Reference `bson:"custodian,omitempty" json:"custodian,omitempty" xml:"custodian,omitempty"`
}

type CapabilityStatementRest struct {
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
	// Identifies whether this portion of the statement is describing the ability to initiate or receive restful operations.
	Mode RestfulCapabilityMode `bson:"mode" json:"mode" xml:"mode"`
	// Information about the system's restful capabilities that apply across all applications, such as security.
	Documentation *string `bson:"documentation,omitempty" json:"documentation,omitempty" xml:"documentation,omitempty"`
	// Information about security implementation from an interface perspective - what a client needs to know.
	Security *CapabilityStatementRestSecurity `bson:"security,omitempty" json:"security,omitempty" xml:"security,omitempty"`
	// A specification of the restful capabilities of the solution for a specific resource type.
	//
	// Comment: Max of one repetition per resource type.
	Resource []CapabilityStatementRestResource `bson:"resource,omitempty" json:"resource,omitempty" xml:"resource,omitempty"`
	// A specification of restful operations supported by the system.
	Interaction []CapabilityStatementRestInteraction `bson:"interaction,omitempty" json:"interaction,omitempty" xml:"interaction,omitempty"`
	// Search parameters that are supported for searching all resources for implementations to support and/or make use of - either references to ones defined in the specification, or additional ones defined for/by the implementation.
	//
	// This is only for searches executed against the system-level endpoint.
	//
	// Comment: Typically, the only search parameters supported for all searches are those that apply to all resources - tags, profiles, text search etc. These search parameters should include the control search parameters such as _sort, _count, etc. that also apply to this resource (though many will be listed at [CapabilityStatement.rest.searchParam](capabilitystatement-definitions.html#CapabilityStatement.rest.searchParam)). The behavior of some search parameters may be further described by other code or extension elements, or narrative within the capability statement or linked [SearchParameter](searchparameter.html#) definitions.
	SearchParam []CapabilityStatementRestResourceSearchParam `bson:"searchParam,omitempty" json:"searchParam,omitempty"`
	// Definition of an operation or a named query together with its parameters and their meaning and type.
	//
	// Comment: CapabilityStatement.rest.operation is for operations invoked at the system level, or for operations that are supported across multiple resource types. Operations linked from CapabilityStatement.rest.operation must have OperationDefinition.system = true, or more than one Operation.resource.
	Operation []CapabilityStatementRestResourceOperation `bson:"operation,omitempty" json:"operation,omitempty"`
	// An absolute URI which is a reference to the definition of a compartment that the system supports.
	//
	// The reference is to a CompartmentDefinition resource by its canonical URL .
	//
	// Comment: At present, the only defined compartments are at [CompartmentDefinition](compartmentdefinition.html).
	Compartment []string `bson:"compartment,omitempty" json:"compartment,omitempty" xml:"compartment,omitempty"`
}

type CapabilityStatementRestSecurity struct {
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
	// Server adds CORS headers when responding to requests - this enables Javascript applications to use the server.
	//
	// Comment: The easiest CORS headers to add are Access-Control-Allow-Origin: * & Access-Control-Request-Method: GET, POST, PUT, DELETE. All servers SHOULD support CORS.
	Cors *bool `bson:"cors,omitempty" json:"cors,omitempty" xml:"cors,omitempty"`
	// Types of security services that are supported/required by the system.
	Service []CodeableConcept `bson:"service,omitempty" json:"service,omitempty" xml:"service,omitempty"`
	// General description of how security works.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
}

type CapabilityStatementRestResource struct {
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
	// A type of resource exposed via the restful interface.
	Type ResourceType `bson:"type" json:"type" xml:"type"`
	// A system-wide profile that is applied across *all* instances of the resource supported by the system.
	//
	// For example, if declared on Observation, this profile is the "superset" of capabilities for laboratory *and* vitals *and* other domains. See further discussion in [Using Profiles](profiling.html#profile-uses).
	//
	// Comment: All other profiles for this type that are listed in `.rest.resource.supportedProfile` must conform to this profile.
	Profile *string `bson:"profile,omitempty" json:"profile,omitempty" xml:"profile,omitempty"`
	// A list of profiles representing different use cases the system hosts/produces.
	//
	// A supported profile is a statement about the functionality of the data and services provided by the server (or the client) for supported use cases. For example, a system can define and declare multiple Observation profiles for laboratory observations, vital sign observations, etc. By declaring supported profiles, systems provide a way to determine whether individual resources are conformant. See further discussion in [Using Profiles](profiling.html#profile-uses).
	//
	// Comment: Supported profiles must conform to the resource profile in the `.rest.resource.profile` element if it is present. The resource profile is a system-wide profile applied across *all* instances of the resource supported by the system. A supported profile is a statement about the functionality of the data and services provided by the server (or used by the client) for a particular set of use cases and will not necessarily apply to all data consumed or exposed by the server.
	SupportedProfile []string `bson:"supportedProfile,omitempty" json:"supportedProfile,omitempty" xml:"supportedProfile,omitempty"`
	// Additional information about the resource type used by the system.
	Documentation *string `bson:"documentation,omitempty" json:"documentation,omitempty" xml:"documentation,omitempty"`
	// Identifies a restful operation supported by the solution.
	//
	// Comment: In general, a Resource will only appear in a CapabilityStatement if the server actually has some capabilities - e.g. there is at least one interaction supported. However interactions can be omitted to support summarization (_summary = true).
	Interaction []CapabilityStatementRestResourceInteraction `bson:"interaction,omitempty" json:"interaction,omitempty" xml:"interaction,omitempty"`
	// This field is set to no-version to specify that the system does not support (server) or use (client) versioning for this resource type.
	//
	// If this has some other value, the server must at least correctly track and populate the versionId meta-property on resources. If the value is 'versioned-update', then the server supports all the versioning features, including using e-tags for version integrity in the API.
	//
	// Comment: If a server supports versionIds correctly, it SHOULD support vread too, but is not required to do so.
	Versioning *ResourceVersionPolicy `bson:"versioning,omitempty" json:"versioning,omitempty" xml:"versioning,omitempty"`
	// A flag for whether the server is able to return past versions as part of the vRead operation.
	//
	// Comment: It is useful to support the vRead operation for current operations, even if past versions aren't available.
	ReadHistory *bool `bson:"readHistory,omitempty" json:"readHistory,omitempty" xml:"readHistory,omitempty"`
	// A flag to indicate that the server allows or needs to allow the client to create new identities on the server (that is, the client PUTs to a location where there is no existing resource).
	//
	// Allowing this operation means that the server allows the client to create new identities on the server.
	//
	// Comment: Allowing the clients to create new identities on the server means that the system administrator needs to have confidence that the clients do not create clashing identities between them. Obviously, if there is only one client, this won't happen. While creating identities on the client means that the clients need to be managed, it's much more convenient for many scenarios if such management can be put in place.
	UpdateCreate *bool `bson:"updateCreate,omitempty" json:"updateCreate,omitempty" xml:"updateCreate,omitempty"`
	// A flag that indicates that the server supports conditional create.
	//
	// Comment: Conditional Create is mainly appropriate for interface engine scripts converting from other formats, such as json.
	ConditionalCreate *bool `bson:"conditionalCreate,omitempty" json:"conditionalCreate,omitempty" xml:"conditionalCreate,omitempty"`
	// A code that indicates how the server supports conditional read.
	//
	// Comment: Conditional Read is mainly appropriate for interface engine scripts converting from other formats, such as json.
	ConditionalRead *ConditionalReadStatus `bson:"conditionalRead,omitempty" json:"conditionalRead,omitempty" xml:"conditionalRead,omitempty"`
	// A flag that indicates that the server supports conditional update.
	//
	// Comment: Conditional Update is mainly appropriate for interface engine scripts converting from other formats, such as json.
	ConditionalUpdate *bool `bson:"conditionalUpdate,omitempty" json:"conditionalUpdate,omitempty" xml:"conditionalUpdate,omitempty"`
	// A flag that indicates that the server supports conditional patch.
	//
	// Comment: Conditional Patch is mainly appropriate for interface engine scripts converting from other formats, such as json.
	ConditionalPatch *bool `bson:"conditionalPatch,omitempty" json:"conditionalPatch,omitempty" xml:"conditionalPatch,omitempty"`
	// A code that indicates how the server supports conditional delete.
	//
	// Comment: Conditional Delete is mainly appropriate for interface engine scripts converting from other formats, such as json.
	ConditionalDelete *ConditionalDeleteStatus `bson:"conditionalDelete,omitempty" json:"conditionalDelete,omitempty" xml:"conditionalDelete,omitempty"`
	// A set of flags that defines how references are supported.
	ReferencePolicy []ReferenceHandlingPolicy `bson:"referencePolicy,omitempty" json:"referencePolicy,omitempty" xml:"referencePolicy,omitempty"`
	// A list of _include values supported by the server.
	//
	// Comment: Documenting [`_include`](http://hl7.org/fhir/R5/search.html#revinclude) support helps set conformance expectations for the desired system. Still, it is a level of detail that might not be exposed by production servers or clients when using CapabilityStatement to describe an actual implementation. If this list is empty, the server does not support includes. Support for *iterative* (a.k.a., recursive) `_include` is communicated by listing the iterative includes values supported by the server in the `searchInclude` element of the "root" resource type. For example, to support the following search:
	//
	// `GET [base]/CarePlan?_include=CarePlan:activity-reference:DeviceRequest&_include:iterate=DeviceRequest:device`
	//
	// These values would be listed as part of capabilities for "CarePlan":
	//
	// "searchInclude" : ["CarePlan:activity-reference:DeviceRequest","DeviceRequest:device"],
	//
	//
	SearchInclude []string `bson:"searchInclude,omitempty" json:"searchInclude,omitempty" xml:"searchInclude,omitempty"`
	// A list of _revinclude (reverse include) values supported by the server.
	//
	// Comment: See `CapabilityStatement.rest.resource.searchInclude` comments.
	SearchRevInclude []string `bson:"searchRevInclude,omitempty" json:"searchRevInclude,omitempty" xml:"searchRevInclude,omitempty"`
	// Search parameters for implementations to support and/or make use of - either references to ones defined in the specification, or additional ones defined for/by the implementation.
	//
	// Comment: The search parameters should include the control search parameters such as _sort, _count, etc. that also apply to this resource (though many will be listed at [CapabilityStatement.rest.searchParam](capabilitystatement-definitions.html#CapabilityStatement.rest.searchParam)). The behavior of some search parameters may be further described by other code or extension elements, or narrative within the capability statement or linked [SearchParameter](searchparameter.html#) definitions.
	SearchParam []CapabilityStatementRestResourceSearchParam `bson:"searchParam,omitempty" json:"searchParam,omitempty" xml:"searchParam,omitempty"`
	// Definition of an operation or a named query together with its parameters and their meaning and type.
	//
	// Consult the definition of the operation for details about how to invoke the operation, and the parameters.
	//
	// Comment: Operations linked from CapabilityStatement.rest.resource.operation must have OperationDefinition.type = true or OperationDefinition.instance = true.
	//
	// If an operation that is listed in multiple CapabilityStatement.rest.resource.operation (e.g. for different resource types), then clients should understand that the operation is only supported on the specified resource types, and that may be a subset of those listed in OperationDefinition.resource.
	Operation []CapabilityStatementRestResourceOperation `bson:"operation,omitempty" json:"operation,omitempty" xml:"operation,omitempty"`
}

type CapabilityStatementRestResourceInteraction struct {
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
	// Coded identifier of the operation, supported by the system resource.
	Code TypeRestfulInteraction `bson:"code" json:"code" xml:"code"`
	// Guidance specific to the implementation of this operation, such as 'delete is a logical delete' or 'updates are only allowed with version id' or 'creates permitted from pre-authorized certificates only'.
	Documentation *string `bson:"documentation,omitempty" json:"documentation,omitempty" xml:"documentation,omitempty"`
}

type CapabilityStatementRestResourceSearchParam struct {
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
	// The label used for the search parameter in this particular system's API - i.e.
	//
	// the 'name' portion of the name-value pair that will appear as part of the search URL.  This SHOULD be the same as the SearchParameter.code of the defining SearchParameter.  However, it can sometimes differ if necessary to disambiguate when a server supports multiple SearchParameters that happen to share the same code.
	//
	// Comment: Parameter names cannot overlap with standard parameter names, and standard parameters cannot be redefined. There is no correspondence whatsoever between CapabilityStatement's searchParam.name and SearchParameter.name - the latter is used as a class name when generating code for the search parameter.
	Name string `bson:"name" json:"name" xml:"name"`
	// An absolute URI that is a formal reference to where this parameter was first defined, so that a client can be confident of the meaning of the search parameter (a reference to [SearchParameter.url](searchparameter-definitions.html#SearchParameter.url)).
	//
	// This element SHALL be populated if the search parameter refers to a SearchParameter defined by the FHIR core specification or externally defined IGs.
	//
	// Comment: This SHOULD be present, and matches refers to a SearchParameter by its canonical URL. If systems wish to document their support for modifiers, comparators, target resource types, and chained parameters, they should do using a search parameter resource. This element SHALL be populated if the search parameter refers to a SearchParameter defined by the FHIR core specification or externally defined IGs.
	Definition *string `bson:"definition,omitempty" json:"definition,omitempty" xml:"definition,omitempty"`
	// The type of value a search parameter refers to, and how the content is interpreted.
	//
	// Comment: While this can be looked up from the definition, it is included here as a convenience for systems that autogenerate a query interface based on the server capability statement.  It SHALL be the same as the type in the search parameter definition.
	Type SearchParamType `bson:"type" json:"type" xml:"type"`
	// This allows documentation of any distinct behaviors about how the search parameter is used.
	//
	//  For example, text matching algorithms.
	Documentation *string `bson:"documentation,omitempty" json:"documentation,omitempty" xml:"documentation,omitempty"`
}

type CapabilityStatementRestResourceOperation struct {
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
	// The name of the operation or query.
	//
	// For an operation, this name is prefixed with $ and used in the URL. For a query, this is the name used in the _query parameter when the query is called. This SHOULD be the same as the OperationDefinition.code of the defining OperationDefinition.  However, it can sometimes differ if necessary to disambiguate when a server supports multiple OperationDefinition that happen to share the same code.
	//
	// Comment: The name here SHOULD be the same as the OperationDefinition.code in the referenced OperationDefinition, unless there is a name clash and the OperationDefinition.code cannot be used. The name does not include the "$" portion that is always included in the URL. There is no correspondence whatsoever between CapabilityStatement's operation.name and OperationDefinition.name - the latter is used as a class name when generating code for the operation. HL7 will never define operations that have conflicting names.
	Name string `bson:"name" json:"name" xml:"name"`
	// Where the formal definition can be found.
	//
	// If a server references the base definition of an Operation (i.e. from the specification itself such as ```http://hl7.org/fhir/OperationDefinition/ValueSet-expand```), that means it supports the full capabilities of the operation - e.g. both GET and POST invocation.  If it only supports a subset, it must define its own custom [OperationDefinition](operationdefinition.html#) with a 'base' of the original OperationDefinition.  The custom definition would describe the specific subset of functionality supported.
	//
	// Comment: This can be used to build an HTML form to invoke the operation, for instance.
	Definition string `bson:"definition" json:"definition" xml:"definition"`
	// Documentation that describes anything special about the operation behavior, possibly detailing different behavior for system, type and instance-level invocation of the operation.
	Documentation *string `bson:"documentation,omitempty" json:"documentation,omitempty" xml:"documentation,omitempty"`
}

type CapabilityStatementRestInteraction struct {
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
	// A coded identifier of the operation, supported by the system.
	Code SystemRestfulInteraction `bson:"code" json:"code" xml:"code"`
	// Guidance specific to the implementation of this operation, such as limitations on the kind of transactions allowed, or information about system wide search is implemented.
	Documentation *string `bson:"documentation,omitempty" json:"documentation,omitempty" xml:"documentation,omitempty"`
}

type CapabilityStatementMessaging struct {
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
	// An endpoint (network accessible address) to which messages and/or replies are to be sent.
	Endpoint []CapabilityStatementMessagingEndpoint `bson:"endpoint,omitempty" json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// Length if the receiver's reliable messaging cache in minutes (if a receiver) or how long the cache length on the receiver should be (if a sender).
	//
	// Comment: If this value is missing then the application does not implement (receiver) or depend on (sender) reliable messaging.
	ReliableCache *int `bson:"reliableCache,omitempty" json:"reliableCache,omitempty" xml:"reliableCache,omitempty"`
	// Documentation about the system's messaging capabilities for this endpoint not otherwise documented by the capability statement.
	//
	//  For example, the process for becoming an authorized messaging exchange partner.
	Documentation *string `bson:"documentation,omitempty" json:"documentation,omitempty" xml:"documentation,omitempty"`
	// References to message definitions for messages this system can send or receive.
	//
	// Comment: This is a proposed alternative to the messaging.event structure.
	SupportedMessage []CapabilityStatementMessagingSupportedMessage `bson:"supportedMessage,omitempty" json:"supportedMessage,omitempty" xml:"supportedMessage,omitempty"`
}

type CapabilityStatementMessagingEndpoint struct {
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
	// A list of the messaging transport protocol(s) identifiers, supported by this endpoint.
	Protocol Coding `bson:"protocol" json:"protocol" xml:"protocol"`
	// The network address of the endpoint.
	//
	// For solutions that do not use network addresses for routing, it can be just an identifier.
	Address string `bson:"address" json:"address" xml:"address"`
}

type CapabilityStatementMessagingSupportedMessage struct {
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
	// The mode of this event declaration - whether application is sender or receiver.
	Mode EventCapabilityMode `bson:"mode" json:"mode" xml:"mode"`
	// Points to a message definition that identifies the messaging event, message structure, allowed responses, etc.
	Definition string `bson:"definition" json:"definition" xml:"definition"`
}

type CapabilityStatementDocument struct {
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
	// Mode of this document declaration - whether an application is a producer or consumer.
	Mode DocumentMode `bson:"mode" json:"mode" xml:"mode"`
	// A description of how the application supports or uses the specified document profile.
	//
	//  For example, when documents are created, what action is taken with consumed documents, etc.
	Documentation *string `bson:"documentation,omitempty" json:"documentation,omitempty" xml:"documentation,omitempty"`
	// A profile on the document Bundle that constrains which resources are present, and their contents.
	//
	// Comment: The profile is actually on the Bundle.
	Profile string `bson:"profile" json:"profile" xml:"profile"`
}

func (r CapabilityStatement) ResourceType() string {
	return "CapabilityStatement"
}

// MarshalJSON implements json.Marshaler
func (r CapabilityStatement) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		CapabilityStatement
		ResourceType string `json:"resourceType"`
	}{
		CapabilityStatement: r,
		ResourceType:        r.ResourceType(),
	})
}
