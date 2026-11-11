package fhires

// Used to distinguish different roles a resource can play within a set of linked resources.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/linkage-type
type LinkageType string

const (
	// The resource represents the "source of truth" (from the perspective of this Linkage resource) for the underlying event/condition/etc.
	LinkageTypeSource LinkageType = "source"
	// The resource represents an alternative view of the underlying event/condition/etc.  The resource may still be actively maintained, even though it is not considered to be the source of truth.
	LinkageTypeAlternate LinkageType = "alternate"
	// The resource represents an obsolete record of the underlying event/condition/etc.  It is not expected to be actively maintained.
	LinkageTypeHistorical LinkageType = "historical"
)

func (code LinkageType) String() string {
	return code.Code()
}

func (code LinkageType) Code() string {
	return string(code)
}

func (code LinkageType) Display() string {
	switch code {
	case LinkageTypeSource:
		return "Source of Truth"
	case LinkageTypeAlternate:
		return "Alternate Record"
	case LinkageTypeHistorical:
		return "Historical/Obsolete Record"
	}
	return "<unknown>"
}

func (code LinkageType) Definition() string {
	switch code {
	case LinkageTypeSource:
		return "The resource represents the \"source of truth\" (from the perspective of this Linkage resource) for the underlying event/condition/etc."
	case LinkageTypeAlternate:
		return "The resource represents an alternative view of the underlying event/condition/etc.  The resource may still be actively maintained, even though it is not considered to be the source of truth."
	case LinkageTypeHistorical:
		return "The resource represents an obsolete record of the underlying event/condition/etc.  It is not expected to be actively maintained."
	}
	return "<unknown>"
}
