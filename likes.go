package activitypub

type (
	// LikesCollection is a list of all Like activities with this object as the object property,
	// added as a side effect. The likes collection MUST be either an OrderedCollection or a Collection
	// and MAY be filtered on privileges of an authenticated user or as appropriate when
	// no authentication is given.
	LikesCollection Likes

	// Likes is a type alias for an Ordered Collection
	Likes OrderedCollection
)

// LikesCollection initializes a new Outbox
func LikesNew() *Likes {
	id := ObjectID("likes")

	l := Likes{ID: id, Type: CollectionType}
	l.Name = NaturalLanguageValuesNew()
	l.Content = NaturalLanguageValuesNew()

	l.TotalItems = 0

	return &l
}

// Append adds an element to an LikesCollection
func (l *LikesCollection) Append(o Item) error {
	l.OrderedItems = append(l.OrderedItems, o)
	l.TotalItems++
	return nil
}

// Append adds an element to an Outbox
func (l *Likes) Append(ob Item) error {
	l.OrderedItems = append(l.OrderedItems, ob)
	l.TotalItems++
	return nil
}

// GetID returns the ObjectID corresponding to the LikesCollection
func (l LikesCollection) GetID() ObjectID {
	return l.Collection().GetID()
}

// GetLink returns the IRI corresponding to the current LikesCollection object
func (l LikesCollection) GetLink() IRI {
	return IRI(l.ID)
}

// GetType returns the LikesCollection's type
func (l LikesCollection) GetType() ActivityVocabularyType {
	return l.Type
}

// IsLink returns false for an LikesCollection object
func (l LikesCollection) IsLink() bool {
	return false
}

// IsObject returns true for a LikesCollection object
func (l LikesCollection) IsObject() bool {
	return true
}

// GetID returns the ObjectID corresponding to the Likes
func (l Likes) GetID() ObjectID {
	return l.Collection().GetID()
}

// GetLink returns the IRI corresponding to the current Likes object
func (l Likes) GetLink() IRI {
	return IRI(l.ID)
}

// GetType returns the Likes's type
func (l Likes) GetType() ActivityVocabularyType {
	return l.Type
}

// IsLink returns false for an Likes object
func (l Likes) IsLink() bool {
	return false
}

// IsObject returns true for a Likes object
func (l Likes) IsObject() bool {
	return true
}

// Collection returns the underlying Collection type
func (l Likes) Collection() CollectionInterface {
	c := OrderedCollection(l)
	return &c
}

// Collection returns the underlying Collection type
func (l LikesCollection) Collection() CollectionInterface {
	c := OrderedCollection(l)
	return &c
}
