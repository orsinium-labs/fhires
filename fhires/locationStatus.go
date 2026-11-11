package fhires

// Indicates whether the location is still in use.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/location-status
type LocationStatus string

const (
	// The location is operational.
	LocationStatusActive LocationStatus = "active"
	// The location is temporarily closed.
	LocationStatusSuspended LocationStatus = "suspended"
	// The location is no longer used.
	LocationStatusInactive LocationStatus = "inactive"
)

func (code LocationStatus) String() string {
	return code.Code()
}

func (code LocationStatus) Code() string {
	return string(code)
}

func (code LocationStatus) Display() string {
	switch code {
	case LocationStatusActive:
		return "Active"
	case LocationStatusSuspended:
		return "Suspended"
	case LocationStatusInactive:
		return "Inactive"
	}
	return "<unknown>"
}

func (code LocationStatus) Definition() string {
	switch code {
	case LocationStatusActive:
		return "The location is operational."
	case LocationStatusSuspended:
		return "The location is temporarily closed."
	case LocationStatusInactive:
		return "The location is no longer used."
	}
	return "<unknown>"
}
