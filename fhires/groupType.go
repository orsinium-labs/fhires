package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// Types of resources that are part of group.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/group-type
type GroupType int

const (
	GroupTypePerson GroupType = iota
	// Group contains "animal" Patient resources.
	GroupTypeAnimal
	// Group contains healthcare practitioner resources (Practitioner or PractitionerRole).
	GroupTypePractitioner
	// Group contains Device resources.
	GroupTypeDevice
	// Group contains CareTeam resources.
	GroupTypeCareteam
	// Group contains HealthcareService resources.
	GroupTypeHealthcareservice
	// Group contains Location resources.
	GroupTypeLocation
	// Group contains Organization resources.
	GroupTypeOrganization
	// Group contains RelatedPerson resources.
	GroupTypeRelatedperson
	// Group contains Specimen resources.
	GroupTypeSpecimen
)

func (code GroupType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *GroupType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "person":
		*code = GroupTypePerson
	case "animal":
		*code = GroupTypeAnimal
	case "practitioner":
		*code = GroupTypePractitioner
	case "device":
		*code = GroupTypeDevice
	case "careteam":
		*code = GroupTypeCareteam
	case "healthcareservice":
		*code = GroupTypeHealthcareservice
	case "location":
		*code = GroupTypeLocation
	case "organization":
		*code = GroupTypeOrganization
	case "relatedperson":
		*code = GroupTypeRelatedperson
	case "specimen":
		*code = GroupTypeSpecimen
	default:
		return fmt.Errorf("unknown GroupType code `%s`", s)
	}
	return nil
}

func (code GroupType) String() string {
	return code.Code()
}

func (code GroupType) Code() string {
	switch code {
	case GroupTypePerson:
		return "person"
	case GroupTypeAnimal:
		return "animal"
	case GroupTypePractitioner:
		return "practitioner"
	case GroupTypeDevice:
		return "device"
	case GroupTypeCareteam:
		return "careteam"
	case GroupTypeHealthcareservice:
		return "healthcareservice"
	case GroupTypeLocation:
		return "location"
	case GroupTypeOrganization:
		return "organization"
	case GroupTypeRelatedperson:
		return "relatedperson"
	case GroupTypeSpecimen:
		return "specimen"
	}
	return "<unknown>"
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
