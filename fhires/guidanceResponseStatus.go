package fhires

// The status of a guidance response.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/guidance-response-status
type GuidanceResponseStatus string

const (
	// The request was processed successfully.
	GuidanceResponseStatusSuccess GuidanceResponseStatus = "success"
	// The request was processed successfully, but more data may result in a more complete evaluation.
	GuidanceResponseStatusDataRequested GuidanceResponseStatus = "data-requested"
	// The request was processed, but more data is required to complete the evaluation.
	GuidanceResponseStatusDataRequired GuidanceResponseStatus = "data-required"
	// The request is currently being processed.
	GuidanceResponseStatusInProgress GuidanceResponseStatus = "in-progress"
	// The request was not processed successfully.
	GuidanceResponseStatusFailure GuidanceResponseStatus = "failure"
	// The response was entered in error.
	GuidanceResponseStatusEnteredInError GuidanceResponseStatus = "entered-in-error"
)

func (code GuidanceResponseStatus) String() string {
	return code.Code()
}

func (code GuidanceResponseStatus) Code() string {
	return string(code)
}

func (code GuidanceResponseStatus) Display() string {
	switch code {
	case GuidanceResponseStatusSuccess:
		return "Success"
	case GuidanceResponseStatusDataRequested:
		return "Data Requested"
	case GuidanceResponseStatusDataRequired:
		return "Data Required"
	case GuidanceResponseStatusInProgress:
		return "In Progress"
	case GuidanceResponseStatusFailure:
		return "Failure"
	case GuidanceResponseStatusEnteredInError:
		return "Entered In Error"
	}
	return "<unknown>"
}

func (code GuidanceResponseStatus) Definition() string {
	switch code {
	case GuidanceResponseStatusSuccess:
		return "The request was processed successfully."
	case GuidanceResponseStatusDataRequested:
		return "The request was processed successfully, but more data may result in a more complete evaluation."
	case GuidanceResponseStatusDataRequired:
		return "The request was processed, but more data is required to complete the evaluation."
	case GuidanceResponseStatusInProgress:
		return "The request is currently being processed."
	case GuidanceResponseStatusFailure:
		return "The request was not processed successfully."
	case GuidanceResponseStatusEnteredInError:
		return "The response was entered in error."
	}
	return "<unknown>"
}
