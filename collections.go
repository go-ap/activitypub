package activitystreams

import (
	"github.com/buger/jsonparser"
)

var validCollectionTypes = [...]ActivityVocabularyType{
	CollectionType,
	OrderedCollectionType,
	CollectionPageType,
	OrderedCollectionPageType,
}

type CollectionInterface interface {
	ObjectOrLink
	Collection() CollectionInterface
	Append(ob Item) error
}

// Collection is a subtype of Activity Pub Object that represents ordered or unordered sets of Activity Pub Object or Link instances.
type Collection struct {
	Parent
	// In a paged Collection, indicates the page that contains the most recently updated member items.
	Current ObjectOrLink `jsonld:"current,omitempty"`
	// In a paged Collection, indicates the furthest preceeding page of items in the collection.
	First ObjectOrLink `jsonld:"first,omitempty"`
	// In a paged Collection, indicates the furthest proceeding page of the collection.
	Last ObjectOrLink `jsonld:"last,omitempty"`
	// A non-negative integer specifying the total number of objects contained by the logical view of the collection.
	// This number might not reflect the actual number of items serialized within the Collection object instance.
	TotalItems uint `jsonld:"totalItems,omitempty"`
	// Identifies the items contained in a collection. The items might be ordered or unordered.
	Items ItemCollection `jsonld:"items,omitempty"`
}

// OrderedCollection is a subtype of Collection in which members of the logical
// collection are assumed to always be strictly ordered.
type OrderedCollection struct {
	Parent
	// In a paged Collection, indicates the page that contains the most recently updated member items.
	Current ObjectOrLink `jsonld:"current,omitempty"`
	// In a paged Collection, indicates the furthest preceeding page of items in the collection.
	First ObjectOrLink `jsonld:"first,omitempty"`
	// In a paged Collection, indicates the furthest proceeding page of the collection.
	Last ObjectOrLink `jsonld:"last,omitempty"`
	// A non-negative integer specifying the total number of objects contained by the logical view of the collection.
	// This number might not reflect the actual number of items serialized within the Collection object instance.
	TotalItems uint `jsonld:"totalItems,omitempty"`
	// Identifies the items contained in a collection. The items might be ordered or unordered.
	OrderedItems ItemCollection `jsonld:"orderedItems,omitempty"`
}

// CollectionPage is a Collection that contains a large number of items and when it becomes impractical
// for an implementation to serialize every item contained by a Collection using the items (or orderedItems)
// property alone. In such cases, the items within a Collection can be divided into distinct subsets or "pages".
type CollectionPage struct {
	Collection
	// Identifies the Collection to which a CollectionPage objects items belong.
	PartOf Item `jsonld:"partOf,omitempty"`
	// In a paged Collection, indicates the next page of items.
	Next Item `jsonld:"next,omitempty"`
	// In a paged Collection, identifies the previous page of items.
	Prev Item `jsonld:"prev,omitempty"`
}

// OrderedCollectionPage type extends from both CollectionPage and OrderedCollection.
// In addition to the properties inherited from each of those, the OrderedCollectionPage
// may contain an additional startIndex property whose value indicates the relative index position
// of the first item contained by the page within the OrderedCollection to which the page belongs.
type OrderedCollectionPage struct {
	OrderedCollection
	// Identifies the Collection to which a CollectionPage objects items belong.
	PartOf Item `jsonld:"partOf,omitempty"`
	// In a paged Collection, indicates the next page of items.
	Next Item `jsonld:"next,omitempty"`
	// In a paged Collection, identifies the previous page of items.
	Prev Item `jsonld:"prev,omitempty"`
	// A non-negative integer value identifying the relative position within the logical view of a strictly ordered collection.
	StartIndex uint `jsonld:"startIndex,omitempty"`
}

// ValidCollectionType validates against the valid collection types
func ValidCollectionType(typ ActivityVocabularyType) bool {
	for _, v := range validCollectionTypes {
		if v == typ {
			return true
		}
	}
	return false
}

// CollectionNew initializes a new Collection
func CollectionNew(id ObjectID) *Collection {
	c := Collection{Parent: Parent{ID: id, Type: CollectionType}}
	c.Name = NaturalLanguageValueNew()
	c.Content = NaturalLanguageValueNew()
	c.Summary = NaturalLanguageValueNew()
	return &c
}

// OrderedCollectionNew initializes a new OrderedCollection
func OrderedCollectionNew(id ObjectID) *OrderedCollection {
	o := OrderedCollection{Parent: Parent{ID: id, Type: OrderedCollectionType}}
	o.Name = NaturalLanguageValueNew()
	o.Content = NaturalLanguageValueNew()

	return &o
}

