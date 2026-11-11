package fhires

// Indicator for type of action performed during the event that generated the event.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/audit-event-action
type AuditEventAction string

const (
	// Create a new database object, such as placing an order.
	AuditEventActionC AuditEventAction = "C"
	// Read data, such as to print or display to a doctor.
	AuditEventActionR AuditEventAction = "R"
	// Update data, such as revise patient information.
	AuditEventActionU AuditEventAction = "U"
	// Delete items, such as a doctor master file record.
	AuditEventActionD AuditEventAction = "D"
	// Perform a system or application function such as log-on, program execution or use of an object's method, or perform a query/search operation.
	AuditEventActionE AuditEventAction = "E"
)

func (code AuditEventAction) String() string {
	return code.Code()
}

func (code AuditEventAction) Code() string {
	return string(code)
}

func (code AuditEventAction) Display() string {
	switch code {
	case AuditEventActionC:
		return "Create"
	case AuditEventActionR:
		return "Read"
	case AuditEventActionU:
		return "Update"
	case AuditEventActionD:
		return "Delete"
	case AuditEventActionE:
		return "Execute"
	}
	return "<unknown>"
}

func (code AuditEventAction) Definition() string {
	switch code {
	case AuditEventActionC:
		return "Create a new database object, such as placing an order."
	case AuditEventActionR:
		return "Read data, such as to print or display to a doctor."
	case AuditEventActionU:
		return "Update data, such as revise patient information."
	case AuditEventActionD:
		return "Delete items, such as a doctor master file record."
	case AuditEventActionE:
		return "Perform a system or application function such as log-on, program execution or use of an object's method, or perform a query/search operation."
	}
	return "<unknown>"
}
