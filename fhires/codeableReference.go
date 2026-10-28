package fhires

// CodeableReference Type: A reference to a resource (by instance), or instead, a reference to a concept defined in a terminology or ontology (by class).
//
// Purpose: This is a common pattern in record keeping - a reference may be made to a specific condition, observation, plan, or definition, or a reference may be made to a general concept defined in a knowledge base somewhere.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/CodeableReference
type CodeableReference struct {
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
	// A reference to a concept - e.g.
	//
	// the information is identified by its general class to the degree of precision found in the terminology.
	Concept *CodeableConcept `bson:"concept,omitempty" json:"concept,omitempty" xml:"concept,omitempty"`
	// A reference to a resource the provides exact details about the information being referenced.
	Reference *Reference `bson:"reference,omitempty" json:"reference,omitempty" xml:"reference,omitempty"`
}
