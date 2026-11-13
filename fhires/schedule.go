package fhires

import "encoding/xml"

// A container for slots of time that may be available for booking appointments.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Schedule
type Schedule struct {
	DomainResource
	// External Ids for this item.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Whether this schedule record is in active use or should not be used (such as was entered in error).
	//
	// Comment: This element is labeled as a modifier because it may be used to mark that the resource was created in error.
	Active *bool `bson:"active,omitempty" json:"active,omitempty" xml:"active,omitempty"`
	// A broad categorization of the service that is to be performed during this appointment.
	ServiceCategory []CodeableConcept `bson:"serviceCategory,omitempty" json:"serviceCategory,omitempty" xml:"serviceCategory,omitempty"`
	// The specific service that is to be performed during this appointment.
	ServiceType []CodeableReference `bson:"serviceType,omitempty" json:"serviceType,omitempty" xml:"serviceType,omitempty"`
	// The specialty of a practitioner that would be required to perform the service requested in this appointment.
	Specialty []CodeableConcept `bson:"specialty,omitempty" json:"specialty,omitempty" xml:"specialty,omitempty"`
	// Further description of the schedule as it would be presented to a consumer while searching.
	//
	// Comment: This MAY be used to describe what the schedule is for where it is clearer than just the name of the single actor.
	Name *string `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Slots that reference this schedule resource provide the availability details to these referenced resource(s).
	//
	// Comment: The capacity to support multiple referenced resource types should be used in cases where the specific resources themselves cannot be scheduled without the other, and thus only make sense to the system exposing them as a group. Common examples of this are where the combination of a practitioner and a room (Location) are always required by a system.
	Actor []Reference `bson:"actor" json:"actor" xml:"actor"`
	// The period of time that the slots that reference this Schedule resource cover (even if none exist).
	//
	// These  cover the amount of time that an organization's planning horizon; the interval for which they are currently accepting appointments. This does not define a "template" for planning outside these dates.
	PlanningHorizon *Period `bson:"planningHorizon,omitempty" json:"planningHorizon,omitempty" xml:"planningHorizon,omitempty"`
	// Comments on the availability to describe any extended information.
	//
	// Such as custom constraints on the slots that may be associated.
	Comment *string `bson:"comment,omitempty" json:"comment,omitempty" xml:"comment,omitempty"`
}

func (r Schedule) ResourceType() string {
	return "Schedule"
}

// MarshalJSON implements [json.Marshaler]
func (r Schedule) MarshalJSON() ([]byte, error) {
	type Wrapper Schedule
	return marshalJSON(r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Marshaler].
func (r Schedule) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Wrapper Schedule
	return marshalXML(e, start, r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Unmarshaler].
func (r *Schedule) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper Schedule
	w := Wrapper(*r)
	err := unmarshalXML(d, start, r.ResourceType(), &w)
	*r = Schedule(w)
	return err
}
