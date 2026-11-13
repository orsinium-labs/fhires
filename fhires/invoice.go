package fhires

import "encoding/xml"

// Invoice containing collected ChargeItems from an Account with calculated individual and total price for Billing purpose.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Invoice
type Invoice struct {
	DomainResource
	// Identifier of this Invoice, often used for reference in correspondence about this invoice or for tracking of payments.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The current state of the Invoice.
	Status InvoiceStatus `bson:"status" json:"status" xml:"status"`
	// In case of Invoice cancellation a reason must be given (entered in error, superseded by corrected invoice etc.).
	//
	// Comment: Derived Profiles may choose to add invariants requiring this field to be populated if either priceOverride or factorOverride have been filled.
	CancelledReason *string `bson:"cancelledReason,omitempty" json:"cancelledReason,omitempty" xml:"cancelledReason,omitempty"`
	// Type of Invoice depending on domain, realm an usage (e.g.
	//
	// internal/external, dental, preliminary).
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The individual or set of individuals receiving the goods and services billed in this invoice.
	Subject *Reference `bson:"subject,omitempty" json:"subject,omitempty" xml:"subject,omitempty"`
	// The individual or Organization responsible for balancing of this invoice.
	Recipient *Reference `bson:"recipient,omitempty" json:"recipient,omitempty" xml:"recipient,omitempty"`
	// Depricared by the element below.
	Date *string `bson:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// Date/time(s) of when this Invoice was posted.
	//
	// Comment: The list of types may be constrained as appropriate for the type of charge item.
	Creation *string `bson:"creation,omitempty" json:"creation,omitempty" xml:"creation,omitempty"`
	// Date/time(s) range of services included in this invoice.
	PeriodDate   *string `bson:"periodDate,omitempty" json:"periodDate,omitempty" xml:"periodDate,omitempty"`
	PeriodPeriod *Period `bson:"periodPeriod,omitempty" json:"periodPeriod,omitempty" xml:"periodPeriod,omitempty"`
	// Indicates who or what performed or participated in the charged service.
	Participant []InvoiceParticipant `bson:"participant,omitempty" json:"participant,omitempty" xml:"participant,omitempty"`
	// The organizationissuing the Invoice.
	//
	// Comment: Practitioners and Devices can be associated with multiple organizations. It has to be made clear, on behalf of which Organization the services have been rendered.
	Issuer *Reference `bson:"issuer,omitempty" json:"issuer,omitempty" xml:"issuer,omitempty"`
	// Account which is supposed to be balanced with this Invoice.
	//
	// Comment: Systems posting the ChargeItems might not always be able to determine, which accounts the Items need to be places into. It is up to the potprocessing Financial System to apply internal rules to decide based on the Encounter/EpisodeOfCare/Patient/Coverage context and the type of ChargeItem, which Account is appropriate.
	Account *Reference `bson:"account,omitempty" json:"account,omitempty" xml:"account,omitempty"`
	// Each line item represents one charge for goods and services rendered.
	//
	// Details such.ofType(date), code and amount are found in the referenced ChargeItem resource.
	LineItem []InvoiceLineItem `bson:"lineItem,omitempty" json:"lineItem,omitempty" xml:"lineItem,omitempty"`
	// The total amount for the Invoice may be calculated as the sum of the line items with surcharges/deductions that apply in certain conditions.
	//
	//  The priceComponent element can be used to offer transparency to the recipient of the Invoice of how the total price was calculated.
	TotalPriceComponent []MonetaryComponent `bson:"totalPriceComponent,omitempty" json:"totalPriceComponent,omitempty" xml:"totalPriceComponent,omitempty"`
	// Invoice total , taxes excluded.
	//
	// Comment: There is no reason to carry the price in the instance of a ChargeItem unless circumstances require a manual override. The list prices or are usually defined in a back catalogue of the billing codes  (see ChargeItem.definition). Derived profiles may require a ChargeItem.overrideReason to be provided if either factor or price are manually overridden.
	TotalNet *Money `bson:"totalNet,omitempty" json:"totalNet,omitempty" xml:"totalNet,omitempty"`
	// Invoice total, tax included.
	//
	// Comment: There is no reason to carry the price in the instance of a ChargeItem unless circumstances require a manual override. The list prices or are usually defined in a back catalogue of the billing codes  (see ChargeItem.definition). Derived profiles may require a ChargeItem.overrideReason to be provided if either factor or price are manually overridden.
	TotalGross *Money `bson:"totalGross,omitempty" json:"totalGross,omitempty" xml:"totalGross,omitempty"`
	// Payment details such as banking details, period of payment, deductibles, methods of payment.
	//
	// Comment: Derived Profiles may chose to add invariants requiring this field to be populated if either priceOverride or factorOverride have been filled.
	PaymentTerms *string `bson:"paymentTerms,omitempty" json:"paymentTerms,omitempty" xml:"paymentTerms,omitempty"`
	// Comments made about the invoice by the issuer, subject, or other participants.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
}

type InvoiceParticipant struct {
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
	// Describes the type of involvement (e.g.
	//
	// transcriptionist, creator etc.). If the invoice has been created automatically, the Participant may be a billing engine or another kind of device.
	Role *CodeableConcept `bson:"role,omitempty" json:"role,omitempty" xml:"role,omitempty"`
	// The device, practitioner, etc. who performed or participated in the service.
	Actor Reference `bson:"actor" json:"actor" xml:"actor"`
}

type InvoiceLineItem struct {
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
	// Sequence in which the items appear on the invoice.
	Sequence *int `bson:"sequence,omitempty" json:"sequence,omitempty" xml:"sequence,omitempty"`
	// Date/time(s) range when this service was delivered or completed.
	ServicedDate   *string `bson:"servicedDate,omitempty" json:"servicedDate,omitempty" xml:"servicedDate,omitempty"`
	ServicedPeriod *Period `bson:"servicedPeriod,omitempty" json:"servicedPeriod,omitempty" xml:"servicedPeriod,omitempty"`
	// The ChargeItem contains information such as the billing code, date, amount etc.
	//
	// If no further details are required for the lineItem, inline billing codes can be added using the CodeableConcept data type instead of the Reference.
	ChargeItemReference       Reference       `bson:"chargeItemReference" json:"chargeItemReference" xml:"chargeItemReference"`
	ChargeItemCodeableConcept CodeableConcept `bson:"chargeItemCodeableConcept" json:"chargeItemCodeableConcept" xml:"chargeItemCodeableConcept"`
	// The price for a ChargeItem may be calculated as a base price with surcharges/deductions that apply in certain conditions.
	//
	// A ChargeItemDefinition resource that defines the prices, factors and conditions that apply to a billing code is currently under development. The priceComponent element can be used to offer transparency to the recipient of the Invoice as to how the prices have been calculated.
	PriceComponent []MonetaryComponent `bson:"priceComponent,omitempty" json:"priceComponent,omitempty" xml:"priceComponent,omitempty"`
}

func (r Invoice) ResourceType() string {
	return "Invoice"
}

// MarshalJSON implements [json.Marshaler]
func (r Invoice) MarshalJSON() ([]byte, error) {
	type Wrapper Invoice
	return marshalJSON(r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Marshaler].
func (r Invoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Wrapper Invoice
	return marshalXML(e, start, r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Unmarshaler].
func (r *Invoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper Invoice
	w := Wrapper(*r)
	err := unmarshalXML(d, start, r.ResourceType(), &w)
	*r = Invoice(w)
	return err
}
