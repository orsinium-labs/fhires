package fhires

// The status of the endpoint.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/endpoint-status
type EndpointStatus string

const (
	// This endpoint is expected to be active and can be used.
	EndpointStatusActive EndpointStatus = "active"
	// This endpoint is temporarily unavailable.
	EndpointStatusSuspended EndpointStatus = "suspended"
	// This endpoint has exceeded connectivity thresholds and is considered in an error state and should no longer be attempted to connect to until corrective action is taken.
	EndpointStatusError EndpointStatus = "error"
	// This endpoint is no longer to be used.
	EndpointStatusOff EndpointStatus = "off"
	// This instance should not have been part of this patient's medical record.
	EndpointStatusEnteredInError EndpointStatus = "entered-in-error"
)

func (code EndpointStatus) String() string {
	return code.Code()
}

func (code EndpointStatus) Code() string {
	return string(code)
}

func (code EndpointStatus) Display() string {
	switch code {
	case EndpointStatusActive:
		return "Active"
	case EndpointStatusSuspended:
		return "Suspended"
	case EndpointStatusError:
		return "Error"
	case EndpointStatusOff:
		return "Off"
	case EndpointStatusEnteredInError:
		return "Entered in error"
	}
	return "<unknown>"
}

func (code EndpointStatus) Definition() string {
	switch code {
	case EndpointStatusActive:
		return "This endpoint is expected to be active and can be used."
	case EndpointStatusSuspended:
		return "This endpoint is temporarily unavailable."
	case EndpointStatusError:
		return "This endpoint has exceeded connectivity thresholds and is considered in an error state and should no longer be attempted to connect to until corrective action is taken."
	case EndpointStatusOff:
		return "This endpoint is no longer to be used."
	case EndpointStatusEnteredInError:
		return "This instance should not have been part of this patient's medical record."
	}
	return "<unknown>"
}
