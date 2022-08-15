package events

// EventObjectType holds the event type we are working.
type EventObjectType string

const (
	// SubscriptionDataType describes the type of the data we are holding.
	// It allow us easily map data from the database to a struct
	SubscriptionDataType EventObjectType = "subscription"

	// PurchaseDataType describes the type of the data we are holding.
	// It allow us easily map data from the database to a struct
	PurchaseDataType EventObjectType = "purchase"

	// BundleActionDataType describes the type of the data we are holding.
	// It allow us easily map data from the database to a struct
	BundleActionDataType EventObjectType = "bundleAction"
)
