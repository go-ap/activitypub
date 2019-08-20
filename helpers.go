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

// OnObject
func OnObject(it activitystreams.Item, fn withObjectFn) error {
	if !activitystreams.ObjectTypes.Contains(it.GetType()) {
		return errors.New(fmt.Sprintf("%T[%s] can't be converted to Object", it, it.GetType()))
	}
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
