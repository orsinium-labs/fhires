package fhires

// RelatedArtifact Type: Related artifacts such as additional documentation, justification, or bibliographic references.
//
// Purpose: Knowledge resources must be able to provide enough information for consumers of the content (and/or interventions or results produced by the content) to be able to determine and understand the justification for and evidence in support of the content.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/RelatedArtifact
type RelatedArtifact struct {
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
	// The type of relationship to the related artifact.
	//
	// Comment: The presence of both sides of a relationship type (e.g. successor and predecessor) is required to support use cases where one side of a relationship is not represented in FHIR. However, this feature SHALL NOT be used to create bi-directional resource links in FHIR instances. Specifically, following the methodology of "new points to old" and "many points to one", when using related artifact elements to describe and reference FHIR resources, the type element SHALL be drawn from the fhir-related-artifact-type ValueSet.
	Type RelatedArtifactType `bson:"type" json:"type" xml:"type"`
	// Provides additional classifiers of the related artifact.
	Classifier []CodeableConcept `bson:"classifier,omitempty" json:"classifier,omitempty" xml:"classifier,omitempty"`
	// A short label that can be used to reference the citation from elsewhere in the containing artifact, such as a footnote index.
	Label *string `bson:"label,omitempty" json:"label,omitempty" xml:"label,omitempty"`
	// A brief description of the document or knowledge resource being referenced, suitable for display to a consumer.
	Display *string `bson:"display,omitempty" json:"display,omitempty" xml:"display,omitempty"`
	// A bibliographic citation for the related artifact.
	//
	// This text SHOULD be formatted according to an accepted citation format.
	//
	// Comment: Additional structured information about citations should be captured as extensions.
	Citation *string `bson:"citation,omitempty" json:"citation,omitempty" xml:"citation,omitempty"`
	// The document being referenced, represented as an attachment.
	//
	// This is exclusive with the resource element.
	Document *Attachment `bson:"document,omitempty" json:"document,omitempty" xml:"document,omitempty"`
	// The related artifact, such as a library, value set, profile, or other knowledge resource.
	//
	// Comment: If the type is predecessor, this is a reference to the succeeding knowledge resource. If the type is successor, this is a reference to the prior knowledge resource.
	Resource *string `bson:"resource,omitempty" json:"resource,omitempty" xml:"resource,omitempty"`
	// The related artifact, if the artifact is not a canonical resource, or a resource reference to a canonical resource.
	//
	// Comment: If both resource and resourceReference are present, they SHOULD be consistent and reference the same resource. Although relatedArtifact is intended to support references to definitional resources, there are cases where non-definitional resources can be definitional (such as Location where the kind is mode). Related artifacts SHOULD be used to reference definitional resources, and profiles SHOULD be used to make that explicit for particular use cases.
	ResourceReference *Reference `bson:"resourceReference,omitempty" json:"resourceReference,omitempty" xml:"resourceReference,omitempty"`
	// The publication status of the artifact being referred to.
	PublicationStatus *PublicationStatus `bson:"publicationStatus,omitempty" json:"publicationStatus,omitempty" xml:"publicationStatus,omitempty"`
	// The date of publication of the artifact being referred to.
	PublicationDate *string `bson:"publicationDate,omitempty" json:"publicationDate,omitempty" xml:"publicationDate,omitempty"`
}
