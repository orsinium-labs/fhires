package fhires

import decimal "github.com/shopspring/decimal"

// Extension Type: Optional Extension Element - found in all resources.
//
// Purpose: The ability to add extensions in a structured way is what keeps FHIR resources simple.
//
// Status: active
//
// http://hl7.org/fhir/StructureDefinition/Extension
type Extension struct {
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
	// Source of the definition for the extension code - a logical name or a URL.
	//
	// Comment: The definition may point directly to a computable or human-readable definition of the extensibility codes, or it may be a logical URI as declared in some other specification. The definition SHALL be a URI for the Structure Definition defining the extension.
	URL string `bson:"url" json:"url" xml:"url"`
	// Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
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
}
