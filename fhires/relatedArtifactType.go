package fhires

// The type of relationship to the related artifact.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/related-artifact-type
type RelatedArtifactType string

const (
	// Additional documentation for the knowledge resource. This would include additional instructions on usage as well as additional information on clinical context or appropriateness.
	RelatedArtifactTypeDocumentation RelatedArtifactType = "documentation"
	// The target artifact is a summary of the justification for the knowledge resource including supporting evidence, relevant guidelines, or other clinically important information. This information is intended to provide a way to make the justification for the knowledge resource available to the consumer of interventions or results produced by the knowledge resource.
	RelatedArtifactTypeJustification RelatedArtifactType = "justification"
	// Bibliographic citation for papers, references, or other relevant material for the knowledge resource. This is intended to allow for citation of related material, but that was not necessarily specifically prepared in connection with this knowledge resource.
	RelatedArtifactTypeCitation RelatedArtifactType = "citation"
	// The previous version of the knowledge artifact, used to establish an ordering of versions of an artifact, independent of the status of each version.
	RelatedArtifactTypePredecessor RelatedArtifactType = "predecessor"
	// The subsequent version of the knowledge artfact, used to establish an ordering of versions of an artifact, independent of the status of each version.
	RelatedArtifactTypeSuccessor RelatedArtifactType = "successor"
	// This artifact is derived from the target artifact. This is intended to capture the relationship in which a particular knowledge resource is based on the content of another artifact, but is modified to capture either a different set of overall requirements, or a more specific set of requirements such as those involved in a particular institution or clinical setting. The artifact may be derived from one or more target artifacts.
	RelatedArtifactTypeDerivedFrom RelatedArtifactType = "derived-from"
	// This artifact depends on the target artifact. There is a requirement to use the target artifact in the creation or interpretation of this artifact.
	RelatedArtifactTypeDependsOn RelatedArtifactType = "depends-on"
	// This artifact is composed of the target artifact. This artifact is constructed with the target artifact as a component. The target artifact is a part of this artifact. (A dataset is composed of data.).
	RelatedArtifactTypeComposedOf RelatedArtifactType = "composed-of"
	// This artifact is a part of the target artifact. The target artifact is composed of this artifact (and possibly other artifacts).
	RelatedArtifactTypePartOf RelatedArtifactType = "part-of"
	// This artifact amends or changes the target artifact. This artifact adds additional information that is functionally expected to replace information in the target artifact. This artifact replaces a part but not all of the target artifact.
	RelatedArtifactTypeAmends RelatedArtifactType = "amends"
	// This artifact is amended with or changed by the target artifact. There is information in this artifact that should be functionally replaced with information in the target artifact.
	RelatedArtifactTypeAmendedWith RelatedArtifactType = "amended-with"
	// This artifact adds additional information to the target artifact. The additional information does not replace or change information in the target artifact.
	RelatedArtifactTypeAppends RelatedArtifactType = "appends"
	// This artifact has additional information in the target artifact.
	RelatedArtifactTypeAppendedWith RelatedArtifactType = "appended-with"
	// This artifact cites the target artifact. This may be a bibliographic citation for papers, references, or other relevant material for the knowledge resource. This is intended to allow for citation of related material, but that was not necessarily specifically prepared in connection with this knowledge resource.
	RelatedArtifactTypeCites RelatedArtifactType = "cites"
	// This artifact is cited by the target artifact.
	RelatedArtifactTypeCitedBy RelatedArtifactType = "cited-by"
	// This artifact contains comments about the target artifact.
	RelatedArtifactTypeCommentsOn RelatedArtifactType = "comments-on"
	// This artifact has comments about it in the target artifact.  The type of comments may be expressed in the targetClassifier element such as reply, review, editorial, feedback, solicited, unsolicited, structured, unstructured.
	RelatedArtifactTypeCommentIn RelatedArtifactType = "comment-in"
	// This artifact is a container in which the target artifact is contained. A container is a data structure whose instances are collections of other objects. (A database contains the dataset.).
	RelatedArtifactTypeContains RelatedArtifactType = "contains"
	// This artifact is contained in the target artifact. The target artifact is a data structure whose instances are collections of other objects.
	RelatedArtifactTypeContainedIn RelatedArtifactType = "contained-in"
	// This artifact identifies errors and replacement content for the target artifact.
	RelatedArtifactTypeCorrects RelatedArtifactType = "corrects"
	// This artifact has corrections to it in the target artifact. The target artifact identifies errors and replacement content for this artifact.
	RelatedArtifactTypeCorrectionIn RelatedArtifactType = "correction-in"
	// This artifact replaces or supersedes the target artifact. The target artifact may be considered deprecated.
	RelatedArtifactTypeReplaces RelatedArtifactType = "replaces"
	// This artifact is replaced with or superseded by the target artifact. This artifact may be considered deprecated.
	RelatedArtifactTypeReplacedWith RelatedArtifactType = "replaced-with"
	// This artifact retracts the target artifact. The content that was published in the target artifact should be considered removed from publication and should no longer be considered part of the public record.
	RelatedArtifactTypeRetracts RelatedArtifactType = "retracts"
	// This artifact is retracted by the target artifact. The content that was published in this artifact should be considered removed from publication and should no longer be considered part of the public record.
	RelatedArtifactTypeRetractedBy RelatedArtifactType = "retracted-by"
	// This artifact is a signature of the target artifact.
	RelatedArtifactTypeSigns RelatedArtifactType = "signs"
	// This artifact has characteristics in common with the target artifact. This relationship may be used in systems to “deduplicate” knowledge artifacts from different sources, or in systems to show “similar items”.
	RelatedArtifactTypeSimilarTo RelatedArtifactType = "similar-to"
	// This artifact provides additional support for the target artifact. The type of support  is not documentation as it does not describe, explain, or instruct regarding the target artifact.
	RelatedArtifactTypeSupports RelatedArtifactType = "supports"
	// The target artifact contains additional information related to the knowledge artifact but is not documentation as the additional information does not describe, explain, or instruct regarding the knowledge artifact content or application. This could include an associated dataset.
	RelatedArtifactTypeSupportedWith RelatedArtifactType = "supported-with"
	// This artifact was generated by transforming the target artifact (e.g., format or language conversion). This is intended to capture the relationship in which a particular knowledge resource is based on the content of another artifact, but changes are only apparent in form and there is only one target artifact with the “transforms” relationship type.
	RelatedArtifactTypeTransforms RelatedArtifactType = "transforms"
	// This artifact was transformed into the target artifact (e.g., by format or language conversion).
	RelatedArtifactTypeTransformedInto RelatedArtifactType = "transformed-into"
	// This artifact was generated by transforming a related artifact (e.g., format or language conversion), noted separately with the “transforms” relationship type. This transformation used the target artifact to inform the transformation. The target artifact may be a conversion script or translation guide.
	RelatedArtifactTypeTransformedWith RelatedArtifactType = "transformed-with"
	// This artifact provides additional documentation for the target artifact. This could include additional instructions on usage as well as additional information on clinical context or appropriateness.
	RelatedArtifactTypeDocuments RelatedArtifactType = "documents"
	// The target artifact is a precise description of a concept in this artifact. This may be used when the RelatedArtifact datatype is used in elements contained in this artifact.
	RelatedArtifactTypeSpecificationOf RelatedArtifactType = "specification-of"
	// This artifact was created with the target artifact. The target artifact is a tool or support material used in the creation of the artifact, and not content that the artifact was derived from.
	RelatedArtifactTypeCreatedWith RelatedArtifactType = "created-with"
	// The related artifact is the citation for this artifact.
	RelatedArtifactTypeCiteAs RelatedArtifactType = "cite-as"
)

