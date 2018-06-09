package activitypub

type (
	// InboxStream contains all activities received by the actor.
	// The server SHOULD filter content according to the requester's permission.
	// In general, the owner of an inbox is likely to be able to access all of their inbox contents.
	// Depending on access control, some other content may be public, whereas other content may
	// require authentication for non-owner users, if they can access the inbox at all.
	InboxStream Inbox

	// Inbox is a type alias for an Ordered Collection
	Inbox OrderedCollection
)

// InboxNew initializes a new Inbox
func InboxNew() *Inbox {
	id := ObjectID("inbox")

	i := Inbox{ID: id, Type: OrderedCollectionType}
	i.Name = make(NaturalLanguageValue)
	i.Content = make(NaturalLanguageValue)
	i.Summary = make(NaturalLanguageValue)

	i.TotalItems = 0

	return &i
}

// Append adds an element to an InboxStream
func (i *InboxStream) Append(o ObjectOrLink) error {
	i.OrderedItems = append(i.OrderedItems, o)
	i.TotalItems++
	return nil
}
