package fhires

// Operations supported by REST at the system level.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/system-restful-interaction
type SystemRestfulInteraction string

const (
	// Read the current state of the resource.
	SystemRestfulInteractionRead SystemRestfulInteraction = "read"
	// Read the state of a specific version of the resource.
	SystemRestfulInteractionVread SystemRestfulInteraction = "vread"
	// Update an existing resource by its id (or create it if it is new).
	SystemRestfulInteractionUpdate SystemRestfulInteraction = "update"
	// Update an existing resource by posting a set of changes to it.
	SystemRestfulInteractionPatch SystemRestfulInteraction = "patch"
	// Delete a resource.
	SystemRestfulInteractionDelete SystemRestfulInteraction = "delete"
	// Retrieve the change history for a particular resource, type of resource, or the entire system.
	SystemRestfulInteractionHistory SystemRestfulInteraction = "history"
	// Retrieve the change history for a particular resource.
	SystemRestfulInteractionHistoryInstance SystemRestfulInteraction = "history-instance"
	// Retrieve the change history for all resources of a particular type.
	SystemRestfulInteractionHistoryType SystemRestfulInteraction = "history-type"
	// Retrieve the change history for all resources on a system.
	SystemRestfulInteractionHistorySystem SystemRestfulInteraction = "history-system"
	// Create a new resource with a server assigned id.
	SystemRestfulInteractionCreate SystemRestfulInteraction = "create"
	// Search a resource type or all resources based on some filter criteria.
	SystemRestfulInteractionSearch SystemRestfulInteraction = "search"
	// Search all resources of the specified type based on some filter criteria.
	SystemRestfulInteractionSearchType SystemRestfulInteraction = "search-type"
	// Search all resources based on some filter criteria.
	SystemRestfulInteractionSearchSystem SystemRestfulInteraction = "search-system"
	// Search resources in a compartment based on some filter criteria.
	SystemRestfulInteractionSearchCompartment SystemRestfulInteraction = "search-compartment"
	// Get a Capability Statement for the system.
	SystemRestfulInteractionCapabilities SystemRestfulInteraction = "capabilities"
	// Update, create or delete a set of resources as a single transaction.
	SystemRestfulInteractionTransaction SystemRestfulInteraction = "transaction"
	// perform a set of a separate interactions in a single http operation
	SystemRestfulInteractionBatch SystemRestfulInteraction = "batch"
	// Perform an operation as defined by an OperationDefinition.
	SystemRestfulInteractionOperation SystemRestfulInteraction = "operation"
)

func (code SystemRestfulInteraction) String() string {
	return code.Code()
}

func (code SystemRestfulInteraction) Code() string {
	return string(code)
}

func (code SystemRestfulInteraction) Display() string {
	switch code {
	case SystemRestfulInteractionRead:
		return "read"
	case SystemRestfulInteractionVread:
		return "vread"
	case SystemRestfulInteractionUpdate:
		return "update"
	case SystemRestfulInteractionPatch:
		return "patch"
	case SystemRestfulInteractionDelete:
		return "delete"
	case SystemRestfulInteractionHistory:
		return "history"
	case SystemRestfulInteractionHistoryInstance:
		return "history-instance"
	case SystemRestfulInteractionHistoryType:
		return "history-type"
	case SystemRestfulInteractionHistorySystem:
		return "history-system"
	case SystemRestfulInteractionCreate:
		return "create"
	case SystemRestfulInteractionSearch:
		return "search"
	case SystemRestfulInteractionSearchType:
		return "search-type"
	case SystemRestfulInteractionSearchSystem:
		return "search-system"
	case SystemRestfulInteractionSearchCompartment:
		return "search-compartment"
	case SystemRestfulInteractionCapabilities:
		return "capabilities"
	case SystemRestfulInteractionTransaction:
		return "transaction"
	case SystemRestfulInteractionBatch:
		return "batch"
	case SystemRestfulInteractionOperation:
		return "operation"
	}
	return "<unknown>"
}

func (code SystemRestfulInteraction) Definition() string {
	switch code {
	case SystemRestfulInteractionRead:
		return "Read the current state of the resource."
	case SystemRestfulInteractionVread:
		return "Read the state of a specific version of the resource."
	case SystemRestfulInteractionUpdate:
		return "Update an existing resource by its id (or create it if it is new)."
	case SystemRestfulInteractionPatch:
		return "Update an existing resource by posting a set of changes to it."
	case SystemRestfulInteractionDelete:
		return "Delete a resource."
	case SystemRestfulInteractionHistory:
		return "Retrieve the change history for a particular resource, type of resource, or the entire system."
	case SystemRestfulInteractionHistoryInstance:
		return "Retrieve the change history for a particular resource."
	case SystemRestfulInteractionHistoryType:
		return "Retrieve the change history for all resources of a particular type."
	case SystemRestfulInteractionHistorySystem:
		return "Retrieve the change history for all resources on a system."
	case SystemRestfulInteractionCreate:
		return "Create a new resource with a server assigned id."
	case SystemRestfulInteractionSearch:
		return "Search a resource type or all resources based on some filter criteria."
	case SystemRestfulInteractionSearchType:
		return "Search all resources of the specified type based on some filter criteria."
	case SystemRestfulInteractionSearchSystem:
		return "Search all resources based on some filter criteria."
	case SystemRestfulInteractionSearchCompartment:
		return "Search resources in a compartment based on some filter criteria."
	case SystemRestfulInteractionCapabilities:
		return "Get a Capability Statement for the system."
	case SystemRestfulInteractionTransaction:
		return "Update, create or delete a set of resources as a single transaction."
	case SystemRestfulInteractionBatch:
		return "perform a set of a separate interactions in a single http operation"
	case SystemRestfulInteractionOperation:
		return "Perform an operation as defined by an OperationDefinition."
	}
	return "<unknown>"
}
