package fhires

// The status of the episode of care.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/episode-of-care-status
type EpisodeOfCareStatus string

const (
	// This episode of care is planned to start at the date specified in the period.start. During this status, an organization may perform assessments to determine if the patient is eligible to receive services, or be organizing to make resources available to provide care services.
	EpisodeOfCareStatusPlanned EpisodeOfCareStatus = "planned"
	// This episode has been placed on a waitlist, pending the episode being made active (or cancelled).
	EpisodeOfCareStatusWaitlist EpisodeOfCareStatus = "waitlist"
	// This episode of care is current.
	EpisodeOfCareStatusActive EpisodeOfCareStatus = "active"
	// This episode of care is on hold; the organization has limited responsibility for the patient (such as while on respite).
	EpisodeOfCareStatusOnhold EpisodeOfCareStatus = "onhold"
	// This episode of care is finished and the organization is not expecting to be providing further care to the patient. Can also be known as "closed", "completed" or other similar terms.
	EpisodeOfCareStatusFinished EpisodeOfCareStatus = "finished"
	// The episode of care was cancelled, or withdrawn from service, often selected during the planned stage as the patient may have gone elsewhere, or the circumstances have changed and the organization is unable to provide the care. It indicates that services terminated outside the planned/expected workflow.
	EpisodeOfCareStatusCancelled EpisodeOfCareStatus = "cancelled"
	// This instance should not have been part of this patient's medical record.
	EpisodeOfCareStatusEnteredInError EpisodeOfCareStatus = "entered-in-error"
)

func (code EpisodeOfCareStatus) String() string {
	return code.Code()
}

func (code EpisodeOfCareStatus) Code() string {
	return string(code)
}

func (code EpisodeOfCareStatus) Display() string {
	switch code {
	case EpisodeOfCareStatusPlanned:
		return "Planned"
	case EpisodeOfCareStatusWaitlist:
		return "Waitlist"
	case EpisodeOfCareStatusActive:
		return "Active"
	case EpisodeOfCareStatusOnhold:
		return "On Hold"
	case EpisodeOfCareStatusFinished:
		return "Finished"
	case EpisodeOfCareStatusCancelled:
		return "Cancelled"
	case EpisodeOfCareStatusEnteredInError:
		return "Entered in Error"
	}
	return "<unknown>"
}

func (code EpisodeOfCareStatus) Definition() string {
	switch code {
	case EpisodeOfCareStatusPlanned:
		return "This episode of care is planned to start at the date specified in the period.start. During this status, an organization may perform assessments to determine if the patient is eligible to receive services, or be organizing to make resources available to provide care services."
	case EpisodeOfCareStatusWaitlist:
		return "This episode has been placed on a waitlist, pending the episode being made active (or cancelled)."
	case EpisodeOfCareStatusActive:
		return "This episode of care is current."
	case EpisodeOfCareStatusOnhold:
		return "This episode of care is on hold; the organization has limited responsibility for the patient (such as while on respite)."
	case EpisodeOfCareStatusFinished:
		return "This episode of care is finished and the organization is not expecting to be providing further care to the patient. Can also be known as \"closed\", \"completed\" or other similar terms."
	case EpisodeOfCareStatusCancelled:
		return "The episode of care was cancelled, or withdrawn from service, often selected during the planned stage as the patient may have gone elsewhere, or the circumstances have changed and the organization is unable to provide the care. It indicates that services terminated outside the planned/expected workflow."
	case EpisodeOfCareStatusEnteredInError:
		return "This instance should not have been part of this patient's medical record."
	}
	return "<unknown>"
}
