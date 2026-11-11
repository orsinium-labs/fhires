package fhires

// The kind of coverage: insurance, selfpay or other.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/coverage-kind
type Kind string

const (
	// The Coverage provides the identifiers and card-level details of an insurance policy.
	KindInsurance Kind = "insurance"
	// One or more persons and/or organizations are paying for the services rendered.
	KindSelfPay Kind = "self-pay"
	// Some other organization is paying for the service.
	KindOther Kind = "other"
)

func (code Kind) String() string {
	return code.Code()
}

func (code Kind) Code() string {
	return string(code)
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
