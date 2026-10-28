package fhires

import "encoding/json/v2"

// A kind of specimen with associated set of requirements.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/SpecimenDefinition
type SpecimenDefinition struct {
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
	// An absolute URL that is used to identify this SpecimenDefinition when it is referenced in a specification, model, design or an instance.
	//
	// This SHALL be a URL, SHOULD be globally unique, and SHOULD be an address at which this SpecimenDefinition is (or will be) published. The URL SHOULD include the major version of the SpecimenDefinition. For more information see Technical and Business Versions.
	//
	// Comment: Can be a urn:uuid: or a urn:oid:, but real http: addresses are preferred. This is the URI that will be used when making canonical references to this resource.
	URL *string `bson:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
	// A business identifier assigned to this SpecimenDefinition.
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The identifier that is used to identify this version of the SpecimenDefinition when it is referenced in a specification, model, design or instance.
	//
	// This is an arbitrary value managed by the SpecimenDefinition author and is not expected to be globally unique.
	//
	// Comment: There may be multiple different instances of a SpecimenDefinition that have the same identifier but different versions.
	Version *string `bson:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
	// Indicates the mechanism used to compare versions to determine which is more current.
	//
	// Comment: If set as a string, this is a FHIRPath expression that has two additional context variables passed in - %version1 and %version2 and will return a negative number if version1 is newer, a positive number if version2 and a 0 if the version ordering can't be successfully be determined.
	VersionAlgorithmString *string `bson:"versionAlgorithmString,omitempty" json:"versionAlgorithmString,omitempty" xml:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding `bson:"versionAlgorithmCoding,omitempty" json:"versionAlgorithmCoding,omitempty" xml:"versionAlgorithmCoding,omitempty"`
	// A natural language name identifying the {{title}}.
	//
	// This name should be usable as an identifier for the module by machine processing applications such as code generation.
	//
	// Comment: The name is not expected to be globally unique. The name should be a simple alphanumeric type name to ensure that it is machine-processing friendly.
	Name *string `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// A short, descriptive, user-friendly title for the SpecimenDefinition.
	Title *string `bson:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// The canonical URL pointing to another FHIR-defined SpecimenDefinition that is adhered to in whole or in part by this definition.
	DerivedFromCanonical []string `bson:"derivedFromCanonical,omitempty" json:"derivedFromCanonical,omitempty" xml:"derivedFromCanonical,omitempty"`
	// The URL pointing to an externally-defined type of specimen, guideline or other definition that is adhered to in whole or in part by this definition.
	DerivedFromUri []string `bson:"derivedFromUri,omitempty" json:"derivedFromUri,omitempty" xml:"derivedFromUri,omitempty"`
	// The current state of theSpecimenDefinition.
	//
	// Comment: A nominal state-transition diagram can be found in the [Definition pattern](definition.html#statemachine) documentation
	//
	// Unknown does not represent "other" - one of the defined statuses must apply. Unknown is used when the authoring system is not sure what the current status is.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Status PublicationStatus `bson:"status" json:"status" xml:"status"`
	// A flag to indicate that this SpecimenDefinition is not authored for  genuine usage.
	//
	// Comment: Allows filtering of SpecimenDefinition that are appropriate for use vs. not.
	Experimental *bool `bson:"experimental,omitempty" json:"experimental,omitempty" xml:"experimental,omitempty"`
	// A code or group definition that describes the intended subject  from which this kind of specimen is to be collected.
	//
	// Comment: Examples: person, animal, device, water ….
	SubjectCodeableConcept *CodeableConcept `bson:"subjectCodeableConcept,omitempty" json:"subjectCodeableConcept,omitempty" xml:"subjectCodeableConcept,omitempty"`
	SubjectReference       *Reference       `bson:"subjectReference,omitempty" json:"subjectReference,omitempty" xml:"subjectReference,omitempty"`
	// For draft definitions, indicates the date of initial creation.
	//
	// For active definitions, represents the date of activation. For withdrawn definitions, indicates the date of withdrawal.
	//
	// Comment: See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Date *string `bson:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// Helps establish the "authority/credibility" of the SpecimenDefinition.
	//
	// May also allow for contact.
	Publisher *string `bson:"publisher,omitempty" json:"publisher,omitempty" xml:"publisher,omitempty"`
	// Contact details to assist a user in finding and communicating with the publisher.
	//
	// Comment: May be a web site, an email address, a telephone number, etc.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty" xml:"contact,omitempty"`
	// A free text natural language description of the SpecimenDefinition from the consumer's perspective.
	//
	// Comment: This description can be used to capture details such as comments about misuse, instructions for clinical use and interpretation, literature references, examples from the paper world, etc. It is not a rendering of the module as conveyed in the text field of the resource itself. This item SHOULD be populated unless the information is available from context.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// The content was developed with a focus and intent of supporting the contexts that are listed.
	//
	// These terms may be used to assist with indexing and searching of specimen definitions.
	//
	// Comment: When multiple usageContexts are specified, there is no expectation for whether all or any of the contexts apply.
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty" xml:"useContext,omitempty"`
	// A jurisdiction in which the SpecimenDefinition is intended to be used.
	//
	// Comment: DEPRECATION NOTE: For consistency, implementations are encouraged to migrate to using the new 'jurisdiction' code in the useContext element.  (I.e. useContext.code indicating http://terminology.hl7.org/CodeSystem/usage-context-type#jurisdiction and useContext.valueCodeableConcept indicating the jurisdiction.)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty" xml:"jurisdiction,omitempty"`
	// Explains why this SpecimeDefinition is needed and why it has been designed as it has.
	//
	// Comment: This element does not describe the usage of the SpecimenDefinition. Rather it is for traceability of ''why'' the resource is either needed or ''why'' it is defined as it is. This may be used to point to source materials or specifications that drove the structure of this SpecimenDefinition.
	Purpose *string `bson:"purpose,omitempty" json:"purpose,omitempty" xml:"purpose,omitempty"`
	// Copyright statement relating to the SpecimenDefinition and/or its contents.
	//
	// Copyright statements are generally legal restrictions on the use and publishing of the SpecimenDefinition.
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
	// The period during which the SpecimenDefinition content was or is planned to be effective.
	//
	// Comment: The effective period for a SpecimenDefinition determines when the content is applicable for usage and is independent of publication and review dates.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	EffectivePeriod *Period `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty" xml:"effectivePeriod,omitempty"`
	// The kind of material to be collected.
	TypeCollected *CodeableConcept `bson:"typeCollected,omitempty" json:"typeCollected,omitempty" xml:"typeCollected,omitempty"`
	// Preparation of the patient for specimen collection.
	PatientPreparation []CodeableConcept `bson:"patientPreparation,omitempty" json:"patientPreparation,omitempty" xml:"patientPreparation,omitempty"`
	// Time aspect of specimen collection (duration or offset).
	TimeAspect *string `bson:"timeAspect,omitempty" json:"timeAspect,omitempty" xml:"timeAspect,omitempty"`
	// The action to be performed for collecting the specimen.
	Collection []CodeableConcept `bson:"collection,omitempty" json:"collection,omitempty" xml:"collection,omitempty"`
	// Specimen conditioned in a container as expected by the testing laboratory.
	TypeTested []SpecimenDefinitionTypeTested `bson:"typeTested,omitempty" json:"typeTested,omitempty" xml:"typeTested,omitempty"`
}

