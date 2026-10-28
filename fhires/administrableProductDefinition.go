package fhires

import v2 "encoding/json/v2"

// A medicinal product in the final form which is suitable for administering to a patient (after any mixing of multiple components, dissolution etc.
//
// has been performed).
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/AdministrableProductDefinition
type AdministrableProductDefinition struct {
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
	// An identifier for the administrable product.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The status of this administrable product. Enables tracking the life-cycle of the content.
	//
	// Comment: Allows filtering of administrable products that are appropriate for use versus not.
	Status PublicationStatus `bson:"status" json:"status" xml:"status"`
	// References a product from which one or more of the constituent parts of that product can be prepared and used as described by this administrable product.
	//
	//  If this administrable product describes the administration of a crushed tablet, the 'formOf' would be the product representing a distribution containing tablets and possibly also a cream.  This is distinct from the 'producedFrom' which refers to the specific components of the product that are used in this preparation, rather than the product as a whole.
	FormOf []Reference `bson:"formOf,omitempty" json:"formOf,omitempty" xml:"formOf,omitempty"`
	// The dose form of the final product after necessary reconstitution or processing.
	//
	// Contrasts to the manufactured dose form (see ManufacturedItemDefinition). If the manufactured form was 'powder for solution for injection', the administrable dose form could be 'solution for injection' (once mixed with another item having manufactured form 'solvent for solution for injection').
	AdministrableDoseForm *CodeableConcept `bson:"administrableDoseForm,omitempty" json:"administrableDoseForm,omitempty" xml:"administrableDoseForm,omitempty"`
	// The presentation type in which this item is given to a patient.
	//
	// e.g. for a spray - 'puff' (as in 'contains 100 mcg per puff'), or for a liquid - 'vial' (as in 'contains 5 ml per vial').
	UnitOfPresentation *CodeableConcept `bson:"unitOfPresentation,omitempty" json:"unitOfPresentation,omitempty" xml:"unitOfPresentation,omitempty"`
	// Indicates the specific manufactured items that are part of the 'formOf' product that are used in the preparation of this specific administrable form.
	//
	//  In some cases, an administrable form might use all of the items from the overall product (or there might only be one item), while in other cases, an administrable form might use only a subset of the items available in the overall product.  For example, an administrable form might involve combining a liquid and a powder available as part of an overall product, but not involve applying the also supplied cream.
	ProducedFrom []Reference `bson:"producedFrom,omitempty" json:"producedFrom,omitempty" xml:"producedFrom,omitempty"`
	// The ingredients of this administrable medicinal product.
	//
	// This is only needed if the ingredients are not specified either using ManufacturedItemDefiniton (via AdministrableProductDefinition.producedFrom) to state which component items are used to make this, or using by incoming references from the Ingredient resource, to state in detail which substances exist within this. This element allows a basic coded ingredient to be used.
	Ingredient []CodeableConcept `bson:"ingredient,omitempty" json:"ingredient,omitempty" xml:"ingredient,omitempty"`
	// A device that is integral to the medicinal product, in effect being considered as an "ingredient" of the medicinal product.
	//
	// This is not intended for devices that are just co-packaged.
	Device *Reference `bson:"device,omitempty" json:"device,omitempty" xml:"device,omitempty"`
	// A general description of the product, when in its final form, suitable for administration e.g.
	//
	// effervescent blue liquid, to be swallowed. Intended to be used when the other structured properties of this resource are insufficient or cannot be supported. It is not intended to duplicate information already carried elswehere.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Characteristics e.g. a product's onset of action.
	Property []AdministrableProductDefinitionProperty `bson:"property,omitempty" json:"property,omitempty" xml:"property,omitempty"`
	// The path by which the product is taken into or makes contact with the body.
	//
	// In some regions this is referred to as the licenced or approved route. RouteOfAdministration cannot be used when the 'formOf' product already uses MedicinalProductDefinition.route (and vice versa).
	RouteOfAdministration []AdministrableProductDefinitionRouteOfAdministration `bson:"routeOfAdministration" json:"routeOfAdministration" xml:"routeOfAdministration"`
}

