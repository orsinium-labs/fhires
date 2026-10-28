package fhires

// DataRequirement Type: Describes a required data item for evaluation in terms of the type of data, and optional code or date-based filters of the data.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/DataRequirement
type DataRequirement struct {
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
	// The type of the required data, specified as the type name of a resource.
	//
	// For profiles, this value is set to the type of the base resource of the profile.
	Type FHIRTypes `bson:"type" json:"type" xml:"type"`
	// The profile of the required data, specified as the uri of the profile definition.
	Profile []string `bson:"profile,omitempty" json:"profile,omitempty" xml:"profile,omitempty"`
	// The intended subjects of the data requirement.
	//
	// If this element is not provided, a Patient subject is assumed.
	//
	// Comment: The subject of a data requirement is critical, as the data being specified is determined with respect to a particular subject. This corresponds roughly to the notion of a Compartment in that it limits what data is available based on its relationship to the subject. In CQL, this corresponds to the context declaration.
	SubjectCodeableConcept *CodeableConcept `bson:"subjectCodeableConcept,omitempty" json:"subjectCodeableConcept,omitempty" xml:"subjectCodeableConcept,omitempty"`
	SubjectReference       *Reference       `bson:"subjectReference,omitempty" json:"subjectReference,omitempty" xml:"subjectReference,omitempty"`
	// Indicates that specific elements of the type are referenced by the knowledge module and must be supported by the consumer in order to obtain an effective evaluation.
	//
	// This does not mean that a value is required for this element, only that the consuming system must understand the element and be able to provide values for it if they are available.
	//
	// The value of mustSupport SHALL be a FHIRPath resolvable on the type of the DataRequirement. The path SHALL consist only of identifiers, constant indexers, and .resolve() (see the [Simple FHIRPath Profile](fhirpath.html#simple) for full details).
	MustSupport []string `bson:"mustSupport,omitempty" json:"mustSupport,omitempty" xml:"mustSupport,omitempty"`
	// Code filters specify additional constraints on the data, specifying the value set of interest for a particular element of the data.
	//
	// Each code filter defines an additional constraint on the data, i.e. code filters are AND'ed, not OR'ed.
	CodeFilter []DataRequirementCodeFilter `bson:"codeFilter,omitempty" json:"codeFilter,omitempty" xml:"codeFilter,omitempty"`
	// Date filters specify additional constraints on the data in terms of the applicable date range for specific elements.
	//
	// Each date filter specifies an additional constraint on the data, i.e. date filters are AND'ed, not OR'ed.
	DateFilter []DataRequirementDateFilter `bson:"dateFilter,omitempty" json:"dateFilter,omitempty" xml:"dateFilter,omitempty"`
	// Value filters specify additional constraints on the data for elements other than code-valued or date-valued.
	//
	// Each value filter specifies an additional constraint on the data (i.e. valueFilters are AND'ed, not OR'ed).
	ValueFilter []DataRequirementValueFilter `bson:"valueFilter,omitempty" json:"valueFilter,omitempty" xml:"valueFilter,omitempty"`
	// Specifies a maximum number of results that are required (uses the _count search parameter).
	//
	// Comment: This element can be used in combination with the sort element to specify quota requirements such as "the most recent 5" or "the highest 5".
	Limit *int `bson:"limit,omitempty" json:"limit,omitempty" xml:"limit,omitempty"`
	// Specifies the order of the results to be returned.
	//
	// Comment: This element can be used in combination with the sort element to specify quota requirements such as "the most recent 5" or "the highest 5". When multiple sorts are specified, they are applied in the order they appear in the resource.
	Sort []DataRequirementSort `bson:"sort,omitempty" json:"sort,omitempty" xml:"sort,omitempty"`
}

