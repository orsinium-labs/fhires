package fhires

import v2 "encoding/json/v2"

// The EvidenceReport Resource is a specialized container for a collection of resources and codeable concepts, adapted to support compositions of Evidence, EvidenceVariable, and Citation resources and related concepts.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/EvidenceReport
type EvidenceReport struct {
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
	// An absolute URI that is used to identify this EvidenceReport when it is referenced in a specification, model, design or an instance; also called its canonical identifier.
	//
	// This SHOULD be globally unique and SHOULD be a literal address at which an authoritative instance of this summary is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the summary is stored on different servers.
	//
	// Comment: Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred. Multiple instances may share the same URL if they have a distinct version.
	URL *string `bson:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
	// The status of this summary. Enables tracking the life-cycle of the content.
	//
	// Comment: Allows filtering of summaries that are appropriate for use versus not.
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Status PublicationStatus `bson:"status" json:"status" xml:"status"`
	// The content was developed with a focus and intent of supporting the contexts that are listed.
	//
	// These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate evidence report instances.
	//
	// Comment: When multiple useContexts are specified, there is no expectation that all or any of the contexts apply.
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty" xml:"useContext,omitempty"`
	// A formal identifier that is used to identify this EvidenceReport when it is represented in other formats, or referenced in a specification, model, design or an instance.
	//
	// Comment: This element will contain unique identifiers that support de-duplication of EvidenceReports. This identifier can be valid for only one EvidenceReport resource.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// A formal identifier that is used to identify things closely related to this EvidenceReport.
	//
	// Comment: May include trial registry identifiers, e.g. NCT04372602 from clinicaltrials.gov. This identifier can be valid for multiple EvidenceReport resources.
	RelatedIdentifier []Identifier `bson:"relatedIdentifier,omitempty" json:"relatedIdentifier,omitempty" xml:"relatedIdentifier,omitempty"`
	// Citation Resource or display of suggested citation for this report.
	//
	// Comment: used for reports for which external citation is expected, such as use in support of scholarly publications.
	CiteAsReference *Reference `bson:"citeAsReference,omitempty" json:"citeAsReference,omitempty" xml:"citeAsReference,omitempty"`
	CiteAsMarkdown  *string    `bson:"citeAsMarkdown,omitempty" json:"citeAsMarkdown,omitempty" xml:"citeAsMarkdown,omitempty"`
	// Specifies the kind of report, such as grouping of classifiers, search results, or human-compiled expression.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Used for footnotes and annotations.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// Link, description or reference to artifact associated with the report.
	RelatedArtifact []RelatedArtifact `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty" xml:"relatedArtifact,omitempty"`
	// Specifies the subject or focus of the report. Answers "What is this report about?".
	//
	// Comment: May be used as an expression for search queries and search results
	Subject EvidenceReportSubject `bson:"subject" json:"subject" xml:"subject"`
	// The name of the organization or individual responsible for the release and ongoing maintenance of the evidence report.
	//
	// Comment: Usually an organization but may be an individual. The publisher (or steward) of the evidence report is the organization or individual primarily responsible for the maintenance and upkeep of the evidence report. This is not necessarily the same individual or organization that developed and initially authored the content. The publisher is the primary point of contact for questions or issues with the evidence report. This item SHOULD be populated unless the information is available from context.
	Publisher *string `bson:"publisher,omitempty" json:"publisher,omitempty" xml:"publisher,omitempty"`
	// Contact details to assist a user in finding and communicating with the publisher.
	//
	// Comment: Extensions to ContactDetail include: contactReference, contactAddress, and contributionTime (see [Clinical Reasoning Module](clinicalreasoning-module.html)).
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty" xml:"contact,omitempty"`
	// An individiual, organization, or device primarily involved in the creation and maintenance of the content.
	//
	// Comment: Extensions to ContactDetail include: contactReference, contactAddress, and contributionTime (see [Clinical Reasoning Module](clinicalreasoning-module.html)).
	Author []ContactDetail `bson:"author,omitempty" json:"author,omitempty" xml:"author,omitempty"`
	// An individiual, organization, or device primarily responsible for internal coherence of the content.
	//
	// Comment: Extensions to ContactDetail include: contactReference, contactAddress, and contributionTime (see [Clinical Reasoning Module](clinicalreasoning-module.html)).
	Editor []ContactDetail `bson:"editor,omitempty" json:"editor,omitempty" xml:"editor,omitempty"`
	// An individiual, organization, or device primarily responsible for review of some aspect of the content.
	//
	// Comment: Extensions to ContactDetail include: contactReference, contactAddress, and contributionTime (see [Clinical Reasoning Module](clinicalreasoning-module.html)).
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Reviewer []ContactDetail `bson:"reviewer,omitempty" json:"reviewer,omitempty" xml:"reviewer,omitempty"`
	// An individiual, organization, or device responsible for officially endorsing the content for use in some setting.
	//
	// Comment: Extensions to ContactDetail include: contactReference, contactAddress, and contributionTime (see [Clinical Reasoning Module](clinicalreasoning-module.html)).
	//
	// See guidance around (not) making local changes to elements [here](canonicalresource.html#localization).
	Endorser []ContactDetail `bson:"endorser,omitempty" json:"endorser,omitempty" xml:"endorser,omitempty"`
	// Relationships that this composition has with other compositions or documents that already exist.
	//
	// Comment: A document is a version specific composition.
	RelatesTo []EvidenceReportRelatesTo `bson:"relatesTo,omitempty" json:"relatesTo,omitempty" xml:"relatesTo,omitempty"`
	// The root of the sections that make up the composition.
	Section []EvidenceReportSection `bson:"section,omitempty" json:"section,omitempty" xml:"section,omitempty"`
}

