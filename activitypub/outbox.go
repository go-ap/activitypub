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
func (o OutboxStream) GetID() ObjectID {
	return o.ID
}
func (o OutboxStream)GetType() ActivityVocabularyType {
	return o.Type
}
