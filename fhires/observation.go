package fhires

import v2 "encoding/json/v2"

// Measurements and simple assertions made about a patient, device or other subject.
//
// Purpose: Observations are a key aspect of healthcare.  This resource is used to capture those that do not require more sophisticated mechanisms.
//
// Status: active
//
// http://hl7.org/fhir/StructureDefinition/Observation
type Observation struct {
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
	// A unique identifier assigned to this observation.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The reference to a FHIR ObservationDefinition resource that provides the definition that is adhered to in whole or in part by this Observation instance.
	//
	// Comment: ObservationDefinition can be referenced by its canonical url using instantiatesCanonical, or by a name or an identifier using the appropriate sub-elements of instantiatesReference.
	InstantiatesCanonical *string    `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty" xml:"instantiatesCanonical,omitempty"`
	InstantiatesReference *Reference `bson:"instantiatesReference,omitempty" json:"instantiatesReference,omitempty" xml:"instantiatesReference,omitempty"`
	// A plan, proposal or order that is fulfilled in whole or in part by this event.
	//
	//  For example, a MedicationRequest may require a patient to have laboratory test performed before  it is dispensed.
	BasedOn []Reference `bson:"basedOn,omitempty" json:"basedOn,omitempty" xml:"basedOn,omitempty"`
	// Identifies the observation(s) that triggered the performance of this observation.
	TriggeredBy []ObservationTriggeredBy `bson:"triggeredBy,omitempty" json:"triggeredBy,omitempty" xml:"triggeredBy,omitempty"`
	// A larger event of which this particular Observation is a component or step.
	//
	//  For example,  an observation as part of a procedure.
	//
	// Comment: To link an Observation to an Encounter use `encounter`.  See the  [Notes](observation.html#obsgrouping) below for guidance on referencing another Observation.
	PartOf []Reference `bson:"partOf,omitempty" json:"partOf,omitempty" xml:"partOf,omitempty"`
	// The status of the result value.
	//
	// Comment: This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid.
	Status ObservationStatus `bson:"status" json:"status" xml:"status"`
	// A code that classifies the general type of observation being made.
	//
	// Comment: In addition to the required category valueset, this element allows various categorization schemes based on the owner’s definition of the category and effectively multiple categories can be used at once.  The level of granularity is defined by the category concepts in the value set.
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
	// Describes what was observed. Sometimes this is called the observation "name".
	//
	// Comment: *All* code-value and, if present, component.code-component.value pairs need to be taken into account to correctly understand the meaning of the observation.
	Code CodeableConcept `bson:"code" json:"code" xml:"code"`
	// The patient, or group of patients, location, device, organization, procedure or practitioner this observation is about and into whose or what record the observation is placed.
	//
	// If the actual focus of the observation is different from the subject (or a sample of, part, or region of the subject), the `focus` element or the `code` itself specifies the actual focus of the observation.
	//
	// Comment: One would expect this element to be a cardinality of 1..1. The only circumstance in which the subject can be missing is when the observation is made by a device that does not know the patient. In this case, the observation SHALL be matched to a patient through some context/channel matching technique, and at this point, the observation should be updated. The subject of an Observation may in some cases be a procedure.  This supports the regulatory inspection use case where observations are captured during inspections of a procedure that is being performed (independent of any particular patient or whether patient related at all).
	Subject *Reference `bson:"subject,omitempty" json:"subject,omitempty" xml:"subject,omitempty"`
	// The actual focus of an observation when it is not the patient of record representing something or someone associated with the patient such as a spouse, parent, fetus, or donor.
	//
	// For example, fetus observations in a mother's record.  The focus of an observation could also be an existing condition,  an intervention, the subject's diet,  another observation of the subject,  or a body structure such as tumor or implanted device.   An example use case would be using the Observation resource to capture whether the mother is trained to change her child's tracheostomy tube. In this example, the child is the patient of record and the mother is the focus.
	//
	// Comment: Typically, an observation is made about the subject - a patient, or group of patients, location, or device - and the distinction between the subject and what is directly measured for an observation is specified in the observation code itself ( e.g., "Blood Glucose") and does not need to be represented separately using this element.  Use `specimen` if a reference to a specimen is required.  If a code is required instead of a resource use either  `bodysite` for bodysites or the standard extension [http://hl7.org/fhir/StructureDefinition/observation-focusCode](http://hl7.org/fhir/extensions/StructureDefinition-observation-focusCode.html).
	Focus []Reference `bson:"focus,omitempty" json:"focus,omitempty" xml:"focus,omitempty"`
	// The healthcare event  (e.g.
	//
	// a patient and healthcare provider interaction) during which this observation is made.
	//
	// Comment: This will typically be the encounter the event occurred within, but some events may be initiated prior to or after the official completion of an encounter but still be tied to the context of the encounter (e.g. pre-admission laboratory tests).
	Encounter *Reference `bson:"encounter,omitempty" json:"encounter,omitempty" xml:"encounter,omitempty"`
	// The time or time-period the observed value is asserted as being true.
	//
	// For biological subjects - e.g. human patients - this is usually called the "physiologically relevant time". This is usually either the time of the procedure or of specimen collection, but very often the source of the date/time is not known, only the date/time itself.
	//
	// Comment: At least a date should be present unless this observation is a historical report.  For recording imprecise or "fuzzy" times (For example, a blood glucose measurement taken "after breakfast") use the [Timing](datatypes.html#timing) datatype which allow the measurement to be tied to regular life events.
	EffectiveDateTime *string `bson:"effectiveDateTime,omitempty" json:"effectiveDateTime,omitempty" xml:"effectiveDateTime,omitempty"`
	EffectivePeriod   *Period `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty" xml:"effectivePeriod,omitempty"`
	EffectiveTiming   *Timing `bson:"effectiveTiming,omitempty" json:"effectiveTiming,omitempty" xml:"effectiveTiming,omitempty"`
	EffectiveInstant  *string `bson:"effectiveInstant,omitempty" json:"effectiveInstant,omitempty" xml:"effectiveInstant,omitempty"`
	// The date and time this version of the observation was made available to providers, typically after the results have been reviewed and verified.
	//
	// Comment: For Observations that don't require review and verification, it may be the same as the [`lastUpdated` ](resource-definitions.html#Meta.lastUpdated) time of the resource itself.  For Observations that do require review and verification for certain updates, it might not be the same as the `lastUpdated` time of the resource itself due to a non-clinically significant update that doesn't require the new version to be reviewed and verified again.
	Issued *string `bson:"issued,omitempty" json:"issued,omitempty" xml:"issued,omitempty"`
	// Who was responsible for asserting the observed value as "true".
	Performer []Reference `bson:"performer,omitempty" json:"performer,omitempty" xml:"performer,omitempty"`
	// The information determined as a result of making the observation, if the information has a simple value.
	//
	// Comment: - An observation may have:
	//
	//   1.  a single value here
	//
	//   1.  both a value and a set of related or component values
	//
	//   1.  only a set of related or component values.
	//
	// -  If a value is present, the datatype for this element should be determined by the `code`.
	//
	// -  *CodeableConcept* with just a text would be used instead of a string if the field was usually coded, or if the type associated with the `code` defines a coded value.
	//
	// -  *Attachment* is used if the observation result value is a binary file such as an image.  If the observation result value is derived from the binary file (for example 'X' detected and here is the the proof in this image), the binary file may be directly represented using *DocumentReference* and referenced by `derivedFrom`.
	//
	// - The usage of valueReference is restricted to the MolecularSequence resource when used as a definitional resource, not as a patient-specific finding. .
	//
	// - For additional guidance, see the [Notes section](observation.html#notes) below.
	ValueQuantity        *Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty" xml:"valueQuantity,omitempty"`
	ValueCodeableConcept *CodeableConcept `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty" xml:"valueCodeableConcept,omitempty"`
	ValueString          *string          `bson:"valueString,omitempty" json:"valueString,omitempty" xml:"valueString,omitempty"`
	ValueBoolean         *bool            `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty" xml:"valueBoolean,omitempty"`
	ValueInteger         *int             `bson:"valueInteger,omitempty" json:"valueInteger,omitempty" xml:"valueInteger,omitempty"`
	ValueRange           *Range           `bson:"valueRange,omitempty" json:"valueRange,omitempty" xml:"valueRange,omitempty"`
	ValueRatio           *Ratio           `bson:"valueRatio,omitempty" json:"valueRatio,omitempty" xml:"valueRatio,omitempty"`
	ValueSampledData     *SampledData     `bson:"valueSampledData,omitempty" json:"valueSampledData,omitempty" xml:"valueSampledData,omitempty"`
	ValueTime            *string          `bson:"valueTime,omitempty" json:"valueTime,omitempty" xml:"valueTime,omitempty"`
	ValueDateTime        *string          `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty" xml:"valueDateTime,omitempty"`
	ValuePeriod          *Period          `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty" xml:"valuePeriod,omitempty"`
	ValueAttachment      *Attachment      `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty" xml:"valueAttachment,omitempty"`
	ValueReference       *Reference       `bson:"valueReference,omitempty" json:"valueReference,omitempty" xml:"valueReference,omitempty"`
	// Provides a reason why the expected value in the element Observation.value[x] is missing.
	//
	// Comment: Null or exceptional values can be represented two ways in FHIR Observations.  One way is to simply include them in the value set and represent the exceptions in the value.  For example, measurement values for a serology test could be  "detected", "not detected", "inconclusive", or  "specimen unsatisfactory".
	//
	// The alternate way is to use the value element for actual observations and use the explicit dataAbsentReason element to record exceptional values.  For example, the dataAbsentReason code "error" could be used when the measurement was not completed. Note that an observation may only be reported if there are values to report. For example differential cell counts values may be reported only when > 0.  Because of these options, use-case agreements are required to interpret general observations for null or exceptional values.
	DataAbsentReason *CodeableConcept `bson:"dataAbsentReason,omitempty" json:"dataAbsentReason,omitempty" xml:"dataAbsentReason,omitempty"`
	// A categorical assessment of an observation value.  For example, high, low, normal.
	//
	// Comment: Historically used for laboratory results (known as 'abnormal flag' ),  its use extends to other use cases where coded interpretations  are relevant.  Often reported as one or more simple compact codes this element is often placed adjacent to the result value in reports and flow sheets to signal the meaning/normalcy status of the result.
	Interpretation []CodeableConcept `bson:"interpretation,omitempty" json:"interpretation,omitempty" xml:"interpretation,omitempty"`
	// Comments about the observation or the results.
	//
	// Comment: May include general statements about the observation, or statements about significant, unexpected or unreliable results values, or information about its source when relevant to its interpretation.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// Indicates the site on the subject's body where the observation was made (i.e.
	//
	// the target site).
	//
	// Comment: Only used if not implicit in code found in Observation.code.  In many systems, this may be represented as a related observation instead of an inline component.
	BodySite *CodeableConcept `bson:"bodySite,omitempty" json:"bodySite,omitempty" xml:"bodySite,omitempty"`
	// Indicates the body structure on the subject's body where the observation was made (i.e.
	//
	// the target site).
	//
	// Comment: Only used if not implicit in code found in Observation.code or bodySite is used. In many systems, this may be represented as a related observation instead of an inline component.
	BodyStructure *Reference `bson:"bodyStructure,omitempty" json:"bodyStructure,omitempty" xml:"bodyStructure,omitempty"`
	// Indicates the mechanism used to perform the observation.
	//
	// Comment: Only used if not implicit in code for Observation.code.
	Method *CodeableConcept `bson:"method,omitempty" json:"method,omitempty" xml:"method,omitempty"`
	// The specimen that was used when this observation was made.
	//
	// Comment: Should only be used if not implicit in code found in `Observation.code`.  Observations are not made on specimens themselves; they are made on a subject, but in many cases by the means of a specimen. Note that although specimens are often involved, they are not always tracked and reported explicitly. Also note that observation resources may be used in contexts that track the specimen explicitly (e.g. Diagnostic Report).
	Specimen *Reference `bson:"specimen,omitempty" json:"specimen,omitempty" xml:"specimen,omitempty"`
	// A reference to the device that generates the measurements or the device settings for the device.
	//
	// Comment: Note that this is not meant to represent a device involved in the transmission of the result, e.g., a gateway.  Such devices may be documented using the Provenance resource where relevant.
	Device *Reference `bson:"device,omitempty" json:"device,omitempty" xml:"device,omitempty"`
	// Guidance on how to interpret the value by comparison to a normal or recommended range.
	//
	//  Multiple reference ranges are interpreted as an "OR".   In other words, to represent two distinct target populations, two `referenceRange` elements would be used.
	//
	// Comment: Most observations only have one generic reference range. Systems MAY choose to restrict to only supplying the relevant reference range based on knowledge about the patient (e.g., specific to the patient's age, gender, weight and other factors), but this might not be possible or appropriate. Whenever more than one reference range is supplied, the differences between them SHOULD be provided in the reference range and/or age properties.
	ReferenceRange []ObservationReferenceRange `bson:"referenceRange,omitempty" json:"referenceRange,omitempty" xml:"referenceRange,omitempty"`
	// This observation is a group observation (e.g.
	//
	// a battery, a panel of tests, a set of vital sign measurements) that includes the target as a member of the group.
	//
	// Comment: When using this element, an observation will typically have either a value or a set of related resources, although both may be present in some cases.  For a discussion on the ways Observations can assembled in groups together, see [Notes](observation.html#obsgrouping) below.  Note that a system may calculate results from [QuestionnaireResponse](questionnaireresponse.html)  into a final score and represent the score as an Observation.
	HasMember []Reference `bson:"hasMember,omitempty" json:"hasMember,omitempty" xml:"hasMember,omitempty"`
	// The target resource that represents a measurement from which this observation value is derived.
	//
	// For example, a calculated anion gap or a fetal measurement based on an ultrasound image.
	//
	// Comment: All the reference choices that are listed in this element can represent clinical observations and other measurements that may be the source for a derived value.  The most common reference will be another Observation.  For a discussion on the ways Observations can assembled in groups together, see [Notes](observation.html#obsgrouping) below.
	DerivedFrom []Reference `bson:"derivedFrom,omitempty" json:"derivedFrom,omitempty" xml:"derivedFrom,omitempty"`
	// Some observations have multiple component observations.
	//
	//  These component observations are expressed as separate code value pairs that share the same attributes.  Examples include systolic and diastolic component observations for blood pressure measurement and multiple component observations for genetics observations.
	//
	// Comment: For a discussion on the ways Observations can be assembled in groups together see [Notes](observation.html#notes) below.
	Component []ObservationComponent `bson:"component,omitempty" json:"component,omitempty" xml:"component,omitempty"`
}

