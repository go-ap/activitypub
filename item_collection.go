package activitypub

import (
	"errors"
	"reflect"
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
func (i ItemCollection) Contains(r Item) bool {
	if len(i) == 0 {
		return false
	}
	for _, it := range i {
		if ItemsEqual(it, r) {
			return true
		}
	}
	return false
}

// Contains verifies if IRIs array contains the received one
func (i *ItemCollection) Remove(r Item) {
	li := len(*i)
	if li == 0 {
		return
	}
	if r == nil {
		return
	}
	remIdx := -1
	for idx, it := range *i {
		if ItemsEqual(it, r) {
			remIdx = idx
		}
	}
	if remIdx == -1 {
		return
	}
	if remIdx < li - 1 {
		*i = append((*i)[:remIdx], (*i)[remIdx+1:]...)
	} else {
		*i = (*i)[:remIdx]
	}
}

// ItemCollectionDeduplication normalizes the received arguments lists into a single unified one
func ItemCollectionDeduplication(recCols ...*ItemCollection) ItemCollection {
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
				testIt = cur.GetID()
			} else if cur.IsLink() {
				testIt = cur.GetLink()
			} else {
				continue
			}
			for _, it := range rec {
				if testIt.Equals(it.GetID(), false) {
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
	return rec
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
	default:
		// NOTE(marius): this is an ugly way of dealing with the interface conversion error: types from different scopes
		typ := reflect.TypeOf(new(ItemCollection))
		if reflect.TypeOf(it).ConvertibleTo(typ) {
			if i, ok := reflect.ValueOf(it).Convert(typ).Interface().(*ItemCollection); ok {
				return i, nil
			}
		}
	}
	return nil, errors.New("unable to convert to item collection")
}

// ItemsMatch
func (i ItemCollection) ItemsMatch(col ...Item) bool {
	for _, it := range col {
		if match := i.Contains(it); !match {
			return false
		}
	}
	return true
}

// Equals
func (i ItemCollection) Equals(with Item) bool {
	if with == nil {
		return false
	}
	if !with.IsCollection() {
		return false
	}
	if with.GetType() != CollectionOfItems {
		return false
	}
	result := true
	OnItemCollection(with, func(w *ItemCollection) error {
		if w.Count() != i.Count() {
			result = false
			return nil
		}
		for _, it := range i {
			if !w.Contains(it.GetLink()) {
				result = false
				return nil
			}
		}
		return nil
	})
	return result
}
