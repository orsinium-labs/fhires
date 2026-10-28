package fhires

import v2 "encoding/json/v2"

// Significant health conditions for a person related to the patient relevant in the context of care for the patient.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/FamilyMemberHistory
type FamilyMemberHistory struct {
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
	// Business identifiers assigned to this family member history by the performer or other systems which remain constant as the resource is updated and propagates from server to server.
	//
	// Comment: This is a business identifier, not a resource identifier (see [discussion](resource.html#identifiers)).  It is best practice for the identifier to only appear on a single resource instance, however business practices may occasionally dictate that multiple resource instances with the same identifier can exist - possibly even with different resource types.  For example, multiple Patient and a Person resource instance might share the same social insurance number.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The URL pointing to a FHIR-defined protocol, guideline, orderset or other definition that is adhered to in whole or in part by this FamilyMemberHistory.
	InstantiatesCanonical []string `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty" xml:"instantiatesCanonical,omitempty"`
	// The URL pointing to an externally maintained protocol, guideline, orderset or other definition that is adhered to in whole or in part by this FamilyMemberHistory.
	//
	// Comment: This might be an HTML page, PDF, etc. or could just be a non-resolvable URI identifier.
	InstantiatesUri []string `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty" xml:"instantiatesUri,omitempty"`
	// A code specifying the status of the record of the family history of a specific family member.
	//
	// Comment: This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid.
	Status FamilyHistoryStatus `bson:"status" json:"status" xml:"status"`
	// Describes why the family member's history is not available.
	DataAbsentReason *CodeableConcept `bson:"dataAbsentReason,omitempty" json:"dataAbsentReason,omitempty" xml:"dataAbsentReason,omitempty"`
	// The person who this history concerns.
	//
	// Comment: This is not the family member.
	Patient Reference `bson:"patient" json:"patient" xml:"patient"`
	// The date (and possibly time) when the family member history was recorded or last updated.
	//
	// Comment: This should be captured even if the same as the date on the List aggregating the full family history.
	Date *string `bson:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// Indicates who or what participated in the activities related to the family member history and how they were involved.
	Participant []FamilyMemberHistoryParticipant `bson:"participant,omitempty" json:"participant,omitempty" xml:"participant,omitempty"`
	// This will either be a name or a description; e.g.
	//
	// "Aunt Susan", "my cousin with the red hair".
	Name *string `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// The type of relationship this person has to the patient (father, mother, brother etc.).
	Relationship CodeableConcept `bson:"relationship" json:"relationship" xml:"relationship"`
	// The birth sex of the family member.
	//
	// Comment: This element should ideally reflect whether the individual is genetically male or female.  However, as reported information based on the knowledge of the patient or reporting friend/relative, there may be situations where the reported sex might not be totally accurate.  E.g. 'Aunt Sue' might be XY rather than XX.  Questions soliciting this information should be phrased to encourage capture of genetic sex where known.  However, systems performing analysis should also allow for the possibility of imprecision with this element.
	Sex *CodeableConcept `bson:"sex,omitempty" json:"sex,omitempty" xml:"sex,omitempty"`
	// The actual or approximate date of birth of the relative.
	BornPeriod *Period `bson:"bornPeriod,omitempty" json:"bornPeriod,omitempty" xml:"bornPeriod,omitempty"`
	BornDate   *string `bson:"bornDate,omitempty" json:"bornDate,omitempty" xml:"bornDate,omitempty"`
	BornString *string `bson:"bornString,omitempty" json:"bornString,omitempty" xml:"bornString,omitempty"`
	// The age of the relative at the time the family member history is recorded.
	//
	// Comment: use estimatedAge to indicate whether the age is actual or not.
	AgeAge    *Age    `bson:"ageAge,omitempty" json:"ageAge,omitempty" xml:"ageAge,omitempty"`
	AgeRange  *Range  `bson:"ageRange,omitempty" json:"ageRange,omitempty" xml:"ageRange,omitempty"`
	AgeString *string `bson:"ageString,omitempty" json:"ageString,omitempty" xml:"ageString,omitempty"`
	// If true, indicates that the age value specified is an estimated value.
	//
	// Comment: This element is labeled as a modifier because the fact that age is estimated can/should change the results of any algorithm that calculates based on the specified age.
	EstimatedAge *bool `bson:"estimatedAge,omitempty" json:"estimatedAge,omitempty" xml:"estimatedAge,omitempty"`
	// Deceased flag or the actual or approximate age of the relative at the time of death for the family member history record.
	DeceasedBoolean *bool   `bson:"deceasedBoolean,omitempty" json:"deceasedBoolean,omitempty" xml:"deceasedBoolean,omitempty"`
	DeceasedAge     *Age    `bson:"deceasedAge,omitempty" json:"deceasedAge,omitempty" xml:"deceasedAge,omitempty"`
	DeceasedRange   *Range  `bson:"deceasedRange,omitempty" json:"deceasedRange,omitempty" xml:"deceasedRange,omitempty"`
	DeceasedDate    *string `bson:"deceasedDate,omitempty" json:"deceasedDate,omitempty" xml:"deceasedDate,omitempty"`
	DeceasedString  *string `bson:"deceasedString,omitempty" json:"deceasedString,omitempty" xml:"deceasedString,omitempty"`
	// Describes why the family member history occurred in coded or textual form, or Indicates a Condition, Observation, AllergyIntolerance, or QuestionnaireResponse that justifies this family member history event.
	//
	// Comment: Textual reasons can be captured using reasonCode.text.
	Reason []CodeableReference `bson:"reason,omitempty" json:"reason,omitempty" xml:"reason,omitempty"`
	// This property allows a non condition-specific note to the made about the related person.
	//
	// Ideally, the note would be in the condition property, but this is not always possible.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// The significant Conditions (or condition) that the family member had.
	//
	// This is a repeating section to allow a system to represent more than one condition per resource, though there is nothing stopping multiple resources - one per condition.
	Condition []FamilyMemberHistoryCondition `bson:"condition,omitempty" json:"condition,omitempty" xml:"condition,omitempty"`
	// The significant Procedures (or procedure) that the family member had.
	//
	// This is a repeating section to allow a system to represent more than one procedure per resource, though there is nothing stopping multiple resources - one per procedure.
	Procedure []FamilyMemberHistoryProcedure `bson:"procedure,omitempty" json:"procedure,omitempty" xml:"procedure,omitempty"`
}

