package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// The mode of a message capability statement.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/event-capability-mode
type EventCapabilityMode int

const (
	EventCapabilityModeSender EventCapabilityMode = iota
	// The application receives requests and sends responses.
	EventCapabilityModeReceiver
)

func (code EventCapabilityMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *EventCapabilityMode) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "sender":
		*code = EventCapabilityModeSender
	case "receiver":
		*code = EventCapabilityModeReceiver
	default:
		return fmt.Errorf("unknown EventCapabilityMode code `%s`", s)
	}
	return nil
}

func (code EventCapabilityMode) String() string {
	return code.Code()
}

func (code EventCapabilityMode) Code() string {
	switch code {
	case EventCapabilityModeSender:
		return "sender"
	case EventCapabilityModeReceiver:
		return "receiver"
	}
	return "<unknown>"
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
