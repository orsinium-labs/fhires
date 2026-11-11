package fhires

// If field is a list, how to manage the source.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/map-source-list-mode
type StructureMapSourceListMode string

const (
	// Only process this rule for the first in the list.
	StructureMapSourceListModeFirst StructureMapSourceListMode = "first"
	// Process this rule for all but the first.
	StructureMapSourceListModeNot_first StructureMapSourceListMode = "not_first"
	// Only process this rule for the last in the list.
	StructureMapSourceListModeLast StructureMapSourceListMode = "last"
	// Process this rule for all but the last.
	StructureMapSourceListModeNot_last StructureMapSourceListMode = "not_last"
	// Only process this rule is there is only item.
	StructureMapSourceListModeOnly_one StructureMapSourceListMode = "only_one"
)

func (code StructureMapSourceListMode) String() string {
	return code.Code()
}

func (code StructureMapSourceListMode) Code() string {
	return string(code)
}

func (code StructureMapSourceListMode) Display() string {
	switch code {
	case StructureMapSourceListModeFirst:
		return "First"
	case StructureMapSourceListModeNot_first:
		return "All but the first"
	case StructureMapSourceListModeLast:
		return "Last"
	case StructureMapSourceListModeNot_last:
		return "All but the last"
	case StructureMapSourceListModeOnly_one:
		return "Enforce only one"
	}
	return "<unknown>"
}

func (code StructureMapSourceListMode) Definition() string {
	switch code {
	case StructureMapSourceListModeFirst:
		return "Only process this rule for the first in the list."
	case StructureMapSourceListModeNot_first:
		return "Process this rule for all but the first."
	case StructureMapSourceListModeLast:
		return "Only process this rule for the last in the list."
	case StructureMapSourceListModeNot_last:
		return "Process this rule for all but the last."
	case StructureMapSourceListModeOnly_one:
		return "Only process this rule is there is only item."
	}
	return "<unknown>"
}
