package fhires

// The type of count.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/inventoryreport-counttype
type InventoryCountType string

const (
	// The inventory report is a current absolute snapshot, i.e. it represents the quantities at hand.
	InventoryCountTypeSnapshot InventoryCountType = "snapshot"
	// The inventory report is about the difference between a previous count and a current count, i.e. it represents the items that have been added/subtracted from inventory.
	InventoryCountTypeDifference InventoryCountType = "difference"
)

func (code InventoryCountType) String() string {
	return code.Code()
}

func (code InventoryCountType) Code() string {
	return string(code)
}

func (code InventoryCountType) Display() string {
	switch code {
	case InventoryCountTypeSnapshot:
		return "Snapshot"
	case InventoryCountTypeDifference:
		return "Difference"
	}
	return "<unknown>"
}

func (code InventoryCountType) Definition() string {
	switch code {
	case InventoryCountTypeSnapshot:
		return "The inventory report is a current absolute snapshot, i.e. it represents the quantities at hand."
	case InventoryCountTypeDifference:
		return "The inventory report is about the difference between a previous count and a current count, i.e. it represents the items that have been added/subtracted from inventory."
	}
	return "<unknown>"
}
