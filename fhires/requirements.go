package fhires

import v2 "encoding/json/v2"

// The Requirements resource is used to describe an actor - a human or an application that plays a role in data exchange, and that may have obligations associated with the role the actor plays.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Requirements
type Requirements struct {
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
	// An absolute URI that is used to identify this Requirements when it is referenced in a specification, model, design or an instance; also called its canonical identifier.
	//
	// This SHOULD be globally unique and SHOULD be a literal address at which an authoritative instance of this Requirements is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the Requirements is stored on different servers.
	//
	// Comment: Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred.  Multiple instances may share the same URL if they have a distinct version.
	//
	// The determination of when to create a new version of a resource (same url, new version) vs. defining a new artifact is up to the author.  Considerations for making this decision are found in [Technical and Business Versions](resource.html#versions).
	//
	// In some cases, the resource can no longer be found at the stated url, but the url itself cannot change. Implementations can use the [meta.source](resource.html#meta) element to indicate where the current master source of the resource can be found.
	URL *string `bson:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
	// A formal identifier that is used to identify this Requirements when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The identifier that is used to identify this version of the Requirements when it is referenced in a specification, model, design or instance.
	//
	// This is an arbitrary value managed by the Requirements author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	//
	// Comment: There may be different Requirements instances that have the same identifier but different versions.  The version can be appended to the url in a reference to allow a reference to a particular business version of the Requirements with the format [url]|[version]. The version SHOULD NOT contain a '#' - see [Business Version](resource.html#bv-format).
	Version *string `bson:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
	// Indicates the mechanism used to compare versions to determine which is more current.
	//
	// Comment: If set as a string, this is a FHIRPath expression that has two additional context variables passed in - %version1 and %version2 and will return a negative number if version1 is newer, a positive number if version2 and a 0 if the version ordering can't be successfully be determined.
	VersionAlgorithmString *string `bson:"versionAlgorithmString,omitempty" json:"versionAlgorithmString,omitempty" xml:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding `bson:"versionAlgorithmCoding,omitempty" json:"versionAlgorithmCoding,omitempty" xml:"versionAlgorithmCoding,omitempty"`
	// A natural language name identifying the Requirements.
	//
	// This name should be usable as an identifier for the module by machine processing applications such as code generation.
	//
	// Comment: The name is not expected to be globally unique. The name should be a simple alphanumeric type name to ensure that it is machine-processing friendly.
	Name *string `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// A short, descriptive, user-friendly title for the Requirements.
	//
	// Comment: This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc.
	Title *string `bson:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// The status of this Requirements. Enables tracking the life-cycle of the content.
	//
	// Comment: Allows filtering of Requirements that are appropriate for use versus not.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Status PublicationStatus `bson:"status" json:"status" xml:"status"`
	// A Boolean value to indicate that this Requirements is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	//
	// Comment: Allows filtering of Requirements that are appropriate for use versus not.
	Experimental *bool `bson:"experimental,omitempty" json:"experimental,omitempty" xml:"experimental,omitempty"`
	// The date  (and optionally time) when the Requirements was published.
	//
	// The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the Requirements changes.
	//
	// Comment: Note that this is not the same as the resource last-modified-date, since the resource may be a secondary representation of the Requirements. Additional specific dates may be added as extensions or be found by consulting Provenances associated with past versions of the resource.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Date *string `bson:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// The name of the organization or individual responsible for the release and ongoing maintenance of the Requirements.
	//
	// Comment: Usually an organization but may be an individual. The publisher (or steward) of the Requirements is the organization or individual primarily responsible for the maintenance and upkeep of the Requirements. This is not necessarily the same individual or organization that developed and initially authored the content. The publisher is the primary point of contact for questions or issues with the Requirements. This item SHOULD be populated unless the information is available from context.
	Publisher *string `bson:"publisher,omitempty" json:"publisher,omitempty" xml:"publisher,omitempty"`
	// Contact details to assist a user in finding and communicating with the publisher.
	//
	// Comment: May be a web site, an email address, a telephone number, etc.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty" xml:"contact,omitempty"`
	// A free text natural language description of the requirements.
	//
	// Comment: This description can be used for a general description of the requirements, and which covers why it was defined.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// The content was developed with a focus and intent of supporting the contexts that are listed.
	//
	// These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate Requirements instances.
	//
	// Comment: When multiple useContexts are specified, there is no expectation that all or any of the contexts apply.
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty" xml:"useContext,omitempty"`
	// A legal or geographic region in which the Requirements is intended to be used.
	//
	// Comment: It may be possible for the Requirements to be used in jurisdictions other than those for which it was originally designed or intended
	//
	// DEPRECATION NOTE: For consistency, implementations are encouraged to migrate to using the new 'jurisdiction' code in the useContext element.  (I.e. useContext.code indicating http://terminology.hl7.org/CodeSystem/usage-context-type#jurisdiction and useContext.valueCodeableConcept indicating the jurisdiction.).
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty" xml:"jurisdiction,omitempty"`
	// Explanation of why this Requirements is needed and why it has been designed as it has.
	//
	// Comment: This element does not describe the usage of the Requirements. Instead, it provides traceability of ''why'' the resource is either needed or ''why'' it is defined as it is.  This may be used to point to source materials or specifications that drove the structure of this Requirements.
	Purpose *string `bson:"purpose,omitempty" json:"purpose,omitempty" xml:"purpose,omitempty"`
	// A copyright statement relating to the Requirements and/or its contents.
	//
	// Copyright statements are generally legal restrictions on the use and publishing of the Requirements.
	//
	// Comment: ...
	Copyright *string `bson:"copyright,omitempty" json:"copyright,omitempty" xml:"copyright,omitempty"`
	// A short string (<50 characters), suitable for inclusion in a page footer that identifies the copyright holder, effective period, and optionally whether rights are resctricted.
	//
	// (e.g. 'All rights reserved', 'Some rights reserved').
	//
	// Comment: The (c) symbol should NOT be included in this string. It will be added by software when rendering the notation. Full details about licensing, restrictions, warrantees, etc. goes in the more general 'copyright' element.
	CopyrightLabel *string `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty" xml:"copyrightLabel,omitempty"`
	// Another set of Requirements that this set of Requirements builds on and updates.
	//
	// Comment: Existing statements (by key) may be narrowed or clarified, and additional statements added.
	DerivedFrom []string `bson:"derivedFrom,omitempty" json:"derivedFrom,omitempty" xml:"derivedFrom,omitempty"`
	// A reference to another artifact that created this set of requirements.
	//
	// This could be a Profile, etc., or external regulation, or business requirements expressed elsewhere.
	Reference []string `bson:"reference,omitempty" json:"reference,omitempty" xml:"reference,omitempty"`
	// An actor these requirements are in regard to.
	//
	// Comment: If more than one actor is specified, then it's up to the statements to describe in narrative if they don't apply to all actors.
	Actor []string `bson:"actor,omitempty" json:"actor,omitempty" xml:"actor,omitempty"`
	// The actual statement of requirement, in markdown format.
	Statement []RequirementsStatement `bson:"statement,omitempty" json:"statement,omitempty" xml:"statement,omitempty"`
}

