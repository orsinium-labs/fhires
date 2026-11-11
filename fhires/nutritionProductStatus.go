package fhires

// Codes identifying the lifecycle stage of a product.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/nutritionproduct-status
type NutritionProductStatus string

const (
	// The product can be used.
	NutritionProductStatusActive NutritionProductStatus = "active"
	// The product is not expected or allowed to be used.
	NutritionProductStatusInactive NutritionProductStatus = "inactive"
	// This electronic record should never have existed, though it is possible that real-world decisions were based on it.  (If real-world activity has occurred, the status should be "cancelled" rather than "entered-in-error".).
	NutritionProductStatusEnteredInError NutritionProductStatus = "entered-in-error"
)

func (code NutritionProductStatus) String() string {
	return code.Code()
}

func (code NutritionProductStatus) Code() string {
	return string(code)
}

func (code NutritionProductStatus) Display() string {
	switch code {
	case NutritionProductStatusActive:
		return "Active"
	case NutritionProductStatusInactive:
		return "Inactive"
	case NutritionProductStatusEnteredInError:
		return "Entered in Error"
	}
	return "<unknown>"
}

func (code NutritionProductStatus) Definition() string {
	switch code {
	case NutritionProductStatusActive:
		return "The product can be used."
	case NutritionProductStatusInactive:
		return "The product is not expected or allowed to be used."
	case NutritionProductStatusEnteredInError:
		return "This electronic record should never have existed, though it is possible that real-world decisions were based on it.  (If real-world activity has occurred, the status should be \"cancelled\" rather than \"entered-in-error\".)."
	}
	return "<unknown>"
}
