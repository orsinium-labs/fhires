package fhires

import v2 "encoding/json/v2"

// This resource provides eligibility and plan details from the processing of an CoverageEligibilityRequest resource.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/CoverageEligibilityResponse
type CoverageEligibilityResponse struct {
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
	// A unique identifier assigned to this coverage eligiblity request.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The status of the resource instance.
	//
	// Comment: This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid.
	Status FinancialResourceStatusCodes `bson:"status" json:"status" xml:"status"`
	// Code to specify whether requesting: prior authorization requirements for some service categories or billing codes; benefits for coverages specified or discovered; discovery and return of coverages for the patient; and/or validation that the specified coverage is in-force at the date/period specified or 'now' if not specified.
	Purpose []EligibilityResponsePurpose `bson:"purpose" json:"purpose" xml:"purpose"`
	// The party who is the beneficiary of the supplied coverage and for whom eligibility is sought.
	Patient Reference `bson:"patient" json:"patient" xml:"patient"`
	// Information code for an event with a corresponding date or period.
	Event []CoverageEligibilityResponseEvent `bson:"event,omitempty" json:"event,omitempty" xml:"event,omitempty"`
	// The date or dates when the enclosed suite of services were performed or completed.
	ServicedDate   *string `bson:"servicedDate,omitempty" json:"servicedDate,omitempty" xml:"servicedDate,omitempty"`
	ServicedPeriod *Period `bson:"servicedPeriod,omitempty" json:"servicedPeriod,omitempty" xml:"servicedPeriod,omitempty"`
	// The date this resource was created.
	Created string `bson:"created" json:"created" xml:"created"`
	// The provider which is responsible for the request.
	//
	// Comment: This party is responsible for the claim but not necessarily professionally responsible for the provision of the individual products and services listed below. This field is the Billing Provider, for example, a facility, provider group, lab or practitioner.
	Requestor *Reference `bson:"requestor,omitempty" json:"requestor,omitempty" xml:"requestor,omitempty"`
	// Reference to the original request resource.
	Request Reference `bson:"request" json:"request" xml:"request"`
	// The outcome of the request processing.
	//
	// Comment: The resource may be used to indicate that: the request has been held (queued) for processing; that it has been processed and errors found (error); that no errors were found and that some of the adjudication has been undertaken (partial) or that all of the adjudication has been undertaken (complete).
	Outcome EligibilityOutcome `bson:"outcome" json:"outcome" xml:"outcome"`
	// A human readable description of the status of the adjudication.
	Disposition *string `bson:"disposition,omitempty" json:"disposition,omitempty" xml:"disposition,omitempty"`
	// The Insurer who issued the coverage in question and is the author of the response.
	Insurer Reference `bson:"insurer" json:"insurer" xml:"insurer"`
	// Financial instruments for reimbursement for the health care products and services.
	//
	// Comment: All insurance coverages for the patient which may be applicable for reimbursement, of the products and services listed in the claim, are typically provided in the claim to allow insurers to confirm the ordering of the insurance coverages relative to local 'coordination of benefit' rules. One coverage (and only one) with 'focal=true' is to be used in the adjudication of this claim. Coverages appearing before the focal Coverage in the list, and where 'subrogation=false', should provide a reference to the ClaimResponse containing the adjudication results of the prior claim.
	Insurance []CoverageEligibilityResponseInsurance `bson:"insurance,omitempty" json:"insurance,omitempty" xml:"insurance,omitempty"`
	// A reference from the Insurer to which these services pertain to be used on further communication and as proof that the request occurred.
	PreAuthRef *string `bson:"preAuthRef,omitempty" json:"preAuthRef,omitempty" xml:"preAuthRef,omitempty"`
	// A code for the form to be used for printing the content.
	//
	// Comment: May be needed to identify specific jurisdictional forms.
	Form *CodeableConcept `bson:"form,omitempty" json:"form,omitempty" xml:"form,omitempty"`
	// Errors encountered during the processing of the request.
	Error []CoverageEligibilityResponseError `bson:"error,omitempty" json:"error,omitempty" xml:"error,omitempty"`
}

