package fhires

import (
	"encoding/json/v2"

	decimal "github.com/shopspring/decimal"
)

// This resource provides the adjudication details from the processing of a Claim resource.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/ClaimResponse
type ClaimResponse struct {
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
	// A unique identifier assigned to this claim response.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Trace number for tracking purposes.
	//
	// May be defined at the jurisdiction level or between trading partners.
	TraceNumber []Identifier `bson:"traceNumber,omitempty" json:"traceNumber,omitempty" xml:"traceNumber,omitempty"`
	// The status of the resource instance.
	//
	// Comment: This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid.
	Status FinancialResourceStatusCodes `bson:"status" json:"status" xml:"status"`
	// A finer grained suite of claim type codes which may convey additional information such as Inpatient vs Outpatient and/or a specialty service.
	//
	// Comment: This may contain the local bill type codes, for example the US UB-04 bill type code or the CMS bill type.
	Type CodeableConcept `bson:"type" json:"type" xml:"type"`
	// A finer grained suite of claim type codes which may convey additional information such as Inpatient vs Outpatient and/or a specialty service.
	//
	// Comment: This may contain the local bill type codes, for example the US UB-04 bill type code or the CMS bill type.
	SubType *CodeableConcept `bson:"subType,omitempty" json:"subType,omitempty" xml:"subType,omitempty"`
	// A code to indicate whether the nature of the request is: Claim - A request to an Insurer to adjudicate the supplied charges for health care goods and services under the identified policy and to pay the determined Benefit amount, if any; Preauthorization - A request to an Insurer to adjudicate the supplied proposed future charges for health care goods and services under the identified policy and to approve the services and provide the expected benefit amounts and potentially to reserve funds to pay the benefits when Claims for the indicated services are later submitted; or, Pre-determination - A request to an Insurer to adjudicate the supplied 'what if' charges for health care goods and services under the identified policy and report back what the Benefit payable would be had the services actually been provided.
	Use Use `bson:"use" json:"use" xml:"use"`
	// The party to whom the professional services and/or products have been supplied or are being considered and for whom actual for facast reimbursement is sought.
	Patient Reference `bson:"patient" json:"patient" xml:"patient"`
	// The date this resource was created.
	Created string `bson:"created" json:"created" xml:"created"`
	// The party responsible for authorization, adjudication and reimbursement.
	Insurer *Reference `bson:"insurer,omitempty" json:"insurer,omitempty" xml:"insurer,omitempty"`
	// The provider which is responsible for the claim, predetermination or preauthorization.
	//
	// Comment: Typically this field would be 1..1 where this party is accountable for the data content within the claim but is not necessarily the facility, provider group or practitioner who provided the products and services listed within this claim resource. This field is the Billing Provider, for example, a facility, provider group, lab or practitioner..
	Requestor *Reference `bson:"requestor,omitempty" json:"requestor,omitempty" xml:"requestor,omitempty"`
	// Original request resource reference.
	Request *Reference `bson:"request,omitempty" json:"request,omitempty" xml:"request,omitempty"`
	// The outcome of the claim, predetermination, or preauthorization processing.
	//
	// Comment: The resource may be used to indicate that the Claim/Preauthorization/Pre-determination has been received but processing has not begun (queued); that it has been processed and one or more errors have been detected (error); no errors were detected and some of the adjudication processing has been performed (partial); or all of the adjudication processing has completed without errors (complete).
	Outcome ClaimProcessingCodes `bson:"outcome" json:"outcome" xml:"outcome"`
	// The result of the claim, predetermination, or preauthorization adjudication.
	//
	// Comment: The element is used to indicate the current state of the adjudication overall for the claim resource, for example: the request has been held (pended) for adjudication processing, for manual review or other reasons; that it has been processed and will be paid, or the outstanding paid, as submitted (approved); that no amount will be paid (denied); or that some amount between zero and the submitted amount will be paid (partial).
	Decision *CodeableConcept `bson:"decision,omitempty" json:"decision,omitempty" xml:"decision,omitempty"`
	// A human readable description of the status of the adjudication.
	Disposition *string `bson:"disposition,omitempty" json:"disposition,omitempty" xml:"disposition,omitempty"`
	// Reference from the Insurer which is used in later communications which refers to this adjudication.
	//
	// Comment: This value is only present on preauthorization adjudications.
	PreAuthRef *string `bson:"preAuthRef,omitempty" json:"preAuthRef,omitempty" xml:"preAuthRef,omitempty"`
	// The time frame during which this authorization is effective.
	PreAuthPeriod *Period `bson:"preAuthPeriod,omitempty" json:"preAuthPeriod,omitempty" xml:"preAuthPeriod,omitempty"`
	// Information code for an event with a corresponding date or period.
	Event []ClaimResponseEvent `bson:"event,omitempty" json:"event,omitempty" xml:"event,omitempty"`
	// Type of Party to be reimbursed: subscriber, provider, other.
	PayeeType *CodeableConcept `bson:"payeeType,omitempty" json:"payeeType,omitempty" xml:"payeeType,omitempty"`
	// Healthcare encounters related to this claim.
	//
	// Comment: This will typically be the encounter the event occurred within, but some activities may be initiated prior to or after the official completion of an encounter but still be tied to the context of the encounter.
	Encounter []Reference `bson:"encounter,omitempty" json:"encounter,omitempty" xml:"encounter,omitempty"`
	// A package billing code or bundle code used to group products and services to a particular health condition (such as heart attack) which is based on a predetermined grouping code system.
	//
	// Comment: For example DRG (Diagnosis Related Group) or a bundled billing code. A patient may have a diagnosis of a Myocardial Infarction and a DRG for HeartAttack would be assigned. The Claim item (and possible subsequent claims) would refer to the DRG for those line items that were for services related to the heart attack event.
	DiagnosisRelatedGroup *CodeableConcept `bson:"diagnosisRelatedGroup,omitempty" json:"diagnosisRelatedGroup,omitempty" xml:"diagnosisRelatedGroup,omitempty"`
	// A claim line.
	//
	// Either a simple (a product or service) or a 'group' of details which can also be a simple items or groups of sub-details.
	Item []ClaimResponseItem `bson:"item,omitempty" json:"item,omitempty" xml:"item,omitempty"`
	// The first-tier service adjudications for payor added product or service lines.
	AddItem []ClaimResponseAddItem `bson:"addItem,omitempty" json:"addItem,omitempty" xml:"addItem,omitempty"`
	// The adjudication results which are presented at the header level rather than at the line-item or add-item levels.
	Adjudication []ClaimResponseItemAdjudication `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	// Categorized monetary totals for the adjudication.
	//
	// Comment: Totals for amounts submitted, co-pays, benefits payable etc.
	Total []ClaimResponseTotal `bson:"total,omitempty" json:"total,omitempty" xml:"total,omitempty"`
	// Payment details for the adjudication of the claim.
	Payment *ClaimResponsePayment `bson:"payment,omitempty" json:"payment,omitempty" xml:"payment,omitempty"`
	// A code, used only on a response to a preauthorization, to indicate whether the benefits payable have been reserved and for whom.
	//
	// Comment: Fund would be release by a future claim quoting the preAuthRef of this response. Examples of values include: provider, patient, none.
	FundsReserve *CodeableConcept `bson:"fundsReserve,omitempty" json:"fundsReserve,omitempty" xml:"fundsReserve,omitempty"`
	// A code for the form to be used for printing the content.
	//
	// Comment: May be needed to identify specific jurisdictional forms.
	FormCode *CodeableConcept `bson:"formCode,omitempty" json:"formCode,omitempty" xml:"formCode,omitempty"`
	// The actual form, by reference or inclusion, for printing the content or an EOB.
	//
	// Comment: Needed to permit insurers to include the actual form.
	Form *Attachment `bson:"form,omitempty" json:"form,omitempty" xml:"form,omitempty"`
	// A note that describes or explains adjudication results in a human readable form.
	ProcessNote []ClaimResponseProcessNote `bson:"processNote,omitempty" json:"processNote,omitempty" xml:"processNote,omitempty"`
	// Request for additional supporting or authorizing information.
	//
	// Comment: For example: professional reports, documents, images, clinical resources, or accident reports.
	CommunicationRequest []Reference `bson:"communicationRequest,omitempty" json:"communicationRequest,omitempty" xml:"communicationRequest,omitempty"`
	// Financial instruments for reimbursement for the health care products and services specified on the claim.
	//
	// Comment: All insurance coverages for the patient which may be applicable for reimbursement, of the products and services listed in the claim, are typically provided in the claim to allow insurers to confirm the ordering of the insurance coverages relative to local 'coordination of benefit' rules. One coverage (and only one) with 'focal=true' is to be used in the adjudication of this claim. Coverages appearing before the focal Coverage in the list, and where 'subrogation=false', should provide a reference to the ClaimResponse containing the adjudication results of the prior claim.
	Insurance []ClaimResponseInsurance `bson:"insurance,omitempty" json:"insurance,omitempty" xml:"insurance,omitempty"`
	// Errors encountered during the processing of the adjudication.
	//
	// Comment: If the request contains errors then an error element should be provided and no adjudication related sections (item, addItem, or payment) should be present.
	Error []ClaimResponseError `bson:"error,omitempty" json:"error,omitempty" xml:"error,omitempty"`
}

type ClaimResponseEvent struct {
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
	// A coded event such as when a service is expected or a card printed.
	Type CodeableConcept `bson:"type" json:"type" xml:"type"`
	// A date or period in the past or future indicating when the event occurred or is expectd to occur.
	WhenDateTime string `bson:"whenDateTime" json:"whenDateTime" xml:"whenDateTime"`
	WhenPeriod   Period `bson:"whenPeriod" json:"whenPeriod" xml:"whenPeriod"`
}

type ClaimResponseItem struct {
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
	// A number to uniquely reference the claim item entries.
	ItemSequence int `bson:"itemSequence" json:"itemSequence" xml:"itemSequence"`
	// Trace number for tracking purposes.
	//
	// May be defined at the jurisdiction level or between trading partners.
	TraceNumber []Identifier `bson:"traceNumber,omitempty" json:"traceNumber,omitempty" xml:"traceNumber,omitempty"`
	// The numbers associated with notes below which apply to the adjudication of this item.
	NoteNumber []int `bson:"noteNumber,omitempty" json:"noteNumber,omitempty" xml:"noteNumber,omitempty"`
	// The high-level results of the adjudication if adjudication has been performed.
	ReviewOutcome *ClaimResponseItemReviewOutcome `bson:"reviewOutcome,omitempty" json:"reviewOutcome,omitempty" xml:"reviewOutcome,omitempty"`
	// If this item is a group then the values here are a summary of the adjudication of the detail items.
	//
	// If this item is a simple product or service then this is the result of the adjudication of this item.
	Adjudication []ClaimResponseItemAdjudication `bson:"adjudication,omitempty" json:"adjudication,omitempty" xml:"adjudication,omitempty"`
	// A claim detail.
	//
	// Either a simple (a product or service) or a 'group' of sub-details which are simple items.
	Detail []ClaimResponseItemDetail `bson:"detail,omitempty" json:"detail,omitempty" xml:"detail,omitempty"`
}

type ClaimResponseItemReviewOutcome struct {
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
	// The result of the claim, predetermination, or preauthorization adjudication.
	//
	// Comment: The element is used to indicate the current state of the adjudication overall for the claim resource, for example: the request has been held (pended) for adjudication processing, for manual review or other reasons; that it has been processed and will be paid, or the outstanding paid, as submitted (approved); that no amount will be paid (denied); or that some amount between zero and the submitted amount will be paid (partial).
	Decision *CodeableConcept `bson:"decision,omitempty" json:"decision,omitempty" xml:"decision,omitempty"`
	// The reasons for the result of the claim, predetermination, or preauthorization adjudication.
	Reason []CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty" xml:"reason,omitempty"`
	// Reference from the Insurer which is used in later communications which refers to this adjudication.
	//
	// Comment: This value is only present on preauthorization adjudications.
	PreAuthRef *string `bson:"preAuthRef,omitempty" json:"preAuthRef,omitempty" xml:"preAuthRef,omitempty"`
	// The time frame during which this authorization is effective.
	PreAuthPeriod *Period `bson:"preAuthPeriod,omitempty" json:"preAuthPeriod,omitempty" xml:"preAuthPeriod,omitempty"`
}

