package fhires

import v2 "encoding/json/v2"

// A ResearchSubject is a participant or object which is the recipient of investigative activities in a research study.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/ResearchSubject
type ResearchSubject struct {
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
	// Identifiers assigned to this research subject for a study.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The publication state of the resource (not of the subject).
	Status PublicationStatus `bson:"status" json:"status" xml:"status"`
	// The current state (status) of the subject and resons for status change where appropriate.
	//
	// Comment: This is intended to deal with the confusion routinely created by haing two conflated concepts of being in a particular state and having achieved a particular milestone.  In strict terms a milestone is a point of time event that results in a change from one state to another.  The state before the milestone is achieved is often given the same name as the milestone, and sometimes the state may have the same description.  For instance "Randomised" and "Visit 1" may be different milestones but the state remains at "on study" after each of them.
	//
	// It is likely that more than one "state" pattern will be recorded for a subject and a type has been introduced to allow this simultaneous recording.
	Progress []ResearchSubjectProgress `bson:"progress,omitempty" json:"progress,omitempty" xml:"progress,omitempty"`
	// The dates the subject began and ended their participation in the study.
	Period *Period `bson:"period,omitempty" json:"period,omitempty" xml:"period,omitempty"`
	// Reference to the study the subject is participating in.
	Study Reference `bson:"study" json:"study" xml:"study"`
	// The record of the person, animal or other entity involved in the study.
	Subject Reference `bson:"subject" json:"subject" xml:"subject"`
	// The name of the arm in the study the subject is expected to follow as part of this study.
	//
	// Comment: Maybe this is really PlanDefinition.
	AssignedComparisonGroup *string `bson:"assignedComparisonGroup,omitempty" json:"assignedComparisonGroup,omitempty" xml:"assignedComparisonGroup,omitempty"`
	// The name of the arm in the study the subject actually followed as part of this study.
	ActualComparisonGroup *string `bson:"actualComparisonGroup,omitempty" json:"actualComparisonGroup,omitempty" xml:"actualComparisonGroup,omitempty"`
	// A record of the patient's informed agreement to participate in the study.
	Consent []Reference `bson:"consent,omitempty" json:"consent,omitempty" xml:"consent,omitempty"`
}

type ResearchSubjectProgress struct {
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
	// Identifies the aspect of the subject's journey that the state refers to.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The current state of the subject.
	SubjectState *CodeableConcept `bson:"subjectState,omitempty" json:"subjectState,omitempty" xml:"subjectState,omitempty"`
	// The milestones the subject has passed through.
	//
	// Comment: There can be multiple entries but it is also valid to just have the most recent.  This should npt be rlied upon as the full path the subject has taken.
	Milestone *CodeableConcept `bson:"milestone,omitempty" json:"milestone,omitempty" xml:"milestone,omitempty"`
	// The reason for the state change.
	//
	//  If coded it should follow the formal subject state model.
	Reason *CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty" xml:"reason,omitempty"`
	// The date when the new status started.
	//
	// Comment: This is NOT the date the change in state was recorded.
	StartDate *string `bson:"startDate,omitempty" json:"startDate,omitempty" xml:"startDate,omitempty"`
	// The date when the state ended.
	EndDate *string `bson:"endDate,omitempty" json:"endDate,omitempty" xml:"endDate,omitempty"`
}

func (r ResearchSubject) ResourceType() string {
	return "ResearchSubject"
}

// MarshalJSON implements json.Marshaler
func (r ResearchSubject) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		ResearchSubject
		ResourceType string `json:"resourceType"`
	}{
		ResearchSubject: r,
		ResourceType:    r.ResourceType(),
	})
}
