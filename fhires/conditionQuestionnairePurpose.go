package fhires

// The use of a questionnaire.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/condition-questionnaire-purpose
type ConditionQuestionnairePurpose string

const (
	// A pre-admit questionnaire.
	ConditionQuestionnairePurposePreadmit ConditionQuestionnairePurpose = "preadmit"
	// A questionnaire that helps with diferential diagnosis.
	ConditionQuestionnairePurposeDiffDiagnosis ConditionQuestionnairePurpose = "diff-diagnosis"
	// A questionnaire to check on outcomes for the patient.
	ConditionQuestionnairePurposeOutcome ConditionQuestionnairePurpose = "outcome"
)

func (code ConditionQuestionnairePurpose) String() string {
	return code.Code()
}

func (code ConditionQuestionnairePurpose) Code() string {
	return string(code)
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
