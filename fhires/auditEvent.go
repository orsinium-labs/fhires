package fhires

import v2 "encoding/json/v2"

// A record of an event relevant for purposes such as operations, privacy, security, maintenance, and performance analysis.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/AuditEvent
type AuditEvent struct {
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
	// Classification of the type of event.
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
	// Describes what happened. The most specific code for the event.
	Code CodeableConcept `bson:"code" json:"code" xml:"code"`
	// Indicator for type of action performed during the event that generated the audit.
	Action *AuditEventAction `bson:"action,omitempty" json:"action,omitempty" xml:"action,omitempty"`
	// Indicates and enables segmentation of various severity including debugging from critical.
	//
	// Comment: ATNA will map this to the SYSLOG PRI element.
	Severity *AuditEventSeverity `bson:"severity,omitempty" json:"severity,omitempty" xml:"severity,omitempty"`
	// The time or period during which the activity occurred.
	//
	// Comment: The time or period can be a little arbitrary; where possible, the time should correspond to human assessment of the activity time.
	OccurredPeriod   *Period `bson:"occurredPeriod,omitempty" json:"occurredPeriod,omitempty" xml:"occurredPeriod,omitempty"`
	OccurredDateTime *string `bson:"occurredDateTime,omitempty" json:"occurredDateTime,omitempty" xml:"occurredDateTime,omitempty"`
	// The time when the event was recorded.
	//
	// Comment: In a distributed system, some sort of common time base (e.g. an NTP [RFC1305] server) is a good implementation tactic.
	Recorded string `bson:"recorded" json:"recorded" xml:"recorded"`
	// Indicates whether the event succeeded or failed.
	//
	// A free text descripiton can be given in outcome.text.
	//
	// Comment: In some cases a "success" may be partial, for example, an incomplete or interrupted transfer of a radiological study. For the purpose of establishing accountability, these distinctions are not relevant.
	Outcome *AuditEventOutcome `bson:"outcome,omitempty" json:"outcome,omitempty" xml:"outcome,omitempty"`
	// The authorization (e.g., PurposeOfUse) that was used during the event being recorded.
	//
	// Comment: Use AuditEvent.agent.authorization when you know that it is specific to the agent, otherwise use AuditEvent.authorization. For example, during a machine-to-machine transfer it might not be obvious to the audit system who caused the event, but it does know why.
	Authorization []CodeableConcept `bson:"authorization,omitempty" json:"authorization,omitempty" xml:"authorization,omitempty"`
	// Allows tracing of authorizatino for the events and tracking whether proposals/recommendations were acted upon.
	BasedOn []Reference `bson:"basedOn,omitempty" json:"basedOn,omitempty" xml:"basedOn,omitempty"`
	// The patient element is available to enable deterministic tracking of activities that involve the patient as the subject of the data used in an activity.
	Patient *Reference `bson:"patient,omitempty" json:"patient,omitempty" xml:"patient,omitempty"`
	// This will typically be the encounter the event occurred, but some events may be initiated prior to or after the official completion of an encounter but still be tied to the context of the encounter (e.g.
	//
	// pre-admission lab tests).
	Encounter *Reference `bson:"encounter,omitempty" json:"encounter,omitempty" xml:"encounter,omitempty"`
	// An actor taking an active role in the event or activity that is logged.
	//
	// Comment: Several agents may be associated (i.e. have some responsibility for an activity) with an event or activity.
	//
	// For example, an activity may be initiated by one user for other users or involve more than one user. However, only one user may be the initiator/requestor for the activity.
	//
	// When a network are used in an event being recorded, there should be distinct agent elements for the known actors using the network. The agent with a network detail would be the responsible agent for use of that network.
	Agent []AuditEventAgent `bson:"agent" json:"agent" xml:"agent"`
	// The actor that is reporting the event.
	//
	// Comment: Events are reported by the actor that detected them. This may be one of the participating actors, but may also be different. The actor may be a human such as a medical-records clerk disclosing data manually, that clerk would be the source for the record of disclosure.
	Source AuditEventSource `bson:"source" json:"source" xml:"source"`
	// Specific instances of data or objects that have been accessed.
	//
	// Comment: Required unless the values for event identification, agent identification, and audit source identification are sufficient to document the entire auditable event. Because events may have more than one entity, this group can be a repeating set of values.
	Entity []AuditEventEntity `bson:"entity,omitempty" json:"entity,omitempty" xml:"entity,omitempty"`
}

