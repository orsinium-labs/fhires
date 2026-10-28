package fhires

import (
	"encoding/json"

	decimal "github.com/shopspring/decimal"
)

// A provider issued list of professional services and products which have been provided, or are to be provided, to a patient which is sent to an insurer for reimbursement.
//
// Purpose: The Claim resource is used by providers to exchange services and products rendered to patients or planned to be rendered with insurers for reimbuserment. It is also used by insurers to exchange claims information with statutory reporting and data analytics firms.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Claim
type Claim struct {
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
	// A unique identifier assigned to this claim.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Trace number for tracking purposes.
	//
	// May be defined at the jurisdiction level or between trading partners.
	TraceNumber []Identifier `bson:"traceNumber,omitempty" json:"traceNumber,omitempty" xml:"traceNumber,omitempty"`
	// The status of the resource instance.
	//
	// Comment: This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid.
	Status FinancialResourceStatusCodes `bson:"status" json:"status" xml:"status"`
	// The category of claim, e.g. oral, pharmacy, vision, institutional, professional.
	//
	// Comment: The code system provides oral, pharmacy, vision, professional and institutional claim types. Those supported depends on the requirements of the jurisdiction. The valueset is extensible to accommodate other types of claims as required by the jurisdiction.
	Type CodeableConcept `bson:"type" json:"type" xml:"type"`
	// A finer grained suite of claim type codes which may convey additional information such as Inpatient vs Outpatient and/or a specialty service.
	//
	// Comment: This may contain the local bill type codes, for example the US UB-04 bill type code or the CMS bill type.
	SubType *CodeableConcept `bson:"subType,omitempty" json:"subType,omitempty" xml:"subType,omitempty"`
	// A code to indicate whether the nature of the request is: Claim - A request to an Insurer to adjudicate the supplied charges for health care goods and services under the identified policy and to pay the determined Benefit amount, if any; Preauthorization - A request to an Insurer to adjudicate the supplied proposed future charges for health care goods and services under the identified policy and to approve the services and provide the expected benefit amounts and potentially to reserve funds to pay the benefits when Claims for the indicated services are later submitted; or, Pre-determination - A request to an Insurer to adjudicate the supplied 'what if' charges for health care goods and services under the identified policy and report back what the Benefit payable would be had the services actually been provided.
	Use Use `bson:"use" json:"use" xml:"use"`
	// The party to whom the professional services and/or products have been supplied or are being considered and for whom actual or forecast reimbursement is sought.
	Patient Reference `bson:"patient" json:"patient" xml:"patient"`
	// The period for which charges are being submitted.
	//
	// Comment: Typically this would be today or in the past for a claim, and today or in the future for preauthorizations and predeterminations. Typically line item dates of service should fall within the billing period if one is specified.
	BillablePeriod *Period `bson:"billablePeriod,omitempty" json:"billablePeriod,omitempty" xml:"billablePeriod,omitempty"`
	// The date this resource was created.
	//
	// Comment: This field is independent of the date of creation of the resource as it may reflect the creation date of a source document prior to digitization. Typically for claims all services must be completed as of this date.
	Created string `bson:"created" json:"created" xml:"created"`
	// Individual who created the claim, predetermination or preauthorization.
	Enterer *Reference `bson:"enterer,omitempty" json:"enterer,omitempty" xml:"enterer,omitempty"`
	// The Insurer who is target of the request.
	Insurer *Reference `bson:"insurer,omitempty" json:"insurer,omitempty" xml:"insurer,omitempty"`
	// The provider which is responsible for the claim, predetermination or preauthorization.
	//
	// Comment: Typically this field would be 1..1 where this party is accountable for the data content within the claim but is not necessarily the facility, provider group or practitioner who provided the products and services listed within this claim resource. This field is the Billing Provider, for example, a facility, provider group, lab or practitioner.
	Provider *Reference `bson:"provider,omitempty" json:"provider,omitempty" xml:"provider,omitempty"`
	// The provider-required urgency of processing the request.
	//
	// Typical values include: stat, normal, deferred.
	//
	// Comment: If a claim processor is unable to complete the processing as per the priority then they should generate an error and not process the request.
	Priority *CodeableConcept `bson:"priority,omitempty" json:"priority,omitempty" xml:"priority,omitempty"`
	// A code to indicate whether and for whom funds are to be reserved for future claims.
	//
	// Comment: This field is only used for preauthorizations.
	FundsReserve *CodeableConcept `bson:"fundsReserve,omitempty" json:"fundsReserve,omitempty" xml:"fundsReserve,omitempty"`
	// Other claims which are related to this claim such as prior submissions or claims for related services or for the same event.
	//
	// Comment: For example,  for the original treatment and follow-up exams.
	Related []ClaimRelated `bson:"related,omitempty" json:"related,omitempty" xml:"related,omitempty"`
	// Prescription is the document/authorization given to the claim author for them to provide products and services for which consideration (reimbursement) is sought.
	//
	// Could be a RX for medications, an 'order' for oxygen or wheelchair or physiotherapy treatments.
	Prescription *Reference `bson:"prescription,omitempty" json:"prescription,omitempty" xml:"prescription,omitempty"`
	// Original prescription which has been superseded by this prescription to support the dispensing of pharmacy services, medications or products.
	//
	// Comment: For example, a physician may prescribe a medication which the pharmacy determines is contraindicated, or for which the patient has an intolerance, and therefore issues a new prescription for an alternate medication which has the same therapeutic intent. The prescription from the pharmacy becomes the 'prescription' and that from the physician becomes the 'original prescription'.
	OriginalPrescription *Reference `bson:"originalPrescription,omitempty" json:"originalPrescription,omitempty" xml:"originalPrescription,omitempty"`
	// The party to be reimbursed for cost of the products and services according to the terms of the policy.
	//
	// Comment: Often providers agree to receive the benefits payable to reduce the near-term costs to the patient. The insurer may decline to pay the provider and choose to pay the subscriber instead.
	Payee *ClaimPayee `bson:"payee,omitempty" json:"payee,omitempty" xml:"payee,omitempty"`
	// The referral information received by the claim author, it is not to be used when the author generates a referral for a patient.
	//
	// A copy of that referral may be provided as supporting information. Some insurers require proof of referral to pay for services or to pay specialist rates for services.
	//
	// Comment: The referral resource which lists the date, practitioner, reason and other supporting information.
	Referral *Reference `bson:"referral,omitempty" json:"referral,omitempty" xml:"referral,omitempty"`
	// Healthcare encounters related to this claim.
	//
	// Comment: This will typically be the encounter the event occurred within, but some activities may be initiated prior to or after the official completion of an encounter but still be tied to the context of the encounter.
	Encounter []Reference `bson:"encounter,omitempty" json:"encounter,omitempty" xml:"encounter,omitempty"`
	// Facility where the services were provided.
	Facility *Reference `bson:"facility,omitempty" json:"facility,omitempty" xml:"facility,omitempty"`
	// A package billing code or bundle code used to group products and services to a particular health condition (such as heart attack) which is based on a predetermined grouping code system.
	//
	// Comment: For example DRG (Diagnosis Related Group) or a bundled billing code. A patient may have a diagnosis of a Myocardial Infarction and a DRG for HeartAttack would be assigned. The Claim item (and possible subsequent claims) would refer to the DRG for those line items that were for services related to the heart attack event.
	DiagnosisRelatedGroup *CodeableConcept `bson:"diagnosisRelatedGroup,omitempty" json:"diagnosisRelatedGroup,omitempty" xml:"diagnosisRelatedGroup,omitempty"`
	// Information code for an event with a corresponding date or period.
	Event []ClaimEvent `bson:"event,omitempty" json:"event,omitempty" xml:"event,omitempty"`
	// The members of the team who provided the products and services.
	CareTeam []ClaimCareTeam `bson:"careTeam,omitempty" json:"careTeam,omitempty" xml:"careTeam,omitempty"`
	// Additional information codes regarding exceptions, special considerations, the condition, situation, prior or concurrent issues.
	//
	// Comment: Often there are multiple jurisdiction specific valuesets which are required.
	SupportingInfo []ClaimSupportingInfo `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty" xml:"supportingInfo,omitempty"`
	// Information about diagnoses relevant to the claim items.
	Diagnosis []ClaimDiagnosis `bson:"diagnosis,omitempty" json:"diagnosis,omitempty" xml:"diagnosis,omitempty"`
	// Procedures performed on the patient relevant to the billing items with the claim.
	Procedure []ClaimProcedure `bson:"procedure,omitempty" json:"procedure,omitempty" xml:"procedure,omitempty"`
	// Financial instruments for reimbursement for the health care products and services specified on the claim.
	//
	// Comment: All insurance coverages for the patient which may be applicable for reimbursement, of the products and services listed in the claim, are typically provided in the claim to allow insurers to confirm the ordering of the insurance coverages relative to local 'coordination of benefit' rules. One coverage (and only one) with 'focal=true' is to be used in the adjudication of this claim. Coverages appearing before the focal Coverage in the list, and where 'Coverage.subrogation=false', should provide a reference to the ClaimResponse containing the adjudication results of the prior claim.
	Insurance []ClaimInsurance `bson:"insurance,omitempty" json:"insurance,omitempty" xml:"insurance,omitempty"`
	// Details of an accident which resulted in injuries which required the products and services listed in the claim.
	Accident *ClaimAccident `bson:"accident,omitempty" json:"accident,omitempty" xml:"accident,omitempty"`
	// The amount paid by the patient, in total at the claim claim level or specifically for the item and detail level, to the provider for goods and services.
	PatientPaid *Money `bson:"patientPaid,omitempty" json:"patientPaid,omitempty" xml:"patientPaid,omitempty"`
	// A claim line.
	//
	// Either a simple  product or service or a 'group' of details which can each be a simple items or groups of sub-details.
	Item []ClaimItem `bson:"item,omitempty" json:"item,omitempty" xml:"item,omitempty"`
	// The total value of the all the items in the claim.
	Total *Money `bson:"total,omitempty" json:"total,omitempty" xml:"total,omitempty"`
}

