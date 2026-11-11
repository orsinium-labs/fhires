package fhires

// How an element value is interpreted when discrimination is evaluated.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/discriminator-type
type DiscriminatorType string

const (
	// The slices have different values in the nominated element, as determined by the applicable fixed value, pattern, or required ValueSet binding.
	DiscriminatorTypeValue DiscriminatorType = "value"
	// The slices are differentiated by the presence or absence of the nominated element. There SHALL be no more than two slices. The slices are differentiated by the fact that one must have a max of 0 and the other must have a min of 1 (or more).  The order in which the slices are declared doesn't matter.
	DiscriminatorTypeExists DiscriminatorType = "exists"
	// The slices have different values in the nominated element, as determined by the applicable fixed value, pattern, or required ValueSet binding. This has the same meaning as 'value' and is deprecated.
	DiscriminatorTypePattern DiscriminatorType = "pattern"
	// The slices are differentiated by type of the nominated element.
	DiscriminatorTypeType DiscriminatorType = "type"
	// The slices are differentiated by conformance of the nominated element to a specified profile. Note that if the path specifies .resolve() then the profile is the target profile on the reference. In this case, validation by the possible profiles is required to differentiate the slices.
	DiscriminatorTypeProfile DiscriminatorType = "profile"
	// The slices are differentiated by their index. This is only possible if all but the last slice have min=max cardinality, and the (optional) last slice contains other undifferentiated elements.
	DiscriminatorTypePosition DiscriminatorType = "position"
)

func (code DiscriminatorType) String() string {
	return code.Code()
}

func (code DiscriminatorType) Code() string {
	return string(code)
}

func (code DiscriminatorType) Display() string {
	switch code {
	case DiscriminatorTypeValue:
		return "Value"
	case DiscriminatorTypeExists:
		return "Exists"
	case DiscriminatorTypePattern:
		return "Pattern"
	case DiscriminatorTypeType:
		return "Type"
	case DiscriminatorTypeProfile:
		return "Profile"
	case DiscriminatorTypePosition:
		return "Position"
	}
	return "<unknown>"
}

func (code DiscriminatorType) Definition() string {
	switch code {
	case DiscriminatorTypeValue:
		return "The slices have different values in the nominated element, as determined by the applicable fixed value, pattern, or required ValueSet binding."
	case DiscriminatorTypeExists:
		return "The slices are differentiated by the presence or absence of the nominated element. There SHALL be no more than two slices. The slices are differentiated by the fact that one must have a max of 0 and the other must have a min of 1 (or more).  The order in which the slices are declared doesn't matter."
	case DiscriminatorTypePattern:
		return "The slices have different values in the nominated element, as determined by the applicable fixed value, pattern, or required ValueSet binding. This has the same meaning as 'value' and is deprecated."
	case DiscriminatorTypeType:
		return "The slices are differentiated by type of the nominated element."
	case DiscriminatorTypeProfile:
		return "The slices are differentiated by conformance of the nominated element to a specified profile. Note that if the path specifies .resolve() then the profile is the target profile on the reference. In this case, validation by the possible profiles is required to differentiate the slices."
	case DiscriminatorTypePosition:
		return "The slices are differentiated by their index. This is only possible if all but the last slice have min=max cardinality, and the (optional) last slice contains other undifferentiated elements."
	}
	return "<unknown>"
}
