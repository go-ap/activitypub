package activitypub

import (
	"fmt"
	"reflect"
	"slices"
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

func compareByType(typ Typer, it, with ObjectOrLink) bool {
	result := false
	iTyp := it.GetType()
	if QuestionType.Match(iTyp) {
		_ = OnQuestion(it, func(q *Question) error {
			result = q.Equals(with)
			return nil
		})
	} else if IntransitiveActivityTypes.Match(iTyp) {
		_ = OnIntransitiveActivity(it, func(i *IntransitiveActivity) error {
			result = i.Equals(with)
			return nil
		})
	} else if ActivityTypes.Match(iTyp) {
		_ = OnActivity(it, func(i *Activity) error {
			result = i.Equals(with)
			return nil
		})
	} else if ActorTypes.Match(iTyp) {
		_ = OnActor(it, func(i *Actor) error {
			result = i.Equals(with)
			return nil
		})
	} else if CollectionTypes.Match(iTyp) {
		if CollectionType.Match(it.GetType()) {
			_ = OnCollection(it, func(c *Collection) error {
				result = c.Equals(with)
				return nil
			})
		}
		if OrderedCollectionType.Match(iTyp) {
			_ = OnOrderedCollection(it, func(c *OrderedCollection) error {
				result = c.Equals(with)
				return nil
			})
		}
		if CollectionPageType.Match(iTyp) {
			_ = OnCollectionPage(it, func(c *CollectionPage) error {
				result = c.Equals(with)
				return nil
			})
		}
		if OrderedCollectionPageType.Match(iTyp) {
			_ = OnOrderedCollectionPage(it, func(c *OrderedCollectionPage) error {
				result = c.Equals(with)
				return nil
			})
		}
	} else {
		_ = OnObject(it, func(i *Object) error {
			result = i.Equals(with)
			return nil
		})
	}
	return result
}

// ItemsEqual checks if it and with Items are equal
func ItemsEqual(it, with Item) bool {
	if IsNil(it) || IsNil(with) {
		return IsNil(with) && IsNil(it)
	}

	if ita, ok := it.(interface{ Equals(Item) bool }); ok {
		eq := ita.Equals(with)
		return eq
	}

	result := false
	switch {
	case IsIRI(with) || IsIRI(it):
		ii, ok := it.(IRI)
		iw, wok := with.(IRI)
		result = ok && wok && ii.Equal(iw)
	case IsIRIs(it):
		if !IsIRIs(with) {
			return false
		}
		iIRIs, _ := ToIRIs(it)
		wIRIs, _ := ToIRIs(with)
		return slices.Equal(*iIRIs, *wIRIs)
	case IsItemCollection(it):
		if !IsItemCollection(with) {
			return false
		}
		_ = OnItemCollection(it, func(c *ItemCollection) error {
			result = c.Equals(with)
			return nil
		})
	case IsLink(it):
		_ = OnLink(it, func(l *Link) error {
			result = l.Equals(with)
			return nil
		})
	default:
		itt, oki := it.(ObjectOrLink)
		witht, okw := it.(ObjectOrLink)
		result = (oki && okw) && typedObjectsEqual(itt, witht)
	}
	return result
}

func typedObjectsEqual(it, with ObjectOrLink) bool {
	if !typesEqual(it.GetType(), with.GetType()) {
		return false
	}
	return compareByType(with.GetType(), it, with)
}

func typesEqual(t1, t2 Typer) bool {
	tt1 := t1.AsTypes()
	tt2 := t2.AsTypes()
	if len(tt1) != len(tt2) {
		return false
	}
	slices.Sort(tt1)
	slices.Sort(tt2)
	for i, ti1 := range tt1 {
		if ti1 != tt2[i] {
			return false
		}
	}
	return true
}

// IsCollection returns if the current Item interface holds an ItemCollection, or any of the collection types
func IsCollection(it LinkOrIRI) bool {
	switch it.(type) {
	case Collection:
		return true
	case *Collection:
		return true
	case CollectionPage:
		return true
	case *CollectionPage:
		return true
	case OrderedCollection:
		return true
	case *OrderedCollection:
		return true
	case OrderedCollectionPage:
		return true
	case *OrderedCollectionPage:
		return true
	default:
		return IsItemCollection(it)
	}
}

// IsItemCollection returns if the current Item interface holds a Collection
func IsItemCollection(it LinkOrIRI) bool {
	if _, ok := it.(ItemCollection); ok {
		return ok
	}
	if _, ok := it.(*ItemCollection); ok {
		return ok
	}
	return !IsNil(it) && IsIRIs(it)
}

// IsIRI returns if the current Item interface holds an IRI
func IsIRI(it LinkOrIRI) bool {
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
	_, ok := it.(IRIs)
	_, okp := it.(*IRIs)
	return ok || okp
}

// IsLink returns if the current Item interface holds a Link
func IsLink(it LinkOrIRI) bool {
	if IsNil(it) {
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
	if IsNil(it) {
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
	switch maybeNil := it.(type) {
	case IRI:
		return len(maybeNil) == 0 || strings.EqualFold(maybeNil.String(), NilIRI.String())
	case *IRI:
		return maybeNil == nil || len(*maybeNil) == 0 || strings.EqualFold(maybeNil.String(), NilIRI.String())
	case IRIs:
		return maybeNil == nil
	case *IRIs:
		return maybeNil == nil
	case ItemCollection:
		return maybeNil == nil
	case *ItemCollection:
		return maybeNil == nil
	}
	// NOTE(marius): we're not dealing with a type that we know about, so we use slow reflection
	// as we still care about the result
	v := reflect.ValueOf(it)
	return v.Kind() == reflect.Pointer && v.IsNil()
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
	if IsNil(it) {
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

// NotEmpty tells us if an Item interface value has a non nil value for various types
// that implement
func NotEmpty(it Item) bool {
	if IsNil(it) {
		return false
	}
	var notEmpty bool
	switch {
	case IsIRI(it):
		notEmpty = len(it.GetLink()) > 0
	case IsIRIs(it):
		_ = OnIRIs(it, func(is *IRIs) error {
			notEmpty = len(*is) > 0
			return nil
		})
	case IsItemCollection(it):
		_ = OnItemCollection(it, func(ic *ItemCollection) error {
			notEmpty = len(*ic) > 0
			return nil
		})
	default:
		if itt, ok := it.(ObjectOrLink); ok {
			notEmpty = emptyByType(itt)
		}
	}
	return notEmpty
}

func emptyByType(it ObjectOrLink) bool {
	var notEmpty bool
	typ := it.GetType()
	switch {
	case QuestionType.Match(typ):
		_ = OnQuestion(it, func(q *Question) error {
			notEmpty = notEmptyQuestion(q)
			return nil
		})
	case IntransitiveActivityTypes.Match(typ):
		_ = OnIntransitiveActivity(it, func(a *IntransitiveActivity) error {
			notEmpty = notEmptyIntransitiveActivity(a)
			return nil
		})
	case ActivityTypes.Match(typ):
		_ = OnActivity(it, func(a *Activity) error {
			notEmpty = notEmptyActivity(a)
			return nil
		})
	case CollectionTypes.Match(typ):
		_ = OnCollectionIntf(it, func(c CollectionInterface) error {
			notEmpty = c != nil || len(c.Collection()) > 0
			return nil
		})
	case ActorTypes.Match(typ):
		_ = OnActor(it, func(a *Actor) error {
			notEmpty = notEmptyActor(a)
			return nil
		})
	case LinkTypes.Match(typ):
		_ = OnLink(it, func(l *Link) error {
			notEmpty = notEmptyLink(l)
			return nil
		})
	default:
		_ = OnObject(it, func(o *Object) error {
			notEmpty = notEmptyObject(o)
			return nil
		})
	}
	return notEmpty
}

// DerefItem unpacks an Item into an ItemCollection.
// If the Item is a slice type, like [IRIs], or [ItemCollection], it returns an [ItemCollection] corresponding to that
func DerefItem(it Item) ItemCollection {
	if IsNil(it) {
		return nil
	}

	var items ItemCollection
	switch {
	case IsIRIs(it):
		if col, err := ToIRIs(it); err == nil {
			items = col.Collection()
		}
	case IsItemCollection(it):
		if col, err := ToItemCollection(it); err == nil {
			items = *col
		}
	default:
		items = ItemCollection{it}
	}
	return items
}
