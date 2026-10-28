package fhires

import (
	"encoding/json/v2"

	"github.com/shopspring/decimal"
)

// A structured set of questions and their answers.
//
// The questions are ordered and grouped into coherent subsets, corresponding to the structure of the grouping of the questionnaire being responded to.
//
// Purpose: To support structured, hierarchical reporting of data gathered using digital forms and other questionnaires.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/QuestionnaireResponse
type QuestionnaireResponse struct {
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
	// Business identifiers assigned to this questionnaire response by the performer and/or other systems.
	//
	//  These identifiers remain constant as the resource is updated and propagates from server to server.
	//
	// Comment: Note: This is a business identifier, not a resource identifier (see [discussion](resource.html#identifiers)).
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// A plan, proposal or order that is fulfilled in whole or in part by this questionnaire response.
	//
	//  For example, a ServiceRequest seeking an intake assessment or a decision support recommendation to assess for post-partum depression.
	BasedOn []Reference `bson:"basedOn,omitempty" json:"basedOn,omitempty" xml:"basedOn,omitempty"`
	// A procedure or observation that this questionnaire was performed as part of the execution of.
	//
	//  For example, the surgery a checklist was executed as part of.
	//
	// Comment: Not to be used to link an questionnaire response to an Encounter - use 'context' for that.
	//
	// Composition of questionnaire responses will be handled using the Assemble operation defined in the SDC IG.  For relationships to referrals, and other types of requests, use basedOn.
	PartOf []Reference `bson:"partOf,omitempty" json:"partOf,omitempty" xml:"partOf,omitempty"`
	// The Questionnaire that defines and organizes the questions for which answers are being provided.
	//
	// Comment: If a QuestionnaireResponse references a Questionnaire that can be resolved, then the QuestionnaireResponse structure must be consistent with the Questionnaire (i.e. questions must be organized into the same groups, nested questions must still be nested, etc.).  It is possible to have a QuestionnaireResponse whose 'questionnaire' element does not resolve.  It is also possible for the questionnaire element to not have a value but only extensions (e.g. conveying the title or identifier for the questionnaire).  This may happen for legacy data.  If there is no formally defined Questionnaire, it is undefined what the 'correct' values for the linkId elements should be and it is possible that linkIds might be inconsistent for QuestionnaireResponses for the same form if captured by distinct systems.
	Questionnaire string `bson:"questionnaire" json:"questionnaire" xml:"questionnaire"`
	// The current state of the questionnaire response.
	//
	// Comment: Unknown does not represent "other" - one of the defined statuses must apply.  Unknown is used when the authoring system is not sure what the current status is.
	Status QuestionnaireResponseStatus `bson:"status" json:"status" xml:"status"`
	// The subject of the questionnaire response.
	//
	//  This could be a patient, organization, practitioner, device, etc.  This is who/what the answers apply to, but is not necessarily the source of information.
	//
	// Comment: If the Questionnaire declared a subjectType, the resource pointed to by this element must be an instance of one of the listed types.
	Subject *Reference `bson:"subject,omitempty" json:"subject,omitempty" xml:"subject,omitempty"`
	// The Encounter during which this questionnaire response was created or to which the creation of this record is tightly associated.
	//
	// Comment: This will typically be the encounter the questionnaire response was created during, but some questionnaire responses may be initiated prior to or after the official completion of an encounter but still be tied to the context of the encounter (e.g. pre-admission forms).  A questionnaire that was initiated during an encounter but not fully completed during the encounter would still generally be associated with the encounter.
	Encounter *Reference `bson:"encounter,omitempty" json:"encounter,omitempty" xml:"encounter,omitempty"`
	// The date and/or time that this questionnaire response was last modified by the user - e.g.
	//
	// changing answers or revising status.
	//
	// Comment: May be different from the lastUpdateTime of the resource itself, because that reflects when the data was known to the server, not when the data was captured.
	//
	// This element is optional to allow for systems that might not know the value, however it SHOULD be populated if possible.
	Authored *string `bson:"authored,omitempty" json:"authored,omitempty" xml:"authored,omitempty"`
	// The individual or device that received the answers to the questions in the QuestionnaireResponse and recorded them in the system.
	//
	// Comment: Mapping a subject's answers to multiple choice options and determining what to put in the textual answer is a matter of interpretation. Authoring by device would indicate that some portion of the questionnaire had been auto-populated. Device should only be used if it directly determined the answers, not if it was merely used as a capture tool to record answers provided by others. In the latter case, information about the physical device, software, etc. would be captured using Provenance.
	Author *Reference `bson:"author,omitempty" json:"author,omitempty" xml:"author,omitempty"`
	// The individual or device that answered the questions about the subject.
	//
	// Comment: If not specified, no inference can be made about who provided the data. Device should only be used if it directly determined the answers, not if it was merely used as a capture tool to record answers provided by others. In the latter case, information about the physical device, software, etc. would be captured using Provenance.
	Source *Reference `bson:"source,omitempty" json:"source,omitempty" xml:"source,omitempty"`
	// A group or question item from the original questionnaire for which answers are provided.
	//
	// Comment: Groups cannot have answers and therefore must nest directly within item. When dealing with questions, nesting must occur within each answer because some questions may have multiple answers (and the nesting occurs for each answer).\nWhen dealing with repeating items, each group repetition will be handled by a separate item.  However, repeating questions are handled with a single question item and potentially multiple answers.
	Item []QuestionnaireResponseItem `bson:"item,omitempty" json:"item,omitempty" xml:"item,omitempty"`
}

