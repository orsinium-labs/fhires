package fhires

import v2 "encoding/json/v2"

// A walkthrough of a workflow showing the interaction between systems and the instances shared, possibly including the evolution of instances over time.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/ExampleScenario
type ExampleScenario struct {
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
	// An absolute URI that is used to identify this example scenario when it is referenced in a specification, model, design or an instance; also called its canonical identifier.
	//
	// This SHOULD be globally unique and SHOULD be a literal address at which an authoritative instance of this example scenario is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the example scenario is stored on different servers.
	//
	// Comment: Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred.  Multiple instances may share the same URL if they have a distinct version.
	//
	// The determination of when to create a new version of a resource (same url, new version) vs. defining a new artifact is up to the author.  Considerations for making this decision are found in [Technical and Business Versions](resource.html#versions).
	//
	// In some cases, the resource can no longer be found at the stated url, but the url itself cannot change. Implementations can use the [meta.source](resource.html#meta) element to indicate where the current master source of the resource can be found.
	URL *string `bson:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
	// A formal identifier that is used to identify this example scenario when it is represented in other formats, or referenced in a specification, model, design or an instance.
	//
	// Comment: Typically, this is used for identifiers that can go in an HL7 V3 II (instance identifier) data type, and can then identify this example scenario outside of FHIR, where it is not possible to use the logical URI.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The identifier that is used to identify this version of the example scenario when it is referenced in a specification, model, design or instance.
	//
	// This is an arbitrary value managed by the example scenario author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	//
	// Comment: There may be different example scenario instances that have the same identifier but different versions.  The version can be appended to the url in a reference to allow a reference to a particular business version of the example scenario with the format [url]|[version]. The version SHOULD NOT contain a '#' - see [Business Version](resource.html#bv-format).
	Version *string `bson:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
	// Indicates the mechanism used to compare versions to determine which is more current.
	//
	// Comment: If set as a string, this is a FHIRPath expression that has two additional context variables passed in - %version1 and %version2 and will return a negative number if version1 is newer, a positive number if version2 and a 0 if the version ordering can't be successfully be determined.
	VersionAlgorithmString *string `bson:"versionAlgorithmString,omitempty" json:"versionAlgorithmString,omitempty" xml:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding `bson:"versionAlgorithmCoding,omitempty" json:"versionAlgorithmCoding,omitempty" xml:"versionAlgorithmCoding,omitempty"`
	// Temporarily retained for tooling purposes.
	//
	// Comment: Not needed as code will never be generated from one of these
	Name *string `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// A short, descriptive, user-friendly title for the ExampleScenario.
	//
	// Comment: This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc.
	Title *string `bson:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// The status of this example scenario. Enables tracking the life-cycle of the content.
	//
	// Comment: Allows filtering of example scenarios that are appropriate for use versus not.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Status PublicationStatus `bson:"status" json:"status" xml:"status"`
	// A Boolean value to indicate that this example scenario is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	//
	// Comment: Allows filtering of example scenarios that are appropriate for use versus not.
	Experimental *bool `bson:"experimental,omitempty" json:"experimental,omitempty" xml:"experimental,omitempty"`
	// The date  (and optionally time) when the example scenario was last significantly changed.
	//
	// The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the example scenario changes. (e.g. the 'content logical definition').
	//
	// Comment: The date is often not tracked until the resource is published, but may be present on draft content. Note that this is not the same as the resource last-modified-date, since the resource may be a secondary representation of the example scenario. Additional specific dates may be added as extensions or be found by consulting Provenances associated with past versions of the resource.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Date *string `bson:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// The name of the organization or individual responsible for the release and ongoing maintenance of the example scenario.
	//
	// Comment: Usually an organization but may be an individual. The publisher (or steward) of the example scenario is the organization or individual primarily responsible for the maintenance and upkeep of the example scenario. This is not necessarily the same individual or organization that developed and initially authored the content. The publisher is the primary point of contact for questions or issues with the example scenario. This item SHOULD be populated unless the information is available from context.
	Publisher *string `bson:"publisher,omitempty" json:"publisher,omitempty" xml:"publisher,omitempty"`
	// Contact details to assist a user in finding and communicating with the publisher.
	//
	// Comment: May be a web site, an email address, a telephone number, etc.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty" xml:"contact,omitempty"`
	// A free text natural language description of the ExampleScenario from a consumer's perspective.
	//
	// Comment: This description can be used to capture details such as comments about misuse. It is not a rendering of the ExampleScenario as conveyed in the 'text' field of the resource itself.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// The content was developed with a focus and intent of supporting the contexts that are listed.
	//
	// These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate example scenario instances.
	//
	// Comment: When multiple useContexts are specified, there is no expectation that all or any of the contexts apply.
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty" xml:"useContext,omitempty"`
	// A legal or geographic region in which the example scenario is intended to be used.
	//
	// Comment: It may be possible for the example scenario to be used in jurisdictions other than those for which it was originally designed or intended.
	//
	// DEPRECATION NOTE: For consistency, implementations are encouraged to migrate to using the new 'jurisdiction' code in the useContext element.  (I.e. useContext.code indicating http://terminology.hl7.org/CodeSystem/usage-context-type#jurisdiction and useContext.valueCodeableConcept indicating the jurisdiction.)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty" xml:"jurisdiction,omitempty"`
	// What the example scenario resource is created for.
	//
	// This should not be used to show the business purpose of the scenario itself, but the purpose of documenting a scenario.
	//
	// Comment: This element does not describe the usage of the example scenario. Instead, it provides traceability of ''why'' the resource is either needed or ''why'' it is defined as it is.  This may be used to point to source materials or specifications that drove the structure of this example scenario.
	Purpose *string `bson:"purpose,omitempty" json:"purpose,omitempty" xml:"purpose,omitempty"`
	// A copyright statement relating to the example scenario and/or its contents.
	//
	// Copyright statements are generally legal restrictions on the use and publishing of the example scenario.
	//
	// Comment: ...
	Copyright *string `bson:"copyright,omitempty" json:"copyright,omitempty" xml:"copyright,omitempty"`
	// A short string (<50 characters), suitable for inclusion in a page footer that identifies the copyright holder, effective period, and optionally whether rights are resctricted.
	//
	// (e.g. 'All rights reserved', 'Some rights reserved').
	//
	// Comment: The (c) symbol should NOT be included in this string. It will be added by software when rendering the notation. Full details about licensing, restrictions, warrantees, etc. goes in the more general 'copyright' element.
	CopyrightLabel *string `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty" xml:"copyrightLabel,omitempty"`
	// A system or person who shares or receives an instance within the scenario.
	Actor []ExampleScenarioActor `bson:"actor,omitempty" json:"actor,omitempty" xml:"actor,omitempty"`
	// A single data collection that is shared as part of the scenario.
	Instance []ExampleScenarioInstance `bson:"instance,omitempty" json:"instance,omitempty" xml:"instance,omitempty"`
	// A group of operations that represents a significant step within a scenario.
	//
	// Comment: Some scenarios might describe only one process.
	Process []ExampleScenarioProcess `bson:"process,omitempty" json:"process,omitempty" xml:"process,omitempty"`
}

