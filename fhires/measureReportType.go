package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// The type of the measure report.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/measure-report-type
type MeasureReportType int

const (
	MeasureReportTypeIndividual MeasureReportType = iota
	// A subject list report that includes a listing of subjects that satisfied each population criteria in the measure.
	MeasureReportTypeSubjectList
	// A summary report that returns the number of members in each population criteria for the measure.
	MeasureReportTypeSummary
	// A data exchange report that contains data-of-interest for the measure (i.e. data that is needed to calculate the measure)
	MeasureReportTypeDataExchange
)

func (code MeasureReportType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *MeasureReportType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "individual":
		*code = MeasureReportTypeIndividual
	case "subject-list":
		*code = MeasureReportTypeSubjectList
	case "summary":
		*code = MeasureReportTypeSummary
	case "data-exchange":
		*code = MeasureReportTypeDataExchange
	default:
		return fmt.Errorf("unknown MeasureReportType code `%s`", s)
	}
	return nil
}

func (code MeasureReportType) String() string {
	return code.Code()
}

func (code MeasureReportType) Code() string {
	switch code {
	case MeasureReportTypeIndividual:
		return "individual"
	case MeasureReportTypeSubjectList:
		return "subject-list"
	case MeasureReportTypeSummary:
		return "summary"
	case MeasureReportTypeDataExchange:
		return "data-exchange"
	}
	return "<unknown>"
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
