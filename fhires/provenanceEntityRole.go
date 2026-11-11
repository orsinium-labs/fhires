package fhires

// How an entity was used in an activity.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/provenance-entity-role
type ProvenanceEntityRole string

const (
	// An entity that is used by the activity to produce a new version of that entity.
	ProvenanceEntityRoleRevision ProvenanceEntityRole = "revision"
	// An entity that is copied in full or part by an agent that is not the author of the entity.
	ProvenanceEntityRoleQuotation ProvenanceEntityRole = "quotation"
	// An entity that is used as input to the activity that produced the target.
	ProvenanceEntityRoleSource ProvenanceEntityRole = "source"
	// The record resulting from this event adheres to the protocol, guideline, order set or other definition represented by this entity.
	ProvenanceEntityRoleInstantiates ProvenanceEntityRole = "instantiates"
	// An entity that is removed from accessibility, usually through the DELETE operator.
	ProvenanceEntityRoleRemoval ProvenanceEntityRole = "removal"
)

func (code ProvenanceEntityRole) String() string {
	return code.Code()
}

func (code ProvenanceEntityRole) Code() string {
	return string(code)
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
