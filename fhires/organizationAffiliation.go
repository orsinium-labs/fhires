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
	BaseResource
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human.
	//
	// The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	//
	// Comment: Contained resources do not have a narrative. Resources that are not contained SHOULD have a narrative. In some cases, a resource may only have text with little or no additional discrete data (as long as all minOccurs=1 elements are satisfied).  This may be necessary for data from legacy systems where information is captured as a "text blob" or where text is additionally entered raw or narrated and encoded information is added later.
	Text *Narrative `bson:"text,omitempty" json:"text,omitempty" xml:"text,omitempty"`
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, nor can they have their own independent transaction scope.
	//
	// This is allowed to be a Parameters resource if and only if it is referenced by a resource that provides context/meaning.
	//
	// Comment: This should never be done when the content can be identified properly, as once identification is lost, it is extremely difficult (and context dependent) to restore it again. Contained resources may have profiles and tags in their meta elements, but SHALL NOT have security labels.
	// May be used to represent additional information that is not part of the basic definition of the resource.
	//
	// To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	//
	// Comment: There can be no stigma associated with the use of extensions by any application, project, or standard - regardless of the institution or jurisdiction that uses or defines the extensions.  The use of extensions is what allows the FHIR specification to retain a core level of simplicity for everyone.
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty" xml:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants.
	//
	// Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	//
	// Comment: There can be no stigma associated with the use of extensions by any application, project, or standard - regardless of the institution or jurisdiction that uses or defines the extensions.  The use of extensions is what allows the FHIR specification to retain a core level of simplicity for everyone.
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty" xml:"modifierExtension,omitempty"`
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

// UnmarshalXML implements [xml.Unmarshaler].
func (r *OrganizationAffiliation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper OrganizationAffiliation
	w := Wrapper(*r)
	err := unmarshalXML(d, start, r.ResourceType(), &w)
	*r = OrganizationAffiliation(w)
	return err
}
