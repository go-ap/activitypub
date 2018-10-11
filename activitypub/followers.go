package activitypub

import (
	"fmt"
	as "github.com/mariusor/activitypub.go/activitystreams"
)

type (
	// FollowersCollection is a collection of followers
	FollowersCollection Followers

	// Followers is a Collection type
	Followers as.Collection
)

// FollowersNew initializes a new Followers
func FollowersNew() *Followers {
	id := as.ObjectID("followers")

	i := Followers{ID: id, Type: as.OrderedCollectionType}
	i.Name = as.NaturalLanguageValueNew()
	i.Content = as.NaturalLanguageValueNew()
	i.Summary = as.NaturalLanguageValueNew()

	i.TotalItems = 0

	return &i
}

// Append adds an element to an FollowersCollection
func (i *FollowersCollection) Append(o as.Item) error {
	if i == nil {
		return fmt.Errorf("nil ")
	}
	i.Items = append(i.Items, o)
	i.TotalItems++
	return nil
}

// Append adds an element to an Followers
func (i *Followers) Append(ob as.Item) error {
	i.Items = append(i.Items, ob)
	i.TotalItems++
	return nil
}

// GetID returns the ObjectID corresponding to FollowersCollection
func (i FollowersCollection) GetID() *as.ObjectID {
	return i.Collection().GetID()
}

// GetType returns the FollowersCollection's type
func (i FollowersCollection) GetType() as.ActivityVocabularyType {
	return i.Type
}

// IsLink returns false for an FollowersCollection object
func (i FollowersCollection) IsLink() bool {
	return false
}

// IsObject returns true for a FollowersCollection object
func (i FollowersCollection) IsObject() bool {
	return true
}

// GetID returns the ObjectID corresponding to Followers
func (i Followers) GetID() *as.ObjectID {
	return i.Collection().GetID()
}

// GetType returns the Followers's type
func (i Followers) GetType() as.ActivityVocabularyType {
	return i.Type
}

// IsLink returns false for an Followers object
func (i Followers) IsLink() bool {
	return false
}

// IsObject returns true for a Followers object
func (i Followers) IsObject() bool {
	return true
}

// UnmarshalJSON
func (i *FollowersCollection) UnmarshalJSON(data []byte) error {
	c := as.Collection(*i)
	err := c.UnmarshalJSON(data)

	*i = FollowersCollection(c)

	return err
}

// UnmarshalJSON
func (i *Followers) UnmarshalJSON(data []byte) error {
	c := as.Collection(*i)
	err := c.UnmarshalJSON(data)

	*i = Followers(c)

	return err
}

// Collection returns the underlying Collection type
func (i Followers) Collection() as.CollectionInterface {
	c := as.Collection(i)
	return &c
}

// Collection returns the underlying Collection type
func (i FollowersCollection) Collection() as.CollectionInterface {
	c := as.Collection(i)
	return &c
}
