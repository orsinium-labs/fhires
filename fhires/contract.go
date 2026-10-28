package fhires

import (
	"encoding/json/v2"

	"github.com/shopspring/decimal"
)

// Legally enforceable, formally recorded unilateral or bilateral directive i.e., a policy or agreement.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Contract
type Contract struct {
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
	// Unique identifier for this Contract or a derivative that references a Source Contract.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Canonical identifier for this contract, represented as a URI (globally unique).
	//
	// Comment: Used in a domain that uses a supplied contract repository.
	URL *string `bson:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
	// An edition identifier used for business purposes to label business significant variants.
	//
	// Comment: Note -  This is a business versionId, not a resource version id (see discussion at [Versioning](resource.html#versions)) Comments - There may be different contract instances that have the same identifier but different versions. The version can be appended to the url in a reference to allow a reference to a particular business version of the plan definition with the format [url]|[version]. The version SHOULD NOT contain a '#' - see [Business Version](resource.html#bv-format).
	Version *string `bson:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
	// The status of the resource instance.
	//
	// Comment: This element is labeled as a modifier because the status contains codes that mark the contract as not currently valid or active.
	Status *ContractResourceStatusCodes `bson:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Legal states of the formation of a legal instrument, which is a formally executed written document that can be formally attributed to its author, records and formally expresses a legally enforceable act, process, or contractual duty, obligation, or right, and therefore evidences that act, process, or agreement.
	LegalState *CodeableConcept `bson:"legalState,omitempty" json:"legalState,omitempty" xml:"legalState,omitempty"`
	// The URL pointing to a FHIR-defined Contract Definition that is adhered to in whole or part by this Contract.
	InstantiatesCanonical *Reference `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty" xml:"instantiatesCanonical,omitempty"`
	// The URL pointing to an externally maintained definition that is adhered to in whole or in part by this Contract.
	InstantiatesUri *string `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty" xml:"instantiatesUri,omitempty"`
	// The minimal content derived from the basal information source at a specific stage in its lifecycle.
	ContentDerivative *CodeableConcept `bson:"contentDerivative,omitempty" json:"contentDerivative,omitempty" xml:"contentDerivative,omitempty"`
	// When this  Contract was issued.
	Issued *string `bson:"issued,omitempty" json:"issued,omitempty" xml:"issued,omitempty"`
	// Relevant time or time-period when this Contract is applicable.
	Applies *Period `bson:"applies,omitempty" json:"applies,omitempty" xml:"applies,omitempty"`
	// Event resulting in discontinuation or termination of this Contract instance by one or more parties to the contract.
	ExpirationType *CodeableConcept `bson:"expirationType,omitempty" json:"expirationType,omitempty" xml:"expirationType,omitempty"`
	// The target entity impacted by or of interest to parties to the agreement.
	//
	// Comment: The Contract.subject is an entity that has some role with respect to the Contract.topic and Contract.topic.term, which is of focal interest to the parties to the contract and likely impacted in a significant way by the Contract.action/Contract.action.reason and the Contract.term.action/Contract.action.reason. In many cases, the Contract.subject is a Contract.signer if the subject is an adult; has a legal interest in the contract; and incompetent to participate in the contract agreement.
	Subject []Reference `bson:"subject,omitempty" json:"subject,omitempty" xml:"subject,omitempty"`
	// A formally or informally recognized grouping of people, principals, organizations, or jurisdictions formed for the purpose of achieving some form of collective action such as the promulgation, administration and enforcement of contracts and policies.
	Authority []Reference `bson:"authority,omitempty" json:"authority,omitempty" xml:"authority,omitempty"`
	// Recognized governance framework or system operating with a circumscribed scope in accordance with specified principles, policies, processes or procedures for managing rights, actions, or behaviors of parties or principals relative to resources.
	Domain []Reference `bson:"domain,omitempty" json:"domain,omitempty" xml:"domain,omitempty"`
	// Sites in which the contract is complied with,  exercised, or in force.
	Site []Reference `bson:"site,omitempty" json:"site,omitempty" xml:"site,omitempty"`
	// A natural language name identifying this Contract definition, derivative, or instance in any legal state.
	//
	// Provides additional information about its content. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	//
	// Comment: The name is not expected to be globally unique. The name should be a simple alphanumeric type name to ensure that it is machine-processing friendly.
	Name *string `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// A short, descriptive, user-friendly title for this Contract definition, derivative, or instance in any legal state.
	Title *string `bson:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// A more detailed or qualifying explanatory or alternate user-friendly title for this Contract definition, derivative, or instance in any legal state.
	Subtitle *string `bson:"subtitle,omitempty" json:"subtitle,omitempty" xml:"subtitle,omitempty"`
	// Alternative representation of the title for this Contract definition, derivative, or instance in any legal state., e.g., a domain specific contract number related to legislation.
	Alias []string `bson:"alias,omitempty" json:"alias,omitempty" xml:"alias,omitempty"`
	// The individual or organization that authored the Contract definition, derivative, or instance in any legal state.
	Author *Reference `bson:"author,omitempty" json:"author,omitempty" xml:"author,omitempty"`
	// A selector of legal concerns for this Contract definition, derivative, or instance in any legal state.
	Scope *CodeableConcept `bson:"scope,omitempty" json:"scope,omitempty" xml:"scope,omitempty"`
	// Narrows the range of legal concerns to focus on the achievement of specific contractual objectives.
	//
	// Comment: Contractual areas of concern are very broad. This element supports narrowing the area of concern to more specific term topics within this Contract. Given the wide range of contract topics, implementers need to be cognizant of the business use case for which they are designing a FHIR Contract, and narrowly specify the topic being represented with respect to the Contract.type and any specializing Contract.subtype. The same topic, e.g., an asset such as a good or service, such as a real property, medical supply, insurance, information, a procedure or employment, or a manner of conduct, such adherence to a privacy, trust, or security policy, may be the topic of multiple types of contracts. One way to determine the Contract.topic is to answer the question: "What is the overall objective of this legal instrument?". The Contract.topic is described with more detail by the terms of the Contract.
	TopicCodeableConcept *CodeableConcept `bson:"topicCodeableConcept,omitempty" json:"topicCodeableConcept,omitempty" xml:"topicCodeableConcept,omitempty"`
	TopicReference       *Reference       `bson:"topicReference,omitempty" json:"topicReference,omitempty" xml:"topicReference,omitempty"`
	// A high-level category for the legal instrument, whether constructed as a Contract definition, derivative, or instance in any legal state.
	//
	//  Provides additional information about its content within the context of the Contract's scope to distinguish the kinds of systems that would be interested in the contract.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Sub-category for the Contract that distinguishes the kinds of systems that would be interested in the Contract within the context of the Contract's scope.
	SubType []CodeableConcept `bson:"subType,omitempty" json:"subType,omitempty" xml:"subType,omitempty"`
	// Precusory content developed with a focus and intent of supporting the formation a Contract instance, which may be associated with and transformable into a Contract.
	ContentDefinition *ContractContentDefinition `bson:"contentDefinition,omitempty" json:"contentDefinition,omitempty" xml:"contentDefinition,omitempty"`
	// One or more Contract Provisions, which may be related and conveyed as a group, and may contain nested groups.
	Term []ContractTerm `bson:"term,omitempty" json:"term,omitempty" xml:"term,omitempty"`
	// Information that may be needed by/relevant to the performer in their execution of this term action.
	SupportingInfo []Reference `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty" xml:"supportingInfo,omitempty"`
	// Links to Provenance records for past versions of this Contract definition, derivative, or instance, which identify key state transitions or updates that are likely to be relevant to a user looking at the current version of the Contract.
	//
	//  The Provenance.entity indicates the target that was changed in the update (see [Provenance.entity](provenance-definitions.html#Provenance.entity)).
	RelevantHistory []Reference `bson:"relevantHistory,omitempty" json:"relevantHistory,omitempty" xml:"relevantHistory,omitempty"`
	// Parties with legal standing in the Contract, including the principal parties, the grantor(s) and grantee(s), which are any person or organization bound by the contract, and any ancillary parties, which facilitate the execution of the contract such as a notary or witness.
	//
	// Comment: Signers who are principal parties to the contract are bound by the Contract.activity related to the Contract.topic, and the Contract.term(s), which either extend or restrict the overall action on the topic by, for example, stipulating specific policies or obligations constraining actions, action reason, or agents with respect to some or all of the topic.For example, specifying how policies or obligations shall constrain actions and action reasons permitted or denied on all or a subset of the Contract.topic (e.g., all or a portion of property being transferred by the contract), agents (e.g., who can resell, assign interests, or alter the property being transferred by the contract), actions, and action reasons; or with respect to Contract.terms, stipulating, extending, or limiting the Contract.period of applicability or valuation of items under consideration.
	Signer []ContractSigner `bson:"signer,omitempty" json:"signer,omitempty" xml:"signer,omitempty"`
	// The "patient friendly language" versionof the Contract in whole or in parts.
	//
	// "Patient friendly language" means the representation of the Contract and Contract Provisions in a manner that is readily accessible and understandable by a layperson in accordance with best practices for communication styles that ensure that those agreeing to or signing the Contract understand the roles, actions, obligations, responsibilities, and implication of the agreement.
	Friendly []ContractFriendly `bson:"friendly,omitempty" json:"friendly,omitempty" xml:"friendly,omitempty"`
	// List of Legal expressions or representations of this Contract.
	Legal []ContractLegal `bson:"legal,omitempty" json:"legal,omitempty" xml:"legal,omitempty"`
	// List of Computable Policy Rule Language Representations of this Contract.
	Rule []ContractRule `bson:"rule,omitempty" json:"rule,omitempty" xml:"rule,omitempty"`
	// Legally binding Contract: This is the signed and legally recognized representation of the Contract, which is considered the "source of truth" and which would be the basis for legal action related to enforcement of this Contract.
	LegallyBindingAttachment *Attachment `bson:"legallyBindingAttachment,omitempty" json:"legallyBindingAttachment,omitempty" xml:"legallyBindingAttachment,omitempty"`
	LegallyBindingReference  *Reference  `bson:"legallyBindingReference,omitempty" json:"legallyBindingReference,omitempty" xml:"legallyBindingReference,omitempty"`
}

