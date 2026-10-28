package fhires

import v2 "encoding/json/v2"

// Describes validation requirements, source(s), status and dates for one or more elements.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/VerificationResult
type VerificationResult struct {
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
	// A resource that was validated.
	Target []Reference `bson:"target,omitempty" json:"target,omitempty" xml:"target,omitempty"`
	// The fhirpath location(s) within the resource that was validated.
	TargetLocation []string `bson:"targetLocation,omitempty" json:"targetLocation,omitempty" xml:"targetLocation,omitempty"`
	// The frequency with which the target must be validated (none; initial; periodic).
	Need *CodeableConcept `bson:"need,omitempty" json:"need,omitempty" xml:"need,omitempty"`
	// The validation status of the target (attested; validated; in process; requires revalidation; validation failed; revalidation failed).
	Status VerificationResultStatus `bson:"status" json:"status" xml:"status"`
	// When the validation status was updated.
	StatusDate *string `bson:"statusDate,omitempty" json:"statusDate,omitempty" xml:"statusDate,omitempty"`
	// What the target is validated against (nothing; primary source; multiple sources).
	ValidationType *CodeableConcept `bson:"validationType,omitempty" json:"validationType,omitempty" xml:"validationType,omitempty"`
	// The primary process by which the target is validated (edit check; value set; primary source; multiple sources; standalone; in context).
	ValidationProcess []CodeableConcept `bson:"validationProcess,omitempty" json:"validationProcess,omitempty" xml:"validationProcess,omitempty"`
	// Frequency of revalidation.
	Frequency *Timing `bson:"frequency,omitempty" json:"frequency,omitempty" xml:"frequency,omitempty"`
	// The date/time validation was last completed (including failed validations).
	LastPerformed *string `bson:"lastPerformed,omitempty" json:"lastPerformed,omitempty" xml:"lastPerformed,omitempty"`
	// The date when target is next validated, if appropriate.
	NextScheduled *string `bson:"nextScheduled,omitempty" json:"nextScheduled,omitempty" xml:"nextScheduled,omitempty"`
	// The result if validation fails (fatal; warning; record only; none).
	FailureAction *CodeableConcept `bson:"failureAction,omitempty" json:"failureAction,omitempty" xml:"failureAction,omitempty"`
	// Information about the primary source(s) involved in validation.
	PrimarySource []VerificationResultPrimarySource `bson:"primarySource,omitempty" json:"primarySource,omitempty" xml:"primarySource,omitempty"`
	// Information about the entity attesting to information.
	Attestation *VerificationResultAttestation `bson:"attestation,omitempty" json:"attestation,omitempty" xml:"attestation,omitempty"`
	// Information about the entity validating information.
	Validator []VerificationResultValidator `bson:"validator,omitempty" json:"validator,omitempty" xml:"validator,omitempty"`
}

type VerificationResultPrimarySource struct {
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
	// Reference to the primary source.
	Who *Reference `bson:"who,omitempty" json:"who,omitempty" xml:"who,omitempty"`
	// Type of primary source (License Board; Primary Education; Continuing Education; Postal Service; Relationship owner; Registration Authority; legal source; issuing source; authoritative source).
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Method for communicating with the primary source (manual; API; Push).
	CommunicationMethod []CodeableConcept `bson:"communicationMethod,omitempty" json:"communicationMethod,omitempty" xml:"communicationMethod,omitempty"`
	// Status of the validation of the target against the primary source (successful; failed; unknown).
	ValidationStatus *CodeableConcept `bson:"validationStatus,omitempty" json:"validationStatus,omitempty" xml:"validationStatus,omitempty"`
	// When the target was validated against the primary source.
	ValidationDate *string `bson:"validationDate,omitempty" json:"validationDate,omitempty" xml:"validationDate,omitempty"`
	// Ability of the primary source to push updates/alerts (yes; no; undetermined).
	CanPushUpdates *CodeableConcept `bson:"canPushUpdates,omitempty" json:"canPushUpdates,omitempty" xml:"canPushUpdates,omitempty"`
	// Type of alerts/updates the primary source can send (specific requested changes; any changes; as defined by source).
	PushTypeAvailable []CodeableConcept `bson:"pushTypeAvailable,omitempty" json:"pushTypeAvailable,omitempty" xml:"pushTypeAvailable,omitempty"`
}

type VerificationResultAttestation struct {
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
	// The individual or organization attesting to information.
	Who *Reference `bson:"who,omitempty" json:"who,omitempty" xml:"who,omitempty"`
	// When the who is asserting on behalf of another (organization or individual).
	OnBehalfOf *Reference `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty" xml:"onBehalfOf,omitempty"`
	// The method by which attested information was submitted/retrieved (manual; API; Push).
	CommunicationMethod *CodeableConcept `bson:"communicationMethod,omitempty" json:"communicationMethod,omitempty" xml:"communicationMethod,omitempty"`
	// The date the information was attested to.
	Date *string `bson:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// A digital identity certificate associated with the attestation source.
	SourceIdentityCertificate *string `bson:"sourceIdentityCertificate,omitempty" json:"sourceIdentityCertificate,omitempty" xml:"sourceIdentityCertificate,omitempty"`
	// A digital identity certificate associated with the proxy entity submitting attested information on behalf of the attestation source.
	ProxyIdentityCertificate *string `bson:"proxyIdentityCertificate,omitempty" json:"proxyIdentityCertificate,omitempty" xml:"proxyIdentityCertificate,omitempty"`
	// Signed assertion by the proxy entity indicating that they have the right to submit attested information on behalf of the attestation source.
	ProxySignature *Signature `bson:"proxySignature,omitempty" json:"proxySignature,omitempty" xml:"proxySignature,omitempty"`
	// Signed assertion by the attestation source that they have attested to the information.
	SourceSignature *Signature `bson:"sourceSignature,omitempty" json:"sourceSignature,omitempty" xml:"sourceSignature,omitempty"`
}

type VerificationResultValidator struct {
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
	// Reference to the organization validating information.
	Organization Reference `bson:"organization" json:"organization" xml:"organization"`
	// A digital identity certificate associated with the validator.
	IdentityCertificate *string `bson:"identityCertificate,omitempty" json:"identityCertificate,omitempty" xml:"identityCertificate,omitempty"`
	// Signed assertion by the validator that they have validated the information.
	AttestationSignature *Signature `bson:"attestationSignature,omitempty" json:"attestationSignature,omitempty" xml:"attestationSignature,omitempty"`
}

func (r VerificationResult) ResourceType() string {
	return "VerificationResult"
}

// MarshalJSON implements json.Marshaler
func (r VerificationResult) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		VerificationResult
		ResourceType string `json:"resourceType"`
	}{
		ResourceType:       r.ResourceType(),
		VerificationResult: r,
	})
}
