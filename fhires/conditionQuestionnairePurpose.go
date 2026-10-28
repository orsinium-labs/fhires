package fhires

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ConditionQuestionnairePurpose is documented here http://hl7.org/fhir/ValueSet/condition-questionnaire-purpose
type ConditionQuestionnairePurpose int

const (
	ConditionQuestionnairePurposePreadmit ConditionQuestionnairePurpose = iota
	ConditionQuestionnairePurposeDiffDiagnosis
	ConditionQuestionnairePurposeOutcome
)

func (code ConditionQuestionnairePurpose) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *ConditionQuestionnairePurpose) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "preadmit":
		*code = ConditionQuestionnairePurposePreadmit
	case "diff-diagnosis":
		*code = ConditionQuestionnairePurposeDiffDiagnosis
	case "outcome":
		*code = ConditionQuestionnairePurposeOutcome
	default:
		return fmt.Errorf("unknown ConditionQuestionnairePurpose code `%s`", s)
	}
	return nil
}

func (code ConditionQuestionnairePurpose) String() string {
	return code.Code()
}

func (code ConditionQuestionnairePurpose) Code() string {
	switch code {
	case ConditionQuestionnairePurposePreadmit:
		return "preadmit"
	case ConditionQuestionnairePurposeDiffDiagnosis:
		return "diff-diagnosis"
	case ConditionQuestionnairePurposeOutcome:
		return "outcome"
	}
	return "<unknown>"
}

func (code ConditionQuestionnairePurpose) Display() string {
	switch code {
	case ConditionQuestionnairePurposePreadmit:
		return "Pre-admit"
	case ConditionQuestionnairePurposeDiffDiagnosis:
		return "Diff Diagnosis"
	case ConditionQuestionnairePurposeOutcome:
		return "Outcome"
	}
	return "<unknown>"
}

func (code ConditionQuestionnairePurpose) Definition() string {
	switch code {
	case ConditionQuestionnairePurposePreadmit:
		return "A pre-admit questionnaire."
	case ConditionQuestionnairePurposeDiffDiagnosis:
		return "A questionnaire that helps with diferential diagnosis."
	case ConditionQuestionnairePurposeOutcome:
		return "A questionnaire to check on outcomes for the patient."
	}
	return "<unknown>"
}
