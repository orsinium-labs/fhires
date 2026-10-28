package fhires

import (
	"encoding/json"
	"fmt"
	"strings"
)

// SubmitDataUpdateType is documented here http://hl7.org/fhir/ValueSet/submit-data-update-type
type SubmitDataUpdateType int

const (
	SubmitDataUpdateTypeIncremental SubmitDataUpdateType = iota
	SubmitDataUpdateTypeSnapshot
)

func (code SubmitDataUpdateType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *SubmitDataUpdateType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "incremental":
		*code = SubmitDataUpdateTypeIncremental
	case "snapshot":
		*code = SubmitDataUpdateTypeSnapshot
	default:
		return fmt.Errorf("unknown SubmitDataUpdateType code `%s`", s)
	}
	return nil
}

func (code SubmitDataUpdateType) String() string {
	return code.Code()
}

func (code SubmitDataUpdateType) Code() string {
	switch code {
	case SubmitDataUpdateTypeIncremental:
		return "incremental"
	case SubmitDataUpdateTypeSnapshot:
		return "snapshot"
	}
	return "<unknown>"
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
