package activitypub

import (
	as "github.com/go-ap/activitystreams"
)

type (
	// InboxStream contains all activities received by the actor.
	// The server SHOULD filter content according to the requester's permission.
	// In general, the owner of an inbox is likely to be able to access all of their inbox contents.
	// Depending on access control, some other content may be public, whereas other content may
	// require authentication for non-owner users, if they can access the inbox at all.
	InboxStream = Inbox

	// Inbox is a type alias for an Ordered Collection
	Inbox as.OrderedCollection
)

// InboxNew initializes a new Inbox
func InboxNew() *as.OrderedCollection {
	id := as.ObjectID("inbox")

	i := as.OrderedCollection{Parent: as.Parent{ID: id, Type: as.CollectionType}}
	i.Name = as.NaturalLanguageValuesNew()
	i.Content = as.NaturalLanguageValuesNew()

	i.TotalItems = 0

	return &i
}

// Append adds an element to an Inbox
func (i *Inbox) Append(ob as.Item) error {
	i.OrderedItems = append(i.OrderedItems, ob)
	i.TotalItems++
	return nil
}

// GetID returns the ObjectID corresponding to Inbox
func (i Inbox) GetID() *as.ObjectID {
	return i.Collection().GetID()
}

// GetLink returns the IRI corresponding to the current Inbox object
func (i Inbox) GetLink() as.IRI {
	return as.IRI(i.ID)
}

// GetType returns the Inbox's type
func (i Inbox) GetType() as.ActivityVocabularyType {
	return i.Type
}

// IsLink returns false for an Inbox object
func (i Inbox) IsLink() bool {
	return false
}

// IsObject returns true for a Inbox object
func (i Inbox) IsObject() bool {
	return true
}

// UnmarshalJSON
func (i *Inbox) UnmarshalJSON(data []byte) error {
	if as.ItemTyperFunc == nil {
		as.ItemTyperFunc = JSONGetItemByType
	}
	c := as.OrderedCollection(*i)
	err := c.UnmarshalJSON(data)

	*i = Inbox(c)

	return err
}

// Collection returns the underlying Collection type
func (i Inbox) Collection() as.CollectionInterface {
	c := as.OrderedCollection(i)
	return &c
}
