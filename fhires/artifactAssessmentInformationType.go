package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// ArtifactAssessmentInformationType is documented here http://hl7.org/fhir/ValueSet/artifactassessment-information-type
type ArtifactAssessmentInformationType int

const (
	ArtifactAssessmentInformationTypeComment ArtifactAssessmentInformationType = iota
	ArtifactAssessmentInformationTypeClassifier
	ArtifactAssessmentInformationTypeRating
	ArtifactAssessmentInformationTypeContainer
	ArtifactAssessmentInformationTypeResponse
	ArtifactAssessmentInformationTypeChangeRequest
)

func (code ArtifactAssessmentInformationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *ArtifactAssessmentInformationType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "comment":
		*code = ArtifactAssessmentInformationTypeComment
	case "classifier":
		*code = ArtifactAssessmentInformationTypeClassifier
	case "rating":
		*code = ArtifactAssessmentInformationTypeRating
	case "container":
		*code = ArtifactAssessmentInformationTypeContainer
	case "response":
		*code = ArtifactAssessmentInformationTypeResponse
	case "change-request":
		*code = ArtifactAssessmentInformationTypeChangeRequest
	default:
		return fmt.Errorf("unknown ArtifactAssessmentInformationType code `%s`", s)
	}
	return nil
}

func (code ArtifactAssessmentInformationType) String() string {
	return code.Code()
}

func (code ArtifactAssessmentInformationType) Code() string {
	switch code {
	case ArtifactAssessmentInformationTypeComment:
		return "comment"
	case ArtifactAssessmentInformationTypeClassifier:
		return "classifier"
	case ArtifactAssessmentInformationTypeRating:
		return "rating"
	case ArtifactAssessmentInformationTypeContainer:
		return "container"
	case ArtifactAssessmentInformationTypeResponse:
		return "response"
	case ArtifactAssessmentInformationTypeChangeRequest:
		return "change-request"
	}
	return "<unknown>"
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
