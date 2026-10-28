package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// InventoryReportStatus is documented here http://hl7.org/fhir/ValueSet/inventoryreport-status
type InventoryReportStatus int

const (
	InventoryReportStatusDraft InventoryReportStatus = iota
	InventoryReportStatusRequested
	InventoryReportStatusActive
	InventoryReportStatusEnteredInError
)

func (code InventoryReportStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *InventoryReportStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "draft":
		*code = InventoryReportStatusDraft
	case "requested":
		*code = InventoryReportStatusRequested
	case "active":
		*code = InventoryReportStatusActive
	case "entered-in-error":
		*code = InventoryReportStatusEnteredInError
	default:
		return fmt.Errorf("unknown InventoryReportStatus code `%s`", s)
	}
	return nil
}

func (code InventoryReportStatus) String() string {
	return code.Code()
}

func (code InventoryReportStatus) Code() string {
	switch code {
	case InventoryReportStatusDraft:
		return "draft"
	case InventoryReportStatusRequested:
		return "requested"
	case InventoryReportStatusActive:
		return "active"
	case InventoryReportStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}

func (code InventoryReportStatus) Display() string {
	switch code {
	case InventoryReportStatusDraft:
		return "Draft"
	case InventoryReportStatusRequested:
		return "Requested"
	case InventoryReportStatusActive:
		return "Active"
	case InventoryReportStatusEnteredInError:
		return "Entered in Error"
	}
	return "<unknown>"
}

func (code InventoryReportStatus) Definition() string {
	switch code {
	case InventoryReportStatusDraft:
		return "The existence of the report is registered, but it is still without content or only some preliminary content."
	case InventoryReportStatusRequested:
		return "The inventory report has been requested but there is no data available."
	case InventoryReportStatusActive:
		return "This report is submitted as current."
	case InventoryReportStatusEnteredInError:
		return "The report has been withdrawn following a previous final release.  This electronic record should never have existed, though it is possible that real-world decisions were based on it."
	}
	return "<unknown>"
}
