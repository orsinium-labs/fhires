package fhires

// Kind of precondition for the condition.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/condition-precondition-type
type ConditionPreconditionType string

const (
	// The observation is very sensitive for the condition, but may also indicate other conditions.
	ConditionPreconditionTypeSensitive ConditionPreconditionType = "sensitive"
	// The observation is very specific for this condition, but not particularly sensitive.
	ConditionPreconditionTypeSpecific ConditionPreconditionType = "specific"
)

func (code ConditionPreconditionType) String() string {
	return code.Code()
}

func (code ConditionPreconditionType) Code() string {
	return string(code)
}

func (code ConditionPreconditionType) Display() string {
	switch code {
	case ConditionPreconditionTypeSensitive:
		return "Sensitive"
	case ConditionPreconditionTypeSpecific:
		return "Specific"
	}
	return "<unknown>"
}

func (code ConditionPreconditionType) Definition() string {
	switch code {
	case ConditionPreconditionTypeSensitive:
		return "The observation is very sensitive for the condition, but may also indicate other conditions."
	case ConditionPreconditionTypeSpecific:
		return "The observation is very specific for this condition, but not particularly sensitive."
	}
	return "<unknown>"
}
