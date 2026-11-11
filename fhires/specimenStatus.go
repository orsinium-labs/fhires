package fhires

// Codes providing the status/availability of a specimen.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/specimen-status
type SpecimenStatus string

const (
	// The physical specimen is present and in good condition.
	SpecimenStatusAvailable SpecimenStatus = "available"
	// There is no physical specimen because it is either lost, destroyed or consumed.
	SpecimenStatusUnavailable SpecimenStatus = "unavailable"
	// The specimen cannot be used because of a quality issue such as a broken container, contamination, or too old.
	SpecimenStatusUnsatisfactory SpecimenStatus = "unsatisfactory"
	// The specimen was entered in error and therefore nullified.
	SpecimenStatusEnteredInError SpecimenStatus = "entered-in-error"
)

func (code SpecimenStatus) String() string {
	return code.Code()
}

func (code SpecimenStatus) Code() string {
	return string(code)
}

func (code SpecimenStatus) Display() string {
	switch code {
	case SpecimenStatusAvailable:
		return "Available"
	case SpecimenStatusUnavailable:
		return "Unavailable"
	case SpecimenStatusUnsatisfactory:
		return "Unsatisfactory"
	case SpecimenStatusEnteredInError:
		return "Entered in Error"
	}
	return "<unknown>"
}

func (code SpecimenStatus) Definition() string {
	switch code {
	case SpecimenStatusAvailable:
		return "The physical specimen is present and in good condition."
	case SpecimenStatusUnavailable:
		return "There is no physical specimen because it is either lost, destroyed or consumed."
	case SpecimenStatusUnsatisfactory:
		return "The specimen cannot be used because of a quality issue such as a broken container, contamination, or too old."
	case SpecimenStatusEnteredInError:
		return "The specimen was entered in error and therefore nullified."
	}
	return "<unknown>"
}
