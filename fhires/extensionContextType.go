package fhires

// How an extension context is interpreted.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/extension-context-type
type ExtensionContextType string

const (
	// The context is all elements that match the FHIRPath query found in the expression.
	ExtensionContextTypeFhirpath ExtensionContextType = "fhirpath"
	// The context is any element that has an ElementDefinition.id that matches that found in the expression. This includes ElementDefinition Ids that have slicing identifiers. The full path for the element is [url]#[elementid]. If there is no #, the Element id is one defined in the base specification.
	ExtensionContextTypeElement ExtensionContextType = "element"
	// The context is a particular extension from a particular StructureDefinition, and the expression is just a uri that identifies the extension.
	ExtensionContextTypeExtension ExtensionContextType = "extension"
)

func (code ExtensionContextType) String() string {
	return code.Code()
}

func (code ExtensionContextType) Code() string {
	return string(code)
}

func (code ExtensionContextType) Display() string {
	switch code {
	case ExtensionContextTypeFhirpath:
		return "FHIRPath"
	case ExtensionContextTypeElement:
		return "Element ID"
	case ExtensionContextTypeExtension:
		return "Extension URL"
	}
	return "<unknown>"
}

func (code ExtensionContextType) Definition() string {
	switch code {
	case ExtensionContextTypeFhirpath:
		return "The context is all elements that match the FHIRPath query found in the expression."
	case ExtensionContextTypeElement:
		return "The context is any element that has an ElementDefinition.id that matches that found in the expression. This includes ElementDefinition Ids that have slicing identifiers. The full path for the element is [url]#[elementid]. If there is no #, the Element id is one defined in the base specification."
	case ExtensionContextTypeExtension:
		return "The context is a particular extension from a particular StructureDefinition, and the expression is just a uri that identifies the extension."
	}
	return "<unknown>"
}
