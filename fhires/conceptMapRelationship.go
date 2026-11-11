package fhires

// The relationship between concepts.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/concept-map-relationship
type ConceptMapRelationship string

const (
	// The concepts are related to each other, but the exact relationship is not known.
	ConceptMapRelationshipRelatedTo ConceptMapRelationship = "related-to"
	// The definitions of the concepts mean the same thing.
	ConceptMapRelationshipEquivalent ConceptMapRelationship = "equivalent"
	// The source concept is narrower in meaning than the target concept.
	ConceptMapRelationshipSourceIsNarrowerThanTarget ConceptMapRelationship = "source-is-narrower-than-target"
	// The source concept is broader in meaning than the target concept.
	ConceptMapRelationshipSourceIsBroaderThanTarget ConceptMapRelationship = "source-is-broader-than-target"
	// This is an explicit assertion that the target concept is not related to the source concept.
	ConceptMapRelationshipNotRelatedTo ConceptMapRelationship = "not-related-to"
)

func (code ConceptMapRelationship) String() string {
	return code.Code()
}

func (code ConceptMapRelationship) Code() string {
	return string(code)
}

func (code ConceptMapRelationship) Display() string {
	switch code {
	case ConceptMapRelationshipRelatedTo:
		return "Related To"
	case ConceptMapRelationshipEquivalent:
		return "Equivalent"
	case ConceptMapRelationshipSourceIsNarrowerThanTarget:
		return "Source Is Narrower Than Target"
	case ConceptMapRelationshipSourceIsBroaderThanTarget:
		return "Source Is Broader Than Target"
	case ConceptMapRelationshipNotRelatedTo:
		return "Not Related To"
	}
	return "<unknown>"
}

func (code ConceptMapRelationship) Definition() string {
	switch code {
	case ConceptMapRelationshipRelatedTo:
		return "The concepts are related to each other, but the exact relationship is not known."
	case ConceptMapRelationshipEquivalent:
		return "The definitions of the concepts mean the same thing."
	case ConceptMapRelationshipSourceIsNarrowerThanTarget:
		return "The source concept is narrower in meaning than the target concept."
	case ConceptMapRelationshipSourceIsBroaderThanTarget:
		return "The source concept is broader in meaning than the target concept."
	case ConceptMapRelationshipNotRelatedTo:
		return "This is an explicit assertion that the target concept is not related to the source concept."
	}
	return "<unknown>"
}
