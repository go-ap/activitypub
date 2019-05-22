package storage

import (
	as "github.com/go-ap/activitystreams"
)

// Loader
type Loader interface {
	Load(f Filterable) (as.ItemCollection, int, error)
}

// ActivityLoader
type ActivityLoader interface {
	LoadActivities(f Filterable) (as.ItemCollection, int, error)
}

// ActorLoader
type ActorLoader interface {
	LoadActors(f Filterable) (as.ItemCollection, int, error)
}

// ObjectLoader
type ObjectLoader interface {
	LoadObjects(f Filterable) (as.ItemCollection, int, error)
}

// CollectionLoader
type CollectionLoader interface {
	LoadCollection(f Filterable) (as.CollectionInterface, int, error)
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
