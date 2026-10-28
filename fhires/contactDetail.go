package fhires

// ContactDetail Type: Specifies contact information for a person or organization.
//
// Purpose: Need to track contact information in the same way across multiple resources.
//
// Status: active
//
// http://hl7.org/fhir/StructureDefinition/ContactDetail
type ContactDetail struct {
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
	// The name of an individual to contact.
	//
	// Comment: If there is no named individual, the telecom information is for the organization as a whole.
	Name *string `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// The contact details for the individual (if a name was provided) or the organization.
	Telecom []ContactPoint `bson:"telecom,omitempty" json:"telecom,omitempty" xml:"telecom,omitempty"`
}
