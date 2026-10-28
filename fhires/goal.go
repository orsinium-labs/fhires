package fhires

import v2 "encoding/json/v2"

// Describes the intended objective(s) for a patient, group or organization care, for example, weight loss, restoring an activity of daily living, obtaining herd immunity via immunization, meeting a process improvement objective, etc.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Goal
type Goal struct {
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
	// Business identifiers assigned to this goal by the performer or other systems which remain constant as the resource is updated and propagates from server to server.
	//
	// Comment: This is a business identifier, not a resource identifier (see [discussion](resource.html#identifiers)).  It is best practice for the identifier to only appear on a single resource instance, however business practices may occasionally dictate that multiple resource instances with the same identifier can exist - possibly even with different resource types.  For example, multiple Patient and a Person resource instance might share the same social insurance number.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The state of the goal throughout its lifecycle.
	//
	// Comment: This element is labeled as a modifier because the lifecycleStatus contains codes that mark the resource as not currently valid.
	LifecycleStatus GoalLifecycleStatus `bson:"lifecycleStatus" json:"lifecycleStatus" xml:"lifecycleStatus"`
	// Describes the progression, or lack thereof, towards the goal against the target.
	AchievementStatus *CodeableConcept `bson:"achievementStatus,omitempty" json:"achievementStatus,omitempty" xml:"achievementStatus,omitempty"`
	// Indicates a category the goal falls within.
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
	// After meeting the goal, ongoing activity is needed to sustain the goal objective.
	//
	// Comment: For example, getting a yellow fever vaccination for a planned trip is a goal that is designed to be completed (continuous = false).  A goal to sustain HbA1c levels would not be a one-time goal (continuous = true).
	Continuous *bool `bson:"continuous,omitempty" json:"continuous,omitempty" xml:"continuous,omitempty"`
	// Identifies the mutually agreed level of importance associated with reaching/sustaining the goal.
	//
	// Comment: Extensions are available to track priorities as established by each participant (i.e. Priority from the patient's perspective, different practitioners' perspectives, family member's perspectives)The ordinal extension on Coding can be used to convey a numerically comparable ranking to priority.  (Keep in mind that different coding systems may use a "low value=important".
	Priority *CodeableConcept `bson:"priority,omitempty" json:"priority,omitempty" xml:"priority,omitempty"`
	// Human-readable and/or coded description of a specific desired objective of care, such as "control blood pressure" or "negotiate an obstacle course" or "dance with child at wedding".
	//
	// Comment: If no code is available, use CodeableConcept.text.
	Description CodeableConcept `bson:"description" json:"description" xml:"description"`
	// Identifies the patient, group or organization for whom the goal is being established.
	Subject Reference `bson:"subject" json:"subject" xml:"subject"`
	// The date or event after which the goal should begin being pursued.
	StartDate            *string          `bson:"startDate,omitempty" json:"startDate,omitempty" xml:"startDate,omitempty"`
	StartCodeableConcept *CodeableConcept `bson:"startCodeableConcept,omitempty" json:"startCodeableConcept,omitempty" xml:"startCodeableConcept,omitempty"`
	// Indicates what should be done by when.
	//
	// Comment: When multiple targets are present for a single goal instance, all targets must be met for the overall goal to be met.
	Target []GoalTarget `bson:"target,omitempty" json:"target,omitempty" xml:"target,omitempty"`
	// Identifies when the current status.
	//
	//  I.e. When initially created, when achieved, when cancelled, etc.
	//
	// Comment: To see the date for past statuses, query history.
	StatusDate *string `bson:"statusDate,omitempty" json:"statusDate,omitempty" xml:"statusDate,omitempty"`
	// Captures the reason for the current status.
	//
	// Comment: This will typically be captured for statuses such as rejected, on-hold or cancelled, but could be present for others.
	StatusReason *string `bson:"statusReason,omitempty" json:"statusReason,omitempty" xml:"statusReason,omitempty"`
	// Indicates whose goal this is - patient goal, practitioner goal, etc.
	//
	// Comment: This is the individual or team responsible for establishing the goal, not necessarily who recorded it.  (For that, use the Provenance resource.).
	Source *Reference `bson:"source,omitempty" json:"source,omitempty" xml:"source,omitempty"`
	// The identified conditions and other health record elements that are intended to be addressed by the goal.
	Addresses []Reference `bson:"addresses,omitempty" json:"addresses,omitempty" xml:"addresses,omitempty"`
	// Any comments related to the goal.
	//
	// Comment: May be used for progress notes, concerns or other related information that doesn't actually describe the goal itself.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// Identifies the change (or lack of change) at the point when the status of the goal is assessed.
	//
	// Comment: Note that this should not duplicate the goal status; The goal outcome is independent of the outcome of the related activities.  For example, if the Goal is to achieve a target body weight of 150 lb and a care plan activity is defined to diet, then the care plan’s activity outcome could be calories consumed whereas goal outcome is an observation for the actual body weight measured.
	Outcome []CodeableReference `bson:"outcome,omitempty" json:"outcome,omitempty" xml:"outcome,omitempty"`
}

type GoalTarget struct {
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
	// The parameter whose value is being tracked, e.g.
	//
	// body weight, blood pressure, or hemoglobin A1c level.
	Measure *CodeableConcept `bson:"measure,omitempty" json:"measure,omitempty" xml:"measure,omitempty"`
	// The target value of the focus to be achieved to signify the fulfillment of the goal, e.g.
	//
	// 150 pounds, 7.0%. Either the high or low or both values of the range can be specified. When a low value is missing, it indicates that the goal is achieved at any focus value at or below the high value. Similarly, if the high value is missing, it indicates that the goal is achieved at any focus value at or above the low value.
	//
	// Comment: A CodeableConcept with just a text would be used instead of a string if the field was usually coded, or if the type associated with the Goal.target.measure defines a coded value.
	DetailQuantity        *Quantity        `bson:"detailQuantity,omitempty" json:"detailQuantity,omitempty" xml:"detailQuantity,omitempty"`
	DetailRange           *Range           `bson:"detailRange,omitempty" json:"detailRange,omitempty" xml:"detailRange,omitempty"`
	DetailCodeableConcept *CodeableConcept `bson:"detailCodeableConcept,omitempty" json:"detailCodeableConcept,omitempty" xml:"detailCodeableConcept,omitempty"`
	DetailString          *string          `bson:"detailString,omitempty" json:"detailString,omitempty" xml:"detailString,omitempty"`
	DetailBoolean         *bool            `bson:"detailBoolean,omitempty" json:"detailBoolean,omitempty" xml:"detailBoolean,omitempty"`
	DetailInteger         *int             `bson:"detailInteger,omitempty" json:"detailInteger,omitempty" xml:"detailInteger,omitempty"`
	DetailRatio           *Ratio           `bson:"detailRatio,omitempty" json:"detailRatio,omitempty" xml:"detailRatio,omitempty"`
	// Indicates either the date or the duration after start by which the goal should be met.
	DueDate     *string   `bson:"dueDate,omitempty" json:"dueDate,omitempty" xml:"dueDate,omitempty"`
	DueDuration *Duration `bson:"dueDuration,omitempty" json:"dueDuration,omitempty" xml:"dueDuration,omitempty"`
}

func (r Goal) ResourceType() string {
	return "Goal"
}

// MarshalJSON implements json.Marshaler
func (r Goal) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		Goal
		ResourceType string `json:"resourceType"`
	}{
		Goal:         r,
		ResourceType: r.ResourceType(),
	})
}
