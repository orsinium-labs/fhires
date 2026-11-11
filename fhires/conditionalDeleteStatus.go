package fhires

// A code that indicates how the server supports conditional delete.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/conditional-delete-status
type ConditionalDeleteStatus string

const (
	// No support for conditional deletes.
	ConditionalDeleteStatusNotSupported ConditionalDeleteStatus = "not-supported"
	// Conditional deletes are supported, but only single resources at a time.
	ConditionalDeleteStatusSingle ConditionalDeleteStatus = "single"
	// Conditional deletes are supported, and multiple resources can be deleted in a single interaction.
	ConditionalDeleteStatusMultiple ConditionalDeleteStatus = "multiple"
)

func (code ConditionalDeleteStatus) String() string {
	return code.Code()
}

func (code ConditionalDeleteStatus) Code() string {
	return string(code)
}

func (code ConditionalDeleteStatus) Display() string {
	switch code {
	case ConditionalDeleteStatusNotSupported:
		return "Not Supported"
	case ConditionalDeleteStatusSingle:
		return "Single Deletes Supported"
	case ConditionalDeleteStatusMultiple:
		return "Multiple Deletes Supported"
	}
	return "<unknown>"
}

func (code ConditionalDeleteStatus) Definition() string {
	switch code {
	case ConditionalDeleteStatusNotSupported:
		return "No support for conditional deletes."
	case ConditionalDeleteStatusSingle:
		return "Conditional deletes are supported, but only single resources at a time."
	case ConditionalDeleteStatusMultiple:
		return "Conditional deletes are supported, and multiple resources can be deleted in a single interaction."
	}
	return "<unknown>"
}
