package fhires

// The severity of the audit entry.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/audit-event-severity
type AuditEventSeverity string

const (
	// System is unusable. e.g., This level should only be reported by infrastructure and should not be used by applications.
	AuditEventSeverityEmergency AuditEventSeverity = "emergency"
	// Notification should be sent to trigger action be taken. e.g., Loss of the primary network connection needing attention.
	AuditEventSeverityAlert AuditEventSeverity = "alert"
	// Critical conditions. e.g., A failure in the system's primary application that will reset automatically.
	AuditEventSeverityCritical AuditEventSeverity = "critical"
	// Error conditions. e.g., An application has exceeded its file storage limit and attempts to write are failing.
	AuditEventSeverityError AuditEventSeverity = "error"
	// Warning conditions. May indicate that an error will occur if action is not taken. e.g., A non-root file system has only 2GB remaining.
	AuditEventSeverityWarning AuditEventSeverity = "warning"
	// Notice messages. Normal but significant condition. Events that are unusual, but not error conditions.
	AuditEventSeverityNotice AuditEventSeverity = "notice"
	// Normal operational messages that require no action. e.g., An application has started, paused, or ended successfully.
	AuditEventSeverityInformational AuditEventSeverity = "informational"
	// Debug-level messages. Information useful to developers for debugging the application.
	AuditEventSeverityDebug AuditEventSeverity = "debug"
)

func (code AuditEventSeverity) String() string {
	return code.Code()
}

func (code AuditEventSeverity) Code() string {
	return string(code)
}

func (code AuditEventSeverity) Display() string {
	switch code {
	case AuditEventSeverityEmergency:
		return "Emergency"
	case AuditEventSeverityAlert:
		return "Alert"
	case AuditEventSeverityCritical:
		return "Critical"
	case AuditEventSeverityError:
		return "Error"
	case AuditEventSeverityWarning:
		return "Warning"
	case AuditEventSeverityNotice:
		return "Notice"
	case AuditEventSeverityInformational:
		return "Informational"
	case AuditEventSeverityDebug:
		return "Debug"
	}
	return "<unknown>"
}

func (code AuditEventSeverity) Definition() string {
	switch code {
	case AuditEventSeverityEmergency:
		return "System is unusable. e.g., This level should only be reported by infrastructure and should not be used by applications."
	case AuditEventSeverityAlert:
		return "Notification should be sent to trigger action be taken. e.g., Loss of the primary network connection needing attention."
	case AuditEventSeverityCritical:
		return "Critical conditions. e.g., A failure in the system's primary application that will reset automatically."
	case AuditEventSeverityError:
		return "Error conditions. e.g., An application has exceeded its file storage limit and attempts to write are failing. "
	case AuditEventSeverityWarning:
		return "Warning conditions. May indicate that an error will occur if action is not taken. e.g., A non-root file system has only 2GB remaining."
	case AuditEventSeverityNotice:
		return "Notice messages. Normal but significant condition. Events that are unusual, but not error conditions."
	case AuditEventSeverityInformational:
		return "Normal operational messages that require no action. e.g., An application has started, paused, or ended successfully."
	case AuditEventSeverityDebug:
		return "Debug-level messages. Information useful to developers for debugging the application."
	}
	return "<unknown>"
}
