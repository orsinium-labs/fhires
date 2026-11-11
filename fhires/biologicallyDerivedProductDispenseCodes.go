package fhires

// BiologicallyDerivedProductDispense Status Codes
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/biologicallyderivedproductdispense-status
type BiologicallyDerivedProductDispenseCodes string

const (
	// The dispense process has started but not yet completed.
	BiologicallyDerivedProductDispenseCodesPreparation BiologicallyDerivedProductDispenseCodes = "preparation"
	// The dispense process is in progress.
	BiologicallyDerivedProductDispenseCodesInProgress BiologicallyDerivedProductDispenseCodes = "in-progress"
	// The requested product has been allocated and is ready for transport.
	BiologicallyDerivedProductDispenseCodesAllocated BiologicallyDerivedProductDispenseCodes = "allocated"
	// The dispensed product has been picked up.
	BiologicallyDerivedProductDispenseCodesIssued BiologicallyDerivedProductDispenseCodes = "issued"
	// The dispense could not be completed.
	BiologicallyDerivedProductDispenseCodesUnfulfilled BiologicallyDerivedProductDispenseCodes = "unfulfilled"
	// The dispensed product was returned.
	BiologicallyDerivedProductDispenseCodesReturned BiologicallyDerivedProductDispenseCodes = "returned"
	// The dispense was entered in error and therefore nullified.
	BiologicallyDerivedProductDispenseCodesEnteredInError BiologicallyDerivedProductDispenseCodes = "entered-in-error"
	// The authoring system does not know which of the status values applies for this dispense. Note: this concept is not to be used for other - one of the listed statuses is presumed to apply, it's just not known which one.
	BiologicallyDerivedProductDispenseCodesUnknown BiologicallyDerivedProductDispenseCodes = "unknown"
)

func (code BiologicallyDerivedProductDispenseCodes) String() string {
	return code.Code()
}

func (code BiologicallyDerivedProductDispenseCodes) Code() string {
	return string(code)
}

func (code BiologicallyDerivedProductDispenseCodes) Display() string {
	switch code {
	case BiologicallyDerivedProductDispenseCodesPreparation:
		return "Preparation"
	case BiologicallyDerivedProductDispenseCodesInProgress:
		return "In Progress"
	case BiologicallyDerivedProductDispenseCodesAllocated:
		return "Allocated"
	case BiologicallyDerivedProductDispenseCodesIssued:
		return "Issued"
	case BiologicallyDerivedProductDispenseCodesUnfulfilled:
		return "Unfulfilled"
	case BiologicallyDerivedProductDispenseCodesReturned:
		return "Returned"
	case BiologicallyDerivedProductDispenseCodesEnteredInError:
		return "Entered in Error"
	case BiologicallyDerivedProductDispenseCodesUnknown:
		return "Unknown"
	}
	return "<unknown>"
}

func (code BiologicallyDerivedProductDispenseCodes) Definition() string {
	switch code {
	case BiologicallyDerivedProductDispenseCodesPreparation:
		return "The dispense process has started but not yet completed."
	case BiologicallyDerivedProductDispenseCodesInProgress:
		return "The dispense process is in progress."
	case BiologicallyDerivedProductDispenseCodesAllocated:
		return "The requested product has been allocated and is ready for transport."
	case BiologicallyDerivedProductDispenseCodesIssued:
		return "The dispensed product has been picked up."
	case BiologicallyDerivedProductDispenseCodesUnfulfilled:
		return "The dispense could not be completed."
	case BiologicallyDerivedProductDispenseCodesReturned:
		return "The dispensed product was returned."
	case BiologicallyDerivedProductDispenseCodesEnteredInError:
		return "The dispense was entered in error and therefore nullified."
	case BiologicallyDerivedProductDispenseCodesUnknown:
		return "The authoring system does not know which of the status values applies for this dispense. Note: this concept is not to be used for other - one of the listed statuses is presumed to apply, it's just not known which one."
	}
	return "<unknown>"
}
