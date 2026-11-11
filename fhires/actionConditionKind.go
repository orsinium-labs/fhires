package fhires

// Defines the kinds of conditions that can appear on actions.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/action-condition-kind
type ActionConditionKind string

const (
	// The condition describes whether or not a given action is applicable.
	ActionConditionKindApplicability ActionConditionKind = "applicability"
	// The condition is a starting condition for the action.
	ActionConditionKindStart ActionConditionKind = "start"
	// The condition is a stop, or exit condition for the action.
	ActionConditionKindStop ActionConditionKind = "stop"
)

func (code ActionConditionKind) String() string {
	return code.Code()
}

func (code ActionConditionKind) Code() string {
	return string(code)
}

func (code ActionConditionKind) Display() string {
	switch code {
	case ActionConditionKindApplicability:
		return "Applicability"
	case ActionConditionKindStart:
		return "Start"
	case ActionConditionKindStop:
		return "Stop"
	}
	return "<unknown>"
}

func (code ActionConditionKind) Definition() string {
	switch code {
	case ActionConditionKindApplicability:
		return "The condition describes whether or not a given action is applicable."
	case ActionConditionKindStart:
		return "The condition is a starting condition for the action."
	case ActionConditionKindStop:
		return "The condition is a stop, or exit condition for the action."
	}
	return "<unknown>"
}
