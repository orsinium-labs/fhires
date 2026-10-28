package fhires

import "encoding/json/v2"

// The MeasureReport resource contains the results of the calculation of a measure; and optionally a reference to the resources involved in that calculation.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/MeasureReport
type MeasureReport struct {
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
	// A formal identifier that is used to identify this MeasureReport when it is represented in other formats or referenced in a specification, model, design or an instance.
	//
	// Comment: Typically, this is used for identifiers that can go in an HL7 V3 II data type - e.g. to identify this {{title}} outside of FHIR, where the logical URL is not possible to use.
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The MeasureReport status.
	//
	// No data will be available until the MeasureReport status is complete.
	//
	// Comment: This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid.
	Status MeasureReportStatus `bson:"status" json:"status" xml:"status"`
	// The type of measure report.
	//
	// This may be an individual report, which provides the score for the measure for an individual member of the population; a subject-listing, which returns the list of members that meet the various criteria in the measure; a summary report, which returns a population count for each of the criteria in the measure; or a data-collection, which enables the MeasureReport to be used to exchange the data-of-interest for a quality measure.
	//
	// Comment: Data-exchange reports are used only to communicate data-of-interest for a measure. They do not necessarily include all the data for a particular subject or population, but they may.
	Type MeasureReportType `bson:"type" json:"type" xml:"type"`
	// Indicates whether the data submitted in a data-exchange report represents a snapshot or incremental update.
	//
	// A snapshot update replaces all previously submitted data for the receiver, whereas an incremental update represents only updated and/or changed data and should be applied as a differential update to the existing submitted data for the receiver.
	//
	// Comment: This element only applies to Data-collection reports used to communicate the data-of-interest for a measure
	DataUpdateType *SubmitDataUpdateType `bson:"dataUpdateType,omitempty" json:"dataUpdateType,omitempty" xml:"dataUpdateType,omitempty"`
	// A reference to the Measure that was calculated to produce this report.
	Measure *string `bson:"measure,omitempty" json:"measure,omitempty" xml:"measure,omitempty"`
	// Optional subject identifying the individual or individuals the report is for.
	Subject *Reference `bson:"subject,omitempty" json:"subject,omitempty" xml:"subject,omitempty"`
	// The date this measure was calculated.
	Date *string `bson:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// The individual or organization that is reporting the data.
	Reporter *Reference `bson:"reporter,omitempty" json:"reporter,omitempty" xml:"reporter,omitempty"`
	// A reference to the vendor who queried the data, calculated results and/or generated the report.
	//
	// The ‘reporting vendor’ is intended to represent the submitting entity when it is not the same as the reporting entity. This extension is used when the Receiver is interested in getting vendor information in the report.
	ReportingVendor *Reference `bson:"reportingVendor,omitempty" json:"reportingVendor,omitempty" xml:"reportingVendor,omitempty"`
	// A reference to the location for which the data is being reported.
	Location *Reference `bson:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// The reporting period for which the report was calculated.
	Period Period `bson:"period" json:"period" xml:"period"`
	// A reference to a Parameters resource (typically represented using a contained resource) that represents any input parameters that were provided to the operation that generated the report.
	//
	// Comment: Whenever an operation returns a MeasureReport, the IN parameters from that operation can be returned to the caller of the operation using MeasureReport.inputParameters. Operations include but are not limited to $evaluate-measure, $care-gaps, $data-requirements, and $collect-data. The inputParameters can also be used in a MeasureReport used to define a test case (i.e., not as a result of an operation).
	InputParameters *Reference `bson:"inputParameters,omitempty" json:"inputParameters,omitempty" xml:"inputParameters,omitempty"`
	// Indicates how the calculation is performed for the measure, including proportion, ratio, continuous-variable, and cohort.
	//
	// The value set is extensible, allowing additional measure scoring types to be represented. It is expected to be the same as the scoring element on the referenced Measure.
	Scoring *CodeableConcept `bson:"scoring,omitempty" json:"scoring,omitempty" xml:"scoring,omitempty"`
	// Whether improvement in the measure is noted by an increase or decrease in the measure score.
	//
	// Comment: This element is typically defined by the measure, but reproduced here to ensure the measure score can be interpreted. The element is labeled as a modifier because it changes the interpretation of the reported measure score. Note also that a MeasureReport instance includes the improvementNotation as defined by the Measure being reported. It is duplicated in the MeasureReport because it is a critical aspect of interpreting the measure score but it is not intended to reflect whether the measure report is an increase or decrease. It helps interpret if the measure score is an increase or decrease, I.e., moving in the direction of the desired outcome.
	ImprovementNotation *CodeableConcept `bson:"improvementNotation,omitempty" json:"improvementNotation,omitempty" xml:"improvementNotation,omitempty"`
	// The results of the calculation, one for each population group in the measure.
	Group []MeasureReportGroup `bson:"group,omitempty" json:"group,omitempty" xml:"group,omitempty"`
	// A reference to a Resource that represents additional information collected for the report.
	//
	// If the value of the supplemental data is not a Resource (i.e. evaluating the supplementalData expression for this case in the measure results in a value that is not a FHIR Resource), it is reported as a reference to a contained Observation resource.
	//
	// Comment: For individual measure reports, the supplementalData elements represent the direct result of evaluating the supplementalData expression for the subject of the report. For summary reports, supplementalData is reported as references to contained Observation resources that represent the number of times each value appeared as a supplementalData result for members of the population.
	SupplementalData []Reference `bson:"supplementalData,omitempty" json:"supplementalData,omitempty" xml:"supplementalData,omitempty"`
	// Evaluated resources are used to capture what data was involved in the calculation of a measure.
	//
	// This usage is only allowed for individual reports to ensure that the size of the MeasureReport resource is bounded.
	//
	// Comment: Evaluated resources are only reported for individual reports.
	EvaluatedResource []Reference `bson:"evaluatedResource,omitempty" json:"evaluatedResource,omitempty" xml:"evaluatedResource,omitempty"`
}

