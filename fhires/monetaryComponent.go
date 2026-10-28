package fhires

import "github.com/shopspring/decimal"

// MonetaryComponent Type: Availability data for an {item}.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/MonetaryComponent
type MonetaryComponent struct {
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
	// base | surcharge | deduction | discount | tax | informational.
	Type PriceComponentType `bson:"type" json:"type" xml:"type"`
	// Codes may be used to differentiate between kinds of taxes, surcharges, discounts etc.
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// Factor used for calculating this component.
	Factor *decimal.Decimal `bson:"factor,omitempty" json:"factor,omitempty" xml:"factor,omitempty"`
	// Explicit value amount to be used.
	Amount *Money `bson:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
}
