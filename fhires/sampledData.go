package fhires

import decimal "github.com/shopspring/decimal"

// SampledData Type: A series of measurements taken by a device, with upper and lower limits.
//
// There may be more than one dimension in the data.
//
// Purpose: There is a need for a concise way to handle the data produced by devices that sample a physical state at a high frequency.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/SampledData
type SampledData struct {
	// Unique id for the element within a resource (for internal references).
	//
	// This may be any string value that does not contain spaces.
	ID *string `bson:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element.
	//
	// To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	//
	// Comment: There can be no stigma associated with the use of extensions by any application, project, or standard - regardless of the institution or jurisdiction that uses or defines the extensions.  The use of extensions is what allows the FHIR specification to retain a core level of simplicity for everyone.
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty" xml:"extension,omitempty"`
	// The base quantity that a measured value of zero represents.
	//
	// In addition, this provides the units of the entire measurement series.
	Origin Quantity `bson:"origin" json:"origin" xml:"origin"`
	// Amount of intervalUnits between samples, e.g. milliseconds for time-based sampling.
	//
	// Comment: This is usually a whole number.
	Interval *decimal.Decimal `bson:"interval,omitempty" json:"interval,omitempty" xml:"interval,omitempty"`
	// The measurement unit in which the sample interval is expressed.
	IntervalUnit string `bson:"intervalUnit" json:"intervalUnit" xml:"intervalUnit"`
	// A correction factor that is applied to the sampled data points before they are added to the origin.
	Factor *decimal.Decimal `bson:"factor,omitempty" json:"factor,omitempty" xml:"factor,omitempty"`
	// The lower limit of detection of the measured points.
	//
	// This is needed if any of the data points have the value "L" (lower than detection limit).
	LowerLimit *decimal.Decimal `bson:"lowerLimit,omitempty" json:"lowerLimit,omitempty" xml:"lowerLimit,omitempty"`
	// The upper limit of detection of the measured points.
	//
	// This is needed if any of the data points have the value "U" (higher than detection limit).
	UpperLimit *decimal.Decimal `bson:"upperLimit,omitempty" json:"upperLimit,omitempty" xml:"upperLimit,omitempty"`
	// The number of sample points at each time point.
	//
	// If this value is greater than one, then the dimensions will be interlaced - all the sample points for a point in time will be recorded at once.
	//
	// Comment: If there is more than one dimension, the code for the type of data will define the meaning of the dimensions (typically ECG data).
	Dimensions int `bson:"dimensions" json:"dimensions" xml:"dimensions"`
	// Reference to ConceptMap that defines the codes used in the data.
	//
	// Comment: The ConceptMap cannot define meanings for the codes 'E', 'U', or 'L' (nor 'e', 'u', or 'l').
	CodeMap *string `bson:"codeMap,omitempty" json:"codeMap,omitempty" xml:"codeMap,omitempty"`
	// A series of data points which are decimal values separated by a single space (character u20).
	//
	//  The units in which the offsets are expressed are found in intervalUnit.  The absolute point at which the measurements begin SHALL be conveyed outside the scope of this datatype, e.g. Observation.effectiveDateTime for a timing offset.
	//
	// Comment: If offsets is present, the number of data points must be equal to the number of offsets mlutipled by the dimensions.
	Offsets *string `bson:"offsets,omitempty" json:"offsets,omitempty" xml:"offsets,omitempty"`
	// A series of data points which are decimal values or codes separated by a single space (character u20).
	//
	// The special codes "E" (error), "L" (below detection limit) and "U" (above detection limit) are also defined for used in place of decimal values.
	//
	// Comment: The data may be missing if it is omitted for summarization purposes. In general, data is required for any actual use of a SampledData.
	Data *string `bson:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
}
