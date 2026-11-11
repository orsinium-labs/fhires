package fhires

// Codes identifying the lifecycle stage of an Invoice.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/invoice-status
type InvoiceStatus string

const (
	// the invoice has been prepared but not yet finalized.
	InvoiceStatusDraft InvoiceStatus = "draft"
	// the invoice has been finalized and sent to the recipient.
	InvoiceStatusIssued InvoiceStatus = "issued"
	// the invoice has been balaced / completely paid.
	InvoiceStatusBalanced InvoiceStatus = "balanced"
	// the invoice was cancelled.
	InvoiceStatusCancelled InvoiceStatus = "cancelled"
	// the invoice was determined as entered in error before it was issued.
	InvoiceStatusEnteredInError InvoiceStatus = "entered-in-error"
)

func (code InvoiceStatus) String() string {
	return code.Code()
}

func (code InvoiceStatus) Code() string {
	return string(code)
}

func (code InvoiceStatus) Display() string {
	switch code {
	case InvoiceStatusDraft:
		return "draft"
	case InvoiceStatusIssued:
		return "issued"
	case InvoiceStatusBalanced:
		return "balanced"
	case InvoiceStatusCancelled:
		return "cancelled"
	case InvoiceStatusEnteredInError:
		return "entered in error"
	}
	return "<unknown>"
}

func (code InvoiceStatus) Definition() string {
	switch code {
	case InvoiceStatusDraft:
		return "the invoice has been prepared but not yet finalized."
	case InvoiceStatusIssued:
		return "the invoice has been finalized and sent to the recipient."
	case InvoiceStatusBalanced:
		return "the invoice has been balaced / completely paid."
	case InvoiceStatusCancelled:
		return "the invoice was cancelled."
	case InvoiceStatusEnteredInError:
		return "the invoice was determined as entered in error before it was issued."
	}
	return "<unknown>"
}
