package activitypub

import "testing"

func TestQuestionNew(t *testing.T) {
	var testValue = ObjectID("test")

	a := QuestionNew(testValue)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != QuestionType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, QuestionType)
	}
}

func TestQuestion_GetID(t *testing.T) {
	a := QuestionNew("test")

	if a.GetID() != "test" {
		t.Errorf("%T should return an empty %T object. Received %#v", a, a.GetID(), a.GetID())
	}
}

func TestQuestion_IsObject(t *testing.T) {
	a := QuestionNew("test")

	if !a.IsObject() {
		t.Errorf("%T should respond true to IsObject", a)
	}
}

func TestQuestion_IsLink(t *testing.T) {
	a := QuestionNew("test")

	if a.IsLink() {
		t.Errorf("%T should respond false to IsLink", a)
	}
}

func TestQuestion_GetLink(t *testing.T) {
	a := QuestionNew("test")

	if a.GetLink() != "test" {
		t.Errorf("GetLink should return \"test\" for %T, received %q", a, a.GetLink())
	}
}

func TestQuestion_GetType(t *testing.T) {
	a := QuestionNew("test")

	if a.GetType() != QuestionType {
		t.Errorf("GetType should return %q for %T, received %q", QuestionType, a, a.GetType())
	}
}


func TestToQuestion(t *testing.T) {
	var it Item
	act := QuestionNew("test")
	it = act

	a, err := ToQuestion(it)
	if err != nil {
		t.Error(err)
	}
	if a != act {
		t.Errorf("Invalid activity returned by ToActivity #%v", a)
	}

	ob := ObjectNew(ArticleType)
	it = ob

	o, err := ToQuestion(it)
	if err == nil {
		t.Errorf("Error returned when calling ToActivity with object should not be nil")
	}
	if o != nil {
		t.Errorf("Invalid return by ToActivity #%v, should have been nil", o)
	}
}

func TestQuestion_IsCollection(t *testing.T) {
	t.Skipf("TODO")
}

func TestQuestion_UnmarshalJSON(t *testing.T) {
	t.Skipf("TODO")
}
