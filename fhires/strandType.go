package fhires

import (
	"encoding/json"
	"fmt"
	"strings"
)

// StrandType is documented here http://hl7.org/fhir/ValueSet/strand-type
type StrandType int

const (
	StrandTypeWatson StrandType = iota
	StrandTypeCrick
)

func (code StrandType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *StrandType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "watson":
		*code = StrandTypeWatson
	case "crick":
		*code = StrandTypeCrick
	default:
		return fmt.Errorf("unknown StrandType code `%s`", s)
	}
	return nil
}

func (code StrandType) String() string {
	return code.Code()
}

func (code StrandType) Code() string {
	switch code {
	case StrandTypeWatson:
		return "watson"
	case StrandTypeCrick:
		return "crick"
	}
	return "<unknown>"
}

func (code StrandType) Display() string {
	switch code {
	case StrandTypeWatson:
		return "Watson strand of starting sequence"
	case StrandTypeCrick:
		return "Crick strand of starting sequence"
	}
	return "<unknown>"
}

func (code StrandType) Definition() string {
	switch code {
	case StrandTypeWatson:
		return "Watson strand of starting sequence."
	case StrandTypeCrick:
		return "Crick strand of starting sequence."
	}
	return "<unknown>"
}
