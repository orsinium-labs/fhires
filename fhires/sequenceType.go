package fhires

// Type if a sequence -- DNA, RNA, or amino acid sequence.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/sequence-type
type SequenceType string

const (
	// Amino acid sequence.
	SequenceTypeAa SequenceType = "aa"
	// DNA Sequence.
	SequenceTypeDna SequenceType = "dna"
	// RNA Sequence.
	SequenceTypeRna SequenceType = "rna"
)

func (code SequenceType) String() string {
	return code.Code()
}

func (code SequenceType) Code() string {
	return string(code)
}

func (code SequenceType) Display() string {
	switch code {
	case SequenceTypeAa:
		return "AA Sequence"
	case SequenceTypeDna:
		return "DNA Sequence"
	case SequenceTypeRna:
		return "RNA Sequence"
	}
	return "<unknown>"
}

func (code SequenceType) Definition() string {
	switch code {
	case SequenceTypeAa:
		return "Amino acid sequence."
	case SequenceTypeDna:
		return "DNA Sequence."
	case SequenceTypeRna:
		return "RNA Sequence."
	}
	return "<unknown>"
}