type FamilyMemberHistoryParticipant struct {
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
	// Distinguishes the type of involvement of the actor in the activities related to the family member history.
	Function *CodeableConcept `bson:"function,omitempty" json:"function,omitempty" xml:"function,omitempty"`
	// Indicates who or what participated in the activities related to the family member history.
	Actor Reference `bson:"actor" json:"actor" xml:"actor"`
}

type FamilyMemberHistoryCondition struct {
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
	// The actual condition specified.
	//
	// Could be a coded condition (like MI or Diabetes) or a less specific string like 'cancer' depending on how much is known about the condition and the capabilities of the creating system.
	Code CodeableConcept `bson:"code" json:"code" xml:"code"`
	// Indicates what happened following the condition.
	//
	//  If the condition resulted in death, deceased date is captured on the relation.
	Outcome *CodeableConcept `bson:"outcome,omitempty" json:"outcome,omitempty" xml:"outcome,omitempty"`
	// This condition contributed to the cause of death of the related person.
	//
	// If contributedToDeath is not populated, then it is unknown.
	ContributedToDeath *bool `bson:"contributedToDeath,omitempty" json:"contributedToDeath,omitempty" xml:"contributedToDeath,omitempty"`
	// Either the age of onset, range of approximate age or descriptive string can be recorded.
	//
	//  For conditions with multiple occurrences, this describes the first known occurrence.
	OnsetAge    *Age    `bson:"onsetAge,omitempty" json:"onsetAge,omitempty" xml:"onsetAge,omitempty"`
	OnsetRange  *Range  `bson:"onsetRange,omitempty" json:"onsetRange,omitempty" xml:"onsetRange,omitempty"`
	OnsetPeriod *Period `bson:"onsetPeriod,omitempty" json:"onsetPeriod,omitempty" xml:"onsetPeriod,omitempty"`
	OnsetString *string `bson:"onsetString,omitempty" json:"onsetString,omitempty" xml:"onsetString,omitempty"`
	// An area where general notes can be placed about this specific condition.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
}

type FamilyMemberHistoryProcedure struct {
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
	// The actual procedure specified.
	//
	// Could be a coded procedure or a less specific string depending on how much is known about the procedure and the capabilities of the creating system.
	Code CodeableConcept `bson:"code" json:"code" xml:"code"`
	// Indicates what happened following the procedure.
	//
	// If the procedure resulted in death, deceased date is captured on the relation.
	Outcome *CodeableConcept `bson:"outcome,omitempty" json:"outcome,omitempty" xml:"outcome,omitempty"`
	// This procedure contributed to the cause of death of the related person.
	//
	// If contributedToDeath is not populated, then it is unknown.
	ContributedToDeath *bool `bson:"contributedToDeath,omitempty" json:"contributedToDeath,omitempty" xml:"contributedToDeath,omitempty"`
	// Estimated or actual date, date-time, period, or age when the procedure was performed.
	//
	// Allows a period to support complex procedures that span more than one date, and also allows for the length of the procedure to be captured.
	PerformedAge      *Age    `bson:"performedAge,omitempty" json:"performedAge,omitempty" xml:"performedAge,omitempty"`
	PerformedRange    *Range  `bson:"performedRange,omitempty" json:"performedRange,omitempty" xml:"performedRange,omitempty"`
	PerformedPeriod   *Period `bson:"performedPeriod,omitempty" json:"performedPeriod,omitempty" xml:"performedPeriod,omitempty"`
	PerformedString   *string `bson:"performedString,omitempty" json:"performedString,omitempty" xml:"performedString,omitempty"`
	PerformedDateTime *string `bson:"performedDateTime,omitempty" json:"performedDateTime,omitempty" xml:"performedDateTime,omitempty"`
	// An area where general notes can be placed about this specific procedure.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
}

func (r FamilyMemberHistory) ResourceType() string {
	return "FamilyMemberHistory"
}

// MarshalJSON implements json.Marshaler
func (r FamilyMemberHistory) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		FamilyMemberHistory
		ResourceType string `json:"resourceType"`
	}{
		FamilyMemberHistory: r,
		ResourceType:        r.ResourceType(),
	})
}
