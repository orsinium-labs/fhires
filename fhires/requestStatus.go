package fhires

// Codes identifying the lifecycle stage of a request.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/request-status
type RequestStatus string

const (
	// The request has been created but is not yet complete or ready for action.
	RequestStatusDraft RequestStatus = "draft"
	// The request is in force and ready to be acted upon.
	RequestStatusActive RequestStatus = "active"
	// The request (and any implicit authorization to act) has been temporarily withdrawn but is expected to resume in the future.
	RequestStatusOnHold RequestStatus = "on-hold"
	// The request (and any implicit authorization to act) has been terminated prior to the known full completion of the intended actions.  No further activity should occur.
	RequestStatusRevoked RequestStatus = "revoked"
	// The activity described by the request has been fully performed.  No further activity will occur.
	RequestStatusCompleted RequestStatus = "completed"
	// This request should never have existed and should be considered 'void'.  (It is possible that real-world decisions were based on it.  If real-world activity has occurred, the status should be "revoked" rather than "entered-in-error".).
	RequestStatusEnteredInError RequestStatus = "entered-in-error"
	// The authoring/source system does not know which of the status values currently applies for this request.  Note: This concept is not to be used for "other" - one of the listed statuses is presumed to apply,  but the authoring/source system does not know which.
	RequestStatusUnknown RequestStatus = "unknown"
)

func (code RequestStatus) String() string {
	return code.Code()
}

func (code RequestStatus) Code() string {
	return string(code)
}

func (code RequestStatus) Display() string {
	switch code {
	case RequestStatusDraft:
		return "Draft"
	case RequestStatusActive:
		return "Active"
	case RequestStatusOnHold:
		return "On Hold"
	case RequestStatusRevoked:
		return "Revoked"
	case RequestStatusCompleted:
		return "Completed"
	case RequestStatusEnteredInError:
		return "Entered in Error"
	case RequestStatusUnknown:
		return "Unknown"
	}
	return "<unknown>"
}

func (code RequestStatus) Definition() string {
	switch code {
	case RequestStatusDraft:
		return "The request has been created but is not yet complete or ready for action."
	case RequestStatusActive:
		return "The request is in force and ready to be acted upon."
	case RequestStatusOnHold:
		return "The request (and any implicit authorization to act) has been temporarily withdrawn but is expected to resume in the future."
	case RequestStatusRevoked:
		return "The request (and any implicit authorization to act) has been terminated prior to the known full completion of the intended actions.  No further activity should occur."
	case RequestStatusCompleted:
		return "The activity described by the request has been fully performed.  No further activity will occur."
	case RequestStatusEnteredInError:
		return "This request should never have existed and should be considered 'void'.  (It is possible that real-world decisions were based on it.  If real-world activity has occurred, the status should be \"revoked\" rather than \"entered-in-error\".)."
	case RequestStatusUnknown:
		return "The authoring/source system does not know which of the status values currently applies for this request.  Note: This concept is not to be used for \"other\" - one of the listed statuses is presumed to apply,  but the authoring/source system does not know which."
	}
	return "<unknown>"
}
