package fhires

// Use of contact point.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/contact-point-use
type ContactPointUse string

const (
	// A communication contact point at a home; attempted contacts for business purposes might intrude privacy and chances are one will contact family or other household members instead of the person one wishes to call. Typically used with urgent cases, or if no other contacts are available.
	ContactPointUseHome ContactPointUse = "home"
	// An office contact point. First choice for business related contacts during business hours.
	ContactPointUseWork ContactPointUse = "work"
	// A temporary contact point. The period can provide more detailed information.
	ContactPointUseTemp ContactPointUse = "temp"
	// This contact point is no longer in use (or was never correct, but retained for records).
	ContactPointUseOld ContactPointUse = "old"
	// A telecommunication device that moves and stays with its owner. May have characteristics of all other use codes, suitable for urgent matters, not the first choice for routine business.
	ContactPointUseMobile ContactPointUse = "mobile"
)

func (code ContactPointUse) String() string {
	return code.Code()
}

func (code ContactPointUse) Code() string {
	return string(code)
}

func (code ContactPointUse) Display() string {
	switch code {
	case ContactPointUseHome:
		return "Home"
	case ContactPointUseWork:
		return "Work"
	case ContactPointUseTemp:
		return "Temp"
	case ContactPointUseOld:
		return "Old"
	case ContactPointUseMobile:
		return "Mobile"
	}
	return "<unknown>"
}

func (code ContactPointUse) Definition() string {
	switch code {
	case ContactPointUseHome:
		return "A communication contact point at a home; attempted contacts for business purposes might intrude privacy and chances are one will contact family or other household members instead of the person one wishes to call. Typically used with urgent cases, or if no other contacts are available."
	case ContactPointUseWork:
		return "An office contact point. First choice for business related contacts during business hours."
	case ContactPointUseTemp:
		return "A temporary contact point. The period can provide more detailed information."
	case ContactPointUseOld:
		return "This contact point is no longer in use (or was never correct, but retained for records)."
	case ContactPointUseMobile:
		return "A telecommunication device that moves and stays with its owner. May have characteristics of all other use codes, suitable for urgent matters, not the first choice for routine business."
	}
	return "<unknown>"
}
