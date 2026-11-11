package fhires

// The validation status of the target
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/verificationresult-status
type VerificationResultStatus string

const (
	// ***TODO***
	VerificationResultStatusAttested VerificationResultStatus = "attested"
	// ***TODO***
	VerificationResultStatusValidated VerificationResultStatus = "validated"
	// ***TODO***
	VerificationResultStatusInProcess VerificationResultStatus = "in-process"
	// ***TODO***
	VerificationResultStatusReqRevalid VerificationResultStatus = "req-revalid"
	// ***TODO***
	VerificationResultStatusValFail VerificationResultStatus = "val-fail"
	// ***TODO***
	VerificationResultStatusRevalFail VerificationResultStatus = "reval-fail"
	// The VerificationResult record was created erroneously and is not appropriated for use.
	VerificationResultStatusEnteredInError VerificationResultStatus = "entered-in-error"
)

func (code VerificationResultStatus) String() string {
	return code.Code()
}

func (code VerificationResultStatus) Code() string {
	return string(code)
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
