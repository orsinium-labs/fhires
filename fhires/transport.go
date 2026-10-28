package fhires

import (
	"encoding/json/v2"

	decimal "github.com/shopspring/decimal"
)

// Record of transport.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Transport
type Transport struct {
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
	// Identifier for the transport event that is used to identify it across multiple disparate systems.
	//
	// Comment: This identifier is typically assigned by the dispenser, and may be used to reference the delivery when exchanging information about it with other systems.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The URL pointing to a *FHIR*-defined protocol, guideline, orderset or other definition that is adhered to in whole or in part by this Transport.
	InstantiatesCanonical *string `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty" xml:"instantiatesCanonical,omitempty"`
	// The URL pointing to an *externally* maintained  protocol, guideline, orderset or other definition that is adhered to in whole or in part by this Transport.
	InstantiatesUri *string `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty" xml:"instantiatesUri,omitempty"`
	// BasedOn refers to a higher-level authorization that triggered the creation of the transport.
	//
	//  It references a "request" resource such as a ServiceRequest or Transport, which is distinct from the "request" resource the Transport is seeking to fulfill.  This latter resource is referenced by FocusOn.  For example, based on a ServiceRequest (= BasedOn), a transport is created to fulfill a procedureRequest ( = FocusOn ) to transport a specimen to the lab.
	BasedOn []Reference `bson:"basedOn,omitempty" json:"basedOn,omitempty" xml:"basedOn,omitempty"`
	// A shared identifier common to multiple independent Request instances that were activated/authorized more or less simultaneously by a single author.
	//
	//  The presence of the same identifier on each request ties those requests together and may have business ramifications in terms of reporting of results, billing, etc.  E.g. a requisition number shared by a set of lab tests ordered together, or a prescription number shared by all meds ordered at one time.
	GroupIdentifier *Identifier `bson:"groupIdentifier,omitempty" json:"groupIdentifier,omitempty" xml:"groupIdentifier,omitempty"`
	// A larger event of which this particular event is a component or step.
	//
	// Comment: Not to be used to link an event to an Encounter - use Event.context for that.[The allowed reference resources may be adjusted as appropriate for the event resource].
	PartOf []Reference `bson:"partOf,omitempty" json:"partOf,omitempty" xml:"partOf,omitempty"`
	// A code specifying the state of the transport event.
	Status *TransportStatus `bson:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// An explanation as to why this transport is held, failed, was refused, etc.
	//
	// Comment: This applies to the current status.  Look at the history of the transport to see reasons for past statuses.
	StatusReason *CodeableConcept `bson:"statusReason,omitempty" json:"statusReason,omitempty" xml:"statusReason,omitempty"`
	// Indicates the "level" of actionability associated with the Transport, i.e.
	//
	// i+R[9]Cs this a proposed transport, a planned transport, an actionable transport, etc.
	//
	// Comment: This element is immutable.  Proposed transports, planned transports, etc. must be distinct instances.
	//
	// In most cases, Transports will have an intent of "order".
	Intent string `bson:"intent" json:"intent" xml:"intent"`
	// Indicates how quickly the Transport should be addressed with respect to other requests.
	Priority *RequestPriority `bson:"priority,omitempty" json:"priority,omitempty" xml:"priority,omitempty"`
	// A name or code (or both) briefly describing what the transport involves.
	//
	// Comment: The title (eg "My Transports", "Outstanding Transports for Patient X") should go into the code.
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// A free-text description of what is to be performed.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// The request being actioned or the resource being manipulated by this transport.
	//
	// Comment: If multiple resources need to be manipulated, use sub-transports.  (This ensures that status can be tracked independently for each referenced resource.).
	Focus *Reference `bson:"focus,omitempty" json:"focus,omitempty" xml:"focus,omitempty"`
	// The entity who benefits from the performance of the service specified in the transport (e.g., the patient).
	For *Reference `bson:"for,omitempty" json:"for,omitempty" xml:"for,omitempty"`
	// The healthcare event  (e.g.
	//
	// a patient and healthcare provider interaction) during which this transport was created.
	Encounter *Reference `bson:"encounter,omitempty" json:"encounter,omitempty" xml:"encounter,omitempty"`
	// Identifies the completion time of the event (the occurrence).
	CompletionTime *string `bson:"completionTime,omitempty" json:"completionTime,omitempty" xml:"completionTime,omitempty"`
	// The date and time this transport was created.
	AuthoredOn *string `bson:"authoredOn,omitempty" json:"authoredOn,omitempty" xml:"authoredOn,omitempty"`
	// The date and time of last modification to this transport.
	LastModified *string `bson:"lastModified,omitempty" json:"lastModified,omitempty" xml:"lastModified,omitempty"`
	// The creator of the transport.
	Requester *Reference `bson:"requester,omitempty" json:"requester,omitempty" xml:"requester,omitempty"`
	// The kind of participant that should perform the transport.
	PerformerType []CodeableConcept `bson:"performerType,omitempty" json:"performerType,omitempty" xml:"performerType,omitempty"`
	// Individual organization or Device currently responsible for transport execution.
	//
	// Comment: Transports may be created with an owner not yet identified.
	Owner *Reference `bson:"owner,omitempty" json:"owner,omitempty" xml:"owner,omitempty"`
	// Principal physical location where this transport is performed.
	Location *Reference `bson:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// Insurance plans, coverage extensions, pre-authorizations and/or pre-determinations that may be relevant to the Transport.
	Insurance []Reference `bson:"insurance,omitempty" json:"insurance,omitempty" xml:"insurance,omitempty"`
	// Free-text information captured about the transport as it progresses.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// Links to Provenance records for past versions of this Transport that identify key state transitions or updates that are likely to be relevant to a user looking at the current version of the transport.
	//
	// Comment: This element does not point to the Provenance associated with the *current* version of the resource - as it would be created after this version existed.  The Provenance for the current version can be retrieved with a _revinclude.
	RelevantHistory []Reference `bson:"relevantHistory,omitempty" json:"relevantHistory,omitempty" xml:"relevantHistory,omitempty"`
	// If the Transport.focus is a request resource and the transport is seeking fulfillment (i.e.
	//
	// is asking for the request to be actioned), this element identifies any limitations on what parts of the referenced request should be actioned.
	Restriction *TransportRestriction `bson:"restriction,omitempty" json:"restriction,omitempty" xml:"restriction,omitempty"`
	// Additional information that may be needed in the execution of the transport.
	Input []TransportInput `bson:"input,omitempty" json:"input,omitempty" xml:"input,omitempty"`
	// Outputs produced by the Transport.
	Output []TransportOutput `bson:"output,omitempty" json:"output,omitempty" xml:"output,omitempty"`
	// The desired or final location for the transport.
	RequestedLocation Reference `bson:"requestedLocation" json:"requestedLocation" xml:"requestedLocation"`
	// The current location for the entity to be transported.
	CurrentLocation Reference `bson:"currentLocation" json:"currentLocation" xml:"currentLocation"`
	// A resource reference indicating why this transport needs to be performed.
	//
	// Comment: Transports might be justified based on an Observation, a Condition, a past or planned procedure, etc. This should only be included if there is no focus or if it differs from the reason indicated on the focus.    Use the CodeableConcept text element in `Transport.reasonCode` if the data is free (uncoded) text.
	Reason *CodeableReference `bson:"reason,omitempty" json:"reason,omitempty" xml:"reason,omitempty"`
	// The transport event prior to this one.
	History *Reference `bson:"history,omitempty" json:"history,omitempty" xml:"history,omitempty"`
}

type TransportRestriction struct {
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
	// Over what time-period is fulfillment sought.
	//
	// Comment: Note that period.high is the due date representing the time by which the transport should be completed.
	Period *Period `bson:"period,omitempty" json:"period,omitempty" xml:"period,omitempty"`
	// For requests that are targeted to more than one potential recipient/target, to identify who is fulfillment is sought for.
	Recipient []Reference `bson:"recipient,omitempty" json:"recipient,omitempty" xml:"recipient,omitempty"`
}

type TransportInput struct {
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
	// A code or description indicating how the input is intended to be used as part of the transport execution.
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

type TransportOutput struct {
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

func (r Transport) ResourceType() string {
	return "Transport"
}

// MarshalJSON implements json.Marshaler
func (r Transport) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Transport
		ResourceType string `json:"resourceType"`
	}{
		ResourceType: r.ResourceType(),
		Transport:    r,
	})
}
