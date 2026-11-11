package fhires

// Codes indicating the kind of the price component.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/price-component-type
type PriceComponentType string

const (
	// the amount is the base price used for calculating the total price before applying surcharges, discount or taxes.
	PriceComponentTypeBase PriceComponentType = "base"
	// the amount is a surcharge applied on the base price.
	PriceComponentTypeSurcharge PriceComponentType = "surcharge"
	// the amount is a deduction applied on the base price.
	PriceComponentTypeDeduction PriceComponentType = "deduction"
	// the amount is a discount applied on the base price.
	PriceComponentTypeDiscount PriceComponentType = "discount"
	// the amount is the tax component of the total price.
	PriceComponentTypeTax PriceComponentType = "tax"
	// the amount is of informational character, it has not been applied in the calculation of the total price.
	PriceComponentTypeInformational PriceComponentType = "informational"
)

func (code PriceComponentType) String() string {
	return code.Code()
}

func (code PriceComponentType) Code() string {
	return string(code)
}

func (code PriceComponentType) Display() string {
	switch code {
	case PriceComponentTypeBase:
		return "base price"
	case PriceComponentTypeSurcharge:
		return "surcharge"
	case PriceComponentTypeDeduction:
		return "deduction"
	case PriceComponentTypeDiscount:
		return "discount"
	case PriceComponentTypeTax:
		return "tax"
	case PriceComponentTypeInformational:
		return "informational"
	}
	return "<unknown>"
}

func (code PriceComponentType) Definition() string {
	switch code {
	case PriceComponentTypeBase:
		return "the amount is the base price used for calculating the total price before applying surcharges, discount or taxes."
	case PriceComponentTypeSurcharge:
		return "the amount is a surcharge applied on the base price."
	case PriceComponentTypeDeduction:
		return "the amount is a deduction applied on the base price."
	case PriceComponentTypeDiscount:
		return "the amount is a discount applied on the base price."
	case PriceComponentTypeTax:
		return "the amount is the tax component of the total price."
	case PriceComponentTypeInformational:
		return "the amount is of informational character, it has not been applied in the calculation of the total price."
	}
	return "<unknown>"
}
