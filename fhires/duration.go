package fhires

import decimal "github.com/shopspring/decimal"

// Duration Type: A length of time.
//
// Status: active
//
// http://hl7.org/fhir/StructureDefinition/Duration
type Duration struct {
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
	// The value of the measured amount.
	//
	// The value includes an implicit precision in the presentation of the value.
	//
	// Comment: The implicit precision in the value should always be honored. Monetary values have their own rules for handling precision (refer to standard accounting text books).
	Value *decimal.Decimal `bson:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
	// How the value should be understood and represented - whether the actual value is greater or less than the stated value due to measurement issues; e.g.
	//
	// if the comparator is "<" , then the real value is < stated value.
	Comparator *QuantityComparator `bson:"comparator,omitempty" json:"comparator,omitempty" xml:"comparator,omitempty"`
	// A human-readable form of the unit.
	Unit *string `bson:"unit,omitempty" json:"unit,omitempty" xml:"unit,omitempty"`
	// The identification of the system that provides the coded form of the unit.
	System *string `bson:"system,omitempty" json:"system,omitempty" xml:"system,omitempty"`
	// A computer processable form of the unit in some unit representation system.
	//
	// Comment: The preferred system is UCUM, but SNOMED CT can also be used (for customary units) or ISO 4217 for currency.  The context of use may additionally require a code from a particular system.
	Code *string `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
}
