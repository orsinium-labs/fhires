package fhires

// Meta Type: The metadata about a resource.
//
// This is content in the resource that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
//
// Status: active
//
// http://hl7.org/fhir/StructureDefinition/Meta
type Meta struct {
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
	// The version specific identifier, as it appears in the version portion of the URL.
	//
	// This value changes when the resource is created, updated, or deleted.
	//
	// Comment: The server assigns this value, and ignores what the client specifies, except in the case that the server is imposing version integrity on updates/deletes.
	VersionID *string `bson:"versionId,omitempty" json:"versionId,omitempty" xml:"versionId,omitempty"`
	// When the resource last changed - e.g. when the version changed.
	//
	// Comment: This element is generally omitted in instances submitted in a PUT or POST. Instead, it is populated in the response instance and when retrieving information using a GET. The server / resource manager sets this value; what a client provides is irrelevant. This is equivalent to the HTTP Last-Modified and SHOULD have the same value on a [read](http.html#read) interaction.
	LastUpdated *string `bson:"lastUpdated,omitempty" json:"lastUpdated,omitempty" xml:"lastUpdated,omitempty"`
	// A uri that identifies the source system of the resource.
	//
	// This provides a minimal amount of [Provenance](provenance.html#) information that can be used to track or differentiate the source of information in the resource. The source may identify another FHIR server, document, message, database, etc.
	//
	// Comment: The exact use of the source (and the possible implied Provenance.entity.role and agent.role) is left to implementer discretion. Only one nominated source is allowed; for additional provenance details, a full Provenance resource should be used. The source may correspond to Provenance.entity.what[x] or Provenance.agent.who[x], though it may be a more general or abstract reference.
	//
	// This element can be used to indicate where the current master source of a resource that has a canonical URL if the resource is no longer hosted at the canonical URL.
	Source *string `bson:"source,omitempty" json:"source,omitempty" xml:"source,omitempty"`
	// A list of profiles (references to [StructureDefinition](structuredefinition.html#) resources) that this resource claims to conform to.
	//
	// The URL is a reference to [StructureDefinition.url](structuredefinition-definitions.html#StructureDefinition.url).
	//
	// Comment: It is up to the server and/or other infrastructure of policy to determine whether/how these claims are verified and/or updated over time.  The list of profile URLs is a set.
	Profile []string `bson:"profile,omitempty" json:"profile,omitempty" xml:"profile,omitempty"`
	// Security labels applied to this resource.
	//
	// These tags connect specific resources to the overall security policy and infrastructure.
	//
	// Comment: The security labels can be updated without changing the stated version of the resource. The list of security labels is a set. Uniqueness is based the system/code, and version and display are ignored.
	Security []Coding `bson:"security,omitempty" json:"security,omitempty" xml:"security,omitempty"`
	// Tags applied to this resource.
	//
	// Tags are intended to be used to identify and relate resources to process and workflow, and applications are not required to consider the tags when interpreting the meaning of a resource.
	//
	// Comment: The tags can be updated without changing the stated version of the resource. The list of tags is a set. Uniqueness is based the system/code, and version and display are ignored.
	Tag []Coding `bson:"tag,omitempty" json:"tag,omitempty" xml:"tag,omitempty"`
}
