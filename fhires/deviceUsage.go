package fhires

import "encoding/json/v2"

// A record of a device being used by a patient where the record is the result of a report from the patient or a clinician.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/DeviceUsage
type DeviceUsage struct {
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
	// An external identifier for this statement such as an IRI.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// A plan, proposal or order that is fulfilled in whole or in part by this DeviceUsage.
	BasedOn []Reference `bson:"basedOn,omitempty" json:"basedOn,omitempty" xml:"basedOn,omitempty"`
	// A code representing the patient or other source's judgment about the state of the device used that this statement is about.
	//
	//  Generally this will be active or completed.
	//
	// Comment: DeviceUseStatment is a statement at a point in time.  The status is only representative at the point when it was asserted.  The value set for contains codes that assert the status of the use  by the patient (for example, stopped or on hold) as well as codes that assert the status of the resource itself (for example, entered in error).This element is labeled as a modifier because the status contains the codes that mark the statement as not currently valid.
	Status DeviceUsageStatus `bson:"status" json:"status" xml:"status"`
	// This attribute indicates a category for the statement - The device statement may be made in an inpatient or outpatient settting (inpatient | outpatient | community | patientspecified).
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
	// The patient who used the device.
	Patient Reference `bson:"patient" json:"patient" xml:"patient"`
	// Allows linking the DeviceUsage to the underlying Request, or to other information that supports or is used to derive the DeviceUsage.
	//
	// Comment: The most common use cases for deriving a DeviceUsage comes from creating it from a request or from an observation or a claim. it should be noted that the amount of information that is available varies from the type resource that you derive the DeviceUsage from.
	DerivedFrom []Reference `bson:"derivedFrom,omitempty" json:"derivedFrom,omitempty" xml:"derivedFrom,omitempty"`
	// The encounter or episode of care that establishes the context for this device use statement.
	Context *Reference `bson:"context,omitempty" json:"context,omitempty" xml:"context,omitempty"`
	// How often the device was used.
	TimingTiming   *Timing `bson:"timingTiming,omitempty" json:"timingTiming,omitempty" xml:"timingTiming,omitempty"`
	TimingPeriod   *Period `bson:"timingPeriod,omitempty" json:"timingPeriod,omitempty" xml:"timingPeriod,omitempty"`
	TimingDateTime *string `bson:"timingDateTime,omitempty" json:"timingDateTime,omitempty" xml:"timingDateTime,omitempty"`
	// The time at which the statement was recorded by informationSource.
	DateAsserted *string `bson:"dateAsserted,omitempty" json:"dateAsserted,omitempty" xml:"dateAsserted,omitempty"`
	// The status of the device usage, for example always, sometimes, never.
	//
	// This is not the same as the status of the statement.
	UsageStatus *CodeableConcept `bson:"usageStatus,omitempty" json:"usageStatus,omitempty" xml:"usageStatus,omitempty"`
	// The reason for asserting the usage status - for example forgot, lost, stolen, broken.
	UsageReason []CodeableConcept `bson:"usageReason,omitempty" json:"usageReason,omitempty" xml:"usageReason,omitempty"`
	// This indicates how or if the device is being used.
	Adherence *DeviceUsageAdherence `bson:"adherence,omitempty" json:"adherence,omitempty" xml:"adherence,omitempty"`
	// Who reported the device was being used by the patient.
	InformationSource *Reference `bson:"informationSource,omitempty" json:"informationSource,omitempty" xml:"informationSource,omitempty"`
	// Code or Reference to device used.
	Device CodeableReference `bson:"device" json:"device" xml:"device"`
	// Reason or justification for the use of the device.
	//
	// A coded concept, or another resource whose existence justifies this DeviceUsage.
	//
	// Comment: When the status is not done, the reason code indicates why it was not done.
	Reason []CodeableReference `bson:"reason,omitempty" json:"reason,omitempty" xml:"reason,omitempty"`
	// Indicates the anotomic location on the subject's body where the device was used ( i.e.
	//
	// the target).
	BodySite *CodeableReference `bson:"bodySite,omitempty" json:"bodySite,omitempty" xml:"bodySite,omitempty"`
	// Details about the device statement that were not represented at all or sufficiently in one of the attributes provided in a class.
	//
	// These may include for example a comment, an instruction, or a note associated with the statement.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
}

type DeviceUsageAdherence struct {
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
	// Type of adherence.
	Code CodeableConcept `bson:"code" json:"code" xml:"code"`
	// Reason for adherence type.
	Reason []CodeableConcept `bson:"reason" json:"reason" xml:"reason"`
}

func (r DeviceUsage) ResourceType() string {
	return "DeviceUsage"
}

// MarshalJSON implements json.Marshaler
func (r DeviceUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		DeviceUsage
		ResourceType string `json:"resourceType"`
	}{
		DeviceUsage:  r,
		ResourceType: r.ResourceType(),
	})
}
