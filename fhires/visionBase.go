package fhires

// A coded concept listing the base codes.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/vision-base-codes
type VisionBase string

const (
	// top.
	VisionBaseUp VisionBase = "up"
	// bottom.
	VisionBaseDown VisionBase = "down"
	// inner edge.
	VisionBaseIn VisionBase = "in"
	// outer edge.
	VisionBaseOut VisionBase = "out"
)

func (code VisionBase) String() string {
	return code.Code()
}

func (code VisionBase) Code() string {
	return string(code)
}

func (code VisionBase) Display() string {
	switch code {
	case VisionBaseUp:
		return "Up"
	case VisionBaseDown:
		return "Down"
	case VisionBaseIn:
		return "In"
	case VisionBaseOut:
		return "Out"
	}
	return "<unknown>"
}

func (code VisionBase) Definition() string {
	switch code {
	case VisionBaseUp:
		return "top."
	case VisionBaseDown:
		return "bottom."
	case VisionBaseIn:
		return "inner edge."
	case VisionBaseOut:
		return "outer edge."
	}
	return "<unknown>"
}
