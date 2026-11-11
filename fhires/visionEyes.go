package fhires

// A coded concept listing the eye codes.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/vision-eye-codes
type VisionEyes string

const (
	// Right Eye.
	VisionEyesRight VisionEyes = "right"
	// Left Eye.
	VisionEyesLeft VisionEyes = "left"
)

func (code VisionEyes) String() string {
	return code.Code()
}

func (code VisionEyes) Code() string {
	return string(code)
}

func (code VisionEyes) Display() string {
	switch code {
	case VisionEyesRight:
		return "Right Eye"
	case VisionEyesLeft:
		return "Left Eye"
	}
	return "<unknown>"
}

func (code VisionEyes) Definition() string {
	switch code {
	case VisionEyesRight:
		return "Right Eye."
	case VisionEyesLeft:
		return "Left Eye."
	}
	return "<unknown>"
}
