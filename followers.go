package activitypub

type (
	// FollowersCollection is a collection of followers
	FollowersCollection = Followers

	// Followers is a Collection type
	Followers Collection
)

// FollowersNew initializes a new Followers
func FollowersNew() *Followers {
	id := ObjectID("followers")

	i := Followers{ID: id, Type: CollectionType}
	i.Name = NaturalLanguageValuesNew()
	i.Content = NaturalLanguageValuesNew()
	i.Summary = NaturalLanguageValuesNew()

	i.TotalItems = 0

	return &i
}

// Append adds an element to an Followers
func (f *Followers) Append(ob Item) error {
	f.Items = append(f.Items, ob)
	f.TotalItems++
	return nil
}

// GetID returns the ObjectID corresponding to Followers
func (f Followers) GetID() ObjectID {
	return f.Collection().GetID()
}

// GetLink returns the IRI corresponding to the current Followers object
func (f Followers) GetLink() IRI {
	return IRI(f.ID)
}

// GetType returns the Followers's type
func (f Followers) GetType() ActivityVocabularyType {
	return f.Type
}

// IsLink returns false for an Followers object
func (f Followers) IsLink() bool {
	return false
}

// IsObject returns true for a Followers object
func (f Followers) IsObject() bool {
	return true
}

// UnmarshalJSON
func (f *Followers) UnmarshalJSON(data []byte) error {
	if ItemTyperFunc == nil {
		ItemTyperFunc = JSONGetItemByType
	}
	c := Collection(*f)
	err := c.UnmarshalJSON(data)

	*f = Followers(c)

	return err
}

// Collection returns the underlying Collection type
func (f Followers) Collection() CollectionInterface {
	c := Collection(f)
	return &c
}
