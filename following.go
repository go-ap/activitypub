package activitypub

type (
	// FollowingCollection is a list of everybody that the actor has followed, added as a side effect.
	// The following collection MUST be either an OrderedCollection or a Collection and MAY
	// be filtered on privileges of an authenticated user or as appropriate when no authentication is given.
	FollowingCollection = Following

	// Following is a type alias for a simple Collection
	Following Collection
)

// FollowingNew initializes a new Following
func FollowingNew() *Following {
	id := ObjectID("following")

	i := Following{ID: id, Type: CollectionType}
	i.Name = NaturalLanguageValuesNew()
	i.Content = NaturalLanguageValuesNew()
	i.Summary = NaturalLanguageValuesNew()

	i.TotalItems = 0

	return &i
}

// Append adds an element to an Following
func (f *Following) Append(ob Item) error {
	f.Items = append(f.Items, ob)
	f.TotalItems++
	return nil
}

// GetID returns the ObjectID corresponding to Following
func (f Following) GetID() ObjectID {
	return f.Collection().GetID()
}

// GetLink returns the IRI corresponding to the current Following object
func (f Following) GetLink() IRI {
	return IRI(f.ID)
}

// GetType returns the Following's type
func (f Following) GetType() ActivityVocabularyType {
	return f.Type
}

// IsLink returns false for an Following object
func (f Following) IsLink() bool {
	return false
}

// IsObject returns true for a Following object
func (f Following) IsObject() bool {
	return true
}

// UnmarshalJSON
func (f *Following) UnmarshalJSON(data []byte) error {
	if ItemTyperFunc == nil {
		ItemTyperFunc = JSONGetItemByType
	}
	c := Collection(*f)
	err := c.UnmarshalJSON(data)

	*f = Following(c)

	return err
}

// Collection returns the underlying Collection type
func (f Following) Collection() CollectionInterface {
	c := Collection(f)
	return &c
}
