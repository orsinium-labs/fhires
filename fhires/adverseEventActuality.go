package fhires

// Overall nature of the adverse event, e.g.
//
// real or potential.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/adverse-event-actuality
type AdverseEventActuality string

const (
	// The adverse event actually happened regardless of whether anyone was affected or harmed.
	AdverseEventActualityActual AdverseEventActuality = "actual"
	// A potential adverse event.
	AdverseEventActualityPotential AdverseEventActuality = "potential"
)

func (code AdverseEventActuality) String() string {
	return code.Code()
}

func (code AdverseEventActuality) Code() string {
	return string(code)
}

func (code AdverseEventActuality) Display() string {
	switch code {
	case AdverseEventActualityActual:
		return "Adverse Event"
	case AdverseEventActualityPotential:
		return "Potential Adverse Event"
	}
	return "<unknown>"
}

func (code AdverseEventActuality) Definition() string {
	switch code {
	case AdverseEventActualityActual:
		return "The adverse event actually happened regardless of whether anyone was affected or harmed."
	case AdverseEventActualityPotential:
		return "A potential adverse event."
	}
	return "<unknown>"
}
