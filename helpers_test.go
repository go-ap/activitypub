package activitypub

import (
	"testing"
)

func TestOnObject(t *testing.T) {
	ob := ObjectNew(ArticleType)

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
	ob := ObjectNew(ArticleType)
	act := ActivityNew("test", CreateType, ob)

	err := OnActivity(act, func(a *Activity) error {
		return nil
	})

	if err != nil {
		t.Errorf("Unexpected error returned %s", err)
	}

	err = OnActivity(act, func(a *Activity) error {
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
	act := IntransitiveActivityNew("test", ArriveType)

	err := OnIntransitiveActivity(act, func(a *IntransitiveActivity) error {
		return nil
	})

	if err != nil {
		t.Errorf("Unexpected error returned %s", err)
	}

	err = OnIntransitiveActivity(act, func(a *IntransitiveActivity) error {
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
	act := QuestionNew("test")

	err := OnQuestion(act, func(a *Question) error {
		return nil
	})

	if err != nil {
		t.Errorf("Unexpected error returned %s", err)
	}

	err = OnQuestion(act, func(a *Question) error {
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
	pers := PersonNew("testPerson")
	err := OnActor(pers, func(a *Person) error {
		return nil
	})

	if err != nil {
		t.Errorf("Unexpected error returned %s", err)
	}

	err = OnActor(pers, func(p *Person) error {
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
