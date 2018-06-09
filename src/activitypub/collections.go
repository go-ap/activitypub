package activitypub

var validCollectionTypes = [...]ActivityVocabularyType{CollectionType, OrderedCollectionType}

// Page represents a Web Page.
type Page ObjectOrLink

type CollectionInterface interface {
	Append(o ObjectOrLink) error
}

// Collection is a subtype of GetObject that represents ordered or unordered sets of GetObject or GetLink instances.
type Collection struct {
	BaseObject
	// A non-negative integer specifying the total number of objects contained by the logical view of the collection.
	// This number might not reflect the actual number of items serialized within the Collection object instance.
	TotalItems uint `jsonld:"totalItems,omitempty"`
	// Identifies the items contained in a collection. The items might be ordered or unordered.
	Items ItemCollection `jsonld:"items,omitempty"`
}

// OrderedCollection is a subtype of Collection in which members of the logical
// collection are assumed to always be strictly ordered.
type OrderedCollection struct {
	BaseObject
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
	PartOf *Collection
	// In a paged Collection, indicates the page that contains the most recently updated member items.
	Current Page `jsonld:"current,omitempty"`
	// In a paged Collection, indicates the furthest preceeding page of items in the collection.
	First Page `jsonld:"first,omitempty"`
	// In a paged Collection, indicates the furthest proceeding page of the collection.
	Last Page `jsonld:"last,omitempty"`
	// In a paged Collection, indicates the next page of items.
	Next Page `jsonld:"next,omitempty"`
	// In a paged Collection, identifies the previous page of items.
	Prev Page `jsonld:"prev,omitempty"`
}

// OrderedCollectionPage type extends from both CollectionPage and OrderedCollection.
// In addition to the properties inherited from each of those, the OrderedCollectionPage
// may contain an additional startIndex property whose value indicates the relative index position
// of the first item contained by the page within the OrderedCollection to which the page belongs.
type OrderedCollectionPage struct {
	PartOf *OrderedCollection
	// In a paged Collection, indicates the page that contains the most recently updated member items.
	Current Page `jsonld:"current,omitempty"`
	// In a paged Collection, indicates the furthest preceeding page of items in the collection.
	First Page `jsonld:"first,omitempty"`
	// In a paged Collection, indicates the furthest proceeding page of the collection.
	Last Page `jsonld:"last,omitempty"`
	// In a paged Collection, indicates the next page of items.
	Next Page `jsonld:"next,omitempty"`
	// In a paged Collection, identifies the previous page of items.
	Prev Page `jsonld:"prev,omitempty"`
	// A non-negative integer value identifying the relative position within the logical view of a strictly ordered collection.
	StartIndex uint `jsonld:"startIndex,omitempty"`
}

// ValidCollectionType validates against the valid collection types
func ValidCollectionType(_type ActivityVocabularyType) bool {
	for _, v := range validCollectionTypes {
		if v == _type {
			return true
		}
	}
	return false
}

// CollectionNew initializes a new Collection
func CollectionNew(id ObjectID) *Collection {
	o := ObjectNew(id, CollectionType)

	return &Collection{BaseObject: o}
}

// CollectionNew initializes a new Collection
func OrderedCollectionNew(id ObjectID) *OrderedCollection {
	o := ObjectNew(id, OrderedCollectionType)

	return &OrderedCollection{BaseObject: o}
}

// CollectionNew initializes a new Collection
func CollectionPageNew(parent *Collection) *CollectionPage {
	return &CollectionPage{PartOf: parent}
}

// CollectionNew initializes a new Collection
func OrderedCollectionPageNew(parent *OrderedCollection) *OrderedCollectionPage {
	return &OrderedCollectionPage{PartOf: parent}
}

// Append adds an element to an OrderedCollection
func (c *OrderedCollection) Append(o ObjectOrLink) error {
	c.OrderedItems = append(c.OrderedItems, o)
	c.TotalItems++
	return nil
}

// Append adds an element to an Collection
func (c *Collection) Append(o ObjectOrLink) error {
	c.Items = append(c.Items, o)
	c.TotalItems++
	return nil
}
