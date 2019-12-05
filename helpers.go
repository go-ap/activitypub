package activitypub

import (
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
	ob, err  := ToObject(it)
	if err != nil {
		return err
	}
	return fn(ob)
}

// OnActivity
func OnActivity(it Item, fn withActivityFn) error {
	if !(ActivityTypes.Contains(it.GetType()) || IntransitiveActivityTypes.Contains(it.GetType())) {
		return fmt.Errorf("%T[%s] can't be converted to Activity", it, it.GetType())
	}
	act, err  := ToActivity(it)
	if err != nil {
		return err
	}
	return fn(act)
}

// OnIntransitiveActivity
func OnIntransitiveActivity(it Item, fn withIntransitiveActivityFn) error {
	if it.GetType() == QuestionType {
		fmt.Errorf("for %T[%s] you need to use OnQuestion function", it, it.GetType())
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
		fmt.Errorf("For %T[%s] can't be converted to Question", it, it.GetType())
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
		return fmt.Errorf("%T[%s] can't be converted to Person", it, it.GetType())
	}
	act, err  := ToActor(it)
	if err != nil {
		return err
	}
	return fn(act)
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
		return fmt.Errorf("%T[%s] can't be converted to Collection", it, it.GetType())
	}
}

// OnCollectionPage
func OnCollectionPage(it Item, fn withCollectionPageFn) error {
	if it.GetType() != CollectionPageType {
		return fmt.Errorf("%T[%s] can't be converted to Collection Page", it, it.GetType())
	}
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

// OnOrderedCollectionPage
func OnOrderedCollectionPage(it Item, fn withOrderedCollectionPageFn) error {
	if it.GetType() != OrderedCollectionPageType {
		return fmt.Errorf("%T[%s] can't be converted to OrderedCollection Page", it, it.GetType())
	}
	col, err  := ToOrderedCollectionPage(it)
	if err != nil {
		return err
	}
	return fn(col)
}

// OnItemCollection
func OnItemCollection(it Item, fn withItemCollectionFn) error {
	col, err  := ToItemCollection(it)
	if err != nil {
		return err
	}
	return fn(col)
}
