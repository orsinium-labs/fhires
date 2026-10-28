package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// Defines how a compartment rule is used.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/graph-compartment-use
type GraphCompartmentUse int

const (
	GraphCompartmentUseWhere GraphCompartmentUse = iota
	// This compartment rule is enforced on any relationships that meet the conditions.
	GraphCompartmentUseRequires
)

func (code GraphCompartmentUse) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *GraphCompartmentUse) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "where":
		*code = GraphCompartmentUseWhere
	case "requires":
		*code = GraphCompartmentUseRequires
	default:
		return fmt.Errorf("unknown GraphCompartmentUse code `%s`", s)
	}
	return nil
}

func (code GraphCompartmentUse) String() string {
	return code.Code()
}

func (code GraphCompartmentUse) Code() string {
	switch code {
	case GraphCompartmentUseWhere:
		return "where"
	case GraphCompartmentUseRequires:
		return "requires"
	}
	return "<unknown>"
}

func (code GraphCompartmentUse) Display() string {
	switch code {
	case GraphCompartmentUseWhere:
		return "Where"
	case GraphCompartmentUseRequires:
		return "requires"
	}
	return "<unknown>"
}

func (code GraphCompartmentUse) Definition() string {
	switch code {
	case GraphCompartmentUseWhere:
		return "This compartment rule is a condition for whether the rule applies."
	case GraphCompartmentUseRequires:
		return "This compartment rule is enforced on any relationships that meet the conditions."
	}
	return "<unknown>"
}
