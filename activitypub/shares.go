package activitypub

import as "github.com/mariusor/activitypub.go/activitystreams"

type (
	// SharesCollection is a list of all Announce activities with this object as the object property,
	// added as a side effect. The shares collection MUST be either an OrderedCollection or a Collection
	// and MAY be filtered on privileges of an authenticated user or as appropriate when no authentication
	// is given.
	SharesCollection Shares

	// Shares is a type alias for an Ordered Collection
	Shares as.OrderedCollection
)

// SharesNew initializes a new Shares
func SharesNew() *Shares {
	id := as.ObjectID("Shares")

	i := Shares{ID: id, Type: as.OrderedCollectionType}
	i.Name = as.NaturalLanguageValueNew()
	i.Content = as.NaturalLanguageValueNew()

	i.TotalItems = 0

	return &i
}

// Append adds an element to an SharesCollection
func (o *SharesCollection) Append(ob as.Item) error {
	o.OrderedItems = append(o.OrderedItems, ob)
	o.TotalItems++
	return nil
}

// Append adds an element to an Shares
func (o *Shares) Append(ob as.Item) error {
	o.OrderedItems = append(o.OrderedItems, ob)
	o.TotalItems++
	return nil
}

// GetID returns the ObjectID corresponding to the SharesCollection
func (o SharesCollection) GetID() *as.ObjectID {
	return o.Collection().GetID()
}

// GetLink returns the IRI corresponding to the current SharesCollection object
func (o SharesCollection) GetLink() as.IRI {
	return as.IRI(o.ID)
}

// GetType returns the SharesCollection's type
func (o SharesCollection) GetType() as.ActivityVocabularyType {
	return o.Type
}

// IsLink returns false for an SharesCollection object
func (o SharesCollection) IsLink() bool {
	return false
}

// IsObject returns true for a SharesCollection object
func (o SharesCollection) IsObject() bool {
	return true
}

// GetID returns the ObjectID corresponding to Shares
func (o Shares) GetID() *as.ObjectID {
	return o.Collection().GetID()
}

// GetLink returns the IRI corresponding to the current Shares object
func (o Shares) GetLink() as.IRI {
	return as.IRI(o.ID)
}

// GetType returns the Shares's type
func (o Shares) GetType() as.ActivityVocabularyType {
	return o.Type
}

// IsLink returns false for an Shares object
func (o Shares) IsLink() bool {
	return false
}

// IsObject returns true for a Shares object
func (o Shares) IsObject() bool {
	return true
}

// UnmarshalJSON
func (o *SharesCollection) UnmarshalJSON(data []byte) error {
	c := as.OrderedCollection(*o)
	err := c.UnmarshalJSON(data)

	*o = SharesCollection(c)

	return err
}

// UnmarshalJSON
func (o *Shares) UnmarshalJSON(data []byte) error {
	c := as.OrderedCollection(*o)
	err := c.UnmarshalJSON(data)

	*o = Shares(c)

	return err
}

// Collection returns the underlying Collection type
func (o Shares) Collection() as.CollectionInterface {
	c := as.OrderedCollection(o)
	return &c
}

// Collection returns the underlying Collection type
func (o SharesCollection) Collection() as.CollectionInterface {
	c := as.OrderedCollection(o)
	return &c
}
