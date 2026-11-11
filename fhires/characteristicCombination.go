package fhires

// Logical grouping of characteristics.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/characteristic-combination
type CharacteristicCombination string

const (
	// Combine characteristics with AND.
	CharacteristicCombinationAllOf CharacteristicCombination = "all-of"
	// Combine characteristics with OR.
	CharacteristicCombinationAnyOf CharacteristicCombination = "any-of"
	// Meet at least the threshold number of characteristics for definition.
	CharacteristicCombinationAtLeast CharacteristicCombination = "at-least"
	// Meet at most the threshold number of characteristics for definition.
	CharacteristicCombinationAtMost CharacteristicCombination = "at-most"
	// Combine characteristics statistically. Use method to specify the statistical method.
	CharacteristicCombinationStatistical CharacteristicCombination = "statistical"
	// Combine characteristics by addition of benefits and subtraction of harms.
	CharacteristicCombinationNetEffect CharacteristicCombination = "net-effect"
	// Combine characteristics as a collection used as the dataset.
	CharacteristicCombinationDataset CharacteristicCombination = "dataset"
)

func (code CharacteristicCombination) String() string {
	return code.Code()
}

func (code CharacteristicCombination) Code() string {
	return string(code)
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
