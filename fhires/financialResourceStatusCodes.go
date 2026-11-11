package fhires

// This value set includes Status codes.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/fm-status
type FinancialResourceStatusCodes string

const (
	// The instance is currently in-force.
	FinancialResourceStatusCodesActive FinancialResourceStatusCodes = "active"
	// The instance is withdrawn, rescinded or reversed.
	FinancialResourceStatusCodesCancelled FinancialResourceStatusCodes = "cancelled"
	// A new instance the contents of which is not complete.
	FinancialResourceStatusCodesDraft FinancialResourceStatusCodes = "draft"
	// The instance was entered in error.
	FinancialResourceStatusCodesEnteredInError FinancialResourceStatusCodes = "entered-in-error"
)

func (code FinancialResourceStatusCodes) String() string {
	return code.Code()
}

func (code FinancialResourceStatusCodes) Code() string {
	return string(code)
}

func (code FinancialResourceStatusCodes) Display() string {
	switch code {
	case FinancialResourceStatusCodesActive:
		return "Active"
	case FinancialResourceStatusCodesCancelled:
		return "Cancelled"
	case FinancialResourceStatusCodesDraft:
		return "Draft"
	case FinancialResourceStatusCodesEnteredInError:
		return "Entered in Error"
	}
	return "<unknown>"
}

func (code FinancialResourceStatusCodes) Definition() string {
	switch code {
	case FinancialResourceStatusCodesActive:
		return "The instance is currently in-force."
	case FinancialResourceStatusCodesCancelled:
		return "The instance is withdrawn, rescinded or reversed."
	case FinancialResourceStatusCodesDraft:
		return "A new instance the contents of which is not complete."
	case FinancialResourceStatusCodesEnteredInError:
		return "The instance was entered in error."
	}
	return "<unknown>"
}
