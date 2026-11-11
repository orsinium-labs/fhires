package fhires

// Possible values for the workflow status of the comment or assessment, typically used to coordinate workflow around the process of accepting and rejecting changes and comments on the artifact.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/artifactassessment-workflow-status
type ArtifactAssessmentWorkflowStatus string

const (
	// The comment has been submitted, but the responsible party has not yet been determined, or the responsible party has not yet determined the next steps to be taken.
	ArtifactAssessmentWorkflowStatusSubmitted ArtifactAssessmentWorkflowStatus = "submitted"
	// The comment has been triaged, meaning the responsible party has been determined and next steps have been identified to address the comment.
	ArtifactAssessmentWorkflowStatusTriaged ArtifactAssessmentWorkflowStatus = "triaged"
	// The comment is waiting for input from a specific party before next steps can be taken.
	ArtifactAssessmentWorkflowStatusWaitingForInput ArtifactAssessmentWorkflowStatus = "waiting-for-input"
	// The comment has been resolved and no changes resulted from the resolution
	ArtifactAssessmentWorkflowStatusResolvedNoChange ArtifactAssessmentWorkflowStatus = "resolved-no-change"
	// The comment has been resolved and changes are required to address the comment
	ArtifactAssessmentWorkflowStatusResolvedChangeRequired ArtifactAssessmentWorkflowStatus = "resolved-change-required"
	// The comment is acceptable, but resolution of the comment and application of any associated changes have been deferred
	ArtifactAssessmentWorkflowStatusDeferred ArtifactAssessmentWorkflowStatus = "deferred"
	// The comment is a duplicate of another comment already received
	ArtifactAssessmentWorkflowStatusDuplicate ArtifactAssessmentWorkflowStatus = "duplicate"
	// The comment is resolved and any necessary changes have been applied
	ArtifactAssessmentWorkflowStatusApplied ArtifactAssessmentWorkflowStatus = "applied"
	// The necessary changes to the artifact have been published in a new version of the artifact
	ArtifactAssessmentWorkflowStatusPublished ArtifactAssessmentWorkflowStatus = "published"
	// The assessment was entered in error
	ArtifactAssessmentWorkflowStatusEnteredInError ArtifactAssessmentWorkflowStatus = "entered-in-error"
)

func (code ArtifactAssessmentWorkflowStatus) String() string {
	return code.Code()
}

func (code ArtifactAssessmentWorkflowStatus) Code() string {
	return string(code)
}

func (code ArtifactAssessmentWorkflowStatus) Display() string {
	switch code {
	case ArtifactAssessmentWorkflowStatusSubmitted:
		return "Submitted"
	case ArtifactAssessmentWorkflowStatusTriaged:
		return "Triaged"
	case ArtifactAssessmentWorkflowStatusWaitingForInput:
		return "Waiting for Input"
	case ArtifactAssessmentWorkflowStatusResolvedNoChange:
		return "Resolved - No Change"
	case ArtifactAssessmentWorkflowStatusResolvedChangeRequired:
		return "Resolved - Change Required"
	case ArtifactAssessmentWorkflowStatusDeferred:
		return "Deferred"
	case ArtifactAssessmentWorkflowStatusDuplicate:
		return "Duplicate"
	case ArtifactAssessmentWorkflowStatusApplied:
		return "Applied"
	case ArtifactAssessmentWorkflowStatusPublished:
		return "Published"
	case ArtifactAssessmentWorkflowStatusEnteredInError:
		return "Entered in Error"
	}
	return "<unknown>"
}

func (code ArtifactAssessmentWorkflowStatus) Definition() string {
	switch code {
	case ArtifactAssessmentWorkflowStatusSubmitted:
		return "The comment has been submitted, but the responsible party has not yet been determined, or the responsible party has not yet determined the next steps to be taken."
	case ArtifactAssessmentWorkflowStatusTriaged:
		return "The comment has been triaged, meaning the responsible party has been determined and next steps have been identified to address the comment."
	case ArtifactAssessmentWorkflowStatusWaitingForInput:
		return "The comment is waiting for input from a specific party before next steps can be taken."
	case ArtifactAssessmentWorkflowStatusResolvedNoChange:
		return "The comment has been resolved and no changes resulted from the resolution"
	case ArtifactAssessmentWorkflowStatusResolvedChangeRequired:
		return "The comment has been resolved and changes are required to address the comment"
	case ArtifactAssessmentWorkflowStatusDeferred:
		return "The comment is acceptable, but resolution of the comment and application of any associated changes have been deferred"
	case ArtifactAssessmentWorkflowStatusDuplicate:
		return "The comment is a duplicate of another comment already received"
	case ArtifactAssessmentWorkflowStatusApplied:
		return "The comment is resolved and any necessary changes have been applied"
	case ArtifactAssessmentWorkflowStatusPublished:
		return "The necessary changes to the artifact have been published in a new version of the artifact"
	case ArtifactAssessmentWorkflowStatusEnteredInError:
		return "The assessment was entered in error"
	}
	return "<unknown>"
}