type QuestionnaireResponseItem struct {
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
	// The item from the Questionnaire that corresponds to this item in the QuestionnaireResponse resource.
	LinkID string `bson:"linkId" json:"linkId" xml:"linkId"`
	// A reference to an [ElementDefinition](elementdefinition.html) that provides the details for the item.
	//
	// Comment: The ElementDefinition must be in a [StructureDefinition](structuredefinition.html#), and must have a fragment identifier that identifies the specific data element by its id (Element.id). E.g. http://hl7.org/fhir/StructureDefinition/Observation#Observation.value[x].
	//
	// There is no need for this element if the item pointed to by the linkId has a definition listed.
	Definition *string `bson:"definition,omitempty" json:"definition,omitempty" xml:"definition,omitempty"`
	// Text that is displayed above the contents of the group or as the text of the question being answered.
	//
	// Comment: The text for an item SHOULD be identical to the text from the corresponding Questionnaire.item. This can't be strictly enforced because it's possible for the Questionnaire to be updated subsequent to the QuestionnaireResponse having been created, however the intention is that the text in the QuestionnaireResponse reflects what the user saw when completing the Questionnaire.
	Text *string `bson:"text,omitempty" json:"text,omitempty" xml:"text,omitempty"`
	// The respondent's answer(s) to the question.
	//
	// Comment: The value is nested because we cannot have a repeating structure that has variable type.
	Answer []QuestionnaireResponseItemAnswer `bson:"answer,omitempty" json:"answer,omitempty" xml:"answer,omitempty"`
	// Sub-questions, sub-groups or display items nested beneath a group.
	//
	// Comment: Only used when nesting beneath a group - see item.answer.item for nesting beneath questions
	Item []QuestionnaireResponseItem `bson:"item,omitempty" json:"item,omitempty"`
}

type QuestionnaireResponseItemAnswer struct {
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
	// The answer (or one of the answers) provided by the respondent to the question.
	//
	// Comment: More complex structures (Attachment, Resource and Quantity) will typically be limited to electronic forms that can expose an appropriate user interface to capture the components and enforce the constraints of a complex data type.  Additional complex types can be introduced through extensions. Must match the datatype specified by Questionnaire.item.type in the corresponding Questionnaire.     Note that a question is answered using one of the possible choices defined with answerOption, answerValueSet or some other means and the answer has a complex data type, all elements within the answer in the QuestionnaireResponse **SHOULD** match the elements defined corresponding choice value in the Questionnaire.  However, it is possible that not all elements will be propagated.  Also, some systems might use language translations resulting in different displays.  Comparison of value to the values defined in the Questionnaire (whether by answerOption, answerValueSet or answerExpression) **SHALL NOT** pay attention to Coding.display, Reference.display, Quantity.unit unless those are the only elements present.  As well, systems are not required to check for a match on any extensions (e.g. ordinal values, translations, etc.).  Systems **MAY** enforce that if extensions such as ordinal values are present in both Questionnaire and QuestionnaireResponse, they match.
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
	// Nested groups and/or questions found within this particular answer.
	//
	// Comment: Only used when nesting beneath a question - see item.item for nesting beneath groups
	Item []QuestionnaireResponseItem `bson:"item,omitempty" json:"item,omitempty"`
}

func (r QuestionnaireResponse) ResourceType() string {
	return "QuestionnaireResponse"
}

// MarshalJSON implements json.Marshaler
func (r QuestionnaireResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		QuestionnaireResponse
		ResourceType string `json:"resourceType"`
	}{
		QuestionnaireResponse: r,
		ResourceType:          r.ResourceType(),
	})
}
