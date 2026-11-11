package fhires

// State values for FHIR Subscriptions.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/subscription-status
type SubscriptionStatusCodes string

const (
	// The client has requested the subscription, and the server has not yet set it up.
	SubscriptionStatusCodesRequested SubscriptionStatusCodes = "requested"
	// The subscription is active.
	SubscriptionStatusCodesActive SubscriptionStatusCodes = "active"
	// The server has an error executing the notification.
	SubscriptionStatusCodesError SubscriptionStatusCodes = "error"
	// Too many errors have occurred or the subscription has expired.
	SubscriptionStatusCodesOff SubscriptionStatusCodes = "off"
	// This subscription has been flagged as incorrect.
	SubscriptionStatusCodesEnteredInError SubscriptionStatusCodes = "entered-in-error"
)

func (code SubscriptionStatusCodes) String() string {
	return code.Code()
}

func (code SubscriptionStatusCodes) Code() string {
	return string(code)
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
