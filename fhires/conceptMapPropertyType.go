package fhires

// The type of a ConceptMap mapping property value.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/conceptmap-property-type
type ConceptMapPropertyType string

const (
	// The property  value is a code defined in an external code system. This may be used for translations, but is not the intent.
	ConceptMapPropertyTypeCoding ConceptMapPropertyType = "Coding"
	// The property value is a string.
	ConceptMapPropertyTypeString ConceptMapPropertyType = "string"
	// The property value is an integer (often used to assign ranking values to concepts for supporting score assessments).
	ConceptMapPropertyTypeInteger ConceptMapPropertyType = "integer"
	// The property value is a boolean true | false.
	ConceptMapPropertyTypeBoolean ConceptMapPropertyType = "boolean"
	// The property is a date or a date + time.
	ConceptMapPropertyTypeDateTime ConceptMapPropertyType = "dateTime"
	// The property value is a decimal number.
	ConceptMapPropertyTypeDecimal ConceptMapPropertyType = "decimal"
	// The property value is a code as defined in the CodeSystem in ConceptMap.property.system.
	ConceptMapPropertyTypeCode ConceptMapPropertyType = "code"
)

func (code ConceptMapPropertyType) String() string {
	return code.Code()
}

func (code ConceptMapPropertyType) Code() string {
	return string(code)
}

func (code ConceptMapPropertyType) Display() string {
	switch code {
	case ConceptMapPropertyTypeCoding:
		return "Coding (external reference)"
	case ConceptMapPropertyTypeString:
		return "string"
	case ConceptMapPropertyTypeInteger:
		return "integer"
	case ConceptMapPropertyTypeBoolean:
		return "boolean"
	case ConceptMapPropertyTypeDateTime:
		return "dateTime"
	case ConceptMapPropertyTypeDecimal:
		return "decimal"
	case ConceptMapPropertyTypeCode:
		return "code"
	}
	return "<unknown>"
}

func (code ConceptMapPropertyType) Definition() string {
	switch code {
	case ConceptMapPropertyTypeCoding:
		return "The property  value is a code defined in an external code system. This may be used for translations, but is not the intent."
	case ConceptMapPropertyTypeString:
		return "The property value is a string."
	case ConceptMapPropertyTypeInteger:
		return "The property value is an integer (often used to assign ranking values to concepts for supporting score assessments)."
	case ConceptMapPropertyTypeBoolean:
		return "The property value is a boolean true | false."
	case ConceptMapPropertyTypeDateTime:
		return "The property is a date or a date + time."
	case ConceptMapPropertyTypeDecimal:
		return "The property value is a decimal number."
	case ConceptMapPropertyTypeCode:
		return "The property value is a code as defined in the CodeSystem in ConceptMap.property.system."
	}
	return "<unknown>"
}
