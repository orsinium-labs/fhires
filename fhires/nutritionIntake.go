package fhires

import v2 "encoding/json/v2"

// A record of food or fluid that is being consumed by a patient.
//
//	A NutritionIntake may indicate that the patient may be consuming the food or fluid now or has consumed the food or fluid in the past.  The source of this information can be the patient, significant other (such as a family member or spouse), or a clinician.  A common scenario where this information is captured is during the history taking process during a patient visit or stay or through an app that tracks food or fluids consumed.   The consumption information may come from sources such as the patient's memory, from a nutrition label,  or from a clinician documenting observed intake.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/NutritionIntake
type NutritionIntake struct {
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
	// Identifiers associated with this Nutrition Intake that are defined by business processes and/or used to refer to it when a direct URL reference to the resource itself is not appropriate.
	//
	// They are business identifiers assigned to this resource by the performer or other systems and remain constant as the resource is updated and propagates from server to server.
	//
	// Comment: This is a business identifier, not a resource identifier.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Instantiates FHIR protocol or definition.
	InstantiatesCanonical []string `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty" xml:"instantiatesCanonical,omitempty"`
	// Instantiates external protocol or definition.
	InstantiatesUri []string `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty" xml:"instantiatesUri,omitempty"`
	// A plan, proposal or order that is fulfilled in whole or in part by this event.
	BasedOn []Reference `bson:"basedOn,omitempty" json:"basedOn,omitempty" xml:"basedOn,omitempty"`
	// A larger event of which this particular event is a component or step.
	PartOf []Reference `bson:"partOf,omitempty" json:"partOf,omitempty" xml:"partOf,omitempty"`
	// A code representing the patient or other source's judgment about the state of the intake that this assertion is about.
	//
	//  Generally, this will be active or completed.
	//
	// Comment: NutritionIntake is a statement at a point in time.  The status is only representative at the point when it was asserted.  The value set for NutritionIntake.status contains codes that assert the status of the consumption of the food or fluid by the patient (for example, stopped or on hold) as well as codes that assert the status of the Nutrition Intake itself (for example, entered in error).
	//
	// This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid.
	Status EventStatus `bson:"status" json:"status" xml:"status"`
	// Captures the reason for the current state of the NutritionIntake.
	//
	// Comment: This is generally only used for "exception" statuses such as "not-taken", "on-hold", "cancelled" or "entered-in-error". The reason for performing the event at all is captured in reasonCode, not here.
	StatusReason []CodeableConcept `bson:"statusReason,omitempty" json:"statusReason,omitempty" xml:"statusReason,omitempty"`
	// Overall type of nutrition intake.
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// The person, animal or group who is/was consuming the food or fluid.
	Subject Reference `bson:"subject" json:"subject" xml:"subject"`
	// The encounter that establishes the context for this NutritionIntake.
	Encounter *Reference `bson:"encounter,omitempty" json:"encounter,omitempty" xml:"encounter,omitempty"`
	// The interval of time during which it is being asserted that the patient is/was consuming the food or fluid.
	//
	// Comment: This attribute reflects the period over which the patient consumed the food fluid and is expected to be populated on the majority of NutritionIntake. If the food or fluid is still being taken and is expected to continue indefinitely at the time the usage is recorded, the "end" date will be omitted.  If the end date is known, then it is included as the "end date".  The date/time attribute supports a variety of dates - year, year/month and exact date.  If something more than this is required, this should be conveyed as text.
	OccurrenceDateTime *string `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty" xml:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *Period `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty" xml:"occurrencePeriod,omitempty"`
	// The date when the Nutrition Intake was asserted by the information source.
	Recorded *string `bson:"recorded,omitempty" json:"recorded,omitempty" xml:"recorded,omitempty"`
	// The person or organization that provided the information about the consumption of this food or fluid.
	//
	// Note: Use derivedFrom when a NutritionIntake is derived from other resources.
	ReportedBoolean   *bool      `bson:"reportedBoolean,omitempty" json:"reportedBoolean,omitempty" xml:"reportedBoolean,omitempty"`
	ReportedReference *Reference `bson:"reportedReference,omitempty" json:"reportedReference,omitempty" xml:"reportedReference,omitempty"`
	// What food or fluid product or item was consumed.
	ConsumedItem []NutritionIntakeConsumedItem `bson:"consumedItem" json:"consumedItem" xml:"consumedItem"`
	// Total nutrient amounts for the whole meal, product, serving, etc.
	//
	// Comment: Individual item nutrients are not currently included in the resource and will likely end up as a reference in nutritionProduct to represent the individual items.
	IngredientLabel []NutritionIntakeIngredientLabel `bson:"ingredientLabel,omitempty" json:"ingredientLabel,omitempty" xml:"ingredientLabel,omitempty"`
	// Who performed the intake and how they were involved.
	Performer []NutritionIntakePerformer `bson:"performer,omitempty" json:"performer,omitempty" xml:"performer,omitempty"`
	// Where the intake occurred.
	Location *Reference `bson:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// Allows linking the NutritionIntake to the underlying NutritionOrder, or to other information, such as AllergyIntolerance, that supports or is used to derive the NutritionIntake.
	//
	// Comment: Likely references would be to AllergyIntolerance,  Observation or QuestionnaireAnswers.
	DerivedFrom []Reference `bson:"derivedFrom,omitempty" json:"derivedFrom,omitempty" xml:"derivedFrom,omitempty"`
	// A reason, Condition or observation for why the food or fluid is /was consumed.
	Reason []CodeableReference `bson:"reason,omitempty" json:"reason,omitempty" xml:"reason,omitempty"`
	// Provides extra information about the Nutrition Intake that is not conveyed by the other attributes.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
}

type NutritionIntakeConsumedItem struct {
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
	// Indicates what a category of item that was consumed: e.g., food, fluid, enteral, etc.
	Type CodeableConcept `bson:"type" json:"type" xml:"type"`
	// Identifies the food or fluid product that was consumed.
	//
	// This is potentially a link to a resource representing the details of the food product (TBD) or a simple attribute carrying a code that identifies the food from a known list of foods.
	//
	// Comment: If only a code is specified, then it needs to be a code for a specific product. If more information is required, then the use of a new nutritionProduct resource is recommended.
	NutritionProduct CodeableReference `bson:"nutritionProduct" json:"nutritionProduct" xml:"nutritionProduct"`
	// Scheduled frequency of consumption.
	Schedule *Timing `bson:"schedule,omitempty" json:"schedule,omitempty" xml:"schedule,omitempty"`
	// Quantity of the specified food.
	Amount *Quantity `bson:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
	// Rate at which enteral feeding was administered.
	Rate *Quantity `bson:"rate,omitempty" json:"rate,omitempty" xml:"rate,omitempty"`
	// Indicator when a patient is in a setting where it is helpful to know if food was not consumed, such as it was refused, held (as in tube feedings), or otherwise not provided.
	//
	// If a consumption is being recorded from an app, such as MyFitnessPal, this indicator will likely not be used.
	NotConsumed *bool `bson:"notConsumed,omitempty" json:"notConsumed,omitempty" xml:"notConsumed,omitempty"`
	// Document the reason the food or fluid was not consumed, such as refused, held, etc.
	NotConsumedReason *CodeableConcept `bson:"notConsumedReason,omitempty" json:"notConsumedReason,omitempty" xml:"notConsumedReason,omitempty"`
}

type NutritionIntakeIngredientLabel struct {
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
	// Total nutrient consumed.
	//
	// This could be a macronutrient (protein, fat, carbohydrate), or a vitamin and mineral.
	Nutrient CodeableReference `bson:"nutrient" json:"nutrient" xml:"nutrient"`
	// Total amount of nutrient consumed.
	Amount Quantity `bson:"amount" json:"amount" xml:"amount"`
}

type NutritionIntakePerformer struct {
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
	// Type of performer.
	Function *CodeableConcept `bson:"function,omitempty" json:"function,omitempty" xml:"function,omitempty"`
	// Who performed the intake.
	Actor Reference `bson:"actor" json:"actor" xml:"actor"`
}

func (r NutritionIntake) ResourceType() string {
	return "NutritionIntake"
}

// MarshalJSON implements json.Marshaler
func (r NutritionIntake) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		NutritionIntake
		ResourceType string `json:"resourceType"`
	}{
		NutritionIntake: r,
		ResourceType:    r.ResourceType(),
	})
}