type ClaimRelated struct {
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
	// Reference to a related claim.
	Claim *Reference `bson:"claim,omitempty" json:"claim,omitempty" xml:"claim,omitempty"`
	// A code to convey how the claims are related.
	//
	// Comment: For example, prior claim or umbrella.
	Relationship *CodeableConcept `bson:"relationship,omitempty" json:"relationship,omitempty" xml:"relationship,omitempty"`
	// An alternate organizational reference to the case or file to which this particular claim pertains.
	//
	// Comment: For example, Property/Casualty insurer claim # or Workers Compensation case # .
	Reference *Identifier `bson:"reference,omitempty" json:"reference,omitempty" xml:"reference,omitempty"`
}

type ClaimPayee struct {
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
	// Type of Party to be reimbursed: subscriber, provider, other.
	Type CodeableConcept `bson:"type" json:"type" xml:"type"`
	// Reference to the individual or organization to whom any payment will be made.
	//
	// Comment: Not required if the payee is 'subscriber' or 'provider'.
	Party *Reference `bson:"party,omitempty" json:"party,omitempty" xml:"party,omitempty"`
}

type ClaimEvent struct {
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

type ClaimCareTeam struct {
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
	// A number to uniquely identify care team entries.
	Sequence int `bson:"sequence" json:"sequence" xml:"sequence"`
	// Member of the team who provided the product or service.
	Provider Reference `bson:"provider" json:"provider" xml:"provider"`
	// The party who is billing and/or responsible for the claimed products or services.
	//
	// Comment: Responsible might not be required when there is only a single provider listed.
	Responsible *bool `bson:"responsible,omitempty" json:"responsible,omitempty" xml:"responsible,omitempty"`
	// The lead, assisting or supervising practitioner and their discipline if a multidisciplinary team.
	//
	// Comment: Role might not be required when there is only a single provider listed.
	Role *CodeableConcept `bson:"role,omitempty" json:"role,omitempty" xml:"role,omitempty"`
	// The specialization of the practitioner or provider which is applicable for this service.
	Specialty *CodeableConcept `bson:"specialty,omitempty" json:"specialty,omitempty" xml:"specialty,omitempty"`
}

type ClaimSupportingInfo struct {
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
	// A number to uniquely identify supporting information entries.
	Sequence int `bson:"sequence" json:"sequence" xml:"sequence"`
	// The general class of the information supplied: information; exception; accident, employment; onset, etc.
	//
	// Comment: This may contain a category for the local bill type codes.
	Category CodeableConcept `bson:"category" json:"category" xml:"category"`
	// System and code pertaining to the specific information regarding special conditions relating to the setting, treatment or patient  for which care is sought.
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// The date when or period to which this information refers.
	TimingDate   *string `bson:"timingDate,omitempty" json:"timingDate,omitempty" xml:"timingDate,omitempty"`
	TimingPeriod *Period `bson:"timingPeriod,omitempty" json:"timingPeriod,omitempty" xml:"timingPeriod,omitempty"`
	// Additional data or information such as resources, documents, images etc.
	//
	// including references to the data or the actual inclusion of the data.
	//
	// Comment: Could be used to provide references to other resources, document. For example could contain a PDF in an Attachment of the Police Report for an Accident.
	ValueBoolean    *bool       `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty" xml:"valueBoolean,omitempty"`
	ValueString     *string     `bson:"valueString,omitempty" json:"valueString,omitempty" xml:"valueString,omitempty"`
	ValueQuantity   *Quantity   `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty" xml:"valueQuantity,omitempty"`
	ValueAttachment *Attachment `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty" xml:"valueAttachment,omitempty"`
	ValueReference  *Reference  `bson:"valueReference,omitempty" json:"valueReference,omitempty" xml:"valueReference,omitempty"`
	ValueIdentifier *Identifier `bson:"valueIdentifier,omitempty" json:"valueIdentifier,omitempty" xml:"valueIdentifier,omitempty"`
	// Provides the reason in the situation where a reason code is required in addition to the content.
	//
	// Comment: For example: the reason for the additional stay, or why a tooth is  missing.
	Reason *CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty" xml:"reason,omitempty"`
}

type ClaimDiagnosis struct {
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
	// A number to uniquely identify diagnosis entries.
	//
	// Comment: Diagnosis are presented in list order to their expected importance: primary, secondary, etc.
	Sequence int `bson:"sequence" json:"sequence" xml:"sequence"`
	// The nature of illness or problem in a coded form or as a reference to an external defined Condition.
	DiagnosisCodeableConcept CodeableConcept `bson:"diagnosisCodeableConcept" json:"diagnosisCodeableConcept" xml:"diagnosisCodeableConcept"`
	DiagnosisReference       Reference       `bson:"diagnosisReference" json:"diagnosisReference" xml:"diagnosisReference"`
	// When the condition was observed or the relative ranking.
	//
	// Comment: For example: admitting, primary, secondary, discharge.
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Indication of whether the diagnosis was present on admission to a facility.
	OnAdmission *CodeableConcept `bson:"onAdmission,omitempty" json:"onAdmission,omitempty" xml:"onAdmission,omitempty"`
}

type ClaimProcedure struct {
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
	// A number to uniquely identify procedure entries.
	Sequence int `bson:"sequence" json:"sequence" xml:"sequence"`
	// When the condition was observed or the relative ranking.
	//
	// Comment: For example: primary, secondary.
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Date and optionally time the procedure was performed.
	Date *string `bson:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// The code or reference to a Procedure resource which identifies the clinical intervention performed.
	ProcedureCodeableConcept CodeableConcept `bson:"procedureCodeableConcept" json:"procedureCodeableConcept" xml:"procedureCodeableConcept"`
	ProcedureReference       Reference       `bson:"procedureReference" json:"procedureReference" xml:"procedureReference"`
	// Unique Device Identifiers associated with this line item.
	Udi []Reference `bson:"udi,omitempty" json:"udi,omitempty" xml:"udi,omitempty"`
}

type ClaimInsurance struct {
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
	// The business identifier to be used when the claim is sent for adjudication against this insurance policy.
	//
	// Comment: Only required in jurisdictions where insurers, rather than the provider, are required to send claims to  insurers that appear after them in the list. This element is not required when 'subrogation=true'.
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Reference to the insurance card level information contained in the Coverage resource.
	//
	// The coverage issuing insurer will use these details to locate the patient's actual coverage within the insurer's information system.
	Coverage Reference `bson:"coverage" json:"coverage" xml:"coverage"`
	// A business agreement number established between the provider and the insurer for special business processing purposes.
	BusinessArrangement *string `bson:"businessArrangement,omitempty" json:"businessArrangement,omitempty" xml:"businessArrangement,omitempty"`
	// Reference numbers previously provided by the insurer to the provider to be quoted on subsequent claims containing services or products related to the prior authorization.
	//
	// Comment: This value is an alphanumeric string that may be provided over the phone, via text, via paper, or within a ClaimResponse resource and is not a FHIR Identifier.
	PreAuthRef []string `bson:"preAuthRef,omitempty" json:"preAuthRef,omitempty" xml:"preAuthRef,omitempty"`
	// The result of the adjudication of the line items for the Coverage specified in this insurance.
	//
	// Comment: Must not be specified when 'focal=true' for this insurance.
	ClaimResponse *Reference `bson:"claimResponse,omitempty" json:"claimResponse,omitempty" xml:"claimResponse,omitempty"`
}

type ClaimAccident struct {
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
	// Date of an accident event  related to the products and services contained in the claim.
	//
	// Comment: The date of the accident has to precede the dates of the products and services but within a reasonable timeframe.
	Date string `bson:"date" json:"date" xml:"date"`
	// The type or context of the accident event for the purposes of selection of potential insurance coverages and determination of coordination between insurers.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The physical location of the accident event.
	LocationAddress   *Address   `bson:"locationAddress,omitempty" json:"locationAddress,omitempty" xml:"locationAddress,omitempty"`
	LocationReference *Reference `bson:"locationReference,omitempty" json:"locationReference,omitempty" xml:"locationReference,omitempty"`
}

type ClaimItem struct {
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
	// A number to uniquely identify item entries.
	Sequence int `bson:"sequence" json:"sequence" xml:"sequence"`
	// Trace number for tracking purposes.
	//
	// May be defined at the jurisdiction level or between trading partners.
	TraceNumber []Identifier `bson:"traceNumber,omitempty" json:"traceNumber,omitempty" xml:"traceNumber,omitempty"`
	// CareTeam members related to this service or product.
	CareTeamSequence []int `bson:"careTeamSequence,omitempty" json:"careTeamSequence,omitempty" xml:"careTeamSequence,omitempty"`
	// Diagnosis applicable for this service or product.
	DiagnosisSequence []int `bson:"diagnosisSequence,omitempty" json:"diagnosisSequence,omitempty" xml:"diagnosisSequence,omitempty"`
	// Procedures applicable for this service or product.
	ProcedureSequence []int `bson:"procedureSequence,omitempty" json:"procedureSequence,omitempty" xml:"procedureSequence,omitempty"`
	// Exceptions, special conditions and supporting information applicable for this service or product.
	InformationSequence []int `bson:"informationSequence,omitempty" json:"informationSequence,omitempty" xml:"informationSequence,omitempty"`
	// The type of revenue or cost center providing the product and/or service.
	Revenue *CodeableConcept `bson:"revenue,omitempty" json:"revenue,omitempty" xml:"revenue,omitempty"`
	// Code to identify the general type of benefits under which products and services are provided.
	//
	// Comment: Examples include Medical Care, Periodontics, Renal Dialysis, Vision Coverage.
	Category *CodeableConcept `bson:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
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
	// The amount paid by the patient, in total at the claim claim level or specifically for the item and detail level, to the provider for goods and services.
	PatientPaid *Money `bson:"patientPaid,omitempty" json:"patientPaid,omitempty" xml:"patientPaid,omitempty"`
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
	// The total amount claimed for the group (if a grouper) or the line item.
	//
	// Net = unit price * quantity * factor.
	//
	// Comment: For example, the formula: quantity * unitPrice * factor  = net. Quantity and factor are assumed to be 1 if not supplied.
	Net *Money `bson:"net,omitempty" json:"net,omitempty" xml:"net,omitempty"`
	// Unique Device Identifiers associated with this line item.
	Udi []Reference `bson:"udi,omitempty" json:"udi,omitempty" xml:"udi,omitempty"`
	// Physical location where the service is performed or applies.
	BodySite []ClaimItemBodySite `bson:"bodySite,omitempty" json:"bodySite,omitempty" xml:"bodySite,omitempty"`
	// Healthcare encounters related to this claim.
	//
	// Comment: This will typically be the encounter the event occurred within, but some activities may be initiated prior to or after the official completion of an encounter but still be tied to the context of the encounter.
	Encounter []Reference `bson:"encounter,omitempty" json:"encounter,omitempty" xml:"encounter,omitempty"`
	// A claim detail line.
	//
	// Either a simple (a product or service) or a 'group' of sub-details which are simple items.
	Detail []ClaimItemDetail `bson:"detail,omitempty" json:"detail,omitempty" xml:"detail,omitempty"`
}

type ClaimItemBodySite struct {
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
	// Comment: For example: Providing a tooth code, allows an insurer to identify a provider performing a filling on a tooth that was previously removed.
	Site []CodeableReference `bson:"site" json:"site" xml:"site"`
	// A region or surface of the bodySite, e.g. limb region or tooth surface(s).
	SubSite []CodeableConcept `bson:"subSite,omitempty" json:"subSite,omitempty" xml:"subSite,omitempty"`
}

type ClaimItemDetail struct {
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
	// A number to uniquely identify item entries.
	Sequence int `bson:"sequence" json:"sequence" xml:"sequence"`
	// Trace number for tracking purposes.
	//
	// May be defined at the jurisdiction level or between trading partners.
	TraceNumber []Identifier `bson:"traceNumber,omitempty" json:"traceNumber,omitempty" xml:"traceNumber,omitempty"`
	// The type of revenue or cost center providing the product and/or service.
	Revenue *CodeableConcept `bson:"revenue,omitempty" json:"revenue,omitempty" xml:"revenue,omitempty"`
	// Code to identify the general type of benefits under which products and services are provided.
	//
	// Comment: Examples include Medical Care, Periodontics, Renal Dialysis, Vision Coverage.
	Category *CodeableConcept `bson:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
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
	// Comment: For example in Oral whether the treatment is cosmetic or associated with TMJ, or for Medical whether the treatment was outside the clinic or out of office hours.
	Modifier []CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty" xml:"modifier,omitempty"`
	// Identifies the program under which this may be recovered.
	//
	// Comment: For example: Neonatal program, child dental program or drug users recovery program.
	ProgramCode []CodeableConcept `bson:"programCode,omitempty" json:"programCode,omitempty" xml:"programCode,omitempty"`
	// The amount paid by the patient, in total at the claim claim level or specifically for the item and detail level, to the provider for goods and services.
	PatientPaid *Money `bson:"patientPaid,omitempty" json:"patientPaid,omitempty" xml:"patientPaid,omitempty"`
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
	// The total amount claimed for the group (if a grouper) or the line item.detail.
	//
	// Net = unit price * quantity * factor.
	//
	// Comment: For example, the formula: quantity * unitPrice * factor  = net. Quantity and factor are assumed to be 1 if not supplied.
	Net *Money `bson:"net,omitempty" json:"net,omitempty" xml:"net,omitempty"`
	// Unique Device Identifiers associated with this line item.
	Udi []Reference `bson:"udi,omitempty" json:"udi,omitempty" xml:"udi,omitempty"`
	// A claim detail line.
	//
	// Either a simple (a product or service) or a 'group' of sub-details which are simple items.
	SubDetail []ClaimItemDetailSubDetail `bson:"subDetail,omitempty" json:"subDetail,omitempty" xml:"subDetail,omitempty"`
}

type ClaimItemDetailSubDetail struct {
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
	// A number to uniquely identify item entries.
	Sequence int `bson:"sequence" json:"sequence" xml:"sequence"`
	// Trace number for tracking purposes.
	//
	// May be defined at the jurisdiction level or between trading partners.
	TraceNumber []Identifier `bson:"traceNumber,omitempty" json:"traceNumber,omitempty" xml:"traceNumber,omitempty"`
	// The type of revenue or cost center providing the product and/or service.
	Revenue *CodeableConcept `bson:"revenue,omitempty" json:"revenue,omitempty" xml:"revenue,omitempty"`
	// Code to identify the general type of benefits under which products and services are provided.
	//
	// Comment: Examples include Medical Care, Periodontics, Renal Dialysis, Vision Coverage.
	Category *CodeableConcept `bson:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
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
	// Comment: For example in Oral whether the treatment is cosmetic or associated with TMJ, or for Medical whether the treatment was outside the clinic or out of office hours.
	Modifier []CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty" xml:"modifier,omitempty"`
	// Identifies the program under which this may be recovered.
	//
	// Comment: For example: Neonatal program, child dental program or drug users recovery program.
	ProgramCode []CodeableConcept `bson:"programCode,omitempty" json:"programCode,omitempty" xml:"programCode,omitempty"`
	// The amount paid by the patient, in total at the claim claim level or specifically for the item and detail level, to the provider for goods and services.
	PatientPaid *Money `bson:"patientPaid,omitempty" json:"patientPaid,omitempty" xml:"patientPaid,omitempty"`
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
	// The total amount claimed for line item.detail.subDetail.
	//
	// Net = unit price * quantity * factor.
	//
	// Comment: For example, the formula: quantity * unitPrice * factor  = net. Quantity and factor are assumed to be 1 if not supplied.
	Net *Money `bson:"net,omitempty" json:"net,omitempty" xml:"net,omitempty"`
	// Unique Device Identifiers associated with this line item.
	Udi []Reference `bson:"udi,omitempty" json:"udi,omitempty" xml:"udi,omitempty"`
}

func (r Claim) ResourceType() string {
	return "Claim"
}

// MarshalJSON implements json.Marshaler
func (r Claim) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Claim
		ResourceType string `json:"resourceType"`
	}{
		Claim:        r,
		ResourceType: r.ResourceType(),
	})
}
