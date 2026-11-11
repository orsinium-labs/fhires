package fhires

// Identifies the level of importance to be assigned to actioning the request.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/request-priority
type RequestPriority string

const (
	// The request has normal priority.
	RequestPriorityRoutine RequestPriority = "routine"
	// The request should be actioned promptly - higher priority than routine.
	RequestPriorityUrgent RequestPriority = "urgent"
	// The request should be actioned as soon as possible - higher priority than urgent.
	RequestPriorityAsap RequestPriority = "asap"
	// The request should be actioned immediately - highest possible priority.  E.g. an emergency.
	RequestPriorityStat RequestPriority = "stat"
)

func (code RequestPriority) String() string {
	return code.Code()
}

func (code RequestPriority) Code() string {
	return string(code)
}

func (code RequestPriority) Display() string {
	switch code {
	case RequestPriorityRoutine:
		return "Routine"
	case RequestPriorityUrgent:
		return "Urgent"
	case RequestPriorityAsap:
		return "ASAP"
	case RequestPriorityStat:
		return "STAT"
	}
	return "<unknown>"
}

func (code RequestPriority) Definition() string {
	switch code {
	case RequestPriorityRoutine:
		return "The request has normal priority."
	case RequestPriorityUrgent:
		return "The request should be actioned promptly - higher priority than routine."
	case RequestPriorityAsap:
		return "The request should be actioned as soon as possible - higher priority than urgent."
	case RequestPriorityStat:
		return "The request should be actioned immediately - highest possible priority.  E.g. an emergency."
	}
	return "<unknown>"
}