type ContractContentDefinition struct {
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
	// Precusory content structure and use, i.e., a boilerplate, template, application for a contract such as an insurance policy or benefits under a program, e.g., workers compensation.
	Type CodeableConcept `bson:"type" json:"type" xml:"type"`
	// Detailed Precusory content type.
	SubType *CodeableConcept `bson:"subType,omitempty" json:"subType,omitempty" xml:"subType,omitempty"`
	// The  individual or organization that published the Contract precursor content.
	Publisher *Reference `bson:"publisher,omitempty" json:"publisher,omitempty" xml:"publisher,omitempty"`
	// The date (and optionally time) when the contract was last significantly changed.
	//
	// The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the contract changes.
	PublicationDate *string `bson:"publicationDate,omitempty" json:"publicationDate,omitempty" xml:"publicationDate,omitempty"`
	// amended | appended | cancelled | disputed | entered-in-error | executable +.
	PublicationStatus ContractResourcePublicationStatusCodes `bson:"publicationStatus" json:"publicationStatus" xml:"publicationStatus"`
	// A copyright statement relating to Contract precursor content.
	//
	// Copyright statements are generally legal restrictions on the use and publishing of the Contract precursor content.
	Copyright *string `bson:"copyright,omitempty" json:"copyright,omitempty" xml:"copyright,omitempty"`
}

