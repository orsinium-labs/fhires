package fhires

import v2 "encoding/json/v2"

// A set of related requests that can be used to capture intended activities that have inter-dependencies such as "give this medication after that one".
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/RequestOrchestration
type RequestOrchestration struct {
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
	// Allows a service to provide a unique, business identifier for the request.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// A canonical URL referencing a FHIR-defined protocol, guideline, orderset or other definition that is adhered to in whole or in part by this request.
	InstantiatesCanonical []string `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty" xml:"instantiatesCanonical,omitempty"`
	// A URL referencing an externally defined protocol, guideline, orderset or other definition that is adhered to in whole or in part by this request.
	InstantiatesUri []string `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty" xml:"instantiatesUri,omitempty"`
	// A plan, proposal or order that is fulfilled in whole or in part by this request.
	BasedOn []Reference `bson:"basedOn,omitempty" json:"basedOn,omitempty" xml:"basedOn,omitempty"`
	// Completed or terminated request(s) whose function is taken by this new request.
	//
	// Comment: The replacement could be because the initial request was immediately rejected (due to an issue) or because the previous request was completed, but the need for the action described by the request remains ongoing.
	Replaces []Reference `bson:"replaces,omitempty" json:"replaces,omitempty" xml:"replaces,omitempty"`
	// A shared identifier common to multiple independent Request instances that were activated/authorized more or less simultaneously by a single author.
	//
	//  The presence of the same identifier on each request ties those requests together and may have business ramifications in terms of reporting of results, billing, etc.  E.g. a requisition number shared by a set of lab tests ordered together, or a prescription number shared by all meds ordered at one time.
	//
	// Comment: Requests are linked either by a "basedOn" relationship (i.e. one request is fulfilling another) or by having a common requisition.  Requests that are part of the same requisition are generally treated independently from the perspective of changing their state or maintaining them after initial creation.
	GroupIdentifier *Identifier `bson:"groupIdentifier,omitempty" json:"groupIdentifier,omitempty" xml:"groupIdentifier,omitempty"`
	// The current state of the request.
	//
	// For request orchestrations, the status reflects the status of all the requests in the orchestration.
	Status RequestStatus `bson:"status" json:"status" xml:"status"`
	// Indicates the level of authority/intentionality associated with the request and where the request fits into the workflow chain.
	Intent RequestIntent `bson:"intent" json:"intent" xml:"intent"`
	// Indicates how quickly the request should be addressed with respect to other requests.
	Priority *RequestPriority `bson:"priority,omitempty" json:"priority,omitempty" xml:"priority,omitempty"`
	// A code that identifies what the overall request orchestration is.
	//
	// Comment: This element can be used to provide a code that captures the meaning of the request orchestration as a whole, as opposed to the code of the action element, which captures the meaning of the individual actions within the request orchestration.
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// The subject for which the request orchestration was created.
	Subject *Reference `bson:"subject,omitempty" json:"subject,omitempty" xml:"subject,omitempty"`
	// Describes the context of the request orchestration, if any.
	Encounter *Reference `bson:"encounter,omitempty" json:"encounter,omitempty" xml:"encounter,omitempty"`
	// Indicates when the request orchestration was created.
	AuthoredOn *string `bson:"authoredOn,omitempty" json:"authoredOn,omitempty" xml:"authoredOn,omitempty"`
	// Provides a reference to the author of the request orchestration.
	Author *Reference `bson:"author,omitempty" json:"author,omitempty" xml:"author,omitempty"`
	// Describes the reason for the request orchestration in coded or textual form.
	Reason []CodeableReference `bson:"reason,omitempty" json:"reason,omitempty" xml:"reason,omitempty"`
	// Goals that are intended to be achieved by following the requests in this RequestOrchestration.
	Goal []Reference `bson:"goal,omitempty" json:"goal,omitempty" xml:"goal,omitempty"`
	// Provides a mechanism to communicate additional information about the response.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// The actions, if any, produced by the evaluation of the artifact.
	Action []RequestOrchestrationAction `bson:"action,omitempty" json:"action,omitempty" xml:"action,omitempty"`
}

