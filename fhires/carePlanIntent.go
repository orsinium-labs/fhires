package fhires

// Codes indicating the degree of authority/intentionality associated with a care plan.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/care-plan-intent
type CarePlanIntent string

const (
	// The request is a suggestion made by someone/something that does not have an intention to ensure it occurs and without providing an authorization to act.
	CarePlanIntentProposal CarePlanIntent = "proposal"
	// The request represents an intention to ensure something occurs without providing an authorization for others to act.
	CarePlanIntentPlan CarePlanIntent = "plan"
	// The request represents a legally binding instruction authored by a Patient or RelatedPerson.
	CarePlanIntentDirective CarePlanIntent = "directive"
	// The request represents a request/demand and authorization for action by the requestor.
	CarePlanIntentOrder CarePlanIntent = "order"
	// The request represents an original authorization for action.
	CarePlanIntentOriginalOrder CarePlanIntent = "original-order"
	// The request represents an automatically generated supplemental authorization for action based on a parent authorization together with initial results of the action taken against that parent authorization.
	CarePlanIntentReflexOrder CarePlanIntent = "reflex-order"
	// The request represents the view of an authorization instantiated by a fulfilling system representing the details of the fulfiller's intention to act upon a submitted order.
	CarePlanIntentFillerOrder CarePlanIntent = "filler-order"
	// An order created in fulfillment of a broader order that represents the authorization for a single activity occurrence.  E.g. The administration of a single dose of a drug.
	CarePlanIntentInstanceOrder CarePlanIntent = "instance-order"
	// The request represents a component or option for a RequestOrchestration that establishes timing, conditionality and/or other constraints among a set of requests.  Refer to [[[RequestOrchestration]]] for additional information on how this status is used.
	CarePlanIntentOption CarePlanIntent = "option"
)

func (code CarePlanIntent) String() string {
	return code.Code()
}

func (code CarePlanIntent) Code() string {
	return string(code)
}

func (code CarePlanIntent) Display() string {
	switch code {
	case CarePlanIntentProposal:
		return "Proposal"
	case CarePlanIntentPlan:
		return "Plan"
	case CarePlanIntentDirective:
		return "Directive"
	case CarePlanIntentOrder:
		return "Order"
	case CarePlanIntentOriginalOrder:
		return "Original Order"
	case CarePlanIntentReflexOrder:
		return "Reflex Order"
	case CarePlanIntentFillerOrder:
		return "Filler Order"
	case CarePlanIntentInstanceOrder:
		return "Instance Order"
	case CarePlanIntentOption:
		return "Option"
	}
	return "<unknown>"
}

func (code CarePlanIntent) Definition() string {
	switch code {
	case CarePlanIntentProposal:
		return "The request is a suggestion made by someone/something that does not have an intention to ensure it occurs and without providing an authorization to act."
	case CarePlanIntentPlan:
		return "The request represents an intention to ensure something occurs without providing an authorization for others to act."
	case CarePlanIntentDirective:
		return "The request represents a legally binding instruction authored by a Patient or RelatedPerson."
	case CarePlanIntentOrder:
		return "The request represents a request/demand and authorization for action by the requestor."
	case CarePlanIntentOriginalOrder:
		return "The request represents an original authorization for action."
	case CarePlanIntentReflexOrder:
		return "The request represents an automatically generated supplemental authorization for action based on a parent authorization together with initial results of the action taken against that parent authorization."
	case CarePlanIntentFillerOrder:
		return "The request represents the view of an authorization instantiated by a fulfilling system representing the details of the fulfiller's intention to act upon a submitted order."
	case CarePlanIntentInstanceOrder:
		return "An order created in fulfillment of a broader order that represents the authorization for a single activity occurrence.  E.g. The administration of a single dose of a drug."
	case CarePlanIntentOption:
		return "The request represents a component or option for a RequestOrchestration that establishes timing, conditionality and/or other constraints among a set of requests.  Refer to [[[RequestOrchestration]]] for additional information on how this status is used."
	}
	return "<unknown>"
}