type ContractTerm struct {
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
	// Unique identifier for this particular Contract Provision.
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// When this Contract Provision was issued.
	Issued *string `bson:"issued,omitempty" json:"issued,omitempty" xml:"issued,omitempty"`
	// Relevant time or time-period when this Contract Provision is applicable.
	Applies *Period `bson:"applies,omitempty" json:"applies,omitempty" xml:"applies,omitempty"`
	// The entity that the term applies to.
	TopicCodeableConcept *CodeableConcept `bson:"topicCodeableConcept,omitempty" json:"topicCodeableConcept,omitempty" xml:"topicCodeableConcept,omitempty"`
	TopicReference       *Reference       `bson:"topicReference,omitempty" json:"topicReference,omitempty" xml:"topicReference,omitempty"`
	// A legal clause or condition contained within a contract that requires one or both parties to perform a particular requirement by some specified time or prevents one or both parties from performing a particular requirement by some specified time.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// A specialized legal clause or condition based on overarching contract type.
	SubType *CodeableConcept `bson:"subType,omitempty" json:"subType,omitempty" xml:"subType,omitempty"`
	// Statement of a provision in a policy or a contract.
	Text *string `bson:"text,omitempty" json:"text,omitempty" xml:"text,omitempty"`
	// Security labels that protect the handling of information about the term and its elements, which may be specifically identified.
	//
	// Comment: Within a Contract, a security label may apply to the one to many nested group of terms or to a term, whether inside a group or a singleton.  The security label on the entire set of term provision elements may be different from the security labels on a contained offer, asset, valuedItem, or data such as sensitive information, and must be the high water mark of all security labels within the term. Rationale is that a labelled term, which may be disaggregated from the Contract, and must persist the label on the term and on contained elements within other contexts. If more than one policy dictates a level of confidentiality of the term, then each applicable policy may be represented by a security label specific to its requirements.
	//
	// A Security Label is comprised of 1..1 confidentiality code and 0..* other security label codes (aka "privacy tags". A Security Label is metadata assigned (aka "tagged") to Resources and Resource elements, which specify the confidentiality and sensitivity of the Resource or element and their related attributes, the purposes for which tagged items may be collected, access, used or disclosed, any limitation on the type of entities authorized to engage in permitted purposes, and handling instructions that obligate or prohibit certain actions on the tagged items.
	SecurityLabel []ContractTermSecurityLabel `bson:"securityLabel,omitempty" json:"securityLabel,omitempty" xml:"securityLabel,omitempty"`
	// The matter of concern in the context of this provision of the agrement.
	Offer ContractTermOffer `bson:"offer" json:"offer" xml:"offer"`
	// Contract Term Asset List.
	Asset []ContractTermAsset `bson:"asset,omitempty" json:"asset,omitempty" xml:"asset,omitempty"`
	// An actor taking a role in an activity for which it can be assigned some degree of responsibility for the activity taking place.
	//
	// Comment: Several agents may be associated (i.e. has some responsibility for an activity) with an activity and vice-versa.For example, in cases of actions initiated by one user for other users, or in events that involve more than one user, hardware device, software, or system process. However, only one user may be the initiator/requestor for the event.
	Action []ContractTermAction `bson:"action,omitempty" json:"action,omitempty" xml:"action,omitempty"`
	// Nested group of Contract Provisions.
	Group []ContractTerm `bson:"group,omitempty" json:"group,omitempty"`
}

