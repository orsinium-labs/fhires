package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// Kind is documented here http://hl7.org/fhir/ValueSet/coverage-kind
type Kind int

const (
	KindInsurance Kind = iota
	KindSelfPay
	KindOther
)

func (code Kind) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *Kind) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "insurance":
		*code = KindInsurance
	case "self-pay":
		*code = KindSelfPay
	case "other":
		*code = KindOther
	default:
		return fmt.Errorf("unknown Kind code `%s`", s)
	}
	return nil
}

func (code Kind) String() string {
	return code.Code()
}

func (code Kind) Code() string {
	switch code {
	case KindInsurance:
		return "insurance"
	case KindSelfPay:
		return "self-pay"
	case KindOther:
		return "other"
	}
	return "<unknown>"
}

func (code Kind) Display() string {
	switch code {
	case KindInsurance:
		return "Insurance"
	case KindSelfPay:
		return "Self-pay"
	case KindOther:
		return "Other"
	}
	return "<unknown>"
}

func (code Kind) Definition() string {
	switch code {
	case KindInsurance:
		return "The Coverage provides the identifiers and card-level details of an insurance policy."
	case KindSelfPay:
		return "One or more persons and/or organizations are paying for the services rendered."
	case KindOther:
		return "Some other organization is paying for the service."
	}
	return "<unknown>"
}
