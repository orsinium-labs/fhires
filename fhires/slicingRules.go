package fhires

// How slices are interpreted when evaluating an instance.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/resource-slicing-rules
type SlicingRules string

const (
	// No additional content is allowed other than that described by the slices in this profile.
	SlicingRulesClosed SlicingRules = "closed"
	// Additional content is allowed anywhere in the list.
	SlicingRulesOpen SlicingRules = "open"
	// Additional content is allowed, but only at the end of the list. Note that using this requires that the slices be ordered, which makes it hard to share uses. This should only be done where absolutely required.
	SlicingRulesOpenAtEnd SlicingRules = "openAtEnd"
)

func (code SlicingRules) String() string {
	return code.Code()
}

func (code SlicingRules) Code() string {
	return string(code)
}

func (code SlicingRules) Display() string {
	switch code {
	case SlicingRulesClosed:
		return "Closed"
	case SlicingRulesOpen:
		return "Open"
	case SlicingRulesOpenAtEnd:
		return "Open at End"
	}
	return "<unknown>"
}

func (code SlicingRules) Definition() string {
	switch code {
	case SlicingRulesClosed:
		return "No additional content is allowed other than that described by the slices in this profile."
	case SlicingRulesOpen:
		return "Additional content is allowed anywhere in the list."
	case SlicingRulesOpenAtEnd:
		return "Additional content is allowed, but only at the end of the list. Note that using this requires that the slices be ordered, which makes it hard to share uses. This should only be done where absolutely required."
	}
	return "<unknown>"
}