type ObservationTriggeredBy struct {
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
	// Reference to the triggering observation.
	Observation Reference `bson:"observation" json:"observation" xml:"observation"`
	// The type of trigger.
	//
	// Reflex | Repeat | Re-run.
	Type TriggeredBytype `bson:"type" json:"type" xml:"type"`
	// Provides the reason why this observation was performed as a result of the observation(s) referenced.
	Reason *string `bson:"reason,omitempty" json:"reason,omitempty" xml:"reason,omitempty"`
}

type ObservationReferenceRange struct {
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
	// The value of the low bound of the reference range.
	//
	//  The low bound of the reference range endpoint is inclusive of the value (e.g.  reference range is >=5 - <=9). If the low bound is omitted,  it is assumed to be meaningless (e.g. reference range is <=2.3).
	Low *Quantity `bson:"low,omitempty" json:"low,omitempty" xml:"low,omitempty"`
	// The value of the high bound of the reference range.
	//
	//  The high bound of the reference range endpoint is inclusive of the value (e.g.  reference range is >=5 - <=9). If the high bound is omitted,  it is assumed to be meaningless (e.g. reference range is >= 2.3).
	High *Quantity `bson:"high,omitempty" json:"high,omitempty" xml:"high,omitempty"`
	// The value of the normal value of the reference range.
	NormalValue *CodeableConcept `bson:"normalValue,omitempty" json:"normalValue,omitempty" xml:"normalValue,omitempty"`
	// Codes to indicate the what part of the targeted reference population it applies to.
	//
	// For example, the normal or therapeutic range.
	//
	// Comment: This SHOULD be populated if there is more than one range.  If this element is not present then the normal range is assumed.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Codes to indicate the target population this reference range applies to.
	//
	//  For example, a reference range may be based on the normal population or a particular sex or race.  Multiple `appliesTo`  are interpreted as an "AND" of the target populations.  For example, to represent a target population of African American females, both a code of female and a code for African American would be used.
	//
	// Comment: This SHOULD be populated if there is more than one range.  If this element is not present then the normal population is assumed.
	AppliesTo []CodeableConcept `bson:"appliesTo,omitempty" json:"appliesTo,omitempty" xml:"appliesTo,omitempty"`
	// The age at which this reference range is applicable.
	//
	// This is a neonatal age (e.g. number of weeks at term) if the meaning says so.
	Age *Range `bson:"age,omitempty" json:"age,omitempty" xml:"age,omitempty"`
	// Text based reference range in an observation which may be used when a quantitative range is not appropriate for an observation.
	//
	//  An example would be a reference value of "Negative" or a list or table of "normals".
	Text *string `bson:"text,omitempty" json:"text,omitempty" xml:"text,omitempty"`
}

