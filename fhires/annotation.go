package fhires

// Annotation Type: A  text note which also  contains information about who made the statement and when.
//
// Status: active
//
// http://hl7.org/fhir/StructureDefinition/Annotation
type Annotation struct {
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
	// The individual responsible for making the annotation.
	//
	// Comment: Organization is used when there's no need for specific attribution as to who made the comment.
	AuthorReference *Reference `bson:"authorReference,omitempty" json:"authorReference,omitempty" xml:"authorReference,omitempty"`
	AuthorString    *string    `bson:"authorString,omitempty" json:"authorString,omitempty" xml:"authorString,omitempty"`
	// Indicates when this particular annotation was made.
	Time *string `bson:"time,omitempty" json:"time,omitempty" xml:"time,omitempty"`
	// The text of the annotation in markdown format.
	Text string `bson:"text" json:"text" xml:"text"`
}
