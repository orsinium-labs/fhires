package fhires

// Describes the type of a metric calibration.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/metric-calibration-type
type DeviceMetricCalibrationType string

const (
	// Metric calibration method has not been identified.
	DeviceMetricCalibrationTypeUnspecified DeviceMetricCalibrationType = "unspecified"
	// Offset metric calibration method.
	DeviceMetricCalibrationTypeOffset DeviceMetricCalibrationType = "offset"
	// Gain metric calibration method.
	DeviceMetricCalibrationTypeGain DeviceMetricCalibrationType = "gain"
	// Two-point metric calibration method.
	DeviceMetricCalibrationTypeTwoPoint DeviceMetricCalibrationType = "two-point"
)

func (code DeviceMetricCalibrationType) String() string {
	return code.Code()
}

func (code DeviceMetricCalibrationType) Code() string {
	return string(code)
}

func (code DeviceMetricCalibrationType) Display() string {
	switch code {
	case DeviceMetricCalibrationTypeUnspecified:
		return "Unspecified"
	case DeviceMetricCalibrationTypeOffset:
		return "Offset"
	case DeviceMetricCalibrationTypeGain:
		return "Gain"
	case DeviceMetricCalibrationTypeTwoPoint:
		return "Two Point"
	}
	return "<unknown>"
}

func (code DeviceMetricCalibrationType) Definition() string {
	switch code {
	case DeviceMetricCalibrationTypeUnspecified:
		return "Metric calibration method has not been identified."
	case DeviceMetricCalibrationTypeOffset:
		return "Offset metric calibration method."
	case DeviceMetricCalibrationTypeGain:
		return "Gain metric calibration method."
	case DeviceMetricCalibrationTypeTwoPoint:
		return "Two-point metric calibration method."
	}
	return "<unknown>"
}
