package fhires

import (
	"encoding/json"
	jsontext "encoding/json/jsontext"

	decimal "github.com/shopspring/decimal"
)

// This resource is used to pass information into and back from an operation (whether invoked directly from REST or within a messaging environment).
//
//	It is not persisted or allowed to be referenced by other resources except as described in the definition of the Parameters resource.
//
// Status: active
//
// http://hl7.org/fhir/StructureDefinition/Parameters
type Parameters struct {
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
	// A parameter passed to or received from the operation.
	Parameter []ParametersParameter `bson:"parameter,omitempty" json:"parameter,omitempty" xml:"parameter,omitempty"`
}

type ParametersParameter struct {
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
	// The name of the parameter (reference to the operation definition).
	Name string `bson:"name" json:"name" xml:"name"`
	// Conveys the content if the parameter is a data type.
	ValueBase64Binary          *string                `bson:"valueBase64Binary,omitempty" json:"valueBase64Binary,omitempty" xml:"valueBase64Binary,omitempty"`
	ValueBoolean               *bool                  `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty" xml:"valueBoolean,omitempty"`
	ValueCanonical             *string                `bson:"valueCanonical,omitempty" json:"valueCanonical,omitempty" xml:"valueCanonical,omitempty"`
	ValueCode                  *string                `bson:"valueCode,omitempty" json:"valueCode,omitempty" xml:"valueCode,omitempty"`
	ValueDate                  *string                `bson:"valueDate,omitempty" json:"valueDate,omitempty" xml:"valueDate,omitempty"`
	ValueDateTime              *string                `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty" xml:"valueDateTime,omitempty"`
	ValueDecimal               *decimal.Decimal       `bson:"valueDecimal,omitempty" json:"valueDecimal,omitempty" xml:"valueDecimal,omitempty"`
	ValueID                    *string                `bson:"valueId,omitempty" json:"valueId,omitempty" xml:"valueId,omitempty"`
	ValueInstant               *string                `bson:"valueInstant,omitempty" json:"valueInstant,omitempty" xml:"valueInstant,omitempty"`
	ValueInteger               *int                   `bson:"valueInteger,omitempty" json:"valueInteger,omitempty" xml:"valueInteger,omitempty"`
	ValueInteger64             *int64                 `bson:"valueInteger64,omitempty" json:"valueInteger64,omitempty,string" xml:"valueInteger64,omitempty"`
	ValueMarkdown              *string                `bson:"valueMarkdown,omitempty" json:"valueMarkdown,omitempty" xml:"valueMarkdown,omitempty"`
	ValueOid                   *string                `bson:"valueOid,omitempty" json:"valueOid,omitempty" xml:"valueOid,omitempty"`
	ValuePositiveInt           *int                   `bson:"valuePositiveInt,omitempty" json:"valuePositiveInt,omitempty" xml:"valuePositiveInt,omitempty"`
	ValueString                *string                `bson:"valueString,omitempty" json:"valueString,omitempty" xml:"valueString,omitempty"`
	ValueTime                  *string                `bson:"valueTime,omitempty" json:"valueTime,omitempty" xml:"valueTime,omitempty"`
	ValueUnsignedInt           *int                   `bson:"valueUnsignedInt,omitempty" json:"valueUnsignedInt,omitempty" xml:"valueUnsignedInt,omitempty"`
	ValueUri                   *string                `bson:"valueUri,omitempty" json:"valueUri,omitempty" xml:"valueUri,omitempty"`
	ValueUURL                  *string                `bson:"valueUrl,omitempty" json:"valueUrl,omitempty" xml:"valueUrl,omitempty"`
	ValueUuid                  *string                `bson:"valueUuid,omitempty" json:"valueUuid,omitempty" xml:"valueUuid,omitempty"`
	ValueAddress               *Address               `bson:"valueAddress,omitempty" json:"valueAddress,omitempty" xml:"valueAddress,omitempty"`
	ValueAge                   *Age                   `bson:"valueAge,omitempty" json:"valueAge,omitempty" xml:"valueAge,omitempty"`
	ValueAnnotation            *Annotation            `bson:"valueAnnotation,omitempty" json:"valueAnnotation,omitempty" xml:"valueAnnotation,omitempty"`
	ValueAttachment            *Attachment            `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty" xml:"valueAttachment,omitempty"`
	ValueCodeableConcept       *CodeableConcept       `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty" xml:"valueCodeableConcept,omitempty"`
	ValueCodeableReference     *CodeableReference     `bson:"valueCodeableReference,omitempty" json:"valueCodeableReference,omitempty" xml:"valueCodeableReference,omitempty"`
	ValueCoding                *Coding                `bson:"valueCoding,omitempty" json:"valueCoding,omitempty" xml:"valueCoding,omitempty"`
	ValueContactPoint          *ContactPoint          `bson:"valueContactPoint,omitempty" json:"valueContactPoint,omitempty" xml:"valueContactPoint,omitempty"`
	ValueCount                 *Count                 `bson:"valueCount,omitempty" json:"valueCount,omitempty" xml:"valueCount,omitempty"`
	ValueDistance              *Distance              `bson:"valueDistance,omitempty" json:"valueDistance,omitempty" xml:"valueDistance,omitempty"`
	ValueDuration              *Duration              `bson:"valueDuration,omitempty" json:"valueDuration,omitempty" xml:"valueDuration,omitempty"`
	ValueHumanName             *HumanName             `bson:"valueHumanName,omitempty" json:"valueHumanName,omitempty" xml:"valueHumanName,omitempty"`
	ValueIdentifier            *Identifier            `bson:"valueIdentifier,omitempty" json:"valueIdentifier,omitempty" xml:"valueIdentifier,omitempty"`
	ValueMoney                 *Money                 `bson:"valueMoney,omitempty" json:"valueMoney,omitempty" xml:"valueMoney,omitempty"`
	ValuePeriod                *Period                `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty" xml:"valuePeriod,omitempty"`
	ValueQuantity              *Quantity              `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty" xml:"valueQuantity,omitempty"`
	ValueRange                 *Range                 `bson:"valueRange,omitempty" json:"valueRange,omitempty" xml:"valueRange,omitempty"`
	ValueRatio                 *Ratio                 `bson:"valueRatio,omitempty" json:"valueRatio,omitempty" xml:"valueRatio,omitempty"`
	ValueRatioRange            *RatioRange            `bson:"valueRatioRange,omitempty" json:"valueRatioRange,omitempty" xml:"valueRatioRange,omitempty"`
	ValueReference             *Reference             `bson:"valueReference,omitempty" json:"valueReference,omitempty" xml:"valueReference,omitempty"`
	ValueSampledData           *SampledData           `bson:"valueSampledData,omitempty" json:"valueSampledData,omitempty" xml:"valueSampledData,omitempty"`
	ValueSignature             *Signature             `bson:"valueSignature,omitempty" json:"valueSignature,omitempty" xml:"valueSignature,omitempty"`
	ValueTiming                *Timing                `bson:"valueTiming,omitempty" json:"valueTiming,omitempty" xml:"valueTiming,omitempty"`
	ValueContactDetail         *ContactDetail         `bson:"valueContactDetail,omitempty" json:"valueContactDetail,omitempty" xml:"valueContactDetail,omitempty"`
	ValueDataRequirement       *DataRequirement       `bson:"valueDataRequirement,omitempty" json:"valueDataRequirement,omitempty" xml:"valueDataRequirement,omitempty"`
	ValueExpression            *Expression            `bson:"valueExpression,omitempty" json:"valueExpression,omitempty" xml:"valueExpression,omitempty"`
	ValueParameterDefinition   *ParameterDefinition   `bson:"valueParameterDefinition,omitempty" json:"valueParameterDefinition,omitempty" xml:"valueParameterDefinition,omitempty"`
	ValueRelatedArtifact       *RelatedArtifact       `bson:"valueRelatedArtifact,omitempty" json:"valueRelatedArtifact,omitempty" xml:"valueRelatedArtifact,omitempty"`
	ValueTriggerDefinition     *TriggerDefinition     `bson:"valueTriggerDefinition,omitempty" json:"valueTriggerDefinition,omitempty" xml:"valueTriggerDefinition,omitempty"`
	ValueUsageContext          *UsageContext          `bson:"valueUsageContext,omitempty" json:"valueUsageContext,omitempty" xml:"valueUsageContext,omitempty"`
	ValueAvailability          *Availability          `bson:"valueAvailability,omitempty" json:"valueAvailability,omitempty" xml:"valueAvailability,omitempty"`
	ValueExtendedContactDetail *ExtendedContactDetail `bson:"valueExtendedContactDetail,omitempty" json:"valueExtendedContactDetail,omitempty" xml:"valueExtendedContactDetail,omitempty"`
	ValueDosage                *Dosage                `bson:"valueDosage,omitempty" json:"valueDosage,omitempty" xml:"valueDosage,omitempty"`
	ValueMeta                  *Meta                  `bson:"valueMeta,omitempty" json:"valueMeta,omitempty" xml:"valueMeta,omitempty"`
	// Conveys the content if the parameter is a whole resource.
	//
	// Comment: When resolving references in resources, the operation definition may specify how references may be resolved between parameters. If a reference cannot be resolved between the parameters, the application should fall back to its general resource resolution methods.
	Resource jsontext.Value `bson:"resource,omitempty" json:"resource,omitempty" xml:"resource,omitempty"`
	// A named part of a multi-part parameter.
	Part []ParametersParameter `bson:"part,omitempty" json:"part,omitempty"`
}

func (r Parameters) ResourceType() string {
	return "Parameters"
}

// MarshalJSON implements json.Marshaler
func (r Parameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Parameters
		ResourceType string `json:"resourceType"`
	}{
		Parameters:   r,
		ResourceType: r.ResourceType(),
	})
}
