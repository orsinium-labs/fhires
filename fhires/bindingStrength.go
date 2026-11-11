package fhires

// Indication of the degree of conformance expectations associated with a binding.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/binding-strength
type BindingStrength string

const (
	// To be conformant, the concept in this element SHALL be from the specified value set.
	BindingStrengthRequired BindingStrength = "required"
	// To be conformant, the concept in this element SHALL be from the specified value set if any of the codes within the value set can apply to the concept being communicated.  If the value set does not cover the concept (based on human review), alternate codings (or, data type allowing, text) may be included instead.
	BindingStrengthExtensible BindingStrength = "extensible"
	// Instances are encouraged to draw from the specified codes for interoperability purposes but are not required to do so to be considered conformant.
	BindingStrengthPreferred BindingStrength = "preferred"
	// Instances are not expected or even encouraged to draw from the specified value set.  The value set merely provides examples of the types of concepts intended to be included.
	BindingStrengthExample BindingStrength = "example"
)

func (code BindingStrength) String() string {
	return code.Code()
}

func (code BindingStrength) Code() string {
	return string(code)
}

func (code BindingStrength) Display() string {
	switch code {
	case BindingStrengthRequired:
		return "Required"
	case BindingStrengthExtensible:
		return "Extensible"
	case BindingStrengthPreferred:
		return "Preferred"
	case BindingStrengthExample:
		return "Example"
	}
	return "<unknown>"
}

func (code BindingStrength) Definition() string {
	switch code {
	case BindingStrengthRequired:
		return "To be conformant, the concept in this element SHALL be from the specified value set."
	case BindingStrengthExtensible:
		return "To be conformant, the concept in this element SHALL be from the specified value set if any of the codes within the value set can apply to the concept being communicated.  If the value set does not cover the concept (based on human review), alternate codings (or, data type allowing, text) may be included instead."
	case BindingStrengthPreferred:
		return "Instances are encouraged to draw from the specified codes for interoperability purposes but are not required to do so to be considered conformant."
	case BindingStrengthExample:
		return "Instances are not expected or even encouraged to draw from the specified value set.  The value set merely provides examples of the types of concepts intended to be included."
	}
	return "<unknown>"
}
