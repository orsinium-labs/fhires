package fhires

import (
	"encoding/json"

	decimal "github.com/shopspring/decimal"
)

// A task to be performed.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Task
type Task struct {
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
	// The business identifier for this task.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The URL pointing to a *FHIR*-defined protocol, guideline, orderset or other definition that is adhered to in whole or in part by this Task.
	InstantiatesCanonical *string `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty" xml:"instantiatesCanonical,omitempty"`
	// The URL pointing to an *externally* maintained  protocol, guideline, orderset or other definition that is adhered to in whole or in part by this Task.
	InstantiatesUri *string `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty" xml:"instantiatesUri,omitempty"`
	// BasedOn refers to a higher-level authorization that triggered the creation of the task.
	//
	//  It references a "request" resource such as a ServiceRequest, MedicationRequest, CarePlan, etc. which is distinct from the "request" resource the task is seeking to fulfill.  This latter resource is referenced by focus.  For example, based on a CarePlan (= basedOn), a task is created to fulfill a ServiceRequest ( = focus ) to collect a specimen from a patient.
	//
	// Comment: Task.basedOn is never the same as Task.focus.  Task.basedOn will typically not be present for 'please fulfill' Tasks as a distinct authorization is rarely needed to request fulfillment.  If the Task is seeking fulfillment of an order, the order to be fulfilled is always communicated using `focus`, never basedOn.  However, authorization may be needed to perform other types of Task actions.  As an example of when both would be present, a Task seeking suspension of a prescription might have a Task.basedOn pointing to the ServiceRequest ordering surgery (which is the driver for suspending the MedicationRequest - which would be the Task.focus).
	BasedOn []Reference `bson:"basedOn,omitempty" json:"basedOn,omitempty" xml:"basedOn,omitempty"`
	// A shared identifier common to multiple independent Task and Request instances that were activated/authorized more or less simultaneously by a single author.
	//
	//  The presence of the same identifier on each request ties those requests together and may have business ramifications in terms of reporting of results, billing, etc.  E.g. a requisition number shared by a set of lab tests ordered together, or a prescription number shared by all meds ordered at one time.
	GroupIdentifier *Identifier `bson:"groupIdentifier,omitempty" json:"groupIdentifier,omitempty" xml:"groupIdentifier,omitempty"`
	// Task that this particular task is part of.
	//
	// Comment: This should usually be 0..1.
	PartOf []Reference `bson:"partOf,omitempty" json:"partOf,omitempty" xml:"partOf,omitempty"`
	// The current status of the task.
	Status TaskStatus `bson:"status" json:"status" xml:"status"`
	// An explanation as to why this task is held, failed, was refused, etc.
	//
	// Comment: This applies to the current status.  Look at the history of the task to see reasons for past statuses.
	StatusReason *CodeableReference `bson:"statusReason,omitempty" json:"statusReason,omitempty" xml:"statusReason,omitempty"`
	// Contains business-specific nuances of the business state.
	BusinessStatus *CodeableConcept `bson:"businessStatus,omitempty" json:"businessStatus,omitempty" xml:"businessStatus,omitempty"`
	// Indicates the "level" of actionability associated with the Task, i.e.
	//
	// i+R[9]Cs this a proposed task, a planned task, an actionable task, etc.
	//
	// Comment: This element is immutable.  Proposed tasks, planned tasks, etc. must be distinct instances.
	//
	// In most cases, Tasks will have an intent of "order".
	Intent string `bson:"intent" json:"intent" xml:"intent"`
	// Indicates how quickly the Task should be addressed with respect to other requests.
	Priority *RequestPriority `bson:"priority,omitempty" json:"priority,omitempty" xml:"priority,omitempty"`
	// If true indicates that the Task is asking for the specified action to *not* occur.
	//
	// Comment: The attributes provided with the Task qualify what is not to be done. For example, if a requestedPeriod is provided, the 'do not' request only applies within the specified time. If a requestedPerformer is specified then the 'do not' request only applies to performers of that type. Qualifiers include: code, subject, occurrence, requestedPerformer and performer.
	//
	// In some cases, the Request.code may pre-coordinate prohibition into the requested action. E.g. 'NPO' (nothing by mouth), 'DNR' (do not recussitate). If this happens, doNotPerform SHALL NOT be set to true. I.e. The resource shall not have double negation. (E.g. 'Do not DNR').
	//
	// doNotPerform should ONLY be used with Tasks that are tightly bounded in time or process phase.  E.g. 'Do not fulfill the midnight dose of medication X tonight due to the early morning scheduled procedure, where the nurse could reasonably check off 'Med X not given at midnight as instructed'.  Similarly, a decision support proposal that a patient should not be given a standard intake questionnaire (because the patient is cognitively impaired) would be marked as 'complete' or 'rejected' when the clinician preps the CarePlan or order set after reviewing the decision support results.  If there is a need to create a standing order to not do something that can't be satisfied by a single 'non-action', but rather an ongoing refusal to perform the function, MedicationRequest, ServiceRequest or some other form of authorization should be used.
	DoNotPerform *bool `bson:"doNotPerform,omitempty" json:"doNotPerform,omitempty" xml:"doNotPerform,omitempty"`
	// A name or code (or both) briefly describing what the task involves.
	//
	// Comment: The title (eg "My Tasks", "Outstanding Tasks for Patient X") should go into the code.
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// A free-text description of what is to be performed.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// The request being fulfilled or the resource being manipulated (changed, suspended, etc.) by this task.
	//
	// Comment: If multiple resources need to be manipulated, use sub-tasks.  (This ensures that status can be tracked independently for each referenced resource.).
	Focus *Reference `bson:"focus,omitempty" json:"focus,omitempty" xml:"focus,omitempty"`
	// The entity who benefits from the performance of the service specified in the task (e.g., the patient).
	For *Reference `bson:"for,omitempty" json:"for,omitempty" xml:"for,omitempty"`
	// The healthcare event  (e.g.
	//
	// a patient and healthcare provider interaction) during which this task was created.
	Encounter *Reference `bson:"encounter,omitempty" json:"encounter,omitempty" xml:"encounter,omitempty"`
	// Indicates the start and/or end of the period of time when completion of the task is desired to take place.
	//
	// Comment: This is typically used when the Task is *not* seeking fulfillment of a focus Request, as in that case the period would be specified on the Request and/or in the Task.restriction.period.  Instead, it is used for stand-alone tasks.
	RequestedPeriod *Period `bson:"requestedPeriod,omitempty" json:"requestedPeriod,omitempty" xml:"requestedPeriod,omitempty"`
	// Identifies the time action was first taken against the task (start) and/or the time final action was taken against the task prior to marking it as completed (end).
	ExecutionPeriod *Period `bson:"executionPeriod,omitempty" json:"executionPeriod,omitempty" xml:"executionPeriod,omitempty"`
	// The date and time this task was created.
	AuthoredOn *string `bson:"authoredOn,omitempty" json:"authoredOn,omitempty" xml:"authoredOn,omitempty"`
	// The date and time of last modification to this task.
	LastModified *string `bson:"lastModified,omitempty" json:"lastModified,omitempty" xml:"lastModified,omitempty"`
	// The creator of the task.
	Requester *Reference `bson:"requester,omitempty" json:"requester,omitempty" xml:"requester,omitempty"`
	// The kind of participant or specific participant that should perform the task.
	RequestedPerformer []CodeableReference `bson:"requestedPerformer,omitempty" json:"requestedPerformer,omitempty" xml:"requestedPerformer,omitempty"`
	// Party responsible for managing task execution.
	//
	// Comment: Tasks may be created with an owner not yet identified.
	Owner *Reference `bson:"owner,omitempty" json:"owner,omitempty" xml:"owner,omitempty"`
	// The entity who performed the requested task.
	Performer []TaskPerformer `bson:"performer,omitempty" json:"performer,omitempty" xml:"performer,omitempty"`
	// Principal physical location where this task is performed.
	//
	// Comment: This should only be specified when the Task to be/being performed happens or is expected to happen primarily within the bounds of a single Location.  Other locations (e.g. source, destination, etc.) would either be reflected on the 'basedOn' Request or be conveyed as distinct Task.input values.
	Location *Reference `bson:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// A description, code, or reference indicating why this task needs to be performed.
	//
	// Comment: This will typically not be present for Tasks with a code of 'please fulfill' as, for those, the reason for action is conveyed on the Request pointed to by Task.focus.  Some types of tasks will not need a 'reason'.  E.g. a request to discharge a patient can be inferred to be 'because the patient is ready' and this would not need a reason to be stated on the Task.
	Reason []CodeableReference `bson:"reason,omitempty" json:"reason,omitempty" xml:"reason,omitempty"`
	// Insurance plans, coverage extensions, pre-authorizations and/or pre-determinations that may be relevant to the Task.
	Insurance []Reference `bson:"insurance,omitempty" json:"insurance,omitempty" xml:"insurance,omitempty"`
	// Free-text information captured about the task as it progresses.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// Links to Provenance records for past versions of this Task that identify key state transitions or updates that are likely to be relevant to a user looking at the current version of the task.
	//
	// Comment: This element does not point to the Provenance associated with the *current* version of the resource - as it would be created after this version existed.  The Provenance for the current version can be retrieved with a _revinclude.
	RelevantHistory []Reference `bson:"relevantHistory,omitempty" json:"relevantHistory,omitempty" xml:"relevantHistory,omitempty"`
	// If the Task.focus is a request resource and the task is seeking fulfillment (i.e.
	//
	// is asking for the request to be actioned), this element identifies any limitations on what parts of the referenced request should be actioned.
	//
	// Comment: Task.restriction can only be present if the Task is seeking fulfillment of another Request resource, and the restriction identifies what subset of the authorization conveyed by the request is supposed to be fulfilled by this Task. A possible example could be a standing order (the request) covering a significant time period and/or individuals, while the Task seeks fulfillment for only a subset of that time-period and a single individual.
	Restriction *TaskRestriction `bson:"restriction,omitempty" json:"restriction,omitempty" xml:"restriction,omitempty"`
	// Additional information that may be needed in the execution of the task.
	Input []TaskInput `bson:"input,omitempty" json:"input,omitempty" xml:"input,omitempty"`
	// Outputs produced by the Task.
	Output []TaskOutput `bson:"output,omitempty" json:"output,omitempty" xml:"output,omitempty"`
}

type TaskPerformer struct {
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
	// A code or description of the performer of the task.
	Function *CodeableConcept `bson:"function,omitempty" json:"function,omitempty" xml:"function,omitempty"`
	// The actor or entity who performed the task.
	Actor Reference `bson:"actor" json:"actor" xml:"actor"`
}

type TaskRestriction struct {
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
	// Indicates the number of times the requested action should occur.
	Repetitions *int `bson:"repetitions,omitempty" json:"repetitions,omitempty" xml:"repetitions,omitempty"`
	// The time-period for which fulfillment is sought.
	//
	// This must fall within the overall time period authorized in the referenced request.  E.g. ServiceRequest.occurance[x].
	//
	// Comment: This is distinct from Task.executionPeriod. ExecutionPeriod indicates when the task needs to be initiated, while Task.restriction.period specifies the subset of the overall authorization that this period covers. For example, a MedicationRequest with an overall effective period of 1 year might have a Task whose restriction.period is 2 months (i.e. satisfy 2 months of medication therapy), while the execution period might be 'between now and 5 days from now' - i.e. If you say yes to this, then you're agreeing to supply medication for that 2 month period within the next 5 days.
	//
	// Note that period.high is the due date representing the time by which the task should be completed.
	Period *Period `bson:"period,omitempty" json:"period,omitempty" xml:"period,omitempty"`
	// For requests that are targeted to more than one potential recipient/target, to identify who is fulfillment is sought for.
	Recipient []Reference `bson:"recipient,omitempty" json:"recipient,omitempty" xml:"recipient,omitempty"`
}

type TaskInput struct {
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
	// A code or description indicating how the input is intended to be used as part of the task execution.
	//
	// Comment: If referencing a BPMN workflow or Protocol, the "system" is the URL for the workflow definition and the code is the "name" of the required input.
	Type CodeableConcept `bson:"type" json:"type" xml:"type"`
	// The value of the input parameter as a basic type.
	ValueBase64Binary          string                `bson:"valueBase64Binary" json:"valueBase64Binary" xml:"valueBase64Binary"`
	ValueBoolean               bool                  `bson:"valueBoolean" json:"valueBoolean" xml:"valueBoolean"`
	ValueCanonical             string                `bson:"valueCanonical" json:"valueCanonical" xml:"valueCanonical"`
	ValueCode                  string                `bson:"valueCode" json:"valueCode" xml:"valueCode"`
	ValueDate                  string                `bson:"valueDate" json:"valueDate" xml:"valueDate"`
	ValueDateTime              string                `bson:"valueDateTime" json:"valueDateTime" xml:"valueDateTime"`
	ValueDecimal               decimal.Decimal       `bson:"valueDecimal" json:"valueDecimal" xml:"valueDecimal"`
	ValueID                    string                `bson:"valueId" json:"valueId" xml:"valueId"`
	ValueInstant               string                `bson:"valueInstant" json:"valueInstant" xml:"valueInstant"`
	ValueInteger               int                   `bson:"valueInteger" json:"valueInteger" xml:"valueInteger"`
	ValueInteger64             int64                 `bson:"valueInteger64" json:"valueInteger64,string" xml:"valueInteger64"`
	ValueMarkdown              string                `bson:"valueMarkdown" json:"valueMarkdown" xml:"valueMarkdown"`
	ValueOid                   string                `bson:"valueOid" json:"valueOid" xml:"valueOid"`
	ValuePositiveInt           int                   `bson:"valuePositiveInt" json:"valuePositiveInt" xml:"valuePositiveInt"`
	ValueString                string                `bson:"valueString" json:"valueString" xml:"valueString"`
	ValueTime                  string                `bson:"valueTime" json:"valueTime" xml:"valueTime"`
	ValueUnsignedInt           int                   `bson:"valueUnsignedInt" json:"valueUnsignedInt" xml:"valueUnsignedInt"`
	ValueUri                   string                `bson:"valueUri" json:"valueUri" xml:"valueUri"`
	ValueUURL                  string                `bson:"valueUrl" json:"valueUrl" xml:"valueUrl"`
	ValueUuid                  string                `bson:"valueUuid" json:"valueUuid" xml:"valueUuid"`
	ValueAddress               Address               `bson:"valueAddress" json:"valueAddress" xml:"valueAddress"`
	ValueAge                   Age                   `bson:"valueAge" json:"valueAge" xml:"valueAge"`
	ValueAnnotation            Annotation            `bson:"valueAnnotation" json:"valueAnnotation" xml:"valueAnnotation"`
	ValueAttachment            Attachment            `bson:"valueAttachment" json:"valueAttachment" xml:"valueAttachment"`
	ValueCodeableConcept       CodeableConcept       `bson:"valueCodeableConcept" json:"valueCodeableConcept" xml:"valueCodeableConcept"`
	ValueCodeableReference     CodeableReference     `bson:"valueCodeableReference" json:"valueCodeableReference" xml:"valueCodeableReference"`
	ValueCoding                Coding                `bson:"valueCoding" json:"valueCoding" xml:"valueCoding"`
	ValueContactPoint          ContactPoint          `bson:"valueContactPoint" json:"valueContactPoint" xml:"valueContactPoint"`
	ValueCount                 Count                 `bson:"valueCount" json:"valueCount" xml:"valueCount"`
	ValueDistance              Distance              `bson:"valueDistance" json:"valueDistance" xml:"valueDistance"`
	ValueDuration              Duration              `bson:"valueDuration" json:"valueDuration" xml:"valueDuration"`
	ValueHumanName             HumanName             `bson:"valueHumanName" json:"valueHumanName" xml:"valueHumanName"`
	ValueIdentifier            Identifier            `bson:"valueIdentifier" json:"valueIdentifier" xml:"valueIdentifier"`
	ValueMoney                 Money                 `bson:"valueMoney" json:"valueMoney" xml:"valueMoney"`
	ValuePeriod                Period                `bson:"valuePeriod" json:"valuePeriod" xml:"valuePeriod"`
	ValueQuantity              Quantity              `bson:"valueQuantity" json:"valueQuantity" xml:"valueQuantity"`
	ValueRange                 Range                 `bson:"valueRange" json:"valueRange" xml:"valueRange"`
	ValueRatio                 Ratio                 `bson:"valueRatio" json:"valueRatio" xml:"valueRatio"`
	ValueRatioRange            RatioRange            `bson:"valueRatioRange" json:"valueRatioRange" xml:"valueRatioRange"`
	ValueReference             Reference             `bson:"valueReference" json:"valueReference" xml:"valueReference"`
	ValueSampledData           SampledData           `bson:"valueSampledData" json:"valueSampledData" xml:"valueSampledData"`
	ValueSignature             Signature             `bson:"valueSignature" json:"valueSignature" xml:"valueSignature"`
	ValueTiming                Timing                `bson:"valueTiming" json:"valueTiming" xml:"valueTiming"`
	ValueContactDetail         ContactDetail         `bson:"valueContactDetail" json:"valueContactDetail" xml:"valueContactDetail"`
	ValueDataRequirement       DataRequirement       `bson:"valueDataRequirement" json:"valueDataRequirement" xml:"valueDataRequirement"`
	ValueExpression            Expression            `bson:"valueExpression" json:"valueExpression" xml:"valueExpression"`
	ValueParameterDefinition   ParameterDefinition   `bson:"valueParameterDefinition" json:"valueParameterDefinition" xml:"valueParameterDefinition"`
	ValueRelatedArtifact       RelatedArtifact       `bson:"valueRelatedArtifact" json:"valueRelatedArtifact" xml:"valueRelatedArtifact"`
	ValueTriggerDefinition     TriggerDefinition     `bson:"valueTriggerDefinition" json:"valueTriggerDefinition" xml:"valueTriggerDefinition"`
	ValueUsageContext          UsageContext          `bson:"valueUsageContext" json:"valueUsageContext" xml:"valueUsageContext"`
	ValueAvailability          Availability          `bson:"valueAvailability" json:"valueAvailability" xml:"valueAvailability"`
	ValueExtendedContactDetail ExtendedContactDetail `bson:"valueExtendedContactDetail" json:"valueExtendedContactDetail" xml:"valueExtendedContactDetail"`
	ValueDosage                Dosage                `bson:"valueDosage" json:"valueDosage" xml:"valueDosage"`
	ValueMeta                  Meta                  `bson:"valueMeta" json:"valueMeta" xml:"valueMeta"`
}

type TaskOutput struct {
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
	// The name of the Output parameter.
	Type CodeableConcept `bson:"type" json:"type" xml:"type"`
	// The value of the Output parameter as a basic type.
	ValueBase64Binary          string                `bson:"valueBase64Binary" json:"valueBase64Binary" xml:"valueBase64Binary"`
	ValueBoolean               bool                  `bson:"valueBoolean" json:"valueBoolean" xml:"valueBoolean"`
	ValueCanonical             string                `bson:"valueCanonical" json:"valueCanonical" xml:"valueCanonical"`
	ValueCode                  string                `bson:"valueCode" json:"valueCode" xml:"valueCode"`
	ValueDate                  string                `bson:"valueDate" json:"valueDate" xml:"valueDate"`
	ValueDateTime              string                `bson:"valueDateTime" json:"valueDateTime" xml:"valueDateTime"`
	ValueDecimal               decimal.Decimal       `bson:"valueDecimal" json:"valueDecimal" xml:"valueDecimal"`
	ValueID                    string                `bson:"valueId" json:"valueId" xml:"valueId"`
	ValueInstant               string                `bson:"valueInstant" json:"valueInstant" xml:"valueInstant"`
	ValueInteger               int                   `bson:"valueInteger" json:"valueInteger" xml:"valueInteger"`
	ValueInteger64             int64                 `bson:"valueInteger64" json:"valueInteger64,string" xml:"valueInteger64"`
	ValueMarkdown              string                `bson:"valueMarkdown" json:"valueMarkdown" xml:"valueMarkdown"`
	ValueOid                   string                `bson:"valueOid" json:"valueOid" xml:"valueOid"`
	ValuePositiveInt           int                   `bson:"valuePositiveInt" json:"valuePositiveInt" xml:"valuePositiveInt"`
	ValueString                string                `bson:"valueString" json:"valueString" xml:"valueString"`
	ValueTime                  string                `bson:"valueTime" json:"valueTime" xml:"valueTime"`
	ValueUnsignedInt           int                   `bson:"valueUnsignedInt" json:"valueUnsignedInt" xml:"valueUnsignedInt"`
	ValueUri                   string                `bson:"valueUri" json:"valueUri" xml:"valueUri"`
	ValueUURL                  string                `bson:"valueUrl" json:"valueUrl" xml:"valueUrl"`
	ValueUuid                  string                `bson:"valueUuid" json:"valueUuid" xml:"valueUuid"`
	ValueAddress               Address               `bson:"valueAddress" json:"valueAddress" xml:"valueAddress"`
	ValueAge                   Age                   `bson:"valueAge" json:"valueAge" xml:"valueAge"`
	ValueAnnotation            Annotation            `bson:"valueAnnotation" json:"valueAnnotation" xml:"valueAnnotation"`
	ValueAttachment            Attachment            `bson:"valueAttachment" json:"valueAttachment" xml:"valueAttachment"`
	ValueCodeableConcept       CodeableConcept       `bson:"valueCodeableConcept" json:"valueCodeableConcept" xml:"valueCodeableConcept"`
	ValueCodeableReference     CodeableReference     `bson:"valueCodeableReference" json:"valueCodeableReference" xml:"valueCodeableReference"`
	ValueCoding                Coding                `bson:"valueCoding" json:"valueCoding" xml:"valueCoding"`
	ValueContactPoint          ContactPoint          `bson:"valueContactPoint" json:"valueContactPoint" xml:"valueContactPoint"`
	ValueCount                 Count                 `bson:"valueCount" json:"valueCount" xml:"valueCount"`
	ValueDistance              Distance              `bson:"valueDistance" json:"valueDistance" xml:"valueDistance"`
	ValueDuration              Duration              `bson:"valueDuration" json:"valueDuration" xml:"valueDuration"`
	ValueHumanName             HumanName             `bson:"valueHumanName" json:"valueHumanName" xml:"valueHumanName"`
	ValueIdentifier            Identifier            `bson:"valueIdentifier" json:"valueIdentifier" xml:"valueIdentifier"`
	ValueMoney                 Money                 `bson:"valueMoney" json:"valueMoney" xml:"valueMoney"`
	ValuePeriod                Period                `bson:"valuePeriod" json:"valuePeriod" xml:"valuePeriod"`
	ValueQuantity              Quantity              `bson:"valueQuantity" json:"valueQuantity" xml:"valueQuantity"`
	ValueRange                 Range                 `bson:"valueRange" json:"valueRange" xml:"valueRange"`
	ValueRatio                 Ratio                 `bson:"valueRatio" json:"valueRatio" xml:"valueRatio"`
	ValueRatioRange            RatioRange            `bson:"valueRatioRange" json:"valueRatioRange" xml:"valueRatioRange"`
	ValueReference             Reference             `bson:"valueReference" json:"valueReference" xml:"valueReference"`
	ValueSampledData           SampledData           `bson:"valueSampledData" json:"valueSampledData" xml:"valueSampledData"`
	ValueSignature             Signature             `bson:"valueSignature" json:"valueSignature" xml:"valueSignature"`
	ValueTiming                Timing                `bson:"valueTiming" json:"valueTiming" xml:"valueTiming"`
	ValueContactDetail         ContactDetail         `bson:"valueContactDetail" json:"valueContactDetail" xml:"valueContactDetail"`
	ValueDataRequirement       DataRequirement       `bson:"valueDataRequirement" json:"valueDataRequirement" xml:"valueDataRequirement"`
	ValueExpression            Expression            `bson:"valueExpression" json:"valueExpression" xml:"valueExpression"`
	ValueParameterDefinition   ParameterDefinition   `bson:"valueParameterDefinition" json:"valueParameterDefinition" xml:"valueParameterDefinition"`
	ValueRelatedArtifact       RelatedArtifact       `bson:"valueRelatedArtifact" json:"valueRelatedArtifact" xml:"valueRelatedArtifact"`
	ValueTriggerDefinition     TriggerDefinition     `bson:"valueTriggerDefinition" json:"valueTriggerDefinition" xml:"valueTriggerDefinition"`
	ValueUsageContext          UsageContext          `bson:"valueUsageContext" json:"valueUsageContext" xml:"valueUsageContext"`
	ValueAvailability          Availability          `bson:"valueAvailability" json:"valueAvailability" xml:"valueAvailability"`
	ValueExtendedContactDetail ExtendedContactDetail `bson:"valueExtendedContactDetail" json:"valueExtendedContactDetail" xml:"valueExtendedContactDetail"`
	ValueDosage                Dosage                `bson:"valueDosage" json:"valueDosage" xml:"valueDosage"`
	ValueMeta                  Meta                  `bson:"valueMeta" json:"valueMeta" xml:"valueMeta"`
}

func (r Task) ResourceType() string {
	return "Task"
}

// MarshalJSON implements json.Marshaler
func (r Task) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Task
		ResourceType string `json:"resourceType"`
	}{
		ResourceType: r.ResourceType(),
		Task:         r,
	})
}
