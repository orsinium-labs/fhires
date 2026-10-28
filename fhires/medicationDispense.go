package fhires

import "encoding/json/v2"

// Indicates that a medication product is to be or has been dispensed for a named person/patient.
//
//	This includes a description of the medication product (supply) provided and the instructions for administering the medication.  The medication dispense is the result of a pharmacy system responding to a medication order.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/MedicationDispense
type MedicationDispense struct {
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
	// Identifiers associated with this Medication Dispense that are defined by business processes and/or used to refer to it when a direct URL reference to the resource itself is not appropriate.
	//
	// They are business identifiers assigned to this resource by the performer or other systems and remain constant as the resource is updated and propagates from server to server.
	//
	// Comment: This is a business identifier, not a resource identifier.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// A plan that is fulfilled in whole or in part by this MedicationDispense.
	BasedOn []Reference `bson:"basedOn,omitempty" json:"basedOn,omitempty" xml:"basedOn,omitempty"`
	// The procedure or medication administration that triggered the dispense.
	//
	// Comment: While both a procedure and a medication administration may have triggered the dispense, but it is not expected that multiple procedures and/or multiple medication administrations would be triggers.
	PartOf []Reference `bson:"partOf,omitempty" json:"partOf,omitempty" xml:"partOf,omitempty"`
	// A code specifying the state of the set of dispense events.
	//
	// Comment: This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid.
	Status MedicationDispenseStatusCodes `bson:"status" json:"status" xml:"status"`
	// Indicates the reason why a dispense was not performed.
	NotPerformedReason *CodeableReference `bson:"notPerformedReason,omitempty" json:"notPerformedReason,omitempty" xml:"notPerformedReason,omitempty"`
	// The date (and maybe time) when the status of the dispense record changed.
	StatusChanged *string `bson:"statusChanged,omitempty" json:"statusChanged,omitempty" xml:"statusChanged,omitempty"`
	// Indicates the type of medication dispense (for example, drug classification like ATC, where meds would be administered, legal category of the medication.).
	//
	// Comment: The category can be used to include where the medication is expected to be consumed or other types of dispenses.  Invariants can be used to bind to different value sets when profiling to bind.
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
	// Identifies the medication supplied.
	//
	// This is either a link to a resource representing the details of the medication or a simple attribute carrying a code that identifies the medication from a known list of medications.
	//
	// Comment: If only a code is specified, then it needs to be a code for a specific product. If more information is required, then the use of the medication resource is recommended.  For example, if you require form or lot number, then you must reference the Medication resource.
	Medication CodeableReference `bson:"medication" json:"medication" xml:"medication"`
	// A link to a resource representing the person or the group to whom the medication will be given.
	//
	// Comment: SubstanceAdministration->subject->Patient.
	Subject Reference `bson:"subject" json:"subject" xml:"subject"`
	// The encounter that establishes the context for this event.
	Encounter *Reference `bson:"encounter,omitempty" json:"encounter,omitempty" xml:"encounter,omitempty"`
	// Additional information that supports the medication being dispensed.
	//
	//  For example, there may be requirements that a specific lab test has been completed prior to dispensing or the patient's weight at the time of dispensing is documented.
	SupportingInformation []Reference `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty" xml:"supportingInformation,omitempty"`
	// Indicates who or what performed the event.
	Performer []MedicationDispensePerformer `bson:"performer,omitempty" json:"performer,omitempty" xml:"performer,omitempty"`
	// The principal physical location where the dispense was performed.
	Location *Reference `bson:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// Indicates the medication order that is being dispensed against.
	//
	// Comment: Maps to basedOn in Event logical model.
	AuthorizingPrescription []Reference `bson:"authorizingPrescription,omitempty" json:"authorizingPrescription,omitempty" xml:"authorizingPrescription,omitempty"`
	// Indicates the type of dispensing event that is performed.
	//
	// For example, Trial Fill, Completion of Trial, Partial Fill, Emergency Fill, Samples, etc.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The amount of medication that has been dispensed. Includes unit of measure.
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty" xml:"quantity,omitempty"`
	// The amount of medication expressed as a timing amount.
	DaysSupply *Quantity `bson:"daysSupply,omitempty" json:"daysSupply,omitempty" xml:"daysSupply,omitempty"`
	// The date (and maybe time) when the dispense activity started if whenPrepared or whenHandedOver is not populated.
	Recorded *string `bson:"recorded,omitempty" json:"recorded,omitempty" xml:"recorded,omitempty"`
	// The time when the dispensed product was packaged and reviewed.
	WhenPrepared *string `bson:"whenPrepared,omitempty" json:"whenPrepared,omitempty" xml:"whenPrepared,omitempty"`
	// The time the dispensed product was provided to the patient or their representative.
	WhenHandedOver *string `bson:"whenHandedOver,omitempty" json:"whenHandedOver,omitempty" xml:"whenHandedOver,omitempty"`
	// Identification of the facility/location where the medication was/will be shipped to, as part of the dispense event.
	Destination *Reference `bson:"destination,omitempty" json:"destination,omitempty" xml:"destination,omitempty"`
	// Identifies the person who picked up the medication or the location of where the medication was delivered.
	//
	//  This will usually be a patient or their caregiver, but some cases exist where it can be a healthcare professional or a location.
	Receiver []Reference `bson:"receiver,omitempty" json:"receiver,omitempty" xml:"receiver,omitempty"`
	// Extra information about the dispense that could not be conveyed in the other attributes.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// The full representation of the dose of the medication included in all dosage instructions.
	//
	//  To be used when multiple dosage instructions are included to represent complex dosing such as increasing or tapering doses.
	//
	// Comment: The content of the renderedDosageInstructions must not be different than the dose represented in the dosageInstruction content.
	RenderedDosageInstruction *string `bson:"renderedDosageInstruction,omitempty" json:"renderedDosageInstruction,omitempty" xml:"renderedDosageInstruction,omitempty"`
	// Indicates how the medication is to be used by the patient.
	//
	// Comment: When the dose or rate is intended to change over the entire administration period (e.g. Tapering dose prescriptions), multiple instances of dosage instructions will need to be supplied to convey the different doses/rates.The pharmacist reviews the medication order prior to dispense and updates the dosageInstruction based on the actual product being dispensed.
	DosageInstruction []Dosage `bson:"dosageInstruction,omitempty" json:"dosageInstruction,omitempty" xml:"dosageInstruction,omitempty"`
	// Indicates whether or not substitution was made as part of the dispense.
	//
	//  In some cases, substitution will be expected but does not happen, in other cases substitution is not expected but does happen.  This block explains what substitution did or did not happen and why.  If nothing is specified, substitution was not done.
	Substitution *MedicationDispenseSubstitution `bson:"substitution,omitempty" json:"substitution,omitempty" xml:"substitution,omitempty"`
	// A summary of the events of interest that have occurred, such as when the dispense was verified.
	//
	// Comment: This might not include provenances for all versions of the request – only those deemed “relevant” or important. This SHALL NOT include the Provenance associated with this current version of the resource. (If that provenance is deemed to be a “relevant” change, it will need to be added as part of a later update. Until then, it can be queried directly as the Provenance that points to this version using _revinclude All Provenances should have some historical version of this Request as their subject.).
	EventHistory []Reference `bson:"eventHistory,omitempty" json:"eventHistory,omitempty" xml:"eventHistory,omitempty"`
}

