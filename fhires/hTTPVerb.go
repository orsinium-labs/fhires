package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// HTTP verbs (in the HTTP command line).
//
// See [HTTP rfc](https://tools.ietf.org/html/rfc7231) for details.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/http-verb
type HTTPVerb int

const (
	HTTPVerbGET HTTPVerb = iota
	// HTTP HEAD Command.
	HTTPVerbHEAD
	// HTTP POST Command.
	HTTPVerbPOST
	// HTTP PUT Command.
	HTTPVerbPUT
	// HTTP DELETE Command.
	HTTPVerbDELETE
	// HTTP PATCH Command.
	HTTPVerbPATCH
)

func (code HTTPVerb) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *HTTPVerb) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "GET":
		*code = HTTPVerbGET
	case "HEAD":
		*code = HTTPVerbHEAD
	case "POST":
		*code = HTTPVerbPOST
	case "PUT":
		*code = HTTPVerbPUT
	case "DELETE":
		*code = HTTPVerbDELETE
	case "PATCH":
		*code = HTTPVerbPATCH
	default:
		return fmt.Errorf("unknown HTTPVerb code `%s`", s)
	}
	return nil
}

func (code HTTPVerb) String() string {
	return code.Code()
}

func (code HTTPVerb) Code() string {
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
