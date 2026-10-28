package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// The status of a guidance response.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/guidance-response-status
type GuidanceResponseStatus int

const (
	GuidanceResponseStatusSuccess GuidanceResponseStatus = iota
	// The request was processed successfully, but more data may result in a more complete evaluation.
	GuidanceResponseStatusDataRequested
	// The request was processed, but more data is required to complete the evaluation.
	GuidanceResponseStatusDataRequired
	// The request is currently being processed.
	GuidanceResponseStatusInProgress
	// The request was not processed successfully.
	GuidanceResponseStatusFailure
	// The response was entered in error.
	GuidanceResponseStatusEnteredInError
)

func (code GuidanceResponseStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *GuidanceResponseStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "success":
		*code = GuidanceResponseStatusSuccess
	case "data-requested":
		*code = GuidanceResponseStatusDataRequested
	case "data-required":
		*code = GuidanceResponseStatusDataRequired
	case "in-progress":
		*code = GuidanceResponseStatusInProgress
	case "failure":
		*code = GuidanceResponseStatusFailure
	case "entered-in-error":
		*code = GuidanceResponseStatusEnteredInError
	default:
		return fmt.Errorf("unknown GuidanceResponseStatus code `%s`", s)
	}
	return nil
}

func (code GuidanceResponseStatus) String() string {
	return code.Code()
}

func (code GuidanceResponseStatus) Code() string {
	switch code {
	case GuidanceResponseStatusSuccess:
		return "success"
	case GuidanceResponseStatusDataRequested:
		return "data-requested"
	case GuidanceResponseStatusDataRequired:
		return "data-required"
	case GuidanceResponseStatusInProgress:
		return "in-progress"
	case GuidanceResponseStatusFailure:
		return "failure"
	case GuidanceResponseStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}

func (code GuidanceResponseStatus) Display() string {
	switch code {
	case GuidanceResponseStatusSuccess:
		return "Success"
	case GuidanceResponseStatusDataRequested:
		return "Data Requested"
	case GuidanceResponseStatusDataRequired:
		return "Data Required"
	case GuidanceResponseStatusInProgress:
		return "In Progress"
	case GuidanceResponseStatusFailure:
		return "Failure"
	case GuidanceResponseStatusEnteredInError:
		return "Entered In Error"
	}
	return "<unknown>"
}

func (code GuidanceResponseStatus) Definition() string {
	switch code {
	case GuidanceResponseStatusSuccess:
		return "The request was processed successfully."
	case GuidanceResponseStatusDataRequested:
		return "The request was processed successfully, but more data may result in a more complete evaluation."
	case GuidanceResponseStatusDataRequired:
		return "The request was processed, but more data is required to complete the evaluation."
	case GuidanceResponseStatusInProgress:
		return "The request is currently being processed."
	case GuidanceResponseStatusFailure:
		return "The request was not processed successfully."
	case GuidanceResponseStatusEnteredInError:
		return "The response was entered in error."
	}
	return "<unknown>"
}
