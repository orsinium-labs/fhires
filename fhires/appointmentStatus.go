package fhires

// The free/busy status of an appointment.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/appointmentstatus
type AppointmentStatus string

const (
	// None of the participant(s) have finalized their acceptance of the appointment request, and the start/end time might not be set yet.
	AppointmentStatusProposed AppointmentStatus = "proposed"
	// Some or all of the participant(s) have not finalized their acceptance of the appointment request.
	AppointmentStatusPending AppointmentStatus = "pending"
	// All participant(s) have been considered and the appointment is confirmed to go ahead at the date/times specified.
	AppointmentStatusBooked AppointmentStatus = "booked"
	// The patient/patients has/have arrived and is/are waiting to be seen.
	AppointmentStatusArrived AppointmentStatus = "arrived"
	// The planning stages of the appointment are now complete, the encounter resource will exist and will track further status changes. Note that an encounter may exist before the appointment status is fulfilled for many reasons.
	AppointmentStatusFulfilled AppointmentStatus = "fulfilled"
	// The appointment has been cancelled.
	AppointmentStatusCancelled AppointmentStatus = "cancelled"
	// Some or all of the participant(s) have not/did not appear for the appointment (usually the patient).
	AppointmentStatusNoshow AppointmentStatus = "noshow"
	// This instance should not have been part of this patient's medical record.
	AppointmentStatusEnteredInError AppointmentStatus = "entered-in-error"
	// When checked in, all pre-encounter administrative work is complete, and the encounter may begin. (where multiple patients are involved, they are all present).
	AppointmentStatusCheckedIn AppointmentStatus = "checked-in"
	/*
	   The appointment has been placed on a waitlist, to be scheduled/confirmed in the future when a slot/service is available.
	   A specific time might or might not be pre-allocated.
	*/
	AppointmentStatusWaitlist AppointmentStatus = "waitlist"
)

func (code AppointmentStatus) String() string {
	return code.Code()
}

func (code AppointmentStatus) Code() string {
	return string(code)
}

func (code AppointmentStatus) Display() string {
	switch code {
	case AppointmentStatusProposed:
		return "Proposed"
	case AppointmentStatusPending:
		return "Pending"
	case AppointmentStatusBooked:
		return "Booked"
	case AppointmentStatusArrived:
		return "Arrived"
	case AppointmentStatusFulfilled:
		return "Fulfilled"
	case AppointmentStatusCancelled:
		return "Cancelled"
	case AppointmentStatusNoshow:
		return "No Show"
	case AppointmentStatusEnteredInError:
		return "Entered in error"
	case AppointmentStatusCheckedIn:
		return "Checked In"
	case AppointmentStatusWaitlist:
		return "Waitlisted"
	}
	return "<unknown>"
}

func (code AppointmentStatus) Definition() string {
	switch code {
	case AppointmentStatusProposed:
		return "None of the participant(s) have finalized their acceptance of the appointment request, and the start/end time might not be set yet."
	case AppointmentStatusPending:
		return "Some or all of the participant(s) have not finalized their acceptance of the appointment request."
	case AppointmentStatusBooked:
		return "All participant(s) have been considered and the appointment is confirmed to go ahead at the date/times specified."
	case AppointmentStatusArrived:
		return "The patient/patients has/have arrived and is/are waiting to be seen."
	case AppointmentStatusFulfilled:
		return "The planning stages of the appointment are now complete, the encounter resource will exist and will track further status changes. Note that an encounter may exist before the appointment status is fulfilled for many reasons."
	case AppointmentStatusCancelled:
		return "The appointment has been cancelled."
	case AppointmentStatusNoshow:
		return "Some or all of the participant(s) have not/did not appear for the appointment (usually the patient)."
	case AppointmentStatusEnteredInError:
		return "This instance should not have been part of this patient's medical record."
	case AppointmentStatusCheckedIn:
		return "When checked in, all pre-encounter administrative work is complete, and the encounter may begin. (where multiple patients are involved, they are all present)."
	case AppointmentStatusWaitlist:
		return "The appointment has been placed on a waitlist, to be scheduled/confirmed in the future when a slot/service is available.\nA specific time might or might not be pre-allocated."
	}
	return "<unknown>"
}