type MeasureReportGroup struct {
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
	// The group from the Measure that corresponds to this group in the MeasureReport resource.
	LinkID *string `bson:"linkId,omitempty" json:"linkId,omitempty" xml:"linkId,omitempty"`
	// The meaning of the population group as defined in the measure definition.
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// Optional subject identifying the individual or individuals the report is for.
	Subject *Reference `bson:"subject,omitempty" json:"subject,omitempty" xml:"subject,omitempty"`
	// The populations that make up the population group, one for each type of population appropriate for the measure.
	Population []MeasureReportGroupPopulation `bson:"population,omitempty" json:"population,omitempty" xml:"population,omitempty"`
	// The measure score for this population group, calculated as appropriate for the measure type and scoring method, and based on the contents of the populations defined in the group.
	MeasureScoreQuantity        *Quantity        `bson:"measureScoreQuantity,omitempty" json:"measureScoreQuantity,omitempty" xml:"measureScoreQuantity,omitempty"`
	MeasureScoreDateTime        *string          `bson:"measureScoreDateTime,omitempty" json:"measureScoreDateTime,omitempty" xml:"measureScoreDateTime,omitempty"`
	MeasureScoreCodeableConcept *CodeableConcept `bson:"measureScoreCodeableConcept,omitempty" json:"measureScoreCodeableConcept,omitempty" xml:"measureScoreCodeableConcept,omitempty"`
	MeasureScorePeriod          *Period          `bson:"measureScorePeriod,omitempty" json:"measureScorePeriod,omitempty" xml:"measureScorePeriod,omitempty"`
	MeasureScoreRange           *Range           `bson:"measureScoreRange,omitempty" json:"measureScoreRange,omitempty" xml:"measureScoreRange,omitempty"`
	MeasureScoreDuration        *Duration        `bson:"measureScoreDuration,omitempty" json:"measureScoreDuration,omitempty" xml:"measureScoreDuration,omitempty"`
	// When a measure includes multiple stratifiers, there will be a stratifier group for each stratifier defined by the measure.
	Stratifier []MeasureReportGroupStratifier `bson:"stratifier,omitempty" json:"stratifier,omitempty" xml:"stratifier,omitempty"`
}