// CollectionNew initializes a new CollectionPage
func CollectionPageNew(parent CollectionInterface) *CollectionPage {
	p := CollectionPage{
		PartOf: parent.GetLink(),
	}
	if pc, ok := parent.(*Collection); ok {
		p.Collection = *pc
	}
	p.Type = CollectionPageType
	return &p
}

// OrderedCollectionPageNew initializes a new OrderedCollectionPage
func OrderedCollectionPageNew(parent CollectionInterface) *OrderedCollectionPage {
	p := OrderedCollectionPage{
		PartOf: parent.GetLink(),
	}
	if pc, ok := parent.(*OrderedCollection); ok {
		p.OrderedCollection = *pc
	}
	p.Type = OrderedCollectionPageType
	return &p
}

// Append adds an element to an OrderedCollection
func (o *OrderedCollection) Append(ob Item) error {
	o.OrderedItems = append(o.OrderedItems, ob)
	o.TotalItems++
	return nil
}

// Append adds an element to a Collection
func (c *Collection) Append(ob Item) error {
	c.Items = append(c.Items, ob)
	c.TotalItems++
	return nil
}

// Append adds an element to an OrderedCollectionPage
func (o *OrderedCollectionPage) Append(ob Item) error {
	o.OrderedItems = append(o.OrderedItems, ob)
	o.TotalItems++
	return nil
}

// Append adds an element to a CollectionPage
func (c *CollectionPage) Append(ob Item) error {
	c.Items = append(c.Items, ob)
	c.TotalItems++
	return nil
}

// GetType returns the Collection's type
func (c Collection) GetType() ActivityVocabularyType {
	return c.Type
}

// IsLink returns false for a Collection object
func (c Collection) IsLink() bool {
	return false
}

// GetID returns the ObjectID corresponding to the Collection object
func (c Collection) GetID() *ObjectID {
	return &c.ID
}

// GetLink returns the IRI corresponding to the Collection object
func (c Collection) GetLink() IRI {
	return IRI(c.ID)
}

// IsObject returns true for a Collection object
func (c Collection) IsObject() bool {
	return true
}

// GetType returns the OrderedCollection's type
func (o OrderedCollection) GetType() ActivityVocabularyType {
	return o.Type
}

// IsLink returns false for an OrderedCollection object
func (o OrderedCollection) IsLink() bool {
	return false
}

// GetID returns the ObjectID corresponding to the OrderedCollection
func (o OrderedCollection) GetID() *ObjectID {
	return &o.ID
}

// GetLink returns the IRI corresponding to the OrderedCollection object
func (o OrderedCollection) GetLink() IRI {
	return IRI(o.ID)
}

// IsObject returns true for am OrderedCollection object
func (o OrderedCollection) IsObject() bool {
	return true
}

// UnmarshalJSON
func (o *OrderedCollection) UnmarshalJSON(data []byte) error {
	o.Parent.UnmarshalJSON(data)

	o.TotalItems = uint(getAPInt(data, "totalItems"))
	o.OrderedItems = getAPItems(data, "orderedItems")

	o.Current = getAPItem(data, "current")
	o.First = getAPItem(data, "first")
	o.Last = getAPItem(data, "last")

	return nil
}

// UnmarshalJSON
func (c *Collection) UnmarshalJSON(data []byte) error {
	c.Parent.UnmarshalJSON(data)

	c.TotalItems = uint(getAPInt(data, "totalItems"))
	c.Items = getAPItems(data, "items")

	c.Current = getAPItem(data, "current")
	c.First = getAPItem(data, "first")
	c.Last = getAPItem(data, "last")

	return nil
}

// UnmarshalJSON
func (o *OrderedCollectionPage) UnmarshalJSON(data []byte) error {
	o.OrderedCollection.UnmarshalJSON(data)

	o.Next = getAPItem(data, "next")
	o.Prev = getAPItem(data, "prev")
	o.PartOf = getAPItem(data, "partOf")

	if si, err := jsonparser.GetInt(data, "startIndex"); err != nil {
		o.StartIndex = uint(si)
	}
	return nil
}

// UnmarshalJSON
func (c *CollectionPage) UnmarshalJSON(data []byte) error {
	c.Collection.UnmarshalJSON(data)

	c.Next = getAPItem(data, "next")
	c.Prev = getAPItem(data, "prev")
	c.PartOf = getAPItem(data, "partOf")

	return nil
}

/*
func (c *Collection) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (o *OrderedCollection) MarshalJSON() ([]byte, error) {
	return nil, nil
}
*/

// Collection returns the underlying Collection type
func (c *Collection) Collection() CollectionInterface {
	return c
}

// Collection returns the underlying Collection type
func (o *OrderedCollection) Collection() CollectionInterface {
	return o
}
