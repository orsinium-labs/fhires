package fhires

// Status of the supply delivery.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/supplydelivery-status
type SupplyDeliveryStatus string

const (
	// Supply has been requested, but not delivered.
	SupplyDeliveryStatusInProgress SupplyDeliveryStatus = "in-progress"
	// Supply has been delivered ("completed").
	SupplyDeliveryStatusCompleted SupplyDeliveryStatus = "completed"
	// Delivery was not completed.
	SupplyDeliveryStatusAbandoned SupplyDeliveryStatus = "abandoned"
	// This electronic record should never have existed, though it is possible that real-world decisions were based on it. (If real-world activity has occurred, the status should be "abandoned" rather than "entered-in-error".).
	SupplyDeliveryStatusEnteredInError SupplyDeliveryStatus = "entered-in-error"
)

func (code SupplyDeliveryStatus) String() string {
	return code.Code()
}

func (code SupplyDeliveryStatus) Code() string {
	return string(code)
}

func (code SupplyDeliveryStatus) Display() string {
	switch code {
	case SupplyDeliveryStatusInProgress:
		return "In Progress"
	case SupplyDeliveryStatusCompleted:
		return "Delivered"
	case SupplyDeliveryStatusAbandoned:
		return "Abandoned"
	case SupplyDeliveryStatusEnteredInError:
		return "Entered In Error"
	}
	return "<unknown>"
}

func (code SupplyDeliveryStatus) Definition() string {
	switch code {
	case SupplyDeliveryStatusInProgress:
		return "Supply has been requested, but not delivered."
	case SupplyDeliveryStatusCompleted:
		return "Supply has been delivered (\"completed\")."
	case SupplyDeliveryStatusAbandoned:
		return "Delivery was not completed."
	case SupplyDeliveryStatusEnteredInError:
		return "This electronic record should never have existed, though it is possible that real-world decisions were based on it. (If real-world activity has occurred, the status should be \"abandoned\" rather than \"entered-in-error\".)."
	}
	return "<unknown>"
}
