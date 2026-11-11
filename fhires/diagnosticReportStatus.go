package fhires

// The status of the diagnostic report.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/diagnostic-report-status
type DiagnosticReportStatus string

const (
	// The existence of the report is registered, but there is nothing yet available.
	DiagnosticReportStatusRegistered DiagnosticReportStatus = "registered"
	// This is a partial (e.g. initial, interim or preliminary) report: data in the report may be incomplete or unverified.
	DiagnosticReportStatusPartial DiagnosticReportStatus = "partial"
	// Verified early results are available, but not all results are final.
	DiagnosticReportStatusPreliminary DiagnosticReportStatus = "preliminary"
	// Prior to being final, the report has been modified.  This includes any change in the results, diagnosis, narrative text, or other content of a non-finalized (e.g., preliminary) report that has been issued.
	DiagnosticReportStatusModified DiagnosticReportStatus = "modified"
	// The report is complete and verified by an authorized person.
	DiagnosticReportStatusFinal DiagnosticReportStatus = "final"
	// Subsequent to being final, the report has been modified.  This includes any change in the results, diagnosis, narrative text, or other content of a report that has been issued.
	DiagnosticReportStatusAmended DiagnosticReportStatus = "amended"
	// Subsequent to being final, the report has been modified to correct an error in the report or referenced results.
	DiagnosticReportStatusCorrected DiagnosticReportStatus = "corrected"
	// Subsequent to being final, the report has been modified by adding new content. The existing content is unchanged.
	DiagnosticReportStatusAppended DiagnosticReportStatus = "appended"
	// The report is unavailable because the measurement was not started or not completed (also sometimes called "aborted").
	DiagnosticReportStatusCancelled DiagnosticReportStatus = "cancelled"
	// The report has been withdrawn following a previous final release.  This electronic record should never have existed, though it is possible that real-world decisions were based on it. (If real-world activity has occurred, the status should be "cancelled" rather than "entered-in-error".).
	DiagnosticReportStatusEnteredInError DiagnosticReportStatus = "entered-in-error"
	// The authoring/source system does not know which of the status values currently applies for this observation. Note: This concept is not to be used for "other" - one of the listed statuses is presumed to apply, but the authoring/source system does not know which.
	DiagnosticReportStatusUnknown DiagnosticReportStatus = "unknown"
)

func (code DiagnosticReportStatus) String() string {
	return code.Code()
}

func (code DiagnosticReportStatus) Code() string {
	return string(code)
}

func (code DiagnosticReportStatus) Display() string {
	switch code {
	case DiagnosticReportStatusRegistered:
		return "Registered"
	case DiagnosticReportStatusPartial:
		return "Partial"
	case DiagnosticReportStatusPreliminary:
		return "Preliminary"
	case DiagnosticReportStatusModified:
		return "Modified"
	case DiagnosticReportStatusFinal:
		return "Final"
	case DiagnosticReportStatusAmended:
		return "Amended"
	case DiagnosticReportStatusCorrected:
		return "Corrected"
	case DiagnosticReportStatusAppended:
		return "Appended"
	case DiagnosticReportStatusCancelled:
		return "Cancelled"
	case DiagnosticReportStatusEnteredInError:
		return "Entered in Error"
	case DiagnosticReportStatusUnknown:
		return "Unknown"
	}
	return "<unknown>"
}

func (code DiagnosticReportStatus) Definition() string {
	switch code {
	case DiagnosticReportStatusRegistered:
		return "The existence of the report is registered, but there is nothing yet available."
	case DiagnosticReportStatusPartial:
		return "This is a partial (e.g. initial, interim or preliminary) report: data in the report may be incomplete or unverified."
	case DiagnosticReportStatusPreliminary:
		return "Verified early results are available, but not all results are final."
	case DiagnosticReportStatusModified:
		return "Prior to being final, the report has been modified.  This includes any change in the results, diagnosis, narrative text, or other content of a non-finalized (e.g., preliminary) report that has been issued."
	case DiagnosticReportStatusFinal:
		return "The report is complete and verified by an authorized person."
	case DiagnosticReportStatusAmended:
		return "Subsequent to being final, the report has been modified.  This includes any change in the results, diagnosis, narrative text, or other content of a report that has been issued."
	case DiagnosticReportStatusCorrected:
		return "Subsequent to being final, the report has been modified to correct an error in the report or referenced results."
	case DiagnosticReportStatusAppended:
		return "Subsequent to being final, the report has been modified by adding new content. The existing content is unchanged."
	case DiagnosticReportStatusCancelled:
		return "The report is unavailable because the measurement was not started or not completed (also sometimes called \"aborted\")."
	case DiagnosticReportStatusEnteredInError:
		return "The report has been withdrawn following a previous final release.  This electronic record should never have existed, though it is possible that real-world decisions were based on it. (If real-world activity has occurred, the status should be \"cancelled\" rather than \"entered-in-error\".)."
	case DiagnosticReportStatusUnknown:
		return "The authoring/source system does not know which of the status values currently applies for this observation. Note: This concept is not to be used for \"other\" - one of the listed statuses is presumed to apply, but the authoring/source system does not know which."
	}
	return "<unknown>"
}
