package fhires

// Codes to represent how much resource content to send in the notification payload.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/subscription-payload-content
type SubscriptionPayloadContent string

const (
	// No resource content is transacted in the notification payload.
	SubscriptionPayloadContentEmpty SubscriptionPayloadContent = "empty"
	// Only the resource id is transacted in the notification payload.
	SubscriptionPayloadContentIdOnly SubscriptionPayloadContent = "id-only"
	// The entire resource is transacted in the notification payload.
	SubscriptionPayloadContentFullResource SubscriptionPayloadContent = "full-resource"
)

func (code SubscriptionPayloadContent) String() string {
	return code.Code()
}

func (code SubscriptionPayloadContent) Code() string {
	return string(code)
}

func (code SubscriptionPayloadContent) Display() string {
	switch code {
	case SubscriptionPayloadContentEmpty:
		return "Empty"
	case SubscriptionPayloadContentIdOnly:
		return "Id-only"
	case SubscriptionPayloadContentFullResource:
		return "Full-resource"
	}
	return "<unknown>"
}

func (code SubscriptionPayloadContent) Definition() string {
	switch code {
	case SubscriptionPayloadContentEmpty:
		return "No resource content is transacted in the notification payload."
	case SubscriptionPayloadContentIdOnly:
		return "Only the resource id is transacted in the notification payload."
	case SubscriptionPayloadContentFullResource:
		return "The entire resource is transacted in the notification payload."
	}
	return "<unknown>"
}
