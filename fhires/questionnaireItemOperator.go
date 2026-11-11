package fhires

// The criteria by which a question is enabled.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/questionnaire-enable-operator
type QuestionnaireItemOperator string

const (
	// True if the determination of 'whether an answer exists for the question' is equal to the enableWhen answer (which must be a boolean).
	QuestionnaireItemOperatorExists QuestionnaireItemOperator = "exists"
	// True if at least one answer has a value that is equal to the enableWhen answer.
	QuestionnaireItemOperatorEquals QuestionnaireItemOperator = "="
	// True if no answer has a value that is equal to the enableWhen answer.
	QuestionnaireItemOperatorNotEquals QuestionnaireItemOperator = "!="
	// True if at least one answer has a value that is greater than the enableWhen answer.
	QuestionnaireItemOperatorGreaterThan QuestionnaireItemOperator = ">"
	// True if at least one answer has a value that is less than the enableWhen answer.
	QuestionnaireItemOperatorLessThan QuestionnaireItemOperator = "<"
	// True if at least one answer has a value that is greater or equal to the enableWhen answer.
	QuestionnaireItemOperatorGreaterOrEquals QuestionnaireItemOperator = ">="
	// True if at least one answer has a value that is less or equal to the enableWhen answer.
	QuestionnaireItemOperatorLessOrEquals QuestionnaireItemOperator = "<="
)

func (code QuestionnaireItemOperator) String() string {
	return code.Code()
}

func (code QuestionnaireItemOperator) Code() string {
	return string(code)
}

func (code QuestionnaireItemOperator) Display() string {
	switch code {
	case QuestionnaireItemOperatorExists:
		return "Exists"
	case QuestionnaireItemOperatorEquals:
		return "Equals"
	case QuestionnaireItemOperatorNotEquals:
		return "Not Equals"
	case QuestionnaireItemOperatorGreaterThan:
		return "Greater Than"
	case QuestionnaireItemOperatorLessThan:
		return "Less Than"
	case QuestionnaireItemOperatorGreaterOrEquals:
		return "Greater or Equals"
	case QuestionnaireItemOperatorLessOrEquals:
		return "Less or Equals"
	}
	return "<unknown>"
}

func (code QuestionnaireItemOperator) Definition() string {
	switch code {
	case QuestionnaireItemOperatorExists:
		return "True if the determination of 'whether an answer exists for the question' is equal to the enableWhen answer (which must be a boolean)."
	case QuestionnaireItemOperatorEquals:
		return "True if at least one answer has a value that is equal to the enableWhen answer."
	case QuestionnaireItemOperatorNotEquals:
		return "True if no answer has a value that is equal to the enableWhen answer."
	case QuestionnaireItemOperatorGreaterThan:
		return "True if at least one answer has a value that is greater than the enableWhen answer."
	case QuestionnaireItemOperatorLessThan:
		return "True if at least one answer has a value that is less than the enableWhen answer."
	case QuestionnaireItemOperatorGreaterOrEquals:
		return "True if at least one answer has a value that is greater or equal to the enableWhen answer."
	case QuestionnaireItemOperatorLessOrEquals:
		return "True if at least one answer has a value that is less or equal to the enableWhen answer."
	}
	return "<unknown>"
}
