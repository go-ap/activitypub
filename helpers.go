package activitypub

import (
	"errors"
	"fmt"
	"github.com/go-ap/activitystreams"
)

type withObjectFn func (*activitystreams.Object) error
type withActivityFn func (*activitystreams.Activity) error
type withIntransitiveActivityFn func (*activitystreams.IntransitiveActivity) error
type withQuestionFn func (*activitystreams.Question) error
type withPersonFn func (*Person) error
type withCollectionFn func (*activitystreams.Collection) error
type withCollectionPageFn func (*activitystreams.CollectionPage) error
type withOrderedCollectionFn func (*activitystreams.OrderedCollection) error
type withOrderedCollectionPageFn func (*activitystreams.OrderedCollectionPage) error

// OnObject
func OnObject(it activitystreams.Item, fn withObjectFn) error {
	ob, err  := activitystreams.ToObject(it)
	if err != nil {
		return err
	}
	return fn(ob)
}

// OnActivity
func OnActivity(it activitystreams.Item, fn withActivityFn) error {
	if !activitystreams.ActivityTypes.Contains(it.GetType()) {
		return errors.New(fmt.Sprintf("%T[%s] can't be converted to Activity", it, it.GetType()))
	}
	act, err  := activitystreams.ToActivity(it)
	if err != nil {
		return err
	}
	return fn(act)
}

// OnIntransitiveActivity
func OnIntransitiveActivity(it activitystreams.Item, fn withIntransitiveActivityFn) error {
	if !activitystreams.IntransitiveActivityTypes.Contains(it.GetType()) {
		return errors.New(fmt.Sprintf("%T[%s] can't be converted to Activity", it, it.GetType()))
	}
	if it.GetType() == activitystreams.QuestionType {
		errors.New(fmt.Sprintf("For %T[%s] you need to use OnQuestion function", it, it.GetType()))
	}
	act, err  := activitystreams.ToIntransitiveActivity(it)
	if err != nil {
		return err
	}
	return fn(act)
}

// OnQuestion
func OnQuestion(it activitystreams.Item, fn withQuestionFn) error {
	if it.GetType() != activitystreams.QuestionType {
		errors.New(fmt.Sprintf("For %T[%s] can't be converted to Question", it, it.GetType()))
	}
	act, err  := activitystreams.ToQuestion(it)
	if err != nil {
		return err
	}
	return fn(act)
}

// OnPerson
func OnPerson(it activitystreams.Item, fn withPersonFn) error {
	if !activitystreams.ActorTypes.Contains(it.GetType()) {
		return errors.New(fmt.Sprintf("%T[%s] can't be converted to Person", it, it.GetType()))
	}
	pers, err  := ToPerson(it)
	if err != nil {
		return err
	}
	return fn(pers)
}

// OnCollection
func OnCollection(it activitystreams.Item, fn withCollectionFn) error {
	switch it.GetType() {
	case activitystreams.CollectionType:
		col, err := activitystreams.ToCollection(it)
		if err != nil {
			return err
		}
		return fn(col)
	case activitystreams.CollectionPageType:
		return OnCollectionPage(it, func(p *activitystreams.CollectionPage) error {
			col, err := activitystreams.ToCollection(&p.ParentCollection)
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
func OnCollectionPage(it activitystreams.Item, fn withCollectionPageFn) error {
	if it.GetType() != activitystreams.CollectionPageType {
		return errors.New(fmt.Sprintf("%T[%s] can't be converted to Collection Page", it, it.GetType()))
	}
	col, err  := activitystreams.ToCollectionPage(it)
	if err != nil {
		return err
	}
	return fn(col)
}

// OnOrderedCollection
func OnOrderedCollection(it activitystreams.Item, fn withOrderedCollectionFn) error {
	switch it.GetType() {
	case activitystreams.OrderedCollectionType:
		col, err := activitystreams.ToOrderedCollection(it)
		if err != nil {
			return err
		}
		return fn(col)
	case activitystreams.OrderedCollectionPageType:
		return OnOrderedCollectionPage(it, func(p *activitystreams.OrderedCollectionPage) error {
			col, err := activitystreams.ToOrderedCollection(&p.OrderedCollection)
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
func OnOrderedCollectionPage(it activitystreams.Item, fn withOrderedCollectionPageFn) error {
	if it.GetType() != activitystreams.OrderedCollectionPageType {
		return errors.New(fmt.Sprintf("%T[%s] can't be converted to OrderedCollection Page", it, it.GetType()))
	}
	col, err  := activitystreams.ToOrderedCollectionPage(it)
	if err != nil {
		return err
	}
	return fn(col)
}
