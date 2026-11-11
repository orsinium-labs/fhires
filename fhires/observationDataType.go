package fhires

// Permitted data type for observation value.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/permitted-data-type
type ObservationDataType string

const (
	// A measured amount.
	ObservationDataTypeQuantity ObservationDataType = "Quantity"
	// A coded concept from a reference terminology and/or text.
	ObservationDataTypeCodeableConcept ObservationDataType = "CodeableConcept"
	// A sequence of Unicode characters.
	ObservationDataTypeString ObservationDataType = "string"
	// true or false.
	ObservationDataTypeBoolean ObservationDataType = "boolean"
	// A signed integer.
	ObservationDataTypeInteger ObservationDataType = "integer"
	// A set of values bounded by low and high.
	ObservationDataTypeRange ObservationDataType = "Range"
	// A ratio of two Quantity values - a numerator and a denominator.
	ObservationDataTypeRatio ObservationDataType = "Ratio"
	// A series of measurements taken by a device.
	ObservationDataTypeSampledData ObservationDataType = "SampledData"
	// A time during the day, in the format hh:mm:ss.
	ObservationDataTypeTime ObservationDataType = "time"
	// A date, date-time or partial date (e.g. just year or year + month) as used in human communication.
	ObservationDataTypeDateTime ObservationDataType = "dateTime"
	// A time range defined by start and end date/time.
	ObservationDataTypePeriod ObservationDataType = "Period"
)

func (code ObservationDataType) String() string {
	return code.Code()
}

func (code ObservationDataType) Code() string {
	return string(code)
}

func (code ObservationDataType) Display() string {
	switch code {
	case ObservationDataTypeQuantity:
		return "Quantity"
	case ObservationDataTypeCodeableConcept:
		return "CodeableConcept"
	case ObservationDataTypeString:
		return "string"
	case ObservationDataTypeBoolean:
		return "boolean"
	case ObservationDataTypeInteger:
		return "integer"
	case ObservationDataTypeRange:
		return "Range"
	case ObservationDataTypeRatio:
		return "Ratio"
	case ObservationDataTypeSampledData:
		return "SampledData"
	case ObservationDataTypeTime:
		return "time"
	case ObservationDataTypeDateTime:
		return "dateTime"
	case ObservationDataTypePeriod:
		return "Period"
	}
	return "<unknown>"
}

func (code ObservationDataType) Definition() string {
	switch code {
	case ObservationDataTypeQuantity:
		return "A measured amount."
	case ObservationDataTypeCodeableConcept:
		return "A coded concept from a reference terminology and/or text."
	case ObservationDataTypeString:
		return "A sequence of Unicode characters."
	case ObservationDataTypeBoolean:
		return "true or false."
	case ObservationDataTypeInteger:
		return "A signed integer."
	case ObservationDataTypeRange:
		return "A set of values bounded by low and high."
	case ObservationDataTypeRatio:
		return "A ratio of two Quantity values - a numerator and a denominator."
	case ObservationDataTypeSampledData:
		return "A series of measurements taken by a device."
	case ObservationDataTypeTime:
		return "A time during the day, in the format hh:mm:ss."
	case ObservationDataTypeDateTime:
		return "A date, date-time or partial date (e.g. just year or year + month) as used in human communication."
	case ObservationDataTypePeriod:
		return "A time range defined by start and end date/time."
	}
	return "<unknown>"
}
