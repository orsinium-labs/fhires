package fhires

// The presentation types of notes.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/note-type
type NoteType string

const (
	// Display the note.
	NoteTypeDisplay NoteType = "display"
	// Print the note on the form.
	NoteTypePrint NoteType = "print"
	// Print the note for the operator.
	NoteTypePrintoper NoteType = "printoper"
)

func (code NoteType) String() string {
	return code.Code()
}

func (code NoteType) Code() string {
	return string(code)
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
