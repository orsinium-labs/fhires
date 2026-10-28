package fhires

import "encoding/json/v2"

// A record of a healthcare consumer’s  choices  or choices made on their behalf by a third party, which permits or denies identified recipient(s) or recipient role(s) to perform one or more actions within a given policy context, for specific purposes and periods of time.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Consent
type Consent struct {
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
	// Unique identifier for this copy of the Consent Statement.
	//
	// Comment: This identifier identifies this copy of the consent. Where this identifier is also used elsewhere as the identifier for a consent record (e.g. a CDA consent document) then the consent details are expected to be the same.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Indicates the current state of this Consent resource.
	//
	// Comment: This element is labeled as a modifier because the status contains the codes rejected and entered-in-error that mark the Consent as not currently valid.
	Status ConsentState `bson:"status" json:"status" xml:"status"`
	// A classification of the type of consents found in the statement.
	//
	// This element supports indexing and retrieval of consent statements.
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
	// The patient/healthcare practitioner or group of persons to whom this consent applies.
	Subject *Reference `bson:"subject,omitempty" json:"subject,omitempty" xml:"subject,omitempty"`
	// Date the consent instance was agreed to.
	Date *string `bson:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// Effective period for this Consent Resource and all provisions unless specified in that provision.
	Period *Period `bson:"period,omitempty" json:"period,omitempty" xml:"period,omitempty"`
	// The entity responsible for granting the rights listed in a Consent Directive.
	Grantor []Reference `bson:"grantor,omitempty" json:"grantor,omitempty" xml:"grantor,omitempty"`
	// The entity responsible for complying with the Consent Directive, including any obligations or limitations on authorizations and enforcement of prohibitions.
	//
	// Comment: In a fully computable consent, both grantee and grantor  will be listed as actors within the consent. The Grantee and Grantor elements are for ease of search only.
	Grantee []Reference `bson:"grantee,omitempty" json:"grantee,omitempty" xml:"grantee,omitempty"`
	// The actor that manages the consent through its lifecycle.
	Manager []Reference `bson:"manager,omitempty" json:"manager,omitempty" xml:"manager,omitempty"`
	// The actor that controls/enforces the access according to the consent.
	Controller []Reference `bson:"controller,omitempty" json:"controller,omitempty" xml:"controller,omitempty"`
	// The source on which this consent statement is based.
	//
	// The source might be a scanned original paper form.
	//
	// Comment: The source can be contained inline (Attachment), referenced directly (Consent), referenced in a consent repository (DocumentReference), or simply by an identifier (Identifier), e.g. a CDA document id.
	SourceAttachment []Attachment `bson:"sourceAttachment,omitempty" json:"sourceAttachment,omitempty" xml:"sourceAttachment,omitempty"`
	// A reference to a consent that links back to such a source, a reference to a document repository (e.g.
	//
	// XDS) that stores the original consent document.
	//
	// Comment: The source can be contained inline (Attachment), referenced directly (Consent), referenced in a consent repository (DocumentReference), or simply by an identifier (Identifier), e.g. a CDA document id.
	SourceReference []Reference `bson:"sourceReference,omitempty" json:"sourceReference,omitempty" xml:"sourceReference,omitempty"`
	// A set of codes that indicate the regulatory basis (if any) that this consent supports.
	RegulatoryBasis []CodeableConcept `bson:"regulatoryBasis,omitempty" json:"regulatoryBasis,omitempty" xml:"regulatoryBasis,omitempty"`
	// A Reference or URL used to uniquely identify the policy the organization will enforce for this Consent.
	//
	// This Reference or URL should be specific to the version of the policy and should be dereferencable to a computable policy of some form.
	PolicyBasis *ConsentPolicyBasis `bson:"policyBasis,omitempty" json:"policyBasis,omitempty" xml:"policyBasis,omitempty"`
	// A Reference to the human readable policy explaining the basis for the Consent.
	PolicyText []Reference `bson:"policyText,omitempty" json:"policyText,omitempty" xml:"policyText,omitempty"`
	// Whether a treatment instruction (e.g.
	//
	// artificial respiration: yes or no) was verified with the patient, his/her family or another authorized person.
	Verification []ConsentVerification `bson:"verification,omitempty" json:"verification,omitempty" xml:"verification,omitempty"`
	// Action to take - permit or deny - as default.
	Decision *ConsentProvisionType `bson:"decision,omitempty" json:"decision,omitempty" xml:"decision,omitempty"`
	// An exception to the base policy of this consent.
	//
	// An exception can be an addition or removal of access permissions.
	Provision []ConsentProvision `bson:"provision,omitempty" json:"provision,omitempty" xml:"provision,omitempty"`
}

type ConsentPolicyBasis struct {
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
	// A Reference that identifies the policy the organization will enforce for this Consent.
	//
	// Comment: While any resource may be used, Consent, PlanDefinition and Contract would be most frequent
	Reference *Reference `bson:"reference,omitempty" json:"reference,omitempty" xml:"reference,omitempty"`
	// A URL that links to a computable version of the policy the organization will enforce for this Consent.
	URL *string `bson:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
}

