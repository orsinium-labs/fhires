package fhires

import (
	"encoding/json/v2"

	"github.com/shopspring/decimal"
)

// A statement of relationships from one set of concepts to one or more other concepts - either concepts in code systems, or data element/data element concepts, or classes in class models.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/ConceptMap
type ConceptMap struct {
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
	// An absolute URI that is used to identify this concept map when it is referenced in a specification, model, design or an instance; also called its canonical identifier.
	//
	// This SHOULD be globally unique and SHOULD be a literal address at which an authoritative instance of this concept map is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the concept map is stored on different servers.
	//
	// Comment: Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred.  Multiple instances may share the same URL if they have a distinct version.
	//
	// The determination of when to create a new version of a resource (same url, new version) vs. defining a new artifact is up to the author.  Considerations for making this decision are found in [Technical and Business Versions](resource.html#versions).
	//
	// In some cases, the resource can no longer be found at the stated url, but the url itself cannot change. Implementations can use the [meta.source](resource.html#meta) element to indicate where the current master source of the resource can be found.
	URL *string `bson:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
	// A formal identifier that is used to identify this concept map when it is represented in other formats, or referenced in a specification, model, design or an instance.
	//
	// Comment: Typically, this is used for identifiers that can go in an HL7 V3 II (instance identifier) data type, and can then identify this concept map outside of FHIR, where it is not possible to use the logical URI.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The identifier that is used to identify this version of the concept map when it is referenced in a specification, model, design or instance.
	//
	// This is an arbitrary value managed by the concept map author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	//
	// Comment: There may be different concept map instances that have the same identifier but different versions.  The version can be appended to the url in a reference to allow a reference to a particular business version of the concept map with the format [url]|[version]. The version SHOULD NOT contain a '#' - see [Business Version](resource.html#bv-format).
	Version *string `bson:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
	// Indicates the mechanism used to compare versions to determine which ConceptMap is more current.
	//
	// Comment: If set as a string, this is a FHIRPath expression that has two additional context variables passed in - %version1 and %version2 and will return a negative number if version1 is newer, a positive number if version2 and a 0 if the version ordering can't be successfully be determined.
	VersionAlgorithmString *string `bson:"versionAlgorithmString,omitempty" json:"versionAlgorithmString,omitempty" xml:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding `bson:"versionAlgorithmCoding,omitempty" json:"versionAlgorithmCoding,omitempty" xml:"versionAlgorithmCoding,omitempty"`
	// A natural language name identifying the concept map.
	//
	// This name should be usable as an identifier for the module by machine processing applications such as code generation.
	//
	// Comment: The name is not expected to be globally unique. The name should be a simple alphanumeric type name to ensure that it is machine-processing friendly.
	Name *string `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// A short, descriptive, user-friendly title for the concept map.
	//
	// Comment: This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc.
	Title *string `bson:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// The status of this concept map. Enables tracking the life-cycle of the content.
	//
	// Comment: Allows filtering of concept maps that are appropriate for use versus not.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Status PublicationStatus `bson:"status" json:"status" xml:"status"`
	// A Boolean value to indicate that this concept map is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	//
	// Comment: Allows filtering of concept maps that are appropriate for use versus not.
	Experimental *bool `bson:"experimental,omitempty" json:"experimental,omitempty" xml:"experimental,omitempty"`
	// The date  (and optionally time) when the concept map was last significantly changed.
	//
	// The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the concept map changes.
	//
	// Comment: The date is often not tracked until the resource is published, but may be present on draft content. Note that this is not the same as the resource last-modified-date, since the resource may be a secondary representation of the concept map. Additional specific dates may be added as extensions or be found by consulting Provenances associated with past versions of the resource.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Date *string `bson:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// The name of the organization or individual responsible for the release and ongoing maintenance of the concept map.
	//
	// Comment: Usually an organization but may be an individual. The publisher (or steward) of the concept map is the organization or individual primarily responsible for the maintenance and upkeep of the concept map. This is not necessarily the same individual or organization that developed and initially authored the content. The publisher is the primary point of contact for questions or issues with the concept map. This item SHOULD be populated unless the information is available from context.
	Publisher *string `bson:"publisher,omitempty" json:"publisher,omitempty" xml:"publisher,omitempty"`
	// Contact details to assist a user in finding and communicating with the publisher.
	//
	// Comment: May be a web site, an email address, a telephone number, etc.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty" xml:"contact,omitempty"`
	// A free text natural language description of the concept map from a consumer's perspective.
	//
	// Comment: The description is not intended to describe the semantics of the concept map. The description should capture its intended use, which is needed for ensuring integrity for its use in models across future changes.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// The content was developed with a focus and intent of supporting the contexts that are listed.
	//
	// These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate concept map instances.
	//
	// Comment: When multiple useContexts are specified, there is no expectation that all or any of the contexts apply.
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty" xml:"useContext,omitempty"`
	// A legal or geographic region in which the concept map is intended to be used.
	//
	// Comment: It may be possible for the concept map to be used in jurisdictions other than those for which it was originally designed or intended.
	//
	// DEPRECATION NOTE: For consistency, implementations are encouraged to migrate to using the new 'jurisdiction' code in the useContext element.  (I.e. useContext.code indicating http://terminology.hl7.org/CodeSystem/usage-context-type#jurisdiction and useContext.valueCodeableConcept indicating the jurisdiction.)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty" xml:"jurisdiction,omitempty"`
	// Explanation of why this concept map is needed and why it has been designed as it has.
	//
	// Comment: This element does not describe the usage of the concept map. Instead, it provides traceability of ''why'' the resource is either needed or ''why'' it is defined as it is.  This may be used to point to source materials or specifications that drove the structure of this concept map.
	Purpose *string `bson:"purpose,omitempty" json:"purpose,omitempty" xml:"purpose,omitempty"`
	// A copyright statement relating to the concept map and/or its contents.
	//
	// Copyright statements are generally legal restrictions on the use and publishing of the concept map.
	//
	// Comment: Frequently the copyright differs between the concept map and codes that are included. The copyright statement should clearly differentiate between these when required.
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
	// The period during which the ConceptMap content was or is planned to be in active use.
	//
	// Comment: The effective period for a ConceptMap  determines when the content is applicable for usage and is independent of publication and review dates. For example, a map intended to be used for the year 2016 might be published in 2015.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	EffectivePeriod *Period `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty" xml:"effectivePeriod,omitempty"`
	// Descriptions related to the content of the ConceptMap.
	//
	// Topics provide a high-level categorization as well as keywords for the ConceptMap that can be useful for filtering and searching.
	//
	// Comment: DEPRECATION NOTE: For consistency, implementations are encouraged to migrate to using the new 'topic' code in the useContext element.  (I.e. useContext.code indicating http://terminology.hl7.org/CodeSystem/usage-context-type#topic and useContext.valueCodeableConcept indicating the topic)
	Topic []CodeableConcept `bson:"topic,omitempty" json:"topic,omitempty" xml:"topic,omitempty"`
	// An individiual or organization primarily involved in the creation and maintenance of the ConceptMap.
	Author []ContactDetail `bson:"author,omitempty" json:"author,omitempty" xml:"author,omitempty"`
	// An individual or organization primarily responsible for internal coherence of the ConceptMap.
	Editor []ContactDetail `bson:"editor,omitempty" json:"editor,omitempty" xml:"editor,omitempty"`
	// An individual or organization asserted by the publisher to be primarily responsible for review of some aspect of the ConceptMap.
	//
	// Comment: See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Reviewer []ContactDetail `bson:"reviewer,omitempty" json:"reviewer,omitempty" xml:"reviewer,omitempty"`
	// An individual or organization asserted by the publisher to be responsible for officially endorsing the ConceptMap for use in some setting.
	//
	// Comment: See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Endorser []ContactDetail `bson:"endorser,omitempty" json:"endorser,omitempty" xml:"endorser,omitempty"`
	// Related artifacts such as additional documentation, justification, dependencies, bibliographic references, and predecessor and successor artifacts.
	//
	// Comment: Each related artifact is either an attachment, or a reference to another resource, but not both.
	RelatedArtifact []RelatedArtifact `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty" xml:"relatedArtifact,omitempty"`
	// A property defines a slot through which additional information can be provided about a map from source -> target.
	//
	// Comment: Properties may be used to supply for example, mapping priority, provenance, presentation hints, flag as experimental, and additional documentation. Multiple occurrences of ConceptMap.group.element.target.property may occur for a ConceptMap.property where ConceptMap.group.element.target.property.code is the same and the values in ConceptMap.group.element.target.property.value differ.
	Property []ConceptMapProperty `bson:"property,omitempty" json:"property,omitempty" xml:"property,omitempty"`
	// An additionalAttribute defines an additional data element found in the source or target data model where the data will come from or be mapped to.
	//
	// Some mappings are based on data in addition to the source data element, where codes in multiple fields are combined to a single field (or vice versa).
	//
	// Comment: Additional attributes are used to define additional data elements where mapping data can be found. For an example, see [Specimen Type json.-> SNOMED CT Mapping(conceptmap-example-specimen-type.html)
	AdditionalAttribute []ConceptMapAdditionalAttribute `bson:"additionalAttribute,omitempty" json:"additionalAttribute,omitempty" xml:"additionalAttribute,omitempty"`
	// Identifier for the source value set that contains the concepts that are being mapped and provides context for the mappings.
	//
	//  Limits the scope of the map to source codes (ConceptMap.group.element code or valueSet) that are members of this value set.
	//
	// Comment: Should be a version specific reference. URIs SHOULD be absolute. If there is no sourceScope or targetScope value set, there is no specified context for the map (not recommended).  The sourceScope value set may select codes from either an explicit (standard or local) or implicit code system.
	SourceScopeUri       *string `bson:"sourceScopeUri,omitempty" json:"sourceScopeUri,omitempty" xml:"sourceScopeUri,omitempty"`
	SourceScopeCanonical *string `bson:"sourceScopeCanonical,omitempty" json:"sourceScopeCanonical,omitempty" xml:"sourceScopeCanonical,omitempty"`
	// Identifier for the target value set that provides important context about how the mapping choices are made.
	//
	//  Limits the scope of the map to target codes (ConceptMap.group.element.target code or valueSet) that are members of this value set.
	//
	// Comment: Should be a version specific reference. URIs SHOULD be absolute. If there is no sourceScope or targetScope value set, there is no specified context for the map (not recommended).  The targetScope value set may select codes from either an explicit (standard or local) or implicit code system.
	TargetScopeUri       *string `bson:"targetScopeUri,omitempty" json:"targetScopeUri,omitempty" xml:"targetScopeUri,omitempty"`
	TargetScopeCanonical *string `bson:"targetScopeCanonical,omitempty" json:"targetScopeCanonical,omitempty" xml:"targetScopeCanonical,omitempty"`
	// A group of mappings that all have the same source and target system.
	Group []ConceptMapGroup `bson:"group,omitempty" json:"group,omitempty" xml:"group,omitempty"`
}

