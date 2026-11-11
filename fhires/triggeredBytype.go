package fhires

// Codes providing the type of triggeredBy observation.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/observation-triggeredbytype
type TriggeredBytype string

const (
	// Performance of one or more other tests depending on the results of the initial test.  This may include collection of additional specimen. While a new ServiceRequest is not required to perform the additional test, where it is still needed (e.g., requesting another laboratory to perform the reflex test), the Observation.basedOn would reference the new ServiceRequest that requested the additional test to be performed as well as the original ServiceRequest to reflect the one that provided the authorization.
	TriggeredBytypeReflex TriggeredBytype = "reflex"
	// Performance of the same test again with the same parameters/settings/solution.
	TriggeredBytypeRepeat TriggeredBytype = "repeat"
	// Performance of the same test but with different parameters/settings/solution.
	TriggeredBytypeReRun TriggeredBytype = "re-run"
)

func (code TriggeredBytype) String() string {
	return code.Code()
}

func (code TriggeredBytype) Code() string {
	return string(code)
}

func (code TriggeredBytype) Display() string {
	switch code {
	case TriggeredBytypeReflex:
		return "Reflex"
	case TriggeredBytypeRepeat:
		return "Repeat (per policy)"
	case TriggeredBytypeReRun:
		return "Re-run (per policy)"
	}
	return "<unknown>"
}

func (code TriggeredBytype) Definition() string {
	switch code {
	case TriggeredBytypeReflex:
		return "Performance of one or more other tests depending on the results of the initial test.  This may include collection of additional specimen. While a new ServiceRequest is not required to perform the additional test, where it is still needed (e.g., requesting another laboratory to perform the reflex test), the Observation.basedOn would reference the new ServiceRequest that requested the additional test to be performed as well as the original ServiceRequest to reflect the one that provided the authorization."
	case TriggeredBytypeRepeat:
		return "Performance of the same test again with the same parameters/settings/solution."
	case TriggeredBytypeReRun:
		return "Performance of the same test but with different parameters/settings/solution."
	}
	return "<unknown>"
}
