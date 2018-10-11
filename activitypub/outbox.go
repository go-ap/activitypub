package activitypub

import as "github.com/mariusor/activitypub.go/activitystreams"

type (
	// OutboxStream contains activities the user has published,
	// subject to the ability of the requestor to retrieve the activity (that is,
	// the contents of the outbox are filtered by the permissions of the person reading it).
	OutboxStream Outbox

	// Outbox is a type alias for an Ordered Collection
	Outbox as.OrderedCollection
)

// OutboxNew initializes a new Outbox
func OutboxNew() *Outbox {
	id := as.ObjectID("outbox")

	i := Outbox{ID: id, Type: as.OrderedCollectionType}
	i.Name = as.NaturalLanguageValueNew()
	i.Content = as.NaturalLanguageValueNew()

	i.TotalItems = 0

	return &i
}

// Append adds an element to an OutboxStream
func (o *OutboxStream) Append(ob as.Item) error {
	o.OrderedItems = append(o.OrderedItems, ob)
	o.TotalItems++
	return nil
}

// Append adds an element to an Outbox
func (o *Outbox) Append(ob as.Item) error {
	o.OrderedItems = append(o.OrderedItems, ob)
	o.TotalItems++
	return nil
}

// GetID returns the ObjectID corresponding to the OutboxStream
func (o OutboxStream) GetID() *as.ObjectID {
	return o.Collection().GetID()
}

// GetType returns the OutboxStream's type
func (o OutboxStream) GetType() as.ActivityVocabularyType {
	return o.Type
}

// IsLink returns false for an OutboxStream object
func (o OutboxStream) IsLink() bool {
	return false
}

// IsObject returns true for a OutboxStream object
func (o OutboxStream) IsObject() bool {
	return true
}

// GetID returns the ObjectID corresponding to Outbox
func (o Outbox) GetID() *as.ObjectID {
	return o.Collection().GetID()
}

// GetType returns the Outbox's type
func (o Outbox) GetType() as.ActivityVocabularyType {
	return o.Type
}

// IsLink returns false for an Outbox object
func (o Outbox) IsLink() bool {
	return false
}

// IsObject returns true for a Outbox object
func (o Outbox) IsObject() bool {
	return true
}

// UnmarshalJSON
func (o *OutboxStream) UnmarshalJSON(data []byte) error {
	c := as.OrderedCollection(*o)
	err := c.UnmarshalJSON(data)

	*o = OutboxStream(c)

	return err
}

// UnmarshalJSON
func (o *Outbox) UnmarshalJSON(data []byte) error {
	c := as.OrderedCollection(*o)
	err := c.UnmarshalJSON(data)

	*o = Outbox(c)

	return err
}

// Collection returns the underlying Collection type
func (o Outbox) Collection() as.CollectionInterface {
	c := as.OrderedCollection(o)
	return &c
}

// Collection returns the underlying Collection type
func (o OutboxStream) Collection() as.CollectionInterface {
	c := as.OrderedCollection(o)
	return &c
}