type ClaimResponseItemAdjudication struct {
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
	// A code to indicate the information type of this adjudication record.
	//
	// Information types may include the value submitted, maximum values or percentages allowed or payable under the plan, amounts that: the patient is responsible for in aggregate or pertaining to this item; amounts paid by other coverages; and, the benefit payable for this item.
	//
	// Comment: For example codes indicating: Co-Pay, deductible, eligible, benefit, tax, etc.
	Category CodeableConcept `bson:"category" json:"category" xml:"category"`
	// A code supporting the understanding of the adjudication result and explaining variance from expected amount.
	//
	// Comment: For example may indicate that the funds for this benefit type have been exhausted.
	Reason *CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty" xml:"reason,omitempty"`
	// Monetary amount associated with the category.
	//
	// Comment: For example: amount submitted, eligible amount, co-payment, and benefit payable.
	Amount *Money `bson:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
	// A non-monetary value associated with the category.
	//
	// Mutually exclusive to the amount element above.
	//
	// Comment: For example: eligible percentage or co-payment percentage.
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty" xml:"quantity,omitempty"`
}

type ClaimResponseItemDetail struct {
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
	// A number to uniquely reference the claim detail entry.
	DetailSequence int `bson:"detailSequence" json:"detailSequence" xml:"detailSequence"`
	// Trace number for tracking purposes.
	//
	// May be defined at the jurisdiction level or between trading partners.
	TraceNumber []Identifier `bson:"traceNumber,omitempty" json:"traceNumber,omitempty" xml:"traceNumber,omitempty"`
	// The numbers associated with notes below which apply to the adjudication of this item.
	NoteNumber []int `bson:"noteNumber,omitempty" json:"noteNumber,omitempty" xml:"noteNumber,omitempty"`
	// The high-level results of the adjudication if adjudication has been performed.
	ReviewOutcome *ClaimResponseItemReviewOutcome `bson:"reviewOutcome,omitempty" json:"reviewOutcome,omitempty"`
	// The adjudication results.
	Adjudication []ClaimResponseItemAdjudication `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	// A sub-detail adjudication of a simple product or service.
	SubDetail []ClaimResponseItemDetailSubDetail `bson:"subDetail,omitempty" json:"subDetail,omitempty" xml:"subDetail,omitempty"`
}

type ClaimResponseItemDetailSubDetail struct {
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
	// A number to uniquely reference the claim sub-detail entry.
	SubDetailSequence int `bson:"subDetailSequence" json:"subDetailSequence" xml:"subDetailSequence"`
	// Trace number for tracking purposes.
	//
	// May be defined at the jurisdiction level or between trading partners.
	TraceNumber []Identifier `bson:"traceNumber,omitempty" json:"traceNumber,omitempty" xml:"traceNumber,omitempty"`
	// The numbers associated with notes below which apply to the adjudication of this item.
	NoteNumber []int `bson:"noteNumber,omitempty" json:"noteNumber,omitempty" xml:"noteNumber,omitempty"`
	// The high-level results of the adjudication if adjudication has been performed.
	ReviewOutcome *ClaimResponseItemReviewOutcome `bson:"reviewOutcome,omitempty" json:"reviewOutcome,omitempty"`
	// The adjudication results.
	Adjudication []ClaimResponseItemAdjudication `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
}