type RequestOrchestrationAction struct {
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
	// The linkId of the action from the PlanDefinition that corresponds to this action in the RequestOrchestration resource.
	LinkID *string `bson:"linkId,omitempty" json:"linkId,omitempty" xml:"linkId,omitempty"`
	// A user-visible prefix for the action.
	//
	// For example a section or item numbering such as 1. or A.
	Prefix *string `bson:"prefix,omitempty" json:"prefix,omitempty" xml:"prefix,omitempty"`
	// The title of the action displayed to a user.
	Title *string `bson:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// A short description of the action used to provide a summary to display to the user.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// A text equivalent of the action to be performed.
	//
	// This provides a human-interpretable description of the action when the definition is consumed by a system that might not be capable of interpreting it dynamically.
	TextEquivalent *string `bson:"textEquivalent,omitempty" json:"textEquivalent,omitempty" xml:"textEquivalent,omitempty"`
	// Indicates how quickly the action should be addressed with respect to other actions.
	Priority *RequestPriority `bson:"priority,omitempty" json:"priority,omitempty" xml:"priority,omitempty"`
	// A code that provides meaning for the action or action group.
	//
	// For example, a section may have a LOINC code for a section of a documentation template.
	Code []CodeableConcept `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// Didactic or other informational resources associated with the action that can be provided to the CDS recipient.
	//
	// Information resources can include inline text commentary and links to web resources.
	Documentation []RelatedArtifact `bson:"documentation,omitempty" json:"documentation,omitempty" xml:"documentation,omitempty"`
	// Goals that are intended to be achieved by following the requests in this action.
	Goal []Reference `bson:"goal,omitempty" json:"goal,omitempty" xml:"goal,omitempty"`
	// An expression that describes applicability criteria, or start/stop conditions for the action.
	//
	// Comment: When multiple conditions of the same kind are present, the effects are combined using AND semantics, so the overall condition is true only if all of the conditions are true.
	Condition []RequestOrchestrationActionCondition `bson:"condition,omitempty" json:"condition,omitempty" xml:"condition,omitempty"`
	// Defines input data requirements for the action.
	Input []RequestOrchestrationActionInput `bson:"input,omitempty" json:"input,omitempty" xml:"input,omitempty"`
	// Defines the outputs of the action, if any.
	Output []RequestOrchestrationActionOutput `bson:"output,omitempty" json:"output,omitempty" xml:"output,omitempty"`
	// A relationship to another action such as "before" or "30-60 minutes after start of".
	RelatedAction []RequestOrchestrationActionRelatedAction `bson:"relatedAction,omitempty" json:"relatedAction,omitempty" xml:"relatedAction,omitempty"`
	// An optional value describing when the action should be performed.
	//
	// Comment: The intent of the timing element is to provide timing for when the action should be performed. The timing may be absolute (specified as a dateTime or Period) or relative (specified as an Age, Duration, or Range), or it may be a more complex, potentially repeating timing specified using Timing.
	TimingDateTime *string   `bson:"timingDateTime,omitempty" json:"timingDateTime,omitempty" xml:"timingDateTime,omitempty"`
	TimingAge      *Age      `bson:"timingAge,omitempty" json:"timingAge,omitempty" xml:"timingAge,omitempty"`
	TimingPeriod   *Period   `bson:"timingPeriod,omitempty" json:"timingPeriod,omitempty" xml:"timingPeriod,omitempty"`
	TimingDuration *Duration `bson:"timingDuration,omitempty" json:"timingDuration,omitempty" xml:"timingDuration,omitempty"`
	TimingRange    *Range    `bson:"timingRange,omitempty" json:"timingRange,omitempty" xml:"timingRange,omitempty"`
	TimingTiming   *Timing   `bson:"timingTiming,omitempty" json:"timingTiming,omitempty" xml:"timingTiming,omitempty"`
	// Identifies the facility where the action will occur; e.g.
	//
	// home, hospital, specific clinic, etc.
	//
	// Comment: May reference a specific clinical location or may just identify a type of location.
	Location *CodeableReference `bson:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// The participant that should perform or be responsible for this action.
	//
	// Comment: Because request orchestrations represent potential options for performing activities, some specific participants may still be unknown, so this element allows for both definitional participants (in the same way they are specified in ActivityDefinition and PlanDefinition resources) as well as identifying specific participants when they are known.
	Participant []RequestOrchestrationActionParticipant `bson:"participant,omitempty" json:"participant,omitempty" xml:"participant,omitempty"`
	// The type of action to perform (create, update, remove).
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Defines the grouping behavior for the action and its children.
	GroupingBehavior *ActionGroupingBehavior `bson:"groupingBehavior,omitempty" json:"groupingBehavior,omitempty" xml:"groupingBehavior,omitempty"`
	// Defines the selection behavior for the action and its children.
	SelectionBehavior *ActionSelectionBehavior `bson:"selectionBehavior,omitempty" json:"selectionBehavior,omitempty" xml:"selectionBehavior,omitempty"`
	// Defines expectations around whether an action is required.
	RequiredBehavior *ActionRequiredBehavior `bson:"requiredBehavior,omitempty" json:"requiredBehavior,omitempty" xml:"requiredBehavior,omitempty"`
	// Defines whether the action should usually be preselected.
	PrecheckBehavior *ActionPrecheckBehavior `bson:"precheckBehavior,omitempty" json:"precheckBehavior,omitempty" xml:"precheckBehavior,omitempty"`
	// Defines whether the action can be selected multiple times.
	CardinalityBehavior *ActionCardinalityBehavior `bson:"cardinalityBehavior,omitempty" json:"cardinalityBehavior,omitempty" xml:"cardinalityBehavior,omitempty"`
	// The resource that is the target of the action (e.g. CommunicationRequest).
	//
	// Comment: The target resource SHALL be a [Request](request.html) resource with a Request.intent set to "option".
	Resource *Reference `bson:"resource,omitempty" json:"resource,omitempty" xml:"resource,omitempty"`
	// A reference to an ActivityDefinition that describes the action to be taken in detail, a PlanDefinition that describes a series of actions to be taken, a Questionnaire that should be filled out, a SpecimenDefinition describing a specimen to be collected, or an ObservationDefinition that specifies what observation should be captured.
	//
	// Comment: Note that the definition is optional, and if no definition is specified, a dynamicValue with a root ($this) path can be used to define the entire resource dynamically.
	DefinitionCanonical *string `bson:"definitionCanonical,omitempty" json:"definitionCanonical,omitempty" xml:"definitionCanonical,omitempty"`
	DefinitionUri       *string `bson:"definitionUri,omitempty" json:"definitionUri,omitempty" xml:"definitionUri,omitempty"`
	// A reference to a StructureMap resource that defines a transform that can be executed to produce the intent resource using the ActivityDefinition instance as the input.
	//
	// Comment: Note that when a referenced ActivityDefinition also defines a transform, the transform specified here generally takes precedence. In addition, if both a transform and dynamic values are specific, the dynamic values are applied to the result of the transform.
	Transform *string `bson:"transform,omitempty" json:"transform,omitempty" xml:"transform,omitempty"`
	// Customizations that should be applied to the statically defined resource.
	//
	// For example, if the dosage of a medication must be computed based on the patient's weight, a customization would be used to specify an expression that calculated the weight, and the path on the resource that would contain the result.
	//
	// Comment: Dynamic values are applied in the order in which they are defined in the RequestOrchestration resource. Note that when dynamic values are also specified by a referenced ActivityDefinition, the dynamicValues from the ActivityDefinition are applied first, followed by the dynamicValues specified here. In addition, if both a transform and dynamic values are specific, the dynamic values are applied to the result of the transform.
	DynamicValue []RequestOrchestrationActionDynamicValue `bson:"dynamicValue,omitempty" json:"dynamicValue,omitempty" xml:"dynamicValue,omitempty"`
	// Sub actions.
	Action []RequestOrchestrationAction `bson:"action,omitempty" json:"action,omitempty"`
}

type RequestOrchestrationActionCondition struct {
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
	// The kind of condition.
	//
	// Comment: Applicability criteria are used to determine immediate applicability when a plan definition is applied to a given context. Start and stop criteria are carried through application and used to describe enter/exit criteria for an action.
	Kind ActionConditionKind `bson:"kind" json:"kind" xml:"kind"`
	// An expression that returns true or false, indicating whether or not the condition is satisfied.
	//
	// Comment: The expression may be inlined, or may be a reference to a named expression within a logic library referenced by the library element.
	Expression *Expression `bson:"expression,omitempty" json:"expression,omitempty" xml:"expression,omitempty"`
}

type RequestOrchestrationActionInput struct {
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
	// A human-readable label for the data requirement used to label data flows in BPMN or similar diagrams.
	//
	// Also provides a human readable label when rendering the data requirement that conveys its purpose to human readers.
	Title *string `bson:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Defines the data that is to be provided as input to the action.
	Requirement *DataRequirement `bson:"requirement,omitempty" json:"requirement,omitempty" xml:"requirement,omitempty"`
	// Points to an existing input or output element that provides data to this input.
	//
	// Comment: The relatedData element allows indicating that an input to a parent action is an input to specific child actions. It also allows the output of one action to be identified as the input to a different action
	RelatedData *string `bson:"relatedData,omitempty" json:"relatedData,omitempty" xml:"relatedData,omitempty"`
}