type ConceptMapProperty struct {
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
	// The code is used internally (in ConceptMap.group.element.target.property.code) and also in the $translate operation.
	Code string `bson:"code" json:"code" xml:"code"`
	// Reference to the formal meaning of the property.
	Uri *string `bson:"uri,omitempty" json:"uri,omitempty" xml:"uri,omitempty"`
	// A description of the property - why it is defined, and how its value might be used.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// The type of the property value.
	Type ConceptMapPropertyType `bson:"type" json:"type" xml:"type"`
	// The CodeSystem that defines the codes from which values of type ```code``` in property values.
	System *string `bson:"system,omitempty" json:"system,omitempty" xml:"system,omitempty"`
}

type ConceptMapAdditionalAttribute struct {
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
	// A code that is used to identify this additional data attribute.
	//
	// The code is used internally in ConceptMap.group.element.target.dependsOn.attribute and ConceptMap.group.element.target.product.attribute.
	Code string `bson:"code" json:"code" xml:"code"`
	// Reference to the formal definition of the source/target data element.
	//
	// For elements defined by the FHIR specification, or using a FHIR logical model, the correct format is {canonical-url}#{element-id}.
	Uri *string `bson:"uri,omitempty" json:"uri,omitempty" xml:"uri,omitempty"`
	// A description of the additional attribute and/or the data element it refers to - why it is defined, and how the value might be used in mappings, and a discussion of issues associated with the use of the data element.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// The type of the source data contained in this concept map for this data element.
	//
	// Comment: When the source or target is not content defined by FHIR, some additional application logic might be required to map from the FHIR Datatype to the externally defined content
	Type ConceptMapAttributeType `bson:"type" json:"type" xml:"type"`
}

