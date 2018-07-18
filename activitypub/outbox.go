package activitypub

type (
	// OutboxStream contains activities the user has published,
	// subject to the ability of the requestor to retrieve the activity (that is,
	// the contents of the outbox are filtered by the permissions of the person reading it).
	OutboxStream Outbox

	// Outbox is a type alias for an Ordered Collection
	Outbox OrderedCollection
)

// OutboxNew initializes a new Outbox
func OutboxNew() *Outbox {
	id := ObjectID("outbox")

	i := Outbox{ID: id, Type: OrderedCollectionType}
	i.Name = make(NaturalLanguageValue)
	i.Content = make(NaturalLanguageValue)
	i.Summary = make(NaturalLanguageValue)

	i.TotalItems = 0

	return &i
}

// Append adds an element to an OutboxStream
func (o *OutboxStream) Append(ob ObjectOrLink) error {
	o.OrderedItems = append(o.OrderedItems, ob)
	o.TotalItems++
	return nil
}

// Append adds an element to an Outbox
func (o *Outbox) Append(ob ObjectOrLink) error {
	o.OrderedItems = append(o.OrderedItems, ob)
	o.TotalItems++
	return nil
}

// GetID returns the ObjectID corresponding to the OutboxStream
func (o OutboxStream) GetID() ObjectID {
	return o.ID
}

// GetType returns the OutboxStream's type
func (o OutboxStream) GetType() ActivityVocabularyType {
	return o.Type
}

// IsLink returns false for an OutboxStream object
func (o OutboxStream) IsLink() bool {
	return false
}

// GetID returns the ObjectID corresponding to Outbox
func (o Outbox) GetID() ObjectID {
	return o.ID
}

// GetType returns the Outbox's type
func (o Outbox) GetType() ActivityVocabularyType {
	return o.Type
}

// IsLink returns false for an Outbox object
func (o Outbox) IsLink() bool {
	return false
}
