package fhires

import "encoding/json/v2"

// A request to supply a diet, formula feeding (enteral) or oral nutritional supplement to a patient/resident.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/NutritionOrder
type NutritionOrder struct {
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
	// Identifiers assigned to this order by the order sender or by the order receiver.
	//
	// Comment: The Identifier.type element can be to indicate filler vs. placer if needed.  This is explained in further detail [here](servicerequest.html#notes).
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The URL pointing to a FHIR-defined protocol, guideline, orderset or other definition that is adhered to in whole or in part by this NutritionOrder.
	//
	// Comment: Note: This is a business identifier, not a resource identifier (see [discussion](resource.html#identifiers)).  It is best practice for the identifier to only appear on a single resource instance, however business practices may occasionally dictate that multiple resource instances with the same identifier can exist - possibly even with different resource types.  For example, multiple Patient and a Person resource instance might share the same social insurance number.
	InstantiatesCanonical []string `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty" xml:"instantiatesCanonical,omitempty"`
	// The URL pointing to an externally maintained protocol, guideline, orderset or other definition that is adhered to in whole or in part by this NutritionOrder.
	//
	// Comment: This might be an HTML page, PDF, etc. or could just be a non-resolvable URI identifier.
	InstantiatesUri []string `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty" xml:"instantiatesUri,omitempty"`
	// The URL pointing to a protocol, guideline, orderset or other definition that is adhered to in whole or in part by this NutritionOrder.
	Instantiates []string `bson:"instantiates,omitempty" json:"instantiates,omitempty" xml:"instantiates,omitempty"`
	// A plan or request that is fulfilled in whole or in part by this nutrition order.
	BasedOn []Reference `bson:"basedOn,omitempty" json:"basedOn,omitempty" xml:"basedOn,omitempty"`
	// A shared identifier common to all nutrition orders that were authorized more or less simultaneously by a single author, representing the composite or group identifier.
	GroupIdentifier *Identifier `bson:"groupIdentifier,omitempty" json:"groupIdentifier,omitempty" xml:"groupIdentifier,omitempty"`
	// The workflow status of the nutrition order/request.
	//
	// Comment: Typically the system placing the order sets the status to "requested". Thereafter, the order is maintained by the receiver that updates the status as the request is handled.  This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid.
	Status RequestStatus `bson:"status" json:"status" xml:"status"`
	// Indicates the level of authority/intentionality associated with the NutrionOrder and where the request fits into the workflow chain.
	//
	// Comment: When resources map to this element, they are free to define as many codes as necessary to cover their space and will map to "proposal, plan or order".  Can have multiple codes that map to one of these.  E.g. "original order", "encoded order", "reflex order" would all map to "order".  Expectation is that the set of codes is mutually exclusive or a strict all-encompassing hierarchy.
	Intent RequestIntent `bson:"intent" json:"intent" xml:"intent"`
	// Indicates how quickly the Nutrition Order should be addressed with respect to other        requests.
	Priority *RequestPriority `bson:"priority,omitempty" json:"priority,omitempty" xml:"priority,omitempty"`
	// The person or set of individuals who needs the nutrition order for an oral diet, nutritional supplement and/or enteral or formula feeding.
	Subject Reference `bson:"subject" json:"subject" xml:"subject"`
	// An encounter that provides additional information about the healthcare context in which this request is made.
	Encounter *Reference `bson:"encounter,omitempty" json:"encounter,omitempty" xml:"encounter,omitempty"`
	// Information to support fulfilling (i.e.
	//
	// dispensing or administering) of the nutrition,        for example, patient height and weight).
	SupportingInformation []Reference `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty" xml:"supportingInformation,omitempty"`
	// The date and time that this nutrition order was requested.
	DateTime string `bson:"dateTime" json:"dateTime" xml:"dateTime"`
	// The practitioner that holds legal responsibility for ordering the diet, nutritional supplement, or formula feedings.
	Orderer *Reference `bson:"orderer,omitempty" json:"orderer,omitempty" xml:"orderer,omitempty"`
	// The specified desired performer of the nutrition order.
	Performer []CodeableReference `bson:"performer,omitempty" json:"performer,omitempty" xml:"performer,omitempty"`
	// A link to a record of allergies or intolerances  which should be included in the nutrition order.
	//
	// Comment: Information on a patient's food allergies and intolerances to inform healthcare personnel about the type of foods that the patient shouldn't receive or consume.
	AllergyIntolerance []Reference `bson:"allergyIntolerance,omitempty" json:"allergyIntolerance,omitempty" xml:"allergyIntolerance,omitempty"`
	// This modifier is used to convey order-specific modifiers about the type of food that should be given.
	//
	// These can be derived from patient allergies, intolerances, or preferences such as Halal, Vegan or Kosher. This modifier applies to the entire nutrition order inclusive of the oral diet, nutritional supplements and enteral formula feedings.
	//
	// Comment: Information on a patient's food preferences that inform healthcare personnel about the food that the patient should receive or consume.
	FoodPreferenceModifier []CodeableConcept `bson:"foodPreferenceModifier,omitempty" json:"foodPreferenceModifier,omitempty" xml:"foodPreferenceModifier,omitempty"`
	// This modifier is used to convey Order-specific modifier about the type of oral food or oral fluids that should not be given.
	//
	// These can be derived from patient allergies, intolerances, or preferences such as No Red Meat, No Soy or No Wheat or  Gluten-Free.  While it should not be necessary to repeat allergy or intolerance information captured in the referenced AllergyIntolerance resource in the excludeFoodModifier, this element may be used to convey additional specificity related to foods that should be eliminated from the patient’s diet for any reason.  This modifier applies to the entire nutrition order inclusive of the oral diet, nutritional supplements and enteral formula feedings.
	//
	// Comment: Information on a patient's food allergies, intolerances and preferences to inform healthcare personnel about the type  of foods that the patient shouldn't receive or consume.
	ExcludeFoodModifier []CodeableConcept `bson:"excludeFoodModifier,omitempty" json:"excludeFoodModifier,omitempty" xml:"excludeFoodModifier,omitempty"`
	// This modifier is used to convey whether a food item is allowed to be brought in by the patient and/or family.
	//
	//  If set to true, indicates that the receiving system does not need to supply the food item.
	OutsideFoodAllowed *bool `bson:"outsideFoodAllowed,omitempty" json:"outsideFoodAllowed,omitempty" xml:"outsideFoodAllowed,omitempty"`
	// Diet given orally in contrast to enteral (tube) feeding.
	OralDiet *NutritionOrderOralDiet `bson:"oralDiet,omitempty" json:"oralDiet,omitempty" xml:"oralDiet,omitempty"`
	// Oral nutritional products given in order to add further nutritional value to the patient's diet.
	Supplement []NutritionOrderSupplement `bson:"supplement,omitempty" json:"supplement,omitempty" xml:"supplement,omitempty"`
	// Feeding provided through the gastrointestinal tract via a tube, catheter, or stoma that delivers nutrition distal to the oral cavity.
	EnteralFormula *NutritionOrderEnteralFormula `bson:"enteralFormula,omitempty" json:"enteralFormula,omitempty" xml:"enteralFormula,omitempty"`
	// Comments made about the {{title}} by the requester, performer, subject or other participants.
	//
	// Comment: This element SHALL NOT be used to supply free text instructions for the diet which are represented in the `.oralDiet.instruction`, `supplement.instruction`, or `enteralFormula.administrationInstruction` elements.
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
}

