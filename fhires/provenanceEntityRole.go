package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// ProvenanceEntityRole is documented here http://hl7.org/fhir/ValueSet/provenance-entity-role
type ProvenanceEntityRole int

const (
	ProvenanceEntityRoleRevision ProvenanceEntityRole = iota
	ProvenanceEntityRoleQuotation
	ProvenanceEntityRoleSource
	ProvenanceEntityRoleInstantiates
	ProvenanceEntityRoleRemoval
)

func (code ProvenanceEntityRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *ProvenanceEntityRole) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "revision":
		*code = ProvenanceEntityRoleRevision
	case "quotation":
		*code = ProvenanceEntityRoleQuotation
	case "source":
		*code = ProvenanceEntityRoleSource
	case "instantiates":
		*code = ProvenanceEntityRoleInstantiates
	case "removal":
		*code = ProvenanceEntityRoleRemoval
	default:
		return fmt.Errorf("unknown ProvenanceEntityRole code `%s`", s)
	}
	return nil
}

func (code ProvenanceEntityRole) String() string {
	return code.Code()
}

func (code ProvenanceEntityRole) Code() string {
	switch code {
	case ProvenanceEntityRoleRevision:
		return "revision"
	case ProvenanceEntityRoleQuotation:
		return "quotation"
	case ProvenanceEntityRoleSource:
		return "source"
	case ProvenanceEntityRoleInstantiates:
		return "instantiates"
	case ProvenanceEntityRoleRemoval:
		return "removal"
	}
	return "<unknown>"
}

func (code ProvenanceEntityRole) Display() string {
	switch code {
	case ProvenanceEntityRoleRevision:
		return "Revision"
	case ProvenanceEntityRoleQuotation:
		return "Quotation"
	case ProvenanceEntityRoleSource:
		return "Source"
	case ProvenanceEntityRoleInstantiates:
		return "Instantiates"
	case ProvenanceEntityRoleRemoval:
		return "Removal"
	}
	return "<unknown>"
}

func (code ProvenanceEntityRole) Definition() string {
	switch code {
	case ProvenanceEntityRoleRevision:
		return "An entity that is used by the activity to produce a new version of that entity."
	case ProvenanceEntityRoleQuotation:
		return "An entity that is copied in full or part by an agent that is not the author of the entity."
	case ProvenanceEntityRoleSource:
		return "An entity that is used as input to the activity that produced the target."
	case ProvenanceEntityRoleInstantiates:
		return "The record resulting from this event adheres to the protocol, guideline, order set or other definition represented by this entity."
	case ProvenanceEntityRoleRemoval:
		return "An entity that is removed from accessibility, usually through the DELETE operator."
	}
	return "<unknown>"
}