type EvidenceReportSubject struct {
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
	// Characteristic.
	Characteristic []EvidenceReportSubjectCharacteristic `bson:"characteristic,omitempty" json:"characteristic,omitempty" xml:"characteristic,omitempty"`
	// Used for general notes and annotations not coded elsewhere.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
}

type EvidenceReportSubjectCharacteristic struct {
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
	// Characteristic code.
	//
	// Comment: Example 1 is a Citation. Example 2 is a type of outcome. Example 3 is a specific outcome.
	Code CodeableConcept `bson:"code" json:"code" xml:"code"`
	// Characteristic value.
	//
	// Comment: Example 1 is Citation #37. Example 2 is selecting clinical outcomes. Example 3 is 1-year mortality.
	ValueReference       Reference       `bson:"valueReference" json:"valueReference" xml:"valueReference"`
	ValueCodeableConcept CodeableConcept `bson:"valueCodeableConcept" json:"valueCodeableConcept" xml:"valueCodeableConcept"`
	ValueBoolean         bool            `bson:"valueBoolean" json:"valueBoolean" xml:"valueBoolean"`
	ValueQuantity        Quantity        `bson:"valueQuantity" json:"valueQuantity" xml:"valueQuantity"`
	ValueRange           Range           `bson:"valueRange" json:"valueRange" xml:"valueRange"`
	// Is used to express not the characteristic.
	Exclude *bool `bson:"exclude,omitempty" json:"exclude,omitempty" xml:"exclude,omitempty"`
	// Timeframe for the characteristic.
	Period *Period `bson:"period,omitempty" json:"period,omitempty" xml:"period,omitempty"`
}

type EvidenceReportRelatesTo struct {
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
	// The type of relationship that this composition has with anther composition or document.
	//
	// Comment: If this document appends another document, then the document cannot be fully understood without also accessing the referenced document.
	Code ReportRelationshipType `bson:"code" json:"code" xml:"code"`
	// The target composition/document of this relationship.
	Target EvidenceReportRelatesToTarget `bson:"target" json:"target" xml:"target"`
}