type ConceptMapGroup struct {
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
	// An absolute URI that identifies the source system where the concepts to be mapped are defined.
	//
	// Comment: This is not needed if the source value set is specified and it contains concepts from only a single system.
	Source *string `bson:"source,omitempty" json:"source,omitempty" xml:"source,omitempty"`
	// An absolute URI that identifies the target system that the concepts will be mapped to.
	//
	// Comment: This is not needed if the target value set is specified and it contains concepts from only a single system. The group target may also be omitted if all of the target element relationship values are 'not-related-to'.
	Target *string `bson:"target,omitempty" json:"target,omitempty" xml:"target,omitempty"`
	// Mappings for an individual concept in the source to one or more concepts in the target.
	//
	// Comment: Generally, the ideal is that there would only be one mapping for each concept in the source value set, but a given concept may be mapped multiple times with different comments or dependencies.
	Element []ConceptMapGroupElement `bson:"element" json:"element" xml:"element"`
	// What to do when there is no mapping to a target concept from the source concept and ConceptMap.group.element.noMap is not true.
	//
	// This provides the "default" to be applied when there is no target concept mapping specified or the expansion of ConceptMap.group.element.target.valueSet is empty.
	//
	// Comment: The 'unmapped' element is ignored if a code is specified to have relationship = not-related-to or if ConceptMap.group.element.noMap = true.
	Unmapped *ConceptMapGroupUnmapped `bson:"unmapped,omitempty" json:"unmapped,omitempty" xml:"unmapped,omitempty"`
}