type MeasureReportGroupPopulation struct {
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
	// The population from the Measure that corresponds to this population in the MeasureReport resource.
	LinkID *string `bson:"linkId,omitempty" json:"linkId,omitempty" xml:"linkId,omitempty"`
	// The type of the population.
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// The number of members of the population.
	Count *int `bson:"count,omitempty" json:"count,omitempty" xml:"count,omitempty"`
	// This element refers to a List of individual level MeasureReport resources, one for each subject in this population.
	SubjectResults *Reference `bson:"subjectResults,omitempty" json:"subjectResults,omitempty" xml:"subjectResults,omitempty"`
	// A reference to an individual level MeasureReport resource for a member of the population.
	//
	// Comment: This element SHALL NOT be used in with subjectResults, either systems provide results as a List of MeasureReport, or as references to individual MeasureReports. Note that using this element introduces a constraint on the size of the subject-list MeasureReport; implementations would be free to return a TOO-COSTLY response if a request is made for a subject-list report with too many subjects.
	SubjectReport []Reference `bson:"subjectReport,omitempty" json:"subjectReport,omitempty" xml:"subjectReport,omitempty"`
	// Optional Group identifying the individuals that make up the population.
	//
	// Comment: Note that this element can only be used for summary and subject-list reports where the basis is of the same type as the populations (e.g. a Patient-based measure, rather than an Encounter-based measure). In addition, the referenced Group resource SHALL be an actual (as opposed to a definitional) group.
	Subjects *Reference `bson:"subjects,omitempty" json:"subjects,omitempty" xml:"subjects,omitempty"`
}

type MeasureReportGroupStratifier struct {
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
	// The stratifier from the Measure that corresponds to this stratifier in the MeasureReport resource.
	LinkID *string `bson:"linkId,omitempty" json:"linkId,omitempty" xml:"linkId,omitempty"`
	// The meaning of this stratifier, as defined in the measure definition.
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// This element contains the results for a single stratum within the stratifier.
	//
	// For example, when stratifying on administrative gender, there will be four strata, one for each possible gender value.
	Stratum []MeasureReportGroupStratifierStratum `bson:"stratum,omitempty" json:"stratum,omitempty" xml:"stratum,omitempty"`
}

type MeasureReportGroupStratifierStratum struct {
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
	// The value for this stratum, expressed as a CodeableConcept.
	//
	// When defining stratifiers on complex values, the value must be rendered such that the value for each stratum within the stratifier is unique.
	ValueCodeableConcept *CodeableConcept `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty" xml:"valueCodeableConcept,omitempty"`
	ValueBoolean         *bool            `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty" xml:"valueBoolean,omitempty"`
	ValueQuantity        *Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty" xml:"valueQuantity,omitempty"`
	ValueRange           *Range           `bson:"valueRange,omitempty" json:"valueRange,omitempty" xml:"valueRange,omitempty"`
	ValueReference       *Reference       `bson:"valueReference,omitempty" json:"valueReference,omitempty" xml:"valueReference,omitempty"`
	// A stratifier component value.
	Component []MeasureReportGroupStratifierStratumComponent `bson:"component,omitempty" json:"component,omitempty" xml:"component,omitempty"`
	// The populations that make up the stratum, one for each type of population appropriate to the measure.
	Population []MeasureReportGroupStratifierStratumPopulation `bson:"population,omitempty" json:"population,omitempty" xml:"population,omitempty"`
	// The measure score for this stratum, calculated as appropriate for the measure type and scoring method, and based on only the members of this stratum.
	MeasureScoreQuantity        *Quantity        `bson:"measureScoreQuantity,omitempty" json:"measureScoreQuantity,omitempty" xml:"measureScoreQuantity,omitempty"`
	MeasureScoreDateTime        *string          `bson:"measureScoreDateTime,omitempty" json:"measureScoreDateTime,omitempty" xml:"measureScoreDateTime,omitempty"`
	MeasureScoreCodeableConcept *CodeableConcept `bson:"measureScoreCodeableConcept,omitempty" json:"measureScoreCodeableConcept,omitempty" xml:"measureScoreCodeableConcept,omitempty"`
	MeasureScorePeriod          *Period          `bson:"measureScorePeriod,omitempty" json:"measureScorePeriod,omitempty" xml:"measureScorePeriod,omitempty"`
	MeasureScoreRange           *Range           `bson:"measureScoreRange,omitempty" json:"measureScoreRange,omitempty" xml:"measureScoreRange,omitempty"`
	MeasureScoreDuration        *Duration        `bson:"measureScoreDuration,omitempty" json:"measureScoreDuration,omitempty" xml:"measureScoreDuration,omitempty"`
}

