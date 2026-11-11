package fhires

// Data types allowed to be used for search parameters.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/search-param-type
type SearchParamType string

const (
	// Search parameter SHALL be a number (a whole number, or a decimal).
	SearchParamTypeNumber SearchParamType = "number"
	// Search parameter is on a date/time. The date format is the standard XML format, though other formats may be supported.
	SearchParamTypeDate SearchParamType = "date"
	// Search parameter is a simple string, like a name part. Search is case-insensitive and accent-insensitive. May match just the start of a string. String parameters may contain spaces.
	SearchParamTypeString SearchParamType = "string"
	// Search parameter on a coded element or identifier. May be used to search through the text, display, code and code/codesystem (for codes) and label, system and key (for identifier). Its value is either a string or a pair of namespace and value, separated by a "|", depending on the modifier used.
	SearchParamTypeToken SearchParamType = "token"
	// A reference to another resource (Reference or canonical).
	SearchParamTypeReference SearchParamType = "reference"
	// A composite search parameter that combines a search on two values together.
	SearchParamTypeComposite SearchParamType = "composite"
	// A search parameter that searches on a quantity.
	SearchParamTypeQuantity SearchParamType = "quantity"
	// A search parameter that searches on a URI (RFC 3986).
	SearchParamTypeUri SearchParamType = "uri"
	// Special logic applies to this parameter per the description of the search parameter.
	SearchParamTypeSpecial SearchParamType = "special"
)

func (code SearchParamType) String() string {
	return code.Code()
}

func (code SearchParamType) Code() string {
	return string(code)
}

func (code SearchParamType) Display() string {
	switch code {
	case SearchParamTypeNumber:
		return "Number"
	case SearchParamTypeDate:
		return "Date/DateTime"
	case SearchParamTypeString:
		return "String"
	case SearchParamTypeToken:
		return "Token"
	case SearchParamTypeReference:
		return "Reference"
	case SearchParamTypeComposite:
		return "Composite"
	case SearchParamTypeQuantity:
		return "Quantity"
	case SearchParamTypeUri:
		return "URI"
	case SearchParamTypeSpecial:
		return "Special"
	}
	return "<unknown>"
}

func (code SearchParamType) Definition() string {
	switch code {
	case SearchParamTypeNumber:
		return "Search parameter SHALL be a number (a whole number, or a decimal)."
	case SearchParamTypeDate:
		return "Search parameter is on a date/time. The date format is the standard XML format, though other formats may be supported."
	case SearchParamTypeString:
		return "Search parameter is a simple string, like a name part. Search is case-insensitive and accent-insensitive. May match just the start of a string. String parameters may contain spaces."
	case SearchParamTypeToken:
		return "Search parameter on a coded element or identifier. May be used to search through the text, display, code and code/codesystem (for codes) and label, system and key (for identifier). Its value is either a string or a pair of namespace and value, separated by a \"|\", depending on the modifier used."
	case SearchParamTypeReference:
		return "A reference to another resource (Reference or canonical)."
	case SearchParamTypeComposite:
		return "A composite search parameter that combines a search on two values together."
	case SearchParamTypeQuantity:
		return "A search parameter that searches on a quantity."
	case SearchParamTypeUri:
		return "A search parameter that searches on a URI (RFC 3986)."
	case SearchParamTypeSpecial:
		return "Special logic applies to this parameter per the description of the search parameter."
	}
	return "<unknown>"
}
