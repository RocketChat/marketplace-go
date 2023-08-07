package data

import (
	"time"

	"github.com/RocketChat/marketplace-go/events"
)

// Purchase defines a set of keys for a purchase event.
type Purchase struct {
	// objectType defines a unique key.
	// This helps allow us to create the create when we only have the JSON representation of this data.
	ObjectType   events.EventObjectType `bson:"objectType" json:"objectType"`
	App          events.App             `bson:"app" json:"app"`
	Workspace    events.Workspace       `bson:"workspace" json:"workspace"`
	Pricing      OneTimePricing         `bson:"pricing" json:"pricing"`
	BoughtAt     time.Time              `bson:"boughtAt" json:"boughtAt"`
	IsSelfManaged bool                   `bson:"isSelfManaged" json:"isSelfManaged"`
	IsBundle     bool                   `bson:"isBundle" json:"isBundle"`
	FromBundle   bool                   `bson:"fromBundle" json:"fromBundle"`
	FromBundleID string                 `bson:"fromBundleID,omitempty" json:"fromBundleID,omitempty"`
}

// OneTimePricing defines the pricing and fields for a purchase.
// There is no way to have different tiers for a one-time purchase.
type OneTimePricing struct {
	Price float32 `bson:"price" json:"price"`
}

// SetPurchaseTime sets the BoughtAt field with the specified time in the local timezone.
func (p *Purchase) SetPurchaseTime(t time.Time) {
	p.BoughtAt = t
}

// GetPurchaseTime returns the BoughtAt field as a time.Time value in the local timezone.
func (p *Purchase) GetPurchaseTime() time.Time {
	return p.BoughtAt
}
