package fhires

import "encoding/json/v2"

// Describes the intention of how one or more practitioners intend to deliver care for a particular patient, group or community for a period of time, possibly limited to care for a specific condition or set of conditions.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/CarePlan
type CarePlan struct {
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
	// Business identifiers assigned to this care plan by the performer or other systems which remain constant as the resource is updated and propagates from server to server.
	//
	// Comment: This is a business identifier, not a resource identifier (see [discussion](resource.html#identifiers)).  It is best practice for the identifier to only appear on a single resource instance, however business practices may occasionally dictate that multiple resource instances with the same identifier can exist - possibly even with different resource types.  For example, multiple Patient and a Person resource instance might share the same social insurance number.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The URL pointing to a FHIR-defined protocol, guideline, questionnaire or other definition that is adhered to in whole or in part by this CarePlan.
	InstantiatesCanonical []string `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty" xml:"instantiatesCanonical,omitempty"`
	// The URL pointing to an externally maintained protocol, guideline, questionnaire or other definition that is adhered to in whole or in part by this CarePlan.
	//
	// Comment: This might be an HTML page, PDF, etc. or could just be a non-resolvable URI identifier.
	InstantiatesUri []string `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty" xml:"instantiatesUri,omitempty"`
	// A higher-level request resource (i.e.
	//
	// a plan, proposal or order) that is fulfilled in whole or in part by this care plan.
	BasedOn []Reference `bson:"basedOn,omitempty" json:"basedOn,omitempty" xml:"basedOn,omitempty"`
	// Completed or terminated care plan whose function is taken by this new care plan.
	//
	// Comment: The replacement could be because the initial care plan was immediately rejected (due to an issue) or because the previous care plan was completed, but the need for the action described by the care plan remains ongoing.
	Replaces []Reference `bson:"replaces,omitempty" json:"replaces,omitempty" xml:"replaces,omitempty"`
	// A larger care plan of which this particular care plan is a component or step.
	//
	// Comment: Each care plan is an independent request, such that having a care plan be part of another care plan can cause issues with cascading statuses.  As such, this element is still being discussed.
	PartOf []Reference `bson:"partOf,omitempty" json:"partOf,omitempty" xml:"partOf,omitempty"`
	// Indicates whether the plan is currently being acted upon, represents future intentions or is now a historical record.
	//
	// Comment: The unknown code is not to be used to convey other statuses.  The unknown code should be used when one of the statuses applies, but the authoring system doesn't know the current state of the care plan.
	//
	// This element is labeled as a modifier because the status contains the code entered-in-error that marks the plan as not currently valid.
	Status RequestStatus `bson:"status" json:"status" xml:"status"`
	// Indicates the level of authority/intentionality associated with the care plan and where the care plan fits into the workflow chain.
	//
	// Comment: This element is labeled as a modifier because the intent alters when and how the resource is actually applicable.
	//
	// This element is expected to be immutable. E.g. A "proposal" instance should never change to be a "plan" instance or "order" instance. Instead, a new instance 'basedOn' the prior instance should be created with the new 'intent' value.
	Intent CarePlanIntent `bson:"intent" json:"intent" xml:"intent"`
	// Identifies what "kind" of plan this is to support differentiation between multiple co-existing plans; e.g.
	//
	// "Home health", "psychiatric", "asthma", "disease management", "wellness plan", etc.
	//
	// Comment: There may be multiple axes of categorization and one plan may serve multiple purposes.  In some cases, this may be redundant with references to CarePlan.addresses.
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
	// Human-friendly name for the care plan.
	Title *string `bson:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// A description of the scope and nature of the plan.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Identifies the patient or group whose intended care is described by the plan.
	Subject Reference `bson:"subject" json:"subject" xml:"subject"`
	// The Encounter during which this CarePlan was created or to which the creation of this record is tightly associated.
	//
	// Comment: This will typically be the encounter the event occurred within, but some activities may be initiated prior to or after the official completion of an encounter but still be tied to the context of the encounter. CarePlan activities conducted as a result of the care plan may well occur as part of other encounters.
	Encounter *Reference `bson:"encounter,omitempty" json:"encounter,omitempty" xml:"encounter,omitempty"`
	// Indicates when the plan did (or is intended to) come into effect and end.
	//
	// Comment: Any activities scheduled as part of the plan should be constrained to the specified period regardless of whether the activities are planned within a single encounter/episode or across multiple encounters/episodes (e.g. the longitudinal management of a chronic condition).
	Period *Period `bson:"period,omitempty" json:"period,omitempty" xml:"period,omitempty"`
	// Represents when this particular CarePlan record was created in the system, which is often a system-generated date.
	Created *string `bson:"created,omitempty" json:"created,omitempty" xml:"created,omitempty"`
	// When populated, the custodian is responsible for the care plan.
	//
	// The care plan is attributed to the custodian.
	//
	// Comment: The custodian might or might not be a contributor.
	Custodian *Reference `bson:"custodian,omitempty" json:"custodian,omitempty" xml:"custodian,omitempty"`
	// Identifies the individual(s), organization or device who provided the contents of the care plan.
	//
	// Comment: Collaborative care plans may have multiple contributors.
	Contributor []Reference `bson:"contributor,omitempty" json:"contributor,omitempty" xml:"contributor,omitempty"`
	// Identifies all people and organizations who are expected to be involved in the care envisioned by this plan.
	CareTeam []Reference `bson:"careTeam,omitempty" json:"careTeam,omitempty" xml:"careTeam,omitempty"`
	// Identifies the conditions/problems/concerns/diagnoses/etc.
	//
	// whose management and/or mitigation are handled by this plan.
	//
	// Comment: Use CarePlan.addresses.concept when a code sufficiently describes the concern (e.g. condition, problem, diagnosis, risk). Use CarePlan.addresses.reference when referencing a resource, which allows more information to be conveyed, such as onset date. CarePlan.addresses.concept and CarePlan.addresses.reference are not meant to be duplicative. For a single concern, either CarePlan.addresses.concept or CarePlan.addresses.reference can be used. CarePlan.addresses.concept may be a summary code, or CarePlan.addresses.reference may be used to reference a very precise definition of the concern using Condition. Both CarePlan.addresses.concept and CarePlan.addresses.reference can be used if they are describing different concerns for the care plan.
	Addresses []CodeableReference `bson:"addresses,omitempty" json:"addresses,omitempty" xml:"addresses,omitempty"`
	// Identifies portions of the patient's record that specifically influenced the formation of the plan.
	//
	//  These might include comorbidities, recent procedures, limitations, recent assessments, etc.
	//
	// Comment: Use "concern" to identify specific conditions addressed by the care plan.  supportingInfo can be used to convey one or more Advance Directives or Medical Treatment Consent Directives by referencing Consent or any other request resource with intent = directive.
	SupportingInfo []Reference `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty" xml:"supportingInfo,omitempty"`
	// Describes the intended objective(s) of carrying out the care plan.
	//
	// Comment: Goal can be achieving a particular change or merely maintaining a current state or even slowing a decline.
	Goal []Reference `bson:"goal,omitempty" json:"goal,omitempty" xml:"goal,omitempty"`
	// Identifies an action that has occurred or is a planned action to occur as part of the plan.
	//
	// For example, a medication to be used, lab tests to perform, self-monitoring that has occurred, education etc.
	Activity []CarePlanActivity `bson:"activity,omitempty" json:"activity,omitempty" xml:"activity,omitempty"`
	// General notes about the care plan not covered elsewhere.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
}