type ClaimResponseAddItem struct {
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
	// Claim items which this service line is intended to replace.
	ItemSequence []int `bson:"itemSequence,omitempty" json:"itemSequence,omitempty" xml:"itemSequence,omitempty"`
	// The sequence number of the details within the claim item which this line is intended to replace.
	DetailSequence []int `bson:"detailSequence,omitempty" json:"detailSequence,omitempty" xml:"detailSequence,omitempty"`
	// The sequence number of the sub-details within the details within the claim item which this line is intended to replace.
	SubdetailSequence []int `bson:"subdetailSequence,omitempty" json:"subdetailSequence,omitempty" xml:"subdetailSequence,omitempty"`
	// Trace number for tracking purposes.
	//
	// May be defined at the jurisdiction level or between trading partners.
	TraceNumber []Identifier `bson:"traceNumber,omitempty" json:"traceNumber,omitempty" xml:"traceNumber,omitempty"`
	// The providers who are authorized for the services rendered to the patient.
	Provider []Reference `bson:"provider,omitempty" json:"provider,omitempty" xml:"provider,omitempty"`
	// The type of revenue or cost center providing the product and/or service.
	Revenue *CodeableConcept `bson:"revenue,omitempty" json:"revenue,omitempty" xml:"revenue,omitempty"`
	// When the value is a group code then this item collects a set of related item details, otherwise this contains the product, service, drug or other billing code for the item.
	//
	// This element may be the start of a range of .productOrService codes used in conjunction with .productOrServiceEnd or it may be a solo element where .productOrServiceEnd is not used.
	//
	// Comment: If this is an actual service or product line, i.e. not a Group, then use code to indicate the Professional Service or Product supplied (e.g. CTP, HCPCS, USCLS, ICD10, NCPDP, DIN, RxNorm, ACHI, CCI). If a grouping item then use a group code to indicate the type of thing being grouped e.g. 'glasses' or 'compound'.
	ProductOrService *CodeableConcept `bson:"productOrService,omitempty" json:"productOrService,omitempty" xml:"productOrService,omitempty"`
	// This contains the end of a range of product, service, drug or other billing codes for the item.
	//
	// This element is not used when the .productOrService is a group code. This value may only be present when a .productOfService code has been provided to convey the start of the range. Typically this value may be used only with preauthorizations and not with claims.
	ProductOrServiceEnd *CodeableConcept `bson:"productOrServiceEnd,omitempty" json:"productOrServiceEnd,omitempty" xml:"productOrServiceEnd,omitempty"`
	// Request or Referral for Goods or Service to be rendered.
	Request []Reference `bson:"request,omitempty" json:"request,omitempty" xml:"request,omitempty"`
	// Item typification or modifiers codes to convey additional context for the product or service.
	//
	// Comment: For example in Oral whether the treatment is cosmetic or associated with TMJ, or for Medical whether the treatment was outside the clinic or outside of office hours.
	Modifier []CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty" xml:"modifier,omitempty"`
	// Identifies the program under which this may be recovered.
	//
	// Comment: For example: Neonatal program, child dental program or drug users recovery program.
	ProgramCode []CodeableConcept `bson:"programCode,omitempty" json:"programCode,omitempty" xml:"programCode,omitempty"`
	// The date or dates when the service or product was supplied, performed or completed.
	ServicedDate   *string `bson:"servicedDate,omitempty" json:"servicedDate,omitempty" xml:"servicedDate,omitempty"`
	ServicedPeriod *Period `bson:"servicedPeriod,omitempty" json:"servicedPeriod,omitempty" xml:"servicedPeriod,omitempty"`
	// Where the product or service was provided.
	LocationCodeableConcept *CodeableConcept `bson:"locationCodeableConcept,omitempty" json:"locationCodeableConcept,omitempty" xml:"locationCodeableConcept,omitempty"`
	LocationAddress         *Address         `bson:"locationAddress,omitempty" json:"locationAddress,omitempty" xml:"locationAddress,omitempty"`
	LocationReference       *Reference       `bson:"locationReference,omitempty" json:"locationReference,omitempty" xml:"locationReference,omitempty"`
	// The number of repetitions of a service or product.
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty" xml:"quantity,omitempty"`
	// If the item is not a group then this is the fee for the product or service, otherwise this is the total of the fees for the details of the group.
	UnitPrice *Money `bson:"unitPrice,omitempty" json:"unitPrice,omitempty" xml:"unitPrice,omitempty"`
	// A real number that represents a multiplier used in determining the overall value of services delivered and/or goods received.
	//
	// The concept of a Factor allows for a discount or surcharge multiplier to be applied to a monetary amount.
	//
	// Comment: To show a 10% senior's discount, the value entered is: 0.90 (1.00 - 0.10).
	Factor *decimal.Decimal `bson:"factor,omitempty" json:"factor,omitempty" xml:"factor,omitempty"`
	// The total of taxes applicable for this product or service.
	Tax *Money `bson:"tax,omitempty" json:"tax,omitempty" xml:"tax,omitempty"`
	// The total amount claimed for the group (if a grouper) or the addItem.
	//
	// Net = unit price * quantity * factor.
	//
	// Comment: For example, the formula: quantity * unitPrice * factor  = net. Quantity and factor are assumed to be 1 if not supplied.
	Net *Money `bson:"net,omitempty" json:"net,omitempty" xml:"net,omitempty"`
	// Physical location where the service is performed or applies.
	BodySite []ClaimResponseAddItemBodySite `bson:"bodySite,omitempty" json:"bodySite,omitempty" xml:"bodySite,omitempty"`
	// The numbers associated with notes below which apply to the adjudication of this item.
	NoteNumber []int `bson:"noteNumber,omitempty" json:"noteNumber,omitempty" xml:"noteNumber,omitempty"`
	// The high-level results of the adjudication if adjudication has been performed.
	ReviewOutcome *ClaimResponseItemReviewOutcome `bson:"reviewOutcome,omitempty" json:"reviewOutcome,omitempty"`
	// The adjudication results.
	Adjudication []ClaimResponseItemAdjudication `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	// The second-tier service adjudications for payor added services.
	Detail []ClaimResponseAddItemDetail `bson:"detail,omitempty" json:"detail,omitempty" xml:"detail,omitempty"`
}

type ClaimResponseAddItemBodySite struct {
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
	// Physical service site on the patient (limb, tooth, etc.).
	//
	// Comment: For example: Providing a tooth code allows an insurer to identify a provider performing a filling on a tooth that was previously removed.
	Site []CodeableReference `bson:"site" json:"site" xml:"site"`
	// A region or surface of the bodySite, e.g. limb region or tooth surface(s).
	SubSite []CodeableConcept `bson:"subSite,omitempty" json:"subSite,omitempty" xml:"subSite,omitempty"`
}

type ClaimResponseAddItemDetail struct {
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
	// Trace number for tracking purposes.
	//
	// May be defined at the jurisdiction level or between trading partners.
	TraceNumber []Identifier `bson:"traceNumber,omitempty" json:"traceNumber,omitempty" xml:"traceNumber,omitempty"`
	// The type of revenue or cost center providing the product and/or service.
	Revenue *CodeableConcept `bson:"revenue,omitempty" json:"revenue,omitempty" xml:"revenue,omitempty"`
	// When the value is a group code then this item collects a set of related item details, otherwise this contains the product, service, drug or other billing code for the item.
	//
	// This element may be the start of a range of .productOrService codes used in conjunction with .productOrServiceEnd or it may be a solo element where .productOrServiceEnd is not used.
	//
	// Comment: If this is an actual service or product line, i.e. not a Group, then use code to indicate the Professional Service or Product supplied (e.g. CTP, HCPCS, USCLS, ICD10, NCPDP, DIN, RxNorm, ACHI, CCI). If a grouping item then use a group code to indicate the type of thing being grouped e.g. 'glasses' or 'compound'.
	ProductOrService *CodeableConcept `bson:"productOrService,omitempty" json:"productOrService,omitempty" xml:"productOrService,omitempty"`
	// This contains the end of a range of product, service, drug or other billing codes for the item.
	//
	// This element is not used when the .productOrService is a group code. This value may only be present when a .productOfService code has been provided to convey the start of the range. Typically this value may be used only with preauthorizations and not with claims.
	ProductOrServiceEnd *CodeableConcept `bson:"productOrServiceEnd,omitempty" json:"productOrServiceEnd,omitempty" xml:"productOrServiceEnd,omitempty"`
	// Item typification or modifiers codes to convey additional context for the product or service.
	//
	// Comment: For example in Oral whether the treatment is cosmetic or associated with TMJ, or for Medical whether the treatment was outside the clinic or outside of office hours.
	Modifier []CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty" xml:"modifier,omitempty"`
	// The number of repetitions of a service or product.
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty" xml:"quantity,omitempty"`
	// If the item is not a group then this is the fee for the product or service, otherwise this is the total of the fees for the details of the group.
	UnitPrice *Money `bson:"unitPrice,omitempty" json:"unitPrice,omitempty" xml:"unitPrice,omitempty"`
	// A real number that represents a multiplier used in determining the overall value of services delivered and/or goods received.
	//
	// The concept of a Factor allows for a discount or surcharge multiplier to be applied to a monetary amount.
	//
	// Comment: To show a 10% senior's discount, the value entered is: 0.90 (1.00 - 0.10).
	Factor *decimal.Decimal `bson:"factor,omitempty" json:"factor,omitempty" xml:"factor,omitempty"`
	// The total of taxes applicable for this product or service.
	Tax *Money `bson:"tax,omitempty" json:"tax,omitempty" xml:"tax,omitempty"`
	// The total amount claimed for the group (if a grouper) or the addItem.detail.
	//
	// Net = unit price * quantity * factor.
	//
	// Comment: For example, the formula: quantity * unitPrice * factor  = net. Quantity and factor are assumed to be 1 if not supplied.
	Net *Money `bson:"net,omitempty" json:"net,omitempty" xml:"net,omitempty"`
	// The numbers associated with notes below which apply to the adjudication of this item.
	NoteNumber []int `bson:"noteNumber,omitempty" json:"noteNumber,omitempty" xml:"noteNumber,omitempty"`
	// The high-level results of the adjudication if adjudication has been performed.
	ReviewOutcome *ClaimResponseItemReviewOutcome `bson:"reviewOutcome,omitempty" json:"reviewOutcome,omitempty"`
	// The adjudication results.
	Adjudication []ClaimResponseItemAdjudication `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	// The third-tier service adjudications for payor added services.
	SubDetail []ClaimResponseAddItemDetailSubDetail `bson:"subDetail,omitempty" json:"subDetail,omitempty" xml:"subDetail,omitempty"`
}

type ClaimResponseAddItemDetailSubDetail struct {
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
	// Trace number for tracking purposes.
	//
	// May be defined at the jurisdiction level or between trading partners.
	TraceNumber []Identifier `bson:"traceNumber,omitempty" json:"traceNumber,omitempty" xml:"traceNumber,omitempty"`
	// The type of revenue or cost center providing the product and/or service.
	Revenue *CodeableConcept `bson:"revenue,omitempty" json:"revenue,omitempty" xml:"revenue,omitempty"`
	// When the value is a group code then this item collects a set of related item details, otherwise this contains the product, service, drug or other billing code for the item.
	//
	// This element may be the start of a range of .productOrService codes used in conjunction with .productOrServiceEnd or it may be a solo element where .productOrServiceEnd is not used.
	//
	// Comment: If this is an actual service or product line, i.e. not a Group, then use code to indicate the Professional Service or Product supplied (e.g. CTP, HCPCS, USCLS, ICD10, NCPDP, DIN, RxNorm, ACHI, CCI). If a grouping item then use a group code to indicate the type of thing being grouped e.g. 'glasses' or 'compound'.
	ProductOrService *CodeableConcept `bson:"productOrService,omitempty" json:"productOrService,omitempty" xml:"productOrService,omitempty"`
	// This contains the end of a range of product, service, drug or other billing codes for the item.
	//
	// This element is not used when the .productOrService is a group code. This value may only be present when a .productOfService code has been provided to convey the start of the range. Typically this value may be used only with preauthorizations and not with claims.
	ProductOrServiceEnd *CodeableConcept `bson:"productOrServiceEnd,omitempty" json:"productOrServiceEnd,omitempty" xml:"productOrServiceEnd,omitempty"`
	// Item typification or modifiers codes to convey additional context for the product or service.
	//
	// Comment: For example in Oral whether the treatment is cosmetic or associated with TMJ, or for Medical whether the treatment was outside the clinic or outside of office hours.
	Modifier []CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty" xml:"modifier,omitempty"`
	// The number of repetitions of a service or product.
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty" xml:"quantity,omitempty"`
	// If the item is not a group then this is the fee for the product or service, otherwise this is the total of the fees for the details of the group.
	UnitPrice *Money `bson:"unitPrice,omitempty" json:"unitPrice,omitempty" xml:"unitPrice,omitempty"`
	// A real number that represents a multiplier used in determining the overall value of services delivered and/or goods received.
	//
	// The concept of a Factor allows for a discount or surcharge multiplier to be applied to a monetary amount.
	//
	// Comment: To show a 10% senior's discount, the value entered is: 0.90 (1.00 - 0.10).
	Factor *decimal.Decimal `bson:"factor,omitempty" json:"factor,omitempty" xml:"factor,omitempty"`
	// The total of taxes applicable for this product or service.
	Tax *Money `bson:"tax,omitempty" json:"tax,omitempty" xml:"tax,omitempty"`
	// The total amount claimed for the addItem.detail.subDetail.
	//
	// Net = unit price * quantity * factor.
	//
	// Comment: For example, the formula: quantity * unitPrice * factor  = net. Quantity and factor are assumed to be 1 if not supplied.
	Net *Money `bson:"net,omitempty" json:"net,omitempty" xml:"net,omitempty"`
	// The numbers associated with notes below which apply to the adjudication of this item.
	NoteNumber []int `bson:"noteNumber,omitempty" json:"noteNumber,omitempty" xml:"noteNumber,omitempty"`
	// The high-level results of the adjudication if adjudication has been performed.
	ReviewOutcome *ClaimResponseItemReviewOutcome `bson:"reviewOutcome,omitempty" json:"reviewOutcome,omitempty"`
	// The adjudication results.
	Adjudication []ClaimResponseItemAdjudication `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
}

type ClaimResponseTotal struct {
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
	// A code to indicate the information type of this adjudication record.
	//
	// Information types may include: the value submitted, maximum values or percentages allowed or payable under the plan, amounts that the patient is responsible for in aggregate or pertaining to this item, amounts paid by other coverages, and the benefit payable for this item.
	//
	// Comment: For example codes indicating: Co-Pay, deductible, eligible, benefit, tax, etc.
	Category CodeableConcept `bson:"category" json:"category" xml:"category"`
	// Monetary total amount associated with the category.
	Amount Money `bson:"amount" json:"amount" xml:"amount"`
}

type ClaimResponsePayment struct {
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
	// Whether this represents partial or complete payment of the benefits payable.
	Type CodeableConcept `bson:"type" json:"type" xml:"type"`
	// Total amount of all adjustments to this payment included in this transaction which are not related to this claim's adjudication.
	//
	// Comment: Insurers will deduct amounts owing from the provider (adjustment), such as a prior overpayment, from the amount owing to the provider (benefits payable) when payment is made to the provider.
	Adjustment *Money `bson:"adjustment,omitempty" json:"adjustment,omitempty" xml:"adjustment,omitempty"`
	// Reason for the payment adjustment.
	AdjustmentReason *CodeableConcept `bson:"adjustmentReason,omitempty" json:"adjustmentReason,omitempty" xml:"adjustmentReason,omitempty"`
	// Estimated date the payment will be issued or the actual issue date of payment.
	Date *string `bson:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// Benefits payable less any payment adjustment.
	Amount Money `bson:"amount" json:"amount" xml:"amount"`
	// Issuer's unique identifier for the payment instrument.
	//
	// Comment: For example: EFT number or check number.
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
}

