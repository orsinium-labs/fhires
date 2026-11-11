package fhires

// The status of the document reference.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/document-reference-status
type DocumentReferenceStatus string

const (
	// This is the current reference for this document.
	DocumentReferenceStatusCurrent DocumentReferenceStatus = "current"
	// This reference has been superseded by another reference.
	DocumentReferenceStatusSuperseded DocumentReferenceStatus = "superseded"
	// This reference was created in error.
	DocumentReferenceStatusEnteredInError DocumentReferenceStatus = "entered-in-error"
)

func (code DocumentReferenceStatus) String() string {
	return code.Code()
}

func (code DocumentReferenceStatus) Code() string {
	return string(code)
}

func (code DocumentReferenceStatus) Display() string {
	switch code {
	case DocumentReferenceStatusCurrent:
		return "Current"
	case DocumentReferenceStatusSuperseded:
		return "Superseded"
	case DocumentReferenceStatusEnteredInError:
		return "Entered in Error"
	}
	return "<unknown>"
}

func (code DocumentReferenceStatus) Definition() string {
	switch code {
	case DocumentReferenceStatusCurrent:
		return "This is the current reference for this document."
	case DocumentReferenceStatusSuperseded:
		return "This reference has been superseded by another reference."
	case DocumentReferenceStatusEnteredInError:
		return "This reference was created in error."
	}
	return "<unknown>"
}
