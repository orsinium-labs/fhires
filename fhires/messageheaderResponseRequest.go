package fhires

import (
	"encoding/json"
	"fmt"
	"strings"
)

// MessageheaderResponseRequest is documented here http://hl7.org/fhir/ValueSet/messageheader-response-request
type MessageheaderResponseRequest int

const (
	MessageheaderResponseRequestAlways MessageheaderResponseRequest = iota
	MessageheaderResponseRequestOnError
	MessageheaderResponseRequestNever
	MessageheaderResponseRequestOnSuccess
)

func (code MessageheaderResponseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *MessageheaderResponseRequest) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "always":
		*code = MessageheaderResponseRequestAlways
	case "on-error":
		*code = MessageheaderResponseRequestOnError
	case "never":
		*code = MessageheaderResponseRequestNever
	case "on-success":
		*code = MessageheaderResponseRequestOnSuccess
	default:
		return fmt.Errorf("unknown MessageheaderResponseRequest code `%s`", s)
	}
	return nil
}

func (code MessageheaderResponseRequest) String() string {
	return code.Code()
}

func (code MessageheaderResponseRequest) Code() string {
	switch code {
	case MessageheaderResponseRequestAlways:
		return "always"
	case MessageheaderResponseRequestOnError:
		return "on-error"
	case MessageheaderResponseRequestNever:
		return "never"
	case MessageheaderResponseRequestOnSuccess:
		return "on-success"
	}
	return "<unknown>"
}

func (code MessageheaderResponseRequest) Display() string {
	switch code {
	case MessageheaderResponseRequestAlways:
		return "Always"
	case MessageheaderResponseRequestOnError:
		return "Error/reject conditions only"
	case MessageheaderResponseRequestNever:
		return "Never"
	case MessageheaderResponseRequestOnSuccess:
		return "Successful completion only"
	}
	return "<unknown>"
}

func (code MessageheaderResponseRequest) Definition() string {
	switch code {
	case MessageheaderResponseRequestAlways:
		return "initiator expects a response for this message."
	case MessageheaderResponseRequestOnError:
		return "initiator expects a response only if in error."
	case MessageheaderResponseRequestNever:
		return "initiator does not expect a response."
	case MessageheaderResponseRequestOnSuccess:
		return "initiator expects a response only if successful."
	}
	return "<unknown>"
}
