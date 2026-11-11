package fhires

// FHIR RESTful interaction codes used for SubscriptionTopic trigger.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/interaction-trigger
type InteractionTrigger string

const (
	// Read the current state of the resource.
	InteractionTriggerRead InteractionTrigger = "read"
	// Read the state of a specific version of the resource.
	InteractionTriggerVread InteractionTrigger = "vread"
	// Update an existing resource by its id (or create it if it is new).
	InteractionTriggerUpdate InteractionTrigger = "update"
	// Update an existing resource by posting a set of changes to it.
	InteractionTriggerPatch InteractionTrigger = "patch"
	// Delete a resource.
	InteractionTriggerDelete InteractionTrigger = "delete"
	// Retrieve the change history for a particular resource, type of resource, or the entire system.
	InteractionTriggerHistory InteractionTrigger = "history"
	// Retrieve the change history for a particular resource.
	InteractionTriggerHistoryInstance InteractionTrigger = "history-instance"
	// Retrieve the change history for all resources of a particular type.
	InteractionTriggerHistoryType InteractionTrigger = "history-type"
	// Retrieve the change history for all resources on a system.
	InteractionTriggerHistorySystem InteractionTrigger = "history-system"
	// Create a new resource with a server assigned id.
	InteractionTriggerCreate InteractionTrigger = "create"
	// Search a resource type or all resources based on some filter criteria.
	InteractionTriggerSearch InteractionTrigger = "search"
	// Search all resources of the specified type based on some filter criteria.
	InteractionTriggerSearchType InteractionTrigger = "search-type"
	// Search all resources based on some filter criteria.
	InteractionTriggerSearchSystem InteractionTrigger = "search-system"
	// Search resources in a compartment based on some filter criteria.
	InteractionTriggerSearchCompartment InteractionTrigger = "search-compartment"
	// Get a Capability Statement for the system.
	InteractionTriggerCapabilities InteractionTrigger = "capabilities"
	// Update, create or delete a set of resources as a single transaction.
	InteractionTriggerTransaction InteractionTrigger = "transaction"
	// perform a set of a separate interactions in a single http operation
	InteractionTriggerBatch InteractionTrigger = "batch"
	// Perform an operation as defined by an OperationDefinition.
	InteractionTriggerOperation InteractionTrigger = "operation"
)

func (code InteractionTrigger) String() string {
	return code.Code()
}

func (code InteractionTrigger) Code() string {
	return string(code)
}

func (code InteractionTrigger) Display() string {
	switch code {
	case InteractionTriggerRead:
		return "read"
	case InteractionTriggerVread:
		return "vread"
	case InteractionTriggerUpdate:
		return "update"
	case InteractionTriggerPatch:
		return "patch"
	case InteractionTriggerDelete:
		return "delete"
	case InteractionTriggerHistory:
		return "history"
	case InteractionTriggerHistoryInstance:
		return "history-instance"
	case InteractionTriggerHistoryType:
		return "history-type"
	case InteractionTriggerHistorySystem:
		return "history-system"
	case InteractionTriggerCreate:
		return "create"
	case InteractionTriggerSearch:
		return "search"
	case InteractionTriggerSearchType:
		return "search-type"
	case InteractionTriggerSearchSystem:
		return "search-system"
	case InteractionTriggerSearchCompartment:
		return "search-compartment"
	case InteractionTriggerCapabilities:
		return "capabilities"
	case InteractionTriggerTransaction:
		return "transaction"
	case InteractionTriggerBatch:
		return "batch"
	case InteractionTriggerOperation:
		return "operation"
	}
	return "<unknown>"
}

func (code InteractionTrigger) Definition() string {
	switch code {
	case InteractionTriggerRead:
		return "Read the current state of the resource."
	case InteractionTriggerVread:
		return "Read the state of a specific version of the resource."
	case InteractionTriggerUpdate:
		return "Update an existing resource by its id (or create it if it is new)."
	case InteractionTriggerPatch:
		return "Update an existing resource by posting a set of changes to it."
	case InteractionTriggerDelete:
		return "Delete a resource."
	case InteractionTriggerHistory:
		return "Retrieve the change history for a particular resource, type of resource, or the entire system."
	case InteractionTriggerHistoryInstance:
		return "Retrieve the change history for a particular resource."
	case InteractionTriggerHistoryType:
		return "Retrieve the change history for all resources of a particular type."
	case InteractionTriggerHistorySystem:
		return "Retrieve the change history for all resources on a system."
	case InteractionTriggerCreate:
		return "Create a new resource with a server assigned id."
	case InteractionTriggerSearch:
		return "Search a resource type or all resources based on some filter criteria."
	case InteractionTriggerSearchType:
		return "Search all resources of the specified type based on some filter criteria."
	case InteractionTriggerSearchSystem:
		return "Search all resources based on some filter criteria."
	case InteractionTriggerSearchCompartment:
		return "Search resources in a compartment based on some filter criteria."
	case InteractionTriggerCapabilities:
		return "Get a Capability Statement for the system."
	case InteractionTriggerTransaction:
		return "Update, create or delete a set of resources as a single transaction."
	case InteractionTriggerBatch:
		return "perform a set of a separate interactions in a single http operation"
	case InteractionTriggerOperation:
		return "Perform an operation as defined by an OperationDefinition."
	}
	return "<unknown>"
}
