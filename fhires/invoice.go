package fhires

import v2 "encoding/json/v2"

// Invoice containing collected ChargeItems from an Account with calculated individual and total price for Billing purpose.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Invoice
type Invoice struct {
	// The logical id of the resource, as used in the URL for the resource.
	//
	// Once assigned, this value never changes.
	//
	// Comment: Within the context of the FHIR RESTful interactions, the resource has an id except for cases like the create and conditional update. Otherwise, the use of the resouce id depends on the given use case.
	ID *string `bson:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// The metadata about the resource.
	//
	// This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *Meta `bson:"meta,omitempty" json:"meta,omitempty" xml:"meta,omitempty"`
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content.
	//
	// Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	//
	// Comment: Asserting this rule set restricts the content to be only understood by a limited set of trading partners. This inherently limits the usefulness of the data in the long term. However, the existing health eco-system is highly fractured, and not yet ready to define, collect, and exchange data in a generally computable sense. Wherever possible, implementers and/or specification writers should avoid using this element. Often, when used, the URL is a reference to an implementation guide that defines these special rules as part of its narrative along with other profiles, value sets, etc.
	ImplicitRules *string `bson:"implicitRules,omitempty" json:"implicitRules,omitempty" xml:"implicitRules,omitempty"`
	// The base language in which the resource is written.
	//
	// Comment: Language is provided to support indexing and accessibility (typically, services such as text to speech use the language tag). The html language tag in the narrative applies  to the narrative. The language tag on the resource may be used to specify the language of other presentations generated from the data in the resource. Not all the content has to be in the base language. The Resource.language should not be assumed to apply to the narrative automatically. If a language is specified, it should it also be specified on the div element in the html (see rules in HTML5 for information about the relationship between xml:lang and the html lang attribute).
	Language *string `bson:"language,omitempty" json:"language,omitempty" xml:"language,omitempty"`
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

// MarshalJSON implements json.Marshaler
func (r Invoice) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		Invoice
		ResourceType string `json:"resourceType"`
	}{
		Invoice:      r,
		ResourceType: r.ResourceType(),
	})
}
