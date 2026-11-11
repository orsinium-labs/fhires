package fhires

// The type of trigger.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/trigger-type
type TriggerType string

const (
	// The trigger occurs in response to a specific named event, and no other information about the trigger is specified. Named events are completely pre-coordinated, and the formal semantics of the trigger are not provided.
	TriggerTypeNamedEvent TriggerType = "named-event"
	// The trigger occurs at a specific time or periodically as described by a timing or schedule. A periodic event cannot have any data elements, but may have a name assigned as a shorthand for the event.
	TriggerTypePeriodic TriggerType = "periodic"
	// The trigger occurs whenever data of a particular type is changed in any way, either added, modified, or removed.
	TriggerTypeDataChanged TriggerType = "data-changed"
	// The trigger occurs whenever data of a particular type is added.
	TriggerTypeDataAdded TriggerType = "data-added"
	// The trigger occurs whenever data of a particular type is modified.
	TriggerTypeDataModified TriggerType = "data-modified"
	// The trigger occurs whenever data of a particular type is removed.
	TriggerTypeDataRemoved TriggerType = "data-removed"
	// The trigger occurs whenever data of a particular type is accessed.
	TriggerTypeDataAccessed TriggerType = "data-accessed"
	// The trigger occurs whenever access to data of a particular type is completed.
	TriggerTypeDataAccessEnded TriggerType = "data-access-ended"
)

func (code TriggerType) String() string {
	return code.Code()
}

func (code TriggerType) Code() string {
	return string(code)
}

func (code TriggerType) Display() string {
	switch code {
	case TriggerTypeNamedEvent:
		return "Named Event"
	case TriggerTypePeriodic:
		return "Periodic"
	case TriggerTypeDataChanged:
		return "Data Changed"
	case TriggerTypeDataAdded:
		return "Data Added"
	case TriggerTypeDataModified:
		return "Data Updated"
	case TriggerTypeDataRemoved:
		return "Data Removed"
	case TriggerTypeDataAccessed:
		return "Data Accessed"
	case TriggerTypeDataAccessEnded:
		return "Data Access Ended"
	}
	return "<unknown>"
}

func (code TriggerType) Definition() string {
	switch code {
	case TriggerTypeNamedEvent:
		return "The trigger occurs in response to a specific named event, and no other information about the trigger is specified. Named events are completely pre-coordinated, and the formal semantics of the trigger are not provided."
	case TriggerTypePeriodic:
		return "The trigger occurs at a specific time or periodically as described by a timing or schedule. A periodic event cannot have any data elements, but may have a name assigned as a shorthand for the event."
	case TriggerTypeDataChanged:
		return "The trigger occurs whenever data of a particular type is changed in any way, either added, modified, or removed."
	case TriggerTypeDataAdded:
		return "The trigger occurs whenever data of a particular type is added."
	case TriggerTypeDataModified:
		return "The trigger occurs whenever data of a particular type is modified."
	case TriggerTypeDataRemoved:
		return "The trigger occurs whenever data of a particular type is removed."
	case TriggerTypeDataAccessed:
		return "The trigger occurs whenever data of a particular type is accessed."
	case TriggerTypeDataAccessEnded:
		return "The trigger occurs whenever access to data of a particular type is completed."
	}
	return "<unknown>"
}
