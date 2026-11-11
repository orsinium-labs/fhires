package fhires

// Defines which action to take if there is no match in the group.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/conceptmap-unmapped-mode
type ConceptMapGroupUnmappedMode string

const (
	// Use the code as provided in the $translate request in one of the following input parameters: sourceCode, sourceCoding, sourceCodeableConcept.
	ConceptMapGroupUnmappedModeUseSourceCode ConceptMapGroupUnmappedMode = "use-source-code"
	// Use the code(s) explicitly provided in the group.unmapped 'code' or 'valueSet' element.
	ConceptMapGroupUnmappedModeFixed ConceptMapGroupUnmappedMode = "fixed"
	// Use the map identified by the canonical URL in the url element.
	ConceptMapGroupUnmappedModeOtherMap ConceptMapGroupUnmappedMode = "other-map"
)

func (code ConceptMapGroupUnmappedMode) String() string {
	return code.Code()
}

func (code ConceptMapGroupUnmappedMode) Code() string {
	return string(code)
}

func (code ConceptMapGroupUnmappedMode) Display() string {
	switch code {
	case ConceptMapGroupUnmappedModeUseSourceCode:
		return "Use Provided Source Code"
	case ConceptMapGroupUnmappedModeFixed:
		return "Fixed Code"
	case ConceptMapGroupUnmappedModeOtherMap:
		return "Other Map"
	}
	return "<unknown>"
}

func (code ConceptMapGroupUnmappedMode) Definition() string {
	switch code {
	case ConceptMapGroupUnmappedModeUseSourceCode:
		return "Use the code as provided in the $translate request in one of the following input parameters: sourceCode, sourceCoding, sourceCodeableConcept."
	case ConceptMapGroupUnmappedModeFixed:
		return "Use the code(s) explicitly provided in the group.unmapped 'code' or 'valueSet' element."
	case ConceptMapGroupUnmappedModeOtherMap:
		return "Use the map identified by the canonical URL in the url element."
	}
	return "<unknown>"
}
