package events

import (
	"time"

	"github.com/RocketChat/marketplace-go/events"
)

// Purchase defines a set of keys for a purchase event.
type Purchase struct {
	// Type defines a unique key.
	// This help allow us to create the create when we only have the json representation of this data
	Type events.EventObjectType `bson:"objectType" json:"objectType"`

	App       events.App       `bson:"app" json:"app"`
	Workspace events.Workspace `bson:"workspace" json:"workspace"`

	Pricing OneTimePricing `bson:"pricing" json:"pricing"`

	// BoughtAt defines when the app was purchased
	BoughtAt time.Time `bson:"boughtAt" json:"boughtAt"`

	// IsBundle tells us whether the purchase is a bundle
	IsBundle bool `bson:"isBundle" json:"isBundle"`

	// FromBundle defines if the app is inside a bundle and this event was triggered from a bundle
	FromBundle bool `bson:"fromBundle" json:"fromBundle"`

	// FromBundleID indicates which bundle id triggered this purchase event for the app
	FromBundleID string `bson:"fromBundleId,omitempty" json:"fromBundleId,omitempty"`
}

// OneTimePricing defines the pricing and fields for a purchase
// There is no way to have different tiers for an one time purchase
type OneTimePricing struct {
	Price float32 `bson:"price" json:"price"`
}
