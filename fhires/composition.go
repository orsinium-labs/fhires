package fhires

import v2 "encoding/json/v2"

// A set of healthcare-related information that is assembled together into a single logical package that provides a single coherent statement of meaning, establishes its own context and that has clinical attestation with regard to who is making the statement.
//
// A Composition defines the structure and narrative content necessary for a document. However, a Composition alone does not constitute a document. Rather, the Composition must be the first entry in a Bundle where Bundle.type=document, and any other resources referenced from Composition must be included as subsequent entries in the Bundle (for example Patient, Practitioner, Encounter, etc.).
//
// Purpose: To support documents, and also to capture the EN13606 notion of an attested commit to the patient EHR, and to allow a set of disparate resources at the information/engineering level to be gathered into a clinical statement.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Composition
type Composition struct {
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
	// An absolute URI that is used to identify this Composition when it is referenced in a specification, model, design or an instance; also called its canonical identifier.
	//
	// This SHOULD be globally unique and SHOULD be a literal address at which an authoritative instance of this Composition is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the Composition is stored on different servers.
	//
	// Comment: Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred. Multiple instances may share the same URL if they have a distinct version.
	URL *string `bson:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
	// A version-independent identifier for the Composition.
	//
	// This identifier stays constant as the composition is changed over time.
	//
	// Comment: Similar to ClinicalDocument/setId in CDA. See discussion in resource definition for how these relate.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// An explicitly assigned identifer of a variation of the content in the Composition.
	//
	// Comment: While each resource, including the composition itself, has its own version identifier, this is a formal identifier for the logical version of the Composition as a whole. It would remain constant if the resources were moved to a new server, and all got new individual resource versions, for example.
	Version *string `bson:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
	// The workflow/clinical status of this composition.
	//
	// The status is a marker for the clinical standing of the document.
	//
	// Comment: If a composition is marked as withdrawn, the compositions/documents in the series, or data from the composition or document series, should never be displayed to a user without being clearly marked as untrustworthy. The flag "entered-in-error" is why this element is labeled as a modifier of other elements.
	//
	// Some reporting work flows require that the original narrative of a final document never be altered; instead, only new narrative can be added. The composition resource has no explicit status for explicitly noting whether this business rule is in effect. This would be handled by an extension if required.
	Status CompositionStatus `bson:"status" json:"status" xml:"status"`
	// Specifies the particular kind of composition (e.g.
	//
	// History and Physical, Discharge Summary, Progress Note). This usually equates to the purpose of making the composition.
	//
	// Comment: For Composition type, LOINC is ubiquitous and strongly endorsed by HL7. Most implementation guides will require a specific LOINC code, or use LOINC as an extensible binding.
	Type CodeableConcept `bson:"type" json:"type" xml:"type"`
	// A categorization for the type of the composition - helps for indexing and searching.
	//
	// This may be implied by or derived from the code specified in the Composition Type.
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
	// Who or what the composition is about.
	//
	// The composition can be about a person, (patient or healthcare practitioner), a device (e.g. a machine) or even a group of subjects (such as a document about a herd of livestock, or a set of patients that share a common exposure).
	//
	// Comment: For clinical documents, this is usually the patient.
	Subject []Reference `bson:"subject,omitempty" json:"subject,omitempty" xml:"subject,omitempty"`
	// Describes the clinical encounter or type of care this documentation is associated with.
	Encounter *Reference `bson:"encounter,omitempty" json:"encounter,omitempty" xml:"encounter,omitempty"`
	// The composition editing time, when the composition was last logically changed by the author.
	//
	// Comment: The Last Modified Date on the composition may be after the date of the document was attested without being changed. This means that the date on an amended document is the date of the amendment, not the date of original authorship.
	Date string `bson:"date" json:"date" xml:"date"`
	// The content was developed with a focus and intent of supporting the contexts that are listed.
	//
	// These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate Composition instances.
	//
	// Comment: When multiple useContexts are specified, there is no expectation that all or any of the contexts apply.
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty" xml:"useContext,omitempty"`
	// Identifies who is responsible for the information in the composition, not necessarily who typed it in.
	Author []Reference `bson:"author" json:"author" xml:"author"`
	// A natural language name identifying the {{title}}.
	//
	// This name should be usable as an identifier for the module by machine processing applications such as code generation.
	//
	// Comment: The name is not expected to be globally unique. The name should be a simple alphanumeric type name to ensure that it is machine-processing friendly.
	Name *string `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Official human-readable label for the composition.
	//
	// Comment: For many compositions, the title is the same as the text or a display name of Composition.type (e.g. a "consultation" or "progress note"). Note that CDA does not make title mandatory, but there are no known cases where it is useful for title to be omitted, so it is mandatory here. Feedback on this requirement is welcome during the trial use period.
	Title string `bson:"title" json:"title" xml:"title"`
	// For any additional notes.
	//
	// Comment: A comment, general note or annotation not coded elsewhere.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// A participant who has attested to the accuracy of the composition/document.
	//
	// Comment: Only list each attester once.
	Attester []CompositionAttester `bson:"attester,omitempty" json:"attester,omitempty" xml:"attester,omitempty"`
	// Identifies the organization or group who is responsible for ongoing maintenance of and access to the composition/document information.
	//
	// Comment: This is useful when documents are derived from a composition - provides guidance for how to get the latest version of the document. This is optional because this is sometimes not known by the authoring system, and can be inferred by context. However, it is important that this information be known when working with a derived document, so providing a custodian is encouraged.
	Custodian *Reference `bson:"custodian,omitempty" json:"custodian,omitempty" xml:"custodian,omitempty"`
	// Relationships that this composition has with other compositions or documents that already exist.
	//
	// Comment: To reference (using relatesTo) a Bundle of type "document" or another Composition where you only have the identifier, use a resourceReference with an identifier.\
	//
	// When a Composition is an alternate representation of another type of document e.g. CDA, the other document could be referenced in a DocumentReference resource.
	//
	// Here is an example of using relatesTo to reference another Bundle of type "document" or Composition where you only have the identifier:
	//
	// ```
	//
	// "relatesTo": [{
	//
	//   "type": "replaces",
	//
	//   "resourceReference": {
	//
	//     "identifier": {
	//
	//       "system": "...",
	//
	//       "value": "..."
	//
	//     }
	//
	//   }
	//
	// }]
	//
	// ```
	RelatesTo []RelatedArtifact `bson:"relatesTo,omitempty" json:"relatesTo,omitempty" xml:"relatesTo,omitempty"`
	// The clinical service, such as a colonoscopy or an appendectomy, being documented.
	//
	// Comment: The event needs to be consistent with the type element, though can provide further information if desired.
	Event []CompositionEvent `bson:"event,omitempty" json:"event,omitempty" xml:"event,omitempty"`
	// The root of the sections that make up the composition.
	Section []CompositionSection `bson:"section,omitempty" json:"section,omitempty" xml:"section,omitempty"`
}

