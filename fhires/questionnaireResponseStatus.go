package fhires

// Lifecycle status of the questionnaire response.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/questionnaire-answers-status
type QuestionnaireResponseStatus string

const (
	// This QuestionnaireResponse has been partially filled out with answers but changes or additions are still expected to be made to it.
	QuestionnaireResponseStatusInProgress QuestionnaireResponseStatus = "in-progress"
	// This QuestionnaireResponse has been filled out with answers and the current content is regarded as definitive.
	QuestionnaireResponseStatusCompleted QuestionnaireResponseStatus = "completed"
	// This QuestionnaireResponse has been filled out with answers, then marked as complete, yet changes or additions have been made to it afterwards.
	QuestionnaireResponseStatusAmended QuestionnaireResponseStatus = "amended"
	// This QuestionnaireResponse was entered in error and voided.
	QuestionnaireResponseStatusEnteredInError QuestionnaireResponseStatus = "entered-in-error"
	// This QuestionnaireResponse has been partially filled out with answers but has been abandoned. No subsequent changes can be made.
	QuestionnaireResponseStatusStopped QuestionnaireResponseStatus = "stopped"
)

func (code QuestionnaireResponseStatus) String() string {
	return code.Code()
}

func (code QuestionnaireResponseStatus) Code() string {
	return string(code)
}

func (code QuestionnaireResponseStatus) Display() string {
	switch code {
	case QuestionnaireResponseStatusInProgress:
		return "In Progress"
	case QuestionnaireResponseStatusCompleted:
		return "Completed"
	case QuestionnaireResponseStatusAmended:
		return "Amended"
	case QuestionnaireResponseStatusEnteredInError:
		return "Entered in Error"
	case QuestionnaireResponseStatusStopped:
		return "Stopped"
	}
	return "<unknown>"
}

func (code QuestionnaireResponseStatus) Definition() string {
	switch code {
	case QuestionnaireResponseStatusInProgress:
		return "This QuestionnaireResponse has been partially filled out with answers but changes or additions are still expected to be made to it."
	case QuestionnaireResponseStatusCompleted:
		return "This QuestionnaireResponse has been filled out with answers and the current content is regarded as definitive."
	case QuestionnaireResponseStatusAmended:
		return "This QuestionnaireResponse has been filled out with answers, then marked as complete, yet changes or additions have been made to it afterwards."
	case QuestionnaireResponseStatusEnteredInError:
		return "This QuestionnaireResponse was entered in error and voided."
	case QuestionnaireResponseStatusStopped:
		return "This QuestionnaireResponse has been partially filled out with answers but has been abandoned. No subsequent changes can be made."
	}
	return "<unknown>"
}
