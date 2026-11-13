package fhires

import "encoding/xml"

// Defines an affiliation/assotiation/relationship between 2 distinct organizations, that is not a part-of relationship/sub-division relationship.
//
// Purpose: Need to define relationships between organizations that are not sub-divisions of the same organization (part-of relationships).
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/OrganizationAffiliation
type OrganizationAffiliation struct {
	DomainResource
	// Business identifiers that are specific to this role.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Whether this organization affiliation record is in active use.
	//
	// Comment: If this value is false, you may refer to the period to see when the role was in active use. If there is no period specified, no inference can be made about when it was active.
	Active *bool `bson:"active,omitempty" json:"active,omitempty" xml:"active,omitempty"`
	// The period during which the participatingOrganization is affiliated with the primary organization.
	Period *Period `bson:"period,omitempty" json:"period,omitempty" xml:"period,omitempty"`
	// Organization where the role is available (primary organization/has members).
	//
	// Comment: For example, a Spotless Cleaning Services (participatingOrganization) is a supplier (code) to General Hospital (organization).   Similarly, General Hospital (participatingOrganization) is a member (code) of Eastern HIE (organization).  If there are two organizations that do not have a directional relationship, there would normally be a third organization that represents the joint partnership, and the two organizations would have a relationship to that joint Organization via two OrganizationAffiliations.
	Organization *Reference `bson:"organization,omitempty" json:"organization,omitempty" xml:"organization,omitempty"`
	// The Participating Organization provides/performs the role(s) defined by the code to the Primary Organization (e.g.
	//
	// providing services or is a member of).
	//
	// Comment: See comments for OrganizationAffiliation.organization above.
	ParticipatingOrganization *Reference `bson:"participatingOrganization,omitempty" json:"participatingOrganization,omitempty" xml:"participatingOrganization,omitempty"`
	// The network in which the participatingOrganization provides the role's services (if defined) at the indicated locations (if defined).
	//
	// Comment: e.g. Commonly used for Health Insurance provider networks.
	Network []Reference `bson:"network,omitempty" json:"network,omitempty" xml:"network,omitempty"`
	// Definition of the role the participatingOrganization plays in the association.
	Code []CodeableConcept `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// Specific specialty of the participatingOrganization in the context of the role.
	Specialty []CodeableConcept `bson:"specialty,omitempty" json:"specialty,omitempty" xml:"specialty,omitempty"`
	// The location(s) at which the role occurs.
	Location []Reference `bson:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// Healthcare services provided through the role.
	HealthcareService []Reference `bson:"healthcareService,omitempty" json:"healthcareService,omitempty" xml:"healthcareService,omitempty"`
	// The contact details of communication devices available at the participatingOrganization relevant to this Affiliation.
	Contact []ExtendedContactDetail `bson:"contact,omitempty" json:"contact,omitempty" xml:"contact,omitempty"`
	// Technical endpoints providing access to services operated for this role.
	Endpoint []Reference `bson:"endpoint,omitempty" json:"endpoint,omitempty" xml:"endpoint,omitempty"`
}

func (r OrganizationAffiliation) ResourceType() string {
	return "OrganizationAffiliation"
}

// MarshalJSON implements [json.Marshaler]
func (r OrganizationAffiliation) MarshalJSON() ([]byte, error) {
	type Wrapper OrganizationAffiliation
	return marshalJSON(r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Marshaler].
func (r OrganizationAffiliation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Wrapper OrganizationAffiliation
	return marshalXML(e, start, r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Unmarshaler].
func (r *OrganizationAffiliation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper OrganizationAffiliation
	w := Wrapper(*r)
	err := unmarshalXML(d, start, r.ResourceType(), &w)
	*r = OrganizationAffiliation(w)
	return err
}
