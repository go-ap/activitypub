package activitypub

import (
	"errors"
	"sort"
)

// ItemCollection represents an array of items
type ItemCollection []Item

// GetID returns the ID corresponding to ItemCollection
func (i ItemCollection) GetID() ID {
	return EmptyID
}

// GetLink returns the empty IRI
func (i ItemCollection) GetLink() IRI {
	return EmptyIRI
}

// GetType returns the ItemCollection's type
func (i ItemCollection) GetType() ActivityVocabularyType {
	return CollectionOfItems
}

// IsLink returns false for an ItemCollection object
func (i ItemCollection) IsLink() bool {
	return false
}

// IsObject returns true for a ItemCollection object
func (i ItemCollection) IsObject() bool {
	return false
}

func (i ItemCollection) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0)
	writeItemCollectionValue(&b, i)
	return b, nil
}

// Append facilitates adding elements to Item arrays
// and ensures ItemCollection implements the Collection interface
func (i *ItemCollection) Append(o Item) error {
	oldLen := len(*i)
	d := make(ItemCollection, oldLen+1)
	for k, it := range *i {
		d[k] = it
	}
	d[oldLen] = o
	*i = d
	return nil
}

// Count returns the length of Items in the item collection
func (i *ItemCollection) Count() uint {
	return uint(len(*i))
}

// First returns the ID corresponding to ItemCollection
func (i ItemCollection) First() Item {
	if len(i) == 0 {
		return nil
	}
	return i[0]
}

// Collection returns the current object as collection interface
func (i *ItemCollection) Collection() ItemCollection {
	return *i
}

// IsCollection returns true for ItemCollection arrays
func (i ItemCollection) IsCollection() bool {
	return true
}

// Contains verifies if IRIs array contains the received one
func (i ItemCollection) Contains(r IRI) bool {
	if len(i) == 0 {
		return false
	}
	for _, iri := range i {
		if r.Equals(iri.GetLink(), false) {
			return true
		}
	}
	return false
}

// ItemCollectionDeduplication normalizes the received arguments lists into a single unified one
func ItemCollectionDeduplication(recCols ...*ItemCollection) (ItemCollection, error) {
	rec := make(ItemCollection, 0)

	for _, recCol := range recCols {
		if recCol == nil {
			continue
		}

		toRemove := make([]int, 0)
		for i, cur := range *recCol {
			save := true
			if cur == nil {
				continue
			}
			var testIt IRI
			if cur.IsObject() {
				testIt = IRI(cur.GetID())
			} else if cur.IsLink() {
				testIt = cur.GetLink()
			} else {
				continue
			}
			for _, it := range rec {
				if testIt.Equals(IRI(it.GetID()), false) {
					// mark the element for removal
					toRemove = append(toRemove, i)
					save = false
				}
			}
			if save {
				rec = append(rec, testIt)
			}
		}

		sort.Sort(sort.Reverse(sort.IntSlice(toRemove)))
		for _, idx := range toRemove {
			*recCol = append((*recCol)[:idx], (*recCol)[idx+1:]...)
		}
	}
	return rec, nil
}

// FlattenItemCollection flattens the Collection's properties from Object type to IRI
func FlattenItemCollection(c ItemCollection) ItemCollection {
	if c != nil && len(c) > 0 {
		for i, it := range c {
			c[i] = FlattenToIRI(it)
		}
	}
	return c
}

// ToItemCollection
func ToItemCollection(it Item) (*ItemCollection, error) {
	switch i := it.(type) {
	case *ItemCollection:
		return i, nil
	case ItemCollection:
		return &i, nil
	}
	return nil, errors.New("unable to convert to item collection")
}

func (i ItemCollection) ItemMatches(it Item) bool {
	return i.Contains(it.GetLink())
}