type ContractTermSecurityLabel struct {
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
	// Number used to link this term or term element to the applicable Security Label.
	Number []int `bson:"number,omitempty" json:"number,omitempty" xml:"number,omitempty"`
	// Security label privacy tag that specifies the level of confidentiality protection required for this term and/or term elements.
	Classification Coding `bson:"classification" json:"classification" xml:"classification"`
	// Security label privacy tag that specifies the applicable privacy and security policies governing this term and/or term elements.
	Category []Coding `bson:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
	// Security label privacy tag that specifies the manner in which term and/or term elements are to be protected.
	Control []Coding `bson:"control,omitempty" json:"control,omitempty" xml:"control,omitempty"`
}

type ContractTermOffer struct {
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
	// Unique identifier for this particular Contract Provision.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Offer Recipient.
	Party []ContractTermOfferParty `bson:"party,omitempty" json:"party,omitempty" xml:"party,omitempty"`
	// The owner of an asset has the residual control rights over the asset: the right to decide all usages of the asset in any way not inconsistent with a prior contract, custom, or law (Hart, 1995, p.
	//
	// 30).
	//
	// Comment: The Contract.topic may be an application for or offer of a policy or service (e.g., uri to a consent directive form or a health insurance policy), which becomes the Contract once accepted by both the grantor and grantee. The Contract Resource may function simply as the computable representation of the executed contract, which may be the attached to the Contract Resource as the “binding” or as the “friendly” electronic form.  For example, a Contract Resource may be automatically populated with the values expressed in a related QuestionnaireResponse. However, the Contract Resource may be considered the legally binding contract if it is the only “executed” form of this contract, and includes the signatures as *The Contract Resource may function as the computable representation of an application or offer in a pre-executed Contract if the grantor has not entered any values.  In this case, it is populated with values in a “legal” form of the application or offer or by the values in an associated Questionnaire.  If the grantor has filled in the legal form or the associated Questionnaire Response, then these values are used to populate a pre-executed Contract Resource.If the Contract.topic is considered an application or offer, then the policy is often required to be attached as the “legal” basis for the application to ensure “informed consent” to the contract, and that any discrepancy between the application and the policy are interpreted against the policy.  Implementers should check organizational and jurisdictional policies to determine the relationship among multiple representations of a contract pre- and post-execution.
	Topic *Reference `bson:"topic,omitempty" json:"topic,omitempty" xml:"topic,omitempty"`
	// Type of Contract Provision such as specific requirements, purposes for actions, obligations, prohibitions, e.g.
	//
	// life time maximum benefit.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Type of choice made by accepting party with respect to an offer made by an offeror/ grantee.
	Decision *CodeableConcept `bson:"decision,omitempty" json:"decision,omitempty" xml:"decision,omitempty"`
	// How the decision about a Contract was conveyed.
	DecisionMode []CodeableConcept `bson:"decisionMode,omitempty" json:"decisionMode,omitempty" xml:"decisionMode,omitempty"`
	// Response to offer text.
	Answer []ContractTermOfferAnswer `bson:"answer,omitempty" json:"answer,omitempty" xml:"answer,omitempty"`
	// Human readable form of this Contract Offer.
	Text *string `bson:"text,omitempty" json:"text,omitempty" xml:"text,omitempty"`
	// The id of the clause or question text of the offer in the referenced questionnaire/response.
	LinkID []string `bson:"linkId,omitempty" json:"linkId,omitempty" xml:"linkId,omitempty"`
	// Security labels that protects the offer.
	SecurityLabelNumber []int `bson:"securityLabelNumber,omitempty" json:"securityLabelNumber,omitempty" xml:"securityLabelNumber,omitempty"`
}

type ContractTermOfferParty struct {
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
	// Participant in the offer.
	Reference []Reference `bson:"reference" json:"reference" xml:"reference"`
	// How the party participates in the offer.
	Role CodeableConcept `bson:"role" json:"role" xml:"role"`
}

type ContractTermOfferAnswer struct {
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
	// Response to an offer clause or question text,  which enables selection of values to be agreed to, e.g., the period of participation, the date of occupancy of a rental, warranty duration, or whether biospecimen may be used for further research.
	ValueBoolean    bool            `bson:"valueBoolean" json:"valueBoolean" xml:"valueBoolean"`
	ValueDecimal    decimal.Decimal `bson:"valueDecimal" json:"valueDecimal" xml:"valueDecimal"`
	ValueInteger    int             `bson:"valueInteger" json:"valueInteger" xml:"valueInteger"`
	ValueDate       string          `bson:"valueDate" json:"valueDate" xml:"valueDate"`
	ValueDateTime   string          `bson:"valueDateTime" json:"valueDateTime" xml:"valueDateTime"`
	ValueTime       string          `bson:"valueTime" json:"valueTime" xml:"valueTime"`
	ValueString     string          `bson:"valueString" json:"valueString" xml:"valueString"`
	ValueUri        string          `bson:"valueUri" json:"valueUri" xml:"valueUri"`
	ValueAttachment Attachment      `bson:"valueAttachment" json:"valueAttachment" xml:"valueAttachment"`
	ValueCoding     Coding          `bson:"valueCoding" json:"valueCoding" xml:"valueCoding"`
	ValueQuantity   Quantity        `bson:"valueQuantity" json:"valueQuantity" xml:"valueQuantity"`
	ValueReference  Reference       `bson:"valueReference" json:"valueReference" xml:"valueReference"`
}

type ContractTermAsset struct {
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
	// Differentiates the kind of the asset .
	Scope *CodeableConcept `bson:"scope,omitempty" json:"scope,omitempty" xml:"scope,omitempty"`
	// Target entity type about which the term may be concerned.
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Associated entities.
	TypeReference []Reference `bson:"typeReference,omitempty" json:"typeReference,omitempty" xml:"typeReference,omitempty"`
	// May be a subtype or part of an offered asset.
	Subtype []CodeableConcept `bson:"subtype,omitempty" json:"subtype,omitempty" xml:"subtype,omitempty"`
	// Specifies the applicability of the term to an asset resource instance, and instances it refers to or instances that refer to it, and/or are owned by the offeree.
	Relationship *Coding `bson:"relationship,omitempty" json:"relationship,omitempty" xml:"relationship,omitempty"`
	// Circumstance of the asset.
	Context []ContractTermAssetContext `bson:"context,omitempty" json:"context,omitempty" xml:"context,omitempty"`
	// Description of the quality and completeness of the asset that may be a factor in its valuation.
	Condition *string `bson:"condition,omitempty" json:"condition,omitempty" xml:"condition,omitempty"`
	// Type of Asset availability for use or ownership.
	PeriodType []CodeableConcept `bson:"periodType,omitempty" json:"periodType,omitempty" xml:"periodType,omitempty"`
	// Asset relevant contractual time period.
	Period []Period `bson:"period,omitempty" json:"period,omitempty" xml:"period,omitempty"`
	// Time period of asset use.
	UsePeriod []Period `bson:"usePeriod,omitempty" json:"usePeriod,omitempty" xml:"usePeriod,omitempty"`
	// Clause or question text (Prose Object) concerning the asset in a linked form, such as a QuestionnaireResponse used in the formation of the contract.
	Text *string `bson:"text,omitempty" json:"text,omitempty" xml:"text,omitempty"`
	// Id [identifier??] of the clause or question text about the asset in the referenced form or QuestionnaireResponse.
	LinkID []string `bson:"linkId,omitempty" json:"linkId,omitempty" xml:"linkId,omitempty"`
	// Response to assets.
	Answer []ContractTermOfferAnswer `bson:"answer,omitempty" json:"answer,omitempty"`
	// Security labels that protects the asset.
	SecurityLabelNumber []int `bson:"securityLabelNumber,omitempty" json:"securityLabelNumber,omitempty" xml:"securityLabelNumber,omitempty"`
	// Contract Valued Item List.
	ValuedItem []ContractTermAssetValuedItem `bson:"valuedItem,omitempty" json:"valuedItem,omitempty" xml:"valuedItem,omitempty"`
}

type ContractTermAssetContext struct {
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
	// Asset context reference may include the creator, custodian, or owning Person or Organization (e.g., bank, repository),  location held, e.g., building,  jurisdiction.
	Reference *Reference `bson:"reference,omitempty" json:"reference,omitempty" xml:"reference,omitempty"`
	// Coded representation of the context generally or of the Referenced entity, such as the asset holder type or location.
	Code []CodeableConcept `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// Context description.
	Text *string `bson:"text,omitempty" json:"text,omitempty" xml:"text,omitempty"`
}

