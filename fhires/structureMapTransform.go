package fhires

// How data is copied/created.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/map-transform
type StructureMapTransform string

const (
	// create(type : string) - type is passed through to the application on the standard API, and must be known by it.
	StructureMapTransformCreate StructureMapTransform = "create"
	// copy(source).
	StructureMapTransformCopy StructureMapTransform = "copy"
	// truncate(source, length) - source must be stringy type.
	StructureMapTransformTruncate StructureMapTransform = "truncate"
	// escape(source, fmt1, fmt2) - change source from one kind of escaping to another (plain, java, xml, json). note that this is for when the string itself is escaped.
	StructureMapTransformEscape StructureMapTransform = "escape"
	// cast(source, type?) - cast (convert) source from one type to another. Target type can be left as implicit if there is one and only one target type known. The default namespace for the type is 'FHIR' (see [FHIRPath type specifiers](http://hl7.org/fhirpath/N1/#is-type-specifier))
	StructureMapTransformCast StructureMapTransform = "cast"
	// append(source...) - source is element or string.
	StructureMapTransformAppend StructureMapTransform = "append"
	// translate(source, uri_of_map) - use the translate operation.
	StructureMapTransformTranslate StructureMapTransform = "translate"
	// reference(source : object) - return a string that references the provided tree properly.
	StructureMapTransformReference StructureMapTransform = "reference"
	// Perform a date operation. *Parameters to be documented*.
	StructureMapTransformDateOp StructureMapTransform = "dateOp"
	// Generate a random UUID (in lowercase). No Parameters.
	StructureMapTransformUuid StructureMapTransform = "uuid"
	// Return the appropriate string to put in a reference that refers to the resource provided as a parameter.
	StructureMapTransformPointer StructureMapTransform = "pointer"
	// Execute the supplied FHIRPath expression and use the value returned by that.
	StructureMapTransformEvaluate StructureMapTransform = "evaluate"
	// Create a CodeableConcept. Parameters = (text) or (system. Code[, display]).
	StructureMapTransformCc StructureMapTransform = "cc"
	// Create a Coding. Parameters = (system. Code[, display]).
	StructureMapTransformC StructureMapTransform = "c"
	// Create a quantity. Parameters = (text) or (value, unit, [system, code]) where text is the natural representation e.g. [comparator]value[space]unit.
	StructureMapTransformQty StructureMapTransform = "qty"
	// Create an identifier. Parameters = (system, value[, type]) where type is a code from the identifier type value set.
	StructureMapTransformId StructureMapTransform = "id"
	// Create a contact details. Parameters = (value) or (system, value). If no system is provided, the system should be inferred from the content of the value.
	StructureMapTransformCp StructureMapTransform = "cp"
)

func (code StructureMapTransform) String() string {
	return code.Code()
}

func (code StructureMapTransform) Code() string {
	return string(code)
}

func (code StructureMapTransform) Display() string {
	switch code {
	case StructureMapTransformCreate:
		return "create"
	case StructureMapTransformCopy:
		return "copy"
	case StructureMapTransformTruncate:
		return "truncate"
	case StructureMapTransformEscape:
		return "escape"
	case StructureMapTransformCast:
		return "cast"
	case StructureMapTransformAppend:
		return "append"
	case StructureMapTransformTranslate:
		return "translate"
	case StructureMapTransformReference:
		return "reference"
	case StructureMapTransformDateOp:
		return "dateOp"
	case StructureMapTransformUuid:
		return "uuid"
	case StructureMapTransformPointer:
		return "pointer"
	case StructureMapTransformEvaluate:
		return "evaluate"
	case StructureMapTransformCc:
		return "cc"
	case StructureMapTransformC:
		return "c"
	case StructureMapTransformQty:
		return "qty"
	case StructureMapTransformId:
		return "id"
	case StructureMapTransformCp:
		return "cp"
	}
	return "<unknown>"
}

func (code StructureMapTransform) Definition() string {
	switch code {
	case StructureMapTransformCreate:
		return "create(type : string) - type is passed through to the application on the standard API, and must be known by it."
	case StructureMapTransformCopy:
		return "copy(source)."
	case StructureMapTransformTruncate:
		return "truncate(source, length) - source must be stringy type."
	case StructureMapTransformEscape:
		return "escape(source, fmt1, fmt2) - change source from one kind of escaping to another (plain, java, xml, json). note that this is for when the string itself is escaped."
	case StructureMapTransformCast:
		return "cast(source, type?) - cast (convert) source from one type to another. Target type can be left as implicit if there is one and only one target type known. The default namespace for the type is 'FHIR' (see [FHIRPath type specifiers](http://hl7.org/fhirpath/N1/#is-type-specifier))"
	case StructureMapTransformAppend:
		return "append(source...) - source is element or string."
	case StructureMapTransformTranslate:
		return "translate(source, uri_of_map) - use the translate operation."
	case StructureMapTransformReference:
		return "reference(source : object) - return a string that references the provided tree properly."
	case StructureMapTransformDateOp:
		return "Perform a date operation. *Parameters to be documented*."
	case StructureMapTransformUuid:
		return "Generate a random UUID (in lowercase). No Parameters."
	case StructureMapTransformPointer:
		return "Return the appropriate string to put in a reference that refers to the resource provided as a parameter."
	case StructureMapTransformEvaluate:
		return "Execute the supplied FHIRPath expression and use the value returned by that."
	case StructureMapTransformCc:
		return "Create a CodeableConcept. Parameters = (text) or (system. Code[, display])."
	case StructureMapTransformC:
		return "Create a Coding. Parameters = (system. Code[, display])."
	case StructureMapTransformQty:
		return "Create a quantity. Parameters = (text) or (value, unit, [system, code]) where text is the natural representation e.g. [comparator]value[space]unit."
	case StructureMapTransformId:
		return "Create an identifier. Parameters = (system, value[, type]) where type is a code from the identifier type value set."
	case StructureMapTransformCp:
		return "Create a contact details. Parameters = (value) or (system, value). If no system is provided, the system should be inferred from the content of the value."
	}
	return "<unknown>"
}
