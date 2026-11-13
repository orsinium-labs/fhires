package fhires

import "encoding/xml"

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
	DomainResource
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

// MarshalJSON implements [json.Marshaler]
func (r Coverage) MarshalJSON() ([]byte, error) {
	type Wrapper Coverage
	return marshalJSON(r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Marshaler].
func (r Coverage) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Wrapper Coverage
	return marshalXML(e, start, r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Unmarshaler].
func (r *Coverage) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper Coverage
	w := Wrapper(*r)
	err := unmarshalXML(d, start, r.ResourceType(), &w)
	*r = Coverage(w)
	return err
}
