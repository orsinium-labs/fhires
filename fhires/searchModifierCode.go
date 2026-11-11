package fhires

// A supported modifier for a search parameter.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/search-modifier-code
type SearchModifierCode string

const (
	// The search parameter returns resources that have a value or not.
	SearchModifierCodeMissing SearchModifierCode = "missing"
	// The search parameter returns resources that have a value that exactly matches the supplied parameter (the whole string, including casing and accents).
	SearchModifierCodeExact SearchModifierCode = "exact"
	// The search parameter returns resources that include the supplied parameter value anywhere within the field being searched.
	SearchModifierCodeContains SearchModifierCode = "contains"
	// The search parameter returns resources that do not contain a match.
	SearchModifierCodeNot SearchModifierCode = "not"
	// The search parameter is processed as a string that searches text associated with the code/value - either CodeableConcept.text, Coding.display, Identifier.type.text, or Reference.display.
	SearchModifierCodeText SearchModifierCode = "text"
	// The search parameter is a URI (relative or absolute) that identifies a value set, and the search parameter tests whether the coding is in the specified value set.
	SearchModifierCodeIn SearchModifierCode = "in"
	// The search parameter is a URI (relative or absolute) that identifies a value set, and the search parameter tests whether the coding is not in the specified value set.
	SearchModifierCodeNotIn SearchModifierCode = "not-in"
	// The search parameter tests whether the value in a resource is subsumed by the specified value (is-a, or hierarchical relationships).
	SearchModifierCodeBelow SearchModifierCode = "below"
	// The search parameter tests whether the value in a resource subsumes the specified value (is-a, or hierarchical relationships).
	SearchModifierCodeAbove SearchModifierCode = "above"
	// The search parameter only applies to the Resource Type specified as a modifier (e.g. the modifier is not actually :type, but :Patient etc.).
	SearchModifierCodeType SearchModifierCode = "type"
	// The search parameter applies to the identifier on the resource, not the reference.
	SearchModifierCodeIdentifier SearchModifierCode = "identifier"
	// The search parameter has the format system|code|value, where the system and code refer to an Identifier.type.coding.system and .code, and match if any of the type codes match. All 3 parts must be present.
	SearchModifierCodeOfType SearchModifierCode = "of-type"
	// Tests whether the textual display value in a resource (e.g., CodeableConcept.text, Coding.display, or Reference.display) matches the supplied parameter value.
	SearchModifierCodeCodeText SearchModifierCode = "code-text"
	// Tests whether the value in a resource matches the supplied parameter value using advanced text handling that searches text associated with the code/value - e.g., CodeableConcept.text, Coding.display, or Identifier.type.text.
	SearchModifierCodeTextAdvanced SearchModifierCode = "text-advanced"
	// The search parameter indicates an inclusion directive (_include, _revinclude) that is applied to an included resource instead of the matching resource.
	SearchModifierCodeIterate SearchModifierCode = "iterate"
)

func (code SearchModifierCode) String() string {
	return code.Code()
}

func (code SearchModifierCode) Code() string {
	return string(code)
}

func (code SearchModifierCode) Display() string {
	switch code {
	case SearchModifierCodeMissing:
		return "Missing"
	case SearchModifierCodeExact:
		return "Exact"
	case SearchModifierCodeContains:
		return "Contains"
	case SearchModifierCodeNot:
		return "Not"
	case SearchModifierCodeText:
		return "Text"
	case SearchModifierCodeIn:
		return "In"
	case SearchModifierCodeNotIn:
		return "Not In"
	case SearchModifierCodeBelow:
		return "Below"
	case SearchModifierCodeAbove:
		return "Above"
	case SearchModifierCodeType:
		return "Type"
	case SearchModifierCodeIdentifier:
		return "Identifier"
	case SearchModifierCodeOfType:
		return "Of Type"
	case SearchModifierCodeCodeText:
		return "Code Text"
	case SearchModifierCodeTextAdvanced:
		return "Text Advanced"
	case SearchModifierCodeIterate:
		return "Iterate"
	}
	return "<unknown>"
}

func (code SearchModifierCode) Definition() string {
	switch code {
	case SearchModifierCodeMissing:
		return "The search parameter returns resources that have a value or not."
	case SearchModifierCodeExact:
		return "The search parameter returns resources that have a value that exactly matches the supplied parameter (the whole string, including casing and accents)."
	case SearchModifierCodeContains:
		return "The search parameter returns resources that include the supplied parameter value anywhere within the field being searched."
	case SearchModifierCodeNot:
		return "The search parameter returns resources that do not contain a match."
	case SearchModifierCodeText:
		return "The search parameter is processed as a string that searches text associated with the code/value - either CodeableConcept.text, Coding.display, Identifier.type.text, or Reference.display."
	case SearchModifierCodeIn:
		return "The search parameter is a URI (relative or absolute) that identifies a value set, and the search parameter tests whether the coding is in the specified value set."
	case SearchModifierCodeNotIn:
		return "The search parameter is a URI (relative or absolute) that identifies a value set, and the search parameter tests whether the coding is not in the specified value set."
	case SearchModifierCodeBelow:
		return "The search parameter tests whether the value in a resource is subsumed by the specified value (is-a, or hierarchical relationships)."
	case SearchModifierCodeAbove:
		return "The search parameter tests whether the value in a resource subsumes the specified value (is-a, or hierarchical relationships)."
	case SearchModifierCodeType:
		return "The search parameter only applies to the Resource Type specified as a modifier (e.g. the modifier is not actually :type, but :Patient etc.)."
	case SearchModifierCodeIdentifier:
		return "The search parameter applies to the identifier on the resource, not the reference."
	case SearchModifierCodeOfType:
		return "The search parameter has the format system|code|value, where the system and code refer to an Identifier.type.coding.system and .code, and match if any of the type codes match. All 3 parts must be present."
	case SearchModifierCodeCodeText:
		return "Tests whether the textual display value in a resource (e.g., CodeableConcept.text, Coding.display, or Reference.display) matches the supplied parameter value."
	case SearchModifierCodeTextAdvanced:
		return "Tests whether the value in a resource matches the supplied parameter value using advanced text handling that searches text associated with the code/value - e.g., CodeableConcept.text, Coding.display, or Identifier.type.text."
	case SearchModifierCodeIterate:
		return "The search parameter indicates an inclusion directive (_include, _revinclude) that is applied to an included resource instead of the matching resource."
	}
	return "<unknown>"
}