type ObservationComponent struct {
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
	// Describes what was observed. Sometimes this is called the observation "code".
	//
	// Comment: *All* code-value and  component.code-component.value pairs need to be taken into account to correctly understand the meaning of the observation.
	Code CodeableConcept `bson:"code" json:"code" xml:"code"`
	// The information determined as a result of making the observation, if the information has a simple value.
	//
	// Comment: Used when observation has a set of component observations:
	//
	// -  An observation may have both a value (e.g. an  Apgar score) and component observations (the observations from which the Apgar score was derived).
	//
	// -  If a value is present, the datatype for this element should be determined by the `code`.
	//
	// -  *CodeableConcept* with just a text would be used instead of a string if the field was usually coded, or if the type associated with the `code` defines a coded value.
	//
	// -  *Attachment* is used if the observation result value is a binary file such as an image.  If the observation result value is derived from the binary file (for example 'X' detected and here is the the proof in this image), the binary file may be directly represented using *DocumentReference* and referenced by `derivedFrom`.
	//
	// - If a value is present, the datatype for this element should be determined by the `code`.
	//
	// -  *CodeableConcept* with just a text would be used instead of a string if the field was usually coded, or if the type associated with the `code` defines a coded value.
	//
	// -  *Attachment* is used if the observation result value is a binary file such as an image.  If the observation result value is derived from the binary file (for example 'X' detected and here is the the proof in this image), the binary file may be directly represented using *DocumentReference* and referenced by `derivedFrom`.
	//
	// - The usage of valueReference is restricted to the MolecularSequence resource when used as a definitional resource, not as a patient-specific finding. .
	//
	// - For additional guidance, see the [Notes section](observation.html#notes) below.
	ValueQuantity        *Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty" xml:"valueQuantity,omitempty"`
	ValueCodeableConcept *CodeableConcept `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty" xml:"valueCodeableConcept,omitempty"`
	ValueString          *string          `bson:"valueString,omitempty" json:"valueString,omitempty" xml:"valueString,omitempty"`
	ValueBoolean         *bool            `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty" xml:"valueBoolean,omitempty"`
	ValueInteger         *int             `bson:"valueInteger,omitempty" json:"valueInteger,omitempty" xml:"valueInteger,omitempty"`
	ValueRange           *Range           `bson:"valueRange,omitempty" json:"valueRange,omitempty" xml:"valueRange,omitempty"`
	ValueRatio           *Ratio           `bson:"valueRatio,omitempty" json:"valueRatio,omitempty" xml:"valueRatio,omitempty"`
	ValueSampledData     *SampledData     `bson:"valueSampledData,omitempty" json:"valueSampledData,omitempty" xml:"valueSampledData,omitempty"`
	ValueTime            *string          `bson:"valueTime,omitempty" json:"valueTime,omitempty" xml:"valueTime,omitempty"`
	ValueDateTime        *string          `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty" xml:"valueDateTime,omitempty"`
	ValuePeriod          *Period          `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty" xml:"valuePeriod,omitempty"`
	ValueAttachment      *Attachment      `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty" xml:"valueAttachment,omitempty"`
	ValueReference       *Reference       `bson:"valueReference,omitempty" json:"valueReference,omitempty" xml:"valueReference,omitempty"`
	// Provides a reason why the expected value in the element Observation.component.value[x] is missing.
	//
	// Comment: "Null" or exceptional values can be represented two ways in FHIR Observations.  One way is to simply include them in the value set and represent the exceptions in the value.  For example, measurement values for a serology test could be  "detected", "not detected", "inconclusive", or  "test not done".
	//
	// The alternate way is to use the value element for actual observations and use the explicit dataAbsentReason element to record exceptional values.  For example, the dataAbsentReason code "error" could be used when the measurement was not completed.  Because of these options, use-case agreements are required to interpret general observations for exceptional values.
	DataAbsentReason *CodeableConcept `bson:"dataAbsentReason,omitempty" json:"dataAbsentReason,omitempty" xml:"dataAbsentReason,omitempty"`
	// A categorical assessment of an observation value.  For example, high, low, normal.
	//
	// Comment: Historically used for laboratory results (known as 'abnormal flag' ),  its use extends to other use cases where coded interpretations  are relevant.  Often reported as one or more simple compact codes this element is often placed adjacent to the result value in reports and flow sheets to signal the meaning/normalcy status of the result.
	Interpretation []CodeableConcept `bson:"interpretation,omitempty" json:"interpretation,omitempty" xml:"interpretation,omitempty"`
	// Guidance on how to interpret the value by comparison to a normal or recommended range.
	//
	// Comment: Most observations only have one generic reference range. Systems MAY choose to restrict to only supplying the relevant reference range based on knowledge about the patient (e.g., specific to the patient's age, gender, weight and other factors), but this might not be possible or appropriate. Whenever more than one reference range is supplied, the differences between them SHOULD be provided in the reference range and/or age properties.
	ReferenceRange []ObservationReferenceRange `bson:"referenceRange,omitempty" json:"referenceRange,omitempty"`
}

func (r Observation) ResourceType() string {
	return "Observation"
}

// MarshalJSON implements json.Marshaler
func (r Observation) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		Observation
		ResourceType string `json:"resourceType"`
	}{
		Observation:  r,
		ResourceType: r.ResourceType(),
	})
}
