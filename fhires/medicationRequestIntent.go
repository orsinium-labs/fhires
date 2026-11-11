package fhires

// MedicationRequest Intent Codes
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/medicationrequest-intent
type MedicationRequestIntent string

const (
	// The request is a suggestion made by someone/something that doesn't have an intention to ensure it occurs and without providing an authorization to act
	MedicationRequestIntentProposal MedicationRequestIntent = "proposal"
	// The request represents an intention to ensure something occurs without providing an authorization for others to act.
	MedicationRequestIntentPlan MedicationRequestIntent = "plan"
	// The request represents a request/demand and authorization for action
	MedicationRequestIntentOrder MedicationRequestIntent = "order"
	// The request represents the original authorization for the medication request.
	MedicationRequestIntentOriginalOrder MedicationRequestIntent = "original-order"
	// The request represents an automatically generated supplemental authorization for action based on a parent authorization together with initial results of the action taken against that parent authorization..
	MedicationRequestIntentReflexOrder MedicationRequestIntent = "reflex-order"
	// The request represents the view of an authorization instantiated by a fulfilling system representing the details of the fulfiller's intention to act upon a submitted order.
	MedicationRequestIntentFillerOrder MedicationRequestIntent = "filler-order"
	// The request represents an instance for the particular order and is used to generate a schedule of requests on a medication administration record (MAR).
	MedicationRequestIntentInstanceOrder MedicationRequestIntent = "instance-order"
	// The request represents a component or option for a RequestOrchestration that establishes timing, conditionality and/or  other constraints among a set of requests.
	MedicationRequestIntentOption MedicationRequestIntent = "option"
)

func (code MedicationRequestIntent) String() string {
	return code.Code()
}

func (code MedicationRequestIntent) Code() string {
	return string(code)
}

func (code MedicationRequestIntent) Display() string {
	switch code {
	case MedicationRequestIntentProposal:
		return "Proposal"
	case MedicationRequestIntentPlan:
		return "Plan"
	case MedicationRequestIntentOrder:
		return "Order"
	case MedicationRequestIntentOriginalOrder:
		return "Original Order"
	case MedicationRequestIntentReflexOrder:
		return "Reflex Order"
	case MedicationRequestIntentFillerOrder:
		return "Filler Order"
	case MedicationRequestIntentInstanceOrder:
		return "Instance Order"
	case MedicationRequestIntentOption:
		return "Option"
	}
	return "<unknown>"
}

func (code MedicationRequestIntent) Definition() string {
	switch code {
	case MedicationRequestIntentProposal:
		return "The request is a suggestion made by someone/something that doesn't have an intention to ensure it occurs and without providing an authorization to act"
	case MedicationRequestIntentPlan:
		return "The request represents an intention to ensure something occurs without providing an authorization for others to act."
	case MedicationRequestIntentOrder:
		return "The request represents a request/demand and authorization for action"
	case MedicationRequestIntentOriginalOrder:
		return "The request represents the original authorization for the medication request."
	case MedicationRequestIntentReflexOrder:
		return "The request represents an automatically generated supplemental authorization for action based on a parent authorization together with initial results of the action taken against that parent authorization.."
	case MedicationRequestIntentFillerOrder:
		return "The request represents the view of an authorization instantiated by a fulfilling system representing the details of the fulfiller's intention to act upon a submitted order."
	case MedicationRequestIntentInstanceOrder:
		return "The request represents an instance for the particular order and is used to generate a schedule of requests on a medication administration record (MAR)."
	case MedicationRequestIntentOption:
		return "The request represents a component or option for a RequestOrchestration that establishes timing, conditionality and/or  other constraints among a set of requests."
	}
	return "<unknown>"
}
