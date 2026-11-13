package fhires

import "encoding/xml"

// A slot of time on a schedule that may be available for booking appointments.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Slot
type Slot struct {
	DomainResource
	// External Ids for this item.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// A broad categorization of the service that is to be performed during this appointment.
	ServiceCategory []CodeableConcept `bson:"serviceCategory,omitempty" json:"serviceCategory,omitempty" xml:"serviceCategory,omitempty"`
	// The type of appointments that can be booked into this slot (ideally this would be an identifiable service - which is at a location, rather than the location itself).
	//
	// If provided then this overrides the value provided on the Schedule resource.
	ServiceType []CodeableReference `bson:"serviceType,omitempty" json:"serviceType,omitempty" xml:"serviceType,omitempty"`
	// The specialty of a practitioner that would be required to perform the service requested in this appointment.
	Specialty []CodeableConcept `bson:"specialty,omitempty" json:"specialty,omitempty" xml:"specialty,omitempty"`
	// The style of appointment or patient that may be booked in the slot (not service type).
	//
	// Comment: A slot may be allow multiple appointment types, but when booked, would only be used for one of the given appointment types.
	AppointmentType []CodeableConcept `bson:"appointmentType,omitempty" json:"appointmentType,omitempty" xml:"appointmentType,omitempty"`
	// The schedule resource that this slot defines an interval of status information.
	Schedule Reference `bson:"schedule" json:"schedule" xml:"schedule"`
	// busy | free | busy-unavailable | busy-tentative | entered-in-error.
	Status SlotStatus `bson:"status" json:"status" xml:"status"`
	// Date/Time that the slot is to begin.
	Start string `bson:"start" json:"start" xml:"start"`
	// Date/Time that the slot is to conclude.
	End string `bson:"end" json:"end" xml:"end"`
	// This slot has already been overbooked, appointments are unlikely to be accepted for this time.
	Overbooked *bool `bson:"overbooked,omitempty" json:"overbooked,omitempty" xml:"overbooked,omitempty"`
	// Comments on the slot to describe any extended information.
	//
	// Such as custom constraints on the slot.
	Comment *string `bson:"comment,omitempty" json:"comment,omitempty" xml:"comment,omitempty"`
}

func (r Slot) ResourceType() string {
	return "Slot"
}

// MarshalJSON implements [json.Marshaler]
func (r Slot) MarshalJSON() ([]byte, error) {
	type Wrapper Slot
	return marshalJSON(r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Marshaler].
func (r Slot) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Wrapper Slot
	return marshalXML(e, start, r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Unmarshaler].
func (r *Slot) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper Slot
	w := Wrapper(*r)
	err := unmarshalXML(d, start, r.ResourceType(), &w)
	*r = Slot(w)
	return err
}
