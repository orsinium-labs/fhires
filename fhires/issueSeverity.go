package fhires

// How the issue affects the success of the action.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/issue-severity
type IssueSeverity string

const (
	// The issue caused the action to fail and no further checking could be performed.
	IssueSeverityFatal IssueSeverity = "fatal"
	// The issue is sufficiently important to cause the action to fail.
	IssueSeverityError IssueSeverity = "error"
	// The issue is not important enough to cause the action to fail but may cause it to be performed suboptimally or in a way that is not as desired.
	IssueSeverityWarning IssueSeverity = "warning"
	// The issue has no relation to the degree of success of the action.
	IssueSeverityInformation IssueSeverity = "information"
	// The operation completed successfully.
	IssueSeveritySuccess IssueSeverity = "success"
)

func (code IssueSeverity) String() string {
	return code.Code()
}

func (code IssueSeverity) Code() string {
	return string(code)
}

func (code IssueSeverity) Display() string {
	switch code {
	case IssueSeverityFatal:
		return "Fatal"
	case IssueSeverityError:
		return "Error"
	case IssueSeverityWarning:
		return "Warning"
	case IssueSeverityInformation:
		return "Information"
	case IssueSeveritySuccess:
		return "Operation Successful"
	}
	return "<unknown>"
}

func (code IssueSeverity) Definition() string {
	switch code {
	case IssueSeverityFatal:
		return "The issue caused the action to fail and no further checking could be performed."
	case IssueSeverityError:
		return "The issue is sufficiently important to cause the action to fail."
	case IssueSeverityWarning:
		return "The issue is not important enough to cause the action to fail but may cause it to be performed suboptimally or in a way that is not as desired."
	case IssueSeverityInformation:
		return "The issue has no relation to the degree of success of the action."
	case IssueSeveritySuccess:
		return "The operation completed successfully."
	}
	return "<unknown>"
}
