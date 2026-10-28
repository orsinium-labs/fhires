package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// Overall defining type of this clinical use definition.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/clinical-use-definition-type
type ClinicalUseDefinitionType int

const (
	ClinicalUseDefinitionTypeIndication ClinicalUseDefinitionType = iota
	// A reason for not giving the medication.
	ClinicalUseDefinitionTypeContraindication
	// Interactions between the medication and other substances.
	ClinicalUseDefinitionTypeInteraction
	// Side effects or adverse effects associated with the medication.
	ClinicalUseDefinitionTypeUndesirableEffect
	// A general warning or issue that is not specifically one of the other types.
	ClinicalUseDefinitionTypeWarning
)

func (code ClinicalUseDefinitionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *ClinicalUseDefinitionType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "indication":
		*code = ClinicalUseDefinitionTypeIndication
	case "contraindication":
		*code = ClinicalUseDefinitionTypeContraindication
	case "interaction":
		*code = ClinicalUseDefinitionTypeInteraction
	case "undesirable-effect":
		*code = ClinicalUseDefinitionTypeUndesirableEffect
	case "warning":
		*code = ClinicalUseDefinitionTypeWarning
	default:
		return fmt.Errorf("unknown ClinicalUseDefinitionType code `%s`", s)
	}
	return nil
}

func (code ClinicalUseDefinitionType) String() string {
	return code.Code()
}

func (code ClinicalUseDefinitionType) Code() string {
	switch code {
	case ClinicalUseDefinitionTypeIndication:
		return "indication"
	case ClinicalUseDefinitionTypeContraindication:
		return "contraindication"
	case ClinicalUseDefinitionTypeInteraction:
		return "interaction"
	case ClinicalUseDefinitionTypeUndesirableEffect:
		return "undesirable-effect"
	case ClinicalUseDefinitionTypeWarning:
		return "warning"
	}
	return "<unknown>"
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
