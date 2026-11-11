package fhires

import decimal "github.com/shopspring/decimal"

// ElementDefinition Type: Captures constraints on each element within the resource, profile, or extension.
//
// Status: active
//
// http://hl7.org/fhir/StructureDefinition/ElementDefinition
type ElementDefinition struct {
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
	// The path identifies the element and is expressed as a "."-separated list of ancestor elements, beginning with the name of the resource or extension.
	Path string `bson:"path" json:"path" xml:"path"`
	// Codes that define how this element is represented in instances, when the deviation varies from the normal case.
	//
	// No extensions are allowed on elements with a representation of 'xmlAttr', no matter what FHIR serialization format is used.
	//
	// Comment: In resources, this is rarely used except for special cases where the representation deviates from the normal, and can only be done in the base standard (and profiles must reproduce what the base standard does). This element is used quite commonly in Logical models when the logical models represent a specific serialization format (e.g. CDA, v2 etc.).
	Representation []PropertyRepresentation `bson:"representation,omitempty" json:"representation,omitempty" xml:"representation,omitempty"`
	// The name of this element definition slice, when slicing is working.
	//
	// The name must be a token with no dots or spaces. This is a unique name referring to a specific set of constraints applied to this element, used to provide a name to different slices of the same element.
	//
	// Comment: The name SHALL be unique within the structure within the context of the constrained resource element.  (Though to avoid confusion, uniqueness across all elements is recommended.).
	SliceName *string `bson:"sliceName,omitempty" json:"sliceName,omitempty" xml:"sliceName,omitempty"`
	// If true, indicates that this slice definition is constraining a slice definition with the same name in an inherited profile.
	//
	// If false, the slice is not overriding any slice in an inherited profile. If missing, the slice might or might not be overriding a slice in an inherited profile, depending on the sliceName.
	//
	// Comment: If set to true, an ancestor profile SHALL have a slicing definition with this name.  If set to false, no ancestor profile is permitted to have a slicing definition with this name.
	SliceIsConstraining *bool `bson:"sliceIsConstraining,omitempty" json:"sliceIsConstraining,omitempty" xml:"sliceIsConstraining,omitempty"`
	// A single preferred label which is the text to display beside the element indicating its meaning or to use to prompt for the element in a user display or form.
	//
	// Comment: See also the extension [http://hl7.org/fhir/StructureDefinition/elementdefinition-question](http://hl7.org/fhir/extensions/StructureDefinition-elementdefinition-question.html).
	Label *string `bson:"label,omitempty" json:"label,omitempty" xml:"label,omitempty"`
	// A code that has the same meaning as the element in a particular terminology.
	//
	// Comment: The concept SHALL be properly aligned with the data element definition and other constraints, as defined in the code system, including relationships, of any code listed here.  Where multiple codes exist in a terminology that could correspond to the data element, the most granular code(s) should be selected, so long as they are not more restrictive than the data element itself. The mappings may be used to provide more or less granular or structured equivalences in the code system.
	Code []Coding `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// Indicates that the element is sliced into a set of alternative definitions (i.e.
	//
	// in a structure definition, there are multiple different constraints on a single element in the base resource). Slicing can be used in any resource that has cardinality ..* on the base resource, or any resource with a choice of types. The set of slices is any elements that come after this in the element sequence that have the same path, until a shorter path occurs (the shorter path terminates the set).
	//
	// Comment: The first element in the sequence, the one that carries the slicing, is the definition that applies to all the slices. This is based on the unconstrained element, but can apply any constraints as appropriate. This may include the common constraints on the children of the element.
	Slicing *ElementDefinitionSlicing `bson:"slicing,omitempty" json:"slicing,omitempty" xml:"slicing,omitempty"`
	// A concise description of what this element means (e.g.
	//
	// for use in autogenerated summaries).
	//
	// Comment: It is easy for a different short definition to change the meaning of an element and this can have nasty downstream consequences. Please be careful when providing short definitions in a profile.
	Short *string `bson:"short,omitempty" json:"short,omitempty" xml:"short,omitempty"`
	// Provides a complete explanation of the meaning of the data element for human readability.
	//
	//  For the case of elements derived from existing elements (e.g. constraints), the definition SHALL be consistent with the base definition, but convey the meaning of the element in the particular context of use of the resource. (Note: The text you are reading is specified in ElementDefinition.definition).
	//
	// Comment: It is easy for a different definition to change the meaning of an element and this can have nasty downstream consequences. Please be careful when providing definitions in a profile.
	Definition *string `bson:"definition,omitempty" json:"definition,omitempty" xml:"definition,omitempty"`
	// Explanatory notes and implementation guidance about the data element, including notes about how to use the data properly, exceptions to proper use, etc.
	//
	// (Note: The text you are reading is specified in ElementDefinition.comment).
	//
	// Comment: If it is possible to capture usage rules using constraints, that mechanism should be used in preference to this element.
	Comment *string `bson:"comment,omitempty" json:"comment,omitempty" xml:"comment,omitempty"`
	// This element is for traceability of why the element was created and why the constraints exist as they do.
	//
	// This may be used to point to source materials or specifications that drove the structure of this element.
	//
	// Comment: This element does not describe the usage of the element (that's done in comments), rather it's for traceability of *why* the element is either needed or why the constraints exist as they do.  This may be used to point to source materials or specifications that drove the structure of this data element.
	Requirements *string `bson:"requirements,omitempty" json:"requirements,omitempty" xml:"requirements,omitempty"`
	// Identifies additional names by which this element might also be known.
	Alias []string `bson:"alias,omitempty" json:"alias,omitempty" xml:"alias,omitempty"`
	// The minimum number of times this element SHALL appear in the instance.
	Min *int `bson:"min,omitempty" json:"min,omitempty" xml:"min,omitempty"`
	// The maximum number of times this element is permitted to appear in the instance.
	Max *string `bson:"max,omitempty" json:"max,omitempty" xml:"max,omitempty"`
	// Information about the base definition of the element, provided to make it unnecessary for tools to trace the deviation of the element through the derived and related profiles.
	//
	// When the element definition is not the original definition of an element - e.g. either in a constraint on another type, or for elements from a super type in a snap shot - then the information in provided in the element definition may be different to the base definition. On the original definition of the element, it will be same.
	//
	// Comment: The base information does not carry any information that could not be determined from the path and related profiles, but making this determination requires both that the related profiles are available, and that the algorithm to determine them be available. For tooling simplicity, the base information must always be populated in element definitions in snap shots, even if it is the same.
	Base *ElementDefinitionBase `bson:"base,omitempty" json:"base,omitempty" xml:"base,omitempty"`
	// Identifies an element defined elsewhere in the definition whose content rules should be applied to the current element.
	//
	// ContentReferences bring across all the rules that are in the ElementDefinition for the element, including definitions, cardinality constraints, bindings, invariants etc.
	//
	// Comment: ContentReferences can only be defined in specializations, not constrained types, and they cannot be changed and always reference the non-constrained definition.
	ContentReference *string `bson:"contentReference,omitempty" json:"contentReference,omitempty" xml:"contentReference,omitempty"`
	// The data type or resource that the value of this element is permitted to be.
	//
	// Comment: The Type of the element can be left blank in a differential constraint, in which case the type is inherited from the resource. Abstract types are not permitted to appear as a type when multiple types are listed.  (I.e. Abstract types cannot be part of a choice).
	Type []ElementDefinitionType `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The value that should be used if there is no value stated in the instance (e.g.
	//
	// 'if not otherwise specified, the abstract is false').
	//
	// Comment: Specifying a default value means that the property can never been unknown - it must always have a value. Further, the default value can never be changed, or changed in constraints on content models. Defining default values creates many difficulties in implementation (e.g. when is a value missing?). For these reasons, default values are (and should be) used extremely sparingly.
	//
	// No default values are ever defined in the FHIR specification, nor can they be defined in constraints ("profiles") on data types or resources. This element only exists so that default values may be defined in logical models.
	DefaultValueBase64Binary          *string                `bson:"defaultValueBase64Binary,omitempty" json:"defaultValueBase64Binary,omitempty" xml:"defaultValueBase64Binary,omitempty"`
	DefaultValueBoolean               *bool                  `bson:"defaultValueBoolean,omitempty" json:"defaultValueBoolean,omitempty" xml:"defaultValueBoolean,omitempty"`
	DefaultValueCanonical             *string                `bson:"defaultValueCanonical,omitempty" json:"defaultValueCanonical,omitempty" xml:"defaultValueCanonical,omitempty"`
	DefaultValueCode                  *string                `bson:"defaultValueCode,omitempty" json:"defaultValueCode,omitempty" xml:"defaultValueCode,omitempty"`
	DefaultValueDate                  *string                `bson:"defaultValueDate,omitempty" json:"defaultValueDate,omitempty" xml:"defaultValueDate,omitempty"`
	DefaultValueDateTime              *string                `bson:"defaultValueDateTime,omitempty" json:"defaultValueDateTime,omitempty" xml:"defaultValueDateTime,omitempty"`
	DefaultValueDecimal               *decimal.Decimal       `bson:"defaultValueDecimal,omitempty" json:"defaultValueDecimal,omitempty" xml:"defaultValueDecimal,omitempty"`
	DefaultValueID                    *string                `bson:"defaultValueId,omitempty" json:"defaultValueId,omitempty" xml:"defaultValueId,omitempty"`
	DefaultValueInstant               *string                `bson:"defaultValueInstant,omitempty" json:"defaultValueInstant,omitempty" xml:"defaultValueInstant,omitempty"`
	DefaultValueInteger               *int                   `bson:"defaultValueInteger,omitempty" json:"defaultValueInteger,omitempty" xml:"defaultValueInteger,omitempty"`
	DefaultValueInteger64             *int64                 `bson:"defaultValueInteger64,omitempty" json:"defaultValueInteger64,omitempty,string" xml:"defaultValueInteger64,omitempty"`
	DefaultValueMarkdown              *string                `bson:"defaultValueMarkdown,omitempty" json:"defaultValueMarkdown,omitempty" xml:"defaultValueMarkdown,omitempty"`
	DefaultValueOid                   *string                `bson:"defaultValueOid,omitempty" json:"defaultValueOid,omitempty" xml:"defaultValueOid,omitempty"`
	DefaultValuePositiveInt           *int                   `bson:"defaultValuePositiveInt,omitempty" json:"defaultValuePositiveInt,omitempty" xml:"defaultValuePositiveInt,omitempty"`
	DefaultValueString                *string                `bson:"defaultValueString,omitempty" json:"defaultValueString,omitempty" xml:"defaultValueString,omitempty"`
	DefaultValueTime                  *string                `bson:"defaultValueTime,omitempty" json:"defaultValueTime,omitempty" xml:"defaultValueTime,omitempty"`
	DefaultValueUnsignedInt           *int                   `bson:"defaultValueUnsignedInt,omitempty" json:"defaultValueUnsignedInt,omitempty" xml:"defaultValueUnsignedInt,omitempty"`
	DefaultValueUri                   *string                `bson:"defaultValueUri,omitempty" json:"defaultValueUri,omitempty" xml:"defaultValueUri,omitempty"`
	DefaultValueUURL                  *string                `bson:"defaultValueUrl,omitempty" json:"defaultValueUrl,omitempty" xml:"defaultValueUrl,omitempty"`
	DefaultValueUuid                  *string                `bson:"defaultValueUuid,omitempty" json:"defaultValueUuid,omitempty" xml:"defaultValueUuid,omitempty"`
	DefaultValueAddress               *Address               `bson:"defaultValueAddress,omitempty" json:"defaultValueAddress,omitempty" xml:"defaultValueAddress,omitempty"`
	DefaultValueAge                   *Age                   `bson:"defaultValueAge,omitempty" json:"defaultValueAge,omitempty" xml:"defaultValueAge,omitempty"`
	DefaultValueAnnotation            *Annotation            `bson:"defaultValueAnnotation,omitempty" json:"defaultValueAnnotation,omitempty" xml:"defaultValueAnnotation,omitempty"`
	DefaultValueAttachment            *Attachment            `bson:"defaultValueAttachment,omitempty" json:"defaultValueAttachment,omitempty" xml:"defaultValueAttachment,omitempty"`
	DefaultValueCodeableConcept       *CodeableConcept       `bson:"defaultValueCodeableConcept,omitempty" json:"defaultValueCodeableConcept,omitempty" xml:"defaultValueCodeableConcept,omitempty"`
	DefaultValueCodeableReference     *CodeableReference     `bson:"defaultValueCodeableReference,omitempty" json:"defaultValueCodeableReference,omitempty" xml:"defaultValueCodeableReference,omitempty"`
	DefaultValueCoding                *Coding                `bson:"defaultValueCoding,omitempty" json:"defaultValueCoding,omitempty" xml:"defaultValueCoding,omitempty"`
	DefaultValueContactPoint          *ContactPoint          `bson:"defaultValueContactPoint,omitempty" json:"defaultValueContactPoint,omitempty" xml:"defaultValueContactPoint,omitempty"`
	DefaultValueCount                 *Count                 `bson:"defaultValueCount,omitempty" json:"defaultValueCount,omitempty" xml:"defaultValueCount,omitempty"`
	DefaultValueDistance              *Distance              `bson:"defaultValueDistance,omitempty" json:"defaultValueDistance,omitempty" xml:"defaultValueDistance,omitempty"`
	DefaultValueDuration              *Duration              `bson:"defaultValueDuration,omitempty" json:"defaultValueDuration,omitempty" xml:"defaultValueDuration,omitempty"`
	DefaultValueHumanName             *HumanName             `bson:"defaultValueHumanName,omitempty" json:"defaultValueHumanName,omitempty" xml:"defaultValueHumanName,omitempty"`
	DefaultValueIdentifier            *Identifier            `bson:"defaultValueIdentifier,omitempty" json:"defaultValueIdentifier,omitempty" xml:"defaultValueIdentifier,omitempty"`
	DefaultValueMoney                 *Money                 `bson:"defaultValueMoney,omitempty" json:"defaultValueMoney,omitempty" xml:"defaultValueMoney,omitempty"`
	DefaultValuePeriod                *Period                `bson:"defaultValuePeriod,omitempty" json:"defaultValuePeriod,omitempty" xml:"defaultValuePeriod,omitempty"`
	DefaultValueQuantity              *Quantity              `bson:"defaultValueQuantity,omitempty" json:"defaultValueQuantity,omitempty" xml:"defaultValueQuantity,omitempty"`
	DefaultValueRange                 *Range                 `bson:"defaultValueRange,omitempty" json:"defaultValueRange,omitempty" xml:"defaultValueRange,omitempty"`
	DefaultValueRatio                 *Ratio                 `bson:"defaultValueRatio,omitempty" json:"defaultValueRatio,omitempty" xml:"defaultValueRatio,omitempty"`
	DefaultValueRatioRange            *RatioRange            `bson:"defaultValueRatioRange,omitempty" json:"defaultValueRatioRange,omitempty" xml:"defaultValueRatioRange,omitempty"`
	DefaultValueReference             *Reference             `bson:"defaultValueReference,omitempty" json:"defaultValueReference,omitempty" xml:"defaultValueReference,omitempty"`
	DefaultValueSampledData           *SampledData           `bson:"defaultValueSampledData,omitempty" json:"defaultValueSampledData,omitempty" xml:"defaultValueSampledData,omitempty"`
	DefaultValueSignature             *Signature             `bson:"defaultValueSignature,omitempty" json:"defaultValueSignature,omitempty" xml:"defaultValueSignature,omitempty"`
	DefaultValueTiming                *Timing                `bson:"defaultValueTiming,omitempty" json:"defaultValueTiming,omitempty" xml:"defaultValueTiming,omitempty"`
	DefaultValueContactDetail         *ContactDetail         `bson:"defaultValueContactDetail,omitempty" json:"defaultValueContactDetail,omitempty" xml:"defaultValueContactDetail,omitempty"`
	DefaultValueDataRequirement       *DataRequirement       `bson:"defaultValueDataRequirement,omitempty" json:"defaultValueDataRequirement,omitempty" xml:"defaultValueDataRequirement,omitempty"`
	DefaultValueExpression            *Expression            `bson:"defaultValueExpression,omitempty" json:"defaultValueExpression,omitempty" xml:"defaultValueExpression,omitempty"`
	DefaultValueParameterDefinition   *ParameterDefinition   `bson:"defaultValueParameterDefinition,omitempty" json:"defaultValueParameterDefinition,omitempty" xml:"defaultValueParameterDefinition,omitempty"`
	DefaultValueRelatedArtifact       *RelatedArtifact       `bson:"defaultValueRelatedArtifact,omitempty" json:"defaultValueRelatedArtifact,omitempty" xml:"defaultValueRelatedArtifact,omitempty"`
	DefaultValueTriggerDefinition     *TriggerDefinition     `bson:"defaultValueTriggerDefinition,omitempty" json:"defaultValueTriggerDefinition,omitempty" xml:"defaultValueTriggerDefinition,omitempty"`
	DefaultValueUsageContext          *UsageContext          `bson:"defaultValueUsageContext,omitempty" json:"defaultValueUsageContext,omitempty" xml:"defaultValueUsageContext,omitempty"`
	DefaultValueAvailability          *Availability          `bson:"defaultValueAvailability,omitempty" json:"defaultValueAvailability,omitempty" xml:"defaultValueAvailability,omitempty"`
	DefaultValueExtendedContactDetail *ExtendedContactDetail `bson:"defaultValueExtendedContactDetail,omitempty" json:"defaultValueExtendedContactDetail,omitempty" xml:"defaultValueExtendedContactDetail,omitempty"`
	DefaultValueDosage                *Dosage                `bson:"defaultValueDosage,omitempty" json:"defaultValueDosage,omitempty" xml:"defaultValueDosage,omitempty"`
	DefaultValueMeta                  *Meta                  `bson:"defaultValueMeta,omitempty" json:"defaultValueMeta,omitempty" xml:"defaultValueMeta,omitempty"`
	// The Implicit meaning that is to be understood when this element is missing (e.g.
	//
	// 'when this element is missing, the period is ongoing').
	//
	// Comment: Implicit meanings for missing values can only be specified on a resource, data type, or extension definition, and never in a profile that applies to one of these. An implicit meaning for a missing value can never be changed, and specifying one has the consequence that constraining its use in profiles eliminates use cases as possibilities, not merely moving them out of scope.
	MeaningWhenMissing *string `bson:"meaningWhenMissing,omitempty" json:"meaningWhenMissing,omitempty" xml:"meaningWhenMissing,omitempty"`
	// If present, indicates that the order of the repeating element has meaning and describes what that meaning is.
	//
	//  If absent, it means that the order of the element has no meaning.
	//
	// Comment: This element can only be asserted on repeating elements and can only be introduced when defining resources or data types.  It can be further refined profiled elements but if absent in the base type, a profile cannot assert meaning.
	OrderMeaning *string `bson:"orderMeaning,omitempty" json:"orderMeaning,omitempty" xml:"orderMeaning,omitempty"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance, if present.
	//
	// For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	//
	// Comment: This is not recommended for Coding and CodeableConcept since these often have highly contextual properties such as version or display.
	FixedBase64Binary          *string                `bson:"fixedBase64Binary,omitempty" json:"fixedBase64Binary,omitempty" xml:"fixedBase64Binary,omitempty"`
	FixedBoolean               *bool                  `bson:"fixedBoolean,omitempty" json:"fixedBoolean,omitempty" xml:"fixedBoolean,omitempty"`
	FixedCanonical             *string                `bson:"fixedCanonical,omitempty" json:"fixedCanonical,omitempty" xml:"fixedCanonical,omitempty"`
	FixedCode                  *string                `bson:"fixedCode,omitempty" json:"fixedCode,omitempty" xml:"fixedCode,omitempty"`
	FixedDate                  *string                `bson:"fixedDate,omitempty" json:"fixedDate,omitempty" xml:"fixedDate,omitempty"`
	FixedDateTime              *string                `bson:"fixedDateTime,omitempty" json:"fixedDateTime,omitempty" xml:"fixedDateTime,omitempty"`
	FixedDecimal               *decimal.Decimal       `bson:"fixedDecimal,omitempty" json:"fixedDecimal,omitempty" xml:"fixedDecimal,omitempty"`
	FixedID                    *string                `bson:"fixedId,omitempty" json:"fixedId,omitempty" xml:"fixedId,omitempty"`
	FixedInstant               *string                `bson:"fixedInstant,omitempty" json:"fixedInstant,omitempty" xml:"fixedInstant,omitempty"`
	FixedInteger               *int                   `bson:"fixedInteger,omitempty" json:"fixedInteger,omitempty" xml:"fixedInteger,omitempty"`
	FixedInteger64             *int64                 `bson:"fixedInteger64,omitempty" json:"fixedInteger64,omitempty,string" xml:"fixedInteger64,omitempty"`
	FixedMarkdown              *string                `bson:"fixedMarkdown,omitempty" json:"fixedMarkdown,omitempty" xml:"fixedMarkdown,omitempty"`
	FixedOid                   *string                `bson:"fixedOid,omitempty" json:"fixedOid,omitempty" xml:"fixedOid,omitempty"`
	FixedPositiveInt           *int                   `bson:"fixedPositiveInt,omitempty" json:"fixedPositiveInt,omitempty" xml:"fixedPositiveInt,omitempty"`
	FixedString                *string                `bson:"fixedString,omitempty" json:"fixedString,omitempty" xml:"fixedString,omitempty"`
	FixedTime                  *string                `bson:"fixedTime,omitempty" json:"fixedTime,omitempty" xml:"fixedTime,omitempty"`
	FixedUnsignedInt           *int                   `bson:"fixedUnsignedInt,omitempty" json:"fixedUnsignedInt,omitempty" xml:"fixedUnsignedInt,omitempty"`
	FixedUri                   *string                `bson:"fixedUri,omitempty" json:"fixedUri,omitempty" xml:"fixedUri,omitempty"`
	FixedUURL                  *string                `bson:"fixedUrl,omitempty" json:"fixedUrl,omitempty" xml:"fixedUrl,omitempty"`
	FixedUuid                  *string                `bson:"fixedUuid,omitempty" json:"fixedUuid,omitempty" xml:"fixedUuid,omitempty"`
	FixedAddress               *Address               `bson:"fixedAddress,omitempty" json:"fixedAddress,omitempty" xml:"fixedAddress,omitempty"`
	FixedAge                   *Age                   `bson:"fixedAge,omitempty" json:"fixedAge,omitempty" xml:"fixedAge,omitempty"`
	FixedAnnotation            *Annotation            `bson:"fixedAnnotation,omitempty" json:"fixedAnnotation,omitempty" xml:"fixedAnnotation,omitempty"`
	FixedAttachment            *Attachment            `bson:"fixedAttachment,omitempty" json:"fixedAttachment,omitempty" xml:"fixedAttachment,omitempty"`
	FixedCodeableConcept       *CodeableConcept       `bson:"fixedCodeableConcept,omitempty" json:"fixedCodeableConcept,omitempty" xml:"fixedCodeableConcept,omitempty"`
	FixedCodeableReference     *CodeableReference     `bson:"fixedCodeableReference,omitempty" json:"fixedCodeableReference,omitempty" xml:"fixedCodeableReference,omitempty"`
	FixedCoding                *Coding                `bson:"fixedCoding,omitempty" json:"fixedCoding,omitempty" xml:"fixedCoding,omitempty"`
	FixedContactPoint          *ContactPoint          `bson:"fixedContactPoint,omitempty" json:"fixedContactPoint,omitempty" xml:"fixedContactPoint,omitempty"`
	FixedCount                 *Count                 `bson:"fixedCount,omitempty" json:"fixedCount,omitempty" xml:"fixedCount,omitempty"`
	FixedDistance              *Distance              `bson:"fixedDistance,omitempty" json:"fixedDistance,omitempty" xml:"fixedDistance,omitempty"`
	FixedDuration              *Duration              `bson:"fixedDuration,omitempty" json:"fixedDuration,omitempty" xml:"fixedDuration,omitempty"`
	FixedHumanName             *HumanName             `bson:"fixedHumanName,omitempty" json:"fixedHumanName,omitempty" xml:"fixedHumanName,omitempty"`
	FixedIdentifier            *Identifier            `bson:"fixedIdentifier,omitempty" json:"fixedIdentifier,omitempty" xml:"fixedIdentifier,omitempty"`
	FixedMoney                 *Money                 `bson:"fixedMoney,omitempty" json:"fixedMoney,omitempty" xml:"fixedMoney,omitempty"`
	FixedPeriod                *Period                `bson:"fixedPeriod,omitempty" json:"fixedPeriod,omitempty" xml:"fixedPeriod,omitempty"`
	FixedQuantity              *Quantity              `bson:"fixedQuantity,omitempty" json:"fixedQuantity,omitempty" xml:"fixedQuantity,omitempty"`
	FixedRange                 *Range                 `bson:"fixedRange,omitempty" json:"fixedRange,omitempty" xml:"fixedRange,omitempty"`
	FixedRatio                 *Ratio                 `bson:"fixedRatio,omitempty" json:"fixedRatio,omitempty" xml:"fixedRatio,omitempty"`
	FixedRatioRange            *RatioRange            `bson:"fixedRatioRange,omitempty" json:"fixedRatioRange,omitempty" xml:"fixedRatioRange,omitempty"`
	FixedReference             *Reference             `bson:"fixedReference,omitempty" json:"fixedReference,omitempty" xml:"fixedReference,omitempty"`
	FixedSampledData           *SampledData           `bson:"fixedSampledData,omitempty" json:"fixedSampledData,omitempty" xml:"fixedSampledData,omitempty"`
	FixedSignature             *Signature             `bson:"fixedSignature,omitempty" json:"fixedSignature,omitempty" xml:"fixedSignature,omitempty"`
	FixedTiming                *Timing                `bson:"fixedTiming,omitempty" json:"fixedTiming,omitempty" xml:"fixedTiming,omitempty"`
	FixedContactDetail         *ContactDetail         `bson:"fixedContactDetail,omitempty" json:"fixedContactDetail,omitempty" xml:"fixedContactDetail,omitempty"`
	FixedDataRequirement       *DataRequirement       `bson:"fixedDataRequirement,omitempty" json:"fixedDataRequirement,omitempty" xml:"fixedDataRequirement,omitempty"`
	FixedExpression            *Expression            `bson:"fixedExpression,omitempty" json:"fixedExpression,omitempty" xml:"fixedExpression,omitempty"`
	FixedParameterDefinition   *ParameterDefinition   `bson:"fixedParameterDefinition,omitempty" json:"fixedParameterDefinition,omitempty" xml:"fixedParameterDefinition,omitempty"`
	FixedRelatedArtifact       *RelatedArtifact       `bson:"fixedRelatedArtifact,omitempty" json:"fixedRelatedArtifact,omitempty" xml:"fixedRelatedArtifact,omitempty"`
	FixedTriggerDefinition     *TriggerDefinition     `bson:"fixedTriggerDefinition,omitempty" json:"fixedTriggerDefinition,omitempty" xml:"fixedTriggerDefinition,omitempty"`
	FixedUsageContext          *UsageContext          `bson:"fixedUsageContext,omitempty" json:"fixedUsageContext,omitempty" xml:"fixedUsageContext,omitempty"`
	FixedAvailability          *Availability          `bson:"fixedAvailability,omitempty" json:"fixedAvailability,omitempty" xml:"fixedAvailability,omitempty"`
	FixedExtendedContactDetail *ExtendedContactDetail `bson:"fixedExtendedContactDetail,omitempty" json:"fixedExtendedContactDetail,omitempty" xml:"fixedExtendedContactDetail,omitempty"`
	FixedDosage                *Dosage                `bson:"fixedDosage,omitempty" json:"fixedDosage,omitempty" xml:"fixedDosage,omitempty"`
	FixedMeta                  *Meta                  `bson:"fixedMeta,omitempty" json:"fixedMeta,omitempty" xml:"fixedMeta,omitempty"`
	// Specifies a value that each occurrence of the element in the instance SHALL follow - that is, any value in the pattern must be found in the instance, if the element has a value.
	//
	// Other additional values may be found too. This is effectively constraint by example.
	//
	// When pattern[x] is used to constrain a primitive, it means that the value provided in the pattern[x] must match the instance value exactly.
	//
	// When an element within a pattern[x] is used to constrain an array, it means that each element provided in the pattern[x] must (recursively) match at least one element from the instance array.
	//
	// When pattern[x] is used to constrain a complex object, it means that each property in the pattern must be present in the complex object, and its value must recursively match -- i.e.,
	//
	// 1. If primitive: it must match exactly the pattern value
	//
	// 2. If a complex object: it must match (recursively) the pattern value
	//
	// 3. If an array: it must match (recursively) the pattern value
	//
	// If a pattern[x] is declared on a repeating element, the pattern applies to all repetitions.  If the desire is for a pattern to apply to only one element or a subset of elements, slicing must be used. See [Examples of Patterns](elementdefinition-examples.html#pattern-examples) for examples of pattern usage and the effect it will have.
	//
	// Comment: Mostly used for fixing values of CodeableConcept. In general, pattern[x] is not intended for use with primitive types, where is has the same meaning as fixed[x].
	PatternBase64Binary          *string                `bson:"patternBase64Binary,omitempty" json:"patternBase64Binary,omitempty" xml:"patternBase64Binary,omitempty"`
	PatternBoolean               *bool                  `bson:"patternBoolean,omitempty" json:"patternBoolean,omitempty" xml:"patternBoolean,omitempty"`
	PatternCanonical             *string                `bson:"patternCanonical,omitempty" json:"patternCanonical,omitempty" xml:"patternCanonical,omitempty"`
	PatternCode                  *string                `bson:"patternCode,omitempty" json:"patternCode,omitempty" xml:"patternCode,omitempty"`
	PatternDate                  *string                `bson:"patternDate,omitempty" json:"patternDate,omitempty" xml:"patternDate,omitempty"`
	PatternDateTime              *string                `bson:"patternDateTime,omitempty" json:"patternDateTime,omitempty" xml:"patternDateTime,omitempty"`
	PatternDecimal               *decimal.Decimal       `bson:"patternDecimal,omitempty" json:"patternDecimal,omitempty" xml:"patternDecimal,omitempty"`
	PatternID                    *string                `bson:"patternId,omitempty" json:"patternId,omitempty" xml:"patternId,omitempty"`
	PatternInstant               *string                `bson:"patternInstant,omitempty" json:"patternInstant,omitempty" xml:"patternInstant,omitempty"`
	PatternInteger               *int                   `bson:"patternInteger,omitempty" json:"patternInteger,omitempty" xml:"patternInteger,omitempty"`
	PatternInteger64             *int64                 `bson:"patternInteger64,omitempty" json:"patternInteger64,omitempty,string" xml:"patternInteger64,omitempty"`
	PatternMarkdown              *string                `bson:"patternMarkdown,omitempty" json:"patternMarkdown,omitempty" xml:"patternMarkdown,omitempty"`
	PatternOid                   *string                `bson:"patternOid,omitempty" json:"patternOid,omitempty" xml:"patternOid,omitempty"`
	PatternPositiveInt           *int                   `bson:"patternPositiveInt,omitempty" json:"patternPositiveInt,omitempty" xml:"patternPositiveInt,omitempty"`
	PatternString                *string                `bson:"patternString,omitempty" json:"patternString,omitempty" xml:"patternString,omitempty"`
	PatternTime                  *string                `bson:"patternTime,omitempty" json:"patternTime,omitempty" xml:"patternTime,omitempty"`
	PatternUnsignedInt           *int                   `bson:"patternUnsignedInt,omitempty" json:"patternUnsignedInt,omitempty" xml:"patternUnsignedInt,omitempty"`
	PatternUri                   *string                `bson:"patternUri,omitempty" json:"patternUri,omitempty" xml:"patternUri,omitempty"`
	PatternUURL                  *string                `bson:"patternUrl,omitempty" json:"patternUrl,omitempty" xml:"patternUrl,omitempty"`
	PatternUuid                  *string                `bson:"patternUuid,omitempty" json:"patternUuid,omitempty" xml:"patternUuid,omitempty"`
	PatternAddress               *Address               `bson:"patternAddress,omitempty" json:"patternAddress,omitempty" xml:"patternAddress,omitempty"`
	PatternAge                   *Age                   `bson:"patternAge,omitempty" json:"patternAge,omitempty" xml:"patternAge,omitempty"`
	PatternAnnotation            *Annotation            `bson:"patternAnnotation,omitempty" json:"patternAnnotation,omitempty" xml:"patternAnnotation,omitempty"`
	PatternAttachment            *Attachment            `bson:"patternAttachment,omitempty" json:"patternAttachment,omitempty" xml:"patternAttachment,omitempty"`
	PatternCodeableConcept       *CodeableConcept       `bson:"patternCodeableConcept,omitempty" json:"patternCodeableConcept,omitempty" xml:"patternCodeableConcept,omitempty"`
	PatternCodeableReference     *CodeableReference     `bson:"patternCodeableReference,omitempty" json:"patternCodeableReference,omitempty" xml:"patternCodeableReference,omitempty"`
	PatternCoding                *Coding                `bson:"patternCoding,omitempty" json:"patternCoding,omitempty" xml:"patternCoding,omitempty"`
	PatternContactPoint          *ContactPoint          `bson:"patternContactPoint,omitempty" json:"patternContactPoint,omitempty" xml:"patternContactPoint,omitempty"`
	PatternCount                 *Count                 `bson:"patternCount,omitempty" json:"patternCount,omitempty" xml:"patternCount,omitempty"`
	PatternDistance              *Distance              `bson:"patternDistance,omitempty" json:"patternDistance,omitempty" xml:"patternDistance,omitempty"`
	PatternDuration              *Duration              `bson:"patternDuration,omitempty" json:"patternDuration,omitempty" xml:"patternDuration,omitempty"`
	PatternHumanName             *HumanName             `bson:"patternHumanName,omitempty" json:"patternHumanName,omitempty" xml:"patternHumanName,omitempty"`
	PatternIdentifier            *Identifier            `bson:"patternIdentifier,omitempty" json:"patternIdentifier,omitempty" xml:"patternIdentifier,omitempty"`
	PatternMoney                 *Money                 `bson:"patternMoney,omitempty" json:"patternMoney,omitempty" xml:"patternMoney,omitempty"`
	PatternPeriod                *Period                `bson:"patternPeriod,omitempty" json:"patternPeriod,omitempty" xml:"patternPeriod,omitempty"`
	PatternQuantity              *Quantity              `bson:"patternQuantity,omitempty" json:"patternQuantity,omitempty" xml:"patternQuantity,omitempty"`
	PatternRange                 *Range                 `bson:"patternRange,omitempty" json:"patternRange,omitempty" xml:"patternRange,omitempty"`
	PatternRatio                 *Ratio                 `bson:"patternRatio,omitempty" json:"patternRatio,omitempty" xml:"patternRatio,omitempty"`
	PatternRatioRange            *RatioRange            `bson:"patternRatioRange,omitempty" json:"patternRatioRange,omitempty" xml:"patternRatioRange,omitempty"`
	PatternReference             *Reference             `bson:"patternReference,omitempty" json:"patternReference,omitempty" xml:"patternReference,omitempty"`
	PatternSampledData           *SampledData           `bson:"patternSampledData,omitempty" json:"patternSampledData,omitempty" xml:"patternSampledData,omitempty"`
	PatternSignature             *Signature             `bson:"patternSignature,omitempty" json:"patternSignature,omitempty" xml:"patternSignature,omitempty"`
	PatternTiming                *Timing                `bson:"patternTiming,omitempty" json:"patternTiming,omitempty" xml:"patternTiming,omitempty"`
	PatternContactDetail         *ContactDetail         `bson:"patternContactDetail,omitempty" json:"patternContactDetail,omitempty" xml:"patternContactDetail,omitempty"`
	PatternDataRequirement       *DataRequirement       `bson:"patternDataRequirement,omitempty" json:"patternDataRequirement,omitempty" xml:"patternDataRequirement,omitempty"`
	PatternExpression            *Expression            `bson:"patternExpression,omitempty" json:"patternExpression,omitempty" xml:"patternExpression,omitempty"`
	PatternParameterDefinition   *ParameterDefinition   `bson:"patternParameterDefinition,omitempty" json:"patternParameterDefinition,omitempty" xml:"patternParameterDefinition,omitempty"`
	PatternRelatedArtifact       *RelatedArtifact       `bson:"patternRelatedArtifact,omitempty" json:"patternRelatedArtifact,omitempty" xml:"patternRelatedArtifact,omitempty"`
	PatternTriggerDefinition     *TriggerDefinition     `bson:"patternTriggerDefinition,omitempty" json:"patternTriggerDefinition,omitempty" xml:"patternTriggerDefinition,omitempty"`
	PatternUsageContext          *UsageContext          `bson:"patternUsageContext,omitempty" json:"patternUsageContext,omitempty" xml:"patternUsageContext,omitempty"`
	PatternAvailability          *Availability          `bson:"patternAvailability,omitempty" json:"patternAvailability,omitempty" xml:"patternAvailability,omitempty"`
	PatternExtendedContactDetail *ExtendedContactDetail `bson:"patternExtendedContactDetail,omitempty" json:"patternExtendedContactDetail,omitempty" xml:"patternExtendedContactDetail,omitempty"`
	PatternDosage                *Dosage                `bson:"patternDosage,omitempty" json:"patternDosage,omitempty" xml:"patternDosage,omitempty"`
	PatternMeta                  *Meta                  `bson:"patternMeta,omitempty" json:"patternMeta,omitempty" xml:"patternMeta,omitempty"`
	// A sample value for this element demonstrating the type of information that would typically be found in the element.
	//
	// Comment: Examples will most commonly be present for data where it's not implicitly obvious from either the data type or value set what the values might be.  (I.e. Example values for dates or quantities would generally be unnecessary.)  If the example value is fully populated, the publication tool can generate an instance automatically.
	Example []ElementDefinitionExample `bson:"example,omitempty" json:"example,omitempty" xml:"example,omitempty"`
	// The minimum allowed value for the element.
	//
	// The value is inclusive. This is allowed for the types date, dateTime, instant, time, decimal, integer, and Quantity.
	//
	// Comment: Except for date/date/instant, the type of the minValue[x] SHALL be the same as the specified type of the element. For the date/dateTime/instant values, the type of minValue[x] SHALL be either the same, or a [Duration](datatypes.html#Duration) which specifies a relative time limit to the current time. The duration value is positive, and is subtracted from the current clock to determine the minimum allowable value.   A minimum value for a Quantity is interpreted as a canonical minimum - e.g. you cannot provide 100mg if the minimum value is 10g.
	MinValueDate        *string          `bson:"minValueDate,omitempty" json:"minValueDate,omitempty" xml:"minValueDate,omitempty"`
	MinValueDateTime    *string          `bson:"minValueDateTime,omitempty" json:"minValueDateTime,omitempty" xml:"minValueDateTime,omitempty"`
	MinValueInstant     *string          `bson:"minValueInstant,omitempty" json:"minValueInstant,omitempty" xml:"minValueInstant,omitempty"`
	MinValueTime        *string          `bson:"minValueTime,omitempty" json:"minValueTime,omitempty" xml:"minValueTime,omitempty"`
	MinValueDecimal     *decimal.Decimal `bson:"minValueDecimal,omitempty" json:"minValueDecimal,omitempty" xml:"minValueDecimal,omitempty"`
	MinValueInteger     *int             `bson:"minValueInteger,omitempty" json:"minValueInteger,omitempty" xml:"minValueInteger,omitempty"`
	MinValueInteger64   *int64           `bson:"minValueInteger64,omitempty" json:"minValueInteger64,omitempty,string" xml:"minValueInteger64,omitempty"`
	MinValuePositiveInt *int             `bson:"minValuePositiveInt,omitempty" json:"minValuePositiveInt,omitempty" xml:"minValuePositiveInt,omitempty"`
	MinValueUnsignedInt *int             `bson:"minValueUnsignedInt,omitempty" json:"minValueUnsignedInt,omitempty" xml:"minValueUnsignedInt,omitempty"`
	MinValueQuantity    *Quantity        `bson:"minValueQuantity,omitempty" json:"minValueQuantity,omitempty" xml:"minValueQuantity,omitempty"`
	// The maximum allowed value for the element.
	//
	// The value is inclusive. This is allowed for the types date, dateTime, instant, time, decimal, integer, and Quantity.
	//
	// Comment: Except for date/date/instant, the type of the maxValue[x] SHALL be the same as the specified type of the element. For the date/dateTime/instant values, the type of maxValue[x] SHALL be either the same, or a [Duration](datatypes.html#Duration) which specifies a relative time limit to the current time. The duration value is positive, and is added to the current clock to determine the maximum allowable value.   A maximum value for a Quantity is interpreted as a canonical maximum - e.g. you cannot provide 10g if the maximum value is 50mg.
	MaxValueDate        *string          `bson:"maxValueDate,omitempty" json:"maxValueDate,omitempty" xml:"maxValueDate,omitempty"`
	MaxValueDateTime    *string          `bson:"maxValueDateTime,omitempty" json:"maxValueDateTime,omitempty" xml:"maxValueDateTime,omitempty"`
	MaxValueInstant     *string          `bson:"maxValueInstant,omitempty" json:"maxValueInstant,omitempty" xml:"maxValueInstant,omitempty"`
	MaxValueTime        *string          `bson:"maxValueTime,omitempty" json:"maxValueTime,omitempty" xml:"maxValueTime,omitempty"`
	MaxValueDecimal     *decimal.Decimal `bson:"maxValueDecimal,omitempty" json:"maxValueDecimal,omitempty" xml:"maxValueDecimal,omitempty"`
	MaxValueInteger     *int             `bson:"maxValueInteger,omitempty" json:"maxValueInteger,omitempty" xml:"maxValueInteger,omitempty"`
	MaxValueInteger64   *int64           `bson:"maxValueInteger64,omitempty" json:"maxValueInteger64,omitempty,string" xml:"maxValueInteger64,omitempty"`
	MaxValuePositiveInt *int             `bson:"maxValuePositiveInt,omitempty" json:"maxValuePositiveInt,omitempty" xml:"maxValuePositiveInt,omitempty"`
	MaxValueUnsignedInt *int             `bson:"maxValueUnsignedInt,omitempty" json:"maxValueUnsignedInt,omitempty" xml:"maxValueUnsignedInt,omitempty"`
	MaxValueQuantity    *Quantity        `bson:"maxValueQuantity,omitempty" json:"maxValueQuantity,omitempty" xml:"maxValueQuantity,omitempty"`
	// Indicates the maximum length in characters that is permitted to be present in conformant instances and which is expected to be supported by conformant consumers that support the element.
	//
	// ```maxLength``` SHOULD only be used on primitive data types that have a string representation (see [http://hl7.org/fhir/StructureDefinition/structuredefinition-type-characteristics](http://hl7.org/fhir/extensions/StructureDefinition-structuredefinition-type-characteristics.html)).
	//
	// Comment: Receivers are not required to reject instances that exceed the maximum length.  The full length could be stored.  In some cases, data might be truncated, though truncation should be undertaken with care and an understanding of the consequences of doing so. If not specified, there is no conformance expectation for length support.
	MaxLength *int `bson:"maxLength,omitempty" json:"maxLength,omitempty" xml:"maxLength,omitempty"`
	// A reference to an invariant that may make additional statements about the cardinality or value in the instance.
	Condition []string `bson:"condition,omitempty" json:"condition,omitempty" xml:"condition,omitempty"`
	// Formal constraints such as co-occurrence and other constraints that can be computationally evaluated within the context of the instance.
	//
	// Comment: Constraints should be declared on the "context" element - the lowest element in the hierarchy that is common to all nodes referenced by the constraint.
	Constraint []ElementDefinitionConstraint `bson:"constraint,omitempty" json:"constraint,omitempty" xml:"constraint,omitempty"`
	// Specifies for a primitive data type that the value of the data type cannot be replaced by an extension.
	//
	// Comment: Typical extensions that can be provided in place of a the value are [http://hl7.org/fhir/StructureDefinition/data-absent-reason](http://hl7.org/fhir/extensions/StructureDefinition-data-absent-reason.html),  [http://hl7.org/fhir/StructureDefinition/iso21090-nullFlavor](http://hl7.org/fhir/extensions/StructureDefinition-iso21090-nullFlavor.html),  [http://hl7.org/fhir/StructureDefinition/originalText](http://hl7.org/fhir/extensions/StructureDefinition-originalText.html), and [http://hl7.org/fhir/StructureDefinition/cqf-expression](http://hl7.org/fhir/extensions/StructureDefinition-cqf-expression.html). Note that this element has no meaning when a value is provided, and extensions can be provided in addition to a value.
	MustHaveValue *bool `bson:"mustHaveValue,omitempty" json:"mustHaveValue,omitempty" xml:"mustHaveValue,omitempty"`
	// Specifies a list of extensions that can appear in place of a primitive value.
	//
	// Comment: if mustHaveValue is true, then this element should not be present, since no extensions are allowed in place of the value. Note that this element has no impact if the value is present.
	ValueAlternatives []string `bson:"valueAlternatives,omitempty" json:"valueAlternatives,omitempty" xml:"valueAlternatives,omitempty"`
	// If true, implementations that produce or consume resources SHALL provide "support" for the element in some meaningful way.
	//
	// Note that this is being phased out and replaced by obligations (see below).  If false, the element may be ignored and not supported. If false, whether to populate or use the data element in any way is at the discretion of the implementation.
	//
	// Comment: "Something meaningful" is context dependent and impossible to describe in the base FHIR specification. For this reason, the mustSupport flag is never set to true by the FHIR specification itself - it is only set to true in profiles.  A profile on a type can always make mustSupport = true if it is false in the base type but cannot make mustSupport = false if it is true in the base type.   This is done in [Resource Profiles](profiling.html#mustsupport), where the profile labels an element as mustSupport=true. When a profile does this, it SHALL also make clear exactly what kind of "support" is required, as this can mean many things.    Note that an element that has the property IsModifier is not necessarily a "key" element (e.g. one of the important elements to make use of the resource), nor is it automatically mustSupport - however both of these things are more likely to be true for IsModifier elements than for other elements.
	MustSupport *bool `bson:"mustSupport,omitempty" json:"mustSupport,omitempty" xml:"mustSupport,omitempty"`
	// If true, the value of this element affects the interpretation of the element or resource that contains it, and the value of the element cannot be ignored.
	//
	// Typically, this is used for status, negation and qualification codes. The effect of this is that the element cannot be ignored by systems: they SHALL either recognize the element and process it, and/or a pre-determination has been made that it is not relevant to their particular system. When used on the root element in an extension definition, this indicates whether or not the extension is a modifier extension.
	//
	// Comment: Only the definition of an element can set IsModifier true - either the specification itself or where an extension is originally defined. Once set, it cannot be changed in derived profiles (except in the special case of the defining a new extension). An element/extension that has isModifier=true SHOULD also have a minimum cardinality of 1, so that there is no lack of clarity about what to do if it is missing. If it can be missing, the definition SHALL make the meaning of a missing element clear.
	IsModifier *bool `bson:"isModifier,omitempty" json:"isModifier,omitempty" xml:"isModifier,omitempty"`
	// Explains how that element affects the interpretation of the resource or element that contains it.
	IsModifierReason *string `bson:"isModifierReason,omitempty" json:"isModifierReason,omitempty" xml:"isModifierReason,omitempty"`
	// Whether the element should be included if a client requests a search with the parameter _summary=true.
	//
	// Comment: Some resources include a set of simple metadata, and some very large data. This element is used to reduce the quantity of data returned in searches. Note that servers may pre-cache summarized resources for optimal performance. When a request is made with _summary=true, serializers only include elements marked as 'isSummary = true'. With a few exceptions (listed below), all datatype properties are included in the summary form. In resource and datatype definitions, if an element has a descendant marked as isSummary=true or if it satisfies both of the following conditions, it must be marked as isSummary=true:
	//
	// * is at the root or has a parent that is 'mustSupport'
	//
	// * has a minimum cardinality of 1 or is a modifier element
	//
	// The following datatype properties are exceptions, and are not marked as isSummary:
	//
	// * Attachment.data
	//
	// * Signature.data, Signature.targetFormat, Signature.sigFormat
	//
	// * Narrative.status, Narrative.div
	//
	// * SampledData.data.
	IsSummary *bool `bson:"isSummary,omitempty" json:"isSummary,omitempty" xml:"isSummary,omitempty"`
	// Binds to a value set if this element is coded (code, Coding, CodeableConcept, Quantity), or the data types (string, uri).
	//
	// Comment: For a CodeableConcept, when no codes are allowed - only text, use a binding of strength "required" with a description explaining that no coded values are allowed and what sort of information to put in the "text" element.
	Binding *ElementDefinitionBinding `bson:"binding,omitempty" json:"binding,omitempty" xml:"binding,omitempty"`
	// Identifies a concept from an external specification that roughly corresponds to this element.
	//
	// Comment: Mappings are not necessarily specific enough for safe translation.
	Mapping []ElementDefinitionMapping `bson:"mapping,omitempty" json:"mapping,omitempty" xml:"mapping,omitempty"`
}

type ElementDefinitionSlicing struct {
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
	// Designates which child elements are used to discriminate between the slices when processing an instance.
	//
	// If one or more discriminators are provided, the value of the child elements in the instance data SHALL completely distinguish which slice the element in the resource matches based on the allowed values for those elements in each of the slices.
	//
	// Comment: If there is no discriminator, the content is hard to process, so this should be avoided.
	Discriminator []ElementDefinitionSlicingDiscriminator `bson:"discriminator,omitempty" json:"discriminator,omitempty" xml:"discriminator,omitempty"`
	// A human-readable text description of how the slicing works.
	//
	// If there is no discriminator, this is required to be present to provide whatever information is possible about how the slices can be differentiated.
	//
	// Comment: If it's really not possible to differentiate them, the design should be re-evaluated to make the content usable.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// If the matching elements have to occur in the same order as defined in the profile.
	//
	// Comment: Order should only be required when it is a pressing concern for presentation. Profile authors should consider making the order a feature of the rules about the narrative, not the rules about the data - requiring ordered data makes the profile much less re-usable.
	Ordered *bool `bson:"ordered,omitempty" json:"ordered,omitempty" xml:"ordered,omitempty"`
	// Whether additional slices are allowed or not.
	//
	// When the slices are ordered, profile authors can also say that additional slices are only allowed at the end.
	//
	// Comment: Allowing additional elements makes for a much for flexible template - it's open for use in wider contexts, but also means that the content of the resource is not closed, and applications have to decide how to handle content not described by the profile.
	Rules SlicingRules `bson:"rules" json:"rules" xml:"rules"`
}

type ElementDefinitionSlicingDiscriminator struct {
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
	// How the element value is interpreted when discrimination is evaluated.
	//
	// Comment: 'pattern' is deprecated - it works exactly the same as 'value'.
	Type DiscriminatorType `bson:"type" json:"type" xml:"type"`
	// A FHIRPath expression, using [the simple subset of FHIRPath](fhirpath.html#simple), that is used to identify the element on which discrimination is based.
	Path string `bson:"path" json:"path" xml:"path"`
}

type ElementDefinitionBase struct {
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
	// The Path that identifies the base element - this matches the ElementDefinition.path for that element.
	//
	// Across FHIR, there is only one base definition of any element - that is, an element definition on a [StructureDefinition](structuredefinition.html#) without a StructureDefinition.base.
	Path string `bson:"path" json:"path" xml:"path"`
	// Minimum cardinality of the base element identified by the path.
	//
	// Comment: This is provided for consistency with max, and may affect code generation of mandatory elements of the base resource are generated differently (some reference implementations have done this).
	Min int `bson:"min" json:"min" xml:"min"`
	// Maximum cardinality of the base element identified by the path.
	//
	// Comment: This is provided to code generation, since the serialization representation in JSON differs depending on whether the base element has max > 1. Also, some forms of code generation may differ.
	Max string `bson:"max" json:"max" xml:"max"`
}

type ElementDefinitionType struct {
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
	// URL of Data type or Resource that is a(or the) type used for this element.
	//
	// References are URLs that are relative to http://hl7.org/fhir/StructureDefinition e.g. "string" is a reference to http://hl7.org/fhir/StructureDefinition/string. Absolute URLs are only allowed in logical models.
	//
	// Comment: If the element is a reference to another resource, this element contains "Reference", and the targetProfile element defines what resources can be referenced. The targetProfile may be a reference to the general definition of a resource (e.g. http://hl7.org/fhir/StructureDefinition/Patient).
	Code string `bson:"code" json:"code" xml:"code"`
	// Identifies a profile structure or implementation Guide that applies to the datatype this element refers to.
	//
	// If any profiles are specified, then the content must conform to at least one of them. The URL can be a local reference - to a contained StructureDefinition, or a reference to another StructureDefinition or Implementation Guide by a canonical URL. When an implementation guide is specified, the type SHALL conform to at least one profile defined in the implementation guide.
	//
	// Comment: It is possible to profile  backbone element (e.g. part of a resource), using the [http://hl7.org/fhir/StructureDefinition/elementdefinition-profile-element](http://hl7.org/fhir/extensions/StructureDefinition-elementdefinition-profile-element.html) extension.
	Profile []string `bson:"profile,omitempty" json:"profile,omitempty" xml:"profile,omitempty"`
	// Used when the type is "Reference" or "canonical", and identifies a profile structure or implementation Guide that applies to the target of the reference this element refers to.
	//
	// If any profiles are specified, then the content must conform to at least one of them. The URL can be a local reference - to a contained StructureDefinition, or a reference to another StructureDefinition or Implementation Guide by a canonical URL. When an implementation guide is specified, the target resource SHALL conform to at least one profile defined in the implementation guide.
	TargetProfile []string `bson:"targetProfile,omitempty" json:"targetProfile,omitempty" xml:"targetProfile,omitempty"`
	// If the type is a reference to another resource, how the resource is or can be aggregated - is it a contained resource, or a reference, and if the context is a bundle, is it included in the bundle.
	//
	// Comment: See [Aggregation Rules](elementdefinition.html#aggregation) for further clarification.
	Aggregation []AggregationMode `bson:"aggregation,omitempty" json:"aggregation,omitempty" xml:"aggregation,omitempty"`
	// Whether this reference needs to be version specific or version independent, or whether either can be used.
	//
	// Comment: The base specification never makes a rule as to which form is allowed, but implementation guides may do this. See [Aggregation Rules](elementdefinition.html#aggregation) for further clarification.
	Versioning *ReferenceVersionRules `bson:"versioning,omitempty" json:"versioning,omitempty" xml:"versioning,omitempty"`
}

type ElementDefinitionExample struct {
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
	// Describes the purpose of this example among the set of examples.
	Label string `bson:"label" json:"label" xml:"label"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueBase64Binary          string                `bson:"valueBase64Binary" json:"valueBase64Binary" xml:"valueBase64Binary"`
	ValueBoolean               bool                  `bson:"valueBoolean" json:"valueBoolean" xml:"valueBoolean"`
	ValueCanonical             string                `bson:"valueCanonical" json:"valueCanonical" xml:"valueCanonical"`
	ValueCode                  string                `bson:"valueCode" json:"valueCode" xml:"valueCode"`
	ValueDate                  string                `bson:"valueDate" json:"valueDate" xml:"valueDate"`
	ValueDateTime              string                `bson:"valueDateTime" json:"valueDateTime" xml:"valueDateTime"`
	ValueDecimal               decimal.Decimal       `bson:"valueDecimal" json:"valueDecimal" xml:"valueDecimal"`
	ValueID                    string                `bson:"valueId" json:"valueId" xml:"valueId"`
	ValueInstant               string                `bson:"valueInstant" json:"valueInstant" xml:"valueInstant"`
	ValueInteger               int                   `bson:"valueInteger" json:"valueInteger" xml:"valueInteger"`
	ValueInteger64             int64                 `bson:"valueInteger64" json:"valueInteger64,string" xml:"valueInteger64"`
	ValueMarkdown              string                `bson:"valueMarkdown" json:"valueMarkdown" xml:"valueMarkdown"`
	ValueOid                   string                `bson:"valueOid" json:"valueOid" xml:"valueOid"`
	ValuePositiveInt           int                   `bson:"valuePositiveInt" json:"valuePositiveInt" xml:"valuePositiveInt"`
	ValueString                string                `bson:"valueString" json:"valueString" xml:"valueString"`
	ValueTime                  string                `bson:"valueTime" json:"valueTime" xml:"valueTime"`
	ValueUnsignedInt           int                   `bson:"valueUnsignedInt" json:"valueUnsignedInt" xml:"valueUnsignedInt"`
	ValueUri                   string                `bson:"valueUri" json:"valueUri" xml:"valueUri"`
	ValueUURL                  string                `bson:"valueUrl" json:"valueUrl" xml:"valueUrl"`
	ValueUuid                  string                `bson:"valueUuid" json:"valueUuid" xml:"valueUuid"`
	ValueAddress               Address               `bson:"valueAddress" json:"valueAddress" xml:"valueAddress"`
	ValueAge                   Age                   `bson:"valueAge" json:"valueAge" xml:"valueAge"`
	ValueAnnotation            Annotation            `bson:"valueAnnotation" json:"valueAnnotation" xml:"valueAnnotation"`
	ValueAttachment            Attachment            `bson:"valueAttachment" json:"valueAttachment" xml:"valueAttachment"`
	ValueCodeableConcept       CodeableConcept       `bson:"valueCodeableConcept" json:"valueCodeableConcept" xml:"valueCodeableConcept"`
	ValueCodeableReference     CodeableReference     `bson:"valueCodeableReference" json:"valueCodeableReference" xml:"valueCodeableReference"`
	ValueCoding                Coding                `bson:"valueCoding" json:"valueCoding" xml:"valueCoding"`
	ValueContactPoint          ContactPoint          `bson:"valueContactPoint" json:"valueContactPoint" xml:"valueContactPoint"`
	ValueCount                 Count                 `bson:"valueCount" json:"valueCount" xml:"valueCount"`
	ValueDistance              Distance              `bson:"valueDistance" json:"valueDistance" xml:"valueDistance"`
	ValueDuration              Duration              `bson:"valueDuration" json:"valueDuration" xml:"valueDuration"`
	ValueHumanName             HumanName             `bson:"valueHumanName" json:"valueHumanName" xml:"valueHumanName"`
	ValueIdentifier            Identifier            `bson:"valueIdentifier" json:"valueIdentifier" xml:"valueIdentifier"`
	ValueMoney                 Money                 `bson:"valueMoney" json:"valueMoney" xml:"valueMoney"`
	ValuePeriod                Period                `bson:"valuePeriod" json:"valuePeriod" xml:"valuePeriod"`
	ValueQuantity              Quantity              `bson:"valueQuantity" json:"valueQuantity" xml:"valueQuantity"`
	ValueRange                 Range                 `bson:"valueRange" json:"valueRange" xml:"valueRange"`
	ValueRatio                 Ratio                 `bson:"valueRatio" json:"valueRatio" xml:"valueRatio"`
	ValueRatioRange            RatioRange            `bson:"valueRatioRange" json:"valueRatioRange" xml:"valueRatioRange"`
	ValueReference             Reference             `bson:"valueReference" json:"valueReference" xml:"valueReference"`
	ValueSampledData           SampledData           `bson:"valueSampledData" json:"valueSampledData" xml:"valueSampledData"`
	ValueSignature             Signature             `bson:"valueSignature" json:"valueSignature" xml:"valueSignature"`
	ValueTiming                Timing                `bson:"valueTiming" json:"valueTiming" xml:"valueTiming"`
	ValueContactDetail         ContactDetail         `bson:"valueContactDetail" json:"valueContactDetail" xml:"valueContactDetail"`
	ValueDataRequirement       DataRequirement       `bson:"valueDataRequirement" json:"valueDataRequirement" xml:"valueDataRequirement"`
	ValueExpression            Expression            `bson:"valueExpression" json:"valueExpression" xml:"valueExpression"`
	ValueParameterDefinition   ParameterDefinition   `bson:"valueParameterDefinition" json:"valueParameterDefinition" xml:"valueParameterDefinition"`
	ValueRelatedArtifact       RelatedArtifact       `bson:"valueRelatedArtifact" json:"valueRelatedArtifact" xml:"valueRelatedArtifact"`
	ValueTriggerDefinition     TriggerDefinition     `bson:"valueTriggerDefinition" json:"valueTriggerDefinition" xml:"valueTriggerDefinition"`
	ValueUsageContext          UsageContext          `bson:"valueUsageContext" json:"valueUsageContext" xml:"valueUsageContext"`
	ValueAvailability          Availability          `bson:"valueAvailability" json:"valueAvailability" xml:"valueAvailability"`
	ValueExtendedContactDetail ExtendedContactDetail `bson:"valueExtendedContactDetail" json:"valueExtendedContactDetail" xml:"valueExtendedContactDetail"`
	ValueDosage                Dosage                `bson:"valueDosage" json:"valueDosage" xml:"valueDosage"`
	ValueMeta                  Meta                  `bson:"valueMeta" json:"valueMeta" xml:"valueMeta"`
}

type ElementDefinitionConstraint struct {
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
	// Allows identification of which elements have their cardinalities impacted by the constraint.
	//
	//  Will not be referenced for constraints that do not affect cardinality.
	Key string `bson:"key" json:"key" xml:"key"`
	// Description of why this constraint is necessary or appropriate.
	//
	// Comment: To be used if the reason for the constraint might not be intuitive to all implementers.
	Requirements *string `bson:"requirements,omitempty" json:"requirements,omitempty" xml:"requirements,omitempty"`
	// Identifies the impact constraint violation has on the conformance of the instance.
	//
	// Comment: This allows constraints to be asserted as "shall" (error) and "should" (warning).
	Severity ConstraintSeverity `bson:"severity" json:"severity" xml:"severity"`
	// If true, indicates that the warning or best practice guideline should be suppressed.
	//
	// Comment: This element should only be present in a derived profile where a warning or hint has been determined to be spurious/incorrect. E.g. a warning encouraging the inclusion of an element that the profile explicitly prohibits.
	Suppress *bool `bson:"suppress,omitempty" json:"suppress,omitempty" xml:"suppress,omitempty"`
	// Text that can be used to describe the constraint in messages identifying that the constraint has been violated.
	//
	// Comment: Should be expressed in business terms as much as possible.
	Human string `bson:"human" json:"human" xml:"human"`
	// A [FHIRPath](fhirpath.html) expression of constraint that can be executed to see if this constraint is met.
	//
	// Comment: In the absense of an expression, the expression is likely not enforceable by validators, and might be missed by many systems.
	Expression *string `bson:"expression,omitempty" json:"expression,omitempty" xml:"expression,omitempty"`
	// A reference to the original source of the constraint, for traceability purposes.
	//
	// Comment: This is used when, e.g. rendering, where it is not useful to present inherited constraints when rendering the snapshot.
	Source *string `bson:"source,omitempty" json:"source,omitempty" xml:"source,omitempty"`
}

type ElementDefinitionBinding struct {
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
	// Indicates the degree of conformance expectations associated with this binding - that is, the degree to which the provided value set must be adhered to in the instances.
	//
	// Comment: For further discussion, see [Using Terminologies](terminologies.html).
	Strength BindingStrength `bson:"strength" json:"strength" xml:"strength"`
	// Describes the intended use of this particular set of codes.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Refers to the value set that identifies the set of codes the binding refers to.
	//
	// Comment: The reference may be version-specific or not (e.g. have a |[version] at the end of the canonical URL).
	ValueSet *string `bson:"valueSet,omitempty" json:"valueSet,omitempty" xml:"valueSet,omitempty"`
	// Additional bindings that help applications implementing this element.
	//
	// Additional bindings do not replace the main binding but provide more information and/or context.
	Additional []ElementDefinitionBindingAdditional `bson:"additional,omitempty" json:"additional,omitempty" xml:"additional,omitempty"`
}

type ElementDefinitionBindingAdditional struct {
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
	// The use of this additional binding.
	//
	// Comment: Conformance bindings are in addition to the base binding, not instead of it.
	Purpose AdditionalBindingPurposeVS `bson:"purpose" json:"purpose" xml:"purpose"`
	// The valueSet that is being bound for the purpose.
	ValueSet string `bson:"valueSet" json:"valueSet" xml:"valueSet"`
	// Documentation of the purpose of use of the bindingproviding additional information about how it is intended to be used.
	Documentation *string `bson:"documentation,omitempty" json:"documentation,omitempty" xml:"documentation,omitempty"`
	// Concise documentation - for summary tables.
	ShortDoco *string `bson:"shortDoco,omitempty" json:"shortDoco,omitempty" xml:"shortDoco,omitempty"`
	// Qualifies the usage of the binding.
	//
	// Typically bindings are qualified by jurisdiction, but they may also be qualified by gender, workflow status, clinical domain etc. The information to decide whether a usege context applies is usually outside the resource, determined by context, and this might present challenges for validation tooling.
	//
	// Comment: This specification does not define exactly how the context is determined for the usage; typically, this will be described in implementation guides.
	Usage []UsageContext `bson:"usage,omitempty" json:"usage,omitempty" xml:"usage,omitempty"`
	// Whether the binding applies to all repeats, or just to any one of them.
	//
	// This is only relevant for elements that can repeat.
	//
	// Comment: A common pattern is to make a conformance binding to all repeats, and then make a binding that one of the repeats must conform to.
	Any *bool `bson:"any,omitempty" json:"any,omitempty" xml:"any,omitempty"`
}

type ElementDefinitionMapping struct {
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
	// An internal reference to the definition of a mapping.
	Identity string `bson:"identity" json:"identity" xml:"identity"`
	// Identifies the computable language in which mapping.map is expressed.
	//
	// Comment: If omitted, then there can be no expectation of computational interpretation of the mapping.
	Language *string `bson:"language,omitempty" json:"language,omitempty" xml:"language,omitempty"`
	// Expresses what part of the target specification corresponds to this element.
	//
	// Comment: For most mappings, the syntax is undefined.  Syntax will be provided for mappings to the RIM.  Multiple mappings may be possible and may include constraints on other resource elements that identify when a particular mapping applies.
	Map string `bson:"map" json:"map" xml:"map"`
	// Comments that provide information about the mapping or its use.
	Comment *string `bson:"comment,omitempty" json:"comment,omitempty" xml:"comment,omitempty"`
}
