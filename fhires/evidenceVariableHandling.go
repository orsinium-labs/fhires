package fhires

import (
	"encoding/json"
	"fmt"
	"strings"
)

// EvidenceVariableHandling is documented here http://hl7.org/fhir/ValueSet/variable-handling
type EvidenceVariableHandling int

const (
	EvidenceVariableHandlingContinuous EvidenceVariableHandling = iota
	EvidenceVariableHandlingDichotomous
	EvidenceVariableHandlingOrdinal
	EvidenceVariableHandlingPolychotomous
)

func (code EvidenceVariableHandling) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *EvidenceVariableHandling) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "continuous":
		*code = EvidenceVariableHandlingContinuous
	case "dichotomous":
		*code = EvidenceVariableHandlingDichotomous
	case "ordinal":
		*code = EvidenceVariableHandlingOrdinal
	case "polychotomous":
		*code = EvidenceVariableHandlingPolychotomous
	default:
		return fmt.Errorf("unknown EvidenceVariableHandling code `%s`", s)
	}
	return nil
}

func (code EvidenceVariableHandling) String() string {
	return code.Code()
}

func (code EvidenceVariableHandling) Code() string {
	switch code {
	case EvidenceVariableHandlingContinuous:
		return "continuous"
	case EvidenceVariableHandlingDichotomous:
		return "dichotomous"
	case EvidenceVariableHandlingOrdinal:
		return "ordinal"
	case EvidenceVariableHandlingPolychotomous:
		return "polychotomous"
	}
	return "<unknown>"
}

func (code EvidenceVariableHandling) Display() string {
	switch code {
	case EvidenceVariableHandlingContinuous:
		return "continuous variable"
	case EvidenceVariableHandlingDichotomous:
		return "dichotomous variable"
	case EvidenceVariableHandlingOrdinal:
		return "ordinal variable"
	case EvidenceVariableHandlingPolychotomous:
		return "polychotomous variable"
	}
	return "<unknown>"
}

func (code EvidenceVariableHandling) Definition() string {
	switch code {
	case EvidenceVariableHandlingContinuous:
		return "A continuous variable is one for which, within the limits the variable ranges, any value is possible (from STATO http://purl.obolibrary.org/obo/STATO_0000251)."
	case EvidenceVariableHandlingDichotomous:
		return "A dichotomous variable is a categorical variable which is defined to have only 2 categories or possible values (from STATO http://purl.obolibrary.org/obo/STATO_0000090)."
	case EvidenceVariableHandlingOrdinal:
		return "An ordinal variable is a categorical variable where the discrete possible values are ordered or correspond to an implicit ranking (from STATO http://purl.obolibrary.org/obo/STATO_0000228)."
	case EvidenceVariableHandlingPolychotomous:
		return "A polychotomous variable is a categorical variable which is defined to have minimally 2 categories or possible values. (from STATO  http://purl.obolibrary.org/obo/STATO_0000087).  Suggestion to limit code use to situations when neither dichotomous nor ordinal variables apply."
	}
	return "<unknown>"
}
