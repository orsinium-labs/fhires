package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// The status of the GenomicStudy.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/genomicstudy-status
type GenomicStudyStatus int

const (
	GenomicStudyStatusRegistered GenomicStudyStatus = iota
	// At least one instance has been associated with this genomic study.
	GenomicStudyStatusAvailable
	// The genomic study is unavailable because the genomic study was not started or not completed (also sometimes called "aborted").
	GenomicStudyStatusCancelled
	// The genomic study has been withdrawn following a previous final release.  This electronic record should never have existed, though it is possible that real-world decisions were based on it. (If real-world activity has occurred, the status should be "cancelled" rather than "entered-in-error".).
	GenomicStudyStatusEnteredInError
	// The system does not know which of the status values currently applies for this request. Note: This concept is not to be used for "other" - one of the listed statuses is presumed to apply, it's just not known which one.
	GenomicStudyStatusUnknown
)

func (code GenomicStudyStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *GenomicStudyStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "registered":
		*code = GenomicStudyStatusRegistered
	case "available":
		*code = GenomicStudyStatusAvailable
	case "cancelled":
		*code = GenomicStudyStatusCancelled
	case "entered-in-error":
		*code = GenomicStudyStatusEnteredInError
	case "unknown":
		*code = GenomicStudyStatusUnknown
	default:
		return fmt.Errorf("unknown GenomicStudyStatus code `%s`", s)
	}
	return nil
}

func (code GenomicStudyStatus) String() string {
	return code.Code()
}

func (code GenomicStudyStatus) Code() string {
	switch code {
	case GenomicStudyStatusRegistered:
		return "registered"
	case GenomicStudyStatusAvailable:
		return "available"
	case GenomicStudyStatusCancelled:
		return "cancelled"
	case GenomicStudyStatusEnteredInError:
		return "entered-in-error"
	case GenomicStudyStatusUnknown:
		return "unknown"
	}
	return "<unknown>"
}

func (code GenomicStudyStatus) Display() string {
	switch code {
	case GenomicStudyStatusRegistered:
		return "Registered"
	case GenomicStudyStatusAvailable:
		return "Available"
	case GenomicStudyStatusCancelled:
		return "Cancelled"
	case GenomicStudyStatusEnteredInError:
		return "Entered in Error"
	case GenomicStudyStatusUnknown:
		return "Unknown"
	}
	return "<unknown>"
}

func (code GenomicStudyStatus) Definition() string {
	switch code {
	case GenomicStudyStatusRegistered:
		return "The existence of the genomic study is registered, but there is nothing yet available."
	case GenomicStudyStatusAvailable:
		return "At least one instance has been associated with this genomic study."
	case GenomicStudyStatusCancelled:
		return "The genomic study is unavailable because the genomic study was not started or not completed (also sometimes called \"aborted\")."
	case GenomicStudyStatusEnteredInError:
		return "The genomic study has been withdrawn following a previous final release.  This electronic record should never have existed, though it is possible that real-world decisions were based on it. (If real-world activity has occurred, the status should be \"cancelled\" rather than \"entered-in-error\".)."
	case GenomicStudyStatusUnknown:
		return "The system does not know which of the status values currently applies for this request. Note: This concept is not to be used for \"other\" - one of the listed statuses is presumed to apply, it's just not known which one."
	}
	return "<unknown>"
}
