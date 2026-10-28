package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// The type of count.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/inventoryreport-counttype
type InventoryCountType int

const (
	InventoryCountTypeSnapshot InventoryCountType = iota
	// The inventory report is about the difference between a previous count and a current count, i.e. it represents the items that have been added/subtracted from inventory.
	InventoryCountTypeDifference
)

func (code InventoryCountType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *InventoryCountType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "snapshot":
		*code = InventoryCountTypeSnapshot
	case "difference":
		*code = InventoryCountTypeDifference
	default:
		return fmt.Errorf("unknown InventoryCountType code `%s`", s)
	}
	return nil
}

func (code InventoryCountType) String() string {
	return code.Code()
}

func (code InventoryCountType) Code() string {
	switch code {
	case InventoryCountTypeSnapshot:
		return "snapshot"
	case InventoryCountTypeDifference:
		return "difference"
	}
	return "<unknown>"
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
