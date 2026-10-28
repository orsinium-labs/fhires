package fhires

// ContactPoint Type: Details for all kinds of technology mediated contact points for a person or organization, including telephone, email, etc.
//
// Purpose: Need to track phone, fax, mobile, sms numbers, email addresses, twitter tags, etc.
//
// Status: active
//
// http://hl7.org/fhir/StructureDefinition/ContactPoint
type ContactPoint struct {
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
	// Telecommunications form for contact point - what communications system is required to make use of the contact.
	System *ContactPointSystem `bson:"system,omitempty" json:"system,omitempty" xml:"system,omitempty"`
	// The actual contact point details, in a form that is meaningful to the designated communication system (i.e.
	//
	// phone number or email address).
	//
	// Comment: Additional text data such as phone extension numbers, or notes about use of the contact are sometimes included in the value.
	Value *string `bson:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
	// Identifies the purpose for the contact point.
	//
	// Comment: Applications can assume that a contact is current unless it explicitly says that it is temporary or old.
	Use *ContactPointUse `bson:"use,omitempty" json:"use,omitempty" xml:"use,omitempty"`
	// Specifies a preferred order in which to use a set of contacts.
	//
	// ContactPoints with lower rank values are more preferred than those with higher rank values.
	//
	// Comment: Note that rank does not necessarily follow the order in which the contacts are represented in the instance.
	//
	// Ranks need not be unique.  E.g. it's possible to have multiple contacts with rank=1.  If the ranks have different systems or uses, this would be interpreted to mean "X is my most preferred phone number, Y is my most preferred email address" or "X is my preferred home email and Y is my preferred work email".  If the system and use for equally-ranked contacts are the same, then the level of preference is equivalent for both repetitions.
	//
	// Ranks need not be sequential and not all repetitions must have a rank.  For example, it's possible to have 4 contacts with ranks of 2, 5 and two with no rank specified.  That would be interpreted to mean the first is preferred over the second and no preference stated for the remaining contacts.
	Rank *int `bson:"rank,omitempty" json:"rank,omitempty" xml:"rank,omitempty"`
	// Time period when the contact point was/is in use.
	Period *Period `bson:"period,omitempty" json:"period,omitempty" xml:"period,omitempty"`
}