type ClaimResponseProcessNote struct {
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
	// A number to uniquely identify a note entry.
	Number *int `bson:"number,omitempty" json:"number,omitempty" xml:"number,omitempty"`
	// The business purpose of the note text.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The explanation or description associated with the processing.
	Text string `bson:"text" json:"text" xml:"text"`
	// A code to define the language used in the text of the note.
	//
	// Comment: Only required if the language is different from the resource language.
	Language *CodeableConcept `bson:"language,omitempty" json:"language,omitempty" xml:"language,omitempty"`
}

type ClaimResponseInsurance struct {
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
	// A number to uniquely identify insurance entries and provide a sequence of coverages to convey coordination of benefit order.
	Sequence int `bson:"sequence" json:"sequence" xml:"sequence"`
	// A flag to indicate that this Coverage is to be used for adjudication of this claim when set to true.
	//
	// Comment: A patient may (will) have multiple insurance policies which provide reimbursement for healthcare services and products. For example a person may also be covered by their spouse's policy and both appear in the list (and may be from the same insurer). This flag will be set to true for only one of the listed policies and that policy will be used for adjudicating this claim. Other claims would be created to request adjudication against the other listed policies.
	Focal bool `bson:"focal" json:"focal" xml:"focal"`
	// Reference to the insurance card level information contained in the Coverage resource.
	//
	// The coverage issuing insurer will use these details to locate the patient's actual coverage within the insurer's information system.
	Coverage Reference `bson:"coverage" json:"coverage" xml:"coverage"`
	// A business agreement number established between the provider and the insurer for special business processing purposes.
	BusinessArrangement *string `bson:"businessArrangement,omitempty" json:"businessArrangement,omitempty" xml:"businessArrangement,omitempty"`
	// The result of the adjudication of the line items for the Coverage specified in this insurance.
	//
	// Comment: Must not be specified when 'focal=true' for this insurance.
	ClaimResponse *Reference `bson:"claimResponse,omitempty" json:"claimResponse,omitempty" xml:"claimResponse,omitempty"`
}

