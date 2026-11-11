package fhires

// The status of the measure report.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/measure-report-status
type MeasureReportStatus string

const (
	// The report is complete and ready for use.
	MeasureReportStatusComplete MeasureReportStatus = "complete"
	// The report is currently being generated.
	MeasureReportStatusPending MeasureReportStatus = "pending"
	// An error occurred attempting to generate the report.
	MeasureReportStatusError MeasureReportStatus = "error"
)

func (code MeasureReportStatus) String() string {
	return code.Code()
}

func (code MeasureReportStatus) Code() string {
	return string(code)
}

func (code MeasureReportStatus) Display() string {
	switch code {
	case MeasureReportStatusComplete:
		return "Complete"
	case MeasureReportStatusPending:
		return "Pending"
	case MeasureReportStatusError:
		return "Error"
	}
	return "<unknown>"
}

func (code MeasureReportStatus) Definition() string {
	switch code {
	case MeasureReportStatusComplete:
		return "The report is complete and ready for use."
	case MeasureReportStatusPending:
		return "The report is currently being generated."
	case MeasureReportStatusError:
		return "An error occurred attempting to generate the report."
	}
	return "<unknown>"
}
