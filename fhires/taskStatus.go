package fhires

// The current status of the task.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/task-status
type TaskStatus string

const (
	// The task is not yet ready to be acted upon.
	TaskStatusDraft TaskStatus = "draft"
	// The task is ready to be acted upon and action is sought.
	TaskStatusRequested TaskStatus = "requested"
	// A potential performer has claimed ownership of the task and is evaluating whether to perform it.
	TaskStatusReceived TaskStatus = "received"
	// The potential performer has agreed to execute the task but has not yet started work.
	TaskStatusAccepted TaskStatus = "accepted"
	// The potential performer who claimed ownership of the task has decided not to execute it prior to performing any action.
	TaskStatusRejected TaskStatus = "rejected"
	// The task is ready to be performed, but no action has yet been taken.  Used in place of requested/received/accepted/rejected when request assignment and acceptance is a given.
	TaskStatusReady TaskStatus = "ready"
	// The task was not completed.
	TaskStatusCancelled TaskStatus = "cancelled"
	// The task has been started but is not yet complete.
	TaskStatusInProgress TaskStatus = "in-progress"
	// The task has been started but work has been paused.
	TaskStatusOnHold TaskStatus = "on-hold"
	// The task was attempted but could not be completed due to some error.
	TaskStatusFailed TaskStatus = "failed"
	// The task has been completed.
	TaskStatusCompleted TaskStatus = "completed"
	// The task should never have existed and is retained only because of the possibility it may have used.
	TaskStatusEnteredInError TaskStatus = "entered-in-error"
)

func (code TaskStatus) String() string {
	return code.Code()
}

func (code TaskStatus) Code() string {
	return string(code)
}

func (code TaskStatus) Display() string {
	switch code {
	case TaskStatusDraft:
		return "Draft"
	case TaskStatusRequested:
		return "Requested"
	case TaskStatusReceived:
		return "Received"
	case TaskStatusAccepted:
		return "Accepted"
	case TaskStatusRejected:
		return "Rejected"
	case TaskStatusReady:
		return "Ready"
	case TaskStatusCancelled:
		return "Cancelled"
	case TaskStatusInProgress:
		return "In Progress"
	case TaskStatusOnHold:
		return "On Hold"
	case TaskStatusFailed:
		return "Failed"
	case TaskStatusCompleted:
		return "Completed"
	case TaskStatusEnteredInError:
		return "Entered in Error"
	}
	return "<unknown>"
}

func (code TaskStatus) Definition() string {
	switch code {
	case TaskStatusDraft:
		return "The task is not yet ready to be acted upon."
	case TaskStatusRequested:
		return "The task is ready to be acted upon and action is sought."
	case TaskStatusReceived:
		return "A potential performer has claimed ownership of the task and is evaluating whether to perform it."
	case TaskStatusAccepted:
		return "The potential performer has agreed to execute the task but has not yet started work."
	case TaskStatusRejected:
		return "The potential performer who claimed ownership of the task has decided not to execute it prior to performing any action."
	case TaskStatusReady:
		return "The task is ready to be performed, but no action has yet been taken.  Used in place of requested/received/accepted/rejected when request assignment and acceptance is a given."
	case TaskStatusCancelled:
		return "The task was not completed."
	case TaskStatusInProgress:
		return "The task has been started but is not yet complete."
	case TaskStatusOnHold:
		return "The task has been started but work has been paused."
	case TaskStatusFailed:
		return "The task was attempted but could not be completed due to some error."
	case TaskStatusCompleted:
		return "The task has been completed."
	case TaskStatusEnteredInError:
		return "The task should never have existed and is retained only because of the possibility it may have used."
	}
	return "<unknown>"
}
