package fhires

// The possible sort directions, ascending or descending.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/sort-direction
type SortDirection string

const (
	// Sort by the value ascending, so that lower values appear first.
	SortDirectionAscending SortDirection = "ascending"
	// Sort by the value descending, so that lower values appear last.
	SortDirectionDescending SortDirection = "descending"
)

func (code SortDirection) String() string {
	return code.Code()
}

func (code SortDirection) Code() string {
	return string(code)
}

func (code SortDirection) Display() string {
	switch code {
	case SortDirectionAscending:
		return "Ascending"
	case SortDirectionDescending:
		return "Descending"
	}
	return "<unknown>"
}

func (code SortDirection) Definition() string {
	switch code {
	case SortDirectionAscending:
		return "Sort by the value ascending, so that lower values appear first."
	case SortDirectionDescending:
		return "Sort by the value descending, so that lower values appear last."
	}
	return "<unknown>"
}