type SpecimenDefinitionTypeTested struct {
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
	// Primary of secondary specimen.
	IsDerived *bool `bson:"isDerived,omitempty" json:"isDerived,omitempty" xml:"isDerived,omitempty"`
	// The kind of specimen conditioned for testing expected by lab.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The preference for this type of conditioned specimen.
	Preference SpecimenContainedPreference `bson:"preference" json:"preference" xml:"preference"`
	// The specimen's container.
	Container *SpecimenDefinitionTypeTestedContainer `bson:"container,omitempty" json:"container,omitempty" xml:"container,omitempty"`
	// Requirements for delivery and special handling of this kind of conditioned specimen.
	Requirement *string `bson:"requirement,omitempty" json:"requirement,omitempty" xml:"requirement,omitempty"`
	// The usual time that a specimen of this kind is retained after the ordered tests are completed, for the purpose of additional testing.
	RetentionTime *Duration `bson:"retentionTime,omitempty" json:"retentionTime,omitempty" xml:"retentionTime,omitempty"`
	// Specimen can be used by only one test or panel if the value is "true".
	SingleUse *bool `bson:"singleUse,omitempty" json:"singleUse,omitempty" xml:"singleUse,omitempty"`
	// Criterion for rejection of the specimen in its container by the laboratory.
	RejectionCriterion []CodeableConcept `bson:"rejectionCriterion,omitempty" json:"rejectionCriterion,omitempty" xml:"rejectionCriterion,omitempty"`
	// Set of instructions for preservation/transport of the specimen at a defined temperature interval, prior the testing process.
	Handling []SpecimenDefinitionTypeTestedHandling `bson:"handling,omitempty" json:"handling,omitempty" xml:"handling,omitempty"`
	// Where the specimen will be tested: e.g., lab, sector, device or any combination of these.
	TestingDestination []CodeableConcept `bson:"testingDestination,omitempty" json:"testingDestination,omitempty" xml:"testingDestination,omitempty"`
}

