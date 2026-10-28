package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// ConceptMapGroupUnmappedMode is documented here http://hl7.org/fhir/ValueSet/conceptmap-unmapped-mode
type ConceptMapGroupUnmappedMode int

const (
	ConceptMapGroupUnmappedModeUseSourceCode ConceptMapGroupUnmappedMode = iota
	ConceptMapGroupUnmappedModeFixed
	ConceptMapGroupUnmappedModeOtherMap
)

func (code ConceptMapGroupUnmappedMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *ConceptMapGroupUnmappedMode) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "use-source-code":
		*code = ConceptMapGroupUnmappedModeUseSourceCode
	case "fixed":
		*code = ConceptMapGroupUnmappedModeFixed
	case "other-map":
		*code = ConceptMapGroupUnmappedModeOtherMap
	default:
		return fmt.Errorf("unknown ConceptMapGroupUnmappedMode code `%s`", s)
	}
	return nil
}

func (code ConceptMapGroupUnmappedMode) String() string {
	return code.Code()
}

func (code ConceptMapGroupUnmappedMode) Code() string {
	switch code {
	case ConceptMapGroupUnmappedModeUseSourceCode:
		return "use-source-code"
	case ConceptMapGroupUnmappedModeFixed:
		return "fixed"
	case ConceptMapGroupUnmappedModeOtherMap:
		return "other-map"
	}
	return "<unknown>"
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
