package fhires

import v2 "encoding/json/v2"

// A record of a request for service such as diagnostic investigations, treatments, or operations to be performed.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/ServiceRequest
type ServiceRequest struct {
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
	// Identifiers assigned to this order instance by the orderer and/or the receiver and/or order fulfiller.
	//
	// Comment: The identifier.type element is used to distinguish between the identifiers assigned by the orderer (known as the 'Placer' in HL7 V2) and the producer of the observations in response to the order (known as the 'Filler' in HL7 V2).  For further discussion and examples see the resource notes section below.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The URL pointing to a FHIR-defined protocol, guideline, orderset or other definition that is adhered to in whole or in part by this ServiceRequest.
	//
	// Comment: The PlanDefinition resource is used to describe series, sequences, or groups of actions to be taken, while the ActivityDefinition resource is used to define each specific step or activity to be performed. More information can be found in the [Boundaries and Relationships](plandefinition.html#12.23.2) section for PlanDefinition.
	InstantiatesCanonical []string `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty" xml:"instantiatesCanonical,omitempty"`
	// The URL pointing to an externally maintained protocol, guideline, orderset or other definition that is adhered to in whole or in part by this ServiceRequest.
	//
	// Comment: This might be an HTML page, PDF, etc. or could just be a non-resolvable URI identifier.
	InstantiatesUri []string `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty" xml:"instantiatesUri,omitempty"`
	// Plan/proposal/order fulfilled by this request.
	BasedOn []Reference `bson:"basedOn,omitempty" json:"basedOn,omitempty" xml:"basedOn,omitempty"`
	// The request takes the place of the referenced completed or terminated request(s).
	Replaces []Reference `bson:"replaces,omitempty" json:"replaces,omitempty" xml:"replaces,omitempty"`
	// A shared identifier common to all service requests that were authorized more or less simultaneously by a single author, representing the composite or group identifier.
	//
	// Comment: Requests are linked either by a "basedOn" relationship (i.e. one request is fulfilling another) or by having a common requisition. Requests that are part of the same requisition are generally treated independently from the perspective of changing their state or maintaining them after initial creation.
	Requisition *Identifier `bson:"requisition,omitempty" json:"requisition,omitempty" xml:"requisition,omitempty"`
	// The status of the order.
	//
	// Comment: The status is generally fully in the control of the requester - they determine whether the order is draft or active and, after it has been activated, competed, revoked or placed on-hold. States relating to the activities of the performer are reflected on either the corresponding event (see [Event Pattern](event.html) for general discussion) or using the [Task](task.html) resource.
	Status RequestStatus `bson:"status" json:"status" xml:"status"`
	// Whether the request is a proposal, plan, an original order or a reflex order.
	//
	// Comment: This element is labeled as a modifier because the intent alters when and how the resource is actually applicable.
	Intent RequestIntent `bson:"intent" json:"intent" xml:"intent"`
	// A code that classifies the service for searching, sorting and display purposes (e.g.
	//
	// "Surgical Procedure").
	//
	// Comment: There may be multiple axis of categorization depending on the context or use case for retrieving or displaying the resource.  The level of granularity is defined by the category concepts in the value set.
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
	// Indicates how quickly the ServiceRequest should be addressed with respect to other requests.
	Priority *RequestPriority `bson:"priority,omitempty" json:"priority,omitempty" xml:"priority,omitempty"`
	// Set this to true if the record is saying that the service/procedure should NOT be performed.
	//
	// Comment: In general, only the code and timeframe will be present, though occasional additional qualifiers such as body site or even performer could be included to narrow the scope of the prohibition.  If the ServiceRequest.code and ServiceRequest.doNotPerform both contain negation, that will reinforce prohibition and should not have a double negative interpretation.
	DoNotPerform *bool `bson:"doNotPerform,omitempty" json:"doNotPerform,omitempty" xml:"doNotPerform,omitempty"`
	// A code or reference that identifies a particular service (i.e., procedure, diagnostic investigation, or panel of investigations) that have been requested.
	//
	// Comment: Many laboratory and radiology procedure codes embed the specimen/organ system in the test order name, for example,  serum or serum/plasma glucose, or a chest x-ray. The specimen might not be recorded separately from the test code.
	Code *CodeableReference `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// Additional details and instructions about the how the services are to be delivered.
	//
	//   For example, and order for a urinary catheter may have an order detail for an external or indwelling catheter, or an order for a bandage may require additional instructions specifying how the bandage should be applied.
	//
	// Comment: For information from the medical record intended to support the delivery of the requested services, use the `supportingInformation` element.
	OrderDetail []ServiceRequestOrderDetail `bson:"orderDetail,omitempty" json:"orderDetail,omitempty" xml:"orderDetail,omitempty"`
	// An amount of service being requested which can be a quantity ( for example $1,500 home modification), a ratio ( for example, 20 half day visits per month), or a range (2.0 to 1.8 Gy per fraction).
	QuantityQuantity *Quantity `bson:"quantityQuantity,omitempty" json:"quantityQuantity,omitempty" xml:"quantityQuantity,omitempty"`
	QuantityRatio    *Ratio    `bson:"quantityRatio,omitempty" json:"quantityRatio,omitempty" xml:"quantityRatio,omitempty"`
	QuantityRange    *Range    `bson:"quantityRange,omitempty" json:"quantityRange,omitempty" xml:"quantityRange,omitempty"`
	// On whom or what the service is to be performed.
	//
	// This is usually a human patient, but can also be requested on animals, groups of humans or animals, devices such as dialysis machines, or even locations (typically for environmental scans).
	Subject Reference `bson:"subject" json:"subject" xml:"subject"`
	// The actual focus of a service request when it is not the subject of record representing something or someone associated with the subject such as a spouse, parent, fetus, or donor.
	//
	// The focus of a service request could also be an existing condition,  an intervention, the subject's diet,  another service request on the subject,  or a body structure such as tumor or implanted device.
	Focus []Reference `bson:"focus,omitempty" json:"focus,omitempty" xml:"focus,omitempty"`
	// An encounter that provides additional information about the healthcare context in which this request is made.
	Encounter *Reference `bson:"encounter,omitempty" json:"encounter,omitempty" xml:"encounter,omitempty"`
	// The date/time at which the requested service should occur.
	OccurrenceDateTime *string `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty" xml:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *Period `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty" xml:"occurrencePeriod,omitempty"`
	OccurrenceTiming   *Timing `bson:"occurrenceTiming,omitempty" json:"occurrenceTiming,omitempty" xml:"occurrenceTiming,omitempty"`
	// If a CodeableConcept is present, it indicates the pre-condition for performing the service.
	//
	//  For example "pain", "on flare-up", etc.
	AsNeededBoolean         *bool            `bson:"asNeededBoolean,omitempty" json:"asNeededBoolean,omitempty" xml:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept *CodeableConcept `bson:"asNeededCodeableConcept,omitempty" json:"asNeededCodeableConcept,omitempty" xml:"asNeededCodeableConcept,omitempty"`
	// When the request transitioned to being actionable.
	AuthoredOn *string `bson:"authoredOn,omitempty" json:"authoredOn,omitempty" xml:"authoredOn,omitempty"`
	// The individual who initiated the request and has responsibility for its activation.
	//
	// Comment: This not the dispatcher, but rather who is the authorizer.  This element is not intended to handle delegation which would generally be managed through the Provenance resource.
	Requester *Reference `bson:"requester,omitempty" json:"requester,omitempty" xml:"requester,omitempty"`
	// Desired type of performer for doing the requested service.
	//
	// Comment: This is a  role, not  a participation type.  In other words, does not describe the task but describes the capacity.  For example, “compounding pharmacy”, “psychiatrist” or “internal referral”.
	PerformerType *CodeableConcept `bson:"performerType,omitempty" json:"performerType,omitempty" xml:"performerType,omitempty"`
	// The desired performer for doing the requested service.
	//
	//  For example, the surgeon, dermatopathologist, endoscopist, etc.
	//
	// Comment: If multiple performers are present, it is interpreted as a list of *alternative* performers without any preference regardless of order.  If order of preference is needed use the [http://hl7.org/fhir/StructureDefinition/request-performerOrder](http://hl7.org/fhir/extensions/StructureDefinition-request-performerOrder.html).  Use CareTeam to represent a group of performers (for example, Practitioner A *and* Practitioner B).
	Performer []Reference `bson:"performer,omitempty" json:"performer,omitempty" xml:"performer,omitempty"`
	// The preferred location(s) where the procedure should actually happen in coded or free text form.
	//
	// E.g. at home or nursing day care center.
	Location []CodeableReference `bson:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// An explanation or justification for why this service is being requested in coded or textual form.
	//
	//   This is often for billing purposes.  May relate to the resources referred to in `supportingInfo`.
	//
	// Comment: This element represents why the referral is being made and may be used to decide how the service will be performed, or even if it will be performed at all. To be as specific as possible,  a reference to  *Observation* or *Condition* should be used if available.  Otherwise, use `concept.text` element if the data is free (uncoded) text as shown in the [CT Scan example](servicerequest-example-di.html).
	Reason []CodeableReference `bson:"reason,omitempty" json:"reason,omitempty" xml:"reason,omitempty"`
	// Insurance plans, coverage extensions, pre-authorizations and/or pre-determinations that may be needed for delivering the requested service.
	Insurance []Reference `bson:"insurance,omitempty" json:"insurance,omitempty" xml:"insurance,omitempty"`
	// Additional clinical information about the patient or specimen that may influence the services or their interpretations.
	//
	//     This information includes diagnosis, clinical findings and other observations.  In laboratory ordering these are typically referred to as "ask at order entry questions (AOEs)".  This includes observations explicitly requested by the producer (filler) to provide context or supporting information needed to complete the order. For example,  reporting the amount of inspired oxygen for blood gas measurements.
	//
	// Comment: To represent information about how the services are to be delivered use the `instructions` element.
	SupportingInfo []CodeableReference `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty" xml:"supportingInfo,omitempty"`
	// One or more specimens that the laboratory procedure will use.
	//
	// Comment: The purpose of the ServiceRequest.specimen is to reflect the actual specimen that the requested test/procedure is asked to be performed on, whether the lab already has it or not.  References to specimens for purposes other than to perform a test/procedure on should be made using the ServiceRequest.supportingInfo or the Specimen.parent where the Specimen.parent would enable descendency and ServiceRequest.supportingInfo a general reference for context.
	Specimen []Reference `bson:"specimen,omitempty" json:"specimen,omitempty" xml:"specimen,omitempty"`
	// Anatomic location where the procedure should be performed. This is the target site.
	//
	// Comment: Only used if not implicit in the code found in ServiceRequest.code.  If the use case requires BodySite to be handled as a separate resource instead of an inline coded element (e.g. to identify and track separately)  then use the standard extension [http://hl7.org/fhir/StructureDefinition/procedure-targetBodyStructure](http://hl7.org/fhir/extensions/StructureDefinition-procedure-targetBodyStructure.html).
	BodySite []CodeableConcept `bson:"bodySite,omitempty" json:"bodySite,omitempty" xml:"bodySite,omitempty"`
	// Anatomic location where the procedure should be performed. This is the target site.
	BodyStructure *Reference `bson:"bodyStructure,omitempty" json:"bodyStructure,omitempty" xml:"bodyStructure,omitempty"`
	// Any other notes and comments made about the service request.
	//
	// For example, internal billing notes.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// Instructions in terms that are understood by the patient or consumer.
	PatientInstruction []ServiceRequestPatientInstruction `bson:"patientInstruction,omitempty" json:"patientInstruction,omitempty" xml:"patientInstruction,omitempty"`
	// Key events in the history of the request.
	//
	// Comment: This might not include provenances for all versions of the request – only those deemed “relevant” or important.This SHALL NOT include the Provenance associated with this current version of the resource.  (If that provenance is deemed to be a “relevant” change, it will need to be added as part of a later update.  Until then, it can be queried directly as the Provenance that points to this version using _revincludeAll Provenances should have some historical version of this Request as their subject.
	RelevantHistory []Reference `bson:"relevantHistory,omitempty" json:"relevantHistory,omitempty" xml:"relevantHistory,omitempty"`
}

