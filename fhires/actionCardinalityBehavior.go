package fhires

// Defines behavior for an action or a group for how many times that item may be repeated.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/action-cardinality-behavior
type ActionCardinalityBehavior string

const (
	// The action may only be selected one time.
	ActionCardinalityBehaviorSingle ActionCardinalityBehavior = "single"
	// The action may be selected multiple times.
	ActionCardinalityBehaviorMultiple ActionCardinalityBehavior = "multiple"
)

func (code ActionCardinalityBehavior) String() string {
	return code.Code()
}

func (code ActionCardinalityBehavior) Code() string {
	return string(code)
}

func (code ActionCardinalityBehavior) Display() string {
	switch code {
	case ActionCardinalityBehaviorSingle:
		return "Single"
	case ActionCardinalityBehaviorMultiple:
		return "Multiple"
	}
	return "<unknown>"
}

func (code ActionCardinalityBehavior) Definition() string {
	switch code {
	case ActionCardinalityBehaviorSingle:
		return "The action may only be selected one time."
	case ActionCardinalityBehaviorMultiple:
		return "The action may be selected multiple times."
	}
	return "<unknown>"
}