func (code RelatedArtifactType) String() string {
	return code.Code()
}

func (code RelatedArtifactType) Code() string {
	return string(code)
}

func (code RelatedArtifactType) Display() string {
	switch code {
	case RelatedArtifactTypeDocumentation:
		return "Documentation"
	case RelatedArtifactTypeJustification:
		return "Justification"
	case RelatedArtifactTypeCitation:
		return "Citation"
	case RelatedArtifactTypePredecessor:
		return "Predecessor"
	case RelatedArtifactTypeSuccessor:
		return "Successor"
	case RelatedArtifactTypeDerivedFrom:
		return "Derived From"
	case RelatedArtifactTypeDependsOn:
		return "Depends On"
	case RelatedArtifactTypeComposedOf:
		return "Composed Of"
	case RelatedArtifactTypePartOf:
		return "Part Of"
	case RelatedArtifactTypeAmends:
		return "Amends"
	case RelatedArtifactTypeAmendedWith:
		return "Amended With"
	case RelatedArtifactTypeAppends:
		return "Appends"
	case RelatedArtifactTypeAppendedWith:
		return "Appended With"
	case RelatedArtifactTypeCites:
		return "Cites"
	case RelatedArtifactTypeCitedBy:
		return "Cited By"
	case RelatedArtifactTypeCommentsOn:
		return "Is Comment On"
	case RelatedArtifactTypeCommentIn:
		return "Has Comment In"
	case RelatedArtifactTypeContains:
		return "Contains"
	case RelatedArtifactTypeContainedIn:
		return "Contained In"
	case RelatedArtifactTypeCorrects:
		return "Corrects"
	case RelatedArtifactTypeCorrectionIn:
		return "Correction In"
	case RelatedArtifactTypeReplaces:
		return "Replaces"
	case RelatedArtifactTypeReplacedWith:
		return "Replaced With"
	case RelatedArtifactTypeRetracts:
		return "Retracts"
	case RelatedArtifactTypeRetractedBy:
		return "Retracted By"
	case RelatedArtifactTypeSigns:
		return "Signs"
	case RelatedArtifactTypeSimilarTo:
		return "Similar To"
	case RelatedArtifactTypeSupports:
		return "Supports"
	case RelatedArtifactTypeSupportedWith:
		return "Supported With"
	case RelatedArtifactTypeTransforms:
		return "Transforms"
	case RelatedArtifactTypeTransformedInto:
		return "Transformed Into"
	case RelatedArtifactTypeTransformedWith:
		return "Transformed With"
	case RelatedArtifactTypeDocuments:
		return "Documents"
	case RelatedArtifactTypeSpecificationOf:
		return "Specification Of"
	case RelatedArtifactTypeCreatedWith:
		return "Created With"
	case RelatedArtifactTypeCiteAs:
		return "Cite As"
	}
	return "<unknown>"
}

