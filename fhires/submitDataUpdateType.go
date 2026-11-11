package fhires

// Concepts for how a measure report consumer and receiver coordinate data exchange updates.
//
// # The choices are snapshot or incremental updates
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/submit-data-update-type
type SubmitDataUpdateType string

const (
	// In contrast to the Snapshot Update, the FHIR Parameters resource used in a Submit Data or the Collect Data scenario contains only the new and updated DEQM and QI Core Profiles since the last transaction. If the Consumer supports incremental updates, the contents of the updated payload updates the previous payload data.
	SubmitDataUpdateTypeIncremental SubmitDataUpdateType = "incremental"
	// In contrast to the Incremental Update, the FHIR Parameters resource used in a Submit Data or the Collect Data scenario contains all the DEQM and QI Core Profiles for each transaction.  If the Consumer supports snapshot updates, the contents of the updated payload entirely replaces the previous payload
	SubmitDataUpdateTypeSnapshot SubmitDataUpdateType = "snapshot"
)

func (code SubmitDataUpdateType) String() string {
	return code.Code()
}

func (code SubmitDataUpdateType) Code() string {
	return string(code)
}

func (code SubmitDataUpdateType) Display() string {
	switch code {
	case SubmitDataUpdateTypeIncremental:
		return "Incremental"
	case SubmitDataUpdateTypeSnapshot:
		return "Snapshot"
	}
	return "<unknown>"
}

func (code SubmitDataUpdateType) Definition() string {
	switch code {
	case SubmitDataUpdateTypeIncremental:
		return "In contrast to the Snapshot Update, the FHIR Parameters resource used in a Submit Data or the Collect Data scenario contains only the new and updated DEQM and QI Core Profiles since the last transaction. If the Consumer supports incremental updates, the contents of the updated payload updates the previous payload data."
	case SubmitDataUpdateTypeSnapshot:
		return "In contrast to the Incremental Update, the FHIR Parameters resource used in a Submit Data or the Collect Data scenario contains all the DEQM and QI Core Profiles for each transaction.  If the Consumer supports snapshot updates, the contents of the updated payload entirely replaces the previous payload"
	}
	return "<unknown>"
}
