package fhires

import (
	"encoding/json/v2"

	decimal "github.com/shopspring/decimal"
)

// A selection of DICOM SOP instances and/or frames within a single Study and Series.
//
// This might include additional specifics such as an image region, an Observation UID or a Segmentation Number, allowing linkage to an Observation Resource or transferring this information along with the ImagingStudy Resource.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/ImagingSelection
type ImagingSelection struct {
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
	// A unique identifier assigned to this imaging selection.
	//
	// Comment: This is a business identifier, not a resource identifier -- see discussion under [Business Identifiers](resource.html#identifiers).
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The current state of the ImagingSelection resource.
	//
	// This is not the status of any ImagingStudy, ServiceRequest, or Task resources associated with the ImagingSelection.
	//
	// Comment: Unknown does not represent "other" - one of the defined statuses must apply.  Unknown is used when the authoring system is not sure what the current status is.
	Status ImagingSelectionStatus `bson:"status" json:"status" xml:"status"`
	// The patient, or group of patients, location, device, organization, procedure or practitioner this imaging selection is about and into whose or what record the imaging selection is placed.
	Subject *Reference `bson:"subject,omitempty" json:"subject,omitempty" xml:"subject,omitempty"`
	// The date and time this imaging selection was created.
	Issued *string `bson:"issued,omitempty" json:"issued,omitempty" xml:"issued,omitempty"`
	// Selector of the instances – human or machine.
	Performer []ImagingSelectionPerformer `bson:"performer,omitempty" json:"performer,omitempty" xml:"performer,omitempty"`
	// A list of the diagnostic requests that resulted in this imaging selection being performed.
	BasedOn []Reference `bson:"basedOn,omitempty" json:"basedOn,omitempty" xml:"basedOn,omitempty"`
	// Classifies the imaging selection.
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
	// Reason for referencing the selected content.
	//
	// Comment: *All* code-value and, if present, component.code-component.value pairs need to be taken into account to correctly understand the meaning of the observation.
	Code CodeableConcept `bson:"code" json:"code" xml:"code"`
	// The Study Instance UID for the DICOM Study from which the images were selected.
	//
	// Comment: See [DICOM PS3.3 C.7.2](http://dicom.nema.org/medical/dicom/current/output/chtml/part03/sect_C.7.2.html).
	StudyUid *string `bson:"studyUid,omitempty" json:"studyUid,omitempty" xml:"studyUid,omitempty"`
	// The imaging study from which the imaging selection is made.
	DerivedFrom []Reference `bson:"derivedFrom,omitempty" json:"derivedFrom,omitempty" xml:"derivedFrom,omitempty"`
	// The network service providing retrieval access to the selected images, frames, etc.
	//
	// See implementation notes for information about using DICOM endpoints.
	//
	// Comment: Typical endpoint types include DICOM WADO-RS, which is used to retrieve DICOM instances in native or rendered (e.g., JPG, PNG) formats using a RESTful API; DICOM WADO-URI, which can similarly retrieve native or rendered instances, except using an HTTP query-based approach.
	Endpoint []Reference `bson:"endpoint,omitempty" json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// The Series Instance UID for the DICOM Series from which the images were selected.
	//
	// Comment: See [DICOM PS3.3 C.7.3](http://dicom.nema.org/medical/dicom/current/output/chtml/part03/sect_C.7.3.html).
	SeriesUid *string `bson:"seriesUid,omitempty" json:"seriesUid,omitempty" xml:"seriesUid,omitempty"`
	// The Series Number for the DICOM Series from which the images were selected.
	//
	// Comment: See [DICOM PS3.3 C.7.3](http://dicom.nema.org/medical/dicom/current/output/chtml/part03/sect_C.7.3.html).
	SeriesNumber *int `bson:"seriesNumber,omitempty" json:"seriesNumber,omitempty" xml:"seriesNumber,omitempty"`
	// The Frame of Reference UID identifying the coordinate system that conveys spatial and/or temporal information for the selected images or frames.
	//
	// Comment: See [DICOM PS3.3 C.7.4.1](http://dicom.nema.org/medical/dicom/current/output/chtml/part03/sect_C.7.4.html).
	FrameOfReferenceUid *string `bson:"frameOfReferenceUid,omitempty" json:"frameOfReferenceUid,omitempty" xml:"frameOfReferenceUid,omitempty"`
	// The anatomic structures examined.
	//
	// See DICOM Part 16 Annex L (http://dicom.nema.org/medical/dicom/current/output/chtml/part16/chapter_L.html) for DICOM to SNOMED-CT mappings.
	BodySite *CodeableReference `bson:"bodySite,omitempty" json:"bodySite,omitempty" xml:"bodySite,omitempty"`
	// The actual focus of an observation when it is not the patient of record representing something or someone associated with the patient such as a spouse, parent, fetus, or donor.
	//
	// For example, fetus observations in a mother's record.  The focus of an observation could also be an existing condition,  an intervention, the subject's diet,  another observation of the subject,  or a body structure such as tumor or implanted device.   An example use case would be using the Observation resource to capture whether the mother is trained to change her child's tracheostomy tube. In this example, the child is the patient of record and the mother is the focus.
	//
	// Comment: An imaging selection may reference a DICOM resource that itself references other DICOM resources.
	//
	//        e.g. a presentation state references a set of source images or frames.
	Focus []Reference `bson:"focus,omitempty" json:"focus,omitempty" xml:"focus,omitempty"`
	// Each imaging selection includes one or more selected DICOM SOP instances.
	Instance []ImagingSelectionInstance `bson:"instance,omitempty" json:"instance,omitempty" xml:"instance,omitempty"`
}

type ImagingSelectionPerformer struct {
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
	// Distinguishes the type of involvement of the performer.
	Function *CodeableConcept `bson:"function,omitempty" json:"function,omitempty" xml:"function,omitempty"`
	// Author – human or machine.
	Actor *Reference `bson:"actor,omitempty" json:"actor,omitempty" xml:"actor,omitempty"`
}

type ImagingSelectionInstance struct {
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
	// The SOP Instance UID for the selected DICOM instance.
	//
	// Comment: See [DICOM PS3.3 C.12.1](http://dicom.nema.org/medical/dicom/current/output/chtml/part03/sect_C.12.html).
	Uid string `bson:"uid" json:"uid" xml:"uid"`
	// The Instance Number for the selected DICOM instance.
	//
	// Comment: See [DICOM PS3.3 C.7.6.1](https://dicom.nema.org/medical/dicom/current/output/chtml/part03/sect_C.7.6.html#sect_C.7.6.1).
	//
	//        Note: A multiframe instance has a single instance number with each frame identified by a frame number.
	Number *int `bson:"number,omitempty" json:"number,omitempty" xml:"number,omitempty"`
	// The SOP Class UID for the selected DICOM instance.
	//
	// Comment: See [DICOM PS3.3 C.12.1](http://dicom.nema.org/medical/dicom/current/output/chtml/part03/sect_C.12.html).
	SopClass *Coding `bson:"sopClass,omitempty" json:"sopClass,omitempty" xml:"sopClass,omitempty"`
	// Selected subset of the SOP Instance.
	//
	// The content and format of the subset item is determined by the SOP Class of the selected instance.
	//
	//        May be one of:
	//
	//        - A list of frame numbers selected from a multiframe SOP Instance.
	//
	//        - A list of Content Item Observation UID values selected from a DICOM SR or other structured document SOP Instance.
	//
	//        - A list of segment numbers selected from a segmentation SOP Instance.
	//
	//        - A list of Region of Interest (ROI) numbers selected from a radiotherapy structure set SOP Instance.
	Subset []string `bson:"subset,omitempty" json:"subset,omitempty" xml:"subset,omitempty"`
	// Each imaging selection instance or frame list might includes an image region, specified by a region type and a set of 2D coordinates.
	//
	//        If the parent imagingSelection.instance contains a subset element of type frame, the image region applies to all frames in the subset list.
	ImageRegion2D []ImagingSelectionInstanceImageRegion2D `bson:"imageRegion2D,omitempty" json:"imageRegion2D,omitempty" xml:"imageRegion2D,omitempty"`
	// Each imaging selection might includes a 3D image region, specified by a region type and a set of 3D coordinates.
	ImageRegion3D []ImagingSelectionInstanceImageRegion3D `bson:"imageRegion3D,omitempty" json:"imageRegion3D,omitempty" xml:"imageRegion3D,omitempty"`
}

type ImagingSelectionInstanceImageRegion2D struct {
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
	// Specifies the type of image region.
	//
	// Comment: See [DICOM PS3.3 C.10.5.1.2](https://dicom.nema.org/medical/dicom/current/output/chtml/part03/sect_C.10.5.html#sect_C.10.5.1.2).
	RegionType ImagingSelection2DGraphicType `bson:"regionType" json:"regionType" xml:"regionType"`
	// The coordinates describing the image region.
	//
	// Encoded as a set of (column, row) pairs that denote positions in the selected image / frames specified with sub-pixel resolution.
	//
	//        The origin at the TLHC of the TLHC pixel is 0.0\0.0, the BRHC of the TLHC pixel is 1.0\1.0, and the BRHC of the BRHC pixel is the number of columns\rows in the image / frames. The values must be within the range 0\0 to the number of columns\rows in the image / frames.
	//
	// Comment: For a description of how 2D coordinates are encoded, see [DICOM PS3.3 C.18.6](http://dicom.nema.org/medical/dicom/current/output/chtml/part03/sect_C.18.6.html).
	Coordinate []decimal.Decimal `bson:"coordinate" json:"coordinate" xml:"coordinate"`
}

type ImagingSelectionInstanceImageRegion3D struct {
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
	// Specifies the type of image region.
	//
	// Comment: See [DICOM PS3.3 C.18.9.1.2](http://dicom.nema.org/medical/dicom/current/output/chtml/part03/sect_C.18.9.html#sect_C.18.9.1.2).
	RegionType ImagingSelection3DGraphicType `bson:"regionType" json:"regionType" xml:"regionType"`
	// The coordinates describing the image region.
	//
	// Encoded as an ordered set of (x,y,z) triplets (in mm and may be negative) that define a region of interest in the patient-relative Reference Coordinate System defined by ImagingSelection.frameOfReferenceUid element.
	//
	// Comment: For a description of how 3D coordinates are encoded, see [DICOM PS3.3 C.18.9](http://dicom.nema.org/medical/dicom/current/output/chtml/part03/sect_C.18.9.html).
	Coordinate []decimal.Decimal `bson:"coordinate" json:"coordinate" xml:"coordinate"`
}

func (r ImagingSelection) ResourceType() string {
	return "ImagingSelection"
}

// MarshalJSON implements json.Marshaler
func (r ImagingSelection) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ImagingSelection
		ResourceType string `json:"resourceType"`
	}{
		ImagingSelection: r,
		ResourceType:     r.ResourceType(),
	})
}
