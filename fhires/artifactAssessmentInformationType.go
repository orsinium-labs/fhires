package fhires

// The type of information contained in a component of an artifact assessment.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/artifactassessment-information-type
type ArtifactAssessmentInformationType string

const (
	// A comment on the artifact
	ArtifactAssessmentInformationTypeComment ArtifactAssessmentInformationType = "comment"
	// A classifier of the artifact
	ArtifactAssessmentInformationTypeClassifier ArtifactAssessmentInformationType = "classifier"
	// A rating of the artifact
	ArtifactAssessmentInformationTypeRating ArtifactAssessmentInformationType = "rating"
	// A container for multiple components
	ArtifactAssessmentInformationTypeContainer ArtifactAssessmentInformationType = "container"
	// A response to a comment
	ArtifactAssessmentInformationTypeResponse ArtifactAssessmentInformationType = "response"
	// A change request for the artifact
	ArtifactAssessmentInformationTypeChangeRequest ArtifactAssessmentInformationType = "change-request"
)

func (code ArtifactAssessmentInformationType) String() string {
	return code.Code()
}

func (code ArtifactAssessmentInformationType) Code() string {
	return string(code)
}

func (code ArtifactAssessmentInformationType) Display() string {
	switch code {
	case ArtifactAssessmentInformationTypeComment:
		return "Comment"
	case ArtifactAssessmentInformationTypeClassifier:
		return "Classifier"
	case ArtifactAssessmentInformationTypeRating:
		return "Rating"
	case ArtifactAssessmentInformationTypeContainer:
		return "Container"
	case ArtifactAssessmentInformationTypeResponse:
		return "Response"
	case ArtifactAssessmentInformationTypeChangeRequest:
		return "Change Request"
	}
	return "<unknown>"
}

func (code ArtifactAssessmentInformationType) Definition() string {
	switch code {
	case ArtifactAssessmentInformationTypeComment:
		return "A comment on the artifact"
	case ArtifactAssessmentInformationTypeClassifier:
		return "A classifier of the artifact"
	case ArtifactAssessmentInformationTypeRating:
		return "A rating of the artifact"
	case ArtifactAssessmentInformationTypeContainer:
		return "A container for multiple components"
	case ArtifactAssessmentInformationTypeResponse:
		return "A response to a comment"
	case ArtifactAssessmentInformationTypeChangeRequest:
		return "A change request for the artifact"
	}
	return "<unknown>"
}
