package activitystreams

import "testing"

func TestIntransitiveActivityNew(t *testing.T) {
	var testValue = ObjectID("test")
	var testType ActivityVocabularyType = "Arrive"

	a := IntransitiveActivityNew(testValue, testType)

	if a.ID != testValue {
		t.Errorf("IntransitiveActivity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != testType {
		t.Errorf("IntransitiveActivity Type '%v' different than expected '%v'", a.Type, testType)
	}

	g := IntransitiveActivityNew(testValue, "")

	if g.ID != testValue {
		t.Errorf("IntransitiveActivity Id '%v' different than expected '%v'", g.ID, testValue)
	}
	if g.Type != IntransitiveActivityType {
		t.Errorf("IntransitiveActivity Type '%v' different than expected '%v'", g.Type, IntransitiveActivityType)
	}
}

func TestIntransitiveActivityRecipients(t *testing.T) {
	bob := PersonNew("bob")
	alice := PersonNew("alice")
	foo := OrganizationNew("foo")
	bar := GroupNew("bar")

	a := IntransitiveActivityNew("test", "t")

	a.To.Append(bob)
	a.To.Append(alice)
	a.To.Append(foo)
	a.To.Append(bar)
	if len(a.To) != 4 {
		t.Errorf("%T.To should have exactly 4(four) elements, not %d", a, len(a.To))
	}

	a.To.Append(bar)
	a.To.Append(alice)
	a.To.Append(foo)
	a.To.Append(bob)
	if len(a.To) != 8 {
		t.Errorf("%T.To should have exactly 8(eight) elements, not %d", a, len(a.To))
	}

	a.Recipients()
	if len(a.To) != 4 {
		t.Errorf("%T.To should have exactly 4(four) elements, not %d", a, len(a.To))
	}

	b := ActivityNew("t", "test", nil)

	b.To.Append(bar)
	b.To.Append(alice)
	b.To.Append(foo)
	b.To.Append(bob)
	b.Bto.Append(bar)
	b.Bto.Append(alice)
	b.Bto.Append(foo)
	b.Bto.Append(bob)
	b.CC.Append(bar)
	b.CC.Append(alice)
	b.CC.Append(foo)
	b.CC.Append(bob)
	b.BCC.Append(bar)
	b.BCC.Append(alice)
	b.BCC.Append(foo)
	b.BCC.Append(bob)

	b.Recipients()
	if len(b.To) != 4 {
		t.Errorf("%T.To should have exactly 4(four) elements, not %d", b, len(b.To))
	}
	if len(b.Bto) != 0 {
		t.Errorf("%T.Bto should have exactly 0(zero) elements, not %d", b, len(b.Bto))
	}
	if len(b.CC) != 0 {
		t.Errorf("%T.CC should have exactly 0(zero) elements, not %d", b, len(b.CC))
	}
	if len(b.BCC) != 0 {
		t.Errorf("%T.BCC should have exactly 0(zero) elements, not %d", b, len(b.BCC))
	}
	var err error
	recIds := make([]ObjectID, 0)
	err = checkDedup(b.To, &recIds)
	if err != nil {
		t.Error(err)
	}
	err = checkDedup(b.Bto, &recIds)
	if err != nil {
		t.Error(err)
	}
	err = checkDedup(b.CC, &recIds)
	if err != nil {
		t.Error(err)
	}
	err = checkDedup(b.BCC, &recIds)
	if err != nil {
		t.Error(err)
	}
}
func TestIntransitiveActivity_GetLink(t *testing.T) {
	i := IntransitiveActivityNew("test", QuestionType)

	if *i.GetID() != "test" {
		t.Errorf("%T should return an empty %T object. Received %#v", i, i, i)
	}
}
func TestIntransitiveActivity_GetObject(t *testing.T) {
	i := IntransitiveActivityNew("test", QuestionType)

	if *i.GetID() != "test" || i.GetType() != QuestionType {
		t.Errorf("%T should not return an empty %T object. Received %#v", i, i, i)
	}
}
func TestIntransitiveActivity_IsLink(t *testing.T) {
	i := IntransitiveActivityNew("test", QuestionType)

	if i.IsLink() {
		t.Errorf("%T should not respond true to IsLink", i)
	}
}
func TestIntransitiveActivity_IsObject(t *testing.T) {
	i := IntransitiveActivityNew("test", ActivityType)

	if !i.IsObject() {
		t.Errorf("%T should respond true to IsObject", i)
	}
}
func TestIntransitiveActivity_Recipients(t *testing.T) {
	to := PersonNew("bob")
	o := ObjectNew(ArticleType)
	cc := PersonNew("alice")

	o.ID = "something"

	c := IntransitiveActivityNew("act", IntransitiveActivityType)
	c.To.Append(to)
	c.CC.Append(cc)
	c.BCC.Append(cc)

	c.Recipients()

	var err error
	recIds := make([]ObjectID, 0)
	err = checkDedup(c.To, &recIds)
	if err != nil {
		t.Error(err)
	}
	err = checkDedup(c.Bto, &recIds)
	if err != nil {
		t.Error(err)
	}
	err = checkDedup(c.CC, &recIds)
	if err != nil {
		t.Error(err)
	}
	err = checkDedup(c.BCC, &recIds)
	if err != nil {
		t.Error(err)
	}
}

func TestIntransitiveActivity_GetID(t *testing.T) {
	a := IntransitiveActivityNew("test", IntransitiveActivityType)

	if *a.GetID() != "test" {
		t.Errorf("%T should return an empty %T object. Received %#v", a, a.GetID(), *a.GetID())
	}
}

func TestIntransitiveActivity_GetType(t *testing.T) {
	{
		a := IntransitiveActivityNew("test", IntransitiveActivityType)
		if a.GetType() != IntransitiveActivityType {
			t.Errorf("GetType should return %q for %T, received %q", IntransitiveActivityType, a, a.GetType())
		}
	}
	{
		a := IntransitiveActivityNew("test", ArriveType)
		if a.GetType() != ArriveType {
			t.Errorf("GetType should return %q for %T, received %q", ArriveType, a, a.GetType())
		}
	}
	{
		a := IntransitiveActivityNew("test", QuestionType)
		if a.GetType() != QuestionType {
			t.Errorf("GetType should return %q for %T, received %q", QuestionType, a, a.GetType())
		}
	}
}

func TestToIntransitiveActivity(t *testing.T) {
	var it Item
	act := IntransitiveActivityNew("test", TravelType)
	it = act

	a, err := ToIntransitiveActivity(it)
	if err != nil {
		t.Error(err)
	}
	if a != act {
		t.Errorf("Invalid activity returned by ToActivity #%v", a)
	}

	ob := ObjectNew(ArticleType)
	it = ob

	o, err := ToIntransitiveActivity(it)
	if err == nil {
		t.Errorf("Error returned when calling ToActivity with object should not be nil")
	}
	if o != nil {
		t.Errorf("Invalid return by ToActivity #%v, should have been nil", o)
	}
}

func TestFlattenIntransitiveActivityProperties(t *testing.T) {
	t.Skipf("TODO")
}

func TestIntransitiveActivity_Clean(t *testing.T) {
	t.Skipf("TODO")
}

func TestIntransitiveActivity_IsCollection(t *testing.T) {
	t.Skipf("TODO")
}

func TestIntransitiveActivity_UnmarshalJSON(t *testing.T) {
	t.Skipf("TODO")
}

func TestArriveNew(t *testing.T) {
	var testValue = ObjectID("test")

	a := ArriveNew(testValue)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != ArriveType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, ArriveType)
	}
}

func TestTravelNew(t *testing.T) {
	var testValue = ObjectID("test")

	a := TravelNew(testValue)

	if a.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.ID, testValue)
	}
	if a.Type != TravelType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, TravelType)
	}
}