type MeasureReportGroupStratifierStratumComponent struct {
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
	// The stratifier component from the Measure that corresponds to this stratifier component in the MeasureReport resource.
	LinkID *string `bson:"linkId,omitempty" json:"linkId,omitempty" xml:"linkId,omitempty"`
	// The code for the stratum component value.
	Code CodeableConcept `bson:"code" json:"code" xml:"code"`
	// The stratum component value.
	ValueCodeableConcept CodeableConcept `bson:"valueCodeableConcept" json:"valueCodeableConcept" xml:"valueCodeableConcept"`
	ValueBoolean         bool            `bson:"valueBoolean" json:"valueBoolean" xml:"valueBoolean"`
	ValueQuantity        Quantity        `bson:"valueQuantity" json:"valueQuantity" xml:"valueQuantity"`
	ValueRange           Range           `bson:"valueRange" json:"valueRange" xml:"valueRange"`
	ValueReference       Reference       `bson:"valueReference" json:"valueReference" xml:"valueReference"`
}

type MeasureReportGroupStratifierStratumPopulation struct {
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
	// The population from the Measure that corresponds to this population in the MeasureReport resource.
	LinkID *string `bson:"linkId,omitempty" json:"linkId,omitempty" xml:"linkId,omitempty"`
	// The type of the population.
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// The number of members of the population in this stratum.
	Count *int `bson:"count,omitempty" json:"count,omitempty" xml:"count,omitempty"`
	// This element refers to a List of individual level MeasureReport resources, one for each subject in this population in this stratum.
	SubjectResults *Reference `bson:"subjectResults,omitempty" json:"subjectResults,omitempty" xml:"subjectResults,omitempty"`
	// A reference to an individual level MeasureReport resource for a member of the population.
	//
	// Comment: This element SHALL NOT be used in with subjectResults, either systems provide results as a List of MeasureReport, or as references to individual MeasureReports. Note that using this element introduces a constraint on the size of the subject-list MeasureReport; implementations would be free to return a TOO-COSTLY response if a request is made for a subject-list report with too many subjects.
	SubjectReport []Reference `bson:"subjectReport,omitempty" json:"subjectReport,omitempty" xml:"subjectReport,omitempty"`
	// Optional Group identifying the individuals that make up the population.
	//
	// Comment: Note that this element can only be used for summary and subject-list reports where the basis is of the same type as the populations (e.g. a Patient-based measure, rather than an Encounter-based measure). In addition, the referenced Group SHALL be an actual (as opposed to a definitional) group.
	Subjects *Reference `bson:"subjects,omitempty" json:"subjects,omitempty" xml:"subjects,omitempty"`
}

func (r MeasureReport) ResourceType() string {
	return "MeasureReport"
}

// MarshalJSON implements json.Marshaler
func (r MeasureReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		MeasureReport
		ResourceType string `json:"resourceType"`
	}{
		MeasureReport: r,
		ResourceType:  r.ResourceType(),
	})
}
