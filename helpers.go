package activitypub

import (
	"fmt"
	"time"
)

// WithLinkFn represents a function type that can be used as a parameter for OnLink helper function
type WithLinkFn func(*Link) error

// WithObjectFn represents a function type that can be used as a parameter for OnObject helper function
type WithObjectFn func(*Object) error

// WithActivityFn represents a function type that can be used as a parameter for OnActivity helper function
type WithActivityFn func(*Activity) error

// WithIntransitiveActivityFn represents a function type that can be used as a parameter for OnIntransitiveActivity helper function
type WithIntransitiveActivityFn func(*IntransitiveActivity) error

// WithQuestionFn represents a function type that can be used as a parameter for OnQuestion helper function
type WithQuestionFn func(*Question) error

// WithActorFn represents a function type that can be used as a parameter for OnActor helper function
type WithActorFn func(*Actor) error

// WithCollectionInterfaceFn represents a function type that can be used as a parameter for OnCollectionIntf helper function
type WithCollectionInterfaceFn func(CollectionInterface) error

// WithCollectionFn represents a function type that can be used as a parameter for OnCollection helper function
type WithCollectionFn func(*Collection) error

// WithCollectionPageFn represents a function type that can be used as a parameter for OnCollectionPage helper function
type WithCollectionPageFn func(*CollectionPage) error

// WithOrderedCollectionFn represents a function type that can be used as a parameter for OnOrderedCollection helper function
type WithOrderedCollectionFn func(*OrderedCollection) error

// WithOrderedCollectionPageFn represents a function type that can be used as a parameter for OnOrderedCollectionPage helper function
type WithOrderedCollectionPageFn func(*OrderedCollectionPage) error

// WithItemCollectionFn represents a function type that can be used as a parameter for OnItemCollection helper function
type WithItemCollectionFn func(*ItemCollection) error

// WithIRIsFn represents a function type that can be used as a parameter for OnIRIs helper function
type WithIRIsFn func(*IRIs) error

// OnLink calls function fn on it Item if it can be asserted to type *Link
//
// This function should be safe to use for all types with a structure compatible
// with the Link type
func OnLink(it LinkOrIRI, fn WithLinkFn) error {
	if it == nil {
		return nil
	}
	ob, err := ToLink(it)
	if err != nil {
		return err
	}
	return fn(ob)
}

func To[T Item](it Item) (*T, error) {
	if ob, ok := it.(T); ok {
		return &ob, nil
	}
	return nil, fmt.Errorf("invalid cast for object %T", it)
}

// On handles in a generic way the call to fn(*T) if the "it" Item can be asserted to one of the Objects type.
// It also covers the case where "it" is a collection of items that match the assertion.
func On[T Item](it Item, fn func(*T) error) error {
	if !IsItemCollection(it) {
		ob, err := To[T](it)
		if err != nil {
			return err
		}
		return fn(ob)
	}
	return OnItemCollection(it, func(col *ItemCollection) error {
		for _, it := range *col {
			if err := On(it, fn); err != nil {
				return err
			}
		}
		return nil
	})
}

// OnObject calls function fn on it Item if it can be asserted to type *Object
//
// This function should be safe to be called for all types with a structure compatible
// to the Object type.
func OnObject(it Item, fn WithObjectFn) error {
	if it == nil {
		return nil
	}
	if IsItemCollection(it) {
		return OnItemCollection(it, func(col *ItemCollection) error {
			for _, it := range *col {
				if IsLink(it) {
					continue
				}
				if err := OnObject(it, fn); err != nil {
					return err
				}
			}
			return nil
		})
	}
	ob, err := ToObject(it)
	if err != nil {
		return err
	}
	return fn(ob)
}

