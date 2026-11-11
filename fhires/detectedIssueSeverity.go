package fhires

// Indicates the potential degree of impact of the identified issue on the patient.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/detectedissue-severity
type DetectedIssueSeverity string

const (
	// Indicates the issue may be life-threatening or has the potential to cause permanent injury.
	DetectedIssueSeverityHigh DetectedIssueSeverity = "high"
	// Indicates the issue may result in noticeable adverse consequences but is unlikely to be life-threatening or cause permanent injury.
	DetectedIssueSeverityModerate DetectedIssueSeverity = "moderate"
	// Indicates the issue may result in some adverse consequences but is unlikely to substantially affect the situation of the subject.
	DetectedIssueSeverityLow DetectedIssueSeverity = "low"
)

func (code DetectedIssueSeverity) String() string {
	return code.Code()
}

func (code DetectedIssueSeverity) Code() string {
	return string(code)
}

func (code DetectedIssueSeverity) Display() string {
	switch code {
	case DetectedIssueSeverityHigh:
		return "High"
	case DetectedIssueSeverityModerate:
		return "Moderate"
	case DetectedIssueSeverityLow:
		return "Low"
	}
	return "<unknown>"
}

func (code DetectedIssueSeverity) Definition() string {
	switch code {
	case DetectedIssueSeverityHigh:
		return "Indicates the issue may be life-threatening or has the potential to cause permanent injury."
	case DetectedIssueSeverityModerate:
		return "Indicates the issue may result in noticeable adverse consequences but is unlikely to be life-threatening or cause permanent injury."
	case DetectedIssueSeverityLow:
		return "Indicates the issue may result in some adverse consequences but is unlikely to substantially affect the situation of the subject."
	}
	return "<unknown>"
}
