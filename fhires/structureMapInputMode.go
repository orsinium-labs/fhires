package fhires

// Mode for this instance of data.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/map-input-mode
type StructureMapInputMode string

const (
	// Names an input instance used a source for mapping.
	StructureMapInputModeSource StructureMapInputMode = "source"
	// Names an instance that is being populated.
	StructureMapInputModeTarget StructureMapInputMode = "target"
)

func (code StructureMapInputMode) String() string {
	return code.Code()
}

func (code StructureMapInputMode) Code() string {
	return string(code)
}

func (code StructureMapInputMode) Display() string {
	switch code {
	case StructureMapInputModeSource:
		return "Source Instance"
	case StructureMapInputModeTarget:
		return "Target Instance"
	}
	return "<unknown>"
}

func (code StructureMapInputMode) Definition() string {
	switch code {
	case StructureMapInputModeSource:
		return "Names an input instance used a source for mapping."
	case StructureMapInputModeTarget:
		return "Names an instance that is being populated."
	}
	return "<unknown>"
}