type ContractTermAssetValuedItem struct {
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
	// Specific type of Contract Valued Item that may be priced.
	EntityCodeableConcept *CodeableConcept `bson:"entityCodeableConcept,omitempty" json:"entityCodeableConcept,omitempty" xml:"entityCodeableConcept,omitempty"`
	EntityReference       *Reference       `bson:"entityReference,omitempty" json:"entityReference,omitempty" xml:"entityReference,omitempty"`
	// Identifies a Contract Valued Item instance.
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Indicates the time during which this Contract ValuedItem information is effective.
	EffectiveTime *string `bson:"effectiveTime,omitempty" json:"effectiveTime,omitempty" xml:"effectiveTime,omitempty"`
	// Specifies the units by which the Contract Valued Item is measured or counted, and quantifies the countable or measurable Contract Valued Item instances.
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty" xml:"quantity,omitempty"`
	// A Contract Valued Item unit valuation measure.
	UnitPrice *Money `bson:"unitPrice,omitempty" json:"unitPrice,omitempty" xml:"unitPrice,omitempty"`
	// A real number that represents a multiplier used in determining the overall value of the Contract Valued Item delivered.
	//
	// The concept of a Factor allows for a discount or surcharge multiplier to be applied to a monetary amount.
	Factor *decimal.Decimal `bson:"factor,omitempty" json:"factor,omitempty" xml:"factor,omitempty"`
	// An amount that expresses the weighting (based on difficulty, cost and/or resource intensiveness) associated with the Contract Valued Item delivered.
	//
	// The concept of Points allows for assignment of point values for a Contract Valued Item, such that a monetary amount can be assigned to each point.
	Points *decimal.Decimal `bson:"points,omitempty" json:"points,omitempty" xml:"points,omitempty"`
	// Expresses the product of the Contract Valued Item unitQuantity and the unitPriceAmt.
	//
	// For example, the formula: unit Quantity * unit Price (Cost per Point) * factor Number  * points = net Amount. Quantity, factor and points are assumed to be 1 if not supplied.
	Net *Money `bson:"net,omitempty" json:"net,omitempty" xml:"net,omitempty"`
	// Terms of valuation.
	Payment *string `bson:"payment,omitempty" json:"payment,omitempty" xml:"payment,omitempty"`
	// When payment is due.
	PaymentDate *string `bson:"paymentDate,omitempty" json:"paymentDate,omitempty" xml:"paymentDate,omitempty"`
	// Who will make payment.
	Responsible *Reference `bson:"responsible,omitempty" json:"responsible,omitempty" xml:"responsible,omitempty"`
	// Who will receive payment.
	Recipient *Reference `bson:"recipient,omitempty" json:"recipient,omitempty" xml:"recipient,omitempty"`
	// Id  of the clause or question text related to the context of this valuedItem in the referenced form or QuestionnaireResponse.
	LinkID []string `bson:"linkId,omitempty" json:"linkId,omitempty" xml:"linkId,omitempty"`
	// A set of security labels that define which terms are controlled by this condition.
	SecurityLabelNumber []int `bson:"securityLabelNumber,omitempty" json:"securityLabelNumber,omitempty" xml:"securityLabelNumber,omitempty"`
}

