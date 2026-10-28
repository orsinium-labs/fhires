package fhires

import v2 "encoding/json/v2"

// A formal computable definition of an operation (on the RESTful interface) or a named query (using the search interaction).
//
// Status: active
//
// http://hl7.org/fhir/StructureDefinition/OperationDefinition
type OperationDefinition struct {
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
	// An absolute URI that is used to identify this operation definition when it is referenced in a specification, model, design or an instance; also called its canonical identifier.
	//
	// This SHOULD be globally unique and SHOULD be a literal address at which an authoritative instance of this operation definition is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the operation definition is stored on different servers.
	//
	// Comment: Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred.  Multiple instances may share the same URL if they have a distinct version.
	//
	// The determination of when to create a new version of a resource (same url, new version) vs. defining a new artifact is up to the author.  Considerations for making this decision are found in [Technical and Business Versions](resource.html#versions).
	//
	// In some cases, the resource can no longer be found at the stated url, but the url itself cannot change. Implementations can use the [meta.source](resource.html#meta) element to indicate where the current master source of the resource can be found.
	URL *string `bson:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
	// A formal identifier that is used to identify this implementation guide when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The identifier that is used to identify this version of the operation definition when it is referenced in a specification, model, design or instance.
	//
	// This is an arbitrary value managed by the operation definition author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	//
	// Comment: There may be different operation definitions that have the same url but different versions.  The version can be appended to the url in a reference to allow a reference to a particular business version of the operation definition with the format. The version SHOULD NOT contain a '#' - see [Business Version](resource.html#bv-format).
	Version *string `bson:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
	// Indicates the mechanism used to compare versions to determine which is more current.
	//
	// Comment: If set as a string, this is a FHIRPath expression that has two additional context variables passed in - %version1 and %version2 and will return a negative number if version1 is newer, a positive number if version2 and a 0 if the version ordering can't be successfully be determined.
	VersionAlgorithmString *string `bson:"versionAlgorithmString,omitempty" json:"versionAlgorithmString,omitempty" xml:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding `bson:"versionAlgorithmCoding,omitempty" json:"versionAlgorithmCoding,omitempty" xml:"versionAlgorithmCoding,omitempty"`
	// A natural language name identifying the operation definition.
	//
	// This name should be usable as an identifier for the module by machine processing applications such as code generation.
	//
	// Comment: The name is not expected to be globally unique. The name should be a simple alphanumeric type no-whitespace name to ensure that it is machine-processing friendly.
	Name string `bson:"name" json:"name" xml:"name"`
	// A short, descriptive, user-friendly title for the operation definition.
	//
	// Comment: This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc.
	Title *string `bson:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// The current state of this operation definition.
	//
	// Comment: A nominal state-transition diagram can be found in the] documentation
	//
	// Unknown does not represent 'other' - one of the defined statuses must apply.  Unknown is used when the authoring system is not sure what the current status is.
	Status PublicationStatus `bson:"status" json:"status" xml:"status"`
	// Whether this is an operation or a named query.
	//
	// Comment: Named queries are invoked differently, and have different capabilities.
	Kind OperationKind `bson:"kind" json:"kind" xml:"kind"`
	// A Boolean value to indicate that this operation definition is authored for testing purposes (or education/evaluation/marketing) and is not intended for genuine usage.
	//
	// Comment: Allows filtering of operation definitions that are appropriate for use versus not.
	Experimental *bool `bson:"experimental,omitempty" json:"experimental,omitempty" xml:"experimental,omitempty"`
	// The date  (and optionally time) when the operation definition was last significantly changed.
	//
	// The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the operation definition changes.
	//
	// Comment: The date is often not tracked until the resource is published, but may be present on draft content. Note that this is not the same as the resource last-modified-date, since the resource may be a secondary representation of the operation definition. Additional specific dates may be added as extensions or be found by consulting Provenances associated with past versions of the resource.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Date *string `bson:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// The name of the organization or individual responsible for the release and ongoing maintenance of the operation definition.
	//
	// Comment: Usually an organization but may be an individual. The publisher (or steward) of the operation definition is the organization or individual primarily responsible for the maintenance and upkeep of the operation definition. This is not necessarily the same individual or organization that developed and initially authored the content. The publisher is the primary point of contact for questions or issues with the operation definition. This item SHOULD be populated unless the information is available from context.
	Publisher *string `bson:"publisher,omitempty" json:"publisher,omitempty" xml:"publisher,omitempty"`
	// Contact details to assist a user in finding and communicating with the publisher.
	//
	// Comment: May be a web site, an email address, a telephone number, etc.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty" xml:"contact,omitempty"`
	// A free text natural language description of the operation definition from a consumer's perspective.
	//
	// Comment: This description can be used to capture details such as comments about misuse, instructions for clinical use and interpretation, literature references, examples from the paper world, etc. It is not a rendering of the operation definition as conveyed in the 'text' field of the resource itself. This item SHOULD be populated unless the information is available from context (e.g. the language of the operation definition is presumed to be the predominant language in the place the operation definition was created).
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// The content was developed with a focus and intent of supporting the contexts that are listed.
	//
	// These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate operation definition.
	//
	// Comment: When multiple useContexts are specified, there is no expectation that all or even any of the contexts apply.
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty" xml:"useContext,omitempty"`
	// A legal or geographic region in which the operation definition is intended to be used.
	//
	// Comment: It may be possible for the operation definition to be used in jurisdictions other than those for which it was originally designed or intended.
	//
	// DEPRECATION NOTE: For consistency, implementations are encouraged to migrate to using the new 'jurisdiction' code in the useContext element.  (I.e. useContext.code indicating http://terminology.hl7.org/CodeSystem/usage-context-type#jurisdiction and useContext.valueCodeableConcept indicating the jurisdiction.)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty" xml:"jurisdiction,omitempty"`
	// Explanation of why this operation definition is needed and why it has been designed as it has.
	//
	// Comment: This element does not describe the usage of the operation definition. Instead, it provides traceability of ''why'' the resource is either needed or ''why'' it is defined as it is.  This may be used to point to source materials or specifications that drove the structure of this operation definition.
	Purpose *string `bson:"purpose,omitempty" json:"purpose,omitempty" xml:"purpose,omitempty"`
	// A copyright statement relating to the operation definition and/or its contents.
	//
	// Copyright statements are generally legal restrictions on the use and publishing of the operation definition.
	Copyright *string `bson:"copyright,omitempty" json:"copyright,omitempty" xml:"copyright,omitempty"`
	// A short string (<50 characters), suitable for inclusion in a page footer that identifies the copyright holder, effective period, and optionally whether rights are resctricted.
	//
	// (e.g. 'All rights reserved', 'Some rights reserved').
	//
	// Comment: The (c) symbol should NOT be included in this string. It will be added by software when rendering the notation. Full details about licensing, restrictions, warrantees, etc. goes in the more general 'copyright' element.
	CopyrightLabel *string `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty" xml:"copyrightLabel,omitempty"`
	// Whether the operation affects state.
	//
	// Side effects such as producing audit trail entries do not count as 'affecting  state'.
	//
	// Comment: What http methods can be used for the operation depends on the .affectsState value and whether the input parameters are primitive or complex:
	//
	// 1. Servers SHALL support POST method for all operations.
	//
	// 2. Servers SHALL support the GET method when the operation has affectsState = false and all required parameters for the operation are primitive.
	AffectsState *bool `bson:"affectsState,omitempty" json:"affectsState,omitempty" xml:"affectsState,omitempty"`
	// The label that is recommended to be used in the URL for this operation.
	//
	// In some cases, servers may need to use a different CapabilityStatement operation.name to differentiate between multiple SearchParameters that happen to have the same code.
	//
	// Comment: For maximum compatibility, use only lowercase ASCII characters. Note that HL7 will never define multiple operations with the same code. See [Naming Rules & Guidelines](https://confluence.hl7.org/display/FHIR/Guide+to+Designing+Resources#GuidetoDesigningResources-NamingRules&Guidelines) for the internal HL7 rules around codes, which are useful additional advice for all implementers
	Code string `bson:"code" json:"code" xml:"code"`
	// Additional information about how to use this operation or named query.
	Comment *string `bson:"comment,omitempty" json:"comment,omitempty" xml:"comment,omitempty"`
	// Indicates that this operation definition is a constraining profile on the base.
	//
	// Comment: A constrained profile can make optional parameters required or not used, introduce new parameters (required or optional), and clarify documentation.
	Base *string `bson:"base,omitempty" json:"base,omitempty" xml:"base,omitempty"`
	// The types on which this operation can be executed.
	//
	// Comment: If the type is an abstract resource ("Resource" or "DomainResource") then the operation can be invoked on any concrete specialization.
	Resource []string `bson:"resource,omitempty" json:"resource,omitempty" xml:"resource,omitempty"`
	// Indicates whether this operation or named query can be invoked at the system level (e.g.
	//
	// without needing to choose a resource type for the context).
	System bool `bson:"system" json:"system" xml:"system"`
	// Indicates whether this operation or named query can be invoked at the resource type level for any given resource type level (e.g.
	//
	// without needing to choose a specific resource id for the context).
	Type bool `bson:"type" json:"type" xml:"type"`
	// Indicates whether this operation can be invoked on a particular instance of one of the given types.
	Instance bool `bson:"instance" json:"instance" xml:"instance"`
	// Additional validation information for the in parameters - a single profile that covers all the parameters.
	//
	// The profile is a constraint on the parameters resource as a whole.
	//
	// Comment: If present the profile shall not conflict with what is specified in the parameters in the operation definition (max/min etc.), though it may provide additional constraints. The constraints expressed in the profile apply whether the operation is invoked by a POST wih parameters or not.
	InputProfile *string `bson:"inputProfile,omitempty" json:"inputProfile,omitempty" xml:"inputProfile,omitempty"`
	// Additional validation information for the out parameters - a single profile that covers all the parameters.
	//
	// The profile is a constraint on the parameters resource.
	//
	// Comment: If present the profile shall not conflict with what is specified in the parameters in the operation definition (max/min etc.), though it may provide additional constraints. The constraints expressed in the profile apply whether the operation is invoked by a POST wih parameters or not.
	OutputProfile *string `bson:"outputProfile,omitempty" json:"outputProfile,omitempty" xml:"outputProfile,omitempty"`
	// The parameters for the operation/query.
	//
	// Comment: Query Definitions only have one output parameter, named "result". This might not be described, but can be to allow a profile to be defined.
	Parameter []OperationDefinitionParameter `bson:"parameter,omitempty" json:"parameter,omitempty" xml:"parameter,omitempty"`
	// Defines an appropriate combination of parameters to use when invoking this operation, to help code generators when generating overloaded parameter sets for this operation.
	//
	// Comment: The combinations are suggestions as to which sets of parameters to use together, but the combinations are not intended to be authoritative.
	Overload []OperationDefinitionOverload `bson:"overload,omitempty" json:"overload,omitempty" xml:"overload,omitempty"`
}

