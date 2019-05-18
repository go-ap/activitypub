package activitypub

import (
	as "github.com/go-ap/activitystreams"
)

type (
	// FollowingCollection is a list of everybody that the actor has followed, added as a side effect.
	// The following collection MUST be either an OrderedCollection or a Collection and MAY
	// be filtered on privileges of an authenticated user or as appropriate when no authentication is given.
	FollowingCollection = Following

	// Following is a type alias for a simple Collection
	Following as.Collection
)

// FollowingNew initializes a new Following
func FollowingNew() *Following {
	id := as.ObjectID("following")

	i := Following{Parent: as.Parent{ID: id, Type: as.CollectionType}}
	i.Name = as.NaturalLanguageValuesNew()
	i.Content = as.NaturalLanguageValuesNew()
	i.Summary = as.NaturalLanguageValuesNew()

	i.TotalItems = 0

	return &i
}

// Append adds an element to an Following
func (f *Following) Append(ob as.Item) error {
	f.Items = append(f.Items, ob)
	f.TotalItems++
	return nil
}

// GetID returns the ObjectID corresponding to Following
func (f Following) GetID() *as.ObjectID {
	return f.Collection().GetID()
}

// GetLink returns the IRI corresponding to the current Following object
func (f Following) GetLink() as.IRI {
	return as.IRI(f.ID)
}

// GetType returns the Following's type
func (f Following) GetType() as.ActivityVocabularyType {
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
	if as.ItemTyperFunc == nil {
		as.ItemTyperFunc = JSONGetItemByType
	}
	c := as.Collection(*f)
	err := c.UnmarshalJSON(data)

	*f = Following(c)

	return err
}

// Collection returns the underlying Collection type
func (f Following) Collection() as.CollectionInterface {
	c := as.Collection(f)
	return &c
}
