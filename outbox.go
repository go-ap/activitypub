package activitypub

type (
	// OutboxStream contains activities the user has published,
	// subject to the ability of the requestor to retrieve the activity (that is,
	// the contents of the outbox are filtered by the permissions of the person reading it).
	OutboxStream = Outbox

	// Outbox is a type alias for an Ordered Collection
	Outbox = OrderedCollection
)

// OutboxNew initializes a new Outbox
func OutboxNew() *Outbox {
	id := ObjectID("outbox")

	i := Outbox{ID: id, Type: OrderedCollectionType}
	i.Name = NaturalLanguageValuesNew()
	i.Content = NaturalLanguageValuesNew()
	i.TotalItems = 0
	i.OrderedItems = make(ItemCollection, 0)

	return &i
}
