package fhires

import v2 "encoding/json/v2"

// Representation of the content produced in a DICOM imaging study.
//
// A study comprises a set of series, each of which includes a set of Service-Object Pair Instances (SOP Instances - images or other data) acquired or produced in a common context.  A series is of only one modality (e.g. X-ray, CT, MR, ultrasound), but a study may have multiple series of different modalities.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/ImagingStudy
type ImagingStudy struct {
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
	// Identifiers for the ImagingStudy such as DICOM Study Instance UID.
	//
	// Comment: See discussion under [Imaging Study Implementation Notes](imagingstudy.html#notes) for encoding of DICOM Study Instance UID.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The current state of the ImagingStudy resource.
	//
	// This is not the status of any ServiceRequest or Task resources associated with the ImagingStudy.
	//
	// Comment: Unknown does not represent "other" - one of the defined statuses must apply.  Unknown is used when the authoring system is not sure what the current status is.
	Status ImagingStudyStatus `bson:"status" json:"status" xml:"status"`
	// A list of all the distinct values of series.modality.
	//
	// This may include both acquisition and non-acquisition modalities.
	Modality []CodeableConcept `bson:"modality,omitempty" json:"modality,omitempty" xml:"modality,omitempty"`
	// The subject, typically a patient, of the imaging study.
	//
	// Comment: QA phantoms can be recorded with a Device; multiple subjects (such as mice) can be recorded with a Group.
	Subject Reference `bson:"subject" json:"subject" xml:"subject"`
	// The healthcare event (e.g.
	//
	// a patient and healthcare provider interaction) during which this ImagingStudy is made.
	//
	// Comment: This will typically be the encounter the event occurred within, but some events may be initiated prior to or after the official completion of an encounter but still be tied to the context of the encounter (e.g. pre-admission test).
	Encounter *Reference `bson:"encounter,omitempty" json:"encounter,omitempty" xml:"encounter,omitempty"`
	// Date and time the study started.
	Started *string `bson:"started,omitempty" json:"started,omitempty" xml:"started,omitempty"`
	// A list of the diagnostic requests that resulted in this imaging study being performed.
	BasedOn []Reference `bson:"basedOn,omitempty" json:"basedOn,omitempty" xml:"basedOn,omitempty"`
	// A larger event of which this particular ImagingStudy is a component or step.
	//
	//  For example,  an ImagingStudy as part of a procedure.
	//
	// Comment: To link an ImagingStudy to an Encounter use `encounter`.
	PartOf []Reference `bson:"partOf,omitempty" json:"partOf,omitempty" xml:"partOf,omitempty"`
	// The requesting/referring physician.
	Referrer *Reference `bson:"referrer,omitempty" json:"referrer,omitempty" xml:"referrer,omitempty"`
	// The network service providing access (e.g., query, view, or retrieval) for the study.
	//
	// See implementation notes for information about using DICOM endpoints. A study-level endpoint applies to each series in the study, unless overridden by a series-level endpoint with the same Endpoint.connectionType.
	//
	// Comment: Typical endpoint types include DICOM WADO-RS, which is used to retrieve DICOM instances in native or rendered (e.g., JPG, PNG), formats using a RESTful API; DICOM WADO-URI, which can similarly retrieve native or rendered instances, except using an HTTP query-based approach; DICOM QIDO-RS, which allows RESTful query for DICOM information without retrieving the actual instances; or IHE Invoke Image Display (IID), which provides standard invocation of an imaging web viewer.
	Endpoint []Reference `bson:"endpoint,omitempty" json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// Number of Series in the Study.
	//
	// This value given may be larger than the number of series elements this Resource contains due to resource availability, security, or other factors. This element should be present if any series elements are present.
	NumberOfSeries *int `bson:"numberOfSeries,omitempty" json:"numberOfSeries,omitempty" xml:"numberOfSeries,omitempty"`
	// Number of SOP Instances in Study.
	//
	// This value given may be larger than the number of instance elements this resource contains due to resource availability, security, or other factors. This element should be present if any instance elements are present.
	NumberOfInstances *int `bson:"numberOfInstances,omitempty" json:"numberOfInstances,omitempty" xml:"numberOfInstances,omitempty"`
	// This field corresponds to the DICOM Procedure Code Sequence (0008,1032).
	//
	// This is different from the FHIR Procedure resource that may include the ImagingStudy.
	Procedure []CodeableReference `bson:"procedure,omitempty" json:"procedure,omitempty" xml:"procedure,omitempty"`
	// The principal physical location where the ImagingStudy was performed.
	Location *Reference `bson:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// Description of clinical condition indicating why the ImagingStudy was requested, and/or Indicates another resource whose existence justifies this Study.
	Reason []CodeableReference `bson:"reason,omitempty" json:"reason,omitempty" xml:"reason,omitempty"`
	// Per the recommended DICOM mapping, this element is derived from the Study Description attribute (0008,1030).
	//
	// Observations or findings about the imaging study should be recorded in another resource, e.g. Observation, and not in this element.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// The Imaging Manager description of the study.
	//
	// Institution-generated description or classification of the Study (component) performed.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Each study has one or more series of images or other content.
	Series []ImagingStudySeries `bson:"series,omitempty" json:"series,omitempty" xml:"series,omitempty"`
}

