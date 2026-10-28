package fhires

import "encoding/json/v2"

// Financial instrument which may be used to reimburse or pay for health care products and services.
//
// Includes both insurance and self-payment.
//
// Purpose: Coverage provides a link between covered parties (patients) and the payors of their healthcare costs (both insurance and self-pay).
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Coverage
type Coverage struct {
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
	// The identifier of the coverage as issued by the insurer.
	//
	// Comment: The main (and possibly only) identifier for the coverage - often referred to as a Member Id, Certificate number, Personal Health Number or Case ID. May be constructed as the concatenation of the Coverage.SubscriberID and the Coverage.dependant. Note that not all insurers issue unique member IDs therefore searches may result in multiple responses.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The status of the resource instance.
	//
	// Comment: This element is labeled as a modifier because the status contains the code entered-in-error that marks the coverage as not currently valid.
	Status FinancialResourceStatusCodes `bson:"status" json:"status" xml:"status"`
	// The nature of the coverage be it insurance, or cash payment such as self-pay.
	Kind Kind `bson:"kind" json:"kind" xml:"kind"`
	// Link to the paying party and optionally what specifically they will be responsible to pay.
	PaymentBy []CoveragePaymentBy `bson:"paymentBy,omitempty" json:"paymentBy,omitempty" xml:"paymentBy,omitempty"`
	// The type of coverage: social program, medical plan, accident coverage (workers compensation, auto), group health or payment by an individual or organization.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The party who 'owns' the insurance policy.
	//
	// Comment: For example: may be an individual, corporation or the subscriber's employer.
	PolicyHolder *Reference `bson:"policyHolder,omitempty" json:"policyHolder,omitempty" xml:"policyHolder,omitempty"`
	// The party who has signed-up for or 'owns' the contractual relationship to the policy or to whom the benefit of the policy for services rendered to them or their family is due.
	//
	// Comment: May be self or a parent in the case of dependants. A subscriber is only required on certain types of policies not all policies and that it is appropriate to have just a policyholder and a beneficiary when not other party can join that policy instance.
	Subscriber *Reference `bson:"subscriber,omitempty" json:"subscriber,omitempty" xml:"subscriber,omitempty"`
	// The insurer assigned ID for the Subscriber.
	SubscriberID []Identifier `bson:"subscriberId,omitempty" json:"subscriberId,omitempty" xml:"subscriberId,omitempty"`
	// The party who benefits from the insurance coverage; the patient when products and/or services are provided.
	Beneficiary Reference `bson:"beneficiary" json:"beneficiary" xml:"beneficiary"`
	// A designator for a dependent under the coverage.
	//
	// Comment: Sometimes the member number is constructed from the subscriberId and the dependant number.
	Dependent *string `bson:"dependent,omitempty" json:"dependent,omitempty" xml:"dependent,omitempty"`
	// The relationship of beneficiary (patient) to the subscriber.
	//
	// Comment: Typically, an individual uses policies which are theirs (relationship='self') before policies owned by others.
	Relationship *CodeableConcept `bson:"relationship,omitempty" json:"relationship,omitempty" xml:"relationship,omitempty"`
	// Time period during which the coverage is in force.
	//
	// A missing start date indicates the start date isn't known, a missing end date means the coverage is continuing to be in force.
	Period *Period `bson:"period,omitempty" json:"period,omitempty" xml:"period,omitempty"`
	// The program or plan underwriter, payor, insurance company.
	//
	// Comment: May provide multiple identifiers such as insurance company identifier or business identifier (BIN number).
	Insurer *Reference `bson:"insurer,omitempty" json:"insurer,omitempty" xml:"insurer,omitempty"`
	// A suite of underwriter specific classifiers.
	//
	// Comment: For example, class may be used to identify a class of coverage or employer group, policy, or plan.
	Class []CoverageClass `bson:"class,omitempty" json:"class,omitempty" xml:"class,omitempty"`
	// The order of applicability of this coverage relative to other coverages which are currently in force.
	//
	// Note, there may be gaps in the numbering and this does not imply primary, secondary etc. as the specific positioning of coverages depends upon the episode of care. For example; a patient might have (0) auto insurance (1) their own health insurance and (2) spouse's health insurance. When claiming for treatments which were not the result of an auto accident then only coverages (1) and (2) above would be applicatble and would apply in the order specified in parenthesis.
	Order *int `bson:"order,omitempty" json:"order,omitempty" xml:"order,omitempty"`
	// The insurer-specific identifier for the insurer-defined network of providers to which the beneficiary may seek treatment which will be covered at the 'in-network' rate, otherwise 'out of network' terms and conditions apply.
	Network *string `bson:"network,omitempty" json:"network,omitempty" xml:"network,omitempty"`
	// A suite of codes indicating the cost category and associated amount which have been detailed in the policy and may have been  included on the health card.
	//
	// Comment: For example by knowing the patient visit co-pay, the provider can collect the amount prior to undertaking treatment.
	CostToBeneficiary []CoverageCostToBeneficiary `bson:"costToBeneficiary,omitempty" json:"costToBeneficiary,omitempty" xml:"costToBeneficiary,omitempty"`
	// When 'subrogation=true' this insurance instance has been included not for adjudication but to provide insurers with the details to recover costs.
	//
	// Comment: Typically, automotive and worker's compensation policies would be flagged with 'subrogation=true' to enable healthcare payors to collect against accident claims.
	Subrogation *bool `bson:"subrogation,omitempty" json:"subrogation,omitempty" xml:"subrogation,omitempty"`
	// The policy(s) which constitute this insurance coverage.
	Contract []Reference `bson:"contract,omitempty" json:"contract,omitempty" xml:"contract,omitempty"`
	// The insurance plan details, benefits and costs, which constitute this insurance coverage.
	InsurancePlan *Reference `bson:"insurancePlan,omitempty" json:"insurancePlan,omitempty" xml:"insurancePlan,omitempty"`
}

