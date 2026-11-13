package fhires

import "encoding/xml"

// A specific set of Roles/Locations/specialties/services that a practitioner may perform, or has performed at an organization during a period of time.
//
// Purpose: Need to track services that a healthcare provider is able to provide at an organization's location, and the services that they can perform there.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/PractitionerRole
type PractitionerRole struct {
	DomainResource
	// Business Identifiers that are specific to a role/location.
	//
	// Comment: A specific identifier value (e.g. Minnesota Registration ID) may appear on multiple PractitionerRole instances which could be for different periods, or different Roles (or other reasons).  A specific identifier may be included multiple times in the same PractitionerRole instance with different identifier validity periods.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	//  Whether this practitioner role record is in active use.
	//
	// Some systems may use this property to mark non-active practitioners, such as those that are not currently employed.
	//
	// Comment: If this value is false, you may refer to the period to see when the role was in active use. If there is no period specified, no inference can be made about when it was active.
	Active *bool `bson:"active,omitempty" json:"active,omitempty" xml:"active,omitempty"`
	// The period during which the person is authorized to act as a practitioner in these role(s) for the organization.
	//
	// Comment: If a practitioner is performing a role within an organization over multiple, non-adjacent periods, there should be a distinct PractitionerRole instance for each period.  For example, if a nurse is employed at a hospital, leaves the organization for a period of time due to pandemic related stress, but returns post-pandemic to the same job role, there would be two PractitionerRole instances for the different periods of employment.
	Period *Period `bson:"period,omitempty" json:"period,omitempty" xml:"period,omitempty"`
	// Practitioner that is able to provide the defined services for the organization.
	Practitioner *Reference `bson:"practitioner,omitempty" json:"practitioner,omitempty" xml:"practitioner,omitempty"`
	// The organization where the Practitioner performs the roles associated.
	Organization *Reference `bson:"organization,omitempty" json:"organization,omitempty" xml:"organization,omitempty"`
	// Roles which this practitioner is authorized to perform for the organization.
	//
	// Comment: A person may have more than one role.
	Code []CodeableConcept `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// The specialty of a practitioner that describes the functional role they are practicing at a given organization or location.
	//
	// Comment: The specialty represents the functional role a practitioner is playing within an organization/location. This role may require the individual have certain qualifications, which would be represented in the Practitioner.qualifications property.  Note that qualifications may or might not relate to or be required by the practicing specialty.
	Specialty []CodeableConcept `bson:"specialty,omitempty" json:"specialty,omitempty" xml:"specialty,omitempty"`
	// The location(s) at which this practitioner provides care.
	Location []Reference `bson:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// The list of healthcare services that this worker provides for this role's Organization/Location(s).
	HealthcareService []Reference `bson:"healthcareService,omitempty" json:"healthcareService,omitempty" xml:"healthcareService,omitempty"`
	// The contact details of communication devices available relevant to the specific PractitionerRole.
	//
	// This can include addresses, phone numbers, fax numbers, mobile numbers, email addresses and web sites.
	Contact []ExtendedContactDetail `bson:"contact,omitempty" json:"contact,omitempty" xml:"contact,omitempty"`
	// Collection of characteristics (attributes).
	//
	// Comment: These could be such things as is the service mode used by this role.
	Characteristic []CodeableConcept `bson:"characteristic,omitempty" json:"characteristic,omitempty" xml:"characteristic,omitempty"`
	// A language the practitioner can use in patient communication.
	//
	// The practitioner may know several languages (listed in practitioner.communication), however these are the languages that could be advertised in a directory for a patient to search.
	//
	// Comment: The structure aa-BB with this exact casing is one the most widely used notations for locale. However not all systems code this but instead have it as free text. Hence CodeableConcept instead of code as the data type.Note that for non-patient oriented communication, see Practitioner.communication.  Note that all 'person' type resources (Person, RelatedPerson, Patient, Practitioner) have a communication structure that includes preferences.  Role or service oriented resources such as HealthcareService and PractitionerRole only include languages that are available for interacting with patients.
	Communication []CodeableConcept `bson:"communication,omitempty" json:"communication,omitempty" xml:"communication,omitempty"`
	// A collection of times the practitioner is available or performing this role at the location and/or healthcareservice.
	//
	// Comment: More detailed availability information may be provided in associated Schedule/Slot resources.
	//
	// Systems may choose to render availability differently than it is exchanged on the interface. For example, rather than "Mon, Tue, Wed, Thur, Fri from 9am-12am; Mon, Tue, Wed, Thur, Fri from 1pm-5pm" as would be implied by two availableTime repetitions, an application could render this information as "Mon-Fri 9-12am and 1-5pm".The NotAvailableTime(s) included indicate the general days/periods where the practitioner is not available (for things such as vacation time, or public holidays).
	Availability []Availability `bson:"availability,omitempty" json:"availability,omitempty" xml:"availability,omitempty"`
	//  Technical endpoints providing access to services operated for the practitioner with this role.
	//
	// Commonly used for locating scheduling services, or identifying where to send referrals electronically.
	Endpoint []Reference `bson:"endpoint,omitempty" json:"endpoint,omitempty" xml:"endpoint,omitempty"`
}

func (r PractitionerRole) ResourceType() string {
	return "PractitionerRole"
}

// MarshalJSON implements [json.Marshaler]
func (r PractitionerRole) MarshalJSON() ([]byte, error) {
	type Wrapper PractitionerRole
	return marshalJSON(r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Marshaler].
func (r PractitionerRole) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Wrapper PractitionerRole
	return marshalXML(e, start, r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Unmarshaler].
func (r *PractitionerRole) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper PractitionerRole
	w := Wrapper(*r)
	err := unmarshalXML(d, start, r.ResourceType(), &w)
	*r = PractitionerRole(w)
	return err
}
