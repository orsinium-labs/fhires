package fhires

// The impact of the content of a message.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/message-significance-category
type MessageSignificanceCategory string

const (
	// The message represents/requests a change that should not be processed more than once; e.g., making a booking for an appointment.
	MessageSignificanceCategoryConsequence MessageSignificanceCategory = "consequence"
	// The message represents a response to query for current information. Retrospective processing is wrong and/or wasteful.
	MessageSignificanceCategoryCurrency MessageSignificanceCategory = "currency"
	// The content is not necessarily intended to be current, and it can be reprocessed, though there may be version issues created by processing old notifications.
	MessageSignificanceCategoryNotification MessageSignificanceCategory = "notification"
)

func (code MessageSignificanceCategory) String() string {
	return code.Code()
}

func (code MessageSignificanceCategory) Code() string {
	return string(code)
}

func (code MessageSignificanceCategory) Display() string {
	switch code {
	case MessageSignificanceCategoryConsequence:
		return "Consequence"
	case MessageSignificanceCategoryCurrency:
		return "Currency"
	case MessageSignificanceCategoryNotification:
		return "Notification"
	}
	return "<unknown>"
}

func (code MessageSignificanceCategory) Definition() string {
	switch code {
	case MessageSignificanceCategoryConsequence:
		return "The message represents/requests a change that should not be processed more than once; e.g., making a booking for an appointment."
	case MessageSignificanceCategoryCurrency:
		return "The message represents a response to query for current information. Retrospective processing is wrong and/or wasteful."
	case MessageSignificanceCategoryNotification:
		return "The content is not necessarily intended to be current, and it can be reprocessed, though there may be version issues created by processing old notifications."
	}
	return "<unknown>"
}