// OnActivity calls function fn on it Item if it can be asserted to type *Activity
//
// This function should be called if trying to access the Activity specific properties
// like "object", for the other properties OnObject, or OnIntransitiveActivity
// should be used instead.
func OnActivity(it Item, fn WithActivityFn) error {
	if it == nil {
		return nil
	}
	if IsItemCollection(it) {
		return OnItemCollection(it, func(col *ItemCollection) error {
			for _, it := range *col {
				if IsLink(it) {
					continue
				}
				if err := OnActivity(it, fn); err != nil {
					return err
				}
			}
			return nil
		})
	}
	act, err := ToActivity(it)
	if err != nil {
		return err
	}
	return fn(act)
}

// OnIntransitiveActivity calls function fn on it Item if it can be asserted
// to type *IntransitiveActivity
//
// This function should be called if trying to access the IntransitiveActivity
// specific properties like "actor", for the other properties OnObject
// should be used instead.
func OnIntransitiveActivity(it Item, fn WithIntransitiveActivityFn) error {
	if it == nil {
		return nil
	}
	if IsItemCollection(it) {
		return OnItemCollection(it, func(col *ItemCollection) error {
			for _, it := range *col {
				if err := OnIntransitiveActivity(it, fn); err != nil {
					return err
				}
			}
			return nil
		})
	}
	act, err := ToIntransitiveActivity(it)
	if err != nil {
		return err
	}
	return fn(act)
}

// OnQuestion calls function fn on it Item if it can be asserted to type Question
//
// This function should be called if trying to access the Questions specific
// properties like "anyOf", "oneOf", "closed", etc. For the other properties
// OnObject or OnIntransitiveActivity should be used instead.
func OnQuestion(it Item, fn WithQuestionFn) error {
	if it == nil {
		return nil
	}
	if IsItemCollection(it) {
		return OnItemCollection(it, func(col *ItemCollection) error {
			for _, it := range *col {
				if err := OnQuestion(it, fn); err != nil {
					return err
				}
			}
			return nil
		})
	}
	act, err := ToQuestion(it)
	if err != nil {
		return err
	}
	return fn(act)
}

// OnActor calls function fn on it Item if it can be asserted to type *Actor
//
// This function should be called if trying to access the Actor specific
// properties like "preferredName", "publicKey", etc. For the other properties
// OnObject should be used instead.
func OnActor(it Item, fn WithActorFn) error {
	if it == nil {
		return nil
	}
	if IsItemCollection(it) {
		return OnItemCollection(it, func(col *ItemCollection) error {
			for _, it := range *col {
				if IsLink(it) {
					continue
				}
				if err := OnActor(it, fn); err != nil {
					return err
				}
			}
			return nil
		})
	}
	act, err := ToActor(it)
	if err != nil {
		return err
	}
	return fn(act)
}

// OnItemCollection calls function fn on it Item if it can be asserted to type ItemCollection
//
// It should be used when Item represents an Item collection and it's usually used as a way
// to wrap functionality for other functions that will be called on each item in the collection.
func OnItemCollection(it Item, fn WithItemCollectionFn) error {
	if it == nil {
		return nil
	}
	col, err := ToItemCollection(it)
	if err != nil {
		return err
	}
	return fn(col)
}

// OnIRIs calls function fn on it Item if it can be asserted to type IRIs
//
// It should be used when Item represents an IRI slice.
func OnIRIs(it Item, fn WithIRIsFn) error {
	if it == nil {
		return nil
	}
	col, err := ToIRIs(it)
	if err != nil {
		return err
	}
	return fn(col)
}

// OnCollectionIntf calls function fn on it Item if it can be asserted to a type
// that implements the CollectionInterface
//
// This function should be called if Item represents a collection of ActivityPub
// objects. It basically wraps functionality for the different collection types
// supported by the package.
func OnCollectionIntf(it Item, fn WithCollectionInterfaceFn) error {
	if it == nil {
		return nil
	}
	switch it.GetType() {
	case CollectionOfItems:
		col, err := ToItemCollection(it)
		if err != nil {
			return err
		}
		return fn(col)
	case CollectionOfIRIs:
		col, err := ToIRIs(it)
		if err != nil {
			return err
		}
		itCol := col.Collection()
		return fn(&itCol)
	case CollectionType:
		col, err := ToCollection(it)
		if err != nil {
			return err
		}
		return fn(col)
	case CollectionPageType:
		return OnCollectionPage(it, func(p *CollectionPage) error {
			col, err := ToCollectionPage(p)
			if err != nil {
				return err
			}
			return fn(col)
		})
	case OrderedCollectionType:
		col, err := ToOrderedCollection(it)
		if err != nil {
			return err
		}
		return fn(col)
	case OrderedCollectionPageType:
		return OnOrderedCollectionPage(it, func(p *OrderedCollectionPage) error {
			col, err := ToOrderedCollectionPage(p)
			if err != nil {
				return err
			}
			return fn(col)
		})
	default:
		return fmt.Errorf("%T[%s] can't be converted to a Collection type", it, it.GetType())
	}
}