type CompositionAttester struct {
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
	// The type of attestation the authenticator offers.
	Mode CodeableConcept `bson:"mode" json:"mode" xml:"mode"`
	// When the composition was attested by the party.
	Time *string `bson:"time,omitempty" json:"time,omitempty" xml:"time,omitempty"`
	// Who attested the composition in the specified way.
	Party *Reference `bson:"party,omitempty" json:"party,omitempty" xml:"party,omitempty"`
}

type CompositionEvent struct {
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
	// The period of time covered by the documentation.
	//
	// There is no assertion that the documentation is a complete representation for this period, only that it documents events during this time.
	Period *Period `bson:"period,omitempty" json:"period,omitempty" xml:"period,omitempty"`
	// Represents the main clinical acts, such as a colonoscopy or an appendectomy, being documented.
	//
	// In some cases, the event is inherent in the typeCode, such as a "History and Physical Report" in which case the procedure being documented is necessarily a "History and Physical" act. The events may be included as a code or as a reference to an other resource.
	//
	// Comment: An event can further specialize the act inherent in the typeCode, such as where it is simply "Procedure Report" and the procedure was a "colonoscopy". If one or more events are included, they SHALL NOT conflict with the values inherent in the classCode, practiceSettingCode or typeCode, as such a conflict would create an ambiguous situation. This short list of codes is provided to be used as key words for certain types of queries.
	Detail []CodeableReference `bson:"detail,omitempty" json:"detail,omitempty" xml:"detail,omitempty"`
}