type ClaimResponseError struct {
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
	// The sequence number of the line item submitted which contains the error.
	//
	// This value is omitted when the error occurs outside of the item structure.
	ItemSequence *int `bson:"itemSequence,omitempty" json:"itemSequence,omitempty" xml:"itemSequence,omitempty"`
	// The sequence number of the detail within the line item submitted which contains the error.
	//
	// This value is omitted when the error occurs outside of the item structure.
	DetailSequence *int `bson:"detailSequence,omitempty" json:"detailSequence,omitempty" xml:"detailSequence,omitempty"`
	// The sequence number of the sub-detail within the detail within the line item submitted which contains the error.
	//
	// This value is omitted when the error occurs outside of the item structure.
	SubDetailSequence *int `bson:"subDetailSequence,omitempty" json:"subDetailSequence,omitempty" xml:"subDetailSequence,omitempty"`
	// An error code, from a specified code system, which details why the claim could not be adjudicated.
	Code CodeableConcept `bson:"code" json:"code" xml:"code"`
	// A [simple subset of FHIRPath](fhirpath.html#simple) limited to element names, repetition indicators and the default child accessor that identifies one of the elements in the resource that caused this issue to be raised.
	//
	// Comment: The root of the FHIRPath is the resource or bundle that generated OperationOutcome.  Each FHIRPath SHALL resolve to a single node.
	Expression []string `bson:"expression,omitempty" json:"expression,omitempty" xml:"expression,omitempty"`
}

func (r ClaimResponse) ResourceType() string {
	return "ClaimResponse"
}

// MarshalJSON implements json.Marshaler
func (r ClaimResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ClaimResponse
		ResourceType string `json:"resourceType"`
	}{
		ClaimResponse: r,
		ResourceType:  r.ResourceType(),
	})
}
