package fhires

import v2 "encoding/json/v2"

// Represents a request a device to be provided to a specific patient.
//
// The device may be an implantable device to be subsequently implanted, or an external assistive device, such as a walker, to be delivered and subsequently be used.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/DeviceRequest
type DeviceRequest struct {
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
	// Identifiers assigned to this order by the orderer or by the receiver.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The URL pointing to a FHIR-defined protocol, guideline, orderset or other definition that is adhered to in whole or in part by this DeviceRequest.
	//
	// Comment: Note: This is a business identifier, not a resource identifier (see [discussion](resource.html#identifiers)).  It is best practice for the identifier to only appear on a single resource instance, however business practices may occasionally dictate that multiple resource instances with the same identifier can exist - possibly even with different resource types.  For example, multiple Patient and a Person resource instance might share the same social insurance number.
	InstantiatesCanonical []string `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty" xml:"instantiatesCanonical,omitempty"`
	// The URL pointing to an externally maintained protocol, guideline, orderset or other definition that is adhered to in whole or in part by this DeviceRequest.
	//
	// Comment: This might be an HTML page, PDF, etc. or could just be a non-resolvable URI identifier.
	InstantiatesUri []string `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty" xml:"instantiatesUri,omitempty"`
	// Plan/proposal/order fulfilled by this request.
	BasedOn []Reference `bson:"basedOn,omitempty" json:"basedOn,omitempty" xml:"basedOn,omitempty"`
	// The request takes the place of the referenced completed or terminated request(s).
	Replaces []Reference `bson:"replaces,omitempty" json:"replaces,omitempty" xml:"replaces,omitempty"`
	// A shared identifier common to multiple independent Request instances that were activated/authorized more or less simultaneously by a single author.
	//
	//  The presence of the same identifier on each request ties those requests together and may have business ramifications in terms of reporting of results, billing, etc.  E.g. a requisition number shared by a set of lab tests ordered together, or a prescription number shared by all meds ordered at one time.
	GroupIdentifier *Identifier `bson:"groupIdentifier,omitempty" json:"groupIdentifier,omitempty" xml:"groupIdentifier,omitempty"`
	// The status of the request.
	//
	// Comment: This element is labeled as a modifier because the status contains the codes revoked and entered-in-error that mark the request as not currently valid.
	Status *RequestStatus `bson:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Whether the request is a proposal, plan, an original order or a reflex order.
	Intent RequestIntent `bson:"intent" json:"intent" xml:"intent"`
	// Indicates how quickly the request should be addressed with respect to other requests.
	Priority *RequestPriority `bson:"priority,omitempty" json:"priority,omitempty" xml:"priority,omitempty"`
	// If true, indicates that the provider is asking for the patient to either stop using or to not start using the specified device or category of devices.
	//
	// For example, the patient has undergone surgery and the provider is indicating that the patient should not wear contact lenses.
	//
	// Comment: If do not perform is not specified, the request is a positive request e.g. "do perform". DeviceRequest.reasonCode includes the reason for marking the DeviceRequest as 'do not perform'.
	DoNotPerform *bool `bson:"doNotPerform,omitempty" json:"doNotPerform,omitempty" xml:"doNotPerform,omitempty"`
	// The details of the device to be used.
	Code CodeableReference `bson:"code" json:"code" xml:"code"`
	// The number of devices to be provided.
	Quantity *int `bson:"quantity,omitempty" json:"quantity,omitempty" xml:"quantity,omitempty"`
	// Specific parameters for the ordered item.  For example, the prism value for lenses.
	Parameter []DeviceRequestParameter `bson:"parameter,omitempty" json:"parameter,omitempty" xml:"parameter,omitempty"`
	// The patient who will use the device.
	Subject Reference `bson:"subject" json:"subject" xml:"subject"`
	// An encounter that provides additional context in which this request is made.
	Encounter *Reference `bson:"encounter,omitempty" json:"encounter,omitempty" xml:"encounter,omitempty"`
	// The timing schedule for the use of the device.
	//
	// The Schedule data type allows many different expressions, for example. "Every 8 hours"; "Three times a day"; "1/2 an hour before breakfast for 10 days from 23-Dec 2011:"; "15 Oct 2013, 17 Oct 2013 and 1 Nov 2013".
	OccurrenceDateTime *string `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty" xml:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *Period `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty" xml:"occurrencePeriod,omitempty"`
	OccurrenceTiming   *Timing `bson:"occurrenceTiming,omitempty" json:"occurrenceTiming,omitempty" xml:"occurrenceTiming,omitempty"`
	// When the request transitioned to being actionable.
	AuthoredOn *string `bson:"authoredOn,omitempty" json:"authoredOn,omitempty" xml:"authoredOn,omitempty"`
	// The individual or entity who initiated the request and has responsibility for its activation.
	Requester *Reference `bson:"requester,omitempty" json:"requester,omitempty" xml:"requester,omitempty"`
	// The desired individual or entity to provide the device to the subject of the request (e.g., patient, location).
	Performer *CodeableReference `bson:"performer,omitempty" json:"performer,omitempty" xml:"performer,omitempty"`
	// Reason or justification for the use of this device.
	Reason []CodeableReference `bson:"reason,omitempty" json:"reason,omitempty" xml:"reason,omitempty"`
	// This status is to indicate whether the request is a PRN or may be given as needed.
	AsNeeded *bool `bson:"asNeeded,omitempty" json:"asNeeded,omitempty" xml:"asNeeded,omitempty"`
	// The reason for using the device.
	AsNeededFor *CodeableConcept `bson:"asNeededFor,omitempty" json:"asNeededFor,omitempty" xml:"asNeededFor,omitempty"`
	// Insurance plans, coverage extensions, pre-authorizations and/or pre-determinations that may be required for delivering the requested service.
	Insurance []Reference `bson:"insurance,omitempty" json:"insurance,omitempty" xml:"insurance,omitempty"`
	// Additional clinical information about the patient that may influence the request fulfilment.
	//
	//  For example, this may include where on the subject's body the device will be used (i.e. the target site).
	SupportingInfo []Reference `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty" xml:"supportingInfo,omitempty"`
	// Details about this request that were not represented at all or sufficiently in one of the attributes provided in a class.
	//
	// These may include for example a comment, an instruction, or a note associated with the statement.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// Key events in the history of the request.
	//
	// Comment: This might not include provenances for all versions of the request - only those deemed "relevant" or important.This SHALL NOT include the Provenance associated with this current version of the resource.  (If that provenance is deemed to be a "relevant" change, it will need to be added as part of a later update.  Until then, it can be queried directly as the Provenance that points to this version using _revincludeAll Provenances should have some historical version of this Request as their subject.
	RelevantHistory []Reference `bson:"relevantHistory,omitempty" json:"relevantHistory,omitempty" xml:"relevantHistory,omitempty"`
}

type DeviceRequestParameter struct {
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
	// A code or string that identifies the device detail being asserted.
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// The value of the device detail.
	//
	// Comment: Range means device should have a value that falls somewhere within the specified range.
	ValueCodeableConcept *CodeableConcept `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty" xml:"valueCodeableConcept,omitempty"`
	ValueQuantity        *Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty" xml:"valueQuantity,omitempty"`
	ValueRange           *Range           `bson:"valueRange,omitempty" json:"valueRange,omitempty" xml:"valueRange,omitempty"`
	ValueBoolean         *bool            `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty" xml:"valueBoolean,omitempty"`
}

func (r DeviceRequest) ResourceType() string {
	return "DeviceRequest"
}

// MarshalJSON implements json.Marshaler
func (r DeviceRequest) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		DeviceRequest
		ResourceType string `json:"resourceType"`
	}{
		DeviceRequest: r,
		ResourceType:  r.ResourceType(),
	})
}
