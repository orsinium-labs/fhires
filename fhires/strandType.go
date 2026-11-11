package fhires

// Type for strand.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/strand-type
type StrandType string

const (
	// Watson strand of starting sequence.
	StrandTypeWatson StrandType = "watson"
	// Crick strand of starting sequence.
	StrandTypeCrick StrandType = "crick"
)

func (code StrandType) String() string {
	return code.Code()
}

func (code StrandType) Code() string {
	return string(code)
}

func (code StrandType) Display() string {
	switch code {
	case StrandTypeWatson:
		return "Watson strand of starting sequence"
	case StrandTypeCrick:
		return "Crick strand of starting sequence"
	}
	return "<unknown>"
}

func (code StrandType) Definition() string {
	switch code {
	case StrandTypeWatson:
		return "Watson strand of starting sequence."
	case StrandTypeCrick:
		return "Crick strand of starting sequence."
	}
	return "<unknown>"
}
