package activitypub

import (
	"errors"
	"fmt"
	"github.com/go-ap/activitystreams"
)

type withObjectFn func (*activitystreams.Object) error
type withActivityFn func (*activitystreams.Activity) error
type withPersonFn func (*Person) error

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
