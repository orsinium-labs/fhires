package fhires

// If this is the default rule set to apply for the source type, or this combination of types.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/map-group-type-mode
type StructureMapGroupTypeMode string

const (
	// This group is a default mapping group for the specified types and for the primary source type.
	StructureMapGroupTypeModeTypes StructureMapGroupTypeMode = "types"
	// This group is a default mapping group for the specified types.
	StructureMapGroupTypeModeTypeAndTypes StructureMapGroupTypeMode = "type-and-types"
)

func (code StructureMapGroupTypeMode) String() string {
	return code.Code()
}

func (code StructureMapGroupTypeMode) Code() string {
	return string(code)
}

func (code StructureMapGroupTypeMode) Display() string {
	switch code {
	case StructureMapGroupTypeModeTypes:
		return "Default for Type Combination"
	case StructureMapGroupTypeModeTypeAndTypes:
		return "Default for type + combination"
	}
	return "<unknown>"
}

func (code StructureMapGroupTypeMode) Definition() string {
	switch code {
	case StructureMapGroupTypeModeTypes:
		return "This group is a default mapping group for the specified types and for the primary source type."
	case StructureMapGroupTypeModeTypeAndTypes:
		return "This group is a default mapping group for the specified types."
	}
	return "<unknown>"
}