func (code RelatedArtifactType) Definition() string {
	switch code {
	case RelatedArtifactTypeDocumentation:
		return "Additional documentation for the knowledge resource. This would include additional instructions on usage as well as additional information on clinical context or appropriateness."
	case RelatedArtifactTypeJustification:
		return "The target artifact is a summary of the justification for the knowledge resource including supporting evidence, relevant guidelines, or other clinically important information. This information is intended to provide a way to make the justification for the knowledge resource available to the consumer of interventions or results produced by the knowledge resource."
	case RelatedArtifactTypeCitation:
		return "Bibliographic citation for papers, references, or other relevant material for the knowledge resource. This is intended to allow for citation of related material, but that was not necessarily specifically prepared in connection with this knowledge resource."
	case RelatedArtifactTypePredecessor:
		return "The previous version of the knowledge artifact, used to establish an ordering of versions of an artifact, independent of the status of each version."
	case RelatedArtifactTypeSuccessor:
		return "The subsequent version of the knowledge artfact, used to establish an ordering of versions of an artifact, independent of the status of each version."
	case RelatedArtifactTypeDerivedFrom:
		return "This artifact is derived from the target artifact. This is intended to capture the relationship in which a particular knowledge resource is based on the content of another artifact, but is modified to capture either a different set of overall requirements, or a more specific set of requirements such as those involved in a particular institution or clinical setting. The artifact may be derived from one or more target artifacts."
	case RelatedArtifactTypeDependsOn:
		return "This artifact depends on the target artifact. There is a requirement to use the target artifact in the creation or interpretation of this artifact."
	case RelatedArtifactTypeComposedOf:
		return "This artifact is composed of the target artifact. This artifact is constructed with the target artifact as a component. The target artifact is a part of this artifact. (A dataset is composed of data.)."
	case RelatedArtifactTypePartOf:
		return "This artifact is a part of the target artifact. The target artifact is composed of this artifact (and possibly other artifacts)."
	case RelatedArtifactTypeAmends:
		return "This artifact amends or changes the target artifact. This artifact adds additional information that is functionally expected to replace information in the target artifact. This artifact replaces a part but not all of the target artifact."
	case RelatedArtifactTypeAmendedWith:
		return "This artifact is amended with or changed by the target artifact. There is information in this artifact that should be functionally replaced with information in the target artifact."
	case RelatedArtifactTypeAppends:
		return "This artifact adds additional information to the target artifact. The additional information does not replace or change information in the target artifact."
	case RelatedArtifactTypeAppendedWith:
		return "This artifact has additional information in the target artifact."
	case RelatedArtifactTypeCites:
		return "This artifact cites the target artifact. This may be a bibliographic citation for papers, references, or other relevant material for the knowledge resource. This is intended to allow for citation of related material, but that was not necessarily specifically prepared in connection with this knowledge resource."
	case RelatedArtifactTypeCitedBy:
		return "This artifact is cited by the target artifact."
	case RelatedArtifactTypeCommentsOn:
		return "This artifact contains comments about the target artifact."
	case RelatedArtifactTypeCommentIn:
		return "This artifact has comments about it in the target artifact.  The type of comments may be expressed in the targetClassifier element such as reply, review, editorial, feedback, solicited, unsolicited, structured, unstructured."
	case RelatedArtifactTypeContains:
		return "This artifact is a container in which the target artifact is contained. A container is a data structure whose instances are collections of other objects. (A database contains the dataset.)."
	case RelatedArtifactTypeContainedIn:
		return "This artifact is contained in the target artifact. The target artifact is a data structure whose instances are collections of other objects."
	case RelatedArtifactTypeCorrects:
		return "This artifact identifies errors and replacement content for the target artifact."
	case RelatedArtifactTypeCorrectionIn:
		return "This artifact has corrections to it in the target artifact. The target artifact identifies errors and replacement content for this artifact."
	case RelatedArtifactTypeReplaces:
		return "This artifact replaces or supersedes the target artifact. The target artifact may be considered deprecated."
	case RelatedArtifactTypeReplacedWith:
		return "This artifact is replaced with or superseded by the target artifact. This artifact may be considered deprecated."
	case RelatedArtifactTypeRetracts:
		return "This artifact retracts the target artifact. The content that was published in the target artifact should be considered removed from publication and should no longer be considered part of the public record."
	case RelatedArtifactTypeRetractedBy:
		return "This artifact is retracted by the target artifact. The content that was published in this artifact should be considered removed from publication and should no longer be considered part of the public record."
	case RelatedArtifactTypeSigns:
		return "This artifact is a signature of the target artifact."
	case RelatedArtifactTypeSimilarTo:
		return "This artifact has characteristics in common with the target artifact. This relationship may be used in systems to “deduplicate” knowledge artifacts from different sources, or in systems to show “similar items”."
	case RelatedArtifactTypeSupports:
		return "This artifact provides additional support for the target artifact. The type of support  is not documentation as it does not describe, explain, or instruct regarding the target artifact."
	case RelatedArtifactTypeSupportedWith:
		return "The target artifact contains additional information related to the knowledge artifact but is not documentation as the additional information does not describe, explain, or instruct regarding the knowledge artifact content or application. This could include an associated dataset."
	case RelatedArtifactTypeTransforms:
		return "This artifact was generated by transforming the target artifact (e.g., format or language conversion). This is intended to capture the relationship in which a particular knowledge resource is based on the content of another artifact, but changes are only apparent in form and there is only one target artifact with the “transforms” relationship type."
	case RelatedArtifactTypeTransformedInto:
		return "This artifact was transformed into the target artifact (e.g., by format or language conversion)."
	case RelatedArtifactTypeTransformedWith:
		return "This artifact was generated by transforming a related artifact (e.g., format or language conversion), noted separately with the “transforms” relationship type. This transformation used the target artifact to inform the transformation. The target artifact may be a conversion script or translation guide."
	case RelatedArtifactTypeDocuments:
		return "This artifact provides additional documentation for the target artifact. This could include additional instructions on usage as well as additional information on clinical context or appropriateness."
	case RelatedArtifactTypeSpecificationOf:
		return "The target artifact is a precise description of a concept in this artifact. This may be used when the RelatedArtifact datatype is used in elements contained in this artifact."
	case RelatedArtifactTypeCreatedWith:
		return "This artifact was created with the target artifact. The target artifact is a tool or support material used in the creation of the artifact, and not content that the artifact was derived from."
	case RelatedArtifactTypeCiteAs:
		return "The related artifact is the citation for this artifact."
	}
	return "<unknown>"
}
