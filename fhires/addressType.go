package fhires

// The type of an address (physical / postal).
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/address-type
type AddressType string

const (
	// Mailing addresses - PO Boxes and care-of addresses.
	AddressTypePostal AddressType = "postal"
	// A physical address that can be visited.
	AddressTypePhysical AddressType = "physical"
	// An address that is both physical and postal.
	AddressTypeBoth AddressType = "both"
)

func (code AddressType) String() string {
	return code.Code()
}

func (code AddressType) Code() string {
	return string(code)
}

func (code AddressType) Display() string {
	switch code {
	case AddressTypePostal:
		return "Postal"
	case AddressTypePhysical:
		return "Physical"
	case AddressTypeBoth:
		return "Postal & Physical"
	}
	return "<unknown>"
}

func (code AddressType) Definition() string {
	switch code {
	case AddressTypePostal:
		return "Mailing addresses - PO Boxes and care-of addresses."
	case AddressTypePhysical:
		return "A physical address that can be visited."
	case AddressTypeBoth:
		return "An address that is both physical and postal."
	}
	return "<unknown>"
}
