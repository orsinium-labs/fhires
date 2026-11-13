package fhires

import "encoding/xml"

// Representation of the content produced in a DICOM imaging study.
//
// A study comprises a set of series, each of which includes a set of Service-Object Pair Instances (SOP Instances - images or other data) acquired or produced in a common context.  A series is of only one modality (e.g. X-ray, CT, MR, ultrasound), but a study may have multiple series of different modalities.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/ImagingStudy
type ImagingStudy struct {
	DomainResource
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

// MarshalJSON implements [json.Marshaler]
func (r ImagingStudy) MarshalJSON() ([]byte, error) {
	type Wrapper ImagingStudy
	return marshalJSON(r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Marshaler].
func (r ImagingStudy) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Wrapper ImagingStudy
	return marshalXML(e, start, r.ResourceType(), Wrapper(r))
}

// UnmarshalXML implements [xml.Unmarshaler].
func (r *ImagingStudy) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Wrapper ImagingStudy
	w := Wrapper(*r)
	err := unmarshalXML(d, start, r.ResourceType(), &w)
	*r = ImagingStudy(w)
	return err
}
