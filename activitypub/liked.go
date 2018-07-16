package activitypub

type (
	// LikedCollection is a list of every object from all of the actor's Like activities,
	// added as a side effect. The liked collection MUST be either an OrderedCollection or
	// a Collection and MAY be filtered on privileges of an authenticated user or as
	// appropriate when no authentication is given.
	LikedCollection Liked

	// Liked is a type alias for an Ordered Collection
	Liked OrderedCollection
)

// LikedCollection initializes a new Outbox
func LikedNew() *Liked {
	id := ObjectID("liked")

	l := Liked{ID: id, Type: OrderedCollectionType}
	l.Name = make(NaturalLanguageValue)
	l.Content = make(NaturalLanguageValue)
	l.Summary = make(NaturalLanguageValue)

	l.TotalItems = 0

	return &l
}

// Append adds an element to an LikedCollection
func (l *LikedCollection) Append(o ObjectOrLink) error {
	l.OrderedItems = append(l.OrderedItems, o)
	l.TotalItems++
	return nil
}

// Append adds an element to an Outbox
func (l *Liked) Append(ob ObjectOrLink) error {
	l.OrderedItems = append(l.OrderedItems, ob)
	l.TotalItems++
	return nil
}

// GetID returns the LikedCollection's object ID
func (l LikedCollection) GetID() ObjectID {
	return l.ID
}

// GetType returns the LikedCollection's type
func (l LikedCollection) GetType() ActivityVocabularyType {
	return l.Type
}
