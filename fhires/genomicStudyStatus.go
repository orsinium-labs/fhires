package fhires

// The status of the GenomicStudy.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/genomicstudy-status
type GenomicStudyStatus string

const (
	// The existence of the genomic study is registered, but there is nothing yet available.
	GenomicStudyStatusRegistered GenomicStudyStatus = "registered"
	// At least one instance has been associated with this genomic study.
	GenomicStudyStatusAvailable GenomicStudyStatus = "available"
	// The genomic study is unavailable because the genomic study was not started or not completed (also sometimes called "aborted").
	GenomicStudyStatusCancelled GenomicStudyStatus = "cancelled"
	// The genomic study has been withdrawn following a previous final release.  This electronic record should never have existed, though it is possible that real-world decisions were based on it. (If real-world activity has occurred, the status should be "cancelled" rather than "entered-in-error".).
	GenomicStudyStatusEnteredInError GenomicStudyStatus = "entered-in-error"
	// The system does not know which of the status values currently applies for this request. Note: This concept is not to be used for "other" - one of the listed statuses is presumed to apply, it's just not known which one.
	GenomicStudyStatusUnknown GenomicStudyStatus = "unknown"
)

func (code GenomicStudyStatus) String() string {
	return code.Code()
}

func (code GenomicStudyStatus) Code() string {
	return string(code)
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