type OperationDefinitionParameter struct {
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
	// The name of used to identify the parameter.
	//
	// Comment: This name must be a token (start with a letter in a..z, and only contain letters, numerals, and underscore. Note that for search parameters (type = string, with a search type), the name may be altered by the search modifiers.
	Name string `bson:"name" json:"name" xml:"name"`
	// Whether this is an input or an output parameter.
	//
	// Comment: If a parameter name is used for both an input and an output parameter, the parameter should be defined twice.
	Use OperationParameterUse `bson:"use" json:"use" xml:"use"`
	// If present, indicates that the parameter applies when the operation is being invoked at the specified level.
	Scope []OperationParameterScope `bson:"scope,omitempty" json:"scope,omitempty" xml:"scope,omitempty"`
	// The minimum number of times this parameter SHALL appear in the request or response.
	Min int `bson:"min" json:"min" xml:"min"`
	// The maximum number of times this element is permitted to appear in the request or response.
	Max string `bson:"max" json:"max" xml:"max"`
	// Describes the meaning or use of this parameter.
	Documentation *string `bson:"documentation,omitempty" json:"documentation,omitempty" xml:"documentation,omitempty"`
	// The type for this parameter.
	//
	// Comment: if there is no stated parameter, then the parameter is a multi-part parameter type and must have at least one part defined.
	Type *FHIRTypes `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Support for polymorphic types.
	//
	// If the parameter type is abstract, this element lists allowed sub-types for the parameter.
	//
	// Comment: In previous versions of FHIR, there was an extension for this: http://hl7.org/fhir/StructureDefinition/operationdefinition-allowed-type
	AllowedType []FHIRTypes `bson:"allowedType,omitempty" json:"allowedType,omitempty" xml:"allowedType,omitempty"`
	// Used when the type is "Reference" or "canonical", and identifies a profile structure or implementation Guide that applies to the target of the reference this parameter refers to.
	//
	// If any profiles are specified, then the content must conform to at least one of them. The URL can be a local reference - to a contained StructureDefinition, or a reference to another StructureDefinition or Implementation Guide by a canonical URL. When an implementation guide is specified, the target resource SHALL conform to at least one profile defined in the implementation guide.
	//
	// Comment: Often, these profiles are the base definitions from the spec (e.g. http://hl7.org/fhir/StructureDefinition/Patient).
	TargetProfile []string `bson:"targetProfile,omitempty" json:"targetProfile,omitempty" xml:"targetProfile,omitempty"`
	// How the parameter is understood if/when it used as search parameter.
	//
	// This is only used if the parameter is a string.
	//
	// Comment: Parameters that are search parameter usually - but not always - are also defined for use with general searches, and function in the operation to filter the set of resources in scope for the operation just like the do when searching. The searchType describes how the server processes the value. For an example, see the [Evaluate Measure Operation](measure-operation-evaluate-measure.html)
	SearchType *SearchParamType `bson:"searchType,omitempty" json:"searchType,omitempty" xml:"searchType,omitempty"`
	// Binds to a value set if this parameter is coded (code, Coding, CodeableConcept).
	Binding *OperationDefinitionParameterBinding `bson:"binding,omitempty" json:"binding,omitempty" xml:"binding,omitempty"`
	// Identifies other resource parameters within the operation invocation that are expected to resolve to this resource.
	//
	// Comment: Resolution applies if the referenced parameter exists.
	ReferencedFrom []OperationDefinitionParameterReferencedFrom `bson:"referencedFrom,omitempty" json:"referencedFrom,omitempty" xml:"referencedFrom,omitempty"`
	// The parts of a nested Parameter.
	//
	// Comment: Query Definitions only have one output parameter, named "result". This might not be described, but can be to allow a profile to be defined.
	Part []OperationDefinitionParameter `bson:"part,omitempty" json:"part,omitempty"`
}

type OperationDefinitionParameterBinding struct {
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
	// Indicates the degree of conformance expectations associated with this binding - that is, the degree to which the provided value set must be adhered to in the instances.
	//
	// Comment: For further discussion, see [Using Terminologies](terminologies.html).
	Strength BindingStrength `bson:"strength" json:"strength" xml:"strength"`
	// Points to the value set or external definition (e.g.
	//
	// implicit value set) that identifies the set of codes to be used.
	//
	// Comment: For value sets with a referenceResource, the display can contain the value set description.  The reference may be version-specific or not.
	ValueSet string `bson:"valueSet" json:"valueSet" xml:"valueSet"`
}

type OperationDefinitionParameterReferencedFrom struct {
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
	// The name of the parameter or dot-separated path of parameter names pointing to the resource parameter that is expected to contain a reference to this resource.
	Source string `bson:"source" json:"source" xml:"source"`
	// The id of the element in the referencing resource that is expected to resolve to this resource.
	SourceID *string `bson:"sourceId,omitempty" json:"sourceId,omitempty" xml:"sourceId,omitempty"`
}

type OperationDefinitionOverload struct {
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
	// Name of parameter to include in overload.
	ParameterName []string `bson:"parameterName,omitempty" json:"parameterName,omitempty" xml:"parameterName,omitempty"`
	// Comments to go on overload.
	Comment *string `bson:"comment,omitempty" json:"comment,omitempty" xml:"comment,omitempty"`
}

func (r OperationDefinition) ResourceType() string {
	return "OperationDefinition"
}

// MarshalJSON implements json.Marshaler
func (r OperationDefinition) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		OperationDefinition
		ResourceType string `json:"resourceType"`
	}{
		OperationDefinition: r,
		ResourceType:        r.ResourceType(),
	})
}
