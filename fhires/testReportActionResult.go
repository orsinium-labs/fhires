package fhires

// The results of executing an action.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/report-action-result-codes
type TestReportActionResult string

const (
	// The action was successful.
	TestReportActionResultPass TestReportActionResult = "pass"
	// The action was skipped.
	TestReportActionResultSkip TestReportActionResult = "skip"
	// The action failed.
	TestReportActionResultFail TestReportActionResult = "fail"
	// The action passed but with warnings.
	TestReportActionResultWarning TestReportActionResult = "warning"
	// The action encountered a fatal error and the engine was unable to process.
	TestReportActionResultError TestReportActionResult = "error"
)

func (code TestReportActionResult) String() string {
	return code.Code()
}

func (code TestReportActionResult) Code() string {
	return string(code)
}

func (code TestReportActionResult) Display() string {
	switch code {
	case TestReportActionResultPass:
		return "Pass"
	case TestReportActionResultSkip:
		return "Skip"
	case TestReportActionResultFail:
		return "Fail"
	case TestReportActionResultWarning:
		return "Warning"
	case TestReportActionResultError:
		return "Error"
	}
	return "<unknown>"
}

func (code TestReportActionResult) Definition() string {
	switch code {
	case TestReportActionResultPass:
		return "The action was successful."
	case TestReportActionResultSkip:
		return "The action was skipped."
	case TestReportActionResultFail:
		return "The action failed."
	case TestReportActionResultWarning:
		return "The action passed but with warnings."
	case TestReportActionResultError:
		return "The action encountered a fatal error and the engine was unable to process."
	}
	return "<unknown>"
}
