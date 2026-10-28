package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// Type if a sequence -- DNA, RNA, or amino acid sequence.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/sequence-type
type SequenceType int

const (
	SequenceTypeAa SequenceType = iota
	// DNA Sequence.
	SequenceTypeDna
	// RNA Sequence.
	SequenceTypeRna
)

func (code SequenceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *SequenceType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "aa":
		*code = SequenceTypeAa
	case "dna":
		*code = SequenceTypeDna
	case "rna":
		*code = SequenceTypeRna
	default:
		return fmt.Errorf("unknown SequenceType code `%s`", s)
	}
	return nil
}

func (code SequenceType) String() string {
	return code.Code()
}

func (code SequenceType) Code() string {
	switch code {
	case SequenceTypeAa:
		return "aa"
	case SequenceTypeDna:
		return "dna"
	case SequenceTypeRna:
		return "rna"
	}
	return "<unknown>"
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
