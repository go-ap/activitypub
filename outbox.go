package activitypub

type (
	// OutboxStream contains activities the user has published,
	// subject to the ability of the requestor to retrieve the activity (that is,
	// the contents of the outbox are filtered by the permissions of the person reading it).
	OutboxStream = Outbox

	// Outbox is a type alias for an Ordered Collection
	Outbox OrderedCollection
)

// OutboxNew initializes a new Outbox
func OutboxNew() *Outbox {
	id := ObjectID("outbox")

	i := Outbox{ID: id, Type: CollectionType}
	i.Name = NaturalLanguageValuesNew()
	i.Content = NaturalLanguageValuesNew()

	i.TotalItems = 0

	return &i
}

// Append adds an element to an Outbox
func (o *Outbox) Append(ob Item) error {
	o.OrderedItems = append(o.OrderedItems, ob)
	o.TotalItems++
	return nil
}

// GetID returns the ObjectID corresponding to Outbox
func (o Outbox) GetID() ObjectID {
	return o.Collection().GetID()
}

// GetLink returns the IRI corresponding to the current Outbox object
func (o Outbox) GetLink() IRI {
	return IRI(o.ID)
}

// GetType returns the Outbox's type
func (o Outbox) GetType() ActivityVocabularyType {
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
func (o *Outbox) UnmarshalJSON(data []byte) error {
	c := OrderedCollection(*o)
	err := c.UnmarshalJSON(data)

	*o = Outbox(c)

	return err
}

// Collection returns the underlying Collection type
func (o Outbox) Collection() CollectionInterface {
	c := OrderedCollection(o)
	return &c
}
