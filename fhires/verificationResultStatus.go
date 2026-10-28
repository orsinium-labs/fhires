package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// VerificationResultStatus is documented here http://hl7.org/fhir/ValueSet/verificationresult-status
type VerificationResultStatus int

const (
	VerificationResultStatusAttested VerificationResultStatus = iota
	VerificationResultStatusValidated
	VerificationResultStatusInProcess
	VerificationResultStatusReqRevalid
	VerificationResultStatusValFail
	VerificationResultStatusRevalFail
	VerificationResultStatusEnteredInError
)

func (code VerificationResultStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *VerificationResultStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "attested":
		*code = VerificationResultStatusAttested
	case "validated":
		*code = VerificationResultStatusValidated
	case "in-process":
		*code = VerificationResultStatusInProcess
	case "req-revalid":
		*code = VerificationResultStatusReqRevalid
	case "val-fail":
		*code = VerificationResultStatusValFail
	case "reval-fail":
		*code = VerificationResultStatusRevalFail
	case "entered-in-error":
		*code = VerificationResultStatusEnteredInError
	default:
		return fmt.Errorf("unknown VerificationResultStatus code `%s`", s)
	}
	return nil
}

func (code VerificationResultStatus) String() string {
	return code.Code()
}

func (code VerificationResultStatus) Code() string {
	switch code {
	case VerificationResultStatusAttested:
		return "attested"
	case VerificationResultStatusValidated:
		return "validated"
	case VerificationResultStatusInProcess:
		return "in-process"
	case VerificationResultStatusReqRevalid:
		return "req-revalid"
	case VerificationResultStatusValFail:
		return "val-fail"
	case VerificationResultStatusRevalFail:
		return "reval-fail"
	case VerificationResultStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}

func (code VerificationResultStatus) Display() string {
	switch code {
	case VerificationResultStatusAttested:
		return "Attested"
	case VerificationResultStatusValidated:
		return "Validated"
	case VerificationResultStatusInProcess:
		return "In process"
	case VerificationResultStatusReqRevalid:
		return "Requires revalidation"
	case VerificationResultStatusValFail:
		return "Validation failed"
	case VerificationResultStatusRevalFail:
		return "Re-Validation failed"
	case VerificationResultStatusEnteredInError:
		return "Entered in Error"
	}
	return "<unknown>"
}

func (code VerificationResultStatus) Definition() string {
	switch code {
	case VerificationResultStatusAttested:
		return "***TODO***"
	case VerificationResultStatusValidated:
		return "***TODO***"
	case VerificationResultStatusInProcess:
		return "***TODO***"
	case VerificationResultStatusReqRevalid:
		return "***TODO***"
	case VerificationResultStatusValFail:
		return "***TODO***"
	case VerificationResultStatusRevalFail:
		return "***TODO***"
	case VerificationResultStatusEnteredInError:
		return "The VerificationResult record was created erroneously and is not appropriated for use."
	}
	return "<unknown>"
}
