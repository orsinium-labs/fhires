package fhires

// The type of link between this Patient resource and another Patient/RelatedPerson resource.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/link-type
type LinkType string

const (
	// The patient resource containing this link must no longer be used. The link points forward to another patient resource that must be used in lieu of the patient resource that contains this link.
	LinkTypeReplacedBy LinkType = "replaced-by"
	// The patient resource containing this link is the current active patient record. The link points back to an inactive patient resource that has been merged into this resource, and should be consulted to retrieve additional referenced information.
	LinkTypeReplaces LinkType = "replaces"
	// The patient resource containing this link is in use and valid but not considered the main source of information about a patient. The link points forward to another patient resource that should be consulted to retrieve additional patient information.
	LinkTypeRefer LinkType = "refer"
	// The patient resource containing this link is in use and valid, but points to another Patient or RelatedPerson resource that is known to contain data about the same person. Data in this resource might overlap or contradict information found in the other Patient/RelatedPerson resource. This link does not indicate any relative importance of the resources concerned, and both should be regarded as equally valid.
	LinkTypeSeealso LinkType = "seealso"
)

func (code LinkType) String() string {
	return code.Code()
}

func (code LinkType) Code() string {
	return string(code)
}

func (code LinkType) Display() string {
	switch code {
	case LinkTypeReplacedBy:
		return "Replaced-by"
	case LinkTypeReplaces:
		return "Replaces"
	case LinkTypeRefer:
		return "Refer"
	case LinkTypeSeealso:
		return "See also"
	}
	return "<unknown>"
}

func (code LinkType) Definition() string {
	switch code {
	case LinkTypeReplacedBy:
		return "The patient resource containing this link must no longer be used. The link points forward to another patient resource that must be used in lieu of the patient resource that contains this link."
	case LinkTypeReplaces:
		return "The patient resource containing this link is the current active patient record. The link points back to an inactive patient resource that has been merged into this resource, and should be consulted to retrieve additional referenced information."
	case LinkTypeRefer:
		return "The patient resource containing this link is in use and valid but not considered the main source of information about a patient. The link points forward to another patient resource that should be consulted to retrieve additional patient information."
	case LinkTypeSeealso:
		return "The patient resource containing this link is in use and valid, but points to another Patient or RelatedPerson resource that is known to contain data about the same person. Data in this resource might overlap or contradict information found in the other Patient/RelatedPerson resource. This link does not indicate any relative importance of the resources concerned, and both should be regarded as equally valid."
	}
	return "<unknown>"
}