type ConceptMapGroupElement struct {
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
	// Identity (code or path) or the element/item being mapped.
	Code *string `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// The display for the code.
	//
	// The display is only provided to help editors when editing the concept map.
	//
	// Comment: The display is ignored when processing the map.
	Display *string `bson:"display,omitempty" json:"display,omitempty" xml:"display,omitempty"`
	// The set of concepts from the ConceptMap.group.source code system which are all being mapped to the target as part of this mapping rule.
	//
	// Comment: If the value set expansion is empty then there are no source concepts to map in this rule.
	ValueSet *string `bson:"valueSet,omitempty" json:"valueSet,omitempty" xml:"valueSet,omitempty"`
	// If noMap = true this indicates that no mapping to a target concept exists for this source concept.
	//
	// Comment: If noMap = true, then the behavior of ConceptMap.group.unmapped (if present) does not apply.
	NoMap *bool `bson:"noMap,omitempty" json:"noMap,omitempty" xml:"noMap,omitempty"`
	// A concept from the target value set that this concept maps to.
	//
	// Comment: Ideally there would only be one map, with an 'equivalent' mapping. But multiple maps are allowed for several narrower (i.e. source-is-broader-than-target) options, or to assert that other concepts are not related.
	Target []ConceptMapGroupElementTarget `bson:"target,omitempty" json:"target,omitempty" xml:"target,omitempty"`
}

type ConceptMapGroupElementTarget struct {
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
	// Identity (code or path) or the element/item that the map refers to.
	Code *string `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// The display for the code.
	//
	// The display is only provided to help editors when editing the concept map.
	//
	// Comment: The display is ignored when processing the map.
	Display *string `bson:"display,omitempty" json:"display,omitempty" xml:"display,omitempty"`
	// The set of concepts from the ConceptMap.group.target code system which are all being mapped to as part of this mapping rule.
	//
	// The effect of using this data element is the same as having multiple ConceptMap.group.element.target elements with one for each concept in the ConceptMap.group.element.target.valueSet value set.
	//
	// Comment: If the the value set expansion is empty then the source concept(s) are unmapped and the behavior of ConceptMap.group.unmapped (if present) applies.
	ValueSet *string `bson:"valueSet,omitempty" json:"valueSet,omitempty" xml:"valueSet,omitempty"`
	// The relationship between the source and target concepts.
	//
	// The relationship is read from source to target (e.g. source-is-narrower-than-target).
	//
	// Comment: This element is labeled as a modifier because it may indicate that there is no mapping.
	Relationship ConceptMapRelationship `bson:"relationship" json:"relationship" xml:"relationship"`
	// A description of status/issues in mapping that conveys additional information not represented in  the structured data.
	Comment *string `bson:"comment,omitempty" json:"comment,omitempty" xml:"comment,omitempty"`
	// A property value for this source -> target mapping.
	Property []ConceptMapGroupElementTargetProperty `bson:"property,omitempty" json:"property,omitempty" xml:"property,omitempty"`
	// A set of additional dependencies for this mapping to hold.
	//
	// This mapping is only applicable if the specified data attribute can be resolved, and it has the specified value.
	DependsOn []ConceptMapGroupElementTargetDependsOn `bson:"dependsOn,omitempty" json:"dependsOn,omitempty" xml:"dependsOn,omitempty"`
	// Product is the output of a ConceptMap that provides additional values that go in other attributes / data elemnts of the target data.
	Product []ConceptMapGroupElementTargetDependsOn `bson:"product,omitempty" json:"product,omitempty"`
}

