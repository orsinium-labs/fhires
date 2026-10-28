package fhires

import (
	"encoding/json"
	"fmt"
	"strings"
)

// QuestionnaireAnswerConstraint is documented here http://hl7.org/fhir/ValueSet/questionnaire-answer-constraint
type QuestionnaireAnswerConstraint int

const (
	QuestionnaireAnswerConstraintOptionsOnly QuestionnaireAnswerConstraint = iota
	QuestionnaireAnswerConstraintOptionsOrType
	QuestionnaireAnswerConstraintOptionsOrString
)

func (code QuestionnaireAnswerConstraint) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *QuestionnaireAnswerConstraint) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "optionsOnly":
		*code = QuestionnaireAnswerConstraintOptionsOnly
	case "optionsOrType":
		*code = QuestionnaireAnswerConstraintOptionsOrType
	case "optionsOrString":
		*code = QuestionnaireAnswerConstraintOptionsOrString
	default:
		return fmt.Errorf("unknown QuestionnaireAnswerConstraint code `%s`", s)
	}
	return nil
}

func (code QuestionnaireAnswerConstraint) String() string {
	return code.Code()
}

func (code QuestionnaireAnswerConstraint) Code() string {
	switch code {
	case QuestionnaireAnswerConstraintOptionsOnly:
		return "optionsOnly"
	case QuestionnaireAnswerConstraintOptionsOrType:
		return "optionsOrType"
	case QuestionnaireAnswerConstraintOptionsOrString:
		return "optionsOrString"
	}
	return "<unknown>"
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
