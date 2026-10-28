package fhires

import v2 "encoding/json/v2"

// Regulatory approval, clearance or licencing related to a regulated product, treatment, facility or activity that is cited in a guidance, regulation, rule or legislative act.
//
// An example is Market Authorization relating to a Medicinal Product.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/RegulatedAuthorization
type RegulatedAuthorization struct {
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
	// Business identifier for the authorization, typically assigned by the authorizing body.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The product type, treatment, facility or activity that is being authorized.
	Subject []Reference `bson:"subject,omitempty" json:"subject,omitempty" xml:"subject,omitempty"`
	// Overall type of this authorization, for example drug marketing approval, orphan drug designation.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// General textual supporting information.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// The territory (e.g., country, jurisdiction etc.) in which the authorization has been granted.
	Region []CodeableConcept `bson:"region,omitempty" json:"region,omitempty" xml:"region,omitempty"`
	// The status that is authorised e.g.
	//
	// approved. Intermediate states and actions can be tracked with cases and applications.
	Status *CodeableConcept `bson:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// The date at which the current status was assigned.
	StatusDate *string `bson:"statusDate,omitempty" json:"statusDate,omitempty" xml:"statusDate,omitempty"`
	// The time period in which the regulatory approval, clearance or licencing is in effect.
	//
	// As an example, a Marketing Authorization includes the date of authorization and/or an expiration date.
	ValidityPeriod *Period `bson:"validityPeriod,omitempty" json:"validityPeriod,omitempty" xml:"validityPeriod,omitempty"`
	// Condition for which the use of the regulated product applies.
	Indication []CodeableReference `bson:"indication,omitempty" json:"indication,omitempty" xml:"indication,omitempty"`
	// The intended use of the product, e.g. prevention, treatment, diagnosis.
	IntendedUse *CodeableConcept `bson:"intendedUse,omitempty" json:"intendedUse,omitempty" xml:"intendedUse,omitempty"`
	// The legal or regulatory framework against which this authorization is granted, or other reasons for it.
	Basis []CodeableConcept `bson:"basis,omitempty" json:"basis,omitempty" xml:"basis,omitempty"`
	// The organization that has been granted this authorization, by some authoritative body (the 'regulator').
	Holder *Reference `bson:"holder,omitempty" json:"holder,omitempty" xml:"holder,omitempty"`
	// The regulatory authority or authorizing body granting the authorization.
	//
	// For example, European Medicines Agency (EMA), Food and Drug Administration (FDA), Health Canada (HC), etc.
	Regulator *Reference `bson:"regulator,omitempty" json:"regulator,omitempty" xml:"regulator,omitempty"`
	// Additional information or supporting documentation about the authorization.
	AttachedDocument []Reference `bson:"attachedDocument,omitempty" json:"attachedDocument,omitempty" xml:"attachedDocument,omitempty"`
	// The case or regulatory procedure for granting or amending a regulated authorization.
	//
	// An authorization is granted in response to submissions/applications by those seeking authorization. A case is the administrative process that deals with the application(s) that relate to this and assesses them. Note: This area is subject to ongoing review and the workgroup is seeking implementer feedback on its use (see link at bottom of page).
	Case *RegulatedAuthorizationCase `bson:"case,omitempty" json:"case,omitempty" xml:"case,omitempty"`
}

type RegulatedAuthorizationCase struct {
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
	// Identifier by which this case can be referenced.
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The defining type of case.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The status associated with the case.
	Status *CodeableConcept `bson:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Relevant date for this case.
	DatePeriod   *Period `bson:"datePeriod,omitempty" json:"datePeriod,omitempty" xml:"datePeriod,omitempty"`
	DateDateTime *string `bson:"dateDateTime,omitempty" json:"dateDateTime,omitempty" xml:"dateDateTime,omitempty"`
	// A regulatory submission from an organization to a regulator, as part of an assessing case.
	//
	// Multiple applications may occur over time, with more or different information to support or modify the submission or the authorization. The applications can be considered as steps within the longer running case or procedure for this authorization process.
	Application []RegulatedAuthorizationCase `bson:"application,omitempty" json:"application,omitempty"`
}

func (r RegulatedAuthorization) ResourceType() string {
	return "RegulatedAuthorization"
}

// MarshalJSON implements json.Marshaler
func (r RegulatedAuthorization) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		RegulatedAuthorization
		ResourceType string `json:"resourceType"`
	}{
		RegulatedAuthorization: r,
		ResourceType:           r.ResourceType(),
	})
}
