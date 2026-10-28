package fhires

import "encoding/json/v2"

// A financial tool for tracking value accrued for a particular purpose.
//
//	In the healthcare field, used to track charges for a patient, cost centers, etc.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Account
type Account struct {
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
	// Unique identifier used to reference the account.
	//
	//  Might or might not be intended for human use (e.g. credit card number).
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Indicates whether the account is presently used/usable or not.
	//
	// Comment: This element is labeled as a modifier because the status contains the codes inactive and entered-in-error that mark the Account as not currently valid.
	Status AccountStatus `bson:"status" json:"status" xml:"status"`
	// The BillingStatus tracks the lifecycle of the account through the billing process.
	//
	// It indicates how transactions are treated when they are allocated to the account.
	BillingStatus *CodeableConcept `bson:"billingStatus,omitempty" json:"billingStatus,omitempty" xml:"billingStatus,omitempty"`
	// Categorizes the account for reporting and searching purposes.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Name used for the account when displaying it to humans in reports, etc.
	Name *string `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Identifies the entity which incurs the expenses.
	//
	// While the immediate recipients of services or goods might be entities related to the subject, the expenses were ultimately incurred by the subject of the Account.
	//
	// Comment: Accounts can be applied to non-patients for tracking other non-patient related activities, such as group services (patients not tracked, and costs charged to another body), or might not be allocated.
	Subject []Reference `bson:"subject,omitempty" json:"subject,omitempty" xml:"subject,omitempty"`
	// The date range of services associated with this account.
	//
	// Comment: It is possible for transactions to be posted outside the service period, as long as the service was provided within the defined service period.
	ServicePeriod *Period `bson:"servicePeriod,omitempty" json:"servicePeriod,omitempty" xml:"servicePeriod,omitempty"`
	// The party(s) that are responsible for covering the payment of this account, and what order should they be applied to the account.
	//
	// Comment: Typically. this may be some form of insurance, internal charges, or self-pay.
	//
	// Local or jurisdictional business rules may determine which coverage covers which types of billable items charged to the account, and in which order.
	//
	// Where the order is important, a local/jurisdictional extension may be defined to specify the order for the type of charge.
	Coverage []AccountCoverage `bson:"coverage,omitempty" json:"coverage,omitempty" xml:"coverage,omitempty"`
	// Indicates the service area, hospital, department, etc.
	//
	// with responsibility for managing the Account.
	Owner *Reference `bson:"owner,omitempty" json:"owner,omitempty" xml:"owner,omitempty"`
	// Provides additional information about what the account tracks and how it is used.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// The parties responsible for balancing the account if other payment options fall short.
	Guarantor []AccountGuarantor `bson:"guarantor,omitempty" json:"guarantor,omitempty" xml:"guarantor,omitempty"`
	// When using an account for billing a specific Encounter the set of diagnoses that are relevant for billing are stored here on the account where they are able to be sequenced appropriately prior to processing to produce claim(s).
	Diagnosis []AccountDiagnosis `bson:"diagnosis,omitempty" json:"diagnosis,omitempty" xml:"diagnosis,omitempty"`
	// When using an account for billing a specific Encounter the set of procedures that are relevant for billing are stored here on the account where they are able to be sequenced appropriately prior to processing to produce claim(s).
	Procedure []AccountProcedure `bson:"procedure,omitempty" json:"procedure,omitempty" xml:"procedure,omitempty"`
	// Other associated accounts related to this account.
	RelatedAccount []AccountRelatedAccount `bson:"relatedAccount,omitempty" json:"relatedAccount,omitempty" xml:"relatedAccount,omitempty"`
	// The default currency for the account.
	Currency *CodeableConcept `bson:"currency,omitempty" json:"currency,omitempty" xml:"currency,omitempty"`
	// The calculated account balances - these are calculated and processed by the finance system.The balances with a `term` that is not current are usually generated/updated by an invoicing or similar process.
	Balance []AccountBalance `bson:"balance,omitempty" json:"balance,omitempty" xml:"balance,omitempty"`
	// Time the balance amount was calculated.
	//
	// Comment: The total(s) are not typically updated using the restful API but more often by back-end processes included in a typical finance system.
	CalculatedAt *string `bson:"calculatedAt,omitempty" json:"calculatedAt,omitempty" xml:"calculatedAt,omitempty"`
}

type AccountCoverage struct {
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
	// The party(s) that contribute to payment (or part of) of the charges applied to this account (including self-pay).
	//
	// A coverage may only be responsible for specific types of charges, and the sequence of the coverages in the account could be important when processing billing.
	Coverage Reference `bson:"coverage" json:"coverage" xml:"coverage"`
	// The priority of the coverage in the context of this account.
	//
	// Comment: It is common in some jurisdictions for there to be multiple coverages allocated to an account, and a sequence is required to order the settling of the account (often with insurance claiming).
	Priority *int `bson:"priority,omitempty" json:"priority,omitempty" xml:"priority,omitempty"`
}

type AccountGuarantor struct {
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
	// The entity who is responsible.
	Party Reference `bson:"party" json:"party" xml:"party"`
	// A guarantor may be placed on credit hold or otherwise have their role temporarily suspended.
	OnHold *bool `bson:"onHold,omitempty" json:"onHold,omitempty" xml:"onHold,omitempty"`
	// The timeframe during which the guarantor accepts responsibility for the account.
	Period *Period `bson:"period,omitempty" json:"period,omitempty" xml:"period,omitempty"`
}

type AccountDiagnosis struct {
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
	// Ranking of the diagnosis (for each type).
	Sequence *int `bson:"sequence,omitempty" json:"sequence,omitempty" xml:"sequence,omitempty"`
	// The diagnosis relevant to the account.
	Condition CodeableReference `bson:"condition" json:"condition" xml:"condition"`
	// Ranking of the diagnosis (for each type).
	DateOfDiagnosis *string `bson:"dateOfDiagnosis,omitempty" json:"dateOfDiagnosis,omitempty" xml:"dateOfDiagnosis,omitempty"`
	// Type that this diagnosis has relevant to the account (e.g.
	//
	// admission, billing, discharge …).
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Was the Diagnosis present on Admission in the related Encounter.
	OnAdmission *bool `bson:"onAdmission,omitempty" json:"onAdmission,omitempty" xml:"onAdmission,omitempty"`
	// The package code can be used to group diagnoses that may be priced or delivered as a single product.
	//
	// Such as DRGs.
	PackageCode []CodeableConcept `bson:"packageCode,omitempty" json:"packageCode,omitempty" xml:"packageCode,omitempty"`
}

type AccountProcedure struct {
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
	// Ranking of the procedure (for each type).
	Sequence *int `bson:"sequence,omitempty" json:"sequence,omitempty" xml:"sequence,omitempty"`
	// The procedure relevant to the account.
	Code CodeableReference `bson:"code" json:"code" xml:"code"`
	// Date of the procedure when using a coded procedure.
	//
	// If using a reference to a procedure, then the date on the procedure should be used.
	DateOfService *string `bson:"dateOfService,omitempty" json:"dateOfService,omitempty" xml:"dateOfService,omitempty"`
	// How this procedure value should be used in charging the account.
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The package code can be used to group procedures that may be priced or delivered as a single product.
	//
	// Such as DRGs.
	PackageCode []CodeableConcept `bson:"packageCode,omitempty" json:"packageCode,omitempty" xml:"packageCode,omitempty"`
	// Any devices that were associated with the procedure relevant to the account.
	Device []Reference `bson:"device,omitempty" json:"device,omitempty" xml:"device,omitempty"`
}

type AccountRelatedAccount struct {
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
	// Relationship of the associated Account.
	Relationship *CodeableConcept `bson:"relationship,omitempty" json:"relationship,omitempty" xml:"relationship,omitempty"`
	// Reference to an associated Account.
	Account Reference `bson:"account" json:"account" xml:"account"`
}

type AccountBalance struct {
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
	// Who is expected to pay this part of the balance.
	Aggregate *CodeableConcept `bson:"aggregate,omitempty" json:"aggregate,omitempty" xml:"aggregate,omitempty"`
	// The term of the account balances - The balance value is the amount that was outstanding for this age.
	Term *CodeableConcept `bson:"term,omitempty" json:"term,omitempty" xml:"term,omitempty"`
	// The amount is only an estimated value - this is likely common for `current` term balances, but not with known terms (that were generated by a backend process).
	Estimate *bool `bson:"estimate,omitempty" json:"estimate,omitempty" xml:"estimate,omitempty"`
	// The actual balance value calculated for the age defined in the term property.
	Amount Money `bson:"amount" json:"amount" xml:"amount"`
}

func (r Account) ResourceType() string {
	return "Account"
}

// MarshalJSON implements json.Marshaler
func (r Account) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Account
		ResourceType string `json:"resourceType"`
	}{
		Account:      r,
		ResourceType: r.ResourceType(),
	})
}
