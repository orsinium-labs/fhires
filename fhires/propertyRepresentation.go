package fhires

// How a property is represented when serialized.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/property-representation
type PropertyRepresentation string

const (
	// In XML, this property is represented as an attribute not an element.
	PropertyRepresentationXmlAttr PropertyRepresentation = "xmlAttr"
	// This element is represented using the XML text attribute (primitives only).
	PropertyRepresentationXmlText PropertyRepresentation = "xmlText"
	// The type of this element is indicated using xsi:type.
	PropertyRepresentationTypeAttr PropertyRepresentation = "typeAttr"
	// Use CDA narrative instead of XHTML.
	PropertyRepresentationCdaText PropertyRepresentation = "cdaText"
	// The property is represented using XHTML.
	PropertyRepresentationXhtml PropertyRepresentation = "xhtml"
)

func (code PropertyRepresentation) String() string {
	return code.Code()
}

func (code PropertyRepresentation) Code() string {
	return string(code)
}

func (code PropertyRepresentation) Display() string {
	switch code {
	case PropertyRepresentationXmlAttr:
		return "XML Attribute"
	case PropertyRepresentationXmlText:
		return "XML Text"
	case PropertyRepresentationTypeAttr:
		return "Type Attribute"
	case PropertyRepresentationCdaText:
		return "CDA Text Format"
	case PropertyRepresentationXhtml:
		return "XHTML"
	}
	return "<unknown>"
}

func (code PropertyRepresentation) Definition() string {
	switch code {
	case PropertyRepresentationXmlAttr:
		return "In XML, this property is represented as an attribute not an element."
	case PropertyRepresentationXmlText:
		return "This element is represented using the XML text attribute (primitives only)."
	case PropertyRepresentationTypeAttr:
		return "The type of this element is indicated using xsi:type."
	case PropertyRepresentationCdaText:
		return "Use CDA narrative instead of XHTML."
	case PropertyRepresentationXhtml:
		return "The property is represented using XHTML."
	}
	return "<unknown>"
}