type ImagingStudySeries struct {
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
	// The DICOM Series Instance UID for the series.
	//
	// Comment: See [DICOM PS3.3 C.7.3](http://dicom.nema.org/medical/dicom/current/output/chtml/part03/sect_C.7.3.html).
	Uid string `bson:"uid" json:"uid" xml:"uid"`
	// The numeric identifier of this series in the study.
	Number *int `bson:"number,omitempty" json:"number,omitempty" xml:"number,omitempty"`
	// The distinct modality for this series.
	//
	// This may include both acquisition and non-acquisition modalities.
	Modality CodeableConcept `bson:"modality" json:"modality" xml:"modality"`
	// A description of the series.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Number of SOP Instances in the Study.
	//
	// The value given may be larger than the number of instance elements this resource contains due to resource availability, security, or other factors. This element should be present if any instance elements are present.
	NumberOfInstances *int `bson:"numberOfInstances,omitempty" json:"numberOfInstances,omitempty" xml:"numberOfInstances,omitempty"`
	// The network service providing access (e.g., query, view, or retrieval) for this series.
	//
	// See implementation notes for information about using DICOM endpoints. A series-level endpoint, if present, has precedence over a study-level endpoint with the same Endpoint.connectionType.
	//
	// Comment: Typical endpoint types include DICOM WADO-RS, which is used to retrieve DICOM instances in native or rendered (e.g., JPG, PNG) formats using a RESTful API; DICOM WADO-URI, which can similarly retrieve native or rendered instances, except using an HTTP query-based approach; and DICOM QIDO-RS, which allows RESTful query for DICOM information without retrieving the actual instances.
	Endpoint []Reference `bson:"endpoint,omitempty" json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// The anatomic structures examined.
	//
	// See DICOM Part 16 Annex L (http://dicom.nema.org/medical/dicom/current/output/chtml/part16/chapter_L.html) for DICOM to SNOMED-CT mappings. The bodySite may indicate the laterality of body part imaged; if so, it shall be consistent with any content of ImagingStudy.series.laterality.
	BodySite *CodeableReference `bson:"bodySite,omitempty" json:"bodySite,omitempty" xml:"bodySite,omitempty"`
	// The laterality of the (possibly paired) anatomic structures examined.
	//
	// E.g., the left knee, both lungs, or unpaired abdomen. If present, shall be consistent with any laterality information indicated in ImagingStudy.series.bodySite.
	Laterality *CodeableConcept `bson:"laterality,omitempty" json:"laterality,omitempty" xml:"laterality,omitempty"`
	// The specimen imaged, e.g., for whole slide imaging of a biopsy.
	Specimen []Reference `bson:"specimen,omitempty" json:"specimen,omitempty" xml:"specimen,omitempty"`
	// The date and time the series was started.
	Started *string `bson:"started,omitempty" json:"started,omitempty" xml:"started,omitempty"`
	// Indicates who or what performed the series and how they were involved.
	//
	// Comment: If the person who performed the series is not known, their Organization may be recorded. A patient, or related person, may be the performer, e.g. for patient-captured images.
	Performer []ImagingStudySeriesPerformer `bson:"performer,omitempty" json:"performer,omitempty" xml:"performer,omitempty"`
	// A single SOP instance within the series, e.g. an image, or presentation state.
	Instance []ImagingStudySeriesInstance `bson:"instance,omitempty" json:"instance,omitempty" xml:"instance,omitempty"`
}

type ImagingStudySeriesPerformer struct {
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
	// Distinguishes the type of involvement of the performer in the series.
	Function *CodeableConcept `bson:"function,omitempty" json:"function,omitempty" xml:"function,omitempty"`
	// Indicates who or what performed the series.
	Actor Reference `bson:"actor" json:"actor" xml:"actor"`
}

type ImagingStudySeriesInstance struct {
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
	// The DICOM SOP Instance UID for this image or other DICOM content.
	//
	// Comment: See  [DICOM PS3.3 C.12.1](http://dicom.nema.org/medical/dicom/current/output/chtml/part03/sect_C.12.html#sect_C.12.1).
	Uid string `bson:"uid" json:"uid" xml:"uid"`
	// DICOM instance  type.
	SopClass Coding `bson:"sopClass" json:"sopClass" xml:"sopClass"`
	// The number of instance in the series.
	Number *int `bson:"number,omitempty" json:"number,omitempty" xml:"number,omitempty"`
	// The description of the instance.
	//
	// Comment: Particularly for post-acquisition analytic objects, such as SR, presentation states, value mapping, etc.
	Title *string `bson:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
}

func (r ImagingStudy) ResourceType() string {
	return "ImagingStudy"
}

// MarshalJSON implements json.Marshaler
func (r ImagingStudy) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		ImagingStudy
		ResourceType string `json:"resourceType"`
	}{
		ImagingStudy: r,
		ResourceType: r.ResourceType(),
	})
}
