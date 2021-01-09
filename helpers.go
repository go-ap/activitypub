package activitypub

import (
	"fmt"
)

type withLinkFn func (*Link) error
type withObjectFn func (*Object) error
type withActivityFn func (*Activity) error
type withIntransitiveActivityFn func (*IntransitiveActivity) error
type withQuestionFn func (*Question) error
type withActorFn func (*Actor) error
type withCollectionInterfaceFn func (collection CollectionInterface) error
type withCollectionFn func (collection *Collection) error
type withCollectionPageFn func (*CollectionPage) error
type withOrderedCollectionFn func (*OrderedCollection) error
type withOrderedCollectionPageFn func (*OrderedCollectionPage) error
type withItemCollectionFn func (collection *ItemCollection) error

// OnLink
func OnLink(it Item, fn withLinkFn) error {
	ob, err := ToLink(it)
	if err != nil {
		return err
	}
	return fn(ob)
}

// OnObject
func OnObject(it Item, fn withObjectFn) error {
	ob, err := ToObject(it)
	if err != nil {
		return err
	}
	return fn(ob)
}

// OnActivity
func OnActivity(it Item, fn withActivityFn) error {
	act, err  := ToActivity(it)
	if err != nil {
		return err
	}
	return fn(act)
}

// OnIntransitiveActivity
func OnIntransitiveActivity(it Item, fn withIntransitiveActivityFn) error {
	act, err  := ToIntransitiveActivity(it)
	if err != nil {
		return err
	}
	return fn(act)
}

// OnQuestion
func OnQuestion(it Item, fn withQuestionFn) error {
	act, err  := ToQuestion(it)
	if err != nil {
		return err
	}
	return fn(act)
}

// OnActor
func OnActor(it Item, fn withActorFn) error {
	act, err  := ToActor(it)
	if err != nil {
		return err
	}
	return fn(act)
}
// OnCollection
func OnCollection (it Item, fn withCollectionFn) error {
	col, err := ToCollection(it)
	if err != nil {
		return err
	}
	return fn(col)
}

// OnCollectionIntf
func OnCollectionIntf(it Item, fn withCollectionInterfaceFn) error {
	switch it.GetType() {
	case CollectionOfItems:
		col, err := ToItemCollection(it)
		if err != nil {
			return err
		}
		return fn(col)
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
		return fmt.Errorf("%T[%s] can't be converted to Collection", it, it.GetType())
	}
}

// OnCollectionPage
func OnCollectionPage(it Item, fn withCollectionPageFn) error {
	col, err  := ToCollectionPage(it)
	if err != nil {
		return err
	}
	return fn(col)
}

// OnOrderedCollection
func OnOrderedCollection(it Item, fn withOrderedCollectionFn) error {
	col, err := ToOrderedCollection(it)
	if err != nil {
		return err
	}
	return fn(col)
}

// OnOrderedCollectionPage executes a function on an ordered collection page type item
func OnOrderedCollectionPage(it Item, fn withOrderedCollectionPageFn) error {
	col, err  := ToOrderedCollectionPage(it)
	if err != nil {
		return err
	}
	return fn(col)
}

// OnItemCollection executes a function on a collection type item
func OnItemCollection(it Item, fn withItemCollectionFn) error {
	col, err  := ToItemCollection(it)
	if err != nil {
		return err
	}
	return fn(col)
}

// ItemOrderTimestamp is used for ordering a ItemCollection slice using the slice.Sort function
// It orders i1 and i2 based on their Published and Updated timestamps.
func ItemOrderTimestamp(i1, i2 Item) bool {
	o1, e1 := ToObject(i1)
	o2, e2 := ToObject(i2)
	if e1 != nil || e2 != nil {
		return false
	}
	t1 := o1.Published
	if !o1.Updated.IsZero() {
		t1 = o1.Updated
	}
	t2 := o2.Published
	if !o2.Updated.IsZero() {
		t2 = o2.Updated
	}
	return t1.Sub(t2) > 0
}

func notEmptyLink (l *Link) bool {
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

func notEmptyInstransitiveActivity (i *IntransitiveActivity) bool {
	return i.Actor != nil ||
		i.Target != nil ||
		i.Result != nil ||
		i.Origin != nil ||
		i.Instrument != nil
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
		len(a.PublicKey.ID) > 0 ||
		(a.PublicKey.Owner != nil &&
		len(a.PublicKey.PublicKeyPem) > 0)
}

func NotEmpty(i Item) bool {
	if i == nil {
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