package fhires

// Defines selection behavior of a group.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/action-selection-behavior
type ActionSelectionBehavior string

const (
	// Any number of the actions in the group may be chosen, from zero to all.
	ActionSelectionBehaviorAny ActionSelectionBehavior = "any"
	// All the actions in the group must be selected as a single unit.
	ActionSelectionBehaviorAll ActionSelectionBehavior = "all"
	// All the actions in the group are meant to be chosen as a single unit: either all must be selected by the end user, or none may be selected.
	ActionSelectionBehaviorAllOrNone ActionSelectionBehavior = "all-or-none"
	// The end user must choose one and only one of the selectable actions in the group. The user SHALL NOT choose none of the actions in the group.
	ActionSelectionBehaviorExactlyOne ActionSelectionBehavior = "exactly-one"
	// The end user may choose zero or at most one of the actions in the group.
	ActionSelectionBehaviorAtMostOne ActionSelectionBehavior = "at-most-one"
	// The end user must choose a minimum of one, and as many additional as desired.
	ActionSelectionBehaviorOneOrMore ActionSelectionBehavior = "one-or-more"
)

func (code ActionSelectionBehavior) String() string {
	return code.Code()
}

func (code ActionSelectionBehavior) Code() string {
	return string(code)
}

func (code ActionSelectionBehavior) Display() string {
	switch code {
	case ActionSelectionBehaviorAny:
		return "Any"
	case ActionSelectionBehaviorAll:
		return "All"
	case ActionSelectionBehaviorAllOrNone:
		return "All Or None"
	case ActionSelectionBehaviorExactlyOne:
		return "Exactly One"
	case ActionSelectionBehaviorAtMostOne:
		return "At Most One"
	case ActionSelectionBehaviorOneOrMore:
		return "One Or More"
	}
	return "<unknown>"
}

func (code ActionSelectionBehavior) Definition() string {
	switch code {
	case ActionSelectionBehaviorAny:
		return "Any number of the actions in the group may be chosen, from zero to all."
	case ActionSelectionBehaviorAll:
		return "All the actions in the group must be selected as a single unit."
	case ActionSelectionBehaviorAllOrNone:
		return "All the actions in the group are meant to be chosen as a single unit: either all must be selected by the end user, or none may be selected."
	case ActionSelectionBehaviorExactlyOne:
		return "The end user must choose one and only one of the selectable actions in the group. The user SHALL NOT choose none of the actions in the group."
	case ActionSelectionBehaviorAtMostOne:
		return "The end user may choose zero or at most one of the actions in the group."
	case ActionSelectionBehaviorOneOrMore:
		return "The end user must choose a minimum of one, and as many additional as desired."
	}
	return "<unknown>"
}
