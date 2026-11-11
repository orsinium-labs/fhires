package fhires

// The way in which this manufacturer is associated with the ingredient.
//
// For example whether it is a possible one (others allowed), or an exclusive authorized one for this ingredient. Note that this is not the manufacturing process role.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/ingredient-manufacturer-role
type IngredientManufacturerRole string

const (
	IngredientManufacturerRoleAllowed  IngredientManufacturerRole = "allowed"
	IngredientManufacturerRolePossible IngredientManufacturerRole = "possible"
	IngredientManufacturerRoleActual   IngredientManufacturerRole = "actual"
)

func (code IngredientManufacturerRole) String() string {
	return code.Code()
}

func (code IngredientManufacturerRole) Code() string {
	return string(code)
}

func (code IngredientManufacturerRole) Display() string {
	switch code {
	case IngredientManufacturerRoleAllowed:
		return "Manufacturer is specifically allowed for this ingredient"
	case IngredientManufacturerRolePossible:
		return "Manufacturer is known to make this ingredient in general"
	case IngredientManufacturerRoleActual:
		return "Manufacturer actually makes this particular ingredient"
	}
	return "<unknown>"
}

func (code IngredientManufacturerRole) Definition() string {
	switch code {
	}
	return "<unknown>"
}