type ExampleScenarioActor struct {
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
	// A unique string within the scenario that is used to reference the actor.
	Key string `bson:"key" json:"key" xml:"key"`
	// The category of actor - person or system.
	Type ExampleScenarioActorType `bson:"type" json:"type" xml:"type"`
	// The human-readable name for the actor used when rendering the scenario.
	Title string `bson:"title" json:"title" xml:"title"`
	// An explanation of who/what the actor is and its role in the scenario.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
}

type ExampleScenarioInstance struct {
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
	// A unique string within the scenario that is used to reference the instance.
	Key string `bson:"key" json:"key" xml:"key"`
	// A code indicating the kind of data structure (FHIR resource or some other standard) this is an instance of.
	StructureType Coding `bson:"structureType" json:"structureType" xml:"structureType"`
	// Conveys the version of the data structure instantiated.
	//
	//  I.e. what release of FHIR, X12, OpenEHR, etc. is instance compliant with.
	StructureVersion *string `bson:"structureVersion,omitempty" json:"structureVersion,omitempty" xml:"structureVersion,omitempty"`
	// Refers to a profile, template or other ruleset the instance adheres to.
	StructureProfileCanonical *string `bson:"structureProfileCanonical,omitempty" json:"structureProfileCanonical,omitempty" xml:"structureProfileCanonical,omitempty"`
	StructureProfileUri       *string `bson:"structureProfileUri,omitempty" json:"structureProfileUri,omitempty" xml:"structureProfileUri,omitempty"`
	// A short descriptive label the instance to be used in tables or diagrams.
	Title string `bson:"title" json:"title" xml:"title"`
	// An explanation of what the instance contains and what it's for.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Points to an instance (typically an example) that shows the data that would corespond to this instance.
	//
	// Comment: If not conveying FHIR data or not using the same version of FHIR as this ExampleScenario instance, the reference must be to a Binary.
	Content *Reference `bson:"content,omitempty" json:"content,omitempty" xml:"content,omitempty"`
	// Represents the instance as it was at a specific time-point.
	//
	// Comment: Not used if an instance doesn't change
	Version []ExampleScenarioInstanceVersion `bson:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
	// References to other instances that can be found within this instance (e.g.
	//
	// the observations contained in a bundle).
	ContainedInstance []ExampleScenarioInstanceContainedInstance `bson:"containedInstance,omitempty" json:"containedInstance,omitempty" xml:"containedInstance,omitempty"`
}

type ExampleScenarioInstanceVersion struct {
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
	// A unique string within the instance that is used to reference the version of the instance.
	Key string `bson:"key" json:"key" xml:"key"`
	// A short descriptive label the version to be used in tables or diagrams.
	Title string `bson:"title" json:"title" xml:"title"`
	// An explanation of what this specific version of the instance contains and represents.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Points to an instance (typically an example) that shows the data that would flow at this point in the scenario.
	//
	// Comment: If not conveying FHIR data or not using the same version of FHIR as this ExampleScenario instance, the reference must be to a Binary.
	Content *Reference `bson:"content,omitempty" json:"content,omitempty" xml:"content,omitempty"`
}

type ExampleScenarioInstanceContainedInstance struct {
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
	// A reference to the key of an instance found within this one.
	InstanceReference string `bson:"instanceReference" json:"instanceReference" xml:"instanceReference"`
	// A reference to the key of a specific version of an instance in this instance.
	//
	// Comment: Required if the referenced instance has versions
	VersionReference *string `bson:"versionReference,omitempty" json:"versionReference,omitempty" xml:"versionReference,omitempty"`
}

type ExampleScenarioProcess struct {
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
	// A short descriptive label the process to be used in tables or diagrams.
	Title string `bson:"title" json:"title" xml:"title"`
	// An explanation of what the process represents and what it does.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Description of the initial state of the actors, environment and data before the process starts.
	PreConditions *string `bson:"preConditions,omitempty" json:"preConditions,omitempty" xml:"preConditions,omitempty"`
	// Description of the final state of the actors, environment and data after the process has been successfully completed.
	//
	// Comment: Alternate steps might not result in all post conditions holding
	PostConditions *string `bson:"postConditions,omitempty" json:"postConditions,omitempty" xml:"postConditions,omitempty"`
	// A significant action that occurs as part of the process.
	Step []ExampleScenarioProcessStep `bson:"step,omitempty" json:"step,omitempty" xml:"step,omitempty"`
}

type ExampleScenarioProcessStep struct {
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
	// The sequential number of the step, e.g. 1.2.5.
	//
	// Comment: If step numbers are simultaneous, they will be the same.  Numbers don't have to be numeric.  E.g. '2c)' is a valid step number
	Number *string `bson:"number,omitempty" json:"number,omitempty" xml:"number,omitempty"`
	// Indicates that the step is a complex sub-process with its own steps.
	Process *ExampleScenarioProcess `bson:"process,omitempty" json:"process,omitempty"`
	// Indicates that the step is defined by a seaparate scenario instance.
	Workflow *string `bson:"workflow,omitempty" json:"workflow,omitempty" xml:"workflow,omitempty"`
	// The step represents a single operation invoked on receiver by sender.
	Operation *ExampleScenarioProcessStepOperation `bson:"operation,omitempty" json:"operation,omitempty" xml:"operation,omitempty"`
	// Indicates an alternative step that can be taken instead of the sub-process, scenario or operation.
	//
	//  E.g. to represent non-happy-path/exceptional/atypical circumstances.
	Alternative []ExampleScenarioProcessStepAlternative `bson:"alternative,omitempty" json:"alternative,omitempty" xml:"alternative,omitempty"`
	// If true, indicates that, following this step, there is a pause in the flow and the subsequent step will occur at some later time (triggered by some event).
	Pause *bool `bson:"pause,omitempty" json:"pause,omitempty" xml:"pause,omitempty"`
}

type ExampleScenarioProcessStepOperation struct {
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
	// The standardized type of action (FHIR or otherwise).
	Type *Coding `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// A short descriptive label the step to be used in tables or diagrams.
	Title string `bson:"title" json:"title" xml:"title"`
	// The system that invokes the action/transmits the data.
	//
	// Comment: This must either be the 'key' of one of the actors defined in this scenario or the special keyword 'OTHER' if the initiator is not one of the actors defined for the scenario.  (Multiple references to 'OTHER' don't necessarily indicate the same actor.)
	Initiator *string `bson:"initiator,omitempty" json:"initiator,omitempty" xml:"initiator,omitempty"`
	// The system on which the action is invoked/receives the data.
	//
	// Comment: This must either be the 'key' of one of the actors defined in this scenario or the special keyword 'OTHER' if the receiver is not one of the actors defined for the scenario.  (Multiple references to 'OTHER' don't necessarily indicate the same actor.)  In some cases, receiver could be same as sender if representing a local operation
	Receiver *string `bson:"receiver,omitempty" json:"receiver,omitempty" xml:"receiver,omitempty"`
	// An explanation of what the operation represents and what it does.
	//
	// Comment: This should contain information not already present in the process step.  It's more of a description of what the operation in general does - if not already evident from the operation.type
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// If false, the initiator is deactivated right after the operation.
	//
	// Comment: De-activation of an actor means they have no further role until such time as they are the recipient of an operation.
	InitiatorActive *bool `bson:"initiatorActive,omitempty" json:"initiatorActive,omitempty" xml:"initiatorActive,omitempty"`
	// If false, the receiver is deactivated right after the operation.
	//
	// Comment: De-activation of an actor means they have no further role until such time as they are the recipient of an operation.
	ReceiverActive *bool `bson:"receiverActive,omitempty" json:"receiverActive,omitempty" xml:"receiverActive,omitempty"`
	// A reference to the instance that is transmitted from requester to receiver as part of the invocation of the operation.
	Request *ExampleScenarioInstanceContainedInstance `bson:"request,omitempty" json:"request,omitempty"`
	// A reference to the instance that is transmitted from receiver to requester as part of the operation's synchronous response (if any).
	Response *ExampleScenarioInstanceContainedInstance `bson:"response,omitempty" json:"response,omitempty"`
}

type ExampleScenarioProcessStepAlternative struct {
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
	// The label to display for the alternative that gives a sense of the circumstance in which the alternative should be invoked.
	Title string `bson:"title" json:"title" xml:"title"`
	// A human-readable description of the alternative explaining when the alternative should occur rather than the base step.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Indicates the operation, sub-process or scenario that happens if the alternative option is selected.
	Step []ExampleScenarioProcessStep `bson:"step,omitempty" json:"step,omitempty"`
}

func (r ExampleScenario) ResourceType() string {
	return "ExampleScenario"
}

// MarshalJSON implements json.Marshaler
func (r ExampleScenario) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		ExampleScenario
		ResourceType string `json:"resourceType"`
	}{
		ExampleScenario: r,
		ResourceType:    r.ResourceType(),
	})
}
