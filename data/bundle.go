package data

import "github.com/RocketChat/marketplace-go/events"

// BundleAction defines the data we will send when something happens to a bundle.
// Right now, we have: add/remove an app to bundle, add/remove restriction to an app in a bundle
type BundleAction struct {
	Type events.EventObjectType `bson:"objectType" json:"objectType"`

	// Bundle defines the bundle that is affected by the action
	Bundle events.App `bson:"bundle" json:"bundle"`

	// App defines the app that is affected by the action
	App events.App `bson:"app" json:"app"`
}