type ConceptMapGroupElementTargetProperty struct {
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
	// A reference to a mapping property defined in ConceptMap.property.
	Code string `bson:"code" json:"code" xml:"code"`
	// The value of this property.
	//
	// If the type chosen for this element is 'code', then the property SHALL be defined in a ConceptMap.property element.
	ValueCoding   Coding          `bson:"valueCoding" json:"valueCoding" xml:"valueCoding"`
	ValueString   string          `bson:"valueString" json:"valueString" xml:"valueString"`
	ValueInteger  int             `bson:"valueInteger" json:"valueInteger" xml:"valueInteger"`
	ValueBoolean  bool            `bson:"valueBoolean" json:"valueBoolean" xml:"valueBoolean"`
	ValueDateTime string          `bson:"valueDateTime" json:"valueDateTime" xml:"valueDateTime"`
	ValueDecimal  decimal.Decimal `bson:"valueDecimal" json:"valueDecimal" xml:"valueDecimal"`
	ValueCode     string          `bson:"valueCode" json:"valueCode" xml:"valueCode"`
}

type ConceptMapGroupElementTargetDependsOn struct {
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
	// A reference to the additional attribute that holds a value the map depends on.
	Attribute string `bson:"attribute" json:"attribute" xml:"attribute"`
	// Data element value that the map depends on / produces.
	//
	// Comment: If the data type is a code, then the code system is .group.source for .dependsOn.valueCode and .group.target for .product.valueCode.
	ValueCode     *string   `bson:"valueCode,omitempty" json:"valueCode,omitempty" xml:"valueCode,omitempty"`
	ValueCoding   *Coding   `bson:"valueCoding,omitempty" json:"valueCoding,omitempty" xml:"valueCoding,omitempty"`
	ValueString   *string   `bson:"valueString,omitempty" json:"valueString,omitempty" xml:"valueString,omitempty"`
	ValueBoolean  *bool     `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty" xml:"valueBoolean,omitempty"`
	ValueQuantity *Quantity `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty" xml:"valueQuantity,omitempty"`
	// This mapping applies if the data element value is a code from this value set.
	ValueSet *string `bson:"valueSet,omitempty" json:"valueSet,omitempty" xml:"valueSet,omitempty"`
}

type ConceptMapGroupUnmapped struct {
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
	// Defines which action to take if there is no match for the source concept in the target system designated for the group.
	//
	// One of 3 actions are possible: use the unmapped source code (this is useful when doing a mapping between versions, and only a few codes have changed), use a fixed code (a default code), or alternatively, a reference to a different concept map can be provided (by canonical URL).
	Mode ConceptMapGroupUnmappedMode `bson:"mode" json:"mode" xml:"mode"`
	// The fixed code to use when the mode = 'fixed'  - all unmapped codes are mapped to a single fixed code.
	Code *string `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// The display for the code.
	//
	// The display is only provided to help editors when editing the concept map.
	//
	// Comment: The display is ignored when processing the map.
	Display *string `bson:"display,omitempty" json:"display,omitempty" xml:"display,omitempty"`
	// The set of fixed codes to use when the mode = 'fixed'  - all unmapped codes are mapped to each of the fixed codes.
	ValueSet *string `bson:"valueSet,omitempty" json:"valueSet,omitempty" xml:"valueSet,omitempty"`
	// The default relationship value to apply between the source and target concepts when the source code is unmapped and the mode is 'fixed' or 'use-source-code'.
	//
	// Comment: This element is labeled as a modifier because it may indicate that there is no mapping.
	Relationship *ConceptMapRelationship `bson:"relationship,omitempty" json:"relationship,omitempty" xml:"relationship,omitempty"`
	// The canonical reference to an additional ConceptMap resource instance to use for mapping if this ConceptMap resource contains no matching mapping for the source concept.
	OtherMap *string `bson:"otherMap,omitempty" json:"otherMap,omitempty" xml:"otherMap,omitempty"`
}

func (r ConceptMap) ResourceType() string {
	return "ConceptMap"
}

// MarshalJSON implements json.Marshaler
func (r ConceptMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ConceptMap
		ResourceType string `json:"resourceType"`
	}{
		ConceptMap:   r,
		ResourceType: r.ResourceType(),
	})
}
