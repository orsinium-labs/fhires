package fhires

// The status of the InventoryReport.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/inventoryreport-status
type InventoryReportStatus string

const (
	// The existence of the report is registered, but it is still without content or only some preliminary content.
	InventoryReportStatusDraft InventoryReportStatus = "draft"
	// The inventory report has been requested but there is no data available.
	InventoryReportStatusRequested InventoryReportStatus = "requested"
	// This report is submitted as current.
	InventoryReportStatusActive InventoryReportStatus = "active"
	// The report has been withdrawn following a previous final release.  This electronic record should never have existed, though it is possible that real-world decisions were based on it.
	InventoryReportStatusEnteredInError InventoryReportStatus = "entered-in-error"
)

func (code InventoryReportStatus) String() string {
	return code.Code()
}

func (code InventoryReportStatus) Code() string {
	return string(code)
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
