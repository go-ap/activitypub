package storage

import (
	as "github.com/go-ap/activitystreams"
)

// Loader
type Loader interface {
	Load(f Filterable) (as.ItemCollection, error)
}
// ActivityLoader
type ActivityLoader interface {
	LoadActivities(f Filterable) (as.ItemCollection, error)
}
// ActorLoader
type ActorLoader interface {
	LoadActors(f Filterable) (as.ItemCollection, error)
}
// ObjectLoader
type ObjectLoader interface {
	LoadObjects(f Filterable) (as.ItemCollection, error)
}
// ActivitySaver
type ActivitySaver interface {
	SaveActivity(as.Item) (as.Item, error)
}
// ActorSaver
type ActorSaver interface {
	SaveActor(as.Item) (as.Item, error)
}
// ObjectSaver
type ObjectSaver interface {
	SaveObject(as.Item) (as.Item, error)
}

