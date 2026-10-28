package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// The type of actor - system or human.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/examplescenario-actor-type
type ExampleScenarioActorType int

const (
	ExampleScenarioActorTypePerson ExampleScenarioActorType = iota
	// A software application or other system
	ExampleScenarioActorTypeSystem
)

func (code ExampleScenarioActorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *ExampleScenarioActorType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "person":
		*code = ExampleScenarioActorTypePerson
	case "system":
		*code = ExampleScenarioActorTypeSystem
	default:
		return fmt.Errorf("unknown ExampleScenarioActorType code `%s`", s)
	}
	return nil
}

func (code ExampleScenarioActorType) String() string {
	return code.Code()
}

func (code ExampleScenarioActorType) Code() string {
	switch code {
	case ExampleScenarioActorTypePerson:
		return "person"
	case ExampleScenarioActorTypeSystem:
		return "system"
	}
	return "<unknown>"
}

func (code ExampleScenarioActorType) Display() string {
	switch code {
	case ExampleScenarioActorTypePerson:
		return "Person"
	case ExampleScenarioActorTypeSystem:
		return "System"
	}
	return "<unknown>"
}

func (code ExampleScenarioActorType) Definition() string {
	switch code {
	case ExampleScenarioActorTypePerson:
		return "A human actor"
	case ExampleScenarioActorTypeSystem:
		return "A software application or other system"
	}
	return "<unknown>"
}
