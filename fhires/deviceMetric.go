package fhires

import "encoding/xml"

// Describes a measurement, calculation or setting capability of a device.
//
//	The DeviceMetric resource is derived from the ISO/IEEE 11073-10201 Domain Information Model standard, but is more widely applicable.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/DeviceMetric
type DeviceMetric struct {
	BaseResource
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human.
	//
	// The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	//
	// Comment: Contained resources do not have a narrative. Resources that are not contained SHOULD have a narrative. In some cases, a resource may only have text with little or no additional discrete data (as long as all minOccurs=1 elements are satisfied).  This may be necessary for data from legacy systems where information is captured as a "text blob" or where text is additionally entered raw or narrated and encoded information is added later.
	Text *Narrative `bson:"text,omitempty" json:"text,omitempty" xml:"text,omitempty"`
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, nor can they have their own independent transaction scope.
	//
	// This is allowed to be a Parameters resource if and only if it is referenced by a resource that provides context/meaning.
	//
	// Comment: This should never be done when the content can be identified properly, as once identification is lost, it is extremely difficult (and context dependent) to restore it again. Contained resources may have profiles and tags in their meta elements, but SHALL NOT have security labels.
	// May be used to represent additional information that is not part of the basic definition of the resource.
	//
	// To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	//
	// Comment: There can be no stigma associated with the use of extensions by any application, project, or standard - regardless of the institution or jurisdiction that uses or defines the extensions.  The use of extensions is what allows the FHIR specification to retain a core level of simplicity for everyone.
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty" xml:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants.
	//
	// Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	//
	// Comment: There can be no stigma associated with the use of extensions by any application, project, or standard - regardless of the institution or jurisdiction that uses or defines the extensions.  The use of extensions is what allows the FHIR specification to retain a core level of simplicity for everyone.
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty" xml:"modifierExtension,omitempty"`
	// Instance identifiers assigned to a device, by the device or gateway software, manufacturers, other organizations or owners.
	//
	// For example, handle ID.
	//
	// Comment: For identifiers assigned to a device by the device or gateway software, the `system` element of the identifier should be set to any identifier of the device.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Describes the type of the metric. For example: Heart Rate, PEEP Setting, etc.
	//
	// Comment: DeviceMetric.type can be referred to either IEEE 11073-10101 or LOINC.
	Type CodeableConcept `bson:"type" json:"type" xml:"type"`
	// Describes the unit that an observed value determined for this metric will have.
	//
	// For example: Percent, Seconds, etc.
	//
	// Comment: DeviceMetric.unit can refer to either UCUM or preferable a RTMMS coding system.
	Unit *CodeableConcept `bson:"unit,omitempty" json:"unit,omitempty" xml:"unit,omitempty"`
	// Describes the link to the Device.  This is also known as a channel device.
	//
	// Comment: Note if this is not the top level device - also know as the source device, it may be referenced in the device.parent element.
	Device Reference `bson:"device" json:"device" xml:"device"`
	// Indicates current operational state of the device. For example: On, Off, Standby, etc.
	OperationalStatus *DeviceMetricOperationalStatus `bson:"operationalStatus,omitempty" json:"operationalStatus,omitempty" xml:"operationalStatus,omitempty"`
	// The preferred color associated with the metric (e.g., display color).
	//
	// This is often used to aid clinicians to track and identify parameter types by color. In practice, consider a Patient Monitor that has ECG/HR and Pleth; the metrics are displayed in different characteristic colors, such as HR in blue, BP in green, and PR and SpO2 in magenta.
	Color *string `bson:"color,omitempty" json:"color,omitempty" xml:"color,omitempty"`
	// Indicates the category of the observation generation process.
	//
	// A DeviceMetric can be for example a setting, measurement, or calculation.
	Category DeviceMetricCategory `bson:"category" json:"category" xml:"category"`
	// The frequency at which the metric is taken or recorded.
	//
	// Devices measure metrics at a wide range of frequencies; for example, an ECG might sample measurements in the millisecond range, while an NIBP might trigger only once an hour. Less often, the measurementFrequency may be based on a unit other than time, such as distance (e.g. for a measuring wheel). The update period may be different than the measurement frequency, if the device does not update the published observed value with the same frequency as it was measured.
	MeasurementFrequency *Quantity `bson:"measurementFrequency,omitempty" json:"measurementFrequency,omitempty" xml:"measurementFrequency,omitempty"`
	// Describes the calibrations that have been performed or that are required to be performed.
	Calibration []DeviceMetricCalibration `bson:"calibration,omitempty" json:"calibration,omitempty" xml:"calibration,omitempty"`
}

type DeviceMetricCalibration struct {
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
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants.
	//
	// Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	//
	// Comment: There can be no stigma associated with the use of extensions by any application, project, or standard - regardless of the institution or jurisdiction that uses or defines the extensions.  The use of extensions is what allows the FHIR specification to retain a core level of simplicity for everyone.
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty" xml:"modifierExtension,omitempty"`
	// Describes the type of the calibration method.
	Type *DeviceMetricCalibrationType `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Describes the state of the calibration.
	State *DeviceMetricCalibrationState `bson:"state,omitempty" json:"state,omitempty" xml:"state,omitempty"`
	// Describes the time last calibration has been performed.
	Time *string `bson:"time,omitempty" json:"time,omitempty" xml:"time,omitempty"`
}

func (r DeviceMetric) ResourceType() string {
	return "DeviceMetric"
}

// MarshalJSON implements json.Marshaler
func (r DeviceMetric) MarshalJSON() ([]byte, error) {
	type Wrapper DeviceMetric
	return marshalJSON(r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Unmarshaler].
func (r *DeviceMetric) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper DeviceMetric
	w := Wrapper(*r)
	err := unmarshalXML(d, start, &w)
	*r = DeviceMetric(w)
	return err
}
