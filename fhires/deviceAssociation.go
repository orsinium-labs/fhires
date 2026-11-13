package fhires

import "encoding/xml"

// A record of association of a device.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/DeviceAssociation
type DeviceAssociation struct {
	DomainResource
	// Instance identifier.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Reference to the devices associated with the patient or group.
	Device Reference `bson:"device" json:"device" xml:"device"`
	// Describes the relationship between the device and subject.
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
	// Indicates the state of the Device association.
	Status CodeableConcept `bson:"status" json:"status" xml:"status"`
	// The reasons given for the current association status.
	StatusReason []CodeableConcept `bson:"statusReason,omitempty" json:"statusReason,omitempty" xml:"statusReason,omitempty"`
	// The individual, group of individuals or device that the device is on or associated with.
	Subject *Reference `bson:"subject,omitempty" json:"subject,omitempty" xml:"subject,omitempty"`
	// Current anatomical location of the device in/on subject.
	BodyStructure *Reference `bson:"bodyStructure,omitempty" json:"bodyStructure,omitempty" xml:"bodyStructure,omitempty"`
	// Begin and end dates and times for the device association.
	Period *Period `bson:"period,omitempty" json:"period,omitempty" xml:"period,omitempty"`
	// The details about the device when it is in use to describe its operation.
	Operation []DeviceAssociationOperation `bson:"operation,omitempty" json:"operation,omitempty" xml:"operation,omitempty"`
}

type DeviceAssociationOperation struct {
	// Unique id for the element within a resource (for internal references).
	//
	// This may be any string value that does not contain spaces.
	ID *string `bson:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element.
	//
	// To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	//
	// Comment: There can be no stigma associated with the use of extensions by any application, project, or standard - regardless of the institution or jurisdiction that uses or defines the extensions.  The use of extensions is what allows the FHIR specification to retain a core level of simplicity for everyone.
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty" xml:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants.
	//
	// Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	//
	// Comment: There can be no stigma associated with the use of extensions by any application, project, or standard - regardless of the institution or jurisdiction that uses or defines the extensions.  The use of extensions is what allows the FHIR specification to retain a core level of simplicity for everyone.
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty" xml:"modifierExtension,omitempty"`
	// Device operational condition corresponding to the association.
	Status CodeableConcept `bson:"status" json:"status" xml:"status"`
	// The individual performing the action enabled by the device.
	Operator []Reference `bson:"operator,omitempty" json:"operator,omitempty" xml:"operator,omitempty"`
	// Begin and end dates and times for the device's operation.
	Period *Period `bson:"period,omitempty" json:"period,omitempty" xml:"period,omitempty"`
}

func (r DeviceAssociation) ResourceType() string {
	return "DeviceAssociation"
}

// MarshalJSON implements [json.Marshaler]
func (r DeviceAssociation) MarshalJSON() ([]byte, error) {
	type Wrapper DeviceAssociation
	return marshalJSON(r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Marshaler].
func (r DeviceAssociation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Wrapper DeviceAssociation
	return marshalXML(e, start, r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Unmarshaler].
func (r *DeviceAssociation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper DeviceAssociation
	w := Wrapper(*r)
	err := unmarshalXML(d, start, r.ResourceType(), &w)
	*r = DeviceAssociation(w)
	return err
}