type CarePlanActivity struct {
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
	// Identifies the activity that was performed.
	//
	// For example, an activity could be patient education, exercise, or a medication administration. The reference to an "event" resource, such as Procedure or Encounter or Observation, represents the activity that was performed. The requested activity can be conveyed using the CarePlan.activity.plannedActivityReference (a reference to a “request” resource).
	//
	// Comment: Note that this should not duplicate the activity status (e.g. completed or in progress). The activity performed is independent of the outcome of the related goal(s). For example, if the goal is to achieve a target body weight of 150 lbs and an activity is defined to exercise, then the activity performed could be amount and intensity of exercise performed whereas the goal outcome is an observation for the actual body weight measured.
	PerformedActivity []CodeableReference `bson:"performedActivity,omitempty" json:"performedActivity,omitempty" xml:"performedActivity,omitempty"`
	// Notes about the adherence/status/progress of the activity.
	//
	// Comment: This element should NOT be used to describe the activity to be performed - that occurs either within the resource pointed to by activity.detail.reference or in activity.detail.description.
	Progress []Annotation `bson:"progress,omitempty" json:"progress,omitempty" xml:"progress,omitempty"`
	// The details of the proposed activity represented in a specific resource.
	//
	// Comment: Standard extension exists ([http://hl7.org/fhir/StructureDefinition/resource-pertainsToGoal](http://hl7.org/fhir/extensions/StructureDefinition-resource-pertainsToGoal.html)) that allows goals to be referenced from any of the referenced resources in CarePlan.activity.plannedActivityReference.
	//
	// The goal should be visible when the resource referenced by CarePlan.activity.plannedActivityReference is viewed independently from the CarePlan.  Requests that are pointed to by a CarePlan using this element should *not* point to this CarePlan using the "basedOn" element.  i.e. Requests that are part of a CarePlan are not "based on" the CarePlan.
	PlannedActivityReference *Reference `bson:"plannedActivityReference,omitempty" json:"plannedActivityReference,omitempty" xml:"plannedActivityReference,omitempty"`
}

func (r CarePlan) ResourceType() string {
	return "CarePlan"
}

// MarshalJSON implements json.Marshaler
func (r CarePlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		CarePlan
		ResourceType string `json:"resourceType"`
	}{
		CarePlan:     r,
		ResourceType: r.ResourceType(),
	})
}
