package fhires

// Codes that guide the display of disabled questionnaire items
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/questionnaire-disabled-display
type QuestionnaireItemDisabledDisplay string

const (
	// The item (and its children) should not be visible to the user at all.
	QuestionnaireItemDisabledDisplayHidden QuestionnaireItemDisabledDisplay = "hidden"
	// The item (and possibly its children) should not be selectable or editable but should still be visible - to allow the user to see what questions *could* have been completed had other answers caused the item to be enabled.
	QuestionnaireItemDisabledDisplayProtected QuestionnaireItemDisabledDisplay = "protected"
)

func (code QuestionnaireItemDisabledDisplay) String() string {
	return code.Code()
}

func (code QuestionnaireItemDisabledDisplay) Code() string {
	return string(code)
}

func (code QuestionnaireItemDisabledDisplay) Display() string {
	switch code {
	case QuestionnaireItemDisabledDisplayHidden:
		return "Hidden"
	case QuestionnaireItemDisabledDisplayProtected:
		return "Protected"
	}
	return "<unknown>"
}

func (code QuestionnaireItemDisabledDisplay) Definition() string {
	switch code {
	case QuestionnaireItemDisabledDisplayHidden:
		return "The item (and its children) should not be visible to the user at all."
	case QuestionnaireItemDisabledDisplayProtected:
		return "The item (and possibly its children) should not be selectable or editable but should still be visible - to allow the user to see what questions *could* have been completed had other answers caused the item to be enabled."
	}
	return "<unknown>"
}
