package fhires

// A code that indicates how the server supports conditional read.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/conditional-read-status
type ConditionalReadStatus string

const (
	// No support for conditional reads.
	ConditionalReadStatusNotSupported ConditionalReadStatus = "not-supported"
	// Conditional reads are supported, but only with the If-Modified-Since HTTP Header.
	ConditionalReadStatusModifiedSince ConditionalReadStatus = "modified-since"
	// Conditional reads are supported, but only with the If-None-Match HTTP Header.
	ConditionalReadStatusNotMatch ConditionalReadStatus = "not-match"
	// Conditional reads are supported, with both If-Modified-Since and If-None-Match HTTP Headers.
	ConditionalReadStatusFullSupport ConditionalReadStatus = "full-support"
)

func (code ConditionalReadStatus) String() string {
	return code.Code()
}

func (code ConditionalReadStatus) Code() string {
	return string(code)
}

func (code ConditionalReadStatus) Display() string {
	switch code {
	case ConditionalReadStatusNotSupported:
		return "Not Supported"
	case ConditionalReadStatusModifiedSince:
		return "If-Modified-Since"
	case ConditionalReadStatusNotMatch:
		return "If-None-Match"
	case ConditionalReadStatusFullSupport:
		return "Full Support"
	}
	return "<unknown>"
}

func (code ConditionalReadStatus) Definition() string {
	switch code {
	case ConditionalReadStatusNotSupported:
		return "No support for conditional reads."
	case ConditionalReadStatusModifiedSince:
		return "Conditional reads are supported, but only with the If-Modified-Since HTTP Header."
	case ConditionalReadStatusNotMatch:
		return "Conditional reads are supported, but only with the If-None-Match HTTP Header."
	case ConditionalReadStatusFullSupport:
		return "Conditional reads are supported, with both If-Modified-Since and If-None-Match HTTP Headers."
	}
	return "<unknown>"
}
