package activitypub

import as "github.com/go-ap/activitystreams"

type (
	// LikedCollection is a list of every object from all of the actor's Like activities,
	// added as a side effect. The liked collection MUST be either an OrderedCollection or
	// a Collection and MAY be filtered on privileges of an authenticated user or as
	// appropriate when no authentication is given.
	LikedCollection Liked

	// Liked is a type alias for an Ordered Collection
	Liked as.OrderedCollection
)

// LikedCollection initializes a new Outbox
func LikedNew() *as.OrderedCollection {
	id := as.ObjectID("liked")

	l := as.OrderedCollection{Parent: as.Parent{ID: id, Type: as.CollectionType}}
	l.Name = as.NaturalLanguageValueNew()
	l.Content = as.NaturalLanguageValueNew()

	l.TotalItems = 0

	return &l
}

// Append adds an element to an LikedCollection
func (l *LikedCollection) Append(o as.Item) error {
	l.OrderedItems = append(l.OrderedItems, o)
	l.TotalItems++
	return nil
}

// Append adds an element to an Outbox
func (l *Liked) Append(ob as.Item) error {
	l.OrderedItems = append(l.OrderedItems, ob)
	l.TotalItems++
	return nil
}

// GetID returns the ObjectID corresponding to the LikedCollection
func (l LikedCollection) GetID() *as.ObjectID {
	return l.Collection().GetID()
}

// GetLink returns the IRI corresponding to the current LikedCollection object
func (l LikedCollection) GetLink() as.IRI {
	return as.IRI(l.ID)
}

// GetType returns the LikedCollection's type
func (l LikedCollection) GetType() as.ActivityVocabularyType {
	return l.Type
}

// IsLink returns false for an LikedCollection object
func (l LikedCollection) IsLink() bool {
	return false
}

// IsObject returns true for a LikedCollection object
func (l LikedCollection) IsObject() bool {
	return true
}

// GetID returns the ObjectID corresponding to the Liked
func (l Liked) GetID() *as.ObjectID {
	return l.Collection().GetID()
}

// GetLink returns the IRI corresponding to the current Liked object
func (l Liked) GetLink() as.IRI {
	return as.IRI(l.ID)
}

// GetType returns the Liked's type
func (l Liked) GetType() as.ActivityVocabularyType {
	return l.Type
}

// IsLink returns false for an Liked object
func (l Liked) IsLink() bool {
	return false
}

// IsObject returns true for a Liked object
func (l Liked) IsObject() bool {
	return true
}

// Collection returns the underlying Collection type
func (l Liked) Collection() as.CollectionInterface {
	c := as.OrderedCollection(l)
	return &c
}

// Collection returns the underlying Collection type
func (l LikedCollection) Collection() as.CollectionInterface {
	c := as.OrderedCollection(l)
	return &c
}
