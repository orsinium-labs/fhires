package fhires

// The days of the week.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/days-of-week
type DaysOfWeek string

const (
	// Monday.
	DaysOfWeekMon DaysOfWeek = "mon"
	// Tuesday.
	DaysOfWeekTue DaysOfWeek = "tue"
	// Wednesday.
	DaysOfWeekWed DaysOfWeek = "wed"
	// Thursday.
	DaysOfWeekThu DaysOfWeek = "thu"
	// Friday.
	DaysOfWeekFri DaysOfWeek = "fri"
	// Saturday.
	DaysOfWeekSat DaysOfWeek = "sat"
	// Sunday.
	DaysOfWeekSun DaysOfWeek = "sun"
)

func (code DaysOfWeek) String() string {
	return code.Code()
}

func (code DaysOfWeek) Code() string {
	return string(code)
}

func (code DaysOfWeek) Display() string {
	switch code {
	case DaysOfWeekMon:
		return "Monday"
	case DaysOfWeekTue:
		return "Tuesday"
	case DaysOfWeekWed:
		return "Wednesday"
	case DaysOfWeekThu:
		return "Thursday"
	case DaysOfWeekFri:
		return "Friday"
	case DaysOfWeekSat:
		return "Saturday"
	case DaysOfWeekSun:
		return "Sunday"
	}
	return "<unknown>"
}

func (code DaysOfWeek) Definition() string {
	switch code {
	case DaysOfWeekMon:
		return "Monday."
	case DaysOfWeekTue:
		return "Tuesday."
	case DaysOfWeekWed:
		return "Wednesday."
	case DaysOfWeekThu:
		return "Thursday."
	case DaysOfWeekFri:
		return "Friday."
	case DaysOfWeekSat:
		return "Saturday."
	case DaysOfWeekSun:
		return "Sunday."
	}
	return "<unknown>"
}
