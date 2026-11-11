package fhires

// The type of notification represented by the status message.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/subscription-notification-type
type SubscriptionNotificationType string

const (
	// The status was generated as part of the setup or verification of a communications channel.
	SubscriptionNotificationTypeHandshake SubscriptionNotificationType = "handshake"
	// The status was generated to perform a heartbeat notification to the subscriber.
	SubscriptionNotificationTypeHeartbeat SubscriptionNotificationType = "heartbeat"
	// The status was generated for an event to the subscriber.
	SubscriptionNotificationTypeEventNotification SubscriptionNotificationType = "event-notification"
	// The status was generated in response to a status query/request.
	SubscriptionNotificationTypeQueryStatus SubscriptionNotificationType = "query-status"
	// The status was generated in response to an event query/request.
	SubscriptionNotificationTypeQueryEvent SubscriptionNotificationType = "query-event"
)

func (code SubscriptionNotificationType) String() string {
	return code.Code()
}

func (code SubscriptionNotificationType) Code() string {
	return string(code)
}

func (code SubscriptionNotificationType) Display() string {
	switch code {
	case SubscriptionNotificationTypeHandshake:
		return "Handshake"
	case SubscriptionNotificationTypeHeartbeat:
		return "Heartbeat"
	case SubscriptionNotificationTypeEventNotification:
		return "Event Notification"
	case SubscriptionNotificationTypeQueryStatus:
		return "Query Status"
	case SubscriptionNotificationTypeQueryEvent:
		return "Query Event"
	}
	return "<unknown>"
}

func (code SubscriptionNotificationType) Definition() string {
	switch code {
	case SubscriptionNotificationTypeHandshake:
		return "The status was generated as part of the setup or verification of a communications channel."
	case SubscriptionNotificationTypeHeartbeat:
		return "The status was generated to perform a heartbeat notification to the subscriber."
	case SubscriptionNotificationTypeEventNotification:
		return "The status was generated for an event to the subscriber."
	case SubscriptionNotificationTypeQueryStatus:
		return "The status was generated in response to a status query/request."
	case SubscriptionNotificationTypeQueryEvent:
		return "The status was generated in response to an event query/request."
	}
	return "<unknown>"
}
