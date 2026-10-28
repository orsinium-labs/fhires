package fhires

// HumanName Type: A name, normally of a human, that can be used for other living entities (e.g.
//
// animals but not organizations) that have been assigned names by a human and may need the use of name parts or the need for usage information.
//
// Purpose: Need to be able to record names, along with notes about their use.
//
// Status: active
//
// http://hl7.org/fhir/StructureDefinition/HumanName
type HumanName struct {
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
	// Identifies the purpose for this name.
	//
	// Comment: Applications can assume that a name is current unless it explicitly says that it is temporary or old.
	Use *NameUse `bson:"use,omitempty" json:"use,omitempty" xml:"use,omitempty"`
	// Specifies the entire name as it should be displayed e.g.
	//
	// on an application UI. This may be provided instead of or as well as the specific parts.
	//
	// Comment: Can provide both a text representation and parts. Applications updating a name SHALL ensure that when both text and parts are present,  no content is included in the text that isn't found in a part.
	Text *string `bson:"text,omitempty" json:"text,omitempty" xml:"text,omitempty"`
	// The part of a name that links to the genealogy.
	//
	// In some cultures (e.g. Eritrea) the family name of a son is the first name of his father.
	//
	// Comment: Family Name may be decomposed into specific parts using extensions (de, nl, es related cultures).
	Family *string `bson:"family,omitempty" json:"family,omitempty" xml:"family,omitempty"`
	// Given name.
	//
	// Comment: If only initials are recorded, they may be used in place of the full name parts. Initials may be separated into multiple given names but often aren't due to paractical limitations.  This element is not called "first name" since given names do not always come first.
	Given []string `bson:"given,omitempty" json:"given,omitempty" xml:"given,omitempty"`
	// Part of the name that is acquired as a title due to academic, legal, employment or nobility status, etc.
	//
	// and that appears at the start of the name.
	Prefix []string `bson:"prefix,omitempty" json:"prefix,omitempty" xml:"prefix,omitempty"`
	// Part of the name that is acquired as a title due to academic, legal, employment or nobility status, etc.
	//
	// and that appears at the end of the name.
	Suffix []string `bson:"suffix,omitempty" json:"suffix,omitempty" xml:"suffix,omitempty"`
	// Indicates the period of time when this name was valid for the named person.
	Period *Period `bson:"period,omitempty" json:"period,omitempty" xml:"period,omitempty"`
}
