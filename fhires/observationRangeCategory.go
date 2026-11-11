package fhires

// Codes identifying the category of observation range.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/observation-range-category
type ObservationRangeCategory string

const (
	// Reference (Normal) Range for Ordinal and Continuous Observations.
	ObservationRangeCategoryReference ObservationRangeCategory = "reference"
	// Critical Range for Ordinal and Continuous Observations. Results outside this range are critical.
	ObservationRangeCategoryCritical ObservationRangeCategory = "critical"
	// Absolute Range for Ordinal and Continuous Observations. Results outside this range are not possible.
	ObservationRangeCategoryAbsolute ObservationRangeCategory = "absolute"
)

func (code ObservationRangeCategory) String() string {
	return code.Code()
}

func (code ObservationRangeCategory) Code() string {
	return string(code)
}

func (code ObservationRangeCategory) Display() string {
	switch code {
	case ObservationRangeCategoryReference:
		return "reference range"
	case ObservationRangeCategoryCritical:
		return "critical range"
	case ObservationRangeCategoryAbsolute:
		return "absolute range"
	}
	return "<unknown>"
}

func (code ObservationRangeCategory) Definition() string {
	switch code {
	case ObservationRangeCategoryReference:
		return "Reference (Normal) Range for Ordinal and Continuous Observations."
	case ObservationRangeCategoryCritical:
		return "Critical Range for Ordinal and Continuous Observations. Results outside this range are critical."
	case ObservationRangeCategoryAbsolute:
		return "Absolute Range for Ordinal and Continuous Observations. Results outside this range are not possible."
	}
	return "<unknown>"
}
