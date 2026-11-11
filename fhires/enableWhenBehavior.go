package fhires

// Controls how multiple enableWhen values are interpreted -  whether all or any must be true.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/questionnaire-enable-behavior
type EnableWhenBehavior string

const (
	// Enable the question when all the enableWhen criteria are satisfied.
	EnableWhenBehaviorAll EnableWhenBehavior = "all"
	// Enable the question when any of the enableWhen criteria are satisfied.
	EnableWhenBehaviorAny EnableWhenBehavior = "any"
)

func (code EnableWhenBehavior) String() string {
	return code.Code()
}

func (code EnableWhenBehavior) Code() string {
	return string(code)
}

func (code EnableWhenBehavior) Display() string {
	switch code {
	case EnableWhenBehaviorAll:
		return "All"
	case EnableWhenBehaviorAny:
		return "Any"
	}
	return "<unknown>"
}

func (code EnableWhenBehavior) Definition() string {
	switch code {
	case EnableWhenBehaviorAll:
		return "Enable the question when all the enableWhen criteria are satisfied."
	case EnableWhenBehaviorAny:
		return "Enable the question when any of the enableWhen criteria are satisfied."
	}
	return "<unknown>"
}
