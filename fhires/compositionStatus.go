package fhires

// The workflow/clinical status of the composition.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/composition-status
type CompositionStatus string

const (
	// The existence of the composition is registered, but there is nothing yet available.
	CompositionStatusRegistered CompositionStatus = "registered"
	// This is a partial (e.g. initial, interim or preliminary) composition: data in the composition may be incomplete or unverified.
	CompositionStatusPartial CompositionStatus = "partial"
	// Verified early results are available, but not all results are final.
	CompositionStatusPreliminary CompositionStatus = "preliminary"
	// This version of the composition is complete and verified by an appropriate person and no further work is planned. Any subsequent updates would be on a new version of the composition.
	CompositionStatusFinal CompositionStatus = "final"
	// The composition content or the referenced resources have been modified (edited or added to) subsequent to being released as "final" and the composition is complete and verified by an authorized person.
	CompositionStatusAmended CompositionStatus = "amended"
	// Subsequent to being final, the composition content has been modified to correct an error in the composition or referenced results.
	CompositionStatusCorrected CompositionStatus = "corrected"
	// Subsequent to being final, the composition content has been modified by adding new content. The existing content is unchanged.
	CompositionStatusAppended CompositionStatus = "appended"
	// The composition is unavailable because the measurement was not started or not completed (also sometimes called "aborted").
	CompositionStatusCancelled CompositionStatus = "cancelled"
	// The composition or document was originally created/issued in error, and this is an amendment that marks that the entire series should not be considered as valid.
	CompositionStatusEnteredInError CompositionStatus = "entered-in-error"
	// This composition has been withdrawn or superseded and should no longer be used.
	CompositionStatusDeprecated CompositionStatus = "deprecated"
	// The authoring/source system does not know which of the status values currently applies for this observation. Note: This concept is not to be used for "other" - one of the listed statuses is presumed to apply, but the authoring/source system does not know which.
	CompositionStatusUnknown CompositionStatus = "unknown"
)

func (code CompositionStatus) String() string {
	return code.Code()
}

func (code CompositionStatus) Code() string {
	return string(code)
}

func (code CompositionStatus) Display() string {
	switch code {
	case CompositionStatusRegistered:
		return "Registered"
	case CompositionStatusPartial:
		return "Partial"
	case CompositionStatusPreliminary:
		return "Preliminary"
	case CompositionStatusFinal:
		return "Final"
	case CompositionStatusAmended:
		return "Amended"
	case CompositionStatusCorrected:
		return "Corrected"
	case CompositionStatusAppended:
		return "Appended"
	case CompositionStatusCancelled:
		return "Cancelled"
	case CompositionStatusEnteredInError:
		return "Entered in Error"
	case CompositionStatusDeprecated:
		return "Deprecated"
	case CompositionStatusUnknown:
		return "Unknown"
	}
	return "<unknown>"
}

func (code CompositionStatus) Definition() string {
	switch code {
	case CompositionStatusRegistered:
		return "The existence of the composition is registered, but there is nothing yet available."
	case CompositionStatusPartial:
		return "This is a partial (e.g. initial, interim or preliminary) composition: data in the composition may be incomplete or unverified."
	case CompositionStatusPreliminary:
		return "Verified early results are available, but not all results are final."
	case CompositionStatusFinal:
		return "This version of the composition is complete and verified by an appropriate person and no further work is planned. Any subsequent updates would be on a new version of the composition."
	case CompositionStatusAmended:
		return "The composition content or the referenced resources have been modified (edited or added to) subsequent to being released as \"final\" and the composition is complete and verified by an authorized person."
	case CompositionStatusCorrected:
		return "Subsequent to being final, the composition content has been modified to correct an error in the composition or referenced results."
	case CompositionStatusAppended:
		return "Subsequent to being final, the composition content has been modified by adding new content. The existing content is unchanged."
	case CompositionStatusCancelled:
		return "The composition is unavailable because the measurement was not started or not completed (also sometimes called \"aborted\")."
	case CompositionStatusEnteredInError:
		return "The composition or document was originally created/issued in error, and this is an amendment that marks that the entire series should not be considered as valid."
	case CompositionStatusDeprecated:
		return "This composition has been withdrawn or superseded and should no longer be used."
	case CompositionStatusUnknown:
		return "The authoring/source system does not know which of the status values currently applies for this observation. Note: This concept is not to be used for \"other\" - one of the listed statuses is presumed to apply, but the authoring/source system does not know which."
	}
	return "<unknown>"
}
