package fhires

// HL7-defined table of codes which identify conditions under which acknowledgments are required to be returned in response to a message.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/messageheader-response-request
type MessageheaderResponseRequest string

const (
	// initiator expects a response for this message.
	MessageheaderResponseRequestAlways MessageheaderResponseRequest = "always"
	// initiator expects a response only if in error.
	MessageheaderResponseRequestOnError MessageheaderResponseRequest = "on-error"
	// initiator does not expect a response.
	MessageheaderResponseRequestNever MessageheaderResponseRequest = "never"
	// initiator expects a response only if successful.
	MessageheaderResponseRequestOnSuccess MessageheaderResponseRequest = "on-success"
)

func (code MessageheaderResponseRequest) String() string {
	return code.Code()
}

func (code MessageheaderResponseRequest) Code() string {
	return string(code)
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
