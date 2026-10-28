package fhires

import "encoding/json/v2"

// Describes the event of a patient being administered a vaccine or a record of an immunization as reported by a patient, a clinician or another party.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Immunization
type Immunization struct {
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
	// A unique identifier assigned to this immunization record.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// A plan, order or recommendation fulfilled in whole or in part by this immunization.
	//
	// Comment: Allows tracing of an authorization for the Immunization.
	BasedOn []Reference `bson:"basedOn,omitempty" json:"basedOn,omitempty" xml:"basedOn,omitempty"`
	// Indicates the current status of the immunization event.
	//
	// Comment: Will generally be set to show that the immunization has been completed or not done.  This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid.
	Status ImmunizationStatusCodes `bson:"status" json:"status" xml:"status"`
	// Indicates the reason the immunization event was not performed.
	//
	// Comment: This is generally only used for the status of "not-done". The reason for performing the immunization event is captured in reasonCode, not here.
	StatusReason *CodeableConcept `bson:"statusReason,omitempty" json:"statusReason,omitempty" xml:"statusReason,omitempty"`
	// Vaccine that was administered or was to be administered.
	//
	// Comment: The code for the administered vaccine may be reported at various levels of granularity from a very generic code for a general type of vaccine (e.g. the SNOMED or CVX code for the flu vaccine) to a very discrete code for a specific product including the form and packaging (e.g. an NDC code). Because vaccineCode uses the CodeableConcept data type, multiple codes at different levels of granularity may be reported for a single immunization event provided that all codes are appropriate for the same vaccine. For example, both NDC and CVX codes may be provided for a given immunization event by repeating the coding element within vaccineCode.
	VaccineCode CodeableConcept `bson:"vaccineCode" json:"vaccineCode" xml:"vaccineCode"`
	// An indication of which product was administered to the patient.
	//
	// This is typically a more detailed representation of the concept conveyed by the vaccineCode data element. If a Medication resource is referenced, it may be to a stand-alone resource or a contained resource within the Immunization resource.
	AdministeredProduct *CodeableReference `bson:"administeredProduct,omitempty" json:"administeredProduct,omitempty" xml:"administeredProduct,omitempty"`
	// Name of vaccine manufacturer.
	Manufacturer *CodeableReference `bson:"manufacturer,omitempty" json:"manufacturer,omitempty" xml:"manufacturer,omitempty"`
	// Lot number of the  vaccine product.
	LotNumber *string `bson:"lotNumber,omitempty" json:"lotNumber,omitempty" xml:"lotNumber,omitempty"`
	// Date vaccine batch expires.
	ExpirationDate *string `bson:"expirationDate,omitempty" json:"expirationDate,omitempty" xml:"expirationDate,omitempty"`
	// The patient who either received or did not receive the immunization.
	Patient Reference `bson:"patient" json:"patient" xml:"patient"`
	// The visit or admission or other contact between patient and health care provider the immunization was performed as part of.
	Encounter *Reference `bson:"encounter,omitempty" json:"encounter,omitempty" xml:"encounter,omitempty"`
	// Additional information that is relevant to the immunization (e.g.
	//
	// for a vaccine recipient who is pregnant, the gestational age of the fetus). The reason why a vaccine was given (e.g. occupation, underlying medical condition) should be conveyed in Immunization.reason, not as supporting information. The reason why a vaccine was not given (e.g. contraindication) should be conveyed in Immunization.statusReason, not as supporting information.
	SupportingInformation []Reference `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty" xml:"supportingInformation,omitempty"`
	// Date vaccine administered or was to be administered.
	//
	// Comment: When immunizations are given a specific date and time should always be known.   When immunizations are patient reported, a specific date might not be known.  Although partial dates are allowed, an adult patient might not be able to recall the year a childhood immunization was given. An exact date is always preferable, but the use of the String data type is acceptable when an exact date is not known. A small number of vaccines (e.g. live oral typhoid vaccine) are given as a series of patient self-administered dose over a span of time. In cases like this, often, only the first dose (typically a provider supervised dose) is recorded with the occurrence indicating the date/time of the first dose.
	OccurrenceDateTime string `bson:"occurrenceDateTime" json:"occurrenceDateTime" xml:"occurrenceDateTime"`
	OccurrenceString   string `bson:"occurrenceString" json:"occurrenceString" xml:"occurrenceString"`
	// Indicates whether the data contained in the resource was captured by the individual/organization which was responsible for the administration of the vaccine rather than as 'secondary reported' data documented by a third party.
	//
	// A value of 'true' means this data originated with the individual/organization which was responsible for the administration of the vaccine.
	//
	// Comment: Reflects the “reliability” of the content.
	PrimarySource *bool `bson:"primarySource,omitempty" json:"primarySource,omitempty" xml:"primarySource,omitempty"`
	// Typically the source of the data when the report of the immunization event is not based on information from the person who administered the vaccine.
	//
	// Comment: Typically will not be populated if primarySource = True, not required even if primarySource = False.
	InformationSource *CodeableReference `bson:"informationSource,omitempty" json:"informationSource,omitempty" xml:"informationSource,omitempty"`
	// The service delivery location where the vaccine administration occurred.
	Location *Reference `bson:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// Body site where vaccine was administered.
	Site *CodeableConcept `bson:"site,omitempty" json:"site,omitempty" xml:"site,omitempty"`
	// The path by which the vaccine product is taken into the body.
	Route *CodeableConcept `bson:"route,omitempty" json:"route,omitempty" xml:"route,omitempty"`
	// The quantity of vaccine product that was administered.
	DoseQuantity *Quantity `bson:"doseQuantity,omitempty" json:"doseQuantity,omitempty" xml:"doseQuantity,omitempty"`
	// Indicates who performed the immunization event.
	Performer []ImmunizationPerformer `bson:"performer,omitempty" json:"performer,omitempty" xml:"performer,omitempty"`
	// Extra information about the immunization that is not conveyed by the other attributes.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// Describes why the immunization occurred in coded or textual form, or Indicates another resource (Condition, Observation or DiagnosticReport) whose existence justifies this immunization.
	Reason []CodeableReference `bson:"reason,omitempty" json:"reason,omitempty" xml:"reason,omitempty"`
	// Indication if a dose is considered to be subpotent.
	//
	// By default, a dose should be considered to be potent.
	//
	// Comment: Typically, the recognition of the dose being sub-potent is retrospective, after the administration (ex. notification of a manufacturer recall after administration). However, in the case of a partial administration (the patient moves unexpectedly and only some of the dose is actually administered), subpotency may be recognized immediately, but it is still important to record the event.
	IsSubpotent *bool `bson:"isSubpotent,omitempty" json:"isSubpotent,omitempty" xml:"isSubpotent,omitempty"`
	// Reason why a dose is considered to be subpotent.
	SubpotentReason []CodeableConcept `bson:"subpotentReason,omitempty" json:"subpotentReason,omitempty" xml:"subpotentReason,omitempty"`
	// Indicates a patient's eligibility for a funding program.
	ProgramEligibility []ImmunizationProgramEligibility `bson:"programEligibility,omitempty" json:"programEligibility,omitempty" xml:"programEligibility,omitempty"`
	// Indicates the source of the vaccine actually administered.
	//
	// This may be different than the patient eligibility (e.g. the patient may be eligible for a publically purchased vaccine but due to inventory issues, vaccine purchased with private funds was actually administered).
	FundingSource *CodeableConcept `bson:"fundingSource,omitempty" json:"fundingSource,omitempty" xml:"fundingSource,omitempty"`
	// Categorical data indicating that an adverse event is associated in time to an immunization.
	//
	// Comment: A reaction may be an indication of an allergy or intolerance and, if this is determined to be the case, it should be recorded as a new AllergyIntolerance resource instance as most systems will not query against past Immunization.reaction elements.
	Reaction []ImmunizationReaction `bson:"reaction,omitempty" json:"reaction,omitempty" xml:"reaction,omitempty"`
	// The protocol (set of recommendations) being followed by the provider who administered the dose.
	ProtocolApplied []ImmunizationProtocolApplied `bson:"protocolApplied,omitempty" json:"protocolApplied,omitempty" xml:"protocolApplied,omitempty"`
}

type ImmunizationPerformer struct {
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
	// Describes the type of performance (e.g. ordering provider, administering provider, etc.).
	Function *CodeableConcept `bson:"function,omitempty" json:"function,omitempty" xml:"function,omitempty"`
	// The practitioner or organization who performed the action.
	//
	// Comment: When the individual practitioner who performed the action is known, it is best to send.
	Actor Reference `bson:"actor" json:"actor" xml:"actor"`
}

type ImmunizationProgramEligibility struct {
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
	// Indicates which program the patient had their eligility evaluated for.
	Program CodeableConcept `bson:"program" json:"program" xml:"program"`
	// Indicates the patient's eligility status for for a specific payment program.
	ProgramStatus CodeableConcept `bson:"programStatus" json:"programStatus" xml:"programStatus"`
}

type ImmunizationReaction struct {
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
	// Date of reaction to the immunization.
	Date *string `bson:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// Details of the reaction.
	Manifestation *CodeableReference `bson:"manifestation,omitempty" json:"manifestation,omitempty" xml:"manifestation,omitempty"`
	// Self-reported indicator.
	Reported *bool `bson:"reported,omitempty" json:"reported,omitempty" xml:"reported,omitempty"`
}

