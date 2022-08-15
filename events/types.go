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

	// AppBundleRestrictedEventType happens when an app is restricted to a bundle
	// When an app is restricted to a bundle, it means the app can only be used if you subscribe the bundle
	AppBundleRestrictedEventType EventType = "app.bundle.restricted"

	// AppBundleUnrestrictedEventType happens when an app is unrestricted to a bundle
	AppBundleUnrestrictedEventType EventType = "app.bundle.unrestricted"

	// BundleAppAddedEventType happens when a bundle is restricted to an app
	BundleAppAddedEventType EventType = "bundle.app.added"

	// BundleAppRemovedEventType happens when a bundle is unrestricted to an app
	BundleAppRemovedEventType EventType = "bundle.app.removed"
)

// Valid events will look for all available events type we have. If the event is one of them, than it's valid
func (et EventType) Valid() bool {
	return et.ValidAppEvent() || et.ValidBundleEvent()
}

// ValidAppEvent will only look for app related events, so if you want to check if it's an app related event
// you can use this method
func (et EventType) ValidAppEvent() bool {
	return et == AppSubscribedEventType ||
		et == AppPurchasedEventType ||
		et == AppSubscriptionCancelledEventType ||
		et == AppSubscriptionModifiedEventType ||
		et == AppSubscriptionMigratedEventType ||
		et == AppBundleRestrictedEventType ||
		et == AppBundleUnrestrictedEventType
}

// ValidBundleEvent will only look for bundle related events, so if you want to check if it's a bundle related event
func (et EventType) ValidBundleEvent() bool {
	return et == BundleAppAddedEventType || et == BundleAppRemovedEventType
}
