package fhires

// Defines the types of relationships between actions.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/action-relationship-type
type ActionRelationshipType string

const (
	// The action must be performed before the related action.
	ActionRelationshipTypeBefore ActionRelationshipType = "before"
	// The action must be performed before the start of the related action.
	ActionRelationshipTypeBeforeStart ActionRelationshipType = "before-start"
	// The action must be performed before the end of the related action.
	ActionRelationshipTypeBeforeEnd ActionRelationshipType = "before-end"
	// The action must be performed concurrent with the related action.
	ActionRelationshipTypeConcurrent ActionRelationshipType = "concurrent"
	// The action must be performed concurrent with the start of the related action.
	ActionRelationshipTypeConcurrentWithStart ActionRelationshipType = "concurrent-with-start"
	// The action must be performed concurrent with the end of the related action.
	ActionRelationshipTypeConcurrentWithEnd ActionRelationshipType = "concurrent-with-end"
	// The action must be performed after the related action.
	ActionRelationshipTypeAfter ActionRelationshipType = "after"
	// The action must be performed after the start of the related action.
	ActionRelationshipTypeAfterStart ActionRelationshipType = "after-start"
	// The action must be performed after the end of the related action.
	ActionRelationshipTypeAfterEnd ActionRelationshipType = "after-end"
)

func (code ActionRelationshipType) String() string {
	return code.Code()
}

func (code ActionRelationshipType) Code() string {
	return string(code)
}

func (code ActionRelationshipType) Display() string {
	switch code {
	case ActionRelationshipTypeBefore:
		return "Before"
	case ActionRelationshipTypeBeforeStart:
		return "Before Start"
	case ActionRelationshipTypeBeforeEnd:
		return "Before End"
	case ActionRelationshipTypeConcurrent:
		return "Concurrent"
	case ActionRelationshipTypeConcurrentWithStart:
		return "Concurrent With Start"
	case ActionRelationshipTypeConcurrentWithEnd:
		return "Concurrent With End"
	case ActionRelationshipTypeAfter:
		return "After"
	case ActionRelationshipTypeAfterStart:
		return "After Start"
	case ActionRelationshipTypeAfterEnd:
		return "After End"
	}
	return "<unknown>"
}

func (code ActionRelationshipType) Definition() string {
	switch code {
	case ActionRelationshipTypeBefore:
		return "The action must be performed before the related action."
	case ActionRelationshipTypeBeforeStart:
		return "The action must be performed before the start of the related action."
	case ActionRelationshipTypeBeforeEnd:
		return "The action must be performed before the end of the related action."
	case ActionRelationshipTypeConcurrent:
		return "The action must be performed concurrent with the related action."
	case ActionRelationshipTypeConcurrentWithStart:
		return "The action must be performed concurrent with the start of the related action."
	case ActionRelationshipTypeConcurrentWithEnd:
		return "The action must be performed concurrent with the end of the related action."
	case ActionRelationshipTypeAfter:
		return "The action must be performed after the related action."
	case ActionRelationshipTypeAfterStart:
		return "The action must be performed after the start of the related action."
	case ActionRelationshipTypeAfterEnd:
		return "The action must be performed after the end of the related action."
	}
	return "<unknown>"
}
