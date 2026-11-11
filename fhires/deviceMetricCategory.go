package fhires

// Describes the category of the metric.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/metric-category
type DeviceMetricCategory string

const (
	// Observations generated for this DeviceMetric are measured.
	DeviceMetricCategoryMeasurement DeviceMetricCategory = "measurement"
	// Observations generated for this DeviceMetric is a setting that will influence the behavior of the Device.
	DeviceMetricCategorySetting DeviceMetricCategory = "setting"
	// Observations generated for this DeviceMetric are calculated.
	DeviceMetricCategoryCalculation DeviceMetricCategory = "calculation"
	// The category of this DeviceMetric is unspecified.
	DeviceMetricCategoryUnspecified DeviceMetricCategory = "unspecified"
)

func (code DeviceMetricCategory) String() string {
	return code.Code()
}

func (code DeviceMetricCategory) Code() string {
	return string(code)
}

func (code DeviceMetricCategory) Display() string {
	switch code {
	case DeviceMetricCategoryMeasurement:
		return "Measurement"
	case DeviceMetricCategorySetting:
		return "Setting"
	case DeviceMetricCategoryCalculation:
		return "Calculation"
	case DeviceMetricCategoryUnspecified:
		return "Unspecified"
	}
	return "<unknown>"
}

func (code DeviceMetricCategory) Definition() string {
	switch code {
	case DeviceMetricCategoryMeasurement:
		return "Observations generated for this DeviceMetric are measured."
	case DeviceMetricCategorySetting:
		return "Observations generated for this DeviceMetric is a setting that will influence the behavior of the Device."
	case DeviceMetricCategoryCalculation:
		return "Observations generated for this DeviceMetric are calculated."
	case DeviceMetricCategoryUnspecified:
		return "The category of this DeviceMetric is unspecified."
	}
	return "<unknown>"
}
