package fhires

// Behavior a server can exhibit when a criteria state does not exist (e.g., state prior to a create or after a delete).
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/subscriptiontopic-cr-behavior
type CriteriaNotExistsBehavior string

const (
	// The requested conditional statement will pass if a matching state does not exist (e.g., previous state during create).
	CriteriaNotExistsBehaviorTestPasses CriteriaNotExistsBehavior = "test-passes"
	// The requested conditional statement will fail if a matching state does not exist (e.g., previous state during create).
	CriteriaNotExistsBehaviorTestFails CriteriaNotExistsBehavior = "test-fails"
)

func (code CriteriaNotExistsBehavior) String() string {
	return code.Code()
}

func (code CriteriaNotExistsBehavior) Code() string {
	return string(code)
}

func (code CriteriaNotExistsBehavior) Display() string {
	switch code {
	case CriteriaNotExistsBehaviorTestPasses:
		return "Test passes"
	case CriteriaNotExistsBehaviorTestFails:
		return "Test fails"
	}
	return "<unknown>"
}

func (code CriteriaNotExistsBehavior) Definition() string {
	switch code {
	case CriteriaNotExistsBehaviorTestPasses:
		return "The requested conditional statement will pass if a matching state does not exist (e.g., previous state during create)."
	case CriteriaNotExistsBehaviorTestFails:
		return "The requested conditional statement will fail if a matching state does not exist (e.g., previous state during create)."
	}
	return "<unknown>"
}
