package fhires

// Describes the state of a metric calibration.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/metric-calibration-state
type DeviceMetricCalibrationState string

const (
	// The metric has not been calibrated.
	DeviceMetricCalibrationStateNotCalibrated DeviceMetricCalibrationState = "not-calibrated"
	// The metric needs to be calibrated.
	DeviceMetricCalibrationStateCalibrationRequired DeviceMetricCalibrationState = "calibration-required"
	// The metric has been calibrated.
	DeviceMetricCalibrationStateCalibrated DeviceMetricCalibrationState = "calibrated"
	// The state of calibration of this metric is unspecified.
	DeviceMetricCalibrationStateUnspecified DeviceMetricCalibrationState = "unspecified"
)

func (code DeviceMetricCalibrationState) String() string {
	return code.Code()
}

func (code DeviceMetricCalibrationState) Code() string {
	return string(code)
}

func (code DeviceMetricCalibrationState) Display() string {
	switch code {
	case DeviceMetricCalibrationStateNotCalibrated:
		return "Not Calibrated"
	case DeviceMetricCalibrationStateCalibrationRequired:
		return "Calibration Required"
	case DeviceMetricCalibrationStateCalibrated:
		return "Calibrated"
	case DeviceMetricCalibrationStateUnspecified:
		return "Unspecified"
	}
	return "<unknown>"
}

func (code DeviceMetricCalibrationState) Definition() string {
	switch code {
	case DeviceMetricCalibrationStateNotCalibrated:
		return "The metric has not been calibrated."
	case DeviceMetricCalibrationStateCalibrationRequired:
		return "The metric needs to be calibrated."
	case DeviceMetricCalibrationStateCalibrated:
		return "The metric has been calibrated."
	case DeviceMetricCalibrationStateUnspecified:
		return "The state of calibration of this metric is unspecified."
	}
	return "<unknown>"
}
