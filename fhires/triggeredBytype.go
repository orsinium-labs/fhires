package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// Codes providing the type of triggeredBy observation.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/observation-triggeredbytype
type TriggeredBytype int

const (
	TriggeredBytypeReflex TriggeredBytype = iota
	// Performance of the same test again with the same parameters/settings/solution.
	TriggeredBytypeRepeat
	// Performance of the same test but with different parameters/settings/solution.
	TriggeredBytypeReRun
)

func (code TriggeredBytype) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *TriggeredBytype) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "reflex":
		*code = TriggeredBytypeReflex
	case "repeat":
		*code = TriggeredBytypeRepeat
	case "re-run":
		*code = TriggeredBytypeReRun
	default:
		return fmt.Errorf("unknown TriggeredBytype code `%s`", s)
	}
	return nil
}

func (code TriggeredBytype) String() string {
	return code.Code()
}

func (code TriggeredBytype) Code() string {
	switch code {
	case TriggeredBytypeReflex:
		return "reflex"
	case TriggeredBytypeRepeat:
		return "repeat"
	case TriggeredBytypeReRun:
		return "re-run"
	}
	return "<unknown>"
}

func (code TriggeredBytype) Display() string {
	switch code {
	case TriggeredBytypeReflex:
		return "Reflex"
	case TriggeredBytypeRepeat:
		return "Repeat (per policy)"
	case TriggeredBytypeReRun:
		return "Re-run (per policy)"
	}
	return "<unknown>"
}

func (code TriggeredBytype) Definition() string {
	switch code {
	case TriggeredBytypeReflex:
		return "Performance of one or more other tests depending on the results of the initial test.  This may include collection of additional specimen. While a new ServiceRequest is not required to perform the additional test, where it is still needed (e.g., requesting another laboratory to perform the reflex test), the Observation.basedOn would reference the new ServiceRequest that requested the additional test to be performed as well as the original ServiceRequest to reflect the one that provided the authorization."
	case TriggeredBytypeRepeat:
		return "Performance of the same test again with the same parameters/settings/solution."
	case TriggeredBytypeReRun:
		return "Performance of the same test but with different parameters/settings/solution."
	}
	return "<unknown>"
}