type ContractTermAction struct {
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
	// True if the term prohibits the  action.
	DoNotPerform *bool `bson:"doNotPerform,omitempty" json:"doNotPerform,omitempty" xml:"doNotPerform,omitempty"`
	// Activity or service obligation to be done or not done, performed or not performed, effectuated or not by this Contract term.
	Type CodeableConcept `bson:"type" json:"type" xml:"type"`
	// Entity of the action.
	Subject []ContractTermActionSubject `bson:"subject,omitempty" json:"subject,omitempty" xml:"subject,omitempty"`
	// Reason or purpose for the action stipulated by this Contract Provision.
	Intent CodeableConcept `bson:"intent" json:"intent" xml:"intent"`
	// Id [identifier??] of the clause or question text related to this action in the referenced form or QuestionnaireResponse.
	LinkID []string `bson:"linkId,omitempty" json:"linkId,omitempty" xml:"linkId,omitempty"`
	// Current state of the term action.
	Status CodeableConcept `bson:"status" json:"status" xml:"status"`
	// Encounter or Episode with primary association to the specified term activity.
	Context *Reference `bson:"context,omitempty" json:"context,omitempty" xml:"context,omitempty"`
	// Id [identifier??] of the clause or question text related to the requester of this action in the referenced form or QuestionnaireResponse.
	ContextLinkID []string `bson:"contextLinkId,omitempty" json:"contextLinkId,omitempty" xml:"contextLinkId,omitempty"`
	// When action happens.
	OccurrenceDateTime *string `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty" xml:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *Period `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty" xml:"occurrencePeriod,omitempty"`
	OccurrenceTiming   *Timing `bson:"occurrenceTiming,omitempty" json:"occurrenceTiming,omitempty" xml:"occurrenceTiming,omitempty"`
	// Who or what initiated the action and has responsibility for its activation.
	Requester []Reference `bson:"requester,omitempty" json:"requester,omitempty" xml:"requester,omitempty"`
	// Id [identifier??] of the clause or question text related to the requester of this action in the referenced form or QuestionnaireResponse.
	RequesterLinkID []string `bson:"requesterLinkId,omitempty" json:"requesterLinkId,omitempty" xml:"requesterLinkId,omitempty"`
	// The type of individual that is desired or required to perform or not perform the action.
	PerformerType []CodeableConcept `bson:"performerType,omitempty" json:"performerType,omitempty" xml:"performerType,omitempty"`
	// The type of role or competency of an individual desired or required to perform or not perform the action.
	PerformerRole *CodeableConcept `bson:"performerRole,omitempty" json:"performerRole,omitempty" xml:"performerRole,omitempty"`
	// Indicates who or what is being asked to perform (or not perform) the ction.
	Performer *Reference `bson:"performer,omitempty" json:"performer,omitempty" xml:"performer,omitempty"`
	// Id [identifier??] of the clause or question text related to the reason type or reference of this  action in the referenced form or QuestionnaireResponse.
	PerformerLinkID []string `bson:"performerLinkId,omitempty" json:"performerLinkId,omitempty" xml:"performerLinkId,omitempty"`
	// Rationale for the action to be performed or not performed.
	//
	// Describes why the action is permitted or prohibited. Either a coded concept, or another resource whose existence justifies permitting or not permitting this action.
	Reason []CodeableReference `bson:"reason,omitempty" json:"reason,omitempty" xml:"reason,omitempty"`
	// Id [identifier??] of the clause or question text related to the reason type or reference of this  action in the referenced form or QuestionnaireResponse.
	ReasonLinkID []string `bson:"reasonLinkId,omitempty" json:"reasonLinkId,omitempty" xml:"reasonLinkId,omitempty"`
	// Comments made about the term action made by the requester, performer, subject or other participants.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// Security labels that protects the action.
	SecurityLabelNumber []int `bson:"securityLabelNumber,omitempty" json:"securityLabelNumber,omitempty" xml:"securityLabelNumber,omitempty"`
}