type EvidenceReportRelatesToTarget struct {
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
	// Target of the relationship URL.
	URL *string `bson:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
	// Target of the relationship Identifier.
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Target of the relationship Display.
	Display *string `bson:"display,omitempty" json:"display,omitempty" xml:"display,omitempty"`
	// Target of the relationship Resource reference.
	Resource *Reference `bson:"resource,omitempty" json:"resource,omitempty" xml:"resource,omitempty"`
}

type EvidenceReportSection struct {
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
	// Comment: The title identifies the section for a human reader. The title must be consistent with the narrative of the resource that is the target of the section.content reference. Generally, sections SHOULD have titles, but in some documents, it is unnecessary or inappropriate. Typically, this is where a section has subsections that have their own adequately distinguishing title,  or documents that only have a single section.
	Title *string `bson:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// A code identifying the kind of content contained within the section.
	//
	// This should be consistent with the section title.
	//
	// Comment: The code identifies the section for an automated processor of the document. This is particularly relevant when using profiles to control the structure of the document.
	//
	// If the section has content (instead of sub-sections), the section.code does not change the meaning or interpretation of the resource that is the content of the section in the comments for the section.code.
	Focus *CodeableConcept `bson:"focus,omitempty" json:"focus,omitempty" xml:"focus,omitempty"`
	// A definitional Resource identifying the kind of content contained within the section.
	//
	// This should be consistent with the section title.
	FocusReference *Reference `bson:"focusReference,omitempty" json:"focusReference,omitempty" xml:"focusReference,omitempty"`
	// Identifies who is responsible for the information in this section, not necessarily who typed it in.
	Author []Reference `bson:"author,omitempty" json:"author,omitempty" xml:"author,omitempty"`
	// A human-readable narrative that contains the attested content of the section, used to represent the content of the resource to a human.
	//
	// The narrative need not encode all the structured data, but is peferred to contain sufficient detail to make it acceptable for a human to just read the narrative.
	//
	// Comment: Document profiles may define what content should be represented in the narrative.
	Text *Narrative `bson:"text,omitempty" json:"text,omitempty" xml:"text,omitempty"`
	// How the entry list was prepared - whether it is a working list that is suitable for being maintained on an ongoing basis, or if it represents a snapshot of a list of items from another source, or whether it is a prepared list where items may be marked as added, modified or deleted.
	//
	// Comment: This element is labeled as a modifier because a change list must not be misunderstood as a complete list.
	Mode *ListMode `bson:"mode,omitempty" json:"mode,omitempty" xml:"mode,omitempty"`
	// Specifies the order applied to the items in the section entries.
	//
	// Comment: Applications SHOULD render ordered lists in the order provided, but MAY allow users to re-order based on their own preferences as well. If there is no order specified, the order is unknown, though there may still be some order.
	OrderedBy *CodeableConcept `bson:"orderedBy,omitempty" json:"orderedBy,omitempty" xml:"orderedBy,omitempty"`
	// Specifies any type of classification of the evidence report.
	EntryClassifier []CodeableConcept `bson:"entryClassifier,omitempty" json:"entryClassifier,omitempty" xml:"entryClassifier,omitempty"`
	// A reference to the actual resource from which the narrative in the section is derived.
	//
	// Comment: If there are no entries in the list, an emptyReason SHOULD be provided.
	EntryReference []Reference `bson:"entryReference,omitempty" json:"entryReference,omitempty" xml:"entryReference,omitempty"`
	// Quantity as content.
	EntryQuantity []Quantity `bson:"entryQuantity,omitempty" json:"entryQuantity,omitempty" xml:"entryQuantity,omitempty"`
	// If the section is empty, why the list is empty.
	//
	// An empty section typically has some text explaining the empty reason.
	//
	// Comment: The various reasons for an empty section make a significant interpretation to its interpretation. Note that this code is for use when the entire section content has been suppressed, and not for when individual items are omitted - implementers may consider using a text note or a flag on an entry in these cases.
	EmptyReason *CodeableConcept `bson:"emptyReason,omitempty" json:"emptyReason,omitempty" xml:"emptyReason,omitempty"`
	// A nested sub-section within this section.
	//
	// Comment: Nested sections are primarily used to help human readers navigate to particular portions of the document.
	Section []EvidenceReportSection `bson:"section,omitempty" json:"section,omitempty"`
}

func (r EvidenceReport) ResourceType() string {
	return "EvidenceReport"
}

// MarshalJSON implements json.Marshaler
func (r EvidenceReport) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		EvidenceReport
		ResourceType string `json:"resourceType"`
	}{
		EvidenceReport: r,
		ResourceType:   r.ResourceType(),
	})
}