type MedicationDispensePerformer struct {
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
	// Distinguishes the type of performer in the dispense.
	//
	//  For example, date enterer, packager, final checker.
	Function *CodeableConcept `bson:"function,omitempty" json:"function,omitempty" xml:"function,omitempty"`
	// The device, practitioner, etc.
	//
	// who performed the action.  It should be assumed that the actor is the dispenser of the medication.
	Actor Reference `bson:"actor" json:"actor" xml:"actor"`
}

type MedicationDispenseSubstitution struct {
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
	// True if the dispenser dispensed a different drug or product from what was prescribed.
	WasSubstituted bool `bson:"wasSubstituted" json:"wasSubstituted" xml:"wasSubstituted"`
	// A code signifying whether a different drug was dispensed from what was prescribed.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Indicates the reason for the substitution (or lack of substitution) from what was prescribed.
	Reason []CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty" xml:"reason,omitempty"`
	// The person or organization that has primary responsibility for the substitution.
	ResponsibleParty *Reference `bson:"responsibleParty,omitempty" json:"responsibleParty,omitempty" xml:"responsibleParty,omitempty"`
}

func (r MedicationDispense) ResourceType() string {
	return "MedicationDispense"
}

// MarshalJSON implements json.Marshaler
func (r MedicationDispense) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		MedicationDispense
		ResourceType string `json:"resourceType"`
	}{
		MedicationDispense: r,
		ResourceType:       r.ResourceType(),
	})
}
