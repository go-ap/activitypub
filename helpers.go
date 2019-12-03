package activitypub

import (
	"errors"
	"fmt"
)

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

// OnObject
func OnObject(it Item, fn withObjectFn) error {
	if ActivityTypes.Contains(it.GetType()) {
		return OnActivity(it, func(a *Activity) error {
			ob, err := ToObject(a)
			if err != nil {
				return err
			}
			return fn(ob)
		})
	} else if ActorTypes.Contains(it.GetType()) {
		return OnActor(it, func(p *Actor) error {
			ob, err := ToObject(p)
			if err != nil {
				return err
			}
			return fn(ob)
		})
	} else if it.IsCollection() {
		return OnCollection(it, func(col CollectionInterface) error {
			for _, it := range col.Collection() {
				err := OnObject(it, fn)
				if err != nil {
					return err
				}
			}
			return nil
		})
	} else {
		ob, err  := ToObject(it)
		if err != nil {
			return err
		}
		return fn(ob)
	}
}

// OnActivity
func OnActivity(it Item, fn withActivityFn) error {
	if !ActivityTypes.Contains(it.GetType()) {
		return errors.New(fmt.Sprintf("%T[%s] can't be converted to Activity", it, it.GetType()))
	}
	act, err  := ToActivity(it)
	if err != nil {
		return err
	}
	return fn(act)
}

// OnIntransitiveActivity
func OnIntransitiveActivity(it Item, fn withIntransitiveActivityFn) error {
	if !IntransitiveActivityTypes.Contains(it.GetType()) {
		return errors.New(fmt.Sprintf("%T[%s] can't be converted to Activity", it, it.GetType()))
	}
	if it.GetType() == QuestionType {
		errors.New(fmt.Sprintf("For %T[%s] you need to use OnQuestion function", it, it.GetType()))
	}
	act, err  := ToIntransitiveActivity(it)
	if err != nil {
		return err
	}
	return fn(act)
}

// OnQuestion
func OnQuestion(it Item, fn withQuestionFn) error {
	if it.GetType() != QuestionType {
		errors.New(fmt.Sprintf("For %T[%s] can't be converted to Question", it, it.GetType()))
	}
	act, err  := ToQuestion(it)
	if err != nil {
		return err
	}
	return fn(act)
}

// OnActor
func OnActor(it Item, fn withActorFn) error {
	if !ActorTypes.Contains(it.GetType()) {
		return errors.New(fmt.Sprintf("%T[%s] can't be converted to Person", it, it.GetType()))
	}
	pers, err  := ToActor(it)
	if err != nil {
		return err
	}
	return fn(pers)
}

// OnCollection
func OnCollection(it Item, fn withCollectionInterfaceFn) error {
	switch it.GetType() {
	case CollectionOfItems:
		col, err := ToItemCollection(it)
		if err != nil {
			return err
		}
		c := Collection{
			TotalItems: uint(len(*col)),
			Items:      *col,
		}
		return fn(&c)
	case CollectionType:
		col, err := ToCollection(it)
		if err != nil {
			return err
		}
		return fn(col)
	case CollectionPageType:
		return OnCollectionPage(it, func(p *CollectionPage) error {
			col, err := ToCollection(p)
			if err != nil {
				return err
			}
			return fn(col)
		})
	default:
		return errors.New(fmt.Sprintf("%T[%s] can't be converted to Collection", it, it.GetType()))
	}
}

// OnCollectionPage
func OnCollectionPage(it Item, fn withCollectionPageFn) error {
	if it.GetType() != CollectionPageType {
		return errors.New(fmt.Sprintf("%T[%s] can't be converted to Collection Page", it, it.GetType()))
	}
	col, err  := ToCollectionPage(it)
	if err != nil {
		return err
	}
	return fn(col)
}

// OnOrderedCollection
func OnOrderedCollection(it Item, fn withOrderedCollectionFn) error {
	switch it.GetType() {
	case OrderedCollectionType:
		col, err := ToOrderedCollection(it)
		if err != nil {
			return err
		}
		return fn(col)
	case OrderedCollectionPageType:
		return OnOrderedCollectionPage(it, func(p *OrderedCollectionPage) error {
			col, err := ToOrderedCollection(p)
			if err != nil {
				return err
			}
			return fn(col)
		})
	default:
		return errors.New(fmt.Sprintf("%T[%s] can't be converted to OrderedCollection", it, it.GetType()))
	}
}

// OnOrderedCollectionPage
func OnOrderedCollectionPage(it Item, fn withOrderedCollectionPageFn) error {
	if it.GetType() != OrderedCollectionPageType {
		return errors.New(fmt.Sprintf("%T[%s] can't be converted to OrderedCollection Page", it, it.GetType()))
	}
	col, err  := ToOrderedCollectionPage(it)
	if err != nil {
		return err
	}
	return fn(col)
}

// OnOrderedCollectionPage
func OnItemCOllection(it Item, fn withOrderedCollectionPageFn) error {
	if it.GetType() != OrderedCollectionPageType {
		return errors.New(fmt.Sprintf("%T[%s] can't be converted to OrderedCollection Page", it, it.GetType()))
	}
	col, err  := ToOrderedCollectionPage(it)
	if err != nil {
		return err
	}
	return fn(col)
}
