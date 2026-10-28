package fhires

import "encoding/json/v2"

// A sample to be used for analysis.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Specimen
type Specimen struct {
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
	// Id for specimen.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The identifier assigned by the lab when accessioning specimen(s).
	//
	// This is not necessarily the same as the specimen identifier, depending on local lab procedures.
	AccessionIdentifier *Identifier `bson:"accessionIdentifier,omitempty" json:"accessionIdentifier,omitempty" xml:"accessionIdentifier,omitempty"`
	// The availability of the specimen.
	//
	// Comment: This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid.
	Status *SpecimenStatus `bson:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// The kind of material that forms the specimen.
	//
	// Comment: The type can change the way that a specimen is handled and drives what kind of analyses can properly be performed on the specimen. It is frequently used in diagnostic work flow decision making systems.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Where the specimen came from.
	//
	// This may be from patient(s), from a location (e.g., the source of an environmental sample), or a sampling of a substance, a biologically-derived product, or a device.
	Subject *Reference `bson:"subject,omitempty" json:"subject,omitempty" xml:"subject,omitempty"`
	// Time when specimen is received by the testing laboratory for processing or testing.
	ReceivedTime *string `bson:"receivedTime,omitempty" json:"receivedTime,omitempty" xml:"receivedTime,omitempty"`
	// Reference to the parent (source) specimen which is used when the specimen was either derived from or a component of another specimen.
	//
	// Comment: The parent specimen could be the source from which the current specimen is derived by some processing step (e.g. an aliquot or isolate or extracted nucleic acids from clinical samples) or one of many specimens that were combined to create a pooled sample.
	Parent []Reference `bson:"parent,omitempty" json:"parent,omitempty" xml:"parent,omitempty"`
	// Details concerning a service request that required a specimen to be collected.
	//
	// Comment: The request may be explicit or implied such with a ServiceRequest that requires a blood draw.
	Request []Reference `bson:"request,omitempty" json:"request,omitempty" xml:"request,omitempty"`
	// This element signifies if the specimen is part of a group or pooled.
	Combined *SpecimenCombined `bson:"combined,omitempty" json:"combined,omitempty" xml:"combined,omitempty"`
	// The role or reason for the specimen in the testing workflow.
	Role []CodeableConcept `bson:"role,omitempty" json:"role,omitempty" xml:"role,omitempty"`
	// A physical feature or landmark on a specimen, highlighted for context by the collector of the specimen (e.g.
	//
	// surgeon), that identifies the type of feature as well as its meaning (e.g. the red ink indicating the resection margin of the right lobe of the excised prostate tissue or wire loop at radiologically suspected tumor location).
	Feature []SpecimenFeature `bson:"feature,omitempty" json:"feature,omitempty" xml:"feature,omitempty"`
	// Details concerning the specimen collection.
	Collection *SpecimenCollection `bson:"collection,omitempty" json:"collection,omitempty" xml:"collection,omitempty"`
	// Details concerning processing and processing steps for the specimen.
	Processing []SpecimenProcessing `bson:"processing,omitempty" json:"processing,omitempty" xml:"processing,omitempty"`
	// The container holding the specimen.
	//
	//  The recursive nature of containers; i.e. blood in tube in tray in rack is not addressed here.
	Container []SpecimenContainer `bson:"container,omitempty" json:"container,omitempty" xml:"container,omitempty"`
	// A mode or state of being that describes the nature of the specimen.
	//
	// Comment: Specimen condition is an observation made about the specimen.  It's a point-in-time assessment.  It can be used to assess its quality or appropriateness for a specific test.
	Condition []CodeableConcept `bson:"condition,omitempty" json:"condition,omitempty" xml:"condition,omitempty"`
	// To communicate any details or issues about the specimen or during the specimen collection.
	//
	// (for example: broken vial, sent with patient, frozen).
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
}

type SpecimenFeature struct {
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
	// The landmark or feature being highlighted.
	Type CodeableConcept `bson:"type" json:"type" xml:"type"`
	// Description of the feature of the specimen.
	Description string `bson:"description" json:"description" xml:"description"`
}

type SpecimenCollection struct {
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
	// Person who collected the specimen.
	Collector *Reference `bson:"collector,omitempty" json:"collector,omitempty" xml:"collector,omitempty"`
	// Time when specimen was collected from subject - the physiologically relevant time.
	CollectedDateTime *string `bson:"collectedDateTime,omitempty" json:"collectedDateTime,omitempty" xml:"collectedDateTime,omitempty"`
	CollectedPeriod   *Period `bson:"collectedPeriod,omitempty" json:"collectedPeriod,omitempty" xml:"collectedPeriod,omitempty"`
	// The span of time over which the collection of a specimen occurred.
	Duration *Duration `bson:"duration,omitempty" json:"duration,omitempty" xml:"duration,omitempty"`
	// The quantity of specimen collected; for instance the volume of a blood sample, or the physical measurement of an anatomic pathology sample.
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty" xml:"quantity,omitempty"`
	// A coded value specifying the technique that is used to perform the procedure.
	Method *CodeableConcept `bson:"method,omitempty" json:"method,omitempty" xml:"method,omitempty"`
	// A coded value specifying the technique that is used to perform the procedure.
	Device *CodeableReference `bson:"device,omitempty" json:"device,omitempty" xml:"device,omitempty"`
	// The procedure event during which the specimen was collected (e.g.
	//
	// the surgery leading to the collection of a pathology sample).
	Procedure *Reference `bson:"procedure,omitempty" json:"procedure,omitempty" xml:"procedure,omitempty"`
	// Anatomical location from which the specimen was collected (if subject is a patient).
	//
	// This is the target site.  This element is not used for environmental specimens.
	//
	// Comment: Reasons for using BodyStructure reference include: 1.) Need to identify a specific site instance (e.g. a specific mole), 2.) a single pre-coordinated coded concept for the body site does not exist, 3.) a desire to represent bodySite using multiple attributes (e.g. modifiers).
	BodySite *CodeableReference `bson:"bodySite,omitempty" json:"bodySite,omitempty" xml:"bodySite,omitempty"`
	// Abstinence or reduction from some or all food, drink, or both, for a period of time prior to sample collection.
	//
	// Comment: Representing fasting status using this element is preferred to representing it with an observation using a 'pre-coordinated code'  such as  LOINC 2005-7 (Calcium [Moles/​time] in 2 hour Urine --12 hours fasting), or  using  a component observation ` such as `Observation.component code`  = LOINC 49541-6 (Fasting status - Reported).
	FastingStatusCodeableConcept *CodeableConcept `bson:"fastingStatusCodeableConcept,omitempty" json:"fastingStatusCodeableConcept,omitempty" xml:"fastingStatusCodeableConcept,omitempty"`
	FastingStatusDuration        *Duration        `bson:"fastingStatusDuration,omitempty" json:"fastingStatusDuration,omitempty" xml:"fastingStatusDuration,omitempty"`
}

