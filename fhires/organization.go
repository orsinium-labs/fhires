package fhires

import "encoding/xml"

// A formally or informally recognized grouping of people or organizations formed for the purpose of achieving some form of collective action.
//
//	Includes companies, institutions, corporations, departments, community groups, healthcare practice groups, payer/insurer, etc.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Organization
type Organization struct {
	DomainResource
	// Identifier for the organization that is used to identify the organization across multiple disparate systems.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Whether the organization's record is still in active use.
	//
	// Comment: This active flag is not intended to be used to mark an organization as temporarily closed or under construction. Instead the Location(s) within the Organization should have the suspended status. If further details of the reason for the suspension are required, then an extension on this element should be used.
	//
	// This element is labeled as a modifier because it may be used to mark that the resource was created in error.
	Active *bool `bson:"active,omitempty" json:"active,omitempty" xml:"active,omitempty"`
	// The kind(s) of organization that this is.
	//
	// Comment: Organizations can be corporations, wards, sections, clinical teams, government departments, etc. Note that code is generally a classifier of the type of organization; in many applications, codes are used to identity a particular organization (say, ward) as opposed to another of the same type - these are identifiers, not codes
	//
	// When considering if multiple types are appropriate, you should evaluate if child organizations would be a more appropriate use of the concept, as different types likely are in different sub-areas of the organization. This is most likely to be used where type values have orthogonal values, such as a religious, academic and medical center.
	//
	// We expect that some jurisdictions will profile this optionality to be a single cardinality.
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// A name associated with the organization.
	//
	// Comment: If the name of an organization changes, consider putting the old name in the alias column so that it can still be located through searches.
	Name *string `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// A list of alternate names that the organization is known as, or was known as in the past.
	//
	// Comment: There are no dates associated with the alias/historic names, as this is not intended to track when names were used, but to assist in searching so that older names can still result in identifying the organization.
	Alias []string `bson:"alias,omitempty" json:"alias,omitempty" xml:"alias,omitempty"`
	// Description of the organization, which helps provide additional general context on the organization to ensure that the correct organization is selected.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// The contact details of communication devices available relevant to the specific Organization.
	//
	// This can include addresses, phone numbers, fax numbers, mobile numbers, email addresses and web sites.
	//
	// Comment: The address/telecom use code 'home' are not to be used. Note that these contacts are not the contact details of people who provide the service (that would be through PractitionerRole), these are official contacts for the Organization itself for specific purposes. E.g. Mailing Addresses, Billing Addresses, Contact numbers for Booking or Billing Enquiries, general web address, web address for online bookings etc.
	Contact []ExtendedContactDetail `bson:"contact,omitempty" json:"contact,omitempty" xml:"contact,omitempty"`
	// The organization of which this organization forms a part.
	PartOf *Reference `bson:"partOf,omitempty" json:"partOf,omitempty" xml:"partOf,omitempty"`
	// Technical endpoints providing access to services operated for the organization.
	Endpoint []Reference `bson:"endpoint,omitempty" json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// The official certifications, accreditations, training, designations and licenses that authorize and/or otherwise endorse the provision of care by the organization.For example, an approval to provide a type of services issued by a certifying body (such as the US Joint Commission) to an organization.
	Qualification []OrganizationQualification `bson:"qualification,omitempty" json:"qualification,omitempty" xml:"qualification,omitempty"`
}

type OrganizationQualification struct {
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
	// An identifier allocated to this qualification for this organization.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Coded representation of the qualification.
	Code CodeableConcept `bson:"code" json:"code" xml:"code"`
	// Period during which the qualification is valid.
	Period *Period `bson:"period,omitempty" json:"period,omitempty" xml:"period,omitempty"`
	// Organization that regulates and issues the qualification.
	Issuer *Reference `bson:"issuer,omitempty" json:"issuer,omitempty" xml:"issuer,omitempty"`
}

func (r Organization) ResourceType() string {
	return "Organization"
}

// MarshalJSON implements [json.Marshaler]
func (r Organization) MarshalJSON() ([]byte, error) {
	type Wrapper Organization
	return marshalJSON(r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Marshaler].
func (r Organization) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Wrapper Organization
	return marshalXML(e, start, r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Unmarshaler].
func (r *Organization) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper Organization
	w := Wrapper(*r)
	err := unmarshalXML(d, start, r.ResourceType(), &w)
	*r = Organization(w)
	return err
}
