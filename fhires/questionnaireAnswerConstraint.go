package fhires

// Codes that describe the types of constraints possible on a question item that has a list of permitted answers
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/questionnaire-answer-constraint
type QuestionnaireAnswerConstraint string

const (
	// Only values listed as answerOption or in the expansion of the answerValueSet are permitted
	QuestionnaireAnswerConstraintOptionsOnly QuestionnaireAnswerConstraint = "optionsOnly"
	// In addition to the values listed as answerOption or in the expansion of the answerValueSet, any other values that correspond to the specified item.type are permitted
	QuestionnaireAnswerConstraintOptionsOrType QuestionnaireAnswerConstraint = "optionsOrType"
	// In addition to the values listed as answerOption or in the expansion of the answerValueSet, free-text strings are permitted.  Answers will have a type of 'string'.
	QuestionnaireAnswerConstraintOptionsOrString QuestionnaireAnswerConstraint = "optionsOrString"
)

func (code QuestionnaireAnswerConstraint) String() string {
	return code.Code()
}

func (code QuestionnaireAnswerConstraint) Code() string {
	return string(code)
}

func (code QuestionnaireAnswerConstraint) Display() string {
	switch code {
	case QuestionnaireAnswerConstraintOptionsOnly:
		return "Options only"
	case QuestionnaireAnswerConstraintOptionsOrType:
		return "Options or 'type'"
	case QuestionnaireAnswerConstraintOptionsOrString:
		return "Options or string"
	}
	return "<unknown>"
}

func (code QuestionnaireAnswerConstraint) Definition() string {
	switch code {
	case QuestionnaireAnswerConstraintOptionsOnly:
		return "Only values listed as answerOption or in the expansion of the answerValueSet are permitted"
	case QuestionnaireAnswerConstraintOptionsOrType:
		return "In addition to the values listed as answerOption or in the expansion of the answerValueSet, any other values that correspond to the specified item.type are permitted"
	case QuestionnaireAnswerConstraintOptionsOrString:
		return "In addition to the values listed as answerOption or in the expansion of the answerValueSet, free-text strings are permitted.  Answers will have a type of 'string'."
	}
	return "<unknown>"
}
