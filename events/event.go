package events

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/rs/xid"
)

// App holds information about the application which the event was triggered
type App struct {
	// ID defines the application id that is provided by the publisher
	ID string `bson:"id" json:"id"`

	// Version defines the application version which the event was triggered
	Version string `bson:"version" json:"version"`

	// Name holds the name of the app. It helps us add more context to the payload
	Name string `bson:"name" json:"name"`

	// IsBundle defines if the app is a group of apps.
	IsBundle bool `bson:"isBundle" json:"isBundle"`

	BundledAppIDs []string     `bson:"bundledAppIds,omitempty" json:"bundledAppIds,omitempty"`
	BundledIn     []BundleInfo `bson:"bundledIn,omitempty" json:"bundledIn,omitempty"` // BundledIn only includes id and name
}

// BundleInfo defines information about a bundle when the app is included in a bundle.
// This way we can identify the bundle and the app is included in it
type BundleInfo struct {
	ID   string `bson:"id" json:"id,omitempty"`
	Name string `bson:"name" json:"name,omitempty"`
}

// Workspace defines the workspace where the event was triggered
type Workspace struct {
	ID        string `bson:"id" json:"id"`
	AccountID string `bson:"accountId" json:"accountId"`
	NickName  string `bson:"nickName" json:"nickName"`
	Address   string `bson:"address" json:"address"`
}

//#region APIVersion

// APIVersion defines the current API version
type APIVersion string

const (
	// CurrentAPIVersion specify the lastest API version
	CurrentAPIVersion APIVersion = "1.2.0"
)

//#endregion APIVersion

// Event represent something that happened
// An event is always related to an app. We will always have the application info at the root of the event
type Event struct {
	// unique generated id
	ID string `bson:"_id,omitempty" json:"id"`

	// Type defines an unique string key that identifies an event
	// E.g: application.purchased
	Type EventType `bson:"type" json:"type"`

	// Event definition version. We can use this field handle the event definition schema evolution
	// E.g: 1.0
	APIVersion `bson:"apiVersion" json:"apiVersion"`

	// Defines a UTC timestamp within the time when an event was created
	// E.g: 1651772844
	When time.Time `bson:"when" json:"when"`

	// Data has the event specific payload
	Data json.RawMessage `bson:"data" json:"data"`
}

// New creates a new event with the status of created
func New(name EventType, data interface{}) Event {
	marshalledJson, err := json.Marshal(data)
	if err != nil {
		panic(fmt.Sprintf("error creating a new event: %s", err.Error()))
	}

	return Event{
		ID:         fmt.Sprintf("evt_%s", xid.New().String()),
		Type:       name,
		When:       time.Now(),
		APIVersion: CurrentAPIVersion,
		Data:       marshalledJson,
	}
}

// Time handles the time.Time checking if the current date is a zero-value.
// If it is, it will return a pointer to the current time. Otherwise, we will return nil
// Using that, we can handle empty dates nil, instead of zero-value
func Time(date time.Time) *time.Time {
	if date.IsZero() || date.Year() <= 1970 {
		return nil
	}

	return &date
}
