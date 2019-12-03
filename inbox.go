package activitypub

type (
	// InboxStream contains all activities received by the actor.
	// The server SHOULD filter content according to the requester's permission.
	// In general, the owner of an inbox is likely to be able to access all of their inbox contents.
	// Depending on access control, some other content may be public, whereas other content may
	// require authentication for non-owner users, if they can access the inbox at all.
	InboxStream = Inbox

	// Inbox is a type alias for an Ordered Collection
	Inbox OrderedCollection
)

// InboxNew initializes a new Inbox
func InboxNew() *OrderedCollection {
	id := ObjectID("inbox")

	i := OrderedCollection{ID: id, Type: CollectionType}
	i.Name = NaturalLanguageValuesNew()
	i.Content = NaturalLanguageValuesNew()

	i.TotalItems = 0

	return &i
}

// Append adds an element to an Inbox
func (i *Inbox) Append(ob Item) error {
	i.OrderedItems = append(i.OrderedItems, ob)
	i.TotalItems++
	return nil
}

// GetID returns the ObjectID corresponding to Inbox
func (i Inbox) GetID() ObjectID {
	return i.Collection().GetID()
}

// GetLink returns the IRI corresponding to the current Inbox object
func (i Inbox) GetLink() IRI {
	return IRI(i.ID)
}

// GetType returns the Inbox's type
func (i Inbox) GetType() ActivityVocabularyType {
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
	if ItemTyperFunc == nil {
		ItemTyperFunc = JSONGetItemByType
	}
	c := OrderedCollection(*i)
	err := c.UnmarshalJSON(data)

	*i = Inbox(c)

	return err
}

// Collection returns the underlying Collection type
func (i Inbox) Collection() CollectionInterface {
	c := OrderedCollection(i)
	return &c
}
