package fhires

import "encoding/xml"

// A record of a non-patient specific request for a medication, substance, device, certain types of biologically derived product, and nutrition product used in the healthcare setting.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/SupplyRequest
type SupplyRequest struct {
	DomainResource
	// Business identifiers assigned to this SupplyRequest by the author and/or other systems.
	//
	// These identifiers remain constant as the resource is updated and propagates from server to server.
	//
	// Comment: The identifier.type element is used to distinguish between the identifiers assigned by the requester/placer and the performer/filler.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Status of the supply request.
	Status *SupplyRequestStatus `bson:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Plan/proposal/order fulfilled by this request.
	BasedOn []Reference `bson:"basedOn,omitempty" json:"basedOn,omitempty" xml:"basedOn,omitempty"`
	// Category of supply, e.g.
	//
	//  central, non-stock, etc. This is used to support work flows associated with the supply process.
	Category *CodeableConcept `bson:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
	// Indicates how quickly this SupplyRequest should be addressed with respect to other requests.
	Priority *RequestPriority `bson:"priority,omitempty" json:"priority,omitempty" xml:"priority,omitempty"`
	// The patient to whom the supply will be given or for whom they will be used.
	DeliverFor *Reference `bson:"deliverFor,omitempty" json:"deliverFor,omitempty" xml:"deliverFor,omitempty"`
	// The item that is requested to be supplied.
	//
	// This is either a link to a resource representing the details of the item or a code that identifies the item from a known list.
	//
	// Comment: Note that there's a difference between a prescription - an instruction to take a medication, along with a (sometimes) implicit supply, and an explicit request to supply, with no explicit instructions.
	Item CodeableReference `bson:"item" json:"item" xml:"item"`
	// The amount that is being ordered of the indicated item.
	Quantity Quantity `bson:"quantity" json:"quantity" xml:"quantity"`
	// Specific parameters for the ordered item.  For example, the size of the indicated item.
	Parameter []SupplyRequestParameter `bson:"parameter,omitempty" json:"parameter,omitempty" xml:"parameter,omitempty"`
	// When the request should be fulfilled.
	OccurrenceDateTime *string `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty" xml:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *Period `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty" xml:"occurrencePeriod,omitempty"`
	OccurrenceTiming   *Timing `bson:"occurrenceTiming,omitempty" json:"occurrenceTiming,omitempty" xml:"occurrenceTiming,omitempty"`
	// When the request was made.
	AuthoredOn *string `bson:"authoredOn,omitempty" json:"authoredOn,omitempty" xml:"authoredOn,omitempty"`
	// The device, practitioner, etc. who initiated the request.
	Requester *Reference `bson:"requester,omitempty" json:"requester,omitempty" xml:"requester,omitempty"`
	// Who is intended to fulfill the request.
	Supplier []Reference `bson:"supplier,omitempty" json:"supplier,omitempty" xml:"supplier,omitempty"`
	// The reason why the supply item was requested.
	Reason []CodeableReference `bson:"reason,omitempty" json:"reason,omitempty" xml:"reason,omitempty"`
	// Where the supply is expected to come from.
	DeliverFrom *Reference `bson:"deliverFrom,omitempty" json:"deliverFrom,omitempty" xml:"deliverFrom,omitempty"`
	// Where the supply is destined to go.
	DeliverTo *Reference `bson:"deliverTo,omitempty" json:"deliverTo,omitempty" xml:"deliverTo,omitempty"`
}

type SupplyRequestParameter struct {
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
	// A code or string that identifies the device detail being asserted.
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// The value of the device detail.
	//
	// Comment: Range means device should have a value that falls somewhere within the specified range.
	ValueCodeableConcept *CodeableConcept `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty" xml:"valueCodeableConcept,omitempty"`
	ValueQuantity        *Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty" xml:"valueQuantity,omitempty"`
	ValueRange           *Range           `bson:"valueRange,omitempty" json:"valueRange,omitempty" xml:"valueRange,omitempty"`
	ValueBoolean         *bool            `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty" xml:"valueBoolean,omitempty"`
}

func (r SupplyRequest) ResourceType() string {
	return "SupplyRequest"
}

// MarshalJSON implements [json.Marshaler]
func (r SupplyRequest) MarshalJSON() ([]byte, error) {
	type Wrapper SupplyRequest
	return marshalJSON(r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Marshaler].
func (r SupplyRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Wrapper SupplyRequest
	return marshalXML(e, start, r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Unmarshaler].
func (r *SupplyRequest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper SupplyRequest
	w := Wrapper(*r)
	err := unmarshalXML(d, start, r.ResourceType(), &w)
	*r = SupplyRequest(w)
	return err
}