type ImmunizationProtocolApplied struct {
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
	// One possible path to achieve presumed immunity against a disease - within the context of an authority.
	Series *string `bson:"series,omitempty" json:"series,omitempty" xml:"series,omitempty"`
	// Indicates the authority who published the protocol (e.g. ACIP) that is being followed.
	Authority *Reference `bson:"authority,omitempty" json:"authority,omitempty" xml:"authority,omitempty"`
	// The vaccine preventable disease the dose is being administered against.
	TargetDisease []CodeableConcept `bson:"targetDisease,omitempty" json:"targetDisease,omitempty" xml:"targetDisease,omitempty"`
	// Nominal position in a series as intended by the practitioner administering the dose.
	//
	// Comment: The use of an integer is preferred if known. Text should only be used in cases where an integer is not available (such as when documenting a recurring booster dose).
	DoseNumber string `bson:"doseNumber" json:"doseNumber" xml:"doseNumber"`
	// The recommended number of doses to achieve immunity as intended by the practitioner administering the dose.
	//
	// Comment: The use of an integer is preferred if known.
	SeriesDoses *string `bson:"seriesDoses,omitempty" json:"seriesDoses,omitempty" xml:"seriesDoses,omitempty"`
}

func (r Immunization) ResourceType() string {
	return "Immunization"
}

// MarshalJSON implements json.Marshaler
func (r Immunization) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Immunization
		ResourceType string `json:"resourceType"`
	}{
		Immunization: r,
		ResourceType: r.ResourceType(),
	})
}
