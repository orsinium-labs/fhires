package fhires

// The status of the ImagingSelection.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/imagingselection-status
type ImagingSelectionStatus string

const (
	// The selected resources are available..
	ImagingSelectionStatusAvailable ImagingSelectionStatus = "available"
	// The imaging selection has been withdrawn following a release.  This electronic record should never have existed, though it is possible that real-world decisions were based on it. (If real-world activity has occurred, the status should be "cancelled" rather than "entered-in-error".).
	ImagingSelectionStatusEnteredInError ImagingSelectionStatus = "entered-in-error"
	// The system does not know which of the status values currently applies for this request. Note: This concept is not to be used for "other" - one of the listed statuses is presumed to apply, it's just not known which one.
	ImagingSelectionStatusUnknown ImagingSelectionStatus = "unknown"
)

func (code ImagingSelectionStatus) String() string {
	return code.Code()
}

func (code ImagingSelectionStatus) Code() string {
	return string(code)
}

func (code ImagingSelectionStatus) Display() string {
	switch code {
	case ImagingSelectionStatusAvailable:
		return "Available"
	case ImagingSelectionStatusEnteredInError:
		return "Entered in Error"
	case ImagingSelectionStatusUnknown:
		return "Unknown"
	}
	return "<unknown>"
}

func (code ImagingSelectionStatus) Definition() string {
	switch code {
	case ImagingSelectionStatusAvailable:
		return "The selected resources are available.."
	case ImagingSelectionStatusEnteredInError:
		return "The imaging selection has been withdrawn following a release.  This electronic record should never have existed, though it is possible that real-world decisions were based on it. (If real-world activity has occurred, the status should be \"cancelled\" rather than \"entered-in-error\".)."
	case ImagingSelectionStatusUnknown:
		return "The system does not know which of the status values currently applies for this request. Note: This concept is not to be used for \"other\" - one of the listed statuses is presumed to apply, it's just not known which one."
	}
	return "<unknown>"
}