type NutritionOrderOralDiet struct {
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
	// The kind of diet or dietary restriction such as fiber restricted diet or diabetic diet.
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Schedule information for an oral diet.
	Schedule *NutritionOrderOralDietSchedule `bson:"schedule,omitempty" json:"schedule,omitempty" xml:"schedule,omitempty"`
	// Class that defines the quantity and type of nutrient modifications (for example carbohydrate, fiber or sodium) required for the oral diet.
	Nutrient []NutritionOrderOralDietNutrient `bson:"nutrient,omitempty" json:"nutrient,omitempty" xml:"nutrient,omitempty"`
	// Class that describes any texture modifications required for the patient to safely consume various types of solid foods.
	Texture []NutritionOrderOralDietTexture `bson:"texture,omitempty" json:"texture,omitempty" xml:"texture,omitempty"`
	// The required consistency (e.g.
	//
	// honey-thick, nectar-thick, thin, thickened.) of liquids or fluids served to the patient.
	FluidConsistencyType []CodeableConcept `bson:"fluidConsistencyType,omitempty" json:"fluidConsistencyType,omitempty" xml:"fluidConsistencyType,omitempty"`
	// Free text or additional instructions or information pertaining to the oral diet.
	//
	// Comment: Free text dosage instructions can be used for cases where the instructions are too complex to code.
	Instruction *string `bson:"instruction,omitempty" json:"instruction,omitempty" xml:"instruction,omitempty"`
}

