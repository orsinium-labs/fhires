package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// CharacteristicCombination is documented here http://hl7.org/fhir/ValueSet/characteristic-combination
type CharacteristicCombination int

const (
	CharacteristicCombinationAllOf CharacteristicCombination = iota
	CharacteristicCombinationAnyOf
	CharacteristicCombinationAtLeast
	CharacteristicCombinationAtMost
	CharacteristicCombinationStatistical
	CharacteristicCombinationNetEffect
	CharacteristicCombinationDataset
)

func (code CharacteristicCombination) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *CharacteristicCombination) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "all-of":
		*code = CharacteristicCombinationAllOf
	case "any-of":
		*code = CharacteristicCombinationAnyOf
	case "at-least":
		*code = CharacteristicCombinationAtLeast
	case "at-most":
		*code = CharacteristicCombinationAtMost
	case "statistical":
		*code = CharacteristicCombinationStatistical
	case "net-effect":
		*code = CharacteristicCombinationNetEffect
	case "dataset":
		*code = CharacteristicCombinationDataset
	default:
		return fmt.Errorf("unknown CharacteristicCombination code `%s`", s)
	}
	return nil
}

func (code CharacteristicCombination) String() string {
	return code.Code()
}

func (code CharacteristicCombination) Code() string {
	switch code {
	case CharacteristicCombinationAllOf:
		return "all-of"
	case CharacteristicCombinationAnyOf:
		return "any-of"
	case CharacteristicCombinationAtLeast:
		return "at-least"
	case CharacteristicCombinationAtMost:
		return "at-most"
	case CharacteristicCombinationStatistical:
		return "statistical"
	case CharacteristicCombinationNetEffect:
		return "net-effect"
	case CharacteristicCombinationDataset:
		return "dataset"
	}
	return "<unknown>"
}

func (code CharacteristicCombination) Display() string {
	switch code {
	case CharacteristicCombinationAllOf:
		return "All of"
	case CharacteristicCombinationAnyOf:
		return "Any of"
	case CharacteristicCombinationAtLeast:
		return "At least"
	case CharacteristicCombinationAtMost:
		return "At most"
	case CharacteristicCombinationStatistical:
		return "Statistical"
	case CharacteristicCombinationNetEffect:
		return "Net effect"
	case CharacteristicCombinationDataset:
		return "Dataset"
	}
	return "<unknown>"
}

func (code CharacteristicCombination) Definition() string {
	switch code {
	case CharacteristicCombinationAllOf:
		return "Combine characteristics with AND."
	case CharacteristicCombinationAnyOf:
		return "Combine characteristics with OR."
	case CharacteristicCombinationAtLeast:
		return "Meet at least the threshold number of characteristics for definition."
	case CharacteristicCombinationAtMost:
		return "Meet at most the threshold number of characteristics for definition."
	case CharacteristicCombinationStatistical:
		return "Combine characteristics statistically. Use method to specify the statistical method."
	case CharacteristicCombinationNetEffect:
		return "Combine characteristics by addition of benefits and subtraction of harms."
	case CharacteristicCombinationDataset:
		return "Combine characteristics as a collection used as the dataset."
	}
	return "<unknown>"
}
