package fhires

// The type of a property value.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/concept-property-type
type PropertyType string

const (
	// The property value is a code that identifies a concept defined in the code system.
	PropertyTypeCode PropertyType = "code"
	// The property  value is a code defined in an external code system. This may be used for translations, but is not the intent.
	PropertyTypeCoding PropertyType = "Coding"
	// The property value is a string.
	PropertyTypeString PropertyType = "string"
	// The property value is an integer (often used to assign ranking values to concepts for supporting score assessments).
	PropertyTypeInteger PropertyType = "integer"
	// The property value is a boolean true | false.
	PropertyTypeBoolean PropertyType = "boolean"
	// The property is a date or a date + time.
	PropertyTypeDateTime PropertyType = "dateTime"
	// The property value is a decimal number.
	PropertyTypeDecimal PropertyType = "decimal"
)

func (code PropertyType) String() string {
	return code.Code()
}

func (code PropertyType) Code() string {
	return string(code)
}

func (code PropertyType) Display() string {
	switch code {
	case PropertyTypeCode:
		return "code (internal reference)"
	case PropertyTypeCoding:
		return "Coding (external reference)"
	case PropertyTypeString:
		return "string"
	case PropertyTypeInteger:
		return "integer"
	case PropertyTypeBoolean:
		return "boolean"
	case PropertyTypeDateTime:
		return "dateTime"
	case PropertyTypeDecimal:
		return "decimal"
	}
	return "<unknown>"
}

func (code PropertyType) Definition() string {
	switch code {
	case PropertyTypeCode:
		return "The property value is a code that identifies a concept defined in the code system."
	case PropertyTypeCoding:
		return "The property  value is a code defined in an external code system. This may be used for translations, but is not the intent."
	case PropertyTypeString:
		return "The property value is a string."
	case PropertyTypeInteger:
		return "The property value is an integer (often used to assign ranking values to concepts for supporting score assessments)."
	case PropertyTypeBoolean:
		return "The property value is a boolean true | false."
	case PropertyTypeDateTime:
		return "The property is a date or a date + time."
	case PropertyTypeDecimal:
		return "The property value is a decimal number."
	}
	return "<unknown>"
}