type CompositionSection struct {
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
	// The label for this particular section.
	//
	//  This will be part of the rendered content for the document, and is often used to build a table of contents.
	//
	// Comment: The title identifies the section for a human reader. The title must be consistent with the narrative of the resource that is the target of the section.content reference. Generally, sections SHOULD have titles, but in some documents, it is unnecessary or inappropriate. Typically, this is where a section has subsections that have their own adequately distinguishing title,  or documents that only have a single section. Most Implementation Guides will make section title to be a required element.
	Title *string `bson:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// A code identifying the kind of content contained within the section.
	//
	// This must be consistent with the section title.
	//
	// Comment: The code identifies the section for an automated processor of the document. This is particularly relevant when using profiles to control the structure of the document.
	//
	// If the section has content (instead of sub-sections), the section.code does not change the meaning or interpretation of the resource that is the content of the section in the comments for the section.code.
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// Identifies who is responsible for the information in this section, not necessarily who typed it in.
	Author []Reference `bson:"author,omitempty" json:"author,omitempty" xml:"author,omitempty"`
	// The actual focus of the section when it is not the subject of the composition, but instead represents something or someone associated with the subject such as (for a patient subject) a spouse, parent, fetus, or donor.
	//
	// If not focus is specified, the focus is assumed to be focus of the parent section, or, for a section in the Composition itself, the subject of the composition. Sections with a focus SHALL only include resources where the logical subject (patient, subject, focus, etc.) matches the section focus, or the resources have no logical subject (few resources).
	//
	// Comment: Typically, sections in a doument are about the subject of the document, whether that is a  patient, or group of patients, location, or device, or whatever. For some kind of documents, some sections actually contain data about related entities. Typical examples are  a section in a newborn discharge summary concerning the mother, or family history documents, with a section about each family member, though there are many other examples.
	Focus *Reference `bson:"focus,omitempty" json:"focus,omitempty" xml:"focus,omitempty"`
	// A human-readable narrative that contains the attested content of the section, used to represent the content of the resource to a human.
	//
	// The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative.
	//
	// Comment: Document profiles may define what content should be represented in the narrative to ensure clinical safety.
	Text *Narrative `bson:"text,omitempty" json:"text,omitempty" xml:"text,omitempty"`
	// Specifies the order applied to the items in the section entries.
	//
	// Comment: Applications SHOULD render ordered lists in the order provided, but MAY allow users to re-order based on their own preferences as well. If there is no order specified, the order is unknown, though there may still be some order.
	OrderedBy *CodeableConcept `bson:"orderedBy,omitempty" json:"orderedBy,omitempty" xml:"orderedBy,omitempty"`
	// A reference to the actual resource from which the narrative in the section is derived.
	//
	// Comment: If there are no entries in the list, an emptyReason SHOULD be provided.
	Entry []Reference `bson:"entry,omitempty" json:"entry,omitempty" xml:"entry,omitempty"`
	// If the section is empty, why the list is empty.
	//
	// An empty section typically has some text explaining the empty reason.
	//
	// Comment: The various reasons for an empty section make a significant interpretation to its interpretation. Note that this code is for use when the entire section content has been suppressed, and not for when individual items are omitted - implementers may consider using a text note or a flag on an entry in these cases.
	EmptyReason *CodeableConcept `bson:"emptyReason,omitempty" json:"emptyReason,omitempty" xml:"emptyReason,omitempty"`
	// A nested sub-section within this section.
	//
	// Comment: Nested sections are primarily used to help human readers navigate to particular portions of the document.
	Section []CompositionSection `bson:"section,omitempty" json:"section,omitempty"`
}

func (r Composition) ResourceType() string {
	return "Composition"
}

// MarshalJSON implements json.Marshaler
func (r Composition) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		Composition
		ResourceType string `json:"resourceType"`
	}{
		Composition:  r,
		ResourceType: r.ResourceType(),
	})
}
