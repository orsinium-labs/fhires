package fhires

// How a resource reference is interpreted when testing consent restrictions.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/consent-data-meaning
type ConsentDataMeaning string

const (
	// The consent applies directly to the instance of the resource.
	ConsentDataMeaningInstance ConsentDataMeaning = "instance"
	// The consent applies directly to the instance of the resource and instances it refers to.
	ConsentDataMeaningRelated ConsentDataMeaning = "related"
	// The consent applies directly to the instance of the resource and instances that refer to it.
	ConsentDataMeaningDependents ConsentDataMeaning = "dependents"
	// The consent applies to instances of resources that are authored by.
	ConsentDataMeaningAuthoredby ConsentDataMeaning = "authoredby"
)

func (code ConsentDataMeaning) String() string {
	return code.Code()
}

func (code ConsentDataMeaning) Code() string {
	return string(code)
}

func (code ConsentDataMeaning) Display() string {
	switch code {
	case ConsentDataMeaningInstance:
		return "Instance"
	case ConsentDataMeaningRelated:
		return "Related"
	case ConsentDataMeaningDependents:
		return "Dependents"
	case ConsentDataMeaningAuthoredby:
		return "AuthoredBy"
	}
	return "<unknown>"
}

func (code ConsentDataMeaning) Definition() string {
	switch code {
	case ConsentDataMeaningInstance:
		return "The consent applies directly to the instance of the resource."
	case ConsentDataMeaningRelated:
		return "The consent applies directly to the instance of the resource and instances it refers to."
	case ConsentDataMeaningDependents:
		return "The consent applies directly to the instance of the resource and instances that refer to it."
	case ConsentDataMeaningAuthoredby:
		return "The consent applies to instances of resources that are authored by."
	}
	return "<unknown>"
}
