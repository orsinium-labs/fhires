package fhires

// The type of the measure report.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/measure-report-type
type MeasureReportType string

const (
	// An individual report that provides information on the performance for a given measure with respect to a single subject.
	MeasureReportTypeIndividual MeasureReportType = "individual"
	// A subject list report that includes a listing of subjects that satisfied each population criteria in the measure.
	MeasureReportTypeSubjectList MeasureReportType = "subject-list"
	// A summary report that returns the number of members in each population criteria for the measure.
	MeasureReportTypeSummary MeasureReportType = "summary"
	// A data exchange report that contains data-of-interest for the measure (i.e. data that is needed to calculate the measure)
	MeasureReportTypeDataExchange MeasureReportType = "data-exchange"
)

func (code MeasureReportType) String() string {
	return code.Code()
}

func (code MeasureReportType) Code() string {
	return string(code)
}

func (code MeasureReportType) Display() string {
	switch code {
	case MeasureReportTypeIndividual:
		return "Individual"
	case MeasureReportTypeSubjectList:
		return "Subject List"
	case MeasureReportTypeSummary:
		return "Summary"
	case MeasureReportTypeDataExchange:
		return "Data Exchange"
	}
	return "<unknown>"
}

func (code MeasureReportType) Definition() string {
	switch code {
	case MeasureReportTypeIndividual:
		return "An individual report that provides information on the performance for a given measure with respect to a single subject."
	case MeasureReportTypeSubjectList:
		return "A subject list report that includes a listing of subjects that satisfied each population criteria in the measure."
	case MeasureReportTypeSummary:
		return "A summary report that returns the number of members in each population criteria for the measure."
	case MeasureReportTypeDataExchange:
		return "A data exchange report that contains data-of-interest for the measure (i.e. data that is needed to calculate the measure)"
	}
	return "<unknown>"
}
