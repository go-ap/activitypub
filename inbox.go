package activitypub

type (
	// InboxStream contains all activities received by the actor.
	// The server SHOULD filter content according to the requester's permission.
	// In general, the owner of an inbox is likely to be able to access all of their inbox contents.
	// Depending on access control, some other content may be public, whereas other content may
	// require authentication for non-owner users, if they can access the inbox at all.
	InboxStream = Inbox

	// Inbox is a type alias for an Ordered Collection
	Inbox = OrderedCollection
)

// InboxNew initializes a new Inbox
func InboxNew() *OrderedCollection {
	id := ID("inbox")

	i := OrderedCollection{ID: id, Type: CollectionType}
	i.Name = NaturalLanguageValuesNew()
	i.Content = NaturalLanguageValuesNew()

	i.TotalItems = 0

	return &i
}