type CoveragePaymentBy struct {
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
	// The list of parties providing non-insurance payment for the treatment costs.
	Party Reference `bson:"party" json:"party" xml:"party"`
	//  Description of the financial responsibility.
	Responsibility *string `bson:"responsibility,omitempty" json:"responsibility,omitempty" xml:"responsibility,omitempty"`
}

type CoverageClass struct {
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
	// The type of classification for which an insurer-specific class label or number and optional name is provided.
	//
	//  For example, type may be used to identify a class of coverage or employer group, policy, or plan.
	Type CodeableConcept `bson:"type" json:"type" xml:"type"`
	// The alphanumeric identifier associated with the insurer issued label.
	//
	// Comment: For example, the Group or Plan number.
	Value Identifier `bson:"value" json:"value" xml:"value"`
	// A short description for the class.
	Name *string `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

type CoverageCostToBeneficiary struct {
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
	// The category of patient centric costs associated with treatment.
	//
	// Comment: For example visit, specialist visits, emergency, inpatient care, etc.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Code to identify the general type of benefits under which products and services are provided.
	//
	// Comment: Examples include Medical Care, Periodontics, Renal Dialysis, Vision Coverage.
	Category *CodeableConcept `bson:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
	// Is a flag to indicate whether the benefits refer to in-network providers or out-of-network providers.
	Network *CodeableConcept `bson:"network,omitempty" json:"network,omitempty" xml:"network,omitempty"`
	// Indicates if the benefits apply to an individual or to the family.
	Unit *CodeableConcept `bson:"unit,omitempty" json:"unit,omitempty" xml:"unit,omitempty"`
	// The term or period of the values such as 'maximum lifetime benefit' or 'maximum annual visits'.
	Term *CodeableConcept `bson:"term,omitempty" json:"term,omitempty" xml:"term,omitempty"`
	// The amount due from the patient for the cost category.
	//
	// Comment: Amount may be expressed as a percentage of the service/product cost or a fixed amount of currency.
	ValueQuantity *Quantity `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty" xml:"valueQuantity,omitempty"`
	ValueMoney    *Money    `bson:"valueMoney,omitempty" json:"valueMoney,omitempty" xml:"valueMoney,omitempty"`
	// A suite of codes indicating exceptions or reductions to patient costs and their effective periods.
	Exception []CoverageCostToBeneficiaryException `bson:"exception,omitempty" json:"exception,omitempty" xml:"exception,omitempty"`
}

type CoverageCostToBeneficiaryException struct {
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
	// The code for the specific exception.
	Type CodeableConcept `bson:"type" json:"type" xml:"type"`
	// The timeframe the exception is in force.
	Period *Period `bson:"period,omitempty" json:"period,omitempty" xml:"period,omitempty"`
}

func (r Coverage) ResourceType() string {
	return "Coverage"
}

// MarshalJSON implements json.Marshaler
func (r Coverage) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Coverage
		ResourceType string `json:"resourceType"`
	}{
		Coverage:     r,
		ResourceType: r.ResourceType(),
	})
}
