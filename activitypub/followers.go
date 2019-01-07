package activitypub

import (
	"fmt"
	as "github.com/go-ap/activitypub.go/activitystreams"
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

	i := Followers{Parent: as.Parent{ID: id, Type: as.CollectionType}}
	i.Name = as.NaturalLanguageValueNew()
	i.Content = as.NaturalLanguageValueNew()
	i.Summary = as.NaturalLanguageValueNew()

	i.TotalItems = 0

	return &i
}

// Append adds an element to an FollowersCollection
func (f *FollowersCollection) Append(o as.Item) error {
	if f == nil {
		return fmt.Errorf("nil ")
	}
	f.Items = append(f.Items, o)
	f.TotalItems++
	return nil
}

// Append adds an element to an Followers
func (f *Followers) Append(ob as.Item) error {
	f.Items = append(f.Items, ob)
	f.TotalItems++
	return nil
}

// GetID returns the ObjectID corresponding to FollowersCollection
func (f FollowersCollection) GetID() *as.ObjectID {
	return f.Collection().GetID()
}

// GetLink returns the IRI corresponding to the current FollowersCollection object
func (f FollowersCollection) GetLink() as.IRI {
	return as.IRI(f.ID)
}

// GetType returns the FollowersCollection's type
func (f FollowersCollection) GetType() as.ActivityVocabularyType {
	return f.Type
}

// IsLink returns false for an FollowersCollection object
func (f FollowersCollection) IsLink() bool {
	return false
}

// IsObject returns true for a FollowersCollection object
func (f FollowersCollection) IsObject() bool {
	return true
}

// GetID returns the ObjectID corresponding to Followers
func (f Followers) GetID() *as.ObjectID {
	return f.Collection().GetID()
}

// GetLink returns the IRI corresponding to the current Followers object
func (f Followers) GetLink() as.IRI {
	return as.IRI(f.ID)
}

// GetType returns the Followers's type
func (f Followers) GetType() as.ActivityVocabularyType {
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
func (f *FollowersCollection) UnmarshalJSON(data []byte) error {
	c := as.Collection(*f)
	err := c.UnmarshalJSON(data)

	*f = FollowersCollection(c)

	return err
}

// UnmarshalJSON
func (f *Followers) UnmarshalJSON(data []byte) error {
	c := as.Collection(*f)
	err := c.UnmarshalJSON(data)

	*f = Followers(c)

	return err
}

// Collection returns the underlying Collection type
func (f Followers) Collection() as.CollectionInterface {
	c := as.Collection(f)
	return &c
}

// Collection returns the underlying Collection type
func (f FollowersCollection) Collection() as.CollectionInterface {
	c := as.Collection(f)
	return &c
}
