package fhires

// The kind of response to a message.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/response-code
type ResponseType string

const (
	// The message was accepted and processed without error.
	ResponseTypeOk ResponseType = "ok"
	// Some internal unexpected error occurred - wait and try again. Note - this is usually used for things like database unavailable, which may be expected to resolve, though human intervention may be required.
	ResponseTypeTransientError ResponseType = "transient-error"
	// The message was rejected because of a problem with the content. There is no point in re-sending without change. The response narrative SHALL describe the issue.
	ResponseTypeFatalError ResponseType = "fatal-error"
)

func (code ResponseType) String() string {
	return code.Code()
}

func (code ResponseType) Code() string {
	return string(code)
}

func (code ResponseType) Display() string {
	switch code {
	case ResponseTypeOk:
		return "OK"
	case ResponseTypeTransientError:
		return "Transient Error"
	case ResponseTypeFatalError:
		return "Fatal Error"
	}
	return "<unknown>"
}

func (code ResponseType) Definition() string {
	switch code {
	case ResponseTypeOk:
		return "The message was accepted and processed without error."
	case ResponseTypeTransientError:
		return "Some internal unexpected error occurred - wait and try again. Note - this is usually used for things like database unavailable, which may be expected to resolve, though human intervention may be required."
	case ResponseTypeFatalError:
		return "The message was rejected because of a problem with the content. There is no point in re-sending without change. The response narrative SHALL describe the issue."
	}
	return "<unknown>"
}
