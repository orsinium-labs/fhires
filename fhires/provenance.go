package fhires

import "encoding/json/v2"

// Provenance of a resource is a record that describes entities and processes involved in producing and delivering or otherwise influencing that resource.
//
// Provenance provides a critical foundation for assessing authenticity, enabling trust, and allowing reproducibility. Provenance assertions are a form of contextual metadata and can themselves become important records with their own provenance. Provenance statement indicates clinical significance in terms of confidence in authenticity, reliability, and trustworthiness, integrity, and stage in lifecycle (e.g. Document Completion - has the artifact been legally authenticated), all of which may impact security, privacy, and trust policies.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Provenance
type Provenance struct {
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
	// The Reference(s) that were generated or updated by  the activity described in this resource.
	//
	// A provenance can point to more than one target if multiple resources were created/updated by the same activity.
	//
	// Comment: Target references are usually version specific, but might not be, if a version has not been assigned or if the provenance information is part of the set of resources being maintained (i.e. a document). When using the RESTful API, the identity of the resource might not be known (especially not the version specific one); the client may either submit the resource first, and then the provenance, or it may submit both using a single transaction. See the notes on transaction for further discussion.
	Target []Reference `bson:"target" json:"target" xml:"target"`
	// The period during which the activity occurred.
	//
	// Comment: The period can be a little arbitrary; where possible, the time should correspond to human assessment of the activity time.
	OccurredPeriod   *Period `bson:"occurredPeriod,omitempty" json:"occurredPeriod,omitempty" xml:"occurredPeriod,omitempty"`
	OccurredDateTime *string `bson:"occurredDateTime,omitempty" json:"occurredDateTime,omitempty" xml:"occurredDateTime,omitempty"`
	// The instant of time at which the activity was recorded.
	//
	// Comment: This can be a little different from the time stamp on the resource if there is a delay between recording the event and updating the provenance and target resource.
	Recorded *string `bson:"recorded,omitempty" json:"recorded,omitempty" xml:"recorded,omitempty"`
	// Policy or plan the activity was defined by.
	//
	// Typically, a single activity may have multiple applicable policy documents, such as patient consent, guarantor funding, etc.
	//
	// Comment: For example: Where an OAuth token authorizes, the unique identifier from the OAuth token is placed into the policy element Where a policy engine (e.g. XACML) holds policy logic, the unique policy identifier is placed into the policy element.
	Policy []string `bson:"policy,omitempty" json:"policy,omitempty" xml:"policy,omitempty"`
	// Where the activity occurred, if relevant.
	Location *Reference `bson:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// The authorization (e.g., PurposeOfUse) that was used during the event being recorded.
	Authorization []CodeableReference `bson:"authorization,omitempty" json:"authorization,omitempty" xml:"authorization,omitempty"`
	// An activity is something that occurs over a period of time and acts upon or with entities; it may include consuming, processing, transforming, modifying, relocating, using, or generating entities.
	Activity *CodeableConcept `bson:"activity,omitempty" json:"activity,omitempty" xml:"activity,omitempty"`
	// Allows tracing of authorizatino for the events and tracking whether proposals/recommendations were acted upon.
	BasedOn []Reference `bson:"basedOn,omitempty" json:"basedOn,omitempty" xml:"basedOn,omitempty"`
	// The patient element is available to enable deterministic tracking of activities that involve the patient as the subject of the data used in an activity.
	Patient *Reference `bson:"patient,omitempty" json:"patient,omitempty" xml:"patient,omitempty"`
	// This will typically be the encounter the event occurred, but some events may be initiated prior to or after the official completion of an encounter but still be tied to the context of the encounter (e.g.
	//
	// pre-admission lab tests).
	Encounter *Reference `bson:"encounter,omitempty" json:"encounter,omitempty" xml:"encounter,omitempty"`
	// An actor taking a role in an activity  for which it can be assigned some degree of responsibility for the activity taking place.
	//
	// Comment: Several agents may be associated (i.e. has some responsibility for an activity) with an activity and vice-versa.
	Agent []ProvenanceAgent `bson:"agent" json:"agent" xml:"agent"`
	// An entity used in this activity.
	Entity []ProvenanceEntity `bson:"entity,omitempty" json:"entity,omitempty" xml:"entity,omitempty"`
	// A digital signature on the target Reference(s).
	//
	// The signer should match a Provenance.agent. The purpose of the signature is indicated.
	Signature []Signature `bson:"signature,omitempty" json:"signature,omitempty" xml:"signature,omitempty"`
}

type ProvenanceAgent struct {
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
	// The Functional Role of the agent with respect to the activity.
	//
	// Comment: For example: assembler, author, prescriber, signer, investigator, etc.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The structural roles of the agent indicating the agent's competency.
	//
	// The security role enabling the agent with respect to the activity.
	//
	// Comment: For example: Chief-of-Radiology, Nurse, Physician, Medical-Student, etc.
	Role []CodeableConcept `bson:"role,omitempty" json:"role,omitempty" xml:"role,omitempty"`
	// Indicates who or what performed in the event.
	Who Reference `bson:"who" json:"who" xml:"who"`
	// The agent that delegated authority to perform the activity performed by the agent.who element.
	OnBehalfOf *Reference `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty" xml:"onBehalfOf,omitempty"`
}

type ProvenanceEntity struct {
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
	// How the entity was used during the activity.
	Role ProvenanceEntityRole `bson:"role" json:"role" xml:"role"`
	// Identity of the  Entity used.
	//
	// May be a logical or physical uri and maybe absolute or relative.
	//
	// Comment: whatIdentity should be used for entities that are not a Resource type.
	What Reference `bson:"what" json:"what" xml:"what"`
	// The entity is attributed to an agent to express the agent's responsibility for that entity, possibly along with other agents.
	//
	// This description can be understood as shorthand for saying that the agent was responsible for the activity which used the entity.
	//
	// Comment: A usecase where one Provenance.entity.agent is used where the Entity that was used in the creation/updating of the Target, is not in the context of the same custodianship as the Target, and thus the meaning of Provenance.entity.agent is to say that the entity referenced is managed elsewhere and that this Agent provided access to it.  This would be similar to where the Entity being referenced is managed outside FHIR, such as through HL7 V2, v3, or XDS. This might be where the Entity being referenced is managed in another FHIR resource server. Thus it explains the Provenance of that Entity's use in the context of this Provenance activity.
	Agent []ProvenanceAgent `bson:"agent,omitempty" json:"agent,omitempty"`
}

func (r Provenance) ResourceType() string {
	return "Provenance"
}

// MarshalJSON implements json.Marshaler
func (r Provenance) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Provenance
		ResourceType string `json:"resourceType"`
	}{
		Provenance:   r,
		ResourceType: r.ResourceType(),
	})
}
