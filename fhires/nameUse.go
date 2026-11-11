package fhires

// The use of a human name.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/name-use
type NameUse string

const (
	// Known as/conventional/the one you normally use.
	NameUseUsual NameUse = "usual"
	// The formal name as registered in an official (government) registry, but which name might not be commonly used. May be called "legal name".
	NameUseOfficial NameUse = "official"
	// A temporary name. Name.period can provide more detailed information. This may also be used for temporary names assigned at birth or in emergency situations.
	NameUseTemp NameUse = "temp"
	// A name that is used to address the person in an informal manner, but is not part of their formal or usual name.
	NameUseNickname NameUse = "nickname"
	// Anonymous assigned name, alias, or pseudonym (used to protect a person's identity for privacy reasons).
	NameUseAnonymous NameUse = "anonymous"
	// This name is no longer in use (or was never correct, but retained for records).
	NameUseOld NameUse = "old"
	// A name used prior to changing name because of marriage. This name use is for use by applications that collect and store names that were used prior to a marriage. Marriage naming customs vary greatly around the world, and are constantly changing. This term is not gender specific. The use of this term does not imply any particular history for a person's name.
	NameUseMaiden NameUse = "maiden"
)

func (code NameUse) String() string {
	return code.Code()
}

func (code NameUse) Code() string {
	return string(code)
}

func (code NameUse) Display() string {
	switch code {
	case NameUseUsual:
		return "Usual"
	case NameUseOfficial:
		return "Official"
	case NameUseTemp:
		return "Temp"
	case NameUseNickname:
		return "Nickname"
	case NameUseAnonymous:
		return "Anonymous"
	case NameUseOld:
		return "Old"
	case NameUseMaiden:
		return "Name changed for Marriage"
	}
	return "<unknown>"
}

func (code NameUse) Definition() string {
	switch code {
	case NameUseUsual:
		return "Known as/conventional/the one you normally use."
	case NameUseOfficial:
		return "The formal name as registered in an official (government) registry, but which name might not be commonly used. May be called \"legal name\"."
	case NameUseTemp:
		return "A temporary name. Name.period can provide more detailed information. This may also be used for temporary names assigned at birth or in emergency situations."
	case NameUseNickname:
		return "A name that is used to address the person in an informal manner, but is not part of their formal or usual name."
	case NameUseAnonymous:
		return "Anonymous assigned name, alias, or pseudonym (used to protect a person's identity for privacy reasons)."
	case NameUseOld:
		return "This name is no longer in use (or was never correct, but retained for records)."
	case NameUseMaiden:
		return "A name used prior to changing name because of marriage. This name use is for use by applications that collect and store names that were used prior to a marriage. Marriage naming customs vary greatly around the world, and are constantly changing. This term is not gender specific. The use of this term does not imply any particular history for a person's name."
	}
	return "<unknown>"
}
