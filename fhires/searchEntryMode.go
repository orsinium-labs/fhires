package fhires

// Why an entry is in the result set - whether it's included as a match or because of an _include requirement, or to convey information or warning information about the search process.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/search-entry-mode
type SearchEntryMode string

const (
	// This resource matched the search specification.
	SearchEntryModeMatch SearchEntryMode = "match"
	// This resource is returned because it is referred to from another resource in the search set.
	SearchEntryModeInclude SearchEntryMode = "include"
	// An OperationOutcome that provides additional information about the processing of a search.
	SearchEntryModeOutcome SearchEntryMode = "outcome"
)

func (code SearchEntryMode) String() string {
	return code.Code()
}

func (code SearchEntryMode) Code() string {
	return string(code)
}

func (code SearchEntryMode) Display() string {
	switch code {
	case SearchEntryModeMatch:
		return "Match"
	case SearchEntryModeInclude:
		return "Include"
	case SearchEntryModeOutcome:
		return "Outcome"
	}
	return "<unknown>"
}

func (code SearchEntryMode) Definition() string {
	switch code {
	case SearchEntryModeMatch:
		return "This resource matched the search specification."
	case SearchEntryModeInclude:
		return "This resource is returned because it is referred to from another resource in the search set."
	case SearchEntryModeOutcome:
		return "An OperationOutcome that provides additional information about the processing of a search."
	}
	return "<unknown>"
}
