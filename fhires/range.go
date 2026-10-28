package fhires

// Range Type: A set of ordered Quantities defined by a low and high limit.
//
// Purpose: Need to be able to specify ranges of values.
//
// Status: active
//
// http://hl7.org/fhir/StructureDefinition/Range
type Range struct {
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
	// The low limit. The boundary is inclusive.
	//
	// Comment: If the low element is missing, the low boundary is not known.
	Low *Quantity `bson:"low,omitempty" json:"low,omitempty" xml:"low,omitempty"`
	// The high limit. The boundary is inclusive.
	//
	// Comment: If the high element is missing, the high boundary is not known.
	High *Quantity `bson:"high,omitempty" json:"high,omitempty" xml:"high,omitempty"`
}
