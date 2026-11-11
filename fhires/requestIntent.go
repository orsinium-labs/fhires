package fhires

// Codes indicating the degree of authority/intentionality associated with a request.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/request-intent
type RequestIntent string

const (
	// The request is a suggestion made by someone/something that does not have an intention to ensure it occurs and without providing an authorization to act.
	RequestIntentProposal RequestIntent = "proposal"
	// The request represents an intention to ensure something occurs without providing an authorization for others to act.
	RequestIntentPlan RequestIntent = "plan"
	// The request represents a legally binding instruction authored by a Patient or RelatedPerson.
	RequestIntentDirective RequestIntent = "directive"
	// The request represents a request/demand and authorization for action by the requestor.
	RequestIntentOrder RequestIntent = "order"
	// The request represents an original authorization for action.
	RequestIntentOriginalOrder RequestIntent = "original-order"
	// The request represents an automatically generated supplemental authorization for action based on a parent authorization together with initial results of the action taken against that parent authorization.
	RequestIntentReflexOrder RequestIntent = "reflex-order"
	// The request represents the view of an authorization instantiated by a fulfilling system representing the details of the fulfiller's intention to act upon a submitted order.
	RequestIntentFillerOrder RequestIntent = "filler-order"
	// An order created in fulfillment of a broader order that represents the authorization for a single activity occurrence.  E.g. The administration of a single dose of a drug.
	RequestIntentInstanceOrder RequestIntent = "instance-order"
	// The request represents a component or option for a RequestOrchestration that establishes timing, conditionality and/or other constraints among a set of requests.  Refer to [[[RequestOrchestration]]] for additional information on how this status is used.
	RequestIntentOption RequestIntent = "option"
)

func (code RequestIntent) String() string {
	return code.Code()
}

func (code RequestIntent) Code() string {
	return string(code)
}

func (code RequestIntent) Display() string {
	switch code {
	case RequestIntentProposal:
		return "Proposal"
	case RequestIntentPlan:
		return "Plan"
	case RequestIntentDirective:
		return "Directive"
	case RequestIntentOrder:
		return "Order"
	case RequestIntentOriginalOrder:
		return "Original Order"
	case RequestIntentReflexOrder:
		return "Reflex Order"
	case RequestIntentFillerOrder:
		return "Filler Order"
	case RequestIntentInstanceOrder:
		return "Instance Order"
	case RequestIntentOption:
		return "Option"
	}
	return "<unknown>"
}

func (code RequestIntent) Definition() string {
	switch code {
	case RequestIntentProposal:
		return "The request is a suggestion made by someone/something that does not have an intention to ensure it occurs and without providing an authorization to act."
	case RequestIntentPlan:
		return "The request represents an intention to ensure something occurs without providing an authorization for others to act."
	case RequestIntentDirective:
		return "The request represents a legally binding instruction authored by a Patient or RelatedPerson."
	case RequestIntentOrder:
		return "The request represents a request/demand and authorization for action by the requestor."
	case RequestIntentOriginalOrder:
		return "The request represents an original authorization for action."
	case RequestIntentReflexOrder:
		return "The request represents an automatically generated supplemental authorization for action based on a parent authorization together with initial results of the action taken against that parent authorization."
	case RequestIntentFillerOrder:
		return "The request represents the view of an authorization instantiated by a fulfilling system representing the details of the fulfiller's intention to act upon a submitted order."
	case RequestIntentInstanceOrder:
		return "An order created in fulfillment of a broader order that represents the authorization for a single activity occurrence.  E.g. The administration of a single dose of a drug."
	case RequestIntentOption:
		return "The request represents a component or option for a RequestOrchestration that establishes timing, conditionality and/or other constraints among a set of requests.  Refer to [[[RequestOrchestration]]] for additional information on how this status is used."
	}
	return "<unknown>"
}
