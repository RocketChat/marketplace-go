package data

import (
	"time"

	"github.com/RocketChat/marketplace-go/events"
)

// SubscriptionStatus defines the current subscription status.
type SubscriptionStatus string

const (
	SubscriptionStatusTrialing    SubscriptionStatus = "trialing"
	SubscriptionStatusActive     SubscriptionStatus = "active"
	SubscriptionStatusCancelled  SubscriptionStatus = "cancelled"
	SubscriptionStatusCancelling SubscriptionStatus = "cancelling"
	SubscriptionStatusPastDue    SubscriptionStatus = "pastDue"
)

// Subscription represents a user's subscription to an app.
type Subscription struct {
	Type             events.EventObjectType `bson:"objectType" json:"objectType"`
	App              events.App             `bson:"app" json:"app"`
	Workspace        events.Workspace       `bson:"workspace" json:"workspace"`
	Pricing          Pricing                `bson:"pricing" json:"pricing"`
	Status           SubscriptionStatus     `bson:"status" json:"status"`
	PeriodStart      time.Time              `bson:"periodStart" json:"periodStart"`
	PeriodEnd        *time.Time             `bson:"periodEnd,omitempty" json:"periodEnd,omitempty"`
	TerminationDate  *time.Time             `bson:"terminationDate,omitempty" json:"terminationDate,omitempty"`
	IsSelfManaged    bool                   `bson:"isSelfManaged" json:"isSelfManaged"`
	IsBundle         bool                   `bson:"isBundle" json:"isBundle"`
	FromBundle       bool                   `bson:"fromBundle" json:"fromBundle"`
	FromBundleID     string                 `bson:"fromBundleId,omitempty" json:"fromBundleId,omitempty"`
	MigratedToBundle bool                   `bson:"migratedToBundle,omitempty" json:"migratedToBundle,omitempty"`
}

// Pricing represents the pricing details for a subscription plan.
type Pricing struct {
	Name          string            `bson:"name,omitempty" json:"name,omitempty"`
	Price         float32           `bson:"price" json:"price"`
	IsUsageBased  bool              `bson:"isUsageBased,omitempty" json:"isUsageBased,omitempty"`
	Recurrence    PricingRecurrence `bson:"recurrence" json:"recurrence"`
	TrialEnabled  bool              `bson:"trialEnabled" json:"trialEnabled"`
	TrialDuration int64             `bson:"trialDuration,omitempty" json:"trialDuration,omitempty"`
}

// PricingRecurrence represents the recurrence options for pricing.
type PricingRecurrence string

const (
	PricingRecurrenceMonthly PricingRecurrence = "monthly"
	PricingRecurrenceYearly  PricingRecurrence = "yearly"
)
