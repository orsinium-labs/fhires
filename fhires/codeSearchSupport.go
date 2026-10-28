package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// CodeSearchSupport is documented here http://hl7.org/fhir/ValueSet/code-search-support
type CodeSearchSupport int

const (
	CodeSearchSupportInCompose CodeSearchSupport = iota
	CodeSearchSupportInExpansion
	CodeSearchSupportInComposeOrExpansion
)

func (code CodeSearchSupport) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *CodeSearchSupport) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "in-compose":
		*code = CodeSearchSupportInCompose
	case "in-expansion":
		*code = CodeSearchSupportInExpansion
	case "in-compose-or-expansion":
		*code = CodeSearchSupportInComposeOrExpansion
	default:
		return fmt.Errorf("unknown CodeSearchSupport code `%s`", s)
	}
	return nil
}

func (code CodeSearchSupport) String() string {
	return code.Code()
}

func (code CodeSearchSupport) Code() string {
	switch code {
	case CodeSearchSupportInCompose:
		return "in-compose"
	case CodeSearchSupportInExpansion:
		return "in-expansion"
	case CodeSearchSupportInComposeOrExpansion:
		return "in-compose-or-expansion"
	}
	return "<unknown>"
}

func (code CodeSearchSupport) Display() string {
	switch code {
	case CodeSearchSupportInCompose:
		return "In Compose"
	case CodeSearchSupportInExpansion:
		return "In Expansion"
	case CodeSearchSupportInComposeOrExpansion:
		return "In Compose Or Expansion"
	}
	return "<unknown>"
}

func (code CodeSearchSupport) Definition() string {
	switch code {
	case CodeSearchSupportInCompose:
		return "The search for code on ValueSet returns ValueSet resources where the code is included in the extensional definition of the ValueSet."
	case CodeSearchSupportInExpansion:
		return "The search for code on ValueSet returns ValueSet resources where the code is contained in the  ValueSet expansion."
	case CodeSearchSupportInComposeOrExpansion:
		return "The search for code on ValueSet returns ValueSet resources where the code is included in the extensional definition or contained in the ValueSet expansion."
	}
	return "<unknown>"
}
