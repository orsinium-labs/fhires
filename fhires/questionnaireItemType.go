package fhires

// Distinguishes groups from questions and display text and indicates data type for questions.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/item-type
type QuestionnaireItemType string

const (
	// An item with no direct answer but should have at least one child item.
	QuestionnaireItemTypeGroup QuestionnaireItemType = "group"
	// Text for display that will not capture an answer or have child items.
	QuestionnaireItemTypeDisplay QuestionnaireItemType = "display"
	// An item that defines a specific answer to be captured, and which may have child items. (the answer provided in the QuestionnaireResponse should be of the defined datatype).
	QuestionnaireItemTypeQuestion QuestionnaireItemType = "question"
	// Question with a yes/no answer (valueBoolean).
	QuestionnaireItemTypeBoolean QuestionnaireItemType = "boolean"
	// Question with is a real number answer (valueDecimal).  There is an extension 'http://hl7.org/fhir/StructureDefinition/questionnaire-unit' that can be used to computably convey the unit of measure associated with the answer for use when performing data extraction to an element of type Quantity.
	QuestionnaireItemTypeDecimal QuestionnaireItemType = "decimal"
	// Question with an integer answer (valueInteger).  There is an extension 'http://hl7.org/fhir/StructureDefinition/questionnaire-unit' that can be used to computably convey the unit of measure associated with the answer for use when performing data extraction to an element of type Quantity.
	QuestionnaireItemTypeInteger QuestionnaireItemType = "integer"
	// Question with a date answer (valueDate).
	QuestionnaireItemTypeDate QuestionnaireItemType = "date"
	// Question with a date and time answer (valueDateTime).
	QuestionnaireItemTypeDateTime QuestionnaireItemType = "dateTime"
	// Question with a time (hour:minute:second) answer independent of date. (valueTime).
	QuestionnaireItemTypeTime QuestionnaireItemType = "time"
	// Question with a short (few words to short sentence) free-text entry answer (valueString).  Strings SHOULD NOT contain carriage return or newline characters.  If multi-line answers are needed, use the 'text' type.
	QuestionnaireItemTypeString QuestionnaireItemType = "string"
	// Question with a long (potentially multi-paragraph) free-text entry answer (valueString).
	QuestionnaireItemTypeText QuestionnaireItemType = "text"
	// Question with a URL (website, FTP site, etc.) answer (valueUri).
	QuestionnaireItemTypeUrl QuestionnaireItemType = "url"
	// Question with a Coding - generally drawn from a list of possible answers (valueCoding)
	QuestionnaireItemTypeCoding QuestionnaireItemType = "coding"
	// Question with binary content such as an image, PDF, etc. as an answer (valueAttachment).
	QuestionnaireItemTypeAttachment QuestionnaireItemType = "attachment"
	// Question with a reference to another resource (practitioner, organization, etc.) as an answer (valueReference).
	QuestionnaireItemTypeReference QuestionnaireItemType = "reference"
	// Question with a combination of a numeric value and unit as an answer. (valueSimpleQuantity)  There are two extensions ('http://hl7.org/fhir/StructureDefinition/questionnaire-unitOption' and 'http://hl7.org/fhir/StructureDefinition/questionnaire-unitValueSet')  that can be used to define what unit should be selected for the Quantity.code and Quantity.system.
	QuestionnaireItemTypeQuantity QuestionnaireItemType = "quantity"
)

func (code QuestionnaireItemType) String() string {
	return code.Code()
}

func (code QuestionnaireItemType) Code() string {
	return string(code)
}

func (code QuestionnaireItemType) Display() string {
	switch code {
	case QuestionnaireItemTypeGroup:
		return "Group"
	case QuestionnaireItemTypeDisplay:
		return "Display"
	case QuestionnaireItemTypeQuestion:
		return "Question"
	case QuestionnaireItemTypeBoolean:
		return "Boolean"
	case QuestionnaireItemTypeDecimal:
		return "Decimal"
	case QuestionnaireItemTypeInteger:
		return "Integer"
	case QuestionnaireItemTypeDate:
		return "Date"
	case QuestionnaireItemTypeDateTime:
		return "Date Time"
	case QuestionnaireItemTypeTime:
		return "Time"
	case QuestionnaireItemTypeString:
		return "String"
	case QuestionnaireItemTypeText:
		return "Text"
	case QuestionnaireItemTypeUrl:
		return "Url"
	case QuestionnaireItemTypeCoding:
		return "Coding"
	case QuestionnaireItemTypeAttachment:
		return "Attachment"
	case QuestionnaireItemTypeReference:
		return "Reference"
	case QuestionnaireItemTypeQuantity:
		return "Quantity"
	}
	return "<unknown>"
}

func (code QuestionnaireItemType) Definition() string {
	switch code {
	case QuestionnaireItemTypeGroup:
		return "An item with no direct answer but should have at least one child item."
	case QuestionnaireItemTypeDisplay:
		return "Text for display that will not capture an answer or have child items."
	case QuestionnaireItemTypeQuestion:
		return "An item that defines a specific answer to be captured, and which may have child items. (the answer provided in the QuestionnaireResponse should be of the defined datatype)."
	case QuestionnaireItemTypeBoolean:
		return "Question with a yes/no answer (valueBoolean)."
	case QuestionnaireItemTypeDecimal:
		return "Question with is a real number answer (valueDecimal).  There is an extension 'http://hl7.org/fhir/StructureDefinition/questionnaire-unit' that can be used to computably convey the unit of measure associated with the answer for use when performing data extraction to an element of type Quantity."
	case QuestionnaireItemTypeInteger:
		return "Question with an integer answer (valueInteger).  There is an extension 'http://hl7.org/fhir/StructureDefinition/questionnaire-unit' that can be used to computably convey the unit of measure associated with the answer for use when performing data extraction to an element of type Quantity."
	case QuestionnaireItemTypeDate:
		return "Question with a date answer (valueDate)."
	case QuestionnaireItemTypeDateTime:
		return "Question with a date and time answer (valueDateTime)."
	case QuestionnaireItemTypeTime:
		return "Question with a time (hour:minute:second) answer independent of date. (valueTime)."
	case QuestionnaireItemTypeString:
		return "Question with a short (few words to short sentence) free-text entry answer (valueString).  Strings SHOULD NOT contain carriage return or newline characters.  If multi-line answers are needed, use the 'text' type."
	case QuestionnaireItemTypeText:
		return "Question with a long (potentially multi-paragraph) free-text entry answer (valueString)."
	case QuestionnaireItemTypeUrl:
		return "Question with a URL (website, FTP site, etc.) answer (valueUri)."
	case QuestionnaireItemTypeCoding:
		return "Question with a Coding - generally drawn from a list of possible answers (valueCoding)"
	case QuestionnaireItemTypeAttachment:
		return "Question with binary content such as an image, PDF, etc. as an answer (valueAttachment)."
	case QuestionnaireItemTypeReference:
		return "Question with a reference to another resource (practitioner, organization, etc.) as an answer (valueReference)."
	case QuestionnaireItemTypeQuantity:
		return "Question with a combination of a numeric value and unit as an answer. (valueSimpleQuantity)  There are two extensions ('http://hl7.org/fhir/StructureDefinition/questionnaire-unitOption' and 'http://hl7.org/fhir/StructureDefinition/questionnaire-unitValueSet')  that can be used to define what unit should be selected for the Quantity.code and Quantity.system."
	}
	return "<unknown>"
}
