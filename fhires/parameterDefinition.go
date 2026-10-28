package fhires

// ParameterDefinition Type: The parameters to the module.
//
// This collection specifies both the input and output parameters. Input parameters are provided by the caller as part of the $evaluate operation. Output parameters are included in the GuidanceResponse.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/ParameterDefinition
type ParameterDefinition struct {
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
	// The name of the parameter used to allow access to the value of the parameter in evaluation contexts.
	Name *string `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Whether the parameter is input or output for the module.
	Use OperationParameterUse `bson:"use" json:"use" xml:"use"`
	// The minimum number of times this parameter SHALL appear in the request or response.
	Min *int `bson:"min,omitempty" json:"min,omitempty" xml:"min,omitempty"`
	// The maximum number of times this element is permitted to appear in the request or response.
	Max *string `bson:"max,omitempty" json:"max,omitempty" xml:"max,omitempty"`
	// A brief discussion of what the parameter is for and how it is used by the module.
	Documentation *string `bson:"documentation,omitempty" json:"documentation,omitempty" xml:"documentation,omitempty"`
	// The type of the parameter.
	Type FHIRTypes `bson:"type" json:"type" xml:"type"`
	// If specified, this indicates a profile that the input data must conform to, or that the output data will conform to.
	Profile *string `bson:"profile,omitempty" json:"profile,omitempty" xml:"profile,omitempty"`
}
