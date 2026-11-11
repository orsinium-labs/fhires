package fhires

// Whether a reference needs to be version specific or version independent, or whether either can be used.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/reference-version-rules
type ReferenceVersionRules string

const (
	// The reference may be either version independent or version specific.
	ReferenceVersionRulesEither ReferenceVersionRules = "either"
	// The reference must be version independent.
	ReferenceVersionRulesIndependent ReferenceVersionRules = "independent"
	// The reference must be version specific.
	ReferenceVersionRulesSpecific ReferenceVersionRules = "specific"
)

func (code ReferenceVersionRules) String() string {
	return code.Code()
}

func (code ReferenceVersionRules) Code() string {
	return string(code)
}

func (code ReferenceVersionRules) Display() string {
	switch code {
	case ReferenceVersionRulesEither:
		return "Either Specific or independent"
	case ReferenceVersionRulesIndependent:
		return "Version independent"
	case ReferenceVersionRulesSpecific:
		return "Version Specific"
	}
	return "<unknown>"
}

func (code ReferenceVersionRules) Definition() string {
	switch code {
	case ReferenceVersionRulesEither:
		return "The reference may be either version independent or version specific."
	case ReferenceVersionRulesIndependent:
		return "The reference must be version independent."
	case ReferenceVersionRulesSpecific:
		return "The reference must be version specific."
	}
	return "<unknown>"
}
