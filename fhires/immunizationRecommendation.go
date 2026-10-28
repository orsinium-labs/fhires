package fhires

import v2 "encoding/json/v2"

// A patient's point-in-time set of recommendations (i.e.
//
// forecasting) according to a published schedule with optional supporting justification.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/ImmunizationRecommendation
type ImmunizationRecommendation struct {
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
	// A unique identifier assigned to this particular recommendation record.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The patient the recommendation(s) are for.
	Patient Reference `bson:"patient" json:"patient" xml:"patient"`
	// The date the immunization recommendation(s) were created.
	Date string `bson:"date" json:"date" xml:"date"`
	// Indicates the authority who published the protocol (e.g. ACIP).
	Authority *Reference `bson:"authority,omitempty" json:"authority,omitempty" xml:"authority,omitempty"`
	// Vaccine administration recommendations.
	//
	// Comment: A given instance of the .recommendation backbone element should correspond to a single recommended administration.
	Recommendation []ImmunizationRecommendationRecommendation `bson:"recommendation" json:"recommendation" xml:"recommendation"`
}

type ImmunizationRecommendationRecommendation struct {
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
	// Vaccine(s) or vaccine group that pertain to the recommendation.
	VaccineCode []CodeableConcept `bson:"vaccineCode,omitempty" json:"vaccineCode,omitempty" xml:"vaccineCode,omitempty"`
	// The targeted disease for the recommendation.
	//
	// Comment: A given instance of the .recommendation backbone element should correspond to a single recommended administration but sometimes that administration (e.g. MMR) will encompass multiple diseases, thus the targetDisease element is allowed to repeat.
	TargetDisease []CodeableConcept `bson:"targetDisease,omitempty" json:"targetDisease,omitempty" xml:"targetDisease,omitempty"`
	// Vaccine(s) which should not be used to fulfill the recommendation.
	ContraindicatedVaccineCode []CodeableConcept `bson:"contraindicatedVaccineCode,omitempty" json:"contraindicatedVaccineCode,omitempty" xml:"contraindicatedVaccineCode,omitempty"`
	// Indicates the patient status with respect to the path to immunity for the target disease.
	ForecastStatus CodeableConcept `bson:"forecastStatus" json:"forecastStatus" xml:"forecastStatus"`
	// The reason for the assigned forecast status.
	ForecastReason []CodeableConcept `bson:"forecastReason,omitempty" json:"forecastReason,omitempty" xml:"forecastReason,omitempty"`
	// Vaccine date recommendations.
	//
	//  For example, earliest date to administer, latest date to administer, etc.
	DateCriterion []ImmunizationRecommendationRecommendationDateCriterion `bson:"dateCriterion,omitempty" json:"dateCriterion,omitempty" xml:"dateCriterion,omitempty"`
	// Contains the description about the protocol under which the vaccine was administered.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// One possible path to achieve presumed immunity against a disease - within the context of an authority.
	Series *string `bson:"series,omitempty" json:"series,omitempty" xml:"series,omitempty"`
	// Nominal position of the recommended dose in a series as determined by the evaluation and forecasting process (e.g.
	//
	// dose 2 is the next recommended dose).
	//
	// Comment: The use of an integer is preferred if known. A string should only be used in cases where an integer is not available (such as when documenting a recurring booster dose).
	DoseNumber *string `bson:"doseNumber,omitempty" json:"doseNumber,omitempty" xml:"doseNumber,omitempty"`
	// The recommended number of doses to achieve immunity as determined by the evaluation and forecasting process.
	//
	// Comment: The use of an integer is preferred if known.
	SeriesDoses *string `bson:"seriesDoses,omitempty" json:"seriesDoses,omitempty" xml:"seriesDoses,omitempty"`
	// Immunization event history and/or evaluation that supports the status and recommendation.
	SupportingImmunization []Reference `bson:"supportingImmunization,omitempty" json:"supportingImmunization,omitempty" xml:"supportingImmunization,omitempty"`
	// Patient Information that supports the status and recommendation.
	//
	//  This includes patient observations, adverse reactions and allergy/intolerance information.
	SupportingPatientInformation []Reference `bson:"supportingPatientInformation,omitempty" json:"supportingPatientInformation,omitempty" xml:"supportingPatientInformation,omitempty"`
}

type ImmunizationRecommendationRecommendationDateCriterion struct {
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
	// Date classification of recommendation.
	//
	//  For example, earliest date to give, latest date to give, etc.
	Code CodeableConcept `bson:"code" json:"code" xml:"code"`
	// The date whose meaning is specified by dateCriterion.code.
	Value string `bson:"value" json:"value" xml:"value"`
}

func (r ImmunizationRecommendation) ResourceType() string {
	return "ImmunizationRecommendation"
}

// MarshalJSON implements json.Marshaler
func (r ImmunizationRecommendation) MarshalJSON() ([]byte, error) {
	return v2.Marshal(struct {
		ImmunizationRecommendation
		ResourceType string `json:"resourceType"`
	}{
		ImmunizationRecommendation: r,
		ResourceType:               r.ResourceType(),
	})
}
