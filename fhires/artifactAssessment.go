package fhires

import v2 "encoding/json/v2"

// This Resource provides one or more comments, classifiers or ratings about a Resource and supports attribution and rights management metadata for the added content.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/ArtifactAssessment
type ArtifactAssessment struct {
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
	// A formal identifier that is used to identify this artifact assessment when it is represented in other formats, or referenced in a specification, model, design or an instance.
	//
	// Comment: Typically, this is used for identifiers that can go in an HL7 V3 II (instance identifier) data type, and can then identify this activity definition outside of FHIR, where it is not possible to use the logical URI.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// A short title for the assessment for use in displaying and selecting.
	Title *string `bson:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Display of or reference to the bibliographic citation of the comment, classifier, or rating.
	CiteAsReference *Reference `bson:"citeAsReference,omitempty" json:"citeAsReference,omitempty" xml:"citeAsReference,omitempty"`
	CiteAsMarkdown  *string    `bson:"citeAsMarkdown,omitempty" json:"citeAsMarkdown,omitempty" xml:"citeAsMarkdown,omitempty"`
	// The date  (and optionally time) when the artifact assessment was published.
	//
	// The date must change when the disposition changes and it must change if the workflow status code changes. In addition, it should change when the substantive content of the artifact assessment changes.
	//
	// Comment: The date is often not tracked until the resource is published, but may be present on draft content. Note that this is not the same as the resource last-modified-date, since the resource may be a secondary representation of the artifact assessment. Additional specific dates may be added as extensions or be found by consulting Provenances associated with past versions of the resource.
	Date *string `bson:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// A copyright statement relating to the artifact assessment and/or its contents.
	//
	// Copyright statements are generally legal restrictions on the use and publishing of the artifact assessment.
	Copyright *string `bson:"copyright,omitempty" json:"copyright,omitempty" xml:"copyright,omitempty"`
	// The date on which the resource content was approved by the publisher.
	//
	// Approval happens once when the content is officially approved for usage.
	//
	// Comment: The 'date' element may be more recent than the approval date because of minor changes or editorial corrections.
	ApprovalDate *string `bson:"approvalDate,omitempty" json:"approvalDate,omitempty" xml:"approvalDate,omitempty"`
	// The date on which the resource content was last reviewed.
	//
	// Review happens periodically after approval but does not change the original approval date.
	//
	// Comment: If specified, this date follows the original approval date.
	LastReviewDate *string `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty" xml:"lastReviewDate,omitempty"`
	// A reference to a resource, canonical resource, or non-FHIR resource which the comment or assessment is about.
	ArtifactReference Reference `bson:"artifactReference" json:"artifactReference" xml:"artifactReference"`
	ArtifactCanonical string    `bson:"artifactCanonical" json:"artifactCanonical" xml:"artifactCanonical"`
	ArtifactUri       string    `bson:"artifactUri" json:"artifactUri" xml:"artifactUri"`
	// A component comment, classifier, or rating of the artifact.
	Content []ArtifactAssessmentContent `bson:"content,omitempty" json:"content,omitempty" xml:"content,omitempty"`
	// Indicates the workflow status of the comment or change request.
	WorkflowStatus *ArtifactAssessmentWorkflowStatus `bson:"workflowStatus,omitempty" json:"workflowStatus,omitempty" xml:"workflowStatus,omitempty"`
	// Indicates the disposition of the responsible party to the comment or change request.
	Disposition *ArtifactAssessmentDisposition `bson:"disposition,omitempty" json:"disposition,omitempty" xml:"disposition,omitempty"`
}

type ArtifactAssessmentContent struct {
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
	// The type of information this component of the content represents.
	InformationType *ArtifactAssessmentInformationType `bson:"informationType,omitempty" json:"informationType,omitempty" xml:"informationType,omitempty"`
	// A brief summary of the content of this component.
	Summary *string `bson:"summary,omitempty" json:"summary,omitempty" xml:"summary,omitempty"`
	// Indicates what type of content this component represents.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Represents a rating, classifier, or assessment of the artifact.
	Classifier []CodeableConcept `bson:"classifier,omitempty" json:"classifier,omitempty" xml:"classifier,omitempty"`
	// A quantitative rating of the artifact.
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty" xml:"quantity,omitempty"`
	// Indicates who or what authored the content.
	Author *Reference `bson:"author,omitempty" json:"author,omitempty" xml:"author,omitempty"`
	// A URI that points to what the comment is about, such as a line of text in the CQL, or a specific element in a resource.
	//
	// Comment: The target element is used to point the comment to aspect of the artifact, such as a text range within a CQL library (e.g. #content?0:0-120:80).
	Path []string `bson:"path,omitempty" json:"path,omitempty" xml:"path,omitempty"`
	// Additional related artifacts that provide supporting documentation, additional evidence, or further information related to the content.
	RelatedArtifact []RelatedArtifact `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty" xml:"relatedArtifact,omitempty"`
	// Acceptable to publicly share the comment, classifier or rating.
	FreeToShare *bool `bson:"freeToShare,omitempty" json:"freeToShare,omitempty" xml:"freeToShare,omitempty"`
	// If the informationType is container, the components of the content.
	Component []ArtifactAssessmentContent `bson:"component,omitempty" json:"component,omitempty"`
}

func (r ArtifactAssessment) ResourceType() string {
	return "ArtifactAssessment"
}

// MarshalJSON implements json.Marshaler
func (r ArtifactAssessment) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		ArtifactAssessment
		ResourceType string `json:"resourceType"`
	}{
		ArtifactAssessment: r,
		ResourceType:       r.ResourceType(),
	})
}
