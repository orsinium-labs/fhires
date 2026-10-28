package fhires

// CodeableConcept Type: A concept that may be defined by a formal reference to a terminology or ontology or may be provided by text.
//
// Purpose: This is a common pattern in healthcare - a concept that may be defined by one or more codes from formal definitions including LOINC and SNOMED CT, and/or defined by the provision of text that captures a human sense of the concept.
//
// Status: active
//
// http://hl7.org/fhir/StructureDefinition/CodeableConcept
type CodeableConcept struct {
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
	// A reference to a code defined by a terminology system.
	//
	// Comment: Codes may be defined very casually in enumerations, or code lists, up to very formal definitions such as SNOMED CT - see the HL7 v3 Core Principles for more information.  Ordering of codings is undefined and SHALL NOT be used to infer meaning. Generally, at most only one of the coding values will be labeled as UserSelected = true.
	Coding []Coding `bson:"coding,omitempty" json:"coding,omitempty" xml:"coding,omitempty"`
	// A human language representation of the concept as seen/selected/uttered by the user who entered the data and/or which represents the intended meaning of the user.
	//
	// Comment: Very often the text is the same as a displayName of one of the codings.
	Text *string `bson:"text,omitempty" json:"text,omitempty" xml:"text,omitempty"`
}
