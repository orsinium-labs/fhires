package fhires

// The use of an address.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/address-use
type AddressUse string

const (
	// A communication address at a home.
	AddressUseHome AddressUse = "home"
	// An office address. First choice for business related contacts during business hours.
	AddressUseWork AddressUse = "work"
	// A temporary address. The period can provide more detailed information.
	AddressUseTemp AddressUse = "temp"
	// This address is no longer in use (or was never correct but retained for records).
	AddressUseOld AddressUse = "old"
	// An address to be used to send bills, invoices, receipts etc.
	AddressUseBilling AddressUse = "billing"
)

func (code AddressUse) String() string {
	return code.Code()
}

func (code AddressUse) Code() string {
	return string(code)
}

func (code AddressUse) Display() string {
	switch code {
	case AddressUseHome:
		return "Home"
	case AddressUseWork:
		return "Work"
	case AddressUseTemp:
		return "Temporary"
	case AddressUseOld:
		return "Old / Incorrect"
	case AddressUseBilling:
		return "Billing"
	}
	return "<unknown>"
}

func (code AddressUse) Definition() string {
	switch code {
	case AddressUseHome:
		return "A communication address at a home."
	case AddressUseWork:
		return "An office address. First choice for business related contacts during business hours."
	case AddressUseTemp:
		return "A temporary address. The period can provide more detailed information."
	case AddressUseOld:
		return "This address is no longer in use (or was never correct but retained for records)."
	case AddressUseBilling:
		return "An address to be used to send bills, invoices, receipts etc."
	}
	return "<unknown>"
}
