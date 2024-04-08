package activitypub

import (
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

// MarshalJSON encodes the receiver object to a JSON document.
func (i ItemCollection) MarshalJSON() ([]byte, error) {
	if i == nil {
		return nil, nil
	}
	b := make([]byte, 0)
	JSONWriteItemCollectionValue(&b, i, true)
	return b, nil
}

// Append facilitates adding elements to Item arrays
// and ensures ItemCollection implements the Collection interface
func (i *ItemCollection) Append(it ...Item) error {
	for _, ob := range it {
		if i.Contains(ob) {
			continue
		}
		*i = append(*i, ob)
	}
	return nil
}

// Count returns the length of Items in the item collection
func (i *ItemCollection) Count() uint {
	if i == nil {
		return 0
	}
	return uint(len(*i))
}

// First returns the ID corresponding to ItemCollection
func (i ItemCollection) First() Item {
	if len(i) == 0 {
		return nil
	}
	return i[0]
}

// Normalize returns the first item if the collection contains only one,
// the full collection if the collection contains more than one item,
// or nil
func (i ItemCollection) Normalize() Item {
	if len(i) == 0 {
		return nil
	}
	if len(i) == 1 {
		return i[0]
	}
	return i
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

// Remove removes the r Item from the i ItemCollection if it contains it
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
	if remIdx < li-1 {
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

// ToItemCollection returns the item collection contained as part of OrderedCollection, OrderedCollectionPage,
// Collection and CollectionPage.
// It also converts an IRI slice into an equivalent ItemCollection.
func ToItemCollection(it Item) (*ItemCollection, error) {
	switch i := it.(type) {
	case *ItemCollection:
		return i, nil
	case ItemCollection:
		return &i, nil
	case *OrderedCollection:
		return &i.OrderedItems, nil
	case *OrderedCollectionPage:
		return &i.OrderedItems, nil
	case *Collection:
		return &i.Items, nil
	case *CollectionPage:
		return &i.Items, nil
	case IRIs:
		iris := make(ItemCollection, len(i))
		for j, ob := range i {
			iris[j] = ob
		}
		return &iris, nil
	case *IRIs:
		iris := make(ItemCollection, len(*i))
		for j, ob := range *i {
			iris[j] = ob
		}
		return &iris, nil
	default:
		return reflectItemToType[ItemCollection](it)
	}
	return nil, ErrorInvalidType[ItemCollection](it)
}

// ToIRIs
func ToIRIs(it Item) (*IRIs, error) {
	switch i := it.(type) {
	case *IRIs:
		return i, nil
	case IRIs:
		return &i, nil
	case ItemCollection:
		iris := make(IRIs, len(i))
		for j, ob := range i {
			iris[j] = ob.GetLink()
		}
		return &iris, nil
	case *ItemCollection:
		iris := make(IRIs, len(*i))
		for j, ob := range *i {
			iris[j] = ob.GetLink()
		}
		return &iris, nil
	default:
		return reflectItemToType[IRIs](it)
	}
	return nil, ErrorInvalidType[IRIs](it)
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
	if IsNil(with) {
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

// Clean removes Bto and BCC properties on all the members of the collection
func (i ItemCollection) Clean() {
	for j, it := range i {
		i[j] = CleanRecipients(it)
	}
}

func (i ItemCollection) Recipients() ItemCollection {
	all := make(ItemCollection, 0)
	for _, it := range i {
		_ = OnObject(it, func(ob *Object) error {
			aud := ob.Audience
			_ = all.Append(ItemCollectionDeduplication(&ob.To, &ob.CC, &ob.Bto, &ob.BCC, &aud)...)
			return nil
		})
	}
	return ItemCollectionDeduplication(&all)
}
