package activitypub

import (
	"github.com/go-ap/activitystreams"
	"testing"
)

func TestOnObject(t *testing.T) {
	ob := activitystreams.ObjectNew(activitystreams.ArticleType)

	err := OnObject(ob, func(o *Object) error {
		return nil
	})

	if err != nil {
		t.Errorf("Unexpected error returned %s", err)
	}

	err = OnObject(ob, func(o *Object) error {
		if o.Type != ob.Type {
			t.Errorf("In function type %s different than expected, %s", o.Type, ob.Type)
		}
		return nil
	})
	if err != nil {
		t.Errorf("Unexpected error returned %s", err)
	}
}

func TestOnActivity(t *testing.T) {
	ob := activitystreams.ObjectNew(activitystreams.ArticleType)
	act := activitystreams.ActivityNew("test", activitystreams.CreateType, ob)

	err := OnActivity(act, func(a *activitystreams.Activity) error {
		return nil
	})

	if err != nil {
		t.Errorf("Unexpected error returned %s", err)
	}

	err = OnActivity(act, func(a *activitystreams.Activity) error {
		if a.Type != act.Type {
			t.Errorf("In function type %s different than expected, %s", a.Type, act.Type)
		}
		if a.ID != act.ID {
			t.Errorf("In function ID %s different than expected, %s", a.ID, act.ID)
		}
		if a.Object != act.Object {
			t.Errorf("In function object %s different than expected, %s", a.Object, act.Object)
		}
		return nil
	})
	if err != nil {
		t.Errorf("Unexpected error returned %s", err)
	}
}

func TestOnIntransitiveActivity(t *testing.T) {
	act := activitystreams.IntransitiveActivityNew("test", activitystreams.ArriveType)

	err := OnIntransitiveActivity(act, func(a *activitystreams.IntransitiveActivity) error {
		return nil
	})

	if err != nil {
		t.Errorf("Unexpected error returned %s", err)
	}

	err = OnIntransitiveActivity(act, func(a *activitystreams.IntransitiveActivity) error {
		if a.Type != act.Type {
			t.Errorf("In function type %s different than expected, %s", a.Type, act.Type)
		}
		if a.ID != act.ID {
			t.Errorf("In function ID %s different than expected, %s", a.ID, act.ID)
		}
		return nil
	})
	if err != nil {
		t.Errorf("Unexpected error returned %s", err)
	}
}

func TestOnQuestion(t *testing.T) {
	act := activitystreams.QuestionNew("test")

	err := OnQuestion(act, func(a *activitystreams.Question) error {
		return nil
	})

	if err != nil {
		t.Errorf("Unexpected error returned %s", err)
	}

	err = OnQuestion(act, func(a *activitystreams.Question) error {
		if a.Type != act.Type {
			t.Errorf("In function type %s different than expected, %s", a.Type, act.Type)
		}
		if a.ID != act.ID {
			t.Errorf("In function ID %s different than expected, %s", a.ID, act.ID)
		}
		return nil
	})
	if err != nil {
		t.Errorf("Unexpected error returned %s", err)
	}
}

func TestOnPerson(t *testing.T) {
	pers := activitystreams.PersonNew("testPerson")
	err := OnPerson(pers, func(a *Person) error {
		return nil
	})

	if err != nil {
		t.Errorf("Unexpected error returned %s", err)
	}

	err = OnPerson(pers, func(p *Person) error {
		if p.Type != pers.Type {
			t.Errorf("In function type %s different than expected, %s", p.Type, pers.Type)
		}
		if p.ID != pers.ID {
			t.Errorf("In function ID %s different than expected, %s", p.ID, pers.ID)
		}
		return nil
	})
	if err != nil {
		t.Errorf("Unexpected error returned %s", err)
	}
}

func TestOnCollection(t *testing.T) {
	t.Skipf("TODO")
}

func TestOnCollectionPage(t *testing.T) {
	t.Skipf("TODO")
}

func TestOnOrderedCollection(t *testing.T) {
	t.Skipf("TODO")
}

func TestOnOrderedCollectionPage(t *testing.T) {
	t.Skipf("TODO")
}
