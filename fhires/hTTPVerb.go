package fhires

// HTTP verbs (in the HTTP command line).
//
// See [HTTP rfc](https://tools.ietf.org/html/rfc7231) for details.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/http-verb
type HTTPVerb string

const (
	// HTTP GET Command.
	HTTPVerbGET HTTPVerb = "GET"
	// HTTP HEAD Command.
	HTTPVerbHEAD HTTPVerb = "HEAD"
	// HTTP POST Command.
	HTTPVerbPOST HTTPVerb = "POST"
	// HTTP PUT Command.
	HTTPVerbPUT HTTPVerb = "PUT"
	// HTTP DELETE Command.
	HTTPVerbDELETE HTTPVerb = "DELETE"
	// HTTP PATCH Command.
	HTTPVerbPATCH HTTPVerb = "PATCH"
)

func (code HTTPVerb) String() string {
	return code.Code()
}

func (code HTTPVerb) Code() string {
	return string(code)
}

func (code HTTPVerb) Display() string {
	switch code {
	case HTTPVerbGET:
		return "GET"
	case HTTPVerbHEAD:
		return "HEAD"
	case HTTPVerbPOST:
		return "POST"
	case HTTPVerbPUT:
		return "PUT"
	case HTTPVerbDELETE:
		return "DELETE"
	case HTTPVerbPATCH:
		return "PATCH"
	}
	return "<unknown>"
}

func (code HTTPVerb) Definition() string {
	switch code {
	case HTTPVerbGET:
		return "HTTP GET Command."
	case HTTPVerbHEAD:
		return "HTTP HEAD Command."
	case HTTPVerbPOST:
		return "HTTP POST Command."
	case HTTPVerbPUT:
		return "HTTP PUT Command."
	case HTTPVerbDELETE:
		return "HTTP DELETE Command."
	case HTTPVerbPATCH:
		return "HTTP PATCH Command."
	}
	return "<unknown>"
}
