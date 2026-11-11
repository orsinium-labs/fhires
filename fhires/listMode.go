package fhires

// The processing mode that applies to this list.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/list-mode
type ListMode string

const (
	// This list is the master list, maintained in an ongoing fashion with regular updates as the real-world list it is tracking changes.
	ListModeWorking ListMode = "working"
	// This list was prepared as a snapshot. It should not be assumed to be current.
	ListModeSnapshot ListMode = "snapshot"
	// A point-in-time list that shows what changes have been made or recommended.  E.g. a discharge medication list showing what was added and removed during an encounter.
	ListModeChanges ListMode = "changes"
)

func (code ListMode) String() string {
	return code.Code()
}

func (code ListMode) Code() string {
	return string(code)
}

func (code ListMode) Display() string {
	switch code {
	case ListModeWorking:
		return "Working List"
	case ListModeSnapshot:
		return "Snapshot List"
	case ListModeChanges:
		return "Change List"
	}
	return "<unknown>"
}

func (code ListMode) Definition() string {
	switch code {
	case ListModeWorking:
		return "This list is the master list, maintained in an ongoing fashion with regular updates as the real-world list it is tracking changes."
	case ListModeSnapshot:
		return "This list was prepared as a snapshot. It should not be assumed to be current."
	case ListModeChanges:
		return "A point-in-time list that shows what changes have been made or recommended.  E.g. a discharge medication list showing what was added and removed during an encounter."
	}
	return "<unknown>"
}