type SpecimenDefinitionTypeTestedContainer struct {
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
	// The type of material of the container.
	//
	// Comment: Examples: glass, plastic, metal.
	Material *CodeableConcept `bson:"material,omitempty" json:"material,omitempty" xml:"material,omitempty"`
	// The type of container used to contain this kind of specimen.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Color of container cap.
	Cap *CodeableConcept `bson:"cap,omitempty" json:"cap,omitempty" xml:"cap,omitempty"`
	// The textual description of the kind of container.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// The capacity (volume or other measure) of this kind of container.
	Capacity *Quantity `bson:"capacity,omitempty" json:"capacity,omitempty" xml:"capacity,omitempty"`
	// The minimum volume to be conditioned in the container.
	MinimumVolumeQuantity *Quantity `bson:"minimumVolumeQuantity,omitempty" json:"minimumVolumeQuantity,omitempty" xml:"minimumVolumeQuantity,omitempty"`
	MinimumVolumeString   *string   `bson:"minimumVolumeString,omitempty" json:"minimumVolumeString,omitempty" xml:"minimumVolumeString,omitempty"`
	// Substance introduced in the kind of container to preserve, maintain or enhance the specimen.
	//
	// Examples: Formalin, Citrate, EDTA.
	Additive []SpecimenDefinitionTypeTestedContainerAdditive `bson:"additive,omitempty" json:"additive,omitempty" xml:"additive,omitempty"`
	// Special processing that should be applied to the container for this kind of specimen.
	Preparation *string `bson:"preparation,omitempty" json:"preparation,omitempty" xml:"preparation,omitempty"`
}

type SpecimenDefinitionTypeTestedContainerAdditive struct {
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
	// Substance introduced in the kind of container to preserve, maintain or enhance the specimen.
	//
	// Examples: Formalin, Citrate, EDTA.
	AdditiveCodeableConcept CodeableConcept `bson:"additiveCodeableConcept" json:"additiveCodeableConcept" xml:"additiveCodeableConcept"`
	AdditiveReference       Reference       `bson:"additiveReference" json:"additiveReference" xml:"additiveReference"`
}

type SpecimenDefinitionTypeTestedHandling struct {
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
	// It qualifies the interval of temperature, which characterizes an occurrence of handling.
	//
	// Conditions that are not related to temperature may be handled in the instruction element.
	TemperatureQualifier *CodeableConcept `bson:"temperatureQualifier,omitempty" json:"temperatureQualifier,omitempty" xml:"temperatureQualifier,omitempty"`
	// The temperature interval for this set of handling instructions.
	TemperatureRange *Range `bson:"temperatureRange,omitempty" json:"temperatureRange,omitempty" xml:"temperatureRange,omitempty"`
	// The maximum time interval of preservation of the specimen with these conditions.
	MaxDuration *Duration `bson:"maxDuration,omitempty" json:"maxDuration,omitempty" xml:"maxDuration,omitempty"`
	// Additional textual instructions for the preservation or transport of the specimen.
	//
	// For instance, 'Protect from light exposure'.
	Instruction *string `bson:"instruction,omitempty" json:"instruction,omitempty" xml:"instruction,omitempty"`
}

func (r SpecimenDefinition) ResourceType() string {
	return "SpecimenDefinition"
}

// MarshalJSON implements json.Marshaler
func (r SpecimenDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		SpecimenDefinition
		ResourceType string `json:"resourceType"`
	}{
		ResourceType:       r.ResourceType(),
		SpecimenDefinition: r,
	})
}
