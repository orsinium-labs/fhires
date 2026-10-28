package fhires

import "encoding/json/v2"

// A scientific study of nature that sometimes includes processes involved in health and disease.
//
// For example, clinical trials are research studies that involve people. These studies may be related to new ways to screen, prevent, diagnose, and treat disease. They may also study certain outcomes and certain groups of people by looking at data collected in the past or future.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/ResearchStudy
type ResearchStudy struct {
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
	// Canonical identifier for this study resource, represented as a globally unique URI.
	URL *string `bson:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
	// Identifiers assigned to this research study by the sponsor or other systems.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The business version for the study record.
	Version *string `bson:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
	// Name for this study (computer friendly).
	Name *string `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// The human readable name of the research study.
	Title *string `bson:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Additional names for the study.
	Label []ResearchStudyLabel `bson:"label,omitempty" json:"label,omitempty" xml:"label,omitempty"`
	// The set of steps expected to be performed as part of the execution of the study.
	Protocol []Reference `bson:"protocol,omitempty" json:"protocol,omitempty" xml:"protocol,omitempty"`
	// A larger research study of which this particular study is a component or step.
	PartOf []Reference `bson:"partOf,omitempty" json:"partOf,omitempty" xml:"partOf,omitempty"`
	// Citations, references, URLs and other related documents.
	//
	//  When using relatedArtifact to share URLs, the relatedArtifact.type will often be set to one of "documentation" or "supported-with" and the URL value will often be in relatedArtifact.document.url but another possible location is relatedArtifact.resource when it is a canonical URL.
	RelatedArtifact []RelatedArtifact `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty" xml:"relatedArtifact,omitempty"`
	// The date (and optionally time) when the ResearchStudy Resource was last significantly changed.
	//
	// The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the ResearchStudy Resource changes.
	Date *string `bson:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// The publication state of the resource (not of the study).
	Status PublicationStatus `bson:"status" json:"status" xml:"status"`
	// The type of study based upon the intent of the study activities.
	//
	// A classification of the intent of the study.
	PrimaryPurposeType *CodeableConcept `bson:"primaryPurposeType,omitempty" json:"primaryPurposeType,omitempty" xml:"primaryPurposeType,omitempty"`
	// The stage in the progression of a therapy from initial experimental use in humans in clinical trials to post-market evaluation.
	Phase *CodeableConcept `bson:"phase,omitempty" json:"phase,omitempty" xml:"phase,omitempty"`
	// Codes categorizing the type of study such as investigational vs.
	//
	// observational, type of blinding, type of randomization, safety vs. efficacy, etc.
	StudyDesign []CodeableConcept `bson:"studyDesign,omitempty" json:"studyDesign,omitempty" xml:"studyDesign,omitempty"`
	// The medication(s), food(s), therapy(ies), device(s) or other concerns or interventions that the study is seeking to gain more information about.
	Focus []CodeableReference `bson:"focus,omitempty" json:"focus,omitempty" xml:"focus,omitempty"`
	// The condition that is the focus of the study.
	//
	//  For example, In a study to examine risk factors for Lupus, might have as an inclusion criterion "healthy volunteer", but the target condition code would be a Lupus SNOMED code.
	Condition []CodeableConcept `bson:"condition,omitempty" json:"condition,omitempty" xml:"condition,omitempty"`
	// Key terms to aid in searching for or filtering the study.
	Keyword []CodeableConcept `bson:"keyword,omitempty" json:"keyword,omitempty" xml:"keyword,omitempty"`
	// A country, state or other area where the study is taking place rather than its precise geographic location or address.
	Region []CodeableConcept `bson:"region,omitempty" json:"region,omitempty" xml:"region,omitempty"`
	// A brief text for explaining the study.
	DescriptionSummary *string `bson:"descriptionSummary,omitempty" json:"descriptionSummary,omitempty" xml:"descriptionSummary,omitempty"`
	// A detailed and human-readable narrative of the study. E.g., study abstract.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Identifies the start date and the expected (or actual, depending on status) end date for the study.
	Period *Period `bson:"period,omitempty" json:"period,omitempty" xml:"period,omitempty"`
	// A facility in which study activities are conducted.
	Site []Reference `bson:"site,omitempty" json:"site,omitempty" xml:"site,omitempty"`
	// Comments made about the study by the performer, subject or other participants.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// Additional grouping mechanism or categorization of a research study.
	//
	// Example: FDA regulated device, FDA regulated drug, MPG Paragraph 23b (a German legal requirement), IRB-exempt, etc. Implementation Note: do not use the classifier element to support existing semantics that are already supported thru explicit elements in the resource.
	Classifier []CodeableConcept `bson:"classifier,omitempty" json:"classifier,omitempty" xml:"classifier,omitempty"`
	// Sponsors, collaborators, and other parties.
	//
	// Comment: For a Sponsor or a PrincipalInvestigator use the dedicated attributes provided.
	AssociatedParty []ResearchStudyAssociatedParty `bson:"associatedParty,omitempty" json:"associatedParty,omitempty" xml:"associatedParty,omitempty"`
	// Status of study with time for that status.
	ProgressStatus []ResearchStudyProgressStatus `bson:"progressStatus,omitempty" json:"progressStatus,omitempty" xml:"progressStatus,omitempty"`
	// A description and/or code explaining the premature termination of the study.
	WhyStopped *CodeableConcept `bson:"whyStopped,omitempty" json:"whyStopped,omitempty" xml:"whyStopped,omitempty"`
	// Target or actual group of participants enrolled in study.
	Recruitment *ResearchStudyRecruitment `bson:"recruitment,omitempty" json:"recruitment,omitempty" xml:"recruitment,omitempty"`
	// Describes an expected event or sequence of events for one of the subjects of a study.
	//
	// E.g. for a living subject: exposure to drug A, wash-out, exposure to drug B, wash-out, follow-up. E.g. for a stability study: {store sample from lot A at 25 degrees for 1 month}, {store sample from lot A at 40 degrees for 1 month}.
	//
	// Comment: In many clinical trials this is refered to as the ARM of the study, but such a term is not used in other sorts of trials even when there is a comparison between two or more groups.
	ComparisonGroup []ResearchStudyComparisonGroup `bson:"comparisonGroup,omitempty" json:"comparisonGroup,omitempty" xml:"comparisonGroup,omitempty"`
	// A goal that the study is aiming to achieve in terms of a scientific question to be answered by the analysis of data collected during the study.
	Objective []ResearchStudyObjective `bson:"objective,omitempty" json:"objective,omitempty" xml:"objective,omitempty"`
	// An "outcome measure", "endpoint", "effect measure" or "measure of effect" is a specific measurement or observation used to quantify the effect of experimental variables on the participants in a study, or for observational studies, to describe patterns of diseases or traits or associations with exposures, risk factors or treatment.
	//
	// Comment: A study may have multiple distinct outcome measures that can be used to assess the overall goal for a study. The goal of a study is in the objective whereas the metric by which the goal is assessed is the outcomeMeasure. Examples: Time to Local Recurrence (TLR), Disease-free Survival (DFS), 30 Day Mortality, Systolic BP
	OutcomeMeasure []ResearchStudyOutcomeMeasure `bson:"outcomeMeasure,omitempty" json:"outcomeMeasure,omitempty" xml:"outcomeMeasure,omitempty"`
	// Link to one or more sets of results generated by the study.
	//
	//  Could also link to a research registry holding the results such as ClinicalTrials.gov.
	Result []Reference `bson:"result,omitempty" json:"result,omitempty" xml:"result,omitempty"`
}

type ResearchStudyLabel struct {
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
	// Kind of name.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The name.
	Value *string `bson:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

type ResearchStudyAssociatedParty struct {
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
	// Name of associated party.
	Name *string `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Type of association.
	Role CodeableConcept `bson:"role" json:"role" xml:"role"`
	// Identifies the start date and the end date of the associated party in the role.
	//
	// Comment: The cardinality is 0..* due to the fact that an associated party may be intermittently active in a given role over multiple time periods.
	Period []Period `bson:"period,omitempty" json:"period,omitempty" xml:"period,omitempty"`
	// A categorization other than role for the associated party.
	Classifier []CodeableConcept `bson:"classifier,omitempty" json:"classifier,omitempty" xml:"classifier,omitempty"`
	// Individual or organization associated with study (use practitionerRole to specify their organisation).
	//
	// Comment: Suggestions of a better attribute name are appreciated
	Party *Reference `bson:"party,omitempty" json:"party,omitempty" xml:"party,omitempty"`
}

type ResearchStudyProgressStatus struct {
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
	// Label for status or state (e.g. recruitment status).
	State CodeableConcept `bson:"state" json:"state" xml:"state"`
	// An indication of whether or not the date is a known date when the state changed or will change.
	//
	// A value of true indicates a known date. A value of false indicates an estimated date.
	Actual *bool `bson:"actual,omitempty" json:"actual,omitempty" xml:"actual,omitempty"`
	// Date range.
	Period *Period `bson:"period,omitempty" json:"period,omitempty" xml:"period,omitempty"`
}

type ResearchStudyRecruitment struct {
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
	// Estimated total number of participants to be enrolled.
	TargetNumber *int `bson:"targetNumber,omitempty" json:"targetNumber,omitempty" xml:"targetNumber,omitempty"`
	// Actual total number of participants enrolled in study.
	ActualNumber *int `bson:"actualNumber,omitempty" json:"actualNumber,omitempty" xml:"actualNumber,omitempty"`
	// Inclusion and exclusion criteria.
	Eligibility *Reference `bson:"eligibility,omitempty" json:"eligibility,omitempty" xml:"eligibility,omitempty"`
	// Group of participants who were enrolled in study.
	ActualGroup *Reference `bson:"actualGroup,omitempty" json:"actualGroup,omitempty" xml:"actualGroup,omitempty"`
}

type ResearchStudyComparisonGroup struct {
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
	// Allows the comparisonGroup for the study and the comparisonGroup for the subject to be linked easily.
	LinkID *string `bson:"linkId,omitempty" json:"linkId,omitempty" xml:"linkId,omitempty"`
	// Unique, human-readable label for this comparisonGroup of the study.
	Name string `bson:"name" json:"name" xml:"name"`
	// Categorization of study comparisonGroup, e.g.
	//
	// experimental, active comparator, placebo comparater.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// A succinct description of the path through the study that would be followed by a subject adhering to this comparisonGroup.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Interventions or exposures in this comparisonGroup or cohort.
	IntendedExposure []Reference `bson:"intendedExposure,omitempty" json:"intendedExposure,omitempty" xml:"intendedExposure,omitempty"`
	// Group of participants who were enrolled in study comparisonGroup.
	ObservedGroup *Reference `bson:"observedGroup,omitempty" json:"observedGroup,omitempty" xml:"observedGroup,omitempty"`
}

type ResearchStudyObjective struct {
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
	// Unique, human-readable label for this objective of the study.
	Name *string `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// The kind of study objective.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Free text description of the objective of the study.
	//
	//  This is what the study is trying to achieve rather than how it is going to achieve it (see ResearchStudy.description).
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
}

type ResearchStudyOutcomeMeasure struct {
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
	// Label for the outcome.
	Name *string `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// The parameter or characteristic being assessed as one of the values by which the study is assessed.
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Description of the outcome.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Structured outcome definition.
	Reference *Reference `bson:"reference,omitempty" json:"reference,omitempty" xml:"reference,omitempty"`
}

func (r ResearchStudy) ResourceType() string {
	return "ResearchStudy"
}

// MarshalJSON implements json.Marshaler
func (r ResearchStudy) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ResearchStudy
		ResourceType string `json:"resourceType"`
	}{
		ResearchStudy: r,
		ResourceType:  r.ResourceType(),
	})
}
