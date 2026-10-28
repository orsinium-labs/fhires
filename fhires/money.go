package fhires

import "github.com/shopspring/decimal"

// Money Type: An amount of economic utility in some recognized currency.
//
// Status: active
//
// http://hl7.org/fhir/StructureDefinition/Money
type Money struct {
	// Unique id for the element within a resource (for internal references).
	//
	// This may be any string value that does not contain spaces.
	ID *string `bson:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element.
	//
	// To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	//
	// Comment: There can be no stigma associated with the use of extensions by any application, project, or standard - regardless of the institution or jurisdiction that uses or defines the extensions.  The use of extensions is what allows the FHIR specification to retain a core level of simplicity for everyone.
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty" xml:"extension,omitempty"`
	// Numerical value (with implicit precision).
	//
	// Comment: Monetary values have their own rules for handling precision (refer to standard accounting text books).
	Value *decimal.Decimal `bson:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
	// ISO 4217 Currency Code.
	Currency *string `bson:"currency,omitempty" json:"currency,omitempty" xml:"currency,omitempty"`
}
