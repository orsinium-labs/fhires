package fhires

// Period Type: A time period defined by a start and end date and optionally time.
//
// Status: active
//
// http://hl7.org/fhir/StructureDefinition/Period
type Period struct {
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
	// The start of the period. The boundary is inclusive.
	//
	// Comment: If the low element is missing, the meaning is that the low boundary is not known.
	Start *string `bson:"start,omitempty" json:"start,omitempty" xml:"start,omitempty"`
	// The end of the period.
	//
	// If the end of the period is missing, it means no end was known or planned at the time the instance was created. The start may be in the past, and the end date in the future, which means that period is expected/planned to end at that time.
	//
	// Comment: The end value includes any matching date/time. i.e. 2012-02-03T10:00:00 is in a period that has an end value of 2012-02-03.
	End *string `bson:"end,omitempty" json:"end,omitempty" xml:"end,omitempty"`
}
