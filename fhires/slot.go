package fhires

import "encoding/xml"

// A slot of time on a schedule that may be available for booking appointments.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Slot
type Slot struct {
	BaseResource
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human.
	//
	// The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	//
	// Comment: Contained resources do not have a narrative. Resources that are not contained SHOULD have a narrative. In some cases, a resource may only have text with little or no additional discrete data (as long as all minOccurs=1 elements are satisfied).  This may be necessary for data from legacy systems where information is captured as a "text blob" or where text is additionally entered raw or narrated and encoded information is added later.
	Text *Narrative `bson:"text,omitempty" json:"text,omitempty" xml:"text,omitempty"`
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, nor can they have their own independent transaction scope.
	//
	// This is allowed to be a Parameters resource if and only if it is referenced by a resource that provides context/meaning.
	//
	// Comment: This should never be done when the content can be identified properly, as once identification is lost, it is extremely difficult (and context dependent) to restore it again. Contained resources may have profiles and tags in their meta elements, but SHALL NOT have security labels.
	// May be used to represent additional information that is not part of the basic definition of the resource.
	//
	// To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	//
	// Comment: There can be no stigma associated with the use of extensions by any application, project, or standard - regardless of the institution or jurisdiction that uses or defines the extensions.  The use of extensions is what allows the FHIR specification to retain a core level of simplicity for everyone.
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty" xml:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants.
	//
	// Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	//
	// Comment: There can be no stigma associated with the use of extensions by any application, project, or standard - regardless of the institution or jurisdiction that uses or defines the extensions.  The use of extensions is what allows the FHIR specification to retain a core level of simplicity for everyone.
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty" xml:"modifierExtension,omitempty"`
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

// UnmarshalXML implements [xml.Unmarshaler].
func (r *Slot) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper Slot
	w := Wrapper(*r)
	err := unmarshalXML(d, start, r.ResourceType(), &w)
	*r = Slot(w)
	return err
}
