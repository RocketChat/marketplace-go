package events

// EventType is a unique key defined by event
type EventType string

const (
	// AppSubscribedEventType defines the string constant name for the event that happens when
	// a user subscribes to a plan
	AppSubscribedEventType EventType = "app.subscription.created"

	// AppSubscriptionCancelledEventType defines the string constant name for the event that happens when
	// a user cancels a subscription
	AppSubscriptionCancelledEventType EventType = "app.subscription.cancelled"

	// AppSubscriptionModifiedEventType defines the string constant name for the event that happens when
	// a subscription is modified
	AppSubscriptionModifiedEventType EventType = "app.subscription.modified"

	// AppSubscriptionMigratedEventType defines the event related to when a subscription is migrate to use the bundle
	AppSubscriptionMigratedEventType EventType = "app.subscription.migrated"

	// AppPurchasedEventType indicates which type the struct is. Could make it easier for us data map this event from the database
	AppPurchasedEventType EventType = "app.purchase.created"
)

func (et EventType) Valid() bool {
	return et == AppSubscribedEventType ||
		et == AppPurchasedEventType ||
		et == AppSubscriptionCancelledEventType ||
		et == AppSubscriptionModifiedEventType ||
		et == AppSubscriptionMigratedEventType
}