type ContractTermActionSubject struct {
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
	// The entity the action is performed or not performed on or for.
	Reference []Reference `bson:"reference" json:"reference" xml:"reference"`
	// Role type of agent assigned roles in this Contract.
	Role *CodeableConcept `bson:"role,omitempty" json:"role,omitempty" xml:"role,omitempty"`
}

type ContractSigner struct {
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
	// Role of this Contract signer, e.g. notary, grantee.
	Type Coding `bson:"type" json:"type" xml:"type"`
	// Party which is a signator to this Contract.
	Party Reference `bson:"party" json:"party" xml:"party"`
	// Legally binding Contract DSIG signature contents in Base64.
	Signature []Signature `bson:"signature" json:"signature" xml:"signature"`
}

type ContractFriendly struct {
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
	// Human readable rendering of this Contract in a format and representation intended to enhance comprehension and ensure understandability.
	ContentAttachment Attachment `bson:"contentAttachment" json:"contentAttachment" xml:"contentAttachment"`
	ContentReference  Reference  `bson:"contentReference" json:"contentReference" xml:"contentReference"`
}

type ContractLegal struct {
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
	// Contract legal text in human renderable form.
	ContentAttachment Attachment `bson:"contentAttachment" json:"contentAttachment" xml:"contentAttachment"`
	ContentReference  Reference  `bson:"contentReference" json:"contentReference" xml:"contentReference"`
}

type ContractRule struct {
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
	// Computable Contract conveyed using a policy rule language (e.g. XACML, DKAL, SecPal).
	ContentAttachment Attachment `bson:"contentAttachment" json:"contentAttachment" xml:"contentAttachment"`
	ContentReference  Reference  `bson:"contentReference" json:"contentReference" xml:"contentReference"`
}

func (r Contract) ResourceType() string {
	return "Contract"
}

// MarshalJSON implements json.Marshaler
func (r Contract) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Contract
		ResourceType string `json:"resourceType"`
	}{
		Contract:     r,
		ResourceType: r.ResourceType(),
	})
}
