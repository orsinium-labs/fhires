package fhires

import decimal "github.com/shopspring/decimal"

// Timing Type: Specifies an event that may occur multiple times.
//
// Timing schedules are used to record when things are planned, expected or requested to occur. The most common usage is in dosage instructions for medications. They are also used when planning care of various kinds, and may be used for reporting the schedule to which past regular activities were carried out.
//
// Purpose: Need to able to track proposed timing schedules. There are several different ways to do this: one or more specified times, a simple rules like three times a day, or  before/after meals.
//
// Status: active
//
// http://hl7.org/fhir/StructureDefinition/Timing
type Timing struct {
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
	// Identifies specific times when the event occurs.
	Event []string `bson:"event,omitempty" json:"event,omitempty" xml:"event,omitempty"`
	// A set of rules that describe when the event is scheduled.
	Repeat *TimingRepeat `bson:"repeat,omitempty" json:"repeat,omitempty" xml:"repeat,omitempty"`
	// A code for the timing schedule (or just text in code.text).
	//
	// Some codes such as BID are ubiquitous, but many institutions define their own additional codes. If a code is provided, the code is understood to be a complete statement of whatever is specified in the structured timing data, and either the code or the data may be used to interpret the Timing, with the exception that .repeat.bounds still applies over the code (and is not contained in the code).
	//
	// Comment: BID etc. are defined as 'at institutionally specified times'. For example, an institution may choose that BID is "always at 7am and 6pm".  If it is inappropriate for this choice to be made, the code BID should not be used. Instead, a distinct organization-specific code should be used in place of the HL7-defined BID code and/or a structured representation should be used (in this case, specifying the two event times).
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
}

type TimingRepeat struct {
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
	// Either a duration for the length of the timing schedule, a range of possible length, or outer bounds for start and/or end limits of the timing schedule.
	BoundsDuration *Duration `bson:"boundsDuration,omitempty" json:"boundsDuration,omitempty" xml:"boundsDuration,omitempty"`
	BoundsRange    *Range    `bson:"boundsRange,omitempty" json:"boundsRange,omitempty" xml:"boundsRange,omitempty"`
	BoundsPeriod   *Period   `bson:"boundsPeriod,omitempty" json:"boundsPeriod,omitempty" xml:"boundsPeriod,omitempty"`
	// A total count of the desired number of repetitions across the duration of the entire timing specification.
	//
	// If countMax is present, this element indicates the lower bound of the allowed range of count values.
	//
	// Comment: If you have both bounds and count, then this should be understood as within the bounds period, until count times happens.
	Count *int `bson:"count,omitempty" json:"count,omitempty" xml:"count,omitempty"`
	// If present, indicates that the count is a range - so to perform the action between [count] and [countMax] times.
	CountMax *int `bson:"countMax,omitempty" json:"countMax,omitempty" xml:"countMax,omitempty"`
	// How long this thing happens for when it happens.
	//
	// If durationMax is present, this element indicates the lower bound of the allowed range of the duration.
	//
	// Comment: For some events the duration is part of the definition of the event (e.g. IV infusions, where the duration is implicit in the specified quantity and rate). For others, it's part of the timing specification (e.g. exercise).
	Duration *decimal.Decimal `bson:"duration,omitempty" json:"duration,omitempty" xml:"duration,omitempty"`
	// If present, indicates that the duration is a range - so to perform the action between [duration] and [durationMax] time length.
	//
	// Comment: For some events the duration is part of the definition of the event (e.g. IV infusions, where the duration is implicit in the specified quantity and rate). For others, it's part of the timing specification (e.g. exercise).
	DurationMax *decimal.Decimal `bson:"durationMax,omitempty" json:"durationMax,omitempty" xml:"durationMax,omitempty"`
	// The units of time for the duration, in UCUM units
	//
	// Normal practice is to use the 'mo' code as a calendar month when calculating the next occurrence.
	DurationUnit *string `bson:"durationUnit,omitempty" json:"durationUnit,omitempty" xml:"durationUnit,omitempty"`
	// The number of times to repeat the action within the specified period.
	//
	// If frequencyMax is present, this element indicates the lower bound of the allowed range of the frequency.
	Frequency *int `bson:"frequency,omitempty" json:"frequency,omitempty" xml:"frequency,omitempty"`
	// If present, indicates that the frequency is a range - so to repeat between [frequency] and [frequencyMax] times within the period or period range.
	FrequencyMax *int `bson:"frequencyMax,omitempty" json:"frequencyMax,omitempty" xml:"frequencyMax,omitempty"`
	// Indicates the duration of time over which repetitions are to occur; e.g.
	//
	// to express "3 times per day", 3 would be the frequency and "1 day" would be the period. If periodMax is present, this element indicates the lower bound of the allowed range of the period length.
	Period *decimal.Decimal `bson:"period,omitempty" json:"period,omitempty" xml:"period,omitempty"`
	// If present, indicates that the period is a range from [period] to [periodMax], allowing expressing concepts such as "do this once every 3-5 days.
	PeriodMax *decimal.Decimal `bson:"periodMax,omitempty" json:"periodMax,omitempty" xml:"periodMax,omitempty"`
	// The units of time for the period in UCUM units
	//
	// Normal practice is to use the 'mo' code as a calendar month when calculating the next occurrence.
	PeriodUnit *string `bson:"periodUnit,omitempty" json:"periodUnit,omitempty" xml:"periodUnit,omitempty"`
	// If one or more days of week is provided, then the action happens only on the specified day(s).
	//
	// Comment: If no days are specified, the action is assumed to happen every day as otherwise specified.
	DayOfWeek []DaysOfWeek `bson:"dayOfWeek,omitempty" json:"dayOfWeek,omitempty" xml:"dayOfWeek,omitempty"`
	// Specified time of day for action to take place.
	//
	// Comment: When time of day is specified, it is inferred that the action happens every day (as filtered by dayofWeek) on the specified times. If there's a timeOfDay, there cannot be a when, or vice versa.
	TimeOfDay []string `bson:"timeOfDay,omitempty" json:"timeOfDay,omitempty" xml:"timeOfDay,omitempty"`
	// An approximate time period during the day, potentially linked to an event of daily living that indicates when the action should occur.
	//
	// Comment: When more than one event is listed, the event is tied to the union of the specified events.
	When []string `bson:"when,omitempty" json:"when,omitempty" xml:"when,omitempty"`
	// The number of minutes from the event.
	//
	// If the event code does not indicate whether the minutes is before or after the event, then the offset is assumed to be after the event.
	Offset *int `bson:"offset,omitempty" json:"offset,omitempty" xml:"offset,omitempty"`
}
