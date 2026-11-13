package fhires

import "encoding/xml"

// A record of significant events/milestones key data throughout the history of an Encounter
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/EncounterHistory
type EncounterHistory struct {
	DomainResource
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

// MarshalJSON implements [json.Marshaler]
func (r EncounterHistory) MarshalJSON() ([]byte, error) {
	type Wrapper EncounterHistory
	return marshalJSON(r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Marshaler].
func (r EncounterHistory) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Wrapper EncounterHistory
	return marshalXML(e, start, r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Unmarshaler].
func (r *EncounterHistory) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper EncounterHistory
	w := Wrapper(*r)
	err := unmarshalXML(d, start, r.ResourceType(), &w)
	*r = EncounterHistory(w)
	return err
}