type NutritionOrderOralDietSchedule struct {
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
	// The time period and frequency at which the diet should be given.
	//
	//  The diet should be given for the combination of all schedules if more than one schedule is present.
	Timing []Timing `bson:"timing,omitempty" json:"timing,omitempty" xml:"timing,omitempty"`
	// Indicates whether the product is only taken when needed within a specific dosing schedule.
	AsNeeded *bool `bson:"asNeeded,omitempty" json:"asNeeded,omitempty" xml:"asNeeded,omitempty"`
	// Indicates whether the product is only taken based on a precondition for taking the product.
	AsNeededFor *CodeableConcept `bson:"asNeededFor,omitempty" json:"asNeededFor,omitempty" xml:"asNeededFor,omitempty"`
}

type NutritionOrderOralDietNutrient struct {
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
	// The nutrient that is being modified such as carbohydrate or sodium.
	Modifier *CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty" xml:"modifier,omitempty"`
	// The quantity of the specified nutrient to include in diet.
	Amount *Quantity `bson:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
}

type NutritionOrderOralDietTexture struct {
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
	// Any texture modifications (for solid foods) that should be made, e.g.
	//
	// easy to chew, chopped, ground, and pureed.
	//
	// Comment: Coupled with the foodType (Meat).
	Modifier *CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty" xml:"modifier,omitempty"`
	// The food type(s) (e.g.
	//
	// meats, all foods)  that the texture modification applies to.  This could be all foods types.
	//
	// Comment: Coupled with the `texture.modifier`; could be (All Foods).
	FoodType *CodeableConcept `bson:"foodType,omitempty" json:"foodType,omitempty" xml:"foodType,omitempty"`
}

type NutritionOrderSupplement struct {
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
	// The kind of nutritional supplement product required such as a high protein or pediatric clear liquid supplement.
	Type *CodeableReference `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The product or brand name of the nutritional supplement such as "Acme Protein Shake".
	ProductName *string `bson:"productName,omitempty" json:"productName,omitempty" xml:"productName,omitempty"`
	// Schedule information for a supplement.
	Schedule *NutritionOrderSupplementSchedule `bson:"schedule,omitempty" json:"schedule,omitempty" xml:"schedule,omitempty"`
	// The amount of the nutritional supplement to be given.
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty" xml:"quantity,omitempty"`
	// Free text or additional instructions or information pertaining to the oral supplement.
	//
	// Comment: Free text dosage instructions can be used for cases where the instructions are too complex to code.
	Instruction *string `bson:"instruction,omitempty" json:"instruction,omitempty" xml:"instruction,omitempty"`
}

type NutritionOrderSupplementSchedule struct {
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
	// The time period and frequency at which the supplement should be given.
	//
	//  The supplement should be given for the combination of all schedules if more than one schedule is present.
	Timing []Timing `bson:"timing,omitempty" json:"timing,omitempty" xml:"timing,omitempty"`
	// Indicates whether the supplement is only taken when needed within a specific dosing schedule.
	AsNeeded *bool `bson:"asNeeded,omitempty" json:"asNeeded,omitempty" xml:"asNeeded,omitempty"`
	// Indicates whether the supplement is only taken based on a precondition for taking the supplement.
	AsNeededFor *CodeableConcept `bson:"asNeededFor,omitempty" json:"asNeededFor,omitempty" xml:"asNeededFor,omitempty"`
}

