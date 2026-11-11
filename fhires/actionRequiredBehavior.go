package fhires

// Defines expectations around whether an action or action group is required.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/action-required-behavior
type ActionRequiredBehavior string

const (
	// An action with this behavior must be included in the actions processed by the end user; the end user SHALL NOT choose not to include this action.
	ActionRequiredBehaviorMust ActionRequiredBehavior = "must"
	// An action with this behavior may be included in the set of actions processed by the end user.
	ActionRequiredBehaviorCould ActionRequiredBehavior = "could"
	// An action with this behavior must be included in the set of actions processed by the end user, unless the end user provides documentation as to why the action was not included.
	ActionRequiredBehaviorMustUnlessDocumented ActionRequiredBehavior = "must-unless-documented"
)

func (code ActionRequiredBehavior) String() string {
	return code.Code()
}

func (code ActionRequiredBehavior) Code() string {
	return string(code)
}

func (code ActionRequiredBehavior) Display() string {
	switch code {
	case ActionRequiredBehaviorMust:
		return "Must"
	case ActionRequiredBehaviorCould:
		return "Could"
	case ActionRequiredBehaviorMustUnlessDocumented:
		return "Must Unless Documented"
	}
	return "<unknown>"
}

func (code ActionRequiredBehavior) Definition() string {
	switch code {
	case ActionRequiredBehaviorMust:
		return "An action with this behavior must be included in the actions processed by the end user; the end user SHALL NOT choose not to include this action."
	case ActionRequiredBehaviorCould:
		return "An action with this behavior may be included in the set of actions processed by the end user."
	case ActionRequiredBehaviorMustUnlessDocumented:
		return "An action with this behavior must be included in the set of actions processed by the end user, unless the end user provides documentation as to why the action was not included."
	}
	return "<unknown>"
}
