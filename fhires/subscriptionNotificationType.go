package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// SubscriptionNotificationType is documented here http://hl7.org/fhir/ValueSet/subscription-notification-type
type SubscriptionNotificationType int

const (
	SubscriptionNotificationTypeHandshake SubscriptionNotificationType = iota
	SubscriptionNotificationTypeHeartbeat
	SubscriptionNotificationTypeEventNotification
	SubscriptionNotificationTypeQueryStatus
	SubscriptionNotificationTypeQueryEvent
)

func (code SubscriptionNotificationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *SubscriptionNotificationType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "handshake":
		*code = SubscriptionNotificationTypeHandshake
	case "heartbeat":
		*code = SubscriptionNotificationTypeHeartbeat
	case "event-notification":
		*code = SubscriptionNotificationTypeEventNotification
	case "query-status":
		*code = SubscriptionNotificationTypeQueryStatus
	case "query-event":
		*code = SubscriptionNotificationTypeQueryEvent
	default:
		return fmt.Errorf("unknown SubscriptionNotificationType code `%s`", s)
	}
	return nil
}

func (code SubscriptionNotificationType) String() string {
	return code.Code()
}

func (code SubscriptionNotificationType) Code() string {
	switch code {
	case SubscriptionNotificationTypeHandshake:
		return "handshake"
	case SubscriptionNotificationTypeHeartbeat:
		return "heartbeat"
	case SubscriptionNotificationTypeEventNotification:
		return "event-notification"
	case SubscriptionNotificationTypeQueryStatus:
		return "query-status"
	case SubscriptionNotificationTypeQueryEvent:
		return "query-event"
	}
	return "<unknown>"
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