type NutritionOrderEnteralFormula struct {
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
	// The type of enteral or infant formula such as an adult standard formula with fiber or a soy-based infant formula.
	BaseFormulaType *CodeableReference `bson:"baseFormulaType,omitempty" json:"baseFormulaType,omitempty" xml:"baseFormulaType,omitempty"`
	// The product or brand name of the enteral or infant formula product such as "ACME Adult Standard Formula".
	BaseFormulaProductName *string `bson:"baseFormulaProductName,omitempty" json:"baseFormulaProductName,omitempty" xml:"baseFormulaProductName,omitempty"`
	// The intended type of device that is to be used for the administration of the enteral formula.
	DeliveryDevice []CodeableReference `bson:"deliveryDevice,omitempty" json:"deliveryDevice,omitempty" xml:"deliveryDevice,omitempty"`
	// Indicates modular components to be provided in addition or mixed with the base formula.
	Additive []NutritionOrderEnteralFormulaAdditive `bson:"additive,omitempty" json:"additive,omitempty" xml:"additive,omitempty"`
	// The amount of energy (calories) that the formula should provide per specified volume, typically per mL or fluid oz.
	//
	//  For example, an infant may require a formula that provides 24 calories per fluid ounce or an adult may require an enteral formula that provides 1.5 calorie/mL.
	CaloricDensity *Quantity `bson:"caloricDensity,omitempty" json:"caloricDensity,omitempty" xml:"caloricDensity,omitempty"`
	// The route or physiological path of administration into the patient's gastrointestinal  tract for purposes of providing the formula feeding, e.g.
	//
	// nasogastric tube.
	RouteOfAdministration *CodeableConcept `bson:"routeOfAdministration,omitempty" json:"routeOfAdministration,omitempty" xml:"routeOfAdministration,omitempty"`
	// Formula administration instructions as structured data.
	//
	//  This repeating structure allows for changing the administration rate or volume over time for both bolus and continuous feeding.  An example of this would be an instruction to increase the rate of continuous feeding every 2 hours.
	//
	// Comment: See implementation notes below for further discussion on how to order continuous vs bolus enteral feeding using this resource.
	Administration []NutritionOrderEnteralFormulaAdministration `bson:"administration,omitempty" json:"administration,omitempty" xml:"administration,omitempty"`
	// The maximum total quantity of formula that may be administered to a subject over the period of time, e.g.
	//
	// 1440 mL over 24 hours.
	MaxVolumeToDeliver *Quantity `bson:"maxVolumeToDeliver,omitempty" json:"maxVolumeToDeliver,omitempty" xml:"maxVolumeToDeliver,omitempty"`
	// Free text formula administration, feeding instructions or additional instructions or information.
	//
	// Comment: Free text dosage instructions can be used for cases where the instructions are too complex to code.
	AdministrationInstruction *string `bson:"administrationInstruction,omitempty" json:"administrationInstruction,omitempty" xml:"administrationInstruction,omitempty"`
}

type NutritionOrderEnteralFormulaAdditive struct {
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
	// Indicates the type of modular component such as protein, carbohydrate, fat or fiber to be provided in addition to or mixed with the base formula.
	Type *CodeableReference `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The product or brand name of the type of modular component to be added to the formula.
	ProductName *string `bson:"productName,omitempty" json:"productName,omitempty" xml:"productName,omitempty"`
	// The amount of additive to be given in addition or to be mixed in with the base formula.
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty" xml:"quantity,omitempty"`
}

type NutritionOrderEnteralFormulaAdministration struct {
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
	// Schedule information for an enteral formula.
	Schedule *NutritionOrderEnteralFormulaAdministrationSchedule `bson:"schedule,omitempty" json:"schedule,omitempty" xml:"schedule,omitempty"`
	// The volume of formula to provide to the patient per the specified administration schedule.
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty" xml:"quantity,omitempty"`
	// The rate of administration of formula via a feeding pump, e.g.
	//
	// 60 mL per hour, according to the specified schedule.
	//
	// Comment: Ratio is used when the quantity value in the denominator is not "1", otherwise use Quantity. For example, the Ratio datatype is used for "200 mL/4 hrs" versus the Quantity datatype for "50 mL/hr".
	RateQuantity *Quantity `bson:"rateQuantity,omitempty" json:"rateQuantity,omitempty" xml:"rateQuantity,omitempty"`
	RateRatio    *Ratio    `bson:"rateRatio,omitempty" json:"rateRatio,omitempty" xml:"rateRatio,omitempty"`
}

type NutritionOrderEnteralFormulaAdministrationSchedule struct {
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
	// The time period and frequency at which the enteral formula should be given.
	//
	//  The enteral formula should be given for the combination of all schedules if more than one schedule is present.
	Timing []Timing `bson:"timing,omitempty" json:"timing,omitempty" xml:"timing,omitempty"`
	// Indicates whether the enteral formula is only taken when needed within a specific dosing schedule.
	AsNeeded *bool `bson:"asNeeded,omitempty" json:"asNeeded,omitempty" xml:"asNeeded,omitempty"`
	// Indicates whether the enteral formula is only taken based on a precondition for taking the enteral formula.
	AsNeededFor *CodeableConcept `bson:"asNeededFor,omitempty" json:"asNeededFor,omitempty" xml:"asNeededFor,omitempty"`
}

func (r NutritionOrder) ResourceType() string {
	return "NutritionOrder"
}

// MarshalJSON implements json.Marshaler
func (r NutritionOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		NutritionOrder
		ResourceType string `json:"resourceType"`
	}{
		NutritionOrder: r,
		ResourceType:   r.ResourceType(),
	})
}
