package events_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/RocketChat/marketplace-go/data"
	"github.com/RocketChat/marketplace-go/events"
)

func TestNewPurchase(t *testing.T) {
	app := events.App{
		ID:       "app-id",
		Version:  "1.0.0",
		Name:     "my test app",
		IsBundle: false,
	}

	w := events.Workspace{
		ID:       "workspace-id",
		NickName: "my test workspace",
		Address:  "https://mytestworkspace.example.com",
	}

	d := data.Purchase{
		Type:      events.PurchaseDataType,
		App:       app,
		Workspace: w,
		Pricing:   data.OneTimePricing{Price: 1.0},
		BoughtAt:  time.Now(),
		IsBundle:  false,
	}

	ev := events.New(events.AppSubscribedEventType, d)

	// Checking the event type
	if ev.Type != events.AppSubscribedEventType {
		t.Errorf("Expecting event type %s, got %s", events.AppSubscribedEventType, ev.Type)
	}

	// checking the prefix
	if strings.HasPrefix(ev.ID, "evt_") == false {
		t.Errorf("Expecting event id to start with evt_, got %s", ev.ID)
	}

	var purchaseData data.Purchase
	if err := json.Unmarshal(ev.Data, &purchaseData); err != nil {
		t.Errorf("Unable to unmarshall the purchase data: %s", err.Error())
	}

	// Checking if we can unmarshal the event back to the original data
	if purchaseData.Type != events.PurchaseDataType {
		t.Errorf("Expecting event data type %s, got %s", events.PurchaseDataType, purchaseData.Type)
	}
}

func TestNewSubscription(t *testing.T) {
	app := events.App{
		ID:       "app-id",
		Version:  "1.0.0",
		Name:     "my test app",
		IsBundle: false,
	}

	w := events.Workspace{
		ID:       "workspace-id",
		NickName: "my test workspace",
		Address:  "https://mytestworkspace.example.com",
	}

	d := data.Subscription{
		Type:      events.SubscriptionDataType,
		App:       app,
		Workspace: w,
		Pricing: data.Pricing{
			Name:          "my example pricing",
			Price:         1.0,
			IsUsageBased:  false,
			Recurrence:    data.MonthlyRecurrence,
			TrialEnabled:  true,
			TrialDuration: 15,
		},
		Status:      data.SubscriptionStatusActive,
		PeriodStart: time.Now(),
		IsBundle:    false,
	}

	ev := events.New(events.AppSubscribedEventType, d)

	// Checking the event type
	if ev.Type != events.AppSubscribedEventType {
		t.Errorf("Expecting event type %s, got %s", events.AppSubscribedEventType, ev.Type)
	}

	// checking the prefix
	if strings.HasPrefix(ev.ID, "evt_") == false {
		t.Errorf("Expecting event id to start with evt_, got %s", ev.ID)
	}

	var subscriptionData data.Subscription
	if err := json.Unmarshal(ev.Data, &subscriptionData); err != nil {
		t.Errorf("Unable to unmarshall the subscription data: %s", err.Error())
	}

	// Checking if we can unmarshal the event back to the original data
	if subscriptionData.Type != events.SubscriptionDataType {
		t.Errorf("Expecting event data type %s, got %s", events.SubscriptionDataType, subscriptionData.Type)
	}
}

func TestTime(t *testing.T) {
	zeroValue := time.Time{}
	unixZeroValue := time.Unix(0, 0).UTC()

	timeAfter := events.Time(zeroValue)
	if timeAfter != nil {
		t.Errorf("Expecting nil date as we passed a zero-value date, got %v", timeAfter)
	}

	unixTimeAfter := events.Time(unixZeroValue)
	if unixTimeAfter != nil {
		t.Errorf("Expecting nil date as we passed a zero-value date, got %v", unixTimeAfter)
	}

	validDate := time.Now()
	validDateAfter := events.Time(validDate)
	if validDateAfter == nil {
		t.Errorf("Expecting date as we passed a valid date, got %v", validDateAfter)
	}
}
