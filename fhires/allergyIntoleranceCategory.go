package fhires

// Category of an identified substance associated with allergies or intolerances.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/allergy-intolerance-category
type AllergyIntoleranceCategory string

const (
	// Any substance consumed to provide nutritional support for the body.
	AllergyIntoleranceCategoryFood AllergyIntoleranceCategory = "food"
	// Substances administered to achieve a physiological effect.
	AllergyIntoleranceCategoryMedication AllergyIntoleranceCategory = "medication"
	// Any substances that are encountered in the environment, including any substance not already classified as food, medication, or biologic.
	AllergyIntoleranceCategoryEnvironment AllergyIntoleranceCategory = "environment"
	// A preparation that is synthesized from living organisms or their products, especially a human or animal protein, such as a hormone or antitoxin, that is used as a diagnostic, preventive, or therapeutic agent. Examples of biologic medications include: vaccines; allergenic extracts, which are used for both diagnosis and treatment (for example, allergy shots); gene therapies; cellular therapies.  There are other biologic products, such as tissues, which are not typically associated with allergies.
	AllergyIntoleranceCategoryBiologic AllergyIntoleranceCategory = "biologic"
)

func (code AllergyIntoleranceCategory) String() string {
	return code.Code()
}

func (code AllergyIntoleranceCategory) Code() string {
	return string(code)
}

func (code AllergyIntoleranceCategory) Display() string {
	switch code {
	case AllergyIntoleranceCategoryFood:
		return "Food"
	case AllergyIntoleranceCategoryMedication:
		return "Medication"
	case AllergyIntoleranceCategoryEnvironment:
		return "Environment"
	case AllergyIntoleranceCategoryBiologic:
		return "Biologic"
	}
	return "<unknown>"
}

func (code AllergyIntoleranceCategory) Definition() string {
	switch code {
	case AllergyIntoleranceCategoryFood:
		return "Any substance consumed to provide nutritional support for the body."
	case AllergyIntoleranceCategoryMedication:
		return "Substances administered to achieve a physiological effect."
	case AllergyIntoleranceCategoryEnvironment:
		return "Any substances that are encountered in the environment, including any substance not already classified as food, medication, or biologic."
	case AllergyIntoleranceCategoryBiologic:
		return "A preparation that is synthesized from living organisms or their products, especially a human or animal protein, such as a hormone or antitoxin, that is used as a diagnostic, preventive, or therapeutic agent. Examples of biologic medications include: vaccines; allergenic extracts, which are used for both diagnosis and treatment (for example, allergy shots); gene therapies; cellular therapies.  There are other biologic products, such as tissues, which are not typically associated with allergies."
	}
	return "<unknown>"
}
