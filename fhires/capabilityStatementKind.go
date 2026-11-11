package fhires

// How a capability statement is intended to be used.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/capability-statement-kind
type CapabilityStatementKind string

const (
	// The CapabilityStatement instance represents the present capabilities of a specific system instance.  This is the kind returned by /metadata for a FHIR server end-point.
	CapabilityStatementKindInstance CapabilityStatementKind = "instance"
	// The CapabilityStatement instance represents the capabilities of a system or piece of software, independent of a particular installation.
	CapabilityStatementKindCapability CapabilityStatementKind = "capability"
	// The CapabilityStatement instance represents a set of requirements for other systems to meet; e.g. as part of an implementation guide or 'request for proposal'.
	CapabilityStatementKindRequirements CapabilityStatementKind = "requirements"
)

func (code CapabilityStatementKind) String() string {
	return code.Code()
}

func (code CapabilityStatementKind) Code() string {
	return string(code)
}

func (code CapabilityStatementKind) Display() string {
	switch code {
	case CapabilityStatementKindInstance:
		return "Instance"
	case CapabilityStatementKindCapability:
		return "Capability"
	case CapabilityStatementKindRequirements:
		return "Requirements"
	}
	return "<unknown>"
}

func (code CapabilityStatementKind) Definition() string {
	switch code {
	case CapabilityStatementKindInstance:
		return "The CapabilityStatement instance represents the present capabilities of a specific system instance.  This is the kind returned by /metadata for a FHIR server end-point."
	case CapabilityStatementKindCapability:
		return "The CapabilityStatement instance represents the capabilities of a system or piece of software, independent of a particular installation."
	case CapabilityStatementKindRequirements:
		return "The CapabilityStatement instance represents a set of requirements for other systems to meet; e.g. as part of an implementation guide or 'request for proposal'."
	}
	return "<unknown>"
}
