package fhires

// Types of resources that are part of group.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/group-type
type GroupType string

const (
	// Group contains "person" Patient resources.
	GroupTypePerson GroupType = "person"
	// Group contains "animal" Patient resources.
	GroupTypeAnimal GroupType = "animal"
	// Group contains healthcare practitioner resources (Practitioner or PractitionerRole).
	GroupTypePractitioner GroupType = "practitioner"
	// Group contains Device resources.
	GroupTypeDevice GroupType = "device"
	// Group contains CareTeam resources.
	GroupTypeCareteam GroupType = "careteam"
	// Group contains HealthcareService resources.
	GroupTypeHealthcareservice GroupType = "healthcareservice"
	// Group contains Location resources.
	GroupTypeLocation GroupType = "location"
	// Group contains Organization resources.
	GroupTypeOrganization GroupType = "organization"
	// Group contains RelatedPerson resources.
	GroupTypeRelatedperson GroupType = "relatedperson"
	// Group contains Specimen resources.
	GroupTypeSpecimen GroupType = "specimen"
)

func (code GroupType) String() string {
	return code.Code()
}

func (code GroupType) Code() string {
	return string(code)
}

func (code GroupType) Display() string {
	switch code {
	case GroupTypePerson:
		return "Person"
	case GroupTypeAnimal:
		return "Animal"
	case GroupTypePractitioner:
		return "Practitioner"
	case GroupTypeDevice:
		return "Device"
	case GroupTypeCareteam:
		return "CareTeam"
	case GroupTypeHealthcareservice:
		return "HealthcareService"
	case GroupTypeLocation:
		return "Location"
	case GroupTypeOrganization:
		return "Organization"
	case GroupTypeRelatedperson:
		return "RelatedPerson"
	case GroupTypeSpecimen:
		return "Specimen"
	}
	return "<unknown>"
}

func (code GroupType) Definition() string {
	switch code {
	case GroupTypePerson:
		return "Group contains \"person\" Patient resources."
	case GroupTypeAnimal:
		return "Group contains \"animal\" Patient resources."
	case GroupTypePractitioner:
		return "Group contains healthcare practitioner resources (Practitioner or PractitionerRole)."
	case GroupTypeDevice:
		return "Group contains Device resources."
	case GroupTypeCareteam:
		return "Group contains CareTeam resources."
	case GroupTypeHealthcareservice:
		return "Group contains HealthcareService resources."
	case GroupTypeLocation:
		return "Group contains Location resources."
	case GroupTypeOrganization:
		return "Group contains Organization resources."
	case GroupTypeRelatedperson:
		return "Group contains RelatedPerson resources."
	case GroupTypeSpecimen:
		return "Group contains Specimen resources."
	}
	return "<unknown>"
}
