package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// The status of the ImagingSelection.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/imagingselection-status
type ImagingSelectionStatus int

const (
	ImagingSelectionStatusAvailable ImagingSelectionStatus = iota
	// The imaging selection has been withdrawn following a release.  This electronic record should never have existed, though it is possible that real-world decisions were based on it. (If real-world activity has occurred, the status should be "cancelled" rather than "entered-in-error".).
	ImagingSelectionStatusEnteredInError
	// The system does not know which of the status values currently applies for this request. Note: This concept is not to be used for "other" - one of the listed statuses is presumed to apply, it's just not known which one.
	ImagingSelectionStatusUnknown
)

func (code ImagingSelectionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *ImagingSelectionStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "available":
		*code = ImagingSelectionStatusAvailable
	case "entered-in-error":
		*code = ImagingSelectionStatusEnteredInError
	case "unknown":
		*code = ImagingSelectionStatusUnknown
	default:
		return fmt.Errorf("unknown ImagingSelectionStatus code `%s`", s)
	}
	return nil
}

func (code ImagingSelectionStatus) String() string {
	return code.Code()
}

func (code ImagingSelectionStatus) Code() string {
	switch code {
	case ImagingSelectionStatusAvailable:
		return "available"
	case ImagingSelectionStatusEnteredInError:
		return "entered-in-error"
	case ImagingSelectionStatusUnknown:
		return "unknown"
	}
	return "<unknown>"
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
