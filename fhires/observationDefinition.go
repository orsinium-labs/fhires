package fhires

import v2 "encoding/json/v2"

// Set of definitional characteristics for a kind of observation or measurement produced or consumed by an orderable health care service.
//
// Purpose: In a catalog of health-related services that use or produce observations and measurements, this resource describes the expected characteristics of these observation / measurements.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/ObservationDefinition
type ObservationDefinition struct {
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
	// An absolute URL that is used to identify this ObservationDefinition when it is referenced in a specification, model, design or an instance.
	//
	// This SHALL be a URL, SHOULD be globally unique, and SHOULD be an address at which this ObservationDefinition is (or will be) published. The URL SHOULD include the major version of the ObservationDefinition. For more information see Technical and Business Versions.
	//
	// Comment: Can be a urn:uuid: or a urn:oid:, but real http: addresses are preferred. This is the URI that will be used when making canonical references to this resource.
	URL *string `bson:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
	// Business identifiers assigned to this ObservationDefinition.
	//
	// by the performer and/or other systems. These identifiers remain constant as the resource is updated and propagates from server to server.
	//
	// Comment: This is a business identifier, not a resource identifier. It is best practice for the identifier to only appear on a single resource instance, however business practices may occasionally dictate that multiple resource instances with the same identifier can exist - possibly even with different resource types.
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The identifier that is used to identify this version of the ObservationDefinition when it is referenced in a specification, model, design or instance.
	//
	// This is an arbitrary value managed by the ObservationDefinition author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions are orderable.
	//
	// Comment: There may be multiple different instances of an observationDefinition that have the same identifier but different versions.
	Version *string `bson:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
	// Indicates the mechanism used to compare versions to determine which is more current.
	//
	// Comment: If set as a string, this is a FHIRPath expression that has two additional context variables passed in - %version1 and %version2 and will return a negative number if version1 is newer, a positive number if version2 and a 0 if the version ordering can't be successfully be determined.
	VersionAlgorithmString *string `bson:"versionAlgorithmString,omitempty" json:"versionAlgorithmString,omitempty" xml:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding `bson:"versionAlgorithmCoding,omitempty" json:"versionAlgorithmCoding,omitempty" xml:"versionAlgorithmCoding,omitempty"`
	// A natural language name identifying the ObservationDefinition.
	//
	// This name should be usable as an identifier for the module by machine processing applications such as code generation.
	//
	// Comment: The name is not expected to be globally unique. The name should be a simple alphanumeric type name to ensure that it is machine-processing friendly.
	Name *string `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// A short, descriptive, user-friendly title for the ObservationDefinition.
	Title *string `bson:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// The current state of the ObservationDefinition.
	//
	// Comment: A nominal state-transition diagram can be found in the [Definition pattern](definition.html#statemachine) documentation
	//
	// Unknown does not represent "other" - one of the defined statuses must apply. Unknown is used when the authoring system is not sure what the current status is.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Status PublicationStatus `bson:"status" json:"status" xml:"status"`
	// A flag to indicate that this ObservationDefinition is authored for testing purposes (or education/evaluation/marketing), and is not intended to be used for genuine usage.
	//
	// Comment: Allows filtering of ObservationDefinition that are appropriate for use vs. not.
	Experimental *bool `bson:"experimental,omitempty" json:"experimental,omitempty" xml:"experimental,omitempty"`
	// The date (and optionally time) when the ObservationDefinition was last significantly changed.
	//
	// The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the ObservationDefinition changes.
	//
	// Comment: The date is often not tracked until the resource is published, but may be present on draft content. Note that this is not the same as the resource last-modified-date, since the resource may be a secondary representation of the ObservationDefinition. Additional specific dates may be added as extensions or be found by consulting Provenances associated with past versions of the resource.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Date *string `bson:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// Helps establish the "authority/credibility" of the ObservationDefinition.
	//
	// May also allow for contact.
	Publisher *string `bson:"publisher,omitempty" json:"publisher,omitempty" xml:"publisher,omitempty"`
	// Contact details to assist a user in finding and communicating with the publisher.
	//
	// Comment: May be a web site, an email address, a telephone number, etc.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty" xml:"contact,omitempty"`
	// A free text natural language description of the ObservationDefinition from the consumer's perspective.
	//
	// Comment: This description can be used to capture details such as comments about misuse, instructions for clinical use and interpretation, literature references, examples from the paper world, etc. It is not a rendering of the module as conveyed in the text field of the resource itself. This item SHOULD be populated unless the information is available from context.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// The content was developed with a focus and intent of supporting the contexts that are listed.
	//
	// These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate ObservationDefinition instances.
	//
	// Comment: When multiple usageContexts are specified, there is no expectation for whether all or any of the contexts apply.
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty" xml:"useContext,omitempty"`
	// A jurisdiction in which the ObservationDefinition is intended to be used.
	//
	// Comment: It may be possible for the ObservationDefinition to be used in jurisdictions other than those for which it was originally designed or intended.
	//
	// DEPRECATION NOTE: For consistency, implementations are encouraged to migrate to using the new 'jurisdiction' code in the useContext element.  (I.e. useContext.code indicating http://terminology.hl7.org/CodeSystem/usage-context-type#jurisdiction and useContext.valueCodeableConcept indicating the jurisdiction.)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty" xml:"jurisdiction,omitempty"`
	// Explains why this ObservationDefinition is needed and why it has been designed as it has.
	//
	// Comment: This element does not describe the usage of the ObservationDefinition. Rather it is for traceability of ''why'' the resource is either needed or ''why'' it is defined as it is. This may be used to point to source materials or specifications that drove the structure of this ObservationDefinition.
	Purpose *string `bson:"purpose,omitempty" json:"purpose,omitempty" xml:"purpose,omitempty"`
	// Copyright statement relating to the ObservationDefinition and/or its contents.
	//
	// Copyright statements are generally legal restrictions on the use and publishing of the ObservationDefinition.
	Copyright *string `bson:"copyright,omitempty" json:"copyright,omitempty" xml:"copyright,omitempty"`
	// A short string (<50 characters), suitable for inclusion in a page footer that identifies the copyright holder, effective period, and optionally whether rights are resctricted.
	//
	// (e.g. 'All rights reserved', 'Some rights reserved').
	//
	// Comment: The (c) symbol should NOT be included in this string. It will be added by software when rendering the notation. Full details about licensing, restrictions, warrantees, etc. goes in the more general 'copyright' element.
	CopyrightLabel *string `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty" xml:"copyrightLabel,omitempty"`
	// The date on which the asset content was approved by the publisher.
	//
	// Approval happens once when the content is officially approved for usage.
	//
	// Comment: The date may be more recent than the approval date because of minor changes / editorial corrections.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	ApprovalDate *string `bson:"approvalDate,omitempty" json:"approvalDate,omitempty" xml:"approvalDate,omitempty"`
	// The date on which the asset content was last reviewed.
	//
	// Review happens periodically after that, but doesn't change the original approval date.
	//
	// Comment: If specified, this is usually after the approval date.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	LastReviewDate *string `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty" xml:"lastReviewDate,omitempty"`
	// The period during which the ObservationDefinition content was or is planned to be effective.
	//
	// Comment: The effective period for an ObservationDefinition determines when the content is applicable for usage and is independent of publication and review dates. For example, an observation definition intended to be used for the year 2021 might be published in 2020.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	EffectivePeriod *Period `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty" xml:"effectivePeriod,omitempty"`
	// The canonical URL pointing to another FHIR-defined ObservationDefinition that is adhered to in whole or in part by this definition.
	DerivedFromCanonical []string `bson:"derivedFromCanonical,omitempty" json:"derivedFromCanonical,omitempty" xml:"derivedFromCanonical,omitempty"`
	// The URL pointing to an externally-defined observation definition, guideline or other definition that is adhered to in whole or in part by this definition.
	DerivedFromUri []string `bson:"derivedFromUri,omitempty" json:"derivedFromUri,omitempty" xml:"derivedFromUri,omitempty"`
	// A code that describes the intended kind of subject of Observation instances conforming to this ObservationDefinition.
	//
	// Comment: Examples: person, animal, device, air, surface ….
	Subject []CodeableConcept `bson:"subject,omitempty" json:"subject,omitempty" xml:"subject,omitempty"`
	// The type of individual/organization/device that is expected to act upon instances of this definition.
	PerformerType *CodeableConcept `bson:"performerType,omitempty" json:"performerType,omitempty" xml:"performerType,omitempty"`
	// A code that classifies the general type of observation.
	//
	// Comment: This element allows various categorization schemes based on the owner’s definition of the category and effectively multiple categories can be used for one instance of ObservationDefinition. The level of granularity is defined by the category concepts in the value set.
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
	// Describes what will be observed. Sometimes this is called the observation "name".
	Code CodeableConcept `bson:"code" json:"code" xml:"code"`
	// The data types allowed for the value element of the instance observations conforming to this ObservationDefinition.
	PermittedDataType []ObservationDataType `bson:"permittedDataType,omitempty" json:"permittedDataType,omitempty" xml:"permittedDataType,omitempty"`
	// Multiple results allowed for observations conforming to this ObservationDefinition.
	//
	// Comment: An example of observation allowing multiple results is "bacteria identified by culture". Conversely, the measurement of a potassium level allows a single result.
	MultipleResultsAllowed *bool `bson:"multipleResultsAllowed,omitempty" json:"multipleResultsAllowed,omitempty" xml:"multipleResultsAllowed,omitempty"`
	// The site on the subject's body where the  observation is to be made.
	//
	// Comment: Only used if the defined observation is to be made directly on a body part and if this body part is not implied by code found in ObservationDefinition.code. Not used for in vitro diagnostic observations.
	//
	// Example: 14975008 |Forearm structure|.
	BodySite *CodeableConcept `bson:"bodySite,omitempty" json:"bodySite,omitempty" xml:"bodySite,omitempty"`
	// The method or technique used to perform the observation.
	//
	// Comment: Only used if not implicit in observation code
	//
	// Example: 702659008 |Automated count technique|.
	Method *CodeableConcept `bson:"method,omitempty" json:"method,omitempty" xml:"method,omitempty"`
	// The kind of specimen that this type of observation is produced on.
	//
	// Comment: Only used for in vitro observations. When multiple occurrences of specimen are present, they are to be combined with a logical AND: all specified specimens  are needed for the observation.
	Specimen []Reference `bson:"specimen,omitempty" json:"specimen,omitempty" xml:"specimen,omitempty"`
	// The measurement model of device or actual device used to produce observations of this type.
	//
	// Comment: When multiple occurrences of device are present, they are to be combined with a logical OR: at least one of the specified devices is used to produce the observation.
	Device []Reference `bson:"device,omitempty" json:"device,omitempty" xml:"device,omitempty"`
	// The preferred name to be used when reporting the results of observations conforming to this ObservationDefinition.
	PreferredReportName *string `bson:"preferredReportName,omitempty" json:"preferredReportName,omitempty" xml:"preferredReportName,omitempty"`
	// Units allowed for the valueQuantity element in the instance observations conforming to this ObservationDefinition.
	PermittedUnit []Coding `bson:"permittedUnit,omitempty" json:"permittedUnit,omitempty" xml:"permittedUnit,omitempty"`
	// A set of qualified values associated with a context and a set of conditions -  provides a range for quantitative and ordinal observations and a collection of value sets for qualitative observations.
	QualifiedValue []ObservationDefinitionQualifiedValue `bson:"qualifiedValue,omitempty" json:"qualifiedValue,omitempty" xml:"qualifiedValue,omitempty"`
	// This ObservationDefinition defines a group  observation (e.g.
	//
	// a battery, a panel of tests, a set of vital sign measurements) that includes the target as a member of the group.
	HasMember []Reference `bson:"hasMember,omitempty" json:"hasMember,omitempty" xml:"hasMember,omitempty"`
	// Some observations have multiple component observations, expressed as separate code value pairs.
	Component []ObservationDefinitionComponent `bson:"component,omitempty" json:"component,omitempty" xml:"component,omitempty"`
}

