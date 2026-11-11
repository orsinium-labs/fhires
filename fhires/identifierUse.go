package fhires

// Identifies the purpose for this identifier, if known .
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/identifier-use
type IdentifierUse string

const (
	// The identifier recommended for display and use in real-world interactions which should be used when such identifier is different from the "official" identifier.
	IdentifierUseUsual IdentifierUse = "usual"
	// The identifier considered to be most trusted for the identification of this item. Sometimes also known as "primary" and "main". The determination of "official" is subjective and implementation guides often provide additional guidelines for use.
	IdentifierUseOfficial IdentifierUse = "official"
	// A temporary identifier.
	IdentifierUseTemp IdentifierUse = "temp"
	// An identifier that was assigned in secondary use - it serves to identify the object in a relative context, but cannot be consistently assigned to the same object again in a different context.
	IdentifierUseSecondary IdentifierUse = "secondary"
	// The identifier id no longer considered valid, but may be relevant for search purposes.  E.g. Changes to identifier schemes, account merges, etc.
	IdentifierUseOld IdentifierUse = "old"
)

func (code IdentifierUse) String() string {
	return code.Code()
}

func (code IdentifierUse) Code() string {
	return string(code)
}

func (code IdentifierUse) Display() string {
	switch code {
	case IdentifierUseUsual:
		return "Usual"
	case IdentifierUseOfficial:
		return "Official"
	case IdentifierUseTemp:
		return "Temp"
	case IdentifierUseSecondary:
		return "Secondary"
	case IdentifierUseOld:
		return "Old"
	}
	return "<unknown>"
}

func (code IdentifierUse) Definition() string {
	switch code {
	case IdentifierUseUsual:
		return "The identifier recommended for display and use in real-world interactions which should be used when such identifier is different from the \"official\" identifier."
	case IdentifierUseOfficial:
		return "The identifier considered to be most trusted for the identification of this item. Sometimes also known as \"primary\" and \"main\". The determination of \"official\" is subjective and implementation guides often provide additional guidelines for use."
	case IdentifierUseTemp:
		return "A temporary identifier."
	case IdentifierUseSecondary:
		return "An identifier that was assigned in secondary use - it serves to identify the object in a relative context, but cannot be consistently assigned to the same object again in a different context."
	case IdentifierUseOld:
		return "The identifier id no longer considered valid, but may be relevant for search purposes.  E.g. Changes to identifier schemes, account merges, etc."
	}
	return "<unknown>"
}
