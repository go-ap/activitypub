package activitypub

import (
	"fmt"
	as "github.com/mariusor/activitypub.go/activitystreams"
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

	i := Following{ID: id, Type: as.OrderedCollectionType}
	i.Name = as.NaturalLanguageValueNew()
	i.Content = as.NaturalLanguageValueNew()
	i.Summary = as.NaturalLanguageValueNew()

	i.TotalItems = 0

	return &i
}

// Append adds an element to an FollowingCollection
func (i *FollowingCollection) Append(o as.Item) error {
	if i == nil {
		return fmt.Errorf("nil ")
	}
	i.Items = append(i.Items, o)
	i.TotalItems++
	return nil
}

// Append adds an element to an Following
func (i *Following) Append(ob as.Item) error {
	i.Items = append(i.Items, ob)
	i.TotalItems++
	return nil
}

// GetID returns the ObjectID corresponding to FollowingCollection
func (i FollowingCollection) GetID() *as.ObjectID {
	return i.Collection().GetID()
}

// GetType returns the FollowingCollection's type
func (i FollowingCollection) GetType() as.ActivityVocabularyType {
	return i.Type
}

// IsLink returns false for an FollowingCollection object
func (i FollowingCollection) IsLink() bool {
	return false
}

// IsObject returns true for a FollowingCollection object
func (i FollowingCollection) IsObject() bool {
	return true
}

// GetID returns the ObjectID corresponding to Following
func (i Following) GetID() *as.ObjectID {
	return i.Collection().GetID()
}

// GetType returns the Following's type
func (i Following) GetType() as.ActivityVocabularyType {
	return i.Type
}

// IsLink returns false for an Following object
func (i Following) IsLink() bool {
	return false
}

// IsObject returns true for a Following object
func (i Following) IsObject() bool {
	return true
}

// UnmarshalJSON
func (i *FollowingCollection) UnmarshalJSON(data []byte) error {
	c := as.Collection(*i)
	err := c.UnmarshalJSON(data)

	*i = FollowingCollection(c)

	return err
}

// UnmarshalJSON
func (i *Following) UnmarshalJSON(data []byte) error {
	c := as.Collection(*i)
	err := c.UnmarshalJSON(data)

	*i = Following(c)

	return err
}

// Collection returns the underlying Collection type
func (i Following) Collection() as.CollectionInterface {
	c := as.Collection(i)
	return &c
}

// Collection returns the underlying Collection type
func (i FollowingCollection) Collection() as.CollectionInterface {
	c := as.Collection(i)
	return &c
}
