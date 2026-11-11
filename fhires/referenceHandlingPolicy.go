package fhires

// A set of flags that defines how references are supported.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/reference-handling-policy
type ReferenceHandlingPolicy string

const (
	// The server supports and populates Literal references (i.e. using Reference.reference) where they are known (this code does not guarantee that all references are literal; see 'enforced').
	ReferenceHandlingPolicyLiteral ReferenceHandlingPolicy = "literal"
	// The server allows logical references (i.e. using Reference.identifier).
	ReferenceHandlingPolicyLogical ReferenceHandlingPolicy = "logical"
	// The server will attempt to resolve logical references to literal references - i.e. converting Reference.identifier to Reference.reference (if resolution fails, the server may still accept resources; see logical).
	ReferenceHandlingPolicyResolves ReferenceHandlingPolicy = "resolves"
	// The server enforces that references have integrity - e.g. it ensures that references can always be resolved. This is typically the case for clinical record systems, but often not the case for middleware/proxy systems.
	ReferenceHandlingPolicyEnforced ReferenceHandlingPolicy = "enforced"
	// The server does not support references that point to other servers.
	ReferenceHandlingPolicyLocal ReferenceHandlingPolicy = "local"
)

func (code ReferenceHandlingPolicy) String() string {
	return code.Code()
}

func (code ReferenceHandlingPolicy) Code() string {
	return string(code)
}

func (code ReferenceHandlingPolicy) Display() string {
	switch code {
	case ReferenceHandlingPolicyLiteral:
		return "Literal References"
	case ReferenceHandlingPolicyLogical:
		return "Logical References"
	case ReferenceHandlingPolicyResolves:
		return "Resolves References"
	case ReferenceHandlingPolicyEnforced:
		return "Reference Integrity Enforced"
	case ReferenceHandlingPolicyLocal:
		return "Local References Only"
	}
	return "<unknown>"
}

func (code ReferenceHandlingPolicy) Definition() string {
	switch code {
	case ReferenceHandlingPolicyLiteral:
		return "The server supports and populates Literal references (i.e. using Reference.reference) where they are known (this code does not guarantee that all references are literal; see 'enforced')."
	case ReferenceHandlingPolicyLogical:
		return "The server allows logical references (i.e. using Reference.identifier)."
	case ReferenceHandlingPolicyResolves:
		return "The server will attempt to resolve logical references to literal references - i.e. converting Reference.identifier to Reference.reference (if resolution fails, the server may still accept resources; see logical)."
	case ReferenceHandlingPolicyEnforced:
		return "The server enforces that references have integrity - e.g. it ensures that references can always be resolved. This is typically the case for clinical record systems, but often not the case for middleware/proxy systems."
	case ReferenceHandlingPolicyLocal:
		return "The server does not support references that point to other servers."
	}
	return "<unknown>"
}