type ConsentVerification struct {
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
	// Has the instruction been verified.
	Verified bool `bson:"verified" json:"verified" xml:"verified"`
	// Extensible list of verification type starting with verification and re-validation.
	//
	// Comment: This allows the verification element to hold multiple use cases including RelatedPerson verification of the Grantee decision and periodic re-validation of the consent.
	VerificationType *CodeableConcept `bson:"verificationType,omitempty" json:"verificationType,omitempty" xml:"verificationType,omitempty"`
	// The person who conducted the verification/validation of the Grantor decision.
	VerifiedBy *Reference `bson:"verifiedBy,omitempty" json:"verifiedBy,omitempty" xml:"verifiedBy,omitempty"`
	// Who verified the instruction (Patient, Relative or other Authorized Person).
	VerifiedWith *Reference `bson:"verifiedWith,omitempty" json:"verifiedWith,omitempty" xml:"verifiedWith,omitempty"`
	// Date(s) verification was collected.
	//
	// Comment: Allows for history of verification/validation.
	VerificationDate []string `bson:"verificationDate,omitempty" json:"verificationDate,omitempty" xml:"verificationDate,omitempty"`
}

type ConsentProvision struct {
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
	// Timeframe for this provision.
	//
	// Comment: This is the bound effective time of the consent and should be in the base provision in the Consent resource.
	Period *Period `bson:"period,omitempty" json:"period,omitempty" xml:"period,omitempty"`
	// Who or what is controlled by this provision.
	//
	// Use group to identify a set of actors by some property they share (e.g. 'admitting officers').
	Actor []ConsentProvisionActor `bson:"actor,omitempty" json:"actor,omitempty" xml:"actor,omitempty"`
	// Actions controlled by this provision.
	//
	// Comment: Note that this is the direct action (not the grounds for the action covered in the purpose element). At present, the only action in the understood and tested scope of this resource is 'read'.
	Action []CodeableConcept `bson:"action,omitempty" json:"action,omitempty" xml:"action,omitempty"`
	// A security label, comprised of 0..* security label fields (Privacy tags), which define which resources are controlled by this exception.
	//
	// Comment: If the consent specifies a security label of "R" then it applies to all resources that are labeled "R" or lower. E.g. for Confidentiality, it's a high water mark. For other kinds of security labels, subsumption logic applies. When the purpose of use tag is on the data, access request purpose of use shall not conflict.
	SecurityLabel []Coding `bson:"securityLabel,omitempty" json:"securityLabel,omitempty" xml:"securityLabel,omitempty"`
	// The context of the activities a user is taking - why the user is accessing the data - that are controlled by this provision.
	//
	// Comment: When the purpose of use tag is on the data, access request purpose of use shall not conflict.
	Purpose []Coding `bson:"purpose,omitempty" json:"purpose,omitempty" xml:"purpose,omitempty"`
	// The documentType(s) covered by this provision.
	//
	// The type can be a CDA document, or some other type that indicates what sort of information the consent relates to.
	//
	// Comment: Multiple types are or'ed together. The intention of the documentType element is that the codes refer to document types defined in a standard somewhere.
	DocumentType []Coding `bson:"documentType,omitempty" json:"documentType,omitempty" xml:"documentType,omitempty"`
	// The resourceType(s) covered by this provision.
	//
	// The type can be a FHIR resource type or a profile on a type that indicates what information the consent relates to.
	//
	// Comment: Multiple types are or'ed together. The intention of the resourceType element is that the codes refer to profiles or document types defined in a standard or an implementation guide somewhere.
	ResourceType []Coding `bson:"resourceType,omitempty" json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// If this code is found in an instance, then the provision applies.
	Code []CodeableConcept `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// Clinical or Operational Relevant period of time that bounds the data controlled by this provision.
	//
	// Comment: This has a different sense to the Consent.period - that is when the consent agreement holds. This is the time period of the data that is controlled by the agreement.
	DataPeriod *Period `bson:"dataPeriod,omitempty" json:"dataPeriod,omitempty" xml:"dataPeriod,omitempty"`
	// The resources controlled by this provision if specific resources are referenced.
	Data []ConsentProvisionData `bson:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	// A computable (FHIRPath or other) definition of what is controlled by this consent.
	//
	// Comment: Constraining the expression type for a specific implementation via profile is recommended
	Expression *Expression `bson:"expression,omitempty" json:"expression,omitempty" xml:"expression,omitempty"`
	// Provisions which provide exceptions to the base provision or subprovisions.
	Provision []ConsentProvision `bson:"provision,omitempty" json:"provision,omitempty"`
}

type ConsentProvisionActor struct {
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
	// How the individual is involved in the resources content that is described in the exception.
	Role *CodeableConcept `bson:"role,omitempty" json:"role,omitempty" xml:"role,omitempty"`
	// The resource that identifies the actor.
	//
	// To identify actors by type, use group to identify a set of actors by some property they share (e.g. 'admitting officers').
	Reference *Reference `bson:"reference,omitempty" json:"reference,omitempty" xml:"reference,omitempty"`
}

type ConsentProvisionData struct {
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
	// How the resource reference is interpreted when testing consent restrictions.
	Meaning ConsentDataMeaning `bson:"meaning" json:"meaning" xml:"meaning"`
	// A reference to a specific resource that defines which resources are covered by this consent.
	Reference Reference `bson:"reference" json:"reference" xml:"reference"`
}

func (r Consent) ResourceType() string {
	return "Consent"
}

// MarshalJSON implements json.Marshaler
func (r Consent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Consent
		ResourceType string `json:"resourceType"`
	}{
		Consent:      r,
		ResourceType: r.ResourceType(),
	})
}
