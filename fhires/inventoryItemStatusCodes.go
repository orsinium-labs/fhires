package fhires

// InventoryItem Status Codes
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/inventoryitem-status
type InventoryItemStatusCodes string

const (
	// The item is active and can be referenced.
	InventoryItemStatusCodesActive InventoryItemStatusCodes = "active"
	// The item is presently inactive - there may be references to it but the item is not expected to be used.
	InventoryItemStatusCodesInactive InventoryItemStatusCodes = "inactive"
	// The item record was entered in error.
	InventoryItemStatusCodesEnteredInError InventoryItemStatusCodes = "entered-in-error"
	// The item status has not been determined.
	InventoryItemStatusCodesUnknown InventoryItemStatusCodes = "unknown"
)

func (code InventoryItemStatusCodes) String() string {
	return code.Code()
}

func (code InventoryItemStatusCodes) Code() string {
	return string(code)
}

func (code InventoryItemStatusCodes) Display() string {
	switch code {
	case InventoryItemStatusCodesActive:
		return "Active"
	case InventoryItemStatusCodesInactive:
		return "Inactive"
	case InventoryItemStatusCodesEnteredInError:
		return "Entered in Error"
	case InventoryItemStatusCodesUnknown:
		return "Unknown"
	}
	return "<unknown>"
}

func (code InventoryItemStatusCodes) Definition() string {
	switch code {
	case InventoryItemStatusCodesActive:
		return "The item is active and can be referenced."
	case InventoryItemStatusCodesInactive:
		return "The item is presently inactive - there may be references to it but the item is not expected to be used."
	case InventoryItemStatusCodesEnteredInError:
		return "The item record was entered in error."
	case InventoryItemStatusCodesUnknown:
		return "The item status has not been determined."
	}
	return "<unknown>"
}
