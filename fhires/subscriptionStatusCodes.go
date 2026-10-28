package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// State values for FHIR Subscriptions.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/subscription-status
type SubscriptionStatusCodes int

const (
	SubscriptionStatusCodesRequested SubscriptionStatusCodes = iota
	// The subscription is active.
	SubscriptionStatusCodesActive
	// The server has an error executing the notification.
	SubscriptionStatusCodesError
	// Too many errors have occurred or the subscription has expired.
	SubscriptionStatusCodesOff
	// This subscription has been flagged as incorrect.
	SubscriptionStatusCodesEnteredInError
)

func (code SubscriptionStatusCodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *SubscriptionStatusCodes) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "requested":
		*code = SubscriptionStatusCodesRequested
	case "active":
		*code = SubscriptionStatusCodesActive
	case "error":
		*code = SubscriptionStatusCodesError
	case "off":
		*code = SubscriptionStatusCodesOff
	case "entered-in-error":
		*code = SubscriptionStatusCodesEnteredInError
	default:
		return fmt.Errorf("unknown SubscriptionStatusCodes code `%s`", s)
	}
	return nil
}

func (code SubscriptionStatusCodes) String() string {
	return code.Code()
}

func (code SubscriptionStatusCodes) Code() string {
	switch code {
	case SubscriptionStatusCodesRequested:
		return "requested"
	case SubscriptionStatusCodesActive:
		return "active"
	case SubscriptionStatusCodesError:
		return "error"
	case SubscriptionStatusCodesOff:
		return "off"
	case SubscriptionStatusCodesEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}

func (code SubscriptionStatusCodes) Display() string {
	switch code {
	case SubscriptionStatusCodesRequested:
		return "Requested"
	case SubscriptionStatusCodesActive:
		return "Active"
	case SubscriptionStatusCodesError:
		return "Error"
	case SubscriptionStatusCodesOff:
		return "Off"
	case SubscriptionStatusCodesEnteredInError:
		return "Entered in Error"
	}
	return "<unknown>"
}

func (code SubscriptionStatusCodes) Definition() string {
	switch code {
	case SubscriptionStatusCodesRequested:
		return "The client has requested the subscription, and the server has not yet set it up."
	case SubscriptionStatusCodesActive:
		return "The subscription is active."
	case SubscriptionStatusCodesError:
		return "The server has an error executing the notification."
	case SubscriptionStatusCodesOff:
		return "Too many errors have occurred or the subscription has expired."
	case SubscriptionStatusCodesEnteredInError:
		return "This subscription has been flagged as incorrect."
	}
	return "<unknown>"
}
