package fhires

// Type for orientation.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/orientation-type
type OrientationType string

const (
	// Sense orientation of reference sequence.
	OrientationTypeSense OrientationType = "sense"
	// Antisense orientation of reference sequence.
	OrientationTypeAntisense OrientationType = "antisense"
)

func (code OrientationType) String() string {
	return code.Code()
}

func (code OrientationType) Code() string {
	return string(code)
}

func (code OrientationType) Display() string {
	switch code {
	case OrientationTypeSense:
		return "Sense orientation of referenceSeq"
	case OrientationTypeAntisense:
		return "Antisense orientation of referenceSeq"
	}
	return "<unknown>"
}

func (code OrientationType) Definition() string {
	switch code {
	case OrientationTypeSense:
		return "Sense orientation of reference sequence."
	case OrientationTypeAntisense:
		return "Antisense orientation of reference sequence."
	}
	return "<unknown>"
}
