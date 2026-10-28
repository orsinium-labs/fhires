package fhires

import "encoding/json/v2"

// This resource provides the details including amount of a payment and allocates the payment items being paid.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/PaymentReconciliation
type PaymentReconciliation struct {
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
	// A unique identifier assigned to this payment reconciliation.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Code to indicate the nature of the payment such as payment, adjustment.
	Type CodeableConcept `bson:"type" json:"type" xml:"type"`
	// The status of the resource instance.
	//
	// Comment: This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid.
	Status FinancialResourceStatusCodes `bson:"status" json:"status" xml:"status"`
	// The workflow or activity which gave rise to or during which the payment ocurred such as a kiosk, deposit on account, periodic payment etc.
	Kind *CodeableConcept `bson:"kind,omitempty" json:"kind,omitempty" xml:"kind,omitempty"`
	// The period of time for which payments have been gathered into this bulk payment for settlement.
	Period *Period `bson:"period,omitempty" json:"period,omitempty" xml:"period,omitempty"`
	// The date when the resource was created.
	Created string `bson:"created" json:"created" xml:"created"`
	// Payment enterer if not the actual payment issuer.
	Enterer *Reference `bson:"enterer,omitempty" json:"enterer,omitempty" xml:"enterer,omitempty"`
	// The type of the source such as patient or insurance.
	IssuerType *CodeableConcept `bson:"issuerType,omitempty" json:"issuerType,omitempty" xml:"issuerType,omitempty"`
	// The party who generated the payment.
	//
	// Comment: This party is also responsible for the reconciliation.
	PaymentIssuer *Reference `bson:"paymentIssuer,omitempty" json:"paymentIssuer,omitempty" xml:"paymentIssuer,omitempty"`
	// Original request resource reference.
	Request *Reference `bson:"request,omitempty" json:"request,omitempty" xml:"request,omitempty"`
	// The practitioner who is responsible for the services rendered to the patient.
	Requestor *Reference `bson:"requestor,omitempty" json:"requestor,omitempty" xml:"requestor,omitempty"`
	// The outcome of a request for a reconciliation.
	//
	// Comment: The resource may be used to indicate that: the request has been held (queued) for processing; that it has been processed and errors found (error); that no errors were found and that some of the adjudication has been undertaken (partial) or that all of the adjudication has been undertaken (complete).
	Outcome *PaymentOutcome `bson:"outcome,omitempty" json:"outcome,omitempty" xml:"outcome,omitempty"`
	// A human readable description of the status of the request for the reconciliation.
	Disposition *string `bson:"disposition,omitempty" json:"disposition,omitempty" xml:"disposition,omitempty"`
	// The date of payment as indicated on the financial instrument.
	Date string `bson:"date" json:"date" xml:"date"`
	// The location of the site or device for electronic transfers or physical location for cash payments.
	Location *Reference `bson:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// The means of payment such as check, card cash, or electronic funds transfer.
	Method *CodeableConcept `bson:"method,omitempty" json:"method,omitempty" xml:"method,omitempty"`
	// The card brand such as debit, Visa, Amex etc. used if a card is the method of payment.
	CardBrand *string `bson:"cardBrand,omitempty" json:"cardBrand,omitempty" xml:"cardBrand,omitempty"`
	// A portion of the account number, often the last 4 digits, used for verification not charging purposes.
	AccountNumber *string `bson:"accountNumber,omitempty" json:"accountNumber,omitempty" xml:"accountNumber,omitempty"`
	// The year and month (YYYY-MM) when the instrument, typically card, expires.
	ExpirationDate *string `bson:"expirationDate,omitempty" json:"expirationDate,omitempty" xml:"expirationDate,omitempty"`
	// The name of the card processor, etf processor, bank for checks.
	Processor *string `bson:"processor,omitempty" json:"processor,omitempty" xml:"processor,omitempty"`
	// The check number, eft reference, car processor reference.
	ReferenceNumber *string `bson:"referenceNumber,omitempty" json:"referenceNumber,omitempty" xml:"referenceNumber,omitempty"`
	// An alphanumeric issued by the processor to confirm the successful issuance of payment.
	Authorization *string `bson:"authorization,omitempty" json:"authorization,omitempty" xml:"authorization,omitempty"`
	// The amount offered by the issuer, typically applies to cash when the issuer provides an amount in bank note denominations equal to or excess of the amount actually being paid.
	TenderedAmount *Money `bson:"tenderedAmount,omitempty" json:"tenderedAmount,omitempty" xml:"tenderedAmount,omitempty"`
	// The amount returned by the receiver which is excess to the amount payable, often referred to as 'change'.
	ReturnedAmount *Money `bson:"returnedAmount,omitempty" json:"returnedAmount,omitempty" xml:"returnedAmount,omitempty"`
	// Total payment amount as indicated on the financial instrument.
	Amount Money `bson:"amount" json:"amount" xml:"amount"`
	// Issuer's unique identifier for the payment instrument.
	//
	// Comment: For example: EFT number or check number.
	PaymentIdentifier *Identifier `bson:"paymentIdentifier,omitempty" json:"paymentIdentifier,omitempty" xml:"paymentIdentifier,omitempty"`
	// Distribution of the payment amount for a previously acknowledged payable.
	Allocation []PaymentReconciliationAllocation `bson:"allocation,omitempty" json:"allocation,omitempty" xml:"allocation,omitempty"`
	// A code for the form to be used for printing the content.
	//
	// Comment: May be needed to identify specific jurisdictional forms.
	FormCode *CodeableConcept `bson:"formCode,omitempty" json:"formCode,omitempty" xml:"formCode,omitempty"`
	// A note that describes or explains the processing in a human readable form.
	ProcessNote []PaymentReconciliationProcessNote `bson:"processNote,omitempty" json:"processNote,omitempty" xml:"processNote,omitempty"`
}

type PaymentReconciliationAllocation struct {
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
	// Unique identifier for the current payment item for the referenced payable.
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Unique identifier for the prior payment item for the referenced payable.
	Predecessor *Identifier `bson:"predecessor,omitempty" json:"predecessor,omitempty" xml:"predecessor,omitempty"`
	// Specific resource to which the payment/adjustment/advance applies.
	Target *Reference `bson:"target,omitempty" json:"target,omitempty" xml:"target,omitempty"`
	//  Identifies the claim line item, encounter or other sub-element being paid.
	//
	// Note payment may be partial, that is not match the then outstanding balance or amount incurred.
	TargetItemString      *string     `bson:"targetItemString,omitempty" json:"targetItemString,omitempty" xml:"targetItemString,omitempty"`
	TargetItemIdentifier  *Identifier `bson:"targetItemIdentifier,omitempty" json:"targetItemIdentifier,omitempty" xml:"targetItemIdentifier,omitempty"`
	TargetItemPositiveInt *int        `bson:"targetItemPositiveInt,omitempty" json:"targetItemPositiveInt,omitempty" xml:"targetItemPositiveInt,omitempty"`
	// The Encounter to which this payment applies, may be completed by the receiver, used for search.
	Encounter *Reference `bson:"encounter,omitempty" json:"encounter,omitempty" xml:"encounter,omitempty"`
	// The Account to which this payment applies, may be completed by the receiver, used for search.
	Account *Reference `bson:"account,omitempty" json:"account,omitempty" xml:"account,omitempty"`
	// Code to indicate the nature of the payment.
	//
	// Comment: For example: payment, adjustment, funds advance, etc.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The party which submitted the claim or financial transaction.
	Submitter *Reference `bson:"submitter,omitempty" json:"submitter,omitempty" xml:"submitter,omitempty"`
	// A resource, such as a ClaimResponse, which contains a commitment to payment.
	Response *Reference `bson:"response,omitempty" json:"response,omitempty" xml:"response,omitempty"`
	// The date from the response resource containing a commitment to pay.
	Date *string `bson:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// A reference to the individual who is responsible for inquiries regarding the response and its payment.
	Responsible *Reference `bson:"responsible,omitempty" json:"responsible,omitempty" xml:"responsible,omitempty"`
	// The party which is receiving the payment.
	Payee *Reference `bson:"payee,omitempty" json:"payee,omitempty" xml:"payee,omitempty"`
	// The monetary amount allocated from the total payment to the payable.
	Amount *Money `bson:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
}

type PaymentReconciliationProcessNote struct {
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
	// The business purpose of the note text.
	Type *NoteType `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The explanation or description associated with the processing.
	Text *string `bson:"text,omitempty" json:"text,omitempty" xml:"text,omitempty"`
}

func (r PaymentReconciliation) ResourceType() string {
	return "PaymentReconciliation"
}

// MarshalJSON implements json.Marshaler
func (r PaymentReconciliation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		PaymentReconciliation
		ResourceType string `json:"resourceType"`
	}{
		PaymentReconciliation: r,
		ResourceType:          r.ResourceType(),
	})
}
