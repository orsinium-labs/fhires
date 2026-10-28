package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// The type of a ConceptMap mapping attribute value.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/conceptmap-attribute-type
type ConceptMapAttributeType int

const (
	ConceptMapAttributeTypeCode ConceptMapAttributeType = iota
	// The attribute value is a code defined in a code system.
	ConceptMapAttributeTypeCoding
	// The attribute value is a string.
	ConceptMapAttributeTypeString
	// The attribute value is a boolean true | false.
	ConceptMapAttributeTypeBoolean
	// The attribute is a Quantity (may represent an integer or a decimal with no units).
	ConceptMapAttributeTypeQuantity
)

func (code ConceptMapAttributeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *ConceptMapAttributeType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "code":
		*code = ConceptMapAttributeTypeCode
	case "Coding":
		*code = ConceptMapAttributeTypeCoding
	case "string":
		*code = ConceptMapAttributeTypeString
	case "boolean":
		*code = ConceptMapAttributeTypeBoolean
	case "Quantity":
		*code = ConceptMapAttributeTypeQuantity
	default:
		return fmt.Errorf("unknown ConceptMapAttributeType code `%s`", s)
	}
	return nil
}

func (code ConceptMapAttributeType) String() string {
	return code.Code()
}

func (code ConceptMapAttributeType) Code() string {
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
