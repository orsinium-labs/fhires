package fhires

// The status of the ImagingStudy.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/imagingstudy-status
type ImagingStudyStatus string

const (
	// The existence of the imaging study is registered, but there is nothing yet available.
	ImagingStudyStatusRegistered ImagingStudyStatus = "registered"
	// At least one instance has been associated with this imaging study.
	ImagingStudyStatusAvailable ImagingStudyStatus = "available"
	// The imaging study is unavailable because the imaging study was not started or not completed (also sometimes called "aborted").
	ImagingStudyStatusCancelled ImagingStudyStatus = "cancelled"
	// The imaging study has been withdrawn following a previous final release.  This electronic record should never have existed, though it is possible that real-world decisions were based on it. (If real-world activity has occurred, the status should be "cancelled" rather than "entered-in-error".).
	ImagingStudyStatusEnteredInError ImagingStudyStatus = "entered-in-error"
	// The system does not know which of the status values currently applies for this request. Note: This concept is not to be used for "other" - one of the listed statuses is presumed to apply, it's just not known which one.
	ImagingStudyStatusUnknown ImagingStudyStatus = "unknown"
)

func (code ImagingStudyStatus) String() string {
	return code.Code()
}

func (code ImagingStudyStatus) Code() string {
	return string(code)
}

func (code ImagingStudyStatus) Display() string {
	switch code {
	case ImagingStudyStatusRegistered:
		return "Registered"
	case ImagingStudyStatusAvailable:
		return "Available"
	case ImagingStudyStatusCancelled:
		return "Cancelled"
	case ImagingStudyStatusEnteredInError:
		return "Entered in Error"
	case ImagingStudyStatusUnknown:
		return "Unknown"
	}
	return "<unknown>"
}

func (code ImagingStudyStatus) Definition() string {
	switch code {
	case ImagingStudyStatusRegistered:
		return "The existence of the imaging study is registered, but there is nothing yet available."
	case ImagingStudyStatusAvailable:
		return "At least one instance has been associated with this imaging study."
	case ImagingStudyStatusCancelled:
		return "The imaging study is unavailable because the imaging study was not started or not completed (also sometimes called \"aborted\")."
	case ImagingStudyStatusEnteredInError:
		return "The imaging study has been withdrawn following a previous final release.  This electronic record should never have existed, though it is possible that real-world decisions were based on it. (If real-world activity has occurred, the status should be \"cancelled\" rather than \"entered-in-error\".)."
	case ImagingStudyStatusUnknown:
		return "The system does not know which of the status values currently applies for this request. Note: This concept is not to be used for \"other\" - one of the listed statuses is presumed to apply, it's just not known which one."
	}
	return "<unknown>"
}
