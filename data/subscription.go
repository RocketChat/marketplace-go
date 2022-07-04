package data

import (
	"time"

	"github.com/RocketChat/marketplace-go/events"
)

// SubscriptionStatus defines the current subscription status
type SubscriptionStatus string

var (
	// SubscriptionStatusTrialing is when the subscription is in the trial phase
	SubscriptionStatusTrialing SubscriptionStatus = "trialing"

	// SubscriptionStatusActive is when the subscription is active and being billed for
	SubscriptionStatusActive SubscriptionStatus = "active"

	// SubscriptionStatusCancelled is when the subscription is inactive due to being canceled
	SubscriptionStatusCancelled SubscriptionStatus = "cancelled" //nolint:misspell //cancelled is how we spell it

	// SubscriptionStatusCancelling is when the subscription is being canceled
	SubscriptionStatusCancelling SubscriptionStatus = "cancelling" //nolint:misspell //cancelled is how we spell it

	// SubscriptionStatusPastDue is when the subscription was active but is now past due as a result of incorrect billing information
	SubscriptionStatusPastDue SubscriptionStatus = "pastDue"
)

// Subscription happens when a user subscribed to an app
type Subscription struct {
	// Type defines a unique key.
	// This help allow us to create the create when we only have the json representation of this data
	Type events.EventObjectType `bson:"objectType" json:"objectType"`

	App events.App `bson:"app" json:"app"`

	Workspace events.Workspace `bson:"workspace" json:"workspace"`

	Pricing Pricing `bson:"pricing"  json:"pricing"`

	// Status defines the status of the subscription
	Status SubscriptionStatus `bson:"status" json:"status"`

	// PeriodStart defines when the the subscription started
	PeriodStart time.Time `bson:"periodStart" json:"periodStart"`

	// PeriodEnd defines when the subscription will end
	PeriodEnd *time.Time `bson:"periodEnd,omitempty" json:"periodEnd,omitempty"`

	TerminationDate *time.Time `bson:"terminationDate,omitempty" json:"terminationDate,omitempty"`

	// IsBundle tells us whether the subscription is a bundle
	IsBundle bool `bson:"isBundle" json:"isBundle"`

	// FromBundle defines if the app is inside a bundle and this event was triggered from a bundle
	FromBundle bool `bson:"fromBundle" json:"fromBundle"`

	// FromBundleID indicates which bundle id triggered this subscription event for the app
	FromBundleID string `bson:"fromBundleId,omitempty" json:"fromBundleId,omitempty"`

	// MigratedToBundle indicates if the subscription was migrated from a single app subscription to a bundle subscription
	MigratedToBundle bool `bson:"migratedToBundle,omitempty" json:"migratedToBundle,omitempty"`
}

//#region PricingRecurrence

// PricingRecurrence defines different prices for monthly and yearly billing options
type PricingRecurrence string

const (
	// MonthlyRecurrence const represents the string for monthly billing recurrence
	MonthlyRecurrence PricingRecurrence = "monthly"

	// YearlyRecurrence const represents the string for yearly billing recurrence
	YearlyRecurrence PricingRecurrence = "yearly"
)

//#endregion PricingRecurrence

// Pricing defines some of possible plans an subscription can have
type Pricing struct {
	// Name defines the plan string name
	Name string `bson:"name,omitempty" json:"name,omitempty"`

	// Price defines the current selected tier plan price
	Price float32 `bson:"price" json:"price"`

	// IsUsageBased indicates if the app is usage based or not,
	// subscription apps can take advantage of this field
	IsUsageBased bool `bson:"isUsageBased,omitempty" json:"isUsageBased,omitempty"`

	Recurrence PricingRecurrence `bson:"recurrence" json:"recurrence"`

	// TrialEnabled indicates if this pricing plan can trial
	TrialEnabled bool `bson:"trialEnabled" json:"trialEnabled"`

	// TrialDuration indicates how many days of trial a pricing plan can have
	TrialDuration int64 `bson:"trialDuration,omitempty" json:"trialDuration,omitempty"`
}
