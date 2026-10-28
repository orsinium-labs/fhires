package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// The presentation types of notes.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/note-type
type NoteType int

const (
	NoteTypeDisplay NoteType = iota
	// Print the note on the form.
	NoteTypePrint
	// Print the note for the operator.
	NoteTypePrintoper
)

func (code NoteType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *NoteType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "display":
		*code = NoteTypeDisplay
	case "print":
		*code = NoteTypePrint
	case "printoper":
		*code = NoteTypePrintoper
	default:
		return fmt.Errorf("unknown NoteType code `%s`", s)
	}
	return nil
}

func (code NoteType) String() string {
	return code.Code()
}

func (code NoteType) Code() string {
	switch code {
	case NoteTypeDisplay:
		return "display"
	case NoteTypePrint:
		return "print"
	case NoteTypePrintoper:
		return "printoper"
	}
	return "<unknown>"
}

func (code NoteType) Display() string {
	switch code {
	case NoteTypeDisplay:
		return "Display"
	case NoteTypePrint:
		return "Print (Form)"
	case NoteTypePrintoper:
		return "Print (Operator)"
	}
	return "<unknown>"
}

func (code NoteType) Definition() string {
	switch code {
	case NoteTypeDisplay:
		return "Display the note."
	case NoteTypePrint:
		return "Print the note on the form."
	case NoteTypePrintoper:
		return "Print the note for the operator."
	}
	return "<unknown>"
}
