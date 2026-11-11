package fhires

// The mode of a message capability statement.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/event-capability-mode
type EventCapabilityMode string

const (
	// The application sends requests and receives responses.
	EventCapabilityModeSender EventCapabilityMode = "sender"
	// The application receives requests and sends responses.
	EventCapabilityModeReceiver EventCapabilityMode = "receiver"
)

func (code EventCapabilityMode) String() string {
	return code.Code()
}

func (code EventCapabilityMode) Code() string {
	return string(code)
}

func (code EventCapabilityMode) Display() string {
	switch code {
	case EventCapabilityModeSender:
		return "Sender"
	case EventCapabilityModeReceiver:
		return "Receiver"
	}
	return "<unknown>"
}

func (code EventCapabilityMode) Definition() string {
	switch code {
	case EventCapabilityModeSender:
		return "The application sends requests and receives responses."
	case EventCapabilityModeReceiver:
		return "The application receives requests and sends responses."
	}
	return "<unknown>"
}