// OnCollection calls function fn on it Item if it can be asserted to type *Collection
//
// This function should be called if trying to access the Collection specific
// properties like "totalItems", "items", etc. For the other properties
// OnObject should be used instead.
func OnCollection(it Item, fn WithCollectionFn) error {
	if it == nil {
		return nil
	}
	col, err := ToCollection(it)
	if err != nil {
		return err
	}
	return fn(col)
}

// OnCollectionPage calls function fn on it Item if it can be asserted to
// type *CollectionPage
//
// This function should be called if trying to access the CollectionPage specific
// properties like "partOf", "next", "perv". For the other properties
// OnObject or OnCollection should be used instead.
func OnCollectionPage(it Item, fn WithCollectionPageFn) error {
	if it == nil {
		return nil
	}
	col, err := ToCollectionPage(it)
	if err != nil {
		return err
	}
	return fn(col)
}

// OnOrderedCollection calls function fn on it Item if it can be asserted
// to type *OrderedCollection
//
// This function should be called if trying to access the Collection specific
// properties like "totalItems", "orderedItems", etc. For the other properties
// OnObject should be used instead.
func OnOrderedCollection(it Item, fn WithOrderedCollectionFn) error {
	if it == nil {
		return nil
	}
	col, err := ToOrderedCollection(it)
	if err != nil {
		return err
	}
	return fn(col)
}

// OnOrderedCollectionPage calls function fn on it Item if it can be asserted
// to type *OrderedCollectionPage
//
// This function should be called if trying to access the OrderedCollectionPage specific
// properties like "partOf", "next", "perv". For the other properties
// OnObject or OnOrderedCollection should be used instead.
func OnOrderedCollectionPage(it Item, fn WithOrderedCollectionPageFn) error {
	if it == nil {
		return nil
	}
	col, err := ToOrderedCollectionPage(it)
	if err != nil {
		return err
	}
	return fn(col)
}

// ItemOrderTimestamp is used for ordering a ItemCollection slice using the slice.Sort function
// It orders i1 and i2 based on their Published and Updated timestamps, whichever is later.
func ItemOrderTimestamp(i1, i2 LinkOrIRI) bool {
	if IsNil(i1) {
		return !IsNil(i2)
	} else if IsNil(i2) {
		return false
	}

	var t1 time.Time
	var t2 time.Time
	if IsObject(i1) {
		o1, e1 := ToObject(i1)
		if e1 != nil {
			return false
		}
		t1 = o1.Published
		if o1.Updated.After(t1) {
			t1 = o1.Updated
		}
	}
	if IsObject(i2) {
		o2, e2 := ToObject(i2)
		if e2 != nil {
			return false
		}
		t2 = o2.Published
		if o2.Updated.After(t2) {
			t2 = o2.Updated
		}
	}
	return t1.After(t2)
}

func notEmptyLink(l *Link) bool {
	return len(l.ID) > 0 ||
		LinkTypes.Contains(l.Type) ||
		len(l.MediaType) > 0 ||
		l.Preview != nil ||
		l.Name != nil ||
		len(l.Href) > 0 ||
		len(l.Rel) > 0 ||
		len(l.HrefLang) > 0 ||
		l.Height > 0 ||
		l.Width > 0
}