type RequestOrchestrationActionOutput struct {
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
	// A human-readable label for the data requirement used to label data flows in BPMN or similar diagrams.
	//
	// Also provides a human readable label when rendering the data requirement that conveys its purpose to human readers.
	Title *string `bson:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Defines the data that results as output from the action.
	Requirement *DataRequirement `bson:"requirement,omitempty" json:"requirement,omitempty" xml:"requirement,omitempty"`
	// Points to an existing input or output element that is results as output from the action.
	//
	// Comment: The relatedData element allows indicating that the output of child action is also the output of a parent action. It also allows the input of one action to be identified as the output of a different action
	RelatedData *string `bson:"relatedData,omitempty" json:"relatedData,omitempty" xml:"relatedData,omitempty"`
}

type RequestOrchestrationActionRelatedAction struct {
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
	// The element id of the target related action.
	TargetID string `bson:"targetId" json:"targetId" xml:"targetId"`
	// The relationship of this action to the related action.
	Relationship ActionRelationshipType `bson:"relationship" json:"relationship" xml:"relationship"`
	// The relationship of the end of this action to the related action.
	EndRelationship *ActionRelationshipType `bson:"endRelationship,omitempty" json:"endRelationship,omitempty" xml:"endRelationship,omitempty"`
	// A duration or range of durations to apply to the relationship.
	//
	// For example, 30-60 minutes before.
	OffsetDuration *Duration `bson:"offsetDuration,omitempty" json:"offsetDuration,omitempty" xml:"offsetDuration,omitempty"`
	OffsetRange    *Range    `bson:"offsetRange,omitempty" json:"offsetRange,omitempty" xml:"offsetRange,omitempty"`
}

type RequestOrchestrationActionParticipant struct {
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
	// The type of participant in the action.
	Type *ActionParticipantType `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The type of participant in the action.
	TypeCanonical *string `bson:"typeCanonical,omitempty" json:"typeCanonical,omitempty" xml:"typeCanonical,omitempty"`
	// The type of participant in the action.
	//
	// Comment: When this element is a reference, it SHOULD be a reference to a definitional resource (for example, a location type, rather than a specific location).
	TypeReference *Reference `bson:"typeReference,omitempty" json:"typeReference,omitempty" xml:"typeReference,omitempty"`
	// The role the participant should play in performing the described action.
	Role *CodeableConcept `bson:"role,omitempty" json:"role,omitempty" xml:"role,omitempty"`
	// Indicates how the actor will be involved in the action - author, reviewer, witness, etc.
	Function *CodeableConcept `bson:"function,omitempty" json:"function,omitempty" xml:"function,omitempty"`
	// A reference to the actual participant.
	ActorCanonical *string    `bson:"actorCanonical,omitempty" json:"actorCanonical,omitempty" xml:"actorCanonical,omitempty"`
	ActorReference *Reference `bson:"actorReference,omitempty" json:"actorReference,omitempty" xml:"actorReference,omitempty"`
}