type ObservationDefinitionQualifiedValue struct {
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
	// A concept defining the context for this set of qualified values.
	Context *CodeableConcept `bson:"context,omitempty" json:"context,omitempty" xml:"context,omitempty"`
	// The target population this  set of qualified values applies to.
	//
	// Comment: If this element is not present then the global population is assumed.
	AppliesTo []CodeableConcept `bson:"appliesTo,omitempty" json:"appliesTo,omitempty" xml:"appliesTo,omitempty"`
	// The gender this  set of qualified values applies to.
	Gender *AdministrativeGender `bson:"gender,omitempty" json:"gender,omitempty" xml:"gender,omitempty"`
	// The age range this  set of qualified values applies to.
	//
	// Comment: Some analytes vary greatly over age.
	Age *Range `bson:"age,omitempty" json:"age,omitempty" xml:"age,omitempty"`
	// The gestational age this  set of qualified values applies to.
	GestationalAge *Range `bson:"gestationalAge,omitempty" json:"gestationalAge,omitempty" xml:"gestationalAge,omitempty"`
	// Text based condition for which the the set of qualified values is valid.
	Condition *string `bson:"condition,omitempty" json:"condition,omitempty" xml:"condition,omitempty"`
	// The category of range of values for continuous or ordinal observations that match the criteria of this set of qualified values.
	RangeCategory *ObservationRangeCategory `bson:"rangeCategory,omitempty" json:"rangeCategory,omitempty" xml:"rangeCategory,omitempty"`
	// The range of values defined for continuous or ordinal observations that match the criteria of this set of qualified values.
	Range *Range `bson:"range,omitempty" json:"range,omitempty" xml:"range,omitempty"`
	// The set of valid coded results for qualitative observations  that match the criteria of this set of qualified values.
	ValidCodedValueSet *string `bson:"validCodedValueSet,omitempty" json:"validCodedValueSet,omitempty" xml:"validCodedValueSet,omitempty"`
	// The set of normal coded results for qualitative observations  that match the criteria of this set of qualified values.
	NormalCodedValueSet *string `bson:"normalCodedValueSet,omitempty" json:"normalCodedValueSet,omitempty" xml:"normalCodedValueSet,omitempty"`
	// The set of abnormal coded results for qualitative observations  that match the criteria of this set of qualified values.
	AbnormalCodedValueSet *string `bson:"abnormalCodedValueSet,omitempty" json:"abnormalCodedValueSet,omitempty" xml:"abnormalCodedValueSet,omitempty"`
	// The set of critical coded results for qualitative observations  that match the criteria of this set of qualified values.
	CriticalCodedValueSet *string `bson:"criticalCodedValueSet,omitempty" json:"criticalCodedValueSet,omitempty" xml:"criticalCodedValueSet,omitempty"`
}

type ObservationDefinitionComponent struct {
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
	// Describes what will be observed.
	Code CodeableConcept `bson:"code" json:"code" xml:"code"`
	// The data types allowed for the value element of the instance of this component observations.
	PermittedDataType []ObservationDataType `bson:"permittedDataType,omitempty" json:"permittedDataType,omitempty" xml:"permittedDataType,omitempty"`
	// Units allowed for the valueQuantity element in the instance observations conforming to this ObservationDefinition.
	PermittedUnit []Coding `bson:"permittedUnit,omitempty" json:"permittedUnit,omitempty" xml:"permittedUnit,omitempty"`
	// A set of qualified values associated with a context and a set of conditions -  provides a range for quantitative and ordinal observations and a collection of value sets for qualitative observations.
	QualifiedValue []ObservationDefinitionQualifiedValue `bson:"qualifiedValue,omitempty" json:"qualifiedValue,omitempty"`
}

func (r ObservationDefinition) ResourceType() string {
	return "ObservationDefinition"
}

// MarshalJSON implements json.Marshaler
func (r ObservationDefinition) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		ObservationDefinition
		ResourceType string `json:"resourceType"`
	}{
		ObservationDefinition: r,
		ResourceType:          r.ResourceType(),
	})
}