type AuditEventOutcome struct {
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
	// Indicates whether the event succeeded or failed.
	//
	// Comment: In some cases a "success" may be partial, for example, an incomplete or interrupted transfer of a radiological study. For the purpose of establishing accountability, these distinctions are not relevant.
	Code Coding `bson:"code" json:"code" xml:"code"`
	// Additional details about the error.
	//
	// This may be a text description of the error or a system code that identifies the error.
	//
	// Comment: A human readable description of the error issue SHOULD be placed in details.text.
	Detail []CodeableConcept `bson:"detail,omitempty" json:"detail,omitempty" xml:"detail,omitempty"`
}

type AuditEventAgent struct {
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
	// The Functional Role of the user when performing the event.
	//
	// Comment: For example: assembler, author, prescriber, signer, investigator, etc.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The structural roles of the agent indicating the agent's competency.
	//
	// The security role enabling the agent with respect to the activity.
	//
	// Comment: For example: Chief-of-Radiology, Nurse, Physician, Medical-Student, etc.
	Role []CodeableConcept `bson:"role,omitempty" json:"role,omitempty" xml:"role,omitempty"`
	// Reference to who this agent is that was involved in the event.
	//
	// Comment: Where a User ID is available it will go into who.identifier. Where a name of the user (human readable) it will go into who.display.
	Who Reference `bson:"who" json:"who" xml:"who"`
	// Indicator that the user is or is not the requestor, or initiator, for the event being audited.
	//
	// Comment: There can only be one initiator. If the initiator is not clear, then do not choose any one agent as the initiator.
	Requestor *bool `bson:"requestor,omitempty" json:"requestor,omitempty" xml:"requestor,omitempty"`
	// Where the agent location is known, the agent location when the event occurred.
	Location *Reference `bson:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// Where the policy(ies) are known that authorized the agent participation in the event.
	//
	// Typically, a single activity may have multiple applicable policies, such as patient consent, guarantor funding, etc. The policy would also indicate the security token used.
	//
	// Comment: For example: Where an OAuth token authorizes, the unique identifier from the OAuth token is placed into the policy element Where a policy engine (e.g. XACML) holds policy logic, the unique policy identifier is placed into the policy element.
	Policy []string `bson:"policy,omitempty" json:"policy,omitempty" xml:"policy,omitempty"`
	// When the event utilizes a network there should be an agent describing the local system, and an agent describing remote system, with the network interface details.
	//
	// Comment: When remote network endpoint is known, another agent representing the remote agent would indicate the remote network endpoint used. Convention is to indicate data flowing from Source to Destination. The convention for Search, given data flows both ways (query parameters vs results), is to have the Source as the initiator of the Search Transaction, and the Destination the responder to the Search transaction.
	NetworkReference *Reference `bson:"networkReference,omitempty" json:"networkReference,omitempty" xml:"networkReference,omitempty"`
	NetworkUri       *string    `bson:"networkUri,omitempty" json:"networkUri,omitempty" xml:"networkUri,omitempty"`
	NetworkString    *string    `bson:"networkString,omitempty" json:"networkString,omitempty" xml:"networkString,omitempty"`
	// The authorization (e.g., PurposeOfUse) that was used during the event being recorded.
	//
	// Comment: Use AuditEvent.agent.authorization when you know that is specific to the agent, otherwise use AuditEvent.authorization. For example, during a machine-to-machine transfer it might not be obvious to the audit system who caused the event, but it does know why.
	Authorization []CodeableConcept `bson:"authorization,omitempty" json:"authorization,omitempty" xml:"authorization,omitempty"`
}

type AuditEventSource struct {
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
	// Logical source location within the healthcare enterprise network.
	//
	//  For example, a hospital or other provider location within a multi-entity provider group.
	Site *Reference `bson:"site,omitempty" json:"site,omitempty" xml:"site,omitempty"`
	// Identifier of the source where the event was detected.
	Observer Reference `bson:"observer" json:"observer" xml:"observer"`
	// Code specifying the type of source where event originated.
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
}

type AuditEventEntity struct {
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
	// Identifies a specific instance of the entity.
	//
	// The reference should be version specific. This is allowed to be a Parameters resource.
	//
	// Comment: Use .what.display when all you have is a string (e.g. ParticipantObjectName).
	What *Reference `bson:"what,omitempty" json:"what,omitempty" xml:"what,omitempty"`
	// Code representing the role the entity played in the event being audited.
	Role *CodeableConcept `bson:"role,omitempty" json:"role,omitempty" xml:"role,omitempty"`
	// Security labels for the identified entity.
	//
	// Comment: Copied from entity meta security tags.
	SecurityLabel []CodeableConcept `bson:"securityLabel,omitempty" json:"securityLabel,omitempty" xml:"securityLabel,omitempty"`
	// The query parameters for a query-type entities.
	//
	// Comment: The meaning and secondary-encoding of the content of base64 encoded blob is specific to the AuditEvent.type, AuditEvent.subtype, and AuditEvent.entity.role.  The base64 is a general-use and safe container for event specific data blobs regardless of the encoding used by the transaction being recorded.  An AuditEvent consuming application must understand the event it is consuming and the formats used by the event. For example, if auditing an Oracle network database access, the Oracle formats must be understood as they will be simply encoded in the base64binary blob.
	//
	// The DICOM AuditMessage schema does not support both .name and .query being populated.
	Query *string `bson:"query,omitempty" json:"query,omitempty" xml:"query,omitempty"`
	// Tagged value pairs for conveying additional information about the entity.
	Detail []AuditEventEntityDetail `bson:"detail,omitempty" json:"detail,omitempty" xml:"detail,omitempty"`
	// The entity is attributed to an agent to express the agent's responsibility for that entity in the activity.
	//
	// This is most used to indicate when persistence media (the entity) are used by an agent. For example when importing data from a device, the device would be described in an entity, and the user importing data from that media would be indicated as the entity.agent.
	//
	// Comment: A usecase where one AuditEvent.entity.agent is used where the Entity that was used in the creation/updating of a target resource, is not in the context of the same custodianship as the target resource, and thus the meaning of AuditEvent.entity.agent is to say that the entity referenced is managed elsewhere and that this Agent provided access to it.  This would be similar to where the Entity being referenced is managed outside FHIR, such as through HL7 V2, v3, or XDS. This might be where the Entity being referenced is managed in another FHIR resource server. Thus it explains the provenance of that Entity's use in the context of this AuditEvent activity.
	Agent []AuditEventAgent `bson:"agent,omitempty" json:"agent,omitempty"`
}

type AuditEventEntityDetail struct {
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
	// The type of extra detail provided in the value.
	Type CodeableConcept `bson:"type" json:"type" xml:"type"`
	// The  value of the extra detail.
	ValueQuantity        Quantity        `bson:"valueQuantity" json:"valueQuantity" xml:"valueQuantity"`
	ValueCodeableConcept CodeableConcept `bson:"valueCodeableConcept" json:"valueCodeableConcept" xml:"valueCodeableConcept"`
	ValueString          string          `bson:"valueString" json:"valueString" xml:"valueString"`
	ValueBoolean         bool            `bson:"valueBoolean" json:"valueBoolean" xml:"valueBoolean"`
	ValueInteger         int             `bson:"valueInteger" json:"valueInteger" xml:"valueInteger"`
	ValueRange           Range           `bson:"valueRange" json:"valueRange" xml:"valueRange"`
	ValueRatio           Ratio           `bson:"valueRatio" json:"valueRatio" xml:"valueRatio"`
	ValueTime            string          `bson:"valueTime" json:"valueTime" xml:"valueTime"`
	ValueDateTime        string          `bson:"valueDateTime" json:"valueDateTime" xml:"valueDateTime"`
	ValuePeriod          Period          `bson:"valuePeriod" json:"valuePeriod" xml:"valuePeriod"`
	ValueBase64Binary    string          `bson:"valueBase64Binary" json:"valueBase64Binary" xml:"valueBase64Binary"`
}

func (r AuditEvent) ResourceType() string {
	return "AuditEvent"
}

// MarshalJSON implements json.Marshaler
func (r AuditEvent) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		AuditEvent
		ResourceType string `json:"resourceType"`
	}{
		AuditEvent:   r,
		ResourceType: r.ResourceType(),
	})
}
