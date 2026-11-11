package fhires

// The type of actor - system or human.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/examplescenario-actor-type
type ExampleScenarioActorType string

const (
	// A human actor
	ExampleScenarioActorTypePerson ExampleScenarioActorType = "person"
	// A software application or other system
	ExampleScenarioActorTypeSystem ExampleScenarioActorType = "system"
)

func (code ExampleScenarioActorType) String() string {
	return code.Code()
}

func (code ExampleScenarioActorType) Code() string {
	return string(code)
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
