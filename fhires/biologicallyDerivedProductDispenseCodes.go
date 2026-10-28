package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// BiologicallyDerivedProductDispenseCodes is documented here http://hl7.org/fhir/ValueSet/biologicallyderivedproductdispense-status
type BiologicallyDerivedProductDispenseCodes int

const (
	BiologicallyDerivedProductDispenseCodesPreparation BiologicallyDerivedProductDispenseCodes = iota
	BiologicallyDerivedProductDispenseCodesInProgress
	BiologicallyDerivedProductDispenseCodesAllocated
	BiologicallyDerivedProductDispenseCodesIssued
	BiologicallyDerivedProductDispenseCodesUnfulfilled
	BiologicallyDerivedProductDispenseCodesReturned
	BiologicallyDerivedProductDispenseCodesEnteredInError
	BiologicallyDerivedProductDispenseCodesUnknown
)

func (code BiologicallyDerivedProductDispenseCodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *BiologicallyDerivedProductDispenseCodes) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "preparation":
		*code = BiologicallyDerivedProductDispenseCodesPreparation
	case "in-progress":
		*code = BiologicallyDerivedProductDispenseCodesInProgress
	case "allocated":
		*code = BiologicallyDerivedProductDispenseCodesAllocated
	case "issued":
		*code = BiologicallyDerivedProductDispenseCodesIssued
	case "unfulfilled":
		*code = BiologicallyDerivedProductDispenseCodesUnfulfilled
	case "returned":
		*code = BiologicallyDerivedProductDispenseCodesReturned
	case "entered-in-error":
		*code = BiologicallyDerivedProductDispenseCodesEnteredInError
	case "unknown":
		*code = BiologicallyDerivedProductDispenseCodesUnknown
	default:
		return fmt.Errorf("unknown BiologicallyDerivedProductDispenseCodes code `%s`", s)
	}
	return nil
}

func (code BiologicallyDerivedProductDispenseCodes) String() string {
	return code.Code()
}

func (code BiologicallyDerivedProductDispenseCodes) Code() string {
	switch code {
	case BiologicallyDerivedProductDispenseCodesPreparation:
		return "preparation"
	case BiologicallyDerivedProductDispenseCodesInProgress:
		return "in-progress"
	case BiologicallyDerivedProductDispenseCodesAllocated:
		return "allocated"
	case BiologicallyDerivedProductDispenseCodesIssued:
		return "issued"
	case BiologicallyDerivedProductDispenseCodesUnfulfilled:
		return "unfulfilled"
	case BiologicallyDerivedProductDispenseCodesReturned:
		return "returned"
	case BiologicallyDerivedProductDispenseCodesEnteredInError:
		return "entered-in-error"
	case BiologicallyDerivedProductDispenseCodesUnknown:
		return "unknown"
	}
	return "<unknown>"
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