func notEmptyObject(o *Object) bool {
	if o == nil {
		return false
	}
	return len(o.ID) > 0 ||
		len(o.Type) > 0 ||
		ActivityTypes.Contains(o.Type) ||
		o.Content != nil ||
		o.Attachment != nil ||
		o.AttributedTo != nil ||
		o.Audience != nil ||
		o.BCC != nil ||
		o.Bto != nil ||
		o.CC != nil ||
		o.Context != nil ||
		o.Duration > 0 ||
		!o.EndTime.IsZero() ||
		o.Generator != nil ||
		o.Icon != nil ||
		o.Image != nil ||
		o.InReplyTo != nil ||
		o.Likes != nil ||
		o.Location != nil ||
		len(o.MediaType) > 0 ||
		o.Name != nil ||
		o.Preview != nil ||
		!o.Published.IsZero() ||
		o.Replies != nil ||
		o.Shares != nil ||
		o.Source.MediaType != "" ||
		o.Source.Content != nil ||
		!o.StartTime.IsZero() ||
		o.Summary != nil ||
		o.Tag != nil ||
		o.To != nil ||
		!o.Updated.IsZero() ||
		o.URL != nil
}

func notEmptyInstransitiveActivity(i *IntransitiveActivity) bool {
	notEmpty := i.Actor != nil ||
		i.Target != nil ||
		i.Result != nil ||
		i.Origin != nil ||
		i.Instrument != nil
	if notEmpty {
		return true
	}
	OnObject(i, func(ob *Object) error {
		notEmpty = notEmptyObject(ob)
		return nil
	})
	return notEmpty
}

func notEmptyActivity(a *Activity) bool {
	var notEmpty bool
	OnIntransitiveActivity(a, func(i *IntransitiveActivity) error {
		notEmpty = notEmptyInstransitiveActivity(i)
		return nil
	})
	return notEmpty || a.Object != nil
}

func notEmptyActor(a *Actor) bool {
	var notEmpty bool
	OnObject(a, func(o *Object) error {
		notEmpty = notEmptyObject(o)
		return nil
	})
	return notEmpty ||
		a.Inbox != nil ||
		a.Outbox != nil ||
		a.Following != nil ||
		a.Followers != nil ||
		a.Liked != nil ||
		a.PreferredUsername != nil ||
		a.Endpoints != nil ||
		a.Streams != nil ||
		len(a.PublicKey.ID)+len(a.PublicKey.Owner)+len(a.PublicKey.PublicKeyPem) > 0
}

// NotEmpty tells us if a Item interface value has a non nil value for various types
// that implement
func NotEmpty(i Item) bool {
	if IsNil(i) {
		return false
	}
	var notEmpty bool
	if IsIRI(i) {
		notEmpty = len(i.GetLink()) > 0
	}
	if i.IsCollection() {
		OnCollectionIntf(i, func(c CollectionInterface) error {
			notEmpty = c != nil || len(c.Collection()) > 0
			return nil
		})
	}
	if ActivityTypes.Contains(i.GetType()) {
		OnActivity(i, func(a *Activity) error {
			notEmpty = notEmptyActivity(a)
			return nil
		})
	} else if ActorTypes.Contains(i.GetType()) {
		OnActor(i, func(a *Actor) error {
			notEmpty = notEmptyActor(a)
			return nil
		})
	} else if i.IsLink() {
		OnLink(i, func(l *Link) error {
			notEmpty = notEmptyLink(l)
			return nil
		})
	} else {
		OnObject(i, func(o *Object) error {
			notEmpty = notEmptyObject(o)
			return nil
		})
	}
	return notEmpty
}

// DerefItem dereferences
func DerefItem(it Item) ItemCollection {
	if IsNil(it) {
		return nil
	}

	var items ItemCollection
	if IsIRIs(it) {
		_ = OnIRIs(it, func(col *IRIs) error {
			items = col.Collection()
			return nil
		})
	} else if IsItemCollection(it) {
		_ = OnItemCollection(it, func(col *ItemCollection) error {
			items = col.Collection()
			return nil
		})
	} else {
		items = ItemCollection{it}
	}
	return items
}
