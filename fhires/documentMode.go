package fhires

// Whether the application produces or consumes documents.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/document-mode
type DocumentMode string

const (
	// The application produces documents of the specified type.
	DocumentModeProducer DocumentMode = "producer"
	// The application consumes documents of the specified type.
	DocumentModeConsumer DocumentMode = "consumer"
)

func (code DocumentMode) String() string {
	return code.Code()
}

func (code DocumentMode) Code() string {
	return string(code)
}

func (code DocumentMode) Display() string {
	switch code {
	case DocumentModeProducer:
		return "Producer"
	case DocumentModeConsumer:
		return "Consumer"
	}
	return "<unknown>"
}

func (code DocumentMode) Definition() string {
	switch code {
	case DocumentModeProducer:
		return "The application produces documents of the specified type."
	case DocumentModeConsumer:
		return "The application consumes documents of the specified type."
	}
	return "<unknown>"
}
