package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// GroupMembershipBasis is documented here http://hl7.org/fhir/ValueSet/group-membership-basis
type GroupMembershipBasis int

const (
	GroupMembershipBasisDefinitional GroupMembershipBasis = iota
	GroupMembershipBasisEnumerated
)

func (code GroupMembershipBasis) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *GroupMembershipBasis) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "definitional":
		*code = GroupMembershipBasisDefinitional
	case "enumerated":
		*code = GroupMembershipBasisEnumerated
	default:
		return fmt.Errorf("unknown GroupMembershipBasis code `%s`", s)
	}
	return nil
}

func (code GroupMembershipBasis) String() string {
	return code.Code()
}

func (code GroupMembershipBasis) Code() string {
	switch code {
	case GroupMembershipBasisDefinitional:
		return "definitional"
	case GroupMembershipBasisEnumerated:
		return "enumerated"
	}
	return "<unknown>"
}

func (code GroupMembershipBasis) Display() string {
	switch code {
	case GroupMembershipBasisDefinitional:
		return "Definitional"
	case GroupMembershipBasisEnumerated:
		return "Enumerated"
	}
	return "<unknown>"
}

func (code GroupMembershipBasis) Definition() string {
	switch code {
	case GroupMembershipBasisDefinitional:
		return "The Group.characteristics specified are both necessary and sufficient to determine membership. All entities that meet the criteria are considered to be members of the group, whether referenced by the group or not. If members are present, they are individuals that happen to be known as meeting the Group.characteristics. The list cannot be presumed to be complete."
	case GroupMembershipBasisEnumerated:
		return "The Group.characteristics are necessary but not sufficient to determine membership. Membership is determined by being listed as one of the Group.member."
	}
	return "<unknown>"
}
