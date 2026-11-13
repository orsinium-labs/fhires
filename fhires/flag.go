package fhires

import "encoding/xml"

// Prospective warnings of potential issues when providing care to the patient.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Flag
type Flag struct {
	DomainResource
	// Business identifiers assigned to this flag by the performer or other systems which remain constant as the resource is updated and propagates from server to server.
	//
	// Comment: This is a business identifier, not a resource identifier (see [discussion](resource.html#identifiers)).  It is best practice for the identifier to only appear on a single resource instance, however business practices may occasionally dictate that multiple resource instances with the same identifier can exist - possibly even with different resource types.  For example, multiple Patient and a Person resource instance might share the same social insurance number.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Supports basic workflow.
	//
	// Comment: This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid.
	Status FlagStatus `bson:"status" json:"status" xml:"status"`
	// Allows a flag to be divided into different categories like clinical, administrative etc.
	//
	// Intended to be used as a means of filtering which flags are displayed to particular user or in a given context.
	//
	// Comment: The value set will often need to be adjusted based on local business rules and usage context.
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
	// The coded value or textual component of the flag to display to the user.
	//
	// Comment: If non-coded, use CodeableConcept.text.  This element should always be included in the narrative.
	Code CodeableConcept `bson:"code" json:"code" xml:"code"`
	// The patient, related person, location, group, organization, or practitioner etc.
	//
	// this is about record this flag is associated with.
	Subject Reference `bson:"subject" json:"subject" xml:"subject"`
	// The period of time from the activation of the flag to inactivation of the flag.
	//
	// If the flag is active, the end of the period should be unspecified.
	Period *Period `bson:"period,omitempty" json:"period,omitempty" xml:"period,omitempty"`
	// This alert is only relevant during the encounter.
	//
	// Comment: If both Flag.encounter and Flag.period are valued, then Flag.period.start shall not be before Encounter.period.start and Flag.period.end shall not be after Encounter.period.end.
	Encounter *Reference `bson:"encounter,omitempty" json:"encounter,omitempty" xml:"encounter,omitempty"`
	// The person, organization or device that created the flag.
	Author *Reference `bson:"author,omitempty" json:"author,omitempty" xml:"author,omitempty"`
}

func (r Flag) ResourceType() string {
	return "Flag"
}

// MarshalJSON implements [json.Marshaler]
func (r Flag) MarshalJSON() ([]byte, error) {
	type Wrapper Flag
	return marshalJSON(r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Marshaler].
func (r Flag) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Wrapper Flag
	return marshalXML(e, start, r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Unmarshaler].
func (r *Flag) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper Flag
	w := Wrapper(*r)
	err := unmarshalXML(d, start, r.ResourceType(), &w)
	*r = Flag(w)
	return err
}
