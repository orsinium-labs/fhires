package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// Codes indicating the degree of authority/intentionality associated with a care plan.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/care-plan-intent
type CarePlanIntent int

const (
	CarePlanIntentProposal CarePlanIntent = iota
	// The request represents an intention to ensure something occurs without providing an authorization for others to act.
	CarePlanIntentPlan
	// The request represents a legally binding instruction authored by a Patient or RelatedPerson.
	CarePlanIntentDirective
	// The request represents a request/demand and authorization for action by the requestor.
	CarePlanIntentOrder
	// The request represents an original authorization for action.
	CarePlanIntentOriginalOrder
	// The request represents an automatically generated supplemental authorization for action based on a parent authorization together with initial results of the action taken against that parent authorization.
	CarePlanIntentReflexOrder
	// The request represents the view of an authorization instantiated by a fulfilling system representing the details of the fulfiller's intention to act upon a submitted order.
	CarePlanIntentFillerOrder
	// An order created in fulfillment of a broader order that represents the authorization for a single activity occurrence.  E.g. The administration of a single dose of a drug.
	CarePlanIntentInstanceOrder
	// The request represents a component or option for a RequestOrchestration that establishes timing, conditionality and/or other constraints among a set of requests.  Refer to [[[RequestOrchestration]]] for additional information on how this status is used.
	CarePlanIntentOption
)

func (code CarePlanIntent) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *CarePlanIntent) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "proposal":
		*code = CarePlanIntentProposal
	case "plan":
		*code = CarePlanIntentPlan
	case "directive":
		*code = CarePlanIntentDirective
	case "order":
		*code = CarePlanIntentOrder
	case "original-order":
		*code = CarePlanIntentOriginalOrder
	case "reflex-order":
		*code = CarePlanIntentReflexOrder
	case "filler-order":
		*code = CarePlanIntentFillerOrder
	case "instance-order":
		*code = CarePlanIntentInstanceOrder
	case "option":
		*code = CarePlanIntentOption
	default:
		return fmt.Errorf("unknown CarePlanIntent code `%s`", s)
	}
	return nil
}

func (code CarePlanIntent) String() string {
	return code.Code()
}

func (code CarePlanIntent) Code() string {
	switch code {
	case CarePlanIntentProposal:
		return "proposal"
	case CarePlanIntentPlan:
		return "plan"
	case CarePlanIntentDirective:
		return "directive"
	case CarePlanIntentOrder:
		return "order"
	case CarePlanIntentOriginalOrder:
		return "original-order"
	case CarePlanIntentReflexOrder:
		return "reflex-order"
	case CarePlanIntentFillerOrder:
		return "filler-order"
	case CarePlanIntentInstanceOrder:
		return "instance-order"
	case CarePlanIntentOption:
		return "option"
	}
	return "<unknown>"
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
