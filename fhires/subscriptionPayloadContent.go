package fhires

import (
	"encoding/json"
	"fmt"
	"strings"
)

// SubscriptionPayloadContent is documented here http://hl7.org/fhir/ValueSet/subscription-payload-content
type SubscriptionPayloadContent int

const (
	SubscriptionPayloadContentEmpty SubscriptionPayloadContent = iota
	SubscriptionPayloadContentIdOnly
	SubscriptionPayloadContentFullResource
)

func (code SubscriptionPayloadContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *SubscriptionPayloadContent) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "empty":
		*code = SubscriptionPayloadContentEmpty
	case "id-only":
		*code = SubscriptionPayloadContentIdOnly
	case "full-resource":
		*code = SubscriptionPayloadContentFullResource
	default:
		return fmt.Errorf("unknown SubscriptionPayloadContent code `%s`", s)
	}
	return nil
}

func (code SubscriptionPayloadContent) String() string {
	return code.Code()
}

func (code SubscriptionPayloadContent) Code() string {
	switch code {
	case SubscriptionPayloadContentEmpty:
		return "empty"
	case SubscriptionPayloadContentIdOnly:
		return "id-only"
	case SubscriptionPayloadContentFullResource:
		return "full-resource"
	}
	return "<unknown>"
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