type CoverageEligibilityResponseEvent struct {
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

type CoverageEligibilityResponseInsurance struct {
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
	// Reference to the insurance card level information contained in the Coverage resource.
	//
	// The coverage issuing insurer will use these details to locate the patient's actual coverage within the insurer's information system.
	Coverage Reference `bson:"coverage" json:"coverage" xml:"coverage"`
	// Flag indicating if the coverage provided is inforce currently if no service date(s) specified or for the whole duration of the service dates.
	Inforce *bool `bson:"inforce,omitempty" json:"inforce,omitempty" xml:"inforce,omitempty"`
	// The term of the benefits documented in this response.
	BenefitPeriod *Period `bson:"benefitPeriod,omitempty" json:"benefitPeriod,omitempty" xml:"benefitPeriod,omitempty"`
	// Benefits and optionally current balances, and authorization details by category or service.
	Item []CoverageEligibilityResponseInsuranceItem `bson:"item,omitempty" json:"item,omitempty" xml:"item,omitempty"`
}

type CoverageEligibilityResponseInsuranceItem struct {
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
	// Code to identify the general type of benefits under which products and services are provided.
	//
	// Comment: Examples include Medical Care, Periodontics, Renal Dialysis, Vision Coverage.
	Category *CodeableConcept `bson:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
	// This contains the product, service, drug or other billing code for the item.
	//
	// Comment: Code to indicate the Professional Service or Product supplied (e.g. CTP, HCPCS, USCLS, ICD10, NCPDP, DIN, RxNorm, ACHI, CCI).
	ProductOrService *CodeableConcept `bson:"productOrService,omitempty" json:"productOrService,omitempty" xml:"productOrService,omitempty"`
	// Item typification or modifiers codes to convey additional context for the product or service.
	//
	// Comment: For example in Oral whether the treatment is cosmetic or associated with TMJ, or for Medical whether the treatment was outside the clinic or out of office hours.
	Modifier []CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty" xml:"modifier,omitempty"`
	// The practitioner who is eligible for the provision of the product or service.
	Provider *Reference `bson:"provider,omitempty" json:"provider,omitempty" xml:"provider,omitempty"`
	// True if the indicated class of service is excluded from the plan, missing or False indicates the product or service is included in the coverage.
	Excluded *bool `bson:"excluded,omitempty" json:"excluded,omitempty" xml:"excluded,omitempty"`
	// A short name or tag for the benefit.
	//
	// Comment: For example: MED01, or DENT2.
	Name *string `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// A richer description of the benefit or services covered.
	//
	// Comment: For example 'DENT2 covers 100% of basic, 50% of major but excludes Ortho, Implants and Cosmetic services'.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Is a flag to indicate whether the benefits refer to in-network providers or out-of-network providers.
	Network *CodeableConcept `bson:"network,omitempty" json:"network,omitempty" xml:"network,omitempty"`
	// Indicates if the benefits apply to an individual or to the family.
	Unit *CodeableConcept `bson:"unit,omitempty" json:"unit,omitempty" xml:"unit,omitempty"`
	// The term or period of the values such as 'maximum lifetime benefit' or 'maximum annual visits'.
	Term *CodeableConcept `bson:"term,omitempty" json:"term,omitempty" xml:"term,omitempty"`
	// Benefits used to date.
	Benefit []CoverageEligibilityResponseInsuranceItemBenefit `bson:"benefit,omitempty" json:"benefit,omitempty" xml:"benefit,omitempty"`
	// A boolean flag indicating whether a preauthorization is required prior to actual service delivery.
	AuthorizationRequired *bool `bson:"authorizationRequired,omitempty" json:"authorizationRequired,omitempty" xml:"authorizationRequired,omitempty"`
	// Codes or comments regarding information or actions associated with the preauthorization.
	AuthorizationSupporting []CodeableConcept `bson:"authorizationSupporting,omitempty" json:"authorizationSupporting,omitempty" xml:"authorizationSupporting,omitempty"`
	// A web location for obtaining requirements or descriptive information regarding the preauthorization.
	AuthorizationUURL *string `bson:"authorizationUrl,omitempty" json:"authorizationUrl,omitempty" xml:"authorizationUrl,omitempty"`
}

type CoverageEligibilityResponseInsuranceItemBenefit struct {
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
	// Classification of benefit being provided.
	//
	// Comment: For example: deductible, visits, benefit amount.
	Type CodeableConcept `bson:"type" json:"type" xml:"type"`
	// The quantity of the benefit which is permitted under the coverage.
	AllowedUnsignedInt *int    `bson:"allowedUnsignedInt,omitempty" json:"allowedUnsignedInt,omitempty" xml:"allowedUnsignedInt,omitempty"`
	AllowedString      *string `bson:"allowedString,omitempty" json:"allowedString,omitempty" xml:"allowedString,omitempty"`
	AllowedMoney       *Money  `bson:"allowedMoney,omitempty" json:"allowedMoney,omitempty" xml:"allowedMoney,omitempty"`
	// The quantity of the benefit which have been consumed to date.
	UsedUnsignedInt *int    `bson:"usedUnsignedInt,omitempty" json:"usedUnsignedInt,omitempty" xml:"usedUnsignedInt,omitempty"`
	UsedString      *string `bson:"usedString,omitempty" json:"usedString,omitempty" xml:"usedString,omitempty"`
	UsedMoney       *Money  `bson:"usedMoney,omitempty" json:"usedMoney,omitempty" xml:"usedMoney,omitempty"`
}

type CoverageEligibilityResponseError struct {
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
	// An error code,from a specified code system, which details why the eligibility check could not be performed.
	Code CodeableConcept `bson:"code" json:"code" xml:"code"`
	// A [simple subset of FHIRPath](fhirpath.html#simple) limited to element names, repetition indicators and the default child accessor that identifies one of the elements in the resource that caused this issue to be raised.
	//
	// Comment: The root of the FHIRPath is the resource or bundle that generated OperationOutcome.  Each FHIRPath SHALL resolve to a single node.
	Expression []string `bson:"expression,omitempty" json:"expression,omitempty" xml:"expression,omitempty"`
}

func (r CoverageEligibilityResponse) ResourceType() string {
	return "CoverageEligibilityResponse"
}

// MarshalJSON implements json.Marshaler
func (r CoverageEligibilityResponse) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		CoverageEligibilityResponse
		ResourceType string `json:"resourceType"`
	}{
		CoverageEligibilityResponse: r,
		ResourceType:                r.ResourceType(),
	})
}
