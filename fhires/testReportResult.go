package fhires

// The reported execution result.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/report-result-codes
type TestReportResult string

const (
	// All test operations successfully passed all asserts.
	TestReportResultPass TestReportResult = "pass"
	// One or more test operations failed one or more asserts.
	TestReportResultFail TestReportResult = "fail"
	// One or more test operations is pending execution completion.
	TestReportResultPending TestReportResult = "pending"
)

func (code TestReportResult) String() string {
	return code.Code()
}

func (code TestReportResult) Code() string {
	return string(code)
}

func (code TestReportResult) Display() string {
	switch code {
	case TestReportResultPass:
		return "Pass"
	case TestReportResultFail:
		return "Fail"
	case TestReportResultPending:
		return "Pending"
	}
	return "<unknown>"
}

func (code TestReportResult) Definition() string {
	switch code {
	case TestReportResultPass:
		return "All test operations successfully passed all asserts."
	case TestReportResultFail:
		return "One or more test operations failed one or more asserts."
	case TestReportResultPending:
		return "One or more test operations is pending execution completion."
	}
	return "<unknown>"
}
