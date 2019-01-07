package activitypub

import (
	"fmt"
	as "github.com/go-ap/activitypub.go/activitystreams"
)

type (
	// FollowingCollection is a list of everybody that the actor has followed, added as a side effect.
	// The following collection MUST be either an OrderedCollection or a Collection and MAY
	// be filtered on privileges of an authenticated user or as appropriate when no authentication is given.
	FollowingCollection Following

	// Following is a type alias for a simple Collection
	Following as.Collection
)

// FollowingNew initializes a new Following
func FollowingNew() *Following {
	id := as.ObjectID("following")

	i := Following{Parent: as.Parent{ID: id, Type: as.CollectionType}}
	i.Name = as.NaturalLanguageValueNew()
	i.Content = as.NaturalLanguageValueNew()
	i.Summary = as.NaturalLanguageValueNew()

	i.TotalItems = 0

	return &i
}

// Append adds an element to an FollowingCollection
func (f *FollowingCollection) Append(o as.Item) error {
	if f == nil {
		return fmt.Errorf("nil ")
	}
	f.Items = append(f.Items, o)
	f.TotalItems++
	return nil
}

// Append adds an element to an Following
func (f *Following) Append(ob as.Item) error {
	f.Items = append(f.Items, ob)
	f.TotalItems++
	return nil
}

// GetID returns the ObjectID corresponding to FollowingCollection
func (f FollowingCollection) GetID() *as.ObjectID {
	return f.Collection().GetID()
}

// GetLink returns the IRI corresponding to the current FollowingCollection object
func (f FollowingCollection) GetLink() as.IRI {
	return as.IRI(f.ID)
}

// GetType returns the FollowingCollection's type
func (f FollowingCollection) GetType() as.ActivityVocabularyType {
	return f.Type
}

// IsLink returns false for an FollowingCollection object
func (f FollowingCollection) IsLink() bool {
	return false
}

// IsObject returns true for a FollowingCollection object
func (f FollowingCollection) IsObject() bool {
	return true
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
func (f *FollowingCollection) UnmarshalJSON(data []byte) error {
	c := as.Collection(*f)
	err := c.UnmarshalJSON(data)

	*f = FollowingCollection(c)

	return err
}

// UnmarshalJSON
func (f *Following) UnmarshalJSON(data []byte) error {
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

// Collection returns the underlying Collection type
func (f FollowingCollection) Collection() as.CollectionInterface {
	c := as.Collection(f)
	return &c
}
