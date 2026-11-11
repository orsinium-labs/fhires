package fhires

// Status of the transport
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/transport-status
type TransportStatus string

const (
	// Transport has started but not completed.
	TransportStatusInProgress TransportStatus = "in-progress"
	// Transport has been completed.
	TransportStatusCompleted TransportStatus = "completed"
	// Transport was started but not completed.
	TransportStatusAbandoned TransportStatus = "abandoned"
	// Transport was cancelled before started.
	TransportStatusCancelled TransportStatus = "cancelled"
	// Planned transport that is not yet requested.
	TransportStatusPlanned TransportStatus = "planned"
	// This electronic record should never have existed, though it is possible that real-world decisions were based on it. (If real-world activity has occurred, the status should be "abandoned" rather than "entered-in-error".).
	TransportStatusEnteredInError TransportStatus = "entered-in-error"
)

func (code TransportStatus) String() string {
	return code.Code()
}

func (code TransportStatus) Code() string {
	return string(code)
}

func (code TransportStatus) Display() string {
	switch code {
	case TransportStatusInProgress:
		return "In Progress"
	case TransportStatusCompleted:
		return "Completed"
	case TransportStatusAbandoned:
		return "Abandoned"
	case TransportStatusCancelled:
		return "Cancelled"
	case TransportStatusPlanned:
		return "Planned"
	case TransportStatusEnteredInError:
		return "Entered In Error"
	}
	return "<unknown>"
}

func (code TransportStatus) Definition() string {
	switch code {
	case TransportStatusInProgress:
		return "Transport has started but not completed."
	case TransportStatusCompleted:
		return "Transport has been completed."
	case TransportStatusAbandoned:
		return "Transport was started but not completed."
	case TransportStatusCancelled:
		return "Transport was cancelled before started."
	case TransportStatusPlanned:
		return "Planned transport that is not yet requested."
	case TransportStatusEnteredInError:
		return "This electronic record should never have existed, though it is possible that real-world decisions were based on it. (If real-world activity has occurred, the status should be \"abandoned\" rather than \"entered-in-error\".)."
	}
	return "<unknown>"
}