type RequirementsStatement struct {
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
	// Key that identifies this statement (unique within this resource).
	//
	// Comment: Refer directly to the statement by {url}#{key}
	Key string `bson:"key" json:"key" xml:"key"`
	// A short human usable label for this statement.
	//
	// Comment: The key is intended for computers, while this is intended for humans. In its absence, the key should be convenient for a human to say (e.g. not a UUID)
	Label *string `bson:"label,omitempty" json:"label,omitempty" xml:"label,omitempty"`
	// A short human usable label for this statement.
	//
	// Comment: The conformance code is extracted from the requirement to make indexing and display easier. The requirement needs to express the conformance verbs directly in the markdown content. It's not unusual to mix verbs in a single sentence (e.g. System SHALL do X and SHOULD do Y)
	Conformance []ConformanceExpectation `bson:"conformance,omitempty" json:"conformance,omitempty" xml:"conformance,omitempty"`
	// This boolean flag is set to true of the text of the requirement is conditional on something e.g.
	//
	// it includes lanauage like 'if x then y'. This conditionality flag is introduced for purposes of filtering and colour highlighting etc.
	//
	// Comment: There is no intent that the conditional statement by evaluatable
	Conditionality *bool `bson:"conditionality,omitempty" json:"conditionality,omitempty" xml:"conditionality,omitempty"`
	// The actual requirement for human consumption.
	Requirement string `bson:"requirement" json:"requirement" xml:"requirement"`
	// Another statement on one of the requirements that this requirement clarifies or restricts.
	//
	// Comment: This can be just a key, if there is only one matching statement in the requirements statements referenced in Requirements.derivedFrom. If there is more than one, it should be in the format ```url#key```
	DerivedFrom *string `bson:"derivedFrom,omitempty" json:"derivedFrom,omitempty" xml:"derivedFrom,omitempty"`
	// A larger requirement that this requirement helps to refine and enable.
	//
	// Comment: Follows the same convention as statement.derivedFrom, where a reference is the canonical of the Requirements followed by '#' and the 'key' of the specific requirement.
	Parent *string `bson:"parent,omitempty" json:"parent,omitempty" xml:"parent,omitempty"`
	// A reference to another artifact that satisfies this requirement.
	//
	// This could be a Profile, extension, or an element in one of those, or a CapabilityStatement, OperationDefinition, SearchParameter, CodeSystem(/code), ValueSet, Libary etc.
	//
	// Comment: This is for use when the requirement is met be a pre-existing artifact e.g. a rule that's met by the base FHIR spec, or a national implementation guide.
	SatisfiedBy []string `bson:"satisfiedBy,omitempty" json:"satisfiedBy,omitempty" xml:"satisfiedBy,omitempty"`
	// A reference to another artifact that created this requirement.
	//
	// This could be a Profile, etc., or external regulation, or business requirements expressed elsewhere.
	Reference []string `bson:"reference,omitempty" json:"reference,omitempty" xml:"reference,omitempty"`
	// Who asked for this statement to be a requirement.
	//
	// By default, it's assumed that the publisher knows who it is if it matters.
	Source []Reference `bson:"source,omitempty" json:"source,omitempty" xml:"source,omitempty"`
}

func (r Requirements) ResourceType() string {
	return "Requirements"
}

// MarshalJSON implements json.Marshaler
func (r Requirements) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		Requirements
		ResourceType string `json:"resourceType"`
	}{
		Requirements: r,
		ResourceType: r.ResourceType(),
	})
}
