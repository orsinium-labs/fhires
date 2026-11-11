package fhires

// The type of participant.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/report-participant-type
type TestReportParticipantType string

const (
	// The test execution engine.
	TestReportParticipantTypeTestEngine TestReportParticipantType = "test-engine"
	// A FHIR Client.
	TestReportParticipantTypeClient TestReportParticipantType = "client"
	// A FHIR Server.
	TestReportParticipantTypeServer TestReportParticipantType = "server"
)

func (code TestReportParticipantType) String() string {
	return code.Code()
}

func (code TestReportParticipantType) Code() string {
	return string(code)
}

func (code TestReportParticipantType) Display() string {
	switch code {
	case TestReportParticipantTypeTestEngine:
		return "Test Engine"
	case TestReportParticipantTypeClient:
		return "Client"
	case TestReportParticipantTypeServer:
		return "Server"
	}
	return "<unknown>"
}

func (code TestReportParticipantType) Definition() string {
	switch code {
	case TestReportParticipantTypeTestEngine:
		return "The test execution engine."
	case TestReportParticipantTypeClient:
		return "A FHIR Client."
	case TestReportParticipantTypeServer:
		return "A FHIR Server."
	}
	return "<unknown>"
}
