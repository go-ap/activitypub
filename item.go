package activitypub

import (
	"fmt"
	"reflect"
	"strings"
)

// Item struct
type Item = ObjectOrLink

const (
	// EmptyIRI represents a zero length IRI
	EmptyIRI IRI = ""
	// NilIRI represents by convention an IRI which is nil
	// Its use is mostly to check if a property of an ActivityPub Item is nil
	NilIRI IRI = "-"

	// EmptyID represents a zero length ID
	EmptyID = EmptyIRI
	// NilID represents by convention an ID which is nil, see details of NilIRI
	NilID = NilIRI
)

func itemsNeedSwapping(i1, i2 Item) bool {
	if IsIRI(i1) && !IsIRI(i2) {
		return true
	}
	t1 := i1.GetType()
	t2 := i2.GetType()
	if ObjectTypes.MatchOther(t2) {
		return !ObjectTypes.MatchOther(t1)
	}
	return false
}

// ItemsEqual checks if it and with Items are equal
func ItemsEqual(it, with Item) bool {
	if IsNil(it) || IsNil(with) {
		return IsNil(with) && IsNil(it)
	}
	if itemsNeedSwapping(it, with) {
		return ItemsEqual(with, it)
	}
	result := false
	if IsIRI(with) || IsIRI(it) {
		// NOTE(marius): I'm not sure this logic is sound:
		// if only one item is an IRI it should not be equal to the other even if it has the same ID
		result = it.GetLink().Equals(with.GetLink(), false)
	} else if IsItemCollection(it) {
		if !IsItemCollection(with) {
			return false
		}
		_ = OnItemCollection(it, func(c *ItemCollection) error {
			result = c.Equals(with)
			return nil
		})
	} else if IsLink(it) {
		_ = OnLink(it, func(l *Link) error {
			result = l.Equals(with)
			return nil
		})
	} else if IsObject(it) {
		_ = OnObject(it, func(i *Object) error {
			result = i.Equals(with)
			return nil
		})
		if ActivityTypes.MatchOther(with.GetType()) {
			_ = OnActivity(it, func(i *Activity) error {
				result = i.Equals(with)
				return nil
			})
		} else if ActorTypes.MatchOther(with.GetType()) {
			_ = OnActor(it, func(i *Actor) error {
				result = i.Equals(with)
				return nil
			})
		} else if it.IsCollection() {
			if it.GetType().Matches(CollectionType) {
				_ = OnCollection(it, func(c *Collection) error {
					result = c.Equals(with)
					return nil
				})
			}
			if it.GetType().Matches(OrderedCollectionType) {
				_ = OnOrderedCollection(it, func(c *OrderedCollection) error {
					result = c.Equals(with)
					return nil
				})
			}
			if it.GetType().Matches(CollectionPageType) {
				_ = OnCollectionPage(it, func(c *CollectionPage) error {
					result = c.Equals(with)
					return nil
				})
			}
			if it.GetType().Matches(OrderedCollectionPageType) {
				_ = OnOrderedCollectionPage(it, func(c *OrderedCollectionPage) error {
					result = c.Equals(with)
					return nil
				})
			}
		}
	}
	return result
}

// IsItemCollection returns if the current Item interface holds a Collection
func IsItemCollection(it LinkOrIRI) bool {
	if _, ok := it.(ItemCollection); ok {
		return ok
	}
	if _, ok := it.(*ItemCollection); ok {
		return ok
	}
	return IsIRIs(it)
}

// IsIRI returns if the current Item interface holds an IRI
func IsIRI(it LinkOrIRI) bool {
	if it == nil {
		return false
	}
	if _, ok := it.(IRI); ok {
		return true
	}
	if iri, ok := it.(*IRI); ok {
		return iri != nil
	}
	return false
}

// IsIRIs returns if the current Item interface holds an IRI slice
func IsIRIs(it LinkOrIRI) bool {
	if it == nil {
		return false
	}
	if iris, ok := it.(IRIs); ok {
		return iris != nil
	}
	if iris, ok := it.(*IRIs); ok {
		return iris != nil
	}
	return false
}

