package fhires

// Availability Type: Availability data for an {item}.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Availability
type Availability struct {
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
	// Times the {item} is available.
	AvailableTime []AvailabilityAvailableTime `bson:"availableTime,omitempty" json:"availableTime,omitempty" xml:"availableTime,omitempty"`
	// Not available during this time due to provided reason.
	NotAvailableTime []AvailabilityNotAvailableTime `bson:"notAvailableTime,omitempty" json:"notAvailableTime,omitempty" xml:"notAvailableTime,omitempty"`
}

type AvailabilityAvailableTime struct {
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
	// mon | tue | wed | thu | fri | sat | sun.
	DaysOfWeek []DaysOfWeek `bson:"daysOfWeek,omitempty" json:"daysOfWeek,omitempty" xml:"daysOfWeek,omitempty"`
	// Always available? i.e. 24 hour service.
	AllDay *bool `bson:"allDay,omitempty" json:"allDay,omitempty" xml:"allDay,omitempty"`
	// Opening time of day (ignored if allDay = true).
	//
	// Comment: The timezone is expected to be specified or implied by the context this datatype is used.
	AvailableStartTime *string `bson:"availableStartTime,omitempty" json:"availableStartTime,omitempty" xml:"availableStartTime,omitempty"`
	// Closing time of day (ignored if allDay = true).
	//
	// Comment: The timezone is expected to be specified or implied by the context this datatype is used.
	AvailableEndTime *string `bson:"availableEndTime,omitempty" json:"availableEndTime,omitempty" xml:"availableEndTime,omitempty"`
}

type AvailabilityNotAvailableTime struct {
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
	// Reason presented to the user explaining why time not available.
	//
	// Comment: The reason will generally be provided to give the textual reason for displaying when the {item} is not available, e.g. 'Closed public holidays' or 'Independence Day'. In cases such as this, the `during` might not be included and local knowledge would be required in such cases (as don't desire to keep updating when the holiday occurs each year).
	//
	// e.g.2: 'Closed for maintenance over the summer' for this example you would want to include the `during` period, unless this was a university hospital and the "summer" period was well known, but would recommend its inclusion anyway.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Service not available during this period.
	During *Period `bson:"during,omitempty" json:"during,omitempty" xml:"during,omitempty"`
}
