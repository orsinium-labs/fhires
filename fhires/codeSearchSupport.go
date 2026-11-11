package fhires

// The degree to which the server supports the code search parameter on ValueSet, if it is supported.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/code-search-support
type CodeSearchSupport string

const (
	// The search for code on ValueSet returns ValueSet resources where the code is included in the extensional definition of the ValueSet.
	CodeSearchSupportInCompose CodeSearchSupport = "in-compose"
	// The search for code on ValueSet returns ValueSet resources where the code is contained in the  ValueSet expansion.
	CodeSearchSupportInExpansion CodeSearchSupport = "in-expansion"
	// The search for code on ValueSet returns ValueSet resources where the code is included in the extensional definition or contained in the ValueSet expansion.
	CodeSearchSupportInComposeOrExpansion CodeSearchSupport = "in-compose-or-expansion"
)

func (code CodeSearchSupport) String() string {
	return code.Code()
}

func (code CodeSearchSupport) Code() string {
	return string(code)
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
