package fhires

import "encoding/xml"

// Indicates that a device is to be or has been dispensed for a named person/patient.
//
//	This includes a description of the product (supply) provided and the instructions for using the device.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/DeviceDispense
type DeviceDispense struct {
	DomainResource
	// Business identifier for this dispensation.
	//
	// Comment: This is a business identifier, not a resource identifier.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The order or request that this dispense is fulfilling.
	BasedOn []Reference `bson:"basedOn,omitempty" json:"basedOn,omitempty" xml:"basedOn,omitempty"`
	// The bigger event that this dispense is a part of.
	PartOf []Reference `bson:"partOf,omitempty" json:"partOf,omitempty" xml:"partOf,omitempty"`
	// A code specifying the state of the set of dispense events.
	//
	// Comment: This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid.
	Status DeviceDispenseStatusCodes `bson:"status" json:"status" xml:"status"`
	// Indicates the reason why a dispense was or was not performed.
	StatusReason *CodeableReference `bson:"statusReason,omitempty" json:"statusReason,omitempty" xml:"statusReason,omitempty"`
	// Indicates the type of device dispense.
	//
	// Comment: The category can be used to include where the device is expected to be consumed or other types of dispenses.  Invariants can be used to bind to different value sets when profiling to bind.
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
	// Identifies the device being dispensed.
	//
	// This is either a link to a resource representing the details of the device or a simple attribute carrying a code that identifies the device from a known list of devices.
	Device CodeableReference `bson:"device" json:"device" xml:"device"`
	// A link to a resource representing the person to whom the device is intended.
	//
	// Comment: What about 'this measuring Device is now asigned to Dr X who needs it now'?
	Subject Reference `bson:"subject" json:"subject" xml:"subject"`
	// Identifies the person who picked up the device or the person or location where the device was delivered.
	//
	//  This may be a patient or their caregiver, but some cases exist where it can be a healthcare professional or a location.
	Receiver *Reference `bson:"receiver,omitempty" json:"receiver,omitempty" xml:"receiver,omitempty"`
	// The encounter that establishes the context for this event.
	Encounter *Reference `bson:"encounter,omitempty" json:"encounter,omitempty" xml:"encounter,omitempty"`
	// Additional information that supports the device being dispensed.
	SupportingInformation []Reference `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty" xml:"supportingInformation,omitempty"`
	// Indicates who or what performed the event.
	Performer []DeviceDispensePerformer `bson:"performer,omitempty" json:"performer,omitempty" xml:"performer,omitempty"`
	// The principal physical location where the dispense was performed.
	Location *Reference `bson:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// Indicates the type of dispensing event that is performed.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The number of devices that have been dispensed.
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty" xml:"quantity,omitempty"`
	// The time when the dispensed product was packaged and reviewed.
	PreparedDate *string `bson:"preparedDate,omitempty" json:"preparedDate,omitempty" xml:"preparedDate,omitempty"`
	// The time the dispensed product was made available to the patient or their representative.
	WhenHandedOver *string `bson:"whenHandedOver,omitempty" json:"whenHandedOver,omitempty" xml:"whenHandedOver,omitempty"`
	// Identification of the facility/location where the device was /should be shipped to, as part of the dispense process.
	Destination *Reference `bson:"destination,omitempty" json:"destination,omitempty" xml:"destination,omitempty"`
	// Extra information about the dispense that could not be conveyed in the other attributes.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// The full representation of the instructions.
	//
	// Comment: The content of the renderedDosageInstructions must not be different than the dose represented in the dosageInstruction content.
	UsageInstruction *string `bson:"usageInstruction,omitempty" json:"usageInstruction,omitempty" xml:"usageInstruction,omitempty"`
	// A summary of the events of interest that have occurred, such as when the dispense was verified.
	//
	// Comment: This might not include provenances for all versions of the request – only those deemed “relevant” or important. This SHALL NOT include the Provenance associated with this current version of the resource. (If that provenance is deemed to be a “relevant” change, it will need to be added as part of a later update. Until then, it can be queried directly as the Provenance that points to this version using _revinclude All Provenances should have some historical version of this Request as their subject.).
	EventHistory []Reference `bson:"eventHistory,omitempty" json:"eventHistory,omitempty" xml:"eventHistory,omitempty"`
}

type DeviceDispensePerformer struct {
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
	// Distinguishes the type of performer in the dispense.
	//
	//  For example, date enterer, packager, final checker.
	Function *CodeableConcept `bson:"function,omitempty" json:"function,omitempty" xml:"function,omitempty"`
	// The device, practitioner, etc.
	//
	// who performed the action.  It should be assumed that the actor is the dispenser of the device.
	Actor Reference `bson:"actor" json:"actor" xml:"actor"`
}

func (r DeviceDispense) ResourceType() string {
	return "DeviceDispense"
}

// MarshalJSON implements [json.Marshaler]
func (r DeviceDispense) MarshalJSON() ([]byte, error) {
	type Wrapper DeviceDispense
	return marshalJSON(r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Marshaler].
func (r DeviceDispense) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Wrapper DeviceDispense
	return marshalXML(e, start, r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Unmarshaler].
func (r *DeviceDispense) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper DeviceDispense
	w := Wrapper(*r)
	err := unmarshalXML(d, start, r.ResourceType(), &w)
	*r = DeviceDispense(w)
	return err
}
