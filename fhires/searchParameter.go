package fhires

import v2 "encoding/json/v2"

// A search parameter that defines a named search item that can be used to search/filter on a resource.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/SearchParameter
type SearchParameter struct {
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
	// An absolute URI that is used to identify this search parameter when it is referenced in a specification, model, design or an instance; also called its canonical identifier.
	//
	// This SHOULD be globally unique and SHOULD be a literal address at which an authoritative instance of this search parameter is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the search parameter is stored on different servers.
	//
	// Comment: Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred.  Multiple instances may share the same URL if they have a distinct version.
	//
	// The determination of when to create a new version of a resource (same url, new version) vs. defining a new artifact is up to the author.  Considerations for making this decision are found in [Technical and Business Versions](resource.html#versions).
	//
	// In some cases, the resource can no longer be found at the stated url, but the url itself cannot change. Implementations can use the [meta.source](resource.html#meta) element to indicate where the current master source of the resource can be found.
	URL string `bson:"url" json:"url" xml:"url"`
	// A formal identifier that is used to identify this search parameter when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The identifier that is used to identify this version of the search parameter when it is referenced in a specification, model, design or instance.
	//
	// This is an arbitrary value managed by the search parameter author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	//
	// Comment: There may be different search parameter instances that have the same identifier but different versions.  The version can be appended to the url in a reference to allow a reference to a particular business version of the search parameter with the format [url]|[version]. The version SHOULD NOT contain a '#' - see [Business Version](resource.html#bv-format).
	Version *string `bson:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
	// Indicates the mechanism used to compare versions to determine which is more current.
	//
	// Comment: If set as a string, this is a FHIRPath expression that has two additional context variables passed in - %version1 and %version2 and will return a negative number if version1 is newer, a positive number if version2 and a 0 if the version ordering can't be successfully be determined.
	VersionAlgorithmString *string `bson:"versionAlgorithmString,omitempty" json:"versionAlgorithmString,omitempty" xml:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding `bson:"versionAlgorithmCoding,omitempty" json:"versionAlgorithmCoding,omitempty" xml:"versionAlgorithmCoding,omitempty"`
	// A natural language name identifying the search parameter.
	//
	// This name should be usable as an identifier for the module by machine processing applications such as code generation.
	//
	// Comment: The name is not expected to be globally unique. The name should be a simple alphanumeric type name to ensure that it is machine-processing friendly.
	Name string `bson:"name" json:"name" xml:"name"`
	// A short, descriptive, user-friendly title for the search parameter.
	//
	// Comment: This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc.
	Title *string `bson:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Where this search parameter is originally defined.
	//
	// If a derivedFrom is provided, then the details in the search parameter must be consistent with the definition from which it is defined. i.e. the parameter should have the same meaning, and (usually) the functionality should be a proper subset of the underlying search parameter.
	//
	// Comment: The intent of this is that a server can designate that it provides support for a search parameter defined in the specification itself (e.g.  [`value-quantity`](http://hl7.org/fhir/SearchParameter/Observation-value-quantity), but detail how it is supported by the server.
	DerivedFrom *string `bson:"derivedFrom,omitempty" json:"derivedFrom,omitempty" xml:"derivedFrom,omitempty"`
	// The status of this search parameter. Enables tracking the life-cycle of the content.
	//
	// Comment: Allows filtering of search parameters that are appropriate for use versus not.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Status PublicationStatus `bson:"status" json:"status" xml:"status"`
	// A Boolean value to indicate that this search parameter is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	//
	// Comment: Allows filtering of search parameters that are appropriate for use versus not.
	Experimental *bool `bson:"experimental,omitempty" json:"experimental,omitempty" xml:"experimental,omitempty"`
	// The date  (and optionally time) when the search parameter was last significantly changed.
	//
	// The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the search parameter changes.
	//
	// Comment: The date is often not tracked until the resource is published, but may be present on draft content. Note that this is not the same as the resource last-modified-date, since the resource may be a secondary representation of the search parameter. Additional specific dates may be added as extensions or be found by consulting Provenances associated with past versions of the resource.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Date *string `bson:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// The name of the organization or individual tresponsible for the release and ongoing maintenance of the search parameter.
	//
	// Comment: Usually an organization but may be an individual. The publisher (or steward) of the search parameter is the organization or individual primarily responsible for the maintenance and upkeep of the search parameter. This is not necessarily the same individual or organization that developed and initially authored the content. The publisher is the primary point of contact for questions or issues with the search parameter. This item SHOULD be populated unless the information is available from context.
	Publisher *string `bson:"publisher,omitempty" json:"publisher,omitempty" xml:"publisher,omitempty"`
	// Contact details to assist a user in finding and communicating with the publisher.
	//
	// Comment: May be a web site, an email address, a telephone number, etc.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty" xml:"contact,omitempty"`
	// And how it used.
	//
	// Comment: This description can be used to capture details such as comments about misuse, instructions for clinical use and interpretation, literature references, examples from the paper world, etc. It is not a rendering of the search parameter as conveyed in the 'text' field of the resource itself. This item SHOULD be populated unless the information is available from context (e.g. the language of the search parameter is presumed to be the predominant language in the place the search parameter was created).
	Description string `bson:"description" json:"description" xml:"description"`
	// The content was developed with a focus and intent of supporting the contexts that are listed.
	//
	// These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate search parameter instances.
	//
	// Comment: When multiple useContexts are specified, there is no expectation that all or any of the contexts apply.
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty" xml:"useContext,omitempty"`
	// A legal or geographic region in which the search parameter is intended to be used.
	//
	// Comment: It may be possible for the search parameter to be used in jurisdictions other than those for which it was originally designed or intended.
	//
	// DEPRECATION NOTE: For consistency, implementations are encouraged to migrate to using the new 'jurisdiction' code in the useContext element.  (I.e. useContext.code indicating http://terminology.hl7.org/CodeSystem/usage-context-type#jurisdiction and useContext.valueCodeableConcept indicating the jurisdiction.)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty" xml:"jurisdiction,omitempty"`
	// Explanation of why this search parameter is needed and why it has been designed as it has.
	//
	// Comment: This element does not describe the usage of the search parameter. Instead, it provides traceability of ''why'' the resource is either needed or ''why'' it is defined as it is.  This may be used to point to source materials or specifications that drove the structure of this search parameter.
	Purpose *string `bson:"purpose,omitempty" json:"purpose,omitempty" xml:"purpose,omitempty"`
	// A copyright statement relating to the search parameter and/or its contents.
	//
	// Copyright statements are generally legal restrictions on the use and publishing of the search parameter.
	//
	// Comment: ...
	Copyright *string `bson:"copyright,omitempty" json:"copyright,omitempty" xml:"copyright,omitempty"`
	// A short string (<50 characters), suitable for inclusion in a page footer that identifies the copyright holder, effective period, and optionally whether rights are resctricted.
	//
	// (e.g. 'All rights reserved', 'Some rights reserved').
	//
	// Comment: The (c) symbol should NOT be included in this string. It will be added by software when rendering the notation. Full details about licensing, restrictions, warrantees, etc. goes in the more general 'copyright' element.
	CopyrightLabel *string `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty" xml:"copyrightLabel,omitempty"`
	// The label that is recommended to be used in the URL or the parameter name in a parameters resource for this search parameter.
	//
	//  In some cases, servers may need to use a different CapabilityStatement searchParam.name to differentiate between multiple SearchParameters that happen to have the same code.
	//
	// Comment: For maximum compatibility, use only lowercase ASCII characters. Note that HL7 will never define multiple search parameters with the same code.
	Code string `bson:"code" json:"code" xml:"code"`
	// The base resource type(s) that this search parameter can be used against.
	//
	// Comment: A search parameter must always apply to at least one resource type. When search parameters apply to more than one resource type, they can be used against any of the listed resource types, or in a cross-type search (see [Cross Resource Search](http.html#xres-search)).
	Base []string `bson:"base" json:"base" xml:"base"`
	// The type of value that a search parameter may contain, and how the content is interpreted.
	Type SearchParamType `bson:"type" json:"type" xml:"type"`
	// A FHIRPath expression that returns a set of elements for the search parameter.
	//
	// Comment: Note that the elements returned by the expression are sometimes complex elements where logic is required to determine quite how to handle them; e.g. CodeableConcepts may contain text and/or multiple codings, where the codings themselves contain a code and a system. For composite search parameters, the outcome of the expression must a collection of base elements from which the composites are derived.
	Expression *string `bson:"expression,omitempty" json:"expression,omitempty" xml:"expression,omitempty"`
	// How the search parameter relates to the set of elements returned by evaluating the expression query.
	ProcessingMode *SearchProcessingModeType `bson:"processingMode,omitempty" json:"processingMode,omitempty" xml:"processingMode,omitempty"`
	// FHIRPath expression that defines/sets a complex constraint for when this SearchParameter is applicable.
	Constraint *string `bson:"constraint,omitempty" json:"constraint,omitempty" xml:"constraint,omitempty"`
	// Types of resource (if a resource is referenced).
	//
	// Comment: In some cases, targets may exist for a search parameter that do not exist for the specified FHIRPath for all of the resources identified in SearchParameter.base.  For example, the core clinical-encounter search parameter allows both Encounter and EpisodeOfCare as targets, even though not all of the base resources have EpisodeOfCare as a valid target.
	//
	// However, the list of targets SHOULD cover all targets that might appear that are permitted by the specified FHIRPath.
	Target []string `bson:"target,omitempty" json:"target,omitempty" xml:"target,omitempty"`
	// Whether multiple values are allowed for each time the parameter exists.
	//
	// Values are separated by commas, and the parameter matches if any of the values match.
	MultipleOr *bool `bson:"multipleOr,omitempty" json:"multipleOr,omitempty" xml:"multipleOr,omitempty"`
	// Whether multiple parameters are allowed - e.g.
	//
	// more than one parameter with the same name. The search matches if all the parameters match.
	MultipleAnd *bool `bson:"multipleAnd,omitempty" json:"multipleAnd,omitempty" xml:"multipleAnd,omitempty"`
	// Comparators supported for the search parameter.
	//
	// Comment: If no comparators are listed, clients should not expect servers to support any comparators.
	Comparator []SearchComparator `bson:"comparator,omitempty" json:"comparator,omitempty" xml:"comparator,omitempty"`
	// A modifier supported for the search parameter.
	Modifier []SearchModifierCode `bson:"modifier,omitempty" json:"modifier,omitempty" xml:"modifier,omitempty"`
	// Contains the names of any search parameters which may be chained to the containing search parameter.
	//
	// Chained parameters may be added to search parameters of type reference and specify that resources will only be returned if they contain a reference to a resource which matches the chained parameter value. Values for this field should be drawn from SearchParameter.code for a parameter on the target resource type.
	//
	// Comment: Notes:
	//
	// * Not all systems will declare chain values.  If no chain values are specified at all, then no presumptions can be made about the degree of chaining support available
	//
	// * A server may use a single chain value of '*' to mean that all reference search types supported can be chained rather than enumerating all chain values
	//
	// *  If a particular reference search parameter is supported, but not subsequent chaining, then just list that search parameter name - e.g. for Patient.practitioner, the chain value would be ```practitioner```
	//
	// * If a particular reference search parameter is supported with unlimited chaining beneath it, then follow the search parameter name with '.'.  E.g. ```practitioner.```
	//
	// * If there are specific (not unlimited) chaining paths allowed they can be explicitly enumerated.  E.g. ```practitioner.organization```
	Chain []string `bson:"chain,omitempty" json:"chain,omitempty" xml:"chain,omitempty"`
	// Used to define the parts of a composite search parameter.
	Component []SearchParameterComponent `bson:"component,omitempty" json:"component,omitempty" xml:"component,omitempty"`
}

type SearchParameterComponent struct {
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
	// The definition of the search parameter that describes this part.
	Definition string `bson:"definition" json:"definition" xml:"definition"`
	// A sub-expression that defines how to extract values for this component from the output of the main SearchParameter.expression.
	//
	// Comment: This expression overrides the expression in the definition and extracts the index values from the outcome of the composite expression.
	Expression string `bson:"expression" json:"expression" xml:"expression"`
}

func (r SearchParameter) ResourceType() string {
	return "SearchParameter"
}

// MarshalJSON implements json.Marshaler
func (r SearchParameter) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		SearchParameter
		ResourceType string `json:"resourceType"`
	}{
		ResourceType:    r.ResourceType(),
		SearchParameter: r,
	})
}