type AdministrableProductDefinitionProperty struct {
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
	// A code expressing the type of characteristic.
	Type CodeableConcept `bson:"type" json:"type" xml:"type"`
	// A value for the characteristic.
	ValueCodeableConcept *CodeableConcept `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty" xml:"valueCodeableConcept,omitempty"`
	ValueQuantity        *Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty" xml:"valueQuantity,omitempty"`
	ValueDate            *string          `bson:"valueDate,omitempty" json:"valueDate,omitempty" xml:"valueDate,omitempty"`
	ValueBoolean         *bool            `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty" xml:"valueBoolean,omitempty"`
	ValueMarkdown        *string          `bson:"valueMarkdown,omitempty" json:"valueMarkdown,omitempty" xml:"valueMarkdown,omitempty"`
	ValueAttachment      *Attachment      `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty" xml:"valueAttachment,omitempty"`
	ValueReference       *Reference       `bson:"valueReference,omitempty" json:"valueReference,omitempty" xml:"valueReference,omitempty"`
	// The status of characteristic e.g. assigned or pending.
	Status *CodeableConcept `bson:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
}

type AdministrableProductDefinitionRouteOfAdministration struct {
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
	// Coded expression for the route.
	Code CodeableConcept `bson:"code" json:"code" xml:"code"`
	// The first dose (dose quantity) administered can be specified for the product, using a numerical value and its unit of measurement.
	FirstDose *Quantity `bson:"firstDose,omitempty" json:"firstDose,omitempty" xml:"firstDose,omitempty"`
	// The maximum single dose that can be administered, specified using a numerical value and its unit of measurement.
	MaxSingleDose *Quantity `bson:"maxSingleDose,omitempty" json:"maxSingleDose,omitempty" xml:"maxSingleDose,omitempty"`
	// The maximum dose per day (maximum dose quantity to be administered in any one 24-h period) that can be administered.
	MaxDosePerDay *Quantity `bson:"maxDosePerDay,omitempty" json:"maxDosePerDay,omitempty" xml:"maxDosePerDay,omitempty"`
	// The maximum dose per treatment period that can be administered.
	MaxDosePerTreatmentPeriod *Ratio `bson:"maxDosePerTreatmentPeriod,omitempty" json:"maxDosePerTreatmentPeriod,omitempty" xml:"maxDosePerTreatmentPeriod,omitempty"`
	// The maximum treatment period during which the product can be administered.
	MaxTreatmentPeriod *Duration `bson:"maxTreatmentPeriod,omitempty" json:"maxTreatmentPeriod,omitempty" xml:"maxTreatmentPeriod,omitempty"`
	// A species for which this route applies.
	TargetSpecies []AdministrableProductDefinitionRouteOfAdministrationTargetSpecies `bson:"targetSpecies,omitempty" json:"targetSpecies,omitempty" xml:"targetSpecies,omitempty"`
}

type AdministrableProductDefinitionRouteOfAdministrationTargetSpecies struct {
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
	// Coded expression for the species.
	Code CodeableConcept `bson:"code" json:"code" xml:"code"`
	// A species specific time during which consumption of animal product is not appropriate.
	WithdrawalPeriod []AdministrableProductDefinitionRouteOfAdministrationTargetSpeciesWithdrawalPeriod `bson:"withdrawalPeriod,omitempty" json:"withdrawalPeriod,omitempty" xml:"withdrawalPeriod,omitempty"`
}

type AdministrableProductDefinitionRouteOfAdministrationTargetSpeciesWithdrawalPeriod struct {
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
	// Coded expression for the type of tissue for which the withdrawal period applies, e.g.
	//
	// meat, milk.
	Tissue CodeableConcept `bson:"tissue" json:"tissue" xml:"tissue"`
	// A value for the time.
	Value Quantity `bson:"value" json:"value" xml:"value"`
	// Extra information about the withdrawal period.
	SupportingInformation *string `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty" xml:"supportingInformation,omitempty"`
}

func (r AdministrableProductDefinition) ResourceType() string {
	return "AdministrableProductDefinition"
}

// MarshalJSON implements json.Marshaler
func (r AdministrableProductDefinition) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		AdministrableProductDefinition
		ResourceType string `json:"resourceType"`
	}{
		AdministrableProductDefinition: r,
		ResourceType:                   r.ResourceType(),
	})
}
