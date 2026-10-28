package fhires

// ExtendedContactDetail Type: Specifies contact information for a specific purpose over a period of time, might be handled/monitored by a specific named person or organization.
//
// Purpose: Need to track contact and address information in the same way across multiple resources.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/ExtendedContactDetail
type ExtendedContactDetail struct {
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
	// The purpose/type of contact.
	//
	// Comment: If no purpose is defined, then these contact details may be used for any purpose.
	Purpose *CodeableConcept `bson:"purpose,omitempty" json:"purpose,omitempty" xml:"purpose,omitempty"`
	// The name of an individual to contact, some types of contact detail are usually blank.
	//
	// Comment: If there is no named individual, the telecom/address information is not generally monitored by a specific individual.
	Name []HumanName `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// The contact details application for the purpose defined.
	Telecom []ContactPoint `bson:"telecom,omitempty" json:"telecom,omitempty" xml:"telecom,omitempty"`
	// Address for the contact.
	//
	// Comment: More than 1 address would be for different purposes, and thus should be entered as a different entry,.
	Address *Address `bson:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
	// This contact detail is handled/monitored by a specific organization.
	//
	// If the name is provided in the contact, then it is referring to the named individual within this organization.
	Organization *Reference `bson:"organization,omitempty" json:"organization,omitempty" xml:"organization,omitempty"`
	// Period that this contact was valid for usage.
	//
	// Comment: If the details have multiple periods, then enter in a new ExtendedContact with the new period.
	Period *Period `bson:"period,omitempty" json:"period,omitempty" xml:"period,omitempty"`
}
