package fhires

// The type of a ConceptMap mapping attribute value.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/conceptmap-attribute-type
type ConceptMapAttributeType string

const (
	// The attribute value is a code defined in the code system in context.
	ConceptMapAttributeTypeCode ConceptMapAttributeType = "code"
	// The attribute value is a code defined in a code system.
	ConceptMapAttributeTypeCoding ConceptMapAttributeType = "Coding"
	// The attribute value is a string.
	ConceptMapAttributeTypeString ConceptMapAttributeType = "string"
	// The attribute value is a boolean true | false.
	ConceptMapAttributeTypeBoolean ConceptMapAttributeType = "boolean"
	// The attribute is a Quantity (may represent an integer or a decimal with no units).
	ConceptMapAttributeTypeQuantity ConceptMapAttributeType = "Quantity"
)

func (code ConceptMapAttributeType) String() string {
	return code.Code()
}

func (code ConceptMapAttributeType) Code() string {
	return string(code)
}

func (code ConceptMapAttributeType) Display() string {
	switch code {
	case ConceptMapAttributeTypeCode:
		return "code"
	case ConceptMapAttributeTypeCoding:
		return "Coding"
	case ConceptMapAttributeTypeString:
		return "string"
	case ConceptMapAttributeTypeBoolean:
		return "boolean"
	case ConceptMapAttributeTypeQuantity:
		return "Quantity"
	}
	return "<unknown>"
}

func (code ConceptMapAttributeType) Definition() string {
	switch code {
	case ConceptMapAttributeTypeCode:
		return "The attribute value is a code defined in the code system in context."
	case ConceptMapAttributeTypeCoding:
		return "The attribute value is a code defined in a code system."
	case ConceptMapAttributeTypeString:
		return "The attribute value is a string."
	case ConceptMapAttributeTypeBoolean:
		return "The attribute value is a boolean true | false."
	case ConceptMapAttributeTypeQuantity:
		return "The attribute is a Quantity (may represent an integer or a decimal with no units)."
	}
	return "<unknown>"
}
