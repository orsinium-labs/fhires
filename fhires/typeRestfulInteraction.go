package fhires

// Operations supported by REST at the type or instance level.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/type-restful-interaction
type TypeRestfulInteraction string

const (
	// Read the current state of the resource.
	TypeRestfulInteractionRead TypeRestfulInteraction = "read"
	// Read the state of a specific version of the resource.
	TypeRestfulInteractionVread TypeRestfulInteraction = "vread"
	// Update an existing resource by its id (or create it if it is new).
	TypeRestfulInteractionUpdate TypeRestfulInteraction = "update"
	// Update an existing resource by posting a set of changes to it.
	TypeRestfulInteractionPatch TypeRestfulInteraction = "patch"
	// Delete a resource.
	TypeRestfulInteractionDelete TypeRestfulInteraction = "delete"
	// Retrieve the change history for a particular resource, type of resource, or the entire system.
	TypeRestfulInteractionHistory TypeRestfulInteraction = "history"
	// Retrieve the change history for a particular resource.
	TypeRestfulInteractionHistoryInstance TypeRestfulInteraction = "history-instance"
	// Retrieve the change history for all resources of a particular type.
	TypeRestfulInteractionHistoryType TypeRestfulInteraction = "history-type"
	// Retrieve the change history for all resources on a system.
	TypeRestfulInteractionHistorySystem TypeRestfulInteraction = "history-system"
	// Create a new resource with a server assigned id.
	TypeRestfulInteractionCreate TypeRestfulInteraction = "create"
	// Search a resource type or all resources based on some filter criteria.
	TypeRestfulInteractionSearch TypeRestfulInteraction = "search"
	// Search all resources of the specified type based on some filter criteria.
	TypeRestfulInteractionSearchType TypeRestfulInteraction = "search-type"
	// Search all resources based on some filter criteria.
	TypeRestfulInteractionSearchSystem TypeRestfulInteraction = "search-system"
	// Search resources in a compartment based on some filter criteria.
	TypeRestfulInteractionSearchCompartment TypeRestfulInteraction = "search-compartment"
	// Get a Capability Statement for the system.
	TypeRestfulInteractionCapabilities TypeRestfulInteraction = "capabilities"
	// Update, create or delete a set of resources as a single transaction.
	TypeRestfulInteractionTransaction TypeRestfulInteraction = "transaction"
	// perform a set of a separate interactions in a single http operation
	TypeRestfulInteractionBatch TypeRestfulInteraction = "batch"
	// Perform an operation as defined by an OperationDefinition.
	TypeRestfulInteractionOperation TypeRestfulInteraction = "operation"
)

func (code TypeRestfulInteraction) String() string {
	return code.Code()
}

func (code TypeRestfulInteraction) Code() string {
	return string(code)
}

func (code TypeRestfulInteraction) Display() string {
	switch code {
	case TypeRestfulInteractionRead:
		return "read"
	case TypeRestfulInteractionVread:
		return "vread"
	case TypeRestfulInteractionUpdate:
		return "update"
	case TypeRestfulInteractionPatch:
		return "patch"
	case TypeRestfulInteractionDelete:
		return "delete"
	case TypeRestfulInteractionHistory:
		return "history"
	case TypeRestfulInteractionHistoryInstance:
		return "history-instance"
	case TypeRestfulInteractionHistoryType:
		return "history-type"
	case TypeRestfulInteractionHistorySystem:
		return "history-system"
	case TypeRestfulInteractionCreate:
		return "create"
	case TypeRestfulInteractionSearch:
		return "search"
	case TypeRestfulInteractionSearchType:
		return "search-type"
	case TypeRestfulInteractionSearchSystem:
		return "search-system"
	case TypeRestfulInteractionSearchCompartment:
		return "search-compartment"
	case TypeRestfulInteractionCapabilities:
		return "capabilities"
	case TypeRestfulInteractionTransaction:
		return "transaction"
	case TypeRestfulInteractionBatch:
		return "batch"
	case TypeRestfulInteractionOperation:
		return "operation"
	}
	return "<unknown>"
}

func (code TypeRestfulInteraction) Definition() string {
	switch code {
	case TypeRestfulInteractionRead:
		return "Read the current state of the resource."
	case TypeRestfulInteractionVread:
		return "Read the state of a specific version of the resource."
	case TypeRestfulInteractionUpdate:
		return "Update an existing resource by its id (or create it if it is new)."
	case TypeRestfulInteractionPatch:
		return "Update an existing resource by posting a set of changes to it."
	case TypeRestfulInteractionDelete:
		return "Delete a resource."
	case TypeRestfulInteractionHistory:
		return "Retrieve the change history for a particular resource, type of resource, or the entire system."
	case TypeRestfulInteractionHistoryInstance:
		return "Retrieve the change history for a particular resource."
	case TypeRestfulInteractionHistoryType:
		return "Retrieve the change history for all resources of a particular type."
	case TypeRestfulInteractionHistorySystem:
		return "Retrieve the change history for all resources on a system."
	case TypeRestfulInteractionCreate:
		return "Create a new resource with a server assigned id."
	case TypeRestfulInteractionSearch:
		return "Search a resource type or all resources based on some filter criteria."
	case TypeRestfulInteractionSearchType:
		return "Search all resources of the specified type based on some filter criteria."
	case TypeRestfulInteractionSearchSystem:
		return "Search all resources based on some filter criteria."
	case TypeRestfulInteractionSearchCompartment:
		return "Search resources in a compartment based on some filter criteria."
	case TypeRestfulInteractionCapabilities:
		return "Get a Capability Statement for the system."
	case TypeRestfulInteractionTransaction:
		return "Update, create or delete a set of resources as a single transaction."
	case TypeRestfulInteractionBatch:
		return "perform a set of a separate interactions in a single http operation"
	case TypeRestfulInteractionOperation:
		return "Perform an operation as defined by an OperationDefinition."
	}
	return "<unknown>"
}
