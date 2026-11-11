package fhires

// Overall defining type of this clinical use definition.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/clinical-use-definition-type
type ClinicalUseDefinitionType string

const (
	// A reason for giving the medication.
	ClinicalUseDefinitionTypeIndication ClinicalUseDefinitionType = "indication"
	// A reason for not giving the medication.
	ClinicalUseDefinitionTypeContraindication ClinicalUseDefinitionType = "contraindication"
	// Interactions between the medication and other substances.
	ClinicalUseDefinitionTypeInteraction ClinicalUseDefinitionType = "interaction"
	// Side effects or adverse effects associated with the medication.
	ClinicalUseDefinitionTypeUndesirableEffect ClinicalUseDefinitionType = "undesirable-effect"
	// A general warning or issue that is not specifically one of the other types.
	ClinicalUseDefinitionTypeWarning ClinicalUseDefinitionType = "warning"
)

func (code ClinicalUseDefinitionType) String() string {
	return code.Code()
}

func (code ClinicalUseDefinitionType) Code() string {
	return string(code)
}

func (code ClinicalUseDefinitionType) Display() string {
	switch code {
	case ClinicalUseDefinitionTypeIndication:
		return "Indication"
	case ClinicalUseDefinitionTypeContraindication:
		return "Contraindication"
	case ClinicalUseDefinitionTypeInteraction:
		return "Interaction"
	case ClinicalUseDefinitionTypeUndesirableEffect:
		return "Undesirable Effect"
	case ClinicalUseDefinitionTypeWarning:
		return "Warning"
	}
	return "<unknown>"
}

func (code ClinicalUseDefinitionType) Definition() string {
	switch code {
	case ClinicalUseDefinitionTypeIndication:
		return "A reason for giving the medication."
	case ClinicalUseDefinitionTypeContraindication:
		return "A reason for not giving the medication."
	case ClinicalUseDefinitionTypeInteraction:
		return "Interactions between the medication and other substances."
	case ClinicalUseDefinitionTypeUndesirableEffect:
		return "Side effects or adverse effects associated with the medication."
	case ClinicalUseDefinitionTypeWarning:
		return "A general warning or issue that is not specifically one of the other types."
	}
	return "<unknown>"
}