type RequestOrchestrationActionDynamicValue struct {
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
	// The path to the element to be customized.
	//
	// This is the path on the resource that will hold the result of the calculation defined by the expression. The specified path SHALL be a FHIRPath resolvable on the specified target type of the ActivityDefinition, and SHALL consist only of identifiers, constant indexers, and a restricted subset of functions. The path is allowed to contain qualifiers (.) to traverse sub-elements, as well as indexers ([x]) to traverse multiple-cardinality sub-elements (see the [Simple FHIRPath Profile](fhirpath.html#simple) for full details).
	//
	// Comment: To specify the path to the current action being realized, the %action environment variable is available in this path. For example, to specify the description element of the target action, the path would be %action.description. The path attribute contains a [Simple FHIRPath Subset](fhirpath.html#simple) that allows path traversal, but not calculation.
	Path *string `bson:"path,omitempty" json:"path,omitempty" xml:"path,omitempty"`
	// An expression specifying the value of the customized element.
	//
	// Comment: The expression may be inlined or may be a reference to a named expression within a logic library referenced by the library element.
	Expression *Expression `bson:"expression,omitempty" json:"expression,omitempty" xml:"expression,omitempty"`
}

func (r RequestOrchestration) ResourceType() string {
	return "RequestOrchestration"
}

// MarshalJSON implements json.Marshaler
func (r RequestOrchestration) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		RequestOrchestration
		ResourceType string `json:"resourceType"`
	}{
		RequestOrchestration: r,
		ResourceType:         r.ResourceType(),
	})
}