// IsLink returns if the current Item interface holds a Link
func IsLink(it LinkOrIRI) bool {
	if it == nil {
		return false
	}
	if _, ok := it.(Link); ok {
		return true
	}
	if l, ok := it.(*Link); ok {
		return l != nil
	}
	return false
}

// IsObject returns if the current Item interface holds an Object
func IsObject(it LinkOrIRI) bool {
	if it == nil {
		return false
	}
	switch ob := it.(type) {
	case Object:
		return true
	case *Object:
		return ob != nil
	case Actor:
		return true
	case *Actor:
		return ob != nil
	case Profile:
		return true
	case *Profile:
		return ob != nil
	case Place:
		return true
	case *Place:
		return ob != nil
	case Relationship:
		return true
	case *Relationship:
		return ob != nil
	case Tombstone:
		return true
	case *Tombstone:
		return ob != nil
	case Activity:
		return true
	case *Activity:
		return ob != nil
	case IntransitiveActivity:
		return true
	case *IntransitiveActivity:
		return ob != nil
	case Question:
		return true
	case *Question:
		return ob != nil
	case Collection:
		return true
	case *Collection:
		return ob != nil
	case CollectionPage:
		return true
	case *CollectionPage:
		return ob != nil
	case OrderedCollection:
		return true
	case *OrderedCollection:
		return ob != nil
	case OrderedCollectionPage:
		return true
	case *OrderedCollectionPage:
		return ob != nil
	default:
		return false
	}
}

// IsNil checks if the object matching an ObjectOrLink interface is nil
func IsNil(it LinkOrIRI) bool {
	if it == nil {
		return true
	}
	// This is the default if the argument can't be cast to Object, as is the case for an ItemCollection
	isNil := false
	if IsIRI(it) {
		var l IRI
		if lp, ok := it.(*IRI); ok {
			l = *lp
		} else {
			l, _ = it.(IRI)
		}
		isNil = len(l) == 0 || strings.EqualFold(l.String(), NilIRI.String())
	} else if IsItemCollection(it) {
		if v, ok := it.(ItemCollection); ok {
			return v == nil
		}
		if v, ok := it.(*ItemCollection); ok {
			return v == nil
		}
		if v, ok := it.(IRIs); ok {
			return v == nil
		}
		if v, ok := it.(*IRIs); ok {
			return v == nil
		}
	} else if IsObject(it) {
		if ob, ok := it.(Item); ok {
			_ = OnObject(ob, func(o *Object) error {
				isNil = o == nil
				return nil
			})
		}
	} else if IsLink(it) {
		_ = OnLink(it, func(l *Link) error {
			isNil = l == nil
			return nil
		})
	} else {
		// NOTE(marius): we're not dealing with a type that we know about, so we use slow reflection
		// as we still care about the result
		v := reflect.ValueOf(it)
		isNil = v.Kind() == reflect.Pointer && v.IsNil()
	}
	return isNil
}

func ErrorInvalidType[T Objects | Links](received LinkOrIRI) error {
	return fmt.Errorf("unable to convert %T to %T", received, new(T))
}

// OnItem runs function "fn" on the Item "it", with the benefit of destructuring "it" to individual
// items if it's actually an ItemCollection or an object holding an ItemCollection
//
// It is expected that the caller handles the logic of dealing with different Item implementations
// internally in "fn".
func OnItem(it Item, fn func(Item) error) error {
	if it == nil {
		return nil
	}
	if !IsItemCollection(it) {
		return fn(it)
	}
	return OnItemCollection(it, func(col *ItemCollection) error {
		for _, it := range *col {
			if err := OnItem(it, fn); err != nil {
				return err
			}
		}
		return nil
	})
}

// OnItemCollection calls function fn on it Item if it can be asserted to type ItemCollection
//
// It should be used when Item represents an Item collection and it's usually used as a way
// to wrap functionality for other functions that will be called on each item in the collection.
func OnItemCollection(it LinkOrIRI, fn WithItemCollectionFn) error {
	if it == nil {
		return nil
	}
	col, err := ToItemCollection(it)
	if err != nil {
		return err
	}
	return fn(col)
}
