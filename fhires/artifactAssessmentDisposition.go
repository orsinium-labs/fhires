package fhires

// Possible values for the disposition of a comment or change request, typically used for comments and change requests, to indicate the disposition of the responsible party towards the changes suggested by the comment or change request.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/artifactassessment-disposition
type ArtifactAssessmentDisposition string

const (
	// The comment is unresolved
	ArtifactAssessmentDispositionUnresolved ArtifactAssessmentDisposition = "unresolved"
	// The comment is not persuasive (rejected in full)
	ArtifactAssessmentDispositionNotPersuasive ArtifactAssessmentDisposition = "not-persuasive"
	// The comment is persuasive (accepted in full)
	ArtifactAssessmentDispositionPersuasive ArtifactAssessmentDisposition = "persuasive"
	// The comment is persuasive with modification (partially accepted)
	ArtifactAssessmentDispositionPersuasiveWithModification ArtifactAssessmentDisposition = "persuasive-with-modification"
	// The comment is not persuasive with modification (partially rejected)
	ArtifactAssessmentDispositionNotPersuasiveWithModification ArtifactAssessmentDisposition = "not-persuasive-with-modification"
)

func (code ArtifactAssessmentDisposition) String() string {
	return code.Code()
}

func (code ArtifactAssessmentDisposition) Code() string {
	return string(code)
}

func (code ArtifactAssessmentDisposition) Display() string {
	switch code {
	case ArtifactAssessmentDispositionUnresolved:
		return "Unresolved"
	case ArtifactAssessmentDispositionNotPersuasive:
		return "Not Persuasive"
	case ArtifactAssessmentDispositionPersuasive:
		return "Persuasive"
	case ArtifactAssessmentDispositionPersuasiveWithModification:
		return "Persuasive with Modification"
	case ArtifactAssessmentDispositionNotPersuasiveWithModification:
		return "Not Persuasive with Modification"
	}
	return "<unknown>"
}

func (code ArtifactAssessmentDisposition) Definition() string {
	switch code {
	case ArtifactAssessmentDispositionUnresolved:
		return "The comment is unresolved"
	case ArtifactAssessmentDispositionNotPersuasive:
		return "The comment is not persuasive (rejected in full)"
	case ArtifactAssessmentDispositionPersuasive:
		return "The comment is persuasive (accepted in full)"
	case ArtifactAssessmentDispositionPersuasiveWithModification:
		return "The comment is persuasive with modification (partially accepted)"
	case ArtifactAssessmentDispositionNotPersuasiveWithModification:
		return "The comment is not persuasive with modification (partially rejected)"
	}
	return "<unknown>"
}