type DataRequirementCodeFilter struct {
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
	// The code-valued attribute of the filter.
	//
	// The specified path SHALL be a FHIRPath resolvable on the specified type of the DataRequirement, and SHALL consist only of identifiers, constant indexers, and .resolve(). The path is allowed to contain qualifiers (.) to traverse sub-elements, as well as indexers ([x]) to traverse multiple-cardinality sub-elements (see the [Simple FHIRPath Profile](fhirpath.html#simple) for full details). Note that the index must be an integer constant. The path must resolve to an element of type code, Coding, or CodeableConcept.
	//
	// Comment: The path attribute contains a [Simple FHIRPath Subset](fhirpath.html#simple) that allows path traversal, but not calculation.
	Path *string `bson:"path,omitempty" json:"path,omitempty" xml:"path,omitempty"`
	// A token parameter that refers to a search parameter defined on the specified type of the DataRequirement, and which searches on elements of type code, Coding, or CodeableConcept.
	SearchParam *string `bson:"searchParam,omitempty" json:"searchParam,omitempty" xml:"searchParam,omitempty"`
	// The valueset for the code filter.
	//
	// The valueSet and code elements are additive. If valueSet is specified, the filter will return only those data items for which the value of the code-valued element specified in the path is a member of the specified valueset.
	ValueSet *string `bson:"valueSet,omitempty" json:"valueSet,omitempty" xml:"valueSet,omitempty"`
	// The codes for the code filter.
	//
	// If values are given, the filter will return only those data items for which the code-valued attribute specified by the path has a value that is one of the specified codes. If codes are specified in addition to a value set, the filter returns items matching a code in the value set or one of the specified codes.
	Code []Coding `bson:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
}

type DataRequirementDateFilter struct {
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
	// The date-valued attribute of the filter.
	//
	// The specified path SHALL be a FHIRPath resolvable on the specified type of the DataRequirement, and SHALL consist only of identifiers, constant indexers, and .resolve(). The path is allowed to contain qualifiers (.) to traverse sub-elements, as well as indexers ([x]) to traverse multiple-cardinality sub-elements (see the [Simple FHIRPath Profile](fhirpath.html#simple) for full details). Note that the index must be an integer constant. The path must resolve to an element of type date, dateTime, Period, Schedule, or Timing.
	//
	// Comment: The path attribute contains a [Simple FHIR Subset](fhirpath.html#simple) that allows path traversal, but not calculation.
	Path *string `bson:"path,omitempty" json:"path,omitempty" xml:"path,omitempty"`
	// A date parameter that refers to a search parameter defined on the specified type of the DataRequirement, and which searches on elements of type date, dateTime, Period, Schedule, or Timing.
	SearchParam *string `bson:"searchParam,omitempty" json:"searchParam,omitempty" xml:"searchParam,omitempty"`
	// The value of the filter.
	//
	// If period is specified, the filter will return only those data items that fall within the bounds determined by the Period, inclusive of the period boundaries. If dateTime is specified, the filter will return only those data items that are equal to the specified dateTime. If a Duration is specified, the filter will return only those data items that fall within Duration before now.
	ValueDateTime *string   `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty" xml:"valueDateTime,omitempty"`
	ValuePeriod   *Period   `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty" xml:"valuePeriod,omitempty"`
	ValueDuration *Duration `bson:"valueDuration,omitempty" json:"valueDuration,omitempty" xml:"valueDuration,omitempty"`
}

type DataRequirementValueFilter struct {
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
	// The attribute of the filter.
	//
	// The specified path SHALL be a FHIRPath resolvable on the specified type of the DataRequirement, and SHALL consist only of identifiers, constant indexers, and .resolve(). The path is allowed to contain qualifiers (.) to traverse sub-elements, as well as indexers ([x]) to traverse multiple-cardinality sub-elements (see the [Simple FHIRPath Profile](fhirpath.html#simple) for full details). Note that the index must be an integer constant. The path must resolve to an element of a type that is comparable to the valueFilter.value[x] element for the filter.
	//
	// Comment: The path attribute contains a [Simple FHIR Subset](fhirpath.html#simple) that allows path traversal, but not calculation.
	Path *string `bson:"path,omitempty" json:"path,omitempty" xml:"path,omitempty"`
	// A search parameter defined on the specified type of the DataRequirement, and which searches on elements of a type compatible with the type of the valueFilter.value[x] for the filter.
	SearchParam *string `bson:"searchParam,omitempty" json:"searchParam,omitempty" xml:"searchParam,omitempty"`
	// The comparator to be used to determine whether the value is matching.
	Comparator *ValueFilterComparator `bson:"comparator,omitempty" json:"comparator,omitempty" xml:"comparator,omitempty"`
	// The value of the filter.
	ValueDateTime *string   `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty" xml:"valueDateTime,omitempty"`
	ValuePeriod   *Period   `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty" xml:"valuePeriod,omitempty"`
	ValueDuration *Duration `bson:"valueDuration,omitempty" json:"valueDuration,omitempty" xml:"valueDuration,omitempty"`
}

type DataRequirementSort struct {
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
	// The attribute of the sort.
	//
	// The specified path must be resolvable from the type of the required data. The path is allowed to contain qualifiers (.) to traverse sub-elements, as well as indexers ([x]) to traverse multiple-cardinality sub-elements. Note that the index must be an integer constant.
	Path string `bson:"path" json:"path" xml:"path"`
	// The direction of the sort, ascending or descending.
	Direction SortDirection `bson:"direction" json:"direction" xml:"direction"`
}
