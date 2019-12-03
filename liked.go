package activitypub

type (
	// LikedCollection is a list of every object from all of the actor's Like activities,
	// added as a side effect. The liked collection MUST be either an OrderedCollection or
	// a Collection and MAY be filtered on privileges of an authenticated user or as
	// appropriate when no authentication is given.
	LikedCollection = Liked

	// Liked is a type alias for an Ordered Collection
	Liked = OrderedCollection
)

// LikedCollection initializes a new Outbox
func LikedNew() *OrderedCollection {
	id := ObjectID("liked")

	l := OrderedCollection{ID: id, Type: CollectionType}
	l.Name = NaturalLanguageValuesNew()
	l.Content = NaturalLanguageValuesNew()

	l.TotalItems = 0

	return &l
}
