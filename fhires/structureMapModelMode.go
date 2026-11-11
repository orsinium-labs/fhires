package fhires

// How the referenced structure is used in this mapping.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/map-model-mode
type StructureMapModelMode string

const (
	// This structure describes an instance passed to the mapping engine that is used a source of data.
	StructureMapModelModeSource StructureMapModelMode = "source"
	// This structure describes an instance that the mapping engine may ask for that is used a source of data.
	StructureMapModelModeQueried StructureMapModelMode = "queried"
	// This structure describes an instance passed to the mapping engine that is used a target of data.
	StructureMapModelModeTarget StructureMapModelMode = "target"
	// This structure describes an instance that the mapping engine may ask to create that is used a target of data.
	StructureMapModelModeProduced StructureMapModelMode = "produced"
)

func (code StructureMapModelMode) String() string {
	return code.Code()
}

func (code StructureMapModelMode) Code() string {
	return string(code)
}

func (code StructureMapModelMode) Display() string {
	switch code {
	case StructureMapModelModeSource:
		return "Source Structure Definition"
	case StructureMapModelModeQueried:
		return "Queried Structure Definition"
	case StructureMapModelModeTarget:
		return "Target Structure Definition"
	case StructureMapModelModeProduced:
		return "Produced Structure Definition"
	}
	return "<unknown>"
}

func (code StructureMapModelMode) Definition() string {
	switch code {
	case StructureMapModelModeSource:
		return "This structure describes an instance passed to the mapping engine that is used a source of data."
	case StructureMapModelModeQueried:
		return "This structure describes an instance that the mapping engine may ask for that is used a source of data."
	case StructureMapModelModeTarget:
		return "This structure describes an instance passed to the mapping engine that is used a target of data."
	case StructureMapModelModeProduced:
		return "This structure describes an instance that the mapping engine may ask to create that is used a target of data."
	}
	return "<unknown>"
}
