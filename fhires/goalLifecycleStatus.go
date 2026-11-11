package fhires

// Codes that reflect the current state of a goal and whether the goal is still being targeted.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/goal-status
type GoalLifecycleStatus string

const (
	// A goal is proposed for this patient.
	GoalLifecycleStatusProposed GoalLifecycleStatus = "proposed"
	// A goal is planned for this patient.
	GoalLifecycleStatusPlanned GoalLifecycleStatus = "planned"
	// A proposed goal was accepted or acknowledged.
	GoalLifecycleStatusAccepted GoalLifecycleStatus = "accepted"
	// The goal is being sought actively.
	GoalLifecycleStatusActive GoalLifecycleStatus = "active"
	// The goal remains a long term objective but is no longer being actively pursued for a temporary period of time.
	GoalLifecycleStatusOnHold GoalLifecycleStatus = "on-hold"
	// The goal is no longer being sought.
	GoalLifecycleStatusCompleted GoalLifecycleStatus = "completed"
	// The goal has been abandoned.
	GoalLifecycleStatusCancelled GoalLifecycleStatus = "cancelled"
	// The goal was entered in error and voided.
	GoalLifecycleStatusEnteredInError GoalLifecycleStatus = "entered-in-error"
	// A proposed goal was rejected.
	GoalLifecycleStatusRejected GoalLifecycleStatus = "rejected"
)

func (code GoalLifecycleStatus) String() string {
	return code.Code()
}

func (code GoalLifecycleStatus) Code() string {
	return string(code)
}

func (code GoalLifecycleStatus) Display() string {
	switch code {
	case GoalLifecycleStatusProposed:
		return "Proposed"
	case GoalLifecycleStatusPlanned:
		return "Planned"
	case GoalLifecycleStatusAccepted:
		return "Accepted"
	case GoalLifecycleStatusActive:
		return "Active"
	case GoalLifecycleStatusOnHold:
		return "On Hold"
	case GoalLifecycleStatusCompleted:
		return "Completed"
	case GoalLifecycleStatusCancelled:
		return "Cancelled"
	case GoalLifecycleStatusEnteredInError:
		return "Entered in Error"
	case GoalLifecycleStatusRejected:
		return "Rejected"
	}
	return "<unknown>"
}

func (code GoalLifecycleStatus) Definition() string {
	switch code {
	case GoalLifecycleStatusProposed:
		return "A goal is proposed for this patient."
	case GoalLifecycleStatusPlanned:
		return "A goal is planned for this patient."
	case GoalLifecycleStatusAccepted:
		return "A proposed goal was accepted or acknowledged."
	case GoalLifecycleStatusActive:
		return "The goal is being sought actively."
	case GoalLifecycleStatusOnHold:
		return "The goal remains a long term objective but is no longer being actively pursued for a temporary period of time."
	case GoalLifecycleStatusCompleted:
		return "The goal is no longer being sought."
	case GoalLifecycleStatusCancelled:
		return "The goal has been abandoned."
	case GoalLifecycleStatusEnteredInError:
		return "The goal was entered in error and voided."
	case GoalLifecycleStatusRejected:
		return "A proposed goal was rejected."
	}
	return "<unknown>"
}
