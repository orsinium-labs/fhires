package fhires

// Basis for membership in a group
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/group-membership-basis
type GroupMembershipBasis string

const (
	// The Group.characteristics specified are both necessary and sufficient to determine membership. All entities that meet the criteria are considered to be members of the group, whether referenced by the group or not. If members are present, they are individuals that happen to be known as meeting the Group.characteristics. The list cannot be presumed to be complete.
	GroupMembershipBasisDefinitional GroupMembershipBasis = "definitional"
	// The Group.characteristics are necessary but not sufficient to determine membership. Membership is determined by being listed as one of the Group.member.
	GroupMembershipBasisEnumerated GroupMembershipBasis = "enumerated"
)

func (code GroupMembershipBasis) String() string {
	return code.Code()
}

func (code GroupMembershipBasis) Code() string {
	return string(code)
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
