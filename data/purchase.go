package data

import (
	"time"

	"github.com/RocketChat/marketplace-go/events"
)

// Purchase defines a set of keys for a purchase event.
type Purchase struct {
	Type          events.EventObjectType `bson:"objectType" json:"objectType"`
	App           events.App             `bson:"app" json:"app"`
	Workspace     events.Workspace       `bson:"workspace" json:"workspace"`
	Pricing       OneTimePricing         `bson:"pricing" json:"pricing"`
	BoughtAt      time.Time              `bson:"boughtAt" json:"boughtAt"`
	IsSelfManaged bool                   `bson:"isSelfManaged" json:"isSelfManaged"`
	IsBundle      bool                   `bson:"isBundle" json:"isBundle"`
	FromBundle    bool                   `bson:"fromBundle" json:"fromBundle"`
	FromBundleID  string                 `bson:"fromBundleId,omitempty" json:"fromBundleId,omitempty"`
}

// OneTimePricing defines the pricing and fields for a one-time purchase.
type OneTimePricing struct {
	Price float32 `bson:"price" json:"price"`
}

// NewPurchase creates a new Purchase event with the specified values.
func NewPurchase(purchaseType events.EventObjectType, app events.App, workspace events.Workspace, pricing OneTimePricing, boughtAt time.Time, isSelfManaged, isBundle, fromBundle bool, fromBundleID string, timezone string) (Purchase, error) {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return Purchase{}, err
	}

	return Purchase{
		Type:          purchaseType,
		App:           app,
		Workspace:     workspace,
		Pricing:       pricing,
		BoughtAt:      boughtAt.In(location),
		IsSelfManaged: isSelfManaged,
		IsBundle:      isBundle,
		FromBundle:    fromBundle,
		FromBundleID:  fromBundleID,
	}, nil
}

// GetBoughtAtInTimezone returns the BoughtAt time in the specified timezone.
func (p *Purchase) GetBoughtAtInTimezone(timezone string) (time.Time, error) {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{}, err
	}
	return p.BoughtAt.In(location), nil
}
