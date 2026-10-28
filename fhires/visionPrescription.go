package fhires

import (
	"encoding/json"

	decimal "github.com/shopspring/decimal"
)

// An authorization for the provision of glasses and/or contact lenses to a patient.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/VisionPrescription
type VisionPrescription struct {
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
	// A unique identifier assigned to this vision prescription.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The status of the resource instance.
	//
	// Comment: This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid.
	Status FinancialResourceStatusCodes `bson:"status" json:"status" xml:"status"`
	// The date this resource was created.
	Created string `bson:"created" json:"created" xml:"created"`
	// A resource reference to the person to whom the vision prescription applies.
	Patient Reference `bson:"patient" json:"patient" xml:"patient"`
	// A reference to a resource that identifies the particular occurrence of contact between patient and health care provider during which the prescription was issued.
	Encounter *Reference `bson:"encounter,omitempty" json:"encounter,omitempty" xml:"encounter,omitempty"`
	// The date (and perhaps time) when the prescription was written.
	//
	// Comment: Jurisdictions determine the valid lifetime of a prescription. Typically vision prescriptions are valid for two years from the date written.
	DateWritten string `bson:"dateWritten" json:"dateWritten" xml:"dateWritten"`
	// The healthcare professional responsible for authorizing the prescription.
	Prescriber Reference `bson:"prescriber" json:"prescriber" xml:"prescriber"`
	// Contain the details of  the individual lens specifications and serves as the authorization for the fullfillment by certified professionals.
	LensSpecification []VisionPrescriptionLensSpecification `bson:"lensSpecification" json:"lensSpecification" xml:"lensSpecification"`
}

type VisionPrescriptionLensSpecification struct {
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
	// Identifies the type of vision correction product which is required for the patient.
	Product CodeableConcept `bson:"product" json:"product" xml:"product"`
	// The eye for which the lens specification applies.
	//
	// Comment: May also appear on the paper claim form or in the Medical Records as as OD (oculus dexter) for the right eye and OS (oculus sinister) for the left eye.
	Eye VisionEyes `bson:"eye" json:"eye" xml:"eye"`
	// Lens power measured in dioptres (0.25 units).
	//
	// Comment: The value is negative for near-sighted and positive for far sighted.
	//
	// Often insurance will not cover a lens with power between +75 and -75.
	Sphere *decimal.Decimal `bson:"sphere,omitempty" json:"sphere,omitempty" xml:"sphere,omitempty"`
	// Power adjustment for astigmatism measured in dioptres (0.25 units).
	Cylinder *decimal.Decimal `bson:"cylinder,omitempty" json:"cylinder,omitempty" xml:"cylinder,omitempty"`
	// Adjustment for astigmatism measured in integer degrees.
	//
	// Comment: The limits are +180 and -180 degrees.
	Axis *int `bson:"axis,omitempty" json:"axis,omitempty" xml:"axis,omitempty"`
	// Allows for adjustment on two axis.
	Prism []VisionPrescriptionLensSpecificationPrism `bson:"prism,omitempty" json:"prism,omitempty" xml:"prism,omitempty"`
	// Power adjustment for multifocal lenses measured in dioptres (0.25 units).
	Add *decimal.Decimal `bson:"add,omitempty" json:"add,omitempty" xml:"add,omitempty"`
	// Contact lens power measured in dioptres (0.25 units).
	Power *decimal.Decimal `bson:"power,omitempty" json:"power,omitempty" xml:"power,omitempty"`
	// Back curvature measured in millimetres.
	BackCurve *decimal.Decimal `bson:"backCurve,omitempty" json:"backCurve,omitempty" xml:"backCurve,omitempty"`
	// Contact lens diameter measured in millimetres.
	Diameter *decimal.Decimal `bson:"diameter,omitempty" json:"diameter,omitempty" xml:"diameter,omitempty"`
	// The recommended maximum wear period for the lens.
	Duration *Quantity `bson:"duration,omitempty" json:"duration,omitempty" xml:"duration,omitempty"`
	// Special color or pattern.
	Color *string `bson:"color,omitempty" json:"color,omitempty" xml:"color,omitempty"`
	// Brand recommendations or restrictions.
	Brand *string `bson:"brand,omitempty" json:"brand,omitempty" xml:"brand,omitempty"`
	// Notes for special requirements such as coatings and lens materials.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
}

type VisionPrescriptionLensSpecificationPrism struct {
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
	// Amount of prism to compensate for eye alignment in fractional units.
	Amount decimal.Decimal `bson:"amount" json:"amount" xml:"amount"`
	// The relative base, or reference lens edge, for the prism.
	Base VisionBase `bson:"base" json:"base" xml:"base"`
}

func (r VisionPrescription) ResourceType() string {
	return "VisionPrescription"
}

// MarshalJSON implements json.Marshaler
func (r VisionPrescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		VisionPrescription
		ResourceType string `json:"resourceType"`
	}{
		ResourceType:       r.ResourceType(),
		VisionPrescription: r,
	})
}
