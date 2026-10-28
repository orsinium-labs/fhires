package fhires

// Coding Type: A reference to a code defined by a terminology system.
//
// Purpose: References to codes are very common in healthcare models.
//
// Status: active
//
// http://hl7.org/fhir/StructureDefinition/Coding
type Coding struct {
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
	// The identification of the code system that defines the meaning of the symbol in the code.
	//
	// Comment: The URI may be an OID (urn:oid:...) or a UUID (urn:uuid:...).  OIDs and UUIDs SHALL be references to the HL7 OID registry. Otherwise, the URI should come from HL7's list of FHIR defined special URIs or it should be an absolute reference to some definition that establishes the system clearly and unambiguously.
	System *string `bson:"system,omitempty" json:"system,omitempty" xml:"system,omitempty"`
	// The version of the code system which was used when choosing this code.
	//
	// Note that a well-maintained code system does not need the version reported, because the meaning of codes is consistent across versions. However this cannot consistently be assured, and when the meaning is not guaranteed to be consistent, the version SHOULD be exchanged.
	//
	// Comment: Where the terminology does not clearly define what string should be used to identify code system versions, the recommendation is to use the date (expressed in FHIR date format) on which that version was officially published as the version date.
	Version *string `bson:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
	// A symbol in syntax defined by the system.
	//
	// The symbol may be a predefined code or an expression in a syntax defined by the coding system (e.g. post-coordination).
	Code *string `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// A representation of the meaning of the code in the system, following the rules of the system.
	Display *string `bson:"display,omitempty" json:"display,omitempty" xml:"display,omitempty"`
	// Indicates that this coding was chosen by a user directly - e.g.
	//
	// off a pick list of available items (codes or displays).
	//
	// Comment: Amongst a set of alternatives, a directly chosen code is the most appropriate starting point for new translations. There is some ambiguity about what exactly 'directly chosen' implies, and trading partner agreement may be needed to clarify the use of this element and its consequences more completely.
	UserSelected *bool `bson:"userSelected,omitempty" json:"userSelected,omitempty" xml:"userSelected,omitempty"`
}
