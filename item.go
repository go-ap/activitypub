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
	if ObjectTypes.Contains(t2) {
		return !ObjectTypes.Contains(t1)
	}
	return false
}

// ItemsEqual checks if it and with Items are equal
func ItemsEqual(it, with Item) bool {
	if IsNil(it) || IsNil(with) {
		return with == it
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
		_ = OnItemCollection(it, func(c *ItemCollection) error {
			result = c.Equals(with)
			return nil
		})
	} else if IsObject(it) {
		_ = OnObject(it, func(i *Object) error {
			result = i.Equals(with)
			return nil
		})
		if ActivityTypes.Contains(with.GetType()) {
			_ = OnActivity(it, func(i *Activity) error {
				result = i.Equals(with)
				return nil
			})
		} else if ActorTypes.Contains(with.GetType()) {
			_ = OnActor(it, func(i *Actor) error {
				result = i.Equals(with)
				return nil
			})
		} else if it.IsCollection() {
			if it.GetType() == CollectionType {
				_ = OnCollection(it, func(c *Collection) error {
					result = c.Equals(with)
					return nil
				})
			}
			if it.GetType() == OrderedCollectionType {
				_ = OnOrderedCollection(it, func(c *OrderedCollection) error {
					result = c.Equals(with)
					return nil
				})
			}
			if it.GetType() == CollectionPageType {
				_ = OnCollectionPage(it, func(c *CollectionPage) error {
					result = c.Equals(with)
					return nil
				})
			}
			if it.GetType() == OrderedCollectionPageType {
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
func IsItemCollection(it Item) bool {
	_, ok := it.(ItemCollection)
	_, okP := it.(*ItemCollection)
	return ok || okP || IsIRIs(it)
}

// IsIRI returns if the current Item interface holds an IRI
func IsIRI(it Item) bool {
	_, okV := it.(IRI)
	_, okP := it.(*IRI)
	return okV || okP
}

// IsIRIs returns if the current Item interface holds an IRI slice
func IsIRIs(it Item) bool {
	_, okV := it.(IRIs)
	_, okP := it.(*IRIs)
	return okV || okP
}

// IsLink returns if the current Item interface holds a Link
func IsLink(it Item) bool {
	_, okV := it.(Link)
	_, okP := it.(*Link)
	return okV || okP
}

// IsObject returns if the current Item interface holds an Object
func IsObject(it Item) bool {
	switch ob := it.(type) {
	case Actor, *Actor,
		Object, *Object, Profile, *Profile, Place, *Place, Relationship, *Relationship, Tombstone, *Tombstone,
		Activity, *Activity, IntransitiveActivity, *IntransitiveActivity, Question, *Question,
		Collection, *Collection, CollectionPage, *CollectionPage,
		OrderedCollection, *OrderedCollection, OrderedCollectionPage, *OrderedCollectionPage:
		return ob != nil
	default:
		return false
	}
}

// IsNil checks if the object matching an ObjectOrLink interface is nil
func IsNil(it Item) bool {
	if it == nil {
		return true
	}
	// This is the default if the argument can't be cast to Object, as is the case for an ItemCollection
	isNil := false
	if IsIRI(it) {
		isNil = len(it.GetLink()) == 0 || strings.EqualFold(it.GetLink().String(), NilIRI.String())
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
		OnObject(it, func(o *Object) error {
			isNil = o == nil
			return nil
		})
	} else if IsLink(it) {
		OnLink(it, func(l *Link) error {
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

func ErrorInvalidType[T Objects | Links | IRIs](received Item) error {
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
