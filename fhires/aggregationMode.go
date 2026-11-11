package fhires

// How resource references can be aggregated.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/resource-aggregation-mode
type AggregationMode string

const (
	// The reference is a local reference to a contained resource.
	AggregationModeContained AggregationMode = "contained"
	// The reference to a resource that has to be resolved externally to the resource that includes the reference.
	AggregationModeReferenced AggregationMode = "referenced"
	// When the resource is in a Bundle, the resource the reference points to will be found in the same bundle as the resource that includes the reference.
	AggregationModeBundled AggregationMode = "bundled"
)

func (code AggregationMode) String() string {
	return code.Code()
}

func (code AggregationMode) Code() string {
	return string(code)
}

func (code AggregationMode) Display() string {
	switch code {
	case AggregationModeContained:
		return "Contained"
	case AggregationModeReferenced:
		return "Referenced"
	case AggregationModeBundled:
		return "Bundled"
	}
	return "<unknown>"
}

func (code AggregationMode) Definition() string {
	switch code {
	case AggregationModeContained:
		return "The reference is a local reference to a contained resource."
	case AggregationModeReferenced:
		return "The reference to a resource that has to be resolved externally to the resource that includes the reference."
	case AggregationModeBundled:
		return "When the resource is in a Bundle, the resource the reference points to will be found in the same bundle as the resource that includes the reference."
	}
	return "<unknown>"
}