type ServiceRequestOrderDetail struct {
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
	// Indicates the context of the order details by reference.
	ParameterFocus *CodeableReference `bson:"parameterFocus,omitempty" json:"parameterFocus,omitempty" xml:"parameterFocus,omitempty"`
	// The parameter details for the service being requested.
	Parameter []ServiceRequestOrderDetailParameter `bson:"parameter" json:"parameter" xml:"parameter"`
}

type ServiceRequestOrderDetailParameter struct {
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
	// A value representing the additional detail or instructions for the order (e.g., catheter insertion, body elevation, descriptive device configuration and/or setting instructions).
	Code CodeableConcept `bson:"code" json:"code" xml:"code"`
	// Indicates a value for the order detail.
	//
	// Comment: CodeableConcept values are indented to express concepts that would normally be coded - when a code is not available for a concept, CodeableConcept.text can be used. When the data is a text or not a single identifiable concept, string should be used.
	ValueQuantity        Quantity        `bson:"valueQuantity" json:"valueQuantity" xml:"valueQuantity"`
	ValueRatio           Ratio           `bson:"valueRatio" json:"valueRatio" xml:"valueRatio"`
	ValueRange           Range           `bson:"valueRange" json:"valueRange" xml:"valueRange"`
	ValueBoolean         bool            `bson:"valueBoolean" json:"valueBoolean" xml:"valueBoolean"`
	ValueCodeableConcept CodeableConcept `bson:"valueCodeableConcept" json:"valueCodeableConcept" xml:"valueCodeableConcept"`
	ValueString          string          `bson:"valueString" json:"valueString" xml:"valueString"`
	ValuePeriod          Period          `bson:"valuePeriod" json:"valuePeriod" xml:"valuePeriod"`
}

type ServiceRequestPatientInstruction struct {
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
	// Instructions in terms that are understood by the patient or consumer.
	InstructionMarkdown  *string    `bson:"instructionMarkdown,omitempty" json:"instructionMarkdown,omitempty" xml:"instructionMarkdown,omitempty"`
	InstructionReference *Reference `bson:"instructionReference,omitempty" json:"instructionReference,omitempty" xml:"instructionReference,omitempty"`
}

func (r ServiceRequest) ResourceType() string {
	return "ServiceRequest"
}

// MarshalJSON implements json.Marshaler
func (r ServiceRequest) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		ServiceRequest
		ResourceType string `json:"resourceType"`
	}{
		ResourceType:   r.ResourceType(),
		ServiceRequest: r,
	})
}