type SpecimenProcessing struct {
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
	// Textual description of procedure.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// A coded value specifying the method used to process the specimen.
	Method *CodeableConcept `bson:"method,omitempty" json:"method,omitempty" xml:"method,omitempty"`
	// Material used in the processing step.
	Additive []Reference `bson:"additive,omitempty" json:"additive,omitempty" xml:"additive,omitempty"`
	// A record of the time or period when the specimen processing occurred.
	//
	//  For example the time of sample fixation or the period of time the sample was in formalin.
	TimeDateTime *string `bson:"timeDateTime,omitempty" json:"timeDateTime,omitempty" xml:"timeDateTime,omitempty"`
	TimePeriod   *Period `bson:"timePeriod,omitempty" json:"timePeriod,omitempty" xml:"timePeriod,omitempty"`
}

type SpecimenContainer struct {
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
	// The device resource for the the container holding the specimen.
	//
	// If the container is in a holder then the referenced device will point to a parent device.
	Device Reference `bson:"device" json:"device" xml:"device"`
	// The location of the container holding the specimen.
	Location *Reference `bson:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// The quantity of specimen in the container; may be volume, dimensions, or other appropriate measurements, depending on the specimen type.
	SpecimenQuantity *Quantity `bson:"specimenQuantity,omitempty" json:"specimenQuantity,omitempty" xml:"specimenQuantity,omitempty"`
}

func (r Specimen) ResourceType() string {
	return "Specimen"
}

// MarshalJSON implements json.Marshaler
func (r Specimen) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Specimen
		ResourceType string `json:"resourceType"`
	}{
		ResourceType: r.ResourceType(),
		Specimen:     r,
	})
}
