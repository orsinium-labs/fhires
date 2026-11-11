package fhires

// Describes the operational status of the DeviceMetric.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/metric-operational-status
type DeviceMetricOperationalStatus string

const (
	// The DeviceMetric is operating and will generate Observations.
	DeviceMetricOperationalStatusOn DeviceMetricOperationalStatus = "on"
	// The DeviceMetric is not operating.
	DeviceMetricOperationalStatusOff DeviceMetricOperationalStatus = "off"
	// The DeviceMetric is operating, but will not generate any Observations.
	DeviceMetricOperationalStatusStandby DeviceMetricOperationalStatus = "standby"
	// The DeviceMetric was entered in error.
	DeviceMetricOperationalStatusEnteredInError DeviceMetricOperationalStatus = "entered-in-error"
)

func (code DeviceMetricOperationalStatus) String() string {
	return code.Code()
}

func (code DeviceMetricOperationalStatus) Code() string {
	return string(code)
}

func (code DeviceMetricOperationalStatus) Display() string {
	switch code {
	case DeviceMetricOperationalStatusOn:
		return "On"
	case DeviceMetricOperationalStatusOff:
		return "Off"
	case DeviceMetricOperationalStatusStandby:
		return "Standby"
	case DeviceMetricOperationalStatusEnteredInError:
		return "Entered In Error"
	}
	return "<unknown>"
}

func (code DeviceMetricOperationalStatus) Definition() string {
	switch code {
	case DeviceMetricOperationalStatusOn:
		return "The DeviceMetric is operating and will generate Observations."
	case DeviceMetricOperationalStatusOff:
		return "The DeviceMetric is not operating."
	case DeviceMetricOperationalStatusStandby:
		return "The DeviceMetric is operating, but will not generate any Observations."
	case DeviceMetricOperationalStatusEnteredInError:
		return "The DeviceMetric was entered in error."
	}
	return "<unknown>"
}
