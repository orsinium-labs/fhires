package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// The current status of the test report.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/report-status-codes
type TestReportStatus int

const (
	TestReportStatusCompleted TestReportStatus = iota
	// A test operations is currently executing.
	TestReportStatusInProgress
	// A test operation is waiting for an external client request.
	TestReportStatusWaiting
	// The test script execution was manually stopped.
	TestReportStatusStopped
	// This test report was entered or created in error.
	TestReportStatusEnteredInError
)

func (code TestReportStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *TestReportStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "completed":
		*code = TestReportStatusCompleted
	case "in-progress":
		*code = TestReportStatusInProgress
	case "waiting":
		*code = TestReportStatusWaiting
	case "stopped":
		*code = TestReportStatusStopped
	case "entered-in-error":
		*code = TestReportStatusEnteredInError
	default:
		return fmt.Errorf("unknown TestReportStatus code `%s`", s)
	}
	return nil
}

func (code TestReportStatus) String() string {
	return code.Code()
}

func (code TestReportStatus) Code() string {
	switch code {
	case TestReportStatusCompleted:
		return "completed"
	case TestReportStatusInProgress:
		return "in-progress"
	case TestReportStatusWaiting:
		return "waiting"
	case TestReportStatusStopped:
		return "stopped"
	case TestReportStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
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
