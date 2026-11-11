package fhires

// The meaning of the hierarchy of concepts in a code system.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/codesystem-hierarchy-meaning
type CodeSystemHierarchyMeaning string

const (
	// No particular relationship between the concepts can be assumed, except what can be determined by inspection of the definitions of the elements (possible reasons to use this: importing from a source where this is not defined, or where various parts of the hierarchy have different meanings).
	CodeSystemHierarchyMeaningGroupedBy CodeSystemHierarchyMeaning = "grouped-by"
	// A hierarchy where the child concepts have an IS-A relationship with the parents - that is, all the properties of the parent are also true for its child concepts. Not that is-a is a property of the concepts, so additional subsumption relationships may be defined using properties.
	CodeSystemHierarchyMeaningIsA CodeSystemHierarchyMeaning = "is-a"
	// Child elements list the individual parts of a composite whole (e.g. body site).
	CodeSystemHierarchyMeaningPartOf CodeSystemHierarchyMeaning = "part-of"
	// Child concepts in the hierarchy may have only one parent, and there is a presumption that the code system is a "closed world" meaning all things must be in the hierarchy. This results in concepts such as "not otherwise classified.".
	CodeSystemHierarchyMeaningClassifiedWith CodeSystemHierarchyMeaning = "classified-with"
)

func (code CodeSystemHierarchyMeaning) String() string {
	return code.Code()
}

func (code CodeSystemHierarchyMeaning) Code() string {
	return string(code)
}

func (code CodeSystemHierarchyMeaning) Display() string {
	switch code {
	case CodeSystemHierarchyMeaningGroupedBy:
		return "Grouped By"
	case CodeSystemHierarchyMeaningIsA:
		return "Is-A"
	case CodeSystemHierarchyMeaningPartOf:
		return "Part Of"
	case CodeSystemHierarchyMeaningClassifiedWith:
		return "Classified With"
	}
	return "<unknown>"
}

func (code CodeSystemHierarchyMeaning) Definition() string {
	switch code {
	case CodeSystemHierarchyMeaningGroupedBy:
		return "No particular relationship between the concepts can be assumed, except what can be determined by inspection of the definitions of the elements (possible reasons to use this: importing from a source where this is not defined, or where various parts of the hierarchy have different meanings)."
	case CodeSystemHierarchyMeaningIsA:
		return "A hierarchy where the child concepts have an IS-A relationship with the parents - that is, all the properties of the parent are also true for its child concepts. Not that is-a is a property of the concepts, so additional subsumption relationships may be defined using properties."
	case CodeSystemHierarchyMeaningPartOf:
		return "Child elements list the individual parts of a composite whole (e.g. body site)."
	case CodeSystemHierarchyMeaningClassifiedWith:
		return "Child concepts in the hierarchy may have only one parent, and there is a presumption that the code system is a \"closed world\" meaning all things must be in the hierarchy. This results in concepts such as \"not otherwise classified.\"."
	}
	return "<unknown>"
}
