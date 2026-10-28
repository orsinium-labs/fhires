package fhires

import v2 "encoding/json/v2"

// A record of significant events/milestones key data throughout the history of an Encounter
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/EncounterHistory
type EncounterHistory struct {
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
	// The Encounter associated with this set of historic values.
	//
	// Comment: This is also used for associating a child's encounter back to the mother's encounter.Refer to the Notes section in the Patient resource for further details.
	Encounter *Reference `bson:"encounter,omitempty" json:"encounter,omitempty" xml:"encounter,omitempty"`
	// Identifier(s) by which this encounter is known.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// planned | in-progress | on-hold | discharged | completed | cancelled | discontinued | entered-in-error | unknown.
	//
	// Comment: Note that internal business rules will determine the appropriate transitions that may occur between statuses (and also classes).
	Status EncounterStatus `bson:"status" json:"status" xml:"status"`
	// Concepts representing classification of patient encounter such as ambulatory (outpatient), inpatient, emergency, home health or others due to local variations.
	Class CodeableConcept `bson:"class" json:"class" xml:"class"`
	// Specific type of encounter (e.g.
	//
	// e-mail consultation, surgical day-care, skilled nursing, rehabilitation).
	//
	// Comment: Since there are many ways to further classify encounters, this element is 0..*.
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Broad categorization of the service that is to be provided (e.g. cardiology).
	ServiceType []CodeableReference `bson:"serviceType,omitempty" json:"serviceType,omitempty" xml:"serviceType,omitempty"`
	// The patient or group related to this encounter.
	//
	// In some use-cases the patient MAY not be present, such as a case meeting about a patient between several practitioners or a careteam.
	//
	// Comment: While the encounter is always about the patient, the patient might not actually be known in all contexts of use, and there may be a group of patients that could be anonymous (such as in a group therapy for Alcoholics Anonymous - where the recording of the encounter could be used for billing on the number of people/staff and not important to the context of the specific patients) or alternately in veterinary care a herd of sheep receiving treatment (where the animals are not individually tracked).
	Subject *Reference `bson:"subject,omitempty" json:"subject,omitempty" xml:"subject,omitempty"`
	// The subjectStatus value can be used to track the patient's status within the encounter.
	//
	// It details whether the patient has arrived or departed, has been triaged or is currently in a waiting status.
	SubjectStatus *CodeableConcept `bson:"subjectStatus,omitempty" json:"subjectStatus,omitempty" xml:"subjectStatus,omitempty"`
	// The start and end time associated with this set of values associated with the encounter, may be different to the planned times for various reasons.
	ActualPeriod *Period `bson:"actualPeriod,omitempty" json:"actualPeriod,omitempty" xml:"actualPeriod,omitempty"`
	// The planned start date/time (or admission date) of the encounter.
	PlannedStartDate *string `bson:"plannedStartDate,omitempty" json:"plannedStartDate,omitempty" xml:"plannedStartDate,omitempty"`
	// The planned end date/time (or discharge date) of the encounter.
	PlannedEndDate *string `bson:"plannedEndDate,omitempty" json:"plannedEndDate,omitempty" xml:"plannedEndDate,omitempty"`
	// Actual quantity of time the encounter lasted.
	//
	// This excludes the time during leaves of absence.When missing it is the time in between the start and end values.
	//
	// Comment: If the precision on these values is low (e.g. to the day only) then this may be considered was an all day (or multi-day) encounter, unless the duration is included, where that amount of time occurred sometime during the interval.May differ from the time in `Encounter.period` due to leave of absence(s).
	Length *Duration `bson:"length,omitempty" json:"length,omitempty" xml:"length,omitempty"`
	// The location of the patient at this point in the encounter, the multiple cardinality permits de-normalizing the levels of the location hierarchy, such as site/ward/room/bed.
	//
	// Comment: Virtual encounters can be recorded in the Encounter by specifying a location reference to a location of type "kind" such as "client's home" and an encounter.class = "virtual".
	Location []EncounterHistoryLocation `bson:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
}

type EncounterHistoryLocation struct {
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
	// The location where the encounter takes place.
	Location Reference `bson:"location" json:"location" xml:"location"`
	// This will be used to specify the required levels (bed/ward/room/etc.) desired to be recorded to simplify either messaging or query.
	//
	// Comment: This information is de-normalized from the Location resource to support the easier understanding of the encounter resource and processing in messaging or query.
	//
	// There may be many levels in the hierachy, and this may only pic specific levels that are required for a specific usage scenario.
	Form *CodeableConcept `bson:"form,omitempty" json:"form,omitempty" xml:"form,omitempty"`
}

func (r EncounterHistory) ResourceType() string {
	return "EncounterHistory"
}

// MarshalJSON implements json.Marshaler
func (r EncounterHistory) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		EncounterHistory
		ResourceType string `json:"resourceType"`
	}{
		EncounterHistory: r,
		ResourceType:     r.ResourceType(),
	})
}
