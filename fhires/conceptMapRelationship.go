package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// ConceptMapRelationship is documented here http://hl7.org/fhir/ValueSet/concept-map-relationship
type ConceptMapRelationship int

const (
	ConceptMapRelationshipRelatedTo ConceptMapRelationship = iota
	ConceptMapRelationshipEquivalent
	ConceptMapRelationshipSourceIsNarrowerThanTarget
	ConceptMapRelationshipSourceIsBroaderThanTarget
	ConceptMapRelationshipNotRelatedTo
)

func (code ConceptMapRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *ConceptMapRelationship) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "related-to":
		*code = ConceptMapRelationshipRelatedTo
	case "equivalent":
		*code = ConceptMapRelationshipEquivalent
	case "source-is-narrower-than-target":
		*code = ConceptMapRelationshipSourceIsNarrowerThanTarget
	case "source-is-broader-than-target":
		*code = ConceptMapRelationshipSourceIsBroaderThanTarget
	case "not-related-to":
		*code = ConceptMapRelationshipNotRelatedTo
	default:
		return fmt.Errorf("unknown ConceptMapRelationship code `%s`", s)
	}
	return nil
}

func (code ConceptMapRelationship) String() string {
	return code.Code()
}

func (code ConceptMapRelationship) Code() string {
	switch code {
	case ConceptMapRelationshipRelatedTo:
		return "related-to"
	case ConceptMapRelationshipEquivalent:
		return "equivalent"
	case ConceptMapRelationshipSourceIsNarrowerThanTarget:
		return "source-is-narrower-than-target"
	case ConceptMapRelationshipSourceIsBroaderThanTarget:
		return "source-is-broader-than-target"
	case ConceptMapRelationshipNotRelatedTo:
		return "not-related-to"
	}
	return "<unknown>"
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
