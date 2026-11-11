package fhires

// SHALL applications comply with this constraint?
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/constraint-severity
type ConstraintSeverity string

const (
	// If the constraint is violated, the resource is not conformant.
	ConstraintSeverityError ConstraintSeverity = "error"
	// If the constraint is violated, the resource is conformant, but it is not necessarily following best practice.
	ConstraintSeverityWarning ConstraintSeverity = "warning"
)

func (code ConstraintSeverity) String() string {
	return code.Code()
}

func (code ConstraintSeverity) Code() string {
	return string(code)
}

func (code ConstraintSeverity) Display() string {
	switch code {
	case ConstraintSeverityError:
		return "Error"
	case ConstraintSeverityWarning:
		return "Warning"
	}
	return "<unknown>"
}

func (code ConstraintSeverity) Definition() string {
	switch code {
	case ConstraintSeverityError:
		return "If the constraint is violated, the resource is not conformant."
	case ConstraintSeverityWarning:
		return "If the constraint is violated, the resource is conformant, but it is not necessarily following best practice."
	}
	return "<unknown>"
}
