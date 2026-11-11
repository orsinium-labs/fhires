package fhires

// The current status of the test report.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/report-status-codes
type TestReportStatus string

const (
	// All test operations have completed.
	TestReportStatusCompleted TestReportStatus = "completed"
	// A test operations is currently executing.
	TestReportStatusInProgress TestReportStatus = "in-progress"
	// A test operation is waiting for an external client request.
	TestReportStatusWaiting TestReportStatus = "waiting"
	// The test script execution was manually stopped.
	TestReportStatusStopped TestReportStatus = "stopped"
	// This test report was entered or created in error.
	TestReportStatusEnteredInError TestReportStatus = "entered-in-error"
)

func (code TestReportStatus) String() string {
	return code.Code()
}

func (code TestReportStatus) Code() string {
	return string(code)
}

func (code TestReportStatus) Display() string {
	switch code {
	case TestReportStatusCompleted:
		return "Completed"
	case TestReportStatusInProgress:
		return "In Progress"
	case TestReportStatusWaiting:
		return "Waiting"
	case TestReportStatusStopped:
		return "Stopped"
	case TestReportStatusEnteredInError:
		return "Entered In Error"
	}
	return "<unknown>"
}

func (code TestReportStatus) Definition() string {
	switch code {
	case TestReportStatusCompleted:
		return "All test operations have completed."
	case TestReportStatusInProgress:
		return "A test operations is currently executing."
	case TestReportStatusWaiting:
		return "A test operation is waiting for an external client request."
	case TestReportStatusStopped:
		return "The test script execution was manually stopped."
	case TestReportStatusEnteredInError:
		return "This test report was entered or created in error."
	}
	return "<unknown>"
}
